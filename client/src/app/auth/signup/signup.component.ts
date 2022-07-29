import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss']
})
export class SignupComponent implements OnInit {

  signupForm!: FormGroup;
  loading = false;
  submitted = false;

  constructor(
    private formBuilder: FormBuilder,
    private router: Router,
    private authService: AuthService
  ) {
    // redirect to home if already logged in
    // if (this.authService.currentUser) {
    //   this.router.navigate(['/']);
    // }
   }

  ngOnInit(): void {
    this.signupForm = this.formBuilder.group({
      email: ['', Validators.required],
      password: ['', [Validators.required, Validators.minLength(6)]],
    });
  }

  get f() { return this.signupForm.controls; }

  onSubmit() {
    this.submitted = true;
    this.loading = true;

    if (this.signupForm.invalid) {
      this.loading = false;
      return;
    }

    // make a post request if we succeed forward to login page
    console.log(this.signupForm.value);

    this.authService.signup(this.signupForm.value.email, this.signupForm.value.password).subscribe((_) => {
      this.loading = false;
      this.router.navigate(['/']);
      return;
    }, (_) => {
      this.loading = false;
      return;
    })
    // otherwise show error
  }

}
