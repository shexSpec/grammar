/* $Id: Langname_Scanner.ll,v 1.1 2008/04/06 17:10:46 eric Exp spec_2_1Scanner.ll 28 2007-08-20 10:27:39Z tb $ -*- mode: c++ -*- */
/** \file shex-parse-to-XMLScanner.ll Define the Flex lexical scanner */

%{ /*** C/C++ Declarations ***/

#include "shex-parse-to-XMLParser.hh"
#include "shex-parse-to-XMLScanner.hh"

/* import the parser's token type into a local typedef */
typedef shexParseToXml::shexParseToXmlParser::token token;
typedef shexParseToXml::shexParseToXmlParser::token_type token_type;

/* Work around an incompatibility in flex (at least versions 2.5.31 through
 * 2.5.33): it generates code that does not conform to C89.  See Debian bug
 * 333231 <http://bugs.debian.org/cgi-bin/bugreport.cgi?bug=333231>.  */
#undef yywrap
#define yywrap()	1

/* By default yylex returns int, we use token_type. Unfortunately yyterminate
 * by default returns 0, which is not of token_type. */
#define yyterminate() return token::__EOF__

/* This disables inclusion of unistd.h, which is not available under Visual C++
 * on Win32. The C++ scanner uses STL streams instead. */
#define YY_NO_UNISTD_H

%}

/*** Flex Declarations and Options ***/

/* enable c++ scanner class generation */
%option c++

/* change the name of the scanner class. results in "shexParseToXmlFlexLexer" */
%option prefix="shexParseToXml"

/* the manual says "somewhat more optimized" */
%option batch

/* enable scanner to generate debug output. disable this for release
 * versions. */
%option debug

/* no support for include files is planned */
%option noyywrap nounput 

/* enables the use of start condition stacks */
%option stack

/* The following paragraph suffices to track locations accurately. Each time
 * yylex is invoked, the begin position is moved onto the end position. */
%{
#define YY_USER_ACTION  yylloc->columns(yyleng);
%}

