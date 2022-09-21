import { Component, OnInit } from '@angular/core';
import { UntypedFormBuilder, UntypedFormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  loginForm!: UntypedFormGroup;
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

    this.loginForm = this.formBuilder.group({
      email: ['', Validators.required],
      password: ['', [Validators.required]]
    });
  }

  get f() { return this.loginForm.controls; }

  onSubmit() {
    this.submitted = true;
    this.loading = true;

    if (this.loginForm.invalid) {
      this.loading = false;
      return;
    }

    // make a post request if we succeed forward to home page
    console.log(this.loginForm.value);
    
    this.authService.signin(this.loginForm.value.email, this.loginForm.value.password).subscribe((response: any) => {
        
        this.authService.identity().subscribe((_: any) => {
          this.loading = false;
          this.router.navigate(['/home']); 
        });
        
    }, (_: any) => {
        // show error
        this.loading = false;
        console.error(_);
        return;
    })
  }
}
