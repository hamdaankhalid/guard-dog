import { Component, OnInit } from '@angular/core';
import { InferenceNotification } from 'src/app/models/inference-notification';
import { InferenceNotificationService } from 'src/app/services/inference-notification.service';

@Component({
  selector: 'app-notification-center',
  templateUrl: './notification-center.component.html',
  styleUrls: ['./notification-center.component.scss']
})
export class NotificationCenterComponent implements OnInit {

  inferenceNotifications: InferenceNotification[] = [];
  queryInterval: any;

  constructor(private inferenceNotificationService: InferenceNotificationService) { }

  ngOnInit(): void {
    this.getNotifications();
    this.queryInterval = setInterval(() => this.getNotifications(), 10_000);
  }

  ngOnDestroy(): void {
    clearInterval(this.queryInterval);
  }

  getNotifications() {
    this.inferenceNotificationService.getAll().subscribe((inferences: InferenceNotification[]) => {
      this.inferenceNotifications = inferences;
    });
  }
}
