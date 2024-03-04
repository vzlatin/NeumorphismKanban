import { NgFor } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { WebSocketService } from '../../services/web-socket.service';
import { Board } from '../../interfaces/board';
import { Dialog, DialogModule } from '@angular/cdk/dialog';
import { BoardDialogComponent } from '../board-dialog/board-dialog.component';
import { MessagePayload } from '../../interfaces/messagePayload';

@Component({
    selector: 'app-sidebar',
    standalone: true,
    imports: [DialogModule, NgFor],
    templateUrl: './sidebar.component.html',
    styleUrl: './sidebar.component.css'
})
export class SidebarComponent implements OnInit {

    constructor(private ws: WebSocketService, public dialog: Dialog) {

        console.log("Service is present: ", this.ws.subject)

    }

    boards: Board[] = [];

    openDialog(): void {
        this.dialog.open(BoardDialogComponent, {
            minWidth: '300px',
        });
    }

    ngOnInit(): void {

        this.ws.subject.subscribe((board: MessagePayload) => {
            // This is not ok, but there's no way anything else can
            // be created / received in this component.
            const boardPayload = board.payload as Board;
            this.boards.push(boardPayload);
        });
    }
}
