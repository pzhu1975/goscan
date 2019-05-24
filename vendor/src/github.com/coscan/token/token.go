// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token defines constants representing the lexical tokens of the Go
// programming language and basic operations on tokens (printing, predicates).
//
package token

import "strconv"

// Token is the set of lexical tokens of the Go programming language.
type Token int

// The list of tokens.
const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT

	literal_beg
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT  // main
	INT    // 12345
	FLOAT  // 123.45
	// IMAG   // 123.45i
	CHAR   // 'a'
	STRING // "abc"
	literal_end

	operator_beg
	// Operators and delimiters
	ADD   // +
	SUB   // -
	MUL   // *
	QUO   // /
	REM   // %
	TILDE // ~

	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^

	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	QUO_ASSIGN // /=
	REM_ASSIGN // %=

	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	XOR_ASSIGN     // ^=
	SHL_ASSIGN     // <<=
	SHR_ASSIGN     // >>=
	// AND_NOT_ASSIGN // &^=

	LAND  // &&
	LOR   // ||
	ARROW // <-
	INC   // ++
	DEC   // --

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !
	QUESTIONMARK // ?
	PREPROCESSOR // #
	BACKSLASH // \\
	MINUSGREATER // ->
	HASHHASH // ##

	NEQ      // !=
	LEQ      // <=
	GEQ      // >=
	DEFINE   // :=
	ELLIPSIS // ...

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
	operator_end

	keyword_beg
	// Keywords
	AUTO
	BREAK
	CASE
	// CHAN
	CONST
	CONTINUE

	DEFAULT
	DO
	DOUBLE
	// DEFER
	ELSE
	ENUM
	EXTERN
	// FALLTHROUGH
	FOR

	// FUNC
	// GO
	GOTO
	IF
	IMPORT
	WHILE

	// INTERFACE
	// MAP
	PACKAGE
	// RANGE
	RETURN

	// SELECT
	STRUCT
	SWITCH
	TYPE
	LONG
	REGISTER
	SIZEOF
	STATIC
	TYPEDEF
	UNION
	SHORT
	SIGNED
	UNSIGNED
	VOID
	VOLATILE
	PRAGMA
	__VOLATILE
	__VOLATILE__
	_COMPLEX
	__COMPLEX__
	__COMPLEX
	__FUNC__
	_IMAGINARY
	INLINE
	__INLINE
	__INLINE__
	__RESTRICT__
	__RESTRICT__
	__RESTRICT
	_ATOMIC
	__THREAD
	ASM
	__ASM__
	__ASM
	TYPEOF
	__TYPEOF__
	__ATTRIBUTE__
	ALIGNED
	PACK
	PACKED
	VAR
	keyword_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",

	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT:  "IDENT",
	INT:    "INT",
	FLOAT:  "FLOAT",
	// IMAG:   "IMAG",
	CHAR:   "CHAR",
	STRING: "STRING",

	ADD:   "+",
	SUB:   "-",
	MUL:   "*",
	QUO:   "/",
	REM:   "%",
	TILDE: "~",

	AND:     "&",
	OR:      "|",
	XOR:     "^",
	SHL:     "<<",
	SHR:     ">>",
	AND_NOT: "&^",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	QUO_ASSIGN: "/=",
	REM_ASSIGN: "%=",

	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	// AND_NOT_ASSIGN: "&^=",

	LAND:  "&&",
	LOR:   "||",
	ARROW: "<-",
	INC:   "++",
	DEC:   "--",

	EQL:    "==",
	LSS:    "<",
	GTR:    ">",
	ASSIGN: "=",
	NOT:    "!",
	QUESTIONMARK: "?"
	PREPROCESSOR: "#"
	BACKSLASH: "\\"
	MINUSGREATER: "->"
	HASHHASH: "##"

	NEQ:      "!=",
	LEQ:      "<=",
	GEQ:      ">=",
	DEFINE:   ":=",
	ELLIPSIS: "...",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",

	AUTO:	  "auto"
	BREAK:    "break",
	CASE:     "case",
	// CHAN:     "chan",
	CONST:    "const",
	CONTINUE: "continue",

	DEFAULT:     "default",
	DO:			"do",
	// DEFER:       "defer",
	DOUBLE:		"double"
	ELSE:        "else",
	ENUM:		 "enum",
	EXTERN:		 "extern",
	// FALLTHROUGH: "fallthrough",
	FOR:         "for",

	// FUNC:   "func",
	// GO:     "go",
	GOTO:   "goto",
	IF:     "if",
	IMPORT: "import",
	WHILE:	"while",

	// INTERFACE: "interface",
	// MAP:       "map",
	PACKAGE:   "package",
	// RANGE:     "range",
	RETURN:    "return",

	// SELECT: "select",
	STRUCT: "struct",
	SWITCH: "switch",
	TYPE:   "type",
	LONG:	"long",
	REGISTER: "register",
	SHORT: "short",
	SIGNED: "signed",
	SIZEOF: "sizeof"
	STATIC: "static"
	TYPEDEF: "typedef",
	UNION: "union",
	UNSIGNED: "unsigned",
	VOID:	"void",
	VOLATILE: "volatile",
	PRAGMA: "pragma",
	__VOLATILE: "__volatile",
	__VOLATILE__: "__volatile__",
	_BOOL: "_Bool",
	_COMPLEX: "_Complex",
	__COMPLEX__: "__complex__",
	__COMPLEX: "__complex",
	__FUNC__: "__func__",
	__FUNCTION__: "__FUNCTION__",
	_IMAGINARY: "_Imaginary",
	INLINE: "inline",
	__INLINE: "__inline",
	__INLINE__: "__inline__",
	__RESTRICT__: "__restrict__",
	__RESTRICT: "__restrict",
	_ATOMIC: "_Atomic",
	__THREAD: "__thread",
	ASM: "asm",
	__ASM__: "__asm__",
	__ASM: "__asm",
	TYPEOF: "typeof",
	__TYPEOF__: "__typeof__",
	__ATTRIBUTE__: "__attribute__",
	ALIGNED: "aligned",
	PACK: "pack",
	PACKED: "packed",
	VAR:    "var",
}

