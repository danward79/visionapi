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

Returns a JSON encoded string at the moment.

TODO
- Watch path for new files -w command flag.
- Add options around what is detected. For example: LANDMARK_DETECTION, LOGO_DETECTION, LABEL_DETECTION, etc
