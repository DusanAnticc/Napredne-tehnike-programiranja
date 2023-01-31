import { Component, OnInit } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { UpdateUser, User } from '../../model/User';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss'],
})
export class ProfileComponent implements OnInit {
  user: UpdateUser | any;
  role: any;
  email: any;
  id: any;
  constructor(
    private userService: UserService,
    private toastrService: ToastrService
  ) {}

  ngOnInit(): void {
    this.role = localStorage.getItem('currentUserRole');
    this.email = localStorage.getItem('currentEmail');
    this.id = localStorage.getItem('currentId');

    this.user = {};
    this.user.Id = '';
    this.user.FirstName = '';
    this.user.LastName = '';
    this.user.EmailAddress = '';
  }

  saveChanges() {
    if (
      (<HTMLInputElement>document.getElementById('FirstName')).value === '' ||
      (<HTMLInputElement>document.getElementById('LastName')).value === '' ||
      (<HTMLInputElement>document.getElementById('EmailAddress')).value === ''
    ) {
      this.toastrService.error('All fields must be filled!');
    } else {
      this.user.FirstName = (<HTMLInputElement>(
        document.getElementById('FirstName')
      )).value;
      this.user.LastName = (<HTMLInputElement>(
        document.getElementById('LastName')
      )).value;
      this.user.EmailAddress = (<HTMLInputElement>(
        document.getElementById('EmailAddress')
      )).value;
      this.user.Id = this.id;
      this.userService.update(this.user);
      window.location.href = 'http://localhost:4200/login';
    }
  }
}
