import { TestBed } from '@angular/core/testing';

import { InferenceNotificationService } from './inference-notification.service';

describe('InferenceNotificationService', () => {
  let service: InferenceNotificationService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(InferenceNotificationService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
