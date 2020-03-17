/* $Id: Langname_Scanner.ll,v 1.1 2008/04/06 17:10:46 eric Exp spec_2_1Scanner.ll 28 2007-08-20 10:27:39Z tb $ -*- mode: c++ -*- */
/** \file shex-to-jsonScanner.ll Define the Flex lexical scanner */

%{ /*** C/C++ Declarations ***/

#include "shex-to-jsonParser.hh"
#include "shex-to-jsonScanner.hh"

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
/*{IRIREF}		{yylval->str = yytext; return token::IRIREF;}*/

/* START semantic actions for shexParseToXml terminals */
%%
{PASSED_TOKENS}		{ /* yylloc->step(); @@ needed? useful? */ }
{IT_BASE}		{return token::IT_BASE;}
{IT_PREFIX}		{return token::IT_PREFIX;}
{IT_IMPORT}		{return token::IT_IMPORT;}
{IT_START}		{return token::IT_START;}
{GT_EQUAL}		{return token::GT_EQUAL;}
{IT_EXTERNAL}		{return token::IT_EXTERNAL;}
{IT_OR}			{return token::IT_OR;}
{IT_AND}		{return token::IT_AND;}
{IT_NOT}		{return token::IT_NOT;}
{GT_LPAREN}		{return token::GT_LPAREN;}
{GT_RPAREN}		{return token::GT_RPAREN;}
{GT_DOT}		{return token::GT_DOT;}
{GT_AT}			{return token::GT_AT;}
{IT_LITERAL}		{return token::IT_LITERAL;}
{IT_IRI}		{return token::IT_IRI;}
{IT_BNODE}		{return token::IT_BNODE;}
{IT_NONLITERAL}		{return token::IT_NONLITERAL;}
{IT_LENGTH}		{return token::IT_LENGTH;}
{IT_MINLENGTH}		{return token::IT_MINLENGTH;}
{IT_MAXLENGTH}		{return token::IT_MAXLENGTH;}
{IT_MININCLUSIVE}	{return token::IT_MININCLUSIVE;}
{IT_MINEXCLUSIVE}	{return token::IT_MINEXCLUSIVE;}
{IT_MAXINCLUSIVE}	{return token::IT_MAXINCLUSIVE;}
{IT_MAXEXCLUSIVE}	{return token::IT_MAXEXCLUSIVE;}
{IT_TOTALDIGITS}	{return token::IT_TOTALDIGITS;}
{IT_FRACTIONDIGITS}	{return token::IT_FRACTIONDIGITS;}
{GT_LCURLEY}		{return token::GT_LCURLEY;}
{GT_RCURLEY}		{return token::GT_RCURLEY;}
{IT_CLOSED}		{return token::IT_CLOSED;}
{IT_EXTRA}		{return token::IT_EXTRA;}
{GT_PIPE}		{return token::GT_PIPE;}
{GT_SEMI}		{return token::GT_SEMI;}
{GT_DOLLAR}		{return token::GT_DOLLAR;}
{GT_TIMES}		{return token::GT_TIMES;}
{GT_PLUS}		{return token::GT_PLUS;}
{GT_OPT}		{return token::GT_OPT;}
{GT_CARROT}		{return token::GT_CARROT;}
{GT_LBRACKET}		{return token::GT_LBRACKET;}
{GT_RBRACKET}		{return token::GT_RBRACKET;}
{GT_KINDA}		{return token::GT_KINDA;}
{GT_MINUS}		{return token::GT_MINUS;}
{GT_AMP}		{return token::GT_AMP;}
{GT_DIVIDE_DIVIDE}	{return token::GT_DIVIDE_DIVIDE;}
{GT_MODULO}		{return token::GT_MODULO;}
{GT_DTYPE}		{return token::GT_DTYPE;}
{IT_true}		{return token::IT_true;}
{IT_false}		{return token::IT_false;}
{CODE}		{yylval->str = yytext; return token::CODE;}
{REPEAT_RANGE}		{return token::REPEAT_RANGE;}
{RDF_TYPE}		{return token::RDF_TYPE;}
{IRIREF}		{yylval->p_IRIREF = shex::IRIREF(yytext); return token::IRIREF;}
{PNAME_NS}		{return token::PNAME_NS;}
{PNAME_LN}		{return token::PNAME_LN;}
{ATPNAME_NS}		{return token::ATPNAME_NS;}
{ATPNAME_LN}		{return token::ATPNAME_LN;}
{REGEXP}		{return token::REGEXP;}
{BLANK_NODE_LABEL}	{return token::BLANK_NODE_LABEL;}
{LANGTAG}		{return token::LANGTAG;}
{INTEGER}		{return token::INTEGER;}
{DECIMAL}		{return token::DECIMAL;}
{DOUBLE}		{return token::DOUBLE;}
{STRING_LITERAL1}	{return token::STRING_LITERAL1;}
{STRING_LITERAL2}	{return token::STRING_LITERAL2;}
{STRING_LITERAL_LONG1}	{return token::STRING_LITERAL_LONG1;}
{STRING_LITERAL_LONG2}	{return token::STRING_LITERAL_LONG2;}
{LANG_STRING_LITERAL1}		{return token::LANG_STRING_LITERAL1;}
{LANG_STRING_LITERAL2}		{return token::LANG_STRING_LITERAL2;}
{LANG_STRING_LITERAL_LONG1}	{return token::LANG_STRING_LITERAL_LONG1;}
{LANG_STRING_LITERAL_LONG2}	{return token::LANG_STRING_LITERAL_LONG2;}

<<EOF>>			{ yyterminate();}
.			{ throw shexParseToXml::shexParseToXmlParser::syntax_error(*yylloc, "invalid character: " + std::string(yytext)); }
%%
//.			{ printf( "Unrecognized character: %s\n", yytext ); yyterminate(); }
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

