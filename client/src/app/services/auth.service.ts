import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, firstValueFrom, map, Observable, of } from 'rxjs';
import { environment } from 'src/environments/environment';
import { HttpHandlerService } from './http-handler.service';

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    constructor(private httpHandler : HttpHandlerService) {}

    login(username: string, password: string): Promise<{}> {
        return firstValueFrom(this.httpHandler.loginRequest(username, password));
    }

    async isLoggedIn(): Promise<boolean> {
        let isLoggedIn = false;
        await firstValueFrom(this.httpHandler.validateAuth())
            .then(() => (isLoggedIn = true))
            .catch();
        return isLoggedIn;
    }
}
