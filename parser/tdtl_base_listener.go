// Generated from TDTL.g4 by ANTLR 4.7.

package parser // TDTL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTDTLListener is a complete listener for a parse tree produced by TDTLParser.
type BaseTDTLListener struct{}

var _ TDTLListener = &BaseTDTLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTDTLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTDTLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTDTLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTDTLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseTDTLListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseTDTLListener) ExitRoot(ctx *RootContext) {}

// EnterTarget is called when production target is entered.
func (s *BaseTDTLListener) EnterTarget(ctx *TargetContext) {}

// ExitTarget is called when production target is exited.
func (s *BaseTDTLListener) ExitTarget(ctx *TargetContext) {}

// EnterFields is called when production fields is entered.
func (s *BaseTDTLListener) EnterFields(ctx *FieldsContext) {}

// ExitFields is called when production fields is exited.
func (s *BaseTDTLListener) ExitFields(ctx *FieldsContext) {}

// EnterFieldElemAs is called when production FieldElemAs is entered.
func (s *BaseTDTLListener) EnterFieldElemAs(ctx *FieldElemAsContext) {}

// ExitFieldElemAs is called when production FieldElemAs is exited.
func (s *BaseTDTLListener) ExitFieldElemAs(ctx *FieldElemAsContext) {}

// EnterFieldElemSource is called when production FieldElemSource is entered.
func (s *BaseTDTLListener) EnterFieldElemSource(ctx *FieldElemSourceContext) {}

// ExitFieldElemSource is called when production FieldElemSource is exited.
func (s *BaseTDTLListener) ExitFieldElemSource(ctx *FieldElemSourceContext) {}

// EnterFieldElemExpr is called when production FieldElemExpr is entered.
func (s *BaseTDTLListener) EnterFieldElemExpr(ctx *FieldElemExprContext) {}

// ExitFieldElemExpr is called when production FieldElemExpr is exited.
func (s *BaseTDTLListener) ExitFieldElemExpr(ctx *FieldElemExprContext) {}

// EnterTargetAsElem is called when production TargetAsElem is entered.
func (s *BaseTDTLListener) EnterTargetAsElem(ctx *TargetAsElemContext) {}

// ExitTargetAsElem is called when production TargetAsElem is exited.
func (s *BaseTDTLListener) ExitTargetAsElem(ctx *TargetAsElemContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseTDTLListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseTDTLListener) ExitFilter(ctx *FilterContext) {}

// EnterFilter_condition is called when production filter_condition is entered.
func (s *BaseTDTLListener) EnterFilter_condition(ctx *Filter_conditionContext) {}

// ExitFilter_condition is called when production filter_condition is exited.
func (s *BaseTDTLListener) ExitFilter_condition(ctx *Filter_conditionContext) {}

// EnterFilter_condition_or is called when production filter_condition_or is entered.
func (s *BaseTDTLListener) EnterFilter_condition_or(ctx *Filter_condition_orContext) {}

// ExitFilter_condition_or is called when production filter_condition_or is exited.
func (s *BaseTDTLListener) ExitFilter_condition_or(ctx *Filter_condition_orContext) {}

// EnterFilter_condition_not is called when production filter_condition_not is entered.
func (s *BaseTDTLListener) EnterFilter_condition_not(ctx *Filter_condition_notContext) {}

// ExitFilter_condition_not is called when production filter_condition_not is exited.
func (s *BaseTDTLListener) ExitFilter_condition_not(ctx *Filter_condition_notContext) {}

