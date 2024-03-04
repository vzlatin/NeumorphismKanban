import { Component } from '@angular/core';
import { CdkDropListGroup } from '@angular/cdk/drag-drop';
import { ColumnComponent } from '../column/column.component';
import { NgFor } from '@angular/common';
import { Column } from '../../interfaces/column';
import { WebSocketService } from '../../services/web-socket.service';
import { MessageType } from '../../enums/new-message-type';





@Component({
    selector: 'app-board',
    standalone: true,
    imports: [CdkDropListGroup, ColumnComponent, NgFor],
    templateUrl: './board.component.html',
    styleUrl: './board.component.css'
})
export class BoardComponent {

    constructor(private ws: WebSocketService) { }

    SendColumn(): void {
        const column = {
            title: "New Board",
            boardId: "some uuid"
        }
        this.ws.sendMessage(column, MessageType.NewColumn);
    }

    columns: Column[] = [
        {
            title: "Backlog", tasks: [
                { title: "Task: 1" },
                { title: "Task: 2" },
                { title: "Task: 3" },
                { title: "Task: 4" }
            ]
        },
        {
            title: "Testing", tasks: [
                { title: "Task: 3" },
                { title: "Task: 4" },
                { title: "Task: 5" },
                { title: "Task: 6" }
            ]
        },
        {
            title: "Done", tasks: [
                { title: "Task: 7" },
                { title: "Task: 8" },
                { title: "Task: 9" },
                { title: "Task: 10" }
            ]
        }
    ]

}
