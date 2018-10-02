from typing import Optional, Union, List

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

    # ------------------
    # shapeOr
    # ------------------
    def _shapeOr(self, ands: Union[List[ShExDocParser.ShapeAndContext], List[ShExDocParser.InlineShapeAndContext]]):
        if len(ands) > 1:
            exprs = []
            for sa in ands:
                sep = ShexShapeExpressionParser(self.context)
                sep.visit(sa)
                exprs.append(sep.expr)
            self.expr = ShapeOr(id=self.label, shapeExprs=exprs)
        else:
            self.visit(ands[0])

    def visitShapeOr(self, ctx: ShExDocParser.ShapeOrContext):
        """ shapeOr: shapeAnd (KW_OR shapeAnd)* """
        self._shapeOr(ctx.shapeAnd())

    def visitInlineShapeOr(self, ctx: ShExDocParser.InlineShapeOrContext):
        """ inlineShapeOr: inlineShapeAnd (KW_OR inlineShapeAnd)* """
        self._shapeOr(ctx.inlineShapeAnd())

    # ------------------
    # shapeAnd
    # ------------------
    def _shapeAnd(self,
                  nots: Union[List[ShExDocParser.ShapeNotContext], List[ShExDocParser.InlineShapeNotContext]],
                  isInline: bool):
        if len(nots) > 1:
            exprs = []
            for sa in nots:
                sep = ShexShapeExpressionParser(self.context)
                sep.visit(sa)
                if isInline:
                    self._collapse_ands(exprs, sep.expr)
                else:
                    exprs.append(sep.expr)
            self.expr = ShapeAnd(id=self.label, shapeExprs=exprs)
        else:
            self.visit(nots[0])

    @staticmethod
    def _collapse_ands(exprs: List, new_expr: shapeExpr) -> None:
        """ For various nefarious reaons, the reference parser has decided to implement
            And(And(a, b), c) --> And(a, b, c).  We've got to do the same
        """
        if isinstance(new_expr, ShapeAnd):
            for expr in new_expr.shapeExprs:
                exprs.append(expr)
        else:
            exprs.append(new_expr)

    def visitShapeAnd(self, ctx: ShExDocParser.ShapeAndContext):
        """ shapeAnd: shapeNot (KW_AND shapeNot)* """
        self._shapeAnd(ctx.shapeNot(), False)

    def visitInlineShapeAnd(self, ctx: ShExDocParser.InlineShapeAndContext):
        """ inlineShapeAnd: inlineShapeNot (KW_AND inlineShapeNot)* """
        self._shapeAnd(ctx.inlineShapeNot(), True)

    # ------------------
    # shapeNot
    # ------------------
    def _shapeNot(self,
                  ctx: Union[ShExDocParser.ShapeNotContext, ShExDocParser.InlineShapeNotContext],
                  isInline: bool) -> None:
        if ctx.KW_NOT():
            self.expr = ShapeNot(id=self.label)
            sn = ShexShapeExpressionParser(self.context)
            sn.visit(ctx.shapeAtom() if not isInline else ctx.inlineShapeAtom())
            self.expr.shapeExpr = sn.expr if sn.expr is not None else Shape()
        else:
            self.visitChildren(ctx)

    def visitShapeNot(self, ctx: ShExDocParser.ShapeNotContext):
        """ shapeNot: KW_NOT? shapeAtom """
        self._shapeNot(ctx, False)

    def visitInlineShapeNot(self, ctx: ShExDocParser.InlineShapeNotContext):
        """ inlineShapeNot: KW_NOT? inlineShapeAtom """
        self._shapeNot(ctx, True)

    # ------------------
    # shapeAtom
    # ------------------
    def _shapeAtom(self, ctx: Union[ShExDocParser.ShapeAtomNonLitNodeConstraintContext,
                                    ShExDocParser.ShapeAtomLitNodeConstraintContext,
                                    ShExDocParser.InlineShapeAtomNonLitNodeConstraintContext,
                                    ShExDocParser.InlineShapeAtomLitNodeConstraintContext],
                   isInline: bool, isLit: bool):
        """ One of nonLitNodeConstraint shapeOrRef?  or shapeOrRef nonLitNodeConstraint? """

        # Process the node constraint if it exists
        nodeConstraint = ctx.litNodeConstraint() if isLit else ctx.nonLitNodeConstraint()
        if nodeConstraint:
            nc = ShexNodeExpressionParser(self.context, self.label if isInline else None)
            nc.visit(nodeConstraint)
            exprs = [nc.nodeconstraint]
        else:
            exprs = []

        # Process the shape or ref if it exists
        if not isLit:
            shapeOrRef = ctx.inlineShapeOrRef() if isInline else ctx.shapeOrRef()
        else:
            shapeOrRef = None
        if shapeOrRef:
            sorref_parser = ShexShapeExpressionParser(self.context)
            sorref_parser.visit(shapeOrRef)
            exprs.append(sorref_parser.expr)

        if len(exprs) > 1:
            self.expr = ShapeAnd(id=self.label if not isInline else None, shapeExprs=exprs)
        else:
            self.expr = exprs[0]
        if not isInline:
            self.expr.id = self.label

    def visitShapeAtomNonLitNodeConstraint(self, ctx: ShExDocParser.ShapeAtomNonLitNodeConstraintContext):
        """ shapeAtom : nonLitNodeConstraint shapeOrRef?             # shapeAtomNonLitNodeConstraint """
        self._shapeAtom(ctx, False, False)

    def visitShapeAtomLitNodeConstraint(self, ctx: ShExDocParser.ShapeAtomLitNodeConstraintContext):
        """ shapeAtom : litNodeConstraint             # shapeAtomLitNodeConstraint"""
        self._shapeAtom(ctx, False, True)

    def visitInlineShapeAtomNonLitNodeConstraint(self, ctx: ShExDocParser.InlineShapeAtomNonLitNodeConstraintContext):
        """ inlineShapeAtom : nonLitNodeConstraint inlineShapeOrRef? # inlineShapeAtomNonLitNodeConstraint """
        self._shapeAtom(ctx, True, False)

    def visitInlineShapeAtomLitNodeConstraint(self, ctx: ShExDocParser.InlineShapeAtomLitNodeConstraintContext):
        """ inlineShapeAtom : | litNodeConstraint             # inlineShapeAtomLitNodeConstraint """
        self._shapeAtom(ctx, True, True)

    def visitShapeAtomShapeOrRef(self, ctx: ShExDocParser.ShapeAtomShapeOrRefContext):
        """ shapeAtom :  shapeOrRef nonLitNodeConstraint?            # shapeAtomShapeOrRef """
        self._shapeAtom(ctx, False, False)

    def visitInlineShapeAtomShapeOrRef(self, ctx: ShExDocParser.InlineShapeAtomShapeOrRefContext):
        """ inlineShapeAtom : inlineShapeOrRef nonLitNodeConstraint? # inlineShapeAtomShapeOrRef """
        self._shapeAtom(ctx, True, False)

    # ------------------
    # shapeOrRef
    # ------------------
    def _shapeOrRef(self, ctx: Union[ShExDocParser.ShapeOrRefContext, ShExDocParser.InlineShapeOrRefContext],
                    isInline: bool):
        defn = ctx.inlineShapeDefinition() if isInline else ctx.shapeDefinition()
        if defn:
            from pyshexc.parser_impl.shex_shape_definition_parser import ShexShapeDefinitionParser
            shdef_parser = ShexShapeDefinitionParser(self.context, self.label)
            shdef_parser.visitChildren(ctx)
            self.expr = shdef_parser.shape
        else:
            self.expr = self.context.shapeRef_to_iriref(ctx.shapeRef())

    def visitShapeOrRef(self, ctx: ShExDocParser.ShapeOrRefContext):
        """ shapeOrRef: shapeDefinition | shapeRef """
        self._shapeOrRef(ctx, False)

    def visitInlineShapeOrRef(self, ctx: ShExDocParser.InlineShapeOrRefContext):
        """ inlineShapeOrRef: inlineShapeDefinition | shapeRef """
        self._shapeOrRef(ctx, True)
