import { Injectable } from '@angular/core';
import { finalize, firstValueFrom, Observable, of, startWith, Subject, switchMap, tap } from 'rxjs';
import { HttpHandlerService } from './http-handler.service';



export function tapOnSubscribe<T>(callback: () => void) {
  return (source: Observable<T>) =>
    of({}).pipe(tap(callback), switchMap(() => source));
}

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    ongoingRequestObservable : Subject<boolean>;
    constructor(private httpHandler: HttpHandlerService) {
        this.ongoingRequestObservable = new Subject();
    }

    login(username: string, password: string): Promise<{}> {
        return firstValueFrom(this.httpHandler.loginRequest(username, password).pipe(tapOnSubscribe(() => this.ongoingRequestObservable.next(true)), finalize(()=> this.ongoingRequestObservable.next(false))));
    }

    createAccount(username: string, email: string, password: string): Observable<{}> {
        return this.httpHandler.createAccountRequest(username, email, password).pipe(tapOnSubscribe(() => this.ongoingRequestObservable.next(true)), finalize(()=> this.ongoingRequestObservable.next(false)));
    }

    async isLoggedIn(): Promise<boolean> {
        let isLoggedIn = false;
        await firstValueFrom(this.httpHandler.validateAuth())
            .then(() => (isLoggedIn = true))
            .catch();
        return isLoggedIn;
    }
}
