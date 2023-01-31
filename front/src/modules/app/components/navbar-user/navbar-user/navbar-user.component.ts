import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/modules/app/services/auth.service';

@Component({
  selector: 'app-navbar-user',
  templateUrl: './navbar-user.component.html',
  styleUrls: ['./navbar-user.component.scss'],
})
export class NavbarUserComponent implements OnInit {
  constructor(private authService: AuthService, private router: Router) {}

  ngOnInit(): void {}

  logout() {
    this.authService.logout();
    this.router.navigate(['login']);
  }
}
