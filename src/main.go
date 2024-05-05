package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"

	parse "card-builder/src/parse"
)

func drawTextNode(ctx *canvas.Context, spec parse.TextSpec, content parse.TextContent) {

	face := fontFaceFromFile("../resources/fonts/"+spec.Font, spec.FontSize, spec.GetColor())

	text := canvas.NewTextLine(face, content.Text, canvas.Center)

	ctx.DrawText(spec.X, spec.Y, text)

}

func paintNodes(cxt *canvas.Context, specs map[string]parse.SpecNode, contents map[string]parse.ContentNode) {

	for _, contentNode := range contents {

		startTime := time.Now()

		specNode, ok := specs[contentNode.Name]
		if !ok {
			fmt.Println("Spec not found: " + contentNode.Name)
			return
		}

		switch contentNode.Type {

		case "text":

			if textContent, ok := contentNode.Content.(parse.TextContent); ok {

				if textSpec, ok := specNode.Spec.(parse.TextSpec); ok {
					drawTextNode(cxt, textSpec, textContent)
				} else {
					fmt.Println("Spec is not TextSpec: " + specNode.Name)
				}

			} else {
				fmt.Println("Content is not TextContent: " + contentNode.Name)
			}
		}

		elapsedTime := time.Since(startTime)

		fmt.Printf(contentNode.Name+" took %s to complete.\n", elapsedTime)

	}
}

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

func fontFaceFromFile(path string, size float64, color color.Color) *canvas.FontFace {
	fontFamily := canvas.NewFontFamily("generic")
	if err := fontFamily.LoadFontFile(path, canvas.FontRegular); err != nil {
		fmt.Println("Failed to load font")
		panic(err)
	}

	face := fontFamily.Face(size, color, canvas.FontRegular, canvas.FontNormal)

	return face
}

func main() {
	const inputPath = "../resources/images/coolwater_sprite-_1_TEMPLATE.png"
	const outputPath = "../resources/images/output.png"

	c := canvas.New(750, 1050)

	ctx := canvas.NewContext(c)

	specNodes := parse.GetSpecification("../spec_hexblitz_familiar.json")
	contentNodes := parse.GetContent("../content_coolwatersprite.json")

	// fmt.Println(specNodes)
	// fmt.Println(contentNodes)

	img := createImage(inputPath)
	ctx.DrawImage(0, 0, img, canvas.DPMM(1))

	paintNodes(ctx, specNodes, contentNodes)

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
