package editorjs

import "encoding/json"

// Data
type Data struct {
	Blocks  []Block `json:"blocks"`
	Version string  `json:"version"`
}

type Block struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type List struct {
	Style string   `json:"style"`
	Items []string `json:"items"`
}

type Image struct {
	File struct {
		Url string `json:"url"`
	} `json:"file"`
	Caption   string `json:"caption"`
	Stretched bool   `json:"stretched"`
}
