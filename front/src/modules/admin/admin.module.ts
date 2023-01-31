import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { AdminRoutes } from './admin.routes';
import { AdminPageComponent } from './pages/admin-page/admin-page.component';
import { FormsModule } from '@angular/forms';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
import { BanComponent } from './components/ban/ban.component';
import { CreateRepairmanComponent } from './components/create-repairman/create-repairman.component';
import { ReportsComponent } from './components/reports/reports.component';

@NgModule({
  declarations: [
    AdminPageComponent,
    BanComponent,
    CreateRepairmanComponent,
    ReportsComponent,
  ],
  imports: [
    CommonModule,
    FormsModule,
    Ng2SearchPipeModule,
    RouterModule.forChild(AdminRoutes),
  ],
})
export class AdminModule {}
