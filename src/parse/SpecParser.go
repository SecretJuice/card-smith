package parse

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image/color"
	"os"
)

type Spec interface {
	isSpec() // This method is just for making Spec an interface
}

type TextSpec struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Font     string  `json:"font"`
	FontSize float64 `json:"font_size"`
	MaxWidth float64 `json:"max_width"`
	Color    string  `json:"color"`
}

type RGBA struct {
	R uint32
	G uint32
	B uint32
	A uint32
}

func (t TextSpec) isSpec() {}
func (t TextSpec) GetColor() color.RGBA {
	bytes, err := hex.DecodeString(t.Color)
	if err != nil {
		panic(err)
	}
	var c color.RGBA

	c.R = bytes[0]
	c.G = bytes[1]
	c.B = bytes[2]
	c.A = bytes[3]

	return c
}

type ImageSpec struct {
	X         int    `json:"x"`
	Y         int    `json:"y"`
	URL       string `json:"url"`
	MaxHeight int    `json:"max_height"`
	MaxWidth  int    `json:"max_width"`
}

func (i ImageSpec) isSpec() {}

type SpecNode struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Spec Spec   `json:"spec"`
}

type SpecRoot struct {
	Nodes []SpecNode `json:"nodes"`
}

func (n *SpecNode) UnmarshalJSON(data []byte) error {
	// Create a temporary struct to hold the raw spec as raw JSON
	type Alias SpecNode
	temp := &struct {
		Spec json.RawMessage `json:"spec"`
		*Alias
	}{
		Alias: (*Alias)(n),
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Determine the spec type based on the node type
	switch n.Type {
	case "text":
		var spec TextSpec
		if err := json.Unmarshal(temp.Spec, &spec); err != nil {
			return err
		}
		n.Spec = spec
	default:
		return fmt.Errorf("unknown type: %s", n.Type)
	}

	return nil
}

func GetSpecification(spec string) map[string]SpecNode {

	jsonFile, err := os.Open(spec)

	var specs = make(map[string]SpecNode)

	if err != nil {
		fmt.Println(err)
		return specs
	}

	defer jsonFile.Close()

	byteArray, err := os.ReadFile(spec)

	if err != nil {
		fmt.Println(err)
		return specs
	}

	var root SpecRoot

	if err := json.Unmarshal(byteArray, &root); err != nil {
		fmt.Println(err)
		return specs
	}

	for _, node := range root.Nodes {
		specs[node.Name] = node
	}

	return specs
}
