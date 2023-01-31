import { Component, OnInit } from '@angular/core';

import { GetRepairman } from 'src/modules/app/model/User';
import { UserService } from 'src/modules/app/services/user.service';

@Component({
  selector: 'app-hire-repairman',
  templateUrl: './hire-repairman.component.html',
  styleUrls: ['./hire-repairman.component.scss'],
})
export class HireRepairmanComponent implements OnInit {
  users: GetRepairman[] = [];
  searchTerm: string = '';
  term: string = '';
  currentRole: any = '';
  constructor(private userService: UserService) {}

  ngOnInit(): void {
    this.userService.getAllRepairmans().subscribe((response) => {
      this.users = response;
    });

    this.currentRole = localStorage.getItem('currentUserRole');
  }
}
