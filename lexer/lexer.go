package lexer

import (
	"fmt"
	"unicode"
)

type TokenType string

const (
	TokenNumber   TokenType = "Number"
	TokenPlus     TokenType = "+"
	TokenMinus    TokenType = "-"
	TokenMultiply TokenType = "*"
	TokenDivide   TokenType = "/"
	TokenLParen   TokenType = "("
	TokenRParen   TokenType = ")"
	TokenLCurly   TokenType = "{"
	TokenRCurly   TokenType = "}"
	TokenAssign   TokenType = "="
	TokenEqual    TokenType = "=="
	TokenLess     TokenType = "<"
	TokenGreater  TokenType = ">"
	TokenIdent    TokenType = "Identifier"
	TokenEOF      TokenType = "EOF"
	TokenComma    TokenType = ","
	TokenVar      TokenType = "var"
	TokenIf       TokenType = "if"
	TokenElse     TokenType = "else"
	TokenFor      TokenType = "for"
	TokenFunc     TokenType = "func"
	TokenReturn   TokenType = "return"
)

type Token struct {
	Type  TokenType
	Value string
}

func Lex(input string) []Token {
	var tokens []Token
	pos := 0

	for pos < len(input) {
		ch := input[pos]

		if unicode.IsSpace(rune(ch)) {
			pos++
			continue
		}

		switch ch {
		case '+':
			tokens = append(tokens, Token{Type: TokenPlus, Value: "+"})
			pos++
		case '-':
			tokens = append(tokens, Token{Type: TokenMinus, Value: "-"})
			pos++
		case '*':
			tokens = append(tokens, Token{Type: TokenMultiply, Value: "*"})
			pos++
		case '/':
			tokens = append(tokens, Token{Type: TokenDivide, Value: "/"})
			pos++
		case '(':
			tokens = append(tokens, Token{Type: TokenLParen, Value: "("})
			pos++
		case ')':
			tokens = append(tokens, Token{Type: TokenRParen, Value: ")"})
			pos++
		case '{':
			tokens = append(tokens, Token{Type: TokenLCurly, Value: "{"})
			pos++
		case '}':
			tokens = append(tokens, Token{Type: TokenRCurly, Value: "}"})
			pos++
		case ',':
			tokens = append(tokens, Token{Type: TokenComma, Value: ","})
			pos++
		case '=':
			if pos+1 < len(input) && input[pos+1] == '=' {
				tokens = append(tokens, Token{Type: TokenEqual, Value: "=="})
				pos += 2
			} else {
				tokens = append(tokens, Token{Type: TokenAssign, Value: "="})
				pos++
			}
		case '<':
			tokens = append(tokens, Token{Type: TokenLess, Value: "<"})
			pos++
		case '>':
			tokens = append(tokens, Token{Type: TokenGreater, Value: ">"})
			pos++
		default:
			if unicode.IsDigit(rune(ch)) {
				start := pos
				for pos < len(input) && unicode.IsDigit(rune(input[pos])) {
					pos++
				}
				tokens = append(tokens, Token{Type: TokenNumber, Value: input[start:pos]})
			} else if unicode.IsLetter(rune(ch)) {
				start := pos
				for pos < len(input) && (unicode.IsLetter(rune(input[pos])) || unicode.IsDigit(rune(input[pos]))) {
					pos++
				}
				value := input[start:pos]
				switch value {
				case "var":
					tokens = append(tokens, Token{Type: TokenVar, Value: value})
				case "if":
					tokens = append(tokens, Token{Type: TokenIf, Value: value})
				case "else":
					tokens = append(tokens, Token{Type: TokenElse, Value: value})
				case "for":
					tokens = append(tokens, Token{Type: TokenFor, Value: value})
				case "func":
					tokens = append(tokens, Token{Type: TokenFunc, Value: value})
				case "return":
					tokens = append(tokens, Token{Type: TokenReturn, Value: value})
				default:
					tokens = append(tokens, Token{Type: TokenIdent, Value: value})
				}
			} else {
				panic(fmt.Sprintf("Unexpected character: %c", ch))
			}
		}
	}

	tokens = append(tokens, Token{Type: TokenEOF, Value: ""})
	return tokens
}
