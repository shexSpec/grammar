from typing import Optional

from pyshexc.parser.ShExDocParser import ShExDocParser
from pyshexc.parser.ShExDocVisitor import ShExDocVisitor

from pyshexc.parser_impl.parser_context import ParserContext
from pyshexc.parser_impl.shex_annotations_and_semacts_parser import ShexAnnotationAndSemactsParser
from pyshexc.parser_impl.shex_shape_expression_parser import ShexShapeExpressionParser
from ShExJSG.ShExJ import ShapeExternal, IRIREF, ShapeDecl


class ShexDocParser(ShExDocVisitor):
    """ parser for sheExDoc production """
    def __init__(self, default_base: Optional[str]=None):
        ShExDocVisitor.__init__(self)
        self.context = ParserContext()
        self.context.base = IRIREF(default_base) if default_base else None

    def visitShExDoc(self, ctx: ShExDocParser.ShExDocContext):
        """ shExDoc: directive* ((notStartAction | startActions) statement*)? EOF """
        super().visitShExDoc(ctx)

    def visitBaseDecl(self, ctx: ShExDocParser.BaseDeclContext):
        """ baseDecl: KW_BASE IRIREF """
        self.context.base = None
        self.context.base = self.context.iriref_to_shexj_iriref(ctx.IRIREF())

    def visitPrefixDecl(self, ctx: ShExDocParser.PrefixDeclContext):
        """ prefixDecl: KW_PREFIX PNAME_NS IRIREF """
        iri = self.context.iriref_to_shexj_iriref(ctx.IRIREF())
        prefix = ctx.PNAME_NS().getText()
        if iri not in self.context.ld_prefixes:
            self.context.prefixes.setdefault(prefix, iri.val)

    def visitImportDecl(self, ctx: ShExDocParser.ImportDeclContext):
        """ importDecl : KW_IMPORT IRIREF """
        if self.context.schema.imports is None:
            self.context.schema.imports = [self.context.iriref_to_shexj_iriref(ctx.IRIREF())]
        else:
            self.context.schema.imports.append(self.context.iriref_to_shexj_iriref(ctx.IRIREF()))

    def visitStart(self, ctx: ShExDocParser.StartContext):
        """ start: KW_START '=' shapeExpression """
        shexpr = ShexShapeExpressionParser(self.context, None)
        shexpr.visit(ctx.shapeExpression())
        self.context.schema.start = shexpr.expr

    def visitShapeExprDecl(self, ctx: ShExDocParser.ShapeExprDeclContext):
        """ shapeExprDecl: KW_ABSTRACT? shapeExprLabel restrictions* (shapeExpression | KW_EXTERNAL) """
        label = self.context.shapeexprlabel_to_IRI(ctx.shapeExprLabel())
        # CHANGED
        # if self.context.schema.shapes is None:
        #     self.context.schema.shapes = []
        if ctx.KW_ABSTRACT() or ctx.restrictions():
            decl = ShapeDecl(id=label)
            label = None
            if ctx.KW_ABSTRACT():
                decl.abstract = True
            if ctx.restrictions():
                decl.restricts = [self.context.shapeexprlabel_to_IRI(lbl.shapeExprLabel()) for
                                  lbl in ctx.restrictions()]
            if self.context.schema.shapes is None:
                self.context.schema.shapes = [decl]
            else:
                self.context.schema.shapes.append(decl)
        else:
            decl = None
        if ctx.KW_EXTERNAL():
            shape = ShapeExternal(id=label)
        else:
            shexpr = ShexShapeExpressionParser(self.context, label)
            shexpr.visit(ctx.shapeExpression())
            shape = shexpr.expr
        if not decl:
            if self.context.schema.shapes is None:
                self.context.schema.shapes = [shape]
            else:
                self.context.schema.shapes.append(shape)
        else:
            decl.shapeExpr = shape

    def visitStartActions(self, ctx: ShExDocParser.StartActionsContext):
        """ startActions: codeDecl+ """
        startacts = []
        for cd in ctx.codeDecl():
            cdparser = ShexAnnotationAndSemactsParser(self.context)
            cdparser.visit(cd)
            startacts += cdparser.semacts
        self.context.schema.startActs = startacts
