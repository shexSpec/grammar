# Copyright (c) 2016, Mayo Clinic
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
#     Neither the name of the <ORGANIZATION> nor the names of its contributors
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
from typing import Optional

from pyshexc.parser.ShExDocParser import ShExDocParser
from pyshexc.parser.ShExDocVisitor import ShExDocVisitor

from pyshexc.parser_impl.parser_context import ParserContext
from pyshexc.parser_impl.shex_annotations_and_semacts_parser import ShexAnnotationAndSemactsParser
from pyshexc.parser_impl.shex_shape_expression_parser import ShexShapeExpressionParser
from ShExJSG.ShExJ import ShapeExternal, IRIREF


class ShexDocParser(ShExDocVisitor):
    """ parser for sheExDoc production """
    def __init__(self, default_base: Optional[str]=None):
        ShExDocVisitor.__init__(self)
        self.context = ParserContext()
        self.context.base = IRIREF(default_base) if default_base else None

    def visitShExDoc(self, ctx: ShExDocParser.ShExDocContext):
        """ shExDoc: directive* ((notStartAction | startActions) statement*)? EOF """
        self.visitChildren(ctx)

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

    def visitStart(self, ctx: ShExDocParser.StartContext):
        """ start: KW_START '=' shapeExpression """
        shexpr = ShexShapeExpressionParser(self.context, None)
        shexpr.visit(ctx.shapeExpression())
        self.context.schema.start = shexpr.expr

    def visitShapeExprDecl(self, ctx: ShExDocParser.ShapeExprDeclContext):
        """ shapeExprDecl: shapeExprLabel (shapeExpression | KW_EXTERNAL) """
        label = self.context.shapeexprlabel_to_IRI(ctx.shapeExprLabel())
        if self.context.schema.shapes is None:
            self.context.schema.shapes = []
        if ctx.KW_EXTERNAL():
            shape = ShapeExternal(id=label)
        else:
            shexpr = ShexShapeExpressionParser(self.context, label)
            shexpr.visit(ctx.shapeExpression())
            shape = shexpr.expr
        self.context.schema.shapes.append(shape)

    def visitStartActions(self, ctx: ShExDocParser.StartActionsContext):
        """ startActions: codeDecl+ """
        self.context.schema.startActs = []
        for cd in ctx.codeDecl():
            cdparser = ShexAnnotationAndSemactsParser(self.context)
            cdparser.visit(cd)
            self.context.schema.startActs += cdparser.semacts
