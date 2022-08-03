import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { Metadata } from 'src/app/models/metadata';
import { VideoStorageService } from 'src/app/services/video-storage.service';

declare var MediaRecorder: any;

const NO_DEVICE_NAME_ERROR = "Enter a device name to begin recording!";

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
  public errors: any[] = [];
  public timer: Date = new Date();
  public isNameDisabled: boolean = false;

  private rounds = 0;
  private chunks: any[] = [];
  private stream: MediaStream | null = null;
  private mediaRecorder: any = null;

  private session: Date | undefined | null;;

  @ViewChild('myVideo', { read: ElementRef })
  myVideo!: ElementRef<HTMLVideoElement>;

  @ViewChild('deviceName', { read: ElementRef })
  deviceName!: ElementRef<HTMLInputElement>;

  // @ViewChild('recordedVideo', { read: ElementRef })
  // recordedVideo!: ElementRef<HTMLVideoElement>;

  constructor(private videoStorageService: VideoStorageService) {}

  ngOnInit(): void {
    setInterval(() => this.timer = new Date(), 1000);
  }

  public toggleWebcam(): void {
    this.showWebcam = !this.showWebcam;
    if (this.showWebcam) {
      if (!this.deviceName.nativeElement.value || this.deviceName.nativeElement.value.includes(" ")) {
        if (!this.errors.includes(NO_DEVICE_NAME_ERROR)) {
          this.errors.push(NO_DEVICE_NAME_ERROR);
        }
        this.showWebcam = !this.showWebcam;
        return;
      } 
      const idx = this.errors.findIndex((val) => val === NO_DEVICE_NAME_ERROR);
      if (idx !== -1) {
        this.errors.splice(idx);
      }
      this.isNameDisabled = true;
      navigator.mediaDevices
        .getUserMedia({ audio: true, video: true })
        .then((stream) => {
          this.stream = stream;
          this.myVideo.nativeElement.srcObject = stream;
          this.mediaRecorder = new MediaRecorder(stream) as any;
          this.mediaRecorder.ondataavailable = async (e: any) => {
            this.chunks.push(e.data);
            if (this.chunks.length === this.ONE_MINUTE_BY_RATE_OF_RECORDING) {
              this.transformAndUpload(this.rounds, this.chunks, `${this.deviceName.nativeElement.value}_video_${this.rounds}`, this.session!);
              this.chunks = [];
              this.rounds++;
            }
          };
          // the rate of recording decides the time we have to upload the data before the next batch comes in
          this.mediaRecorder.start(this.RATE_OF_RECORDING);
        })
        .catch((err) => {
          this.errors.push(err);
        });
      this.session = new Date();
    } else {
      this.isNameDisabled = false;
      this.mediaRecorder.stop();
      if (this.stream) {
        this.stream.getTracks().forEach((track) => track.stop());
      }
      this.myVideo.nativeElement.srcObject = null;
      if (this.chunks.length > 0) {
        this.transformAndUpload(-1, this.chunks, `${this.deviceName.nativeElement.value}_video_video_end`, this.session!);
      }
      this.chunks = [];
      // this.recordedVideo.nativeElement.src = window.URL.createObjectURL(new Blob(this.chunks));
      this.session = null;
      this.rounds = 0;
    }
  }

  private transformAndUpload(round: number, chunks: any[], filename: string, session: Date) {
    const file = new File([new Blob(chunks)], filename);
    const metadata: Metadata = {
      name: filename,
      part: round,
      deviceName: this.deviceName.nativeElement.value,
      durationInSeconds: chunks.length * 2,
      session: this.session!
    };

    this.videoStorageService.uploadFile(file, metadata).subscribe(console.log, console.error);
  }
}
