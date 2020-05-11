package editorjs

import (
	"encoding/json"
	"io"
	"strings"
	"sync"

	"bytes"
)

// Writer :
type Writer interface {
	io.StringWriter
	io.Writer
}

// Flusher :
type Flusher interface {
	Flush() error
}

// ParseFunc :
type ParseFunc func(b []byte, w Writer) error

// EditorJS :
type EditorJS struct {
	atomic  sync.Mutex
	parsers map[string]ParseFunc
}

// NewEditorJS :
func NewEditorJS() *EditorJS {
	ejs := new(EditorJS)
	ejs.parsers = make(map[string]ParseFunc)
	// register default parser
	registerDefaultParsers(ejs, DefaultParser{})
	return ejs
}

// RegisterParser :
func (ejs *EditorJS) RegisterParser(name string, p ParseFunc) {
	ejs.atomic.Lock()
	defer ejs.atomic.Unlock()
	ejs.parsers[name] = p
}

// ParseTo :
func (ejs *EditorJS) ParseTo(b []byte, w Writer) error {
	r := bytes.NewBuffer(b)
	data := Data{}
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return err
	}
	f, flusher := w.(Flusher)
	for _, blk := range data.Blocks {
		blk.Type = strings.ToLower(strings.TrimSpace(blk.Type))
		parseData, ok := ejs.parsers[blk.Type]
		if !ok {
			continue
		}
		if err := parseData(blk.Data, w); err != nil {
			return err
		}
		if flusher {
			f.Flush()
		}
	}
	return nil
}
