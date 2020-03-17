/* $Id: Langname_Parser.yy,v 1.11 2014-04-07 13:02:07 eric Exp spec_2_1Parser.yy 19 2007-08-19 20:36:24Z tb $ -*- mode: c++ -*- */
/** \file shex-to-jsonParser.yy Contains the Bison parser source */

/*** yacc/bison Declarations ***/

/* Require bison 2.3 or later */
%require "2.3"

/* add debug output code to generated parser. disable this for release
 * versions. */
%debug

/* start symbol is named "start" */
%start shexDoc

/* write out a header file containing the token defines */
%defines

%language "c++"

/* use newer C++ skeleton file */
%skeleton "lalr1.cc"

/* namespace to enclose parser in */
%define api.prefix {shexParseToXml}

/* set the parser's class identifier */
%define api.parser.class {shexParseToXmlParser}

%define api.value.type {variant}

/* keep track of the current position within the input */
%locations
%initial-action
{
    // initialize the initial location object
    @$.begin.filename = @$.end.filename = &driver.streamname;
};

/* The driver is passed by reference to the parser and to the scanner. This
 * provides a simple but effective pure interface, not relying on global
 * variables. */
%parse-param { class Driver& driver }

/* verbose error messages */
%define parse.error verbose

%code requires { // ##bison2
//%{ /*** C/C++ Declarations ***/

#include <fstream>
#include <iostream>
#include <sstream>
#include <vector>
#include <exception>
#include <cassert>
#include <string.h>
#include <memory>
#include <boost/variant.hpp>
  // typedef boost::variant variant;
  // using namespace boost;
using semtype = boost::variant<int, std::string>;

#define SHEXPARSETOXMLSTYPE semtype;

#include "shex.hh"
extern std::ostream* _Trace;

/* START ClassBlock */
namespace yacker {

} // namespace yacker

namespace shexParseToXml {



}
class ProgramFlowException : public std::exception {
private:
    const char* msg;
public:
    ProgramFlowException(const char* msg) : msg(msg) { }

    // This declaration is not useless:
    // http://gcc.gnu.org/onlinedocs/gcc-3.0.2/gcc_6.html#SEC118
    virtual ~ProgramFlowException() throw() { }

    // See comment in eh_exception.cc.
    virtual const char* what () const throw() { return msg; }
  // @@ was virtual const char* what () { return msg; }
};
/* END ClassBlock */

namespace shexParseToXml {

/** The Driver class brings together all components. It creates an instance of
 * the shexParseToXmlParser and shexParseToXmlScanner classes and connects them. Then the input stream is
 * fed into the scanner object and the parser gets it's token
 * sequence. Furthermore the driver object is available in the grammar rules as
 * a parameter. Therefore the driver class contains a reference to the
 * structure into which the parsed data is saved. */
class shexParseToXmlContext {
public:
    ~shexParseToXmlContext()
    {
    }
};

class Driver
{
public:
    /// construct a new parser driver context
    Driver(shexParseToXmlContext& context);

    /// enable debug output in the flex scanner
    bool trace_scanning;

    /// enable debug output in the bison parser
    bool trace_parsing;

    /// stream name (file or input stream) used for error messages.
    std::string streamname;

    /** Invoke the scanner and parser for a stream.
     * @param in	input stream
     * @param sname	stream name for error messages
     * @return		true if successfully parsed
     */
    bool parse_stream(std::istream& in,
		      const std::string& sname = "stream input");

    /** Invoke the scanner and parser on an input string.
     * @param input	input string
     * @param sname	stream name for error messages
     * @return		true if successfully parsed
     */
    bool parse_string(const std::string& input,
		      const std::string& sname = "string stream");

    /** Invoke the scanner and parser on a file. Use parse_stream with a
     * std::ifstream if detection of file reading errors is required.
     * @param filename	input file name
     * @return		true if successfully parsed
     */
    bool parse_file(const std::string& filename);

    // To demonstrate pure handling of parse errors, instead of
    // simply dumping them on the standard error output, we will pass
    // them to the driver using the following two member functions.

    /** Error handling with associated line number. This can be modified to
     * output the error e.g. to a dialog box. */
    void error(const class location& l, const std::string& m);

    /** General error handling. This can be modified to output the error
     * e.g. to a dialog box. */
    void error(const std::string& m);

    /** Pointer to the current lexer instance, this is used to connect the
     * parser to the scanner. It is used in the yylex macro. */
    class shexParseToXmlScanner* lexer;

    /** Reference to the context filled during parsing of the expressions. */
    shexParseToXmlContext& context;
    // shexDoc* root;
};

} // namespace shexParseToXml

//%}
} // bison 2

 /*** BEGIN shexParseToXml - Change the grammar's tokens below ***/
