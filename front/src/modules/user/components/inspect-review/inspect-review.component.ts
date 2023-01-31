import { Component, OnInit } from '@angular/core';
import { Review } from '../../model/Review';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-inspect-review',
  templateUrl: './inspect-review.component.html',
  styleUrls: ['./inspect-review.component.scss'],
})
export class InspectReviewComponent implements OnInit {
  reviews: Review[] = [];
  review = <Review>{};
  term: string = '';
  currentId: any = '';
  currentReviewId: any = '';
  currentUserUsername: any = '';
  currentRepairmanUsername: any = '';
  val: any = '';
  constructor(private userService: UserService) {}

  ngOnInit(): void {
    this.currentId = localStorage.getItem('currentId');
    this.userService.getUsername(this.currentId).subscribe((response) => {
      this.currentUserUsername = response.Username;
      this.userService
        .getAllReviews(response.Username)
        .subscribe((response) => {
          this.reviews = response;
        });
    });
  }

  response() {
    //this.review.ClientUsername = this.currentUserUsername;
    this.review.ClientUsername = '';
    this.review.RepairmanUsername = this.currentRepairmanUsername;
    this.review.Deleted = false;
    this.review.Report = false;
    this.review.ResponseId = '';
    this.review.Text = (<HTMLInputElement>(
      document.getElementById('response')
    )).value;
    this.userService.response(this.currentReviewId, this.review);

    this.currentRepairmanUsername = '';
    this.val = '';
  }

  report(id: string) {
    this.userService.report(id);
  }

  response1(id: string, name: string) {
    this.currentReviewId = id;
    this.currentRepairmanUsername = name;
  }
}
