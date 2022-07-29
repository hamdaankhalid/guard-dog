import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map } from 'rxjs/operators';
import { BehaviorSubject, Observable } from 'rxjs';
import { User } from '../models/user';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  
  private currentUserSubject: BehaviorSubject<User | null>;
  public currentUser: Observable<User | null>;

  constructor(private httpClient: HttpClient) {
    const userFromLocal = localStorage.getItem('currentUser') == null ? null : JSON.parse(localStorage.getItem('currentUser') || '{}');
    this.currentUserSubject = new BehaviorSubject<any>(userFromLocal);
    this.currentUser = this.currentUserSubject.asObservable();
   }

  signup(email: string, password: string) {
    return this.httpClient.post("users/signup", {
      email, password
    });
  }

  signin(email: string, password: string) {
    return this.httpClient.post<User>("users/signin", { email, password })
    .pipe(map(user => {
        // store user details and jwt token in local storage to keep user logged in between page refreshes
        localStorage.setItem('currentUser', JSON.stringify(user));
        this.currentUserSubject.next(user);
        return user;
      }));
  }

  logout() {
    return this.httpClient.post("users/signout", {}).pipe(map(_ => {
      localStorage.removeItem('currentUser');
      this.currentUserSubject.next(null);
    }));
  }
}
