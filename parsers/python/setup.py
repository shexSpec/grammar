import sys
try:
    from setuptools import setup
except ImportError:
    from distutils.core import setup

# typing library was introduced as a core module in version 3.5.0
# NOTE: the antlr4-python3-runtime must be the same version that was used to create the parser library
with open('requirements.txt') as reqs:
    requires = [l for l in reqs.readline()]
if sys.version_info < (3, 5):
    requires.append("typing")

setup(
    name='PyShExC',
    version='0.3.0',
    packages=['pyshexc.parser', 'pyshexc.parser_impl'],
    url="http://github.com/shexSpec/grammar/parsers/python",
    license='Apache 2.0',
    author='Harold Solbrig',
    author_email='solbrig.harold@mayo.edu',
    description='"PyShExC - a Python ShExC parser',
    long_description='PyShExC - a ShExC to PyJSG, ShExJ and ShExR parser',
    install_requires=requires,
    test_requires=['dict_compare'],
    scripts=['scripts/shexc_to_shexj'],
    classifiers=[
        'Development Status :: 3 - Alpha',
        'Environment :: Console',
        'Intended Audience :: Developers',
        'Topic :: Software Development :: Compilers',
        'Programming Language :: Python :: 3 :: Only'],
)
