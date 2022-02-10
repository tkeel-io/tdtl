// Generated from QingQL.g4 by ANTLR 4.7.

package parser // QingQL

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 219,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 55, 10, 4, 12, 4, 14,
	4, 58, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5,
	5, 69, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 76, 10, 7, 12, 7, 14,
	7, 79, 11, 7, 3, 8, 3, 8, 3, 8, 7, 8, 84, 10, 8, 12, 8, 14, 8, 87, 11,
	8, 3, 9, 5, 9, 90, 10, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 103, 10, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 114, 10, 10, 12, 10, 14,
	10, 117, 11, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 6, 13,
	126, 10, 13, 13, 13, 14, 13, 127, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 5, 14, 137, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 150, 10, 15, 12, 15, 14, 15,
	153, 11, 15, 3, 15, 3, 15, 5, 15, 157, 10, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 7, 16, 164, 10, 16, 12, 16, 14, 16, 167, 11, 16, 5, 16, 169,
	10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 6, 18, 178, 10,
	18, 13, 18, 14, 18, 179, 3, 18, 6, 18, 183, 10, 18, 13, 18, 14, 18, 184,
	3, 18, 3, 18, 5, 18, 189, 10, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20,
	206, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 5, 21, 217, 10, 21, 3, 21, 2, 3, 18, 22, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 2, 6, 3, 2, 33, 35, 3,
	2, 36, 37, 4, 2, 19, 19, 21, 25, 4, 2, 41, 41, 46, 46, 2, 233, 2, 42, 3,
	2, 2, 2, 4, 49, 3, 2, 2, 2, 6, 51, 3, 2, 2, 2, 8, 68, 3, 2, 2, 2, 10, 70,
	3, 2, 2, 2, 12, 72, 3, 2, 2, 2, 14, 80, 3, 2, 2, 2, 16, 89, 3, 2, 2, 2,
	18, 102, 3, 2, 2, 2, 20, 118, 3, 2, 2, 2, 22, 121, 3, 2, 2, 2, 24, 125,
	3, 2, 2, 2, 26, 136, 3, 2, 2, 2, 28, 138, 3, 2, 2, 2, 30, 158, 3, 2, 2,
	2, 32, 172, 3, 2, 2, 2, 34, 188, 3, 2, 2, 2, 36, 190, 3, 2, 2, 2, 38, 205,
	3, 2, 2, 2, 40, 216, 3, 2, 2, 2, 42, 43, 7, 12, 2, 2, 43, 44, 7, 13, 2,
	2, 44, 45, 5, 4, 3, 2, 45, 46, 7, 29, 2, 2, 46, 47, 5, 6, 4, 2, 47, 48,
	7, 2, 2, 3, 48, 3, 3, 2, 2, 2, 49, 50, 7, 41, 2, 2, 50, 5, 3, 2, 2, 2,
	51, 56, 5, 8, 5, 2, 52, 53, 7, 3, 2, 2, 53, 55, 5, 8, 5, 2, 54, 52, 3,
	2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57,
	7, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 60, 5, 18, 10, 2, 60, 61, 7, 14,
	2, 2, 61, 62, 5, 34, 18, 2, 62, 69, 3, 2, 2, 2, 63, 64, 5, 22, 12, 2, 64,
	65, 7, 38, 2, 2, 65, 66, 5, 32, 17, 2, 66, 69, 3, 2, 2, 2, 67, 69, 5, 34,
	18, 2, 68, 59, 3, 2, 2, 2, 68, 63, 3, 2, 2, 2, 68, 67, 3, 2, 2, 2, 69,
	9, 3, 2, 2, 2, 70, 71, 5, 12, 7, 2, 71, 11, 3, 2, 2, 2, 72, 77, 5, 14,
	8, 2, 73, 74, 7, 15, 2, 2, 74, 76, 5, 14, 8, 2, 75, 73, 3, 2, 2, 2, 76,
	79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 13, 3, 2, 2,
	2, 79, 77, 3, 2, 2, 2, 80, 85, 5, 16, 9, 2, 81, 82, 7, 28, 2, 2, 82, 84,
	5, 16, 9, 2, 83, 81, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2,
	85, 86, 3, 2, 2, 2, 86, 15, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 90, 7,
	26, 2, 2, 89, 88, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91,
	92, 5, 18, 10, 2, 92, 17, 3, 2, 2, 2, 93, 94, 8, 10, 1, 2, 94, 103, 5,
	26, 14, 2, 95, 96, 7, 4, 2, 2, 96, 97, 5, 18, 10, 2, 97, 98, 7, 5, 2, 2,
	98, 103, 3, 2, 2, 2, 99, 103, 5, 30, 16, 2, 100, 103, 5, 28, 15, 2, 101,
	103, 5, 20, 11, 2, 102, 93, 3, 2, 2, 2, 102, 95, 3, 2, 2, 2, 102, 99, 3,
	2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 101, 3, 2, 2, 2, 103, 115, 3, 2, 2,
	2, 104, 105, 12, 8, 2, 2, 105, 106, 9, 2, 2, 2, 106, 114, 5, 18, 10, 9,
	107, 108, 12, 7, 2, 2, 108, 109, 9, 3, 2, 2, 109, 114, 5, 18, 10, 8, 110,
	111, 12, 6, 2, 2, 111, 112, 9, 4, 2, 2, 112, 114, 5, 18, 10, 7, 113, 104,
	3, 2, 2, 2, 113, 107, 3, 2, 2, 2, 113, 110, 3, 2, 2, 2, 114, 117, 3, 2,
	2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 19, 3, 2, 2, 2,
	117, 115, 3, 2, 2, 2, 118, 119, 5, 22, 12, 2, 119, 120, 5, 24, 13, 2, 120,
	21, 3, 2, 2, 2, 121, 122, 7, 41, 2, 2, 122, 23, 3, 2, 2, 2, 123, 124, 7,
	38, 2, 2, 124, 126, 7, 41, 2, 2, 125, 123, 3, 2, 2, 2, 126, 127, 3, 2,
	2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 25, 3, 2, 2, 2,
	129, 137, 7, 39, 2, 2, 130, 137, 7, 40, 2, 2, 131, 137, 7, 42, 2, 2, 132,
	137, 7, 43, 2, 2, 133, 137, 7, 44, 2, 2, 134, 137, 7, 48, 2, 2, 135, 137,
	5, 34, 18, 2, 136, 129, 3, 2, 2, 2, 136, 130, 3, 2, 2, 2, 136, 131, 3,
	2, 2, 2, 136, 132, 3, 2, 2, 2, 136, 133, 3, 2, 2, 2, 136, 134, 3, 2, 2,
	2, 136, 135, 3, 2, 2, 2, 137, 27, 3, 2, 2, 2, 138, 139, 7, 16, 2, 2, 139,
	140, 5, 18, 10, 2, 140, 141, 7, 32, 2, 2, 141, 142, 5, 18, 10, 2, 142,
	143, 7, 30, 2, 2, 143, 151, 5, 18, 10, 2, 144, 145, 7, 32, 2, 2, 145, 146,
	5, 18, 10, 2, 146, 147, 7, 30, 2, 2, 147, 148, 5, 18, 10, 2, 148, 150,
	3, 2, 2, 2, 149, 144, 3, 2, 2, 2, 150, 153, 3, 2, 2, 2, 151, 149, 3, 2,
	2, 2, 151, 152, 3, 2, 2, 2, 152, 156, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2,
	154, 155, 7, 17, 2, 2, 155, 157, 5, 18, 10, 2, 156, 154, 3, 2, 2, 2, 156,
	157, 3, 2, 2, 2, 157, 29, 3, 2, 2, 2, 158, 159, 7, 41, 2, 2, 159, 168,
	7, 4, 2, 2, 160, 165, 5, 18, 10, 2, 161, 162, 7, 3, 2, 2, 162, 164, 5,
	18, 10, 2, 163, 161, 3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165, 163, 3, 2,
	2, 2, 165, 166, 3, 2, 2, 2, 166, 169, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2,
	168, 160, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170,
	171, 7, 5, 2, 2, 171, 31, 3, 2, 2, 2, 172, 173, 7, 33, 2, 2, 173, 33, 3,
	2, 2, 2, 174, 189, 5, 32, 17, 2, 175, 189, 5, 36, 19, 2, 176, 178, 7, 6,
	2, 2, 177, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 177, 3, 2, 2, 2,
	179, 180, 3, 2, 2, 2, 180, 182, 3, 2, 2, 2, 181, 183, 5, 36, 19, 2, 182,
	181, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 184, 185,
	3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 187, 7, 6, 2, 2, 187, 189, 3, 2,
	2, 2, 188, 174, 3, 2, 2, 2, 188, 175, 3, 2, 2, 2, 188, 177, 3, 2, 2, 2,
	189, 35, 3, 2, 2, 2, 190, 191, 9, 5, 2, 2, 191, 37, 3, 2, 2, 2, 192, 193,
	7, 46, 2, 2, 193, 194, 7, 7, 2, 2, 194, 206, 7, 8, 2, 2, 195, 196, 7, 46,
	2, 2, 196, 197, 7, 7, 2, 2, 197, 198, 7, 42, 2, 2, 198, 206, 7, 8, 2, 2,
	199, 200, 7, 46, 2, 2, 200, 201, 7, 7, 2, 2, 201, 202, 7, 9, 2, 2, 202,
	206, 7, 8, 2, 2, 203, 206, 7, 46, 2, 2, 204, 206, 7, 44, 2, 2, 205, 192,
	3, 2, 2, 2, 205, 195, 3, 2, 2, 2, 205, 199, 3, 2, 2, 2, 205, 203, 3, 2,
	2, 2, 205, 204, 3, 2, 2, 2, 206, 39, 3, 2, 2, 2, 207, 208, 7, 41, 2, 2,
	208, 217, 7, 10, 2, 2, 209, 210, 7, 41, 2, 2, 210, 211, 7, 7, 2, 2, 211,
	212, 7, 42, 2, 2, 212, 217, 7, 8, 2, 2, 213, 214, 7, 41, 2, 2, 214, 217,
	7, 11, 2, 2, 215, 217, 7, 41, 2, 2, 216, 207, 3, 2, 2, 2, 216, 209, 3,
	2, 2, 2, 216, 213, 3, 2, 2, 2, 216, 215, 3, 2, 2, 2, 217, 41, 3, 2, 2,
	2, 21, 56, 68, 77, 85, 89, 102, 113, 115, 127, 136, 151, 156, 165, 168,
	179, 184, 188, 205, 216,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'('", "')'", "'\"'", "'['", "']'", "'#'", "'[]'", "'[#]'",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "'*'", "'/'", "'%'", "'+'", "'-'", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "INSERT", "INTO", "AS", "AND",
	"CASE", "ELSE", "END", "EQ", "FROM", "GT", "GTE", "LT", "LTE", "NE", "NOT",
	"NULL", "OR", "SELECT", "THEN", "WHERE", "WHEN", "MUL", "DIV", "MOD", "ADD",
	"SUB", "DOT", "TRUE", "FALSE", "INDENTIFIER", "NUMBER", "INTEGER", "FLOAT",
	"TOPICITEM", "PATHITEM", "ARRAYITEM", "STRING", "WHITESPACE",
}

