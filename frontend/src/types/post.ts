export interface Post {
  id: number
  title: string
  content: string
  category: string
  status_id: number
}

export interface PostStatus {
  id: number
  name: string
}
