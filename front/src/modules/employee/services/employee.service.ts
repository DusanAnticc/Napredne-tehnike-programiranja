import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { from } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class EmployeeService {
  private headers = new HttpHeaders({ 'Content-Type': 'application/json' });

  constructor(private http: HttpClient, private toastr: ToastrService) {}

  deposit(username: string, money: string): void {
    from(
      fetch(
        'http://localhost:8080/api/users/payment/' + username + '/' + money,
        {
          method: 'POST',
        }
      ).then((response) => response.json())
    ).subscribe((response) => {
      this.toastr.success('Money deposit successfull!');
    });
  }
}
