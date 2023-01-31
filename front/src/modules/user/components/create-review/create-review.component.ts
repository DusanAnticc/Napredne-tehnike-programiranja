import { Component, OnInit } from '@angular/core';
import { Review } from '../../model/Review';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-create-review',
  templateUrl: './create-review.component.html',
  styleUrls: ['./create-review.component.scss'],
})
export class CreateReviewComponent implements OnInit {
  review = <Review>{};
  constructor(private userService: UserService) {}

  ngOnInit(): void {}

  saveChanges() {
    this.review.ClientUsername = (<HTMLInputElement>(
      document.getElementById('ClientUsername')
    )).value;

    this.review.RepairmanUsername = (<HTMLInputElement>(
      document.getElementById('RepairmanUsername')
    )).value;

    this.review.Text = (<HTMLInputElement>(
      document.getElementById('Text')
    )).value;

    this.review.ResponseId = '';
    this.review.Report = false;

    this.userService.createReview(this.review);
    window.location.reload();
  }
}
