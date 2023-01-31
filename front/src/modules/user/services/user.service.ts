import { Injectable } from '@angular/core';
import { from, Observable } from 'rxjs';
import { GetUser } from 'src/modules/app/model/User';
import { GetAppointments } from 'src/modules/repairman/model/appointment';
import { Assessment } from '../model/Assessment';
import { Review } from '../model/Review';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  toastr: any;
  constructor() {}

  assessment(assessment: Assessment): void {
    from(
      fetch(
        'http://localhost:8080/api/users/assessment/' +
          assessment.Username +
          '/' +
          assessment.Grade,
        {
          method: 'POST',
        }
      ).then((response) => response.json())
    );
  }

  createReview(rev: Review): void {
    from(
      fetch('http://localhost:8080/api/admin/createReview', {
        method: 'POST',
        body: JSON.stringify(rev),
      }).then((response) => response.json())
    );
  }

  getAllReviews(username: string): Observable<Review[]> {
    return from(
      fetch('http://localhost:8080/api/admin/get-all-Reviews/' + username, {
        method: 'GET',
      }).then((response) => response.json())
    );
  }

  getUsername(id: string): Observable<GetUser> {
    return from(
      fetch('http://localhost:8080/api/users/find/' + id, {
        method: 'GET',
      }).then((response) => response.json())
    );
  }

  report(id: string): void {
    from(
      fetch('http://localhost:8080/api/admin/reportReview/' + id, {
        method: 'POST',
      }).then((response) => response.json())
    ).subscribe((response) => {
      this.toastr.success('Report successfull!');
    });
  }

  response(id: string, review: Review): void {
    from(
      fetch('http://localhost:8080/api/admin/createReviewResponse/' + id, {
        method: 'POST',
        body: JSON.stringify(review),
      }).then((response) => response.json())
    ).subscribe((response) => {
      this.toastr.success('Responsed successfull!');
    });
  }

  getAllAppointments(id: string): Observable<GetAppointments[]> {
    return from(
      fetch(
        'http://localhost:8080/api/repairman/findAllAppointmentUser/' + id,
        {
          method: 'GET',
        }
      ).then((response) => response.json())
    );
  }

  payBill(username: string, money: number): void {
    from(
      fetch(
        'http://localhost:8080/api/users/payBill/' + username + '/' + money,
        {
          method: 'POST',
        }
      ).then((response) => response.json())
    ).subscribe((response) => {
      this.toastr.success('Accepted successfull!');
    });
  }
}
