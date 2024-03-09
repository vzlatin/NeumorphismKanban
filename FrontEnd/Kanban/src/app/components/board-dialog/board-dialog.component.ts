import { Component } from '@angular/core';
import { WebSocketService } from '../../services/web-socket.service';
import { Board } from '../../interfaces/board';
import { FormsModule } from '@angular/forms';
import { MessageType } from '../../enums/new-message-type';
import { DialogRef } from '@angular/cdk/dialog';

@Component({
  selector: 'app-board-dialog',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './board-dialog.component.html',
  styleUrl: './board-dialog.component.css'
})
export class BoardDialogComponent {

    protected boardName: string = ""

    constructor(private ws: WebSocketService,
                        public dialogRef: DialogRef) {}

    createBoard(): void {
        const newBoard: Board = {

            Title: this.boardName
        }
        this.ws.sendMessage(newBoard, MessageType.NewBoard);
        this.dialogRef.close();
    }
}
