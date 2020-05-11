// Code generated from ../../ShExDoc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ShExDoc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseShExDocListener is a complete listener for a parse tree produced by ShExDocParser.
type BaseShExDocListener struct{}

var _ ShExDocListener = &BaseShExDocListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseShExDocListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseShExDocListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseShExDocListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseShExDocListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterShExDoc is called when production shExDoc is entered.
func (s *BaseShExDocListener) EnterShExDoc(ctx *ShExDocContext) {}

// ExitShExDoc is called when production shExDoc is exited.
func (s *BaseShExDocListener) ExitShExDoc(ctx *ShExDocContext) {}

// EnterDirective is called when production directive is entered.
func (s *BaseShExDocListener) EnterDirective(ctx *DirectiveContext) {}

// ExitDirective is called when production directive is exited.
func (s *BaseShExDocListener) ExitDirective(ctx *DirectiveContext) {}

// EnterBaseDecl is called when production baseDecl is entered.
func (s *BaseShExDocListener) EnterBaseDecl(ctx *BaseDeclContext) {}

// ExitBaseDecl is called when production baseDecl is exited.
func (s *BaseShExDocListener) ExitBaseDecl(ctx *BaseDeclContext) {}

// EnterPrefixDecl is called when production prefixDecl is entered.
func (s *BaseShExDocListener) EnterPrefixDecl(ctx *PrefixDeclContext) {}

