import * as h from './support/helpers'
import { assertBigNum } from './support/matchers'

contract('Coordinator', () => {
  const sourcePath = 'Coordinator.sol'
  let coordinator, link

  beforeEach(async () => {
    link = await h.linkContract()
    coordinator = await h.deploy(sourcePath, link.address)
  })

  it('has a limited public interface', () => {
    h.checkPublicABI(artifacts.require(sourcePath), [
      'EXPIRY_TIME',
      'cancelOracleRequest',
      'fulfillOracleRequest',
      'getId',
      'initiateServiceAgreement',
      'onTokenTransfer',
      'oracleRequest',
      'serviceAgreements',
      'withdraw',
      'withdrawableTokens'
    ])
  })

  const endAt = h.sixMonthsFromNow()

  describe('#getId', () => {
    const agreedPayment = 1
    const agreedExpiration = 2
    const agreedOracles = [
      '0x70AEc4B9CFFA7b55C0711b82DD719049d615E21d',
      '0xd26114cd6EE289AccF82350c8d8487fedB8A0C07'
    ]
    const requestDigest =
      '0x85820c5ec619a1f517ee6cfeff545ec0ca1a90206e1a38c47f016d4137e801dd'
    const args = [
      agreedPayment,
      agreedExpiration,
      endAt,
      agreedOracles,
      requestDigest
    ]
    const expectedBinaryArgs = [
      '0x',
      ...[agreedPayment, agreedExpiration, endAt].map(h.padNumTo256Bit),
      ...agreedOracles.map(h.pad0xHexTo256Bit),
      h.strip0x(requestDigest)
    ]
      .join('')
      .toLowerCase()

    it.skip('matches the ID generated by the oracle off-chain', async () => {
      const expectedBinaryArgsSha3 = h.keccak(expectedBinaryArgs, {
        encoding: 'hex'
      })
      const result = await coordinator.getId.call(...args)

      assert.equal(result, expectedBinaryArgsSha3)
    })
  })

  describe('#initiateServiceAgreement', () => {
    let agreement
    before(async () => {
      agreement = await h.newServiceAgreement({ oracles: [h.oracleNode] })
    })

    context('with valid oracle signatures', () => {
      it('saves a service agreement struct from the parameters', async () => {
        let tx = await h.initiateServiceAgreement(coordinator, agreement)
        await h.checkServiceAgreementPresent(coordinator, agreement)
      })

      it('returns the SAID', async () => {
        const sAID = await h.initiateServiceAgreementCall(
          coordinator,
          agreement
        )
        assert.equal(sAID, agreement.id)
      })

      it('logs an event', async () => {
        await h.initiateServiceAgreement(coordinator, agreement)
        const event = await h.getLatestEvent(coordinator)
        assert.equal(agreement.id, event.args.said)
      })
    })

    context('with an invalid oracle signatures', () => {
      let badOracleSignature, badRequestDigestAddr
      before(async () => {
        const sAID = h.calculateSAID(agreement)
        badOracleSignature = await h.personalSign(h.stranger, sAID)
        badRequestDigestAddr = h.recoverPersonalSignature(
          sAID,
          badOracleSignature
        )
        assert.equal(h.stranger.toLowerCase(), h.toHex(badRequestDigestAddr))
      })

      it('saves no service agreement struct, if signatures invalid', async () => {
        await h.assertActionThrows(async () => {
          await h.initiateServiceAgreement(
            coordinator,
            Object.assign(agreement, { oracleSignatures: [badOracleSignature] })
          )
        })
        await h.checkServiceAgreementAbsent(coordinator, agreement.id)
      })
    })

    context('Validation of service agreement deadlines', () => {
      it('Rejects a service agreement with an endAt date in the past', async () => {
        await h.assertActionThrows(async () =>
          h.initiateServiceAgreement(
            coordinator,
            Object.assign(agreement, { endAt: 1 })
          )
        )
        await h.checkServiceAgreementAbsent(coordinator, agreement.id)
      })
    })
  })

  describe('#oracleRequest', () => {
    const fHash = h.functionSelector('requestedBytes32(bytes32,bytes32)')
    const to = '0x80e29acb842498fe6591f020bd82766dce619d43'
    let agreement
    before(async () => {
      agreement = await h.newServiceAgreement({ oracles: [h.oracleNode] })
    })

    beforeEach(async () => {
      await h.initiateServiceAgreement(coordinator, agreement)
      await link.transfer(h.consumer, h.toWei(1000))
    })

    context('when called through the LINK token with enough payment', () => {
      let tx
      beforeEach(async () => {
        const payload = h.executeServiceAgreementBytes(
          agreement.id,
          to,
          fHash,
          '1',
          ''
        )
        tx = await link.transferAndCall(
          coordinator.address,
          agreement.payment,
          payload,
          { from: h.consumer }
        )
      })

      it('logs an event', async () => {
        const log = tx.receipt.rawLogs[2]
        assert.equal(coordinator.address, log.address)

        // If updating this test, be sure to update
        // services.ServiceAgreementExecutionLogTopic. (Which see for the
        // calculation of this hash.)
        const eventSignature =
          '0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65'
        assert.equal(eventSignature, log.topics[0])

        assert.equal(agreement.id, log.topics[1])
        const req = h.decodeRunRequest(tx.receipt.rawLogs[2])
        assertBigNum(
          h.consumer,
          req.requester,
          "Logged consumer address doesn't match"
        )
        assertBigNum(
          agreement.payment,
          req.payment,
          "Logged payment amount doesn't match"
        )
      })
    })

    context(
      'when called through the LINK token with not enough payment',
      () => {
        it('throws an error', async () => {
          const calldata = h.executeServiceAgreementBytes(
            agreement.id,
            to,
            fHash,
            '1',
            ''
          )
          const underPaid = h
            .bigNum(agreement.payment)
            .sub(h.bigNum(1))
            .toString()

          await h.assertActionThrows(async () => {
            await link.transferAndCall(
              coordinator.address,
              underPaid,
              calldata,
              { from: h.consumer }
            )
          })
        })
      }
    )

    context('when not called through the LINK token', () => {
      it('reverts', async () => {
        await h.assertActionThrows(async () => {
          await coordinator.oracleRequest(
            '0x0000000000000000000000000000000000000000',
            0,
            agreement.id,
            to,
            fHash,
            1,
            1,
            '0x',
            { from: h.consumer }
          )
        })
      })
    })
  })

  describe('#fulfillOracleRequest', () => {
    let agreement, mock, request
    beforeEach(async () => {
      agreement = await h.newServiceAgreement({ oracles: [h.oracleNode] })
      const tx = await h.initiateServiceAgreement(coordinator, agreement)
      assert.equal(tx.logs[0].args.said, agreement.id)
    })

    context('cooperative consumer', () => {
      beforeEach(async () => {
        mock = await h.deploy('examples/GetterSetter.sol')
        const fHash = h.functionSelector('requestedBytes32(bytes32,bytes32)')

        const payload = h.executeServiceAgreementBytes(
          agreement.id,
          mock.address,
          fHash,
          1,
          ''
        )
        const tx = await link.transferAndCall(
          coordinator.address,
          agreement.payment,
          payload,
          { value: 0 }
        )
        request = h.decodeRunRequest(tx.receipt.rawLogs[2])
      })

      context('when called by a non-owner', () => {
        // Turn this test on when multiple-oracle response aggregation is enabled
        xit('raises an error', async () => {
          await h.assertActionThrows(async () => {
            await coordinator.fulfillOracleRequest(
              request.id,
              h.toHex('Hello World!'),
              { from: h.stranger }
            )
          })
        })
      })

      context('when called by an owner', () => {
        it('raises an error if the request ID does not exist', async () => {
          await h.assertActionThrows(async () => {
            await coordinator.fulfillOracleRequest(
              '0xdeadbeef',
              h.toHex('Hello World!'),
              { from: h.oracleNode }
            )
          })
        })

        it('sets the value on the requested contract', async () => {
          await coordinator.fulfillOracleRequest(
            request.id,
            h.toHex('Hello World!'),
            { from: h.oracleNode }
          )

          const mockRequestId = await mock.requestId.call()
          assert.equal(request.id, mockRequestId)

          const currentValue = await mock.getBytes32.call()
          assert.equal('Hello World!', h.toUtf8(currentValue))
        })

        it('does not allow a request to be fulfilled twice', async () => {
          await coordinator.fulfillOracleRequest(
            request.id,
            h.toHex('First message!'),
            { from: h.oracleNode }
          )
          await h.assertActionThrows(async () => {
            await coordinator.fulfillOracleRequest(
              request.id,
              h.toHex('Second message!!'),
              { from: h.oracleNode }
            )
          })
        })
      })
    })

    context('with a malicious requester', () => {
      const paymentAmount = h.toWei(1)

      beforeEach(async () => {
        mock = await h.deploy(
          'examples/MaliciousRequester.sol',
          link.address,
          coordinator.address
        )
        await link.transfer(mock.address, paymentAmount)
      })

      xit('cannot cancel before the expiration', async () => {
        await h.assertActionThrows(async () => {
          await mock.maliciousRequestCancel(
            agreement.id,
            'doesNothing(bytes32,bytes32)'
          )
        })
      })

      it('cannot call functions on the LINK token through callbacks', async () => {
        await h.assertActionThrows(async () => {
          await mock.request(
            agreement.id,
            link.address,
            h.toHex('transfer(address,uint256)')
          )
        })
      })

      context('requester lies about amount of LINK sent', () => {
        it('the oracle uses the amount of LINK actually paid', async () => {
          const tx = await mock.maliciousPrice(agreement.id)
          const req = h.decodeRunRequest(tx.receipt.rawLogs[3])
          assertBigNum(
            paymentAmount,
            req.payment,
            [
              'Malicious data request tricked oracle into refunding more than',
              'the requester paid, by claiming a larger amount',
              `(${req.payment}) than the requester paid (${paymentAmount})`
            ].join(' ')
          )
        })
      })
    })

    context('with a malicious consumer', () => {
      const paymentAmount = h.toWei(1)

      beforeEach(async () => {
        mock = await h.deploy(
          'examples/MaliciousConsumer.sol',
          link.address,
          coordinator.address
        )
        await link.transfer(mock.address, paymentAmount)
      })

      context('fails during fulfillment', () => {
        beforeEach(async () => {
          const tx = await mock.requestData(
            agreement.id,
            h.toHex('assertFail(bytes32,bytes32)')
          )
          request = h.decodeRunRequest(tx.receipt.rawLogs[3])
        })

        // needs coordinator withdrawal functionality to meet parity
        xit('allows the oracle node to receive their payment', async () => {
          await coordinator.fulfillOracleRequest(
            request.id,
            h.toHex('hack the planet 101'),
            { from: h.oracleNode }
          )

          const balance = await link.balanceOf.call(h.oracleNode)
          assert.isTrue(balance.equals(0))

          await coordinator.withdraw(h.oracleNode, paymentAmount, {
            from: h.oracleNode
          })
          const newBalance = await link.balanceOf.call(h.oracleNode)
          assert.isTrue(paymentAmount.equals(newBalance))
        })

        it("can't fulfill the data again", async () => {
          await coordinator.fulfillOracleRequest(
            request.id,
            h.toHex('hack the planet 101'),
            { from: h.oracleNode }
          )
          await h.assertActionThrows(async () => {
            await coordinator.fulfillOracleRequest(
              request.id,
              h.toHex('hack the planet 102'),
              { from: h.oracleNode }
            )
          })
        })
      })

      context('calls selfdestruct', () => {
        beforeEach(async () => {
          const tx = await mock.requestData(
            agreement.id,
            'doesNothing(bytes32,bytes32)'
          )
          request = h.decodeRunRequest(tx.receipt.rawLogs[3])
          await mock.remove()
        })

        // needs coordinator withdrawal functionality to meet parity
        xit('allows the oracle node to receive their payment', async () => {
          await coordinator.fulfillOracleRequest(
            request.id,
            h.toHex('hack the planet 101'),
            { from: h.oracleNode }
          )

          const balance = await link.balanceOf.call(h.oracleNode)
          assert.isTrue(balance.equals(0))

          await coordinator.withdraw(h.oracleNode, paymentAmount, {
            from: h.oracleNode
          })
          const newBalance = await link.balanceOf.call(h.oracleNode)
          assert.isTrue(paymentAmount.equals(newBalance))
        })
      })

      context('request is canceled during fulfillment', () => {
        beforeEach(async () => {
          const tx = await mock.requestData(
            agreement.id,
            h.toHex('cancelRequestOnFulfill(bytes32,bytes32)')
          )
          request = h.decodeRunRequest(tx.receipt.rawLogs[3])

          const mockBalance = await link.balanceOf.call(mock.address)
          assertBigNum(mockBalance, h.bigNum(0))
        })

        // needs coordinator withdrawal functionality to meet parity
        xit('allows the oracle node to receive their payment', async () => {
          await coordinator.fulfillOracleRequest(
            request.id,
            h.toHex('hack the planet 101'),
            { from: h.oracleNode }
          )

          const mockBalance = await link.balanceOf.call(mock.address)
          assertBigNum(mockBalance, h.bigNum(0))

          const balance = await link.balanceOf.call(h.oracleNode)
          assert.isTrue(balance.equals(0))

          await coordinator.withdraw(h.oracleNode, paymentAmount, {
            from: h.oracleNode
          })
          const newBalance = await link.balanceOf.call(h.oracleNode)
          assert.isTrue(paymentAmount.equals(newBalance))
        })

        it("can't fulfill the data again", async () => {
          await coordinator.fulfillOracleRequest(
            request.id,
            h.toHex('hack the planet 101'),
            { from: h.oracleNode }
          )
          await h.assertActionThrows(async () => {
            await coordinator.fulfillOracleRequest(
              request.id,
              h.toHex('hack the planet 102'),
              { from: h.oracleNode }
            )
          })
        })
      })
    })

    context('when aggregating answers', () => {
      let oracle1, oracle2, oracle3, request, strangerOracle

      beforeEach(async () => {
        strangerOracle = h.stranger
        oracle1 = h.oracleNode1
        oracle2 = h.oracleNode2
        oracle3 = h.oracleNode3

        agreement = await h.newServiceAgreement({
          oracles: [oracle1, oracle2, oracle3]
        })
        let tx = await h.initiateServiceAgreement(coordinator, agreement)
        assert.equal(tx.logs[0].args.said, agreement.id)

        mock = await h.deploy('examples/GetterSetter.sol')
        const fHash = h.functionSelector('requestedUint256(bytes32,uint256)')

        const payload = h.executeServiceAgreementBytes(
          agreement.id,
          mock.address,
          fHash,
          1,
          ''
        )
        tx = await link.transferAndCall(
          coordinator.address,
          agreement.payment,
          payload,
          { value: 0 }
        )
        request = h.decodeRunRequest(tx.receipt.rawLogs[2])
      })

      it('does not set the value with only one oracle', async () => {
        const tx = await coordinator.fulfillOracleRequest(
          request.id,
          h.toHex(17),
          { from: oracle1 }
        )

        assert.equal(tx.receipt.rawLogs.length, 0) // No logs emitted = consuming contract not called
      })

      it('sets the average of the reported values', async () => {
        await coordinator.fulfillOracleRequest(request.id, h.toHex(16), {
          from: oracle1
        })
        await coordinator.fulfillOracleRequest(request.id, h.toHex(17), {
          from: oracle2
        })
        const lastTx = await coordinator.fulfillOracleRequest(
          request.id,
          h.toHex(18),
          { from: oracle3 }
        )

        assert.equal(lastTx.receipt.rawLogs.length, 1)
        const currentValue = await mock.getUint256.call()
        assertBigNum(h.bigNum(17), currentValue)
      })

      context('when large values are provided in response', async () => {
        // (uint256(-1) / 2) - 1
        const largeValue1 =
          '0x7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe'
        // (uint256(-1) / 2)
        const largeValue2 =
          '0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'
        // (uint256(-1) / 2) + 1
        const largeValue3 =
          '0x8000000000000000000000000000000000000000000000000000000000000000'

        beforeEach(async () => {
          await coordinator.fulfillOracleRequest(request.id, largeValue1, {
            from: oracle1
          })
          await coordinator.fulfillOracleRequest(request.id, largeValue2, {
            from: oracle2
          })
        })

        it('does not overflow', async () => {
          await coordinator.fulfillOracleRequest(request.id, largeValue3, {
            from: oracle3
          })
        })

        it('sets the average of the reported values', async () => {
          await coordinator.fulfillOracleRequest(request.id, largeValue3, {
            from: oracle3
          })
          const currentValue = await mock.getUint256.call()
          assertBigNum(h.bigNum(largeValue2), currentValue)
          assert.notEqual(0, await mock.requestId.call()) // check if called
        })
      })

      it('successfully sets average when responses equal largest uint256', async () => {
        const largest =
          '0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'

        await coordinator.fulfillOracleRequest(request.id, largest, {
          from: oracle1
        })
        await coordinator.fulfillOracleRequest(request.id, largest, {
          from: oracle2
        })
        await coordinator.fulfillOracleRequest(request.id, largest, {
          from: oracle3
        })
        const currentValue = await mock.getUint256.call()
        assertBigNum(h.bigNum(largest), currentValue)
        assert.notEqual(0, await mock.requestId.call()) // check if called
      })

      it('rejects oracles not part of the service agreement', async () => {
        await h.assertActionThrows(async () => {
          await coordinator.fulfillOracleRequest(request.id, h.toHex(18), {
            from: strangerOracle
          })
        })
      })

      context('when an oracle reports multiple times', async () => {
        beforeEach(async () => {
          await coordinator.fulfillOracleRequest(request.id, h.toHex(16), {
            from: oracle1
          })
          await coordinator.fulfillOracleRequest(request.id, h.toHex(17), {
            from: oracle2
          })

          await h.assertActionThrows(async () => {
            await coordinator.fulfillOracleRequest(request.id, h.toHex(18), {
              from: oracle2
            })
          })
        })

        it('does not set the average', async () => {
          assert.equal(0, await mock.requestId.call()) // check if called
        })

        it('still allows the other oracles to report', async () => {
          await coordinator.fulfillOracleRequest(request.id, h.toHex(18), {
            from: oracle3
          })
          const currentValue = await mock.getUint256.call()
          assertBigNum(h.bigNum(17), currentValue)
          assert.notEqual(0, await mock.requestId.call()) // check if called
        })
      })
    })

    context('after aggregation', () => {
      let oracle1, oracle2, oracle3, request

      beforeEach(async () => {
        oracle1 = h.oracleNode1
        oracle2 = h.oracleNode2
        oracle3 = h.oracleNode3

        agreement = await h.newServiceAgreement({
          oracles: [oracle1, oracle2, oracle3]
        })
        let tx = await h.initiateServiceAgreement(coordinator, agreement)
        assert.equal(tx.logs[0].args.said, agreement.id)

        mock = await h.deploy('examples/GetterSetter.sol')
        const fHash = h.functionSelector('requestedUint256(bytes32,uint256)')

        const payload = h.executeServiceAgreementBytes(
          agreement.id,
          mock.address,
          fHash,
          1,
          ''
        )
        tx = await link.transferAndCall(
          coordinator.address,
          agreement.payment,
          payload,
          { value: 0 }
        )
        request = h.decodeRunRequest(tx.receipt.rawLogs[2])

        await coordinator.fulfillOracleRequest(request.id, h.toHex(16), {
          from: oracle1
        })
        await coordinator.fulfillOracleRequest(request.id, h.toHex(17), {
          from: oracle2
        })
        await coordinator.fulfillOracleRequest(request.id, h.toHex(18), {
          from: oracle3
        })

        const currentValue = await mock.getUint256.call()
        assertBigNum(h.bigNum(17), currentValue)
      })

      it('oracle balances are updated', async () => {
        // Given the 3 oracles from the SA, each should have the following balance after fulfillment
        const expected = h.bigNum('333333333333333333')
        const balance1 = await coordinator.withdrawableTokens.call(oracle1)
        assertBigNum(expected, balance1)
      })
    })

    context('withdraw', () => {
      let oracle1, oracle2, oracle3, request

      beforeEach(async () => {
        oracle1 = h.oracleNode1
        oracle2 = h.oracleNode2
        oracle3 = h.oracleNode3

        agreement = await h.newServiceAgreement({
          oracles: [oracle1, oracle2, oracle3]
        })
        let tx = await h.initiateServiceAgreement(coordinator, agreement)
        assert.equal(tx.logs[0].args.said, agreement.id)

        mock = await h.deploy('examples/GetterSetter.sol')
        const fHash = h.functionSelector('requestedUint256(bytes32,uint256)')

        const payload = h.executeServiceAgreementBytes(
          agreement.id,
          mock.address,
          fHash,
          1,
          ''
        )
        tx = await link.transferAndCall(
          coordinator.address,
          agreement.payment,
          payload,
          { value: 0 }
        )
        request = h.decodeRunRequest(tx.receipt.rawLogs[2])

        await coordinator.fulfillOracleRequest(request.id, h.toHex(16), {
          from: oracle1
        })
        await coordinator.fulfillOracleRequest(request.id, h.toHex(17), {
          from: oracle2
        })
        await coordinator.fulfillOracleRequest(request.id, h.toHex(18), {
          from: oracle3
        })

        const currentValue = await mock.getUint256.call()
        assertBigNum(h.bigNum(17), currentValue)
      })

      it('allows the oracle to withdraw their full amount', async () => {
        const coordBalance1 = await link.balanceOf.call(coordinator.address)
        const withdrawAmount = await coordinator.withdrawableTokens.call(
          oracle1
        )
        await coordinator.withdraw(oracle1, withdrawAmount.toString(), {
          from: oracle1
        })
        const oracleBalance = await link.balanceOf.call(oracle1)
        const afterWithdrawBalance = await coordinator.withdrawableTokens.call(
          oracle1
        )
        const coordBalance2 = await link.balanceOf.call(coordinator.address)
        const expectedCoordFinalBalance = coordBalance1.sub(withdrawAmount)
        assertBigNum(withdrawAmount, oracleBalance)
        assertBigNum(expectedCoordFinalBalance, coordBalance2)
        assertBigNum(h.bigNum(0), afterWithdrawBalance)
      })

      it('rejects amounts greater than allowed', async () => {
        const oracleBalance = await coordinator.withdrawableTokens.call(oracle1)
        const withdrawAmount = oracleBalance.add(h.bigNum(1))
        await h.assertActionThrows(async () => {
          await coordinator.withdraw(oracle1, withdrawAmount.toString(), {
            from: oracle1
          })
        })
      })
    })
  })
})
