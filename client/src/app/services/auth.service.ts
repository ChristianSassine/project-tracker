import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { firstValueFrom, Observable } from 'rxjs';

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    constructor(private httpClient: HttpClient) {}

    login(username: string, password: string): Promise<Object> {
        return firstValueFrom(this.httpClient.post('http://localhost:8080/', { username, password }, { withCredentials: true }));
    }

    async isLoggedIn(): Promise<boolean> {
        let logged = false;
        await firstValueFrom(this.httpClient.get('http://localhost:8080/', { withCredentials: true })).then(()=>{
            logged = true
        }).catch(()=> console.log("not valid"))
        return logged;
    }
}
