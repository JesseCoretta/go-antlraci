# go-antlraci

[![GoDoc](https://godoc.org/github.com/JesseCoretta/go-antlraci?status.svg)](https://godoc.org/github.com/JesseCoretta/go-antlraci) ![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)

Package antlraci contains an ANTLR4-generated codebase meant to facilitate the parsing of ACIv3 definitions. The codebase is created through use of ANTLR4 parser/lexer grammar rules that honor the complete third (3rd) version of the Access Control Instruction syntax -- a form of permission expressions favored and supported by multiple X.500/LDAP server implementations on the market today.

## Advisory

Most of the Go codebase within this repository is **AUTO-GENERATED** through ANTLR4 routines. Only a small portion was actually written, and serves only to extend more convenient functions to an importing application. All extensions to the codebase are kept in separate files.

At NO POINT is any portion of the **AUTO-GENERATED** ANTLR4 codebase edited for any reason whatsoever.

## Purpose of this repository

As there are some negative views in the community regarding the quality and efficiency of the Go code that ANTLR4 generates at this time, this `antlraci` package was created to keep cyclomatic penalties wholly separate from the importing [`go-aci`](https://github.com/JesseCoretta/go-aci) package, as opposed to direct integration or even through `internal` import.

## ANTLR4/ACI Codebase Generation Procedure

For those interested, the following command was used to generate this codebase (not counting the custom `.go` files **added** manually):

```
$ antlr4 -Dlanguage=Go -package antlraci ACI*.g4
```

Two (2) grammar files -- `ACIParser.g4` and `ACILexer.g4` -- must be present within the directory in which the above `bash` command is executed. They have been included in this repository for reference and, optionally, use elsewhere.

For information on setting up ANTLR4 on your system, see the following resources:

 - [Official ANTLR Website](http://www.antlr.org)
 - The Definitive ANTLR 4 Reference, 2nd Edition (Book)
