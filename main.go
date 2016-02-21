package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func init() {
	fmt.Println("Google Vision API - Interface")
}

func main() {
	apiKey := flag.String("k", "", "-k api key")
	pipe := flag.Bool("p", false, "-p use unix pipe to pass a file into Vision")
	//watch := flag.String("w", "", "-w watch a path for new files")
	file := flag.String("f", "", "-f file path")
	flag.Parse()

	if *apiKey == "" && (*pipe || *file == "") {
		fmt.Println("No command line arguments specified, usage: ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *pipe {
		fmt.Println("Piping it!")

		info, err := os.Stdin.Stat()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(info.Name())
		fmt.Println(info)

		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		fmt.Println(string(l))
		*file = string(l)
	}

	encodedImage, err := encodeBase64(*file)
	if err != nil {
		log.Fatal(err)
	}

	var url = "https://vision.googleapis.com/v1/images:annotate?key="

	s, err := postRequest(url+*apiKey, marshalJSON(encodedImage))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)

}

//postRequest
func postRequest(u string, j []byte) (string, error) {

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "", fmt.Errorf("Request Err: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body), nil
	}

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return "", fmt.Errorf("Request Err: %v", err)
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

//marshalJSON JSON message
func marshalJSON(image string) []byte {
	return []byte(fmt.Sprintf(
		`{
    "requests": [{
      "image": {
        "content": "%s"
      },
      "features": [{
        "type": "TEXT_DETECTION",
        "maxResults": "100"
      }]
    }]
  }`, image))
}
