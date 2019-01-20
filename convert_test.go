package md2tex

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	type args struct {
		r io.Reader
		c Conf
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			"Header",
			args{
				strings.NewReader("# Header 1\n\n## Header 2"),
				Conf{
					InlineLines: []InlineLine{
						{"# ", Target{"\\chapter{", "}"}},
						{"## ", Target{"\\section{", "}"}},
					},
				},
			},
			"\\chapter{Header 1}\n\n\\section{Header 2}\n",
			false,
		},
		{
			"Code Block",
			args{
				strings.NewReader("```\ncode\n```\n"),
				Conf{
					Blocks: []Block{
						{"```", "```", Target{"begin", "end"}},
					},
				},
			},
			"begin\ncode\nend\n",
			false,
		},
		{
			"Code Inline",
			args{
				strings.NewReader("abc `code` def\nmore `code2` comes `here`\n"),
				Conf{
					Inlines: []Inline{
						{"`", "`", Target{"code{", "}"}},
					},
				},
			},
			"abc code{code} def\nmore code{code2} comes code{here}\n",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := Convert(tt.args.r, w, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Convert() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