// %union {
//     /* Terminals */
//   const char* str;
//   shex::STRING p_STRING;
//   shex::IRIREF p_IRIREF;
//     /* Productions */
  
// }

%{
#include "shex-to-jsonScanner.hh"
%}
%token			__EOF__	     0	"end of file"
/* START TokenBlock */
/* Terminals */
%token <p_IT_BASE> IT_BASE
%token <p_IT_PREFIX> IT_PREFIX
%token <p_IT_IMPORT> IT_IMPORT
%token <p_IT_START> IT_START
%token <p_GT_EQUAL> GT_EQUAL
%token <p_IT_EXTERNAL> IT_EXTERNAL
%token <p_IT_OR> IT_OR
%token <p_IT_AND> IT_AND
%token <p_IT_NOT> IT_NOT
%token <p_GT_LPAREN> GT_LPAREN
%token <p_GT_RPAREN> GT_RPAREN
%token <p_GT_DOT> GT_DOT
%token <p_GT_AT> GT_AT
%token <p_IT_LITERAL> IT_LITERAL
%token <p_IT_IRI> IT_IRI
%token <p_IT_BNODE> IT_BNODE
%token <p_IT_NONLITERAL> IT_NONLITERAL
%token <p_IT_LENGTH> IT_LENGTH
%token <p_IT_MINLENGTH> IT_MINLENGTH
%token <p_IT_MAXLENGTH> IT_MAXLENGTH
%token <p_IT_MININCLUSIVE> IT_MININCLUSIVE
%token <p_IT_MINEXCLUSIVE> IT_MINEXCLUSIVE
%token <p_IT_MAXINCLUSIVE> IT_MAXINCLUSIVE
%token <p_IT_MAXEXCLUSIVE> IT_MAXEXCLUSIVE
%token <p_IT_TOTALDIGITS> IT_TOTALDIGITS
%token <p_IT_FRACTIONDIGITS> IT_FRACTIONDIGITS
%token <p_GT_LCURLEY> GT_LCURLEY
%token <p_GT_RCURLEY> GT_RCURLEY
%token <p_IT_CLOSED> IT_CLOSED
%token <p_IT_EXTRA> IT_EXTRA
%token <p_GT_PIPE> GT_PIPE
%token <p_GT_SEMI> GT_SEMI
%token <p_GT_DOLLAR> GT_DOLLAR
%token <p_GT_TIMES> GT_TIMES
%token <p_GT_PLUS> GT_PLUS
%token <p_GT_OPT> GT_OPT
%token <p_GT_CARROT> GT_CARROT
%token <p_GT_LBRACKET> GT_LBRACKET
%token <p_GT_RBRACKET> GT_RBRACKET
%token <p_GT_KINDA> GT_KINDA
%token <p_GT_MINUS> GT_MINUS
%token <p_GT_AMP> GT_AMP
%token <p_GT_DIVIDE_DIVIDE> GT_DIVIDE_DIVIDE
%token <p_GT_MODULO> GT_MODULO
%token <p_GT_DTYPE> GT_DTYPE
%token <p_IT_true> IT_true
%token <p_IT_false> IT_false
%token <p_CODE> CODE
%token <p_REPEAT_RANGE> REPEAT_RANGE
%token <p_RDF_TYPE> RDF_TYPE
%token <p_IRIREF> IRIREF
%token <p_PNAME_NS> PNAME_NS
%token <p_PNAME_LN> PNAME_LN
%token <p_ATPNAME_NS> ATPNAME_NS
%token <p_ATPNAME_LN> ATPNAME_LN
%token <p_REGEXP> REGEXP
%token <p_BLANK_NODE_LABEL> BLANK_NODE_LABEL
%token <p_LANGTAG> LANGTAG
%token <p_INTEGER> INTEGER
%token <p_DECIMAL> DECIMAL
%token <p_DOUBLE> DOUBLE
%token <p_STRING_LITERAL1> STRING_LITERAL1
%token <p_STRING_LITERAL2> STRING_LITERAL2
%token <p_STRING_LITERAL_LONG1> STRING_LITERAL_LONG1
%token <p_STRING_LITERAL_LONG2> STRING_LITERAL_LONG2
%token <p_LANG_STRING_LITERAL1> LANG_STRING_LITERAL1
%token <p_LANG_STRING_LITERAL2> LANG_STRING_LITERAL2
%token <p_LANG_STRING_LITERAL_LONG1> LANG_STRING_LITERAL_LONG1
%token <p_LANG_STRING_LITERAL_LONG2> LANG_STRING_LITERAL_LONG2

