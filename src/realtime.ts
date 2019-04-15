import { getDb } from './database'
import http from 'http'
import { JobRun, fromString } from './entity/JobRun'
import { TaskRun } from './entity/TaskRun'
import WebSocket from 'ws'

export const bootstrapRealtime = (server: http.Server) => {
  const db = getDb()
  let clnodeCount = 0

  const wss = new WebSocket.Server({ server, perMessageDeflate: false })
  wss.on('connection', (ws: WebSocket) => {
    clnodeCount = clnodeCount + 1
    console.log(
      `websocket connected, total chainlink nodes connected: ${clnodeCount}`
    )
    ws.on('message', async (message: WebSocket.Data) => {
      let result

      try {
        const jobRun = fromString(message as string)
        await db.manager
          .save(jobRun)
          .then(entity => {
            jobRun.taskRuns.map(async (tr: TaskRun) => {
              tr.jobRun = entity
              await db.manager.save(tr).catch(console.error)
            })
          })
          .catch(console.error)
        result = { status: 201 }
      } catch (e) {
        console.error(e)
        result = { status: 422 }
      }

      ws.send(JSON.stringify(result))
    })

    ws.on('close', () => {
      clnodeCount = clnodeCount - 1
      console.log(
        `websocket disconnected, total chainlink nodes connected: ${clnodeCount}`
      )
    })
  })
}
