// Generated from QingQL.g4 by ANTLR 4.7.

package parser // QingQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQingQLListener is a complete listener for a parse tree produced by QingQLParser.
type BaseQingQLListener struct{}

var _ QingQLListener = &BaseQingQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQingQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQingQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQingQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQingQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseQingQLListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseQingQLListener) ExitRoot(ctx *RootContext) {}

// EnterTarget is called when production target is entered.
func (s *BaseQingQLListener) EnterTarget(ctx *TargetContext) {}

// ExitTarget is called when production target is exited.
func (s *BaseQingQLListener) ExitTarget(ctx *TargetContext) {}

// EnterFields is called when production fields is entered.
func (s *BaseQingQLListener) EnterFields(ctx *FieldsContext) {}

// ExitFields is called when production fields is exited.
func (s *BaseQingQLListener) ExitFields(ctx *FieldsContext) {}

// EnterFieldElemAs is called when production FieldElemAs is entered.
func (s *BaseQingQLListener) EnterFieldElemAs(ctx *FieldElemAsContext) {}

// ExitFieldElemAs is called when production FieldElemAs is exited.
func (s *BaseQingQLListener) ExitFieldElemAs(ctx *FieldElemAsContext) {}

// EnterFieldElemSource is called when production FieldElemSource is entered.
func (s *BaseQingQLListener) EnterFieldElemSource(ctx *FieldElemSourceContext) {}

// ExitFieldElemSource is called when production FieldElemSource is exited.
func (s *BaseQingQLListener) ExitFieldElemSource(ctx *FieldElemSourceContext) {}

// EnterFieldElemExpr is called when production FieldElemExpr is entered.
func (s *BaseQingQLListener) EnterFieldElemExpr(ctx *FieldElemExprContext) {}

// ExitFieldElemExpr is called when production FieldElemExpr is exited.
func (s *BaseQingQLListener) ExitFieldElemExpr(ctx *FieldElemExprContext) {}

// EnterTargetAsElem is called when production TargetAsElem is entered.
func (s *BaseQingQLListener) EnterTargetAsElem(ctx *TargetAsElemContext) {}

// ExitTargetAsElem is called when production TargetAsElem is exited.
func (s *BaseQingQLListener) ExitTargetAsElem(ctx *TargetAsElemContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseQingQLListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseQingQLListener) ExitFilter(ctx *FilterContext) {}

// EnterFilter_condition is called when production filter_condition is entered.
func (s *BaseQingQLListener) EnterFilter_condition(ctx *Filter_conditionContext) {}

// ExitFilter_condition is called when production filter_condition is exited.
func (s *BaseQingQLListener) ExitFilter_condition(ctx *Filter_conditionContext) {}

// EnterFilter_condition_or is called when production filter_condition_or is entered.
func (s *BaseQingQLListener) EnterFilter_condition_or(ctx *Filter_condition_orContext) {}

// ExitFilter_condition_or is called when production filter_condition_or is exited.
func (s *BaseQingQLListener) ExitFilter_condition_or(ctx *Filter_condition_orContext) {}

// EnterFilter_condition_not is called when production filter_condition_not is entered.
func (s *BaseQingQLListener) EnterFilter_condition_not(ctx *Filter_condition_notContext) {}

// ExitFilter_condition_not is called when production filter_condition_not is exited.
func (s *BaseQingQLListener) ExitFilter_condition_not(ctx *Filter_condition_notContext) {}

