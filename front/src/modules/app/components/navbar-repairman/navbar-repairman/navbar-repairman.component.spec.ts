import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NavbarRepairmanComponent } from './navbar-repairman.component';

describe('NavbarRepairmanComponent', () => {
  let component: NavbarRepairmanComponent;
  let fixture: ComponentFixture<NavbarRepairmanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NavbarRepairmanComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NavbarRepairmanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
