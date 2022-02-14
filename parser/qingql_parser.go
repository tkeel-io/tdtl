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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 232,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 57, 10,
	4, 12, 4, 14, 4, 60, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 68,
	10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 79,
	10, 8, 12, 8, 14, 8, 82, 11, 8, 3, 9, 3, 9, 3, 9, 7, 9, 87, 10, 9, 12,
	9, 14, 9, 90, 11, 9, 3, 10, 5, 10, 93, 10, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 105, 10, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 116, 10,
	11, 12, 11, 14, 11, 119, 11, 11, 3, 12, 3, 12, 3, 13, 3, 13, 6, 13, 125,
	10, 13, 13, 13, 14, 13, 126, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 5, 14, 136, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 149, 10, 15, 12, 15, 14, 15, 152,
	11, 15, 3, 15, 3, 15, 5, 15, 156, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 7, 16, 163, 10, 16, 12, 16, 14, 16, 166, 11, 16, 5, 16, 168, 10, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 6, 18, 176, 10, 18, 13, 18, 14,
	18, 177, 3, 18, 6, 18, 181, 10, 18, 13, 18, 14, 18, 182, 3, 18, 3, 18,
	5, 18, 187, 10, 18, 3, 19, 3, 19, 6, 19, 191, 10, 19, 13, 19, 14, 19, 192,
	3, 19, 6, 19, 196, 10, 19, 13, 19, 14, 19, 197, 3, 19, 3, 19, 5, 19, 202,
	10, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 219, 10, 21, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 230, 10, 22,
	3, 22, 2, 3, 20, 23, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 2, 6, 3, 2, 33, 35, 3, 2, 36, 37, 4, 2, 19,
	19, 21, 25, 4, 2, 41, 41, 46, 46, 2, 246, 2, 44, 3, 2, 2, 2, 4, 51, 3,
	2, 2, 2, 6, 53, 3, 2, 2, 2, 8, 67, 3, 2, 2, 2, 10, 69, 3, 2, 2, 2, 12,
	73, 3, 2, 2, 2, 14, 75, 3, 2, 2, 2, 16, 83, 3, 2, 2, 2, 18, 92, 3, 2, 2,
	2, 20, 104, 3, 2, 2, 2, 22, 120, 3, 2, 2, 2, 24, 124, 3, 2, 2, 2, 26, 135,
	3, 2, 2, 2, 28, 137, 3, 2, 2, 2, 30, 157, 3, 2, 2, 2, 32, 171, 3, 2, 2,
	2, 34, 186, 3, 2, 2, 2, 36, 201, 3, 2, 2, 2, 38, 203, 3, 2, 2, 2, 40, 218,
	3, 2, 2, 2, 42, 229, 3, 2, 2, 2, 44, 45, 7, 12, 2, 2, 45, 46, 7, 13, 2,
	2, 46, 47, 5, 4, 3, 2, 47, 48, 7, 29, 2, 2, 48, 49, 5, 6, 4, 2, 49, 50,
	7, 2, 2, 3, 50, 3, 3, 2, 2, 2, 51, 52, 7, 41, 2, 2, 52, 5, 3, 2, 2, 2,
	53, 58, 5, 8, 5, 2, 54, 55, 7, 3, 2, 2, 55, 57, 5, 8, 5, 2, 56, 54, 3,
	2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59,
	7, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 68, 5, 10, 6, 2, 62, 63, 5, 22,
	12, 2, 63, 64, 7, 38, 2, 2, 64, 65, 5, 32, 17, 2, 65, 68, 3, 2, 2, 2, 66,
	68, 5, 20, 11, 2, 67, 61, 3, 2, 2, 2, 67, 62, 3, 2, 2, 2, 67, 66, 3, 2,
	2, 2, 68, 9, 3, 2, 2, 2, 69, 70, 5, 20, 11, 2, 70, 71, 7, 14, 2, 2, 71,
	72, 5, 36, 19, 2, 72, 11, 3, 2, 2, 2, 73, 74, 5, 14, 8, 2, 74, 13, 3, 2,
	2, 2, 75, 80, 5, 16, 9, 2, 76, 77, 7, 15, 2, 2, 77, 79, 5, 16, 9, 2, 78,
	76, 3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2,
	2, 81, 15, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 83, 88, 5, 18, 10, 2, 84, 85,
	7, 28, 2, 2, 85, 87, 5, 18, 10, 2, 86, 84, 3, 2, 2, 2, 87, 90, 3, 2, 2,
	2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 17, 3, 2, 2, 2, 90, 88,
	3, 2, 2, 2, 91, 93, 7, 26, 2, 2, 92, 91, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2,
	93, 94, 3, 2, 2, 2, 94, 95, 5, 20, 11, 2, 95, 19, 3, 2, 2, 2, 96, 97, 8,
	11, 1, 2, 97, 105, 5, 26, 14, 2, 98, 99, 7, 4, 2, 2, 99, 100, 5, 20, 11,
	2, 100, 101, 7, 5, 2, 2, 101, 105, 3, 2, 2, 2, 102, 105, 5, 30, 16, 2,
	103, 105, 5, 28, 15, 2, 104, 96, 3, 2, 2, 2, 104, 98, 3, 2, 2, 2, 104,
	102, 3, 2, 2, 2, 104, 103, 3, 2, 2, 2, 105, 117, 3, 2, 2, 2, 106, 107,
	12, 7, 2, 2, 107, 108, 9, 2, 2, 2, 108, 116, 5, 20, 11, 8, 109, 110, 12,
	6, 2, 2, 110, 111, 9, 3, 2, 2, 111, 116, 5, 20, 11, 7, 112, 113, 12, 5,
	2, 2, 113, 114, 9, 4, 2, 2, 114, 116, 5, 20, 11, 6, 115, 106, 3, 2, 2,
	2, 115, 109, 3, 2, 2, 2, 115, 112, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2, 117,
	115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 21, 3, 2, 2, 2, 119, 117, 3,
	2, 2, 2, 120, 121, 7, 41, 2, 2, 121, 23, 3, 2, 2, 2, 122, 123, 7, 38, 2,
	2, 123, 125, 7, 41, 2, 2, 124, 122, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126,
	124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 25, 3, 2, 2, 2, 128, 136, 7,
	39, 2, 2, 129, 136, 7, 40, 2, 2, 130, 136, 7, 42, 2, 2, 131, 136, 7, 43,
	2, 2, 132, 136, 7, 44, 2, 2, 133, 136, 7, 48, 2, 2, 134, 136, 5, 34, 18,
	2, 135, 128, 3, 2, 2, 2, 135, 129, 3, 2, 2, 2, 135, 130, 3, 2, 2, 2, 135,
	131, 3, 2, 2, 2, 135, 132, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135, 134,
	3, 2, 2, 2, 136, 27, 3, 2, 2, 2, 137, 138, 7, 16, 2, 2, 138, 139, 5, 20,
	11, 2, 139, 140, 7, 32, 2, 2, 140, 141, 5, 20, 11, 2, 141, 142, 7, 30,
	2, 2, 142, 150, 5, 20, 11, 2, 143, 144, 7, 32, 2, 2, 144, 145, 5, 20, 11,
	2, 145, 146, 7, 30, 2, 2, 146, 147, 5, 20, 11, 2, 147, 149, 3, 2, 2, 2,
	148, 143, 3, 2, 2, 2, 149, 152, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 150,
	151, 3, 2, 2, 2, 151, 155, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 153, 154,
	7, 17, 2, 2, 154, 156, 5, 20, 11, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3,
	2, 2, 2, 156, 29, 3, 2, 2, 2, 157, 158, 7, 41, 2, 2, 158, 167, 7, 4, 2,
	2, 159, 164, 5, 20, 11, 2, 160, 161, 7, 3, 2, 2, 161, 163, 5, 20, 11, 2,
	162, 160, 3, 2, 2, 2, 163, 166, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 164,
	165, 3, 2, 2, 2, 165, 168, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 167, 159,
	3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 7, 5,
	2, 2, 170, 31, 3, 2, 2, 2, 171, 172, 7, 33, 2, 2, 172, 33, 3, 2, 2, 2,
	173, 187, 5, 38, 20, 2, 174, 176, 7, 6, 2, 2, 175, 174, 3, 2, 2, 2, 176,
	177, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 180,
	3, 2, 2, 2, 179, 181, 5, 38, 20, 2, 180, 179, 3, 2, 2, 2, 181, 182, 3,
	2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 184, 3, 2, 2,
	2, 184, 185, 7, 6, 2, 2, 185, 187, 3, 2, 2, 2, 186, 173, 3, 2, 2, 2, 186,
	175, 3, 2, 2, 2, 187, 35, 3, 2, 2, 2, 188, 202, 5, 38, 20, 2, 189, 191,
	7, 6, 2, 2, 190, 189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 190, 3, 2,
	2, 2, 192, 193, 3, 2, 2, 2, 193, 195, 3, 2, 2, 2, 194, 196, 5, 38, 20,
	2, 195, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 197,
	198, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 200, 7, 6, 2, 2, 200, 202,
	3, 2, 2, 2, 201, 188, 3, 2, 2, 2, 201, 190, 3, 2, 2, 2, 202, 37, 3, 2,
	2, 2, 203, 204, 9, 5, 2, 2, 204, 39, 3, 2, 2, 2, 205, 206, 7, 46, 2, 2,
	206, 207, 7, 7, 2, 2, 207, 219, 7, 8, 2, 2, 208, 209, 7, 46, 2, 2, 209,
	210, 7, 7, 2, 2, 210, 211, 7, 42, 2, 2, 211, 219, 7, 8, 2, 2, 212, 213,
	7, 46, 2, 2, 213, 214, 7, 7, 2, 2, 214, 215, 7, 9, 2, 2, 215, 219, 7, 8,
	2, 2, 216, 219, 7, 46, 2, 2, 217, 219, 7, 44, 2, 2, 218, 205, 3, 2, 2,
	2, 218, 208, 3, 2, 2, 2, 218, 212, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218,
	217, 3, 2, 2, 2, 219, 41, 3, 2, 2, 2, 220, 221, 7, 41, 2, 2, 221, 230,
	7, 10, 2, 2, 222, 223, 7, 41, 2, 2, 223, 224, 7, 7, 2, 2, 224, 225, 7,
	42, 2, 2, 225, 230, 7, 8, 2, 2, 226, 227, 7, 41, 2, 2, 227, 230, 7, 11,
	2, 2, 228, 230, 7, 41, 2, 2, 229, 220, 3, 2, 2, 2, 229, 222, 3, 2, 2, 2,
	229, 226, 3, 2, 2, 2, 229, 228, 3, 2, 2, 2, 230, 43, 3, 2, 2, 2, 24, 58,
	67, 80, 88, 92, 104, 115, 117, 126, 135, 150, 155, 164, 167, 177, 182,
	186, 192, 197, 201, 218, 229,
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
	"root", "target", "fields", "field_elem", "field_elem_with_as", "filter",
	"filter_condition", "filter_condition_or", "filter_condition_not", "expr",
	"sourceEntity", "propertyEntity", "constant", "switch_stmt", "call_expr",
	"asterisk", "xpath_name", "target_name", "dotnotation", "identifierWithTOPICITEM",
	"identifierWithQualifier",
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
	QingQLParserRULE_field_elem_with_as      = 4
	QingQLParserRULE_filter                  = 5
	QingQLParserRULE_filter_condition        = 6
	QingQLParserRULE_filter_condition_or     = 7
	QingQLParserRULE_filter_condition_not    = 8
	QingQLParserRULE_expr                    = 9
	QingQLParserRULE_sourceEntity            = 10
	QingQLParserRULE_propertyEntity          = 11
	QingQLParserRULE_constant                = 12
	QingQLParserRULE_switch_stmt             = 13
	QingQLParserRULE_call_expr               = 14
	QingQLParserRULE_asterisk                = 15
	QingQLParserRULE_xpath_name              = 16
	QingQLParserRULE_target_name             = 17
	QingQLParserRULE_dotnotation             = 18
	QingQLParserRULE_identifierWithTOPICITEM = 19
	QingQLParserRULE_identifierWithQualifier = 20
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
		p.SetState(42)
		p.Match(QingQLParserINSERT)
	}
	{
		p.SetState(43)
		p.Match(QingQLParserINTO)
	}
	{
		p.SetState(44)
		p.Target()
	}
	{
		p.SetState(45)
		p.Match(QingQLParserSELECT)
	}
	{
		p.SetState(46)
		p.Fields()
	}
	{
		p.SetState(47)
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
		p.SetState(49)
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
		p.SetState(51)
		p.Field_elem()
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QingQLParserT__0 {
		{
			p.SetState(52)
			p.Match(QingQLParserT__0)
		}
		{
			p.SetState(53)
			p.Field_elem()
		}

		p.SetState(58)
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

type FieldElemExprContext struct {
	*Field_elemContext
}

func NewFieldElemExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldElemExprContext {
	var p = new(FieldElemExprContext)

	p.Field_elemContext = NewEmptyField_elemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Field_elemContext))

	return p
}

