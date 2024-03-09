import { Component, OnInit } from '@angular/core';
import { CdkDropListGroup } from '@angular/cdk/drag-drop';
import { ColumnComponent } from '../column/column.component';
import { NgFor } from '@angular/common';
import { Column } from '../../interfaces/column';
import { WebSocketService } from '../../services/web-socket.service';
import { ActivatedRoute } from '@angular/router';
import { DialogModule, Dialog } from '@angular/cdk/dialog';
import { ColumnDialogComponent } from '../column-dialog/column-dialog.component';
import { Observable } from 'rxjs';
import { MessageType } from '../../enums/new-message-type';
import { MessagePayload } from '../../interfaces/messagePayload';
import { HttpService } from '../../services/http.service';

@Component({
    selector: 'app-board',
    standalone: true,
    imports: [CdkDropListGroup, ColumnComponent, NgFor, DialogModule],
    templateUrl: './board.component.html',
    styleUrl: './board.component.css'
})
export class BoardComponent implements OnInit{

    columns: Column[] | undefined;
    selectedBoardId: string = "";
    
    constructor(private ws: WebSocketService,
        private route: ActivatedRoute,
        private http: HttpService,
        public dialog: Dialog) {}
        

    createColumn(): void {

        const dialogRef = this.dialog.open(ColumnDialogComponent, {
            minWidth: '300px',
        });
        
        // This unsubscribes automatically.
        (dialogRef.closed as Observable<string>).subscribe((columnName: string) => {
            // this.column.title = columnName;
            // this.column.boardId = this.selectedBoardId;
            // console.log(this.column);
            // this.ws.sendMessage(this.column, MessageType.NewColumn);
        });
    }

    createTask(): void {
        
    }

    ngOnInit(): void {
        // Load the columns with the tasks when the component is initialized;
        // No need to load everything in the master component
        // Some boards may never be accessed.
        // ONLY LOAD THE REQUIRED COLUMNS, NOT EVERYTHING.
        this.route.paramMap.subscribe(params => {
            const boardId = params.get("boardID");
            if (boardId !== null) {
                this.selectedBoardId = boardId;
            } // Handle the empty url params [Especially on startup]

           // Get the columns for a specific board. 
            this.http.getBoardIdColumns("/getBoardData", JSON.stringify(this.selectedBoardId)).subscribe((columns: Column[]) => {
                this.columns = columns
                console.log(columns);
            })
        });

        this.ws.subject.subscribe((column: MessagePayload) => {
            // console.log("Logging from the Board Component: ", column);
            // const columnPayload = column.payload as Column;
            // this.columns?.push(columnPayload);
        })
    }

    // columns: Column[] = [
    //     {
    //         title: "Backlog",
    //         boardId: "fad36ddf-3c10-42be-89fa-33636a968fc5"
    //     },
    //     {
    //         title: "Testing",
    //         boardId: "fad36ddf-3c10-42be-89fa-33636a968fc5"
    //     },
    //     {
    //         title: "Done",
    //         boardId: "fad36ddf-3c10-42be-89fa-33636a968fc5"
    //     },
    //     {
    //         title: "Done",
    //         boardId: "4a72812a-fd91-42c7-b460-9d238775ddbb"
    //     },
    //     {
    //         title: "Done",
    //         boardId: "4a72812a-fd91-42c7-b460-9d238775ddbb"
    //     }
    // ]

}
