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
from typing import Union

from pyshexc.parser.ShExDocParser import ShExDocParser
from pyshexc.shexj.ShExJ import IRI, STRING, BOOL, Schema, BNODE, INTEGER, DECIMAL, DOUBLE, ObjectLiteral, Shape, \
    LANGTAG

IRIstr = str
PREFIXstr = str

RDF_TYPE = IRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")

RDF_INTEGER_TYPE = STRING("http://www.w3.org/2001/XMLSchema#integer")
RDF_DOUBLE_TYPE = STRING("http://www.w3.org/2001/XMLSchema#double")
RDF_DECIMAL_TYPE = STRING("http://www.w3.org/2001/XMLSchema#decimal")
RDF_BOOL_TYPE = STRING("http://www.w3.org/2001/XMLSchema#boolean")

BOOL_TRUE = BOOL('true')
BOOL_FALSE = BOOL('false')


class ParserContext:
    """
    Context maintained across ShExC parser implementation modules
    """
    def __init__(self):
        self.schema = Schema()
        self.ld_prefixes = {}       # Dict[PREFIXstr, IRIstr] - prefixes in the JSON-LD module
        self.prefixes = {}          # Dict[PREFIXstr, IRIstr]
        self.base = None            # IRIstr

    def _lookup_prefix(self, prefix: PREFIXstr) -> str:
        if len(prefix) == 0:
            return self.base
        elif prefix in self.ld_prefixes:
            return self.ld_prefixes[prefix]
        else:
            return self.prefixes.get(prefix, "")

    def iriref_to_str(self, ref: ShExDocParser.IRIREF) -> str:
        """  '<' (~[\u0000-\u0020=<>\"{}|^`\\] | UCHAR)* '>' """
        rval = ref.getText()[1:-1].encode('utf-8').decode('unicode-escape')
        return rval if ':' in rval else self.base.val + rval

    def iriref_to_IRI(self, ref: ShExDocParser.IRIREF) -> IRI:
        """  '<' (~[\u0000-\u0020=<>\"{}|^`\\] | UCHAR)* '>' """
        return IRI(self.iriref_to_str(ref))

    def prefixedname_to_str(self, prefix: ShExDocParser.PrefixedNameContext) -> str:
        """ prefixedName: PNAME_LN | PNAME_NS
            PNAME_NS: PN_PREFIX? ':' ;
            PNAME_LN: PNAME_NS PN_LOCAL ;
        """
        if prefix.PNAME_NS():
            return self._lookup_prefix(prefix.PNAME_NS().getText())
        else:
            prefix, local = prefix.PNAME_LN().getText().split(':', 1)
            return self._lookup_prefix(prefix + ':') + (local if local else "")

    def prefixedname_to_IRI(self, prefix: ShExDocParser.PrefixedNameContext) -> IRI:
        """ prefixedName: PNAME_LN | PNAME_NS
            PNAME_NS: PN_PREFIX? ':' ;
            PNAME_LN: PNAME_NS PN_LOCAL ;
        """
        return IRI(self.prefixedname_to_str(prefix))

    def shapeRef_to_IRI(self, ref: ShExDocParser.ShapeRefContext) -> IRI:
        """ shapeOrRef: ATPNAME_NS | ATPNAME_LN | '@' shapeExprLabel """
        if ref.ATPNAME_NS():
            return IRI(self._lookup_prefix(ref.ATPNAME_NS().getText()[1:]))
        elif ref.ATPNAME_LN():
            prefix, local = ref.ATPNAME_LN().getText()[1:].split(':', 1)
            return IRI(self._lookup_prefix(prefix + ':') + (local if local else ""))
        else:
            return self.shapeexprlabel_to_IRI(ref.shapeExprLabel())

    def iri_to_str(self, iri_: ShExDocParser.IriContext) -> str:
        """ iri: IRIREF | prefixedName 
        """
        if iri_.IRIREF():
            return self.iriref_to_str(iri_.IRIREF())
        else:
            return self.prefixedname_to_str(iri_.prefixedName())

    def iri_to_IRI(self, iri_: ShExDocParser.IriContext) -> IRI:
        """ iri: IRIREF | prefixedName 
            prefixedName: PNAME_LN | PNAME_NS 
        """
        return IRI(self.iri_to_str(iri_))

    def tripleexprlabel_to_IRI(self, tripleExprLabel: ShExDocParser.TripleExprLabelContext) -> Union[BNODE, IRI]:
        """ tripleExprLabel: ri | blankNode """
        if tripleExprLabel.iri():
            return self.iri_to_IRI(tripleExprLabel.iri())
        else:
            return BNODE(tripleExprLabel.blankNode().getText())

    def shapeexprlabel_to_IRI(self, shapeExprLabel: ShExDocParser.ShapeExprLabelContext) -> Union[BNODE, IRI]:
        """ shapeExprLabel: iri | blankNode """
        if shapeExprLabel.iri():
            return self.iri_to_IRI(shapeExprLabel.iri())
        else:
            return BNODE(shapeExprLabel.blankNode().getText())

    def predicate_to_IRI(self, predicate: ShExDocParser.PredicateContext) -> IRI:
        """ predicate: iri | rdfType """
        if predicate.iri():
            return self.iri_to_IRI(predicate.iri())
        else:
            return RDF_TYPE

    @staticmethod
    def numeric_literal_to_type(numlit: ShExDocParser.NumericLiteralContext) -> Union[INTEGER, DECIMAL, DOUBLE]:
        """ numericLiteral: INTEGER | DECIMAL | DOUBLE """
        if numlit.INTEGER():
            rval = INTEGER(numlit.INTEGER().getText())
        elif numlit.DECIMAL():
            rval = DECIMAL(numlit.DECIMAL().getText())
        else:
            rval = DOUBLE(numlit.DOUBLE().getText())
        return rval

    def literal_to_ObjectLiteral(self, literal: ShExDocParser.LiteralContext) -> ObjectLiteral:
        """ literal: rdfLiteral | numericLiteral | booleanLiteral """
        rval = ObjectLiteral()
        if literal.rdfLiteral():
            rdflit = literal.rdfLiteral()
            txt = rdflit.string().getText()
            txt = txt[3:-3] \
                if len(txt) > 5 and (txt.startswith("'''") and txt.endswith("'''") or
                                     txt.startswith('"""') and txt.endswith('"""')) else txt[1:-1]
            rval.value = STRING(txt.replace("\\'", "'").replace('\\"', '"'))
            if rdflit.LANGTAG():
                rval.language = LANGTAG(rdflit.LANGTAG().getText()[1:].lower())
            if rdflit.datatype():
                rval.type = self.iri_to_str(rdflit.datatype().iri())
        elif literal.numericLiteral():
            numlit = literal.numericLiteral()
            if numlit.INTEGER():
                rval.value = STRING(numlit.INTEGER().getText())
                rval.type = RDF_INTEGER_TYPE
            elif numlit.DECIMAL():
                rval.value = STRING(numlit.DECIMAL().getText())
                rval.type = RDF_DECIMAL_TYPE
            elif numlit.DOUBLE():
                rval.value = STRING(numlit.DOUBLE().getText())
                rval.type = RDF_DOUBLE_TYPE
        elif literal.booleanLiteral():
            rval.value = STRING(literal.booleanLiteral().getText().lower())
            rval.type = RDF_BOOL_TYPE
        return rval

    @staticmethod
    def is_empty_shape(sh: Shape) -> bool:
        """ Determine whether sh has any value """
        return sh.closed is None and sh.expression is None and sh.extra is None and sh.inherit is None and \
            sh.inherit is None and sh.semActs is None and sh.virtual is None
