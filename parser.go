package fenced_divs

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type fenceData struct {
	char   byte
	indent int
	length int
	node   ast.Node
	stack  *fenceData
}

var fencedDivInfoKey = parser.NewContextKey()

type FencedDivsParser struct {
}

func (s *FencedDivsParser) Trigger() []byte {
	return []byte{':'}
}

func (s *FencedDivsParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, segment := reader.PeekLine()
	pos := pc.BlockOffset()
	if pos < 0 || line[pos] != ':' {
		return nil, parser.NoChildren
	}
	findent := pos
	fenceChar := line[pos]
	i := pos
	for ; i < len(line) && line[i] == fenceChar; i++ {
	}
	oFenceLength := i - pos
	if oFenceLength < 5 {
		return nil, parser.NoChildren
	}
	var attrs parser.Attributes
	if i < len(line)-1 {
		rest := line[i:]
		left := util.TrimLeftSpaceLength(rest)
		right := util.TrimRightSpaceLength(rest)
		if left < len(rest)-right {
			infoStart, infoStop := segment.Start-segment.Padding+i+left, segment.Stop-right
			value := rest[left : len(rest)-right]
			if fenceChar == ':' && bytes.IndexByte(value, ':') > -1 {
				return nil, parser.NoChildren
			} else if infoStart != infoStop {
				reader.Advance(i + 1)
				attrs, _ = parser.ParseAttributes(reader)
			}
		}
	}
	node := &Node{
		BaseBlock: ast.BaseBlock{},
	}
	for _, attr := range attrs {
		node.SetAttribute(attr.Name, attr.Value)
	}

	var previous *fenceData
	if v := pc.Get(fencedDivInfoKey); v != nil {
		previous = v.(*fenceData)
	}
	pc.Set(fencedDivInfoKey, &fenceData{fenceChar, findent, oFenceLength, node, previous})

	return node, parser.NoChildren
}

func (b *FencedDivsParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	line, segment := reader.PeekLine()
	fdata := pc.Get(fencedDivInfoKey).(*fenceData)

	w, pos := util.IndentWidth(line, reader.LineOffset())
	if w < 4 {
		i := pos
		for ; i < len(line) && line[i] == fdata.char; i++ {
		}
		length := i - pos
		if length >= fdata.length && util.IsBlank(line[i:]) {
			newline := 1
			if line[len(line)-1] != '\n' {
				newline = 0
			}
			reader.Advance(segment.Stop - segment.Start - newline + segment.Padding)
			return parser.Close
		}
	}
	return parser.Continue | parser.HasChildren
}

func (b *FencedDivsParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	fdata := pc.Get(fencedDivInfoKey).(*fenceData)
	if fdata.node == node {
		pc.Set(fencedDivInfoKey, fdata.stack)
	}
}

func (b *FencedDivsParser) CanInterruptParagraph() bool {
	return true
}

func (b *FencedDivsParser) CanAcceptIndentedLine() bool {
	return false
}
