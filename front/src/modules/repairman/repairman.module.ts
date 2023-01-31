import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { RepairmanRoutes } from './repairman.routes';
import { RepairmanPageComponent } from './pages/repairman-page/repairman-page.component';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
import { AssessmentPageComponent } from './pages/assessment-page/assessment-page.component';
import { AssessmentComponent } from '../user/components/assessment/assessment.component';
import { AcceptAppointmentComponent } from './components/accept-appointment/accept-appointment.component';

@NgModule({
  declarations: [
    RepairmanPageComponent,
    AssessmentPageComponent,
    AcceptAppointmentComponent,
  ],
  imports: [
    CommonModule,
    FormsModule,
    HttpClientModule,
    Ng2SearchPipeModule,
    RouterModule.forChild(RepairmanRoutes),
  ],
})
export class RepairmanModule {}