/* Productions */
// %type <p_shexDoc> shexDoc
// %type <p_directives> _Qdirective_E_Star
// %type <p__O_QnotStartAction_E_Or_QstartActions_E_C> _O_QnotStartAction_E_Or_QstartActions_E_C
// %type <p_statements> _Qstatement_E_Star
// %type <p__O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C> _O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C
// %type <p__Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt> _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt
// %type <p_directive> directive
// %type <p_baseDecl> baseDecl
// %type <p_prefixDecl> prefixDecl
// %type <p_importDecl> importDecl
// %type <p_notStartAction> notStartAction
// %type <p_start> start
// %type <p_startActions> startActions
// %type <p_codeDecls> _QcodeDecl_E_Plus
// %type <p_statement> statement
// %type <p_shapeExprDecl> shapeExprDecl
// %type <p__O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C> _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C
// %type <p_shapeExpression> shapeExpression
// %type <p_inlineShapeExpression> inlineShapeExpression
// %type <p_shapeOr> shapeOr
// %type <p__O_QIT_OR_E_S_QshapeAnd_E_C> _O_QIT_OR_E_S_QshapeAnd_E_C
// %type <p__O_QIT_OR_E_S_QshapeAnd_E_Cs> _Q_O_QIT_OR_E_S_QshapeAnd_E_C_E_Star
// %type <p_inlineShapeOr> inlineShapeOr
// %type <p__O_QIT_OR_E_S_QinlineShapeAnd_E_C> _O_QIT_OR_E_S_QinlineShapeAnd_E_C
// %type <p__O_QIT_OR_E_S_QinlineShapeAnd_E_Cs> _Q_O_QIT_OR_E_S_QinlineShapeAnd_E_C_E_Star
// %type <p_shapeAnd> shapeAnd
// %type <p__O_QIT_AND_E_S_QshapeNot_E_C> _O_QIT_AND_E_S_QshapeNot_E_C
// %type <p__O_QIT_AND_E_S_QshapeNot_E_Cs> _Q_O_QIT_AND_E_S_QshapeNot_E_C_E_Star
// %type <p_inlineShapeAnd> inlineShapeAnd
// %type <p__O_QIT_AND_E_S_QinlineShapeNot_E_C> _O_QIT_AND_E_S_QinlineShapeNot_E_C
// %type <p__O_QIT_AND_E_S_QinlineShapeNot_E_Cs> _Q_O_QIT_AND_E_S_QinlineShapeNot_E_C_E_Star
// %type <p_shapeNot> shapeNot
// %type <p__QIT_NOT_E_Opt> _QIT_NOT_E_Opt
// %type <p_inlineShapeNot> inlineShapeNot
// %type <p_shapeAtom> shapeAtom
// %type <p__QshapeOrRef_E_Opt> _QshapeOrRef_E_Opt
// %type <p__QnonLitNodeConstraint_E_Opt> _QnonLitNodeConstraint_E_Opt
// %type <p_inlineShapeAtom> inlineShapeAtom
// %type <p__QinlineShapeOrRef_E_Opt> _QinlineShapeOrRef_E_Opt
// %type <p_shapeOrRef> shapeOrRef
// %type <p_inlineShapeOrRef> inlineShapeOrRef
// %type <p_shapeRef> shapeRef
// %type <p_litNodeConstraint> litNodeConstraint
// %type <p_xsFacets> _QxsFacet_E_Star
// %type <p_numericFacets> _QnumericFacet_E_Plus
// %type <p_nonLitNodeConstraint> nonLitNodeConstraint
// %type <p_stringFacets> _QstringFacet_E_Star
// %type <p_stringFacets> _QstringFacet_E_Plus
// %type <p_nonLiteralKind> nonLiteralKind
// %type <p_xsFacet> xsFacet
// %type <p_stringFacet> stringFacet
// %type <p_stringLength> stringLength
// %type <p_numericFacet> numericFacet
// %type <p_numericRange> numericRange
// %type <p_numericLength> numericLength
// %type <p_shapeDefinition> shapeDefinition
// %type <p_annotations> _Qannotation_E_Star
// %type <p_inlineShapeDefinition> inlineShapeDefinition
// %type <p__QtripleExpression_E_Opt> _QtripleExpression_E_Opt
// %type <p_shapeQualifiers> shapeQualifiers
// %type <p__O_QextraPropertySet_E_Or_QIT_CLOSED_E_C> _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C
// %type <p__O_QextraPropertySet_E_Or_QIT_CLOSED_E_Cs> _Q_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_E_Star
// %type <p_extraPropertySet> extraPropertySet
// %type <p_predicates> _Qpredicate_E_Plus
// %type <p_tripleExpression> tripleExpression
// %type <p_oneOfTripleExpr> oneOfTripleExpr
// %type <p_multiElementOneOf> multiElementOneOf
// %type <p__O_QGT_PIPE_E_S_QgroupTripleExpr_E_C> _O_QGT_PIPE_E_S_QgroupTripleExpr_E_C
// %type <p__O_QGT_PIPE_E_S_QgroupTripleExpr_E_Cs> _Q_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C_E_Plus
// %type <p_groupTripleExpr> groupTripleExpr
// %type <p_singleElementGroup> singleElementGroup
// %type <p__QGT_SEMI_E_Opt> _QGT_SEMI_E_Opt
// %type <p_multiElementGroup> multiElementGroup
// %type <p__O_QGT_SEMI_E_S_QunaryTripleExpr_E_C> _O_QGT_SEMI_E_S_QunaryTripleExpr_E_C
// %type <p__O_QGT_SEMI_E_S_QunaryTripleExpr_E_Cs> _Q_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C_E_Plus
// %type <p_unaryTripleExpr> unaryTripleExpr
// %type <p__O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C> _O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C
// %type <p__Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt> _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt
// %type <p__O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C> _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C
// %type <p_bracketedTripleExpr> bracketedTripleExpr
// %type <p__Qcardinality_E_Opt> _Qcardinality_E_Opt
// %type <p_tripleConstraint> tripleConstraint
// %type <p__QsenseFlags_E_Opt> _QsenseFlags_E_Opt
// %type <p_cardinality> cardinality
// %type <p_senseFlags> senseFlags
// %type <p_valueSet> valueSet
// %type <p_valueSetValues> _QvalueSetValue_E_Star
// %type <p_valueSetValue> valueSetValue
// %type <p_iriExclusions> _QiriExclusion_E_Plus
// %type <p_literalExclusions> _QliteralExclusion_E_Plus
// %type <p_languageExclusions> _QlanguageExclusion_E_Plus
// %type <p__O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C> _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C
// %type <p_iriRange> iriRange
// %type <p_iriExclusions> _QiriExclusion_E_Star
// %type <p__O_QGT_KINDA_E_S_QiriExclusion_E_Star_C> _O_QGT_KINDA_E_S_QiriExclusion_E_Star_C
// %type <p__Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt> _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt
// %type <p_iriExclusion> iriExclusion
// %type <p__QGT_KINDA_E_Opt> _QGT_KINDA_E_Opt
// %type <p_literalRange> literalRange
// %type <p_literalExclusions> _QliteralExclusion_E_Star
// %type <p__O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C> _O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C
// %type <p__Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt> _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt
// %type <p_literalExclusion> literalExclusion
// %type <p_languageRange> languageRange
// %type <p_languageExclusions> _QlanguageExclusion_E_Star
// %type <p__O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C> _O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C
// %type <p__Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt> _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt
// %type <p_languageExclusion> languageExclusion
// %type <p_include> include
// %type <p_annotation> annotation
// %type <p__O_Qiri_E_Or_Qliteral_E_C> _O_Qiri_E_Or_Qliteral_E_C
// %type <p_semanticActions> semanticActions
// %type <p_codeDecls> _QcodeDecl_E_Star
// %type <p_codeDecl> codeDecl
// %type <p__O_QCODE_E_Or_QGT_MODULO_E_C> _O_QCODE_E_Or_QGT_MODULO_E_C
// %type <p_literal> literal
// %type <p_predicate> predicate
// %type <p_datatype> datatype
// %type <p_shapeExprLabel> shapeExprLabel
// %type <p_tripleExprLabel> tripleExprLabel
// %type <p_rawNumeric> rawNumeric
// %type <p_numericLiteral> numericLiteral
// %type <p_rdfLiteral> rdfLiteral
// %type <p__O_QGT_DTYPE_E_S_Qdatatype_E_C> _O_QGT_DTYPE_E_S_Qdatatype_E_C
// %type <p__Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt> _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt
// %type <p_booleanLiteral> booleanLiteral
// %type <p_string> string
// %type <p_langString> langString
// %type <p_iri> iri
// %type <p_prefixedName> prefixedName
// %type <p_blankNode> blankNode

