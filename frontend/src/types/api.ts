export interface APIResponse<T> {
  success: boolean
  data: T
  error: APIError
}

interface APIError {
  name: string
  message: string
  details: any
}