import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss']
})
export class NavbarComponent implements OnInit {
  currentUser: any;

  constructor(
    private router: Router,
    private authenticationService: AuthService
  ) {
  }

  ngOnInit() {
    this.authenticationService.currentUser.subscribe((x: any)=> {
      return this.currentUser = x;
    });
  }

  logout() {
      this.authenticationService.logout();
      this.router.navigate(['/login']);
  }

}
