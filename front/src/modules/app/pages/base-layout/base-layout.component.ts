import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-base-layout',
  templateUrl: './base-layout.component.html',
  styleUrls: ['./base-layout.component.scss'],
})
export class BaseLayoutComponent implements OnInit {
  currentRole: any;

  constructor(private authService: AuthService, private router: Router) {}

  ngOnInit(): void {
    this.currentRole = localStorage.getItem('currentUserRole');
  }
}
