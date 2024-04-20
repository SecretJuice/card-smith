package main

import (
	"fmt"
	"log"

	parse "card-builder/src/parse"

	"github.com/fogleman/gg"
)

func drawTextNode(context *gg.Context, spec parse.TextSpec, content parse.TextContent) {

	context.SetRGB(0, 0, 0)
	context.DrawString(content.Text, spec.X, spec.Y)
}

func paintNodes(context *gg.Context, specs map[string]parse.SpecNode, contents map[string]parse.ContentNode) {

	for _, contentNode := range contents {

		specNode, ok := specs[contentNode.Name]
		if !ok {
			fmt.Println("Spec not found: " + contentNode.Name)
			return
		}

		switch contentNode.Type {

		case "text":

			if textContent, ok := contentNode.Content.(parse.TextContent); ok {

				if textSpec, ok := specNode.Spec.(parse.TextSpec); ok {
					drawTextNode(context, textSpec, textContent)
				} else {
					fmt.Println("Spec is not TextSpec: " + specNode.Name)
				}

			} else {
				fmt.Println("Content is not TextContent: " + contentNode.Name)
			}
		}

	}
}

func main() {
	const inputPath = "../resources/images/dog-eyes.jpg"
	const outputPath = "../resources/images/output.png"

	// Load the image
	image, err := gg.LoadImage(inputPath)

	if err != nil {
		log.Fatalf("Error loading image: %v", err)
	}

	dc := gg.NewContextForImage(image) // Create a context with zero width and height, will auto adjust

	// Load the font
	// if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
	// 	log.Fatalf("Error loading font: %v", err)
	// }

	specNodes := parse.GetSpecification("../spec_dev.json")
	contentNodes := parse.GetContent("../content_dev.json")

	fmt.Println(specNodes)
	fmt.Println(contentNodes)

	paintNodes(dc, specNodes, contentNodes)

	// Save the image
	if err := dc.SavePNG(outputPath); err != nil {
		log.Fatalf("Error saving image: %v", err)
	}

	fmt.Println("Image saved to", outputPath)

	// TODO:
	// Get Specification
	// Get Content
	// Get Image
}