import { Routes } from '@angular/router';
import { ProfileComponent } from '../app/components/profile/profile.component';
import { RoleGuard } from '../app/guards/role.guard';
import { CreateAppointmentComponent } from '../repairman/components/create-appointment/create-appointment.component';
import { AssessmentComponent } from './components/assessment/assessment.component';
import { CreateReviewComponent } from './components/create-review/create-review.component';
import { HireRepairmanComponent } from './components/hire-repairman/hire-repairman.component';
import { InspectReviewComponent } from './components/inspect-review/inspect-review.component';
import { PayBillComponent } from './components/pay-bill/pay-bill.component';
import { HireRepairmanPageComponent } from './pages/hire-repairman-page/hire-repairman-page.component';
import { UserPageComponent } from './pages/user-page/user-page.component';

export const UserRoutes: Routes = [
  {
    path: 'hire',
    pathMatch: 'full',
    component: HireRepairmanPageComponent,
  },
  {
    path: 'profile',
    pathMatch: 'full',
    component: ProfileComponent,
  },
  // {
  //   path: 'hire',
  //   pathMatch: 'full',
  //   component: HireRepairmanComponent,
  // },
  // {
  //   path: 'createAppointment',
  //   pathMatch: 'full',
  //   component: CreateAppointmentComponent,
  // },
  {
    path: 'createReview',
    pathMatch: 'full',
    component: CreateReviewComponent,
  },
  {
    path: 'inspectReview',
    pathMatch: 'full',
    component: InspectReviewComponent,
  },
  {
    path: 'assessment',
    pathMatch: 'full',
    component: AssessmentComponent,
  },
  {
    path: 'payBill',
    pathMatch: 'full',
    component: PayBillComponent,
  },
];
