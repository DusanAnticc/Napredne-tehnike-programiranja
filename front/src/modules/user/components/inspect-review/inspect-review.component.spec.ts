import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InspectReviewComponent } from './inspect-review.component';

describe('InspectReviewComponent', () => {
  let component: InspectReviewComponent;
  let fixture: ComponentFixture<InspectReviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InspectReviewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InspectReviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
