import { HttpErrorResponse, HttpStatusCode } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, finalize, lastValueFrom, Subject } from 'rxjs';
import { ErrorShown } from 'src/common/error-shown';
import { ErrorNotificationService } from './error-notification.service';
import { HttpHandlerService, tapOnSubscribe } from './http-handler.service';

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    ongoingRequestObservable: Subject<boolean>;
    loginObservable: Subject<boolean>;
    creationObservable: Subject<boolean>;
    logoutObservable: Subject<boolean>;
    username: string;

    constructor(private http: HttpHandlerService, private errorNotification: ErrorNotificationService) {
        this.ongoingRequestObservable = new Subject();
        this.logoutObservable = new Subject();
        this.loginObservable = new Subject();
        this.creationObservable = new Subject();
        this.username = '';
    }

    private getUsername(): Promise<unknown> {
        return lastValueFrom(this.http.fetchUsername()).then((username) => (this.username = username));
    }

    login(username: string, password: string) {
        this.http
            .loginRequest(username, password)
            .pipe(
                tapOnSubscribe(() => this.ongoingRequestObservable.next(true)),
                finalize(() => this.ongoingRequestObservable.next(false)),
                catchError((err: HttpErrorResponse) => {
                    if (err.status === HttpStatusCode.Unauthorized) this.loginObservable.next(false);
                    else this.errorNotification.showError(ErrorShown.ServerError)
                    throw err;
                }),
            )
            .subscribe(() => {
                this.loginObservable.next(true);
            });
    }

    createAccount(username: string, email: string, password: string) {
        this.http
            .createAccountRequest(username, email, password)
            .pipe(
                tapOnSubscribe(() => this.ongoingRequestObservable.next(true)),
                finalize(() => this.ongoingRequestObservable.next(false)),
                catchError((err: HttpErrorResponse) => {
                    if (err.status === HttpStatusCode.Unauthorized) this.creationObservable.next(false);
                    else this.errorNotification.showError(ErrorShown.ServerError)
                    throw err;
                }),
            )
            .subscribe(() => this.creationObservable.next(true));
    }

    async isLoggedIn(): Promise<unknown> {
        return lastValueFrom(this.http.validateAuth()).then(() => this.getUsername().then());
    }

    logout() {
        this.http.logoutRequest().subscribe(() => {
            setTimeout(() => this.logoutObservable.next(true));
        });
    }
}