/* START patterns for shexParseToXml terminals */
IT_BASE			[Bb][Aa][Ss][Ee]
IT_PREFIX		[Pp][Rr][Ee][Ff][Ii][Xx]
IT_IMPORT		[Ii][Mm][Pp][Oo][Rr][Tt]
IT_START		[Ss][Tt][Aa][Rr][Tt]
GT_EQUAL		"="
IT_EXTERNAL		[Ee][Xx][Tt][Ee][Rr][Nn][Aa][Ll]
IT_OR			[Oo][Rr]
IT_AND			[Aa][Nn][Dd]
IT_NOT			[Nn][Oo][Tt]
GT_LPAREN		"("
GT_RPAREN		")"
GT_DOT			"."
GT_AT			"@"
IT_LITERAL		[Ll][Ii][Tt][Ee][Rr][Aa][Ll]
IT_IRI			[Ii][Rr][Ii]
IT_BNODE		[Bb][Nn][Oo][Dd][Ee]
IT_NONLITERAL		[Nn][Oo][Nn][Ll][Ii][Tt][Ee][Rr][Aa][Ll]
IT_LENGTH		[Ll][Ee][Nn][Gg][Tt][Hh]
IT_MINLENGTH		[Mm][Ii][Nn][Ll][Ee][Nn][Gg][Tt][Hh]
IT_MAXLENGTH		[Mm][Aa][Xx][Ll][Ee][Nn][Gg][Tt][Hh]
IT_MININCLUSIVE		[Mm][Ii][Nn][Ii][Nn][Cc][Ll][Uu][Ss][Ii][Vv][Ee]
IT_MINEXCLUSIVE		[Mm][Ii][Nn][Ee][Xx][Cc][Ll][Uu][Ss][Ii][Vv][Ee]
IT_MAXINCLUSIVE		[Mm][Aa][Xx][Ii][Nn][Cc][Ll][Uu][Ss][Ii][Vv][Ee]
IT_MAXEXCLUSIVE		[Mm][Aa][Xx][Ee][Xx][Cc][Ll][Uu][Ss][Ii][Vv][Ee]
IT_TOTALDIGITS		[Tt][Oo][Tt][Aa][Ll][Dd][Ii][Gg][Ii][Tt][Ss]
IT_FRACTIONDIGITS	[Ff][Rr][Aa][Cc][Tt][Ii][Oo][Nn][Dd][Ii][Gg][Ii][Tt][Ss]
GT_LCURLEY		"{"
GT_RCURLEY		"}"
IT_CLOSED		[Cc][Ll][Oo][Ss][Ee][Dd]
IT_EXTRA		[Ee][Xx][Tt][Rr][Aa]
GT_PIPE			"|"
GT_SEMI			";"
GT_DOLLAR		"$"
GT_TIMES		"*"
GT_PLUS			"+"
GT_OPT			"?"
GT_CARROT		"^"
GT_LBRACKET		"\["
GT_RBRACKET		"\]"
GT_KINDA		"~"
GT_MINUS		"-"
GT_AMP			"&"
GT_DIVIDE_DIVIDE	"//"
GT_MODULO		"%"
GT_DTYPE		"^^"
IT_true			"true"
IT_false		"false"
RDF_TYPE		"a"
REGEXP_FLAGS		[imsx]+
LANGTAG			"@"[A-Za-z]+("-"[0-9A-Za-z]+)*
INTEGER			[+-]?[0-9]+
REPEAT_RANGE		"{"{INTEGER}(","({INTEGER}|"*")?)?"}"
DECIMAL			[+-]?[0-9]*"."[0-9]+
EXPONENT		[Ee][+-]?[0-9]+
DOUBLE			([+-])?(([0-9]+"."[0-9]*{EXPONENT})|("."?[0-9]+{EXPONENT}))
ECHAR			"\\"[\"'\\bfnrt]
PN_CHARS_BASE		[A-Z]|[a-z]|\xC3[\x80-\x96]|\xC3[\x98-\xB6]|\xC3[\xB8-\xBF]|[\xC4-\xCB][\x80-\xBF]|\xCD[\xB0-\xBD]|\xCD\xBF|[\xCE-\xDF][\x80-\xBF]|\xE0[\xA0-\xBF][\x80-\xBF]|\xE1[\x80-\xBF][\x80-\xBF]|\xE2\x80[\x8C-\x8D]|\xE2(\x81[\xB0-\xBF]|[\x82-\x85][\x80-\xBF]|\x86[\x80-\x8F])|\xE2([\xB0-\xBE][\x80-\xBF]|\xBF[\x80-\xAF])|\xE3(\x80[\x81-\xBF]|[\x81-\xBF][\x80-\xBF])|[\xE4-\xEC][\x80-\xBF][\x80-\xBF]|[\xE1-\xEC][\x80-\xBF][\x80-\xBF]|\xED([\x80-\x9F][\x80-\xBF])|\xEF([\xA4-\xB6][\x80-\xBF]|\xB7[\x80-\x8F])|\xEF(\xB7[\xB0-\xBF]|[\xB8-\xBE][\x80-\xBF]|\xBF[\x80-\xBD])|\xF0[\x90-\xBF][\x80-\xBF][\x80-\xBF]|[\xF1-\xF3][\x80-\xBF][\x80-\xBF][\x80-\xBF]
/*
PN_CHARS_BASE		([A-Z])|(([a-z])|(((\xC3[\x80-\x96]))|(((\xC3[\x98-\xB6]))|(((\xC3[\xB8-\xBF])|([\xC4-\xCB][\x80-\xBF]))|(((\xCD[\xB0-\xBD]))|(((\xCD\xBF)|([\xCE-\xDF][\x80-\xBF])|(\xE0([\xA0-\xBF][\x80-\xBF]))|(\xE1([\x80-\xBF][\x80-\xBF])))|(((\xE2(\x80[\x8C-\x8D])))|(((\xE2(\x81[\xB0-\xBF])|([\x82-\x85][\x80-\xBF])|(\x86[\x80-\x8F])))|(((\xE2([\xB0-\xBE][\x80-\xBF])|(\xBF[\x80-\xAF])))|(((\xE3(\x80[\x81-\xBF])|([\x81-\xBF][\x80-\xBF]))|([\xE4-\xEC][\x80-\xBF][\x80-\xBF])|([\xE1-\xEC][\x80-\xBF][\x80-\xBF])|(\xED([\x80-\x9F][\x80-\xBF])))|(((\xEF([\xA4-\xB6][\x80-\xBF])|(\xB7[\x80-\x8F])))|(((\xEF(\xB7[\xB0-\xBF])|([\xB8-\xBE][\x80-\xBF])|(\xBF[\x80-\xBD])))|((\xF0([\x90-\xBF][\x80-\xBF][\x80-\xBF]))|([\xF1-\xF3][\x80-\xBF][\x80-\xBF][\x80-\xBF]))))))))))))))|\xc2\xb7\xcc\x80\xcd\xaf\xe2\x80\xbf\x2e\xe2\x81\x80
*/
PN_CHARS_U		{PN_CHARS_BASE}|"_"
PN_CHARS		{PN_CHARS_U}|"-"|[0-9]|\xC2\xB7|\xCC[\x80-\xBF]|\xCD[\x80-\xAF]|\xE2(\x80\xBF|\x81\x80)
BLANK_NODE_LABEL	"_:"({PN_CHARS_U}|[0-9])(({PN_CHARS}|".")*{PN_CHARS})?
PN_PREFIX		{PN_CHARS_BASE}((({PN_CHARS}|"."))*{PN_CHARS})?
PNAME_NS		(({PN_PREFIX}))?":"
ATPNAME_NS		"@"({PNAME_NS})
HEX			[0-9]|[A-F]|[a-f]
PERCENT			"%"{HEX}{HEX}
UCHAR			("\\u"({HEX})({HEX})({HEX})({HEX}))|("\\U"({HEX})({HEX})({HEX})({HEX})({HEX})({HEX})({HEX})({HEX}))
_UCODE			[\xC2-\xDF][\x80-\xBF]|\xE0[\xA0-\xBF][\x80-\xBF]|[\xE1-\xEC][\x80-\xBF][\x80-\xBF]|[\xE1-\xEC][\x80-\xBF][\x80-\xBF]|\xED[\x80-\x9F][\x80-\xBF]|[\xEE-\xEF][\x80-\xBF][\x80-\xBF]|\xF0[\x90-\xBF][\x80-\xBF][\x80-\xBF]|[\xF1-\xF3][\x80-\xBF][\x80-\xBF][\x80-\xBF]|\xF4([\x80-\x8E][\x80-\xBF][\x80-\xBF]|\x8F([\x80-\xBE][\x80-\xBF]|\xBF[\x80-\xBD]))
STRING_LITERAL1		"'"([\x00-\x09\x0B-\x0C\x0E-\x26\x28-\x5B\x5D-\x7F]|{_UCODE}|{ECHAR}|{UCHAR})*"'"
/*
STRING_LITERAL1		"'"((([\x00-\t\x0B-\x0C\x0E-&(-\[\]-\x7F]|([\xC2-\xDF][\x80-\xBF])|(\xE0([\xA0-\xBF][\x80-\xBF]))|([\xE1-\xEC][\x80-\xBF][\x80-\xBF])|([\xE1-\xEC][\x80-\xBF][\x80-\xBF])|(\xED([\x80-\x9F][\x80-\xBF]))|([\xEE-\xEF][\x80-\xBF][\x80-\xBF])|(\xF0([\x90-\xBF][\x80-\xBF][\x80-\xBF]))|([\xF1-\xF3][\x80-\xBF][\x80-\xBF][\x80-\xBF])|(\xF4([\x80-\x8E][\x80-\xBF][\x80-\xBF])|(\x8F([\x80-\xBE][\x80-\xBF])|(\xBF[\x80-\xBD])))])|((({ECHAR}))|(({UCHAR})))))*"'"
*/
STRING_LITERAL2		"\""([\x00-\x09\x0B-\x0C\x0E-\x21\x23-\x5B\x5D-\x7F]|{_UCODE}|{ECHAR}|{UCHAR})*"\""
STRING_LITERAL_LONG1	"'''"("'"|"''")?([\x00-\x26\x28-\x5B\x5D-\x7F]|{_UCODE}]|{ECHAR}|{UCHAR})*"'''"
STRING_LITERAL_LONG2	"\"\"\""("\""|"\"\"")?([\x00-\x21\x23-\x5B\x5D-\x7F]|{_UCODE}]|{ECHAR}|{UCHAR})*"\"\"\""
LANG_STRING_LITERAL1		"'"([\x00-\x09\x0B-\x0C\x0E-\x26\x28-\x5B\x5D-\x7F]|{_UCODE}|{ECHAR}|{UCHAR})*"'"{LANGTAG}
LANG_STRING_LITERAL2		"\""([\x00-\x09\x0B-\x0C\x0E-\x21\x23-\x5B\x5D-\x7F]|{_UCODE}|{ECHAR}|{UCHAR})*"\""{LANGTAG}
LANG_STRING_LITERAL_LONG1	"'''"("'"|"''")?([\x00-\x26\x28-\x5B\x5D-\x7F]|{_UCODE}]|{ECHAR}|{UCHAR})*"'''"{LANGTAG}
LANG_STRING_LITERAL_LONG2	"\"\"\""("\""|"\"\"")?([\x00-\x21\x23-\x5B\x5D-\x7F]|{_UCODE}]|{ECHAR}|{UCHAR})*"\"\"\""{LANGTAG}
CODE			"{"([\x00-\x24\x26-\x5B\x5D-\x7F]|{_UCODE}|"\\"[%\\]|{UCHAR})*"%""}"
IRIREF			"<"([\x00-\x21\x23-\x2F\x30-\x3B\x3D\x3F-\x5B\x5D\x5Fa-z\x7E-\x7F]|{_UCODE}|{UCHAR})*">"
/*
IRIREF			"<"((([\x00-!$-/1-;=?-\[\]_a-z~-\x7F]|([\xC2-\xDF][\x80-\xBF])|(\xE0([\xA0-\xBF][\x80-\xBF]))|([\xE1-\xEC][\x80-\xBF][\x80-\xBF])|([\xE1-\xEC][\x80-\xBF][\x80-\xBF])|(\xED([\x80-\x9F][\x80-\xBF]))|([\xEE-\xEF][\x80-\xBF][\x80-\xBF])|(\xF0([\x90-\xBF][\x80-\xBF][\x80-\xBF]))|([\xF1-\xF3][\x80-\xBF][\x80-\xBF][\x80-\xBF])|(\xF4([\x80-\x8E][\x80-\xBF][\x80-\xBF])|(\x8F([\x80-\xBE][\x80-\xBF])|(\xBF[\x80-\xBD])))])|(({UCHAR}))))*">"
*/
REGEXP			"/"([^\x2F\x5C\x0A\x0D]|{_UCODE}|"\\"[nrt\x5C\x7C\x2E\x3F\x2A\x2B\x28\x29\x7B\x7D\x24\x2D\x5B\x5D\x5E/]|{UCHAR})+"/"[smix]*
/*
REGEXP			"/"((([\x00-\t\x0B-\x0C\x0E-.0-\[\]-\x7F]|([\xC2-\xDF][\x80-\xBF])|(\xE0([\xA0-\xBF][\x80-\xBF]))|([\xE1-\xEC][\x80-\xBF][\x80-\xBF])|([\xE1-\xEC][\x80-\xBF][\x80-\xBF])|(\xED([\x80-\x9F][\x80-\xBF]))|([\xEE-\xEF][\x80-\xBF][\x80-\xBF])|(\xF0([\x90-\xBF][\x80-\xBF][\x80-\xBF]))|([\xF1-\xF3][\x80-\xBF][\x80-\xBF][\x80-\xBF])|(\xF4([\x80-\x8E][\x80-\xBF][\x80-\xBF])|(\x8F([\x80-\xBE][\x80-\xBF])|(\xBF[\x80-\xBD])))])|(("\\"[$-\\()*+./?\\^nrt{|}])|(({UCHAR})))))+"/"(({REGEXP_FLAGS}))?
*/

