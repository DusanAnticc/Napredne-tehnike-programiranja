import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HireRepairmanPageComponent } from './hire-repairman-page.component';

describe('HireRepairmanPageComponent', () => {
  let component: HireRepairmanPageComponent;
  let fixture: ComponentFixture<HireRepairmanPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HireRepairmanPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HireRepairmanPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
