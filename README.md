# md2tex

Simple tool for converting markdown to tex

Installation via go get:

```
go get github.com/as27/md2tex
```

# Configuration

The markdown file has different Types:

* __block__: code block
    * Line starts with a string per Line
    * the content is between start and end line
    * Line ends with a string in a line
* __inline line__: header
    * Line starts with a string
* __inline__: bold
    * There can be more inline elements inside a line