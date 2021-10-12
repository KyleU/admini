// Code generated by projectforge.dev using code from the [marketing] module, which is under license [CC0]
package doc

import (
	"embed"

	"github.com/gomarkdown/markdown"
	"github.com/pkg/errors"
)

//go:embed *
var assetFS embed.FS

func DocContent(path string) ([]byte, error) {
	data, err := assetFS.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading doc asset at [%s]", path)
	}

	return data, nil
}

var htmlCache = make(map[string]string, 0)

func DocHTML(path string) (string, error) {
	if curr, ok := htmlCache[path]; ok {
		return curr, nil
	}
	data, err := DocContent(path)
	if err != nil {
		return "", err
	}
	html := string(markdown.ToHTML(data, nil, nil))
	htmlCache[path] = html
	return html, nil
}
