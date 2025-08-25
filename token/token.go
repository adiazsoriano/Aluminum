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
	Type     TokenType
	SubT     TokenSubType
	Data     TokenContent
	Filename string
	Line     uint64
	Col      uint64
}

type TokenInfo struct {
	Name  string
	Value string
}

type Tokens []Token

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
	MODULO   ArithType = 204

	EQUAL           CompareType = 300
	NOT_EQUAL       CompareType = 301
	GREATER_THAN    CompareType = 302
	LESS_THAN       CompareType = 303
	GREATER_THAN_EQ CompareType = 304
	LESS_THAN_EQ    CompareType = 305

	FN     KeywordType = 400 //func
	ST     KeywordType = 401 //struct
	IF     KeywordType = 402
	ELSE   KeywordType = 403
	ELSEIF KeywordType = 404
	FOR    KeywordType = 405
	WHILE  KeywordType = 406

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

var tokenSubTypeInfo = map[TokenSubType]TokenInfo{
	INVALIDSUB: {
		Name: "INVALID SUBTYPE",
	},
	NOSUB: {
		Name: "NO SUBTYPE",
	},
}

var varTypeInfo = map[VarType]TokenInfo{
	INTTYPE: {
		Name:  "INT TYPE",
		Value: "int",
	},
	FLOATTYPE: {
		Name:  "FLOAT TYPE",
		Value: "float",
	},
	BOOLTYPE: {
		Name:  "BOOL TYPE",
		Value: "bool",
	},
	STRINGTYPE: {
		Name:  "STRING TYPE",
		Value: "string",
	},
}

var arithTypeInfo = map[ArithType]TokenInfo{
	PLUS: {
		Name:  "PLUS",
		Value: "+",
	},
	MINUS: {
		Name:  "MINUS",
		Value: "-",
	},
	DIVIDE: {
		Name:  "DIVIDE",
		Value: "/",
	},
	MULTIPLY: {
		Name:  "MULTIPLY",
		Value: "*",
	},
	MODULO: {
		Name:  "MODULO",
		Value: "%",
	},
}

var compareTypeInfo = map[CompareType]TokenInfo{
	EQUAL: {
		Name:  "EQUAL",
		Value: "==",
	},
	NOT_EQUAL: {
		Name:  "NOT EQUAL",
		Value: "!=",
	},
	GREATER_THAN: {
		Name:  "GREATER THAN",
		Value: ">",
	},
	LESS_THAN: {
		Name:  "LESS THAN",
		Value: "<",
	},
	GREATER_THAN_EQ: {
		Name:  "GREATER THAN OR EQUAL",
		Value: ">=",
	},
	LESS_THAN_EQ: {
		Name:  "LESS THAN OR EQUAL",
		Value: "<=",
	},
}

var keywordTypeInfo = map[KeywordType]TokenInfo{
	FN: {
		Name:  "FUNCTION",
		Value: "func",
	},
	ST: {
		Name:  "STRUCT",
		Value: "struct",
	},
	IF: {
		Name:  "IF",
		Value: "if",
	},
	ELSE: {
		Name:  "ELSE",
		Value: "else",
	},
	ELSEIF: {
		Name:  "ELSEIF",
		Value: "elseif",
	},
	FOR: {
		Name:  "FOR",
		Value: "for",
	},
	WHILE: {
		Name:  "WHILE",
		Value: "while",
	},
}

var literalTypeInfo = map[LiteralType]TokenInfo{
	INTLIT: {
		Name: "INT LTIERAL",
	},
	FLOATLIT: {
		Name: "FLOAT LITERAL",
	},
	BOOLLIT: {
		Name: "BOOL LITERAL",
	},
	STRINGLIT: {
		Name: "STRING LITERAL",
	},
}

var logicalTypeInfo = map[LogicalType]TokenInfo{
	AND: {
		Name:  "AND",
		Value: "&&",
	},
	OR: {
		Name:  "OR",
		Value: "||",
	},
	NOT: {
		Name:  "NOT",
		Value: "!",
	},
}

var delimTypeInfo = map[DelimType]TokenInfo{
	LBRACE: {
		Name:  "LEFT BRACE",
		Value: "{",
	},
	RBRACE: {
		Name:  "RIGHT BRACE",
		Value: "}",
	},
	LBRACKET: {
		Name:  "LEFT BRACKET",
		Value: "[",
	},
	RBRACKET: {
		Name:  "RIGHT BRACKET",
		Value: "]",
	},
	LPAREN: {
		Name:  "LEFT PARENTHESIS",
		Value: "(",
	},
	RPAREN: {
		Name:  "RIGHT PARENTHESIS",
		Value: ")",
	},
}

var tokenTypeInfo = map[TokenType]TokenInfo{
	INVALID: {
		Name: "INVALID",
	},
	VARTYPE: {
		Name: "VARTYPE",
	},
	IDENT: {
		Name: "IDENTIFIER",
	},
	LITERAL: {
		Name: "LITERAL",
	},
	ASSIGN: {
		Name:  "ASSIGNMENT",
		Value: "=",
	},
	ARITH: {
		Name: "ARITHMETIC",
	},
	COMPARE: {
		Name: "COMPARE",
	},
	KEYWORD: {
		Name: "KEYWORD",
	},
	LOGICAL: {
		Name: "LOGICAL",
	},
	DELIM: {
		Name: "DELIMITER",
	},
	COMMA: {
		Name:  "COMMA",
		Value: ",",
	},
	SEMI: {
		Name:  "SEMICOLON",
		Value: ";",
	},
	EOF: {
		Name: "EOF",
	},
}
