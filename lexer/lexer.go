package lexer

type Lexer struct {
	input        string
	position     int  // current index in input string
	readPosition int  // next index
	ch           byte // current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
