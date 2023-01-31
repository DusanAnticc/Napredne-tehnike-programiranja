import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { from, Observable } from 'rxjs';
import { Review } from 'src/modules/user/model/Review';
import { Repairman } from '../model/User';

@Injectable({
  providedIn: 'root',
})
export class AdminService {
  private headers = new HttpHeaders({ 'Content-Type': 'application/json' });

  constructor(private http: HttpClient, private toastr: ToastrService) {}

  createRepairman(user: Repairman): void {
    this.http.post<Repairman>(
      'http://localhost:8080/api/users/createRepairman',
      {
        method: 'POST',
        body: JSON.stringify(user),
      }
    );
  }

  getAllReviews(): Observable<Review[]> {
    return from(
      fetch('http://localhost:8080/api/admin/get-all-Reviews', {
        method: 'GET',
      }).then((response) => response.json())
    );
  }

  delete(id: string): void {
    from(
      fetch('http://localhost:8080/api/admin/deleteReview/' + id, {
        method: 'POST',
      }).then((response) => response.json())
    ).subscribe((response) => {
      this.toastr.success('Delete review successfull!');
    });
  }
}
