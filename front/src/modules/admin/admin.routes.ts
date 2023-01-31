import { Routes } from '@angular/router';
import { ProfileComponent } from '../app/components/profile/profile.component';
import { RoleGuard } from '../app/guards/role.guard';
import { BanComponent } from './components/ban/ban.component';
import { CreateRepairmanComponent } from './components/create-repairman/create-repairman.component';
import { ReportsComponent } from './components/reports/reports.component';
import { AdminPageComponent } from './pages/admin-page/admin-page.component';

export const AdminRoutes: Routes = [
  {
    path: 'adminPage',
    pathMatch: 'full',
    component: AdminPageComponent,
  },
  {
    path: 'profile',
    pathMatch: 'full',
    component: ProfileComponent,
  },
  {
    path: 'reports',
    pathMatch: 'full',
    component: ReportsComponent,
  },
  {
    path: 'createRepairman',
    pathMatch: 'full',
    component: CreateRepairmanComponent,
  },
  {
    path: 'ban',
    pathMatch: 'full',
    component: BanComponent,
  },
];
