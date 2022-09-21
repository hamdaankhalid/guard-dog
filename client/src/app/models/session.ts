
type DurationUnit = | "HOURS" |  "MINUTES" | "SECONDS";

export interface VideoMetadata {
  id: number;
  part: number;
  duration: number;
  filename: string;
  url: string;
}

export interface Session {
  id: number;
  userId: number;
  deviceName: string;
  sessionStart: Date;
  duration: number;
  durationUnit: DurationUnit;
  videoMetadatas: VideoMetadata[];
}


/**
 * [
    {
        "id": 2,
        "userId": 1,
        "deviceName": "test1",
        "sessionStart": "2022-09-21T05:24:13.577+00:00",
        "duration": 2,
        "durationUnit": "SECONDS",
        "videoMetadatas": [
            {
                "id": 4,
                "part": 1,
                "duration": 2,
                "filename": "test1",
                "url": "https://guarddogvideostore.blob.core.windows.net/miniuploads/test1_Wed%2C%2021%20Sep%202022%2005%3A24%3A13%20GMT_video_1.webm"
            },
            {
                "id": 5,
                "part": 2,
                "duration": 2,
                "filename": "test1",
                "url": "https://guarddogvideostore.blob.core.windows.net/miniuploads/test1_Wed%2C%2021%20Sep%202022%2005%3A24%3A13%20GMT_video_2.webm"
            },
            {
                "id": 3,
                "part": 0,
                "duration": 2,
                "filename": "test1",
                "url": "https://guarddogvideostore.blob.core.windows.net/miniuploads/test1_Wed%2C%2021%20Sep%202022%2005%3A24%3A13%20GMT_video_0.webm"
            }
        ]
    }
]
 */