package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
)

func init() {
	fmt.Println("Google Vision API - Command Line Interface")
}

func main() {

	apiKey := flag.String("k", "", "-k api key")
	pipe := flag.Bool("p", false, "-p use unix pipe to pass a file into Vision")
	watch := flag.String("w", "", "-w watch a path for new files")
	file := flag.String("f", "", "-f file path")
	flag.Parse()

	requiredFeatures := buildFeatures(handleSubCommandFlags(flag.Args()[:]))

	if *apiKey == "" && (*pipe || *file == "" || *watch == "") {
		fmt.Println("No command line arguments specified, usage: ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var url = "https://vision.googleapis.com/v1/images:annotate?key=" + *apiKey

	if *file != "" {
		err := processImage(*file, url, &requiredFeatures)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *pipe {

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {

			file := scanner.Text()

			err := processImage(file, url, &requiredFeatures)
			if err != nil {
				log.Fatal(err)
			}

		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error, reading stdin:", err)
		}
	}

	if *watch != "" {
		watcherDone := make(chan struct{})
		err := watchPath(*watch, url, watcherDone, &requiredFeatures)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//watchPath
func watchPath(path, url string, done chan struct{}, requiredFeatures *[]Feature) error {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("NewWatcher Err: %v", err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Events:

				if event.Op&fsnotify.Create == fsnotify.Create {

					err := processImage(event.Name, url, requiredFeatures)
					if err != nil {
						log.Fatal(err)
					}

				}

			case err := <-watcher.Errors:
				log.Println("Watcher error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		return fmt.Errorf("Watcher Add Err: %v", err)
	}
	<-done
	return nil
}

//processImage wrapper for encoding, marshaling and request
func processImage(filePath, url string, requiredFeatures *[]Feature) error {
	log.Println("Processing: ", filePath)

	encodedImage, err := encodeBase64(filePath)
	if err != nil {
		return err
	}

	s, err := postRequest(url, marshalJSON(encodedImage, requiredFeatures))
	if err != nil {
		return err
	}

	fmt.Println("Result: ", s)
	return nil
}

//postRequest to service
func postRequest(url string, json []byte) (string, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return "", fmt.Errorf("Post Request Err: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Request Err: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body), nil
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return "", fmt.Errorf("Other Request Err: %v, %v, %v", err, resp.StatusCode, string(body))
}
