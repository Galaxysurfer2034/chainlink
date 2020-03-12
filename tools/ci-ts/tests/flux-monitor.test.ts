import { FluxAggregatorFactory } from '@chainlink/contracts/ethers/v0.6/FluxAggregatorFactory'
import { contract, helpers as h, matchers } from '@chainlink/test-helpers'
import { assert } from 'chai'
import { ethers } from 'ethers'
import 'isomorphic-unfetch'
import { JobSpec } from '../../../operator_ui/@types/operator_ui'
import ChainlinkClient from '../test-helpers/chainlink-cli'
import fluxMonitorJob from '../fixtures/flux-monitor-job'
import {
  assertAsync,
  createProvider,
  fundAddress,
  getArgs,
  wait,
} from '../test-helpers/common'

const { EXTERNAL_ADAPTER_URL } = getArgs(['EXTERNAL_ADAPTER_URL'])

const provider = createProvider()
const carol = ethers.Wallet.createRandom().connect(provider)
const linkTokenFactory = new contract.LinkTokenFactory(carol)
const fluxAggregatorFactory = new FluxAggregatorFactory(carol)
const adapterURL = new URL('result', EXTERNAL_ADAPTER_URL).href
const deposit = h.toWei('1000')

const [node1URL, node2URL] = ['http://node:6688', 'http://node-2:6688']
const clClient = new ChainlinkClient()

console.log(node2URL)

async function changePriceFeed(value: number) {
  const response = await fetch(adapterURL, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ result: value }),
  })
  assert(response.ok)
}

async function assertJobRun(
  jobId: string,
  count: number,
  errorMessage: string,
) {
  await assertAsync(() => {
    const jobRuns = clClient.connect(node1URL).getJobRuns()
    const jobRun = jobRuns[jobRuns.length - 1]
    return (
      clClient.connect(node1URL).getJobRuns().length === count &&
      jobRun.status === 'completed' &&
      jobRun.jobId === jobId
    )
  }, errorMessage)
}

describe('test', () => {
  it('works', () => {
    assert(true)
  })
})

describe('flux monitor eth client integration', () => {
  let linkToken: contract.Instance<contract.LinkTokenFactory>
  let fluxAggregator: contract.Instance<FluxAggregatorFactory>
  let job: JobSpec
  let node1Address: string

  beforeAll(async () => {
    clClient.connect(node1URL).login()
    node1Address = clClient.connect(node1URL).getAdminInfo()[0].address
    await fundAddress(carol.address)
    await fundAddress(node1Address)
    linkToken = await linkTokenFactory.deploy()
    await linkToken.deployed()
  })

  afterEach(async () => {
    clClient.connect(node1URL).archiveJob(job.id)
    await changePriceFeed(100) // original price
  })

  it('updates the price with a single node', async () => {
    fluxAggregator = await fluxAggregatorFactory.deploy(
      linkToken.address,
      1,
      600,
      1,
      ethers.utils.formatBytes32String('ETH/USD'),
    )
    await fluxAggregator.deployed()

    const tx1 = await fluxAggregator.addOracle(
      node1Address,
      node1Address,
      1,
      1,
      0,
    )
    await tx1.wait()
    const tx2 = await linkToken.transfer(fluxAggregator.address, deposit)
    await tx2.wait()
    const tx3 = await fluxAggregator.updateAvailableFunds()
    await tx3.wait()

    expect(await fluxAggregator.getOracles()).toEqual([node1Address])
    matchers.bigNum(
      await linkToken.balanceOf(fluxAggregator.address),
      deposit,
      'Unable to fund FluxAggregator',
    )

    const initialJobCount = clClient.connect(node1URL).getJobs().length
    const initialRunCount = clClient.connect(node1URL).getJobRuns().length

    // create FM job
    fluxMonitorJob.initiators[0].params.address = fluxAggregator.address
    job = clClient.connect(node1URL).createJob(JSON.stringify(fluxMonitorJob))
    assert.equal(
      clClient.connect(node1URL).getJobs().length,
      initialJobCount + 1,
    )

    // Job should trigger initial FM run
    await assertJobRun(job.id, initialRunCount + 1, 'initial job never run')
    matchers.bigNum(10000, await fluxAggregator.latestAnswer())

    // Nominally change price feed
    await changePriceFeed(101)
    await wait(10000)
    assert.equal(
      clClient.connect(node1URL).getJobRuns().length,
      initialRunCount + 1,
      'Flux Monitor should not run job after nominal price deviation',
    )

    // Significantly change price feed
    await changePriceFeed(110)
    await assertJobRun(job.id, initialRunCount + 2, 'second job never run')
    await wait(10000)
    matchers.bigNum(11000, await fluxAggregator.latestAnswer())
  })
})
