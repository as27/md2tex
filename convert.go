package md2tex

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

func Convert(r io.Reader, w io.Writer, c Conf) error {
	// compile Regexp before scanning
	// inlineRes := make([]*regexp.Regexp, len(c.Inlines))
	var inlineRes []*regexp.Regexp
	for _, il := range c.Inlines {
		re := regexp.MustCompile(il.Start + "(.+?)" + il.End)
		inlineRes = append(inlineRes, re)
	}
	// start scanning
	s := bufio.NewScanner(r)
line:
	for s.Scan() {
		l := s.Text()
		// check blocks
		for _, bl := range c.Blocks {
			if strings.Trim(l, " ") == strings.Trim(bl.Start, " ") {
				writeLine(w, bl.Open)
				for s.Scan() {
					l := s.Text()
					if strings.Trim(l, " ") == strings.Trim(bl.End, " ") {
						writeLine(w, bl.Close)
						continue line
					}
					writeLine(w, l)
				}
			}
		}
		// check InlineLines
		for _, ill := range c.InlineLines {
			if strings.HasPrefix(l, ill.Start) {
				t := l[len(ill.Start):]
				w.Write([]byte(ill.Open))
				w.Write([]byte(t))
				w.Write([]byte(ill.Close))
				w.Write([]byte("\n"))
				continue line
			}
		}
		// check Inlines
		for i, re := range inlineRes {
			il := c.Inlines[i]
			l = re.ReplaceAllString(l, il.Open+"$1"+il.Close)
		}
		writeLine(w, l)

	}
	return nil
}
func writeLine(w io.Writer, s string) {
	w.Write([]byte(s + "\n"))
}

type Conf struct {
	Blocks      []Block
	InlineLines []InlineLine
	Inlines     []Inline
}

type Block struct {
	Start string
	End   string
	Target
}

type InlineLine struct {
	Start string
	Target
}

type Inline struct {
	Start string
	End   string
	Target
}

// Target defines the way how a part is translated
// into tex. The content is then between Open and
// Close part.
type Target struct {
	Open  string
	Close string
}