// String returns the string corresponding to the token tok.
// For operators, delimiters, and keywords the string is the actual
// token character sequence (e.g., for the token ADD, the string is
// "+"). For all other tokens the string corresponds to the token
// constant name (e.g. for the token IDENT, the string is "IDENT").
//
func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

// A set of constants for precedence-based expression parsing.
// Non-operators have lowest precedence, followed by operators
// starting with precedence 1 up to unary operators. The highest
// precedence serves as "catch-all" precedence for selector,
// indexing, and other operator and delimiter tokens.
//
const (
	LowestPrec  = 0 // non-operators
	UnaryPrec   = 6
	HighestPrec = 7
)

// Precedence returns the operator precedence of the binary
// operator op. If op is not a binary operator, the result
// is LowestPrecedence.
//
func (op Token) Precedence() int {
	switch op {
	case LOR:
		return 1
	case LAND:
		return 2
	case EQL, NEQ, LSS, LEQ, GTR, GEQ:
		return 3
	case ADD, SUB, OR, XOR:
		return 4
	case MUL, QUO, REM, SHL, SHR, AND, AND_NOT:
		return 5
	}
	return LowestPrec
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token)
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
//
func Lookup(ident string) Token {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}
	return IDENT
}

// Predicates

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
//
func (tok Token) IsLiteral() bool { return literal_beg < tok && tok < literal_end }

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
//
func (tok Token) IsOperator() bool { return operator_beg < tok && tok < operator_end }

// IsKeyword returns true for tokens corresponding to keywords;
// it returns false otherwise.
//
func (tok Token) IsKeyword() bool { return keyword_beg < tok && tok < keyword_end }