/* END TokenBlock */

//%destructor { delete $$; } BlankNode

 /*** END shexParseToXml - Change the grammar's tokens above ***/

%{
#include <stdarg.h>
#include "shex-to-jsonScanner.hh"

/* this "connects" the bison parser in the driver to the flex scanner class
 * object. it defines the yylex() function call to pull the next token from the
 * current lexer object of the driver context. */
#undef yylex
#define yylex driver.lexer->lex
%}

%% /*** Grammar Rules ***/

 /*** BEGIN shexParseToXml - Change the grammar rules below ***/
shexDoc:
    _Qdirective_E_Star _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt	{
}
;

_Qdirective_E_Star:
    {
}

    | _Qdirective_E_Star directive	{
}
;

_O_QnotStartAction_E_Or_QstartActions_E_C:
    notStartAction	{
}

    | startActions	{
}
;

_Qstatement_E_Star:
    {
}

    | _Qstatement_E_Star statement	{
}
;

_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C:
    _O_QnotStartAction_E_Or_QstartActions_E_C _Qstatement_E_Star	{
}
;

_Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt:
    {
}

    | _O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C	{
}
;

directive:
    baseDecl	{
}

    | prefixDecl	{
}

    | importDecl	{
}
;

baseDecl:
    IT_BASE IRIREF	{
}
;