var ruleNames = []string{
	"root", "target", "fields", "field_elem", "filter", "filter_condition",
	"filter_condition_or", "filter_condition_not", "expr", "source_stmt", "sourceEntity",
	"propertyEntity", "constant", "switch_stmt", "call_expr", "asterisk", "xpath_name",
	"dotnotation", "identifierWithTOPICITEM", "identifierWithQualifier",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type QingQLParser struct {
	*antlr.BaseParser
}

func NewQingQLParser(input antlr.TokenStream) *QingQLParser {
	this := new(QingQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "QingQL.g4"

	return this
}

// QingQLParser tokens.
const (
	QingQLParserEOF         = antlr.TokenEOF
	QingQLParserT__0        = 1
	QingQLParserT__1        = 2
	QingQLParserT__2        = 3
	QingQLParserT__3        = 4
	QingQLParserT__4        = 5
	QingQLParserT__5        = 6
	QingQLParserT__6        = 7
	QingQLParserT__7        = 8
	QingQLParserT__8        = 9
	QingQLParserINSERT      = 10
	QingQLParserINTO        = 11
	QingQLParserAS          = 12
	QingQLParserAND         = 13
	QingQLParserCASE        = 14
	QingQLParserELSE        = 15
	QingQLParserEND         = 16
	QingQLParserEQ          = 17
	QingQLParserFROM        = 18
	QingQLParserGT          = 19
	QingQLParserGTE         = 20
	QingQLParserLT          = 21
	QingQLParserLTE         = 22
	QingQLParserNE          = 23
	QingQLParserNOT         = 24
	QingQLParserNULL        = 25
	QingQLParserOR          = 26
	QingQLParserSELECT      = 27
	QingQLParserTHEN        = 28
	QingQLParserWHERE       = 29
	QingQLParserWHEN        = 30
	QingQLParserMUL         = 31
	QingQLParserDIV         = 32
	QingQLParserMOD         = 33
	QingQLParserADD         = 34
	QingQLParserSUB         = 35
	QingQLParserDOT         = 36
	QingQLParserTRUE        = 37
	QingQLParserFALSE       = 38
	QingQLParserINDENTIFIER = 39
	QingQLParserNUMBER      = 40
	QingQLParserINTEGER     = 41
	QingQLParserFLOAT       = 42
	QingQLParserTOPICITEM   = 43
	QingQLParserPATHITEM    = 44
	QingQLParserARRAYITEM   = 45
	QingQLParserSTRING      = 46
	QingQLParserWHITESPACE  = 47
)

// QingQLParser rules.
const (
	QingQLParserRULE_root                    = 0
	QingQLParserRULE_target                  = 1
	QingQLParserRULE_fields                  = 2
	QingQLParserRULE_field_elem              = 3
	QingQLParserRULE_filter                  = 4
	QingQLParserRULE_filter_condition        = 5
	QingQLParserRULE_filter_condition_or     = 6
	QingQLParserRULE_filter_condition_not    = 7
	QingQLParserRULE_expr                    = 8
	QingQLParserRULE_source_stmt             = 9
	QingQLParserRULE_sourceEntity            = 10
	QingQLParserRULE_propertyEntity          = 11
	QingQLParserRULE_constant                = 12
	QingQLParserRULE_switch_stmt             = 13
	QingQLParserRULE_call_expr               = 14
	QingQLParserRULE_asterisk                = 15
	QingQLParserRULE_xpath_name              = 16
	QingQLParserRULE_dotnotation             = 17
	QingQLParserRULE_identifierWithTOPICITEM = 18
	QingQLParserRULE_identifierWithQualifier = 19
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) INSERT() antlr.TerminalNode {
	return s.GetToken(QingQLParserINSERT, 0)
}

func (s *RootContext) INTO() antlr.TerminalNode {
	return s.GetToken(QingQLParserINTO, 0)
}

func (s *RootContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *RootContext) SELECT() antlr.TerminalNode {
	return s.GetToken(QingQLParserSELECT, 0)
}

func (s *RootContext) Fields() IFieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(QingQLParserEOF, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *QingQLParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QingQLParserRULE_root)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(QingQLParserINSERT)
	}
	{
		p.SetState(41)
		p.Match(QingQLParserINTO)
	}
	{
		p.SetState(42)
		p.Target()
	}
	{
		p.SetState(43)
		p.Match(QingQLParserSELECT)
	}
	{
		p.SetState(44)
		p.Fields()
	}
	{
		p.SetState(45)
		p.Match(QingQLParserEOF)
	}

	return localctx
}

// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_target
	return p
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetContext) INDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QingQLParserINDENTIFIER, 0)
}

func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterTarget(s)
	}
}

func (s *TargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitTarget(s)
	}
}

func (p *QingQLParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QingQLParserRULE_target)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(QingQLParserINDENTIFIER)
	}

	return localctx
}

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_fields
	return p
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) AllField_elem() []IField_elemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IField_elemContext)(nil)).Elem())
	var tst = make([]IField_elemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IField_elemContext)
		}
	}

	return tst
}

func (s *FieldsContext) Field_elem(i int) IField_elemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_elemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IField_elemContext)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFields(s)
	}
}

func (s *FieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFields(s)
	}
}

func (p *QingQLParser) Fields() (localctx IFieldsContext) {
	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QingQLParserRULE_fields)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Field_elem()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QingQLParserT__0 {
		{
			p.SetState(50)
			p.Match(QingQLParserT__0)
		}
		{
			p.SetState(51)
			p.Field_elem()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IField_elemContext is an interface to support dynamic dispatch.
type IField_elemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_elemContext differentiates from other interfaces.
	IsField_elemContext()
}

type Field_elemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_elemContext() *Field_elemContext {
	var p = new(Field_elemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_field_elem
	return p
}

func (*Field_elemContext) IsField_elemContext() {}

func NewField_elemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_elemContext {
	var p = new(Field_elemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_field_elem

	return p
}

func (s *Field_elemContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_elemContext) CopyFrom(ctx *Field_elemContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Field_elemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_elemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SourceFieldContext struct {
	*Field_elemContext
}

func NewSourceFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SourceFieldContext {
	var p = new(SourceFieldContext)

	p.Field_elemContext = NewEmptyField_elemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Field_elemContext))

	return p
}

func (s *SourceFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceFieldContext) SourceEntity() ISourceEntityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceEntityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceEntityContext)
}