// EnterFunction is called when production Function is entered.
func (s *BaseTDTLListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production Function is exited.
func (s *BaseTDTLListener) ExitFunction(ctx *FunctionContext) {}

// EnterBraces is called when production Braces is entered.
func (s *BaseTDTLListener) EnterBraces(ctx *BracesContext) {}

// ExitBraces is called when production Braces is exited.
func (s *BaseTDTLListener) ExitBraces(ctx *BracesContext) {}

// EnterSwitch is called when production Switch is entered.
func (s *BaseTDTLListener) EnterSwitch(ctx *SwitchContext) {}

// ExitSwitch is called when production Switch is exited.
func (s *BaseTDTLListener) ExitSwitch(ctx *SwitchContext) {}

// EnterBinary is called when production Binary is entered.
func (s *BaseTDTLListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production Binary is exited.
func (s *BaseTDTLListener) ExitBinary(ctx *BinaryContext) {}

// EnterSourceEntity is called when production sourceEntity is entered.
func (s *BaseTDTLListener) EnterSourceEntity(ctx *SourceEntityContext) {}

// ExitSourceEntity is called when production sourceEntity is exited.
func (s *BaseTDTLListener) ExitSourceEntity(ctx *SourceEntityContext) {}

// EnterPropertyEntity is called when production propertyEntity is entered.
func (s *BaseTDTLListener) EnterPropertyEntity(ctx *PropertyEntityContext) {}

// ExitPropertyEntity is called when production propertyEntity is exited.
func (s *BaseTDTLListener) ExitPropertyEntity(ctx *PropertyEntityContext) {}

// EnterBoolean is called when production Boolean is entered.
func (s *BaseTDTLListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production Boolean is exited.
func (s *BaseTDTLListener) ExitBoolean(ctx *BooleanContext) {}

// EnterInteger is called when production Integer is entered.
func (s *BaseTDTLListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production Integer is exited.
func (s *BaseTDTLListener) ExitInteger(ctx *IntegerContext) {}

// EnterFloat is called when production Float is entered.
func (s *BaseTDTLListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production Float is exited.
func (s *BaseTDTLListener) ExitFloat(ctx *FloatContext) {}

// EnterString is called when production String is entered.
func (s *BaseTDTLListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseTDTLListener) ExitString(ctx *StringContext) {}

// EnterSource is called when production Source is entered.
func (s *BaseTDTLListener) EnterSource(ctx *SourceContext) {}

// ExitSource is called when production Source is exited.
func (s *BaseTDTLListener) ExitSource(ctx *SourceContext) {}

// EnterSwitch_stmt is called when production switch_stmt is entered.
func (s *BaseTDTLListener) EnterSwitch_stmt(ctx *Switch_stmtContext) {}

// ExitSwitch_stmt is called when production switch_stmt is exited.
func (s *BaseTDTLListener) ExitSwitch_stmt(ctx *Switch_stmtContext) {}

// EnterCall_expr is called when production call_expr is entered.
func (s *BaseTDTLListener) EnterCall_expr(ctx *Call_exprContext) {}

// ExitCall_expr is called when production call_expr is exited.
func (s *BaseTDTLListener) ExitCall_expr(ctx *Call_exprContext) {}

// EnterAsterisk is called when production asterisk is entered.
func (s *BaseTDTLListener) EnterAsterisk(ctx *AsteriskContext) {}

// ExitAsterisk is called when production asterisk is exited.
func (s *BaseTDTLListener) ExitAsterisk(ctx *AsteriskContext) {}

// EnterXpath_name is called when production xpath_name is entered.
func (s *BaseTDTLListener) EnterXpath_name(ctx *Xpath_nameContext) {}

// ExitXpath_name is called when production xpath_name is exited.
func (s *BaseTDTLListener) ExitXpath_name(ctx *Xpath_nameContext) {}

// EnterTarget_name is called when production target_name is entered.
func (s *BaseTDTLListener) EnterTarget_name(ctx *Target_nameContext) {}

// ExitTarget_name is called when production target_name is exited.
func (s *BaseTDTLListener) ExitTarget_name(ctx *Target_nameContext) {}

// EnterDotnotation is called when production dotnotation is entered.
func (s *BaseTDTLListener) EnterDotnotation(ctx *DotnotationContext) {}

// ExitDotnotation is called when production dotnotation is exited.
func (s *BaseTDTLListener) ExitDotnotation(ctx *DotnotationContext) {}

// EnterIdentifierWithTOPICITEM is called when production identifierWithTOPICITEM is entered.
func (s *BaseTDTLListener) EnterIdentifierWithTOPICITEM(ctx *IdentifierWithTOPICITEMContext) {}

// ExitIdentifierWithTOPICITEM is called when production identifierWithTOPICITEM is exited.
func (s *BaseTDTLListener) ExitIdentifierWithTOPICITEM(ctx *IdentifierWithTOPICITEMContext) {}

// EnterIdentifierWithQualifier is called when production identifierWithQualifier is entered.
func (s *BaseTDTLListener) EnterIdentifierWithQualifier(ctx *IdentifierWithQualifierContext) {}

// ExitIdentifierWithQualifier is called when production identifierWithQualifier is exited.
func (s *BaseTDTLListener) ExitIdentifierWithQualifier(ctx *IdentifierWithQualifierContext) {}
