import { MessageType } from "../enums/new-message-type";
import { Board } from "./board";
import { Column } from "./column";
import { Task } from "./task";

export interface MessagePayload {
    type: MessageType
    payload: Board | Column | Task
}