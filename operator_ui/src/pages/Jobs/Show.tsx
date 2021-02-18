import React from 'react'
import { v2 } from 'api'
import { Route, RouteComponentProps, Switch } from 'react-router-dom'
import { useErrorHandler } from 'hooks/useErrorHandler'
import { useLoadingPlaceholder } from 'hooks/useLoadingPlaceholder'
import {
  generateJSONDefinition,
  generateTOMLDefinition,
} from './generateJobSpecDefinition'
import { JobData, JobSpecType, JobV2 } from './sharedTypes'
import DefinitionTab from './Definition'
import { JobsErrors } from './Errors'
import { RecentRuns } from './RecentRuns'
import { RegionalNav } from './RegionalNav'
import { Runs as JobRuns } from './Runs'
import { isOcrJob } from './utils'
import {
  transformDirectRequestJobRun,
  transformPipelineJobRun,
} from './transformJobRuns'

type Props = RouteComponentProps<{
  jobSpecId: string
}>

const DEFAULT_PAGE = 1
const RECENT_RUNS_COUNT = 5

export const JobsShow: React.FC<Props> = ({ match }) => {
  const [state, setState] = React.useState<JobData>({
    recentRuns: [],
    recentRunsCount: 0,
  })
  const { job, jobSpec } = state
  const { error, ErrorComponent, setError } = useErrorHandler()
  const { LoadingPlaceholder } = useLoadingPlaceholder(!error && !jobSpec)

  const { jobSpecId } = match.params

  const getJobSpecRuns = React.useCallback(
    ({ page = DEFAULT_PAGE, size = RECENT_RUNS_COUNT } = {}) => {
      const requestParams = {
        jobSpecId,
        page,
        size,
      }
      if (isOcrJob(jobSpecId)) {
        return v2.ocrRuns
          .getJobSpecRuns(requestParams)
          .then((jobSpecRunsResponse) => {
            setState((s) => ({
              ...s,
              recentRuns: jobSpecRunsResponse.data.map(
                transformPipelineJobRun(jobSpecId),
              ),
              recentRunsCount: jobSpecRunsResponse.meta.count,
            }))
          })
          .catch(setError)
      } else {
        return v2.runs
          .getJobSpecRuns(requestParams)
          .then((jobSpecRunsResponse) => {
            setState((s) => ({
              ...s,
              recentRuns: jobSpecRunsResponse.data.map(
                transformDirectRequestJobRun(jobSpecId),
              ),
              recentRunsCount: jobSpecRunsResponse.meta.count,
            }))
          })
          .catch(setError)
      }
    },
    [jobSpecId, setError],
  )

  const getJobSpec = React.useCallback(async () => {
    if (isOcrJob(jobSpecId)) {
      return v2.jobs
        .getJobSpec(jobSpecId)
        .then((response) => {
          const jobSpec = response.data
          setState((s) => {
            const job = {
              ...jobSpec.attributes.pipelineSpec,
              id: jobSpec.id,
              definition: generateTOMLDefinition(jobSpec.attributes),
              type: 'v2',
              name: jobSpec.attributes.name,
              specType: jobSpec.attributes.type as JobSpecType,
              errors: jobSpec.attributes.errors,
            } as JobV2

            switch (jobSpec.attributes.type) {
              case 'offchainreporting':
                job.createdAt = jobSpec.attributes.offChainReportingOracleSpec
                  ?.createdAt as string

                break
              case 'fluxmonitor':
                job.createdAt = jobSpec.attributes.fluxMonitorSpec
                  ?.createdAt as string

                break
              case 'directRequest':
                job.createdAt = jobSpec.attributes.directRequestSpec
                  ?.createdAt as string

                break
            }

            return {
              ...s,
              jobSpec,
              job,
            }
          })
        })
        .catch(setError)
    } else {
      return v2.specs
        .getJobSpec(jobSpecId)
        .then((response) => {
          const jobSpec = response.data
          setState((s) => ({
            ...s,
            jobSpec,
            job: {
              ...jobSpec.attributes,
              id: jobSpec.id,
              type: 'Direct request',
              definition: generateJSONDefinition(jobSpec.attributes),
            },
          }))
        })
        .catch(setError)
    }
  }, [jobSpecId, setError])

  React.useEffect(() => {
    getJobSpec()
  }, [getJobSpec])

  return (
    <div>
      <RegionalNav
        jobSpecId={jobSpecId}
        job={job}
        getJobSpecRuns={getJobSpecRuns}
        runsCount={state.recentRunsCount}
      />
      <Switch>
        <Route
          exact
          path={`${match.path}/definition`}
          render={() => (
            <DefinitionTab
              {...{
                ...state,
                ErrorComponent,
                LoadingPlaceholder,
                error,
              }}
            />
          )}
        />
        <Route
          exact
          path={`${match.path}/errors`}
          render={() => (
            <JobsErrors
              {...{
                ...state,
                ErrorComponent,
                LoadingPlaceholder,
                error,
                getJobSpec,
                setState,
              }}
            />
          )}
        />
        <Route
          exact
          path={`${match.path}/runs`}
          render={() => (
            <JobRuns
              {...{
                ...state,
                error,
                ErrorComponent,
                LoadingPlaceholder,
                getJobSpecRuns,
              }}
            />
          )}
        />
        <Route
          path={`${match.path}`}
          render={() => (
            <RecentRuns
              {...{
                ...state,
                error,
                ErrorComponent,
                LoadingPlaceholder,
                getJobSpecRuns,
              }}
            />
          )}
        />
      </Switch>
    </div>
  )
}

export default JobsShow
