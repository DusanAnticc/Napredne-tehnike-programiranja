import { Component, OnInit } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { Appointment } from '../../model/appointment';
import { RepairmanService } from '../../services/repairman.service';

@Component({
  selector: 'app-create-appointment',
  templateUrl: './create-appointment.component.html',
  styleUrls: ['./create-appointment.component.scss'],
})
export class CreateAppointmentComponent implements OnInit {
  app = <Appointment>{};
  constructor(
    private repairmanService: RepairmanService,
    private toastrService: ToastrService
  ) {}

  ngOnInit(): void {}

  saveChanges() {
    this.app.Username = (<HTMLInputElement>(
      document.getElementById('username')
    )).value;
    this.app.Repairman = (<HTMLInputElement>(
      document.getElementById('repairman')
    )).value;
    this.app.DateOfAppointment = (<HTMLInputElement>(
      document.getElementById('dateOfAppointment')
    )).value;
    this.app.TimeOfAppointment = (<HTMLInputElement>(
      document.getElementById('timeOfAppointment')
    )).value;
    this.app.Price = Number(
      (<HTMLInputElement>document.getElementById('price')).value
    );

    this.repairmanService.createAppointment(this.app);
    window.location.reload();
  }
}
