import { TestBed } from '@angular/core/testing';

import { ModelRegistryService } from './model-registry.service';

describe('ModelRegistryService', () => {
  let service: ModelRegistryService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ModelRegistryService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
