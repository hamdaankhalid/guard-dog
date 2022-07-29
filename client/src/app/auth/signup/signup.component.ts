import { Component, OnInit } from '@angular/core';
import { UntypedFormBuilder, UntypedFormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss']
})
export class SignupComponent implements OnInit {

  signupForm!: UntypedFormGroup;
  loading = false;
  submitted = false;

  constructor(
    private formBuilder: UntypedFormBuilder,
    private router: Router,
    private authService: AuthService
  ) {}

  ngOnInit(): void {
    this.authService.currentUser.subscribe((val: any) => {
      if (val) {
        this.router.navigate(['/home']);
      }
    });

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

    this.authService.signup(this.signupForm.value.email, this.signupForm.value.password).subscribe((response: any) => {
      if (response.status === 200) {
        this.loading = false;
        this.router.navigate(['/login']);
      } else {
        // show error
        this.loading = false;
      }
    }, (_: any) => {
        // show error
        this.loading = false;
        return;
    });
  }

}
