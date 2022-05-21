import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.scss'],
})
export class LoginComponent {
    form: FormGroup;
	usernameLabel: string = 'Username';
    passwordLabel: string = 'Password';
    usernamePlaceholder: string = 'Ex: Bart';
    passwordPlaceholder: string = 'Ex: 123IloveCookies';
    usernameError: string = 'A username is required';
    passwordError: string = 'A password is required';
    loginButtonText: string = 'Login';
    title: string = 'User login';

    constructor(private fb : FormBuilder, private authService : AuthService, private router : Router) {
		this.form = this.fb.group({
            username: ['', Validators.required],
            password: ['', Validators.required],
        });
	}

	submit() {
		if (this.form.get('username')?.valid && this.form.get('password')?.valid){
			//TODO : Handle catch
			this.authService.login(this.form.value.username, this.form.value.password).then(()=> this.router.navigate(['/home'])).catch(()=> console.log('error occured'));
		}
	}
}