# $Id: Makefile,v 1.5 2014-04-03 16:51:56 eric Exp $

# recipies:
#   normal build:
#     make shex-parse-to-XML
#   force the use of the tracing facilities (and redirect to stdout):
#     HAVE_BOOST=1 TRACE_FD=1 make -W shex-parse-to-XMLParser.yy test
#   have valgrind start a debugger (works as M-x gdb invocation command):
#     valgrind --db-attach=yes --leak-check=yes shex-parse-to-XML shex-parse-to-XML.txt
#   same, if you aren't working in gdb:
#     HAVE_BOOST=1 TRACE_FD=1 make valgrind
#   debugging in emacs:
#     gdb --annotate=3 shex-parse-to-XML    (set args shex-parse-to-XML.txt)

ifdef HAVE_BOOST
LIBS=-lboost_iostreams
DEFS=-DHAVE_BOOST
else
LIBS=
DEFS=
endif

GPP=g++ -DYYTEXT_POINTER=1 $(DEFS) -W -Wall -Wextra -ansi -g -c
LINK=g++ -W -Wall -Wextra -ansi -g -o

shex-parse-to-XMLParser.cc shex-parse-to-XMLParser.hh location.hh position.hh stack.hh: shex-parse-to-XMLParser.yy
	bison -o shex-parse-to-XMLParser.cc shex-parse-to-XMLParser.yy

#/bin/sh ../scripts/ylwrap parser.yy y.tab.c parser.cc y.tab.h parser.h y.output parser.output -- bison -y

shex-parse-to-XMLScanner.cc: shex-parse-to-XMLScanner.ll shex-parse-to-XMLParser.hh
	flex -s -o shex-parse-to-XMLScanner.cc shex-parse-to-XMLScanner.ll

#/bin/sh ../scripts/ylwrap scanner.ll lex.yy.c scanner.cc -- flex  -olex.yy.c

shex-parse-to-XMLParser.o: shex-parse-to-XMLParser.cc shex-parse-to-XMLParser.hh shex-parse-to-XMLScanner.hh
	$(GPP)  -o shex-parse-to-XMLParser.o shex-parse-to-XMLParser.cc

shex-parse-to-XMLScanner.o: shex-parse-to-XMLScanner.cc shex-parse-to-XMLScanner.hh
	$(GPP)  -o shex-parse-to-XMLScanner.o shex-parse-to-XMLScanner.cc

#libshex-parse-to-XML.a: shex-parse-to-XMLParser.o shex-parse-to-XMLScanner.o
#	ar cru libshex-parse-to-XML.a shex-parse-to-XMLParser.o shex-parse-to-XMLScanner.o
#	ranlib libshex-parse-to-XML.a

#shex-parse-to-XMLTest.o: shex-parse-to-XMLTest.cc
#	$(GPP)  -o shex-parse-to-XMLTest.o shex-parse-to-XMLTest.cc

#shex-parse-to-XMLTest: shex-parse-to-XMLTest.o libshex-parse-to-XML.a
#	$(LINK) shex-parse-to-XMLTest shex-parse-to-XMLTest.o libshex-parse-to-XML.a $(LIBS)

shex-parse-to-XML: shex-parse-to-XMLParser.o shex-parse-to-XMLScanner.o
	$(LINK) shex-parse-to-XML shex-parse-to-XMLParser.o shex-parse-to-XMLScanner.o $(LIBS)

test: shex-parse-to-XML
	./shex-parse-to-XML shex-parse-to-XML.txt

valgrind: shex-parse-to-XML
	valgrind --leak-check=yes ./shex-parse-to-XML shex-parse-to-XML.txt

clean:
	rm -f shex-parse-to-XML libshex-parse-to-XML.a shex-parse-to-XMLParser.o shex-parse-to-XMLScanner.o shex-parse-to-XMLParser.cc shex-parse-to-XMLParser.hh shex-parse-to-XMLScanner.cc location.hh position.hh stack.hh

