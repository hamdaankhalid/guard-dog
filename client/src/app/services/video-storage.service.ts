import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Metadata } from '../models/metadata';
import { Session } from '../models/session';

@Injectable({
  providedIn: 'root'
})
export class VideoStorageService {

  private readonly VIDEO_STORAGE_API_URL = "http://localhost:8080/api/video_storage";

  constructor(private httpClient: HttpClient) { }

  uploadFile(file: File, metadata: Metadata) {
    const formData = new FormData();
    formData.append("file", file);
    formData.append("metadata", JSON.stringify(metadata));
    const options = {headers: new HttpHeaders().set("Authorization", `Bearer ${localStorage.getItem('access_token')}`)}

    return this.httpClient.post(this.VIDEO_STORAGE_API_URL+"/miniupload", formData, options).toPromise();
  }

  // returns a list of all sessions, with optional parameter of deviceName
  getListOfSessions(deviceName: string | null = null) {
    const headers = new HttpHeaders().set("Authorization", `Bearer ${localStorage.getItem('access_token')}`);

    const options = deviceName ? { params: { "device": deviceName } } : undefined;
    return this.httpClient.get<Session[]>(`${this.VIDEO_STORAGE_API_URL}/sessions`, {...options, headers});
  }

  // returns a single session and all the child video id's and metadata
  getSession(sessionId: string) {
    const options = {headers: new HttpHeaders().set("Authorization", `Bearer ${localStorage.getItem('access_token')}`)}

    return this.httpClient.get(`${this.VIDEO_STORAGE_API_URL}/sessions/${sessionId}`, options);
  }
}
