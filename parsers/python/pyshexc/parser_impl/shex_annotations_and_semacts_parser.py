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
from ShExJSG.ShExJ import Annotation, SemAct


class ShexAnnotationAndSemactsParser(ShExDocVisitor):
    def __init__(self, context: ParserContext):
        ShExDocVisitor.__init__(self)
        self.context = context
        self.semacts = []                   # List[SemAct]
        self.annotations = []               # List[Annotation]

    def visitAnnotation(self, ctx: ShExDocParser.AnnotationContext):
        """ annotation: '//' predicate (iri | literal) """
        # Annotations apply to the expression, NOT the shape (!)
        annot = Annotation(self.context.predicate_to_IRI(ctx.predicate()))
        if ctx.iri():
            annot.object = self.context.iri_to_iriref(ctx.iri())
        else:
            annot.object = self.context.literal_to_ObjectLiteral(ctx.literal())
        self.annotations.append(annot)

    def visitCodeDecl(self, ctx: ShExDocParser.CodeDeclContext):
        """ codeDecl: '%' iri (CODE | '%') 
            CODE: : '{' (~[%\\] | '\\' [%\\] | UCHAR)* '%' '}' """
        semact = SemAct()
        semact.name = self.context.iri_to_iriref(ctx.iri())
        if ctx.CODE():
            semact.code = ctx.CODE().getText()[1:-2].replace('\\%', '%').encode('utf-8').decode('unicode-escape')
        self.semacts.append(semact)
