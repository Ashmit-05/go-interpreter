package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"
	// Operators
	ASSIGN = "="
	PLUS   = "+"
  MINUS = "-"
  BANG = "!"
  ASTERISK = "*"
  SLASH = "/"
  LT = "<"
  GT = ">"
  EQ = "=="
  NOT_EQ = "!="
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	// 1343456
	FUNCTION = "FUNCTION"
	LET      = "LET"
  TRUE = "true"
  FALSE = "false"
  IF = "if"
  ELSE = "else"
  RETURN = "return"
)

var keywords = map[string]TokenType {
  "fn" : FUNCTION,
  "let" : LET,
  "if" : IF,
  "else" : ELSE,
  "true" : TRUE,
  "false" : FALSE,
  "return" : RETURN,
}

func LookUpIdent(ident string) TokenType {
  if tok,ok := keywords[ident]; ok {
    return tok
  } 
  return IDENT
}
