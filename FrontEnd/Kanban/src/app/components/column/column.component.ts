import { Component, Input } from '@angular/core';
import { TaskComponent } from '../task/task.component';
import { CdkDragDrop, CdkDrag, CdkDropList, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';
import { Column } from '../../interfaces/column';
import { NgFor } from '@angular/common';
import { Task } from '../../interfaces/task';
import { WebSocketService } from '../../services/web-socket.service';

@Component({
    selector: 'app-column',
    standalone: true,
    imports: [CdkDropList, CdkDrag, TaskComponent, NgFor],
    templateUrl: './column.component.html',
    styleUrl: './column.component.css'
})
export class ColumnComponent {
    @Input() column: Column = {};

    constructor(private ws: WebSocketService) {}

    drop(event: CdkDragDrop<Task[] | undefined>) {
        if (!event.container.data || !event.previousContainer.data) {
            return
        }

        if (event.previousContainer === event.container) {
            moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
            // send actual schema objects. Extract into a separate function
            // this.ws.sendMessage("SUP !");
        } else {
            transferArrayItem(
                event.previousContainer.data,
                event.container.data,
                event.previousIndex,
                event.currentIndex,
            );
            // send actual schema objects. Extract into a separate function
            // this.ws.sendMessage("UNSUP !");
        }
    }
}
