import { AsyncPipe, NgIf } from '@angular/common';
import { Component } from '@angular/core';
import { FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { AuthService, SignupFormDto } from '../../auth.service';
import { Router } from '@angular/router';
import { BehaviorSubject } from 'rxjs';

@Component({
  selector: 'app-signup',
  standalone: true,
  imports: [NgIf, FormsModule, ReactiveFormsModule, AsyncPipe],
  providers: [AuthService],
  templateUrl: './signup.component.html',
  styleUrl: './signup.component.css'
})
export class SignupComponent {

  onSuccessfulMessage$: BehaviorSubject<string> = new BehaviorSubject('');
  onErrorMessage$: BehaviorSubject<string> = new BehaviorSubject('');

  signupForm: FormGroup = new FormGroup({
    firstName: new FormControl('', [Validators.required]),
    lastName: new FormControl('', [Validators.required]),
    username: new FormControl('', [Validators.required]),
    email: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required]),
    acceptedTnC: new FormControl(false, [Validators.requiredTrue]),
  })

  constructor(private authService: AuthService, private router: Router) { }

  onSubmit() {
    if (this.signupForm.invalid) return;
    const signupForm: SignupFormDto = this.signupForm.value;
    this.authService.signup$(signupForm).subscribe(
      (data) => {
        console.log(data);
        this.onSuccessfulMessage$.next(data.data);
        this.router.navigate(['/login']);
 
      }, (err) => {
        console.log(err.error);
        this.onErrorMessage$.next(err.error.error)
      }
    )

  }
}