PN_LOCAL_ESC		"\\"("_"|"~"|"."|"-"|"!"|"$"|"&"|"'"|"("|")"|"*"|"+"|","|";"|"="|"/"|"?"|"#"|"@"|"%")
PLX			{PERCENT}|{PN_LOCAL_ESC}
PN_LOCAL		({PN_CHARS_U}|":"|[0-9]|{PLX})({PN_CHARS}|"."|":"|{PLX})*
/*({PN_CHARS}|"."|":"|{PLX})**/
/*
PN_LOCAL		({PN_CHARS_U}|":"|[0-9]|{PLX})({PN_CHARS}|"."|":"|{PLX})*
*/
PNAME_LN		{PNAME_NS}{PN_LOCAL}
ATPNAME_LN		"@"{PNAME_LN}
PASSED_TOKENS		[\t\n\r ]+|"#"([\x00-\t\x0B-\x0C\x0E-\x7F]|{_UCODE})*|("/*"([\x00-\x29\x2B-\x7F]|{_UCODE}|"*"([\x00-.0-\x7F]|{_UCODE}|"\\/"))*"*/")

/* END patterns for shexParseToXml terminals */

/* START semantic actions for shexParseToXml terminals */
%%
{PASSED_TOKENS}		{ /* yylloc->step(); @@ needed? useful? */ }
{IT_BASE}		{yylval->p_IT_BASE = new IT_BASE(yytext); return token::IT_BASE;}
{IT_PREFIX}		{yylval->p_IT_PREFIX = new IT_PREFIX(yytext); return token::IT_PREFIX;}
{IT_IMPORT}		{yylval->p_IT_IMPORT = new IT_IMPORT(yytext); return token::IT_IMPORT;}
{IT_START}		{yylval->p_IT_START = new IT_START(yytext); return token::IT_START;}
{GT_EQUAL}		{yylval->p_GT_EQUAL = new GT_EQUAL(yytext); return token::GT_EQUAL;}
{IT_EXTERNAL}		{yylval->p_IT_EXTERNAL = new IT_EXTERNAL(yytext); return token::IT_EXTERNAL;}
{IT_OR}			{yylval->p_IT_OR = new IT_OR(yytext); return token::IT_OR;}
{IT_AND}		{yylval->p_IT_AND = new IT_AND(yytext); return token::IT_AND;}
{IT_NOT}		{yylval->p_IT_NOT = new IT_NOT(yytext); return token::IT_NOT;}
{GT_LPAREN}		{yylval->p_GT_LPAREN = new GT_LPAREN(yytext); return token::GT_LPAREN;}
{GT_RPAREN}		{yylval->p_GT_RPAREN = new GT_RPAREN(yytext); return token::GT_RPAREN;}
{GT_DOT}		{yylval->p_GT_DOT = new GT_DOT(yytext); return token::GT_DOT;}
{GT_AT}			{yylval->p_GT_AT = new GT_AT(yytext); return token::GT_AT;}
{IT_LITERAL}		{yylval->p_IT_LITERAL = new IT_LITERAL(yytext); return token::IT_LITERAL;}
{IT_IRI}		{yylval->p_IT_IRI = new IT_IRI(yytext); return token::IT_IRI;}
{IT_BNODE}		{yylval->p_IT_BNODE = new IT_BNODE(yytext); return token::IT_BNODE;}
{IT_NONLITERAL}		{yylval->p_IT_NONLITERAL = new IT_NONLITERAL(yytext); return token::IT_NONLITERAL;}
{IT_LENGTH}		{yylval->p_IT_LENGTH = new IT_LENGTH(yytext); return token::IT_LENGTH;}
{IT_MINLENGTH}		{yylval->p_IT_MINLENGTH = new IT_MINLENGTH(yytext); return token::IT_MINLENGTH;}
{IT_MAXLENGTH}		{yylval->p_IT_MAXLENGTH = new IT_MAXLENGTH(yytext); return token::IT_MAXLENGTH;}
{IT_MININCLUSIVE}	{yylval->p_IT_MININCLUSIVE = new IT_MININCLUSIVE(yytext); return token::IT_MININCLUSIVE;}
{IT_MINEXCLUSIVE}	{yylval->p_IT_MINEXCLUSIVE = new IT_MINEXCLUSIVE(yytext); return token::IT_MINEXCLUSIVE;}
{IT_MAXINCLUSIVE}	{yylval->p_IT_MAXINCLUSIVE = new IT_MAXINCLUSIVE(yytext); return token::IT_MAXINCLUSIVE;}
{IT_MAXEXCLUSIVE}	{yylval->p_IT_MAXEXCLUSIVE = new IT_MAXEXCLUSIVE(yytext); return token::IT_MAXEXCLUSIVE;}
{IT_TOTALDIGITS}	{yylval->p_IT_TOTALDIGITS = new IT_TOTALDIGITS(yytext); return token::IT_TOTALDIGITS;}
{IT_FRACTIONDIGITS}	{yylval->p_IT_FRACTIONDIGITS = new IT_FRACTIONDIGITS(yytext); return token::IT_FRACTIONDIGITS;}
{GT_LCURLEY}		{yylval->p_GT_LCURLEY = new GT_LCURLEY(yytext); return token::GT_LCURLEY;}
{GT_RCURLEY}		{yylval->p_GT_RCURLEY = new GT_RCURLEY(yytext); return token::GT_RCURLEY;}
{IT_CLOSED}		{yylval->p_IT_CLOSED = new IT_CLOSED(yytext); return token::IT_CLOSED;}
{IT_EXTRA}		{yylval->p_IT_EXTRA = new IT_EXTRA(yytext); return token::IT_EXTRA;}
{GT_PIPE}		{yylval->p_GT_PIPE = new GT_PIPE(yytext); return token::GT_PIPE;}
{GT_SEMI}		{yylval->p_GT_SEMI = new GT_SEMI(yytext); return token::GT_SEMI;}
{GT_DOLLAR}		{yylval->p_GT_DOLLAR = new GT_DOLLAR(yytext); return token::GT_DOLLAR;}
{GT_TIMES}		{yylval->p_GT_TIMES = new GT_TIMES(yytext); return token::GT_TIMES;}
{GT_PLUS}		{yylval->p_GT_PLUS = new GT_PLUS(yytext); return token::GT_PLUS;}
{GT_OPT}		{yylval->p_GT_OPT = new GT_OPT(yytext); return token::GT_OPT;}
{GT_CARROT}		{yylval->p_GT_CARROT = new GT_CARROT(yytext); return token::GT_CARROT;}
{GT_LBRACKET}		{yylval->p_GT_LBRACKET = new GT_LBRACKET(yytext); return token::GT_LBRACKET;}
{GT_RBRACKET}		{yylval->p_GT_RBRACKET = new GT_RBRACKET(yytext); return token::GT_RBRACKET;}
{GT_KINDA}		{yylval->p_GT_KINDA = new GT_KINDA(yytext); return token::GT_KINDA;}
{GT_MINUS}		{yylval->p_GT_MINUS = new GT_MINUS(yytext); return token::GT_MINUS;}
{GT_AMP}		{yylval->p_GT_AMP = new GT_AMP(yytext); return token::GT_AMP;}
{GT_DIVIDE_DIVIDE}	{yylval->p_GT_DIVIDE_DIVIDE = new GT_DIVIDE_DIVIDE(yytext); return token::GT_DIVIDE_DIVIDE;}
{GT_MODULO}		{yylval->p_GT_MODULO = new GT_MODULO(yytext); return token::GT_MODULO;}
{GT_DTYPE}		{yylval->p_GT_DTYPE = new GT_DTYPE(yytext); return token::GT_DTYPE;}
{IT_true}		{yylval->p_IT_true = new IT_true(yytext); return token::IT_true;}
{IT_false}		{yylval->p_IT_false = new IT_false(yytext); return token::IT_false;}
{CODE}		{yylval->p_CODE = new CODE(yytext); return token::CODE;}
{REPEAT_RANGE}		{yylval->p_REPEAT_RANGE = new REPEAT_RANGE(yytext); return token::REPEAT_RANGE;}
{RDF_TYPE}		{yylval->p_RDF_TYPE = new RDF_TYPE(yytext); return token::RDF_TYPE;}
{IRIREF}		{yylval->p_IRIREF = new IRIREF(yytext); return token::IRIREF;}
{PNAME_NS}		{yylval->p_PNAME_NS = new PNAME_NS(yytext); return token::PNAME_NS;}
{PNAME_LN}		{yylval->p_PNAME_LN = new PNAME_LN(yytext); return token::PNAME_LN;}
{ATPNAME_NS}		{yylval->p_ATPNAME_NS = new ATPNAME_NS(yytext); return token::ATPNAME_NS;}
{ATPNAME_LN}		{yylval->p_ATPNAME_LN = new ATPNAME_LN(yytext); return token::ATPNAME_LN;}
{REGEXP}		{yylval->p_REGEXP = new REGEXP(yytext); return token::REGEXP;}
{BLANK_NODE_LABEL}	{yylval->p_BLANK_NODE_LABEL = new BLANK_NODE_LABEL(yytext); return token::BLANK_NODE_LABEL;}
{LANGTAG}		{yylval->p_LANGTAG = new LANGTAG(yytext); return token::LANGTAG;}
{INTEGER}		{yylval->p_INTEGER = new INTEGER(yytext); return token::INTEGER;}
{DECIMAL}		{yylval->p_DECIMAL = new DECIMAL(yytext); return token::DECIMAL;}
{DOUBLE}		{yylval->p_DOUBLE = new DOUBLE(yytext); return token::DOUBLE;}
{STRING_LITERAL1}	{yylval->p_STRING_LITERAL1 = new STRING_LITERAL1(yytext); return token::STRING_LITERAL1;}
{STRING_LITERAL2}	{yylval->p_STRING_LITERAL2 = new STRING_LITERAL2(yytext); return token::STRING_LITERAL2;}
{STRING_LITERAL_LONG1}	{yylval->p_STRING_LITERAL_LONG1 = new STRING_LITERAL_LONG1(yytext); return token::STRING_LITERAL_LONG1;}
{STRING_LITERAL_LONG2}	{yylval->p_STRING_LITERAL_LONG2 = new STRING_LITERAL_LONG2(yytext); return token::STRING_LITERAL_LONG2;}
{LANG_STRING_LITERAL1}		{yylval->p_LANG_STRING_LITERAL1 = new LANG_STRING_LITERAL1(yytext); return token::LANG_STRING_LITERAL1;}
{LANG_STRING_LITERAL2}		{yylval->p_LANG_STRING_LITERAL2 = new LANG_STRING_LITERAL2(yytext); return token::LANG_STRING_LITERAL2;}
{LANG_STRING_LITERAL_LONG1}	{yylval->p_LANG_STRING_LITERAL_LONG1 = new LANG_STRING_LITERAL_LONG1(yytext); return token::LANG_STRING_LITERAL_LONG1;}
{LANG_STRING_LITERAL_LONG2}	{yylval->p_LANG_STRING_LITERAL_LONG2 = new LANG_STRING_LITERAL_LONG2(yytext); return token::LANG_STRING_LITERAL_LONG2;}

<<EOF>>			{ yyterminate();}
.			{ printf( "Unrecognized character: %s\n", yytext ); yyterminate(); }
%%
/* END semantic actions for shexParseToXml terminals */

/* START shexParseToXmlScanner */
namespace shexParseToXml {

shexParseToXmlScanner::shexParseToXmlScanner(std::istream* in,
		 std::ostream* out)
    : shexParseToXmlFlexLexer(in, out)
{
}

shexParseToXmlScanner::~shexParseToXmlScanner()
{
}

void shexParseToXmlScanner::set_debug(bool b)
{
    yy_flex_debug = b;
}

}
/* END shexParseToXmlScanner */

/* This implementation of shexParseToXmlFlexLexer::yylex() is required to fill the
 * vtable of the class shexParseToXmlFlexLexer. We define the scanner's main yylex
 * function via YY_DECL to reside in the shexParseToXmlScanner class instead. */

#ifdef yylex
#undef yylex
#endif

int shexParseToXmlFlexLexer::yylex()
{
    std::cerr << "in shexParseToXmlFlexLexer::yylex() !" << std::endl;
    return 0;
}

