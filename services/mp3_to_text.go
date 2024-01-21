package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

// ConvertMP3ToText converts an MP3 file to text using Google Cloud Speech-to-Text API.
// The text result is saved to a file, and the file path is returned.
func ConvertMP3ToText(filePath string) (string, error) {
	ctx := context.Background()

	// Read the MP3 file content
	audioContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Initialize the Speech client
	client, err := speech.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Configure the Speech-to-Text request
	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 16000, // Adjust based on your audio file's sample rate
			LanguageCode:    "ja-JP", // Specify Japanese language code
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{
				Content: audioContent,
			},
		},
	}

	// Perform Speech-to-Text recognition
	resp, err := client.Recognize(ctx, req)
	if err != nil {
		return "", err
	}

	// Extract text results
	var textResult string
	for _, result := range resp.Results {
		for _, alternative := range result.Alternatives {
			textResult += alternative.Transcript + "\n"
		}
	}

	// Save the text result to a file
	outputFilePath := filepath.Join("outputs", "text_result.txt")
	if err := ioutil.WriteFile(outputFilePath, []byte(textResult), 0644); err != nil {
		return "", err
	}

	return outputFilePath, nil
}
