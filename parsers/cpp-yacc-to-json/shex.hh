// $Id: Langname_Scanner.hh,v 1.1 2008/04/06 17:10:46 eric Exp $

#ifndef shex_H
#define shex_H

#include <iostream>
#include <boost/optional.hpp>
#include <boost/variant.hpp>
#include <utility>

namespace shex {

template<typename T>
struct ShExTerminal;

  #define TT(T, P) using T = struct T {                           \
    std::string val;                                              \
    T (std::string val) : val(val) {  }                           \
    inline static const char* production = #T;                    \
    inline static const char* pattern = P;                        \
    void hi () { std::cout << #T "(" << val << ")" << std::endl; }\
  };

  // ShEx terminals
  TT(STRING, "^.*$");
  TT(LANGTAG, "^((([a-zA-Z])+((-([a-zA-Z0-9])+))*))$");
  TT(BNODE, "^((_:((([A-Z]|[a-z]|[\\u00C0-\\u00D6]|[\\u00D8-\\u00F6]|[\\u00F8-\\u02FF]|[\\u0370-\\u037D]|[\\u037F-\\u1FFF]|[\\u200C-\\u200D]|[\\u2070-\\u218F]|[\\u2C00-\\u2FEF]|[\\u3001-\\uD7FF]|[\\uF900-\\uFDCF]|[\\uFDF0-\\uFFFD]|[\\u10000-\\uEFFFF])|_)|[0-9])((((((([A-Z]|[a-z]|[\\u00C0-\\u00D6]|[\\u00D8-\\u00F6]|[\\u00F8-\\u02FF]|[\\u0370-\\u037D]|[\\u037F-\\u1FFF]|[\\u200C-\\u200D]|[\\u2070-\\u218F]|[\\u2C00-\\u2FEF]|[\\u3001-\\uD7FF]|[\\uF900-\\uFDCF]|[\\uFDF0-\\uFFFD]|[\\u10000-\\uEFFFF])|_)|-|[0-9]|\\\\u00B7|[\\u0300-\\u036F]|[\\u203F-\\u2040])|.))*((([A-Z]|[a-z]|[\\u00C0-\\u00D6]|[\\u00D8-\\u00F6]|[\\u00F8-\\u02FF]|[\\u0370-\\u037D]|[\\u037F-\\u1FFF]|[\\u200C-\\u200D]|[\\u2070-\\u218F]|[\\u2C00-\\u2FEF]|[\\u3001-\\uD7FF]|[\\uF900-\\uFDCF]|[\\uFDF0-\\uFFFD]|[\\u10000-\\uEFFFF])|_)|-|[0-9]|\\\\u00B7|[\\u0300-\\u036F]|[\\u203F-\\u2040])))?))$");
  TT(IRIREF, "^(((((([A-Z]|[a-z]|[\\u00C0-\\u00D6]|[\\u00D8-\\u00F6]|[\\u00F8-\\u02FF]|[\\u0370-\\u037D]|[\\u037F-\\u1FFF]|[\\u200C-\\u200D]|[\\u2070-\\u218F]|[\\u2C00-\\u2FEF]|[\\u3001-\\uD7FF]|[\\uF900-\\uFDCF]|[\\uFDF0-\\uFFFD]|[\\u10000-\\uEFFFF])|_)|-|[0-9]|\\\\u00B7|[\\u0300-\\u036F]|[\\u203F-\\u2040])|.|:|\\/|\\\\\\\\|#|@|%|&|((\\\\\\\\u([0-9]|[A-F]|[a-f])([0-9]|[A-F]|[a-f])([0-9]|[A-F]|[a-f])([0-9]|[A-F]|[a-f]))|(\\\\\\\\U([0-9]|[A-F]|[a-f])([0-9]|[A-F]|[a-f])([0-9]|[A-F]|[a-f])([0-9]|[A-F]|[a-f])([0-9]|[A-F]|[a-f])([0-9]|[A-F]|[a-f])([0-9]|[A-F]|[a-f])([0-9]|[A-F]|[a-f])))))*)$");

