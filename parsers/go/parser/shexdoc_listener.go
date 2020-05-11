// Code generated from ShExDoc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ShExDoc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ShExDocListener is a complete listener for a parse tree produced by ShExDocParser.
type ShExDocListener interface {
	antlr.ParseTreeListener

	// EnterShExDoc is called when entering the shExDoc production.
	EnterShExDoc(c *ShExDocContext)

	// EnterDirective is called when entering the directive production.
	EnterDirective(c *DirectiveContext)

	// EnterBaseDecl is called when entering the baseDecl production.
	EnterBaseDecl(c *BaseDeclContext)

	// EnterPrefixDecl is called when entering the prefixDecl production.
	EnterPrefixDecl(c *PrefixDeclContext)

	// EnterImportDecl is called when entering the importDecl production.
	EnterImportDecl(c *ImportDeclContext)

	// EnterNotStartAction is called when entering the notStartAction production.
	EnterNotStartAction(c *NotStartActionContext)

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterStartActions is called when entering the startActions production.
	EnterStartActions(c *StartActionsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterShapeExprDecl is called when entering the shapeExprDecl production.
	EnterShapeExprDecl(c *ShapeExprDeclContext)

	// EnterShapeExpression is called when entering the shapeExpression production.
	EnterShapeExpression(c *ShapeExpressionContext)

	// EnterInlineShapeExpression is called when entering the inlineShapeExpression production.
	EnterInlineShapeExpression(c *InlineShapeExpressionContext)

	// EnterShapeOr is called when entering the shapeOr production.
	EnterShapeOr(c *ShapeOrContext)

	// EnterInlineShapeOr is called when entering the inlineShapeOr production.
	EnterInlineShapeOr(c *InlineShapeOrContext)

	// EnterShapeAnd is called when entering the shapeAnd production.
	EnterShapeAnd(c *ShapeAndContext)

	// EnterInlineShapeAnd is called when entering the inlineShapeAnd production.
	EnterInlineShapeAnd(c *InlineShapeAndContext)

	// EnterShapeNot is called when entering the shapeNot production.
	EnterShapeNot(c *ShapeNotContext)

	// EnterInlineShapeNot is called when entering the inlineShapeNot production.
	EnterInlineShapeNot(c *InlineShapeNotContext)

	// EnterShapeAtomNonLitNodeConstraint is called when entering the shapeAtomNonLitNodeConstraint production.
	EnterShapeAtomNonLitNodeConstraint(c *ShapeAtomNonLitNodeConstraintContext)

	// EnterShapeAtomLitNodeConstraint is called when entering the shapeAtomLitNodeConstraint production.
	EnterShapeAtomLitNodeConstraint(c *ShapeAtomLitNodeConstraintContext)

	// EnterShapeAtomShapeOrRef is called when entering the shapeAtomShapeOrRef production.
	EnterShapeAtomShapeOrRef(c *ShapeAtomShapeOrRefContext)

	// EnterShapeAtomShapeExpression is called when entering the shapeAtomShapeExpression production.
	EnterShapeAtomShapeExpression(c *ShapeAtomShapeExpressionContext)

	// EnterShapeAtomAny is called when entering the shapeAtomAny production.
	EnterShapeAtomAny(c *ShapeAtomAnyContext)

	// EnterInlineShapeAtomNonLitNodeConstraint is called when entering the inlineShapeAtomNonLitNodeConstraint production.
	EnterInlineShapeAtomNonLitNodeConstraint(c *InlineShapeAtomNonLitNodeConstraintContext)

	// EnterInlineShapeAtomLitNodeConstraint is called when entering the inlineShapeAtomLitNodeConstraint production.
	EnterInlineShapeAtomLitNodeConstraint(c *InlineShapeAtomLitNodeConstraintContext)

	// EnterInlineShapeAtomShapeOrRef is called when entering the inlineShapeAtomShapeOrRef production.
	EnterInlineShapeAtomShapeOrRef(c *InlineShapeAtomShapeOrRefContext)

	// EnterInlineShapeAtomShapeExpression is called when entering the inlineShapeAtomShapeExpression production.
	EnterInlineShapeAtomShapeExpression(c *InlineShapeAtomShapeExpressionContext)

	// EnterInlineShapeAtomAny is called when entering the inlineShapeAtomAny production.
	EnterInlineShapeAtomAny(c *InlineShapeAtomAnyContext)

	// EnterShapeOrRef is called when entering the shapeOrRef production.
	EnterShapeOrRef(c *ShapeOrRefContext)

	// EnterInlineShapeOrRef is called when entering the inlineShapeOrRef production.
	EnterInlineShapeOrRef(c *InlineShapeOrRefContext)

	// EnterShapeRef is called when entering the shapeRef production.
	EnterShapeRef(c *ShapeRefContext)

	// EnterNodeConstraintLiteral is called when entering the nodeConstraintLiteral production.
	EnterNodeConstraintLiteral(c *NodeConstraintLiteralContext)

	// EnterNodeConstraintNonLiteral is called when entering the nodeConstraintNonLiteral production.
	EnterNodeConstraintNonLiteral(c *NodeConstraintNonLiteralContext)

	// EnterNodeConstraintDatatype is called when entering the nodeConstraintDatatype production.
	EnterNodeConstraintDatatype(c *NodeConstraintDatatypeContext)

	// EnterNodeConstraintValueSet is called when entering the nodeConstraintValueSet production.
	EnterNodeConstraintValueSet(c *NodeConstraintValueSetContext)

	// EnterNodeConstraintNumericFacet is called when entering the nodeConstraintNumericFacet production.
	EnterNodeConstraintNumericFacet(c *NodeConstraintNumericFacetContext)

	// EnterLitNodeConstraint is called when entering the litNodeConstraint production.
	EnterLitNodeConstraint(c *LitNodeConstraintContext)

	// EnterLitNodeConstraintLiteral is called when entering the litNodeConstraintLiteral production.
	EnterLitNodeConstraintLiteral(c *LitNodeConstraintLiteralContext)

	// EnterLitNodeConstraintStringFacet is called when entering the litNodeConstraintStringFacet production.
	EnterLitNodeConstraintStringFacet(c *LitNodeConstraintStringFacetContext)

	// EnterNonLitNodeConstraint is called when entering the nonLitNodeConstraint production.
	EnterNonLitNodeConstraint(c *NonLitNodeConstraintContext)

	// EnterNonLiteralKind is called when entering the nonLiteralKind production.
	EnterNonLiteralKind(c *NonLiteralKindContext)

	// EnterXsFacet is called when entering the xsFacet production.
	EnterXsFacet(c *XsFacetContext)

	// EnterStringFacet is called when entering the stringFacet production.
	EnterStringFacet(c *StringFacetContext)

	// EnterStringLength is called when entering the stringLength production.
	EnterStringLength(c *StringLengthContext)

	// EnterNumericFacet is called when entering the numericFacet production.
	EnterNumericFacet(c *NumericFacetContext)

	// EnterNumericRange is called when entering the numericRange production.
	EnterNumericRange(c *NumericRangeContext)

	// EnterNumericLength is called when entering the numericLength production.
	EnterNumericLength(c *NumericLengthContext)

	// EnterRawNumeric is called when entering the rawNumeric production.
	EnterRawNumeric(c *RawNumericContext)

	// EnterShapeDefinition is called when entering the shapeDefinition production.
	EnterShapeDefinition(c *ShapeDefinitionContext)

	// EnterInlineShapeDefinition is called when entering the inlineShapeDefinition production.
	EnterInlineShapeDefinition(c *InlineShapeDefinitionContext)

	// EnterQualifier is called when entering the qualifier production.
	EnterQualifier(c *QualifierContext)

	// EnterExtraPropertySet is called when entering the extraPropertySet production.
	EnterExtraPropertySet(c *ExtraPropertySetContext)

	// EnterTripleExpression is called when entering the tripleExpression production.
	EnterTripleExpression(c *TripleExpressionContext)

	// EnterOneOfTripleExpr is called when entering the oneOfTripleExpr production.
	EnterOneOfTripleExpr(c *OneOfTripleExprContext)

	// EnterMultiElementOneOf is called when entering the multiElementOneOf production.
	EnterMultiElementOneOf(c *MultiElementOneOfContext)

	// EnterGroupTripleExpr is called when entering the groupTripleExpr production.
	EnterGroupTripleExpr(c *GroupTripleExprContext)

	// EnterSingleElementGroup is called when entering the singleElementGroup production.
	EnterSingleElementGroup(c *SingleElementGroupContext)

	// EnterMultiElementGroup is called when entering the multiElementGroup production.
	EnterMultiElementGroup(c *MultiElementGroupContext)

	// EnterUnaryTripleExpr is called when entering the unaryTripleExpr production.
	EnterUnaryTripleExpr(c *UnaryTripleExprContext)

	// EnterBracketedTripleExpr is called when entering the bracketedTripleExpr production.
	EnterBracketedTripleExpr(c *BracketedTripleExprContext)

	// EnterTripleConstraint is called when entering the tripleConstraint production.
	EnterTripleConstraint(c *TripleConstraintContext)

	// EnterStarCardinality is called when entering the starCardinality production.
	EnterStarCardinality(c *StarCardinalityContext)

	// EnterPlusCardinality is called when entering the plusCardinality production.
	EnterPlusCardinality(c *PlusCardinalityContext)

	// EnterOptionalCardinality is called when entering the optionalCardinality production.
	EnterOptionalCardinality(c *OptionalCardinalityContext)

	// EnterRepeatCardinality is called when entering the repeatCardinality production.
	EnterRepeatCardinality(c *RepeatCardinalityContext)

	// EnterExactRange is called when entering the exactRange production.
	EnterExactRange(c *ExactRangeContext)

	// EnterMinMaxRange is called when entering the minMaxRange production.
	EnterMinMaxRange(c *MinMaxRangeContext)

	// EnterSenseFlags is called when entering the senseFlags production.
	EnterSenseFlags(c *SenseFlagsContext)

	// EnterValueSet is called when entering the valueSet production.
	EnterValueSet(c *ValueSetContext)

	// EnterValueSetValue is called when entering the valueSetValue production.
	EnterValueSetValue(c *ValueSetValueContext)

	// EnterIriRange is called when entering the iriRange production.
	EnterIriRange(c *IriRangeContext)

	// EnterIriExclusion is called when entering the iriExclusion production.
	EnterIriExclusion(c *IriExclusionContext)

	// EnterLiteralRange is called when entering the literalRange production.
	EnterLiteralRange(c *LiteralRangeContext)

	// EnterLiteralExclusion is called when entering the literalExclusion production.
	EnterLiteralExclusion(c *LiteralExclusionContext)

	// EnterLanguageRangeFull is called when entering the languageRangeFull production.
	EnterLanguageRangeFull(c *LanguageRangeFullContext)

	// EnterLanguageRangeAt is called when entering the languageRangeAt production.
	EnterLanguageRangeAt(c *LanguageRangeAtContext)

	// EnterLanguageExclusion is called when entering the languageExclusion production.
	EnterLanguageExclusion(c *LanguageExclusionContext)

	// EnterInclude is called when entering the include production.
	EnterInclude(c *IncludeContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterSemanticAction is called when entering the semanticAction production.
	EnterSemanticAction(c *SemanticActionContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterRdfType is called when entering the rdfType production.
	EnterRdfType(c *RdfTypeContext)

	// EnterDatatype is called when entering the datatype production.
	EnterDatatype(c *DatatypeContext)

	// EnterShapeExprLabel is called when entering the shapeExprLabel production.
	EnterShapeExprLabel(c *ShapeExprLabelContext)

	// EnterTripleExprLabel is called when entering the tripleExprLabel production.
	EnterTripleExprLabel(c *TripleExprLabelContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterRdfLiteral is called when entering the rdfLiteral production.
	EnterRdfLiteral(c *RdfLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterRdfString is called when entering the rdfString production.
	EnterRdfString(c *RdfStringContext)

	// EnterIri is called when entering the iri production.
	EnterIri(c *IriContext)

	// EnterPrefixedName is called when entering the prefixedName production.
	EnterPrefixedName(c *PrefixedNameContext)

	// EnterBlankNode is called when entering the blankNode production.
	EnterBlankNode(c *BlankNodeContext)

	// EnterExtension is called when entering the extension production.
	EnterExtension(c *ExtensionContext)

	// EnterRestrictions is called when entering the restrictions production.
	EnterRestrictions(c *RestrictionsContext)

	// ExitShExDoc is called when exiting the shExDoc production.
	ExitShExDoc(c *ShExDocContext)

	// ExitDirective is called when exiting the directive production.
	ExitDirective(c *DirectiveContext)

	// ExitBaseDecl is called when exiting the baseDecl production.
	ExitBaseDecl(c *BaseDeclContext)

	// ExitPrefixDecl is called when exiting the prefixDecl production.
	ExitPrefixDecl(c *PrefixDeclContext)

	// ExitImportDecl is called when exiting the importDecl production.
	ExitImportDecl(c *ImportDeclContext)

	// ExitNotStartAction is called when exiting the notStartAction production.
	ExitNotStartAction(c *NotStartActionContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitStartActions is called when exiting the startActions production.
	ExitStartActions(c *StartActionsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitShapeExprDecl is called when exiting the shapeExprDecl production.
	ExitShapeExprDecl(c *ShapeExprDeclContext)

	// ExitShapeExpression is called when exiting the shapeExpression production.
	ExitShapeExpression(c *ShapeExpressionContext)

	// ExitInlineShapeExpression is called when exiting the inlineShapeExpression production.
	ExitInlineShapeExpression(c *InlineShapeExpressionContext)

	// ExitShapeOr is called when exiting the shapeOr production.
	ExitShapeOr(c *ShapeOrContext)

	// ExitInlineShapeOr is called when exiting the inlineShapeOr production.
	ExitInlineShapeOr(c *InlineShapeOrContext)

	// ExitShapeAnd is called when exiting the shapeAnd production.
	ExitShapeAnd(c *ShapeAndContext)

	// ExitInlineShapeAnd is called when exiting the inlineShapeAnd production.
	ExitInlineShapeAnd(c *InlineShapeAndContext)

	// ExitShapeNot is called when exiting the shapeNot production.
	ExitShapeNot(c *ShapeNotContext)

	// ExitInlineShapeNot is called when exiting the inlineShapeNot production.
	ExitInlineShapeNot(c *InlineShapeNotContext)

	// ExitShapeAtomNonLitNodeConstraint is called when exiting the shapeAtomNonLitNodeConstraint production.
	ExitShapeAtomNonLitNodeConstraint(c *ShapeAtomNonLitNodeConstraintContext)

	// ExitShapeAtomLitNodeConstraint is called when exiting the shapeAtomLitNodeConstraint production.
	ExitShapeAtomLitNodeConstraint(c *ShapeAtomLitNodeConstraintContext)

	// ExitShapeAtomShapeOrRef is called when exiting the shapeAtomShapeOrRef production.
	ExitShapeAtomShapeOrRef(c *ShapeAtomShapeOrRefContext)

	// ExitShapeAtomShapeExpression is called when exiting the shapeAtomShapeExpression production.
	ExitShapeAtomShapeExpression(c *ShapeAtomShapeExpressionContext)

	// ExitShapeAtomAny is called when exiting the shapeAtomAny production.
	ExitShapeAtomAny(c *ShapeAtomAnyContext)

	// ExitInlineShapeAtomNonLitNodeConstraint is called when exiting the inlineShapeAtomNonLitNodeConstraint production.
	ExitInlineShapeAtomNonLitNodeConstraint(c *InlineShapeAtomNonLitNodeConstraintContext)

	// ExitInlineShapeAtomLitNodeConstraint is called when exiting the inlineShapeAtomLitNodeConstraint production.
	ExitInlineShapeAtomLitNodeConstraint(c *InlineShapeAtomLitNodeConstraintContext)

	// ExitInlineShapeAtomShapeOrRef is called when exiting the inlineShapeAtomShapeOrRef production.
	ExitInlineShapeAtomShapeOrRef(c *InlineShapeAtomShapeOrRefContext)

	// ExitInlineShapeAtomShapeExpression is called when exiting the inlineShapeAtomShapeExpression production.
	ExitInlineShapeAtomShapeExpression(c *InlineShapeAtomShapeExpressionContext)

	// ExitInlineShapeAtomAny is called when exiting the inlineShapeAtomAny production.
	ExitInlineShapeAtomAny(c *InlineShapeAtomAnyContext)

	// ExitShapeOrRef is called when exiting the shapeOrRef production.
	ExitShapeOrRef(c *ShapeOrRefContext)

	// ExitInlineShapeOrRef is called when exiting the inlineShapeOrRef production.
	ExitInlineShapeOrRef(c *InlineShapeOrRefContext)

	// ExitShapeRef is called when exiting the shapeRef production.
	ExitShapeRef(c *ShapeRefContext)

	// ExitNodeConstraintLiteral is called when exiting the nodeConstraintLiteral production.
	ExitNodeConstraintLiteral(c *NodeConstraintLiteralContext)

	// ExitNodeConstraintNonLiteral is called when exiting the nodeConstraintNonLiteral production.
	ExitNodeConstraintNonLiteral(c *NodeConstraintNonLiteralContext)

	// ExitNodeConstraintDatatype is called when exiting the nodeConstraintDatatype production.
	ExitNodeConstraintDatatype(c *NodeConstraintDatatypeContext)

	// ExitNodeConstraintValueSet is called when exiting the nodeConstraintValueSet production.
	ExitNodeConstraintValueSet(c *NodeConstraintValueSetContext)

	// ExitNodeConstraintNumericFacet is called when exiting the nodeConstraintNumericFacet production.
	ExitNodeConstraintNumericFacet(c *NodeConstraintNumericFacetContext)

	// ExitLitNodeConstraint is called when exiting the litNodeConstraint production.
	ExitLitNodeConstraint(c *LitNodeConstraintContext)

	// ExitLitNodeConstraintLiteral is called when exiting the litNodeConstraintLiteral production.
	ExitLitNodeConstraintLiteral(c *LitNodeConstraintLiteralContext)

	// ExitLitNodeConstraintStringFacet is called when exiting the litNodeConstraintStringFacet production.
	ExitLitNodeConstraintStringFacet(c *LitNodeConstraintStringFacetContext)

	// ExitNonLitNodeConstraint is called when exiting the nonLitNodeConstraint production.
	ExitNonLitNodeConstraint(c *NonLitNodeConstraintContext)

	// ExitNonLiteralKind is called when exiting the nonLiteralKind production.
	ExitNonLiteralKind(c *NonLiteralKindContext)

	// ExitXsFacet is called when exiting the xsFacet production.
	ExitXsFacet(c *XsFacetContext)

	// ExitStringFacet is called when exiting the stringFacet production.
	ExitStringFacet(c *StringFacetContext)

	// ExitStringLength is called when exiting the stringLength production.
	ExitStringLength(c *StringLengthContext)

	// ExitNumericFacet is called when exiting the numericFacet production.
	ExitNumericFacet(c *NumericFacetContext)

	// ExitNumericRange is called when exiting the numericRange production.
	ExitNumericRange(c *NumericRangeContext)

	// ExitNumericLength is called when exiting the numericLength production.
	ExitNumericLength(c *NumericLengthContext)

	// ExitRawNumeric is called when exiting the rawNumeric production.
	ExitRawNumeric(c *RawNumericContext)

	// ExitShapeDefinition is called when exiting the shapeDefinition production.
	ExitShapeDefinition(c *ShapeDefinitionContext)

	// ExitInlineShapeDefinition is called when exiting the inlineShapeDefinition production.
	ExitInlineShapeDefinition(c *InlineShapeDefinitionContext)

	// ExitQualifier is called when exiting the qualifier production.
	ExitQualifier(c *QualifierContext)

	// ExitExtraPropertySet is called when exiting the extraPropertySet production.
	ExitExtraPropertySet(c *ExtraPropertySetContext)

	// ExitTripleExpression is called when exiting the tripleExpression production.
	ExitTripleExpression(c *TripleExpressionContext)

	// ExitOneOfTripleExpr is called when exiting the oneOfTripleExpr production.
	ExitOneOfTripleExpr(c *OneOfTripleExprContext)

	// ExitMultiElementOneOf is called when exiting the multiElementOneOf production.
	ExitMultiElementOneOf(c *MultiElementOneOfContext)

	// ExitGroupTripleExpr is called when exiting the groupTripleExpr production.
	ExitGroupTripleExpr(c *GroupTripleExprContext)

	// ExitSingleElementGroup is called when exiting the singleElementGroup production.
	ExitSingleElementGroup(c *SingleElementGroupContext)

	// ExitMultiElementGroup is called when exiting the multiElementGroup production.
	ExitMultiElementGroup(c *MultiElementGroupContext)

	// ExitUnaryTripleExpr is called when exiting the unaryTripleExpr production.
	ExitUnaryTripleExpr(c *UnaryTripleExprContext)

	// ExitBracketedTripleExpr is called when exiting the bracketedTripleExpr production.
	ExitBracketedTripleExpr(c *BracketedTripleExprContext)

	// ExitTripleConstraint is called when exiting the tripleConstraint production.
	ExitTripleConstraint(c *TripleConstraintContext)

	// ExitStarCardinality is called when exiting the starCardinality production.
	ExitStarCardinality(c *StarCardinalityContext)

	// ExitPlusCardinality is called when exiting the plusCardinality production.
	ExitPlusCardinality(c *PlusCardinalityContext)

	// ExitOptionalCardinality is called when exiting the optionalCardinality production.
	ExitOptionalCardinality(c *OptionalCardinalityContext)

	// ExitRepeatCardinality is called when exiting the repeatCardinality production.
	ExitRepeatCardinality(c *RepeatCardinalityContext)

	// ExitExactRange is called when exiting the exactRange production.
	ExitExactRange(c *ExactRangeContext)

	// ExitMinMaxRange is called when exiting the minMaxRange production.
	ExitMinMaxRange(c *MinMaxRangeContext)

	// ExitSenseFlags is called when exiting the senseFlags production.
	ExitSenseFlags(c *SenseFlagsContext)

	// ExitValueSet is called when exiting the valueSet production.
	ExitValueSet(c *ValueSetContext)

	// ExitValueSetValue is called when exiting the valueSetValue production.
	ExitValueSetValue(c *ValueSetValueContext)

	// ExitIriRange is called when exiting the iriRange production.
	ExitIriRange(c *IriRangeContext)

	// ExitIriExclusion is called when exiting the iriExclusion production.
	ExitIriExclusion(c *IriExclusionContext)

	// ExitLiteralRange is called when exiting the literalRange production.
	ExitLiteralRange(c *LiteralRangeContext)

	// ExitLiteralExclusion is called when exiting the literalExclusion production.
	ExitLiteralExclusion(c *LiteralExclusionContext)

	// ExitLanguageRangeFull is called when exiting the languageRangeFull production.
	ExitLanguageRangeFull(c *LanguageRangeFullContext)

	// ExitLanguageRangeAt is called when exiting the languageRangeAt production.
	ExitLanguageRangeAt(c *LanguageRangeAtContext)

	// ExitLanguageExclusion is called when exiting the languageExclusion production.
	ExitLanguageExclusion(c *LanguageExclusionContext)

	// ExitInclude is called when exiting the include production.
	ExitInclude(c *IncludeContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitSemanticAction is called when exiting the semanticAction production.
	ExitSemanticAction(c *SemanticActionContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitRdfType is called when exiting the rdfType production.
	ExitRdfType(c *RdfTypeContext)

	// ExitDatatype is called when exiting the datatype production.
	ExitDatatype(c *DatatypeContext)

	// ExitShapeExprLabel is called when exiting the shapeExprLabel production.
	ExitShapeExprLabel(c *ShapeExprLabelContext)

	// ExitTripleExprLabel is called when exiting the tripleExprLabel production.
	ExitTripleExprLabel(c *TripleExprLabelContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitRdfLiteral is called when exiting the rdfLiteral production.
	ExitRdfLiteral(c *RdfLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitRdfString is called when exiting the rdfString production.
	ExitRdfString(c *RdfStringContext)

	// ExitIri is called when exiting the iri production.
	ExitIri(c *IriContext)

	// ExitPrefixedName is called when exiting the prefixedName production.
	ExitPrefixedName(c *PrefixedNameContext)

	// ExitBlankNode is called when exiting the blankNode production.
	ExitBlankNode(c *BlankNodeContext)

	// ExitExtension is called when exiting the extension production.
	ExitExtension(c *ExtensionContext)

	// ExitRestrictions is called when exiting the restrictions production.
	ExitRestrictions(c *RestrictionsContext)
}
