import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { WebSocketSubject } from 'rxjs/webSocket';

@Injectable({
  providedIn: 'root'
})
export class ClientServerStreamService {

  connection$!: WebSocketSubject<any>;
  RETRY_SECONDS = 10; 

  constructor() { }

  
  connect(): Observable<any> {
}

  sendUpstream(data: string) {
    if (this.connection$) {
      this.connection$.next(data);
    } else {
      console.error('Did not send data, open a connection first');
    }
  }
}
