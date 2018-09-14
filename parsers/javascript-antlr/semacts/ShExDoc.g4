// ANTLR4 Equivalent of accompanying bnf, developed in
// http://www.w3.org/2005/01/yacker/uploads/ShEx3
// Updated to Jul 27 AM ShEx3
// Updated to Aug 23 AM ShEx3 (last change was EGP 20150820)
// Sept 21 AM disallow single internal unary (e.g. {(:p .{2}){3}}
//            Change (non-standard) "codeLabel" to "productionName"
// Oct 26 - change annotation predicate to include rdftype (how did this slip in to the production rules?
// Dec 30 - update to match http://www.w3.org/2005/01/yacker/uploads/ShEx2/bnf with last change "EGP 20151120"
// May 23, 2016 - Update to match http://www.w3.org/2005/01/yacker/uploads/ShEx2/bnf with last change "EGP20160520" AND ';' separator and '//' for annotations
// May 24, 2016 - EGP20150424
// Aug 11, 2016 - EGP20160708
// Sep 14, 2016 - Revised to match Eric's latest reshuffle
// Sep 24, 2016 - Switched to TT grammar (vs inner and outer shapes)
// Sep 26, 2016 - Refactored to match https://raw.githubusercontent.com/shexSpec/shex.js/7eb770fe2b5bab9edfe9558dc07bb6f6dcdf5d23/doc/bnf
// Oct 27, 2016 - Added comments to '*', '*' and '?' to facilitate parsing
// Oct 27, 2016 - Added qualifier rule to be reused by shapeDefinition and inlineShapeDefinition
// Oct 27, 2016 - Added negation rule
// Mar 03, 2017 - removed ^^-style facet arguments per shex#41
// Mar 03, 2017 - switch to ~/regexp/
// Apr 09, 2017 - removed WS fragment (unused)
// Apr 09, 2017 - revise REGEXP definition
// Apr 09, 2017 - factor out REGEXP_FLAGS so we don't have to parse them out
// Apr 09, 2017 - literalRange / languageRange additions
// Apr 09, 2017 - factor out shapeRef to match spec
// Apr 09, 2017 - update repeatRange to allow differentiation of {INTEGER} and {INTEGER,}
// Apr 09, 2017 - add STEM_MARK and UNBOUNDED tokens to eliminate lex token parsingf
// Apr 17, 2018 - add 2.1 rules -- extensions, restrictions and ABSTRACT


grammar ShExDoc;

