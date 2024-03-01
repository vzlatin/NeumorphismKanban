import { Component } from '@angular/core';
import { CdkDropListGroup } from '@angular/cdk/drag-drop';
import { ColumnComponent } from '../column/column.component';
import { NgFor } from '@angular/common';
import { Column } from '../../interfaces/column';


@Component({
  selector: 'app-board',
  standalone: true,
  imports: [CdkDropListGroup, ColumnComponent, NgFor],
  templateUrl: './board.component.html',
  styleUrl: './board.component.css'
})
export class BoardComponent {

    columns: Column[] = [
        {title: "Poop", tasks: [
            {title: "Poop"},
            {title: "Poop"},
            {title: "Poop"},
            {title: "Poop"}
        ]},
        {title: "Poop", tasks: [
            {title: "Poop"},
            {title: "Poop"},
            {title: "Poop"},
            {title: "Poop"}
        ]},
        {title: "Poop", tasks: [
            {title: "Poop"},
            {title: "Poop"},
            {title: "Poop"},
            {title: "Poop"}
        ]}
    ]

}
