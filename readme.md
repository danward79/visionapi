### Google Vision Service Command Line Tool

Small command line tool to use with the google vision service.

Usage:
```
go run main.go -k APIKEYFROMGOOGLE -f imagefilepath
```
e.g.
```
go run main.go -k qwerty1234567890 -f image123.jpg
```
or pipe from Standard in (stdin)
```
go run main.go -k APIKEYFROMGOOGLE -p
```
for example
```
ls | tail -2 | go run ~/go/src/github.com/danward79/visionapi/main.go -k qwerty1234567890 -p
```
or watch a folder
```
go run main.go -k APIKEYFROMGOOGLE -w folderpath
```
for example
```
go run ~/go/src/github.com/danward79/visionapi/main.go -k qwerty1234567890 -w ./tmp
```

Returns a JSON encoded string at the moment.

TODO
- Add options around what is detected. For example:

  - LABEL_DETECTION	Execute Image Content Analysis on the entire image and return
  - TEXT_DETECTION	Perform Optical Character Recognition (OCR) on - text within the image
  - FACE_DETECTION	Detect faces within the image
  - LANDMARK_DETECTION	Detect geographic landmarks within the image
  - LOGO_DETECTION	Detect company logos within the image
  - SAFE_SEARCH_DETECTION	Determine image safe search properties on the image
  - IMAGE_PROPERTIES	Compute a set of properties about the image (such as the image's dominant colors)