  struct TripleConstraint ;
  struct OneOf ;
  struct EachOf ;
  struct Shape ;
  struct Wildcard ;
  struct LanguageStemRange ;
  struct LanguageStem ;
  struct Language ;
  struct LiteralStemRange ;
  struct LiteralStem ;
  struct IriStemRange ;
  struct IriStem ;
  struct NodeConstraint ;
  struct ShapeExternal ;
  struct ShapeNot ;
  struct ShapeAnd ;
  struct ShapeOr ;
  struct Schema ;
  struct ObjectLiteral ;
  using shapeExprLabel = boost::variant< IRIREF, BNODE > ;
  using shapeExprRef = shapeExprLabel;
  using shapeExpr = boost::variant< ShapeOr&, ShapeAnd&, ShapeNot&, NodeConstraint&, Shape&, ShapeExternal&, shapeExprRef& >;
  using tripleExprLabel = boost::variant< IRIREF, BNODE >;
  using tripleExprRef = tripleExprLabel;
  using tripleExpr = boost::variant< EachOf, OneOf, TripleConstraint, tripleExprRef >;
  using objectValue = boost::variant< IRIREF, ObjectLiteral >;

  struct ObjectLiteral {
    STRING value;
    boost::optional< STRING > language;
    boost::optional< STRING > type;
  };

  struct Annotation {
    IRIREF predicate;
    objectValue object;
  };

  struct SemAct {
    IRIREF name;
    boost::optional< STRING > code;
  };

  struct TripleConstraint {
    TripleConstraint (std::string predicate,
                      bool inverse = false,
                      boost::optional< shapeExpr > valueExpr = boost::none,
                      long min = 1,
                      long max = 1,
                      std::vector< SemAct > semActs = std::vector< SemAct >(),
                      std::vector< Annotation > annotations = std::vector< Annotation >(),
                      boost::optional< tripleExprLabel > id = boost::none
                      ) :
      predicate(predicate), inverse(inverse),
      valueExpr(valueExpr), min(min), max(max), 
      semActs(semActs), annotations(annotations), id(id)
    {  }
    IRIREF predicate;
    bool inverse;
    boost::optional< shapeExpr > valueExpr;
    long min;
    long max;
    std::vector< SemAct > semActs;
    std::vector< Annotation > annotations;
    boost::optional< tripleExprLabel > id;
  };

  struct OneOf {
    OneOf (std::vector< tripleExpr > expressions = std::vector< tripleExpr >(),
           long min = 1,
           long max = 1,
           std::vector< SemAct > semActs = std::vector< SemAct >(),
           std::vector< Annotation > annotations = std::vector< Annotation >()
           ) :
      expressions(expressions), min(min), max(max), 
      semActs(semActs), annotations(annotations)
    {  }
    boost::optional< tripleExprLabel > id;
    std::vector< tripleExpr > expressions;
    long min;
    long max;
    std::vector< SemAct > semActs;
    std::vector< Annotation > annotations;
  };

  struct EachOf {
    EachOf (std::vector< tripleExpr > expressions = std::vector< tripleExpr >(),
            long min = 1,
            long max = 1,
            std::vector< SemAct > semActs = std::vector< SemAct >(),
            std::vector< Annotation > annotations = std::vector< Annotation >()
            ) :
      expressions(expressions), min(min), max(max), 
      semActs(semActs), annotations(annotations)
    {  }
    boost::optional< tripleExprLabel > id;
    std::vector< tripleExpr > expressions;
    long min;
    long max;
    std::vector< SemAct > semActs;
    std::vector< Annotation > annotations;
  };

