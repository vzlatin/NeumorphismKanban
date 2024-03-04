import { Injectable } from '@angular/core';
import { WebSocketSubject, webSocket } from 'rxjs/webSocket';
import { Board } from '../interfaces/board';
import { Column } from '../interfaces/column';
import { Task } from '../interfaces/task';
import { MessageType } from '../enums/new-message-type';

type MessagePayload = {type: MessageType, payload: Board | Column | Task}

@Injectable({
    providedIn: 'root'
})
export class WebSocketService {

    private readonly URL = 'ws://172.27.102.53:8080/ws';
    public subject: WebSocketSubject<MessagePayload>;
    
    constructor() { 
        if (typeof WebSocket !== 'undefined') {
            // WebSocket API is available, so create WebSocketSubject
            this.subject = webSocket<MessagePayload>(this.URL);
            this.subject.subscribe();
          } else {
            // WebSocket API is not available, handle accordingly (e.g., show error message)
            console.log('WebSocket API is not available in this environment.');
          }
    }

    sendMessage(msg: Board | Column | Task, msgType: MessageType): void {

        const messageblob = {
            type: msgType,
            payload: msg
        };

        // const serializedMessage = JSON.stringify(messageblob);
        // console.log(serializedMessage);
        this.subject.next(messageblob);
    }
}