func (s *FieldElemExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldElemExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FieldElemExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFieldElemExpr(s)
	}
}

func (s *FieldElemExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFieldElemExpr(s)
	}
}

type FieldElemSourceContext struct {
	*Field_elemContext
}

func NewFieldElemSourceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldElemSourceContext {
	var p = new(FieldElemSourceContext)

	p.Field_elemContext = NewEmptyField_elemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Field_elemContext))

	return p
}

func (s *FieldElemSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldElemSourceContext) SourceEntity() ISourceEntityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceEntityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceEntityContext)
}

func (s *FieldElemSourceContext) Asterisk() IAsteriskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsteriskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsteriskContext)
}

func (s *FieldElemSourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFieldElemSource(s)
	}
}

func (s *FieldElemSourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFieldElemSource(s)
	}
}

type FieldElemAsContext struct {
	*Field_elemContext
}

func NewFieldElemAsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldElemAsContext {
	var p = new(FieldElemAsContext)

	p.Field_elemContext = NewEmptyField_elemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Field_elemContext))

	return p
}

func (s *FieldElemAsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldElemAsContext) Field_elem_with_as() IField_elem_with_asContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_elem_with_asContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_elem_with_asContext)
}

func (s *FieldElemAsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterFieldElemAs(s)
	}
}

