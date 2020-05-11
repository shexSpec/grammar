// Code generated from ShExDoc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ShExDoc

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseShExDocVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseShExDocVisitor) VisitShExDoc(ctx *ShExDocContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitDirective(ctx *DirectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitBaseDecl(ctx *BaseDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitPrefixDecl(ctx *PrefixDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitNotStartAction(ctx *NotStartActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitStartActions(ctx *StartActionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeExprDecl(ctx *ShapeExprDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeExpression(ctx *ShapeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitInlineShapeExpression(ctx *InlineShapeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeOr(ctx *ShapeOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitInlineShapeOr(ctx *InlineShapeOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeAnd(ctx *ShapeAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitInlineShapeAnd(ctx *InlineShapeAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeNot(ctx *ShapeNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitInlineShapeNot(ctx *InlineShapeNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeAtomNonLitNodeConstraint(ctx *ShapeAtomNonLitNodeConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeAtomLitNodeConstraint(ctx *ShapeAtomLitNodeConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeAtomShapeOrRef(ctx *ShapeAtomShapeOrRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeAtomShapeExpression(ctx *ShapeAtomShapeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeAtomAny(ctx *ShapeAtomAnyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitInlineShapeAtomNonLitNodeConstraint(ctx *InlineShapeAtomNonLitNodeConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitInlineShapeAtomLitNodeConstraint(ctx *InlineShapeAtomLitNodeConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitInlineShapeAtomShapeOrRef(ctx *InlineShapeAtomShapeOrRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitInlineShapeAtomShapeExpression(ctx *InlineShapeAtomShapeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitInlineShapeAtomAny(ctx *InlineShapeAtomAnyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeOrRef(ctx *ShapeOrRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitInlineShapeOrRef(ctx *InlineShapeOrRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeRef(ctx *ShapeRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitNodeConstraintLiteral(ctx *NodeConstraintLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitNodeConstraintNonLiteral(ctx *NodeConstraintNonLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitNodeConstraintDatatype(ctx *NodeConstraintDatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitNodeConstraintValueSet(ctx *NodeConstraintValueSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitNodeConstraintNumericFacet(ctx *NodeConstraintNumericFacetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitLitNodeConstraint(ctx *LitNodeConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitLitNodeConstraintLiteral(ctx *LitNodeConstraintLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitLitNodeConstraintStringFacet(ctx *LitNodeConstraintStringFacetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitNonLitNodeConstraint(ctx *NonLitNodeConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitNonLiteralKind(ctx *NonLiteralKindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitXsFacet(ctx *XsFacetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitStringFacet(ctx *StringFacetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitStringLength(ctx *StringLengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitNumericFacet(ctx *NumericFacetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitNumericRange(ctx *NumericRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitNumericLength(ctx *NumericLengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitRawNumeric(ctx *RawNumericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeDefinition(ctx *ShapeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitInlineShapeDefinition(ctx *InlineShapeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitQualifier(ctx *QualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitExtraPropertySet(ctx *ExtraPropertySetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitTripleExpression(ctx *TripleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitOneOfTripleExpr(ctx *OneOfTripleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitMultiElementOneOf(ctx *MultiElementOneOfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitGroupTripleExpr(ctx *GroupTripleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitSingleElementGroup(ctx *SingleElementGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitMultiElementGroup(ctx *MultiElementGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitUnaryTripleExpr(ctx *UnaryTripleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitBracketedTripleExpr(ctx *BracketedTripleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitTripleConstraint(ctx *TripleConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitStarCardinality(ctx *StarCardinalityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitPlusCardinality(ctx *PlusCardinalityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitOptionalCardinality(ctx *OptionalCardinalityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitRepeatCardinality(ctx *RepeatCardinalityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitExactRange(ctx *ExactRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitMinMaxRange(ctx *MinMaxRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitSenseFlags(ctx *SenseFlagsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitValueSet(ctx *ValueSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitValueSetValue(ctx *ValueSetValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitIriRange(ctx *IriRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitIriExclusion(ctx *IriExclusionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitLiteralRange(ctx *LiteralRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitLiteralExclusion(ctx *LiteralExclusionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitLanguageRangeFull(ctx *LanguageRangeFullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitLanguageRangeAt(ctx *LanguageRangeAtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitLanguageExclusion(ctx *LanguageExclusionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitInclude(ctx *IncludeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitSemanticAction(ctx *SemanticActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitRdfType(ctx *RdfTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitDatatype(ctx *DatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitShapeExprLabel(ctx *ShapeExprLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitTripleExprLabel(ctx *TripleExprLabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitRdfLiteral(ctx *RdfLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitRdfString(ctx *RdfStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitIri(ctx *IriContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitPrefixedName(ctx *PrefixedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitBlankNode(ctx *BlankNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitExtension(ctx *ExtensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseShExDocVisitor) VisitRestrictions(ctx *RestrictionsContext) interface{} {
	return v.VisitChildren(ctx)
}
