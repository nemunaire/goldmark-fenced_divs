package fenced_divs

import (
	"log"
	"os"
	"testing"

	"github.com/yuin/goldmark"
)

func TestAttributes(t *testing.T) {
	source := []byte(`
Text 1

::::: {.text-center}

Text *centered*

:::::

Text 2

::::: {#complex .lead}

More complex, with identifier

:::::

Text 3
`)

	var md = goldmark.New(Enable)
	err := md.Convert(source, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

}
