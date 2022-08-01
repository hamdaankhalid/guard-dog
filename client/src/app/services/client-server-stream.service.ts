import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ClientServerStreamService {

  constructor() { }

  sendUpstream(data: string) {
    console.log(data);
  }
}
