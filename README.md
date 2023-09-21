# go-antlraci

[![Go Reference](https://pkg.go.dev/badge/github.com/JesseCoretta/go-antlraci.svg)](https://pkg.go.dev/github.com/JesseCoretta/go-antlraci) [![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](https://github.com/JesseCoretta/go-antlraci/blob/main/LICENSE) [![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/JesseCoretta/go-antlraci/go.yml)](https://github.com/JesseCoretta/go-antlraci/actions/workflows/go.yml) [![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/JesseCoretta/go-antlraci)](https://github.com/JesseCoretta/go-antlraci) [![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/JesseCoretta/go-antlraci/issues) [![Experimental](https://img.shields.io/badge/experimental-blue?logoColor=blue&label=%F0%9F%A7%AA%20%F0%9F%94%AC&labelColor=blue&color=gray)](https://github.com/JesseCoretta/JesseCoretta/blob/main/EXPERIMENTAL.md)



Package antlraci contains an ANTLR4-generated codebase meant to facilitate the parsing of ACIv3 definitions. The codebase is created through use of ANTLR4 parser/lexer grammar rules that honor the complete third (3rd) version of the Access Control Instruction syntax -- a form of permission expressions favored and supported by multiple X.500/LDAP server implementations on the market today.

## Advisory

Most of the Go codebase within this repository is **AUTO-GENERATED** through ANTLR4 routines. Only a small portion was actually written, and serves only to extend more convenient functions to an importing application. All extensions to the codebase are kept in separate files.

At NO POINT is any portion of the **AUTO-GENERATED** ANTLR4 codebase edited for any reason whatsoever.

## Status of this repository

This package is imported and used by its sister package, [`go-aci`](https://github.com/JesseCoretta/go-aci), for all parsing related functionality.

## Purpose of this repository

As there are some negative views in the community regarding the quality and efficiency of the Go code that ANTLR4 generates at this time, this `antlraci` package was created to keep cyclomatic penalties wholly separate from its sister package [`go-aci`](https://github.com/JesseCoretta/go-aci).

## ANTLR4/ACI Codebase Generation Procedure

For those interested, the following command was used to generate this codebase (not counting the custom `.go` files **added** manually):

```
$ antlr4 -Dlanguage=Go -package antlraci ACI*.g4
```

Two (2) grammar files -- [`ACIParser.g4`](https://github.com/JesseCoretta/go-antlraci/blob/main/_grammar/ACIParser.g4) and [`ACILexer.g4`](https://github.com/JesseCoretta/go-antlraci/blob/main/_grammar/ACILexer.g4) -- must be present within the directory in which the above `bash` command is executed. They have been included in this repository for reference and, optionally, use elsewhere.

For information on setting up ANTLR4 on your system, see the following resources:

 - [Official ANTLR Website](http://www.antlr.org)
 - The Definitive ANTLR 4 Reference, 2nd Edition (Book)

## Parser Helpers

Some package-level functions have been added to simplify the parsing process for individual components of an ACIv3 instruction, as well as an instruction as a whole.

 - `ParseTargetRule`
 - `ParseTargetRules`
 - `ParseBindRule`
 - `ParseBindRules`
 - `ParsePermissionBindRule`
 - `ParsePermissionBindRules`
 - `ParseInstruction`

Click the reference badge at the top of this file to learn more about this package.
