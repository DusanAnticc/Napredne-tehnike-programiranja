import { Component, OnInit } from '@angular/core';
import { GetAppointments } from '../../model/appointment';
import { RepairmanService } from '../../services/repairman.service';

@Component({
  selector: 'app-accept-appointment',
  templateUrl: './accept-appointment.component.html',
  styleUrls: ['./accept-appointment.component.scss'],
})
export class AcceptAppointmentComponent implements OnInit {
  appointments: GetAppointments[] = [];
  term: string = '';
  currentId: any = '';
  constructor(private repairmanService: RepairmanService) {}

  ngOnInit(): void {
    this.currentId = localStorage.getItem('currentId');
    this.repairmanService.getUsername(this.currentId).subscribe((response) => {
      this.repairmanService
        .getAllAppointments(response.Username)
        .subscribe((response) => {
          this.appointments = response;
        });
    });
  }

  appointmentBtn(app: GetAppointments, id: number) {
    if (app.Accepted === false) {
      this.repairmanService.acceptAppointment(app.Repairman, id);
      window.location.reload();
    }
  }
}
