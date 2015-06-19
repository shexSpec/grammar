// ANTLR4 Equivalent of accompanying bnf, developed in
// http://www.w3.org/2005/01/yacker/uploads/ShEx2

grammar shExDoc;


shExDoc 		: directive* ((shape | start | code+) statement*)? ;  // leading CODE
statement		: directive | start | shape ;
directive       : baseDecl | prefixDecl ;
baseDecl 		: KW_BASE IRIREF ;
prefixDecl		: KW_PREFIX PNAME_NS IRIREF ;
start           : KW_START '=' ( shapeLabel | shapeDefinition code* ) ;
shape           : KW_VIRTUAL? shapeLabel shapeDefinition code* ;
shapeDefinition : (include | extra | KW_CLOSED)* '{' oneOfShape? '}' ;
include         : '&' shapeLabel ;
extra           : (KW_EXTRA | '+') predicate ;
oneOfShape      : someOfShape ( '|' someOfShape )* ;
someOfShape     : groupShape ( '||' groupShape )* ;
groupShape      : unaryShape ( ',' unaryShape )* ','? ;
unaryShape      : id? ( tripleConstraint | include | '(' oneOfShape ')' cardinality? code* ) ;
id              : '$' shapeLabel ;
shapeLabel      : iri | blankNode ;
tripleConstraint : senseFlags? predicate valueClass annotation* cardinality? code* ;
senseFlags      : '~' '^'? | '^' '!'? ;
predicate       : RDF_TYPE | iri ;
valueClass      : KW_LITERAL xsFacet*
                | (KW_IRI | KW_NONLITERAL) groupShapeConstr? (KW_PATTERN string)?
                | KW_BNODE groupShapeConstr?
                | datatype
                | groupShapeConstr
                | valueSet
                | '.'
                ;
groupShapeConstr : shapeOrRef ((KW_AND | KW_OR) shapeOrRef)* ;
shapeOrRef      : '@' shapeLabel | shapeDefinition ;
xsFacet         : KW_PATTERN string
                | (KW_MININCLUSIVE | KW_MINEXCLUSIVE | KW_MAXINCLUSIVE | KW_MAXEXCLUSIVE) INTEGER
                | (KW_LENGTH | KW_MINLENGTH | KW_MAXLENGTH) INTEGER
                | (KW_TOTALDIGITS | KW_FRACTIONDIGITS) INTEGER
                ;
datatype        : iri ;
annotation      : ';' iri (iri | literal) ;
cardinality     :  '*' | '+' | '?' | REPEAT_RANGE ;
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
booleanLiteral  : KW_TRUE | KW_FALSE ;
string          : STRING ;
iri             : IRIREF
				| prefixedName
				;
prefixedName    : PNAME_LN
				| PNAME_NS
				;
blankNode       : BLANK_NODE_LABEL
				| ANON
				;
code	        : CODE ;

// Keywords
KW_BASE 			: B A S E ;
KW_PREFIX       	: P R E F I X ;
KW_START        	: 'start' ;
KW_VIRTUAL      	: V I R T U A L ;
KW_CLOSED       	: C L O S E D ;
KW_EXTRA        	: E X T R A ;
KW_LITERAL      	: L I T E R A L ;
KW_IRI          	: I R I ;
KW_NONLITERAL   	: N O N L I T E R A L ;
KW_PATTERN      	: P A T T E R N ;
KW_BNODE        	: B N O D E ;
KW_AND          	: A N D ;
KW_OR           	: O R ;
KW_MININCLUSIVE 	: M I N I N C L U S I V E ;
KW_MINEXCLUSIVE 	: M I N E X C L U S I V E ;
KW_MAXINCLUSIVE 	: M A X I N C L U S I V E ;
KW_MAXEXCLUSIVE 	: M A X E X C L U S I V E ;
KW_LENGTH       	: L E N G T H ;
KW_MINLENGTH    	: M I N L E N G T H ;
KW_MAXLENGTH    	: M A X L E N G T H ;
KW_TOTALDIGITS  	: T O T A L D I G I T S ;
KW_FRACTIONDIGITS 	: F R A C T I O N D I G I T S ;
KW_TRUE         	: 'true' ;
KW_FALSE        	: 'false' ;

// terminals
PASS				  : [ \t\r\n]+ -> skip;
COMMENT				  : '#' ~[\r\n]* -> skip;

