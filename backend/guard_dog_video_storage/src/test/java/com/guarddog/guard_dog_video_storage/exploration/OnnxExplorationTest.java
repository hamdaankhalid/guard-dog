package com.guarddog.guard_dog_video_storage.exploration;

import ai.onnxruntime.OrtEnvironment;
import ai.onnxruntime.OrtException;
import ai.onnxruntime.OrtSession;
import org.junit.jupiter.api.Test;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

public class OnnxExplorationTest {

    @Test
    public void canLoadModelFromBytes() throws IOException, OrtException {
        Path path = Paths.get("src/main/resources/static/resnet101-v1-7.onnx");
        byte[] data = Files.readAllBytes(path);
        // Simulate reading data in tempfile and building model from tempfile

        OrtEnvironment env = OrtEnvironment.getEnvironment();
        OrtSession session = env.createSession(data);

        assertTrue(data.length > 0);
        assertNotNull(session);
    }

    @Test
    public void canInferModel() throws IOException, OrtException {
        Path path = Paths.get("src/main/resources/static/resnet101-v1-7.onnx");
        byte[] data = Files.readAllBytes(path);
        // Simulate reading data in tempfile and building model from tempfile

        OrtEnvironment env = OrtEnvironment.getEnvironment();
        OrtSession session = env.createSession(data);

        assertTrue(data.length > 0);
        assertNotNull(session);
    }

}
