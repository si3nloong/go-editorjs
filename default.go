package editorjs

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

// DefaultParser :
type DefaultParser struct{}

func registerDefaultParsers(ejs *EditorJS, def DefaultParser) {
	ejs.RegisterParser("header", def.ParseHeader)
	ejs.RegisterParser("paragraph", def.ParseParagraph)
	ejs.RegisterParser("list", def.ParseList)
	ejs.RegisterParser("image", def.ParseImage)
	ejs.RegisterParser("table", def.ParseTable)
}

// ParseHeader :
func (d DefaultParser) ParseHeader(b []byte, w Writer) error {
	paths := gjson.GetManyBytes(b, "text", "level")
	lvl := paths[1].Int()
	w.WriteString(fmt.Sprintf("<h%d>", lvl))
	w.WriteString(paths[0].String())
	w.WriteString(fmt.Sprintf("</h%d>", lvl))
	return nil
}

// ParseParagraph :
func (d DefaultParser) ParseParagraph(b []byte, w Writer) error {
	w.WriteString("<p>" + gjson.GetBytes(b, "text").String() + "</p>")
	return nil
}

// ParseList :
func (d DefaultParser) ParseList(b []byte, w Writer) error {
	var list List
	if err := json.Unmarshal(b, &list); err != nil {
		return err
	}
	if list.Style == "ordered" {
		w.WriteString("<ol>")
		defer w.WriteString("</ol>")
	} else {
		w.WriteString("<ul>")
		defer w.WriteString("</ul>")
	}
	for _, itm := range list.Items {
		w.WriteString("<li>" + itm + "</li>")
	}
	return nil
}

// ParseImage :
func (d DefaultParser) ParseImage(b []byte, w Writer) error {
	var img Image
	w.WriteString("<figure>")
	defer w.WriteString("</figure>")
	if err := json.Unmarshal(b, &img); err != nil {
		return err
	}
	// alt string cannot have space
	w.WriteString(`<img alt="no-image" src="` + img.File.Url + `"/>`)
	if img.Caption != "" {
		w.WriteString(`<figcaption>` + img.Caption + `</figcaption>`)
	}
	return nil
}

// ParseTable :
func (d DefaultParser) ParseTable(b []byte, w Writer) error {
	var table Table
	w.WriteString("<table>")
	defer w.WriteString("</table>")
	if err := json.Unmarshal(b, &table); err != nil {
		return err
	}

	for i, line := range table.Content {
		w.WriteString("<tr>")
		tag := `td`
		if table.WithHeadings && i == 0 {
			tag = `th`
		}

		for _, info := range line {
			w.WriteString("<" + tag + ">" + info + "</" + tag + ">")

		}

		w.WriteString("</tr>")
	}
	return nil
}
