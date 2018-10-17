import sys
try:
    from setuptools import setup
except ImportError:
    from distutils.core import setup

# typing library was introduced as a core module in version 3.5.0
# NOTE: the antlr4-python3-runtime must be the same version that was used to create the parser library
requires = ['antlr4-python3-runtime>=4.7.1',
            'jsonasobj>=1.2.1',
            'ShExJSG>=0.5.1',
            'requests>=2.18',
            'rdflib>=4.2.2',
            'rdflib-jsonld>=0.4.0',
            'PyJSG>=0.9.0']

setup(
    name='PyShExC',
    version='0.5.1',
    packages=['pyshexc.parser', 'pyshexc.parser_impl'],
    url="http://github.com/shexSpec/grammar/parsers/python",
    license='Apache 2.0',
    author='Harold Solbrig',
    author_email='solbrig.harold@mayo.edu',
    description='"PyShExC - a Python ShExC parser',
    long_description='PyShExC - a ShExC to PyJSG, ShExJ and ShExR parser',
    install_requires=requires,
    tests_require=['yadict-compare>=1.2.0'],
    scripts=['scripts/shexc_to_shexj'],
    classifiers=[
        'Development Status :: 4 - Beta',
        'Environment :: Console',
        'Intended Audience :: Developers',
        'Topic :: Software Development :: Compilers',
        'Programming Language :: Python :: 3 :: Only',
        'Programming Language :: Python :: 3.6',
        'Programming Language :: Python :: 3.7'],
)
