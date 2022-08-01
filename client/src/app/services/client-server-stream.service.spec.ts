import { TestBed } from '@angular/core/testing';

import { ClientServerStreamService } from './client-server-stream.service';

describe('ClientServerStreamService', () => {
  let service: ClientServerStreamService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ClientServerStreamService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
