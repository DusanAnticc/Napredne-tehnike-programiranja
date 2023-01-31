import { TestBed } from '@angular/core/testing';

import { RepairmanService } from './repairman.service';

describe('RepairmanService', () => {
  let service: RepairmanService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RepairmanService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
