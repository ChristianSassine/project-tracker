import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { AuthService } from 'src/app/services/auth.service';

@Component({
    selector: 'app-login-page',
    templateUrl: './login-page.component.html',
    styleUrls: ['./login-page.component.scss'],
})
export class LoginPageComponent implements OnInit, OnDestroy {
    private stateSubscription: Subscription;
    isLoading: boolean;

    constructor(private authService: AuthService) {
        this.isLoading = false;
    }

    ngOnInit(): void {
        this.stateSubscription = this.authService.ongoingRequestObservable.subscribe((state) => (this.isLoading = state));
    }

    ngOnDestroy(): void {
        this.stateSubscription.unsubscribe();
    }
}
