import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, finalize, mergeMap, Observable, of, switchMap, tap } from 'rxjs';
import { environment } from 'src/environments/environment';
import { Project } from '../interfaces/project';

type Callback = ()=> void;

export function tapOnSubscribe<T>(callback: () => void) {
    return (source: Observable<T>) =>
      of({}).pipe(tap(callback), switchMap(() => source));
  }
@Injectable({
    providedIn: 'root',
})
export class HttpHandlerService {
    constructor(private readonly http: HttpClient) {}

    private baseUrl = environment.serverUrl;

    private chainAfterAuth<T>(observable : Observable<T>) : Observable<T>{
        return this.validateAuth().pipe(mergeMap(()=> observable));
    }

    loginRequest(username: string, password: string): Observable<{}> {
        return this.http.post(`${this.baseUrl}/auth/login`, { username, password }, { withCredentials: true });
    }

    validateAuth(): Observable<{}> {
        return this.http.get(`${this.baseUrl}/auth/validate`, { withCredentials: true });
    }
    
    createAccountRequest(username: string, email: string, password: string): Observable<{}> {
        return this.http.post(`${this.baseUrl}/auth/create`, { username, email, password }, { withCredentials: true });
    }

    getAllProjects(): Observable<Project[]>{
        return this.chainAfterAuth(this.http.get<Project[]>(`${environment.serverUrl}/data/projects`, {withCredentials: true}));
    }

    // TODO : implement refresh token
    // refreshToken(): Observable<{}>{
    //     return this.http.get(`${this.baseUrl}/auth/refresh`, { withCredentials: true });
    // }

}
