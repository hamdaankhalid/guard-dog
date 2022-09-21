import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { BehaviorSubject, Observable } from 'rxjs';
import { User } from '../models/user';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private readonly AUTH_API_URL = "http://localhost:8080/api/video_storage";

  private currentUserSubject: BehaviorSubject<User | null>;
  public currentUser: Observable<User | null>;

  constructor(private httpClient: HttpClient, private router: Router) {
    const userFromLocal = localStorage.getItem('currentUser') == null ? null : JSON.parse(localStorage.getItem('currentUser') || '{}');
    this.currentUserSubject = new BehaviorSubject<any>(userFromLocal);
    this.currentUser = this.currentUserSubject.asObservable();
   }

  signup(email: string, password: string) {
    return this.httpClient.post(this.AUTH_API_URL + "/signup", {
      email, password
    });
  }

  signin(email: string, password: string) {
    const body = new URLSearchParams();
    body.set('email', email);
    body.set('password', password);
    
    return this.httpClient.post<any>(this.AUTH_API_URL + "/login", body.toString(), {
      observe: "response",
      headers: new HttpHeaders().set('Content-Type', 'application/x-www-form-urlencoded')
    }).pipe(
      map((res: any) => {
        localStorage.setItem('access_token', res.headers.get('access_token'));
        localStorage.setItem('refresh_token', res.headers.get('refresh_token'));
      })
    );
  }

  identity() {
    return this.httpClient.get(this.AUTH_API_URL + "/identity", {
       headers: new HttpHeaders().set("Authorization", `Bearer ${localStorage.getItem('access_token')}`)
      }).pipe(map(user  => {
      // store user details and jwt token in local storage to keep user logged in between page refreshes
      localStorage.setItem('currentUser', JSON.stringify(user));
      this.currentUserSubject.next(user as any);
      return user;
    }));
  }

  logout() {
      localStorage.removeItem('access_token');
      localStorage.removeItem('refresh_token');
      localStorage.removeItem('currentUser');
      this.currentUserSubject.next(null);
      this.router.navigate(['/login']);
    }
}
