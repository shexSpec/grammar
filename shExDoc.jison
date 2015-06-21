/*
  jison Equivalent of accompanying bnf, developed in
  http://www.w3.org/2005/01/yacker/uploads/ShEx2
*/

/* lexical grammar */
%lex

IT_BASE		[Bb][Aa][Ss][Ee]
IT_PREFIX		[Pp][Rr][Ee][Ff][Ii][Xx]
IT_start		"start"
GT_EQUAL		"="
IT_VIRTUAL		[Vv][Ii][Rr][Tt][Uu][Aa][Ll]
GT_LCURLEY		"{"
GT_RCURLEY		"}"
IT_CLOSED		[Cc][Ll][Oo][Ss][Ee][Dd]
GT_AMP		"&"
IT_EXTRA		[Ee][Xx][Tt][Rr][Aa]
GT_PIPE		"|"
GT_OR		"||"
GT_COMMA		","
GT_LPAREN		"("
GT_RPAREN		")"
GT_DOLLAR		"$"
GT_NOT		"!"
GT_CARROT		"^"
IT_LITERAL		[Ll][Ii][Tt][Ee][Rr][Aa][Ll]
IT_BNODE		[Bb][Nn][Oo][Dd][Ee]
GT_DOT		"."
IT_IRI		[Ii][Rr][Ii]
IT_NONLITERAL		[Nn][Oo][Nn][Ll][Ii][Tt][Ee][Rr][Aa][Ll]
IT_PATTERN		[Pp][Aa][Tt][Tt][Ee][Rr][Nn]
IT_AND		[Aa][Nn][Dd]
IT_OR		[Oo][Rr]
GT_AT		"@"
GT_KINDA		"~"
IT_MININCLUSIVE		[Mm][Ii][Nn][Ii][Nn][Cc][Ll][Uu][Ss][Ii][Vv][Ee]
IT_MINEXCLUSIVE		[Mm][Ii][Nn][Ee][Xx][Cc][Ll][Uu][Ss][Ii][Vv][Ee]
IT_MAXINCLUSIVE		[Mm][Aa][Xx][Ii][Nn][Cc][Ll][Uu][Ss][Ii][Vv][Ee]
IT_MAXEXCLUSIVE		[Mm][Aa][Xx][Ee][Xx][Cc][Ll][Uu][Ss][Ii][Vv][Ee]
IT_LENGTH		[Ll][Ee][Nn][Gg][Tt][Hh]
IT_MINLENGTH		[Mm][Ii][Nn][Ll][Ee][Nn][Gg][Tt][Hh]
IT_MAXLENGTH		[Mm][Aa][Xx][Ll][Ee][Nn][Gg][Tt][Hh]
IT_TOTALDIGITS		[Tt][Oo][Tt][Aa][Ll][Dd][Ii][Gg][Ii][Tt][Ss]
IT_FRACTIONDIGITS		[Ff][Rr][Aa][Cc][Tt][Ii][Oo][Nn][Dd][Ii][Gg][Ii][Tt][Ss]
GT_SEMI		";"
GT_TIMES		"*"
GT_PLUS		"+"
GT_OPT		"?"
GT_MINUS		"-"
GT_DTYPE		"^^"
IT_true		"true"
IT_false		"false"
CODE		"%" ([#+A-Z_a-z][#+0-9A-Z_a-z]*)? "{" ([^%] | '\\' '%')* "%" "}"
RDF_TYPE		"a"
LANGTAG		"@"([A-Za-z])+(("-"([0-9A-Za-z])+))*
INTEGER		([+-])?([0-9])+
REPEAT_RANGE		"{"({INTEGER})((","(({INTEGER}))?))?"}"
DECIMAL		([+-])?([0-9])*"."([0-9])+
EXPONENT		[Ee]([+-])?([0-9])+
DOUBLE		([+-])?((([0-9])+"."([0-9])*({EXPONENT}))|((".")?([0-9])+({EXPONENT})))
ECHAR		"\\"[\"\\bfnrt]
WS		(" ")|(("\t")|(("\r")|("\n")))
ANON		"\["(({WS}))*"\]"
PN_CHARS_BASE           [A-Z] | [a-z] | [\u00c0-\u00d6] | [\u00d8-\u00f6] | [\u00f8-\u02ff] | [\u0370-\u037d] | [\u037f-\u1fff] | [\u200c-\u200d] | [\u2070-\u218f] | [\u2c00-\u2fef] | [\u3001-\ud7ff] | [\uf900-\ufdcf] | [\ufdf0-\ufffd] | [\U00010000-\U000effff]
PN_CHARS_U              {PN_CHARS_BASE} | '_' | '_' /* !!! raise jison bug */
PN_CHARS                {PN_CHARS_U} | '-' | [0-9] | [\u00b7] | [\u0300-\u036f] | [\u203f-\u2040]
BLANK_NODE_LABEL        '_:' ({PN_CHARS_U} | [0-9]) (({PN_CHARS} | '.')* {PN_CHARS})?
PN_PREFIX               {PN_CHARS_BASE} (({PN_CHARS} | '.')* {PN_CHARS})?
PNAME_NS                {PN_PREFIX}? ':'
HEX                     [0-9] | [A-F] | [a-f]
PERCENT                 '%' {HEX} {HEX}
UCHAR                   '\\u' {HEX} {HEX} {HEX} {HEX} | '\\U' {HEX} {HEX} {HEX} {HEX} {HEX} {HEX} {HEX} {HEX}
STRING_LITERAL1         "'" ([^\u0027\u005c\u000a\u000d] | {ECHAR} | {UCHAR})* "'" /* #x27=' #x5C=\ #xA=new line #xD=carriage return */
STRING_LITERAL2         '"' ([^\u0022\u005c\u000a\u000d] | {ECHAR} | {UCHAR})* '"' /* #x22=" #x5C=\ #xA=new line #xD=carriage return */
STRING_LITERAL_LONG1    "'''" (("'" | "''")? ([^\'\\] | {ECHAR} | {UCHAR}))* "'''"
STRING_LITERAL_LONG2    '"""' (('"' | '""')? ([^\"\\] | {ECHAR} | {UCHAR}))* '"""'
IRIREF			'<' ([^\u0000-\u0020<>\"{}|^`\\] | {UCHAR})* '>' /* #x00=NULL #01-#x1F=control codes #x20=space */
PN_LOCAL_ESC            '\\' ('_' | '~' | '.' | '-' | '!' | '$' | '&' | "'" | '(' | ')' | '*' | '+' | ',' | ';' | '=' | '/' | '?' | '#' | '@' | '%')
PLX                     {PERCENT} | {PN_LOCAL_ESC}
PN_LOCAL                ({PN_CHARS_U} | ':' | [0-9] | {PLX}) (({PN_CHARS} | '.' | ':' | {PLX})* ({PN_CHARS} | ':' | {PLX}))?
PNAME_LN                {PNAME_NS} {PN_LOCAL}
COMMENT			('//'|'#') [^\u000a\u000d]*

%%

\s+|{COMMENT} /**/
{PNAME_LN}		return 'PNAME_LN';
{IT_BASE}		return 'IT_BASE';
{IT_PREFIX}		return 'IT_PREFIX';
{IT_start}		return 'IT_start';
{GT_EQUAL}		return 'GT_EQUAL';
{IT_VIRTUAL}		return 'IT_VIRTUAL';
{GT_LCURLEY}		return 'GT_LCURLEY';
{GT_RCURLEY}		return 'GT_RCURLEY';
{IT_CLOSED}		return 'IT_CLOSED';
{GT_AMP}		return 'GT_AMP';
{IT_EXTRA}		return 'IT_EXTRA';
{GT_PIPE}		return 'GT_PIPE';
{GT_OR}		return 'GT_OR';
{GT_COMMA}		return 'GT_COMMA';
{GT_LPAREN}		return 'GT_LPAREN';
{GT_RPAREN}		return 'GT_RPAREN';
{GT_DOLLAR}		return 'GT_DOLLAR';
{GT_NOT}		return 'GT_NOT';
{GT_CARROT}		return 'GT_CARROT';
{IT_LITERAL}		return 'IT_LITERAL';
{IT_BNODE}		return 'IT_BNODE';
{GT_DOT}		return 'GT_DOT';
{IT_IRI}		return 'IT_IRI';
{IT_NONLITERAL}		return 'IT_NONLITERAL';
{IT_PATTERN}		return 'IT_PATTERN';
{IT_AND}		return 'IT_AND';
{IT_OR}		return 'IT_OR';
{GT_AT}		return 'GT_AT';
{GT_KINDA}		return 'GT_KINDA';
{IT_MININCLUSIVE}		return 'IT_MININCLUSIVE';
{IT_MINEXCLUSIVE}		return 'IT_MINEXCLUSIVE';
{IT_MAXINCLUSIVE}		return 'IT_MAXINCLUSIVE';
{IT_MAXEXCLUSIVE}		return 'IT_MAXEXCLUSIVE';
{IT_LENGTH}		return 'IT_LENGTH';
{IT_MINLENGTH}		return 'IT_MINLENGTH';
{IT_MAXLENGTH}		return 'IT_MAXLENGTH';
{IT_TOTALDIGITS}		return 'IT_TOTALDIGITS';
{IT_FRACTIONDIGITS}		return 'IT_FRACTIONDIGITS';
{GT_SEMI}		return 'GT_SEMI';
{GT_TIMES}		return 'GT_TIMES';
{GT_PLUS}		return 'GT_PLUS';
{GT_OPT}		return 'GT_OPT';
{GT_MINUS}		return 'GT_MINUS';
{GT_DTYPE}		return 'GT_DTYPE';
{IT_true}		return 'IT_true';
{IT_false}		return 'IT_false';
{CODE}		return 'CODE';
{RDF_TYPE}		return 'RDF_TYPE';
//{LANGTAG}		return 'LANGTAG';
{INTEGER}		return 'INTEGER';
{REPEAT_RANGE}		return 'REPEAT_RANGE';
{DECIMAL}		return 'DECIMAL';
{EXPONENT}		return 'EXPONENT';
{DOUBLE}		return 'DOUBLE';
//{ECHAR}		return 'ECHAR';
//{WS}		return 'WS';
{ANON}		return 'ANON';
{IRIREF}		return 'IRIREF';
{PNAME_NS}		return 'PNAME_NS';
//{PN_CHARS_BASE}		return 'PN_CHARS_BASE';
//{PN_CHARS_U}		return 'PN_CHARS_U';
//{PN_CHARS}		return 'PN_CHARS';
{BLANK_NODE_LABEL}		return 'BLANK_NODE_LABEL';
//{PN_PREFIX}		return 'PN_PREFIX';
//{HEX}		return 'HEX';
//{PERCENT}		return 'PERCENT';
//{UCHAR}		return 'UCHAR';
{STRING_LITERAL1}		return 'STRING_LITERAL1';
{STRING_LITERAL2}		return 'STRING_LITERAL2';
{STRING_LITERAL_LONG1}		return 'STRING_LITERAL_LONG1';
{STRING_LITERAL_LONG2}		return 'STRING_LITERAL_LONG2';
//{PN_LOCAL_ESC}		return 'PN_LOCAL_ESC';
//{PLX}		return 'PLX';
//{PN_LOCAL}		return 'PN_LOCAL';
<<EOF>>               return 'EOF'
.                     return 'invalid character'

/lex

/* operator associations and precedence */

%start shexDoc

%% /* language grammar */

shexDoc:
    _Qdirective_E_Star _Q_O_Qshape_E_Or_Qstart_E_Or_QCODE_E_Plus_S_Qstatement_E_Star_C_E_Opt EOF	;

_Qdirective_E_Star:
    
    | _Qdirective_E_Star directive	;

_QCODE_E_Plus:
    CODE	
    | _QCODE_E_Plus CODE	;

_O_Qshape_E_Or_Qstart_E_Or_QCODE_E_Plus_C:
    shape	
    | start	
    | _QCODE_E_Plus	;

_Qstatement_E_Star:
    
    | _Qstatement_E_Star statement	;

_O_Qshape_E_Or_Qstart_E_Or_QCODE_E_Plus_S_Qstatement_E_Star_C:
    _O_Qshape_E_Or_Qstart_E_Or_QCODE_E_Plus_C _Qstatement_E_Star	;

_Q_O_Qshape_E_Or_Qstart_E_Or_QCODE_E_Plus_S_Qstatement_E_Star_C_E_Opt:
    
    | _O_Qshape_E_Or_Qstart_E_Or_QCODE_E_Plus_S_Qstatement_E_Star_C	;



statement:
    directive	
    | start	
    | shape	;

directive:
    baseDecl	
    | prefixDecl	;

baseDecl:
    IT_BASE IRIREF	;

prefixDecl:
    IT_PREFIX PNAME_NS IRIREF	;

start:
    IT_start GT_EQUAL _O_QshapeLabel_E_Or_QshapeDefinition_E_S_QCODE_E_Star_C	;

_QCODE_E_Star:
    
    | _QCODE_E_Star CODE	;

_O_QshapeLabel_E_Or_QshapeDefinition_E_S_QCODE_E_Star_C:
    shapeLabel	
    | shapeDefinition _QCODE_E_Star	;

shape:
    // _QIT_VIRTUAL_E_Opt 
    shapeLabel shapeDefinition _QCODE_E_Star	
    | IT_VIRTUAL shapeLabel shapeDefinition _QCODE_E_Star	;

// _QIT_VIRTUAL_E_Opt:
//     
//     | IT_VIRTUAL	;

shapeDefinition:
    _Q_O_Qinclude_E_Or_QinclPropertySet_E_Or_QIT_CLOSED_E_C_E_Star GT_LCURLEY _QoneOfShape_E_Opt GT_RCURLEY	;

_O_Qinclude_E_Or_QinclPropertySet_E_Or_QIT_CLOSED_E_C:
    include	
    | inclPropertySet	
    | IT_CLOSED	;

_Q_O_Qinclude_E_Or_QinclPropertySet_E_Or_QIT_CLOSED_E_C_E_Star:
    
    | _Q_O_Qinclude_E_Or_QinclPropertySet_E_Or_QIT_CLOSED_E_C_E_Star _O_Qinclude_E_Or_QinclPropertySet_E_Or_QIT_CLOSED_E_C	;

_QoneOfShape_E_Opt:
    
    | oneOfShape	;

include:
    GT_AMP shapeLabel	;

inclPropertySet:
    IT_EXTRA _Qpredicate_E_Plus	;

_Qpredicate_E_Plus:
    predicate	
    | _Qpredicate_E_Plus predicate	;

oneOfShape:
    someOfShape _Q_O_QGT_PIPE_E_S_QsomeOfShape_E_C_E_Star	;

_O_QGT_PIPE_E_S_QsomeOfShape_E_C:
    GT_PIPE someOfShape	;

_Q_O_QGT_PIPE_E_S_QsomeOfShape_E_C_E_Star:
    
    | _Q_O_QGT_PIPE_E_S_QsomeOfShape_E_C_E_Star _O_QGT_PIPE_E_S_QsomeOfShape_E_C	;

someOfShape:
    groupShape _Q_O_QGT_OR_E_S_QgroupShape_E_C_E_Star	;

_O_QGT_OR_E_S_QgroupShape_E_C:
    GT_OR groupShape	;

_Q_O_QGT_OR_E_S_QgroupShape_E_C_E_Star:
    
    | _Q_O_QGT_OR_E_S_QgroupShape_E_C_E_Star _O_QGT_OR_E_S_QgroupShape_E_C	;

groupShape:
    unaryShape _Q_O_QGT_COMMA_E_S_QunaryShape_E_C_E_Star _QGT_COMMA_E_Opt	;

_O_QGT_COMMA_E_S_QunaryShape_E_C:
    GT_COMMA unaryShape	;

_Q_O_QGT_COMMA_E_S_QunaryShape_E_C_E_Star:
    
    | _Q_O_QGT_COMMA_E_S_QunaryShape_E_C_E_Star _O_QGT_COMMA_E_S_QunaryShape_E_C	;

_QGT_COMMA_E_Opt:
    
    | GT_COMMA	;

unaryShape:
    // _Qid_E_Opt 
    _O_QtripleConstraint_E_Or_Qinclude_E_Or_QGT_LPAREN_E_S_QoneOfShape_E_S_QGT_RPAREN_E_S_Qcardinality_E_Opt_S_QCODE_E_Star_C	
    | id _O_QtripleConstraint_E_Or_Qinclude_E_Or_QGT_LPAREN_E_S_QoneOfShape_E_S_QGT_RPAREN_E_S_Qcardinality_E_Opt_S_QCODE_E_Star_C	;

// _Qid_E_Opt:
//     
//     | id	;

_Qcardinality_E_Opt:
    
    | cardinality	;

_O_QtripleConstraint_E_Or_Qinclude_E_Or_QGT_LPAREN_E_S_QoneOfShape_E_S_QGT_RPAREN_E_S_Qcardinality_E_Opt_S_QCODE_E_Star_C:
    tripleConstraint	
    | include	
    | GT_LPAREN oneOfShape GT_RPAREN _Qcardinality_E_Opt _QCODE_E_Star	;

id:
    GT_DOLLAR shapeLabel	;

shapeLabel:
    iri	
    | blankNode	;

tripleConstraint:
    // _QsenseFlags_E_Opt 
    predicate valueClass _Qannotation_E_Star _Qcardinality_E_Opt _QCODE_E_Star	
    | senseFlags predicate valueClass _Qannotation_E_Star _Qcardinality_E_Opt _QCODE_E_Star	;

// _QsenseFlags_E_Opt:
//     
//     | senseFlags	;

_Qannotation_E_Star:
    
    | _Qannotation_E_Star annotation	;

senseFlags:
    GT_NOT _QGT_CARROT_E_Opt	
    | GT_CARROT _QGT_NOT_E_Opt	;

_QGT_CARROT_E_Opt:
    
    | GT_CARROT	;

_QGT_NOT_E_Opt:
    
    | GT_NOT	;

predicate:
    RDF_TYPE	
    | iri	;

valueClass:
    IT_LITERAL _QxsFacet_E_Star	
//    | _O_QIT_IRI_E_Or_QIT_NONLITERAL_E_C _QgroupShapeConstr_E_Opt _Q_O_QIT_PATTERN_E_S_Qstring_E_C_E_Opt	
    | _O_QIT_IRI_E_Or_QIT_NONLITERAL_E_C	
    | _O_QIT_IRI_E_Or_QIT_NONLITERAL_E_C _O_QIT_PATTERN_E_S_Qstring_E_C	
    | _O_QIT_IRI_E_Or_QIT_NONLITERAL_E_C groupShapeConstr	
    | _O_QIT_IRI_E_Or_QIT_NONLITERAL_E_C groupShapeConstr _O_QIT_PATTERN_E_S_Qstring_E_C	
//    | IT_BNODE _QgroupShapeConstr_E_Opt	
    | IT_BNODE	
    | IT_BNODE groupShapeConstr	
    | datatype	
    | groupShapeConstr	
    | valueSet	
    | GT_DOT	;

_QxsFacet_E_Star:
    
    | _QxsFacet_E_Star xsFacet	;

_O_QIT_IRI_E_Or_QIT_NONLITERAL_E_C:
    IT_IRI	
    | IT_NONLITERAL	;

// _QgroupShapeConstr_E_Opt:
//     
//     | groupShapeConstr	;

_O_QIT_PATTERN_E_S_Qstring_E_C:
    IT_PATTERN string	;

_Q_O_QIT_PATTERN_E_S_Qstring_E_C_E_Opt:
    
    | _O_QIT_PATTERN_E_S_Qstring_E_C	;

groupShapeConstr:
    shapeOrRef _Q_O_QIT_AND_E_Or_QIT_OR_E_S_QshapeOrRef_E_C_E_Star	;

_O_QIT_AND_E_Or_QIT_OR_E_C:
    IT_AND	
    | IT_OR	;

_O_QIT_AND_E_Or_QIT_OR_E_S_QshapeOrRef_E_C:
    _O_QIT_AND_E_Or_QIT_OR_E_C shapeOrRef	;

_Q_O_QIT_AND_E_Or_QIT_OR_E_S_QshapeOrRef_E_C_E_Star:
    
    | _Q_O_QIT_AND_E_Or_QIT_OR_E_S_QshapeOrRef_E_C_E_Star _O_QIT_AND_E_Or_QIT_OR_E_S_QshapeOrRef_E_C	;

shapeOrRef:
    GT_AT shapeLabel	
    | shapeDefinition	;

xsFacet:
    _O_QIT_PATTERN_E_Or_QGT_KINDA_E_C string	
    | _O_QIT_MININCLUSIVE_E_Or_QIT_MINEXCLUSIVE_E_Or_QIT_MAXINCLUSIVE_E_Or_QIT_MAXEXCLUSIVE_E_C INTEGER	
    | _O_QIT_LENGTH_E_Or_QIT_MINLENGTH_E_Or_QIT_MAXLENGTH_E_C INTEGER	
    | _O_QIT_TOTALDIGITS_E_Or_QIT_FRACTIONDIGITS_E_C INTEGER	;

_O_QIT_PATTERN_E_Or_QGT_KINDA_E_C:
    IT_PATTERN	
    | GT_KINDA	;

_O_QIT_MININCLUSIVE_E_Or_QIT_MINEXCLUSIVE_E_Or_QIT_MAXINCLUSIVE_E_Or_QIT_MAXEXCLUSIVE_E_C:
    IT_MININCLUSIVE	
    | IT_MINEXCLUSIVE	
    | IT_MAXINCLUSIVE	
    | IT_MAXEXCLUSIVE	;

_O_QIT_LENGTH_E_Or_QIT_MINLENGTH_E_Or_QIT_MAXLENGTH_E_C:
    IT_LENGTH	
    | IT_MINLENGTH	
    | IT_MAXLENGTH	;

_O_QIT_TOTALDIGITS_E_Or_QIT_FRACTIONDIGITS_E_C:
    IT_TOTALDIGITS	
    | IT_FRACTIONDIGITS	;

datatype:
    iri	;

annotation:
    GT_SEMI iri _O_Qiri_E_Or_Qliteral_E_C	;

_O_Qiri_E_Or_Qliteral_E_C:
    iri	
    | literal	;

cardinality:
    GT_TIMES	
    | GT_PLUS	
    | GT_OPT	
    | REPEAT_RANGE	;

valueSet:
    GT_LPAREN _Qvalue_E_Star GT_RPAREN	;

_Qvalue_E_Star:
    
    | _Qvalue_E_Star value	;

value:
    iriRange	
    | literal	;

iriRange:
    iri _Q_O_QGT_KINDA_E_S_Qexclusion_E_Star_C_E_Opt	
    | GT_DOT _Qexclusion_E_Plus	;

_Qexclusion_E_Star:
    
    | _Qexclusion_E_Star exclusion	;

_O_QGT_KINDA_E_S_Qexclusion_E_Star_C:
    GT_KINDA _Qexclusion_E_Star	;

_Q_O_QGT_KINDA_E_S_Qexclusion_E_Star_C_E_Opt:
    
    | _O_QGT_KINDA_E_S_Qexclusion_E_Star_C	;

_Qexclusion_E_Plus:
    exclusion	
    | _Qexclusion_E_Plus exclusion	;

exclusion:
    GT_MINUS iri _QGT_KINDA_E_Opt	;

_QGT_KINDA_E_Opt:
    
    | GT_KINDA	;

literal:
    rdfLiteral	
    | numericLiteral	
    | booleanLiteral	;

numericLiteral:
    INTEGER	
    | DECIMAL	
    | DOUBLE	;

rdfLiteral:
    string _Q_O_QLANGTAG_E_Or_QGT_DTYPE_E_S_Qiri_E_C_E_Opt	;

_O_QLANGTAG_E_Or_QGT_DTYPE_E_S_Qiri_E_C:
    LANGTAG	
    | GT_DTYPE iri	;

_Q_O_QLANGTAG_E_Or_QGT_DTYPE_E_S_Qiri_E_C_E_Opt:
    
    | _O_QLANGTAG_E_Or_QGT_DTYPE_E_S_Qiri_E_C	;

booleanLiteral:
    IT_true	
    | IT_false	;

string:
    STRING_LITERAL1	
    | STRING_LITERAL2	
    | STRING_LITERAL_LONG1	
    | STRING_LITERAL_LONG2	;

iri:
    IRIREF	
    | prefixedName	;

prefixedName:
    PNAME_LN	
    | PNAME_NS	;

blankNode:
    BLANK_NODE_LABEL	
    | ANON	;

