import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { Metadata } from 'src/app/models/metadata';
import { VideoStorageService } from 'src/app/services/video-storage.service';

declare var MediaRecorder: any;

const NO_DEVICE_NAME_ERROR = 'Enter a device name to begin recording!';

@Component({
  selector: 'app-streamer',
  templateUrl: './streamer.component.html',
  styleUrls: ['./streamer.component.scss'],
})
export class StreamerComponent implements OnInit {
  private readonly DURATION_WRT_RATE_OF_RECORDING = 7; // 1 minute
  private readonly RATE_OF_TRIGGER = 1000;

  // toggle webcam on/off
  public showWebcam = false;
  public errors: any[] = [];
  public timer: Date = new Date();
  public isNameDisabled: boolean = false;

  private rounds = 0;
  private stream: MediaStream | null = null;
  private mediaRecorder: any = null;
  private mediaRecorderInterval: any;

  private session: Date | undefined | null;

  @ViewChild('myVideo', { read: ElementRef })
  myVideo!: ElementRef<HTMLVideoElement>;

  @ViewChild('deviceName', { read: ElementRef })
  deviceName!: ElementRef<HTMLInputElement>;

  constructor(private videoStorageService: VideoStorageService) {}

  ngOnInit(): void {
    setInterval(() => (this.timer = new Date()), 1000);
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

          this.mediaRecorderInterval = setInterval(() => {
            const chunks: any[] = [];
            const mediaRecorder = new MediaRecorder(stream) as any;
            mediaRecorder.ondataavailable = async (e: any) => chunks.push(e.data);
            mediaRecorder.onstop = async (e: any) => {
              const deviceName = this.deviceName.nativeElement.value;
              await this.transformAndUpload(this.rounds, chunks, `${deviceName}_video_${this.rounds}`, this.session!);
              this.rounds++;
            }
            setTimeout(() => {
              if (mediaRecorder.state !== "inactive") {
                mediaRecorder.stop()
              }
            }, 60_000);
            mediaRecorder.start();
          }, 60_000);
 
        })
        .catch((err) => {
          this.errors.push(err);
        });
      this.session = new Date();
    } else {
      this.isNameDisabled = false;
      
      clearInterval(this.mediaRecorderInterval);
      this.myVideo.nativeElement.srcObject = null;

      if (this.stream) {
        this.stream.getTracks().forEach((track) => track.stop());
      }
      this.session = null;
      this.rounds = 0;
    }
  }



  private async transformAndUpload(
    round: number,
    chunks: any[],
    filename: string,
    session: Date
  ) {
    const blob = new File(chunks, `${filename}.webm`, { type: 'video/webm' });

    const metadata: Metadata = {
      name: `${filename}.webm`,
      part: round,
      deviceName: this.deviceName.nativeElement.value,
      durationInSeconds: chunks.length * 2,
      session: session!,
    };

    await this.videoStorageService.uploadFile(blob, metadata);
  }
}
