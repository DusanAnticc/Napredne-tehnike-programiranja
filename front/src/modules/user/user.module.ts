import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { UserRoutes } from './user.routes';
import { UserPageComponent } from './pages/user-page/user-page.component';
import { FormsModule } from '@angular/forms';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
import { HireRepairmanComponent } from './components/hire-repairman/hire-repairman.component';
import { CreateAppointmentComponent } from '../repairman/components/create-appointment/create-appointment.component';
import { AssessmentComponent } from './components/assessment/assessment.component';
import { ToastrModule } from 'ngx-toastr';
import { CreateReviewComponent } from './components/create-review/create-review.component';
import { InspectReviewComponent } from './components/inspect-review/inspect-review.component';
import { PayBillComponent } from './components/pay-bill/pay-bill.component';
import { HireRepairmanPageComponent } from './pages/hire-repairman-page/hire-repairman-page.component';

@NgModule({
  declarations: [
    UserPageComponent,
    PayBillComponent,
    HireRepairmanComponent,
    CreateAppointmentComponent,
    AssessmentComponent,
    CreateReviewComponent,
    InspectReviewComponent,
    HireRepairmanPageComponent,
  ],
  imports: [
    CommonModule,
    FormsModule,
    Ng2SearchPipeModule,
    RouterModule.forChild(UserRoutes),
    ToastrModule,
  ],
})
export class UserModule {}
