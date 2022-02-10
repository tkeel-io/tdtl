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

	// EnterExprField is called when entering the ExprField production.
	EnterExprField(c *ExprFieldContext)

	// EnterSourceField is called when entering the SourceField production.
	EnterSourceField(c *SourceFieldContext)

	// EnterXpathField is called when entering the XpathField production.
	EnterXpathField(c *XpathFieldContext)

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

	// EnterSource is called when entering the Source production.
	EnterSource(c *SourceContext)

	// EnterSource_stmt is called when entering the source_stmt production.
	EnterSource_stmt(c *Source_stmtContext)

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

	// EnterXPath is called when entering the XPath production.
	EnterXPath(c *XPathContext)

	// EnterSwitch_stmt is called when entering the switch_stmt production.
	EnterSwitch_stmt(c *Switch_stmtContext)

	// EnterCall_expr is called when entering the call_expr production.
	EnterCall_expr(c *Call_exprContext)

	// EnterAsterisk is called when entering the asterisk production.
	EnterAsterisk(c *AsteriskContext)

	// EnterXpath_name is called when entering the xpath_name production.
	EnterXpath_name(c *Xpath_nameContext)

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

	// ExitExprField is called when exiting the ExprField production.
	ExitExprField(c *ExprFieldContext)

	// ExitSourceField is called when exiting the SourceField production.
	ExitSourceField(c *SourceFieldContext)

	// ExitXpathField is called when exiting the XpathField production.
	ExitXpathField(c *XpathFieldContext)

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

	// ExitSource is called when exiting the Source production.
	ExitSource(c *SourceContext)

	// ExitSource_stmt is called when exiting the source_stmt production.
	ExitSource_stmt(c *Source_stmtContext)

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

	// ExitXPath is called when exiting the XPath production.
	ExitXPath(c *XPathContext)

	// ExitSwitch_stmt is called when exiting the switch_stmt production.
	ExitSwitch_stmt(c *Switch_stmtContext)

	// ExitCall_expr is called when exiting the call_expr production.
	ExitCall_expr(c *Call_exprContext)

	// ExitAsterisk is called when exiting the asterisk production.
	ExitAsterisk(c *AsteriskContext)

	// ExitXpath_name is called when exiting the xpath_name production.
	ExitXpath_name(c *Xpath_nameContext)

	// ExitDotnotation is called when exiting the dotnotation production.
	ExitDotnotation(c *DotnotationContext)

	// ExitIdentifierWithTOPICITEM is called when exiting the identifierWithTOPICITEM production.
	ExitIdentifierWithTOPICITEM(c *IdentifierWithTOPICITEMContext)

	// ExitIdentifierWithQualifier is called when exiting the identifierWithQualifier production.
	ExitIdentifierWithQualifier(c *IdentifierWithQualifierContext)
}
