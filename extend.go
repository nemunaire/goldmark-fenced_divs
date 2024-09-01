package fenced_divs

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type Extender struct {
}

func (e *Extender) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(&FencedDivsParser{}, 100),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(&FencedDivsRenderer{}, 100),
		),
	)
}

// Extension is a goldmark.Extender with markdown fenced block attributes support.
var Extension goldmark.Extender = new(Extender)

// Enable is a goldmark.Option with fenced block attributes support.
var Enable = goldmark.WithExtensions(Extension)
