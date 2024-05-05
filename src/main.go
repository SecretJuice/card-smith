package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"

	paint "card-builder/src/paint"
	parse "card-builder/src/parse"
)

func createImage(path string) image.Image {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	return img
}

func getCardImage() image.Image {
	const inputPath = "../resources/images/coolwater_sprite-_1_TEMPLATE.png"

	return createImage(inputPath)
}

func main() {

	const outputPath = "../resources/images/output.png"

	c := canvas.New(750, 1050)

	ctx := canvas.NewContext(c)

	specNodes := parse.GetSpecification("../spec_hexblitz_familiar.json")
	contentNodes := parse.GetContent("../content_coolwatersprite.json")

	// fmt.Println(specNodes)
	// fmt.Println(contentNodes)

	img := getCardImage()
	ctx.DrawImage(0, 0, img, canvas.DPMM(1))

	paint.PaintNodes(ctx, specNodes, contentNodes)

	startTime := time.Now()

	if err := renderers.Write(outputPath, c); err != nil {
		panic(err)
	}

	elapsedTime := time.Since(startTime)

	fmt.Printf("Image saved to %s: took %s after loading specs", outputPath, elapsedTime)

	// TODO:
	// Get Specification
	// Get Content
	// Get Image
}
