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

from ShExJSG.ShExJ import Language, NodeConstraint, Wildcard, IriStemRange, LiteralStemRange, \
    LanguageStemRange, IriStem, LiteralStem, LanguageStem, LANGTAG
from pyjsg.jsglib import jsg

from pyshexc.parser.ShExDocParser import ShExDocParser
from pyshexc.parser.ShExDocVisitor import ShExDocVisitor
from pyshexc.parser_impl.parser_context import ParserContext


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
        self.nodeconstraint.datatype = self.context.iri_to_iriref(ctx.datatype().iri())
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
        else:                               # '.' branch - wild card with exclusions
            if ctx.iriExclusion():
                vs_value = IriStemRange(Wildcard(), [])
                self._iri_exclusions(vs_value, ctx.iriExclusion())
            elif ctx.literalExclusion():
                vs_value = LiteralStemRange(Wildcard(), [])
                self._literal_exclusions(vs_value, ctx.literalExclusion())
            else:
                vs_value = LanguageStemRange(Wildcard(), [])
                self._language_exclusions(vs_value, ctx.languageExclusion())
            self.nodeconstraint.values.append(vs_value)

    def visitIriRange(self, ctx: ShExDocParser.IriRangeContext):
        """ iriRange: iri (STEM_MARK iriExclusion*)? """
        baseiri = self.context.iri_to_iriref(ctx.iri())
        if not ctx.STEM_MARK():
            vsvalue = baseiri                   # valueSetValue = objectValue  / objectValue = IRI
        else:
            if ctx.iriExclusion():              # valueSetValue = IriStemRange / iriStemRange = stem + exclusions
                vsvalue = IriStemRange(baseiri, exclusions=[])
                self._iri_exclusions(vsvalue, ctx.iriExclusion())
            else:
                vsvalue = IriStem(baseiri)      # valueSetValue = IriStem  /  IriStem: {stem:IRI}
        self.nodeconstraint.values.append(vsvalue)

    def _iri_exclusions(self, stemrange: IriStemRange, exclusions: List[ShExDocParser.IriExclusionContext]) -> None:
        for excl in exclusions:
            excl_iri = self.context.iri_to_iriref(excl.iri())
            stemrange.exclusions.append(IriStem(excl_iri) if excl.STEM_MARK() else excl_iri)

    def visitLiteralRange(self, ctx: ShExDocParser.LiteralRangeContext):
        """ ShExC: literalRange: literal (STEM_MARK literalExclusion*)? 
            ShExJ: valueSetValue: objectValue | LiteralStem | LiteralStemRange
                   literalStem: {stem:STRING}
                   literalStemRange: {stem:STRING exclusions:[STRING|LiteralStem+]?
        """
        baseliteral = self.context.literal_to_ObjectLiteral(ctx.literal())
        if not ctx.STEM_MARK():
            vsvalue = baseliteral               # valueSetValue = objectValue / objectValue = ObjectLiteral
        else:
            if ctx.literalExclusion():          # valueSetValue = LiteralStemRange /
                vsvalue = LiteralStemRange(baseliteral.value.val, exclusions=[])
                self._literal_exclusions(vsvalue, ctx.literalExclusion())
            else:
                vsvalue = LiteralStem(baseliteral.value)
        self.nodeconstraint.values.append(vsvalue)

    def _literal_exclusions(self, stem: LiteralStemRange,
                            exclusions: List[ShExDocParser.LiteralExclusionContext]) -> None:
        """ ShExC: literalExclusion = '-' literal STEM_MARK?
            ShExJ: exclusions: [STRING|LiteralStem +]
                   literalStem: {stem:STRING}
        """
        for excl in exclusions:
            excl_literal_v = self.context.literal_to_ObjectLiteral(excl.literal()).value
            stem.exclusions.append(LiteralStem(excl_literal_v) if excl.STEM_MARK() else excl_literal_v)

    def visitLanguageRange(self, ctx: ShExDocParser.LanguageRangeContext):
        """ ShExC: languageRange : LANGTAG (STEM_MARK languagExclusion*)? 
            ShExJ: valueSetValue = objectValue | LanguageStem | LanguageStemRange """
        baselang = ctx.LANGTAG().getText()
        if not ctx.STEM_MARK():                 # valueSetValue = objectValue / objectValue = ObjectLiteral
            vsvalue = Language()
            vsvalue.languageTag = baselang[1:]
        else:
            if ctx.languageExclusion():
                vsvalue = LanguageStemRange(LANGTAG(baselang[1:]), exclusions=[])
                self._language_exclusions(vsvalue, ctx.languageExclusion())
            else:
                vsvalue = LanguageStem(LANGTAG(baselang[1:]))
        self.nodeconstraint.values.append(vsvalue)

    @staticmethod
    def _language_exclusions(stem: LanguageStemRange,
                             exclusions: List[ShExDocParser.LanguageExclusionContext]) -> None:
        """ languageExclusion = '-' LANGTAG STEM_MARK?"""
        for excl in exclusions:
            excl_langtag = LANGTAG(excl.LANGTAG().getText()[1:])
            stem.exclusions.append(LanguageStem(excl_langtag) if excl.STEM_MARK() else excl_langtag)

    def visitStringFacet(self, ctx: ShExDocParser.StringFacetContext):
        """ stringFacet: stringLength INTEGER | REGEXP REGEXP_FLAGS
            stringLength: KW_LENGTH | KW_MINLENGTH | KW_MAXLENGTH """
        if ctx.stringLength():
            slen = jsg.Integer(ctx.INTEGER().getText())
            if ctx.stringLength().KW_LENGTH():
                self.nodeconstraint.length = slen
            elif ctx.stringLength().KW_MINLENGTH():
                self.nodeconstraint.minlength = slen
            else:
                self.nodeconstraint.maxlength = slen

        else:
            self.nodeconstraint.pattern = jsg.String(self.context.fix_re_escapes(ctx.REGEXP().getText()[1:-1]))
            if ctx.REGEXP_FLAGS():
                self.nodeconstraint.flags = jsg.String(ctx.REGEXP_FLAGS().getText())

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
            nlen = jsg.Integer(ctx.INTEGER().getText())
            if ctx.numericLength().KW_TOTALDIGITS():
                self.nodeconstraint.totaldigits = nlen
            elif ctx.numericLength().KW_FRACTIONDIGITS():
                self.nodeconstraint.fractiondigits = nlen
