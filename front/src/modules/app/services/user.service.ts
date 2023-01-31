import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { from, Observable } from 'rxjs';
import { GetRepairman, GetUser, UpdateUser, User } from '../model/User';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  private headers = new HttpHeaders({ 'Content-Type': 'application/json' });

  constructor(private http: HttpClient, private toastr: ToastrService) {}

  update(user: UpdateUser): void {
    from(
      fetch('http://localhost:8080/api/users/update', {
        method: 'POST',
        body: JSON.stringify(user),
      }).then((response) => response.json())
    ).subscribe((response) => {
      this.toastr.success('Update successfull!');
    });
  }

  getAll(): Observable<GetUser[]> {
    return from(
      fetch('http://localhost:8080/api/users/get-all-users', {
        method: 'GET',
      }).then((response) => response.json())
    );
  }
  getAllRepairmans(): Observable<GetRepairman[]> {
    return from(
      fetch('http://localhost:8080/api/users/get-all-repairman', {
        method: 'GET',
      }).then((response) => response.json())
    );
  }

  ban(username: string): void {
    from(
      fetch('http://localhost:8080/api/users/ban/' + username, {
        method: 'POST',
      }).then((response) => response.json())
    ).subscribe((response) => {
      this.toastr.success('Ban successfull!');
    });
  }

  getUsername(id: string): Observable<GetUser> {
    return from(
      fetch('http://localhost:8080/api/users/find/' + id, {
        method: 'GET',
      }).then((response) => response.json())
    );
  }
}
