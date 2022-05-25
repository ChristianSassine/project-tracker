import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, finalize, Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

type Callback = ()=> void;
@Injectable({
    providedIn: 'root',
})
export class HttpHandlerService {
    constructor(private readonly http: HttpClient) {}

    private baseUrl = environment.serverUrl;

    loginRequest(username: string, password: string): Observable<{}> {
        return this.http.post(`${this.baseUrl}/auth/login`, { username, password }, { withCredentials: true });
    }

    validateAuth(): Observable<{}> {
        return this.http.get(`${this.baseUrl}/auth/validate`, { withCredentials: true });
    }

    // TODO : implement refresh token
    // refreshToken(): Observable<{}>{
    //     return this.http.get(`${this.baseUrl}/auth/refresh`, { withCredentials: true });
    // }

    createAccountRequest(username: string, email: string, password: string): Observable<{}> {
        return this.http.post(`${this.baseUrl}/auth/create`, { username, email, password }, { withCredentials: true });
    }

    initiateRequest(callback: Callback) {
        return this.validateAuth().pipe(finalize(callback));
    }
}
