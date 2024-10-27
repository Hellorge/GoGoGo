package metaparser

import (
	"html/template"

	"github.com/BurntSushi/toml"
)

type MetaData struct {
	Template     string                 `toml:"template"`
	InlineStyle  bool                   `toml:"inlineStyle"`
	InlineScript bool                   `toml:"inlineScript"`
	Head         []template.HTML        `toml:"head"`
	CSSImports   []string               `toml:"cssImports"`
	JSImports    []string               `toml:"jsImports"`
	Variables    map[string]interface{} `toml:"variables"`
}

func ParseMetaData(data []byte) (MetaData, error) {
	var meta MetaData
	err := toml.Unmarshal(data, &meta)
	return meta, err
}
