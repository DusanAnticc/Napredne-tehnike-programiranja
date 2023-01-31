import { Component, OnInit } from '@angular/core';
import { Assessment } from '../../model/Assessment';
import { UserService } from '../../services/user.service';

@Component({
  selector: 'app-assessment',
  templateUrl: './assessment.component.html',
  styleUrls: ['./assessment.component.scss'],
})
export class AssessmentComponent implements OnInit {
  user = <Assessment>{};
  constructor(private userService: UserService) {}

  ngOnInit(): void {}

  saveChanges() {
    this.user.Username = (<HTMLInputElement>(
      document.getElementById('username')
    )).value;
    this.user.Grade = Number(
      (<HTMLInputElement>document.getElementById('grade')).value
    );
    if (this.user.Grade < 0 || this.user.Grade > 5) {
      alert('Grade must be between 0(min) and 5(max)');
    } else {
      this.userService.assessment(this.user);
      window.location.reload;
    }
  }
}
