import { Routes } from '@angular/router';
import { ProfileComponent } from '../app/components/profile/profile.component';
import { RoleGuard } from '../app/guards/role.guard';
import { AssessmentComponent } from '../user/components/assessment/assessment.component';
import { AcceptAppointmentComponent } from './components/accept-appointment/accept-appointment.component';
import { RepairmanPageComponent } from './pages/repairman-page/repairman-page.component';

export const RepairmanRoutes: Routes = [
  {
    path: 'repairmanPage',
    pathMatch: 'full',
    component: RepairmanPageComponent,
  },
  {
    path: 'acceptAppointment',
    pathMatch: 'full',
    component: AcceptAppointmentComponent,
  },
  {
    path: 'assessment',
    pathMatch: 'full',
    component: AssessmentComponent,
  },
  {
    path: 'profile',
    pathMatch: 'full',
    component: ProfileComponent,
  },
];