func (s *SourceFieldContext) Asterisk() IAsteriskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsteriskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsteriskContext)
}

func (s *SourceFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterSourceField(s)
	}
}

func (s *SourceFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitSourceField(s)
	}
}

type ExprFieldContext struct {
	*Field_elemContext
}

func NewExprFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprFieldContext {
	var p = new(ExprFieldContext)

	p.Field_elemContext = NewEmptyField_elemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Field_elemContext))

	return p
}

func (s *ExprFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprFieldContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprFieldContext) AS() antlr.TerminalNode {
	return s.GetToken(QingQLParserAS, 0)
}

func (s *ExprFieldContext) Xpath_name() IXpath_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXpath_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXpath_nameContext)
}

func (s *ExprFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterExprField(s)
	}
}

func (s *ExprFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitExprField(s)
	}
}

type XpathFieldContext struct {
	*Field_elemContext
}

func NewXpathFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *XpathFieldContext {
	var p = new(XpathFieldContext)

	p.Field_elemContext = NewEmptyField_elemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Field_elemContext))

	return p
}

func (s *XpathFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XpathFieldContext) Xpath_name() IXpath_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXpath_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXpath_nameContext)
}

func (s *XpathFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterXpathField(s)
	}
}

func (s *XpathFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitXpathField(s)
	}
}

func (p *QingQLParser) Field_elem() (localctx IField_elemContext) {
	localctx = NewField_elemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QingQLParserRULE_field_elem)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.expr(0)
		}
		{
			p.SetState(58)
			p.Match(QingQLParserAS)
		}
		{
			p.SetState(59)
			p.Xpath_name()
		}

	case 2:
		localctx = NewSourceFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.SourceEntity()
		}
		{
			p.SetState(62)
			p.Match(QingQLParserDOT)
		}
		{
			p.SetState(63)
			p.Asterisk()
		}

	case 3:
		localctx = NewXpathFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(65)
			p.Xpath_name()
		}

	}

	return localctx
}

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) Filter_condition() IFilter_conditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilter_conditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilter_conditionContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (p *QingQLParser) Filter() (localctx IFilterContext) {
	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QingQLParserRULE_filter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Filter_condition()
	}

	return localctx
}

// IFilter_conditionContext is an interface to support dynamic dispatch.
type IFilter_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilter_conditionContext differentiates from other interfaces.
	IsFilter_conditionContext()
}

type Filter_conditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilter_conditionContext() *Filter_conditionContext {
	var p = new(Filter_conditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_filter_condition
	return p
}

func (*Filter_conditionContext) IsFilter_conditionContext() {}

func NewFilter_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Filter_conditionContext {
	var p = new(Filter_conditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_filter_condition

	return p
}

func (s *Filter_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *Filter_conditionContext) AllFilter_condition_or() []IFilter_condition_orContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilter_condition_orContext)(nil)).Elem())
	var tst = make([]IFilter_condition_orContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilter_condition_orContext)
		}
	}

	return tst
}

func (s *Filter_conditionContext) Filter_condition_or(i int) IFilter_condition_orContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilter_condition_orContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilter_condition_orContext)
}

func (s *Filter_conditionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(QingQLParserAND)
}

func (s *Filter_conditionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(QingQLParserAND, i)
}

func (s *Filter_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Filter_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Filter_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFilter_condition(s)
	}
}

func (s *Filter_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFilter_condition(s)
	}
}

func (p *QingQLParser) Filter_condition() (localctx IFilter_conditionContext) {
	localctx = NewFilter_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, QingQLParserRULE_filter_condition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Filter_condition_or()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QingQLParserAND {
		{
			p.SetState(71)
			p.Match(QingQLParserAND)
		}
		{
			p.SetState(72)
			p.Filter_condition_or()
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFilter_condition_orContext is an interface to support dynamic dispatch.
type IFilter_condition_orContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilter_condition_orContext differentiates from other interfaces.
	IsFilter_condition_orContext()
}

type Filter_condition_orContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilter_condition_orContext() *Filter_condition_orContext {
	var p = new(Filter_condition_orContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_filter_condition_or
	return p
}

func (*Filter_condition_orContext) IsFilter_condition_orContext() {}

func NewFilter_condition_orContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Filter_condition_orContext {
	var p = new(Filter_condition_orContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_filter_condition_or

	return p
}

func (s *Filter_condition_orContext) GetParser() antlr.Parser { return s.parser }

func (s *Filter_condition_orContext) AllFilter_condition_not() []IFilter_condition_notContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilter_condition_notContext)(nil)).Elem())
	var tst = make([]IFilter_condition_notContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilter_condition_notContext)
		}
	}

	return tst
}

func (s *Filter_condition_orContext) Filter_condition_not(i int) IFilter_condition_notContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilter_condition_notContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilter_condition_notContext)
}

func (s *Filter_condition_orContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(QingQLParserOR)
}

func (s *Filter_condition_orContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(QingQLParserOR, i)
}

func (s *Filter_condition_orContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Filter_condition_orContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Filter_condition_orContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFilter_condition_or(s)
	}
}

func (s *Filter_condition_orContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFilter_condition_or(s)
	}
}

