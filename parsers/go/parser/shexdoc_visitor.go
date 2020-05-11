// Code generated from ShExDoc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ShExDoc

import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by ShExDocParser.
type ShExDocVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ShExDocParser#shExDoc.
	VisitShExDoc(ctx *ShExDocContext) interface{}

	// Visit a parse tree produced by ShExDocParser#directive.
	VisitDirective(ctx *DirectiveContext) interface{}

	// Visit a parse tree produced by ShExDocParser#baseDecl.
	VisitBaseDecl(ctx *BaseDeclContext) interface{}

	// Visit a parse tree produced by ShExDocParser#prefixDecl.
	VisitPrefixDecl(ctx *PrefixDeclContext) interface{}

	// Visit a parse tree produced by ShExDocParser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by ShExDocParser#notStartAction.
	VisitNotStartAction(ctx *NotStartActionContext) interface{}

	// Visit a parse tree produced by ShExDocParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by ShExDocParser#startActions.
	VisitStartActions(ctx *StartActionsContext) interface{}

	// Visit a parse tree produced by ShExDocParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeExprDecl.
	VisitShapeExprDecl(ctx *ShapeExprDeclContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeExpression.
	VisitShapeExpression(ctx *ShapeExpressionContext) interface{}

	// Visit a parse tree produced by ShExDocParser#inlineShapeExpression.
	VisitInlineShapeExpression(ctx *InlineShapeExpressionContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeOr.
	VisitShapeOr(ctx *ShapeOrContext) interface{}

	// Visit a parse tree produced by ShExDocParser#inlineShapeOr.
	VisitInlineShapeOr(ctx *InlineShapeOrContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeAnd.
	VisitShapeAnd(ctx *ShapeAndContext) interface{}

	// Visit a parse tree produced by ShExDocParser#inlineShapeAnd.
	VisitInlineShapeAnd(ctx *InlineShapeAndContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeNot.
	VisitShapeNot(ctx *ShapeNotContext) interface{}

	// Visit a parse tree produced by ShExDocParser#inlineShapeNot.
	VisitInlineShapeNot(ctx *InlineShapeNotContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeAtomNonLitNodeConstraint.
	VisitShapeAtomNonLitNodeConstraint(ctx *ShapeAtomNonLitNodeConstraintContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeAtomLitNodeConstraint.
	VisitShapeAtomLitNodeConstraint(ctx *ShapeAtomLitNodeConstraintContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeAtomShapeOrRef.
	VisitShapeAtomShapeOrRef(ctx *ShapeAtomShapeOrRefContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeAtomShapeExpression.
	VisitShapeAtomShapeExpression(ctx *ShapeAtomShapeExpressionContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeAtomAny.
	VisitShapeAtomAny(ctx *ShapeAtomAnyContext) interface{}

	// Visit a parse tree produced by ShExDocParser#inlineShapeAtomNonLitNodeConstraint.
	VisitInlineShapeAtomNonLitNodeConstraint(ctx *InlineShapeAtomNonLitNodeConstraintContext) interface{}

	// Visit a parse tree produced by ShExDocParser#inlineShapeAtomLitNodeConstraint.
	VisitInlineShapeAtomLitNodeConstraint(ctx *InlineShapeAtomLitNodeConstraintContext) interface{}

	// Visit a parse tree produced by ShExDocParser#inlineShapeAtomShapeOrRef.
	VisitInlineShapeAtomShapeOrRef(ctx *InlineShapeAtomShapeOrRefContext) interface{}

	// Visit a parse tree produced by ShExDocParser#inlineShapeAtomShapeExpression.
	VisitInlineShapeAtomShapeExpression(ctx *InlineShapeAtomShapeExpressionContext) interface{}

	// Visit a parse tree produced by ShExDocParser#inlineShapeAtomAny.
	VisitInlineShapeAtomAny(ctx *InlineShapeAtomAnyContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeOrRef.
	VisitShapeOrRef(ctx *ShapeOrRefContext) interface{}

	// Visit a parse tree produced by ShExDocParser#inlineShapeOrRef.
	VisitInlineShapeOrRef(ctx *InlineShapeOrRefContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeRef.
	VisitShapeRef(ctx *ShapeRefContext) interface{}

	// Visit a parse tree produced by ShExDocParser#nodeConstraintLiteral.
	VisitNodeConstraintLiteral(ctx *NodeConstraintLiteralContext) interface{}

	// Visit a parse tree produced by ShExDocParser#nodeConstraintNonLiteral.
	VisitNodeConstraintNonLiteral(ctx *NodeConstraintNonLiteralContext) interface{}

	// Visit a parse tree produced by ShExDocParser#nodeConstraintDatatype.
	VisitNodeConstraintDatatype(ctx *NodeConstraintDatatypeContext) interface{}

	// Visit a parse tree produced by ShExDocParser#nodeConstraintValueSet.
	VisitNodeConstraintValueSet(ctx *NodeConstraintValueSetContext) interface{}

	// Visit a parse tree produced by ShExDocParser#nodeConstraintNumericFacet.
	VisitNodeConstraintNumericFacet(ctx *NodeConstraintNumericFacetContext) interface{}

	// Visit a parse tree produced by ShExDocParser#litNodeConstraint.
	VisitLitNodeConstraint(ctx *LitNodeConstraintContext) interface{}

	// Visit a parse tree produced by ShExDocParser#litNodeConstraintLiteral.
	VisitLitNodeConstraintLiteral(ctx *LitNodeConstraintLiteralContext) interface{}

	// Visit a parse tree produced by ShExDocParser#litNodeConstraintStringFacet.
	VisitLitNodeConstraintStringFacet(ctx *LitNodeConstraintStringFacetContext) interface{}

	// Visit a parse tree produced by ShExDocParser#nonLitNodeConstraint.
	VisitNonLitNodeConstraint(ctx *NonLitNodeConstraintContext) interface{}

	// Visit a parse tree produced by ShExDocParser#nonLiteralKind.
	VisitNonLiteralKind(ctx *NonLiteralKindContext) interface{}

	// Visit a parse tree produced by ShExDocParser#xsFacet.
	VisitXsFacet(ctx *XsFacetContext) interface{}

	// Visit a parse tree produced by ShExDocParser#stringFacet.
	VisitStringFacet(ctx *StringFacetContext) interface{}

	// Visit a parse tree produced by ShExDocParser#stringLength.
	VisitStringLength(ctx *StringLengthContext) interface{}

	// Visit a parse tree produced by ShExDocParser#numericFacet.
	VisitNumericFacet(ctx *NumericFacetContext) interface{}

	// Visit a parse tree produced by ShExDocParser#numericRange.
	VisitNumericRange(ctx *NumericRangeContext) interface{}

	// Visit a parse tree produced by ShExDocParser#numericLength.
	VisitNumericLength(ctx *NumericLengthContext) interface{}

	// Visit a parse tree produced by ShExDocParser#rawNumeric.
	VisitRawNumeric(ctx *RawNumericContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeDefinition.
	VisitShapeDefinition(ctx *ShapeDefinitionContext) interface{}

	// Visit a parse tree produced by ShExDocParser#inlineShapeDefinition.
	VisitInlineShapeDefinition(ctx *InlineShapeDefinitionContext) interface{}

	// Visit a parse tree produced by ShExDocParser#qualifier.
	VisitQualifier(ctx *QualifierContext) interface{}

	// Visit a parse tree produced by ShExDocParser#extraPropertySet.
	VisitExtraPropertySet(ctx *ExtraPropertySetContext) interface{}

	// Visit a parse tree produced by ShExDocParser#tripleExpression.
	VisitTripleExpression(ctx *TripleExpressionContext) interface{}

	// Visit a parse tree produced by ShExDocParser#oneOfTripleExpr.
	VisitOneOfTripleExpr(ctx *OneOfTripleExprContext) interface{}

	// Visit a parse tree produced by ShExDocParser#multiElementOneOf.
	VisitMultiElementOneOf(ctx *MultiElementOneOfContext) interface{}

	// Visit a parse tree produced by ShExDocParser#groupTripleExpr.
	VisitGroupTripleExpr(ctx *GroupTripleExprContext) interface{}

	// Visit a parse tree produced by ShExDocParser#singleElementGroup.
	VisitSingleElementGroup(ctx *SingleElementGroupContext) interface{}

	// Visit a parse tree produced by ShExDocParser#multiElementGroup.
	VisitMultiElementGroup(ctx *MultiElementGroupContext) interface{}

	// Visit a parse tree produced by ShExDocParser#unaryTripleExpr.
	VisitUnaryTripleExpr(ctx *UnaryTripleExprContext) interface{}

	// Visit a parse tree produced by ShExDocParser#bracketedTripleExpr.
	VisitBracketedTripleExpr(ctx *BracketedTripleExprContext) interface{}

	// Visit a parse tree produced by ShExDocParser#tripleConstraint.
	VisitTripleConstraint(ctx *TripleConstraintContext) interface{}

	// Visit a parse tree produced by ShExDocParser#starCardinality.
	VisitStarCardinality(ctx *StarCardinalityContext) interface{}

	// Visit a parse tree produced by ShExDocParser#plusCardinality.
	VisitPlusCardinality(ctx *PlusCardinalityContext) interface{}

	// Visit a parse tree produced by ShExDocParser#optionalCardinality.
	VisitOptionalCardinality(ctx *OptionalCardinalityContext) interface{}

	// Visit a parse tree produced by ShExDocParser#repeatCardinality.
	VisitRepeatCardinality(ctx *RepeatCardinalityContext) interface{}

	// Visit a parse tree produced by ShExDocParser#exactRange.
	VisitExactRange(ctx *ExactRangeContext) interface{}

	// Visit a parse tree produced by ShExDocParser#minMaxRange.
	VisitMinMaxRange(ctx *MinMaxRangeContext) interface{}

	// Visit a parse tree produced by ShExDocParser#senseFlags.
	VisitSenseFlags(ctx *SenseFlagsContext) interface{}

	// Visit a parse tree produced by ShExDocParser#valueSet.
	VisitValueSet(ctx *ValueSetContext) interface{}

	// Visit a parse tree produced by ShExDocParser#valueSetValue.
	VisitValueSetValue(ctx *ValueSetValueContext) interface{}

	// Visit a parse tree produced by ShExDocParser#iriRange.
	VisitIriRange(ctx *IriRangeContext) interface{}

	// Visit a parse tree produced by ShExDocParser#iriExclusion.
	VisitIriExclusion(ctx *IriExclusionContext) interface{}

	// Visit a parse tree produced by ShExDocParser#literalRange.
	VisitLiteralRange(ctx *LiteralRangeContext) interface{}

	// Visit a parse tree produced by ShExDocParser#literalExclusion.
	VisitLiteralExclusion(ctx *LiteralExclusionContext) interface{}

	// Visit a parse tree produced by ShExDocParser#languageRangeFull.
	VisitLanguageRangeFull(ctx *LanguageRangeFullContext) interface{}

	// Visit a parse tree produced by ShExDocParser#languageRangeAt.
	VisitLanguageRangeAt(ctx *LanguageRangeAtContext) interface{}

	// Visit a parse tree produced by ShExDocParser#languageExclusion.
	VisitLanguageExclusion(ctx *LanguageExclusionContext) interface{}

	// Visit a parse tree produced by ShExDocParser#include.
	VisitInclude(ctx *IncludeContext) interface{}

	// Visit a parse tree produced by ShExDocParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by ShExDocParser#semanticAction.
	VisitSemanticAction(ctx *SemanticActionContext) interface{}

	// Visit a parse tree produced by ShExDocParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ShExDocParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by ShExDocParser#rdfType.
	VisitRdfType(ctx *RdfTypeContext) interface{}

	// Visit a parse tree produced by ShExDocParser#datatype.
	VisitDatatype(ctx *DatatypeContext) interface{}

	// Visit a parse tree produced by ShExDocParser#shapeExprLabel.
	VisitShapeExprLabel(ctx *ShapeExprLabelContext) interface{}

	// Visit a parse tree produced by ShExDocParser#tripleExprLabel.
	VisitTripleExprLabel(ctx *TripleExprLabelContext) interface{}

	// Visit a parse tree produced by ShExDocParser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by ShExDocParser#rdfLiteral.
	VisitRdfLiteral(ctx *RdfLiteralContext) interface{}

	// Visit a parse tree produced by ShExDocParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by ShExDocParser#rdfString.
	VisitRdfString(ctx *RdfStringContext) interface{}

	// Visit a parse tree produced by ShExDocParser#iri.
	VisitIri(ctx *IriContext) interface{}

	// Visit a parse tree produced by ShExDocParser#prefixedName.
	VisitPrefixedName(ctx *PrefixedNameContext) interface{}

	// Visit a parse tree produced by ShExDocParser#blankNode.
	VisitBlankNode(ctx *BlankNodeContext) interface{}

	// Visit a parse tree produced by ShExDocParser#extension.
	VisitExtension(ctx *ExtensionContext) interface{}

	// Visit a parse tree produced by ShExDocParser#restrictions.
	VisitRestrictions(ctx *RestrictionsContext) interface{}

}