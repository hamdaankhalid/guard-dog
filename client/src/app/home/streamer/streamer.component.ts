import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { WebcamImage, WebcamInitError } from 'ngx-webcam';
import { ClientServerStreamService } from 'src/app/services/client-server-stream.service';

declare var MediaRecorder: any;


@Component({
  selector: 'app-streamer',
  templateUrl: './streamer.component.html',
  styleUrls: ['./streamer.component.scss'],
})
export class StreamerComponent implements OnInit {
  // toggle webcam on/off
  public showWebcam = false;
  public errors: WebcamInitError[] = [];
  public timer: Date = new Date();
  
  // private chunks: any[] = [];

  private stream: MediaStream | null = null;

  private mediaRecorder: any = null;

  @ViewChild('myVideo', { read: ElementRef })
  myVideo!: ElementRef<HTMLVideoElement>;

  // @ViewChild('recordedVideo', { read: ElementRef })
  // recordedVideo!: ElementRef<HTMLVideoElement>;

  constructor(private clientServerStreamService: ClientServerStreamService) {}

  ngOnInit(): void {
    setInterval(() => {
      this.timer = new Date();
    }, 1000);
  }

  public toggleWebcam(): void {
    this.showWebcam = !this.showWebcam;
    if (this.showWebcam) {
      console.log('Record');

      navigator.mediaDevices.getUserMedia({ audio: true, video: true }).then((stream) => {
        
        this.stream = stream;
        this.myVideo.nativeElement.srcObject = stream;

        this.mediaRecorder = new MediaRecorder(stream) as any;
        
        this.mediaRecorder.ondataavailable = async (e: any) => {
          const blobText = await (new Blob([e.data])).text();
          this.clientServerStreamService.sendUpstream(blobText);
          // this.chunks.push(e.data);
        };

        this.mediaRecorder.start(1000); 

      }).catch((err) => {
        this.errors.push(err);
      });

    } else {
      console.log('Stop');

      this.mediaRecorder.stop();

      if (this.stream) {
        this.stream.getTracks().forEach(track => track.stop());
      }

      this.myVideo.nativeElement.srcObject = null;

      // this.recordedVideo.nativeElement.src = window.URL.createObjectURL(new Blob(this.chunks));
    }
  }
}
