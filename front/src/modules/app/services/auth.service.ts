import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { ToastrService } from 'ngx-toastr';
import { Observable, from } from 'rxjs';
import { Login } from '../model/Login';
import { Register } from '../model/User';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private headers = new HttpHeaders({ 'Content-Type': 'application/json' });

  constructor(private http: HttpClient, private toastr: ToastrService) {}

  login(auth: Login): Observable<string> {
    return from(
      fetch('http://localhost:8080/api/users/login', {
        method: 'POST',
        body: JSON.stringify(auth),
      }).then((response) => response.json())
    );
  }

  register(newUser: Register): void {
    from(
      fetch('http://localhost:8080/api/users/register', {
        method: 'POST',
        body: JSON.stringify(newUser),
      }).then((response) => response.json())
    ).subscribe(() => {
      this.toastr.success('User created!');
    });
  }

  logout(): Observable<String> {
    localStorage.removeItem('currentUser');
    localStorage.removeItem('userToken');

    return this.http.get('backend/api/auth/logOut', {
      headers: this.headers,
      responseType: 'text',
    });
  }
}
