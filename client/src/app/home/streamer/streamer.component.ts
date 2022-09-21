import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { Metadata } from 'src/app/models/metadata';
import { User } from 'src/app/models/user';
import { AuthService } from 'src/app/services/auth.service';
import { VideoStorageService } from 'src/app/services/video-storage.service';

declare var MediaRecorder: any;

const NO_DEVICE_NAME_ERROR = 'Enter a device name to begin recording!';

@Component({
  selector: 'app-streamer',
  templateUrl: './streamer.component.html',
})
export class StreamerComponent implements OnInit {

  // toggle webcam on/off
  public showWebcam = false;
  public errors: any[] = [];
  public timer: Date = new Date();
  public isNameDisabled: boolean = false;

  private rounds = 0;
  private stream: MediaStream | null = null;
  private mediaRecorderInterval: any;

  private sessionStart: Date | undefined | null;

  @ViewChild('myVideo', { read: ElementRef })
  myVideo!: ElementRef<HTMLVideoElement>;

  @ViewChild('deviceName', { read: ElementRef })
  deviceName!: ElementRef<HTMLInputElement>;

  activeUser: User | null = null;

  constructor(private videoStorageService: VideoStorageService, private authenticationService: AuthService) {
    this.authenticationService.currentUser.subscribe((res: User | null) => this.activeUser = res);
  }

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
              await this.transformAndUpload(
                this.rounds, 
                chunks, 
                `${deviceName}_${this.sessionStart?.toUTCString()}_video_${this.rounds}`, 
                this.sessionStart!
              );
              this.rounds++;
            }
            setTimeout(() => {
              if (mediaRecorder.state !== "inactive") {
                mediaRecorder.stop()
              }
            }, 5_000);
            mediaRecorder.start();
          }, 5_000);
 
        })
        .catch((err) => {
          this.errors.push(err);
        });
      this.sessionStart = new Date();
    } else {
      this.isNameDisabled = false;
      
      clearInterval(this.mediaRecorderInterval);
      this.myVideo.nativeElement.srcObject = null;

      if (this.stream) {
        this.stream.getTracks().forEach((track) => track.stop());
      }

      // this is so the clearInterval has time to run the last one without a messed up rounds param
      setTimeout(() => {
        this.rounds = 0;
        this.sessionStart = null;
      }, 2_000);
    }
  }



  private async transformAndUpload(
    round: number,
    chunks: any[],
    filename: string,
    sessionStart: Date
  ) {
    const blob = new File(chunks, `${filename}.webm`, { type: 'video/webm' });

    const metadata: Metadata = {
      name: `${filename}.webm`,
      part: round,
      deviceName: this.deviceName.nativeElement.value,
      durationInSeconds: chunks.length * 2,
      sessionStart: sessionStart!,
      userId: this.activeUser!.id
    };

    await this.videoStorageService.uploadFile(blob, metadata);
  }
}
