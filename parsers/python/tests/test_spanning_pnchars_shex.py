import os
import unittest

from pyshexc.parser_impl.generate_shexj import parse


class SpanningPnTestCase(unittest.TestCase):
    def test_shex(self):
        test_file_path = os.path.abspath(os.path.join(os.path.dirname(__file__),
                                                      '..',  # python
                                                      '..',  # parsers
                                                      '..',  # grammar
                                                      '..',  # shexSpec
                                                      'shexTest',
                                                      'schemas'
                                                      ))
        with open(os.path.join(test_file_path, '1refbnode_with_spanning_PN_CHARS_BASE1.shex'), 'rb') as f:
            # shex = [c if ord(c) < 65536 else c.encode().decode('utf-16') for c in f.read().decode()]
            shex = f.read().decode()
            parse(shex)
        self.assertTrue(True, "Parsing was valid")


if __name__ == '__main__':
    unittest.main()
