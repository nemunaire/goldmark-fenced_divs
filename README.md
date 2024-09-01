# goldmark-fenced_divs

[GoldMark](https://github.com/yuin/goldmark/) fenced_divs extension.

This implements the [`fenced_divs`](https://pandoc.org/MANUAL.html#extension-fenced_divs) of pandoc.

```markdown
::::: {#special .sidebar}
Here is a paragraph.

And another.
:::::
```

```html
<div id="special" class="sidebar">
<p>Here is a paragraph.</p>
<p>And another.</p>
</div>
```

```go
var md = goldmark.New(fenced_divs.Enable)
var source = []byte(`::::: {#special .sidebar}
Here is a paragraph.

And another.
:::::`)
err := md.Convert(source, os.Stdout)
if err != nil {
    log.Fatal(err)
}
```
