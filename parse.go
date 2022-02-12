/*
 * Copyright (C) 2019 Yunify, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this work except in compliance with the License.
 * You may obtain a copy of the License in the LICENSE file, or at:
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ruleql

import (
	"context"
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tkeel-io/tql/parser"
	"strconv"
	"strings"
)

//QingQLListener construct expr
type QingQLListener struct {
	ctx context.Context
	*parser.BaseQingQLListener
	stack  []Expr
	expr   Expr
	errors []string
}

func (l *QingQLListener) appendErrorf(format string, a ...interface{}) {
	l.errors = append(l.errors, fmt.Sprintf(format, a...))
}

func (l *QingQLListener) push(i Expr) {
	l.stack = append(l.stack, i)
}

func (l *QingQLListener) pop() Expr {
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
func (l *QingQLListener) Expr() Expr {
	return l.stack[len(l.stack)-1]
}

func (l *QingQLListener) ExitRoot(c *parser.RootContext) {
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

//ExitSelect_list construct fields from select statement
func (l *QingQLListener) ExitFields(c *parser.FieldsContext) {
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

func (l *QingQLListener) ExitXpathField(c *parser.XpathFieldContext) {
	//fmt.Println("ExitXpathField")
	var alias string
	xpath := c.GetText()
	//if xpath != "*" {
	//	arr := strings.Split(xpath, ".")
	//	alias = arr[len(arr)-1]
	//}
	l.push(&FieldExpr{
		exp: &JSONPathExpr{
			xpath,
		},
		alias: alias,
	})
}

func (l *QingQLListener) ExitExprField(c *parser.ExprFieldContext) {
	//fmt.Println("ExitExprField", c.Expr(), c.Xpath_name())
	var alias string
	expr := l.pop()
	path := c.Xpath_name()
	if path != nil {
		alias = path.GetText()
		l.push(&FieldExpr{
			exp:   expr,
			alias: alias,
		})
	} else {
		//fmt.Println("+", expr, alias)
		l.push(&FieldExpr{
			exp:   expr,
			alias: "",
		})
	}
}

func (l *QingQLListener) ExitBinary(c *parser.BinaryContext) {
	right, left := l.pop(), l.pop()
	//fmt.Println("ExitBinary", c.GetText(), left, c.GetOp().GetText(), right)
	l.push(&BinaryExpr{
		Op:  c.GetOp().GetTokenType(),
		LHS: left,
		RHS: right,
	})
}

func (l *QingQLListener) ExitString(c *parser.StringContext) {
	//fmt.Println("ExitString", c.GetText())
	str := c.GetText()
	if len(str) >= 2 &&
		str[0] == '\'' &&
		str[len(str)-1] == '\'' {
		l.push(StringNode(str[1 : len(str)-1]))
	}
}

func (l *QingQLListener) ExitInteger(c *parser.IntegerContext) {
	//fmt.Println("ExitInteger", c.GetText())
	i, err := strconv.ParseInt(c.GetText(), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	l.push(IntNode(i))
}

func (l *QingQLListener) ExitFloat(c *parser.FloatContext) {
	//fmt.Println("ExitFloat", c.GetText())
	i, err := strconv.ParseFloat(c.GetText(), 64)
	if err != nil {
		panic(err.Error())
	}
	l.push(FloatNode(i))
}

func (l *QingQLListener) ExitBoolean(c *parser.BooleanContext) {
	//fmt.Println("ExitBoolean", c.GetText())
	i, err := strconv.ParseBool(strings.ToLower(c.GetText()))
	if err != nil {
		panic(err.Error())
	}
	l.push(BoolNode(i))
}

func (l *QingQLListener) ExitXPath(c *parser.XPathContext) {
	//fmt.Println("ExitXpath", c.GetText())
	str := c.GetText()
	if str[0] != '"' {
		l.push(&JSONPathExpr{
			str,
		})
	} else {
		if str[len(str)-1] == '"' {
			l.push(&JSONPathExpr{
				str[1 : len(str)-1],
			})
		}
	}
	//error
}

func (l *QingQLListener) ExitCall_expr(c *parser.Call_exprContext) {
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

func (l *QingQLListener) ExitSwitch_stmt(c *parser.Switch_stmtContext) {
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

func (l *QingQLListener) ExitDotnotation(c *parser.DotnotationContext) {
	//fmt.Println("ExitDotnotation", c.GetText())
}

func (l *QingQLListener) ExitIdentifierWithQualifier(c *parser.IdentifierWithQualifierContext) {
	//fmt.Println("IdentifierWithQualifier", c.GetText())
}

func (l *QingQLListener) ExitIdentifierWithTOPICITEM(c *parser.IdentifierWithTOPICITEMContext) {
	//fmt.Println("IdentifierWithTOPICITEM", c.GetText())
}

func (l *QingQLListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	//fmt.Println("SyntaxError", recognizer, offendingSymbol, line, column, msg, e)
	l.errors = append(l.errors, fmt.Sprintf("[%d:%d]%s", line, column, msg))

	return
}

func (l *QingQLListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//fmt.Println("ReportAmbiguity", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (l *QingQLListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//fmt.Println("ReportAttemptingFullContext", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (l *QingQLListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//fmt.Println("ReportContextSensitivity", recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

func (l *QingQLListener) error() error {
	if len(l.errors) == 0 {
		return nil
	}
	return errors.New(strings.Join(l.errors, "\n"))
}
