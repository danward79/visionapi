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
	Requests *[]Request `json:"requests,omitempty"`
}

//Request top level json request structure
type Request struct {
	Image    *Image     `json:"image,omitempty"`
	Features *[]Feature `json:"features,omitempty"`
}

//Image data encoded as base64
type Image struct {
	Content string `json:"content,omitempty"`
}

//Feature feature items to be identified
type Feature struct {
	Type       string `json:"type,omitempty"`
	MaxResults string `json:"maxresults,omitempty"`
}

const (
	//TypeUnspecified Unspecified feature type.
	TypeUnspecified = iota
	//FaceDetection Run face detection.
	FaceDetection
	//LandmarkDetection Run landmark detection.
	LandmarkDetection
	//LogoDetection Run logo detection.
	LogoDetection
	//LabelDetection Run label detection.
	LabelDetection
	//TextDetection Run OCR.
	TextDetection
	//SafeSearchDetection Run various computer vision models to
	SafeSearchDetection
	//ImageDetection Compute a set of properties about the image (such as the image's dominant colors)
	ImageDetection
)

//marshalJSON JSON message
func marshalJSON(image string, requiredFeatures *[]Feature) []byte {

	i := Image{image}
	r := Request{&i, requiredFeatures}
	bdy := Body{&[]Request{r}}

	b, err := json.MarshalIndent(bdy, "", "  ")
	//b, err := json.Marshal(bdy)
	if err != nil {
		fmt.Println(err)
	}

	return b
}

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

//buildFeatures
func buildFeatures(f features) []Feature {
	features := []Feature{}

	if f.FaceDetection {
		features = append(features, Feature{Type: "FACE_DETECTION"})
	}
	if f.LandmarkDetection {
		features = append(features, Feature{Type: "LANDMARK_DETECTION"})
	}
	if f.LogoDetection {
		features = append(features, Feature{Type: "LOGO_DETECTION"})
	}
	if f.LabelDetection {
		features = append(features, Feature{Type: "LABEL_DETECTION"})
	}
	if f.TextDetection {
		features = append(features, Feature{Type: "TEXT_DETECTION"})
	}
	if f.SafeSearchDetection {
		features = append(features, Feature{Type: "SAFE_SEARCH_DETECTION"})
	}
	if f.ImageDetection {
		features = append(features, Feature{Type: "IMAGE_PROPERTIES"})
	}

	return features
}
