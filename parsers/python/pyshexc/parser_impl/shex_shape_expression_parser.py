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
from typing import Optional, Union

from pyshexc.parser.ShExDocParser import ShExDocParser
from pyshexc.parser.ShExDocVisitor import ShExDocVisitor

from pyshexc.parser_impl.parser_context import ParserContext
from pyshexc.parser_impl.shex_node_expression_parser import ShexNodeExpressionParser
from ShExJSG.ShExJ import BNODE, IRIREF, ShapeOr, ShapeAnd, ShapeNot, Shape, shapeExpr


class ShexShapeExpressionParser(ShExDocVisitor):
    """ Parser for Shape Expressions branch """
    def __init__(self, context: ParserContext, label: Optional[Union[IRIREF, BNODE]]=None):
        ShExDocVisitor.__init__(self)
        self.context = context
        self.label = label
        self.expr = None

    def visitShapeOr(self, ctx: ShExDocParser.ShapeOrContext):
        """ shapeOr: shapeAnd (KW_OR shapeAnd)* """
        if len(ctx.shapeAnd()) > 1:
            self.expr = ShapeOr(id=self.label, shapeExprs=[])
            for sa in ctx.shapeAnd():
                sep = ShexShapeExpressionParser(self.context)
                sep.visit(sa)
                self.expr.shapeExprs.append(sep.expr)
        else:
            self.visit(ctx.shapeAnd(0))

    def visitInlineShapeOr(self, ctx: ShExDocParser.InlineShapeOrContext):
        """ inlineShapeOr: inlineShapeAnd (KW_OR inlineShapeAnd)* """
        if len(ctx.inlineShapeAnd()) > 1:
            self.expr = ShapeOr(id=self.label, shapeExprs=[])
            for sa in ctx.inlineShapeAnd():
                sep = ShexShapeExpressionParser(self.context)
                sep.visit(sa)
                self.expr.shapeExprs.append(sep.expr)
        else:
            self.visit(ctx.inlineShapeAnd(0))

    def visitShapeAnd(self, ctx: ShExDocParser.ShapeAndContext):
        """ shapeAnd: shapeNot (KW_AND shapeNot)* """
        if len(ctx.shapeNot()) > 1:
            self.expr = ShapeAnd(id=self.label, shapeExprs=[])
            for sa in ctx.shapeNot():
                sep = ShexShapeExpressionParser(self.context)
                sep.visit(sa)
                self.expr.shapeExprs.append(sep.expr)
        else:
            self.visit(ctx.shapeNot(0))

    def visitInlineShapeAnd(self, ctx: ShExDocParser.InlineShapeAndContext):
        """ inlineShapeAnd: inlineShapeNot (KW_AND inlineShapeNot)* """
        if len(ctx.inlineShapeNot()) > 1:
            self.expr = ShapeAnd(id=self.label, shapeExprs=[])
            for sa in ctx.inlineShapeNot():
                sep = ShexShapeExpressionParser(self.context)
                sep.visit(sa)
                self._and_collapser(self.expr, sep.expr)
        else:
            self.visit(ctx.inlineShapeNot(0))

    @staticmethod
    def _and_collapser(container: ShapeAnd, containee: shapeExpr) -> None:
        """ For various nefarious reaons, the reference parser has decided to implement
            And(And(a, b), c) --> And(a, b, c).  We've got to do the same
        """
        if isinstance(containee, ShapeAnd):
            for expr in containee.shapeExprs:
                container.shapeExprs.append(expr)
        else:
            container.shapeExprs.append(containee)

    def visitShapeNot(self, ctx: ShExDocParser.ShapeNotContext):
        """ shapeNot: negation? shapeAtom """
        if ctx.negation():
            self.expr = ShapeNot(id=self.label)
            sn = ShexShapeExpressionParser(self.context)
            sn.visit(ctx.shapeAtom())
            self.expr.shapeExpr = sn.expr if sn.expr is not None else Shape()
        else:
            self.visitChildren(ctx)

    def visitInlineShapeNot(self, ctx: ShExDocParser.InlineShapeNotContext):
        """ inlineShapeNot: negation? inlineShapeAtom """
        if ctx.negation():
            self.expr = ShapeNot(id=self.label)
            sn = ShexShapeExpressionParser(self.context)
            ctx.inlineShapeAtom()
            sn.visit(ctx.inlineShapeAtom())
            self.expr.shapeExpr = sn.expr if sn.expr is not None else Shape()
        else:
            self.visitChildren(ctx)

    def visitShapeAtomNodeConstraint(self, ctx: ShExDocParser.ShapeAtomNodeConstraintContext):
        """  shapeAtomNodeConstraint: nodeConstraint shapeOrRef?  # shapeAtomNodeConstraint """
        nc = ShexNodeExpressionParser(self.context)
        nc.visit(ctx.nodeConstraint())
        if ctx.shapeOrRef():
            self.expr = ShapeAnd(id=self.label, shapeExprs=[nc.nodeconstraint])
            sorref_parser = ShexShapeExpressionParser(self.context)
            sorref_parser.visit(ctx.shapeOrRef())
            # if isinstance(sorref_parser.expr, Shape) and self.context.is_empty_shape(sorref_parser.expr):
            #     self.expr = nc.nodeconstraint
            #     self.expr.id = self.label
            # else:
            self.expr.shapeExprs.append(sorref_parser.expr)
        else:
            self.expr = nc.nodeconstraint
            self.expr.id = self.label

    def visitInlineShapeAtomNodeConstraint(self, ctx: ShExDocParser.InlineShapeAtomNodeConstraintContext):
        """  inlineShapeAtomNodeConstraint: nodeConstraint inlineShapeOrRef?  # inlineShapeAtomShapeOrRef """
        nc = ShexNodeExpressionParser(self.context, self.label)
        nc.visit(ctx.nodeConstraint())
        if ctx.inlineShapeOrRef():
            self.expr = ShapeAnd(shapeExprs=[nc.nodeconstraint])
            sorref_parser = ShexShapeExpressionParser(self.context)
            sorref_parser.visit(ctx.inlineShapeOrRef())
            # if isinstance(sorref_parser.expr, Shape) and self.context.is_empty_shape(sorref_parser.expr):
            #     self.expr = nc.nodeconstraint
            # else:
            self.expr.shapeExprs.append(sorref_parser.expr)
        else:
            self.expr = nc.nodeconstraint

    def visitShapeOrRef(self, ctx: ShExDocParser.ShapeOrRefContext):
        """ shapeOrRef: shapeDefinition | shapeRef """
        if ctx.shapeDefinition():
            from pyshexc.parser_impl.shex_shape_definition_parser import ShexShapeDefinitionParser
            shdef_parser = ShexShapeDefinitionParser(self.context, self.label)
            shdef_parser.visitChildren(ctx)
            self.expr = shdef_parser.shape
        else:
            self.expr = self.context.shapeRef_to_iriref(ctx.shapeRef())

    def visitInlineShapeOrRef(self, ctx: ShExDocParser.InlineShapeOrRefContext):
        """ inlineShapeOrRef: inlineShapeDefinition | shapeRef """
        if ctx.inlineShapeDefinition():
            from pyshexc.parser_impl.shex_shape_definition_parser import ShexShapeDefinitionParser
            shdef_parser = ShexShapeDefinitionParser(self.context, self.label)
            shdef_parser.visitChildren(ctx)
            self.expr = shdef_parser.shape
        else:
            self.expr = self.context.shapeRef_to_iriref(ctx.shapeRef())
