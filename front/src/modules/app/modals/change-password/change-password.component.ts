import { Component, OnInit } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { User } from '../../model/User';
import { UserService } from '../../services/user.service';
//import { MdbModalRef } from 'mdb-angular-ui-kit/modal';

@Component({
  selector: 'app-change-password',
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.scss'],
})
export class ChangePasswordComponent implements OnInit {
  user: User | any;
  constructor(
    // public modalRef: MdbModalRef<ChangePasswordComponent>,
    private userService: UserService,
    private toastrService: ToastrService
  ) {}

  ngOnInit(): void {}
  saveChanges() {
    if (
      (<HTMLInputElement>document.getElementById('password1')).value === '' ||
      (<HTMLInputElement>document.getElementById('password2')).value === ''
    ) {
      this.toastrService.error('All fields must be filled!');
    } else if (
      (<HTMLInputElement>document.getElementById('password1')).value !==
      (<HTMLInputElement>document.getElementById('password2')).value
    ) {
      this.toastrService.error('Passwords do not match!');
    } else {
      this.user.password = (<HTMLInputElement>(
        document.getElementById('password1')
      )).value;
      this.userService.update(this.user);
      window.location.href = 'http://localhost:4200/login';
    }
  }
}
