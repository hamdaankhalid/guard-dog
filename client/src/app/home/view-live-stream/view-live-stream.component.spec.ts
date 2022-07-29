import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewLiveStreamComponent } from './view-live-stream.component';

describe('ViewLiveStreamComponent', () => {
  let component: ViewLiveStreamComponent;
  let fixture: ComponentFixture<ViewLiveStreamComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewLiveStreamComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewLiveStreamComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
