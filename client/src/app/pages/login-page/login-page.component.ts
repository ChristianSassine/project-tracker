import {Component} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import { AuthService } from 'src/app/services/auth.service';

@Component({
	selector: 'app-login-page',
	templateUrl: './login-page.component.html',
	styleUrls: ['./login-page.component.scss'],
})
export class LoginPageComponent {
	usernameLabel : string = 'Username';
	passwordLabel : string = 'Password';
	usernamePlaceholder : string = 'Ex: Bart';
	passwordPlaceholder : string = 'Ex: 123IloveCookies';
	usernameError : string = 'A username is required';
	passwordError : string = 'A password is required';
	loginButtonText : string = 'Login';
    title : string = 'Logging in';
	
	form: FormGroup;

	constructor(private fb: FormBuilder, private authService : AuthService) {
		this.form = this.fb.group({
		username: ['', Validators.required],
		password: ['', Validators.required],
		});
	}

	submit() {
		if (this.form.get('username')?.valid && this.form.get('password')?.valid){
			this.authService.login(this.form.value.username, this.form.value.password);
		}
	}
}
