import { Component, ElementRef, Input, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatButton } from '@angular/material/button';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { AuthService } from 'src/app/services/auth.service';
import { Paths } from 'src/common/paths';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.scss'],
})
export class LoginComponent implements OnInit, OnDestroy {
    @ViewChild('loginButton', { read: ElementRef, static: false }) private loginButton!: ElementRef;

    usernameLabel: string = 'Username';
    passwordLabel: string = 'Password';
    usernamePlaceholder: string = 'Ex: Bart';
    passwordPlaceholder: string = 'Ex: 123IloveCookies';
    usernameError: string = 'A username is required';
    passwordError: string = 'A password is required';
    loginButtonText: string = 'Login';
    title: string = 'User login';

    form: FormGroup;
    @Input() isLoading: boolean;
    buttonHeight: number;
    requestFailed: boolean;

    private requestSubscription: Subscription;

    constructor(private fb: FormBuilder, private authService: AuthService, private router: Router) {
        this.form = this.fb.group({
            username: ['', Validators.required],
            password: ['', Validators.required],
        });

        this.isLoading = false;
        this.requestFailed = false;
        this.buttonHeight = 0;
        this.requestSubscription = new Subscription();
    }

    ngOnInit(): void {
        this.requestSubscription = this.authService.loginObservable.subscribe((requestPassed) => {
            if (requestPassed) this.router.navigate([Paths.Home]);
            else this.requestFailed = true;
        });
    }

    ngOnDestroy(): void {
        this.requestSubscription.unsubscribe();
    }

    submit() {
        if (this.form.get('username')?.valid && this.form.get('password')?.valid) {
            //TODO : Handle catch
            this.buttonHeight = this.loginButton.nativeElement.offsetHeight;
            this.authService.login(this.form.value.username, this.form.value.password);
        }
    }
}
