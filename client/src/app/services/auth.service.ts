import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, firstValueFrom, map, Observable, of } from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    constructor(private httpClient: HttpClient) {}

    login(username: string, password: string): Promise<Object> {
        return firstValueFrom(this.httpClient.post('http://localhost:8080/api', { username, password }, { withCredentials: true }));
    }

    async isLoggedIn(): Promise<boolean> {
        let isLoggedIn = false;
        await firstValueFrom(this.httpClient.get('http://localhost:8080/api/auth', { withCredentials: true }))
            .then(() => (isLoggedIn = true))
            .catch();
        return isLoggedIn;
    }
}