// ExitPrefixDecl is called when production prefixDecl is exited.
func (s *BaseShExDocListener) ExitPrefixDecl(ctx *PrefixDeclContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BaseShExDocListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BaseShExDocListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterNotStartAction is called when production notStartAction is entered.
func (s *BaseShExDocListener) EnterNotStartAction(ctx *NotStartActionContext) {}

// ExitNotStartAction is called when production notStartAction is exited.
func (s *BaseShExDocListener) ExitNotStartAction(ctx *NotStartActionContext) {}

// EnterStart is called when production start is entered.
func (s *BaseShExDocListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseShExDocListener) ExitStart(ctx *StartContext) {}

// EnterStartActions is called when production startActions is entered.
func (s *BaseShExDocListener) EnterStartActions(ctx *StartActionsContext) {}

// ExitStartActions is called when production startActions is exited.
func (s *BaseShExDocListener) ExitStartActions(ctx *StartActionsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseShExDocListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseShExDocListener) ExitStatement(ctx *StatementContext) {}

// EnterShapeExprDecl is called when production shapeExprDecl is entered.
func (s *BaseShExDocListener) EnterShapeExprDecl(ctx *ShapeExprDeclContext) {}

// ExitShapeExprDecl is called when production shapeExprDecl is exited.
func (s *BaseShExDocListener) ExitShapeExprDecl(ctx *ShapeExprDeclContext) {}

// EnterShapeExpression is called when production shapeExpression is entered.
func (s *BaseShExDocListener) EnterShapeExpression(ctx *ShapeExpressionContext) {}

// ExitShapeExpression is called when production shapeExpression is exited.
func (s *BaseShExDocListener) ExitShapeExpression(ctx *ShapeExpressionContext) {}

// EnterInlineShapeExpression is called when production inlineShapeExpression is entered.
func (s *BaseShExDocListener) EnterInlineShapeExpression(ctx *InlineShapeExpressionContext) {}

// ExitInlineShapeExpression is called when production inlineShapeExpression is exited.
func (s *BaseShExDocListener) ExitInlineShapeExpression(ctx *InlineShapeExpressionContext) {}

// EnterShapeOr is called when production shapeOr is entered.
func (s *BaseShExDocListener) EnterShapeOr(ctx *ShapeOrContext) {}

// ExitShapeOr is called when production shapeOr is exited.
func (s *BaseShExDocListener) ExitShapeOr(ctx *ShapeOrContext) {}

// EnterInlineShapeOr is called when production inlineShapeOr is entered.
func (s *BaseShExDocListener) EnterInlineShapeOr(ctx *InlineShapeOrContext) {}

// ExitInlineShapeOr is called when production inlineShapeOr is exited.
func (s *BaseShExDocListener) ExitInlineShapeOr(ctx *InlineShapeOrContext) {}

// EnterShapeAnd is called when production shapeAnd is entered.
func (s *BaseShExDocListener) EnterShapeAnd(ctx *ShapeAndContext) {}

// ExitShapeAnd is called when production shapeAnd is exited.
func (s *BaseShExDocListener) ExitShapeAnd(ctx *ShapeAndContext) {}

// EnterInlineShapeAnd is called when production inlineShapeAnd is entered.
func (s *BaseShExDocListener) EnterInlineShapeAnd(ctx *InlineShapeAndContext) {}

// ExitInlineShapeAnd is called when production inlineShapeAnd is exited.
func (s *BaseShExDocListener) ExitInlineShapeAnd(ctx *InlineShapeAndContext) {}

// EnterShapeNot is called when production shapeNot is entered.
func (s *BaseShExDocListener) EnterShapeNot(ctx *ShapeNotContext) {}

// ExitShapeNot is called when production shapeNot is exited.
func (s *BaseShExDocListener) ExitShapeNot(ctx *ShapeNotContext) {}

// EnterInlineShapeNot is called when production inlineShapeNot is entered.
func (s *BaseShExDocListener) EnterInlineShapeNot(ctx *InlineShapeNotContext) {}

// ExitInlineShapeNot is called when production inlineShapeNot is exited.
func (s *BaseShExDocListener) ExitInlineShapeNot(ctx *InlineShapeNotContext) {}

// EnterShapeAtomNonLitNodeConstraint is called when production shapeAtomNonLitNodeConstraint is entered.
func (s *BaseShExDocListener) EnterShapeAtomNonLitNodeConstraint(ctx *ShapeAtomNonLitNodeConstraintContext) {}

// ExitShapeAtomNonLitNodeConstraint is called when production shapeAtomNonLitNodeConstraint is exited.
func (s *BaseShExDocListener) ExitShapeAtomNonLitNodeConstraint(ctx *ShapeAtomNonLitNodeConstraintContext) {}

// EnterShapeAtomLitNodeConstraint is called when production shapeAtomLitNodeConstraint is entered.
func (s *BaseShExDocListener) EnterShapeAtomLitNodeConstraint(ctx *ShapeAtomLitNodeConstraintContext) {}

// ExitShapeAtomLitNodeConstraint is called when production shapeAtomLitNodeConstraint is exited.
func (s *BaseShExDocListener) ExitShapeAtomLitNodeConstraint(ctx *ShapeAtomLitNodeConstraintContext) {}

// EnterShapeAtomShapeOrRef is called when production shapeAtomShapeOrRef is entered.
func (s *BaseShExDocListener) EnterShapeAtomShapeOrRef(ctx *ShapeAtomShapeOrRefContext) {}

// ExitShapeAtomShapeOrRef is called when production shapeAtomShapeOrRef is exited.
func (s *BaseShExDocListener) ExitShapeAtomShapeOrRef(ctx *ShapeAtomShapeOrRefContext) {}

// EnterShapeAtomShapeExpression is called when production shapeAtomShapeExpression is entered.
func (s *BaseShExDocListener) EnterShapeAtomShapeExpression(ctx *ShapeAtomShapeExpressionContext) {}

// ExitShapeAtomShapeExpression is called when production shapeAtomShapeExpression is exited.
func (s *BaseShExDocListener) ExitShapeAtomShapeExpression(ctx *ShapeAtomShapeExpressionContext) {}

// EnterShapeAtomAny is called when production shapeAtomAny is entered.
func (s *BaseShExDocListener) EnterShapeAtomAny(ctx *ShapeAtomAnyContext) {}

// ExitShapeAtomAny is called when production shapeAtomAny is exited.
func (s *BaseShExDocListener) ExitShapeAtomAny(ctx *ShapeAtomAnyContext) {}

// EnterInlineShapeAtomNonLitNodeConstraint is called when production inlineShapeAtomNonLitNodeConstraint is entered.
func (s *BaseShExDocListener) EnterInlineShapeAtomNonLitNodeConstraint(ctx *InlineShapeAtomNonLitNodeConstraintContext) {}

// ExitInlineShapeAtomNonLitNodeConstraint is called when production inlineShapeAtomNonLitNodeConstraint is exited.
func (s *BaseShExDocListener) ExitInlineShapeAtomNonLitNodeConstraint(ctx *InlineShapeAtomNonLitNodeConstraintContext) {}

// EnterInlineShapeAtomLitNodeConstraint is called when production inlineShapeAtomLitNodeConstraint is entered.
func (s *BaseShExDocListener) EnterInlineShapeAtomLitNodeConstraint(ctx *InlineShapeAtomLitNodeConstraintContext) {}

// ExitInlineShapeAtomLitNodeConstraint is called when production inlineShapeAtomLitNodeConstraint is exited.
func (s *BaseShExDocListener) ExitInlineShapeAtomLitNodeConstraint(ctx *InlineShapeAtomLitNodeConstraintContext) {}

// EnterInlineShapeAtomShapeOrRef is called when production inlineShapeAtomShapeOrRef is entered.
func (s *BaseShExDocListener) EnterInlineShapeAtomShapeOrRef(ctx *InlineShapeAtomShapeOrRefContext) {}

// ExitInlineShapeAtomShapeOrRef is called when production inlineShapeAtomShapeOrRef is exited.
func (s *BaseShExDocListener) ExitInlineShapeAtomShapeOrRef(ctx *InlineShapeAtomShapeOrRefContext) {}

// EnterInlineShapeAtomShapeExpression is called when production inlineShapeAtomShapeExpression is entered.
func (s *BaseShExDocListener) EnterInlineShapeAtomShapeExpression(ctx *InlineShapeAtomShapeExpressionContext) {}

// ExitInlineShapeAtomShapeExpression is called when production inlineShapeAtomShapeExpression is exited.
func (s *BaseShExDocListener) ExitInlineShapeAtomShapeExpression(ctx *InlineShapeAtomShapeExpressionContext) {}

// EnterInlineShapeAtomAny is called when production inlineShapeAtomAny is entered.
func (s *BaseShExDocListener) EnterInlineShapeAtomAny(ctx *InlineShapeAtomAnyContext) {}

// ExitInlineShapeAtomAny is called when production inlineShapeAtomAny is exited.
func (s *BaseShExDocListener) ExitInlineShapeAtomAny(ctx *InlineShapeAtomAnyContext) {}

// EnterShapeOrRef is called when production shapeOrRef is entered.
func (s *BaseShExDocListener) EnterShapeOrRef(ctx *ShapeOrRefContext) {}

// ExitShapeOrRef is called when production shapeOrRef is exited.
func (s *BaseShExDocListener) ExitShapeOrRef(ctx *ShapeOrRefContext) {}

// EnterInlineShapeOrRef is called when production inlineShapeOrRef is entered.
func (s *BaseShExDocListener) EnterInlineShapeOrRef(ctx *InlineShapeOrRefContext) {}

// ExitInlineShapeOrRef is called when production inlineShapeOrRef is exited.
func (s *BaseShExDocListener) ExitInlineShapeOrRef(ctx *InlineShapeOrRefContext) {}

// EnterShapeRef is called when production shapeRef is entered.
func (s *BaseShExDocListener) EnterShapeRef(ctx *ShapeRefContext) {}

// ExitShapeRef is called when production shapeRef is exited.
func (s *BaseShExDocListener) ExitShapeRef(ctx *ShapeRefContext) {}

// EnterNodeConstraintLiteral is called when production nodeConstraintLiteral is entered.
func (s *BaseShExDocListener) EnterNodeConstraintLiteral(ctx *NodeConstraintLiteralContext) {}

// ExitNodeConstraintLiteral is called when production nodeConstraintLiteral is exited.
func (s *BaseShExDocListener) ExitNodeConstraintLiteral(ctx *NodeConstraintLiteralContext) {}

// EnterNodeConstraintNonLiteral is called when production nodeConstraintNonLiteral is entered.
func (s *BaseShExDocListener) EnterNodeConstraintNonLiteral(ctx *NodeConstraintNonLiteralContext) {}

// ExitNodeConstraintNonLiteral is called when production nodeConstraintNonLiteral is exited.
func (s *BaseShExDocListener) ExitNodeConstraintNonLiteral(ctx *NodeConstraintNonLiteralContext) {}

// EnterNodeConstraintDatatype is called when production nodeConstraintDatatype is entered.
func (s *BaseShExDocListener) EnterNodeConstraintDatatype(ctx *NodeConstraintDatatypeContext) {}

// ExitNodeConstraintDatatype is called when production nodeConstraintDatatype is exited.
func (s *BaseShExDocListener) ExitNodeConstraintDatatype(ctx *NodeConstraintDatatypeContext) {}

// EnterNodeConstraintValueSet is called when production nodeConstraintValueSet is entered.
func (s *BaseShExDocListener) EnterNodeConstraintValueSet(ctx *NodeConstraintValueSetContext) {}

// ExitNodeConstraintValueSet is called when production nodeConstraintValueSet is exited.
func (s *BaseShExDocListener) ExitNodeConstraintValueSet(ctx *NodeConstraintValueSetContext) {}

// EnterNodeConstraintNumericFacet is called when production nodeConstraintNumericFacet is entered.
func (s *BaseShExDocListener) EnterNodeConstraintNumericFacet(ctx *NodeConstraintNumericFacetContext) {}

// ExitNodeConstraintNumericFacet is called when production nodeConstraintNumericFacet is exited.
func (s *BaseShExDocListener) ExitNodeConstraintNumericFacet(ctx *NodeConstraintNumericFacetContext) {}

// EnterLitNodeConstraint is called when production litNodeConstraint is entered.
func (s *BaseShExDocListener) EnterLitNodeConstraint(ctx *LitNodeConstraintContext) {}

// ExitLitNodeConstraint is called when production litNodeConstraint is exited.
func (s *BaseShExDocListener) ExitLitNodeConstraint(ctx *LitNodeConstraintContext) {}

// EnterLitNodeConstraintLiteral is called when production litNodeConstraintLiteral is entered.
func (s *BaseShExDocListener) EnterLitNodeConstraintLiteral(ctx *LitNodeConstraintLiteralContext) {}

// ExitLitNodeConstraintLiteral is called when production litNodeConstraintLiteral is exited.
func (s *BaseShExDocListener) ExitLitNodeConstraintLiteral(ctx *LitNodeConstraintLiteralContext) {}

// EnterLitNodeConstraintStringFacet is called when production litNodeConstraintStringFacet is entered.
func (s *BaseShExDocListener) EnterLitNodeConstraintStringFacet(ctx *LitNodeConstraintStringFacetContext) {}

// ExitLitNodeConstraintStringFacet is called when production litNodeConstraintStringFacet is exited.
func (s *BaseShExDocListener) ExitLitNodeConstraintStringFacet(ctx *LitNodeConstraintStringFacetContext) {}

// EnterNonLitNodeConstraint is called when production nonLitNodeConstraint is entered.
func (s *BaseShExDocListener) EnterNonLitNodeConstraint(ctx *NonLitNodeConstraintContext) {}

// ExitNonLitNodeConstraint is called when production nonLitNodeConstraint is exited.
func (s *BaseShExDocListener) ExitNonLitNodeConstraint(ctx *NonLitNodeConstraintContext) {}

// EnterNonLiteralKind is called when production nonLiteralKind is entered.
func (s *BaseShExDocListener) EnterNonLiteralKind(ctx *NonLiteralKindContext) {}

// ExitNonLiteralKind is called when production nonLiteralKind is exited.
func (s *BaseShExDocListener) ExitNonLiteralKind(ctx *NonLiteralKindContext) {}

// EnterXsFacet is called when production xsFacet is entered.
func (s *BaseShExDocListener) EnterXsFacet(ctx *XsFacetContext) {}

// ExitXsFacet is called when production xsFacet is exited.
func (s *BaseShExDocListener) ExitXsFacet(ctx *XsFacetContext) {}

// EnterStringFacet is called when production stringFacet is entered.
func (s *BaseShExDocListener) EnterStringFacet(ctx *StringFacetContext) {}

// ExitStringFacet is called when production stringFacet is exited.
func (s *BaseShExDocListener) ExitStringFacet(ctx *StringFacetContext) {}

// EnterStringLength is called when production stringLength is entered.
func (s *BaseShExDocListener) EnterStringLength(ctx *StringLengthContext) {}

// ExitStringLength is called when production stringLength is exited.
func (s *BaseShExDocListener) ExitStringLength(ctx *StringLengthContext) {}

// EnterNumericFacet is called when production numericFacet is entered.
func (s *BaseShExDocListener) EnterNumericFacet(ctx *NumericFacetContext) {}

// ExitNumericFacet is called when production numericFacet is exited.
func (s *BaseShExDocListener) ExitNumericFacet(ctx *NumericFacetContext) {}

// EnterNumericRange is called when production numericRange is entered.
func (s *BaseShExDocListener) EnterNumericRange(ctx *NumericRangeContext) {}

// ExitNumericRange is called when production numericRange is exited.
func (s *BaseShExDocListener) ExitNumericRange(ctx *NumericRangeContext) {}

// EnterNumericLength is called when production numericLength is entered.
func (s *BaseShExDocListener) EnterNumericLength(ctx *NumericLengthContext) {}

// ExitNumericLength is called when production numericLength is exited.
func (s *BaseShExDocListener) ExitNumericLength(ctx *NumericLengthContext) {}

// EnterRawNumeric is called when production rawNumeric is entered.
func (s *BaseShExDocListener) EnterRawNumeric(ctx *RawNumericContext) {}

// ExitRawNumeric is called when production rawNumeric is exited.
func (s *BaseShExDocListener) ExitRawNumeric(ctx *RawNumericContext) {}

// EnterShapeDefinition is called when production shapeDefinition is entered.
func (s *BaseShExDocListener) EnterShapeDefinition(ctx *ShapeDefinitionContext) {}

// ExitShapeDefinition is called when production shapeDefinition is exited.
func (s *BaseShExDocListener) ExitShapeDefinition(ctx *ShapeDefinitionContext) {}

// EnterInlineShapeDefinition is called when production inlineShapeDefinition is entered.
func (s *BaseShExDocListener) EnterInlineShapeDefinition(ctx *InlineShapeDefinitionContext) {}

// ExitInlineShapeDefinition is called when production inlineShapeDefinition is exited.
func (s *BaseShExDocListener) ExitInlineShapeDefinition(ctx *InlineShapeDefinitionContext) {}

// EnterQualifier is called when production qualifier is entered.
func (s *BaseShExDocListener) EnterQualifier(ctx *QualifierContext) {}

// ExitQualifier is called when production qualifier is exited.
func (s *BaseShExDocListener) ExitQualifier(ctx *QualifierContext) {}

// EnterExtraPropertySet is called when production extraPropertySet is entered.
func (s *BaseShExDocListener) EnterExtraPropertySet(ctx *ExtraPropertySetContext) {}

// ExitExtraPropertySet is called when production extraPropertySet is exited.
func (s *BaseShExDocListener) ExitExtraPropertySet(ctx *ExtraPropertySetContext) {}

// EnterTripleExpression is called when production tripleExpression is entered.
func (s *BaseShExDocListener) EnterTripleExpression(ctx *TripleExpressionContext) {}

// ExitTripleExpression is called when production tripleExpression is exited.
func (s *BaseShExDocListener) ExitTripleExpression(ctx *TripleExpressionContext) {}

// EnterOneOfTripleExpr is called when production oneOfTripleExpr is entered.
func (s *BaseShExDocListener) EnterOneOfTripleExpr(ctx *OneOfTripleExprContext) {}

// ExitOneOfTripleExpr is called when production oneOfTripleExpr is exited.
func (s *BaseShExDocListener) ExitOneOfTripleExpr(ctx *OneOfTripleExprContext) {}

// EnterMultiElementOneOf is called when production multiElementOneOf is entered.
func (s *BaseShExDocListener) EnterMultiElementOneOf(ctx *MultiElementOneOfContext) {}

// ExitMultiElementOneOf is called when production multiElementOneOf is exited.
func (s *BaseShExDocListener) ExitMultiElementOneOf(ctx *MultiElementOneOfContext) {}

// EnterGroupTripleExpr is called when production groupTripleExpr is entered.
func (s *BaseShExDocListener) EnterGroupTripleExpr(ctx *GroupTripleExprContext) {}

// ExitGroupTripleExpr is called when production groupTripleExpr is exited.
func (s *BaseShExDocListener) ExitGroupTripleExpr(ctx *GroupTripleExprContext) {}

// EnterSingleElementGroup is called when production singleElementGroup is entered.
func (s *BaseShExDocListener) EnterSingleElementGroup(ctx *SingleElementGroupContext) {}

// ExitSingleElementGroup is called when production singleElementGroup is exited.
func (s *BaseShExDocListener) ExitSingleElementGroup(ctx *SingleElementGroupContext) {}

// EnterMultiElementGroup is called when production multiElementGroup is entered.
func (s *BaseShExDocListener) EnterMultiElementGroup(ctx *MultiElementGroupContext) {}

// ExitMultiElementGroup is called when production multiElementGroup is exited.
func (s *BaseShExDocListener) ExitMultiElementGroup(ctx *MultiElementGroupContext) {}

// EnterUnaryTripleExpr is called when production unaryTripleExpr is entered.
func (s *BaseShExDocListener) EnterUnaryTripleExpr(ctx *UnaryTripleExprContext) {}

// ExitUnaryTripleExpr is called when production unaryTripleExpr is exited.
func (s *BaseShExDocListener) ExitUnaryTripleExpr(ctx *UnaryTripleExprContext) {}

// EnterBracketedTripleExpr is called when production bracketedTripleExpr is entered.
func (s *BaseShExDocListener) EnterBracketedTripleExpr(ctx *BracketedTripleExprContext) {}

// ExitBracketedTripleExpr is called when production bracketedTripleExpr is exited.
func (s *BaseShExDocListener) ExitBracketedTripleExpr(ctx *BracketedTripleExprContext) {}

// EnterTripleConstraint is called when production tripleConstraint is entered.
func (s *BaseShExDocListener) EnterTripleConstraint(ctx *TripleConstraintContext) {}

// ExitTripleConstraint is called when production tripleConstraint is exited.
func (s *BaseShExDocListener) ExitTripleConstraint(ctx *TripleConstraintContext) {}

// EnterStarCardinality is called when production starCardinality is entered.
func (s *BaseShExDocListener) EnterStarCardinality(ctx *StarCardinalityContext) {}

// ExitStarCardinality is called when production starCardinality is exited.
func (s *BaseShExDocListener) ExitStarCardinality(ctx *StarCardinalityContext) {}

// EnterPlusCardinality is called when production plusCardinality is entered.
func (s *BaseShExDocListener) EnterPlusCardinality(ctx *PlusCardinalityContext) {}

// ExitPlusCardinality is called when production plusCardinality is exited.
func (s *BaseShExDocListener) ExitPlusCardinality(ctx *PlusCardinalityContext) {}

// EnterOptionalCardinality is called when production optionalCardinality is entered.
func (s *BaseShExDocListener) EnterOptionalCardinality(ctx *OptionalCardinalityContext) {}

// ExitOptionalCardinality is called when production optionalCardinality is exited.
func (s *BaseShExDocListener) ExitOptionalCardinality(ctx *OptionalCardinalityContext) {}

// EnterRepeatCardinality is called when production repeatCardinality is entered.
func (s *BaseShExDocListener) EnterRepeatCardinality(ctx *RepeatCardinalityContext) {}

// ExitRepeatCardinality is called when production repeatCardinality is exited.
func (s *BaseShExDocListener) ExitRepeatCardinality(ctx *RepeatCardinalityContext) {}

// EnterExactRange is called when production exactRange is entered.
func (s *BaseShExDocListener) EnterExactRange(ctx *ExactRangeContext) {}

// ExitExactRange is called when production exactRange is exited.
func (s *BaseShExDocListener) ExitExactRange(ctx *ExactRangeContext) {}

// EnterMinMaxRange is called when production minMaxRange is entered.
func (s *BaseShExDocListener) EnterMinMaxRange(ctx *MinMaxRangeContext) {}

// ExitMinMaxRange is called when production minMaxRange is exited.
func (s *BaseShExDocListener) ExitMinMaxRange(ctx *MinMaxRangeContext) {}

// EnterSenseFlags is called when production senseFlags is entered.
func (s *BaseShExDocListener) EnterSenseFlags(ctx *SenseFlagsContext) {}

// ExitSenseFlags is called when production senseFlags is exited.
func (s *BaseShExDocListener) ExitSenseFlags(ctx *SenseFlagsContext) {}

// EnterValueSet is called when production valueSet is entered.
func (s *BaseShExDocListener) EnterValueSet(ctx *ValueSetContext) {}

// ExitValueSet is called when production valueSet is exited.
func (s *BaseShExDocListener) ExitValueSet(ctx *ValueSetContext) {}

// EnterValueSetValue is called when production valueSetValue is entered.
func (s *BaseShExDocListener) EnterValueSetValue(ctx *ValueSetValueContext) {}

// ExitValueSetValue is called when production valueSetValue is exited.
func (s *BaseShExDocListener) ExitValueSetValue(ctx *ValueSetValueContext) {}

// EnterIriRange is called when production iriRange is entered.
func (s *BaseShExDocListener) EnterIriRange(ctx *IriRangeContext) {}

// ExitIriRange is called when production iriRange is exited.
func (s *BaseShExDocListener) ExitIriRange(ctx *IriRangeContext) {}

// EnterIriExclusion is called when production iriExclusion is entered.
func (s *BaseShExDocListener) EnterIriExclusion(ctx *IriExclusionContext) {}

// ExitIriExclusion is called when production iriExclusion is exited.
func (s *BaseShExDocListener) ExitIriExclusion(ctx *IriExclusionContext) {}

// EnterLiteralRange is called when production literalRange is entered.
func (s *BaseShExDocListener) EnterLiteralRange(ctx *LiteralRangeContext) {}

// ExitLiteralRange is called when production literalRange is exited.
func (s *BaseShExDocListener) ExitLiteralRange(ctx *LiteralRangeContext) {}

// EnterLiteralExclusion is called when production literalExclusion is entered.
func (s *BaseShExDocListener) EnterLiteralExclusion(ctx *LiteralExclusionContext) {}

// ExitLiteralExclusion is called when production literalExclusion is exited.
func (s *BaseShExDocListener) ExitLiteralExclusion(ctx *LiteralExclusionContext) {}

// EnterLanguageRangeFull is called when production languageRangeFull is entered.
func (s *BaseShExDocListener) EnterLanguageRangeFull(ctx *LanguageRangeFullContext) {}

// ExitLanguageRangeFull is called when production languageRangeFull is exited.
func (s *BaseShExDocListener) ExitLanguageRangeFull(ctx *LanguageRangeFullContext) {}

// EnterLanguageRangeAt is called when production languageRangeAt is entered.
func (s *BaseShExDocListener) EnterLanguageRangeAt(ctx *LanguageRangeAtContext) {}

// ExitLanguageRangeAt is called when production languageRangeAt is exited.
func (s *BaseShExDocListener) ExitLanguageRangeAt(ctx *LanguageRangeAtContext) {}

// EnterLanguageExclusion is called when production languageExclusion is entered.
func (s *BaseShExDocListener) EnterLanguageExclusion(ctx *LanguageExclusionContext) {}

// ExitLanguageExclusion is called when production languageExclusion is exited.
func (s *BaseShExDocListener) ExitLanguageExclusion(ctx *LanguageExclusionContext) {}

// EnterInclude is called when production include is entered.
func (s *BaseShExDocListener) EnterInclude(ctx *IncludeContext) {}

// ExitInclude is called when production include is exited.
func (s *BaseShExDocListener) ExitInclude(ctx *IncludeContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseShExDocListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseShExDocListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterSemanticAction is called when production semanticAction is entered.
func (s *BaseShExDocListener) EnterSemanticAction(ctx *SemanticActionContext) {}

// ExitSemanticAction is called when production semanticAction is exited.
func (s *BaseShExDocListener) ExitSemanticAction(ctx *SemanticActionContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseShExDocListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseShExDocListener) ExitLiteral(ctx *LiteralContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseShExDocListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseShExDocListener) ExitPredicate(ctx *PredicateContext) {}

// EnterRdfType is called when production rdfType is entered.
func (s *BaseShExDocListener) EnterRdfType(ctx *RdfTypeContext) {}

// ExitRdfType is called when production rdfType is exited.
func (s *BaseShExDocListener) ExitRdfType(ctx *RdfTypeContext) {}

// EnterDatatype is called when production datatype is entered.
func (s *BaseShExDocListener) EnterDatatype(ctx *DatatypeContext) {}

// ExitDatatype is called when production datatype is exited.
func (s *BaseShExDocListener) ExitDatatype(ctx *DatatypeContext) {}

// EnterShapeExprLabel is called when production shapeExprLabel is entered.
func (s *BaseShExDocListener) EnterShapeExprLabel(ctx *ShapeExprLabelContext) {}

// ExitShapeExprLabel is called when production shapeExprLabel is exited.
func (s *BaseShExDocListener) ExitShapeExprLabel(ctx *ShapeExprLabelContext) {}

// EnterTripleExprLabel is called when production tripleExprLabel is entered.
func (s *BaseShExDocListener) EnterTripleExprLabel(ctx *TripleExprLabelContext) {}

// ExitTripleExprLabel is called when production tripleExprLabel is exited.
func (s *BaseShExDocListener) ExitTripleExprLabel(ctx *TripleExprLabelContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseShExDocListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseShExDocListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterRdfLiteral is called when production rdfLiteral is entered.
func (s *BaseShExDocListener) EnterRdfLiteral(ctx *RdfLiteralContext) {}

// ExitRdfLiteral is called when production rdfLiteral is exited.
func (s *BaseShExDocListener) ExitRdfLiteral(ctx *RdfLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseShExDocListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseShExDocListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterRdfString is called when production rdfString is entered.
func (s *BaseShExDocListener) EnterRdfString(ctx *RdfStringContext) {}

// ExitRdfString is called when production rdfString is exited.
func (s *BaseShExDocListener) ExitRdfString(ctx *RdfStringContext) {}

// EnterIri is called when production iri is entered.
func (s *BaseShExDocListener) EnterIri(ctx *IriContext) {}

// ExitIri is called when production iri is exited.
func (s *BaseShExDocListener) ExitIri(ctx *IriContext) {}

// EnterPrefixedName is called when production prefixedName is entered.
func (s *BaseShExDocListener) EnterPrefixedName(ctx *PrefixedNameContext) {}

// ExitPrefixedName is called when production prefixedName is exited.
func (s *BaseShExDocListener) ExitPrefixedName(ctx *PrefixedNameContext) {}

// EnterBlankNode is called when production blankNode is entered.
func (s *BaseShExDocListener) EnterBlankNode(ctx *BlankNodeContext) {}

// ExitBlankNode is called when production blankNode is exited.
func (s *BaseShExDocListener) ExitBlankNode(ctx *BlankNodeContext) {}

// EnterExtension is called when production extension is entered.
func (s *BaseShExDocListener) EnterExtension(ctx *ExtensionContext) {}

// ExitExtension is called when production extension is exited.
func (s *BaseShExDocListener) ExitExtension(ctx *ExtensionContext) {}

// EnterRestrictions is called when production restrictions is entered.
func (s *BaseShExDocListener) EnterRestrictions(ctx *RestrictionsContext) {}

// ExitRestrictions is called when production restrictions is exited.
func (s *BaseShExDocListener) ExitRestrictions(ctx *RestrictionsContext) {}