shExDoc
@init {
  this.UNBOUNDED = -1;

  this.unescapeText = function (string, replacements) {
    var regex = /\\u([a-fA-F0-9]{4})|\\U([a-fA-F0-9]{8})|\\(.)/g;
    try {
      string = string.replace(regex, function (sequence, unicode4, unicode8, escapedChar) {
        var charCode;
        if (unicode4) {
          charCode = parseInt(unicode4, 16);
          if (isNaN(charCode)) throw new Error(); // can never happen (regex), but helps performance
          return String.fromCharCode(charCode);
        }
        else if (unicode8) {
          charCode = parseInt(unicode8, 16);
          if (isNaN(charCode)) throw new Error(); // can never happen (regex), but helps performance
          if (charCode < 0xFFFF) return String.fromCharCode(charCode);
          return String.fromCharCode(0xD800 + ((charCode -= 0x10000) >> 10), 0xDC00 + (charCode & 0x3FF));
        }
        else {
          var replacement = replacements[escapedChar];
          if (!replacement) throw new Error("no replacement found for '" + escapedChar + "'");
          return replacement;
        }
      });
      return string;
    }
    catch (error) { console.warn(error); return ''; }
  };

  // Common namespaces and entities
  this.RDF = 'http://www.w3.org/1999/02/22-rdf-syntax-ns#',
  this.RDF_TYPE  = this.RDF + 'type',
  this.RDF_FIRST = this.RDF + 'first',
  this.RDF_REST  = this.RDF + 'rest',
  this.RDF_NIL   = this.RDF + 'nil',
  this.XSD = 'http://www.w3.org/2001/XMLSchema#',
  this.XSD_INTEGER  = this.XSD + 'integer',
  this.XSD_DECIMAL  = this.XSD + 'decimal',
  this.XSD_FLOAT   = this.XSD + 'float',
  this.XSD_DOUBLE   = this.XSD + 'double',
  this.XSD_BOOLEAN  = this.XSD + 'boolean',
  this.XSD_TRUE =  '"true"^^'  + this.XSD_BOOLEAN,
  this.XSD_FALSE = '"false"^^' + this.XSD_BOOLEAN,
  this.XSD_PATTERN        = this.XSD + 'pattern',
  this.XSD_MININCLUSIVE   = this.XSD + 'minInclusive',
  this.XSD_MINEXCLUSIVE   = this.XSD + 'minExclusive',
  this.XSD_MAXINCLUSIVE   = this.XSD + 'maxInclusive',
  this.XSD_MAXEXCLUSIVE   = this.XSD + 'maxExclusive',
  this.XSD_LENGTH         = this.XSD + 'length',
  this.XSD_MINLENGTH      = this.XSD + 'minLength',
  this.XSD_MAXLENGTH      = this.XSD + 'maxLength',
  this.XSD_TOTALDIGITS    = this.XSD + 'totalDigits',
  this.XSD_FRACTIONDIGITS = this.XSD + 'fractionDigits';

  this.numericDatatypes = [
      this.XSD + "integer",
      this.XSD + "decimal",
      this.XSD + "float",
      this.XSD + "double",
      this.XSD + "string",
      this.XSD + "boolean",
      this.XSD + "dateTime",
      this.XSD + "nonPositiveInteger",
      this.XSD + "negativeInteger",
      this.XSD + "long",
      this.XSD + "int",
      this.XSD + "short",
      this.XSD + "byte",
      this.XSD + "nonNegativeInteger",
      this.XSD + "unsignedLong",
      this.XSD + "unsignedInt",
      this.XSD + "unsignedShort",
      this.XSD + "unsignedByte",
      this.XSD + "positiveInteger"
  ];

  this.absoluteIRI = /^[a-z][a-z0-9+.-]*:/i,
    schemeAuthority = /^(?:([a-z][a-z0-9+.-]*:))?(?:\/\/[^\/]*)?/i,
    dotSegments = /(?:^|\/)\.\.?(?:$|[\/#?])/;

  this.numericFacets = ["mininclusive", "minexclusive",
                       "maxinclusive", "maxexclusive"];

  // Returns a lowercase version of the given string
  this.lowercase = function (string) {
    return string.toLowerCase();
  }

  // Appends the item to the array and returns the array
  function appendTo(array, item) {
    return array.push(item), array;
  }

  // Appends the items to the array and returns the array
  function appendAllTo(array, items) {
    return array.push.apply(array, items), array;
  }

  // Extends a base object with properties of other objects
  this.extend = function (base) {
    if (!base) base = {};
    for (var i = 1, l = arguments.length, arg; i < l && (arg = arguments[i] || {}); i++)
      for (var name in arg)
        base[name] = arg[name];
    return base;
  }

  // Creates an array that contains all items of the given arrays
  this.unionAll = function() {
    var union = [];
    for (var i = 0, l = arguments.length; i < l; i++)
      union = union.concat.apply(union, arguments[i]);
    return union;
  }

  // N3.js:lib/N3Parser.js<0.4.5>:58 with
  //   s/this\./this./g
  // ### `_setBase` sets the base IRI to resolve relative IRIs.
  this._setBase = function (baseIRI) {
    if (!baseIRI)
      baseIRI = null;

    // baseIRI '#' check disabled to allow -x 'data:text/shex,...#'
    // else if (baseIRI.indexOf('#') >= 0)
    //   throw new Error('Invalid base IRI ' + baseIRI);

    // Set base IRI and its components
    if (this._base = baseIRI) {
      this._basePath   = baseIRI.replace(/[^\/?]*(?:\?.*)?$/, '');
      baseIRI = baseIRI.match(schemeAuthority);
      this._baseRoot   = baseIRI[0];
      this._baseScheme = baseIRI[1];
    }
  }

  // N3.js:lib/N3this.js<0.4.5>:576 with
  //   s/this\./this./g
  //   s/token/iri/
  // ### `_resolveIRI` resolves a relative IRI token against the base path,
  // assuming that a base path has been set and that the IRI is indeed relative.
  this._resolveIRI = function (iri) {
    switch (iri[0]) {
    // An empty relative IRI indicates the base IRI
    case undefined: return this._base;
    // Resolve relative fragment IRIs against the base IRI
    case '#': return this._base + iri;
    // Resolve relative query string IRIs by replacing the query string
    case '?': return this._base.replace(/(?:\?.*)?$/, iri);
    // Resolve root-relative IRIs at the root of the base IRI
    case '/':
      // Resolve scheme-relative IRIs to the scheme
      return (iri[1] === '/' ? this._baseScheme : this._baseRoot) + _removeDotSegments(iri);
    // Resolve all other IRIs at the base IRI's path
    default: {
      return _removeDotSegments(this._basePath + iri);
    }
    }
  }

  // ### `_removeDotSegments` resolves './' and '../' path segments in an IRI as per RFC3986.
  function _removeDotSegments (iri) {
    // Don't modify the IRI if it does not contain any dot segments
    if (!dotSegments.test(iri))
      return iri;

    // Start with an imaginary slash before the IRI in order to resolve trailing './' and '../'
    var result = '', length = iri.length, i = -1, pathStart = -1, segmentStart = 0, next = '/';

    while (i < length) {
      switch (next) {
      // The path starts with the first slash after the authority
      case ':':
        if (pathStart < 0) {
          // Skip two slashes before the authority
          if (iri[++i] === '/' && iri[++i] === '/')
            // Skip to slash after the authority
            while ((pathStart = i + 1) < length && iri[pathStart] !== '/')
              i = pathStart;
        }
        break;
      // Don't modify a query string or fragment
      case '?':
      case '#':
        i = length;
        break;
      // Handle '/.' or '/..' path segments
      case '/':
        if (iri[i + 1] === '.') {
          next = iri[++i + 1];
          switch (next) {
          // Remove a '/.' segment
          case '/':
            result += iri.substring(segmentStart, i - 1);
            segmentStart = i + 1;
            break;
          // Remove a trailing '/.' segment
          case undefined:
          case '?':
          case '#':
            return result + iri.substring(segmentStart, i) + iri.substr(i + 1);
          // Remove a '/..' segment
          case '.':
            next = iri[++i + 1];
            if (next === undefined || next === '/' || next === '?' || next === '#') {
              result += iri.substring(segmentStart, i - 2);
              // Try to remove the parent path from result
              if ((segmentStart = result.lastIndexOf('/')) >= pathStart)
                result = result.substr(0, segmentStart);
              // Remove a trailing '/..' segment
              if (next !== '/')
                return result + '/' + iri.substr(i + 1);
              segmentStart = i + 1;
            }
          }
        }
      }
      next = iri[++i];
    }
    return result + iri.substring(segmentStart);
  }

  // Creates an expression with the given type and attributes
  function expression(expr, attr) {
    var expression = { expression: expr };
    if (attr)
      for (var a in attr)
        expression[a] = attr[a];
    return expression;
  }

  // Creates a path with the given type and items
  function path(type, items) {
    return { type: 'path', pathType: type, items: items };
  }

  // Creates a literal with the given value and type
  this.createLiteral = function (value, type) {
    return { value: value, type: type };
  }

  // Creates a new blank node identifier
  function blank() {
    return '_:b' + blankId++;
  };
  var blankId = 0;
  this._resetBlanks = function () { blankId = 0; }
  this._reset = function () {
    this._prefixes = this._imports = this.valueExprDefns = this._shapes = this.productions = this._start = this.startActs = null; // Reset state.
    this._base = this._baseIRI = this._baseIRIPath = this._baseIRIRoot = null;
  }
  var _fileName; // for debugging
  this._setFileName = function (fn) { _fileName = fn; }

  // Regular expression and replacement strings to escape strings
  this.stringEscapeReplacements = { '\\': '\\', "'": "'", '"': '"',
                                   't': '\t', 'b': '\b', 'n': '\n', 'r': '\r', 'f': '\f' };
  this.semactEscapeReplacements = { '\\': '\\', '%': '%' };
  this.pnameEscapeReplacements = {
        '\\': '\\', "'": "'", '"': '"',
        'n': '\n', 'r': '\r', 't': '\t', 'f': '\f', 'b': '\b',
        '_': '_', '~': '~', '.': '.', '-': '-', '!': '!', '$': '$', '&': '&',
        '(': '(', ')': ')', '*': '*', '+': '+', ',': ',', ';': ';', '=': '=',
        '/': '/', '?': '?', '#': '#', '@': '@', '%': '%',
      };


  // Translates string escape codes in the string into their textual equivalent
  this.unescapeString = function (string, trimLength) {
    string = string.substring(trimLength, string.length - trimLength);
    return { value: this.unescapeText(string, this.stringEscapeReplacements) };
  }

  this.unescapeLangString = function (string, trimLength) {
    var at = string.lastIndexOf("@");
    var lang = string.substr(at);
    string = string.substr(0, at);
    var u = unescapeString(string, trimLength);
    return this.extend(u, { language: this.lowercase(lang.substr(1)) });
  }

  // Translates regular expression escape codes in the string into their textual equivalent
  this.unescapeRegexp = function (regexp) {
    var end = regexp.lastIndexOf("/");
    var s = regexp.substr(1, end-1);
    var regexpEscapeReplacements = {
      '.': "\\.", '\\': "\\\\", '?': "\\?", '*': "\\*", '+': "\\+",
      '{': "\\{", '}': "\\}", '(': "\\(", ')': "\\)", '|': "\\|",
      '^': "\\^", '$': "\\$", '[': "\\[", ']': "\\]", '/': "\\/",
      't': '\\t', 'n': '\\n', 'r': '\\r', '-': "\\-", '/': '/'
    };
    s = this.unescapeText(s, regexpEscapeReplacements)
    var ret = {
      pattern: s
    };
    if (regexp.length > end+1)
      ret.flags = regexp.substr(end+1);
    return ret;
  }

  // Convenience function to return object with p1 key, value p2
  this.keyValObject = function (key, val) {
    var ret = {};
    ret[key] = val;
    return ret;
  }

  // Return object with p1 key, p2 string value
  this.unescapeSemanticAction = function (key, string) {
    string = string.substring(1, string.length - 2);
    return {
      type: "SemAct",
      name: key,
      code: this.unescapeText(string, this.semactEscapeReplacements)
    };
  }

  this.error = function (msg) {
    this._reset();
    throw new Error(msg);
  }

  // Expand declared prefix or throw Error
  this.expandPrefix = function (prefix) {
    if (!(prefix in this._prefixes))
      this.error('Parse error; unknown prefix: ' + prefix);
    return this._prefixes[prefix];
  }

  // Add a shape to the map
  this.addShape = function (label, shape) {
    if (this.productions && label in this.productions)
      this.error("Structural error: "+label+" is a shape");
    if (!this._shapes)
      this._shapes = [];
    // if (label in this.shapes) {
    //   if (this.options.duplicateShape === "replace")
    //     this.shapes[label] = shape;
    //   else if (this.options.duplicateShape !== "ignore")
    //     this.error("Parse error: "+label+" already defined");
    // } else
    shape.id = label
      this._shapes.push(shape);
  }

  // Add a production to the map
  this.addProduction = function (label, production) {
    if (this.shapes && label in this.shapes)
      this.error("Structural error: "+label+" is a shape");
    if (!this.productions)
      this.productions = {};
    if (label in this.productions) {
      if (this.options.duplicateShape === "replace")
        this.productions[label] = production;
      else if (this.options.duplicateShape !== "ignore")
        this.error("Parse error: "+label+" already defined");
    } else
      this.productions[label] = production;
  }

  this.shapeJunction = function (type, container, elts) {
    if (elts.length === 0) {
      return container;
    } else if (container.type === type) {
      container.shapeExprs = container.shapeExprs.concat(elts);
      return container;
    } else {
      return { type: type, shapeExprs: [container].concat(elts) };
    }
  }

  this.EmptyObject = {  };
  this.EmptyShape = { type: "Shape" };

  this.NC = function (l, r) {
    let facets = r.reduce((acc, p) => Object.assign(acc, p.$$), {})      
    return this.extend({ type: "NodeConstraint"}, l, facets)
  }

            this.$1 = function (ctx) {
                ctx.$$ = ctx.children[1-1].$$;
            }
            this.$2 = function (ctx) {
                ctx.$$ = ctx.children[2-1].$$;
            }

            this._prefixes = Object.create(null);
            this._imports = [];
            this._setBase('');
            this.url = require('url');
        }
 		:  directive* ((notStartAction | startActions) statement*)? EOF {
        var valueExprDefns = this.valueExprDefns ? { valueExprDefns: this.valueExprDefns } : {};
        var startObj = this._start ? { start: this._start } : {};
        var startActs = this.startActs ? { startActs: this.startActs } : {};
        var ret = this.extend({ type: "Schema"},
                         Object.keys(this._prefixes).length ? { prefixes: this._prefixes } : {}, // Properties ordered here to
                         Object.keys(this._imports).length ? { imports: this._imports } : {}, // build return object from
                         valueExprDefns, startActs, startObj,                  // components in parser state
                         this._shapes ? {shapes: this._shapes} : {},         // maintaining intuitve order.
                         this.productions ? {productions: this.productions} : {});
        if (this._base !== null)
          ret.base = this._base;
        this._reset();
        localctx.$$ = ret;
      };  // leading CODE
directive       : baseDecl
				| prefixDecl
				| importDecl
				;
baseDecl 		: KW_BASE  IRIREF {
            this._setBase(this._base === null ||
                    this.absoluteIRI.test(localctx.children[2-1].getText().slice(1, -1)) ? localctx.children[2-1].getText().slice(1, -1) : this._resolveIRI(localctx.children[2-1].getText().slice(1, -1)))
        };
prefixDecl		: KW_PREFIX PNAME_NS IRIREF	{ // t: ShExParser-test.js/with pre-defined prefixes
        this._prefixes[localctx.children[2 - 1].getText().slice(0, -1)] = this._base === null ||
                    this.absoluteIRI.test(localctx.children[3-1].getText().slice(1, -1)) ? localctx.children[3-1].getText().slice(1, -1) : this._resolveIRI(localctx.children[3-1].getText().slice(1, -1));
      } ;
importDecl      : KW_IMPORT IRIREF	{
        this._imports.push(this._base === null ||
                    this.absoluteIRI.test(localctx.children[2-1].getText().slice(1, -1)) ? localctx.children[2-1].getText().slice(1, -1) : this._resolveIRI(localctx.children[2-1].getText().slice(1, -1)));
      } ;
notStartAction  : start | shapeExprDecl ;
start           : KW_START '=' shapeExpression	{ this._start = localctx.children[3-1].$$ } ;
startActions	: semanticAction+	{ this.startActs = localctx.children.map(c => c.$$); } ;
statement 		: directive | notStartAction ;
shapeExprDecl   : /* KW_ABSTRACT? */ shapeExprLabel /* restrictions* */ (shapeExpression | KW_EXTERNAL)	{ // t: 1dot 1val1vsMinusiri3??
        this.addShape(localctx.children[0].$$, localctx.KW_EXTERNAL() ? { type: "ShapeExternal" } : localctx.shapeExpression().$$);
      } ;
shapeExpression : shapeOr	{ this.$1(localctx); } ;
inlineShapeExpression : inlineShapeOr	{ this.$1(localctx); } ;
shapeOr  		: shapeAnd (KW_OR shapeAnd)*	{ localctx.$$ = localctx.shapeAnd().length == 1 ? localctx.shapeAnd()[0].$$ : { type: "ShapeOr", shapeExprs: localctx.shapeAnd().map(c => c.$$) };  };
inlineShapeOr   : inlineShapeAnd (KW_OR inlineShapeAnd)*	{ localctx.$$ = localctx.inlineShapeAnd().length == 1 ? localctx.inlineShapeAnd()[0].$$ : { type: "ShapeOr", shapeExprs: localctx.inlineShapeAnd().map(c => c.$$) }; };
shapeAnd		: shapeNot (KW_AND shapeNot)*	{ localctx.$$ = localctx.shapeNot().length == 1 ? localctx.shapeNot()[0].$$ : { type: "ShapeAnd", shapeExprs: localctx.shapeNot().map(c => c.$$) }; } ;
// inlineShapeAnd  : inlineShapeNot (KW_AND inlineShapeNot)*	{ localctx.$$ = localctx.inlineShapeNot().length == 1 ? localctx.inlineShapeNot()[0].$$ : { type: "ShapeAnd", shapeExprs: localctx.inlineShapeNot().map(c => c.$$) }; } ;
inlineShapeAnd  : inlineShapeNot (KW_AND inlineShapeNot)*	{ localctx.$$ = this.shapeJunction("ShapeAnd", localctx.inlineShapeNot()[0].$$, localctx.inlineShapeNot().slice(1).map(c => c.$$)) } ;
shapeNot	    : KW_NOT? shapeAtom	{ localctx.$$ = localctx.KW_NOT() ? { type: "ShapeNot", "shapeExpr": localctx.shapeAtom().$$ } : localctx.shapeAtom().$$; } ;
shapeAtom		: nonLitNodeConstraint shapeOrRef?	{
        localctx.$$ = localctx.children[2-1] ? { type: "ShapeAnd", shapeExprs: [ this.extend({ type: "NodeConstraint" }, localctx.children[1-1].$$), localctx.children[2-1].$$ ] } : localctx.children[1-1].$$
      }    # shapeAtomNonLitNodeConstraint
                | litNodeConstraint	{ this.$1(localctx); }             # shapeAtomLitNodeConstraint
				| shapeOrRef nonLitNodeConstraint?	{
        localctx.$$ = localctx.children[2-1] ? this.shapeJunction("ShapeAnd", localctx.children[1-1].$$, localctx.children[2-1].$$) /* t: 1dotRef1 */ : localctx.children[1-1].$$ // t:@@
      }    # shapeAtomShapeOrRef
				| '(' shapeExpression ')'	{ this.$2(localctx); }		# shapeAtomShapeExpression
				| '.'	{ localctx.$$ = this.EmptyShape; }							# shapeAtomAny			// no constraint
				;
inlineShapeAtom : inlineNonLitNodeConstraint inlineShapeOrRef?	{
        localctx.$$ = localctx.children[2-1] ? { type: "ShapeAnd", shapeExprs: [ this.extend({ type: "NodeConstraint" }, localctx.children[1-1].$$), localctx.children[2-1].$$ ] } : localctx.children[1-1].$$
      } # inlineShapeAtomNonLitNodeConstraint
                | inlineLitNodeConstraint	{ this.$1(localctx); }             # inlineShapeAtomLitNodeConstraint
				| inlineShapeOrRef inlineNonLitNodeConstraint?	{
        localctx.$$ = localctx.children[2-1] ? this.shapeJunction("ShapeAnd", localctx.children[1-1].$$, localctx.children[2-1].$$) /* t: 1dotRef1 */ : localctx.children[1-1].$$ // t:@@
      } # inlineShapeAtomShapeOrRef
				| '(' shapeExpression ')'	{ this.$2(localctx); }		# inlineShapeAtomShapeExpression
				| '.' { localctx.$$ = this.EmptyShape; }							# inlineShapeAtomAny   // no constraint
				;
shapeOrRef      : shapeDefinition	{ this.$1(localctx); }
				| shapeRef	{ this.$1(localctx); }
				;
inlineShapeOrRef : inlineShapeDefinition	{ this.$1(localctx); }
				| shapeRef	{ this.$1(localctx); }
				;
shapeRef 		: ATPNAME_LN	{ // t: 1dotRefLNex@@
        let ln = localctx.children[1-1].getText();
        ln = ln.substr(1, ln.length-1);
        var namePos = ln.indexOf(':');
        localctx.$$ = this.expandPrefix(ln.substr(0, namePos)) + ln.substr(namePos + 1);
      }
				| ATPNAME_NS	{ // t: 1dotRefNS1@@
        let ns = localctx.children[1-1].getText();
        ns = ns.substr(1, ns.length-1);
        localctx.$$ = this.expandPrefix(ns.substr(0, ns.length - 1));
      }
				| '@' shapeExprLabel	{ localctx.$$ = localctx.children[2-1].$$; } // t: 1dotRef1, 1dotRefSpaceLNex, 1dotRefSpaceNS1
				;
inlineLitNodeConstraint : KW_LITERAL xsFacet*	{ localctx.$$ = this.NC({nodeKind: "literal"}, localctx.children.slice(1)); }	# nodeConstraintLiteral
				| nonLiteralKind stringFacet*	{ localctx.$$ = this.NC(localctx.children[0].$$, localctx.children.slice(1)); }	# nodeConstraintNonLiteral
				| datatype xsFacet*	{ localctx.$$ = this.NC({ datatype: localctx.children[0].$$ }, localctx.children.slice(1)); }				# nodeConstraintDatatype // t: 1datatype
				| valueSet xsFacet*	{ localctx.$$ = this.NC({ values: localctx.children[0].$$ }, localctx.children.slice(1)); }			# nodeConstraintValueSet
				| numericFacet+	{ localctx.$$ = this.NC({}, localctx.children); }					# nodeConstraintNumericFacet
				;
litNodeConstraint : inlineLitNodeConstraint  annotation* semanticAction*	{
        this.$1(localctx);
        if (localctx.annotation().length) { localctx.$$.annotations = localctx.annotation().map(c => c.$$); }
        if (localctx.semanticAction().length) { localctx.$$.semActs = localctx.semanticAction().map(c => c.$$); }
      } ;
inlineNonLitNodeConstraint  : nonLiteralKind stringFacet*	{
        localctx.$$ = localctx.children.slice(1).reduce((acc, p) => Object.assign(acc, p.$$), Object.assign({ type: "NodeConstraint" }, localctx.children[0].$$))
      }	# litNodeConstraintLiteral
                | stringFacet+	{
        localctx.$$ = localctx.children.reduce((acc, c) => Object.assign(acc, c.$$), {type: "NodeConstraint"});
      }                  # litNodeConstraintStringFacet
				;
nonLitNodeConstraint : inlineNonLitNodeConstraint  annotation* semanticAction*	{ // t: !!
        localctx.$$ = localctx.children[0].$$;
        if (localctx.annotation().length) { localctx.$$.annotations = localctx.annotation().map(c => c.$$); } // t: !!
        if (localctx.semanticAction().length) { localctx.$$.semActs = localctx.semanticAction().map(c => c.$$); } // t: !!
      } ;
nonLiteralKind  : KW_IRI	{ localctx.$$ = { nodeKind: "iri" }; }
				| KW_BNODE	{ localctx.$$ = { nodeKind: "bnode" }; }
				| KW_NONLITERAL	{ localctx.$$ = { nodeKind: "nonliteral" }; }
				;
xsFacet			: stringFacet	{ this.$1(localctx); }
				| numericFacet	{ this.$1(localctx); }
				;
stringFacet     : stringLength INTEGER	{ localctx.$$ = this.keyValObject(localctx.children[0].$$, parseInt(localctx.children[1].getText(), 10)); } // t: 1literalLength
				| REGEXP REGEXP_FLAGS?	{ localctx.$$ = this.unescapeRegexp(localctx.children[0].getText()); if (localctx.REGEXP_FLAGS()) {localctx.$$.flags = localctx.REGEXP_FLAGS().getText()} } // t: 1literalPattern
				;
stringLength	: KW_LENGTH	{ localctx.$$ = "length"; }
				| KW_MINLENGTH	{ localctx.$$ = "minlength"; }
				| KW_MAXLENGTH	{ localctx.$$ = "maxlength"; }
                ;
numericFacet	: numericRange rawNumeric	{ localctx.$$ = this.keyValObject(localctx.children[0].$$, localctx.children[1].$$); }
				| numericLength INTEGER	{ localctx.$$ = this.keyValObject(localctx.children[0].$$, parseInt(localctx.children[1].getText(), 10)); }
				;
numericRange	: KW_MININCLUSIVE	{ localctx.$$ = "mininclusive"; }
			    | KW_MINEXCLUSIVE	{ localctx.$$ = "minexclusive"; }
			    | KW_MAXINCLUSIVE	{ localctx.$$ = "maxinclusive"; }
			    | KW_MAXEXCLUSIVE	{ localctx.$$ = "maxexclusive"; }
			    ;
numericLength   : KW_TOTALDIGITS	{ localctx.$$ = "totaldigits"; }
				| KW_FRACTIONDIGITS	{ localctx.$$ = "fractiondigits"; }
				;
rawNumeric		: INTEGER	{ localctx.$$ = parseInt(localctx.children[0].getText()); }
				| DECIMAL	{ localctx.$$ = parseFloat(localctx.children[0].getText()); }
				| DOUBLE	{ localctx.$$ = parseFloat(localctx.children[0].getText()); }
				;
shapeDefinition : inlineShapeDefinition annotation* semanticAction*	{
        localctx.$$ = localctx.children[0].$$;
        if (localctx.annotation().length) { localctx.$$.annotations = localctx.annotation().map(c => c.$$); }
        if (localctx.semanticAction().length) { localctx.$$.semActs = localctx.semanticAction().map(c => c.$$); }
      } ;
inlineShapeDefinition : qualifier* '{' tripleExpression? '}'	{
        var exprObj = localctx.tripleExpression() ? { expression: localctx.tripleExpression().$$ } : this.EmptyObject; // t: 0, 0Inherit1
        localctx.$$ = (exprObj === this.EmptyObject && localctx.qualifier().length === 0) ?
          this.EmptyShape :
          this.extend({ type: "Shape" }, localctx.qualifier().reduce((acc, c) => {
            let k = Object.keys(c.$$)[0]; // only one key/value in qualifier result
            return this.extend(acc, k in acc && k !== "closed" ? this.keyValObject(k, acc[k].concat(c.$$[k])) : c.$$)
          }, {}), exprObj);
      };
qualifier       : /* extension
				| */ extraPropertySet	{ localctx.$$ = { extra: localctx.children[1-1].$$ } }
				| KW_CLOSED	{ localctx.$$ = { closed: true } } ;
extraPropertySet : KW_EXTRA predicate+	{ localctx.$$ = localctx.children.slice(1).map(c => c.$$); } ;
tripleExpression : oneOfTripleExpr	{ this.$1(localctx); } ;
oneOfTripleExpr : groupTripleExpr	{ this.$1(localctx); }
				| multiElementOneOf	{ this.$1(localctx); }
				;
multiElementOneOf : groupTripleExpr ( '|' groupTripleExpr )+	{ localctx.$$ = { type: "OneOf", expressions: localctx.groupTripleExpr().map(c => c.$$) }; } ; // t: 2oneOfdot
groupTripleExpr : singleElementGroup	{ this.$1(localctx); }
				| multiElementGroup	{ this.$1(localctx); }
				;
singleElementGroup : unaryTripleExpr ';'?	{ this.$1(localctx); } ;
multiElementGroup : unaryTripleExpr (';' unaryTripleExpr)+ ';'?	{ localctx.$$ = { type: "EachOf", expressions: localctx.unaryTripleExpr().map(c => c.$$) }; } ; // t: 2groupOfdot
unaryTripleExpr : ('$' tripleExprLabel)? (tripleConstraint | bracketedTripleExpr)	{
        let expr = localctx.tripleConstraint() || localctx.bracketedTripleExpr();
        if (localctx.tripleExprLabel()) {
          localctx.$$ = this.extend({ id: localctx.tripleExprLabel().$$ }, expr.$$);
          this.addProduction(localctx.tripleExprLabel().$$,  localctx.$$);
        } else {
          localctx.$$ = expr.$$
        }
      }
				| include	{ this.$1(localctx); }
				;
bracketedTripleExpr : '(' tripleExpression ')' cardinality? /* onShapeExpr? */ annotation* semanticAction*	{
        localctx.$$ = localctx.tripleExpression().$$;
        let card = localctx.cardinality() ? localctx.cardinality().$$ : {  };
        if ("min" in card) { localctx.$$.min = card.min; } // t: open3groupdotclosecard23Annot3Code2
        if ("max" in card) { localctx.$$.max = card.max; } // t: open3groupdotclosecard23Annot3Code2
        if (localctx.annotation().length) {
          localctx.$$.annotations = (localctx.$$.annotations ? localctx.$$.annotations : []).concat(localctx.annotation().map(c => c.$$));
        }
        if (localctx.semanticAction().length) {
          localctx.$$.semActs = (localctx.$$.semActs ? localctx.$$.semActs : []).concat(localctx.semanticAction().map(c => c.$$));
        }
      } ;
tripleConstraint : senseFlags? predicate inlineShapeExpression cardinality? /* onShapeExpr? */ annotation* semanticAction*	{
        // $6: t: 1dotCode1
        // %6: t: 1inversedotCode1
        localctx.$$ = this.extend({ type: "TripleConstraint" }, localctx.senseFlags() ? localctx.senseFlags().$$ : {}, { predicate: localctx.predicate().$$ }, (localctx.inlineShapeExpression().$$ === this.EmptyShape ? {} : { valueExpr: localctx.inlineShapeExpression().$$ }), localctx.cardinality() ? localctx.cardinality().$$ : {}); // t: 1dot // t: 1inversedot
        if (localctx.annotation().length) { localctx.$$.annotations = localctx.annotation().map(c => c.$$); } // t: 1dotAnnot3 // t: 1inversedotAnnot3
        if (localctx.semanticAction().length) { localctx.$$.semActs = localctx.semanticAction().map(c => c.$$); }
      } ;
cardinality     : '*'        	{ localctx.$$ = { min:0, max:this.UNBOUNDED }; }  # starCardinality
				| '+'        	{ localctx.$$ = { min:1, max:this.UNBOUNDED }; }  # plusCardinality
				| '?'        	{ localctx.$$ = { min:0, max:1 }; }  # optionalCardinality
				| repeatRange	{ this.$1(localctx); }  # repeatCardinality
				;
// BNF: REPEAT_RANGE ::= '{' INTEGER (',' (INTEGER | '*')?)? '}'
repeatRange     : '{' INTEGER '}'								{
        let i = parseInt(localctx.children[1].getText());
        localctx.$$ = { min: i, max: i };
      }  # exactRange
				| '{' INTEGER ',' (INTEGER | UNBOUNDED)? '}'	{
        let j = parseInt(localctx.children[1].getText());
        if (localctx.UNBOUNDED()) {
            localctx.$$ = { min: j, max: this.UNBOUNDED };
        } else if (localctx.INTEGER().length > 1) {
            localctx.$$ = { min: j, max: parseInt(localctx.children[3].getText()) };
        } else {
            localctx.$$ = { min: j, max: -1 };
        }
      }  # minMaxRange
				;
senseFlags      : '^'	{ localctx.$$ = { inverse: true }; } ; // t: 1inversedot
valueSet		: '[' valueSetValue* ']'	{ localctx.$$ = localctx.children.slice(1, localctx.children.length - 1).map(c => c.$$); } ;
valueSetValue   : iriRange	{ this.$1(localctx); }
				| literalRange	{ this.$1(localctx); }
				| languageRange	{ this.$1(localctx); }
				| '.' iriExclusion+	{ localctx.$$ = { type: "IriStemRange", stem: { type: "Wildcard" }, exclusions: localctx.children.slice(1).map(c => c.$$) }; }
				| '.' literalExclusion+	{ localctx.$$ = { type: "LiteralStemRange", stem: { type: "Wildcard" }, exclusions: localctx.children.slice(1).map(c => c.$$) }; }
				| '.' languageExclusion+	{ localctx.$$ = { type: "LanguageStemRange", stem: { type: "Wildcard" }, exclusions: localctx.children.slice(1).map(c => c.$$) }; }
				;
iriRange        : iri (STEM_MARK iriExclusion*)?	{
            if (localctx.STEM_MARK()) {
                localctx.$$ = {  // t: 1val1iriStemMinusiriStem3, 1val1iriStem
                    type: localctx.iriExclusion().length ? "IriStemRange" : "IriStem",
                    stem: localctx.children[1-1].$$
                };
                if (localctx.iriExclusion().length)
                    localctx.$$.exclusions = localctx.iriExclusion().map(c => c.$$); // t: 1val1iriStemMinusiri3
            } else {
                localctx.$$ = localctx.children[1-1].$$; // t: 1val1IRI
            }
        } ;
iriExclusion    : '-' iri STEM_MARK?	{ localctx.$$ = localctx.children[3-1] ? { type: "IriStem", stem: localctx.children[2-1].$$ } /* t: 1val1iriStemMinusiri3 */ : localctx.children[2-1].$$ ; } ; // t: 1val1iriStemMinusiriStem3
literalRange    : literal (STEM_MARK literalExclusion*)?	{
            if (localctx.STEM_MARK()) {
                localctx.$$ = {  // t: 1val1literalStemMinusliteralStem3, 1val1literalStem
                    type: localctx.literalExclusion().length ? "LiteralStemRange" : "LiteralStem",
                    stem: localctx.children[1-1].$$.value
                };
                if (localctx.literalExclusion().length)
                    localctx.$$.exclusions = localctx.literalExclusion().map(c => c.$$); // t: 1val1literalStemMinusliteral3
            } else {
                localctx.$$ = localctx.children[1-1].$$; // t: 1val1LITERAL
            }
        };
literalExclusion : '-' literal STEM_MARK?	{ localctx.$$ = localctx.children[3-1] ? { type: "LiteralStem", stem: localctx.children[2-1].$$.value } /* t: 1val1literalStemMinusliteral3 */ : localctx.children[2-1].$$.value ; } ; // t: 1val1literalStemMinusliteralStem3
languageRange   : LANGTAG (STEM_MARK languageExclusion*)?	{
            if (localctx.STEM_MARK()) {
                localctx.$$ = {  // t: 1val1languageStemMinuslanguageStem3, 1val1languageStem
                    type: localctx.languageExclusion().length ? "LanguageStemRange" : "LanguageStem",
                    stem: localctx.children[1-1].getText().substr(1)
                };
                if (localctx.languageExclusion().length)
                    localctx.$$.exclusions = localctx.languageExclusion().map(c => c.$$); // t: 1val1languageStemMinuslanguage3
            } else {
                localctx.$$ = { type: "Language", languageTag: localctx.children[1-1].getText().substr(1) }; // t: 1val1LANGUAGE
            }
        } ;
languageExclusion : '-' LANGTAG STEM_MARK?	{ localctx.$$ = localctx.children[3-1] ? { type: "LanguageStem", stem: localctx.children[2-1].getText().substr(1) } /* t: 1val1languageStemMinuslanguage3 */ : localctx.children[2-1].getText().substr(1) ; } ; // t: 1val1languageStemMinuslanguageStem3
include			: '&' tripleExprLabel	{ localctx.$$ = localctx.children[1].$$ } ; // t: 2groupInclude1
annotation      : '//' predicate (iri | literal)	{ localctx.$$ = { type: "Annotation", predicate: localctx.children[1].$$, object: localctx.children[2].$$ } } ; // t: 1dotAnnotIRIREF
semanticAction	: '%' iri (CODE | '%')	{
            let i = localctx.children[1].$$;
            localctx.$$ = localctx.CODE() ? this.unescapeSemanticAction(i, localctx.CODE().getText()) /* t: 1dotCode1 */ : { type: "SemAct", name: i } // t: 1dotNoCode1
        } ;
literal         : rdfLiteral	{ this.$1(localctx); }
				| numericLiteral	{ this.$1(localctx); }
				| booleanLiteral	{ this.$1(localctx); }
				;
// BNF: predicate ::= iri | RDF_TYPE
predicate       : iri	{ this.$1(localctx); } // t: 1dot
				| rdfType	{ localctx.$$ = this.RDF_TYPE; } // t: 1AvalA
				;
rdfType			: RDF_TYPE	{ this.$1(localctx); } ;
datatype        : iri	{ this.$1(localctx); } ;
shapeExprLabel  : iri	{ this.$1(localctx); }
				| blankNode	{ this.$1(localctx); }
				;
tripleExprLabel : iri	{ this.$1(localctx); }
				| blankNode	{ this.$1(localctx); }
				;
numericLiteral  : INTEGER	{ localctx.$$ = this.createLiteral(localctx.children[0].getText(), this.XSD_INTEGER); } // t: 1val1INTEGER
				| DECIMAL	{ localctx.$$ = this.createLiteral(localctx.children[0].getText(), this.XSD_DECIMAL); } // t: 1val1DECIMAL
				| DOUBLE	{ localctx.$$ = this.createLiteral(localctx.children[0].getText(), this.XSD_DOUBLE); } // t: 1val1DOUBLE
				;
rdfLiteral      : string (LANGTAG | '^^' datatype)? {
            let s = localctx.children[1 - 1].$$
            localctx.$$ = localctx.LANGTAG()
                ? this.extend(s, { language: this.lowercase(localctx.LANGTAG().getText().substr(1)) })
                : localctx.datatype()
                ? this.extend(s, { type: localctx.datatype().$$ })
                : s
        };
booleanLiteral  : KW_TRUE	{ localctx.$$ = { value: "true", type: this.XSD_BOOLEAN }; } // t: 1val1true
				| KW_FALSE	{ localctx.$$ = { value: "false", type: this.XSD_BOOLEAN }; } // t: 1val1false
				;
string          : STRING_LITERAL_LONG1	{ localctx.$$ = this.unescapeString(localctx.children[1 - 1].getText(), 3); }	// t: 1val1LANGTAG
                | STRING_LITERAL_LONG2	{ localctx.$$ = this.unescapeString(localctx.children[1 - 1].getText(), 3); }	// t: 1val1STRING_LITERAL_LONG1
                | STRING_LITERAL1	{ localctx.$$ = this.unescapeString(localctx.children[1 - 1].getText(), 1); }	// t: 1val1STRING_LITERAL2
				| STRING_LITERAL2	{ localctx.$$ = this.unescapeString(localctx.children[1 - 1].getText(), 1); }	// t: 1val1STRING_LITERAL_LONG2
				;
inlineShapeNot  : KW_NOT? inlineShapeAtom	{ localctx.$$ = localctx.KW_NOT() ? { type: "ShapeNot", "shapeExpr": localctx.inlineShapeAtom().$$ } : localctx.inlineShapeAtom().$$; } ;
onShapeExpr     : KW_ON (KW_SHAPE KW_EXPRESSION)? inlineShapeExpression ;
iri             : IRIREF	{ // t: 1dot
        var unesc = this.unescapeText(localctx.children[0].getText().slice(1,-1), {});
        localctx.$$ = this._base === null || this.absoluteIRI.test(unesc) ? unesc : this._resolveIRI(unesc)
      }
				| prefixedName	{ this.$1(localctx); }
				;
prefixedName    : PNAME_LN	{ // t:1dotPNex, 1dotPNdefault, ShExParser-test.js/with pre-defined prefixes
        var namePos = localctx.children[0].getText().indexOf(':');
        localctx.$$ = this.expandPrefix(localctx.children[0].getText().substr(0, namePos)) + this.unescapeText(localctx.children[0].getText().substr(namePos + 1), this.pnameEscapeReplacements);
      }
				| PNAME_NS	{ // t: 1dotNS2, 1dotNSdefault, ShExParser-test.js/PNAME_NS with pre-defined prefixes
        localctx.$$ = this.expandPrefix(localctx.children[0].getText().substr(0, localctx.children[0].getText().length - 1));
      }
				;
blankNode       : BLANK_NODE_LABEL	{ localctx.$$ = localctx.children[0].getText(); } ;
/*
extension       : KW_EXTENDS shapeExprLabel
                | '&' shapeExprLabel
                ;
restrictions    : KW_RESTRICTS shapeExprLabel
                | '-' shapeExprLabel
                ;
*/

// Keywords
 /* KW_ABSTRACT         : A B S T R A C T ; */
KW_BASE 			: B A S E ;
/* KW_EXTENDS          : E X T E N D S ; */
KW_IMPORT           : I M P O R T ;
/* KW_RESTRICTS        : R E S T R I C T S ; */
KW_EXTERNAL			: E X T E R N A L ;
KW_PREFIX       	: P R E F I X ;
KW_START        	: S T A R T ;
KW_VIRTUAL      	: V I R T U A L ;
KW_CLOSED       	: C L O S E D ;
KW_EXTRA        	: E X T R A ;
KW_LITERAL      	: L I T E R A L ;
KW_IRI          	: I R I ;
KW_NONLITERAL   	: N O N L I T E R A L ;
KW_BNODE        	: B N O D E ;
KW_AND          	: A N D ;
KW_OR           	: O R ;
KW_ON               : O N ;
KW_SHAPE            : S H A P E ;
KW_EXPRESSION       : E X P R E S S I O N ;
KW_MININCLUSIVE 	: M I N I N C L U S I V E ;
KW_MINEXCLUSIVE 	: M I N E X C L U S I V E ;
KW_MAXINCLUSIVE 	: M A X I N C L U S I V E ;
KW_MAXEXCLUSIVE 	: M A X E X C L U S I V E ;
KW_LENGTH       	: L E N G T H ;
KW_MINLENGTH    	: M I N L E N G T H ;
KW_MAXLENGTH    	: M A X L E N G T H ;
KW_TOTALDIGITS  	: T O T A L D I G I T S ;
KW_FRACTIONDIGITS 	: F R A C T I O N D I G I T S ;
KW_NOT				: N O T ;
KW_TRUE         	: 'true' ;
KW_FALSE        	: 'false' ;

// terminals
PASS				  : [ \t\r\n]+ -> skip;
COMMENT				  : ('#' ~[\r\n]*
 					  | '/*' (~[*] | '*' ('\\/' | ~[/]))* '*/') -> skip;

CODE                  : '{' (~[%\\] | '\\' [%\\] | UCHAR)* '%' '}' ;
RDF_TYPE              : 'a' ;
IRIREF                : '<' (~[\u0000-\u0020=<>"{}|^`\\] | UCHAR)* '>' ; /* #x00=NULL #01-#x1F=control codes #x20=space */
PNAME_NS			  : PN_PREFIX? ':' ;
PNAME_LN			  : PNAME_NS PN_LOCAL ;
ATPNAME_NS			  : '@' PN_PREFIX? ':' ;
ATPNAME_LN			  : '@' PNAME_NS PN_LOCAL ;
REGEXP                : '/' (~[/\n\r\\] | '\\' [/nrt\\|.?*+(){}[\]$^-] | UCHAR)+ '/' ;
REGEXP_FLAGS		  : [smix]+ ;
BLANK_NODE_LABEL      : '_:' (PN_CHARS_U | [0-9]) ((PN_CHARS | '.')* PN_CHARS)? ;
LANGTAG               : '@' [a-zA-Z]+ ('-' [a-zA-Z0-9]+)* ;
INTEGER               : [+-]? [0-9]+ ;
DECIMAL               : [+-]? [0-9]* '.' [0-9]+ ;
DOUBLE                : [+-]? ([0-9]+ '.' [0-9]* EXPONENT | '.'? [0-9]+ EXPONENT) ;
STEM_MARK			  : '~' ;
UNBOUNDED             : '*' ;

fragment EXPONENT     : [eE] [+-]? [0-9]+ ;

STRING_LITERAL1       : '\'' (~[\u0027\u005C\u000A\u000D] | ECHAR | UCHAR)* '\'' ; /* #x27=' #x5C=\ #xA=new line #xD=carriage return */
STRING_LITERAL2       : '"' (~[\u0022\u005C\u000A\u000D] | ECHAR | UCHAR)* '"' ;   /* #x22=" #x5C=\ #xA=new line #xD=carriage return */
STRING_LITERAL_LONG1  : '\'\'\'' (('\'' | '\'\'')? (~['\\] | ECHAR | UCHAR))* '\'\'\'' ;
STRING_LITERAL_LONG2  : '"""' (('"' | '""')? (~["\\] | ECHAR | UCHAR))* '"""' ;

fragment UCHAR                 : '\\u' HEX HEX HEX HEX | '\\U' HEX HEX HEX HEX HEX HEX HEX HEX ;
fragment ECHAR                 : '\\' [tbnrf\\"'] ;

fragment PN_CHARS_BASE 		   : [A-Z] | [a-z] | [\u00C0-\u00D6] | [\u00D8-\u00F6] | [\u00F8-\u02FF] | [\u0370-\u037D]
					   		   | [\u037F-\u1FFF] | [\u200C-\u200D] | [\u2070-\u218F] | [\u2C00-\u2FEF] | [\u3001-\uD7FF]
					           | [\uF900-\uFDCF] | [\uFDF0-\uFFFD] | [\uD800-\uDB7F][\uDC00-\uDFFF]
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
