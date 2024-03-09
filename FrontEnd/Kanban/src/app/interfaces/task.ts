export interface Task {
    
    task_id: string,
    task_number: string,
    task_title: string,
    task_description: string,
    created: Date,
    completed?: Date,
    due: Date,
    created_at: Date,
    updated_at: Date,
    task_priority: number,
    task_status: number,
    estimation: string,
    user_id: string
}
