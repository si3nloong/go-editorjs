# Editor.js (Go)

Server-side implementation sample for the [Editor.js](https://github.com/codex-team/editor.js). It contains data validation and converts output from Editor.js to HTML.

## Installation

```bash
go get github.com/si3nloong/go-editorjs
```

## Basic usage

```go
import (
    editorjs "github.com/si3nloong/go-editorjs"
)

func main() {
    b := []byte(`
    {
        "time": 1589098011153,
        "blocks": [
            {
            "type": "header",
            "data": {
                "text": "Editor.js",
                "level": 2
            }
            }
        ],
        "version": "2.17.0"
    }
    `)

    ejs := editorjs.NewEditorJS()
    buf := new(bytes.Buffer)
    if err := ejs.ParseTo(b, buf); err != nil {
        panic(err)
    }

}
```

## Example of Editor.js output

```json
{
  "time": 1589098011153,
  "blocks": [
    {
      "type": "header",
      "data": {
        "text": "Editor.js",
        "level": 2
      }
    }
  ],
  "version": "2.17.0"
}
```
