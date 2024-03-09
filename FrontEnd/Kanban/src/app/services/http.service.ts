import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Board } from '../interfaces/board';
import { Column } from '../interfaces/column';

@Injectable({
    providedIn: 'root'
})
export class HttpService {

    private BASE_URL = "http://localhost:8080";

    constructor(private http: HttpClient) { }

    getBoards(url: string): Observable<Board[]> {
        return this.http.get<Board[]>(`${this.BASE_URL}` + url);
    }

    getBoardIdColumns(url: string, boardId: string): Observable<Column[]> {
        return this.http.post<Column[]>(`${this.BASE_URL}` + url, boardId);
    }
}
