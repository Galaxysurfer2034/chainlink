import { JsonApiResponse } from 'json-api-normalizer'

export interface ErrorItem {
  status: number
  detail: any
}

export interface DocumentWithErrors {
  errors: any
}

export class AuthenticationError extends Error {
  errors: ErrorItem[]

  constructor(response: Response) {
    super(`AuthenticationError(${response.statusText})`)
    this.errors = [
      {
        status: response.status,
        detail: response.statusText,
      },
    ]
  }
}

export class BadRequestError extends Error {
  errors: JsonApiResponse['errors']

  constructor({ errors }: Pick<JsonApiResponse, 'errors'>) {
    super('BadRequestError')
    this.errors = errors
  }
}

export class ServerError extends Error {
  errors: ErrorItem[]

  constructor(errors: ErrorItem[]) {
    super('ServerError')
    this.errors = errors
  }
}

export class UnknownResponseError extends Error {
  errors: ErrorItem[]

  constructor(response: Response) {
    super(`UnknownResponseError(${response.statusText})`)
    this.errors = [
      {
        status: response.status,
        detail: response.statusText,
      },
    ]
  }
}
