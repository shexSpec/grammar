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
from typing import Optional, List

from pyshexc.parser.ShExDocParser import ShExDocParser
from pyshexc.parser.ShExDocVisitor import ShExDocVisitor

from pyshexc.parser_impl.parser_context import ParserContext
from pyshexc.shexj.ShExJ import NodeConstraint, Wildcard, IriStemRange, LiteralStemRange, LanguageStemRange, IriStem, \
    LiteralStem, ObjectLiteral, LanguageStem, INTEGER, STRING


class ShexNodeExpressionParser(ShExDocVisitor):
    """ Parser for nodeConstraint productions """
    def __init__(self, context: ParserContext, label: Optional[str]=None):
        ShExDocVisitor.__init__(self)
        self.context = context
        self.nodeconstraint = NodeConstraint(id=label)

    def visitNodeConstraintLiteral(self, ctx: ShExDocParser.NodeConstraintLiteralContext):
        """ nodeConstraint: KW_LITERAL xsFacet* # nodeConstraintLiteral """
        self.nodeconstraint.nodeKind = 'literal'
        self.visitChildren(ctx)

    def visitNodeConstraintNonLiteral(self, ctx: ShExDocParser.NodeConstraintNonLiteralContext):
        """ nodeConstraint: nonLiteralKind stringFacet*  # nodeConstraintNonLiteral """
        self.nodeconstraint.nodeKind = 'iri' if ctx.nonLiteralKind().KW_IRI() \
            else 'bnode' if ctx.nonLiteralKind().KW_BNODE() \
            else 'nonliteral' if ctx.nonLiteralKind().KW_NONLITERAL() \
            else 'undefined'
        self.visitChildren(ctx)

    def visitNodeConstraintDatatype(self, ctx: ShExDocParser.NodeConstraintDatatypeContext):
        """ nodeConstraint: datatype xsFacet* # nodeConstraintDatatype """
        self.nodeconstraint.datatype = self.context.iri_to_IRI(ctx.datatype().iri())
        self.visitChildren(ctx)

    def visitNodeConstraintValueSet(self, ctx: ShExDocParser.NodeConstraintValueSetContext):
        """ nodeConstraint: valueSet xsFacet* #nodeConstraintValueSet """
        self.nodeconstraint.values = []
        self.visitChildren(ctx)

    def visitValueSetValue(self, ctx: ShExDocParser.ValueSetValueContext):
        """ valueSetValue: iriRange | literalRange | languageRange | 
            '.' (iriExclusion+ | literalExclusion+ | languageExclusion+) """
        if ctx.iriRange() or ctx.literalRange() or ctx.languageRange():
            self.visitChildren(ctx)
        else:
            if ctx.iriExclusion():
                stem = IriStemRange(Wildcard(), [])
                self._iri_exclusions(stem, ctx.iriExclusion())
            elif ctx.literalExclusion():
                stem = LiteralStemRange(Wildcard(), [])
                self._literal_exclusions(stem, ctx.literalExclusion())
            else:
                stem = LanguageStemRange(Wildcard(), [])
                self._language_exclusions(stem, ctx.languageExclusion())
            self.nodeconstraint.values.append(stem)

    def visitIriRange(self, ctx: ShExDocParser.IriRangeContext):
        """ iriRange: iri (STEM_MARK iriExclusion*)? """
        baseiri = self.context.iri_to_IRI(ctx.iri())
        if not ctx.STEM_MARK():
            stem = baseiri
        else:
            stem = IriStemRange(baseiri)
            if ctx.iriExclusion():
                stem.exclusions = []
            self._iri_exclusions(stem, ctx.iriExclusion())
        self.nodeconstraint.values.append(stem)

    def _iri_exclusions(self, stem: IriStemRange, exclusions: List[ShExDocParser.IriExclusionContext]) -> None:
        for excl in exclusions:
            excl_iri = self.context.iri_to_IRI(excl.iri())
            stem.exclusions.append(IriStem(excl_iri) if excl.STEM_MARK() else excl_iri)

    def visitLiteralRange(self, ctx: ShExDocParser.LiteralRangeContext):
        """ literalRange: literal (STEM_MARK literalExclusion*)? """
        baseliteral = self.context.literal_to_ObjectLiteral(ctx.literal())
        if not ctx.STEM_MARK():
            stem = baseliteral
        else:
            stem = LiteralStemRange(baseliteral)
            if ctx.literalExclusion():
                stem.exclusions = []
                self._literal_exclusions(stem, ctx.literalExclusion())
        self.nodeconstraint.values.append(stem)

    def _literal_exclusions(self, stem: LiteralStemRange,
                            exclusions: List[ShExDocParser.LiteralExclusionContext]) -> None:
        for excl in exclusions:
            excl_literal = self.context.literal_to_ObjectLiteral(excl.literal())
            stem.exclusions.append(LiteralStem(excl_literal) if excl.STEM_MARK() else excl_literal)

    def visitLanguageRange(self, ctx: ShExDocParser.LanguageRangeContext):
        """ langRange : LANGTAG (STEM_MARK languagExclusion*)? """
        baselang = ObjectLiteral(value=ctx.LANGTAG().getText()[1:])
        if not ctx.STEM_MARK():
            stem = LanguageStem(baselang)
        else:
            stem = LanguageStemRange(baselang)
            if ctx.languageExclusion():
                stem.exclusions = []
                self._language_exclusions(stem, ctx.languageExclusion())
        self.nodeconstraint.values.append(stem)

    @staticmethod
    def _language_exclusions(stem: LanguageStemRange,
                             exclusions: List[ShExDocParser.LanguageExclusionContext]) -> None:
        for excl in exclusions:
            excl_langtag = ObjectLiteral(value=excl.LANGTAG().getText()[1:])
            stem.exclusions.append(LanguageStem(excl_langtag) if excl.STEM_MARK() else excl_langtag)

    def visitStringFacet(self, ctx: ShExDocParser.StringFacetContext):
        """ stringFacet: stringLength INTEGER | REGEXP REGEXP_FLAGS
            stringLength: KW_LENGTH | KW_MINLENGTH | KW_MAXLENGTH """
        if ctx.stringLength():
            slen = INTEGER(ctx.INTEGER().getText())
            if ctx.stringLength().KW_LENGTH():
                self.nodeconstraint.length = slen
            elif ctx.stringLength().KW_MINLENGTH():
                self.nodeconstraint.minlength = slen
            else:
                self.nodeconstraint.maxlength = slen

        else:
            p = ctx.REGEXP().getText()[1:-1]
            p = p.replace('\/', '/').replace("\\'", "'")
            self.nodeconstraint.pattern = STRING(p)
            # TODO: Clean this up
            # self.nodeconstraint.pattern = STRING(ctx.REGEXP().getText()[1:-1].
            #                                      encode('utf-8').
            #                                      decode('unicode-escape').
            #                                      replace(r'\/', '/').
            #                                      replace(r'\\', '\\\\'))
            if ctx.REGEXP_FLAGS():
                self.nodeconstraint.flags = STRING(ctx.REGEXP_FLAGS().getText())

    def visitNumericFacet(self, ctx: ShExDocParser.NumericFacetContext):
        """ numericFacet: numericRange numericLiteral | numericLength INTEGER
            numericRange: KW_MINEXCLUSIVE | KW_MININCLUSIVE | KW_MAXEXCLUSIVE | KW_MAXINCLUSIVE 
            numericLength: KW_TOTALDIGITS | KW_FRACTIONDIGITS """
        if ctx.numericRange():
            numlit = self.context.numeric_literal_to_type(ctx.numericLiteral())
            if ctx.numericRange().KW_MINEXCLUSIVE():
                self.nodeconstraint.minexclusive = numlit
            elif ctx.numericRange().KW_MAXEXCLUSIVE():
                self.nodeconstraint.maxexclusive = numlit
            elif ctx.numericRange().KW_MININCLUSIVE():
                self.nodeconstraint.mininclusive = numlit
            elif ctx.numericRange().KW_MAXINCLUSIVE():
                self.nodeconstraint.maxinclusive = numlit
        else:
            nlen = INTEGER(ctx.INTEGER().getText())
            if ctx.numericLength().KW_TOTALDIGITS():
                self.nodeconstraint.totaldigits = nlen
            elif ctx.numericLength().KW_FRACTIONDIGITS():
                self.nodeconstraint.fractiondigits = nlen
