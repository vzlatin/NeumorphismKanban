import { Injectable } from '@angular/core';
import { WebSocketSubject, webSocket } from 'rxjs/webSocket';
import { Board } from '../interfaces/board';
import { Column } from '../interfaces/column';
import { Task } from '../interfaces/task';
import { MessageType } from '../enums/new-message-type';
import { MessagePayload } from '../interfaces/messagePayload';

@Injectable({
    providedIn: 'root'
})
export class WebSocketService {

    private readonly URL = 'ws://172.27.102.53:8080/ws';
    public subject: WebSocketSubject<MessagePayload> = webSocket<MessagePayload>(this.URL);;
    
    constructor() { 
        if (typeof WebSocket === 'undefined') {
            // WebSocket API is available, so create WebSocketSubject
            console.log('WebSocket API is not available in this environment.');
          }
    }

    sendMessage(msg: Board | Column | Task, msgType: MessageType): void {
        const messageblob = {
            type: msgType,
            payload: msg
        };
        this.subject.next(messageblob);
    }

}
