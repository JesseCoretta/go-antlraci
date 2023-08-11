package antlraci

import (
	"fmt"
	"strings"

        "github.com/antlr4-go/antlr/v4"
)

var (
	printf func(string, ...any) (int, error) = fmt.Printf
	split func(string, string) []string      = strings.Split
)

/*
aciListener is a private type used to facilitate the
parsing of an ACIv3 instruction expression through ANTLR4.
*/
type aciListener struct {
	tempList []string
	*BaseACIParserListener
}

/*
VisitTerminal is the listener component that bridges the internal
ANTLR4 framework with the go-aci package. It is not intended for
use by users, and exists only to satisfy the particular ANTLR4
interface signature requirements needed for a Listener.

The main purpose of this particular function, for us, is to build
the list of tokens (tempList) which we can read outside the ANTLR
framework as needed.
*/
func (r *aciListener) VisitTerminal(node antlr.TerminalNode) {
	value := node.GetSymbol()
	if value.GetTokenType() == antlr.TokenEOF {
		return
	}

        r.tempList = append(r.tempList, value.GetText())
}

/*
InstructionToTokens parses the input value (instruct) into ANTLR
tokens ([]string). An instance of []string -- hopefully containing
all the tokenized components of the input -- is returned.
*/
func InstructionToTokens(instruct string) (tokens []string) {
        // Create input stream using instruct
        // input string.
        is := antlr.NewInputStream(instruct)

        // prepare lexer and a token streamer
        lx := NewACILexer(is)
        ts := antlr.NewCommonTokenStream(lx, 0)

        // prepare the parser and declare our
        // intent build a parse tree
        p := NewACIParser(ts)
        p.BuildParseTrees = true
	tree := p.Parse()

        // Initialize our instruction listener type and
        // initialize a temporary []string instance for
	// storage of said tokens.
        listen := new(aciListener)
        listen.tempList = make([]string,0)

        // Trigger the antlr parse tree walker which
	// will build said temporary list for the user
	// to traverse.
        walker := antlr.NewParseTreeWalker()
        walker.Walk(listen, tree)

        return listen.tempList
}

/*
processTargetRule processes the tokens found within a target rule.
*/
func processTargetRule(itrc ITargetRuleContext) {
        ct := itrc.GetChildCount()
        encapParenthesis(itrc)
        for k := 0; k < ct; k++ {
                switch tv := itrc.GetChild(k).(type) {

                case *OpeningParenthesisContext,
                        IOpeningParenthesisContext,
                        IClosingParenthesisContext,
                        *ClosingParenthesisContext:
                        encapParenthesis(tv)

                case *TargetRuleContext:
                        processTargetRule(tv)

                case *TargetKeywordContext,
                        ITargetKeywordContext:
                        kw, _ := processTargetRuleKeyword(tv)
                        printf("%s\n", kw)

                case *TargetOperatorContext,
                        ITargetOperatorContext:
                        op, _ := processTargetRuleOperator(tv)
                        printf("%s\n", op)

                case IExpressionValuesContext:
                        processRuleExpression(tv)

                }
        }
        encapParenthesis(itrc)
}

/*
processBindRule processes the tokens found within a bind rule.
*/
func processBindRule(ibrc IBindRuleContext) {
        ct := ibrc.GetChildCount()
	encapParenthesis(ibrc)
        for k := 0; k < ct; k++ {
                switch tv := ibrc.GetChild(k).(type) {

		case *BindRuleContext:
			processBindRule(tv)

		case *BindKeywordContext,
			IBindKeywordContext:
			kw, _ := processBindRuleKeyword(tv)
			printf("%s\n", kw)

		case *BindOperatorContext,
			IBindOperatorContext:
			op, _ := processBindRuleOperator(tv)
			printf("%s\n", op)

		case IExpressionValuesContext:
			processRuleExpression(tv)

		case *WordAndContext,
			*WordOrContext,
			*WordNotContext,
			IWordAndContext,
			IWordOrContext,
			IWordNotContext:
			processBooleanWord(tv)
		}
        }
	encapParenthesis(ibrc)
}

/*
func processInstructionAnchor(ins any) {
	case IAccessControlLabelContext:
		if assert, ok := tv.GetPayload().(*antlr.BaseParserRuleContext); ok {
			printf("%s\n", assert.GetText())
		}

	case *AccessControlLabelContext:
		raw = tv.GetText()
}
*/

/*
processRuleExpression returns slices of (previously) double-quoted expression
(string) values, along with a boolean value indicative of parsing success,
which is qualified by at least one (1) value being found.

This function is called during the parsing of Target/Rule conditions.
*/
func processRuleExpression(quo any) (bre []string, ok bool) {
	var raw string
	switch tv := quo.(type) {

	case IExpressionValuesContext:
		if assert, ok := tv.GetPayload().(*antlr.BaseParserRuleContext); ok {
			printf("%s\n", assert.GetText())
		}

	case *ExpressionValuesContext:
		raw = tv.GetText()

	default:
		return
	}

	var uqraw string
	for i := 0; i < len(raw); i++ {
		if c := rune(raw[i]); c != '"' {
			uqraw += string(c)
		}
	}

        bre = split(uqraw,`||`)
	ok = len(bre) > 0

	return
}

