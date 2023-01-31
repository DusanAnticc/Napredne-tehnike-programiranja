import { Component, OnInit } from '@angular/core';
import { GetUser, User } from 'src/modules/app/model/User';
import { AuthService } from 'src/modules/app/services/auth.service';
import { UserService } from 'src/modules/app/services/user.service';

@Component({
  selector: 'app-admin-page',
  templateUrl: './admin-page.component.html',
  styleUrls: ['./admin-page.component.scss'],
})
export class AdminPageComponent implements OnInit {
  constructor() {}

  ngOnInit(): void {}
}
