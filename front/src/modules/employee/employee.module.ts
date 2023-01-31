import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { EmployeeRoutes } from './employee.routes';
import { EmployeePageComponent } from './pages/employee-page/employee-page.component';
import { FormsModule } from '@angular/forms';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
import { PaymentComponent } from './components/payment/payment.component';

@NgModule({
  declarations: [EmployeePageComponent, PaymentComponent],
  imports: [
    CommonModule,
    FormsModule,

    Ng2SearchPipeModule,
    RouterModule.forChild(EmployeeRoutes),
  ],
})
export class EmployeeModule {}