func (p *QingQLParser) Filter_condition_or() (localctx IFilter_condition_orContext) {
	localctx = NewFilter_condition_orContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, QingQLParserRULE_filter_condition_or)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Filter_condition_not()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QingQLParserOR {
		{
			p.SetState(79)
			p.Match(QingQLParserOR)
		}
		{
			p.SetState(80)
			p.Filter_condition_not()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFilter_condition_notContext is an interface to support dynamic dispatch.
type IFilter_condition_notContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilter_condition_notContext differentiates from other interfaces.
	IsFilter_condition_notContext()
}

type Filter_condition_notContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilter_condition_notContext() *Filter_condition_notContext {
	var p = new(Filter_condition_notContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_filter_condition_not
	return p
}

func (*Filter_condition_notContext) IsFilter_condition_notContext() {}

func NewFilter_condition_notContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Filter_condition_notContext {
	var p = new(Filter_condition_notContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_filter_condition_not

	return p
}

func (s *Filter_condition_notContext) GetParser() antlr.Parser { return s.parser }

func (s *Filter_condition_notContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Filter_condition_notContext) NOT() antlr.TerminalNode {
	return s.GetToken(QingQLParserNOT, 0)
}

func (s *Filter_condition_notContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Filter_condition_notContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Filter_condition_notContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFilter_condition_not(s)
	}
}

func (s *Filter_condition_notContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFilter_condition_not(s)
	}
}

func (p *QingQLParser) Filter_condition_not() (localctx IFilter_condition_notContext) {
	localctx = NewFilter_condition_notContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, QingQLParserRULE_filter_condition_not)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QingQLParserNOT {
		{
			p.SetState(86)
			p.Match(QingQLParserNOT)
		}

	}
	{
		p.SetState(89)
		p.expr(0)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionContext struct {
	*ExprContext
}

func NewFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionContext {
	var p = new(FunctionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) Call_expr() ICall_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_exprContext)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFunction(s)
	}
}

type BracesContext struct {
	*ExprContext
}

func NewBracesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracesContext {
	var p = new(BracesContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BracesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracesContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *BracesContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BracesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterBraces(s)
	}
}

func (s *BracesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitBraces(s)
	}
}

type SwitchContext struct {
	*ExprContext
}

func NewSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchContext {
	var p = new(SwitchContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) Switch_stmt() ISwitch_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitch_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitch_stmtContext)
}

func (s *SwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterSwitch(s)
	}
}

func (s *SwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitSwitch(s)
	}
}

type BinaryContext struct {
	*ExprContext
	op antlr.Token
}

func NewBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryContext {
	var p = new(BinaryContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BinaryContext) GetOp() antlr.Token { return s.op }

func (s *BinaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *BinaryContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BinaryContext) EQ() antlr.TerminalNode {
	return s.GetToken(QingQLParserEQ, 0)
}

func (s *BinaryContext) GT() antlr.TerminalNode {
	return s.GetToken(QingQLParserGT, 0)
}

func (s *BinaryContext) LT() antlr.TerminalNode {
	return s.GetToken(QingQLParserLT, 0)
}

func (s *BinaryContext) GTE() antlr.TerminalNode {
	return s.GetToken(QingQLParserGTE, 0)
}

func (s *BinaryContext) LTE() antlr.TerminalNode {
	return s.GetToken(QingQLParserLTE, 0)
}

func (s *BinaryContext) NE() antlr.TerminalNode {
	return s.GetToken(QingQLParserNE, 0)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitBinary(s)
	}
}

type SourceContext struct {
	*ExprContext
}

func NewSourceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SourceContext {
	var p = new(SourceContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceContext) Source_stmt() ISource_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISource_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISource_stmtContext)
}

func (s *SourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterSource(s)
	}
}

func (s *SourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitSource(s)
	}
}

func (p *QingQLParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *QingQLParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, QingQLParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBracesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(92)
			p.Constant()
		}

	case 2:
		localctx = NewBracesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(93)
			p.Match(QingQLParserT__1)
		}
		{
			p.SetState(94)
			p.expr(0)
		}
		{
			p.SetState(95)
			p.Match(QingQLParserT__2)
		}

	case 3:
		localctx = NewFunctionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(97)
			p.Call_expr()
		}

	case 4:
		localctx = NewSwitchContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(98)
			p.Switch_stmt()
		}

	case 5:
		localctx = NewSourceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(99)
			p.Source_stmt()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QingQLParserRULE_expr)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				p.SetState(103)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*BinaryContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(QingQLParserMUL-31))|(1<<(QingQLParserDIV-31))|(1<<(QingQLParserMOD-31)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*BinaryContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(104)
					p.expr(7)
				}

			case 2:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QingQLParserRULE_expr)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(106)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*BinaryContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == QingQLParserADD || _la == QingQLParserSUB) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*BinaryContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(107)
					p.expr(6)
				}

			case 3:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QingQLParserRULE_expr)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				p.SetState(109)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*BinaryContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QingQLParserEQ)|(1<<QingQLParserGT)|(1<<QingQLParserGTE)|(1<<QingQLParserLT)|(1<<QingQLParserLTE)|(1<<QingQLParserNE))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*BinaryContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(110)
					p.expr(5)
				}

			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// ISource_stmtContext is an interface to support dynamic dispatch.
type ISource_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSource_stmtContext differentiates from other interfaces.
	IsSource_stmtContext()
}

type Source_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySource_stmtContext() *Source_stmtContext {
	var p = new(Source_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_source_stmt
	return p
}

func (*Source_stmtContext) IsSource_stmtContext() {}

func NewSource_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Source_stmtContext {
	var p = new(Source_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_source_stmt

	return p
}

func (s *Source_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Source_stmtContext) SourceEntity() ISourceEntityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceEntityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceEntityContext)
}

func (s *Source_stmtContext) PropertyEntity() IPropertyEntityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyEntityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyEntityContext)
}

func (s *Source_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Source_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Source_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterSource_stmt(s)
	}
}

func (s *Source_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitSource_stmt(s)
	}
}

