import { Component, OnInit } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { Repairman } from '../../model/User';
import { AdminService } from '../../services/admin.service';

@Component({
  selector: 'app-create-repairman',
  templateUrl: './create-repairman.component.html',
  styleUrls: ['./create-repairman.component.scss'],
})
export class CreateRepairmanComponent implements OnInit {
  user = <Repairman>{};
  constructor(
    private adminService: AdminService,
    private toastrService: ToastrService
  ) {}

  ngOnInit(): void {}

  saveChanges() {
    this.user.FirstName = (<HTMLInputElement>(
      document.getElementById('firstName')
    )).value;
    this.user.LastName = (<HTMLInputElement>(
      document.getElementById('lastName')
    )).value;
    this.user.EmailAddress = (<HTMLInputElement>(
      document.getElementById('emailAddress')
    )).value;
    this.user.Address = (<HTMLInputElement>(
      document.getElementById('address')
    )).value;
    this.user.Occupation = (<HTMLInputElement>(
      document.getElementById('occupation')
    )).value;
    this.user.Price = Number(
      (<HTMLInputElement>document.getElementById('price')).value
    );
    this.user.Password = (<HTMLInputElement>(
      document.getElementById('password')
    )).value;
    this.user.Username = (<HTMLInputElement>(
      document.getElementById('username')
    )).value;

    this.adminService.createRepairman(this.user);
    window.location.reload();
  }
}
