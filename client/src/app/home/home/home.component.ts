import { Component, OnInit } from '@angular/core';
import { Session } from '../../models/session';
import { VideoStorageService } from 'src/app/services/video-storage.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
})
export class HomeComponent implements OnInit {
  sessions: Session[] = [];
  showSession: Record<number, boolean> = {};
  
  private queryInterval: any;

  constructor(private videoStorageService: VideoStorageService) {}

  ngOnInit(): void {
    this.querySessions();
    this.queryInterval = setInterval(() => {
        this.querySessions();
    }, 10_000);
  }

  ngOnDestroy(): void {
    clearInterval(this.queryInterval);
  }

  querySessions() {
    this.videoStorageService.getListOfSessions().subscribe((sessions: Session[]) => {
      this.sessions = sessions;

      sessions.forEach((ses: Session) => {
        const showVal = ses.id in this.showSession ? this.showSession[ses.id] : false;
        this.showSession[ses.id] =  showVal;
      });

    });
  }

  viewRecordings(sessionid: number) {
    this.showSession[sessionid] = !this.showSession[sessionid];
    if (this.showSession[sessionid]) {
      this.sessions[sessionid].videoMetadatas.sort((a, b) => b.part - a.part);
    }
  }
}