func (s *FieldElemAsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitFieldElemAs(s)
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

	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFieldElemAsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Field_elem_with_as()
		}

	case 2:
		localctx = NewFieldElemSourceContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.SourceEntity()
		}
		{
			p.SetState(61)
			p.Match(QingQLParserDOT)
		}
		{
			p.SetState(62)
			p.Asterisk()
		}

	case 3:
		localctx = NewFieldElemExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.expr(0)
		}

	}

	return localctx
}

// IField_elem_with_asContext is an interface to support dynamic dispatch.
type IField_elem_with_asContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_elem_with_asContext differentiates from other interfaces.
	IsField_elem_with_asContext()
}

type Field_elem_with_asContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_elem_with_asContext() *Field_elem_with_asContext {
	var p = new(Field_elem_with_asContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_field_elem_with_as
	return p
}

func (*Field_elem_with_asContext) IsField_elem_with_asContext() {}

func NewField_elem_with_asContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_elem_with_asContext {
	var p = new(Field_elem_with_asContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_field_elem_with_as

	return p
}

func (s *Field_elem_with_asContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_elem_with_asContext) CopyFrom(ctx *Field_elem_with_asContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Field_elem_with_asContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_elem_with_asContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TargetAsElemContext struct {
	*Field_elem_with_asContext
}

func NewTargetAsElemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetAsElemContext {
	var p = new(TargetAsElemContext)

	p.Field_elem_with_asContext = NewEmptyField_elem_with_asContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Field_elem_with_asContext))

	return p
}

func (s *TargetAsElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetAsElemContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TargetAsElemContext) AS() antlr.TerminalNode {
	return s.GetToken(QingQLParserAS, 0)
}

func (s *TargetAsElemContext) Target_name() ITarget_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITarget_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITarget_nameContext)
}

