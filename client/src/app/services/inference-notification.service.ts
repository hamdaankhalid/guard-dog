import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { InferenceNotification } from '../models/inference-notification';

@Injectable({
  providedIn: 'root'
})
export class InferenceNotificationService {
  private readonly INFERENCE_NOTIFICATION_API_URL = "http://localhost:8080/api/video_storage/inferences";

  constructor(private httpClient: HttpClient) { }

  getAll() {
    const options = {
      headers: new HttpHeaders().set("Authorization", `Bearer ${localStorage.getItem('access_token')}`)
    };
    return this.httpClient.get<InferenceNotification[]>(this.INFERENCE_NOTIFICATION_API_URL, options);
  }
}
