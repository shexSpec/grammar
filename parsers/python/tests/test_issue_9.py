import os
import unittest

from pyshexc.parser_impl.generate_shexj import generate


class Issue9TestCase(unittest.TestCase):

    """ Test that parser can process a URI """
    def test_fhir_uri(self):
        success = generate('http://build.fhir.org/observation.shex -nr -nj')
        self.assertTrue(success)

    def test_fhir_local_file(self):
        """ Test the schema that has UTF8 BOM """
        datafile = os.path.abspath(os.path.join(os.path.dirname(__file__), 'data', 'observation.shex'))
        success = generate([datafile, '-nr', '-nj'])
        self.assertTrue(success)

    def test_shextest_uri(self):
        success = generate('https://raw.githubusercontent.com/shexSpec/shexTest/master/schemas/1dotNS2.shex -nr -nj')
        self.assertTrue(success)


if __name__ == '__main__':
    unittest.main()
