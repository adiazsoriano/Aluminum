package token

import "math"

type TokenType uint64
type TokenSubType uint64
type TokenContent string

type NilType uint64
type VarType uint64
type ArithType uint64
type CompareType uint64
type KeywordType uint64
type LiteralType uint64
type LogicalType uint64
type DelimType uint64

type Token struct {
	t        TokenType
	sub      TokenSubType
	data     TokenContent
	filename string
	line     uint64
	col      uint64
}

type Tokens []Token
type TokenFile []Tokens

const (
	INVALIDSUB TokenSubType = 0
	NOSUB      TokenSubType = math.MaxUint64

	INTTYPE    VarType = 100
	FLOATTYPE  VarType = 101
	BOOLTYPE   VarType = 102
	STRINGTYPE VarType = 103

	PLUS     ArithType = 200
	MINUS    ArithType = 201
	DIVIDE   ArithType = 202
	MULTIPLY ArithType = 203

	EQUAL           CompareType = 300
	NOT_EQUAL       CompareType = 301
	GREATER_THAN    CompareType = 302
	LESS_THAN       CompareType = 303
	GREATER_THAN_EQ CompareType = 304
	LESS_THAN_EQ    CompareType = 305

	FN    KeywordType = 400 //func
	ST    KeywordType = 401 //struct
	IF    KeywordType = 402
	FOR   KeywordType = 403
	WHILE KeywordType = 404

	INTLIT    LiteralType = 500
	FLOATLIT  LiteralType = 501
	BOOLLIT   LiteralType = 502
	STRINGLIT LiteralType = 503

	AND LogicalType = 600
	OR  LogicalType = 601
	NOT LogicalType = 602

	LBRACE   DelimType = 1000
	RBRACE   DelimType = 1001
	LBRACKET DelimType = 1002
	RBRACKET DelimType = 1003
	LPAREN   DelimType = 1004
	RPAREN   DelimType = 1005

	INVALID TokenType = 0
	VARTYPE TokenType = 1
	IDENT   TokenType = 2
	LITERAL TokenType = 3
	ASSIGN  TokenType = 4
	ARITH   TokenType = 5
	COMPARE TokenType = 6
	KEYWORD TokenType = 7
	LOGICAL TokenType = 8
	DELIM   TokenType = 96
	COMMA   TokenType = 97
	SEMI    TokenType = 98
	EOF     TokenType = 99
)

var TokenTypeNames = map[TokenType]string{
	INVALID: "INVALID",
	VARTYPE: "VARTYPE",
	IDENT:   "IDENTIFIER",
	LITERAL: "LITERAL",
	ASSIGN:  "ASSIGNMENT",
	ARITH:   "ARITHMETIC",
	COMPARE: "COMPARE",
	KEYWORD: "KEYWORD",
	LOGICAL: "LOGICAL",
	DELIM:   "DELIMITER",
	COMMA:   "COMMA",
	SEMI:    "SEMICOLON",
	EOF:     "EOF",
}

var TokenSubTypeNames = map[TokenSubType]string{
	INVALIDSUB: "INVALIDSUBTYPE",
	NOSUB:      "NOSUBTYPE",
}
