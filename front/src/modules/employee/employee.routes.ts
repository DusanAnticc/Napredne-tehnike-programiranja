import { Routes } from '@angular/router';
import { ProfileComponent } from '../app/components/profile/profile.component';
import { RoleGuard } from '../app/guards/role.guard';
import { PaymentComponent } from './components/payment/payment.component';
import { EmployeePageComponent } from './pages/employee-page/employee-page.component';

export const EmployeeRoutes: Routes = [
  {
    path: 'employeePage',
    pathMatch: 'full',
    component: EmployeePageComponent,
  },
  {
    path: 'payment',
    pathMatch: 'full',
    component: PaymentComponent,
  },
  {
    path: 'profile',
    pathMatch: 'full',
    component: ProfileComponent,
  },
];