// EnterFunction is called when production Function is entered.
func (s *BaseQingQLListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production Function is exited.
func (s *BaseQingQLListener) ExitFunction(ctx *FunctionContext) {}

// EnterBraces is called when production Braces is entered.
func (s *BaseQingQLListener) EnterBraces(ctx *BracesContext) {}

// ExitBraces is called when production Braces is exited.
func (s *BaseQingQLListener) ExitBraces(ctx *BracesContext) {}

// EnterSwitch is called when production Switch is entered.
func (s *BaseQingQLListener) EnterSwitch(ctx *SwitchContext) {}

// ExitSwitch is called when production Switch is exited.
func (s *BaseQingQLListener) ExitSwitch(ctx *SwitchContext) {}

// EnterBinary is called when production Binary is entered.
func (s *BaseQingQLListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production Binary is exited.
func (s *BaseQingQLListener) ExitBinary(ctx *BinaryContext) {}

// EnterSourceEntity is called when production sourceEntity is entered.
func (s *BaseQingQLListener) EnterSourceEntity(ctx *SourceEntityContext) {}

// ExitSourceEntity is called when production sourceEntity is exited.
func (s *BaseQingQLListener) ExitSourceEntity(ctx *SourceEntityContext) {}

// EnterPropertyEntity is called when production propertyEntity is entered.
func (s *BaseQingQLListener) EnterPropertyEntity(ctx *PropertyEntityContext) {}

// ExitPropertyEntity is called when production propertyEntity is exited.
func (s *BaseQingQLListener) ExitPropertyEntity(ctx *PropertyEntityContext) {}

// EnterBoolean is called when production Boolean is entered.
func (s *BaseQingQLListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production Boolean is exited.
func (s *BaseQingQLListener) ExitBoolean(ctx *BooleanContext) {}

// EnterInteger is called when production Integer is entered.
func (s *BaseQingQLListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production Integer is exited.
func (s *BaseQingQLListener) ExitInteger(ctx *IntegerContext) {}

// EnterFloat is called when production Float is entered.
func (s *BaseQingQLListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production Float is exited.
func (s *BaseQingQLListener) ExitFloat(ctx *FloatContext) {}

// EnterString is called when production String is entered.
func (s *BaseQingQLListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseQingQLListener) ExitString(ctx *StringContext) {}

// EnterSource is called when production Source is entered.
func (s *BaseQingQLListener) EnterSource(ctx *SourceContext) {}

// ExitSource is called when production Source is exited.
func (s *BaseQingQLListener) ExitSource(ctx *SourceContext) {}

// EnterSwitch_stmt is called when production switch_stmt is entered.
func (s *BaseQingQLListener) EnterSwitch_stmt(ctx *Switch_stmtContext) {}

// ExitSwitch_stmt is called when production switch_stmt is exited.
func (s *BaseQingQLListener) ExitSwitch_stmt(ctx *Switch_stmtContext) {}

// EnterCall_expr is called when production call_expr is entered.
func (s *BaseQingQLListener) EnterCall_expr(ctx *Call_exprContext) {}

// ExitCall_expr is called when production call_expr is exited.
func (s *BaseQingQLListener) ExitCall_expr(ctx *Call_exprContext) {}

// EnterAsterisk is called when production asterisk is entered.
func (s *BaseQingQLListener) EnterAsterisk(ctx *AsteriskContext) {}

// ExitAsterisk is called when production asterisk is exited.
func (s *BaseQingQLListener) ExitAsterisk(ctx *AsteriskContext) {}

// EnterXpath_name is called when production xpath_name is entered.
func (s *BaseQingQLListener) EnterXpath_name(ctx *Xpath_nameContext) {}

// ExitXpath_name is called when production xpath_name is exited.
func (s *BaseQingQLListener) ExitXpath_name(ctx *Xpath_nameContext) {}

// EnterTarget_name is called when production target_name is entered.
func (s *BaseQingQLListener) EnterTarget_name(ctx *Target_nameContext) {}

// ExitTarget_name is called when production target_name is exited.
func (s *BaseQingQLListener) ExitTarget_name(ctx *Target_nameContext) {}

// EnterDotnotation is called when production dotnotation is entered.
func (s *BaseQingQLListener) EnterDotnotation(ctx *DotnotationContext) {}

// ExitDotnotation is called when production dotnotation is exited.
func (s *BaseQingQLListener) ExitDotnotation(ctx *DotnotationContext) {}

// EnterIdentifierWithTOPICITEM is called when production identifierWithTOPICITEM is entered.
func (s *BaseQingQLListener) EnterIdentifierWithTOPICITEM(ctx *IdentifierWithTOPICITEMContext) {}

// ExitIdentifierWithTOPICITEM is called when production identifierWithTOPICITEM is exited.
func (s *BaseQingQLListener) ExitIdentifierWithTOPICITEM(ctx *IdentifierWithTOPICITEMContext) {}

// EnterIdentifierWithQualifier is called when production identifierWithQualifier is entered.
func (s *BaseQingQLListener) EnterIdentifierWithQualifier(ctx *IdentifierWithQualifierContext) {}

// ExitIdentifierWithQualifier is called when production identifierWithQualifier is exited.
func (s *BaseQingQLListener) ExitIdentifierWithQualifier(ctx *IdentifierWithQualifierContext) {}
