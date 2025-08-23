package token

import "math"

const (
	NULL NilType  = math.MaxUint64
	BASE BaseType = math.MaxUint64 - 1

	INTTYPE   VarType = 100
	FLOATTYPE VarType = 101
	BOOLTYPE  VarType = 102

	PLUS     ArithType = 200
	MINUS    ArithType = 201
	DIVIDE   ArithType = 202
	MULTIPLY ArithType = 203

	EQUAL           CompareType = 300
	GREATER_THAN    CompareType = 301
	LESS_THAN       CompareType = 302
	GREATER_THAN_EQ CompareType = 303
	LESS_THAN_EQ    CompareType = 304

	FN KeywordType = 400 //func
	ST KeywordType = 401 //struct

	INTLIT   LiteralType = 500
	FLOATLIT LiteralType = 501
	BOOLLIT  LiteralType = 502

	INVALID TokenType[NilType]     = 0
	VARTYPE TokenType[VarType]     = 1
	IDENT   TokenType[BaseType]    = 2
	LITERAL TokenType[LiteralType] = 3
	ASSIGN  TokenType[BaseType]    = 4
	ARITH   TokenType[ArithType]   = 5
	COMPARE TokenType[CompareType] = 6
	KEYWORD TokenType[KeywordType] = 7
	SEMI    TokenType[BaseType]    = 8
)

type TokenCategory interface {
	~uint64
}

type TokenType[T TokenCategory] uint64
type TokenSubType uint64
type TokenContent string

type NilType uint64
type BaseType uint64
type VarType uint64
type ArithType uint64
type CompareType uint64
type KeywordType uint64
type LiteralType uint64

type Token[T TokenCategory] struct {
	t        TokenType[T]
	sub      TokenSubType
	data     TokenContent
	filename string
	line     uint64
	col      uint64
}
