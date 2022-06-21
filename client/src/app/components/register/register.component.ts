import { Component, ElementRef, Input, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { AuthService } from 'src/app/services/auth.service';
import { Paths } from 'src/common/paths';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.scss'],
})
export class RegisterComponent implements OnInit, OnDestroy {
    @ViewChild('registerButton', { read: ElementRef, static: false }) private registerButton!: ElementRef;
    form: FormGroup;
    title: string = 'User Registration';

    usernameLabel: string = 'Username';
    passwordLabel: string = 'Password';
    emailLabel: string = 'Email';
    usernamePlaceholder: string = 'Ex: Bart';
    passwordPlaceholder: string = 'Ex: 123IloveCookies';
    emailPlaceholder: string = 'Ex: BartCookies@simpsons.com';
    usernameError: string = 'A username is required';
    passwordError: string = 'A password is required';
    emailError: string = "The email isn't a valid format";
    registerButtonText: string = 'Register';

    @Input() isLoading: boolean;
    requestFailed: boolean;
    buttonHeight: number;

    private requestSubscription: Subscription;

    constructor(private fb: FormBuilder, private authService: AuthService, private router: Router) {
        this.form = this.fb.group({
            username: ['', Validators.required],
            password: ['', Validators.required],
            email: ['', Validators.email],
        });
        this.buttonHeight = 0;
        this.requestFailed = false;
        this.requestSubscription = new Subscription();
    }

    ngOnInit(): void {
        this.requestSubscription = this.authService.creationObservable.subscribe((requestPassed) => {
            if (requestPassed) this.router.navigate([Paths.Home]);
            else this.requestFailed = true;
        });
    }

    ngOnDestroy(): void {
        this.requestSubscription.unsubscribe();
    }

    submit() {
        if (this.form.get('username')?.valid && this.form.get('password')?.valid) {
            this.buttonHeight = this.registerButton.nativeElement.offsetHeight;
            this.authService.createAccount(this.form.value.username, this.form.value.email, this.form.value.password);
        }
    }
}
