import { Routes } from '@angular/router';
import { BoardComponent } from './components/board/board.component';

export const routes: Routes = [
    { path: 'boards/:boardID', component: BoardComponent }
];