func (p *QingQLParser) Source_stmt() (localctx ISource_stmtContext) {
	localctx = NewSource_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, QingQLParserRULE_source_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.SourceEntity()
	}
	{
		p.SetState(117)
		p.PropertyEntity()
	}

	return localctx
}

// ISourceEntityContext is an interface to support dynamic dispatch.
type ISourceEntityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceEntityContext differentiates from other interfaces.
	IsSourceEntityContext()
}

type SourceEntityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceEntityContext() *SourceEntityContext {
	var p = new(SourceEntityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_sourceEntity
	return p
}

func (*SourceEntityContext) IsSourceEntityContext() {}

func NewSourceEntityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceEntityContext {
	var p = new(SourceEntityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_sourceEntity

	return p
}

func (s *SourceEntityContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceEntityContext) INDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QingQLParserINDENTIFIER, 0)
}

func (s *SourceEntityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceEntityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceEntityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterSourceEntity(s)
	}
}

func (s *SourceEntityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitSourceEntity(s)
	}
}

func (p *QingQLParser) SourceEntity() (localctx ISourceEntityContext) {
	localctx = NewSourceEntityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, QingQLParserRULE_sourceEntity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(QingQLParserINDENTIFIER)
	}

	return localctx
}

// IPropertyEntityContext is an interface to support dynamic dispatch.
type IPropertyEntityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyEntityContext differentiates from other interfaces.
	IsPropertyEntityContext()
}

type PropertyEntityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyEntityContext() *PropertyEntityContext {
	var p = new(PropertyEntityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_propertyEntity
	return p
}

func (*PropertyEntityContext) IsPropertyEntityContext() {}

func NewPropertyEntityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyEntityContext {
	var p = new(PropertyEntityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_propertyEntity

	return p
}

func (s *PropertyEntityContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyEntityContext) AllINDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(QingQLParserINDENTIFIER)
}

func (s *PropertyEntityContext) INDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(QingQLParserINDENTIFIER, i)
}

func (s *PropertyEntityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyEntityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyEntityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterPropertyEntity(s)
	}
}

func (s *PropertyEntityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitPropertyEntity(s)
	}
}

func (p *QingQLParser) PropertyEntity() (localctx IPropertyEntityContext) {
	localctx = NewPropertyEntityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, QingQLParserRULE_propertyEntity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(121)
				p.Match(QingQLParserDOT)
			}
			{
				p.SetState(122)
				p.Match(QingQLParserINDENTIFIER)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) CopyFrom(ctx *ConstantContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IntegerContext struct {
	*ConstantContext
}

func NewIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerContext {
	var p = new(IntegerContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(QingQLParserNUMBER, 0)
}

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(QingQLParserINTEGER, 0)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitInteger(s)
	}
}

type XPathContext struct {
	*ConstantContext
}

func NewXPathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *XPathContext {
	var p = new(XPathContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *XPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XPathContext) Xpath_name() IXpath_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXpath_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXpath_nameContext)
}

func (s *XPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterXPath(s)
	}
}

func (s *XPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitXPath(s)
	}
}

type FloatContext struct {
	*ConstantContext
}

func NewFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatContext {
	var p = new(FloatContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(QingQLParserFLOAT, 0)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFloat(s)
	}
}

type StringContext struct {
	*ConstantContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(QingQLParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitString(s)
	}
}

type BooleanContext struct {
	*ConstantContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) TRUE() antlr.TerminalNode {
	return s.GetToken(QingQLParserTRUE, 0)
}

func (s *BooleanContext) FALSE() antlr.TerminalNode {
	return s.GetToken(QingQLParserFALSE, 0)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (p *QingQLParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, QingQLParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QingQLParserTRUE:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Match(QingQLParserTRUE)
		}

	case QingQLParserFALSE:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Match(QingQLParserFALSE)
		}

	case QingQLParserNUMBER:
		localctx = NewIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.Match(QingQLParserNUMBER)
		}

	case QingQLParserINTEGER:
		localctx = NewIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)
			p.Match(QingQLParserINTEGER)
		}

	case QingQLParserFLOAT:
		localctx = NewFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(131)
			p.Match(QingQLParserFLOAT)
		}

	case QingQLParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(132)
			p.Match(QingQLParserSTRING)
		}

	case QingQLParserT__3, QingQLParserMUL, QingQLParserINDENTIFIER, QingQLParserPATHITEM:
		localctx = NewXPathContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(133)
			p.Xpath_name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISwitch_stmtContext is an interface to support dynamic dispatch.
type ISwitch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitch_stmtContext differentiates from other interfaces.
	IsSwitch_stmtContext()
}

type Switch_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_stmtContext() *Switch_stmtContext {
	var p = new(Switch_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_switch_stmt
	return p
}

func (*Switch_stmtContext) IsSwitch_stmtContext() {}

func NewSwitch_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_stmtContext {
	var p = new(Switch_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_switch_stmt

	return p
}

func (s *Switch_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_stmtContext) CASE() antlr.TerminalNode {
	return s.GetToken(QingQLParserCASE, 0)
}

func (s *Switch_stmtContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Switch_stmtContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Switch_stmtContext) AllWHEN() []antlr.TerminalNode {
	return s.GetTokens(QingQLParserWHEN)
}

func (s *Switch_stmtContext) WHEN(i int) antlr.TerminalNode {
	return s.GetToken(QingQLParserWHEN, i)
}

func (s *Switch_stmtContext) AllTHEN() []antlr.TerminalNode {
	return s.GetTokens(QingQLParserTHEN)
}

func (s *Switch_stmtContext) THEN(i int) antlr.TerminalNode {
	return s.GetToken(QingQLParserTHEN, i)
}

func (s *Switch_stmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(QingQLParserELSE, 0)
}

func (s *Switch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterSwitch_stmt(s)
	}
}

func (s *Switch_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitSwitch_stmt(s)
	}
}

