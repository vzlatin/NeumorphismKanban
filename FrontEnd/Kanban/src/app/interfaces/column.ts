import { Task } from "./task"

export interface Column {
    column_id: string,
    column_title: string,
    tasks: Task[]
}