  struct Shape {
    Shape (bool closed = false,
           std::vector< IRIREF > extra = std::vector< IRIREF >(),
           boost::optional< tripleExpr > expression = boost::none,
           std::vector< SemAct > semActs = std::vector< SemAct >(),
           std::vector< Annotation > annotations = std::vector< Annotation >()
           ) :
      closed(closed), extra(extra),
      expression(expression), semActs(semActs), annotations(annotations)
    {
    }
    bool closed;
    std::vector< IRIREF > extra;
    boost::optional< tripleExpr > expression;
    std::vector< SemAct > semActs;
    std::vector< Annotation > annotations;
  };

  struct Wildcard {
    /* empty */
  };

  struct LanguageStemRange {
    boost::variant< LANGTAG, Wildcard > stem;
    std::vector< boost::variant< LANGTAG, LanguageStem > > exclusions;
  };

  struct LanguageStem {
    boost::variant< LANGTAG, Wildcard > stem;
  };

  struct Language {
    LANGTAG languageTag;
  };

  struct LiteralStemRange {
    boost::variant< STRING, Wildcard > stem;
    std::vector< boost::variant< STRING, LiteralStem > > exclusions;
  };

  struct LiteralStem {
    STRING stem;
  };

  struct IriStemRange {
    boost::variant< IRIREF, Wildcard > stem;
    std::vector< boost::variant< IRIREF, IriStem > > exclusions;
  };

  struct IriStem {
    IRIREF stem;
  };

  using valueSetValue = boost::variant< objectValue, IriStem, IriStemRange, LiteralStem, LiteralStemRange, Language, LanguageStem, LanguageStemRange >;
  using numericLiteral = boost::variant< long, float, double >;
  struct valueFacet {
    enum {mininclusive, minexclusive, maxinclusive, maxexclusive} property;
    numericLiteral type;
  };

  struct digitsFacet {
    enum {totaldigits, fractiondigits} property;
    long type;
  };

  using numericFacet = boost::variant< struct valueFacet, struct digitsFacet >;
  struct lengthFacet {
    enum {length, minlength, maxlength} property;
    long type;
  };

  struct regexFacet {
    STRING pattern;
    boost::optional< STRING > flags;
  };

  using stringFacet = boost::variant< struct lengthFacet, struct regexFacet >;
  using xsFacet = boost::variant< stringFacet, numericFacet >;
  enum e_nodeKind { iri, bnode, nonliteral, literal };
  struct NodeConstraint {
    boost::optional< enum e_nodeKind > nodeKind;
    boost::optional< IRIREF > datatype;
    std::vector< xsFacet > facet;
    std::vector< valueSetValue > values;
  };

  struct ShapeExternal {
  };

  struct ShapeNot {
    shapeExpr shapeExpr;
  };

  struct ShapeAnd {
    std::vector< shapeExpr > shapeExprs;
  };

  struct ShapeOr {
    std::vector< shapeExpr > shapeExprs;
  };

  struct labeledShapeExpr {
    shapeExprLabel label;
    shapeExpr shapeExpr;
    labeledShapeExpr (std::string label, ::shex::shapeExpr shapeExpr)
      : label(IRIREF(label)), shapeExpr(shapeExpr)
    {
      if (label.rfind("_:") == 0)
        this->label = BNODE(label);
    }

    std::ostream& toJSON (std::ostream& os) const {
      return shapeExpr.toJSON(label);
    }
  };

  struct Schema {
    // boost::optional< std::string("http://www.w3.org/ns/shex.jsonld") > context;
    std::vector< IRIREF > imports;
    std::vector< SemAct > startActs;
    boost::optional< shapeExpr > start;
    std::vector< labeledShapeExpr > shapes;
    struct {
      std::map< shapeExprLabel, shapeExpr& > shapeExprs;
      std::map< shapeExprLabel, tripleExpr& > tripleExprs;
    } _index;

    std::ostream& toJSON (std::ostream& os) const {
      os << "{\"type\":\"Schema\"";
      for (auto &sh : shapes)
        sh.toJSON(os);
      os << "}";
      return os;
    }
  };
  std::ostream& operator<< (std::ostream &os, const Schema &s) {
    return s.toJSON(os);
  }

} // namespace shex

#endif // shex_H