CODE                  : '%' ([a-zA-Z+#_][a-zA-Z0-9+#_]*)? '{' (~[%\\] | '\\%')* '%' '}' ;
REPEAT_RANGE          : '{' INTEGER (',' INTEGER?)? '}' ;
RDF_TYPE              : WS+ 'a' WS+ ;
IRIREF                : '<' (~[\u0000-\u0020=<>\"{}|^`\\] | UCHAR)* '>' ; /* #x00=NULL #01-#x1F=control codes #x20=space */
PNAME_NS              : PN_PREFIX? ':' ;
PNAME_LN              : PNAME_NS PN_LOCAL ;
BLANK_NODE_LABEL      : '_:' (PN_CHARS_U | [0-9]) ((PN_CHARS | '.')* PN_CHARS)? ;
LANGTAG               : '@' [a-zA-Z]+ ('-' [a-zA-Z0-9]+)* ;
INTEGER               : [+-]? [0-9]+ ;
DECIMAL               : [+-]? [0-9]* '.' [0-9]+ ;
DOUBLE                : [+-]? ([0-9]+ '.' [0-9]* EXPONENT | '.'? [0-9]+ EXPONENT) ;
fragment EXPONENT              : [eE] [+-]? [0-9]+ ;
STRING                : STRING_LITERAL1
                      | STRING_LITERAL2
                      | STRING_LITERAL_LONG1
                      | STRING_LITERAL_LONG2
                      ;
fragment STRING_LITERAL1       : '\'' (~[\u0027\u005C\u000A\u000D] | ECHAR | UCHAR)* '\'' ; /* #x27=' #x5C=\ #xA=new line #xD=carriage return */
fragment STRING_LITERAL2       : '"' (~[\u0022\u005C\u000A\u000D] | ECHAR | UCHAR)* '"' ;   /* #x22=" #x5C=\ #xA=new line #xD=carriage return */
fragment STRING_LITERAL_LONG1  : '\'\'\'' (('\'' | '\'\'')? (~[\'\\] | ECHAR | UCHAR))* '\'\'\'' ;
fragment STRING_LITERAL_LONG2  : '"""' (('"' | '""')? (~[\"\\] | ECHAR | UCHAR))* '"""' ;
fragment UCHAR                 : '\\u' HEX HEX HEX HEX | '\\U' HEX HEX HEX HEX HEX HEX HEX HEX ;
fragment ECHAR                 : '\\' [tbnrf\\\"\'] ;
fragment WS                    : [\u0020\u0009\u000D\u000A] ; /* #x20=space #x9=character tabulation #xD=carriage return #xA=new line */
ANON                  		   : '[' WS* ']' ;
fragment PN_CHARS_BASE 		   : [A-Z] | [a-z] | [\u00C0-\u00D6] | [\u00D8-\u00F6] | [\u00F8-\u02FF] | [\u0370-\u037D]
					   		   | [\u037F-\u1FFF] | [\u200C-\u200D] | [\u2070-\u218F] | [\u2C00-\u2FEF] | [\u3001-\uD7FF]
					           | [\uF900-\uFDCF] | [\uFDF0-\uFFFD]
					   		   ;
fragment PN_CHARS_U            : PN_CHARS_BASE | '_' ;
fragment PN_CHARS              : PN_CHARS_U | '-' | [0-9] | [\u00B7] | [\u0300-\u036F] | [\u203F-\u2040] ;
fragment PN_PREFIX             : PN_CHARS_BASE ((PN_CHARS | '.')* PN_CHARS)? ;
fragment PN_LOCAL              : (PN_CHARS_U | ':' | [0-9] | PLX) ((PN_CHARS | '.' | ':' | PLX)* (PN_CHARS | ':' | PLX))? ;
fragment PLX                   : PERCENT | PN_LOCAL_ESC ;
fragment PERCENT               : '%' HEX HEX ;
fragment HEX                   : [0-9] | [A-F] | [a-f] ;
fragment PN_LOCAL_ESC          : '\\' ('_' | '~' | '.' | '-' | '!' | '$' | '&' | '\'' | '(' | ')' | '*' | '+' | ','
					  		   | ';' | '=' | '/' | '?' | '#' | '@' | '%') ;

fragment A:('a'|'A');
fragment B:('b'|'B');
fragment C:('c'|'C');
fragment D:('d'|'D');
fragment E:('e'|'E');
fragment F:('f'|'F');
fragment G:('g'|'G');
fragment H:('h'|'H');
fragment I:('i'|'I');
fragment J:('j'|'J');
fragment K:('k'|'K');
fragment L:('l'|'L');
fragment M:('m'|'M');
fragment N:('n'|'N');
fragment O:('o'|'O');
fragment P:('p'|'P');
fragment Q:('q'|'Q');
fragment R:('r'|'R');
fragment S:('s'|'S');
fragment T:('t'|'T');
fragment U:('u'|'U');
fragment V:('v'|'V');
fragment W:('w'|'W');
fragment X:('x'|'X');
fragment Y:('y'|'Y');
fragment Z:('z'|'Z');