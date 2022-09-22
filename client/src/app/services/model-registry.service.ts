import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { RegisteredModel } from '../models/registered-model';

@Injectable({
  providedIn: 'root'
})
export class ModelRegistryService {
  private readonly MODEL_API_URL = "http://localhost:8080/api/video_storage/model";

  constructor(private httpClient: HttpClient) { }

  registerModel(file: File) {
    const formData = new FormData();
    formData.append("file", file);
    const options = {headers: new HttpHeaders().set("Authorization", `Bearer ${localStorage.getItem('access_token')}`)}

    return this.httpClient.post(this.MODEL_API_URL, formData, options).toPromise();
  }

  getModels() {
    const options = {headers: new HttpHeaders().set("Authorization", `Bearer ${localStorage.getItem('access_token')}`)};
    return this.httpClient.get<RegisteredModel[]>(this.MODEL_API_URL, options);
  }

  deleteModel(id: number) {
    const options = {headers: new HttpHeaders().set("Authorization", `Bearer ${localStorage.getItem('access_token')}`)};
    return this.httpClient.delete(`${this.MODEL_API_URL}/${id}`, options);
  }
}
