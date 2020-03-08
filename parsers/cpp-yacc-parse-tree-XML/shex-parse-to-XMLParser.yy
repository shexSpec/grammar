/* $Id: Langname_Parser.yy,v 1.11 2014-04-07 13:02:07 eric Exp spec_2_1Parser.yy 19 2007-08-19 20:36:24Z tb $ -*- mode: c++ -*- */
/** \file shex-parse-to-XMLParser.yy Contains the Bison parser source */

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

/* use newer C++ skeleton file */
%skeleton "lalr1.cc"

/* namespace to enclose parser in */
%define api.prefix {shexParseToXml}

/* set the parser's class identifier */
%define api.parser.class {shexParseToXmlParser}

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

extern std::ostream* _Trace;

/* START ClassBlock */
namespace yacker {
class _Base {
private:
    const char* productionName;
public:
    _Base (const char* productionName) : productionName(productionName) { }
    const char* getProductionName() { return productionName; };
    virtual std::string toStr(std::ofstream* out = NULL) = 0;
    virtual std::string toXml(size_t depth, std::ofstream* out = NULL) = 0;
    virtual size_t size () { return 1; }
    virtual _Base* getElement (size_t) { return this; }
    virtual ~_Base() { }
};
class _Production : public _Base {
protected:
    _Production (const char* productionName) : _Base(productionName) { }
    void makeArray(_Base** target[], ...);
    void trace(unsigned creditToSelf = 0);
public:
    virtual bool transparent() { return false; } // in general, productions are not transparent.
    virtual _Base* operator [] (size_t i) = 0;
    virtual _Base* getElement (size_t i) = 0;
    virtual std::string toStr(std::ofstream* out = NULL);
    virtual std::string toXml(size_t depth, std::ofstream* out = NULL);
};

class _GenProduction : public _Production {
protected:
    _GenProduction (const char* productionName) : _Production(productionName) { }
    ~_GenProduction () {}
};

extern bool TransparentGroupProductions;
extern bool TransparentMultiplicityProductions;
template <typename T> class _ProductionVector : public _GenProduction {
    std::vector<T> data;
public:
    _ProductionVector (const char* productionName) : _GenProduction(productionName) {
	trace();
    }
    _ProductionVector (const char* productionName, T v) : _GenProduction(productionName) {
	data.push_back(v);
	trace(size()-1);
    }
    ~_ProductionVector () {
	for (size_t i = 0; i < size(); i++)
	    delete data[i];
    }
    virtual bool transparent() { return TransparentMultiplicityProductions; }

    void push_back(T v) {
	data.push_back(v);
	trace(size()-1);
    }
    size_t size() { return data.size(); }
    virtual T operator [] (size_t i) { return data[i]; }
    _Base* getElement (size_t i) { return data[i]; }
};
class _GroupProduction : public _GenProduction {
protected:
    _GroupProduction (const char* productionName) : _GenProduction (productionName) { }
    virtual bool transparent() { return TransparentGroupProductions; }
};
class _Token : public _Base {
private:
    const char* matched;

protected:
    _Token (const char* productionName, const char* matched) : _Base(productionName), matched(copyOf(matched)) { }
public:
    ~_Token () { delete[] matched; }

protected:
    void trace();

private:
    static char* copyOf (const char* copyMe) {
	char* ret = new char[::strlen(copyMe)+1];
	strcpy(ret, copyMe);
	return ret;
    }
public:
    virtual std::string toStr(std::ofstream* out = NULL);
    virtual std::string toXml(size_t depth, std::ofstream* out = NULL);
};
class _Terminal : public _Base {
protected:
    const char* terminal;
    _Terminal (const char* productionName, const char* p) : _Base(productionName) {
	terminal = new char[::strlen(p) + 1];
	strcpy((char*)terminal, p); // @@ should initialize as member
    }
    ~_Terminal () {
	delete[] terminal;
    }
    void trace();
public:
    virtual std::string toStr(std::ofstream* out = NULL);
    virtual std::string toXml(size_t depth, std::ofstream* out = NULL);
};

} // namespace yacker

