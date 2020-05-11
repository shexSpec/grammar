#include <boost/property_tree/ptree.hpp>
#include <boost/property_tree/json_parser.hpp>
#include <boost/foreach.hpp>
#include <iostream>
#include "shex.hh"

using namespace shex;

int main () {
  Shape s1(false,
           std::vector< IRIREF >(),
           boost::none
       // std::vector< SemAct >(),
       // std::vector< Annotation >()
           );
  TripleConstraint tc1("p1");
  TripleConstraint tc2("p2");
  EachOf e;
  e.expressions.push_back(tc1);
  e.expressions.push_back(tc2);
  tripleExpr v = e;
  boost::optional< tripleExpr > expression = v;
  Schema s/* = {
              std::vector< IRIREF >(),
              std::vector< SemAct >(),
              boost::none,
              std::map< const char*, shapeExpr >()
              // {"S1", se}
              }*/;
  //  s.shapes["S1"] = se;
  labeledShapeExpr ls("S1", s1);
  s.shapes.push_back(labeledShapeExpr( "S1", s1 ));
  s1.toJSON(std::cout, NULL) << std::endl;
  s.toJSON(std::cout) << std::endl;
  return 0;
}
