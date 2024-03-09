import { DialogRef } from '@angular/cdk/dialog';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-column-dialog',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './column-dialog.component.html',
  styleUrl: './column-dialog.component.css'
})
export class ColumnDialogComponent {

    constructor(private dialogRef: DialogRef<string, ColumnDialogComponent>) {}

    columnName: string = "";

    closeDialog(): void {
        // Add validation in case no data is sent
        this.dialogRef.close(this.columnName);
    }
}
