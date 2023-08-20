package antlraci

import (
	"errors"
	"fmt"
	"strings"

	"github.com/JesseCoretta/go-stackage"
	"github.com/antlr4-go/antlr/v4"
)

var (
	printf  func(string, ...any) (int, error) = fmt.Printf
	sprintf func(string, ...any) string       = fmt.Sprintf
	split   func(string, string) []string     = strings.Split
	join    func([]string, string) string     = strings.Join
	lc      func(string) string               = strings.ToLower
	eq      func(string, string) bool         = strings.EqualFold
	hasPfx  func(string, string) bool         = strings.HasPrefix
	hasSfx  func(string, string) bool         = strings.HasSuffix
	strct   func(string, string) int          = strings.Count
	trimS   func(string) string               = strings.TrimSpace
)

func initAntlr(raw string) (p *ACIParser, err error) {
	// don't waste time on a zero string ...
	if len(raw) == 0 {
		err = errorf("Zero-length input string; impossible parsing request")
		return
	}

	// create a new input stream using the raw input.
	is := antlr.NewInputStream(raw)

	// prepare lexer
	lx := NewACILexer(is)

	// Prepare the parser using a new CommonTokenStream
	// as input, define as p, and verified for nilness.
	if p = NewACIParser(antlr.NewCommonTokenStream(lx, 0)); p == nil {
		err = errorf("Unknown ANTLR error; received a nil %T instance", p)
		return
	}

	// Declare our intent build a parse tree
	p.BuildParseTrees = true
	return
}

/*
errorf wraps errors.New and returns a non-nil instance of error
based upon a non-nil/non-zero msg input value with optional args.
*/
func errorf(msg any, x ...any) error {
	switch tv := msg.(type) {
	case string:
		if len(tv) > 0 {
			return errors.New(sprintf(tv, x...))
		}
	case error:
		if tv != nil {
			return errors.New(sprintf(tv.Error(), x...))
		}
	}

	return nil
}

func andBindRules() stackage.Stack  { return stackage.And() }
func orBindRules() stackage.Stack   { return stackage.Or() }
func notBindRules() stackage.Stack  { return stackage.Not() }
func targetRules() stackage.Stack   { return stackage.List(9) }
func stackageList() stackage.Stack  { return stackage.List() }
func stackageBasic() stackage.Stack { return stackage.Basic() }

func newCondition() (r stackage.Condition) {
	return
}

func newStack() (r stackage.Stack) {
	return
}

/*
stackByBoolOp is used exclusively in Bind Rules parsing, and
will return a (potentially populated) logical Boolean WORD
based stack (e.g.: AND) within the st variable, alongside a
Boolean true/false value (identified).

The Boolean return value indicates whether the input string
value (op) successfully mapped to a known Boolean WORD-based
stackage.Stack configuration. In such a case, true is returned.
A return value of false indicates no WORD operator was matched,
and that a fallback to a List-style stackage.Stack has occurred.
This should not necessarily be perceived as an error condition.

The contents variadic input argument, which is wholly optional,
will facilitate the addition (by stackage.Stack.Push) of one (1)
or more values -- of any type -- into the return st instance once
it has been initialized. Any added contents arguments are honored
regardless of fallback state.
*/
func stackByBoolOp(op string, contents ...any) (st stackage.Stack, identified bool) {

	switch oper := lc(op); oper {

	// match the Boolean AND logical
	// WORD operator ...
	case `and`:
		identified = true
		st = stackage.And()
		for i := 0; i < len(contents); i++ {
			st.Push(contents[i])
		}
		return

	// match the Boolean OR logical
	// WORD operator ...
	case `or`:
		identified = true
		st = stackage.Or()
		for i := 0; i < len(contents); i++ {
			st.Push(contents[i])
		}
		return

	// match the Boolean NOT (AND NOT)
	// logical WORD operator ...
	case `and not`:
		identified = true
		st = stackage.Not()
		for i := 0; i < len(contents); i++ {
			st.Push(contents[i])
		}
		return
	}

	st = andBindRules()
	for i := 0; i < len(contents); i++ {
		st.Push(contents[i])
	}

	return
}

func printStack(x stackage.Stack, parent int) {
	for i := 0; i < x.Len(); i++ {
		slice, _ := x.Index(i)
		switch tv := slice.(type) {
		case stackage.Stack:
			k := trimS(tv.Kind())
			printf("%T[%d:%d] :: [%d;%s]: %s\n", tv, parent, i, tv.Len(), k, tv)
			printStack(tv, i)
		case stackage.Condition:
			printf("%T[%d:%d] :: %s\n", tv, parent, i, tv)
		}
	}
}

func disenvelopSingleBindRule(tree antlr.Tree) (antlr.Tree, bool) {
	if tree.GetChildCount() == 1 {
		switch tv := tree.GetChild(0).(type) {
		case *BindRuleContext:
			return tv, true
		}
	}
	return tree, false
}

func disenvelopStack(outer stackage.Stack) (inner stackage.Stack, ok bool) {
	if outer.Len() != 1 {
		inner = outer
		return
	}

	slice, _ := outer.Index(0)
	switch tv := slice.(type) {
	case stackage.Stack:
		//if eq(outer.Kind(), tv.Kind()) {
		ok = true
		inner = tv
		//}
	default:
		inner = outer
	}

	return
}

func disenvelopCondition(outer stackage.Stack) (inner stackage.Condition, ok bool) {
	if outer.Len() != 1 {
		return
	}

	slice, _ := outer.Index(0)
	switch tv := slice.(type) {
	case stackage.Condition:
		inner = tv
		ok = true
	}

	return
}

func matchComparisonOperator(x string) (op stackage.ComparisonOperator, ok bool) {
	var k string
	for k, op = range comparisonOperatorMap {
		if x == k {
			ok = true
			return
		}
	}

	return
}

var comparisonOperatorMap map[string]stackage.ComparisonOperator

func init() {
	comparisonOperatorMap = map[string]stackage.ComparisonOperator{
		`=`:  stackage.Eq,
		`!=`: stackage.Ne,
		`<`:  stackage.Lt,
		`>`:  stackage.Gt,
		`<=`: stackage.Le,
		`>=`: stackage.Ge,
	}
}
