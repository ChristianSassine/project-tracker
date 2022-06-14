import { AfterViewInit, Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { AuthService } from 'src/app/services/auth.service';

@Component({
    selector: 'app-login-page',
    templateUrl: './login-page.component.html',
    styleUrls: ['./login-page.component.scss'],
})
export class LoginPageComponent implements OnInit, OnDestroy, AfterViewInit {
    private stateSubscription: Subscription;
    private static fullTitle = 'Project Manager';
    displayedTitle: string;
    isLoading: boolean;

    constructor(private authService: AuthService) {
        this.isLoading = false;
        this.displayedTitle = '';
    }

    ngOnInit(): void {
        this.stateSubscription = this.authService.ongoingRequestObservable.subscribe((state) => (this.isLoading = state));
    }

    ngAfterViewInit(): void {
        this.titleWriter();
    }

    ngOnDestroy(): void {
        this.stateSubscription.unsubscribe();
    }

    private titleWriter(){
        const writingSpeed = 150;
        let i = 0;
        const writer = setInterval(()=>{
            if (i >= LoginPageComponent.fullTitle.length) {
                clearInterval(writer)
                return;
            }
            this.displayedTitle += LoginPageComponent.fullTitle[i];
            i++;
        }, writingSpeed);
    }
}