func (p *QingQLParser) Switch_stmt() (localctx ISwitch_stmtContext) {
	localctx = NewSwitch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, QingQLParserRULE_switch_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(QingQLParserCASE)
	}
	{
		p.SetState(137)
		p.expr(0)
	}
	{
		p.SetState(138)
		p.Match(QingQLParserWHEN)
	}
	{
		p.SetState(139)
		p.expr(0)
	}
	{
		p.SetState(140)
		p.Match(QingQLParserTHEN)
	}
	{
		p.SetState(141)
		p.expr(0)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(142)
				p.Match(QingQLParserWHEN)
			}
			{
				p.SetState(143)
				p.expr(0)
			}
			{
				p.SetState(144)
				p.Match(QingQLParserTHEN)
			}
			{
				p.SetState(145)
				p.expr(0)
			}

		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(152)
			p.Match(QingQLParserELSE)
		}
		{
			p.SetState(153)
			p.expr(0)
		}

	}

	return localctx
}

// ICall_exprContext is an interface to support dynamic dispatch.
type ICall_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// IsCall_exprContext differentiates from other interfaces.
	IsCall_exprContext()
}

type Call_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
}

func NewEmptyCall_exprContext() *Call_exprContext {
	var p = new(Call_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_call_expr
	return p
}

func (*Call_exprContext) IsCall_exprContext() {}

func NewCall_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_exprContext {
	var p = new(Call_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_call_expr

	return p
}

func (s *Call_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_exprContext) GetKey() antlr.Token { return s.key }

func (s *Call_exprContext) SetKey(v antlr.Token) { s.key = v }

func (s *Call_exprContext) INDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QingQLParserINDENTIFIER, 0)
}

func (s *Call_exprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Call_exprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Call_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterCall_expr(s)
	}
}

func (s *Call_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitCall_expr(s)
	}
}

func (p *QingQLParser) Call_expr() (localctx ICall_exprContext) {
	localctx = NewCall_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, QingQLParserRULE_call_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)

		var _m = p.Match(QingQLParserINDENTIFIER)

		localctx.(*Call_exprContext).key = _m
	}
	{
		p.SetState(157)
		p.Match(QingQLParserT__1)
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QingQLParserT__1)|(1<<QingQLParserT__3)|(1<<QingQLParserCASE)|(1<<QingQLParserMUL))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(QingQLParserTRUE-37))|(1<<(QingQLParserFALSE-37))|(1<<(QingQLParserINDENTIFIER-37))|(1<<(QingQLParserNUMBER-37))|(1<<(QingQLParserINTEGER-37))|(1<<(QingQLParserFLOAT-37))|(1<<(QingQLParserPATHITEM-37))|(1<<(QingQLParserSTRING-37)))) != 0) {
		{
			p.SetState(158)
			p.expr(0)
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QingQLParserT__0 {
			{
				p.SetState(159)
				p.Match(QingQLParserT__0)
			}
			{
				p.SetState(160)
				p.expr(0)
			}

			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(168)
		p.Match(QingQLParserT__2)
	}

	return localctx
}

// IAsteriskContext is an interface to support dynamic dispatch.
type IAsteriskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAsteriskContext differentiates from other interfaces.
	IsAsteriskContext()
}

type AsteriskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsteriskContext() *AsteriskContext {
	var p = new(AsteriskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_asterisk
	return p
}

func (*AsteriskContext) IsAsteriskContext() {}

func NewAsteriskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsteriskContext {
	var p = new(AsteriskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_asterisk

	return p
}

func (s *AsteriskContext) GetParser() antlr.Parser { return s.parser }
func (s *AsteriskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsteriskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsteriskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterAsterisk(s)
	}
}

func (s *AsteriskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitAsterisk(s)
	}
}

func (p *QingQLParser) Asterisk() (localctx IAsteriskContext) {
	localctx = NewAsteriskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, QingQLParserRULE_asterisk)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(QingQLParserMUL)
	}

	return localctx
}

// IXpath_nameContext is an interface to support dynamic dispatch.
type IXpath_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXpath_nameContext differentiates from other interfaces.
	IsXpath_nameContext()
}

type Xpath_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXpath_nameContext() *Xpath_nameContext {
	var p = new(Xpath_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_xpath_name
	return p
}

func (*Xpath_nameContext) IsXpath_nameContext() {}

func NewXpath_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Xpath_nameContext {
	var p = new(Xpath_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_xpath_name

	return p
}

func (s *Xpath_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Xpath_nameContext) Asterisk() IAsteriskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsteriskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsteriskContext)
}

func (s *Xpath_nameContext) AllDotnotation() []IDotnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDotnotationContext)(nil)).Elem())
	var tst = make([]IDotnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDotnotationContext)
		}
	}

	return tst
}

func (s *Xpath_nameContext) Dotnotation(i int) IDotnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDotnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDotnotationContext)
}

func (s *Xpath_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Xpath_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Xpath_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterXpath_name(s)
	}
}

func (s *Xpath_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitXpath_name(s)
	}
}

func (p *QingQLParser) Xpath_name() (localctx IXpath_nameContext) {
	localctx = NewXpath_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, QingQLParserRULE_xpath_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(186)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QingQLParserMUL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.Asterisk()
		}

	case QingQLParserINDENTIFIER, QingQLParserPATHITEM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(173)
			p.Dotnotation()
		}

	case QingQLParserT__3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QingQLParserT__3 {
			{
				p.SetState(174)
				p.Match(QingQLParserT__3)
			}

			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QingQLParserINDENTIFIER || _la == QingQLParserPATHITEM {
			{
				p.SetState(179)
				p.Dotnotation()
			}

			p.SetState(182)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(184)
			p.Match(QingQLParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDotnotationContext is an interface to support dynamic dispatch.
type IDotnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDotnotationContext differentiates from other interfaces.
	IsDotnotationContext()
}

type DotnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDotnotationContext() *DotnotationContext {
	var p = new(DotnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_dotnotation
	return p
}

func (*DotnotationContext) IsDotnotationContext() {}

func NewDotnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DotnotationContext {
	var p = new(DotnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_dotnotation

	return p
}

func (s *DotnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *DotnotationContext) INDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QingQLParserINDENTIFIER, 0)
}

func (s *DotnotationContext) PATHITEM() antlr.TerminalNode {
	return s.GetToken(QingQLParserPATHITEM, 0)
}

func (s *DotnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DotnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterDotnotation(s)
	}
}

func (s *DotnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitDotnotation(s)
	}
}

