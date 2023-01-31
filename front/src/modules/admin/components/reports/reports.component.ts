import { Component, OnInit } from '@angular/core';
import { Review } from 'src/modules/user/model/Review';
import { AdminService } from '../../services/admin.service';

@Component({
  selector: 'app-reports',
  templateUrl: './reports.component.html',
  styleUrls: ['./reports.component.scss'],
})
export class ReportsComponent implements OnInit {
  reviews: Review[] = [];
  term: string = '';

  constructor(private adminService: AdminService) {}

  ngOnInit(): void {
    this.adminService.getAllReviews().subscribe((response) => {
      this.reviews = response;
    });
  }

  delete(id: string) {
    this.adminService.delete(id);
    window.location.reload();
  }
}
