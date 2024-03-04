import { NgFor } from '@angular/common';
import { Component } from '@angular/core';
import { WebSocketService } from '../../services/web-socket.service';
import { Board } from '../../interfaces/board';
import { MessageType } from '../../enums/new-message-type';

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [NgFor],
  templateUrl: './sidebar.component.html',
  styleUrl: './sidebar.component.css'
})
export class SidebarComponent {

    constructor(private ws: WebSocketService) {}

    boards = ['Marketing', 'Development', 'Sales', 'Accounting'];

    createNewBoard() {
        const boardObj: Board = {
            title: "Auiti sho board",
        }
        this.ws.sendMessage(boardObj, MessageType.NewBoard);
    }
}
