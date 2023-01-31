import { Component, OnInit } from '@angular/core';
import { GetRepairman } from 'src/modules/app/model/User';
import { UserService } from 'src/modules/app/services/user.service';
import { Appointment } from 'src/modules/repairman/model/appointment';
import { RepairmanService } from 'src/modules/repairman/services/repairman.service';

@Component({
  selector: 'app-hire-repairman-page',
  templateUrl: './hire-repairman-page.component.html',
  styleUrls: ['./hire-repairman-page.component.scss'],
})
export class HireRepairmanPageComponent implements OnInit {
  users: GetRepairman[] = [];
  searchTerm: string = '';
  term: string = '';
  currentRole: any = '';
  app = <Appointment>{};
  currentId: any = '';
  currentUserUsername: any = '';
  currentRepairmanUsername: any = '';
  constructor(
    private userService: UserService,
    private repairmanService: RepairmanService
  ) {}

  ngOnInit(): void {
    this.userService.getAllRepairmans().subscribe((response) => {
      this.users = response;
    });

    this.currentId = localStorage.getItem('currentId');
    this.userService.getUsername(this.currentId).subscribe((response) => {
      this.currentUserUsername = response.Username;
    });

    this.currentRole = localStorage.getItem('currentUserRole');
  }

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

  hire(name: string) {
    this.currentRepairmanUsername = name;
  }
}
