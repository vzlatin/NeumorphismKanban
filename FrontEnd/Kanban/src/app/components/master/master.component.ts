import { Component } from '@angular/core';
import { BoardComponent } from '../board/board.component';
import { SidebarComponent } from '../sidebar/sidebar.component';
import { TopBarComponent } from '../top-bar/top-bar.component';

@Component({
  selector: 'app-master',
  standalone: true,
  imports: [BoardComponent, SidebarComponent, TopBarComponent],
  templateUrl: './master.component.html',
  styleUrl: './master.component.css'
})
export class MasterComponent {

}