func (s *TargetAsElemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterTargetAsElem(s)
	}
}

func (s *TargetAsElemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitTargetAsElem(s)
	}
}

func (p *QingQLParser) Field_elem_with_as() (localctx IField_elem_with_asContext) {
	localctx = NewField_elem_with_asContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QingQLParserRULE_field_elem_with_as)

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

	localctx = NewTargetAsElemContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.expr(0)
	}
	{
		p.SetState(68)
		p.Match(QingQLParserAS)
	}
	{
		p.SetState(69)
		p.Target_name()
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
	p.EnterRule(localctx, 10, QingQLParserRULE_filter)

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
		p.SetState(71)
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
	p.EnterRule(localctx, 12, QingQLParserRULE_filter_condition)
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
		p.SetState(73)
		p.Filter_condition_or()
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QingQLParserAND {
		{
			p.SetState(74)
			p.Match(QingQLParserAND)
		}
		{
			p.SetState(75)
			p.Filter_condition_or()
		}

		p.SetState(80)
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
	p.EnterRule(localctx, 14, QingQLParserRULE_filter_condition_or)
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
		p.SetState(81)
		p.Filter_condition_not()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QingQLParserOR {
		{
			p.SetState(82)
			p.Match(QingQLParserOR)
		}
		{
			p.SetState(83)
			p.Filter_condition_not()
		}

		p.SetState(88)
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
	p.EnterRule(localctx, 16, QingQLParserRULE_filter_condition_not)
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
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == QingQLParserNOT {
		{
			p.SetState(89)
			p.Match(QingQLParserNOT)
		}

	}
	{
		p.SetState(92)
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

func (p *QingQLParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *QingQLParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, QingQLParserRULE_expr, _p)
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
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBracesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(95)
			p.Constant()
		}

	case 2:
		localctx = NewBracesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(96)
			p.Match(QingQLParserT__1)
		}
		{
			p.SetState(97)
			p.expr(0)
		}
		{
			p.SetState(98)
			p.Match(QingQLParserT__2)
		}

	case 3:
		localctx = NewFunctionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(100)
			p.Call_expr()
		}

	case 4:
		localctx = NewSwitchContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(101)
			p.Switch_stmt()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QingQLParserRULE_expr)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(105)

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
					p.SetState(106)
					p.expr(6)
				}

			case 2:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QingQLParserRULE_expr)
				p.SetState(107)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				p.SetState(108)

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
					p.SetState(109)
					p.expr(5)
				}

			case 3:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, QingQLParserRULE_expr)
				p.SetState(110)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(111)

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
					p.SetState(112)
					p.expr(4)
				}

			}

		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
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
		p.SetState(118)
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
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == QingQLParserDOT {
		{
			p.SetState(120)
			p.Match(QingQLParserDOT)
		}
		{
			p.SetState(121)
			p.Match(QingQLParserINDENTIFIER)
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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

type SourceContext struct {
	*ConstantContext
}

func NewSourceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SourceContext {
	var p = new(SourceContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *SourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceContext) Xpath_name() IXpath_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXpath_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXpath_nameContext)
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

	p.SetState(133)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QingQLParserTRUE:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.Match(QingQLParserTRUE)
		}

	case QingQLParserFALSE:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Match(QingQLParserFALSE)
		}

	case QingQLParserNUMBER:
		localctx = NewIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
			p.Match(QingQLParserNUMBER)
		}

	case QingQLParserINTEGER:
		localctx = NewIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(129)
			p.Match(QingQLParserINTEGER)
		}

	case QingQLParserFLOAT:
		localctx = NewFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(130)
			p.Match(QingQLParserFLOAT)
		}

	case QingQLParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(131)
			p.Match(QingQLParserSTRING)
		}

	case QingQLParserT__3, QingQLParserINDENTIFIER, QingQLParserPATHITEM:
		localctx = NewSourceContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(132)
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
		p.SetState(135)
		p.Match(QingQLParserCASE)
	}
	{
		p.SetState(136)
		p.expr(0)
	}
	{
		p.SetState(137)
		p.Match(QingQLParserWHEN)
	}
	{
		p.SetState(138)
		p.expr(0)
	}
	{
		p.SetState(139)
		p.Match(QingQLParserTHEN)
	}
	{
		p.SetState(140)
		p.expr(0)
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(141)
				p.Match(QingQLParserWHEN)
			}
			{
				p.SetState(142)
				p.expr(0)
			}
			{
				p.SetState(143)
				p.Match(QingQLParserTHEN)
			}
			{
				p.SetState(144)
				p.expr(0)
			}

		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(151)
			p.Match(QingQLParserELSE)
		}
		{
			p.SetState(152)
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
		p.SetState(155)

		var _m = p.Match(QingQLParserINDENTIFIER)

		localctx.(*Call_exprContext).key = _m
	}
	{
		p.SetState(156)
		p.Match(QingQLParserT__1)
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<QingQLParserT__1)|(1<<QingQLParserT__3)|(1<<QingQLParserCASE))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(QingQLParserTRUE-37))|(1<<(QingQLParserFALSE-37))|(1<<(QingQLParserINDENTIFIER-37))|(1<<(QingQLParserNUMBER-37))|(1<<(QingQLParserINTEGER-37))|(1<<(QingQLParserFLOAT-37))|(1<<(QingQLParserPATHITEM-37))|(1<<(QingQLParserSTRING-37)))) != 0) {
		{
			p.SetState(157)
			p.expr(0)
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == QingQLParserT__0 {
			{
				p.SetState(158)
				p.Match(QingQLParserT__0)
			}
			{
				p.SetState(159)
				p.expr(0)
			}

			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(167)
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
		p.SetState(169)
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

	p.SetState(184)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QingQLParserINDENTIFIER, QingQLParserPATHITEM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.Dotnotation()
		}

	case QingQLParserT__3:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QingQLParserT__3 {
			{
				p.SetState(172)
				p.Match(QingQLParserT__3)
			}

			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QingQLParserINDENTIFIER || _la == QingQLParserPATHITEM {
			{
				p.SetState(177)
				p.Dotnotation()
			}

			p.SetState(180)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(182)
			p.Match(QingQLParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITarget_nameContext is an interface to support dynamic dispatch.
type ITarget_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTarget_nameContext differentiates from other interfaces.
	IsTarget_nameContext()
}

type Target_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTarget_nameContext() *Target_nameContext {
	var p = new(Target_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QingQLParserRULE_target_name
	return p
}

func (*Target_nameContext) IsTarget_nameContext() {}

func NewTarget_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Target_nameContext {
	var p = new(Target_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QingQLParserRULE_target_name

	return p
}

func (s *Target_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Target_nameContext) AllDotnotation() []IDotnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDotnotationContext)(nil)).Elem())
	var tst = make([]IDotnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDotnotationContext)
		}
	}

	return tst
}

