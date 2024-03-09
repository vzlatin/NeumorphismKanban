import { NgFor } from '@angular/common';
import { Component, Input, OnDestroy, OnInit } from '@angular/core';
import { WebSocketService } from '../../services/web-socket.service';
import { Board } from '../../interfaces/board';
import { Dialog, DialogModule } from '@angular/cdk/dialog';
import { BoardDialogComponent } from '../board-dialog/board-dialog.component';
import { MessagePayload } from '../../interfaces/messagePayload';
import { RouterModule } from '@angular/router';
import { Router } from '@angular/router';


@Component({
    selector: 'app-sidebar',
    standalone: true,
    imports: [DialogModule, NgFor, RouterModule],
    templateUrl: './sidebar.component.html',
    styleUrl: './sidebar.component.css'
})
export class SidebarComponent implements OnInit, OnDestroy {

    @Input() boards: Board[] | undefined;

    constructor(private ws: WebSocketService,
                public dialog: Dialog,
                private router: Router) {}


    openDialog(): void {
        this.dialog.open(BoardDialogComponent, {
            minWidth: '300px',
        });
    }

    selectBoard(boardId: string | undefined) {
        this.router.navigate(['/boards', boardId]);
    }

    ngOnInit(): void {
        // Subscribe to WS connection to listen for messages
        this.ws.subject.subscribe((board: MessagePayload) => {
            // This is not ok, but there's no way anything else can
            // be created / received in this component.
            const boardPayload = board.payload as Board;
            // console.log(boardPayload)
            // this.boards.push(boardPayload); !!!!! -> To broadcast changes to parent master component to update the boards.
        });
    }

    ngOnDestroy(): void {
        this.ws.subject.unsubscribe();
    }
}
