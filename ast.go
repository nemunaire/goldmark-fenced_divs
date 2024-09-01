package fenced_divs

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
)

// Kind is the kind of hashtag AST nodes.
var Kind = ast.NewNodeKind("FencedDivs")

// Node is a parsed Superscript node.
type Node struct {
	ast.BaseBlock
	Attrs parser.Attributes
}

// IsRaw implements Node.IsRaw.
func (n *Node) IsRaw() bool {
	return true
}

// Dump implements Node.Dump .
func (n *Node) Dump(source []byte, level int) {
	m := map[string]string{}
	ast.DumpHelper(n, source, level, m, nil)
}

func (*Node) Kind() ast.NodeKind { return Kind }
