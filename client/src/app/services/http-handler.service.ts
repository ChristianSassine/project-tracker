import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
    providedIn: 'root',
})
export class HttpHandlerService {
    constructor(private readonly http: HttpClient) {}

    private baseUrl = environment.serverUrl;

	loginRequest(username: string, password: string): Observable<{}> {
		return this.http.post(`${this.baseUrl}/auth/login`, { username, password }, { withCredentials: true })
	}

	validateAuth(): Observable<{}> {
		return this.http.get(`${this.baseUrl}/auth/validate`, { withCredentials: true });
	}
}
