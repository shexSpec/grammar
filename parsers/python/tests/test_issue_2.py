# Copyright (c) 2018, Mayo Clinic
# All rights reserved.
#
# Redistribution and use in source and binary forms, with or without modification,
# are permitted provided that the following conditions are met:
#
# Redistributions of source code must retain the above copyright notice, this
#     list of conditions and the following disclaimer.
#
#     Redistributions in binary form must reproduce the above copyright notice,
#     this list of conditions and the following disclaimer in the documentation
#     and/or other materials provided with the distribution.
#
#     Neither the name of the Mayo Clinic nor the names of its contributors
#     may be used to endorse or promote products derived from this software
#     without specific prior written permission.
#
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
# ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
# WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
# IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
# INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
# BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, 
# DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
# LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
# OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED
# OF THE POSSIBILITY OF SUCH DAMAGE.

import unittest
from rdflib import Namespace

from ShExJSG import ShExJ

from pyshexc.parser_impl.generate_shexj import parse

BASE = Namespace("https://raw.githubusercontent.com/shexSpec/shexTest/master/validation/")
FOO = Namespace("/some/location/file/")
EX = Namespace("http://example.org/")


class TestIssue2(unittest.TestCase):
    def test_no_base(self):
        shex_str = '<S1> {<p1> [<o1>]}'
        shex: ShExJ.Schema = parse(shex_str)
        self.assertEqual("S1", str(shex.shapes[0].id))
        self.assertEqual("p1", str(shex.shapes[0].expression.predicate))
        self.assertEqual("o1", str(shex.shapes[0].expression.valueExpr.values[0]))

    def test_default_base(self):
        shex_str = '<S1> {<p1> [<o1>]}'
        shex: ShExJ.Schema = parse(shex_str, str(BASE))
        self.assertEqual(str(BASE.S1),
                         str(shex.shapes[0].id))
        self.assertEqual(str(BASE.p1), str(shex.shapes[0].expression.predicate))
        self.assertEqual(str(BASE.o1), str(shex.shapes[0].expression.valueExpr.values[0]))

    def test_explicit_base(self):
        shex_str = f'BASE <{str(FOO)}>\n<S1> {{<p1> [<o1>]}}'
        shex: ShExJ.Schema = parse(shex_str, str(BASE))
        self.assertEqual(str(FOO.S1),
                         str(shex.shapes[0].id))
        self.assertEqual(str(FOO.p1), str(shex.shapes[0].expression.predicate))
        self.assertEqual(str(FOO.o1), str(shex.shapes[0].expression.valueExpr.values[0]))

    def test_explicit_uris(self):
        shex_str = f"""
BASE <{str(FOO)}>
PREFIX ex: <{EX}>

ex:S1 {{ex:p1 [ex:o1]}}"""
        shex: ShExJ.Schema = parse(shex_str, str(BASE))
        self.assertEqual(str(EX.S1),
                         str(shex.shapes[0].id))
        self.assertEqual(str(EX.p1), str(shex.shapes[0].expression.predicate))
        self.assertEqual(str(EX.o1), str(shex.shapes[0].expression.valueExpr.values[0]))


if __name__ == '__main__':
    unittest.main()
