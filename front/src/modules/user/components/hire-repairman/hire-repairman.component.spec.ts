import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HireRepairmanComponent } from './hire-repairman.component';

describe('HireRepairmanComponent', () => {
  let component: HireRepairmanComponent;
  let fixture: ComponentFixture<HireRepairmanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HireRepairmanComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HireRepairmanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
