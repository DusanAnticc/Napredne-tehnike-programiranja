import { Component, OnInit } from '@angular/core';
import { GetUser } from 'src/modules/app/model/User';
import { GetAppointments } from '../../model/appointment';
import { RepairmanService } from '../../services/repairman.service';

@Component({
  selector: 'app-repairman-page',
  templateUrl: './repairman-page.component.html',
  styleUrls: ['./repairman-page.component.scss'],
})
export class RepairmanPageComponent implements OnInit {
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