func processTargetRuleKeyword(trk any) (kw string, found bool) {
        switch tv := trk.(type) {
        case *TargetKeywordContext:
                kw = tv.GetText()
        case ITargetKeywordContext:
                kw = tv.GetText()
        }

        found = len(kw) > 0
        return
}

func processBindRuleKeyword(brk any) (kw string, found bool) {
	switch tv := brk.(type) {
	case *BindKeywordContext:
		kw = tv.GetText()
	case IBindKeywordContext:
		kw = tv.GetText()
	}

	found = len(kw) > 0
	return
}

func processBindRuleOperator(bro any) (op string, found bool) {
        switch tv := bro.(type) {
        case *BindOperatorContext:
                op = tv.GetText()
        case IBindOperatorContext:
                op = tv.GetText()
        }

	found = len(op) > 0
	return
}

func processTargetRuleOperator(tro any) (op string, found bool) {
        switch tv := tro.(type) {
        case *TargetOperatorContext:
                op = tv.GetText()
        case ITargetOperatorContext:
                op = tv.GetText()
        }

        found = len(op) > 0
        return
}

func processBindPermission(iperm IPermissionContext) {
	printf("%T :: %s\n", iperm, iperm.GetText())
}

func processBooleanAnd(word any) {
	switch tv := word.(type) {
	case *WordAndContext:
		printf("%s\n", tv.GetText())
	case IWordAndContext:
		printf("%s\n", tv.GetText())
	}
}

func processBooleanOr(word any) {
        switch tv := word.(type) {
        case *WordOrContext:
                printf("%s\n", tv.GetText())
        case IWordOrContext:
                printf("%s\n", tv.GetText())
        }
}

func processBooleanNot(word any) {
        switch tv := word.(type) {
        case *WordNotContext:
                printf("%s\n", tv.GetText())
        case IWordNotContext:
                printf("%s\n", tv.GetText())
        }
}

func processBooleanWord(word any) {
	switch tv := word.(type) {

	case *WordAndContext,
		IWordAndContext:
		processBooleanAnd(tv)

	case *WordOrContext,
		IWordOrContext:
		processBooleanOr(tv)

	case *WordNotContext,
		IWordNotContext:
		processBooleanNot(tv)
	}
}

func leftParenthesis(brc any) string {
        switch tv := brc.(type) {

	case *BindRulesContext:
                if paren := tv.OpeningParenthesis(); paren != nil {
                        return paren.GetText()
                }

	case *OpeningParenthesisContext:
		if tv != nil {
			return tv.GetText()
		}
        }

	return ``
}

func rightParenthesis(brc any) string {
	switch tv := brc.(type) {

	case IBindRulesContext:
		if paren := tv.ClosingParenthesis(); paren != nil {
			return paren.GetText()
		}

	case *BindRulesContext:
		if paren := tv.ClosingParenthesis(); paren != nil {
			return paren.GetText()
		}

        case *ClosingParenthesisContext:
                if tv != nil {
                        return tv.GetText()
                }

        case IClosingParenthesisContext:
                if tv != nil {
                        return tv.GetText()
                }
	}

	return ``
}

func encapParenthesis(brc any) string {
	switch tv := brc.(type) {

	case *OpeningParenthesisContext:
		return leftParenthesis(tv)

	case *ClosingParenthesisContext:
		return rightParenthesis(tv)
	}
	return ``
}

func processBindRules(ctx IBindRulesContext) {
	count := ctx.GetChildCount()
	encapParenthesis(ctx)
        for c := 0; c < count; c++ {

		// use the child indices to call the
		// slice from the antlr children array.
                child := ctx.GetChild(c)

		switch uv := child.(type) {

		// context type is a nested bind rules
		// (NOTE: plurality)
		case *BindRulesContext:
			processBindRules(uv)

		// context type is a nested bind rule,
		// (NOTE: singular component of a bind
		// rules context)
                case *BindRuleContext:
			processBindRule(uv)

		// context type(s) are left/opener or
		// right/closer parentheticals.
		case *OpeningParenthesisContext,
			*ClosingParenthesisContext:
			printf("%s\n", encapParenthesis(uv))

		// context type(s) are Boolean WORD
		// operators that "glue" bind rule(s)
		// together in a logical manner.
                case *WordOrContext,
			IWordOrContext,
			*WordAndContext,
			IWordAndContext,
			IWordNotContext,
			*WordNotContext:
			processBooleanWord(uv)

		default:
			printf("UNHANDLED BINDRULES CONTEXT: %T\n", uv)
		}
        }
	encapParenthesis(ctx)
}

func BindRuleToTokens(bindrule string) {
	is := antlr.NewInputStream(bindrule)

        // prepare lexer and a token streamer
        lx := NewACILexer(is)
        ts := antlr.NewCommonTokenStream(lx, 0)

        // prepare the parser and declare our
        // intent build a parse tree
        p := NewACIParser(ts)
        p.BuildParseTrees = true
	ipc := p.Parse()
	ins := ipc.Instruction()

	//processBindPermission(pbr.Permission())
	for _, pbr := range ins.AllPermissionBindRule() {
		brls := pbr.BindRules()
		leftParenthesis(brls)
		processBindRules(brls)
		rightParenthesis(brls)
	}
}

