import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { ToastrModule } from 'ngx-toastr';

import { LoginComponent } from './components/login/login.component';
import { BaseLayoutComponent } from './pages/base-layout/base-layout.component';
import { AuthService } from './services/auth.service';
import { NavbarAdminComponent } from './components/navbar-admin/navbar-admin/navbar-admin.component';
import { AdminModule } from '../admin/admin.module';
import { UserModule } from '../user/user.module';
import { RepairmanModule } from '../repairman/repairman.module';
import { NavbarUserComponent } from './components/navbar-user/navbar-user/navbar-user.component';
import { NavbarRepairmanComponent } from './components/navbar-repairman/navbar-repairman/navbar-repairman.component';
import { EmployeeModule } from '../employee/employee.module';
import { NavbarEmployeeComponent } from './components/navbar-employee/navbar-employee/navbar-employee.component';
import { RegisterComponent } from './components/register/register.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    RegisterComponent,
    BaseLayoutComponent,
    NavbarAdminComponent,
    NavbarUserComponent,
    NavbarRepairmanComponent,
    NavbarEmployeeComponent,
  ],
  imports: [
    AdminModule,
    UserModule,
    RepairmanModule,
    BrowserModule,
    EmployeeModule,
    AppRoutingModule,
    HttpClientModule,
    ToastrModule.forRoot({
      positionClass: 'toast-top-right',
    }),
  ],
  providers: [AuthService],
  bootstrap: [AppComponent],
})
export class AppModule {}