func (s *Target_nameContext) Dotnotation(i int) IDotnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDotnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDotnotationContext)
}

func (s *Target_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Target_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Target_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.EnterTarget_name(s)
	}
}

func (s *Target_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QingQLListener); ok {
		listenerT.ExitTarget_name(s)
	}
}

func (p *QingQLParser) Target_name() (localctx ITarget_nameContext) {
	localctx = NewTarget_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, QingQLParserRULE_target_name)
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

	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QingQLParserINDENTIFIER, QingQLParserPATHITEM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Dotnotation()
		}

	case QingQLParserT__3:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QingQLParserT__3 {
			{
				p.SetState(187)
				p.Match(QingQLParserT__3)
			}

			p.SetState(190)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == QingQLParserINDENTIFIER || _la == QingQLParserPATHITEM {
			{
				p.SetState(192)
				p.Dotnotation()
			}

			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(197)
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
	p.EnterRule(localctx, 36, QingQLParserRULE_dotnotation)
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
	p.SetState(201)
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
	p.EnterRule(localctx, 38, QingQLParserRULE_identifierWithTOPICITEM)

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

	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.Match(QingQLParserPATHITEM)
		}
		{
			p.SetState(204)
			p.Match(QingQLParserT__4)
		}
		{
			p.SetState(205)
			p.Match(QingQLParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Match(QingQLParserPATHITEM)
		}
		{
			p.SetState(207)
			p.Match(QingQLParserT__4)
		}
		{
			p.SetState(208)
			p.Match(QingQLParserNUMBER)
		}
		{
			p.SetState(209)
			p.Match(QingQLParserT__5)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(210)
			p.Match(QingQLParserPATHITEM)
		}
		{
			p.SetState(211)
			p.Match(QingQLParserT__4)
		}
		{
			p.SetState(212)
			p.Match(QingQLParserT__6)
		}
		{
			p.SetState(213)
			p.Match(QingQLParserT__5)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(214)
			p.Match(QingQLParserPATHITEM)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(215)
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
	p.EnterRule(localctx, 40, QingQLParserRULE_identifierWithQualifier)

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

	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)
			p.Match(QingQLParserINDENTIFIER)
		}
		{
			p.SetState(219)
			p.Match(QingQLParserT__7)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.Match(QingQLParserINDENTIFIER)
		}
		{
			p.SetState(221)
			p.Match(QingQLParserT__4)
		}
		{
			p.SetState(222)
			p.Match(QingQLParserNUMBER)
		}
		{
			p.SetState(223)
			p.Match(QingQLParserT__5)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(224)
			p.Match(QingQLParserINDENTIFIER)
		}
		{
			p.SetState(225)
			p.Match(QingQLParserT__8)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(226)
			p.Match(QingQLParserINDENTIFIER)
		}

	}

	return localctx
}

func (p *QingQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
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
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