namespace shexParseToXml {
class shexDoc;
class _O_QnotStartAction_E_Or_QstartActions_E_C;
class _O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C;
class _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt;
class directive;
class baseDecl;
class prefixDecl;
class importDecl;
class notStartAction;
class start;
class startActions;
class statement;
class shapeExprDecl;
class _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C;
class shapeExpression;
class inlineShapeExpression;
class shapeOr;
class _O_QIT_OR_E_S_QshapeAnd_E_C;
class inlineShapeOr;
class _O_QIT_OR_E_S_QinlineShapeAnd_E_C;
class shapeAnd;
class _O_QIT_AND_E_S_QshapeNot_E_C;
class inlineShapeAnd;
class _O_QIT_AND_E_S_QinlineShapeNot_E_C;
class shapeNot;
class _QIT_NOT_E_Opt;
class inlineShapeNot;
class shapeAtom;
class _QshapeOrRef_E_Opt;
class _QnonLitNodeConstraint_E_Opt;
class inlineShapeAtom;
class _QinlineShapeOrRef_E_Opt;
class shapeOrRef;
class inlineShapeOrRef;
class shapeRef;
class litNodeConstraint;
class nonLitNodeConstraint;
class nonLiteralKind;
class xsFacet;
class stringFacet;
class stringLength;
class numericFacet;
class numericRange;
class numericLength;
class shapeDefinition;
class inlineShapeDefinition;
class _QtripleExpression_E_Opt;
class shapeQualifiers;
class _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C;
class extraPropertySet;
class tripleExpression;
class oneOfTripleExpr;
class multiElementOneOf;
class _O_QGT_PIPE_E_S_QgroupTripleExpr_E_C;
class groupTripleExpr;
class singleElementGroup;
class _QGT_SEMI_E_Opt;
class multiElementGroup;
class _O_QGT_SEMI_E_S_QunaryTripleExpr_E_C;
class unaryTripleExpr;
class _O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C;
class _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt;
class _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C;
class bracketedTripleExpr;
class _Qcardinality_E_Opt;
class tripleConstraint;
class _QsenseFlags_E_Opt;
class cardinality;
class senseFlags;
class valueSet;
class valueSetValue;
class _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C;
class iriRange;
class _O_QGT_KINDA_E_S_QiriExclusion_E_Star_C;
class _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt;
class iriExclusion;
class _QGT_KINDA_E_Opt;
class literalRange;
class _O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C;
class _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt;
class literalExclusion;
class languageRange;
class _O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C;
class _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt;
class languageExclusion;
class include;
class annotation;
class _O_Qiri_E_Or_Qliteral_E_C;
class semanticActions;
class codeDecl;
class _O_QCODE_E_Or_QGT_MODULO_E_C;
class literal;
class predicate;
class datatype;
class shapeExprLabel;
class tripleExprLabel;
class rawNumeric;
class numericLiteral;
class rdfLiteral;
class _O_QGT_DTYPE_E_S_Qdatatype_E_C;
class _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt;
class booleanLiteral;
class string;
class langString;
class iri;
class prefixedName;
class blankNode;

class IT_BASE;
class IT_PREFIX;
class IT_IMPORT;
class IT_START;
class GT_EQUAL;
class IT_EXTERNAL;
class IT_OR;
class IT_AND;
class IT_NOT;
class GT_LPAREN;
class GT_RPAREN;
class GT_DOT;
class GT_AT;
class IT_LITERAL;
class IT_IRI;
class IT_BNODE;
class IT_NONLITERAL;
class IT_LENGTH;
class IT_MINLENGTH;
class IT_MAXLENGTH;
class IT_MININCLUSIVE;
class IT_MINEXCLUSIVE;
class IT_MAXINCLUSIVE;
class IT_MAXEXCLUSIVE;
class IT_TOTALDIGITS;
class IT_FRACTIONDIGITS;
class GT_LCURLEY;
class GT_RCURLEY;
class IT_CLOSED;
class IT_EXTRA;
class GT_PIPE;
class GT_SEMI;
class GT_DOLLAR;
class GT_TIMES;
class GT_PLUS;
class GT_OPT;
class GT_CARROT;
class GT_LBRACKET;
class GT_RBRACKET;
class GT_KINDA;
class GT_MINUS;
class GT_AMP;
class GT_DIVIDE_DIVIDE;
class GT_MODULO;
class GT_DTYPE;
class IT_true;
class IT_false;
class CODE;
class REPEAT_RANGE;
class RDF_TYPE;
class IRIREF;
class PNAME_NS;
class PNAME_LN;
class ATPNAME_NS;
class ATPNAME_LN;
class REGEXP;
class BLANK_NODE_LABEL;
class LANGTAG;
class INTEGER;
class DECIMAL;
class DOUBLE;
class STRING_LITERAL1;
class STRING_LITERAL2;
class STRING_LITERAL_LONG1;
class STRING_LITERAL_LONG2;
class LANG_STRING_LITERAL1;
class LANG_STRING_LITERAL2;
class LANG_STRING_LITERAL_LONG1;
class LANG_STRING_LITERAL_LONG2;


class shexDoc : public yacker::_Production {
private:
    _Base** _members[2];
    yacker::_ProductionVector<directive*>* m_directives;
    _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt* m__Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt;
    virtual const char* getProductionName () { return "shexDoc"; }
public:
    shexDoc (yacker::_ProductionVector<directive*>* p_directives, _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt* p__Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt) : yacker::_Production("shexDoc") {
	m_directives = p_directives;
	m__Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt = p__Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt;
	makeArray(_members, &m_directives, &m__Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt);
	trace();
    }
    ~shexDoc();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QnotStartAction_E_Or_QstartActions_E_C : public yacker::_GroupProduction {
public:
    _O_QnotStartAction_E_Or_QstartActions_E_C () : yacker::_GroupProduction ("_O_QnotStartAction_E_Or_QstartActions_E_C") {}
};
class _O_QnotStartAction_E_Or_QstartActions_E_C_rule0 : public _O_QnotStartAction_E_Or_QstartActions_E_C {
private:
    _Base** _members[1];
    notStartAction* m_notStartAction;
public:
    _O_QnotStartAction_E_Or_QstartActions_E_C_rule0 (notStartAction* p_notStartAction) {
	m_notStartAction = p_notStartAction;
	makeArray(_members, &m_notStartAction);
	trace();
    }
    ~_O_QnotStartAction_E_Or_QstartActions_E_C_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QnotStartAction_E_Or_QstartActions_E_C_rule1 : public _O_QnotStartAction_E_Or_QstartActions_E_C {
private:
    _Base** _members[1];
    startActions* m_startActions;
public:
    _O_QnotStartAction_E_Or_QstartActions_E_C_rule1 (startActions* p_startActions) {
	m_startActions = p_startActions;
	makeArray(_members, &m_startActions);
	trace();
    }
    ~_O_QnotStartAction_E_Or_QstartActions_E_C_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C : public yacker::_Production {
private:
    _Base** _members[2];
    _O_QnotStartAction_E_Or_QstartActions_E_C* m__O_QnotStartAction_E_Or_QstartActions_E_C;
    yacker::_ProductionVector<statement*>* m_statements;
    virtual const char* getProductionName () { return "_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C"; }
public:
    _O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C (_O_QnotStartAction_E_Or_QstartActions_E_C* p__O_QnotStartAction_E_Or_QstartActions_E_C, yacker::_ProductionVector<statement*>* p_statements) : yacker::_Production("_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C") {
	m__O_QnotStartAction_E_Or_QstartActions_E_C = p__O_QnotStartAction_E_Or_QstartActions_E_C;
	m_statements = p_statements;
	makeArray(_members, &m__O_QnotStartAction_E_Or_QstartActions_E_C, &m_statements);
	trace();
    }
    ~_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt : public yacker::_GroupProduction {
public:
    _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt () : yacker::_GroupProduction ("_Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt") {}
};
class _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt_rule0 : public _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt {
public:
    _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt_rule0 () {
	trace();
    }
    ~_Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt_rule1 : public _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt {
private:
    _Base** _members[1];
    _O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C* m__O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C;
public:
    _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt_rule1 (_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C* p__O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C) {
	m__O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C = p__O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C;
	makeArray(_members, &m__O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C);
	trace();
    }
    ~_Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class directive : public yacker::_GroupProduction {
public:
    directive () : yacker::_GroupProduction ("directive") {}
};
class directive_rule0 : public directive {
private:
    _Base** _members[1];
    baseDecl* m_baseDecl;
public:
    directive_rule0 (baseDecl* p_baseDecl) {
	m_baseDecl = p_baseDecl;
	makeArray(_members, &m_baseDecl);
	trace();
    }
    ~directive_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class directive_rule1 : public directive {
private:
    _Base** _members[1];
    prefixDecl* m_prefixDecl;
public:
    directive_rule1 (prefixDecl* p_prefixDecl) {
	m_prefixDecl = p_prefixDecl;
	makeArray(_members, &m_prefixDecl);
	trace();
    }
    ~directive_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class directive_rule2 : public directive {
private:
    _Base** _members[1];
    importDecl* m_importDecl;
public:
    directive_rule2 (importDecl* p_importDecl) {
	m_importDecl = p_importDecl;
	makeArray(_members, &m_importDecl);
	trace();
    }
    ~directive_rule2();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class baseDecl : public yacker::_Production {
private:
    _Base** _members[2];
    IT_BASE* m_IT_BASE;
    IRIREF* m_IRIREF;
    virtual const char* getProductionName () { return "baseDecl"; }
public:
    baseDecl (IT_BASE* p_IT_BASE, IRIREF* p_IRIREF) : yacker::_Production("baseDecl") {
	m_IT_BASE = p_IT_BASE;
	m_IRIREF = p_IRIREF;
	makeArray(_members, &m_IT_BASE, &m_IRIREF);
	trace();
    }
    ~baseDecl();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class prefixDecl : public yacker::_Production {
private:
    _Base** _members[3];
    IT_PREFIX* m_IT_PREFIX;
    PNAME_NS* m_PNAME_NS;
    IRIREF* m_IRIREF;
    virtual const char* getProductionName () { return "prefixDecl"; }
public:
    prefixDecl (IT_PREFIX* p_IT_PREFIX, PNAME_NS* p_PNAME_NS, IRIREF* p_IRIREF) : yacker::_Production("prefixDecl") {
	m_IT_PREFIX = p_IT_PREFIX;
	m_PNAME_NS = p_PNAME_NS;
	m_IRIREF = p_IRIREF;
	makeArray(_members, &m_IT_PREFIX, &m_PNAME_NS, &m_IRIREF);
	trace();
    }
    ~prefixDecl();
    virtual size_t size () { return 3; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class importDecl : public yacker::_Production {
private:
    _Base** _members[2];
    IT_IMPORT* m_IT_IMPORT;
    IRIREF* m_IRIREF;
    virtual const char* getProductionName () { return "importDecl"; }
public:
    importDecl (IT_IMPORT* p_IT_IMPORT, IRIREF* p_IRIREF) : yacker::_Production("importDecl") {
	m_IT_IMPORT = p_IT_IMPORT;
	m_IRIREF = p_IRIREF;
	makeArray(_members, &m_IT_IMPORT, &m_IRIREF);
	trace();
    }
    ~importDecl();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class notStartAction : public yacker::_GroupProduction {
public:
    notStartAction () : yacker::_GroupProduction ("notStartAction") {}
};
class notStartAction_rule0 : public notStartAction {
private:
    _Base** _members[1];
    start* m_start;
public:
    notStartAction_rule0 (start* p_start) {
	m_start = p_start;
	makeArray(_members, &m_start);
	trace();
    }
    ~notStartAction_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class notStartAction_rule1 : public notStartAction {
private:
    _Base** _members[1];
    shapeExprDecl* m_shapeExprDecl;
public:
    notStartAction_rule1 (shapeExprDecl* p_shapeExprDecl) {
	m_shapeExprDecl = p_shapeExprDecl;
	makeArray(_members, &m_shapeExprDecl);
	trace();
    }
    ~notStartAction_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class start : public yacker::_Production {
private:
    _Base** _members[3];
    IT_START* m_IT_START;
    GT_EQUAL* m_GT_EQUAL;
    inlineShapeExpression* m_inlineShapeExpression;
    virtual const char* getProductionName () { return "start"; }
public:
    start (IT_START* p_IT_START, GT_EQUAL* p_GT_EQUAL, inlineShapeExpression* p_inlineShapeExpression) : yacker::_Production("start") {
	m_IT_START = p_IT_START;
	m_GT_EQUAL = p_GT_EQUAL;
	m_inlineShapeExpression = p_inlineShapeExpression;
	makeArray(_members, &m_IT_START, &m_GT_EQUAL, &m_inlineShapeExpression);
	trace();
    }
    ~start();
    virtual size_t size () { return 3; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class startActions : public yacker::_Production {
private:
    _Base** _members[1];
    yacker::_ProductionVector<codeDecl*>* m_codeDecls;
    virtual const char* getProductionName () { return "startActions"; }
public:
    startActions (yacker::_ProductionVector<codeDecl*>* p_codeDecls) : yacker::_Production("startActions") {
	m_codeDecls = p_codeDecls;
	makeArray(_members, &m_codeDecls);
	trace();
    }
    ~startActions();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class statement : public yacker::_GroupProduction {
public:
    statement () : yacker::_GroupProduction ("statement") {}
};
class statement_rule0 : public statement {
private:
    _Base** _members[1];
    directive* m_directive;
public:
    statement_rule0 (directive* p_directive) {
	m_directive = p_directive;
	makeArray(_members, &m_directive);
	trace();
    }
    ~statement_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class statement_rule1 : public statement {
private:
    _Base** _members[1];
    notStartAction* m_notStartAction;
public:
    statement_rule1 (notStartAction* p_notStartAction) {
	m_notStartAction = p_notStartAction;
	makeArray(_members, &m_notStartAction);
	trace();
    }
    ~statement_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeExprDecl : public yacker::_Production {
private:
    _Base** _members[2];
    shapeExprLabel* m_shapeExprLabel;
    _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C* m__O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C;
    virtual const char* getProductionName () { return "shapeExprDecl"; }
public:
    shapeExprDecl (shapeExprLabel* p_shapeExprLabel, _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C* p__O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C) : yacker::_Production("shapeExprDecl") {
	m_shapeExprLabel = p_shapeExprLabel;
	m__O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C = p__O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C;
	makeArray(_members, &m_shapeExprLabel, &m__O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C);
	trace();
    }
    ~shapeExprDecl();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C : public yacker::_GroupProduction {
public:
    _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C () : yacker::_GroupProduction ("_O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C") {}
};
class _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C_rule0 : public _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C {
private:
    _Base** _members[1];
    shapeExpression* m_shapeExpression;
public:
    _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C_rule0 (shapeExpression* p_shapeExpression) {
	m_shapeExpression = p_shapeExpression;
	makeArray(_members, &m_shapeExpression);
	trace();
    }
    ~_O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C_rule1 : public _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C {
private:
    _Base** _members[1];
    IT_EXTERNAL* m_IT_EXTERNAL;
public:
    _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C_rule1 (IT_EXTERNAL* p_IT_EXTERNAL) {
	m_IT_EXTERNAL = p_IT_EXTERNAL;
	makeArray(_members, &m_IT_EXTERNAL);
	trace();
    }
    ~_O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeExpression : public yacker::_Production {
private:
    _Base** _members[1];
    shapeOr* m_shapeOr;
    virtual const char* getProductionName () { return "shapeExpression"; }
public:
    shapeExpression (shapeOr* p_shapeOr) : yacker::_Production("shapeExpression") {
	m_shapeOr = p_shapeOr;
	makeArray(_members, &m_shapeOr);
	trace();
    }
    ~shapeExpression();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class inlineShapeExpression : public yacker::_Production {
private:
    _Base** _members[1];
    inlineShapeOr* m_inlineShapeOr;
    virtual const char* getProductionName () { return "inlineShapeExpression"; }
public:
    inlineShapeExpression (inlineShapeOr* p_inlineShapeOr) : yacker::_Production("inlineShapeExpression") {
	m_inlineShapeOr = p_inlineShapeOr;
	makeArray(_members, &m_inlineShapeOr);
	trace();
    }
    ~inlineShapeExpression();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeOr : public yacker::_Production {
private:
    _Base** _members[2];
    shapeAnd* m_shapeAnd;
    yacker::_ProductionVector<_O_QIT_OR_E_S_QshapeAnd_E_C*>* m__O_QIT_OR_E_S_QshapeAnd_E_Cs;
    virtual const char* getProductionName () { return "shapeOr"; }
public:
    shapeOr (shapeAnd* p_shapeAnd, yacker::_ProductionVector<_O_QIT_OR_E_S_QshapeAnd_E_C*>* p__O_QIT_OR_E_S_QshapeAnd_E_Cs) : yacker::_Production("shapeOr") {
	m_shapeAnd = p_shapeAnd;
	m__O_QIT_OR_E_S_QshapeAnd_E_Cs = p__O_QIT_OR_E_S_QshapeAnd_E_Cs;
	makeArray(_members, &m_shapeAnd, &m__O_QIT_OR_E_S_QshapeAnd_E_Cs);
	trace();
    }
    ~shapeOr();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QIT_OR_E_S_QshapeAnd_E_C : public yacker::_Production {
private:
    _Base** _members[2];
    IT_OR* m_IT_OR;
    shapeAnd* m_shapeAnd;
    virtual const char* getProductionName () { return "_O_QIT_OR_E_S_QshapeAnd_E_C"; }
public:
    _O_QIT_OR_E_S_QshapeAnd_E_C (IT_OR* p_IT_OR, shapeAnd* p_shapeAnd) : yacker::_Production("_O_QIT_OR_E_S_QshapeAnd_E_C") {
	m_IT_OR = p_IT_OR;
	m_shapeAnd = p_shapeAnd;
	makeArray(_members, &m_IT_OR, &m_shapeAnd);
	trace();
    }
    ~_O_QIT_OR_E_S_QshapeAnd_E_C();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class inlineShapeOr : public yacker::_Production {
private:
    _Base** _members[2];
    inlineShapeAnd* m_inlineShapeAnd;
    yacker::_ProductionVector<_O_QIT_OR_E_S_QinlineShapeAnd_E_C*>* m__O_QIT_OR_E_S_QinlineShapeAnd_E_Cs;
    virtual const char* getProductionName () { return "inlineShapeOr"; }
public:
    inlineShapeOr (inlineShapeAnd* p_inlineShapeAnd, yacker::_ProductionVector<_O_QIT_OR_E_S_QinlineShapeAnd_E_C*>* p__O_QIT_OR_E_S_QinlineShapeAnd_E_Cs) : yacker::_Production("inlineShapeOr") {
	m_inlineShapeAnd = p_inlineShapeAnd;
	m__O_QIT_OR_E_S_QinlineShapeAnd_E_Cs = p__O_QIT_OR_E_S_QinlineShapeAnd_E_Cs;
	makeArray(_members, &m_inlineShapeAnd, &m__O_QIT_OR_E_S_QinlineShapeAnd_E_Cs);
	trace();
    }
    ~inlineShapeOr();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QIT_OR_E_S_QinlineShapeAnd_E_C : public yacker::_Production {
private:
    _Base** _members[2];
    IT_OR* m_IT_OR;
    inlineShapeAnd* m_inlineShapeAnd;
    virtual const char* getProductionName () { return "_O_QIT_OR_E_S_QinlineShapeAnd_E_C"; }
public:
    _O_QIT_OR_E_S_QinlineShapeAnd_E_C (IT_OR* p_IT_OR, inlineShapeAnd* p_inlineShapeAnd) : yacker::_Production("_O_QIT_OR_E_S_QinlineShapeAnd_E_C") {
	m_IT_OR = p_IT_OR;
	m_inlineShapeAnd = p_inlineShapeAnd;
	makeArray(_members, &m_IT_OR, &m_inlineShapeAnd);
	trace();
    }
    ~_O_QIT_OR_E_S_QinlineShapeAnd_E_C();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeAnd : public yacker::_Production {
private:
    _Base** _members[2];
    shapeNot* m_shapeNot;
    yacker::_ProductionVector<_O_QIT_AND_E_S_QshapeNot_E_C*>* m__O_QIT_AND_E_S_QshapeNot_E_Cs;
    virtual const char* getProductionName () { return "shapeAnd"; }
public:
    shapeAnd (shapeNot* p_shapeNot, yacker::_ProductionVector<_O_QIT_AND_E_S_QshapeNot_E_C*>* p__O_QIT_AND_E_S_QshapeNot_E_Cs) : yacker::_Production("shapeAnd") {
	m_shapeNot = p_shapeNot;
	m__O_QIT_AND_E_S_QshapeNot_E_Cs = p__O_QIT_AND_E_S_QshapeNot_E_Cs;
	makeArray(_members, &m_shapeNot, &m__O_QIT_AND_E_S_QshapeNot_E_Cs);
	trace();
    }
    ~shapeAnd();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QIT_AND_E_S_QshapeNot_E_C : public yacker::_Production {
private:
    _Base** _members[2];
    IT_AND* m_IT_AND;
    shapeNot* m_shapeNot;
    virtual const char* getProductionName () { return "_O_QIT_AND_E_S_QshapeNot_E_C"; }
public:
    _O_QIT_AND_E_S_QshapeNot_E_C (IT_AND* p_IT_AND, shapeNot* p_shapeNot) : yacker::_Production("_O_QIT_AND_E_S_QshapeNot_E_C") {
	m_IT_AND = p_IT_AND;
	m_shapeNot = p_shapeNot;
	makeArray(_members, &m_IT_AND, &m_shapeNot);
	trace();
    }
    ~_O_QIT_AND_E_S_QshapeNot_E_C();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class inlineShapeAnd : public yacker::_Production {
private:
    _Base** _members[2];
    inlineShapeNot* m_inlineShapeNot;
    yacker::_ProductionVector<_O_QIT_AND_E_S_QinlineShapeNot_E_C*>* m__O_QIT_AND_E_S_QinlineShapeNot_E_Cs;
    virtual const char* getProductionName () { return "inlineShapeAnd"; }
public:
    inlineShapeAnd (inlineShapeNot* p_inlineShapeNot, yacker::_ProductionVector<_O_QIT_AND_E_S_QinlineShapeNot_E_C*>* p__O_QIT_AND_E_S_QinlineShapeNot_E_Cs) : yacker::_Production("inlineShapeAnd") {
	m_inlineShapeNot = p_inlineShapeNot;
	m__O_QIT_AND_E_S_QinlineShapeNot_E_Cs = p__O_QIT_AND_E_S_QinlineShapeNot_E_Cs;
	makeArray(_members, &m_inlineShapeNot, &m__O_QIT_AND_E_S_QinlineShapeNot_E_Cs);
	trace();
    }
    ~inlineShapeAnd();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QIT_AND_E_S_QinlineShapeNot_E_C : public yacker::_Production {
private:
    _Base** _members[2];
    IT_AND* m_IT_AND;
    inlineShapeNot* m_inlineShapeNot;
    virtual const char* getProductionName () { return "_O_QIT_AND_E_S_QinlineShapeNot_E_C"; }
public:
    _O_QIT_AND_E_S_QinlineShapeNot_E_C (IT_AND* p_IT_AND, inlineShapeNot* p_inlineShapeNot) : yacker::_Production("_O_QIT_AND_E_S_QinlineShapeNot_E_C") {
	m_IT_AND = p_IT_AND;
	m_inlineShapeNot = p_inlineShapeNot;
	makeArray(_members, &m_IT_AND, &m_inlineShapeNot);
	trace();
    }
    ~_O_QIT_AND_E_S_QinlineShapeNot_E_C();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeNot : public yacker::_Production {
private:
    _Base** _members[2];
    _QIT_NOT_E_Opt* m__QIT_NOT_E_Opt;
    shapeAtom* m_shapeAtom;
    virtual const char* getProductionName () { return "shapeNot"; }
public:
    shapeNot (_QIT_NOT_E_Opt* p__QIT_NOT_E_Opt, shapeAtom* p_shapeAtom) : yacker::_Production("shapeNot") {
	m__QIT_NOT_E_Opt = p__QIT_NOT_E_Opt;
	m_shapeAtom = p_shapeAtom;
	makeArray(_members, &m__QIT_NOT_E_Opt, &m_shapeAtom);
	trace();
    }
    ~shapeNot();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _QIT_NOT_E_Opt : public yacker::_GroupProduction {
public:
    _QIT_NOT_E_Opt () : yacker::_GroupProduction ("_QIT_NOT_E_Opt") {}
};
class _QIT_NOT_E_Opt_rule0 : public _QIT_NOT_E_Opt {
public:
    _QIT_NOT_E_Opt_rule0 () {
	trace();
    }
    ~_QIT_NOT_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _QIT_NOT_E_Opt_rule1 : public _QIT_NOT_E_Opt {
private:
    _Base** _members[1];
    IT_NOT* m_IT_NOT;
public:
    _QIT_NOT_E_Opt_rule1 (IT_NOT* p_IT_NOT) {
	m_IT_NOT = p_IT_NOT;
	makeArray(_members, &m_IT_NOT);
	trace();
    }
    ~_QIT_NOT_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class inlineShapeNot : public yacker::_Production {
private:
    _Base** _members[2];
    _QIT_NOT_E_Opt* m__QIT_NOT_E_Opt;
    inlineShapeAtom* m_inlineShapeAtom;
    virtual const char* getProductionName () { return "inlineShapeNot"; }
public:
    inlineShapeNot (_QIT_NOT_E_Opt* p__QIT_NOT_E_Opt, inlineShapeAtom* p_inlineShapeAtom) : yacker::_Production("inlineShapeNot") {
	m__QIT_NOT_E_Opt = p__QIT_NOT_E_Opt;
	m_inlineShapeAtom = p_inlineShapeAtom;
	makeArray(_members, &m__QIT_NOT_E_Opt, &m_inlineShapeAtom);
	trace();
    }
    ~inlineShapeNot();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeAtom : public yacker::_GroupProduction {
public:
    shapeAtom () : yacker::_GroupProduction ("shapeAtom") {}
};
class shapeAtom_rule0 : public shapeAtom {
private:
    _Base** _members[2];
    nonLitNodeConstraint* m_nonLitNodeConstraint;
    _QshapeOrRef_E_Opt* m__QshapeOrRef_E_Opt;
public:
    shapeAtom_rule0 (nonLitNodeConstraint* p_nonLitNodeConstraint, _QshapeOrRef_E_Opt* p__QshapeOrRef_E_Opt) {
	m_nonLitNodeConstraint = p_nonLitNodeConstraint;
	m__QshapeOrRef_E_Opt = p__QshapeOrRef_E_Opt;
	makeArray(_members, &m_nonLitNodeConstraint, &m__QshapeOrRef_E_Opt);
	trace();
    }
    ~shapeAtom_rule0();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeAtom_rule1 : public shapeAtom {
private:
    _Base** _members[1];
    litNodeConstraint* m_litNodeConstraint;
public:
    shapeAtom_rule1 (litNodeConstraint* p_litNodeConstraint) {
	m_litNodeConstraint = p_litNodeConstraint;
	makeArray(_members, &m_litNodeConstraint);
	trace();
    }
    ~shapeAtom_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeAtom_rule2 : public shapeAtom {
private:
    _Base** _members[2];
    shapeOrRef* m_shapeOrRef;
    _QnonLitNodeConstraint_E_Opt* m__QnonLitNodeConstraint_E_Opt;
public:
    shapeAtom_rule2 (shapeOrRef* p_shapeOrRef, _QnonLitNodeConstraint_E_Opt* p__QnonLitNodeConstraint_E_Opt) {
	m_shapeOrRef = p_shapeOrRef;
	m__QnonLitNodeConstraint_E_Opt = p__QnonLitNodeConstraint_E_Opt;
	makeArray(_members, &m_shapeOrRef, &m__QnonLitNodeConstraint_E_Opt);
	trace();
    }
    ~shapeAtom_rule2();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeAtom_rule3 : public shapeAtom {
private:
    _Base** _members[3];
    GT_LPAREN* m_GT_LPAREN;
    shapeExpression* m_shapeExpression;
    GT_RPAREN* m_GT_RPAREN;
public:
    shapeAtom_rule3 (GT_LPAREN* p_GT_LPAREN, shapeExpression* p_shapeExpression, GT_RPAREN* p_GT_RPAREN) {
	m_GT_LPAREN = p_GT_LPAREN;
	m_shapeExpression = p_shapeExpression;
	m_GT_RPAREN = p_GT_RPAREN;
	makeArray(_members, &m_GT_LPAREN, &m_shapeExpression, &m_GT_RPAREN);
	trace();
    }
    ~shapeAtom_rule3();
    virtual size_t size () { return 3; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeAtom_rule4 : public shapeAtom {
private:
    _Base** _members[1];
    GT_DOT* m_GT_DOT;
public:
    shapeAtom_rule4 (GT_DOT* p_GT_DOT) {
	m_GT_DOT = p_GT_DOT;
	makeArray(_members, &m_GT_DOT);
	trace();
    }
    ~shapeAtom_rule4();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _QshapeOrRef_E_Opt : public yacker::_GroupProduction {
public:
    _QshapeOrRef_E_Opt () : yacker::_GroupProduction ("_QshapeOrRef_E_Opt") {}
};
class _QshapeOrRef_E_Opt_rule0 : public _QshapeOrRef_E_Opt {
public:
    _QshapeOrRef_E_Opt_rule0 () {
	trace();
    }
    ~_QshapeOrRef_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _QshapeOrRef_E_Opt_rule1 : public _QshapeOrRef_E_Opt {
private:
    _Base** _members[1];
    shapeOrRef* m_shapeOrRef;
public:
    _QshapeOrRef_E_Opt_rule1 (shapeOrRef* p_shapeOrRef) {
	m_shapeOrRef = p_shapeOrRef;
	makeArray(_members, &m_shapeOrRef);
	trace();
    }
    ~_QshapeOrRef_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _QnonLitNodeConstraint_E_Opt : public yacker::_GroupProduction {
public:
    _QnonLitNodeConstraint_E_Opt () : yacker::_GroupProduction ("_QnonLitNodeConstraint_E_Opt") {}
};
class _QnonLitNodeConstraint_E_Opt_rule0 : public _QnonLitNodeConstraint_E_Opt {
public:
    _QnonLitNodeConstraint_E_Opt_rule0 () {
	trace();
    }
    ~_QnonLitNodeConstraint_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _QnonLitNodeConstraint_E_Opt_rule1 : public _QnonLitNodeConstraint_E_Opt {
private:
    _Base** _members[1];
    nonLitNodeConstraint* m_nonLitNodeConstraint;
public:
    _QnonLitNodeConstraint_E_Opt_rule1 (nonLitNodeConstraint* p_nonLitNodeConstraint) {
	m_nonLitNodeConstraint = p_nonLitNodeConstraint;
	makeArray(_members, &m_nonLitNodeConstraint);
	trace();
    }
    ~_QnonLitNodeConstraint_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class inlineShapeAtom : public yacker::_GroupProduction {
public:
    inlineShapeAtom () : yacker::_GroupProduction ("inlineShapeAtom") {}
};
class inlineShapeAtom_rule0 : public inlineShapeAtom {
private:
    _Base** _members[2];
    nonLitNodeConstraint* m_nonLitNodeConstraint;
    _QinlineShapeOrRef_E_Opt* m__QinlineShapeOrRef_E_Opt;
public:
    inlineShapeAtom_rule0 (nonLitNodeConstraint* p_nonLitNodeConstraint, _QinlineShapeOrRef_E_Opt* p__QinlineShapeOrRef_E_Opt) {
	m_nonLitNodeConstraint = p_nonLitNodeConstraint;
	m__QinlineShapeOrRef_E_Opt = p__QinlineShapeOrRef_E_Opt;
	makeArray(_members, &m_nonLitNodeConstraint, &m__QinlineShapeOrRef_E_Opt);
	trace();
    }
    ~inlineShapeAtom_rule0();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class inlineShapeAtom_rule1 : public inlineShapeAtom {
private:
    _Base** _members[1];
    litNodeConstraint* m_litNodeConstraint;
public:
    inlineShapeAtom_rule1 (litNodeConstraint* p_litNodeConstraint) {
	m_litNodeConstraint = p_litNodeConstraint;
	makeArray(_members, &m_litNodeConstraint);
	trace();
    }
    ~inlineShapeAtom_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class inlineShapeAtom_rule2 : public inlineShapeAtom {
private:
    _Base** _members[2];
    inlineShapeOrRef* m_inlineShapeOrRef;
    _QnonLitNodeConstraint_E_Opt* m__QnonLitNodeConstraint_E_Opt;
public:
    inlineShapeAtom_rule2 (inlineShapeOrRef* p_inlineShapeOrRef, _QnonLitNodeConstraint_E_Opt* p__QnonLitNodeConstraint_E_Opt) {
	m_inlineShapeOrRef = p_inlineShapeOrRef;
	m__QnonLitNodeConstraint_E_Opt = p__QnonLitNodeConstraint_E_Opt;
	makeArray(_members, &m_inlineShapeOrRef, &m__QnonLitNodeConstraint_E_Opt);
	trace();
    }
    ~inlineShapeAtom_rule2();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class inlineShapeAtom_rule3 : public inlineShapeAtom {
private:
    _Base** _members[3];
    GT_LPAREN* m_GT_LPAREN;
    shapeExpression* m_shapeExpression;
    GT_RPAREN* m_GT_RPAREN;
public:
    inlineShapeAtom_rule3 (GT_LPAREN* p_GT_LPAREN, shapeExpression* p_shapeExpression, GT_RPAREN* p_GT_RPAREN) {
	m_GT_LPAREN = p_GT_LPAREN;
	m_shapeExpression = p_shapeExpression;
	m_GT_RPAREN = p_GT_RPAREN;
	makeArray(_members, &m_GT_LPAREN, &m_shapeExpression, &m_GT_RPAREN);
	trace();
    }
    ~inlineShapeAtom_rule3();
    virtual size_t size () { return 3; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class inlineShapeAtom_rule4 : public inlineShapeAtom {
private:
    _Base** _members[1];
    GT_DOT* m_GT_DOT;
public:
    inlineShapeAtom_rule4 (GT_DOT* p_GT_DOT) {
	m_GT_DOT = p_GT_DOT;
	makeArray(_members, &m_GT_DOT);
	trace();
    }
    ~inlineShapeAtom_rule4();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _QinlineShapeOrRef_E_Opt : public yacker::_GroupProduction {
public:
    _QinlineShapeOrRef_E_Opt () : yacker::_GroupProduction ("_QinlineShapeOrRef_E_Opt") {}
};
class _QinlineShapeOrRef_E_Opt_rule0 : public _QinlineShapeOrRef_E_Opt {
public:
    _QinlineShapeOrRef_E_Opt_rule0 () {
	trace();
    }
    ~_QinlineShapeOrRef_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _QinlineShapeOrRef_E_Opt_rule1 : public _QinlineShapeOrRef_E_Opt {
private:
    _Base** _members[1];
    inlineShapeOrRef* m_inlineShapeOrRef;
public:
    _QinlineShapeOrRef_E_Opt_rule1 (inlineShapeOrRef* p_inlineShapeOrRef) {
	m_inlineShapeOrRef = p_inlineShapeOrRef;
	makeArray(_members, &m_inlineShapeOrRef);
	trace();
    }
    ~_QinlineShapeOrRef_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeOrRef : public yacker::_GroupProduction {
public:
    shapeOrRef () : yacker::_GroupProduction ("shapeOrRef") {}
};
class shapeOrRef_rule0 : public shapeOrRef {
private:
    _Base** _members[1];
    shapeDefinition* m_shapeDefinition;
public:
    shapeOrRef_rule0 (shapeDefinition* p_shapeDefinition) {
	m_shapeDefinition = p_shapeDefinition;
	makeArray(_members, &m_shapeDefinition);
	trace();
    }
    ~shapeOrRef_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeOrRef_rule1 : public shapeOrRef {
private:
    _Base** _members[1];
    shapeRef* m_shapeRef;
public:
    shapeOrRef_rule1 (shapeRef* p_shapeRef) {
	m_shapeRef = p_shapeRef;
	makeArray(_members, &m_shapeRef);
	trace();
    }
    ~shapeOrRef_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class inlineShapeOrRef : public yacker::_GroupProduction {
public:
    inlineShapeOrRef () : yacker::_GroupProduction ("inlineShapeOrRef") {}
};
class inlineShapeOrRef_rule0 : public inlineShapeOrRef {
private:
    _Base** _members[1];
    inlineShapeDefinition* m_inlineShapeDefinition;
public:
    inlineShapeOrRef_rule0 (inlineShapeDefinition* p_inlineShapeDefinition) {
	m_inlineShapeDefinition = p_inlineShapeDefinition;
	makeArray(_members, &m_inlineShapeDefinition);
	trace();
    }
    ~inlineShapeOrRef_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class inlineShapeOrRef_rule1 : public inlineShapeOrRef {
private:
    _Base** _members[1];
    shapeRef* m_shapeRef;
public:
    inlineShapeOrRef_rule1 (shapeRef* p_shapeRef) {
	m_shapeRef = p_shapeRef;
	makeArray(_members, &m_shapeRef);
	trace();
    }
    ~inlineShapeOrRef_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeRef : public yacker::_GroupProduction {
public:
    shapeRef () : yacker::_GroupProduction ("shapeRef") {}
};
class shapeRef_rule0 : public shapeRef {
private:
    _Base** _members[1];
    ATPNAME_LN* m_ATPNAME_LN;
public:
    shapeRef_rule0 (ATPNAME_LN* p_ATPNAME_LN) {
	m_ATPNAME_LN = p_ATPNAME_LN;
	makeArray(_members, &m_ATPNAME_LN);
	trace();
    }
    ~shapeRef_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeRef_rule1 : public shapeRef {
private:
    _Base** _members[1];
    ATPNAME_NS* m_ATPNAME_NS;
public:
    shapeRef_rule1 (ATPNAME_NS* p_ATPNAME_NS) {
	m_ATPNAME_NS = p_ATPNAME_NS;
	makeArray(_members, &m_ATPNAME_NS);
	trace();
    }
    ~shapeRef_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeRef_rule2 : public shapeRef {
private:
    _Base** _members[2];
    GT_AT* m_GT_AT;
    shapeExprLabel* m_shapeExprLabel;
public:
    shapeRef_rule2 (GT_AT* p_GT_AT, shapeExprLabel* p_shapeExprLabel) {
	m_GT_AT = p_GT_AT;
	m_shapeExprLabel = p_shapeExprLabel;
	makeArray(_members, &m_GT_AT, &m_shapeExprLabel);
	trace();
    }
    ~shapeRef_rule2();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class litNodeConstraint : public yacker::_GroupProduction {
public:
    litNodeConstraint () : yacker::_GroupProduction ("litNodeConstraint") {}
};
class litNodeConstraint_rule0 : public litNodeConstraint {
private:
    _Base** _members[2];
    IT_LITERAL* m_IT_LITERAL;
    yacker::_ProductionVector<xsFacet*>* m_xsFacets;
public:
    litNodeConstraint_rule0 (IT_LITERAL* p_IT_LITERAL, yacker::_ProductionVector<xsFacet*>* p_xsFacets) {
	m_IT_LITERAL = p_IT_LITERAL;
	m_xsFacets = p_xsFacets;
	makeArray(_members, &m_IT_LITERAL, &m_xsFacets);
	trace();
    }
    ~litNodeConstraint_rule0();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class litNodeConstraint_rule1 : public litNodeConstraint {
private:
    _Base** _members[2];
    datatype* m_datatype;
    yacker::_ProductionVector<xsFacet*>* m_xsFacets;
public:
    litNodeConstraint_rule1 (datatype* p_datatype, yacker::_ProductionVector<xsFacet*>* p_xsFacets) {
	m_datatype = p_datatype;
	m_xsFacets = p_xsFacets;
	makeArray(_members, &m_datatype, &m_xsFacets);
	trace();
    }
    ~litNodeConstraint_rule1();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class litNodeConstraint_rule2 : public litNodeConstraint {
private:
    _Base** _members[2];
    valueSet* m_valueSet;
    yacker::_ProductionVector<xsFacet*>* m_xsFacets;
public:
    litNodeConstraint_rule2 (valueSet* p_valueSet, yacker::_ProductionVector<xsFacet*>* p_xsFacets) {
	m_valueSet = p_valueSet;
	m_xsFacets = p_xsFacets;
	makeArray(_members, &m_valueSet, &m_xsFacets);
	trace();
    }
    ~litNodeConstraint_rule2();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class litNodeConstraint_rule3 : public litNodeConstraint {
private:
    _Base** _members[1];
    yacker::_ProductionVector<numericFacet*>* m_numericFacets;
public:
    litNodeConstraint_rule3 (yacker::_ProductionVector<numericFacet*>* p_numericFacets) {
	m_numericFacets = p_numericFacets;
	makeArray(_members, &m_numericFacets);
	trace();
    }
    ~litNodeConstraint_rule3();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class nonLitNodeConstraint : public yacker::_GroupProduction {
public:
    nonLitNodeConstraint () : yacker::_GroupProduction ("nonLitNodeConstraint") {}
};
class nonLitNodeConstraint_rule0 : public nonLitNodeConstraint {
private:
    _Base** _members[2];
    nonLiteralKind* m_nonLiteralKind;
    yacker::_ProductionVector<stringFacet*>* m_stringFacets;
public:
    nonLitNodeConstraint_rule0 (nonLiteralKind* p_nonLiteralKind, yacker::_ProductionVector<stringFacet*>* p_stringFacets) {
	m_nonLiteralKind = p_nonLiteralKind;
	m_stringFacets = p_stringFacets;
	makeArray(_members, &m_nonLiteralKind, &m_stringFacets);
	trace();
    }
    ~nonLitNodeConstraint_rule0();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class nonLitNodeConstraint_rule1 : public nonLitNodeConstraint {
private:
    _Base** _members[1];
    yacker::_ProductionVector<stringFacet*>* m_stringFacets;
public:
    nonLitNodeConstraint_rule1 (yacker::_ProductionVector<stringFacet*>* p_stringFacets) {
	m_stringFacets = p_stringFacets;
	makeArray(_members, &m_stringFacets);
	trace();
    }
    ~nonLitNodeConstraint_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class nonLiteralKind : public yacker::_GroupProduction {
public:
    nonLiteralKind () : yacker::_GroupProduction ("nonLiteralKind") {}
};
class nonLiteralKind_rule0 : public nonLiteralKind {
private:
    _Base** _members[1];
    IT_IRI* m_IT_IRI;
public:
    nonLiteralKind_rule0 (IT_IRI* p_IT_IRI) {
	m_IT_IRI = p_IT_IRI;
	makeArray(_members, &m_IT_IRI);
	trace();
    }
    ~nonLiteralKind_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class nonLiteralKind_rule1 : public nonLiteralKind {
private:
    _Base** _members[1];
    IT_BNODE* m_IT_BNODE;
public:
    nonLiteralKind_rule1 (IT_BNODE* p_IT_BNODE) {
	m_IT_BNODE = p_IT_BNODE;
	makeArray(_members, &m_IT_BNODE);
	trace();
    }
    ~nonLiteralKind_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class nonLiteralKind_rule2 : public nonLiteralKind {
private:
    _Base** _members[1];
    IT_NONLITERAL* m_IT_NONLITERAL;
public:
    nonLiteralKind_rule2 (IT_NONLITERAL* p_IT_NONLITERAL) {
	m_IT_NONLITERAL = p_IT_NONLITERAL;
	makeArray(_members, &m_IT_NONLITERAL);
	trace();
    }
    ~nonLiteralKind_rule2();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class xsFacet : public yacker::_GroupProduction {
public:
    xsFacet () : yacker::_GroupProduction ("xsFacet") {}
};
class xsFacet_rule0 : public xsFacet {
private:
    _Base** _members[1];
    stringFacet* m_stringFacet;
public:
    xsFacet_rule0 (stringFacet* p_stringFacet) {
	m_stringFacet = p_stringFacet;
	makeArray(_members, &m_stringFacet);
	trace();
    }
    ~xsFacet_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class xsFacet_rule1 : public xsFacet {
private:
    _Base** _members[1];
    numericFacet* m_numericFacet;
public:
    xsFacet_rule1 (numericFacet* p_numericFacet) {
	m_numericFacet = p_numericFacet;
	makeArray(_members, &m_numericFacet);
	trace();
    }
    ~xsFacet_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class stringFacet : public yacker::_GroupProduction {
public:
    stringFacet () : yacker::_GroupProduction ("stringFacet") {}
};
class stringFacet_rule0 : public stringFacet {
private:
    _Base** _members[2];
    stringLength* m_stringLength;
    INTEGER* m_INTEGER;
public:
    stringFacet_rule0 (stringLength* p_stringLength, INTEGER* p_INTEGER) {
	m_stringLength = p_stringLength;
	m_INTEGER = p_INTEGER;
	makeArray(_members, &m_stringLength, &m_INTEGER);
	trace();
    }
    ~stringFacet_rule0();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class stringFacet_rule1 : public stringFacet {
private:
    _Base** _members[1];
    REGEXP* m_REGEXP;
public:
    stringFacet_rule1 (REGEXP* p_REGEXP) {
	m_REGEXP = p_REGEXP;
	makeArray(_members, &m_REGEXP);
	trace();
    }
    ~stringFacet_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class stringLength : public yacker::_GroupProduction {
public:
    stringLength () : yacker::_GroupProduction ("stringLength") {}
};
class stringLength_rule0 : public stringLength {
private:
    _Base** _members[1];
    IT_LENGTH* m_IT_LENGTH;
public:
    stringLength_rule0 (IT_LENGTH* p_IT_LENGTH) {
	m_IT_LENGTH = p_IT_LENGTH;
	makeArray(_members, &m_IT_LENGTH);
	trace();
    }
    ~stringLength_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class stringLength_rule1 : public stringLength {
private:
    _Base** _members[1];
    IT_MINLENGTH* m_IT_MINLENGTH;
public:
    stringLength_rule1 (IT_MINLENGTH* p_IT_MINLENGTH) {
	m_IT_MINLENGTH = p_IT_MINLENGTH;
	makeArray(_members, &m_IT_MINLENGTH);
	trace();
    }
    ~stringLength_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class stringLength_rule2 : public stringLength {
private:
    _Base** _members[1];
    IT_MAXLENGTH* m_IT_MAXLENGTH;
public:
    stringLength_rule2 (IT_MAXLENGTH* p_IT_MAXLENGTH) {
	m_IT_MAXLENGTH = p_IT_MAXLENGTH;
	makeArray(_members, &m_IT_MAXLENGTH);
	trace();
    }
    ~stringLength_rule2();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class numericFacet : public yacker::_GroupProduction {
public:
    numericFacet () : yacker::_GroupProduction ("numericFacet") {}
};
class numericFacet_rule0 : public numericFacet {
private:
    _Base** _members[2];
    numericRange* m_numericRange;
    rawNumeric* m_rawNumeric;
public:
    numericFacet_rule0 (numericRange* p_numericRange, rawNumeric* p_rawNumeric) {
	m_numericRange = p_numericRange;
	m_rawNumeric = p_rawNumeric;
	makeArray(_members, &m_numericRange, &m_rawNumeric);
	trace();
    }
    ~numericFacet_rule0();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class numericFacet_rule1 : public numericFacet {
private:
    _Base** _members[2];
    numericLength* m_numericLength;
    INTEGER* m_INTEGER;
public:
    numericFacet_rule1 (numericLength* p_numericLength, INTEGER* p_INTEGER) {
	m_numericLength = p_numericLength;
	m_INTEGER = p_INTEGER;
	makeArray(_members, &m_numericLength, &m_INTEGER);
	trace();
    }
    ~numericFacet_rule1();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class numericRange : public yacker::_GroupProduction {
public:
    numericRange () : yacker::_GroupProduction ("numericRange") {}
};
class numericRange_rule0 : public numericRange {
private:
    _Base** _members[1];
    IT_MININCLUSIVE* m_IT_MININCLUSIVE;
public:
    numericRange_rule0 (IT_MININCLUSIVE* p_IT_MININCLUSIVE) {
	m_IT_MININCLUSIVE = p_IT_MININCLUSIVE;
	makeArray(_members, &m_IT_MININCLUSIVE);
	trace();
    }
    ~numericRange_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class numericRange_rule1 : public numericRange {
private:
    _Base** _members[1];
    IT_MINEXCLUSIVE* m_IT_MINEXCLUSIVE;
public:
    numericRange_rule1 (IT_MINEXCLUSIVE* p_IT_MINEXCLUSIVE) {
	m_IT_MINEXCLUSIVE = p_IT_MINEXCLUSIVE;
	makeArray(_members, &m_IT_MINEXCLUSIVE);
	trace();
    }
    ~numericRange_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class numericRange_rule2 : public numericRange {
private:
    _Base** _members[1];
    IT_MAXINCLUSIVE* m_IT_MAXINCLUSIVE;
public:
    numericRange_rule2 (IT_MAXINCLUSIVE* p_IT_MAXINCLUSIVE) {
	m_IT_MAXINCLUSIVE = p_IT_MAXINCLUSIVE;
	makeArray(_members, &m_IT_MAXINCLUSIVE);
	trace();
    }
    ~numericRange_rule2();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class numericRange_rule3 : public numericRange {
private:
    _Base** _members[1];
    IT_MAXEXCLUSIVE* m_IT_MAXEXCLUSIVE;
public:
    numericRange_rule3 (IT_MAXEXCLUSIVE* p_IT_MAXEXCLUSIVE) {
	m_IT_MAXEXCLUSIVE = p_IT_MAXEXCLUSIVE;
	makeArray(_members, &m_IT_MAXEXCLUSIVE);
	trace();
    }
    ~numericRange_rule3();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class numericLength : public yacker::_GroupProduction {
public:
    numericLength () : yacker::_GroupProduction ("numericLength") {}
};
class numericLength_rule0 : public numericLength {
private:
    _Base** _members[1];
    IT_TOTALDIGITS* m_IT_TOTALDIGITS;
public:
    numericLength_rule0 (IT_TOTALDIGITS* p_IT_TOTALDIGITS) {
	m_IT_TOTALDIGITS = p_IT_TOTALDIGITS;
	makeArray(_members, &m_IT_TOTALDIGITS);
	trace();
    }
    ~numericLength_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class numericLength_rule1 : public numericLength {
private:
    _Base** _members[1];
    IT_FRACTIONDIGITS* m_IT_FRACTIONDIGITS;
public:
    numericLength_rule1 (IT_FRACTIONDIGITS* p_IT_FRACTIONDIGITS) {
	m_IT_FRACTIONDIGITS = p_IT_FRACTIONDIGITS;
	makeArray(_members, &m_IT_FRACTIONDIGITS);
	trace();
    }
    ~numericLength_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeDefinition : public yacker::_Production {
private:
    _Base** _members[3];
    inlineShapeDefinition* m_inlineShapeDefinition;
    yacker::_ProductionVector<annotation*>* m_annotations;
    semanticActions* m_semanticActions;
    virtual const char* getProductionName () { return "shapeDefinition"; }
public:
    shapeDefinition (inlineShapeDefinition* p_inlineShapeDefinition, yacker::_ProductionVector<annotation*>* p_annotations, semanticActions* p_semanticActions) : yacker::_Production("shapeDefinition") {
	m_inlineShapeDefinition = p_inlineShapeDefinition;
	m_annotations = p_annotations;
	m_semanticActions = p_semanticActions;
	makeArray(_members, &m_inlineShapeDefinition, &m_annotations, &m_semanticActions);
	trace();
    }
    ~shapeDefinition();
    virtual size_t size () { return 3; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class inlineShapeDefinition : public yacker::_Production {
private:
    _Base** _members[4];
    shapeQualifiers* m_shapeQualifiers;
    GT_LCURLEY* m_GT_LCURLEY;
    _QtripleExpression_E_Opt* m__QtripleExpression_E_Opt;
    GT_RCURLEY* m_GT_RCURLEY;
    virtual const char* getProductionName () { return "inlineShapeDefinition"; }
public:
    inlineShapeDefinition (shapeQualifiers* p_shapeQualifiers, GT_LCURLEY* p_GT_LCURLEY, _QtripleExpression_E_Opt* p__QtripleExpression_E_Opt, GT_RCURLEY* p_GT_RCURLEY) : yacker::_Production("inlineShapeDefinition") {
	m_shapeQualifiers = p_shapeQualifiers;
	m_GT_LCURLEY = p_GT_LCURLEY;
	m__QtripleExpression_E_Opt = p__QtripleExpression_E_Opt;
	m_GT_RCURLEY = p_GT_RCURLEY;
	makeArray(_members, &m_shapeQualifiers, &m_GT_LCURLEY, &m__QtripleExpression_E_Opt, &m_GT_RCURLEY);
	trace();
    }
    ~inlineShapeDefinition();
    virtual size_t size () { return 4; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _QtripleExpression_E_Opt : public yacker::_GroupProduction {
public:
    _QtripleExpression_E_Opt () : yacker::_GroupProduction ("_QtripleExpression_E_Opt") {}
};
class _QtripleExpression_E_Opt_rule0 : public _QtripleExpression_E_Opt {
public:
    _QtripleExpression_E_Opt_rule0 () {
	trace();
    }
    ~_QtripleExpression_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _QtripleExpression_E_Opt_rule1 : public _QtripleExpression_E_Opt {
private:
    _Base** _members[1];
    tripleExpression* m_tripleExpression;
public:
    _QtripleExpression_E_Opt_rule1 (tripleExpression* p_tripleExpression) {
	m_tripleExpression = p_tripleExpression;
	makeArray(_members, &m_tripleExpression);
	trace();
    }
    ~_QtripleExpression_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeQualifiers : public yacker::_Production {
private:
    _Base** _members[1];
    yacker::_ProductionVector<_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C*>* m__O_QextraPropertySet_E_Or_QIT_CLOSED_E_Cs;
    virtual const char* getProductionName () { return "shapeQualifiers"; }
public:
    shapeQualifiers (yacker::_ProductionVector<_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C*>* p__O_QextraPropertySet_E_Or_QIT_CLOSED_E_Cs) : yacker::_Production("shapeQualifiers") {
	m__O_QextraPropertySet_E_Or_QIT_CLOSED_E_Cs = p__O_QextraPropertySet_E_Or_QIT_CLOSED_E_Cs;
	makeArray(_members, &m__O_QextraPropertySet_E_Or_QIT_CLOSED_E_Cs);
	trace();
    }
    ~shapeQualifiers();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C : public yacker::_GroupProduction {
public:
    _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C () : yacker::_GroupProduction ("_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C") {}
};
class _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_rule0 : public _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C {
private:
    _Base** _members[1];
    extraPropertySet* m_extraPropertySet;
public:
    _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_rule0 (extraPropertySet* p_extraPropertySet) {
	m_extraPropertySet = p_extraPropertySet;
	makeArray(_members, &m_extraPropertySet);
	trace();
    }
    ~_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_rule1 : public _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C {
private:
    _Base** _members[1];
    IT_CLOSED* m_IT_CLOSED;
public:
    _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_rule1 (IT_CLOSED* p_IT_CLOSED) {
	m_IT_CLOSED = p_IT_CLOSED;
	makeArray(_members, &m_IT_CLOSED);
	trace();
    }
    ~_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class extraPropertySet : public yacker::_Production {
private:
    _Base** _members[2];
    IT_EXTRA* m_IT_EXTRA;
    yacker::_ProductionVector<predicate*>* m_predicates;
    virtual const char* getProductionName () { return "extraPropertySet"; }
public:
    extraPropertySet (IT_EXTRA* p_IT_EXTRA, yacker::_ProductionVector<predicate*>* p_predicates) : yacker::_Production("extraPropertySet") {
	m_IT_EXTRA = p_IT_EXTRA;
	m_predicates = p_predicates;
	makeArray(_members, &m_IT_EXTRA, &m_predicates);
	trace();
    }
    ~extraPropertySet();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class tripleExpression : public yacker::_Production {
private:
    _Base** _members[1];
    oneOfTripleExpr* m_oneOfTripleExpr;
    virtual const char* getProductionName () { return "tripleExpression"; }
public:
    tripleExpression (oneOfTripleExpr* p_oneOfTripleExpr) : yacker::_Production("tripleExpression") {
	m_oneOfTripleExpr = p_oneOfTripleExpr;
	makeArray(_members, &m_oneOfTripleExpr);
	trace();
    }
    ~tripleExpression();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class oneOfTripleExpr : public yacker::_GroupProduction {
public:
    oneOfTripleExpr () : yacker::_GroupProduction ("oneOfTripleExpr") {}
};
class oneOfTripleExpr_rule0 : public oneOfTripleExpr {
private:
    _Base** _members[1];
    groupTripleExpr* m_groupTripleExpr;
public:
    oneOfTripleExpr_rule0 (groupTripleExpr* p_groupTripleExpr) {
	m_groupTripleExpr = p_groupTripleExpr;
	makeArray(_members, &m_groupTripleExpr);
	trace();
    }
    ~oneOfTripleExpr_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class oneOfTripleExpr_rule1 : public oneOfTripleExpr {
private:
    _Base** _members[1];
    multiElementOneOf* m_multiElementOneOf;
public:
    oneOfTripleExpr_rule1 (multiElementOneOf* p_multiElementOneOf) {
	m_multiElementOneOf = p_multiElementOneOf;
	makeArray(_members, &m_multiElementOneOf);
	trace();
    }
    ~oneOfTripleExpr_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class multiElementOneOf : public yacker::_Production {
private:
    _Base** _members[2];
    groupTripleExpr* m_groupTripleExpr;
    yacker::_ProductionVector<_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C*>* m__O_QGT_PIPE_E_S_QgroupTripleExpr_E_Cs;
    virtual const char* getProductionName () { return "multiElementOneOf"; }
public:
    multiElementOneOf (groupTripleExpr* p_groupTripleExpr, yacker::_ProductionVector<_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C*>* p__O_QGT_PIPE_E_S_QgroupTripleExpr_E_Cs) : yacker::_Production("multiElementOneOf") {
	m_groupTripleExpr = p_groupTripleExpr;
	m__O_QGT_PIPE_E_S_QgroupTripleExpr_E_Cs = p__O_QGT_PIPE_E_S_QgroupTripleExpr_E_Cs;
	makeArray(_members, &m_groupTripleExpr, &m__O_QGT_PIPE_E_S_QgroupTripleExpr_E_Cs);
	trace();
    }
    ~multiElementOneOf();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QGT_PIPE_E_S_QgroupTripleExpr_E_C : public yacker::_Production {
private:
    _Base** _members[2];
    GT_PIPE* m_GT_PIPE;
    groupTripleExpr* m_groupTripleExpr;
    virtual const char* getProductionName () { return "_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C"; }
public:
    _O_QGT_PIPE_E_S_QgroupTripleExpr_E_C (GT_PIPE* p_GT_PIPE, groupTripleExpr* p_groupTripleExpr) : yacker::_Production("_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C") {
	m_GT_PIPE = p_GT_PIPE;
	m_groupTripleExpr = p_groupTripleExpr;
	makeArray(_members, &m_GT_PIPE, &m_groupTripleExpr);
	trace();
    }
    ~_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class groupTripleExpr : public yacker::_GroupProduction {
public:
    groupTripleExpr () : yacker::_GroupProduction ("groupTripleExpr") {}
};
class groupTripleExpr_rule0 : public groupTripleExpr {
private:
    _Base** _members[1];
    singleElementGroup* m_singleElementGroup;
public:
    groupTripleExpr_rule0 (singleElementGroup* p_singleElementGroup) {
	m_singleElementGroup = p_singleElementGroup;
	makeArray(_members, &m_singleElementGroup);
	trace();
    }
    ~groupTripleExpr_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class groupTripleExpr_rule1 : public groupTripleExpr {
private:
    _Base** _members[1];
    multiElementGroup* m_multiElementGroup;
public:
    groupTripleExpr_rule1 (multiElementGroup* p_multiElementGroup) {
	m_multiElementGroup = p_multiElementGroup;
	makeArray(_members, &m_multiElementGroup);
	trace();
    }
    ~groupTripleExpr_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class singleElementGroup : public yacker::_Production {
private:
    _Base** _members[2];
    unaryTripleExpr* m_unaryTripleExpr;
    _QGT_SEMI_E_Opt* m__QGT_SEMI_E_Opt;
    virtual const char* getProductionName () { return "singleElementGroup"; }
public:
    singleElementGroup (unaryTripleExpr* p_unaryTripleExpr, _QGT_SEMI_E_Opt* p__QGT_SEMI_E_Opt) : yacker::_Production("singleElementGroup") {
	m_unaryTripleExpr = p_unaryTripleExpr;
	m__QGT_SEMI_E_Opt = p__QGT_SEMI_E_Opt;
	makeArray(_members, &m_unaryTripleExpr, &m__QGT_SEMI_E_Opt);
	trace();
    }
    ~singleElementGroup();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _QGT_SEMI_E_Opt : public yacker::_GroupProduction {
public:
    _QGT_SEMI_E_Opt () : yacker::_GroupProduction ("_QGT_SEMI_E_Opt") {}
};
class _QGT_SEMI_E_Opt_rule0 : public _QGT_SEMI_E_Opt {
public:
    _QGT_SEMI_E_Opt_rule0 () {
	trace();
    }
    ~_QGT_SEMI_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _QGT_SEMI_E_Opt_rule1 : public _QGT_SEMI_E_Opt {
private:
    _Base** _members[1];
    GT_SEMI* m_GT_SEMI;
public:
    _QGT_SEMI_E_Opt_rule1 (GT_SEMI* p_GT_SEMI) {
	m_GT_SEMI = p_GT_SEMI;
	makeArray(_members, &m_GT_SEMI);
	trace();
    }
    ~_QGT_SEMI_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class multiElementGroup : public yacker::_Production {
private:
    _Base** _members[3];
    unaryTripleExpr* m_unaryTripleExpr;
    yacker::_ProductionVector<_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C*>* m__O_QGT_SEMI_E_S_QunaryTripleExpr_E_Cs;
    _QGT_SEMI_E_Opt* m__QGT_SEMI_E_Opt;
    virtual const char* getProductionName () { return "multiElementGroup"; }
public:
    multiElementGroup (unaryTripleExpr* p_unaryTripleExpr, yacker::_ProductionVector<_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C*>* p__O_QGT_SEMI_E_S_QunaryTripleExpr_E_Cs, _QGT_SEMI_E_Opt* p__QGT_SEMI_E_Opt) : yacker::_Production("multiElementGroup") {
	m_unaryTripleExpr = p_unaryTripleExpr;
	m__O_QGT_SEMI_E_S_QunaryTripleExpr_E_Cs = p__O_QGT_SEMI_E_S_QunaryTripleExpr_E_Cs;
	m__QGT_SEMI_E_Opt = p__QGT_SEMI_E_Opt;
	makeArray(_members, &m_unaryTripleExpr, &m__O_QGT_SEMI_E_S_QunaryTripleExpr_E_Cs, &m__QGT_SEMI_E_Opt);
	trace();
    }
    ~multiElementGroup();
    virtual size_t size () { return 3; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QGT_SEMI_E_S_QunaryTripleExpr_E_C : public yacker::_Production {
private:
    _Base** _members[2];
    GT_SEMI* m_GT_SEMI;
    unaryTripleExpr* m_unaryTripleExpr;
    virtual const char* getProductionName () { return "_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C"; }
public:
    _O_QGT_SEMI_E_S_QunaryTripleExpr_E_C (GT_SEMI* p_GT_SEMI, unaryTripleExpr* p_unaryTripleExpr) : yacker::_Production("_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C") {
	m_GT_SEMI = p_GT_SEMI;
	m_unaryTripleExpr = p_unaryTripleExpr;
	makeArray(_members, &m_GT_SEMI, &m_unaryTripleExpr);
	trace();
    }
    ~_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class unaryTripleExpr : public yacker::_GroupProduction {
public:
    unaryTripleExpr () : yacker::_GroupProduction ("unaryTripleExpr") {}
};
class unaryTripleExpr_rule0 : public unaryTripleExpr {
private:
    _Base** _members[2];
    _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt* m__Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt;
    _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C* m__O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C;
public:
    unaryTripleExpr_rule0 (_Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt* p__Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt, _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C* p__O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C) {
	m__Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt = p__Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt;
	m__O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C = p__O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C;
	makeArray(_members, &m__Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt, &m__O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C);
	trace();
    }
    ~unaryTripleExpr_rule0();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class unaryTripleExpr_rule1 : public unaryTripleExpr {
private:
    _Base** _members[1];
    include* m_include;
public:
    unaryTripleExpr_rule1 (include* p_include) {
	m_include = p_include;
	makeArray(_members, &m_include);
	trace();
    }
    ~unaryTripleExpr_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C : public yacker::_Production {
private:
    _Base** _members[2];
    GT_DOLLAR* m_GT_DOLLAR;
    tripleExprLabel* m_tripleExprLabel;
    virtual const char* getProductionName () { return "_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C"; }
public:
    _O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C (GT_DOLLAR* p_GT_DOLLAR, tripleExprLabel* p_tripleExprLabel) : yacker::_Production("_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C") {
	m_GT_DOLLAR = p_GT_DOLLAR;
	m_tripleExprLabel = p_tripleExprLabel;
	makeArray(_members, &m_GT_DOLLAR, &m_tripleExprLabel);
	trace();
    }
    ~_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt : public yacker::_GroupProduction {
public:
    _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt () : yacker::_GroupProduction ("_Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt") {}
};
class _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt_rule0 : public _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt {
public:
    _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt_rule0 () {
	trace();
    }
    ~_Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt_rule1 : public _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt {
private:
    _Base** _members[1];
    _O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C* m__O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C;
public:
    _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt_rule1 (_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C* p__O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C) {
	m__O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C = p__O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C;
	makeArray(_members, &m__O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C);
	trace();
    }
    ~_Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C : public yacker::_GroupProduction {
public:
    _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C () : yacker::_GroupProduction ("_O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C") {}
};
class _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C_rule0 : public _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C {
private:
    _Base** _members[1];
    tripleConstraint* m_tripleConstraint;
public:
    _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C_rule0 (tripleConstraint* p_tripleConstraint) {
	m_tripleConstraint = p_tripleConstraint;
	makeArray(_members, &m_tripleConstraint);
	trace();
    }
    ~_O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C_rule1 : public _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C {
private:
    _Base** _members[1];
    bracketedTripleExpr* m_bracketedTripleExpr;
public:
    _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C_rule1 (bracketedTripleExpr* p_bracketedTripleExpr) {
	m_bracketedTripleExpr = p_bracketedTripleExpr;
	makeArray(_members, &m_bracketedTripleExpr);
	trace();
    }
    ~_O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class bracketedTripleExpr : public yacker::_Production {
private:
    _Base** _members[6];
    GT_LPAREN* m_GT_LPAREN;
    tripleExpression* m_tripleExpression;
    GT_RPAREN* m_GT_RPAREN;
    _Qcardinality_E_Opt* m__Qcardinality_E_Opt;
    yacker::_ProductionVector<annotation*>* m_annotations;
    semanticActions* m_semanticActions;
    virtual const char* getProductionName () { return "bracketedTripleExpr"; }
public:
    bracketedTripleExpr (GT_LPAREN* p_GT_LPAREN, tripleExpression* p_tripleExpression, GT_RPAREN* p_GT_RPAREN, _Qcardinality_E_Opt* p__Qcardinality_E_Opt, yacker::_ProductionVector<annotation*>* p_annotations, semanticActions* p_semanticActions) : yacker::_Production("bracketedTripleExpr") {
	m_GT_LPAREN = p_GT_LPAREN;
	m_tripleExpression = p_tripleExpression;
	m_GT_RPAREN = p_GT_RPAREN;
	m__Qcardinality_E_Opt = p__Qcardinality_E_Opt;
	m_annotations = p_annotations;
	m_semanticActions = p_semanticActions;
	makeArray(_members, &m_GT_LPAREN, &m_tripleExpression, &m_GT_RPAREN, &m__Qcardinality_E_Opt, &m_annotations, &m_semanticActions);
	trace();
    }
    ~bracketedTripleExpr();
    virtual size_t size () { return 6; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _Qcardinality_E_Opt : public yacker::_GroupProduction {
public:
    _Qcardinality_E_Opt () : yacker::_GroupProduction ("_Qcardinality_E_Opt") {}
};
class _Qcardinality_E_Opt_rule0 : public _Qcardinality_E_Opt {
public:
    _Qcardinality_E_Opt_rule0 () {
	trace();
    }
    ~_Qcardinality_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _Qcardinality_E_Opt_rule1 : public _Qcardinality_E_Opt {
private:
    _Base** _members[1];
    cardinality* m_cardinality;
public:
    _Qcardinality_E_Opt_rule1 (cardinality* p_cardinality) {
	m_cardinality = p_cardinality;
	makeArray(_members, &m_cardinality);
	trace();
    }
    ~_Qcardinality_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class tripleConstraint : public yacker::_Production {
private:
    _Base** _members[6];
    _QsenseFlags_E_Opt* m__QsenseFlags_E_Opt;
    predicate* m_predicate;
    inlineShapeExpression* m_inlineShapeExpression;
    _Qcardinality_E_Opt* m__Qcardinality_E_Opt;
    yacker::_ProductionVector<annotation*>* m_annotations;
    semanticActions* m_semanticActions;
    virtual const char* getProductionName () { return "tripleConstraint"; }
public:
    tripleConstraint (_QsenseFlags_E_Opt* p__QsenseFlags_E_Opt, predicate* p_predicate, inlineShapeExpression* p_inlineShapeExpression, _Qcardinality_E_Opt* p__Qcardinality_E_Opt, yacker::_ProductionVector<annotation*>* p_annotations, semanticActions* p_semanticActions) : yacker::_Production("tripleConstraint") {
	m__QsenseFlags_E_Opt = p__QsenseFlags_E_Opt;
	m_predicate = p_predicate;
	m_inlineShapeExpression = p_inlineShapeExpression;
	m__Qcardinality_E_Opt = p__Qcardinality_E_Opt;
	m_annotations = p_annotations;
	m_semanticActions = p_semanticActions;
	makeArray(_members, &m__QsenseFlags_E_Opt, &m_predicate, &m_inlineShapeExpression, &m__Qcardinality_E_Opt, &m_annotations, &m_semanticActions);
	trace();
    }
    ~tripleConstraint();
    virtual size_t size () { return 6; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _QsenseFlags_E_Opt : public yacker::_GroupProduction {
public:
    _QsenseFlags_E_Opt () : yacker::_GroupProduction ("_QsenseFlags_E_Opt") {}
};
class _QsenseFlags_E_Opt_rule0 : public _QsenseFlags_E_Opt {
public:
    _QsenseFlags_E_Opt_rule0 () {
	trace();
    }
    ~_QsenseFlags_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _QsenseFlags_E_Opt_rule1 : public _QsenseFlags_E_Opt {
private:
    _Base** _members[1];
    senseFlags* m_senseFlags;
public:
    _QsenseFlags_E_Opt_rule1 (senseFlags* p_senseFlags) {
	m_senseFlags = p_senseFlags;
	makeArray(_members, &m_senseFlags);
	trace();
    }
    ~_QsenseFlags_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class cardinality : public yacker::_GroupProduction {
public:
    cardinality () : yacker::_GroupProduction ("cardinality") {}
};
class cardinality_rule0 : public cardinality {
private:
    _Base** _members[1];
    GT_TIMES* m_GT_TIMES;
public:
    cardinality_rule0 (GT_TIMES* p_GT_TIMES) {
	m_GT_TIMES = p_GT_TIMES;
	makeArray(_members, &m_GT_TIMES);
	trace();
    }
    ~cardinality_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class cardinality_rule1 : public cardinality {
private:
    _Base** _members[1];
    GT_PLUS* m_GT_PLUS;
public:
    cardinality_rule1 (GT_PLUS* p_GT_PLUS) {
	m_GT_PLUS = p_GT_PLUS;
	makeArray(_members, &m_GT_PLUS);
	trace();
    }
    ~cardinality_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class cardinality_rule2 : public cardinality {
private:
    _Base** _members[1];
    GT_OPT* m_GT_OPT;
public:
    cardinality_rule2 (GT_OPT* p_GT_OPT) {
	m_GT_OPT = p_GT_OPT;
	makeArray(_members, &m_GT_OPT);
	trace();
    }
    ~cardinality_rule2();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class cardinality_rule3 : public cardinality {
private:
    _Base** _members[1];
    REPEAT_RANGE* m_REPEAT_RANGE;
public:
    cardinality_rule3 (REPEAT_RANGE* p_REPEAT_RANGE) {
	m_REPEAT_RANGE = p_REPEAT_RANGE;
	makeArray(_members, &m_REPEAT_RANGE);
	trace();
    }
    ~cardinality_rule3();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class senseFlags : public yacker::_Production {
private:
    _Base** _members[1];
    GT_CARROT* m_GT_CARROT;
    virtual const char* getProductionName () { return "senseFlags"; }
public:
    senseFlags (GT_CARROT* p_GT_CARROT) : yacker::_Production("senseFlags") {
	m_GT_CARROT = p_GT_CARROT;
	makeArray(_members, &m_GT_CARROT);
	trace();
    }
    ~senseFlags();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class valueSet : public yacker::_Production {
private:
    _Base** _members[3];
    GT_LBRACKET* m_GT_LBRACKET;
    yacker::_ProductionVector<valueSetValue*>* m_valueSetValues;
    GT_RBRACKET* m_GT_RBRACKET;
    virtual const char* getProductionName () { return "valueSet"; }
public:
    valueSet (GT_LBRACKET* p_GT_LBRACKET, yacker::_ProductionVector<valueSetValue*>* p_valueSetValues, GT_RBRACKET* p_GT_RBRACKET) : yacker::_Production("valueSet") {
	m_GT_LBRACKET = p_GT_LBRACKET;
	m_valueSetValues = p_valueSetValues;
	m_GT_RBRACKET = p_GT_RBRACKET;
	makeArray(_members, &m_GT_LBRACKET, &m_valueSetValues, &m_GT_RBRACKET);
	trace();
    }
    ~valueSet();
    virtual size_t size () { return 3; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class valueSetValue : public yacker::_GroupProduction {
public:
    valueSetValue () : yacker::_GroupProduction ("valueSetValue") {}
};
class valueSetValue_rule0 : public valueSetValue {
private:
    _Base** _members[1];
    iriRange* m_iriRange;
public:
    valueSetValue_rule0 (iriRange* p_iriRange) {
	m_iriRange = p_iriRange;
	makeArray(_members, &m_iriRange);
	trace();
    }
    ~valueSetValue_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class valueSetValue_rule1 : public valueSetValue {
private:
    _Base** _members[1];
    literalRange* m_literalRange;
public:
    valueSetValue_rule1 (literalRange* p_literalRange) {
	m_literalRange = p_literalRange;
	makeArray(_members, &m_literalRange);
	trace();
    }
    ~valueSetValue_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class valueSetValue_rule2 : public valueSetValue {
private:
    _Base** _members[1];
    languageRange* m_languageRange;
public:
    valueSetValue_rule2 (languageRange* p_languageRange) {
	m_languageRange = p_languageRange;
	makeArray(_members, &m_languageRange);
	trace();
    }
    ~valueSetValue_rule2();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class valueSetValue_rule3 : public valueSetValue {
private:
    _Base** _members[2];
    GT_DOT* m_GT_DOT;
    _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C* m__O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C;
public:
    valueSetValue_rule3 (GT_DOT* p_GT_DOT, _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C* p__O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C) {
	m_GT_DOT = p_GT_DOT;
	m__O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C = p__O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C;
	makeArray(_members, &m_GT_DOT, &m__O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C);
	trace();
    }
    ~valueSetValue_rule3();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C : public yacker::_GroupProduction {
public:
    _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C () : yacker::_GroupProduction ("_O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C") {}
};
class _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule0 : public _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C {
private:
    _Base** _members[1];
    yacker::_ProductionVector<iriExclusion*>* m_iriExclusions;
public:
    _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule0 (yacker::_ProductionVector<iriExclusion*>* p_iriExclusions) {
	m_iriExclusions = p_iriExclusions;
	makeArray(_members, &m_iriExclusions);
	trace();
    }
    ~_O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule1 : public _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C {
private:
    _Base** _members[1];
    yacker::_ProductionVector<literalExclusion*>* m_literalExclusions;
public:
    _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule1 (yacker::_ProductionVector<literalExclusion*>* p_literalExclusions) {
	m_literalExclusions = p_literalExclusions;
	makeArray(_members, &m_literalExclusions);
	trace();
    }
    ~_O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule2 : public _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C {
private:
    _Base** _members[1];
    yacker::_ProductionVector<languageExclusion*>* m_languageExclusions;
public:
    _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule2 (yacker::_ProductionVector<languageExclusion*>* p_languageExclusions) {
	m_languageExclusions = p_languageExclusions;
	makeArray(_members, &m_languageExclusions);
	trace();
    }
    ~_O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule2();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class iriRange : public yacker::_Production {
private:
    _Base** _members[2];
    iri* m_iri;
    _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt* m__Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt;
    virtual const char* getProductionName () { return "iriRange"; }
public:
    iriRange (iri* p_iri, _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt* p__Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt) : yacker::_Production("iriRange") {
	m_iri = p_iri;
	m__Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt = p__Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt;
	makeArray(_members, &m_iri, &m__Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt);
	trace();
    }
    ~iriRange();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QGT_KINDA_E_S_QiriExclusion_E_Star_C : public yacker::_Production {
private:
    _Base** _members[2];
    GT_KINDA* m_GT_KINDA;
    yacker::_ProductionVector<iriExclusion*>* m_iriExclusions;
    virtual const char* getProductionName () { return "_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C"; }
public:
    _O_QGT_KINDA_E_S_QiriExclusion_E_Star_C (GT_KINDA* p_GT_KINDA, yacker::_ProductionVector<iriExclusion*>* p_iriExclusions) : yacker::_Production("_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C") {
	m_GT_KINDA = p_GT_KINDA;
	m_iriExclusions = p_iriExclusions;
	makeArray(_members, &m_GT_KINDA, &m_iriExclusions);
	trace();
    }
    ~_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt : public yacker::_GroupProduction {
public:
    _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt () : yacker::_GroupProduction ("_Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt") {}
};
class _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt_rule0 : public _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt {
public:
    _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt_rule0 () {
	trace();
    }
    ~_Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt_rule1 : public _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt {
private:
    _Base** _members[1];
    _O_QGT_KINDA_E_S_QiriExclusion_E_Star_C* m__O_QGT_KINDA_E_S_QiriExclusion_E_Star_C;
public:
    _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt_rule1 (_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C* p__O_QGT_KINDA_E_S_QiriExclusion_E_Star_C) {
	m__O_QGT_KINDA_E_S_QiriExclusion_E_Star_C = p__O_QGT_KINDA_E_S_QiriExclusion_E_Star_C;
	makeArray(_members, &m__O_QGT_KINDA_E_S_QiriExclusion_E_Star_C);
	trace();
    }
    ~_Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class iriExclusion : public yacker::_Production {
private:
    _Base** _members[3];
    GT_MINUS* m_GT_MINUS;
    iri* m_iri;
    _QGT_KINDA_E_Opt* m__QGT_KINDA_E_Opt;
    virtual const char* getProductionName () { return "iriExclusion"; }
public:
    iriExclusion (GT_MINUS* p_GT_MINUS, iri* p_iri, _QGT_KINDA_E_Opt* p__QGT_KINDA_E_Opt) : yacker::_Production("iriExclusion") {
	m_GT_MINUS = p_GT_MINUS;
	m_iri = p_iri;
	m__QGT_KINDA_E_Opt = p__QGT_KINDA_E_Opt;
	makeArray(_members, &m_GT_MINUS, &m_iri, &m__QGT_KINDA_E_Opt);
	trace();
    }
    ~iriExclusion();
    virtual size_t size () { return 3; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _QGT_KINDA_E_Opt : public yacker::_GroupProduction {
public:
    _QGT_KINDA_E_Opt () : yacker::_GroupProduction ("_QGT_KINDA_E_Opt") {}
};
class _QGT_KINDA_E_Opt_rule0 : public _QGT_KINDA_E_Opt {
public:
    _QGT_KINDA_E_Opt_rule0 () {
	trace();
    }
    ~_QGT_KINDA_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _QGT_KINDA_E_Opt_rule1 : public _QGT_KINDA_E_Opt {
private:
    _Base** _members[1];
    GT_KINDA* m_GT_KINDA;
public:
    _QGT_KINDA_E_Opt_rule1 (GT_KINDA* p_GT_KINDA) {
	m_GT_KINDA = p_GT_KINDA;
	makeArray(_members, &m_GT_KINDA);
	trace();
    }
    ~_QGT_KINDA_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class literalRange : public yacker::_Production {
private:
    _Base** _members[2];
    literal* m_literal;
    _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt* m__Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt;
    virtual const char* getProductionName () { return "literalRange"; }
public:
    literalRange (literal* p_literal, _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt* p__Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt) : yacker::_Production("literalRange") {
	m_literal = p_literal;
	m__Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt = p__Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt;
	makeArray(_members, &m_literal, &m__Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt);
	trace();
    }
    ~literalRange();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C : public yacker::_Production {
private:
    _Base** _members[2];
    GT_KINDA* m_GT_KINDA;
    yacker::_ProductionVector<literalExclusion*>* m_literalExclusions;
    virtual const char* getProductionName () { return "_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C"; }
public:
    _O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C (GT_KINDA* p_GT_KINDA, yacker::_ProductionVector<literalExclusion*>* p_literalExclusions) : yacker::_Production("_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C") {
	m_GT_KINDA = p_GT_KINDA;
	m_literalExclusions = p_literalExclusions;
	makeArray(_members, &m_GT_KINDA, &m_literalExclusions);
	trace();
    }
    ~_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt : public yacker::_GroupProduction {
public:
    _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt () : yacker::_GroupProduction ("_Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt") {}
};
class _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt_rule0 : public _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt {
public:
    _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt_rule0 () {
	trace();
    }
    ~_Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt_rule1 : public _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt {
private:
    _Base** _members[1];
    _O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C* m__O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C;
public:
    _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt_rule1 (_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C* p__O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C) {
	m__O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C = p__O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C;
	makeArray(_members, &m__O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C);
	trace();
    }
    ~_Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class literalExclusion : public yacker::_Production {
private:
    _Base** _members[3];
    GT_MINUS* m_GT_MINUS;
    literal* m_literal;
    _QGT_KINDA_E_Opt* m__QGT_KINDA_E_Opt;
    virtual const char* getProductionName () { return "literalExclusion"; }
public:
    literalExclusion (GT_MINUS* p_GT_MINUS, literal* p_literal, _QGT_KINDA_E_Opt* p__QGT_KINDA_E_Opt) : yacker::_Production("literalExclusion") {
	m_GT_MINUS = p_GT_MINUS;
	m_literal = p_literal;
	m__QGT_KINDA_E_Opt = p__QGT_KINDA_E_Opt;
	makeArray(_members, &m_GT_MINUS, &m_literal, &m__QGT_KINDA_E_Opt);
	trace();
    }
    ~literalExclusion();
    virtual size_t size () { return 3; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class languageRange : public yacker::_GroupProduction {
public:
    languageRange () : yacker::_GroupProduction ("languageRange") {}
};
class languageRange_rule0 : public languageRange {
private:
    _Base** _members[2];
    LANGTAG* m_LANGTAG;
    _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt* m__Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt;
public:
    languageRange_rule0 (LANGTAG* p_LANGTAG, _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt* p__Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt) {
	m_LANGTAG = p_LANGTAG;
	m__Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt = p__Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt;
	makeArray(_members, &m_LANGTAG, &m__Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt);
	trace();
    }
    ~languageRange_rule0();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class languageRange_rule1 : public languageRange {
private:
    _Base** _members[3];
    GT_AT* m_GT_AT;
    GT_KINDA* m_GT_KINDA;
    yacker::_ProductionVector<languageExclusion*>* m_languageExclusions;
public:
    languageRange_rule1 (GT_AT* p_GT_AT, GT_KINDA* p_GT_KINDA, yacker::_ProductionVector<languageExclusion*>* p_languageExclusions) {
	m_GT_AT = p_GT_AT;
	m_GT_KINDA = p_GT_KINDA;
	m_languageExclusions = p_languageExclusions;
	makeArray(_members, &m_GT_AT, &m_GT_KINDA, &m_languageExclusions);
	trace();
    }
    ~languageRange_rule1();
    virtual size_t size () { return 3; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C : public yacker::_Production {
private:
    _Base** _members[2];
    GT_KINDA* m_GT_KINDA;
    yacker::_ProductionVector<languageExclusion*>* m_languageExclusions;
    virtual const char* getProductionName () { return "_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C"; }
public:
    _O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C (GT_KINDA* p_GT_KINDA, yacker::_ProductionVector<languageExclusion*>* p_languageExclusions) : yacker::_Production("_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C") {
	m_GT_KINDA = p_GT_KINDA;
	m_languageExclusions = p_languageExclusions;
	makeArray(_members, &m_GT_KINDA, &m_languageExclusions);
	trace();
    }
    ~_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt : public yacker::_GroupProduction {
public:
    _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt () : yacker::_GroupProduction ("_Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt") {}
};
class _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt_rule0 : public _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt {
public:
    _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt_rule0 () {
	trace();
    }
    ~_Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt_rule1 : public _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt {
private:
    _Base** _members[1];
    _O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C* m__O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C;
public:
    _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt_rule1 (_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C* p__O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C) {
	m__O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C = p__O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C;
	makeArray(_members, &m__O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C);
	trace();
    }
    ~_Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class languageExclusion : public yacker::_Production {
private:
    _Base** _members[3];
    GT_MINUS* m_GT_MINUS;
    LANGTAG* m_LANGTAG;
    _QGT_KINDA_E_Opt* m__QGT_KINDA_E_Opt;
    virtual const char* getProductionName () { return "languageExclusion"; }
public:
    languageExclusion (GT_MINUS* p_GT_MINUS, LANGTAG* p_LANGTAG, _QGT_KINDA_E_Opt* p__QGT_KINDA_E_Opt) : yacker::_Production("languageExclusion") {
	m_GT_MINUS = p_GT_MINUS;
	m_LANGTAG = p_LANGTAG;
	m__QGT_KINDA_E_Opt = p__QGT_KINDA_E_Opt;
	makeArray(_members, &m_GT_MINUS, &m_LANGTAG, &m__QGT_KINDA_E_Opt);
	trace();
    }
    ~languageExclusion();
    virtual size_t size () { return 3; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class include : public yacker::_Production {
private:
    _Base** _members[2];
    GT_AMP* m_GT_AMP;
    tripleExprLabel* m_tripleExprLabel;
    virtual const char* getProductionName () { return "include"; }
public:
    include (GT_AMP* p_GT_AMP, tripleExprLabel* p_tripleExprLabel) : yacker::_Production("include") {
	m_GT_AMP = p_GT_AMP;
	m_tripleExprLabel = p_tripleExprLabel;
	makeArray(_members, &m_GT_AMP, &m_tripleExprLabel);
	trace();
    }
    ~include();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class annotation : public yacker::_Production {
private:
    _Base** _members[3];
    GT_DIVIDE_DIVIDE* m_GT_DIVIDE_DIVIDE;
    predicate* m_predicate;
    _O_Qiri_E_Or_Qliteral_E_C* m__O_Qiri_E_Or_Qliteral_E_C;
    virtual const char* getProductionName () { return "annotation"; }
public:
    annotation (GT_DIVIDE_DIVIDE* p_GT_DIVIDE_DIVIDE, predicate* p_predicate, _O_Qiri_E_Or_Qliteral_E_C* p__O_Qiri_E_Or_Qliteral_E_C) : yacker::_Production("annotation") {
	m_GT_DIVIDE_DIVIDE = p_GT_DIVIDE_DIVIDE;
	m_predicate = p_predicate;
	m__O_Qiri_E_Or_Qliteral_E_C = p__O_Qiri_E_Or_Qliteral_E_C;
	makeArray(_members, &m_GT_DIVIDE_DIVIDE, &m_predicate, &m__O_Qiri_E_Or_Qliteral_E_C);
	trace();
    }
    ~annotation();
    virtual size_t size () { return 3; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_Qiri_E_Or_Qliteral_E_C : public yacker::_GroupProduction {
public:
    _O_Qiri_E_Or_Qliteral_E_C () : yacker::_GroupProduction ("_O_Qiri_E_Or_Qliteral_E_C") {}
};
class _O_Qiri_E_Or_Qliteral_E_C_rule0 : public _O_Qiri_E_Or_Qliteral_E_C {
private:
    _Base** _members[1];
    iri* m_iri;
public:
    _O_Qiri_E_Or_Qliteral_E_C_rule0 (iri* p_iri) {
	m_iri = p_iri;
	makeArray(_members, &m_iri);
	trace();
    }
    ~_O_Qiri_E_Or_Qliteral_E_C_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_Qiri_E_Or_Qliteral_E_C_rule1 : public _O_Qiri_E_Or_Qliteral_E_C {
private:
    _Base** _members[1];
    literal* m_literal;
public:
    _O_Qiri_E_Or_Qliteral_E_C_rule1 (literal* p_literal) {
	m_literal = p_literal;
	makeArray(_members, &m_literal);
	trace();
    }
    ~_O_Qiri_E_Or_Qliteral_E_C_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class semanticActions : public yacker::_Production {
private:
    _Base** _members[1];
    yacker::_ProductionVector<codeDecl*>* m_codeDecls;
    virtual const char* getProductionName () { return "semanticActions"; }
public:
    semanticActions (yacker::_ProductionVector<codeDecl*>* p_codeDecls) : yacker::_Production("semanticActions") {
	m_codeDecls = p_codeDecls;
	makeArray(_members, &m_codeDecls);
	trace();
    }
    ~semanticActions();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class codeDecl : public yacker::_Production {
private:
    _Base** _members[3];
    GT_MODULO* m_GT_MODULO;
    iri* m_iri;
    _O_QCODE_E_Or_QGT_MODULO_E_C* m__O_QCODE_E_Or_QGT_MODULO_E_C;
    virtual const char* getProductionName () { return "codeDecl"; }
public:
    codeDecl (GT_MODULO* p_GT_MODULO, iri* p_iri, _O_QCODE_E_Or_QGT_MODULO_E_C* p__O_QCODE_E_Or_QGT_MODULO_E_C) : yacker::_Production("codeDecl") {
	m_GT_MODULO = p_GT_MODULO;
	m_iri = p_iri;
	m__O_QCODE_E_Or_QGT_MODULO_E_C = p__O_QCODE_E_Or_QGT_MODULO_E_C;
	makeArray(_members, &m_GT_MODULO, &m_iri, &m__O_QCODE_E_Or_QGT_MODULO_E_C);
	trace();
    }
    ~codeDecl();
    virtual size_t size () { return 3; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QCODE_E_Or_QGT_MODULO_E_C : public yacker::_GroupProduction {
public:
    _O_QCODE_E_Or_QGT_MODULO_E_C () : yacker::_GroupProduction ("_O_QCODE_E_Or_QGT_MODULO_E_C") {}
};
class _O_QCODE_E_Or_QGT_MODULO_E_C_rule0 : public _O_QCODE_E_Or_QGT_MODULO_E_C {
private:
    _Base** _members[1];
    CODE* m_CODE;
public:
    _O_QCODE_E_Or_QGT_MODULO_E_C_rule0 (CODE* p_CODE) {
	m_CODE = p_CODE;
	makeArray(_members, &m_CODE);
	trace();
    }
    ~_O_QCODE_E_Or_QGT_MODULO_E_C_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QCODE_E_Or_QGT_MODULO_E_C_rule1 : public _O_QCODE_E_Or_QGT_MODULO_E_C {
private:
    _Base** _members[1];
    GT_MODULO* m_GT_MODULO;
public:
    _O_QCODE_E_Or_QGT_MODULO_E_C_rule1 (GT_MODULO* p_GT_MODULO) {
	m_GT_MODULO = p_GT_MODULO;
	makeArray(_members, &m_GT_MODULO);
	trace();
    }
    ~_O_QCODE_E_Or_QGT_MODULO_E_C_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class literal : public yacker::_GroupProduction {
public:
    literal () : yacker::_GroupProduction ("literal") {}
};
class literal_rule0 : public literal {
private:
    _Base** _members[1];
    rdfLiteral* m_rdfLiteral;
public:
    literal_rule0 (rdfLiteral* p_rdfLiteral) {
	m_rdfLiteral = p_rdfLiteral;
	makeArray(_members, &m_rdfLiteral);
	trace();
    }
    ~literal_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class literal_rule1 : public literal {
private:
    _Base** _members[1];
    numericLiteral* m_numericLiteral;
public:
    literal_rule1 (numericLiteral* p_numericLiteral) {
	m_numericLiteral = p_numericLiteral;
	makeArray(_members, &m_numericLiteral);
	trace();
    }
    ~literal_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class literal_rule2 : public literal {
private:
    _Base** _members[1];
    booleanLiteral* m_booleanLiteral;
public:
    literal_rule2 (booleanLiteral* p_booleanLiteral) {
	m_booleanLiteral = p_booleanLiteral;
	makeArray(_members, &m_booleanLiteral);
	trace();
    }
    ~literal_rule2();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class predicate : public yacker::_GroupProduction {
public:
    predicate () : yacker::_GroupProduction ("predicate") {}
};
class predicate_rule0 : public predicate {
private:
    _Base** _members[1];
    iri* m_iri;
public:
    predicate_rule0 (iri* p_iri) {
	m_iri = p_iri;
	makeArray(_members, &m_iri);
	trace();
    }
    ~predicate_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class predicate_rule1 : public predicate {
private:
    _Base** _members[1];
    RDF_TYPE* m_RDF_TYPE;
public:
    predicate_rule1 (RDF_TYPE* p_RDF_TYPE) {
	m_RDF_TYPE = p_RDF_TYPE;
	makeArray(_members, &m_RDF_TYPE);
	trace();
    }
    ~predicate_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class datatype : public yacker::_Production {
private:
    _Base** _members[1];
    iri* m_iri;
    virtual const char* getProductionName () { return "datatype"; }
public:
    datatype (iri* p_iri) : yacker::_Production("datatype") {
	m_iri = p_iri;
	makeArray(_members, &m_iri);
	trace();
    }
    ~datatype();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeExprLabel : public yacker::_GroupProduction {
public:
    shapeExprLabel () : yacker::_GroupProduction ("shapeExprLabel") {}
};
class shapeExprLabel_rule0 : public shapeExprLabel {
private:
    _Base** _members[1];
    iri* m_iri;
public:
    shapeExprLabel_rule0 (iri* p_iri) {
	m_iri = p_iri;
	makeArray(_members, &m_iri);
	trace();
    }
    ~shapeExprLabel_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class shapeExprLabel_rule1 : public shapeExprLabel {
private:
    _Base** _members[1];
    blankNode* m_blankNode;
public:
    shapeExprLabel_rule1 (blankNode* p_blankNode) {
	m_blankNode = p_blankNode;
	makeArray(_members, &m_blankNode);
	trace();
    }
    ~shapeExprLabel_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class tripleExprLabel : public yacker::_GroupProduction {
public:
    tripleExprLabel () : yacker::_GroupProduction ("tripleExprLabel") {}
};
class tripleExprLabel_rule0 : public tripleExprLabel {
private:
    _Base** _members[1];
    iri* m_iri;
public:
    tripleExprLabel_rule0 (iri* p_iri) {
	m_iri = p_iri;
	makeArray(_members, &m_iri);
	trace();
    }
    ~tripleExprLabel_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class tripleExprLabel_rule1 : public tripleExprLabel {
private:
    _Base** _members[1];
    blankNode* m_blankNode;
public:
    tripleExprLabel_rule1 (blankNode* p_blankNode) {
	m_blankNode = p_blankNode;
	makeArray(_members, &m_blankNode);
	trace();
    }
    ~tripleExprLabel_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class rawNumeric : public yacker::_GroupProduction {
public:
    rawNumeric () : yacker::_GroupProduction ("rawNumeric") {}
};
class rawNumeric_rule0 : public rawNumeric {
private:
    _Base** _members[1];
    INTEGER* m_INTEGER;
public:
    rawNumeric_rule0 (INTEGER* p_INTEGER) {
	m_INTEGER = p_INTEGER;
	makeArray(_members, &m_INTEGER);
	trace();
    }
    ~rawNumeric_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class rawNumeric_rule1 : public rawNumeric {
private:
    _Base** _members[1];
    DECIMAL* m_DECIMAL;
public:
    rawNumeric_rule1 (DECIMAL* p_DECIMAL) {
	m_DECIMAL = p_DECIMAL;
	makeArray(_members, &m_DECIMAL);
	trace();
    }
    ~rawNumeric_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class rawNumeric_rule2 : public rawNumeric {
private:
    _Base** _members[1];
    DOUBLE* m_DOUBLE;
public:
    rawNumeric_rule2 (DOUBLE* p_DOUBLE) {
	m_DOUBLE = p_DOUBLE;
	makeArray(_members, &m_DOUBLE);
	trace();
    }
    ~rawNumeric_rule2();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class numericLiteral : public yacker::_Production {
private:
    _Base** _members[1];
    rawNumeric* m_rawNumeric;
    virtual const char* getProductionName () { return "numericLiteral"; }
public:
    numericLiteral (rawNumeric* p_rawNumeric) : yacker::_Production("numericLiteral") {
	m_rawNumeric = p_rawNumeric;
	makeArray(_members, &m_rawNumeric);
	trace();
    }
    ~numericLiteral();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class rdfLiteral : public yacker::_GroupProduction {
public:
    rdfLiteral () : yacker::_GroupProduction ("rdfLiteral") {}
};
class rdfLiteral_rule0 : public rdfLiteral {
private:
    _Base** _members[1];
    langString* m_langString;
public:
    rdfLiteral_rule0 (langString* p_langString) {
	m_langString = p_langString;
	makeArray(_members, &m_langString);
	trace();
    }
    ~rdfLiteral_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class rdfLiteral_rule1 : public rdfLiteral {
private:
    _Base** _members[2];
    string* m_string;
    _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt* m__Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt;
public:
    rdfLiteral_rule1 (string* p_string, _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt* p__Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt) {
	m_string = p_string;
	m__Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt = p__Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt;
	makeArray(_members, &m_string, &m__Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt);
	trace();
    }
    ~rdfLiteral_rule1();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _O_QGT_DTYPE_E_S_Qdatatype_E_C : public yacker::_Production {
private:
    _Base** _members[2];
    GT_DTYPE* m_GT_DTYPE;
    datatype* m_datatype;
    virtual const char* getProductionName () { return "_O_QGT_DTYPE_E_S_Qdatatype_E_C"; }
public:
    _O_QGT_DTYPE_E_S_Qdatatype_E_C (GT_DTYPE* p_GT_DTYPE, datatype* p_datatype) : yacker::_Production("_O_QGT_DTYPE_E_S_Qdatatype_E_C") {
	m_GT_DTYPE = p_GT_DTYPE;
	m_datatype = p_datatype;
	makeArray(_members, &m_GT_DTYPE, &m_datatype);
	trace();
    }
    ~_O_QGT_DTYPE_E_S_Qdatatype_E_C();
    virtual size_t size () { return 2; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt : public yacker::_GroupProduction {
public:
    _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt () : yacker::_GroupProduction ("_Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt") {}
};
class _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt_rule0 : public _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt {
public:
    _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt_rule0 () {
	trace();
    }
    ~_Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt_rule0();
    virtual size_t size () { return 0; }
    virtual _Base* operator [] (size_t /*i*/) { return NULL; }
    virtual _Base* getElement (size_t /*i*/) { return NULL; }
};
class _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt_rule1 : public _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt {
private:
    _Base** _members[1];
    _O_QGT_DTYPE_E_S_Qdatatype_E_C* m__O_QGT_DTYPE_E_S_Qdatatype_E_C;
public:
    _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt_rule1 (_O_QGT_DTYPE_E_S_Qdatatype_E_C* p__O_QGT_DTYPE_E_S_Qdatatype_E_C) {
	m__O_QGT_DTYPE_E_S_Qdatatype_E_C = p__O_QGT_DTYPE_E_S_Qdatatype_E_C;
	makeArray(_members, &m__O_QGT_DTYPE_E_S_Qdatatype_E_C);
	trace();
    }
    ~_Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class booleanLiteral : public yacker::_GroupProduction {
public:
    booleanLiteral () : yacker::_GroupProduction ("booleanLiteral") {}
};
class booleanLiteral_rule0 : public booleanLiteral {
private:
    _Base** _members[1];
    IT_true* m_IT_true;
public:
    booleanLiteral_rule0 (IT_true* p_IT_true) {
	m_IT_true = p_IT_true;
	makeArray(_members, &m_IT_true);
	trace();
    }
    ~booleanLiteral_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class booleanLiteral_rule1 : public booleanLiteral {
private:
    _Base** _members[1];
    IT_false* m_IT_false;
public:
    booleanLiteral_rule1 (IT_false* p_IT_false) {
	m_IT_false = p_IT_false;
	makeArray(_members, &m_IT_false);
	trace();
    }
    ~booleanLiteral_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class string : public yacker::_GroupProduction {
public:
    string () : yacker::_GroupProduction ("string") {}
};
class string_rule0 : public string {
private:
    _Base** _members[1];
    STRING_LITERAL1* m_STRING_LITERAL1;
public:
    string_rule0 (STRING_LITERAL1* p_STRING_LITERAL1) {
	m_STRING_LITERAL1 = p_STRING_LITERAL1;
	makeArray(_members, &m_STRING_LITERAL1);
	trace();
    }
    ~string_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class string_rule1 : public string {
private:
    _Base** _members[1];
    STRING_LITERAL_LONG1* m_STRING_LITERAL_LONG1;
public:
    string_rule1 (STRING_LITERAL_LONG1* p_STRING_LITERAL_LONG1) {
	m_STRING_LITERAL_LONG1 = p_STRING_LITERAL_LONG1;
	makeArray(_members, &m_STRING_LITERAL_LONG1);
	trace();
    }
    ~string_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class string_rule2 : public string {
private:
    _Base** _members[1];
    STRING_LITERAL2* m_STRING_LITERAL2;
public:
    string_rule2 (STRING_LITERAL2* p_STRING_LITERAL2) {
	m_STRING_LITERAL2 = p_STRING_LITERAL2;
	makeArray(_members, &m_STRING_LITERAL2);
	trace();
    }
    ~string_rule2();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class string_rule3 : public string {
private:
    _Base** _members[1];
    STRING_LITERAL_LONG2* m_STRING_LITERAL_LONG2;
public:
    string_rule3 (STRING_LITERAL_LONG2* p_STRING_LITERAL_LONG2) {
	m_STRING_LITERAL_LONG2 = p_STRING_LITERAL_LONG2;
	makeArray(_members, &m_STRING_LITERAL_LONG2);
	trace();
    }
    ~string_rule3();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class langString : public yacker::_GroupProduction {
public:
    langString () : yacker::_GroupProduction ("langString") {}
};
class langString_rule0 : public langString {
private:
    _Base** _members[1];
    LANG_STRING_LITERAL1* m_LANG_STRING_LITERAL1;
public:
    langString_rule0 (LANG_STRING_LITERAL1* p_LANG_STRING_LITERAL1) {
	m_LANG_STRING_LITERAL1 = p_LANG_STRING_LITERAL1;
	makeArray(_members, &m_LANG_STRING_LITERAL1);
	trace();
    }
    ~langString_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class langString_rule1 : public langString {
private:
    _Base** _members[1];
    LANG_STRING_LITERAL_LONG1* m_LANG_STRING_LITERAL_LONG1;
public:
    langString_rule1 (LANG_STRING_LITERAL_LONG1* p_LANG_STRING_LITERAL_LONG1) {
	m_LANG_STRING_LITERAL_LONG1 = p_LANG_STRING_LITERAL_LONG1;
	makeArray(_members, &m_LANG_STRING_LITERAL_LONG1);
	trace();
    }
    ~langString_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class langString_rule2 : public langString {
private:
    _Base** _members[1];
    LANG_STRING_LITERAL2* m_LANG_STRING_LITERAL2;
public:
    langString_rule2 (LANG_STRING_LITERAL2* p_LANG_STRING_LITERAL2) {
	m_LANG_STRING_LITERAL2 = p_LANG_STRING_LITERAL2;
	makeArray(_members, &m_LANG_STRING_LITERAL2);
	trace();
    }
    ~langString_rule2();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class langString_rule3 : public langString {
private:
    _Base** _members[1];
    LANG_STRING_LITERAL_LONG2* m_LANG_STRING_LITERAL_LONG2;
public:
    langString_rule3 (LANG_STRING_LITERAL_LONG2* p_LANG_STRING_LITERAL_LONG2) {
	m_LANG_STRING_LITERAL_LONG2 = p_LANG_STRING_LITERAL_LONG2;
	makeArray(_members, &m_LANG_STRING_LITERAL_LONG2);
	trace();
    }
    ~langString_rule3();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class iri : public yacker::_GroupProduction {
public:
    iri () : yacker::_GroupProduction ("iri") {}
};
class iri_rule0 : public iri {
private:
    _Base** _members[1];
    IRIREF* m_IRIREF;
public:
    iri_rule0 (IRIREF* p_IRIREF) {
	m_IRIREF = p_IRIREF;
	makeArray(_members, &m_IRIREF);
	trace();
    }
    ~iri_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class iri_rule1 : public iri {
private:
    _Base** _members[1];
    prefixedName* m_prefixedName;
public:
    iri_rule1 (prefixedName* p_prefixedName) {
	m_prefixedName = p_prefixedName;
	makeArray(_members, &m_prefixedName);
	trace();
    }
    ~iri_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class prefixedName : public yacker::_GroupProduction {
public:
    prefixedName () : yacker::_GroupProduction ("prefixedName") {}
};
class prefixedName_rule0 : public prefixedName {
private:
    _Base** _members[1];
    PNAME_LN* m_PNAME_LN;
public:
    prefixedName_rule0 (PNAME_LN* p_PNAME_LN) {
	m_PNAME_LN = p_PNAME_LN;
	makeArray(_members, &m_PNAME_LN);
	trace();
    }
    ~prefixedName_rule0();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class prefixedName_rule1 : public prefixedName {
private:
    _Base** _members[1];
    PNAME_NS* m_PNAME_NS;
public:
    prefixedName_rule1 (PNAME_NS* p_PNAME_NS) {
	m_PNAME_NS = p_PNAME_NS;
	makeArray(_members, &m_PNAME_NS);
	trace();
    }
    ~prefixedName_rule1();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};
class blankNode : public yacker::_Production {
private:
    _Base** _members[1];
    BLANK_NODE_LABEL* m_BLANK_NODE_LABEL;
    virtual const char* getProductionName () { return "blankNode"; }
public:
    blankNode (BLANK_NODE_LABEL* p_BLANK_NODE_LABEL) : yacker::_Production("blankNode") {
	m_BLANK_NODE_LABEL = p_BLANK_NODE_LABEL;
	makeArray(_members, &m_BLANK_NODE_LABEL);
	trace();
    }
    ~blankNode();
    virtual size_t size () { return 1; }
    virtual _Base* operator [] (size_t i) { return *_members[i]; }
    virtual _Base* getElement (size_t i) { return *_members[i]; }
};

class IT_BASE : public yacker::_Token {
public:
    IT_BASE (const char* text) : yacker::_Token("BASE", text) { trace(); }
};
class IT_PREFIX : public yacker::_Token {
public:
    IT_PREFIX (const char* text) : yacker::_Token("PREFIX", text) { trace(); }
};
class IT_IMPORT : public yacker::_Token {
public:
    IT_IMPORT (const char* text) : yacker::_Token("IMPORT", text) { trace(); }
};
class IT_START : public yacker::_Token {
public:
    IT_START (const char* text) : yacker::_Token("START", text) { trace(); }
};
class GT_EQUAL : public yacker::_Token {
public:
    GT_EQUAL (const char* text) : yacker::_Token("EQUAL", text) { trace(); }
};
class IT_EXTERNAL : public yacker::_Token {
public:
    IT_EXTERNAL (const char* text) : yacker::_Token("EXTERNAL", text) { trace(); }
};
class IT_OR : public yacker::_Token {
public:
    IT_OR (const char* text) : yacker::_Token("OR", text) { trace(); }
};
class IT_AND : public yacker::_Token {
public:
    IT_AND (const char* text) : yacker::_Token("AND", text) { trace(); }
};
class IT_NOT : public yacker::_Token {
public:
    IT_NOT (const char* text) : yacker::_Token("NOT", text) { trace(); }
};
class GT_LPAREN : public yacker::_Token {
public:
    GT_LPAREN (const char* text) : yacker::_Token("LPAREN", text) { trace(); }
};
class GT_RPAREN : public yacker::_Token {
public:
    GT_RPAREN (const char* text) : yacker::_Token("RPAREN", text) { trace(); }
};
class GT_DOT : public yacker::_Token {
public:
    GT_DOT (const char* text) : yacker::_Token("DOT", text) { trace(); }
};
class GT_AT : public yacker::_Token {
public:
    GT_AT (const char* text) : yacker::_Token("AT", text) { trace(); }
};
class IT_LITERAL : public yacker::_Token {
public:
    IT_LITERAL (const char* text) : yacker::_Token("LITERAL", text) { trace(); }
};
class IT_IRI : public yacker::_Token {
public:
    IT_IRI (const char* text) : yacker::_Token("IRI", text) { trace(); }
};
class IT_BNODE : public yacker::_Token {
public:
    IT_BNODE (const char* text) : yacker::_Token("BNODE", text) { trace(); }
};
class IT_NONLITERAL : public yacker::_Token {
public:
    IT_NONLITERAL (const char* text) : yacker::_Token("NONLITERAL", text) { trace(); }
};
class IT_LENGTH : public yacker::_Token {
public:
    IT_LENGTH (const char* text) : yacker::_Token("LENGTH", text) { trace(); }
};
class IT_MINLENGTH : public yacker::_Token {
public:
    IT_MINLENGTH (const char* text) : yacker::_Token("MINLENGTH", text) { trace(); }
};
class IT_MAXLENGTH : public yacker::_Token {
public:
    IT_MAXLENGTH (const char* text) : yacker::_Token("MAXLENGTH", text) { trace(); }
};
class IT_MININCLUSIVE : public yacker::_Token {
public:
    IT_MININCLUSIVE (const char* text) : yacker::_Token("MININCLUSIVE", text) { trace(); }
};
class IT_MINEXCLUSIVE : public yacker::_Token {
public:
    IT_MINEXCLUSIVE (const char* text) : yacker::_Token("MINEXCLUSIVE", text) { trace(); }
};
class IT_MAXINCLUSIVE : public yacker::_Token {
public:
    IT_MAXINCLUSIVE (const char* text) : yacker::_Token("MAXINCLUSIVE", text) { trace(); }
};
class IT_MAXEXCLUSIVE : public yacker::_Token {
public:
    IT_MAXEXCLUSIVE (const char* text) : yacker::_Token("MAXEXCLUSIVE", text) { trace(); }
};
class IT_TOTALDIGITS : public yacker::_Token {
public:
    IT_TOTALDIGITS (const char* text) : yacker::_Token("TOTALDIGITS", text) { trace(); }
};
class IT_FRACTIONDIGITS : public yacker::_Token {
public:
    IT_FRACTIONDIGITS (const char* text) : yacker::_Token("FRACTIONDIGITS", text) { trace(); }
};
class GT_LCURLEY : public yacker::_Token {
public:
    GT_LCURLEY (const char* text) : yacker::_Token("LCURLEY", text) { trace(); }
};
class GT_RCURLEY : public yacker::_Token {
public:
    GT_RCURLEY (const char* text) : yacker::_Token("RCURLEY", text) { trace(); }
};
class IT_CLOSED : public yacker::_Token {
public:
    IT_CLOSED (const char* text) : yacker::_Token("CLOSED", text) { trace(); }
};
class IT_EXTRA : public yacker::_Token {
public:
    IT_EXTRA (const char* text) : yacker::_Token("EXTRA", text) { trace(); }
};
class GT_PIPE : public yacker::_Token {
public:
    GT_PIPE (const char* text) : yacker::_Token("PIPE", text) { trace(); }
};
class GT_SEMI : public yacker::_Token {
public:
    GT_SEMI (const char* text) : yacker::_Token("SEMI", text) { trace(); }
};
class GT_DOLLAR : public yacker::_Token {
public:
    GT_DOLLAR (const char* text) : yacker::_Token("DOLLAR", text) { trace(); }
};
class GT_TIMES : public yacker::_Token {
public:
    GT_TIMES (const char* text) : yacker::_Token("TIMES", text) { trace(); }
};
class GT_PLUS : public yacker::_Token {
public:
    GT_PLUS (const char* text) : yacker::_Token("PLUS", text) { trace(); }
};
class GT_OPT : public yacker::_Token {
public:
    GT_OPT (const char* text) : yacker::_Token("OPT", text) { trace(); }
};
class GT_CARROT : public yacker::_Token {
public:
    GT_CARROT (const char* text) : yacker::_Token("CARROT", text) { trace(); }
};
class GT_LBRACKET : public yacker::_Token {
public:
    GT_LBRACKET (const char* text) : yacker::_Token("LBRACKET", text) { trace(); }
};
class GT_RBRACKET : public yacker::_Token {
public:
    GT_RBRACKET (const char* text) : yacker::_Token("RBRACKET", text) { trace(); }
};
class GT_KINDA : public yacker::_Token {
public:
    GT_KINDA (const char* text) : yacker::_Token("KINDA", text) { trace(); }
};
class GT_MINUS : public yacker::_Token {
public:
    GT_MINUS (const char* text) : yacker::_Token("MINUS", text) { trace(); }
};
class GT_AMP : public yacker::_Token {
public:
    GT_AMP (const char* text) : yacker::_Token("AMP", text) { trace(); }
};
class GT_DIVIDE_DIVIDE : public yacker::_Token {
public:
    GT_DIVIDE_DIVIDE (const char* text) : yacker::_Token("DIVIDE_DIVIDE", text) { trace(); }
};
class GT_MODULO : public yacker::_Token {
public:
    GT_MODULO (const char* text) : yacker::_Token("MODULO", text) { trace(); }
};
class GT_DTYPE : public yacker::_Token {
public:
    GT_DTYPE (const char* text) : yacker::_Token("DTYPE", text) { trace(); }
};
class IT_true : public yacker::_Token {
public:
    IT_true (const char* text) : yacker::_Token("true", text) { trace(); }
};
class IT_false : public yacker::_Token {
public:
    IT_false (const char* text) : yacker::_Token("false", text) { trace(); }
};
class CODE : public yacker::_Terminal {
public:
    CODE(const char* p_CODE) : yacker::_Terminal("CODE", p_CODE) { trace(); }
};
class REPEAT_RANGE : public yacker::_Terminal {
public:
    REPEAT_RANGE(const char* p_REPEAT_RANGE) : yacker::_Terminal("REPEAT_RANGE", p_REPEAT_RANGE) { trace(); }
};
class RDF_TYPE : public yacker::_Terminal {
public:
    RDF_TYPE(const char* p_RDF_TYPE) : yacker::_Terminal("RDF_TYPE", p_RDF_TYPE) { trace(); }
};
class IRIREF : public yacker::_Terminal {
public:
    IRIREF(const char* p_IRIREF) : yacker::_Terminal("IRIREF", p_IRIREF) { trace(); }
};
class PNAME_NS : public yacker::_Terminal {
public:
    PNAME_NS(const char* p_PNAME_NS) : yacker::_Terminal("PNAME_NS", p_PNAME_NS) { trace(); }
};
class PNAME_LN : public yacker::_Terminal {
public:
    PNAME_LN(const char* p_PNAME_LN) : yacker::_Terminal("PNAME_LN", p_PNAME_LN) { trace(); }
};
class ATPNAME_NS : public yacker::_Terminal {
public:
    ATPNAME_NS(const char* p_ATPNAME_NS) : yacker::_Terminal("ATPNAME_NS", p_ATPNAME_NS) { trace(); }
};
class ATPNAME_LN : public yacker::_Terminal {
public:
    ATPNAME_LN(const char* p_ATPNAME_LN) : yacker::_Terminal("ATPNAME_LN", p_ATPNAME_LN) { trace(); }
};
class REGEXP : public yacker::_Terminal {
public:
    REGEXP(const char* p_REGEXP) : yacker::_Terminal("REGEXP", p_REGEXP) { trace(); }
};
class BLANK_NODE_LABEL : public yacker::_Terminal {
public:
    BLANK_NODE_LABEL(const char* p_BLANK_NODE_LABEL) : yacker::_Terminal("BLANK_NODE_LABEL", p_BLANK_NODE_LABEL) { trace(); }
};
class LANGTAG : public yacker::_Terminal {
public:
    LANGTAG(const char* p_LANGTAG) : yacker::_Terminal("LANGTAG", p_LANGTAG) { trace(); }
};
class INTEGER : public yacker::_Terminal {
public:
    INTEGER(const char* p_INTEGER) : yacker::_Terminal("INTEGER", p_INTEGER) { trace(); }
};
class DECIMAL : public yacker::_Terminal {
public:
    DECIMAL(const char* p_DECIMAL) : yacker::_Terminal("DECIMAL", p_DECIMAL) { trace(); }
};
class DOUBLE : public yacker::_Terminal {
public:
    DOUBLE(const char* p_DOUBLE) : yacker::_Terminal("DOUBLE", p_DOUBLE) { trace(); }
};
class STRING_LITERAL1 : public yacker::_Terminal {
public:
    STRING_LITERAL1(const char* p_STRING_LITERAL1) : yacker::_Terminal("STRING_LITERAL1", p_STRING_LITERAL1) { trace(); }
};
class STRING_LITERAL2 : public yacker::_Terminal {
public:
    STRING_LITERAL2(const char* p_STRING_LITERAL2) : yacker::_Terminal("STRING_LITERAL2", p_STRING_LITERAL2) { trace(); }
};
class STRING_LITERAL_LONG1 : public yacker::_Terminal {
public:
    STRING_LITERAL_LONG1(const char* p_STRING_LITERAL_LONG1) : yacker::_Terminal("STRING_LITERAL_LONG1", p_STRING_LITERAL_LONG1) { trace(); }
};
class STRING_LITERAL_LONG2 : public yacker::_Terminal {
public:
    STRING_LITERAL_LONG2(const char* p_STRING_LITERAL_LONG2) : yacker::_Terminal("STRING_LITERAL_LONG2", p_STRING_LITERAL_LONG2) { trace(); }
};
class LANG_STRING_LITERAL1 : public yacker::_Terminal {
public:
    LANG_STRING_LITERAL1(const char* p_LANG_STRING_LITERAL1) : yacker::_Terminal("LANG_STRING_LITERAL1", p_LANG_STRING_LITERAL1) { trace(); }
};
class LANG_STRING_LITERAL2 : public yacker::_Terminal {
public:
    LANG_STRING_LITERAL2(const char* p_LANG_STRING_LITERAL2) : yacker::_Terminal("LANG_STRING_LITERAL2", p_LANG_STRING_LITERAL2) { trace(); }
};
class LANG_STRING_LITERAL_LONG1 : public yacker::_Terminal {
public:
    LANG_STRING_LITERAL_LONG1(const char* p_LANG_STRING_LITERAL_LONG1) : yacker::_Terminal("LANG_STRING_LITERAL_LONG1", p_LANG_STRING_LITERAL_LONG1) { trace(); }
};
class LANG_STRING_LITERAL_LONG2 : public yacker::_Terminal {
public:
    LANG_STRING_LITERAL_LONG2(const char* p_LANG_STRING_LITERAL_LONG2) : yacker::_Terminal("LANG_STRING_LITERAL_LONG2", p_LANG_STRING_LITERAL_LONG2) { trace(); }
};

}
class ProgramFlowException : public std::exception  {
private:
    const char* msg;
public:
    ProgramFlowException(const char* msg) : msg(msg) { }

    // This declaration is not useless:
    // http://gcc.gnu.org/onlinedocs/gcc-3.0.2/gcc_6.html#SEC118
    virtual ~ProgramFlowException() throw() { }

    // See comment in eh_exception.cc.
    virtual const char* what () { return msg; }
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
    shexDoc* root;
};

} // namespace shexParseToXml

//%}
} // bison 2

 /*** BEGIN shexParseToXml - Change the grammar's tokens below ***/

%union {
    /* Terminals */
    IT_BASE* p_IT_BASE;
    IT_PREFIX* p_IT_PREFIX;
    IT_IMPORT* p_IT_IMPORT;
    IT_START* p_IT_START;
    GT_EQUAL* p_GT_EQUAL;
    IT_EXTERNAL* p_IT_EXTERNAL;
    IT_OR* p_IT_OR;
    IT_AND* p_IT_AND;
    IT_NOT* p_IT_NOT;
    GT_LPAREN* p_GT_LPAREN;
    GT_RPAREN* p_GT_RPAREN;
    GT_DOT* p_GT_DOT;
    GT_AT* p_GT_AT;
    IT_LITERAL* p_IT_LITERAL;
    IT_IRI* p_IT_IRI;
    IT_BNODE* p_IT_BNODE;
    IT_NONLITERAL* p_IT_NONLITERAL;
    IT_LENGTH* p_IT_LENGTH;
    IT_MINLENGTH* p_IT_MINLENGTH;
    IT_MAXLENGTH* p_IT_MAXLENGTH;
    IT_MININCLUSIVE* p_IT_MININCLUSIVE;
    IT_MINEXCLUSIVE* p_IT_MINEXCLUSIVE;
    IT_MAXINCLUSIVE* p_IT_MAXINCLUSIVE;
    IT_MAXEXCLUSIVE* p_IT_MAXEXCLUSIVE;
    IT_TOTALDIGITS* p_IT_TOTALDIGITS;
    IT_FRACTIONDIGITS* p_IT_FRACTIONDIGITS;
    GT_LCURLEY* p_GT_LCURLEY;
    GT_RCURLEY* p_GT_RCURLEY;
    IT_CLOSED* p_IT_CLOSED;
    IT_EXTRA* p_IT_EXTRA;
    GT_PIPE* p_GT_PIPE;
    GT_SEMI* p_GT_SEMI;
    GT_DOLLAR* p_GT_DOLLAR;
    GT_TIMES* p_GT_TIMES;
    GT_PLUS* p_GT_PLUS;
    GT_OPT* p_GT_OPT;
    GT_CARROT* p_GT_CARROT;
    GT_LBRACKET* p_GT_LBRACKET;
    GT_RBRACKET* p_GT_RBRACKET;
    GT_KINDA* p_GT_KINDA;
    GT_MINUS* p_GT_MINUS;
    GT_AMP* p_GT_AMP;
    GT_DIVIDE_DIVIDE* p_GT_DIVIDE_DIVIDE;
    GT_MODULO* p_GT_MODULO;
    GT_DTYPE* p_GT_DTYPE;
    IT_true* p_IT_true;
    IT_false* p_IT_false;
    CODE* p_CODE;
    REPEAT_RANGE* p_REPEAT_RANGE;
    RDF_TYPE* p_RDF_TYPE;
    IRIREF* p_IRIREF;
    PNAME_NS* p_PNAME_NS;
    PNAME_LN* p_PNAME_LN;
    ATPNAME_NS* p_ATPNAME_NS;
    ATPNAME_LN* p_ATPNAME_LN;
    REGEXP* p_REGEXP;
    BLANK_NODE_LABEL* p_BLANK_NODE_LABEL;
    LANGTAG* p_LANGTAG;
    INTEGER* p_INTEGER;
    DECIMAL* p_DECIMAL;
    DOUBLE* p_DOUBLE;
    STRING_LITERAL1* p_STRING_LITERAL1;
    STRING_LITERAL2* p_STRING_LITERAL2;
    STRING_LITERAL_LONG1* p_STRING_LITERAL_LONG1;
    STRING_LITERAL_LONG2* p_STRING_LITERAL_LONG2;
    LANG_STRING_LITERAL1* p_LANG_STRING_LITERAL1;
    LANG_STRING_LITERAL2* p_LANG_STRING_LITERAL2;
    LANG_STRING_LITERAL_LONG1* p_LANG_STRING_LITERAL_LONG1;
    LANG_STRING_LITERAL_LONG2* p_LANG_STRING_LITERAL_LONG2;

    /* Productions */
    shexDoc* p_shexDoc;
    yacker::_ProductionVector<directive*>* p_directives;    _O_QnotStartAction_E_Or_QstartActions_E_C* p__O_QnotStartAction_E_Or_QstartActions_E_C;
    yacker::_ProductionVector<statement*>* p_statements;    _O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C* p__O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C;
    _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt* p__Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt;
    directive* p_directive;
    baseDecl* p_baseDecl;
    prefixDecl* p_prefixDecl;
    importDecl* p_importDecl;
    notStartAction* p_notStartAction;
    start* p_start;
    startActions* p_startActions;
    yacker::_ProductionVector<codeDecl*>* p_codeDecls;    statement* p_statement;
    shapeExprDecl* p_shapeExprDecl;
    _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C* p__O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C;
    shapeExpression* p_shapeExpression;
    inlineShapeExpression* p_inlineShapeExpression;
    shapeOr* p_shapeOr;
    _O_QIT_OR_E_S_QshapeAnd_E_C* p__O_QIT_OR_E_S_QshapeAnd_E_C;
    yacker::_ProductionVector<_O_QIT_OR_E_S_QshapeAnd_E_C*>* p__O_QIT_OR_E_S_QshapeAnd_E_Cs;    inlineShapeOr* p_inlineShapeOr;
    _O_QIT_OR_E_S_QinlineShapeAnd_E_C* p__O_QIT_OR_E_S_QinlineShapeAnd_E_C;
    yacker::_ProductionVector<_O_QIT_OR_E_S_QinlineShapeAnd_E_C*>* p__O_QIT_OR_E_S_QinlineShapeAnd_E_Cs;    shapeAnd* p_shapeAnd;
    _O_QIT_AND_E_S_QshapeNot_E_C* p__O_QIT_AND_E_S_QshapeNot_E_C;
    yacker::_ProductionVector<_O_QIT_AND_E_S_QshapeNot_E_C*>* p__O_QIT_AND_E_S_QshapeNot_E_Cs;    inlineShapeAnd* p_inlineShapeAnd;
    _O_QIT_AND_E_S_QinlineShapeNot_E_C* p__O_QIT_AND_E_S_QinlineShapeNot_E_C;
    yacker::_ProductionVector<_O_QIT_AND_E_S_QinlineShapeNot_E_C*>* p__O_QIT_AND_E_S_QinlineShapeNot_E_Cs;    shapeNot* p_shapeNot;
    _QIT_NOT_E_Opt* p__QIT_NOT_E_Opt;
    inlineShapeNot* p_inlineShapeNot;
    shapeAtom* p_shapeAtom;
    _QshapeOrRef_E_Opt* p__QshapeOrRef_E_Opt;
    _QnonLitNodeConstraint_E_Opt* p__QnonLitNodeConstraint_E_Opt;
    inlineShapeAtom* p_inlineShapeAtom;
    _QinlineShapeOrRef_E_Opt* p__QinlineShapeOrRef_E_Opt;
    shapeOrRef* p_shapeOrRef;
    inlineShapeOrRef* p_inlineShapeOrRef;
    shapeRef* p_shapeRef;
    litNodeConstraint* p_litNodeConstraint;
    yacker::_ProductionVector<xsFacet*>* p_xsFacets;    yacker::_ProductionVector<numericFacet*>* p_numericFacets;    nonLitNodeConstraint* p_nonLitNodeConstraint;
    yacker::_ProductionVector<stringFacet*>* p_stringFacets;    nonLiteralKind* p_nonLiteralKind;
    xsFacet* p_xsFacet;
    stringFacet* p_stringFacet;
    stringLength* p_stringLength;
    numericFacet* p_numericFacet;
    numericRange* p_numericRange;
    numericLength* p_numericLength;
    shapeDefinition* p_shapeDefinition;
    yacker::_ProductionVector<annotation*>* p_annotations;    inlineShapeDefinition* p_inlineShapeDefinition;
    _QtripleExpression_E_Opt* p__QtripleExpression_E_Opt;
    shapeQualifiers* p_shapeQualifiers;
    _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C* p__O_QextraPropertySet_E_Or_QIT_CLOSED_E_C;
    yacker::_ProductionVector<_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C*>* p__O_QextraPropertySet_E_Or_QIT_CLOSED_E_Cs;    extraPropertySet* p_extraPropertySet;
    yacker::_ProductionVector<predicate*>* p_predicates;    tripleExpression* p_tripleExpression;
    oneOfTripleExpr* p_oneOfTripleExpr;
    multiElementOneOf* p_multiElementOneOf;
    _O_QGT_PIPE_E_S_QgroupTripleExpr_E_C* p__O_QGT_PIPE_E_S_QgroupTripleExpr_E_C;
    yacker::_ProductionVector<_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C*>* p__O_QGT_PIPE_E_S_QgroupTripleExpr_E_Cs;    groupTripleExpr* p_groupTripleExpr;
    singleElementGroup* p_singleElementGroup;
    _QGT_SEMI_E_Opt* p__QGT_SEMI_E_Opt;
    multiElementGroup* p_multiElementGroup;
    _O_QGT_SEMI_E_S_QunaryTripleExpr_E_C* p__O_QGT_SEMI_E_S_QunaryTripleExpr_E_C;
    yacker::_ProductionVector<_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C*>* p__O_QGT_SEMI_E_S_QunaryTripleExpr_E_Cs;    unaryTripleExpr* p_unaryTripleExpr;
    _O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C* p__O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C;
    _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt* p__Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt;
    _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C* p__O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C;
    bracketedTripleExpr* p_bracketedTripleExpr;
    _Qcardinality_E_Opt* p__Qcardinality_E_Opt;
    tripleConstraint* p_tripleConstraint;
    _QsenseFlags_E_Opt* p__QsenseFlags_E_Opt;
    cardinality* p_cardinality;
    senseFlags* p_senseFlags;
    valueSet* p_valueSet;
    yacker::_ProductionVector<valueSetValue*>* p_valueSetValues;    valueSetValue* p_valueSetValue;
    yacker::_ProductionVector<iriExclusion*>* p_iriExclusions;    yacker::_ProductionVector<literalExclusion*>* p_literalExclusions;    yacker::_ProductionVector<languageExclusion*>* p_languageExclusions;    _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C* p__O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C;
    iriRange* p_iriRange;
    _O_QGT_KINDA_E_S_QiriExclusion_E_Star_C* p__O_QGT_KINDA_E_S_QiriExclusion_E_Star_C;
    _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt* p__Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt;
    iriExclusion* p_iriExclusion;
    _QGT_KINDA_E_Opt* p__QGT_KINDA_E_Opt;
    literalRange* p_literalRange;
    _O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C* p__O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C;
    _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt* p__Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt;
    literalExclusion* p_literalExclusion;
    languageRange* p_languageRange;
    _O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C* p__O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C;
    _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt* p__Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt;
    languageExclusion* p_languageExclusion;
    include* p_include;
    annotation* p_annotation;
    _O_Qiri_E_Or_Qliteral_E_C* p__O_Qiri_E_Or_Qliteral_E_C;
    semanticActions* p_semanticActions;
    codeDecl* p_codeDecl;
    _O_QCODE_E_Or_QGT_MODULO_E_C* p__O_QCODE_E_Or_QGT_MODULO_E_C;
    literal* p_literal;
    predicate* p_predicate;
    datatype* p_datatype;
    shapeExprLabel* p_shapeExprLabel;
    tripleExprLabel* p_tripleExprLabel;
    rawNumeric* p_rawNumeric;
    numericLiteral* p_numericLiteral;
    rdfLiteral* p_rdfLiteral;
    _O_QGT_DTYPE_E_S_Qdatatype_E_C* p__O_QGT_DTYPE_E_S_Qdatatype_E_C;
    _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt* p__Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt;
    booleanLiteral* p_booleanLiteral;
    string* p_string;
    langString* p_langString;
    iri* p_iri;
    prefixedName* p_prefixedName;
    blankNode* p_blankNode;

}

%{
#include "shex-parse-to-XMLScanner.hh"
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
%type <p_shexDoc> shexDoc
%type <p_directives> _Qdirective_E_Star
%type <p__O_QnotStartAction_E_Or_QstartActions_E_C> _O_QnotStartAction_E_Or_QstartActions_E_C
%type <p_statements> _Qstatement_E_Star
%type <p__O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C> _O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C
%type <p__Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt> _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt
%type <p_directive> directive
%type <p_baseDecl> baseDecl
%type <p_prefixDecl> prefixDecl
%type <p_importDecl> importDecl
%type <p_notStartAction> notStartAction
%type <p_start> start
%type <p_startActions> startActions
%type <p_codeDecls> _QcodeDecl_E_Plus
%type <p_statement> statement
%type <p_shapeExprDecl> shapeExprDecl
%type <p__O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C> _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C
%type <p_shapeExpression> shapeExpression
%type <p_inlineShapeExpression> inlineShapeExpression
%type <p_shapeOr> shapeOr
%type <p__O_QIT_OR_E_S_QshapeAnd_E_C> _O_QIT_OR_E_S_QshapeAnd_E_C
%type <p__O_QIT_OR_E_S_QshapeAnd_E_Cs> _Q_O_QIT_OR_E_S_QshapeAnd_E_C_E_Star
%type <p_inlineShapeOr> inlineShapeOr
%type <p__O_QIT_OR_E_S_QinlineShapeAnd_E_C> _O_QIT_OR_E_S_QinlineShapeAnd_E_C
%type <p__O_QIT_OR_E_S_QinlineShapeAnd_E_Cs> _Q_O_QIT_OR_E_S_QinlineShapeAnd_E_C_E_Star
%type <p_shapeAnd> shapeAnd
%type <p__O_QIT_AND_E_S_QshapeNot_E_C> _O_QIT_AND_E_S_QshapeNot_E_C
%type <p__O_QIT_AND_E_S_QshapeNot_E_Cs> _Q_O_QIT_AND_E_S_QshapeNot_E_C_E_Star
%type <p_inlineShapeAnd> inlineShapeAnd
%type <p__O_QIT_AND_E_S_QinlineShapeNot_E_C> _O_QIT_AND_E_S_QinlineShapeNot_E_C
%type <p__O_QIT_AND_E_S_QinlineShapeNot_E_Cs> _Q_O_QIT_AND_E_S_QinlineShapeNot_E_C_E_Star
%type <p_shapeNot> shapeNot
%type <p__QIT_NOT_E_Opt> _QIT_NOT_E_Opt
%type <p_inlineShapeNot> inlineShapeNot
%type <p_shapeAtom> shapeAtom
%type <p__QshapeOrRef_E_Opt> _QshapeOrRef_E_Opt
%type <p__QnonLitNodeConstraint_E_Opt> _QnonLitNodeConstraint_E_Opt
%type <p_inlineShapeAtom> inlineShapeAtom
%type <p__QinlineShapeOrRef_E_Opt> _QinlineShapeOrRef_E_Opt
%type <p_shapeOrRef> shapeOrRef
%type <p_inlineShapeOrRef> inlineShapeOrRef
%type <p_shapeRef> shapeRef
%type <p_litNodeConstraint> litNodeConstraint
%type <p_xsFacets> _QxsFacet_E_Star
%type <p_numericFacets> _QnumericFacet_E_Plus
%type <p_nonLitNodeConstraint> nonLitNodeConstraint
%type <p_stringFacets> _QstringFacet_E_Star
%type <p_stringFacets> _QstringFacet_E_Plus
%type <p_nonLiteralKind> nonLiteralKind
%type <p_xsFacet> xsFacet
%type <p_stringFacet> stringFacet
%type <p_stringLength> stringLength
%type <p_numericFacet> numericFacet
%type <p_numericRange> numericRange
%type <p_numericLength> numericLength
%type <p_shapeDefinition> shapeDefinition
%type <p_annotations> _Qannotation_E_Star
%type <p_inlineShapeDefinition> inlineShapeDefinition
%type <p__QtripleExpression_E_Opt> _QtripleExpression_E_Opt
%type <p_shapeQualifiers> shapeQualifiers
%type <p__O_QextraPropertySet_E_Or_QIT_CLOSED_E_C> _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C
%type <p__O_QextraPropertySet_E_Or_QIT_CLOSED_E_Cs> _Q_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_E_Star
%type <p_extraPropertySet> extraPropertySet
%type <p_predicates> _Qpredicate_E_Plus
%type <p_tripleExpression> tripleExpression
%type <p_oneOfTripleExpr> oneOfTripleExpr
%type <p_multiElementOneOf> multiElementOneOf
%type <p__O_QGT_PIPE_E_S_QgroupTripleExpr_E_C> _O_QGT_PIPE_E_S_QgroupTripleExpr_E_C
%type <p__O_QGT_PIPE_E_S_QgroupTripleExpr_E_Cs> _Q_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C_E_Plus
%type <p_groupTripleExpr> groupTripleExpr
%type <p_singleElementGroup> singleElementGroup
%type <p__QGT_SEMI_E_Opt> _QGT_SEMI_E_Opt
%type <p_multiElementGroup> multiElementGroup
%type <p__O_QGT_SEMI_E_S_QunaryTripleExpr_E_C> _O_QGT_SEMI_E_S_QunaryTripleExpr_E_C
%type <p__O_QGT_SEMI_E_S_QunaryTripleExpr_E_Cs> _Q_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C_E_Plus
%type <p_unaryTripleExpr> unaryTripleExpr
%type <p__O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C> _O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C
%type <p__Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt> _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt
%type <p__O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C> _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C
%type <p_bracketedTripleExpr> bracketedTripleExpr
%type <p__Qcardinality_E_Opt> _Qcardinality_E_Opt
%type <p_tripleConstraint> tripleConstraint
%type <p__QsenseFlags_E_Opt> _QsenseFlags_E_Opt
%type <p_cardinality> cardinality
%type <p_senseFlags> senseFlags
%type <p_valueSet> valueSet
%type <p_valueSetValues> _QvalueSetValue_E_Star
%type <p_valueSetValue> valueSetValue
%type <p_iriExclusions> _QiriExclusion_E_Plus
%type <p_literalExclusions> _QliteralExclusion_E_Plus
%type <p_languageExclusions> _QlanguageExclusion_E_Plus
%type <p__O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C> _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C
%type <p_iriRange> iriRange
%type <p_iriExclusions> _QiriExclusion_E_Star
%type <p__O_QGT_KINDA_E_S_QiriExclusion_E_Star_C> _O_QGT_KINDA_E_S_QiriExclusion_E_Star_C
%type <p__Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt> _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt
%type <p_iriExclusion> iriExclusion
%type <p__QGT_KINDA_E_Opt> _QGT_KINDA_E_Opt
%type <p_literalRange> literalRange
%type <p_literalExclusions> _QliteralExclusion_E_Star
%type <p__O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C> _O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C
%type <p__Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt> _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt
%type <p_literalExclusion> literalExclusion
%type <p_languageRange> languageRange
%type <p_languageExclusions> _QlanguageExclusion_E_Star
%type <p__O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C> _O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C
%type <p__Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt> _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt
%type <p_languageExclusion> languageExclusion
%type <p_include> include
%type <p_annotation> annotation
%type <p__O_Qiri_E_Or_Qliteral_E_C> _O_Qiri_E_Or_Qliteral_E_C
%type <p_semanticActions> semanticActions
%type <p_codeDecls> _QcodeDecl_E_Star
%type <p_codeDecl> codeDecl
%type <p__O_QCODE_E_Or_QGT_MODULO_E_C> _O_QCODE_E_Or_QGT_MODULO_E_C
%type <p_literal> literal
%type <p_predicate> predicate
%type <p_datatype> datatype
%type <p_shapeExprLabel> shapeExprLabel
%type <p_tripleExprLabel> tripleExprLabel
%type <p_rawNumeric> rawNumeric
%type <p_numericLiteral> numericLiteral
%type <p_rdfLiteral> rdfLiteral
%type <p__O_QGT_DTYPE_E_S_Qdatatype_E_C> _O_QGT_DTYPE_E_S_Qdatatype_E_C
%type <p__Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt> _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt
%type <p_booleanLiteral> booleanLiteral
%type <p_string> string
%type <p_langString> langString
%type <p_iri> iri
%type <p_prefixedName> prefixedName
%type <p_blankNode> blankNode

/* END TokenBlock */

//%destructor { delete $$; } BlankNode

 /*** END shexParseToXml - Change the grammar's tokens above ***/

%{
#include <stdarg.h>
#include "shex-parse-to-XMLScanner.hh"

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
    driver.root = new shexDoc($1, $2);
}
;

_Qdirective_E_Star:
    {
    $$ = new yacker::_ProductionVector<directive*>("_Qdirective_E_Star");
}

    | _Qdirective_E_Star directive	{
    $1->push_back($2);
    $$ = $1;
}
;

_O_QnotStartAction_E_Or_QstartActions_E_C:
    notStartAction	{
    $$ = new _O_QnotStartAction_E_Or_QstartActions_E_C_rule0($1);
}

    | startActions	{
    $$ = new _O_QnotStartAction_E_Or_QstartActions_E_C_rule1($1);
}
;

_Qstatement_E_Star:
    {
    $$ = new yacker::_ProductionVector<statement*>("_Qstatement_E_Star");
}

    | _Qstatement_E_Star statement	{
    $1->push_back($2);
    $$ = $1;
}
;

_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C:
    _O_QnotStartAction_E_Or_QstartActions_E_C _Qstatement_E_Star	{
    $$ = new _O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C($1, $2);
}
;

_Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt:
    {
    $$ = new _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt_rule0();
}

    | _O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C	{
    $$ = new _Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt_rule1($1);
}
;

directive:
    baseDecl	{
    $$ = new directive_rule0($1);
}

    | prefixDecl	{
    $$ = new directive_rule1($1);
}

    | importDecl	{
    $$ = new directive_rule2($1);
}
;

baseDecl:
    IT_BASE IRIREF	{
    $$ = new baseDecl($1, $2);
}
;

prefixDecl:
    IT_PREFIX PNAME_NS IRIREF	{
    $$ = new prefixDecl($1, $2, $3);
}
;

importDecl:
    IT_IMPORT IRIREF	{
    $$ = new importDecl($1, $2);
}
;

notStartAction:
    start	{
    $$ = new notStartAction_rule0($1);
}

    | shapeExprDecl	{
    $$ = new notStartAction_rule1($1);
}
;

start:
    IT_START GT_EQUAL inlineShapeExpression	{
    $$ = new start($1, $2, $3);
}
;

startActions:
    _QcodeDecl_E_Plus	{
    $$ = new startActions($1);
}
;

_QcodeDecl_E_Plus:
    codeDecl	{
    $$ = new yacker::_ProductionVector<codeDecl*>("_QcodeDecl_E_Plus", $1);
}

    | _QcodeDecl_E_Plus codeDecl	{
    $1->push_back($2);
    $$ = $1;
}
;

statement:
    directive	{
    $$ = new statement_rule0($1);
}

    | notStartAction	{
    $$ = new statement_rule1($1);
}
;

shapeExprDecl:
    shapeExprLabel _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C	{
    $$ = new shapeExprDecl($1, $2);
}
;

_O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C:
    shapeExpression	{
    $$ = new _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C_rule0($1);
}

    | IT_EXTERNAL	{
    $$ = new _O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C_rule1($1);
}
;

shapeExpression:
    shapeOr	{
    $$ = new shapeExpression($1);
}
;

inlineShapeExpression:
    inlineShapeOr	{
    $$ = new inlineShapeExpression($1);
}
;

shapeOr:
    shapeAnd _Q_O_QIT_OR_E_S_QshapeAnd_E_C_E_Star	{
    $$ = new shapeOr($1, $2);
}
;

_O_QIT_OR_E_S_QshapeAnd_E_C:
    IT_OR shapeAnd	{
    $$ = new _O_QIT_OR_E_S_QshapeAnd_E_C($1, $2);
}
;

_Q_O_QIT_OR_E_S_QshapeAnd_E_C_E_Star:
    {
    $$ = new yacker::_ProductionVector<_O_QIT_OR_E_S_QshapeAnd_E_C*>("_Q_O_QIT_OR_E_S_QshapeAnd_E_C_E_Star");
}

    | _Q_O_QIT_OR_E_S_QshapeAnd_E_C_E_Star _O_QIT_OR_E_S_QshapeAnd_E_C	{
    $1->push_back($2);
    $$ = $1;
}
;

inlineShapeOr:
    inlineShapeAnd _Q_O_QIT_OR_E_S_QinlineShapeAnd_E_C_E_Star	{
    $$ = new inlineShapeOr($1, $2);
}
;

_O_QIT_OR_E_S_QinlineShapeAnd_E_C:
    IT_OR inlineShapeAnd	{
    $$ = new _O_QIT_OR_E_S_QinlineShapeAnd_E_C($1, $2);
}
;

_Q_O_QIT_OR_E_S_QinlineShapeAnd_E_C_E_Star:
    {
    $$ = new yacker::_ProductionVector<_O_QIT_OR_E_S_QinlineShapeAnd_E_C*>("_Q_O_QIT_OR_E_S_QinlineShapeAnd_E_C_E_Star");
}

    | _Q_O_QIT_OR_E_S_QinlineShapeAnd_E_C_E_Star _O_QIT_OR_E_S_QinlineShapeAnd_E_C	{
    $1->push_back($2);
    $$ = $1;
}
;

shapeAnd:
    shapeNot _Q_O_QIT_AND_E_S_QshapeNot_E_C_E_Star	{
    $$ = new shapeAnd($1, $2);
}
;

_O_QIT_AND_E_S_QshapeNot_E_C:
    IT_AND shapeNot	{
    $$ = new _O_QIT_AND_E_S_QshapeNot_E_C($1, $2);
}
;

_Q_O_QIT_AND_E_S_QshapeNot_E_C_E_Star:
    {
    $$ = new yacker::_ProductionVector<_O_QIT_AND_E_S_QshapeNot_E_C*>("_Q_O_QIT_AND_E_S_QshapeNot_E_C_E_Star");
}

    | _Q_O_QIT_AND_E_S_QshapeNot_E_C_E_Star _O_QIT_AND_E_S_QshapeNot_E_C	{
    $1->push_back($2);
    $$ = $1;
}
;

inlineShapeAnd:
    inlineShapeNot _Q_O_QIT_AND_E_S_QinlineShapeNot_E_C_E_Star	{
    $$ = new inlineShapeAnd($1, $2);
}
;

_O_QIT_AND_E_S_QinlineShapeNot_E_C:
    IT_AND inlineShapeNot	{
    $$ = new _O_QIT_AND_E_S_QinlineShapeNot_E_C($1, $2);
}
;

_Q_O_QIT_AND_E_S_QinlineShapeNot_E_C_E_Star:
    {
    $$ = new yacker::_ProductionVector<_O_QIT_AND_E_S_QinlineShapeNot_E_C*>("_Q_O_QIT_AND_E_S_QinlineShapeNot_E_C_E_Star");
}

    | _Q_O_QIT_AND_E_S_QinlineShapeNot_E_C_E_Star _O_QIT_AND_E_S_QinlineShapeNot_E_C	{
    $1->push_back($2);
    $$ = $1;
}
;

shapeNot:
    _QIT_NOT_E_Opt shapeAtom	{
    $$ = new shapeNot($1, $2);
}
;

_QIT_NOT_E_Opt:
    {
    $$ = new _QIT_NOT_E_Opt_rule0();
}

    | IT_NOT	{
    $$ = new _QIT_NOT_E_Opt_rule1($1);
}
;

inlineShapeNot:
    _QIT_NOT_E_Opt inlineShapeAtom	{
    $$ = new inlineShapeNot($1, $2);
}
;

shapeAtom:
    nonLitNodeConstraint _QshapeOrRef_E_Opt	{
    $$ = new shapeAtom_rule0($1, $2);
}

    | litNodeConstraint	{
    $$ = new shapeAtom_rule1($1);
}

    | shapeOrRef _QnonLitNodeConstraint_E_Opt	{
    $$ = new shapeAtom_rule2($1, $2);
}

    | GT_LPAREN shapeExpression GT_RPAREN	{
    $$ = new shapeAtom_rule3($1, $2, $3);
}

    | GT_DOT	{
    $$ = new shapeAtom_rule4($1);
}
;

_QshapeOrRef_E_Opt:
    {
    $$ = new _QshapeOrRef_E_Opt_rule0();
}

    | shapeOrRef	{
    $$ = new _QshapeOrRef_E_Opt_rule1($1);
}
;

_QnonLitNodeConstraint_E_Opt:
    {
    $$ = new _QnonLitNodeConstraint_E_Opt_rule0();
}

    | nonLitNodeConstraint	{
    $$ = new _QnonLitNodeConstraint_E_Opt_rule1($1);
}
;

inlineShapeAtom:
    nonLitNodeConstraint _QinlineShapeOrRef_E_Opt	{
    $$ = new inlineShapeAtom_rule0($1, $2);
}

    | litNodeConstraint	{
    $$ = new inlineShapeAtom_rule1($1);
}

    | inlineShapeOrRef _QnonLitNodeConstraint_E_Opt	{
    $$ = new inlineShapeAtom_rule2($1, $2);
}

    | GT_LPAREN shapeExpression GT_RPAREN	{
    $$ = new inlineShapeAtom_rule3($1, $2, $3);
}

    | GT_DOT	{
    $$ = new inlineShapeAtom_rule4($1);
}
;

_QinlineShapeOrRef_E_Opt:
    {
    $$ = new _QinlineShapeOrRef_E_Opt_rule0();
}

    | inlineShapeOrRef	{
    $$ = new _QinlineShapeOrRef_E_Opt_rule1($1);
}
;

shapeOrRef:
    shapeDefinition	{
    $$ = new shapeOrRef_rule0($1);
}

    | shapeRef	{
    $$ = new shapeOrRef_rule1($1);
}
;

inlineShapeOrRef:
    inlineShapeDefinition	{
    $$ = new inlineShapeOrRef_rule0($1);
}

    | shapeRef	{
    $$ = new inlineShapeOrRef_rule1($1);
}
;

shapeRef:
    ATPNAME_LN	{
    $$ = new shapeRef_rule0($1);
}

    | ATPNAME_NS	{
    $$ = new shapeRef_rule1($1);
}

    | GT_AT shapeExprLabel	{
    $$ = new shapeRef_rule2($1, $2);
}
;

litNodeConstraint:
    IT_LITERAL _QxsFacet_E_Star	{
    $$ = new litNodeConstraint_rule0($1, $2);
}

    | datatype _QxsFacet_E_Star	{
    $$ = new litNodeConstraint_rule1($1, $2);
}

    | valueSet _QxsFacet_E_Star	{
    $$ = new litNodeConstraint_rule2($1, $2);
}

    | _QnumericFacet_E_Plus	{
    $$ = new litNodeConstraint_rule3($1);
}
;

_QxsFacet_E_Star:
    {
    $$ = new yacker::_ProductionVector<xsFacet*>("_QxsFacet_E_Star");
}

    | _QxsFacet_E_Star xsFacet	{
    $1->push_back($2);
    $$ = $1;
}
;

_QnumericFacet_E_Plus:
    numericFacet	{
    $$ = new yacker::_ProductionVector<numericFacet*>("_QnumericFacet_E_Plus", $1);
}

    | _QnumericFacet_E_Plus numericFacet	{
    $1->push_back($2);
    $$ = $1;
}
;

nonLitNodeConstraint:
    nonLiteralKind _QstringFacet_E_Star	{
    $$ = new nonLitNodeConstraint_rule0($1, $2);
}

    | _QstringFacet_E_Plus	{
    $$ = new nonLitNodeConstraint_rule1($1);
}
;

_QstringFacet_E_Star:
    {
    $$ = new yacker::_ProductionVector<stringFacet*>("_QstringFacet_E_Star");
}

    | _QstringFacet_E_Star stringFacet	{
    $1->push_back($2);
    $$ = $1;
}
;

_QstringFacet_E_Plus:
    stringFacet	{
    $$ = new yacker::_ProductionVector<stringFacet*>("_QstringFacet_E_Plus", $1);
}

    | _QstringFacet_E_Plus stringFacet	{
    $1->push_back($2);
    $$ = $1;
}
;

nonLiteralKind:
    IT_IRI	{
    $$ = new nonLiteralKind_rule0($1);
}

    | IT_BNODE	{
    $$ = new nonLiteralKind_rule1($1);
}

    | IT_NONLITERAL	{
    $$ = new nonLiteralKind_rule2($1);
}
;

xsFacet:
    stringFacet	{
    $$ = new xsFacet_rule0($1);
}

    | numericFacet	{
    $$ = new xsFacet_rule1($1);
}
;

stringFacet:
    stringLength INTEGER	{
    $$ = new stringFacet_rule0($1, $2);
}

    | REGEXP	{
    $$ = new stringFacet_rule1($1);
}
;

stringLength:
    IT_LENGTH	{
    $$ = new stringLength_rule0($1);
}

    | IT_MINLENGTH	{
    $$ = new stringLength_rule1($1);
}

    | IT_MAXLENGTH	{
    $$ = new stringLength_rule2($1);
}
;

numericFacet:
    numericRange rawNumeric	{
    $$ = new numericFacet_rule0($1, $2);
}

    | numericLength INTEGER	{
    $$ = new numericFacet_rule1($1, $2);
}
;

numericRange:
    IT_MININCLUSIVE	{
    $$ = new numericRange_rule0($1);
}

    | IT_MINEXCLUSIVE	{
    $$ = new numericRange_rule1($1);
}

    | IT_MAXINCLUSIVE	{
    $$ = new numericRange_rule2($1);
}

    | IT_MAXEXCLUSIVE	{
    $$ = new numericRange_rule3($1);
}
;

numericLength:
    IT_TOTALDIGITS	{
    $$ = new numericLength_rule0($1);
}

    | IT_FRACTIONDIGITS	{
    $$ = new numericLength_rule1($1);
}
;

shapeDefinition:
    inlineShapeDefinition _Qannotation_E_Star semanticActions	{
    $$ = new shapeDefinition($1, $2, $3);
}
;

_Qannotation_E_Star:
    {
    $$ = new yacker::_ProductionVector<annotation*>("_Qannotation_E_Star");
}

    | _Qannotation_E_Star annotation	{
    $1->push_back($2);
    $$ = $1;
}
;

inlineShapeDefinition:
    shapeQualifiers GT_LCURLEY _QtripleExpression_E_Opt GT_RCURLEY	{
    $$ = new inlineShapeDefinition($1, $2, $3, $4);
}
;

_QtripleExpression_E_Opt:
    {
    $$ = new _QtripleExpression_E_Opt_rule0();
}

    | tripleExpression	{
    $$ = new _QtripleExpression_E_Opt_rule1($1);
}
;

shapeQualifiers:
    _Q_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_E_Star	{
    $$ = new shapeQualifiers($1);
}
;

_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C:
    extraPropertySet	{
    $$ = new _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_rule0($1);
}

    | IT_CLOSED	{
    $$ = new _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_rule1($1);
}
;

_Q_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_E_Star:
    {
    $$ = new yacker::_ProductionVector<_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C*>("_Q_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_E_Star");
}

    | _Q_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_E_Star _O_QextraPropertySet_E_Or_QIT_CLOSED_E_C	{
    $1->push_back($2);
    $$ = $1;
}
;

extraPropertySet:
    IT_EXTRA _Qpredicate_E_Plus	{
    $$ = new extraPropertySet($1, $2);
}
;

_Qpredicate_E_Plus:
    predicate	{
    $$ = new yacker::_ProductionVector<predicate*>("_Qpredicate_E_Plus", $1);
}

    | _Qpredicate_E_Plus predicate	{
    $1->push_back($2);
    $$ = $1;
}
;

tripleExpression:
    oneOfTripleExpr	{
    $$ = new tripleExpression($1);
}
;

oneOfTripleExpr:
    groupTripleExpr	{
    $$ = new oneOfTripleExpr_rule0($1);
}

    | multiElementOneOf	{
    $$ = new oneOfTripleExpr_rule1($1);
}
;

multiElementOneOf:
    groupTripleExpr _Q_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C_E_Plus	{
    $$ = new multiElementOneOf($1, $2);
}
;

_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C:
    GT_PIPE groupTripleExpr	{
    $$ = new _O_QGT_PIPE_E_S_QgroupTripleExpr_E_C($1, $2);
}
;

_Q_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C_E_Plus:
    _O_QGT_PIPE_E_S_QgroupTripleExpr_E_C	{
    $$ = new yacker::_ProductionVector<_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C*>("_Q_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C_E_Plus", $1);
}

    | _Q_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C_E_Plus _O_QGT_PIPE_E_S_QgroupTripleExpr_E_C	{
    $1->push_back($2);
    $$ = $1;
}
;

groupTripleExpr:
    singleElementGroup	{
    $$ = new groupTripleExpr_rule0($1);
}

    | multiElementGroup	{
    $$ = new groupTripleExpr_rule1($1);
}
;

singleElementGroup:
    unaryTripleExpr _QGT_SEMI_E_Opt	{
    $$ = new singleElementGroup($1, $2);
}
;

_QGT_SEMI_E_Opt:
    {
    $$ = new _QGT_SEMI_E_Opt_rule0();
}

    | GT_SEMI	{
    $$ = new _QGT_SEMI_E_Opt_rule1($1);
}
;

multiElementGroup:
    unaryTripleExpr _Q_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C_E_Plus _QGT_SEMI_E_Opt	{
    $$ = new multiElementGroup($1, $2, $3);
}
;

_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C:
    GT_SEMI unaryTripleExpr	{
    $$ = new _O_QGT_SEMI_E_S_QunaryTripleExpr_E_C($1, $2);
}
;

_Q_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C_E_Plus:
    _O_QGT_SEMI_E_S_QunaryTripleExpr_E_C	{
    $$ = new yacker::_ProductionVector<_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C*>("_Q_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C_E_Plus", $1);
}

    | _Q_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C_E_Plus _O_QGT_SEMI_E_S_QunaryTripleExpr_E_C	{
    $1->push_back($2);
    $$ = $1;
}
;

unaryTripleExpr:
    _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C	{
    $$ = new unaryTripleExpr_rule0($1, $2);
}

    | include	{
    $$ = new unaryTripleExpr_rule1($1);
}
;

_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C:
    GT_DOLLAR tripleExprLabel	{
    $$ = new _O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C($1, $2);
}
;

_Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt:
    {
    $$ = new _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt_rule0();
}

    | _O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C	{
    $$ = new _Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt_rule1($1);
}
;

_O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C:
    tripleConstraint	{
    $$ = new _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C_rule0($1);
}

    | bracketedTripleExpr	{
    $$ = new _O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C_rule1($1);
}
;

bracketedTripleExpr:
    GT_LPAREN tripleExpression GT_RPAREN _Qcardinality_E_Opt _Qannotation_E_Star semanticActions	{
    $$ = new bracketedTripleExpr($1, $2, $3, $4, $5, $6);
}
;

_Qcardinality_E_Opt:
    {
    $$ = new _Qcardinality_E_Opt_rule0();
}

    | cardinality	{
    $$ = new _Qcardinality_E_Opt_rule1($1);
}
;

tripleConstraint:
    _QsenseFlags_E_Opt predicate inlineShapeExpression _Qcardinality_E_Opt _Qannotation_E_Star semanticActions	{
    $$ = new tripleConstraint($1, $2, $3, $4, $5, $6);
}
;

_QsenseFlags_E_Opt:
    {
    $$ = new _QsenseFlags_E_Opt_rule0();
}

    | senseFlags	{
    $$ = new _QsenseFlags_E_Opt_rule1($1);
}
;

cardinality:
    GT_TIMES	{
    $$ = new cardinality_rule0($1);
}

    | GT_PLUS	{
    $$ = new cardinality_rule1($1);
}

    | GT_OPT	{
    $$ = new cardinality_rule2($1);
}

    | REPEAT_RANGE	{
    $$ = new cardinality_rule3($1);
}
;

senseFlags:
    GT_CARROT	{
    $$ = new senseFlags($1);
}
;

valueSet:
    GT_LBRACKET _QvalueSetValue_E_Star GT_RBRACKET	{
    $$ = new valueSet($1, $2, $3);
}
;

_QvalueSetValue_E_Star:
    {
    $$ = new yacker::_ProductionVector<valueSetValue*>("_QvalueSetValue_E_Star");
}

    | _QvalueSetValue_E_Star valueSetValue	{
    $1->push_back($2);
    $$ = $1;
}
;

valueSetValue:
    iriRange	{
    $$ = new valueSetValue_rule0($1);
}

    | literalRange	{
    $$ = new valueSetValue_rule1($1);
}

    | languageRange	{
    $$ = new valueSetValue_rule2($1);
}

    | GT_DOT _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C	{
    $$ = new valueSetValue_rule3($1, $2);
}
;

_QiriExclusion_E_Plus:
    iriExclusion	{
    $$ = new yacker::_ProductionVector<iriExclusion*>("_QiriExclusion_E_Plus", $1);
}

    | _QiriExclusion_E_Plus iriExclusion	{
    $1->push_back($2);
    $$ = $1;
}
;

_QliteralExclusion_E_Plus:
    literalExclusion	{
    $$ = new yacker::_ProductionVector<literalExclusion*>("_QliteralExclusion_E_Plus", $1);
}

    | _QliteralExclusion_E_Plus literalExclusion	{
    $1->push_back($2);
    $$ = $1;
}
;

_QlanguageExclusion_E_Plus:
    languageExclusion	{
    $$ = new yacker::_ProductionVector<languageExclusion*>("_QlanguageExclusion_E_Plus", $1);
}

    | _QlanguageExclusion_E_Plus languageExclusion	{
    $1->push_back($2);
    $$ = $1;
}
;

_O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C:
    _QiriExclusion_E_Plus	{
    $$ = new _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule0($1);
}

    | _QliteralExclusion_E_Plus	{
    $$ = new _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule1($1);
}

    | _QlanguageExclusion_E_Plus	{
    $$ = new _O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule2($1);
}
;

iriRange:
    iri _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt	{
    $$ = new iriRange($1, $2);
}
;

_QiriExclusion_E_Star:
    {
    $$ = new yacker::_ProductionVector<iriExclusion*>("_QiriExclusion_E_Star");
}

    | _QiriExclusion_E_Star iriExclusion	{
    $1->push_back($2);
    $$ = $1;
}
;

_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C:
    GT_KINDA _QiriExclusion_E_Star	{
    $$ = new _O_QGT_KINDA_E_S_QiriExclusion_E_Star_C($1, $2);
}
;

_Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt:
    {
    $$ = new _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt_rule0();
}

    | _O_QGT_KINDA_E_S_QiriExclusion_E_Star_C	{
    $$ = new _Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt_rule1($1);
}
;

iriExclusion:
    GT_MINUS iri _QGT_KINDA_E_Opt	{
    $$ = new iriExclusion($1, $2, $3);
}
;

_QGT_KINDA_E_Opt:
    {
    $$ = new _QGT_KINDA_E_Opt_rule0();
}

    | GT_KINDA	{
    $$ = new _QGT_KINDA_E_Opt_rule1($1);
}
;

literalRange:
    literal _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt	{
    $$ = new literalRange($1, $2);
}
;

_QliteralExclusion_E_Star:
    {
    $$ = new yacker::_ProductionVector<literalExclusion*>("_QliteralExclusion_E_Star");
}

    | _QliteralExclusion_E_Star literalExclusion	{
    $1->push_back($2);
    $$ = $1;
}
;

_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C:
    GT_KINDA _QliteralExclusion_E_Star	{
    $$ = new _O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C($1, $2);
}
;

_Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt:
    {
    $$ = new _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt_rule0();
}

    | _O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C	{
    $$ = new _Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt_rule1($1);
}
;

literalExclusion:
    GT_MINUS literal _QGT_KINDA_E_Opt	{
    $$ = new literalExclusion($1, $2, $3);
}
;

languageRange:
    LANGTAG _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt	{
    $$ = new languageRange_rule0($1, $2);
}

    | GT_AT GT_KINDA _QlanguageExclusion_E_Star	{
    $$ = new languageRange_rule1($1, $2, $3);
}
;

_QlanguageExclusion_E_Star:
    {
    $$ = new yacker::_ProductionVector<languageExclusion*>("_QlanguageExclusion_E_Star");
}

    | _QlanguageExclusion_E_Star languageExclusion	{
    $1->push_back($2);
    $$ = $1;
}
;

_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C:
    GT_KINDA _QlanguageExclusion_E_Star	{
    $$ = new _O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C($1, $2);
}
;

_Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt:
    {
    $$ = new _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt_rule0();
}

    | _O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C	{
    $$ = new _Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt_rule1($1);
}
;

languageExclusion:
    GT_MINUS LANGTAG _QGT_KINDA_E_Opt	{
    $$ = new languageExclusion($1, $2, $3);
}
;

include:
    GT_AMP tripleExprLabel	{
    $$ = new include($1, $2);
}
;

annotation:
    GT_DIVIDE_DIVIDE predicate _O_Qiri_E_Or_Qliteral_E_C	{
    $$ = new annotation($1, $2, $3);
}
;

_O_Qiri_E_Or_Qliteral_E_C:
    iri	{
    $$ = new _O_Qiri_E_Or_Qliteral_E_C_rule0($1);
}

    | literal	{
    $$ = new _O_Qiri_E_Or_Qliteral_E_C_rule1($1);
}
;

semanticActions:
    _QcodeDecl_E_Star	{
    $$ = new semanticActions($1);
}
;

_QcodeDecl_E_Star:
    {
    $$ = new yacker::_ProductionVector<codeDecl*>("_QcodeDecl_E_Star");
}

    | _QcodeDecl_E_Star codeDecl	{
    $1->push_back($2);
    $$ = $1;
}
;

codeDecl:
    GT_MODULO iri _O_QCODE_E_Or_QGT_MODULO_E_C	{
    $$ = new codeDecl($1, $2, $3);
}
;

_O_QCODE_E_Or_QGT_MODULO_E_C:
    CODE	{
    $$ = new _O_QCODE_E_Or_QGT_MODULO_E_C_rule0($1);
}

    | GT_MODULO	{
    $$ = new _O_QCODE_E_Or_QGT_MODULO_E_C_rule1($1);
}
;

literal:
    rdfLiteral	{
    $$ = new literal_rule0($1);
}

    | numericLiteral	{
    $$ = new literal_rule1($1);
}

    | booleanLiteral	{
    $$ = new literal_rule2($1);
}
;

predicate:
    iri	{
    $$ = new predicate_rule0($1);
}

    | RDF_TYPE	{
    $$ = new predicate_rule1($1);
}
;

datatype:
    iri	{
    $$ = new datatype($1);
}
;

shapeExprLabel:
    iri	{
    $$ = new shapeExprLabel_rule0($1);
}

    | blankNode	{
    $$ = new shapeExprLabel_rule1($1);
}
;

tripleExprLabel:
    iri	{
    $$ = new tripleExprLabel_rule0($1);
}

    | blankNode	{
    $$ = new tripleExprLabel_rule1($1);
}
;

rawNumeric:
    INTEGER	{
    $$ = new rawNumeric_rule0($1);
}

    | DECIMAL	{
    $$ = new rawNumeric_rule1($1);
}

    | DOUBLE	{
    $$ = new rawNumeric_rule2($1);
}
;

numericLiteral:
    rawNumeric	{
    $$ = new numericLiteral($1);
}
;

rdfLiteral:
    langString	{
    $$ = new rdfLiteral_rule0($1);
}

    | string _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt	{
    $$ = new rdfLiteral_rule1($1, $2);
}
;

_O_QGT_DTYPE_E_S_Qdatatype_E_C:
    GT_DTYPE datatype	{
    $$ = new _O_QGT_DTYPE_E_S_Qdatatype_E_C($1, $2);
}
;

_Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt:
    {
    $$ = new _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt_rule0();
}

    | _O_QGT_DTYPE_E_S_Qdatatype_E_C	{
    $$ = new _Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt_rule1($1);
}
;

booleanLiteral:
    IT_true	{
    $$ = new booleanLiteral_rule0($1);
}

    | IT_false	{
    $$ = new booleanLiteral_rule1($1);
}
;

string:
    STRING_LITERAL1	{
    $$ = new string_rule0($1);
}

    | STRING_LITERAL_LONG1	{
    $$ = new string_rule1($1);
}

    | STRING_LITERAL2	{
    $$ = new string_rule2($1);
}

    | STRING_LITERAL_LONG2	{
    $$ = new string_rule3($1);
}
;

langString:
    LANG_STRING_LITERAL1	{
    $$ = new langString_rule0($1);
}

    | LANG_STRING_LITERAL_LONG1	{
    $$ = new langString_rule1($1);
}

    | LANG_STRING_LITERAL2	{
    $$ = new langString_rule2($1);
}

    | LANG_STRING_LITERAL_LONG2	{
    $$ = new langString_rule3($1);
}
;

iri:
    IRIREF	{
    $$ = new iri_rule0($1);
}

    | prefixedName	{
    $$ = new iri_rule1($1);
}
;

prefixedName:
    PNAME_LN	{
    $$ = new prefixedName_rule0($1);
}

    | PNAME_NS	{
    $$ = new prefixedName_rule1($1);
}
;

blankNode:
    BLANK_NODE_LABEL	{
    $$ = new blankNode($1);
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

std::string _Production::toStr(std::ofstream* out) {
    std::stringstream ret;
    for (size_t i = 0; i < size(); i++) {
	if (i > 0)
	    ret << " ";
	ret << getElement(i)->toStr(out);
    }
    return ret.str();
}

#define TAB "  "
char const * yit = "yacker:implicit-terminal";
#define ns "\n xmlns=\"http://www.w3.org/2005/01/yacker/uploads/spec_2_1/\"\n xmlns:yacker=\"http://www.w3.org/2005/01/yacker/\""

std::string _Production::toXml (size_t depth, std::ofstream* out) {
    std::stringstream ret;
    if (!transparent()) {
	for (size_t i = 0; i < depth; i++)
	    ret << TAB;
	ret << "<" << getProductionName() << ">" << std::endl;
    }
    for (size_t i = 0; i < size(); i++) {
	_Base* base = getElement(i);
	if (base != NULL)
	    ret << base->toXml(transparent() ? depth : depth+1, out);
    }
    if (!transparent()) {
	for (size_t i = 0; i < depth; i++)
	    ret << TAB;
	ret << "</" << getProductionName() << ">" << std::endl;
    }
    return ret.str();
}

void _Production::makeArray (_Base** target[], ...) {
    va_list bases;
    va_start(bases, target);
    for (size_t argNo = 0; argNo < size(); argNo++)
	target[argNo] = va_arg(bases, _Base**);
    va_end(bases);    
}
void _Production::trace (unsigned creditToSelf) {
    if (!_Trace)
	return;
    *_Trace << "  " << getProductionName() << ":";
    if (creditToSelf > 0)
	*_Trace << " " << getProductionName() << "(" << creditToSelf << ")";

    for (size_t argNo = creditToSelf; argNo < size(); argNo++) {
	_Base* parm = getElement(argNo);
	*_Trace << " " << parm->getProductionName() << "(" << parm->size() << ")";
    }
    *_Trace << std::endl;

    for (size_t argNo = 0; argNo < size(); argNo++) {
	_Base* parm = getElement(argNo);
	size_t parmSize = parm->size();
	for (size_t j = 0; j < parmSize; j++)
	    *_Trace << "    " << (creditToSelf && argNo < size()-1 ? getProductionName() : parm->getProductionName()) << "(" << (creditToSelf && argNo < size()-1 ? argNo : j) << "): " << parm->getElement(j)->toStr() << std::endl;
    }
}

void _Token::trace() {
    if (!_Trace)
	return;
    *_Trace << "shift (" << getProductionName() << ", " << matched << ")" << std::endl;
}

void _Terminal::trace() {
    if (!_Trace)
	return;
    *_Trace << "shift (" << getProductionName() << ", " << terminal << ")" << std::endl;
}

std::string _Token::toStr(std::ofstream*) {
    std::string ret(matched);
    return ret;
}
std::string _Token::toXml(size_t depth, std::ofstream*) {
    std::stringstream ret;
    for (size_t i = 0; i < depth; i++)
	ret << TAB;
    ret << "<" << yit << ">" << matched << "</" << yit << ">" << std::endl;
    return ret.str();
}

std::string _Terminal::toStr(std::ofstream*) {
    std::string ret(terminal);
    return ret;
}
std::string _Terminal::toXml(size_t depth, std::ofstream*) {
    std::stringstream ret;
    for (size_t i = 0; i < depth; i++)
	ret << TAB;
    ret << "<" << getProductionName() << ">" << terminal << "</" << getProductionName() << ">" << std::endl;
    return ret.str();
}

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

shexDoc::~shexDoc () {
    delete m_directives;
    delete m__Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt;
}
_O_QnotStartAction_E_Or_QstartActions_E_C_rule0::~_O_QnotStartAction_E_Or_QstartActions_E_C_rule0 () {
    delete m_notStartAction;
}
_O_QnotStartAction_E_Or_QstartActions_E_C_rule1::~_O_QnotStartAction_E_Or_QstartActions_E_C_rule1 () {
    delete m_startActions;
}
_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C::~_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C () {
    delete m__O_QnotStartAction_E_Or_QstartActions_E_C;
    delete m_statements;
}
_Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt_rule0::~_Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt_rule0 () {
}
_Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt_rule1::~_Q_O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C_E_Opt_rule1 () {
    delete m__O_QnotStartAction_E_Or_QstartActions_E_S_Qstatement_E_Star_C;
}
directive_rule0::~directive_rule0 () {
    delete m_baseDecl;
}
directive_rule1::~directive_rule1 () {
    delete m_prefixDecl;
}
directive_rule2::~directive_rule2 () {
    delete m_importDecl;
}
baseDecl::~baseDecl () {
    delete m_IT_BASE;
    delete m_IRIREF;
}
prefixDecl::~prefixDecl () {
    delete m_IT_PREFIX;
    delete m_PNAME_NS;
    delete m_IRIREF;
}
importDecl::~importDecl () {
    delete m_IT_IMPORT;
    delete m_IRIREF;
}
notStartAction_rule0::~notStartAction_rule0 () {
    delete m_start;
}
notStartAction_rule1::~notStartAction_rule1 () {
    delete m_shapeExprDecl;
}
start::~start () {
    delete m_IT_START;
    delete m_GT_EQUAL;
    delete m_inlineShapeExpression;
}
startActions::~startActions () {
    delete m_codeDecls;
}
statement_rule0::~statement_rule0 () {
    delete m_directive;
}
statement_rule1::~statement_rule1 () {
    delete m_notStartAction;
}
shapeExprDecl::~shapeExprDecl () {
    delete m_shapeExprLabel;
    delete m__O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C;
}
_O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C_rule0::~_O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C_rule0 () {
    delete m_shapeExpression;
}
_O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C_rule1::~_O_QshapeExpression_E_Or_QIT_EXTERNAL_E_C_rule1 () {
    delete m_IT_EXTERNAL;
}
shapeExpression::~shapeExpression () {
    delete m_shapeOr;
}
inlineShapeExpression::~inlineShapeExpression () {
    delete m_inlineShapeOr;
}
shapeOr::~shapeOr () {
    delete m_shapeAnd;
    delete m__O_QIT_OR_E_S_QshapeAnd_E_Cs;
}
_O_QIT_OR_E_S_QshapeAnd_E_C::~_O_QIT_OR_E_S_QshapeAnd_E_C () {
    delete m_IT_OR;
    delete m_shapeAnd;
}
inlineShapeOr::~inlineShapeOr () {
    delete m_inlineShapeAnd;
    delete m__O_QIT_OR_E_S_QinlineShapeAnd_E_Cs;
}
_O_QIT_OR_E_S_QinlineShapeAnd_E_C::~_O_QIT_OR_E_S_QinlineShapeAnd_E_C () {
    delete m_IT_OR;
    delete m_inlineShapeAnd;
}
shapeAnd::~shapeAnd () {
    delete m_shapeNot;
    delete m__O_QIT_AND_E_S_QshapeNot_E_Cs;
}
_O_QIT_AND_E_S_QshapeNot_E_C::~_O_QIT_AND_E_S_QshapeNot_E_C () {
    delete m_IT_AND;
    delete m_shapeNot;
}
inlineShapeAnd::~inlineShapeAnd () {
    delete m_inlineShapeNot;
    delete m__O_QIT_AND_E_S_QinlineShapeNot_E_Cs;
}
_O_QIT_AND_E_S_QinlineShapeNot_E_C::~_O_QIT_AND_E_S_QinlineShapeNot_E_C () {
    delete m_IT_AND;
    delete m_inlineShapeNot;
}
shapeNot::~shapeNot () {
    delete m__QIT_NOT_E_Opt;
    delete m_shapeAtom;
}
_QIT_NOT_E_Opt_rule0::~_QIT_NOT_E_Opt_rule0 () {
}
_QIT_NOT_E_Opt_rule1::~_QIT_NOT_E_Opt_rule1 () {
    delete m_IT_NOT;
}
inlineShapeNot::~inlineShapeNot () {
    delete m__QIT_NOT_E_Opt;
    delete m_inlineShapeAtom;
}
shapeAtom_rule0::~shapeAtom_rule0 () {
    delete m_nonLitNodeConstraint;
    delete m__QshapeOrRef_E_Opt;
}
shapeAtom_rule1::~shapeAtom_rule1 () {
    delete m_litNodeConstraint;
}
shapeAtom_rule2::~shapeAtom_rule2 () {
    delete m_shapeOrRef;
    delete m__QnonLitNodeConstraint_E_Opt;
}
shapeAtom_rule3::~shapeAtom_rule3 () {
    delete m_GT_LPAREN;
    delete m_shapeExpression;
    delete m_GT_RPAREN;
}
shapeAtom_rule4::~shapeAtom_rule4 () {
    delete m_GT_DOT;
}
_QshapeOrRef_E_Opt_rule0::~_QshapeOrRef_E_Opt_rule0 () {
}
_QshapeOrRef_E_Opt_rule1::~_QshapeOrRef_E_Opt_rule1 () {
    delete m_shapeOrRef;
}
_QnonLitNodeConstraint_E_Opt_rule0::~_QnonLitNodeConstraint_E_Opt_rule0 () {
}
_QnonLitNodeConstraint_E_Opt_rule1::~_QnonLitNodeConstraint_E_Opt_rule1 () {
    delete m_nonLitNodeConstraint;
}
inlineShapeAtom_rule0::~inlineShapeAtom_rule0 () {
    delete m_nonLitNodeConstraint;
    delete m__QinlineShapeOrRef_E_Opt;
}
inlineShapeAtom_rule1::~inlineShapeAtom_rule1 () {
    delete m_litNodeConstraint;
}
inlineShapeAtom_rule2::~inlineShapeAtom_rule2 () {
    delete m_inlineShapeOrRef;
    delete m__QnonLitNodeConstraint_E_Opt;
}
inlineShapeAtom_rule3::~inlineShapeAtom_rule3 () {
    delete m_GT_LPAREN;
    delete m_shapeExpression;
    delete m_GT_RPAREN;
}
inlineShapeAtom_rule4::~inlineShapeAtom_rule4 () {
    delete m_GT_DOT;
}
_QinlineShapeOrRef_E_Opt_rule0::~_QinlineShapeOrRef_E_Opt_rule0 () {
}
_QinlineShapeOrRef_E_Opt_rule1::~_QinlineShapeOrRef_E_Opt_rule1 () {
    delete m_inlineShapeOrRef;
}
shapeOrRef_rule0::~shapeOrRef_rule0 () {
    delete m_shapeDefinition;
}
shapeOrRef_rule1::~shapeOrRef_rule1 () {
    delete m_shapeRef;
}
inlineShapeOrRef_rule0::~inlineShapeOrRef_rule0 () {
    delete m_inlineShapeDefinition;
}
inlineShapeOrRef_rule1::~inlineShapeOrRef_rule1 () {
    delete m_shapeRef;
}
shapeRef_rule0::~shapeRef_rule0 () {
    delete m_ATPNAME_LN;
}
shapeRef_rule1::~shapeRef_rule1 () {
    delete m_ATPNAME_NS;
}
shapeRef_rule2::~shapeRef_rule2 () {
    delete m_GT_AT;
    delete m_shapeExprLabel;
}
litNodeConstraint_rule0::~litNodeConstraint_rule0 () {
    delete m_IT_LITERAL;
    delete m_xsFacets;
}
litNodeConstraint_rule1::~litNodeConstraint_rule1 () {
    delete m_datatype;
    delete m_xsFacets;
}
litNodeConstraint_rule2::~litNodeConstraint_rule2 () {
    delete m_valueSet;
    delete m_xsFacets;
}
litNodeConstraint_rule3::~litNodeConstraint_rule3 () {
    delete m_numericFacets;
}
nonLitNodeConstraint_rule0::~nonLitNodeConstraint_rule0 () {
    delete m_nonLiteralKind;
    delete m_stringFacets;
}
nonLitNodeConstraint_rule1::~nonLitNodeConstraint_rule1 () {
    delete m_stringFacets;
}
nonLiteralKind_rule0::~nonLiteralKind_rule0 () {
    delete m_IT_IRI;
}
nonLiteralKind_rule1::~nonLiteralKind_rule1 () {
    delete m_IT_BNODE;
}
nonLiteralKind_rule2::~nonLiteralKind_rule2 () {
    delete m_IT_NONLITERAL;
}
xsFacet_rule0::~xsFacet_rule0 () {
    delete m_stringFacet;
}
xsFacet_rule1::~xsFacet_rule1 () {
    delete m_numericFacet;
}
stringFacet_rule0::~stringFacet_rule0 () {
    delete m_stringLength;
    delete m_INTEGER;
}
stringFacet_rule1::~stringFacet_rule1 () {
    delete m_REGEXP;
}
stringLength_rule0::~stringLength_rule0 () {
    delete m_IT_LENGTH;
}
stringLength_rule1::~stringLength_rule1 () {
    delete m_IT_MINLENGTH;
}
stringLength_rule2::~stringLength_rule2 () {
    delete m_IT_MAXLENGTH;
}
numericFacet_rule0::~numericFacet_rule0 () {
    delete m_numericRange;
    delete m_rawNumeric;
}
numericFacet_rule1::~numericFacet_rule1 () {
    delete m_numericLength;
    delete m_INTEGER;
}
numericRange_rule0::~numericRange_rule0 () {
    delete m_IT_MININCLUSIVE;
}
numericRange_rule1::~numericRange_rule1 () {
    delete m_IT_MINEXCLUSIVE;
}
numericRange_rule2::~numericRange_rule2 () {
    delete m_IT_MAXINCLUSIVE;
}
numericRange_rule3::~numericRange_rule3 () {
    delete m_IT_MAXEXCLUSIVE;
}
numericLength_rule0::~numericLength_rule0 () {
    delete m_IT_TOTALDIGITS;
}
numericLength_rule1::~numericLength_rule1 () {
    delete m_IT_FRACTIONDIGITS;
}
shapeDefinition::~shapeDefinition () {
    delete m_inlineShapeDefinition;
    delete m_annotations;
    delete m_semanticActions;
}
inlineShapeDefinition::~inlineShapeDefinition () {
    delete m_shapeQualifiers;
    delete m_GT_LCURLEY;
    delete m__QtripleExpression_E_Opt;
    delete m_GT_RCURLEY;
}
_QtripleExpression_E_Opt_rule0::~_QtripleExpression_E_Opt_rule0 () {
}
_QtripleExpression_E_Opt_rule1::~_QtripleExpression_E_Opt_rule1 () {
    delete m_tripleExpression;
}
shapeQualifiers::~shapeQualifiers () {
    delete m__O_QextraPropertySet_E_Or_QIT_CLOSED_E_Cs;
}
_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_rule0::~_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_rule0 () {
    delete m_extraPropertySet;
}
_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_rule1::~_O_QextraPropertySet_E_Or_QIT_CLOSED_E_C_rule1 () {
    delete m_IT_CLOSED;
}
extraPropertySet::~extraPropertySet () {
    delete m_IT_EXTRA;
    delete m_predicates;
}
tripleExpression::~tripleExpression () {
    delete m_oneOfTripleExpr;
}
oneOfTripleExpr_rule0::~oneOfTripleExpr_rule0 () {
    delete m_groupTripleExpr;
}
oneOfTripleExpr_rule1::~oneOfTripleExpr_rule1 () {
    delete m_multiElementOneOf;
}
multiElementOneOf::~multiElementOneOf () {
    delete m_groupTripleExpr;
    delete m__O_QGT_PIPE_E_S_QgroupTripleExpr_E_Cs;
}
_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C::~_O_QGT_PIPE_E_S_QgroupTripleExpr_E_C () {
    delete m_GT_PIPE;
    delete m_groupTripleExpr;
}
groupTripleExpr_rule0::~groupTripleExpr_rule0 () {
    delete m_singleElementGroup;
}
groupTripleExpr_rule1::~groupTripleExpr_rule1 () {
    delete m_multiElementGroup;
}
singleElementGroup::~singleElementGroup () {
    delete m_unaryTripleExpr;
    delete m__QGT_SEMI_E_Opt;
}
_QGT_SEMI_E_Opt_rule0::~_QGT_SEMI_E_Opt_rule0 () {
}
_QGT_SEMI_E_Opt_rule1::~_QGT_SEMI_E_Opt_rule1 () {
    delete m_GT_SEMI;
}
multiElementGroup::~multiElementGroup () {
    delete m_unaryTripleExpr;
    delete m__O_QGT_SEMI_E_S_QunaryTripleExpr_E_Cs;
    delete m__QGT_SEMI_E_Opt;
}
_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C::~_O_QGT_SEMI_E_S_QunaryTripleExpr_E_C () {
    delete m_GT_SEMI;
    delete m_unaryTripleExpr;
}
unaryTripleExpr_rule0::~unaryTripleExpr_rule0 () {
    delete m__Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt;
    delete m__O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C;
}
unaryTripleExpr_rule1::~unaryTripleExpr_rule1 () {
    delete m_include;
}
_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C::~_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C () {
    delete m_GT_DOLLAR;
    delete m_tripleExprLabel;
}
_Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt_rule0::~_Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt_rule0 () {
}
_Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt_rule1::~_Q_O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C_E_Opt_rule1 () {
    delete m__O_QGT_DOLLAR_E_S_QtripleExprLabel_E_C;
}
_O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C_rule0::~_O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C_rule0 () {
    delete m_tripleConstraint;
}
_O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C_rule1::~_O_QtripleConstraint_E_Or_QbracketedTripleExpr_E_C_rule1 () {
    delete m_bracketedTripleExpr;
}
bracketedTripleExpr::~bracketedTripleExpr () {
    delete m_GT_LPAREN;
    delete m_tripleExpression;
    delete m_GT_RPAREN;
    delete m__Qcardinality_E_Opt;
    delete m_annotations;
    delete m_semanticActions;
}
_Qcardinality_E_Opt_rule0::~_Qcardinality_E_Opt_rule0 () {
}
_Qcardinality_E_Opt_rule1::~_Qcardinality_E_Opt_rule1 () {
    delete m_cardinality;
}
tripleConstraint::~tripleConstraint () {
    delete m__QsenseFlags_E_Opt;
    delete m_predicate;
    delete m_inlineShapeExpression;
    delete m__Qcardinality_E_Opt;
    delete m_annotations;
    delete m_semanticActions;
}
_QsenseFlags_E_Opt_rule0::~_QsenseFlags_E_Opt_rule0 () {
}
_QsenseFlags_E_Opt_rule1::~_QsenseFlags_E_Opt_rule1 () {
    delete m_senseFlags;
}
cardinality_rule0::~cardinality_rule0 () {
    delete m_GT_TIMES;
}
cardinality_rule1::~cardinality_rule1 () {
    delete m_GT_PLUS;
}
cardinality_rule2::~cardinality_rule2 () {
    delete m_GT_OPT;
}
cardinality_rule3::~cardinality_rule3 () {
    delete m_REPEAT_RANGE;
}
senseFlags::~senseFlags () {
    delete m_GT_CARROT;
}
valueSet::~valueSet () {
    delete m_GT_LBRACKET;
    delete m_valueSetValues;
    delete m_GT_RBRACKET;
}
valueSetValue_rule0::~valueSetValue_rule0 () {
    delete m_iriRange;
}
valueSetValue_rule1::~valueSetValue_rule1 () {
    delete m_literalRange;
}
valueSetValue_rule2::~valueSetValue_rule2 () {
    delete m_languageRange;
}
valueSetValue_rule3::~valueSetValue_rule3 () {
    delete m_GT_DOT;
    delete m__O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C;
}
_O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule0::~_O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule0 () {
    delete m_iriExclusions;
}
_O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule1::~_O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule1 () {
    delete m_literalExclusions;
}
_O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule2::~_O_QiriExclusion_E_Plus_Or_QliteralExclusion_E_Plus_Or_QlanguageExclusion_E_Plus_C_rule2 () {
    delete m_languageExclusions;
}
iriRange::~iriRange () {
    delete m_iri;
    delete m__Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt;
}
_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C::~_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C () {
    delete m_GT_KINDA;
    delete m_iriExclusions;
}
_Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt_rule0::~_Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt_rule0 () {
}
_Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt_rule1::~_Q_O_QGT_KINDA_E_S_QiriExclusion_E_Star_C_E_Opt_rule1 () {
    delete m__O_QGT_KINDA_E_S_QiriExclusion_E_Star_C;
}
iriExclusion::~iriExclusion () {
    delete m_GT_MINUS;
    delete m_iri;
    delete m__QGT_KINDA_E_Opt;
}
_QGT_KINDA_E_Opt_rule0::~_QGT_KINDA_E_Opt_rule0 () {
}
_QGT_KINDA_E_Opt_rule1::~_QGT_KINDA_E_Opt_rule1 () {
    delete m_GT_KINDA;
}
literalRange::~literalRange () {
    delete m_literal;
    delete m__Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt;
}
_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C::~_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C () {
    delete m_GT_KINDA;
    delete m_literalExclusions;
}
_Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt_rule0::~_Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt_rule0 () {
}
_Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt_rule1::~_Q_O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C_E_Opt_rule1 () {
    delete m__O_QGT_KINDA_E_S_QliteralExclusion_E_Star_C;
}
literalExclusion::~literalExclusion () {
    delete m_GT_MINUS;
    delete m_literal;
    delete m__QGT_KINDA_E_Opt;
}
languageRange_rule0::~languageRange_rule0 () {
    delete m_LANGTAG;
    delete m__Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt;
}
languageRange_rule1::~languageRange_rule1 () {
    delete m_GT_AT;
    delete m_GT_KINDA;
    delete m_languageExclusions;
}
_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C::~_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C () {
    delete m_GT_KINDA;
    delete m_languageExclusions;
}
_Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt_rule0::~_Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt_rule0 () {
}
_Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt_rule1::~_Q_O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C_E_Opt_rule1 () {
    delete m__O_QGT_KINDA_E_S_QlanguageExclusion_E_Star_C;
}
languageExclusion::~languageExclusion () {
    delete m_GT_MINUS;
    delete m_LANGTAG;
    delete m__QGT_KINDA_E_Opt;
}
include::~include () {
    delete m_GT_AMP;
    delete m_tripleExprLabel;
}
annotation::~annotation () {
    delete m_GT_DIVIDE_DIVIDE;
    delete m_predicate;
    delete m__O_Qiri_E_Or_Qliteral_E_C;
}
_O_Qiri_E_Or_Qliteral_E_C_rule0::~_O_Qiri_E_Or_Qliteral_E_C_rule0 () {
    delete m_iri;
}
_O_Qiri_E_Or_Qliteral_E_C_rule1::~_O_Qiri_E_Or_Qliteral_E_C_rule1 () {
    delete m_literal;
}
semanticActions::~semanticActions () {
    delete m_codeDecls;
}
codeDecl::~codeDecl () {
    delete m_GT_MODULO;
    delete m_iri;
    delete m__O_QCODE_E_Or_QGT_MODULO_E_C;
}
_O_QCODE_E_Or_QGT_MODULO_E_C_rule0::~_O_QCODE_E_Or_QGT_MODULO_E_C_rule0 () {
    delete m_CODE;
}
_O_QCODE_E_Or_QGT_MODULO_E_C_rule1::~_O_QCODE_E_Or_QGT_MODULO_E_C_rule1 () {
    delete m_GT_MODULO;
}
literal_rule0::~literal_rule0 () {
    delete m_rdfLiteral;
}
literal_rule1::~literal_rule1 () {
    delete m_numericLiteral;
}
literal_rule2::~literal_rule2 () {
    delete m_booleanLiteral;
}
predicate_rule0::~predicate_rule0 () {
    delete m_iri;
}
predicate_rule1::~predicate_rule1 () {
    delete m_RDF_TYPE;
}
datatype::~datatype () {
    delete m_iri;
}
shapeExprLabel_rule0::~shapeExprLabel_rule0 () {
    delete m_iri;
}
shapeExprLabel_rule1::~shapeExprLabel_rule1 () {
    delete m_blankNode;
}
tripleExprLabel_rule0::~tripleExprLabel_rule0 () {
    delete m_iri;
}
tripleExprLabel_rule1::~tripleExprLabel_rule1 () {
    delete m_blankNode;
}
rawNumeric_rule0::~rawNumeric_rule0 () {
    delete m_INTEGER;
}
rawNumeric_rule1::~rawNumeric_rule1 () {
    delete m_DECIMAL;
}
rawNumeric_rule2::~rawNumeric_rule2 () {
    delete m_DOUBLE;
}
numericLiteral::~numericLiteral () {
    delete m_rawNumeric;
}
rdfLiteral_rule0::~rdfLiteral_rule0 () {
    delete m_langString;
}
rdfLiteral_rule1::~rdfLiteral_rule1 () {
    delete m_string;
    delete m__Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt;
}
_O_QGT_DTYPE_E_S_Qdatatype_E_C::~_O_QGT_DTYPE_E_S_Qdatatype_E_C () {
    delete m_GT_DTYPE;
    delete m_datatype;
}
_Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt_rule0::~_Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt_rule0 () {
}
_Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt_rule1::~_Q_O_QGT_DTYPE_E_S_Qdatatype_E_C_E_Opt_rule1 () {
    delete m__O_QGT_DTYPE_E_S_Qdatatype_E_C;
}
booleanLiteral_rule0::~booleanLiteral_rule0 () {
    delete m_IT_true;
}
booleanLiteral_rule1::~booleanLiteral_rule1 () {
    delete m_IT_false;
}
string_rule0::~string_rule0 () {
    delete m_STRING_LITERAL1;
}
string_rule1::~string_rule1 () {
    delete m_STRING_LITERAL_LONG1;
}
string_rule2::~string_rule2 () {
    delete m_STRING_LITERAL2;
}
string_rule3::~string_rule3 () {
    delete m_STRING_LITERAL_LONG2;
}
langString_rule0::~langString_rule0 () {
    delete m_LANG_STRING_LITERAL1;
}
langString_rule1::~langString_rule1 () {
    delete m_LANG_STRING_LITERAL_LONG1;
}
langString_rule2::~langString_rule2 () {
    delete m_LANG_STRING_LITERAL2;
}
langString_rule3::~langString_rule3 () {
    delete m_LANG_STRING_LITERAL_LONG2;
}
iri_rule0::~iri_rule0 () {
    delete m_IRIREF;
}
iri_rule1::~iri_rule1 () {
    delete m_prefixedName;
}
prefixedName_rule0::~prefixedName_rule0 () {
    delete m_PNAME_LN;
}
prefixedName_rule1::~prefixedName_rule1 () {
    delete m_PNAME_NS;
}
blankNode::~blankNode () {
    delete m_BLANK_NODE_LABEL;
}

} // namespace shexParseToXml

/* END Driver */

/* START main */

#include <stdio.h>
// #include "shex-parse-to-XMLFrob.h"

#include <stdlib.h>
#include <ostream>
#include <fstream>

#ifdef HAVE_BOOST
#include <boost/version.hpp>
#include <boost/iostreams/stream.hpp>
#include <boost/iostreams/device/file_descriptor.hpp>
#endif /* HAVE_BOOST */

bool yacker::TransparentGroupProductions = false;
bool yacker::TransparentMultiplicityProductions = false;

int main(int argc,char** argv) {

    shexParseToXml::shexParseToXmlContext context;
    shexParseToXml::Driver driver(context);

#ifdef HAVE_BOOST
    boost::iostreams::stream_buffer<boost::iostreams::file_descriptor_sink>* buf = NULL;
    if (char* tmp = getenv("TRACE_FD")) {
	std::istringstream is(tmp);
	int fd;
	is >> fd;
	buf = new boost::iostreams::stream_buffer<boost::iostreams::file_descriptor_sink>
	    (fd
#if BOOST_VERSION >= 104400
	     , boost::iostreams::close_handle
#endif
	     );
	if (!(yacker::_Trace = new std::ostream(buf)))
	    std::cerr << "couldn't open trace fd " << fd << std::endl;
    }
#endif /* HAVE_BOOST */

    std::istream* yyin;
    const char* inputId;
    if (argc > 1) {
	inputId = argv[1];
	yyin = new std::ifstream(argv[1], std::ifstream::in);
    } else {
	inputId = "<stdin>";
	yyin = &std::cin;
    }

    std::string str(inputId);
    int result = driver.parse_stream(*yyin, str);

    if (argc > 1)
	delete yyin;

    if (result)
	std::cerr << "Error: " << inputId << " did not contain a valid ShEx 2.1 string." << std::endl;
    else {
	try {
	    std::string xml = driver.root->toXml(0);
	    std::cout << xml;
	    delete driver.root;
	} catch (ProgramFlowException& e) {
	    std::cout << "Standard exception: " << e.what() << std::endl;
	} catch (std::exception& e) {
	    std::cout << "Standard exception: " << e.what() << std::endl;
	}
    }

#ifdef HAVE_BOOST
    delete buf; // this delete breaks cout!
#endif /* HAVE_BOOST */

    if (yacker::_Trace)
	delete yacker::_Trace;

    return result;
};

/* END main */