prefixDecl:
    IT_PREFIX PNAME_NS IRIREF	{
}
;

importDecl:
    IT_IMPORT IRIREF	{
}
;

notStartAction:
    start	{
}

    | shapeExprDecl	{
}
;

start:
    IT_START GT_EQUAL inlineShapeExpression	{
}
;

startActions:
    _QcodeDecl_E_Plus	{
}
;

_QcodeDecl_E_Plus:
    codeDecl	{
}

    | _QcodeDecl_E_Plus codeDecl	{
}
;

statement:
    directive	{
}

    | notStartAction	{
}
;

shapeExprDecl:
    shapeExprLabel _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C	{
}
;

_O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C:
    shapeExpression	{
}

    | IT_EXTERNAL	{
}
;

shapeExpression:
    shapeOr	{
}
;

inlineShapeExpression:
    inlineShapeOr	{
}
;

shapeOr:
    shapeAnd _Q_O_QIT_OR_E_S_QshapeAnd_E_C_E_Star	{
}
;

_O_QIT_OR_E_S_QshapeAnd_E_C:
    IT_OR shapeAnd	{
}
;

_Q_O_QIT_OR_E_S_QshapeAnd_E_C_E_Star:
    {
}

    | _Q_O_QIT_OR_E_S_QshapeAnd_E_C_E_Star _O_QIT_OR_E_S_QshapeAnd_E_C	{
}
;

inlineShapeOr:
    inlineShapeAnd _Q_O_QIT_OR_E_S_QinlineShapeAnd_E_C_E_Star	{
}
;

_O_QIT_OR_E_S_QinlineShapeAnd_E_C:
    IT_OR inlineShapeAnd	{
}
;

_Q_O_QIT_OR_E_S_QinlineShapeAnd_E_C_E_Star:
    {
}

    | _Q_O_QIT_OR_E_S_QinlineShapeAnd_E_C_E_Star _O_QIT_OR_E_S_QinlineShapeAnd_E_C	{
}
;

shapeAnd:
    shapeNot _Q_O_QIT_AND_E_S_QshapeNot_E_C_E_Star	{
}
;

_O_QIT_AND_E_S_QshapeNot_E_C:
    IT_AND shapeNot	{
}
;

_Q_O_QIT_AND_E_S_QshapeNot_E_C_E_Star:
    {
}

    | _Q_O_QIT_AND_E_S_QshapeNot_E_C_E_Star _O_QIT_AND_E_S_QshapeNot_E_C	{
}
;

inlineShapeAnd:
    inlineShapeNot _Q_O_QIT_AND_E_S_QinlineShapeNot_E_C_E_Star	{
}
;

_O_QIT_AND_E_S_QinlineShapeNot_E_C:
    IT_AND inlineShapeNot	{
}
;

_Q_O_QIT_AND_E_S_QinlineShapeNot_E_C_E_Star:
    {
}

    | _Q_O_QIT_AND_E_S_QinlineShapeNot_E_C_E_Star _O_QIT_AND_E_S_QinlineShapeNot_E_C	{
}
;

