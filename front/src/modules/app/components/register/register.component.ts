import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { Register } from '../../model/User';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.scss'],
})
export class RegisterComponent implements OnInit {
  newUser = <Register>{};

  constructor(
    private authService: AuthService,
    private router: Router,
    private toastrService: ToastrService
  ) {}

  ngOnInit(): void {}

  registerBtn() {
    this.newUser.Username = (<HTMLInputElement>(
      document.getElementById('Username')
    )).value;
    this.newUser.Password = (<HTMLInputElement>(
      document.getElementById('Password')
    )).value;
    this.newUser.FirstName = (<HTMLInputElement>(
      document.getElementById('FirstName')
    )).value;
    this.newUser.LastName = (<HTMLInputElement>(
      document.getElementById('LastName')
    )).value;
    this.newUser.EmailAddress = (<HTMLInputElement>(
      document.getElementById('EmailAddress')
    )).value;
    this.newUser.Role = (<HTMLInputElement>(
      document.getElementById('Role')
    )).value;
    this.newUser.MoneyBalance = 0;
    this.newUser.Price = Number(
      (<HTMLInputElement>document.getElementById('Price')).value
    );
    this.newUser.Occupation = (<HTMLInputElement>(
      document.getElementById('Occupation')
    )).value;
    this.newUser.Address = (<HTMLInputElement>(
      document.getElementById('Address')
    )).value;
    this.newUser.Banned = false;
    this.newUser.Review = 0;
    this.newUser.ReviewCounter = 0;

    this.authService.register(this.newUser);
    window.location.href = 'http://localhost:4200/login';
  }
}