func (p *QingQLParser) Dotnotation() (localctx IDotnotationContext) {
	localctx = NewDotnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, QingQLParserRULE_dotnotation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(188)
	_la = p.GetTokenStream().LA(1)

	if !(_la == QingQLParserINDENTIFIER || _la == QingQLParserPATHITEM) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IIdentifierWithTOPICITEMContext is an interface to support dynamic dispatch.
type IIdentifierWithTOPICITEMContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierWithTOPICITEMContext differentiates from other interfaces.
	IsIdentifierWithTOPICITEMContext()
}

type IdentifierWithTOPICITEMContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierWithTOPICITEMContext() *IdentifierWithTOPICITEMContext {
	var p = new(IdentifierWithTOPICITEMContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_identifierWithTOPICITEM
	return p
}

func (*IdentifierWithTOPICITEMContext) IsIdentifierWithTOPICITEMContext() {}

func NewIdentifierWithTOPICITEMContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierWithTOPICITEMContext {
	var p = new(IdentifierWithTOPICITEMContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_identifierWithTOPICITEM

	return p
}

func (s *IdentifierWithTOPICITEMContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierWithTOPICITEMContext) PATHITEM() antlr.TerminalNode {
	return s.GetToken(QingQLParserPATHITEM, 0)
}

func (s *IdentifierWithTOPICITEMContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(QingQLParserNUMBER, 0)
}

func (s *IdentifierWithTOPICITEMContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(QingQLParserFLOAT, 0)
}

func (s *IdentifierWithTOPICITEMContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierWithTOPICITEMContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierWithTOPICITEMContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterIdentifierWithTOPICITEM(s)
	}
}

func (s *IdentifierWithTOPICITEMContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitIdentifierWithTOPICITEM(s)
	}
}

func (p *QingQLParser) IdentifierWithTOPICITEM() (localctx IIdentifierWithTOPICITEMContext) {
	localctx = NewIdentifierWithTOPICITEMContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, QingQLParserRULE_identifierWithTOPICITEM)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(190)
			p.Match(QingQLParserPATHITEM)
		}
		{
			p.SetState(191)
			p.Match(QingQLParserT__4)
		}
		{
			p.SetState(192)
			p.Match(QingQLParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.Match(QingQLParserPATHITEM)
		}
		{
			p.SetState(194)
			p.Match(QingQLParserT__4)
		}
		{
			p.SetState(195)
			p.Match(QingQLParserNUMBER)
		}
		{
			p.SetState(196)
			p.Match(QingQLParserT__5)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(197)
			p.Match(QingQLParserPATHITEM)
		}
		{
			p.SetState(198)
			p.Match(QingQLParserT__4)
		}
		{
			p.SetState(199)
			p.Match(QingQLParserT__6)
		}
		{
			p.SetState(200)
			p.Match(QingQLParserT__5)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(201)
			p.Match(QingQLParserPATHITEM)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(202)
			p.Match(QingQLParserFLOAT)
		}

	}

	return localctx
}

// IIdentifierWithQualifierContext is an interface to support dynamic dispatch.
type IIdentifierWithQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierWithQualifierContext differentiates from other interfaces.
	IsIdentifierWithQualifierContext()
}

type IdentifierWithQualifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierWithQualifierContext() *IdentifierWithQualifierContext {
	var p = new(IdentifierWithQualifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_identifierWithQualifier
	return p
}

func (*IdentifierWithQualifierContext) IsIdentifierWithQualifierContext() {}

func NewIdentifierWithQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierWithQualifierContext {
	var p = new(IdentifierWithQualifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_identifierWithQualifier

	return p
}

func (s *IdentifierWithQualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierWithQualifierContext) INDENTIFIER() antlr.TerminalNode {
	return s.GetToken(QingQLParserINDENTIFIER, 0)
}

func (s *IdentifierWithQualifierContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(QingQLParserNUMBER, 0)
}

func (s *IdentifierWithQualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierWithQualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierWithQualifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterIdentifierWithQualifier(s)
	}
}

func (s *IdentifierWithQualifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitIdentifierWithQualifier(s)
	}
}

func (p *QingQLParser) IdentifierWithQualifier() (localctx IIdentifierWithQualifierContext) {
	localctx = NewIdentifierWithQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, QingQLParserRULE_identifierWithQualifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Match(QingQLParserINDENTIFIER)
		}
		{
			p.SetState(206)
			p.Match(QingQLParserT__7)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
			p.Match(QingQLParserINDENTIFIER)
		}
		{
			p.SetState(208)
			p.Match(QingQLParserT__4)
		}
		{
			p.SetState(209)
			p.Match(QingQLParserNUMBER)
		}
		{
			p.SetState(210)
			p.Match(QingQLParserT__5)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(211)
			p.Match(QingQLParserINDENTIFIER)
		}
		{
			p.SetState(212)
			p.Match(QingQLParserT__8)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(213)
			p.Match(QingQLParserINDENTIFIER)
		}

	}

	return localctx
}

func (p *QingQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *QingQLParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
