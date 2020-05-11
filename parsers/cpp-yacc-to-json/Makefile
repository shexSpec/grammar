# $Id: Makefile,v 1.5 2014-04-03 16:51:56 eric Exp $

# recipies:
#   normal build:
#     make shex-to-json
#   force the use of the tracing facilities (and redirect to stdout):
#     HAVE_BOOST=1 TRACE_FD=1 make -W shex-to-jsonParser.yy test
#   have valgrind start a debugger (works as M-x gdb invocation command):
#     valgrind --db-attach=yes --leak-check=yes shex-to-json shex-to-json.txt
#   same, if you aren't working in gdb:
#     HAVE_BOOST=1 TRACE_FD=1 make valgrind
#   debugging in emacs:
#     gdb --annotate=3 shex-to-json    (set args shex-to-json.txt)

ifdef HAVE_BOOST
  LIBS=-lboost_iostreams
  DEFS=-DHAVE_BOOST
else
  LIBS=
  DEFS=
endif

CXXSTD=-std=c++17

ifdef GCC
  CXX=g++ $(CXXSTD) -DYYTEXT_POINTER=1 $(DEFS) -W -Wall -Wextra -ansi -g -c
  LINK=g++ -W -Wall -Wextra -ansi -g -o
else # llvm (clang) , not using -stdlib=libc++
  CXX=clang $(CXXSTD) -DYYTEXT_POINTER=1 $(DEFS) -W -Wall -Wextra -Wno-c++11-extensions -Wno-c++17-extensions -ansi -g -c
  LINK=clang -L/usr/lib -lstdc++ -W -Wall -Wextra -g -o
endif

shex-to-jsonParser.cc shex-to-jsonParser.hh location.hh position.hh stack.hh: shex-to-jsonParser.yy Makefile
	bison -o shex-to-jsonParser.cc shex-to-jsonParser.yy

#/bin/sh ../scripts/ylwrap parser.yy y.tab.c parser.cc y.tab.h parser.h y.output parser.output -- bison -y

shex-to-jsonScanner.cc: shex-to-jsonScanner.ll shex-to-jsonParser.hh Makefile
	flex -s -o shex-to-jsonScanner.cc shex-to-jsonScanner.ll

#/bin/sh ../scripts/ylwrap scanner.ll lex.yy.c scanner.cc -- flex  -olex.yy.c

shex-to-jsonParser.o: shex-to-jsonParser.cc shex-to-jsonParser.hh shex-to-jsonScanner.hh shex.hh
	$(CXX) -o shex-to-jsonParser.o shex-to-jsonParser.cc

shex-to-jsonScanner.o: shex-to-jsonScanner.cc shex-to-jsonScanner.hh
	$(CXX) -o shex-to-jsonScanner.o shex-to-jsonScanner.cc

#libshex-to-json.a: shex-to-jsonParser.o shex-to-jsonScanner.o
#	ar cru libshex-to-json.a shex-to-jsonParser.o shex-to-jsonScanner.o
#	ranlib libshex-to-json.a

#shex-to-jsonTest.o: shex-to-jsonTest.cc
#	$(CXX)  -o shex-to-jsonTest.o shex-to-jsonTest.cc

#shex-to-jsonTest: shex-to-jsonTest.o libshex-to-json.a
#	$(LINK) shex-to-jsonTest shex-to-jsonTest.o libshex-to-json.a $(LIBS)

shex-to-json.o: shex-to-json.cc shex-to-jsonParser.hh shex.hh
	$(CXX) -o shex-to-json.o shex-to-json.cc

shex-to-json: shex-to-json.o shex-to-jsonParser.o shex-to-jsonScanner.o
	$(LINK) shex-to-json shex-to-json.o shex-to-jsonParser.o shex-to-jsonScanner.o $(LIBS)

test: shex-to-json
	./shex-to-json shex-to-json.txt

valgrind: shex-to-json
	valgrind --leak-check=yes ./shex-to-json shex-to-json.txt

clean:
	rm -f shex-to-json shex-to-json.o libshex-to-json.a shex-to-jsonParser.o shex-to-jsonScanner.o shex-to-jsonParser.cc shex-to-jsonParser.hh shex-to-jsonScanner.cc location.hh position.hh stack.hh

manual.o: manual.cc shex.hh
	$(CXX) -o manual.o manual.cc

manual: manual.o
	$(LINK) manual manual.o

.PHONY: manual-test
manual-test: manual
	manual

