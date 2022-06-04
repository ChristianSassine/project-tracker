import { Injectable } from '@angular/core';
import { catchError, finalize, firstValueFrom, lastValueFrom, Observable, of, Subject, tap, timeout } from 'rxjs';
import { HttpHandlerService, tapOnSubscribe } from './http-handler.service';

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    ongoingRequestObservable: Subject<boolean>;
    logoutObservable: Subject<boolean>;
    username: string;

    constructor(private http: HttpHandlerService) {
        this.ongoingRequestObservable = new Subject();
        this.logoutObservable = new Subject();
        this.username = '';
    }

    login(username: string, password: string): Promise<{}> {
        return firstValueFrom(
            this.http.loginRequest(username, password).pipe(
                tapOnSubscribe(() => this.ongoingRequestObservable.next(true)),
                finalize(() => this.ongoingRequestObservable.next(false)),
            ),
        );
    }

    createAccount(username: string, email: string, password: string): Observable<{}> {
        return this.http.createAccountRequest(username, email, password).pipe(
            tapOnSubscribe(() => this.ongoingRequestObservable.next(true)),
            finalize(() => this.ongoingRequestObservable.next(false)),
        );
    }

    async isLoggedIn(): Promise<unknown> {
        return lastValueFrom(this.http.validateAuth());
    }

    logout() {
        // const timeToExpireToken = 1050;
        this.http.logoutRequest().subscribe(() => {
            setTimeout(() => this.logoutObservable.next(true));
        });
    }
}
