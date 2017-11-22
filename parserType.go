package main

type ParserType int

const (
	unknownParserType ParserType = iota
	jsonParser        ParserType = iota
	htmlParser        ParserType = iota
)
