import { Component, OnInit } from '@angular/core';
import { Repairman } from 'src/modules/admin/model/User';
import { GetAppointments } from 'src/modules/repairman/model/appointment';
import { RepairmanService } from 'src/modules/repairman/services/repairman.service';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-pay-bill',
  templateUrl: './pay-bill.component.html',
  styleUrls: ['./pay-bill.component.scss'],
})
export class PayBillComponent implements OnInit {
  appointments: GetAppointments[] = [];
  term: string = '';
  currentId: any = '';
  currentUsername: any = '';
  constructor(
    private userService: UserService,
    private repairmanService: RepairmanService
  ) {}

  ngOnInit(): void {
    this.currentId = localStorage.getItem('currentId');
    this.userService.getUsername(this.currentId).subscribe((response) => {
      this.currentUsername = response.Username;
      this.userService
        .getAllAppointments(response.Username)
        .subscribe((response) => {
          this.appointments = response;
        });
    });
  }

  appointmentBtn(app: GetAppointments, id: number) {
    this.userService.payBill(this.currentUsername, app.Price);
    this.repairmanService.payAppointment(id);
    window.location.reload();
  }
}
