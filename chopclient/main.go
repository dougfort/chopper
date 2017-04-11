package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	testCases := []struct {
		name           string
		url            string
		expectedStatus int
	}{
		{
			name:           "simple",
			url:            "aaa",
			expectedStatus: 200,
		},
	}

	log.Printf("info: test client starts")

TEST_LOOP:
	for i, testCase := range testCases {
		status, choppedURL, err := testChop(testCase.url)
		if err != nil {
			log.Printf("error: #%d: %s: failed: %s", i+1, testCase.name, err)
			continue TEST_LOOP
		}
		if status != testCase.expectedStatus {
			log.Printf("error: #%d: %s: expected status %d: found %d",
				i+1, testCase.name, testCase.expectedStatus, status)
			continue TEST_LOOP
		}
		log.Printf("debug: chopped url = '%s'", choppedURL)
	}

	log.Printf("info: test client completes")
}

func testChop(url string) (int, string, error) {
	// we marshal the url to escape bad characters
	marshalled, err := json.Marshal(map[string]string{"url": url})
	if err != nil {
		return 0, "", fmt.Errorf("json.Marshal '%s' failed: %s", url, err)
	}
	resp, err := http.Post(
		"http://127.0.0.1:9000/chop",
		"application/json",
		bytes.NewBuffer(marshalled),
	)
	if err != nil {
		return 0, "", fmt.Errorf("request failed: %s", err)
	}
	return resp.StatusCode, "", nil
}
