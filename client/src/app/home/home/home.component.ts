import { Component, OnInit } from '@angular/core';
import { Session } from '../../models/session';
import { VideoStorageService } from 'src/app/services/video-storage.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
})
export class HomeComponent implements OnInit {
  sessions: Session[] = [];

  constructor(private videoStorageService: VideoStorageService) {}

  ngOnInit(): void {
    this.videoStorageService.getListOfSessions().subscribe((sessions: Session[]) => {
      this.sessions = sessions;
    }, alert);
  }
}
