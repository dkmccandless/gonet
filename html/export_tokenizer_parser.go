package html

type TokenStream interface {
	AllowCDATA(bool)
	Err() error
	Next() TokenType
	NextIsNotRawText()
	Token() Token
}

func ParseOptionTokenStream(ts TokenStream) ParseOption {
	return func(p *parser) {
		p.tokenizer = ts
	}
}
