package antlraci

/*
cond.go contains condition related functions and methods, as
well as RuleExpression functionality.
*/

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

func processRuleKeyword(trk any) (kw string, found bool) {
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

	found = len(kw) > 0
	return
}

func processRuleOperator(bro any) (op string, found bool) {
	switch tv := bro.(type) {
	case *BindOperatorContext:
		if tv != nil {
			op = tv.GetText()
		}
	case *TargetOperatorContext:
		if tv != nil {
			op = tv.GetText()
		}
	}

	found = len(op) > 0
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
