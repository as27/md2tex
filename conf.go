package md2tex

var SimpleConf = Conf{
	InlineLines: []InlineLine{
		{"# ", Target{"\\chapter{", "}"}},
		{"## ", Target{"\\section{", "}"}},
		{"### ", Target{"\\subsection{", "}"}},
	},
	Blocks: []Block{
		{"```", "```",
			Target{
				"\\begin{lstlisting}[caption={},label={lst:1}]",
				"\\end{lstlisting}",
			},
		},
		{"```go", "```",
			Target{
				"\\begin{lstlisting}[caption={},label={lst:1}]",
				"\\end{lstlisting}",
			},
		},
		{"```Go", "```",
			Target{
				"\\begin{lstlisting}[caption={},label={lst:1}]",
				"\\end{lstlisting}",
			},
		},
	},
	Inlines: []Inline{
		{"`", "`", Target{"\\code{", "}"}},
		{"_", "_", Target{"\\emph{", "}"}},
	},
}
