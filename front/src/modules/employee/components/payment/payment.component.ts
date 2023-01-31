import { Component, OnInit } from '@angular/core';
import { EmployeeService } from '../../services/employee.service';

@Component({
  selector: 'app-payment',
  templateUrl: './payment.component.html',
  styleUrls: ['./payment.component.scss'],
})
export class PaymentComponent implements OnInit {
  username: any = '';
  money: any = '';
  constructor(private employeeService: EmployeeService) {}

  ngOnInit(): void {}

  deposit() {
    this.username = (<HTMLInputElement>(
      document.getElementById('Username')
    )).value;
    this.money = (<HTMLInputElement>document.getElementById('Money')).value;
    this.employeeService.deposit(this.username, this.money);
    window.location.reload();
  }
}
