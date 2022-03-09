package editorjs

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/require"
)

func TestEditorJS(t *testing.T) {
	var (
		ejs = NewEditorJS()
		b   []byte
		err error
	)

	b, err = ioutil.ReadFile("./examples/data.json")
	require.NoError(t, err)

	buf := new(bytes.Buffer)
	buf.WriteString(`<html>`)
	buf.WriteString(`<head>`)
	buf.WriteString(`
	<title>WeTix</title>
    <meta charset="utf8" />
    <meta
      name="Description"
      content="Malaysia first movie ticketing aggregator."
    />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta http-equiv="refresh" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0, viewport-fit=cover"
    />
	`)
	buf.WriteString(`<link href="https://wetix-assets.oss-ap-southeast-3.aliyuncs.com/css/global.css" rel="stylesheet" />`)
	buf.WriteString(`</head>`)
	buf.WriteString(`<body>`)
	buf.WriteString(`<main class="padding">`)
	err = ejs.ParseTo(b, buf)
	require.NoError(t, err)
	buf.WriteString(`</main>`)
	buf.WriteString(`</body>`)
	buf.WriteString(`</html>`)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(buf)
	require.NoError(t, err)

	// verify the html content is matched
	{
		require.True(t, len(doc.Find("p").Nodes) == 7)
		require.True(t, len(doc.Find("a").Nodes) == 2)
		require.True(t, len(doc.Find("code").Nodes) == 5)
		require.True(t, len(doc.Find("mark").Nodes) == 1)
		require.True(t, len(doc.Find("ol").Nodes) == 1)
		require.True(t, len(doc.Find("ul").Nodes) == 1)
		require.True(t, len(doc.Find("li").Nodes) == 7)
		require.True(t, len(doc.Find("table").Nodes) == 1)
	}
}
