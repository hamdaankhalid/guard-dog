/**
    int id;
    ServiceUser serviceUser;
    VideoMetadata videoMetadata;
    String details;
 */

import { VideoMetadata } from "./session";
import { User } from "./user";

export interface InferenceNotification {
  id: number;
  user: User;
  videoMetadata: VideoMetadata;
  details: string;
}
