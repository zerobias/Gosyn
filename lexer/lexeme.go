package lexer

type MetaToken struct {
	Name string
	Cat  TokenClass
}

type Token struct {
	MetaToken
	Value string
}

//TODO replace Lexeme with token,
//make Lexeme value pointer to real Literals
//and create interface to this types

type Lexeme struct {
	MetaToken
	Value PLString
}

type LexemeSequence []Lexeme
type TokenList []Token

func (tl *TokenList) SelectClass(cat ...TokenClass) (result TokenList) {
	for _, token := range *tl {
		for _, currentCat := range cat {
			if token.Cat == currentCat {
				result = append(result, token)
			}
		}
	}
	return result
}

func NewLexeme(name string, cat TokenClass, text *LiteralString) *Lexeme {
	return &Lexeme{MetaToken{name, cat}, text}
}

func NewToken(name string, cat TokenClass, text string) *Token {
	return &Token{MetaToken{name, cat}, text}
}

func (t Token) deriveLexeme(text *LiteralString) *Lexeme {
	return &Lexeme{t.MetaToken, text}
}

type TokenClass int

const (
	KEYWORD TokenClass = iota
	SIGN
	DELIMETER
	TEXT
	NUMBER
)

var tokenName = map[string]TokenClass{
	"KEYWORD":   KEYWORD,
	"SIGN":      SIGN,
	"DELIMETER": DELIMETER,
	"TEXT":      TEXT,
	"NUMBER":    NUMBER,
}

func GetClassByName(name string) TokenClass {
	return tokenName[name]
}
