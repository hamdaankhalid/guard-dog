
type DurationUnit = | "hours" | "minutes";

export interface Session {
  sessionId: string;
  duration: number;
  durationUnit: DurationUnit;
  timeStarted: Date;
  deviceName: string
}