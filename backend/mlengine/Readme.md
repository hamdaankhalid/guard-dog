# ML Execution Service

## API Routes for Client (all require jwt token)
- Health ```GET /health```
- Upload ONNX model for a user ``` POST /model``` 
- Delete an ONNX model for a user ```DELETE /model```
- Get ML notifications    ```GET /inferences```

## Pub-Sub
- Consumes: Miniupload Created Event
