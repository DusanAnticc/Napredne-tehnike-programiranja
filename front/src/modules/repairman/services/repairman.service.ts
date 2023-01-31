import { Injectable } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { from, Observable } from 'rxjs';
import { GetUser, User } from 'src/modules/app/model/User';
import { Appointment, GetAppointments } from '../model/appointment';

@Injectable({
  providedIn: 'root',
})
export class RepairmanService {
  user: GetUser | undefined;
  constructor(private toastr: ToastrService) {}

  getAllAppointments(id: string): Observable<GetAppointments[]> {
    return from(
      fetch('http://localhost:8080/api/repairman/findAllAppointment/' + id, {
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

  acceptAppointment(username: string, id: number): void {
    from(
      fetch(
        'http://localhost:8080/api/repairman/acceptAppointment/' +
          id +
          '/' +
          username,
        {
          method: 'POST',
        }
      ).then((response) => response.json())
    ).subscribe((response) => {
      this.toastr.success('Accepted successfull!');
    });
  }

  createAppointment(app: Appointment): void {
    from(
      fetch('http://localhost:8080/api/repairman/createAppointment', {
        method: 'POST',
        body: JSON.stringify(app),
      }).then((response) => response.json())
    ).subscribe((response) => {
      this.toastr.success('Accepted successfull!');
    });
  }

  payAppointment(id: number): void {
    from(
      fetch('http://localhost:8080/api/repairman/payAppointment/' + id, {
        method: 'POST',
      }).then((response) => response.json())
    ).subscribe((response) => {
      this.toastr.success('Accepted successfull!');
    });
  }
}
