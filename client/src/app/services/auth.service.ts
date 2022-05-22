import { Injectable } from '@angular/core';
import { firstValueFrom, Observable } from 'rxjs';
import { HttpHandlerService } from './http-handler.service';

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    constructor(private httpHandler: HttpHandlerService) {}

    login(username: string, password: string): Promise<{}> {
        return firstValueFrom(this.httpHandler.loginRequest(username, password));
    }

    createAccount(username: string, email: string, password: string): Observable<{}> {
        return this.httpHandler.createAccountRequest(username, email, password);
    }

    async isLoggedIn(): Promise<boolean> {
        let isLoggedIn = false;
        await firstValueFrom(this.httpHandler.validateAuth())
            .then(() => (isLoggedIn = true))
            .catch();
        return isLoggedIn;
    }
}