shapeNot:
    _QIT_NOT_E_Opt shapeAtom	{
}
;

_QIT_NOT_E_Opt:
    {
}

    | IT_NOT	{
}
;

inlineShapeNot:
    _QIT_NOT_E_Opt inlineShapeAtom	{
}
;

shapeAtom:
    nonLitNodeConstraint _QshapeOrRef_E_Opt	{
}

    | litNodeConstraint	{
}

    | shapeOrRef _QnonLitNodeConstraint_E_Opt	{
}

    | GT_LPAREN shapeExpression GT_RPAREN	{
}

    | GT_DOT	{
}
;

_QshapeOrRef_E_Opt:
    {
}

    | shapeOrRef	{
}
;

_QnonLitNodeConstraint_E_Opt:
    {
}

    | nonLitNodeConstraint	{
}
;

inlineShapeAtom:
    nonLitNodeConstraint _QinlineShapeOrRef_E_Opt	{
}

    | litNodeConstraint	{
}

    | inlineShapeOrRef _QnonLitNodeConstraint_E_Opt	{
}

    | GT_LPAREN shapeExpression GT_RPAREN	{
}

    | GT_DOT	{
}
;

_QinlineShapeOrRef_E_Opt:
    {
}

    | inlineShapeOrRef	{
}
;

shapeOrRef:
    shapeDefinition	{
}

    | shapeRef	{
}
;

inlineShapeOrRef:
    inlineShapeDefinition	{
}

    | shapeRef	{
}
;

shapeRef:
    ATPNAME_LN	{
}

    | ATPNAME_NS	{
}

    | GT_AT shapeExprLabel	{
}
;

litNodeConstraint:
    IT_LITERAL _QxsFacet_E_Star	{
}

    | datatype _QxsFacet_E_Star	{
}

    | valueSet _QxsFacet_E_Star	{
}

    | _QnumericFacet_E_Plus	{
}
;

_QxsFacet_E_Star:
    {
}

    | _QxsFacet_E_Star xsFacet	{
}
;

_QnumericFacet_E_Plus:
    numericFacet	{
}

    | _QnumericFacet_E_Plus numericFacet	{
}
;

nonLitNodeConstraint:
    nonLiteralKind _QstringFacet_E_Star	{
}

    | _QstringFacet_E_Plus	{
}
;

_QstringFacet_E_Star:
    {
}

    | _QstringFacet_E_Star stringFacet	{
}
;

_QstringFacet_E_Plus:
    stringFacet	{
}

    | _QstringFacet_E_Plus stringFacet	{
}
;

nonLiteralKind:
    IT_IRI	{
}

    | IT_BNODE	{
}

    | IT_NONLITERAL	{
}
;

xsFacet:
    stringFacet	{
}

    | numericFacet	{
}
;

stringFacet:
    stringLength INTEGER	{
}

    | REGEXP	{
}
;

stringLength:
    IT_LENGTH	{
}

    | IT_MINLENGTH	{
}

    | IT_MAXLENGTH	{
}
;

numericFacet:
    numericRange rawNumeric	{
}

    | numericLength INTEGER	{
}
;

numericRange:
    IT_MININCLUSIVE	{
}

    | IT_MINEXCLUSIVE	{
}

    | IT_MAXINCLUSIVE	{
}

    | IT_MAXEXCLUSIVE	{
}
;

numericLength:
    IT_TOTALDIGITS	{
}

    | IT_FRACTIONDIGITS	{
}
;

shapeDefinition:
    inlineShapeDefinition _Qannotation_E_Star semanticActions	{
}
;

_Qannotation_E_Star:
    {
}

    | _Qannotation_E_Star annotation	{
}
;

inlineShapeDefinition:
    shapeQualifiers GT_LCURLEY _QtripleExpression_E_Opt GT_RCURLEY	{
}
;

_QtripleExpression_E_Opt:
    {
}

    | tripleExpression	{
}
;

shapeQualifiers:
    _Q_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_E_Star	{
}
;

_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C:
    extraPropertySet	{
}

    | IT_CLOSED	{
}
;

_Q_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_E_Star:
    {
}

    | _Q_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_E_Star _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C	{
}
;

extraPropertySet:
    IT_EXTRA _Qpredicate_E_Plus	{
}
;

_Qpredicate_E_Plus:
    predicate	{
}

    | _Qpredicate_E_Plus predicate	{
}
;

