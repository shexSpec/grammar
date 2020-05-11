
  struct IRIREF {
    std::string i;
    IRIREF (std::string ip) { i = ip; }
    size_t size () { return i.size(); }
    void hi () { std::cout << "IRIREF(" + i + ")" << std::endl; }
  };
  struct Language {
    std::string i;
    Language (std::string ip) { i = ip; }
    size_t size () { return i.size(); }
    void hi () { std::cout << "Language(" + i + ")" << std::endl; }
  };
  /*
Schema	{	startActs:[SemAct+]? start:shapeExpr? imports:[IRIREF]? shapes:[shapeExpr+]? }
shapeExpr	=	ShapeOr | ShapeAnd | ShapeNot | NodeConstraint | Shape | ShapeExternal | shapeExprRef ;
ShapeOr	{	id:shapeExprLabel? shapeExprs:[shapeExpr{2,}] }
ShapeAnd	{	id:shapeExprLabel? shapeExprs:[shapeExpr{2,}] }
ShapeNot	{	id:shapeExprLabel? shapeExpr:shapeExpr }
ShapeExternal	{	id:shapeExprLabel? }
shapeExprRef	=	shapeExprLabel ;
shapeExprLabel	=	IRIREF | BNODE ;
NodeConstraint	{	id:shapeExprLabel? nodeKind:("iri" | "bnode" | "nonliteral" | "literal")? datatype:IRIREF? xsFacet* values:[valueSetValue]? }
xsFacet	=	stringFacet | numericFacet ;
stringFacet	=	(length|minlength|maxlength):INTEGER | pattern:STRING flags:STRING? ;
numericFacet	=	(mininclusive|minexclusive|maxinclusive|maxexclusive):numericLiteral
|	(totaldigits|fractiondigits):INTEGER ;
numericLiteral	=	INTEGER | DECIMAL | DOUBLE ;
valueSetValue	=	objectValue | IriStem | IriStemRange | LiteralStem | LiteralStemRange | Language | LanguageStem | LanguageStemRange ;
  */

struct ObjectLiteral {
  std::string value;
  boost::optional<Language> language;
  boost::optional<IRIREF> type;
};
typedef boost::variant<IRIREF, ObjectLiteral> objectValue;
  /*
IriStem	{	stem:IRIREF }
IriStemRange	{	stem:(IRIREF | Wildcard) exclusions:[IRIREF|IriStem +] }
LiteralStem	{	stem:STRING }
LiteralStemRange	{	stem:(STRING | Wildcard) exclusions:[STRING|LiteralStem +] }
Language	{	languageTag:LANGTAG }
LanguageStem	{	stem:LANGTAG }
LanguageStemRange	{	stem:(LANGTAG | Wildcard) exclusions:[LANGTAG|LanguageStem +] }
Wildcard	{	 }
Shape	{	id:shapeExprLabel? closed:BOOL? extra:[IRIREF]? expression:tripleExpr? semActs:[SemAct+]? annotations:[Annotation+]? }
tripleExpr	=	EachOf | OneOf | TripleConstraint | tripleExprRef ;
EachOf	{	id:tripleExprLabel? expressions:[tripleExpr{2,}] min:INTEGER? max:INTEGER? semActs:[SemAct+]? annotations:[Annotation+]? }
OneOf	{	id:tripleExprLabel? expressions:[tripleExpr{2,}] min:INTEGER? max:INTEGER? semActs:[SemAct+]? annotations:[Annotation+]? }
TripleConstraint	{	id:tripleExprLabel? inverse:BOOL? predicate:IRIREF valueExpr:shapeExpr? min:INTEGER? max:INTEGER? semActs:[SemAct+]? annotations:[Annotation+]? }
tripleExprRef	=	tripleExprLabel ;
tripleExprLabel	=	IRIREF | BNODE ;
SemAct	{	name:IRIREF code:STRING? }
*/
  struct Annotation {
    IRIREF predicate;
    objectValue object;
  };
