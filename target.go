package antlraci

import (
	"github.com/JesseCoretta/go-stackage"
)

/*
ParseTargetRule parses a single target rule expression, e.g.: 'targetattr = "cn"',
and returns a stackage.Condition instance alongside an error.
*/
func ParseTargetRule(raw string) (r stackage.Condition, err error) {
	var p *ACIParser
	if p, err = initAntlr(raw); err != nil {
		return
	}

	// Parse and return our TargetRule instance
	return processTargetRule(p.TargetRule())
}

func processTargetRule(itrc ITargetRuleContext) (r stackage.Condition, err error) {
	if itrc == nil {
		err = errorf("%T instance is nil; cannot proceed", itrc)
		return
	}

	var (
		ct int = itrc.GetChildCount()
		kw string
		op string
		ex RuleExpression
	)

	for k := 0; k < ct; k++ {
		var ok bool

		switch tv := itrc.GetChild(k).(type) {

		case *OpeningParenthesisContext:
			if hasPfx(tv.GetText(), `<missing`) {
				err = errorf("Missing open parenthesis for target rule")
				return
			}

		case *ClosingParenthesisContext:
			if hasPfx(tv.GetText(), `<missing`) {
				err = errorf("Missing close parenthesis for target rule")
				return
			}

		case *TargetKeywordContext:
			if kw, ok = processRuleKeyword(tv); !ok {
				err = errorf("Failed to process %T for Target Rule keyword", tv)
				return
			}

		case *TargetOperatorContext:
			if op, ok = processRuleOperator(tv); !ok {
				err = errorf("Failed to process %T for Target Rule operator", tv)
				return
			}

		case *ExpressionValuesContext:
			if ex, err = processRuleExpression(tv); err != nil {
				return
			}
		}
	}

	cop, found := matchComparisonOperator(op)
	if !found {
		err = errorf("Unknown comparison operator '%s'", op)
		return
	}

	// assemble condition
	r = stackage.Cond(kw, cop, ex).
		Paren(true).
		SetCategory(`target`)

	return
}

/*
ParseTargetRules returns slices of stackage.Condition instances
within a stackage.Stack alongside a boolean value indicative of
a non-zero length return result.

A maximum of nine (9) condition instances is allowed. No target
keyword may be used in more than one (1) condition inside any
given target-based stack.
*/
func ParseTargetRules(raw string) (r stackage.Stack, err error) {
	var p *ACIParser
	if p, err = initAntlr(raw); err != nil {
		return
	}

	return processTargetRules(p.TargetRules())
}

/*
processTargetRule processes the tokens found within a target rule.
*/
func processTargetRules(itrc ITargetRulesContext) (T stackage.Stack, err error) {
	if itrc == nil {
		err = errorf("%T instance is nil; cannot proceed", itrc)
		return
	}

	ct := itrc.GetChildCount()
	T = stackage.List(9) // Cap of nine (9) in-force

	for k := 0; k < ct; k++ {
		switch tv := itrc.GetChild(k).(type) {

		// A TargetRuleContext instance describes a
		// single target rule expression [kw op expr]
		case *TargetRuleContext:
			var t stackage.Condition
			if t, err = processTargetRule(tv); err != nil {
				return
			}

			// in addition to being limited to nine (9)
			// total target rule conditions, we must be
			// certain each keyword is only used in ONE
			// condition. There are a total of nine (9)
			// possible keywords. You get the idea.
			//
			// TODO: use stackage's PushPolicy for this
			// functionality.
			if targetRuleUnique(T, t) {
				T.Push(t)
			} else {
				err = errorf("%T uniqueness violation (%s used more than once)", t.Keyword)
				return
			}
		}
	}

	return
}

/*
targetRuleUnique will scan the current slices within stackage.Stack T
input instance and look for any preexisting Condition that bears the
same keyword as stackage.Condition t input instance. If a match is
found, then a false 'unique' return is sent, thereby preventing the
push of a duplicate condition keyword. A return of true indicates the
push attempt can proceed.
*/
func targetRuleUnique(T stackage.Stack, t stackage.Condition) (unique bool) {
	for i := 0; i < T.Len(); i++ {
		slice, _ := T.Index(i)
		if assert, ok := slice.(stackage.Condition); ok {
			if t.Keyword() == assert.Keyword() {
				return
			}
		}
	}

	unique = true
	return
}