tripleExpression:
    oneOfTripleExpr	{
}
;

oneOfTripleExpr:
    groupTripleExpr	{
}

    | multiElementOneOf	{
}
;

multiElementOneOf:
    groupTripleExpr _Q_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C_E_Plus	{
}
;

_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C:
    GT_PIPE groupTripleExpr	{
}
;

_Q_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C_E_Plus:
    _O_QGT_PIPE_E_S_QgroupTripleExpr_E_C	{
}

    | _Q_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C_E_Plus _O_QGT_PIPE_E_S_QgroupTripleExpr_E_C	{
}
;

groupTripleExpr:
    singleElementGroup	{
}

    | multiElementGroup	{
}
;

singleElementGroup:
    unaryTripleExpr _QGT_SEMI_E_Opt	{
}
;

_QGT_SEMI_E_Opt:
    {
}

    | GT_SEMI	{
}
;

multiElementGroup:
    unaryTripleExpr _Q_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C_E_Plus _QGT_SEMI_E_Opt	{
}
;

_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C:
    GT_SEMI unaryTripleExpr	{
}
;

_Q_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C_E_Plus:
    _O_QGT_SEMI_E_S_QunaryTripleExpr_E_C	{
}

    | _Q_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C_E_Plus _O_QGT_SEMI_E_S_QunaryTripleExpr_E_C	{
}
;

unaryTripleExpr:
    _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C	{
}

    | include	{
}
;

_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C:
    GT_DOLLAR tripleExprLabel	{
}
;

_Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt:
    {
}

    | _O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C	{
}
;

_O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C:
    tripleConstraint	{
}

    | bracketedTripleExpr	{
}
;

bracketedTripleExpr:
    GT_LPAREN tripleExpression GT_RPAREN _Qcardinality_E_Opt _Qannotation_E_Star semanticActions	{
}
;

_Qcardinality_E_Opt:
    {
}

    | cardinality	{
}
;

tripleConstraint:
    _QsenseFlags_E_Opt predicate inlineShapeExpression _Qcardinality_E_Opt _Qannotation_E_Star semanticActions	{
}
;

_QsenseFlags_E_Opt:
    {
}

    | senseFlags	{
}
;

cardinality:
    GT_TIMES	{
}

    | GT_PLUS	{
}

    | GT_OPT	{
}

    | REPEAT_RANGE	{
}
;

senseFlags:
    GT_CARROT	{
}
;

valueSet:
    GT_LBRACKET _QvalueSetValue_E_Star GT_RBRACKET	{
}
;

_QvalueSetValue_E_Star:
    {
}

    | _QvalueSetValue_E_Star valueSetValue	{
}
;

valueSetValue:
    iriRange	{
}

    | literalRange	{
}

    | languageRange	{
}

    | GT_DOT _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C	{
}
;

_QiriExclusion_E_Plus:
    iriExclusion	{
}

    | _QiriExclusion_E_Plus iriExclusion	{
}
;

_QliteralExclusion_E_Plus:
    literalExclusion	{
}

    | _QliteralExclusion_E_Plus literalExclusion	{
}
;

_QlanguageExclusion_E_Plus:
    languageExclusion	{
}

    | _QlanguageExclusion_E_Plus languageExclusion	{
}
;

_O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C:
    _QiriExclusion_E_Plus	{
}

    | _QliteralExclusion_E_Plus	{
}

    | _QlanguageExclusion_E_Plus	{
}
;

iriRange:
    iri _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt	{
}
;

_QiriExclusion_E_Star:
    {
}

    | _QiriExclusion_E_Star iriExclusion	{
}
;

_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C:
    GT_KINDA _QiriExclusion_E_Star	{
}
;

_Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt:
    {
}

    | _O_QGT_KINDA_E_S_QiriExclusion_E_Star_C	{
}
;

iriExclusion:
    GT_MINUS iri _QGT_KINDA_E_Opt	{
}
;

_QGT_KINDA_E_Opt:
    {
}

    | GT_KINDA	{
}
;

literalRange:
    literal _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt	{
}
;

_QliteralExclusion_E_Star:
    {
}

    | _QliteralExclusion_E_Star literalExclusion	{
}
;

_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C:
    GT_KINDA _QliteralExclusion_E_Star	{
}
;

_Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt:
    {
}

    | _O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C	{
}
;

literalExclusion:
    GT_MINUS literal _QGT_KINDA_E_Opt	{
}
;

