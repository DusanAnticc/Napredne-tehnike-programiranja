import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/modules/app/services/auth.service';

@Component({
  selector: 'app-navbar-repairman',
  templateUrl: './navbar-repairman.component.html',
  styleUrls: ['./navbar-repairman.component.scss'],
})
export class NavbarRepairmanComponent implements OnInit {
  constructor(private authService: AuthService, private router: Router) {}

  ngOnInit(): void {}

  logout() {
    this.authService.logout();
    this.router.navigate(['login']);
  }
}
