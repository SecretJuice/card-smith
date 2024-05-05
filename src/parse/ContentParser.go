package parse

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
Root interface for the Card Content Data specification
*/
type Content interface {
	Type() string // This method is just for making Spec an interface
}

type TextContent struct {
	Text string `json:"text"`
}

func (t TextContent) Type() string { return "text" }

type ContentNode struct {
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Content Content `json:"content"`
}

type ContentRoot struct {
	Nodes []ContentNode `json:"nodes"`
}

func (n *ContentNode) UnmarshalJSON(data []byte) error {
	// Create a temporary struct to hold the raw spec as raw JSON
	type Alias ContentNode
	temp := &struct {
		Content json.RawMessage `json:"content"`
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
		var content TextContent
		if err := json.Unmarshal(temp.Content, &content); err != nil {
			return err
		}
		n.Content = content
	default:
		return fmt.Errorf("unknown type: %s", n.Type)
	}

	return nil
}

func GetContent(content string) map[string]ContentNode {

	jsonFile, err := os.Open(content)

	var contents = make(map[string]ContentNode)

	if err != nil {
		fmt.Println(err)
		return contents
	}

	defer jsonFile.Close()

	byteArray, err := os.ReadFile(content)

	if err != nil {
		fmt.Println(err)
		return contents
	}

	var root ContentRoot

	if err := json.Unmarshal(byteArray, &root); err != nil {
		fmt.Println(err)
		return contents
	}

	for _, node := range root.Nodes {
		contents[node.Name] = node
	}

	return contents
}
