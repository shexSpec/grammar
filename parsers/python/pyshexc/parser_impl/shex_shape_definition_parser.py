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
from pyshexc.parser_impl.shex_annotations_and_semacts_parser import ShexAnnotationAndSemactsParser
from pyshexc.parser_impl.shex_oneofshape_parser import ShexOneOfShapeParser
from ShExJSG.ShExJ import IRIREF, BNODE, Shape


class ShexShapeDefinitionParser(ShExDocVisitor):
    def __init__(self, context: ParserContext, label: Optional[Union[IRIREF, BNODE]]=None):
        ShExDocVisitor.__init__(self)
        self.context = context
        self.shape = Shape(label)

    def visitShapeDefinition(self, ctx: ShExDocParser.ShapeDefinitionContext):
        """ shapeDefinition: qualifier* '{' oneOfShape? '}' annotation* semanticActions """
        if ctx.qualifier():
            for q in ctx.qualifier():
                self.visit(q)
        if ctx.oneOfShape():
            oneof_parser = ShexOneOfShapeParser(self.context)
            oneof_parser.visit(ctx.oneOfShape())
            self.shape.expression = oneof_parser.expression
        if ctx.annotation() or ctx.semanticActions():
            ansem_parser = ShexAnnotationAndSemactsParser(self.context)
            for annot in ctx.annotation():
                ansem_parser.visit(annot)
            ansem_parser.visit(ctx.semanticActions())
            if ansem_parser.semacts:
                self.shape.semActs = ansem_parser.semacts
            if ansem_parser.annotations:
                self.shape.annotations = ansem_parser.annotations

    def visitInlineShapeDefinition(self, ctx: ShExDocParser.InlineShapeDefinitionContext):
        """ shapeDefinition: qualifier* '{' oneOfShape? '}' """
        if ctx.qualifier():
            for q in ctx.qualifier():
                self.visit(q)
        if ctx.oneOfShape():
            oneof_parser = ShexOneOfShapeParser(self.context)
            oneof_parser.visit(ctx.oneOfShape())
            self.shape.expression = oneof_parser.expression

    def visitQualifier(self, ctx: ShExDocParser.QualifierContext):
        """ qualifier: includeSet | extraPropertySet | KW_CLOSED """
        if ctx.includeSet():
            if self.shape.inherit is None:
                self.shape.inherit = []
            self.shape.inherit += [self.context.tripleexprlabel_to_iriref(tel)
                                   for tel in ctx.includeSet().tripleExpressionLabel()]
        elif ctx.extraPropertySet():
            if self.shape.extra is None:
                self.shape.extra = []
            self.shape.extra += [self.context.predicate_to_IRI(p) for p in ctx.extraPropertySet().predicate()]

        elif ctx.KW_CLOSED():
            self.shape.closed = True
