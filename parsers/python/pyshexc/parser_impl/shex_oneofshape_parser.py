# Copyright (c) 2017, Mayo Clinic
# All rights reserved.
#
# Redistribution and use in source and binary forms, with or without modification,
# are permitted provided that the following conditions are met:
#
# Redistributions of source code must retain the above copyright notice, this
#     list of conditions and the following disclaimer.
#
#     Redistributions in binary form must reproduce the above copyright notice,
#     this list of conditions and the following disclaimer in the documentation
#     and/or other materials provided with the distribution.
#
#     Neither the name of the Mayo Clinic nor the names of its contributors
#     may be used to endorse or promote products derived from this software
#     without specific prior written permission.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
# ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
# WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
# IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
# INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
# BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, 
# DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
# LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
# OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED
# OF THE POSSIBILITY OF SUCH DAMAGE.
from pyshexc.parser.ShExDocParser import ShExDocParser
from pyshexc.parser.ShExDocVisitor import ShExDocVisitor

from pyshexc.parser_impl.parser_context import ParserContext
from pyshexc.parser_impl.shex_annotations_and_semacts_parser import ShexAnnotationAndSemactsParser
from pyshexc.parser_impl.shex_shape_expression_parser import ShexShapeExpressionParser
from ShExJSG.ShExJ import OneOf, EachOf, TripleConstraint


class ShexOneOfShapeParser(ShExDocVisitor):
    def __init__(self, context: ParserContext):
        ShExDocVisitor.__init__(self)
        self.context = context
        self.expression = None

    def visitMultiElementOneOf(self, ctx: ShExDocParser.MultiElementOneOfContext):
        """ multiElementOneOf: groupShape ('|' groupShape)+ """
        self.expression = OneOf(expressions=[])
        for gs in ctx.groupShape():
            parser = ShexOneOfShapeParser(self.context)
            parser.visit(gs)
            self.expression.expressions.append(parser.expression)

    def visitMultiElementGroup(self, ctx: ShExDocParser.MultiElementGroupContext):
        """ multiElementGroup: unaryShape (';' unaryShape)+ ';'? """
        self.expression = EachOf(expressions=[])
        for us in ctx.unaryShape():
            parser = ShexOneOfShapeParser(self.context)
            parser.visit(us)
            self.expression.expressions.append(parser.expression)

    def visitUnaryShape(self, ctx: ShExDocParser.UnaryShapeContext):
        """ unaryShape: ('$' tripleExprLabel)? (tripleConstraint | encapsulatedShape) | include """
        if ctx.include():
            self.expression = self.context.tripleexprlabel_to_iriref(ctx.include().tripleExprLabel())
        else:
            lbl = self.context.tripleexprlabel_to_iriref(ctx.tripleExprLabel()) if ctx.tripleExprLabel() else None
            if ctx.tripleConstraint():
                self.expression = TripleConstraint(lbl)
                self.visit(ctx.tripleConstraint())
            elif ctx.encapsulatedShape():
                self.visit(ctx.encapsulatedShape())
                self.expression.id = lbl

    def visitEncapsulatedShape(self, ctx: ShExDocParser.EncapsulatedShapeContext):
        """ encapsulatedShape: '(' innerShape ')' cardinality? annotation* semanticActions """
        enc_shape = ShexOneOfShapeParser(self.context)
        enc_shape.visit(ctx.innerShape())
        self.expression = enc_shape.expression
        self._card_annotations_and_semacts(ctx)

    def visitTripleConstraint(self, ctx: ShExDocParser.TripleConstraintContext):
        """ tripleConstraint: senseFlags? predicate inlineShapeExpression cardinality? annotation* semanticActions """
        # This exists because of the predicate within annotation - if we default to visitchildren, we intercept both
        # predicates
        if ctx.senseFlags():
            self.visit(ctx.senseFlags())
        self.visit(ctx.predicate())
        self.visit(ctx.inlineShapeExpression())
        self._card_annotations_and_semacts(ctx)

    def visitStarCardinality(self, ctx: ShExDocParser.StarCardinalityContext):
        """ '*' """
        self.expression.min = 0
        self.expression.max = -1

    def visitPlusCardinality(self, ctx: ShExDocParser.PlusCardinalityContext):
        """ '+' """
        self.expression.min = 1
        self.expression.max = -1

    def visitOptionalCardinality(self, ctx: ShExDocParser.OptionalCardinalityContext):
        """ '?' """
        self.expression.min = 0
        self.expression.max = 1

    def visitExactRange(self, ctx: ShExDocParser.ExactRangeContext):
        """ repeatRange: '{' INTEGER '}' #exactRange """
        self.expression.min = int(ctx.INTEGER().getText())
        self.expression.max = self.expression.min

    def visitMinMaxRange(self, ctx: ShExDocParser.MinMaxRangeContext):
        """ repeatRange: '{' INTEGER (',' (INTEGER | UNBOUNDED)? '}' """
        self.expression.min = int(ctx.INTEGER(0).getText())
        self.expression.max = int(ctx.INTEGER(1).getText()) if len(ctx.INTEGER()) > 1 else -1

    def visitSenseFlags(self, ctx: ShExDocParser.SenseFlagsContext):
        """ '!' '^'? | '^' '!'? """
        if '!' in ctx.getText():
            self.expression.negated = True
        if '^' in ctx.getText():
            self.expression.inverse = True

    def visitPredicate(self, ctx: ShExDocParser.PredicateContext):
        """ predicate: iri | rdfType """
        self.expression.predicate = self.context.predicate_to_IRI(ctx)

    def visitInlineShapeExpression(self, ctx: ShExDocParser.InlineShapeExpressionContext):
        """ inlineShapeExpression: inlineShapeOr """
        expr_parser = ShexShapeExpressionParser(self.context)
        expr_parser.visitChildren(ctx)
        self.expression.valueExpr = expr_parser.expr

    def _card_annotations_and_semacts(self, ctx):
        if ctx.cardinality():
            self.visit(ctx.cardinality())
        annot_parser = ShexAnnotationAndSemactsParser(self.context)
        if ctx.annotation():
            for annot in ctx.annotation():
                annot_parser.visit(annot)
        annot_parser.visit(ctx.semanticActions())
        if annot_parser.annotations:
            self.expression.annotations = annot_parser.annotations
        if annot_parser.semacts:
            self.expression.semActs = annot_parser.semacts
