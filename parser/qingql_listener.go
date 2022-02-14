// Generated from QingQL.g4 by ANTLR 4.7.

package parser // QingQL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// QingQLListener is a complete listener for a parse tree produced by QingQLParser.
type QingQLListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterFields is called when entering the fields production.
	EnterFields(c *FieldsContext)

	// EnterFieldElemAs is called when entering the FieldElemAs production.
	EnterFieldElemAs(c *FieldElemAsContext)

	// EnterFieldElemSource is called when entering the FieldElemSource production.
	EnterFieldElemSource(c *FieldElemSourceContext)

	// EnterFieldElemExpr is called when entering the FieldElemExpr production.
	EnterFieldElemExpr(c *FieldElemExprContext)

	// EnterTargetAsElem is called when entering the TargetAsElem production.
	EnterTargetAsElem(c *TargetAsElemContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterFilter_condition is called when entering the filter_condition production.
	EnterFilter_condition(c *Filter_conditionContext)

	// EnterFilter_condition_or is called when entering the filter_condition_or production.
	EnterFilter_condition_or(c *Filter_condition_orContext)

	// EnterFilter_condition_not is called when entering the filter_condition_not production.
	EnterFilter_condition_not(c *Filter_condition_notContext)

	// EnterFunction is called when entering the Function production.
	EnterFunction(c *FunctionContext)

	// EnterBraces is called when entering the Braces production.
	EnterBraces(c *BracesContext)

	// EnterSwitch is called when entering the Switch production.
	EnterSwitch(c *SwitchContext)

	// EnterBinary is called when entering the Binary production.
	EnterBinary(c *BinaryContext)

	// EnterSourceEntity is called when entering the sourceEntity production.
	EnterSourceEntity(c *SourceEntityContext)

	// EnterPropertyEntity is called when entering the propertyEntity production.
	EnterPropertyEntity(c *PropertyEntityContext)

	// EnterBoolean is called when entering the Boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterInteger is called when entering the Integer production.
	EnterInteger(c *IntegerContext)

	// EnterFloat is called when entering the Float production.
	EnterFloat(c *FloatContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterSource is called when entering the Source production.
	EnterSource(c *SourceContext)

	// EnterSwitch_stmt is called when entering the switch_stmt production.
	EnterSwitch_stmt(c *Switch_stmtContext)

	// EnterCall_expr is called when entering the call_expr production.
	EnterCall_expr(c *Call_exprContext)

	// EnterAsterisk is called when entering the asterisk production.
	EnterAsterisk(c *AsteriskContext)

	// EnterXpath_name is called when entering the xpath_name production.
	EnterXpath_name(c *Xpath_nameContext)

	// EnterTarget_name is called when entering the target_name production.
	EnterTarget_name(c *Target_nameContext)

	// EnterDotnotation is called when entering the dotnotation production.
	EnterDotnotation(c *DotnotationContext)

	// EnterIdentifierWithTOPICITEM is called when entering the identifierWithTOPICITEM production.
	EnterIdentifierWithTOPICITEM(c *IdentifierWithTOPICITEMContext)

	// EnterIdentifierWithQualifier is called when entering the identifierWithQualifier production.
	EnterIdentifierWithQualifier(c *IdentifierWithQualifierContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitFields is called when exiting the fields production.
	ExitFields(c *FieldsContext)

	// ExitFieldElemAs is called when exiting the FieldElemAs production.
	ExitFieldElemAs(c *FieldElemAsContext)

	// ExitFieldElemSource is called when exiting the FieldElemSource production.
	ExitFieldElemSource(c *FieldElemSourceContext)

	// ExitFieldElemExpr is called when exiting the FieldElemExpr production.
	ExitFieldElemExpr(c *FieldElemExprContext)

	// ExitTargetAsElem is called when exiting the TargetAsElem production.
	ExitTargetAsElem(c *TargetAsElemContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitFilter_condition is called when exiting the filter_condition production.
	ExitFilter_condition(c *Filter_conditionContext)

	// ExitFilter_condition_or is called when exiting the filter_condition_or production.
	ExitFilter_condition_or(c *Filter_condition_orContext)

	// ExitFilter_condition_not is called when exiting the filter_condition_not production.
	ExitFilter_condition_not(c *Filter_condition_notContext)

	// ExitFunction is called when exiting the Function production.
	ExitFunction(c *FunctionContext)

	// ExitBraces is called when exiting the Braces production.
	ExitBraces(c *BracesContext)

	// ExitSwitch is called when exiting the Switch production.
	ExitSwitch(c *SwitchContext)

	// ExitBinary is called when exiting the Binary production.
	ExitBinary(c *BinaryContext)

	// ExitSourceEntity is called when exiting the sourceEntity production.
	ExitSourceEntity(c *SourceEntityContext)

	// ExitPropertyEntity is called when exiting the propertyEntity production.
	ExitPropertyEntity(c *PropertyEntityContext)

	// ExitBoolean is called when exiting the Boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitInteger is called when exiting the Integer production.
	ExitInteger(c *IntegerContext)

	// ExitFloat is called when exiting the Float production.
	ExitFloat(c *FloatContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitSource is called when exiting the Source production.
	ExitSource(c *SourceContext)

	// ExitSwitch_stmt is called when exiting the switch_stmt production.
	ExitSwitch_stmt(c *Switch_stmtContext)

	// ExitCall_expr is called when exiting the call_expr production.
	ExitCall_expr(c *Call_exprContext)

	// ExitAsterisk is called when exiting the asterisk production.
	ExitAsterisk(c *AsteriskContext)

	// ExitXpath_name is called when exiting the xpath_name production.
	ExitXpath_name(c *Xpath_nameContext)

	// ExitTarget_name is called when exiting the target_name production.
	ExitTarget_name(c *Target_nameContext)

	// ExitDotnotation is called when exiting the dotnotation production.
	ExitDotnotation(c *DotnotationContext)

	// ExitIdentifierWithTOPICITEM is called when exiting the identifierWithTOPICITEM production.
	ExitIdentifierWithTOPICITEM(c *IdentifierWithTOPICITEMContext)

	// ExitIdentifierWithQualifier is called when exiting the identifierWithQualifier production.
	ExitIdentifierWithQualifier(c *IdentifierWithQualifierContext)
}
