import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateRepairmanComponent } from './create-repairman.component';

describe('CreateRepairmanComponent', () => {
  let component: CreateRepairmanComponent;
  let fixture: ComponentFixture<CreateRepairmanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateRepairmanComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateRepairmanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
