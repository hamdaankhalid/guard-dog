# Video Storage Service

## Endpoints

All endpoints are authenticated only

- File upload endpoint
- Retrieve by datetime queries
- Get List of video sessions for a user sorted by date


## Upload Video Stream Workflow
- Frontend sends 1 minute long video files every 1 minute. This video is related to a user Id, session Id, and device Name. Each file is sent to blob storage and the URL for the blob storage is added to the video metadata table. 
- Frontend also retrieves a list of sessions, device names, durations
- Frontend can make a query to view an entire session at which point we can create a video player that pulls data based on the minute the video player is on.
