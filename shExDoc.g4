grammar shExDoc;


shExDoc 		: directive* ((shape | start | code+) statement*)? ;
statement		: directive | start | shape ;
directive       : baseDecl | prefixDecl ;
baseDecl 		: 'BASE' IRIREF ;
prefixDecl		: 'PREFIX' PNAME_NS IRIREF ;
start           : 'start' '=' ( label | typeSpec CODE* ) ;
shape           : 'VIRTUAL'? label typeSpec CODE* ;
typeSpec        : include* '{' choiceExpr? '}' ;
include         : '&' label;
choiceExpr      : sequenceExpr ( '|' sequenceExpr )* ;
sequenceExpr    : unaryExpr ( ',' unaryExpr )* ;
unaryExpr       : id? ( arc | include | '(' choiceExpr ')' repeatCount? CODE* ) ;
id              : '$' label ;
label           : iri | blankNode ;
arc             : '!'? '^'? (RDF_TYPE | iri) valueClass annotation* repeatCount? CODE* ;
valueClass      : 'LITERAL' xsFacet*
                | ('IRI' | 'NONLITERAL') ('@' label)? ('PATTERN' string)?
                | 'BNODE' ('@' label)?
                | datatype
                | '@' label
                | typeSpec
                | valueSet
                | '.'
                ;
xsFacet         : 'PATTERN' string
                | ('MININCLUSIVE' | 'MINEXCLUSIVE' | 'MAXINCLUSIVE' | 'MAXEXCLUSIVE') INTEGER
                | ('LENGTH' | 'MINLENGTH' | 'MAXLENGTH') INTEGER
                | ('TOTALDIGITS' | 'FRACTIONDIGITS') INTEGER
                ;
datatype        : iri ;
annotation      : ';' iri (iri | literal) ;
repeatCount     :  '*' | '+' | '?' | '{' INTEGER (',' INTEGER?)? '}' ;
valueSet        : '(' value* ')' ;
value           : iriRange
				| literal
				;
iriRange        : iri ('~' exclusion*)? | '.' exclusion+ ;
exclusion       : '-' iri '~'? ;
literal         : rdfLiteral
				| numericLiteral
				| booleanLiteral
				;

numericLiteral  : INTEGER
				| DECIMAL
				| DOUBLE
				;
rdfLiteral      : string (LANGTAG | '^^' iri)? ;
booleanLiteral  : 'true' | 'false' ;
string          : STRING_LITERAL1
				| STRING_LITERAL2
				| STRING_LITERAL_LONG1
				| STRING_LITERAL_LONG2
				;
iri             : IRIREF
				| prefixedName
				;
prefixedName    : PNAME_LN
				| PNAME_NS
				;
blankNode       : BLANK_NODE_LABEL
				| ANON
				;

code			: CODE ;

PASS				  : [ \t\r\n]+ -> skip;
COMMENT				  :  '#' [^\r\n]* -> skip;


CODE                  : '%' ([a-zA-Z+#_][a-zA-Z0-9+#_]*)? '{' (~[%\\] | '\\%')* '%' '}' ;

RDF_TYPE              : 'a' ;
IRIREF                : '<' ([^#x00-#x20<>\"{}|^`\\] | UCHAR)* '>' ; /* #x00=NULL #01-#x1F=control codes #x20=space */
PNAME_NS              : PN_PREFIX? ':' ;
PNAME_LN              : PNAME_NS PN_LOCAL ;
BLANK_NODE_LABEL      : '_:' (PN_CHARS_U | [0-9]) ((PN_CHARS | '.')* PN_CHARS)? ;
LANGTAG               : '@' [a-zA-Z]+ ('-' [a-zA-Z0-9]+)* ;
INTEGER               : [+-]? [0-9]+ ;
DECIMAL               : [+-]? [0-9]* '.' [0-9]+ ;
DOUBLE                : [+-]? ([0-9]+ '.' [0-9]* EXPONENT | '.'? [0-9]+ EXPONENT) ;
EXPONENT              : [eE] [+-]? [0-9]+ ;
STRING_LITERAL1       : '\'' ([^#x27#x5C#xA#xD] | ECHAR | UCHAR)* '\'' ; /* #x27=' #x5C=\ #xA=new line #xD=carriage return */
STRING_LITERAL2       : '"' ([^#x22#x5C#xA#xD] | ECHAR | UCHAR)* '"' ;   /* #x22=" #x5C=\ #xA=new line #xD=carriage return */
STRING_LITERAL_LONG1  : '\'\'\'' (('\'' | '\'\'')? ([^\'\\] | ECHAR | UCHAR))* '\'\'\'' ;
STRING_LITERAL_LONG2  : '"""' (('"' | '""')? ([^\"\\] | ECHAR | UCHAR))* '"""' ;
UCHAR                 : '\\u' HEX HEX HEX HEX | '\\U' HEX HEX HEX HEX HEX HEX HEX HEX ;
ECHAR                 : '\\' [tbnrf\\\"\'] ;
WS                    : [\u0020\u0009\u000D\u000A] ; /* #x20=space #x9=character tabulation #xD=carriage return #xA=new line */
ANON                  : '[' WS* ']' ;
PN_CHARS_BASE         : [A-Z] | [a-z] | [\u00C0-\u00D6] | [\u00D8-\u00F6] | [\u00F8-\u02FF] | [\u0370-\u037D]
					  | [\u037F-\u1FFF] | [\u200C-\u200D] | [\u2070-\u218F] | [\u2C00-\u2FEF] | [\u3001-\uD7FF]
					  | [\uF900-\uFDCF] | [\uFDF0-\uFFFD] | [\u10000-\uEFFFF]
					  ;
PN_CHARS_U            : PN_CHARS_BASE | '_' ;
PN_CHARS              : PN_CHARS_U | '-' | [0-9] | [\u00B7] | [\u0300-\u036F] | [\u203F-\u2040] ;
PN_PREFIX             : PN_CHARS_BASE ((PN_CHARS | '.')* PN_CHARS)? ;
PN_LOCAL              : (PN_CHARS_U | ':' | [0-9] | PLX) ((PN_CHARS | '.' | ':' | PLX)* (PN_CHARS | ':' | PLX))? ;
PLX                   : PERCENT | PN_LOCAL_ESC ;
PERCENT               : '%' HEX HEX ;
HEX                   : [0-9] | [A-F] | [a-f] ;
PN_LOCAL_ESC          : '\\' ('_' | '~' | '.' | '-' | '!' | '$' | '&' | '\'' | '(' | ')' | '*' | '+' | ','
					  | ';' | '=' | '/' | '?' | '#' | '@' | '%') ;



