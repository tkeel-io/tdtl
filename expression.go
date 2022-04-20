package tdtl

import "github.com/antlr/antlr4/runtime/Go/antlr"

type Expression interface {
	Eval(map[string]Node) Node
	Sources() map[string][]string
}

type expr struct {
	sources  map[string][]string
	extFunc  map[string]ContextFunc
	listener *TDTLListener
}

func NewExpr(sql string, extFunc map[string]ContextFunc) (Expression, error) {
	parse, listener := parse(sql)
	antlr.ParseTreeWalkerDefault.Walk(listener, parse.Field_elem())
	err := listener.error()
	if err != nil {
		return nil, err
	}

	return &expr{
		listener: listener,
		sources:  listener.sources,
		extFunc:  extFunc,
	}, nil
}

func (e *expr) expr() Expr {
	return e.listener.Expr()
}

func (e *expr) Eval(in map[string]Node) Node {
	ctx := NewMapContext(in, e.extFunc)
	return EvalRuleQL(ctx, e.expr())
}

func (e *expr) Sources() map[string][]string {
	return e.sources
}
