import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { WebcamInitError } from 'ngx-webcam';

declare var MediaRecorder: any;

@Component({
  selector: 'app-streamer',
  templateUrl: './streamer.component.html',
  styleUrls: ['./streamer.component.scss'],
})
export class StreamerComponent implements OnInit {
  private readonly ONE_MINUTE_BY_RATE_OF_RECORDING = 30;
  private readonly RATE_OF_RECORDING = 2000;

  // toggle webcam on/off
  public showWebcam = false;
  public errors: WebcamInitError[] = [];
  public timer: Date = new Date();

  private rounds = 0;
  private chunks: any[] = [];
  private stream: MediaStream | null = null;
  private mediaRecorder: any = null;

  @ViewChild('myVideo', { read: ElementRef })
  myVideo!: ElementRef<HTMLVideoElement>;

  // @ViewChild('recordedVideo', { read: ElementRef })
  // recordedVideo!: ElementRef<HTMLVideoElement>;

  constructor() {}

  ngOnInit(): void {
    // shows the datetime on recorder screen
    setInterval(() => {
      this.timer = new Date();
    }, 1000);
  }

  public toggleWebcam(): void {
    this.showWebcam = !this.showWebcam;
    if (this.showWebcam) {
      navigator.mediaDevices
        .getUserMedia({ audio: true, video: true })
        .then((stream) => {
          this.stream = stream;
          this.myVideo.nativeElement.srcObject = stream;
          this.mediaRecorder = new MediaRecorder(stream) as any;
          this.mediaRecorder.ondataavailable = async (e: any) => {
            // const blobText = await (new Blob([e.data])).text();
            // this.clientServerStreamService.sendUpstream(blobText);
            this.chunks.push(e.data);
            if (this.chunks.length === this.ONE_MINUTE_BY_RATE_OF_RECORDING) {
              // transform and send remaining data to server
              this.transformAndUpload(this.chunks, `video_${this.rounds}`);
              // refresh chunks to be empty
              this.chunks = [];
              this.rounds++;
            }
          };
          this.mediaRecorder.start(this.RATE_OF_RECORDING);
        })
        .catch((err) => {
          this.errors.push(err);
        });
    } else {
      this.mediaRecorder.stop();
      if (this.stream) {
        this.stream.getTracks().forEach((track) => track.stop());
      }
      this.myVideo.nativeElement.srcObject = null;
      if (this.chunks.length > 0) {
        // transform and send remaining data to server
        this.transformAndUpload(this.chunks, "video_end");
      }
      this.chunks = [];
      // this.recordedVideo.nativeElement.src = window.URL.createObjectURL(new Blob(this.chunks));
    }
  }

  private transformAndUpload(chunks: any[], filename: string) {
    const file = new File([new Blob(chunks)], filename);
    console.log(file);
  }
}
