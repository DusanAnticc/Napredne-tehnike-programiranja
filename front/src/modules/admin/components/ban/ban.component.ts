import { Component, OnInit } from '@angular/core';
import { GetUser, User } from 'src/modules/app/model/User';
import { AuthService } from 'src/modules/app/services/auth.service';
import { UserService } from 'src/modules/app/services/user.service';

@Component({
  selector: 'app-ban',
  templateUrl: './ban.component.html',
  styleUrls: ['./ban.component.scss'],
})
export class BanComponent implements OnInit {
  users: GetUser[] = [];
  searchTerm: string = '';
  term: string = '';
  currentRole: any = '';
  constructor(
    private userService: UserService,
    private authService: AuthService
  ) {}

  ngOnInit(): void {
    this.userService.getAll().subscribe((response) => {
      this.users = response;
    });

    this.currentRole = localStorage.getItem('currentUserRole');
  }

  banBtn(user: GetUser) {
    this.userService.ban(user.Username);
    window.location.reload();
  }
}
