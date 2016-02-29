package main

import "flag"

//Features ...
type features struct {
	//TypeUnspecified Unspecified feature type.
	TypeUnspecified bool
	//FaceDetection Run face detection.
	FaceDetection bool
	//LandmarkDetection Run landmark detection.
	LandmarkDetection bool
	//LogoDetection Run logo detection.
	LogoDetection bool
	//LabelDetection Run label detection.
	LabelDetection bool
	//TextDetection Run OCR.
	TextDetection bool
	//SafeSearchDetection Run various computer vision models to
	SafeSearchDetection bool
	//ImageDetection Compute a set of properties about the image (such as the image's dominant colors)
	ImageDetection bool
}

func handleSubCommandFlags(args []string) features {

	if len(args) == 0 {
		return features{TextDetection: true}
	}

	subFlags := flag.NewFlagSet("Feature Detections", flag.ExitOnError)
	face := subFlags.Bool("face", false, "-face, run face detection")
	land := subFlags.Bool("land", false, "-land, run landmark detection")
	logo := subFlags.Bool("logo", false, "-logo, run logo detection")
	label := subFlags.Bool("label", false, "-label, run label detection")
	text := subFlags.Bool("text", false, "-text, run OCR")
	safe := subFlags.Bool("safe", false, "-safe, run various computer vision models to")
	image := subFlags.Bool("image", false, "-image, compute a set of properties about the image such as the image's dominant colors")
	subFlags.Parse(args)

	f := features{}

	if *face {
		f.FaceDetection = true
	}
	if *land {
		f.LandmarkDetection = true
	}
	if *logo {
		f.LogoDetection = true
	}
	if *label {
		f.LabelDetection = true
	}
	if *text {
		f.TextDetection = true
	}
	if *safe {
		f.SafeSearchDetection = true
	}
	if *image {
		f.ImageDetection = true
	}

	return f
}
