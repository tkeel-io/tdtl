/*
Copyright 2021 The tKeel Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package tdtl

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tkeel-io/tdtl/parser"
)

//TDTLListener construct expr
type TDTLListener struct {
	ctx context.Context
	*parser.BaseTDTLListener
	stack   []Expr
	target  string
	sources map[string][]string
	fields  map[string]string
	//fields:    listener.fields,

	expr   Expr
	errors []string
}

func (l *TDTLListener) setTarget(target string) {
	l.target = target
}

func (l *TDTLListener) addSource(source string, expr string) {
	if l.sources == nil {
		l.sources = map[string][]string{}
	}
	_, ok := l.sources[source]
	if !ok {
		l.sources[source] = []string{}
	}
	l.sources[source] = append(l.sources[source], expr)
}

func (l *TDTLListener) addFields(alias string, expr string) {
	if l.fields == nil {
		l.fields = map[string]string{}
	}
	l.fields[alias] = expr
}

func (l *TDTLListener) appendErrorf(format string, a ...interface{}) {
	l.errors = append(l.errors, fmt.Sprintf(format, a...))
}

func (l *TDTLListener) push(i Expr) {
	l.stack = append(l.stack, i)
}

func (l *TDTLListener) pop() Expr {
	if len(l.stack) < 1 {
		//panic("stack is empty unable to pop")
		return nil
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

//Expr expr
func (l *TDTLListener) Expr() Expr {
	return l.stack[len(l.stack)-1]
}

func (l *TDTLListener) ExitRoot(c *parser.RootContext) {
	//fmt.Println("ExitRoot")
	r := &SelectStatementExpr{
		filter: &FilterExpr{},
	}
	if c.Fields() != nil {
		expr := l.pop()
		switch expr := expr.(type) {
		case FieldsExpr:
			r.fields = expr
		default:
			l.appendErrorf("[+]parse fields error[%s]", typeOf(expr))
		}
	}
	l.push(r)
}

//ExitTarget construct target entity from select statement
func (l *TDTLListener) ExitTarget(c *parser.TargetContext) {
	//fmt.Println("ExitTarget", c.GetText())
	l.setTarget(c.GetText())
}

//ExitFields construct fields from select statement
func (l *TDTLListener) ExitFields(c *parser.FieldsContext) {
	//fmt.Println("ExitFields")
	var (
		size = len(c.AllField_elem())
		data FieldsExpr
	)
	data = make([]*FieldExpr, 0, size)
	for size > 0 {
		elem := l.pop()
		switch elem := elem.(type) {
		case *FieldExpr:
			data = append([]*FieldExpr{elem}, data...)
		}
		size--
	}
	l.push(data)
}

//ExitFieldElemExpr
func (l *TDTLListener) ExitFieldElemExpr(c *parser.FieldElemExprContext) {
	//fmt.Println("ExitFieldElemExpr", c.GetText())

}

func (l *TDTLListener) ExitFieldElemAs(c *parser.FieldElemAsContext) {
	//fmt.Println("ExitFieldElemAs", c.GetText())
}

//TODO
//func (l *TDTLListener) ExitExprElemSource(c *parser.ExprElemSourceContext) {
//	//fmt.Println("ExitExprElemSource", c.GetText(), c.SourceEntity().GetText())
//}

func (l *TDTLListener) ExitFieldElemSource(c *parser.FieldElemSourceContext) {
	//fmt.Println("ExitFieldElemSource", c.GetText())
	l.addSource(c.SourceEntity().GetText(), c.GetText())
}

func (l *TDTLListener) ExitTargetAsElem(c *parser.TargetAsElemContext) {
	// fmt.Println("ExitTargetField", c.GetText(), c.Expr().GetText(), c.Target_name())
	var alias string
	expr := l.pop()
	path := c.Target_name()
	if path != nil {
		alias = path.GetText()
		l.push(&FieldExpr{
			exp:   expr,
			alias: alias,
		})
		l.addFields(alias, c.Expr().GetText())
	} else {
		panic("path could not nil")
	}
}

func (l *TDTLListener) ExitBinary(c *parser.BinaryContext) {
	right, left := l.pop(), l.pop()
	//fmt.Println("ExitBinary", c.GetText(), left, c.GetOp().GetText(), right)
	l.push(&BinaryExpr{
		Op:  c.GetOp().GetTokenType(),
		LHS: left,
		RHS: right,
	})
}

func (l *TDTLListener) ExitString(c *parser.StringContext) {
	//fmt.Println("ExitString", c.GetText())
	str := c.GetText()
	if len(str) >= 2 &&
		str[0] == '\'' &&
		str[len(str)-1] == '\'' {
		l.push(StringNode(str[1 : len(str)-1]))
	}
}

func (l *TDTLListener) ExitInteger(c *parser.IntegerContext) {
	//fmt.Println("ExitInteger", c.GetText())
	i, err := strconv.ParseInt(c.GetText(), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	l.push(IntNode(i))
}

func (l *TDTLListener) ExitFloat(c *parser.FloatContext) {
	//fmt.Println("ExitFloat", c.GetText())
	i, err := strconv.ParseFloat(c.GetText(), 64)
	if err != nil {
		panic(err.Error())
	}
	l.push(FloatNode(i))
}

func (l *TDTLListener) ExitBoolean(c *parser.BooleanContext) {
	//fmt.Println("ExitBoolean", c.GetText())
	i, err := strconv.ParseBool(strings.ToLower(c.GetText()))
	if err != nil {
		panic(err.Error())
	}
	l.push(BoolNode(i))
}

func (l *TDTLListener) ExitXpath_name(c *parser.Xpath_nameContext) {
	// fmt.Println("ExitXpath_name", c.GetText())
	str := c.GetText()
	expr := ""
	if str[0] != '"' {
		expr = str
	} else if str[0] == '"' && str[len(str)-1] == '"' {
		expr = str[1 : len(str)-1]
	} else {
		return
	}
	l.push(&JSONPathExpr{
		expr,
	})
	xpaths := strings.Split(expr, ".")
	if expr != "" && len(xpaths) > 0 {
		l.addSource(xpaths[0], expr)
	}
	//error
}

func (l *TDTLListener) ExitCall_expr(c *parser.Call_exprContext) {
	//fmt.Println("ExitCall_expr", c.GetText(), c.AllExpr())
	n := len(c.AllExpr())
	temp := make([]Expr, 0, n)
	for i := 0; i < n; i++ {
		temp = append(temp, l.pop())
	}
	list := make([]Expr, 0, n)
	for i := n - 1; i >= 0; i-- {
		list = append(list, temp[i])
	}
	l.push(&CallExpr{
		raw:  c.GetText(),
		key:  c.GetKey().GetText(),
		args: list,
	})
}

func (l *TDTLListener) ExitSwitch_stmt(c *parser.Switch_stmtContext) {
	//fmt.Println("[-]ExitSwitch_stmt", c.GetText(), len(c.AllExpr()))
	n := len(c.AllExpr())
	expr := &SwitchExpr{
		list: make([]*CaseExpr, 0, n/2-1),
	}
	if n < 2 {
		return
	}
	if n%2 == 0 {
		expr.last = l.pop()
		n--
	} else {

	}
	for n > 1 {
		then, when := l.pop(), l.pop()
		expr.list = append(expr.list, &CaseExpr{
			then: then,
			when: when,
		})
		n -= 2
	}
	expr.exp = l.pop()

	l.push(expr)
}

func (l *TDTLListener) ExitDotnotation(c *parser.DotnotationContext) {
	//fmt.Println("ExitDotnotation", c.GetText())
}

func (l *TDTLListener) ExitIdentifierWithQualifier(c *parser.IdentifierWithQualifierContext) {
	//fmt.Println("IdentifierWithQualifier", c.GetText())
}

func (l *TDTLListener) ExitIdentifierWithTOPICITEM(c *parser.IdentifierWithTOPICITEMContext) {
	//fmt.Println("IdentifierWithTOPICITEM", c.GetText())
}


//

func (l *TDTLListener) ExitFilter_condition(c *parser.Filter_conditionContext) {
	//fmt.Println("ExitFliter_condition", c.GetText(), c.AllAND())
	level := len(c.AllAND())
	left := l.pop()
	for level > 0 {
		right := l.pop()
		left = &BinaryExpr{
			Op:  parser.TDTLLexerAND,
			LHS: left,
			RHS: right,
		}
		level--
	}
	l.push(left)
}

func (l *TDTLListener) ExitFilter_condition_or(c *parser.Filter_condition_orContext) {
	//fmt.Println("ExitFilter_condition_or", c.GetText(), c.AllOR())
	level := len(c.AllOR())
	left := l.pop()
	for level > 0 {
		right := l.pop()
		left = &BinaryExpr{
			Op:  parser.TDTLLexerOR,
			LHS: left,
			RHS: right,
		}
		level--
	}
	l.push(left)
}

func (l *TDTLListener) ExitFilter_condition_not(c *parser.Filter_condition_notContext) {
	//fmt.Println("ExitFilter_condition_not", c.GetText())
	if c.NOT() != nil {
		right := l.pop()
		l.push(&BinaryExpr{
			Op:  parser.TDTLLexerNOT,
			LHS: nil,
			RHS: right,
		})
	}
}


func (l *TDTLListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	//fmt.Println("SyntaxError", recognizer, offendingSymbol, line, column, msg, e)
	l.errors = append(l.errors, fmt.Sprintf("[%d:%d]%s", line, column, msg))

	return
}

func (l *TDTLListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//fmt.Println("ReportAmbiguity", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (l *TDTLListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//fmt.Println("ReportAttemptingFullContext", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (l *TDTLListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//fmt.Println("ReportContextSensitivity", recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

func (l *TDTLListener) error() error {
	if len(l.errors) == 0 {
		return nil
	}
	return errors.New(strings.Join(l.errors, "\n"))
}
