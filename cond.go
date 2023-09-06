package antlraci

import (
	"github.com/JesseCoretta/go-stackage"
)

/*
cond.go contains condition related functions and methods, as
well as RuleExpression functionality.
*/

var OperatorKeywordMap map[string][]string
var comparisonOperatorMap map[string]stackage.ComparisonOperator

/*
RuleExpression contains the value(s) in a Rule condition expression.

In the case of multiple values, such as when parsing  'userdn' or
'extop' conditions, Style represents the following quotation schemes:

- 0 indicates "value" || "value" || "value" ...

- 1 indicates "value || value || value ..."

This has no meaning for single-valued expressions.
*/
type RuleExpression struct {
	Style  int
	Values []string
}

/*
Len returns the integer length of the receiver instance.
*/
func (r RuleExpression) Len() int {
	return len(r.Values)
}

/*
String is a stringer method that returns the string
representation of the receiver instance. Note that
the underlying quotation "Style" shall be honored
in the output.
*/
func (r RuleExpression) String() string {
	if r.Len() == 0 {
		return ``
	}

	switch r.Style {
	case 0:
		var vals []string
		for i := 0; i < r.Len(); i++ {
			vals = append(vals, sprintf("\"%s\"", r.Values[i]))
		}
		return join(vals, ` || `)
	case 1:
		return sprintf("\"%s\"", join(r.Values, ` || `))
	}

	return ``
}

func processRuleKeyword(trk any, id string) (kw string, err error) {

	switch tv := trk.(type) {

	case *TargetKeywordContext:
		if tv != nil {
			kw = tv.GetText()
		}

	case *BindKeywordContext:
		if tv != nil {
			kw = tv.GetText()
		}
	}

	if hasPfx(kw, `<missing`) {
		err = errorf("Zero-length or unreadable %s keyword, cannot process", id)
		return
	}

	if ctx, known := idKW(kw); !known {
		if kw == `(` {
			// target rules won't match this condition, since the
			// opening parenthesis is part of their rule structure.
			err = errorf("Parenthetical bind rule: only bind rules (plural) should be parenthetical")
		} else {
			err = errorf("Unresolvable keyword [%s]", kw)
		}
		kw = `<invalid_target_keyword>`
	} else if ctx != id {
		err = errorf("Inappropriate keyword context (%s) for type '%s'", ctx, id)
		kw = `<invalid_target_keyword>`
	}

	return
}

func processRuleOperator(btro any, kw string) (cop stackage.ComparisonOperator, err error) {
	var typ string
	var op string = `<unidentified_keyword>`

	switch tv := btro.(type) {
	case *BindOperatorContext:
		typ = `bind`
		if tv != nil {
			if keywordAllowsOperator(kw, tv.GetText()) {
				op = tv.GetText()
			} else {
				err = errorf("Illegal %s keyword/operator combination ('%s' -X-> '%s')",
					typ, kw, op)
				return
			}
		}

	case *TargetOperatorContext:
		typ = `target`
		if tv != nil {
			if keywordAllowsOperator(kw, tv.GetText()) {
				op = tv.GetText()
			} else {
				err = errorf("Illegal %s keyword/operator combination ('%s' -X-> '%s')",
					typ, kw, op)
				return
			}
		}
	}

	var found bool
	if cop, found = matchComparisonOperator(op); !found {
		err = errorf("Unresolvable comparison operator [%s], or not allowed for use with specified %s rule keyword [%s]",
			op, typ, kw)
	}

	return
}

/*
processRuleExpression returns slices of (previously) double-quoted expression (string) values,
along with a boolean value indicative of parsing success, which is qualified by at least one
(1) value being found.

This function is called during the parsing of Target/Rule conditions.
*/
func processRuleExpression(expr *ExpressionValuesContext) (r RuleExpression, err error) {
	if expr == nil {
		err = errorf("%T instance is nil; cannot proceed", expr)
		return
	}
	raw := expr.GetText()
	if rune(raw[0]) != rune(34) {
		err = errorf("%T instance cannot parse; unquoted value '%v'", r, raw)
		return
	}

	// determine the style of quotation, assuming
	// multi-values are present
	if strct(raw, `"`) == 2 {
		r.Style = 1
	}

	var uqraw string
	for i := 0; i < len(raw); i++ {
		if c := rune(raw[i]); c != '"' {
			uqraw += string(c)
		}
	}

	r.Values = split(uqraw, `||`) // split on, and then discard, symbolic OR
	if r.Len() == 0 {
		err = errorf("No expression values were parsed")
		return
	}

	return
}

func keywordAllowsOperator(kw, op string) bool {
	operators, foundkeyword := OperatorKeywordMap[kw]
	if !foundkeyword {
		return false
	}

	return strInSlice(op, operators)
}

func knownKeyword(kw string) (known bool) {
	_, known = OperatorKeywordMap[kw]
	return known
}

func idKW(kw string) (id string, known bool) {
	id = `<unidentified_keyword_context>`
	if known = knownKeyword(kw); !known {
		return
	}

	if hasPfx(kw, `targ`) || kw == `extop` {
		id = `target`
	} else {
		id = `bind`
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

func init() {
	OperatorKeywordMap = map[string][]string{
		// target rule keyword/operators
		`target`:        {`=`, `!=`},
		`target_to`:     {`=`, `!=`},
		`target_from`:   {`=`, `!=`},
		`targetcontrol`: {`=`, `!=`},
		`extop`:         {`=`, `!=`},
		`targetattr`:    {`=`, `!=`},
		`targetfilter`:  {`=`, `!=`},

		// these two target rule keywords disallow negation
		`targetscope`:     {`=`},
		`targattrfilters`: {`=`},

		// bind rule keyword/operators
		`userdn`:     {`=`, `!=`},
		`groupdn`:    {`=`, `!=`},
		`roledn`:     {`=`, `!=`},
		`userattr`:   {`=`, `!=`},
		`groupattr`:  {`=`, `!=`},
		`authmethod`: {`=`, `!=`},
		`dayofweek`:  {`=`, `!=`},
		`dns`:        {`=`, `!=`},
		`ip`:         {`=`, `!=`},

		// these are the only two keywords (of ANY rule) that allow all operators.
		`timeofday`: {`=`, `!=`, `>`, `<`, `>=`, `<=`},
		`ssf`:       {`=`, `!=`, `>`, `<`, `>=`, `<=`},
	}

	comparisonOperatorMap = map[string]stackage.ComparisonOperator{
		`=`:  stackage.Eq,
		`!=`: stackage.Ne,
		`<`:  stackage.Lt,
		`>`:  stackage.Gt,
		`<=`: stackage.Le,
		`>=`: stackage.Ge,
	}
}
