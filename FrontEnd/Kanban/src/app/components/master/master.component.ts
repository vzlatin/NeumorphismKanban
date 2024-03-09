import { Component, OnDestroy, OnInit } from '@angular/core';
import { BoardComponent } from '../board/board.component';
import { SidebarComponent } from '../sidebar/sidebar.component';
import { TopBarComponent } from '../top-bar/top-bar.component';
import { RouterModule } from '@angular/router';
import { HttpService } from '../../services/http.service';
import { Subscription } from 'rxjs';
import { Board } from '../../interfaces/board';

@Component({
    selector: 'app-master',
    standalone: true,
    imports: [BoardComponent, SidebarComponent, TopBarComponent, RouterModule],
    templateUrl: './master.component.html',
    styleUrl: './master.component.css'
})
export class MasterComponent implements OnInit, OnDestroy {

    boards: Board[] = [];
    boardsDataSubscription: Subscription | undefined;
    
    constructor(private http: HttpService) { }

    ngOnInit(): void {
        // Get all the boards at the component init.
        this.boardsDataSubscription = this.http.getBoards("/getBoards").subscribe((data) => {
            console.log(data);
            this.boards = data;
        });
    }

    ngOnDestroy(): void {
        this.boardsDataSubscription?.unsubscribe();
    }
}
