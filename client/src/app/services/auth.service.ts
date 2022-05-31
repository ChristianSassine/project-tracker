import { Injectable } from '@angular/core';
import { finalize, firstValueFrom, lastValueFrom, Observable, Subject, tap, timeout } from 'rxjs';
import { HttpHandlerService, tapOnSubscribe } from './http-handler.service';

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    ongoingRequestObservable: Subject<boolean>;
    logoutObservable : Subject<boolean>;
    username : string;

    constructor(private httpHandler: HttpHandlerService) {
        this.ongoingRequestObservable = new Subject();
        this.logoutObservable = new Subject();
        this.username = '';
    }

    login(username: string, password: string): Promise<{}> {
        return firstValueFrom(
            this.httpHandler.loginRequest(username, password).pipe(
                tapOnSubscribe(() => this.ongoingRequestObservable.next(true)),
                finalize(() => this.ongoingRequestObservable.next(false)),
            ),
        );
    }

    createAccount(username: string, email: string, password: string): Observable<{}> {
        return this.httpHandler.createAccountRequest(username, email, password).pipe(
            tapOnSubscribe(() => this.ongoingRequestObservable.next(true)),
            finalize(() => this.ongoingRequestObservable.next(false)),
        );
    }

    async isLoggedIn(): Promise<boolean> {
        let isLoggedIn = false;
        await lastValueFrom(this.httpHandler.validateAuth())
            .then((username) => {
                this.username = username
                isLoggedIn = true;
            })
        return isLoggedIn;
    }

    logout(){
        const timeToExpireToken = 1050;
        this.httpHandler.logoutRequest().subscribe(()=>{setTimeout(()=> this.logoutObservable.next(true), timeToExpireToken)});
    }
}
