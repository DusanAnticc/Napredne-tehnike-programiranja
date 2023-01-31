import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InspectReviewRepairmanComponent } from './inspect-review-repairman.component';

describe('InspectReviewRepairmanComponent', () => {
  let component: InspectReviewRepairmanComponent;
  let fixture: ComponentFixture<InspectReviewRepairmanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InspectReviewRepairmanComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InspectReviewRepairmanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