languageRange:
    LANGTAG _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt	{
}

    | GT_AT GT_KINDA _QlanguageExclusion_E_Star	{
}
;

_QlanguageExclusion_E_Star:
    {
}

    | _QlanguageExclusion_E_Star languageExclusion	{
}
;

_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C:
    GT_KINDA _QlanguageExclusion_E_Star	{
}
;

_Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt:
    {
}

    | _O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C	{
}
;

languageExclusion:
    GT_MINUS LANGTAG _QGT_KINDA_E_Opt	{
}
;

include:
    GT_AMP tripleExprLabel	{
}
;

annotation:
    GT_DIVIDE_DIVIDE predicate _O_Qiri_E_Or_Qliteral_E_C	{
}
;

_O_Qiri_E_Or_Qliteral_E_C:
    iri	{
}

    | literal	{
}
;

semanticActions:
    _QcodeDecl_E_Star	{
}
;

_QcodeDecl_E_Star:
    {
}

    | _QcodeDecl_E_Star codeDecl	{
}
;

codeDecl:
    GT_MODULO iri _O_QCODE_E_Or_QGT_MODULO_E_C	{
}
;

_O_QCODE_E_Or_QGT_MODULO_E_C:
    CODE	{
}

    | GT_MODULO	{
}
;

literal:
    rdfLiteral	{
}

    | numericLiteral	{
}

    | booleanLiteral	{
}
;

predicate:
    iri	{
}

    | RDF_TYPE	{
}
;

datatype:
    iri	{
}
;

shapeExprLabel:
    iri	{
}

    | blankNode	{
}
;

tripleExprLabel:
    iri	{
}

    | blankNode	{
}
;

rawNumeric:
    INTEGER	{
}

    | DECIMAL	{
}

    | DOUBLE	{
}
;

numericLiteral:
    rawNumeric	{
}
;

rdfLiteral:
    langString	{
}

    | string _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt	{
}
;

_O_QGT_DTYPE_E_S_Qdatatype_E_C:
    GT_DTYPE datatype	{
}
;

_Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt:
    {
}

    | _O_QGT_DTYPE_E_S_Qdatatype_E_C	{
}
;

booleanLiteral:
    IT_true	{
}

    | IT_false	{
}
;

string:
    STRING_LITERAL1	{
}

    | STRING_LITERAL_LONG1	{
}

    | STRING_LITERAL2	{
}

    | STRING_LITERAL_LONG2	{
}
;

langString:
    LANG_STRING_LITERAL1	{
}

    | LANG_STRING_LITERAL_LONG1	{
}

    | LANG_STRING_LITERAL2	{
}

    | LANG_STRING_LITERAL_LONG2	{
}
;

iri:
    IRIREF	{
      std::cout << "IRIREF: " << $1.val << std::endl;
}

    | prefixedName	{
}
;

prefixedName:
    PNAME_LN	{
}

    | PNAME_NS	{
}
;

blankNode:
    BLANK_NODE_LABEL	{
}
;











































































 /*** END shexParseToXml - Change the grammar rules above ***/

%% /*** Additional Code ***/

void shexParseToXml::shexParseToXmlParser::error(const shexParseToXmlParser::location_type& l,
			    const std::string& m)
{
    driver.error(l, m);
}

/* START yacker-specific test harness */

namespace yacker {
std::ostream* _Trace = NULL;
} // namespace yacker

/* END yacker-specific test harness */

/* START Driver (@@ stand-alone would allow it to be shared with other parsers */

namespace shexParseToXml {

Driver::Driver(class shexParseToXmlContext& _context)
    : trace_scanning(false),
      trace_parsing(false),
      context(_context)
{
}

bool Driver::parse_stream(std::istream& in, const std::string& sname)
{
    streamname = sname;

    shexParseToXmlScanner scanner(&in);
    scanner.set_debug(trace_scanning);
    this->lexer = &scanner;

    shexParseToXmlParser parser(*this);
    parser.set_debug_level(trace_parsing);
    return (parser.parse());
}

bool Driver::parse_file(const std::string &filename)
{
    std::ifstream in(filename.c_str());
    return parse_stream(in, filename);
}

bool Driver::parse_string(const std::string &input, const std::string& sname)
{
    std::istringstream iss(input);
    return parse_stream(iss, sname);
}

void Driver::error(const class location& l,
		   const std::string& m)
{
    std::cerr << l << ": " << m << std::endl;
}

void Driver::error(const std::string& m)
{
    std::cerr << m << std::endl;
}

} // namespace shexParseToXml

/* END Driver */
