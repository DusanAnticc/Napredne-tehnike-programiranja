import { Token } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { Login } from '../../model/Login';
import { AuthService } from '../../services/auth.service';
import { JwtHelperService } from '@auth0/angular-jwt';
import { AdminPageComponent } from 'src/modules/admin/pages/admin-page/admin-page.component';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
})
export class LoginComponent {
  constructor(
    private authService: AuthService,
    private router: Router,
    private toastrService: ToastrService
  ) {}

  submitBtn() {
    const auth: Login = {
      Username: (<HTMLInputElement>document.getElementById('username')).value,
      Password: (<HTMLInputElement>document.getElementById('password')).value,
    };
    this.authService.login(auth).subscribe({
      next: (result: any) => {
        localStorage.setItem('userToken', JSON.stringify(result.Token));
        const tokenString = localStorage.getItem('userToken');
        if (tokenString) {
          const jwt: JwtHelperService = new JwtHelperService();
          const info = jwt.decodeToken(tokenString);
          const role = info.role;
          localStorage.setItem('currentUserRole', role);
          localStorage.setItem('currentEmail', info.email);
          localStorage.setItem('currentId', info.id);

          if (role === 'Admin') this.router.navigate(['admin/createRepairman']);
          if (role === 'Standard') this.router.navigate(['user/hire']);
          if (role === 'Repairman')
            this.router.navigate(['repairman/acceptAppointment']);
          if (role === 'Employee') this.router.navigate(['employee/payment']);
        }
      },
      error: (error) => {
        if (error.status === 404)
          this.toastrService.error('Invalid username or password!');
        if (error.status === 400) this.toastrService.error('You are banned!');
      },
    });
  }
}
