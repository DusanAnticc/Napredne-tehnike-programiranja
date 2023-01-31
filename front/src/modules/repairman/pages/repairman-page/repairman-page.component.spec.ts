import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RepairmanPageComponent } from './repairman-page.component';

describe('RepairmanPageComponent', () => {
  let component: RepairmanPageComponent;
  let fixture: ComponentFixture<RepairmanPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RepairmanPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RepairmanPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
