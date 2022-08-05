import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Metadata } from '../models/metadata';

@Injectable({
  providedIn: 'root'
})
export class VideoStorageService {

  private readonly VIDEO_STORAGE_API_URL = "http://localhost:8080/api/video_storage";

  constructor(private httpClient: HttpClient) { }

  uploadFile(file: File, metadata: Metadata) {
    const formData = new FormData();
    formData.append("base64file", file);
    formData.append("metadata", JSON.stringify(metadata));
    // tracer
    //this.httpClient.get(this.VIDEO_STORAGE_API_URL+"/health").subscribe(console.log, console.error);

    return this.httpClient.post(this.VIDEO_STORAGE_API_URL+"/miniupload", formData).toPromise();
  }
}
