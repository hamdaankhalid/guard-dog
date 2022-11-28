# ML Execution Service

## Introduction
- This service handles Models, Model Execution and Inference Notifications from the uploaded models.

## API Routes for Client
### Public Routes:
- GET ```/health```

### Authenticated Routes
- POST  ```/model``` Upload a ML Model
- GET ```/model``` Retrieve all uploaded ML Models except the file data
- GET ```/model/{modelId}``` Retrieve an uploaded ML Model with data
- DELETE ```/model/{modelId}``` Delete an uploaded ML Model
- GET ```/ml-notification``` Get a list of all ml-notifications corresponding to the users id
- GET ```/ml-notification/{mlNotificationId}``` Retrieve an Ml notification

## Kafka Pub-Sub
- Consumes: Topic ```video-upload```

## TODO:
- How to loop in handling training? This feature makes the service into an engine!
