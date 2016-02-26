package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//Body of json request
type Body struct {
	Requests *[]Request
}

//Request top level json request structure
type Request struct {
	Image    *Image
	Features *[]Feature
}

//Image data encoded as base64
type Image struct {
	Content string `json:",omitempty"`
}

//Feature feature items to be identified
type Feature struct {
	Type       string `json:",omitempty"`
	MaxResults string `json:",omitempty"`
}

const (
	TypeUnspecified     = iota //Unspecified feature type.
	FaceDetection              //Run face detection.
	LandmarkDetection          //Run landmark detection.
	LogoDetection              //Run logo detection.
	LabelDetection             //Run label detection.
	TextDetection              //Run OCR.
	SafeSearchDetection        //Run various computer vision models to
	ImageDetection             //Compute a set of properties about the image (such as the image's dominant colors)
)

func testFeature() {
	fmt.Println()
}

//marshalJSON JSON message
func marshalJSON(image string) []byte {

	i := Image{image}

	f := Feature{"TEXT_DETECTION", "10"}
	r := Request{&i, &[]Feature{f, {"TEXT_DETECTION", "10"}}}
	bdy := Body{&[]Request{r}}

	b, err := json.MarshalIndent(bdy, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
	return b
}

/*
//marshalJSON JSON message
func marshalJSON(image string) []byte {
	return []byte(fmt.Sprintf(
		`{
      "requests": [
        {
          "image": {
            "content": "%s"
        },
        "features": [
          {
            "type": "TEXT_DETECTION",
            "maxResults": "100"
          }
        ]
      }
    ]
  }`, image))
}
*/

//encodedImage takes a file and returns a string encoded in base64
func encodeBase64(filePath string) (string, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("Open File: %v", err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("Read File: %v", err)
	}

	return base64.StdEncoding.EncodeToString(b), nil
}
