import { Component, ElementRef, Input, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { catchError, of } from 'rxjs';
import { AuthService } from 'src/app/services/auth.service';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.scss'],
})
export class RegisterComponent {
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
    buttonHeight: number;

    constructor(private fb: FormBuilder, private authService: AuthService, private router: Router) {
        this.form = this.fb.group({
            username: ['', Validators.required],
            password: ['', Validators.required],
            email: ['', Validators.email],
        });
        this.buttonHeight = 0;
    }

    submit() {
        if (this.form.get('username')?.valid && this.form.get('password')?.valid) {
            this.buttonHeight = this.registerButton.nativeElement.offsetHeight;
            //TODO : Handle catch
            this.authService
                .createAccount(this.form.value.username, this.form.value.email, this.form.value.password)
                .pipe(catchError(() => of({})))
                .subscribe(() => this.router.navigate(['/home']));
        }
    }
}
