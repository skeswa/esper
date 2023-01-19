package lexer

// Maps each kind of keyword string to its corresponding token `T`.
var Keywords = map[string]T{
	// Reserved words
	"as":       TAs,
	"async":    TAsync,
	"await":    TAwait,
	"break":    TBreak,
	"continue": TContinue,
	"else":     TElse,
	"enum":     TEnum,
	"extern":   TExtern,
	"false":    TFalse,
	"fn":       TFn,
	"for":      TFor,
	"fork":     TFork,
	"from":     TFrom,
	"if":       TIf,
	"impl":     TImpl,
	"in":       TIn,
	"is":       TIs,
	"let":      TLet,
	"loop":     TLoop,
	"match":    TMatch,
	"mod":      TMod,
	"pub":      TPub,
	"return":   TReturn,
	"self":     TSelf,
	"selftype": TSelfType,
	"show":     TShow,
	"struct":   TStruct,
	"trait":    TTrait,
	"true":     TTrue,
	"try":      TTry,
	"type":     TType,
	"unknown":  TUnknown,
	"use":      TUse,
	"void":     TVoid,
	"where":    TWhere,
	"while":    TWhile,
}
