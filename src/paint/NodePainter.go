package paint

import (
	"fmt"
	"image/color"
	"time"

	"github.com/tdewolff/canvas"

	parse "card-builder/src/parse"
)

func drawTextNode(ctx *canvas.Context, spec parse.TextSpec, content parse.TextContent) {

	face := fontFaceFromFile("../resources/fonts/"+spec.Font, spec.FontSize, spec.GetColor())

	text := canvas.NewTextLine(face, content.Text, canvas.Center)

	ctx.DrawText(spec.X, spec.Y, text)

}

func PaintNodes(cxt *canvas.Context, specs map[string]parse.SpecNode, contents map[string]parse.ContentNode) {

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

func fontFaceFromFile(path string, size float64, color color.Color) *canvas.FontFace {
	fontFamily := canvas.NewFontFamily("generic")
	if err := fontFamily.LoadFontFile(path, canvas.FontRegular); err != nil {
		fmt.Println("Failed to load font")
		panic(err)
	}

	face := fontFamily.Face(size, color, canvas.FontRegular, canvas.FontNormal)

	return face
}
