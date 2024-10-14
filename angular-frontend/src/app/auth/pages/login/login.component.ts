import { Component } from '@angular/core';
import { AuthService, LoginFormDto } from '../../auth.service';
import { FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { AsyncPipe, JsonPipe, NgIf } from '@angular/common';
import { BehaviorSubject } from 'rxjs';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [FormsModule, ReactiveFormsModule, NgIf, AsyncPipe],
  providers: [AuthService],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {

  onErrorMessage$: BehaviorSubject<string> = new BehaviorSubject<string>('');

  loginForm: FormGroup = new FormGroup({
    username: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required]),
  });


  constructor(private authService: AuthService, private router: Router) { }

  onSubmit() {
    const loginDetails: LoginFormDto = this.loginForm.value as LoginFormDto;

    this.authService.login$(loginDetails).subscribe(
      (data) => {
        console.table(data);
        this.loginForm.reset();
        this.router.navigate(['/home']);
      }, (error) => {
        this.onErrorMessage$.next(error.error.error);
        console.error(error);
      }
    )
  }
}


