package antlraci

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

var (
	printf  func(string, ...any) (int, error) = fmt.Printf
	sprintf func(string, ...any) string       = fmt.Sprintf
	split   func(string, string) []string     = strings.Split
	join    func([]string, string) string     = strings.Join
	hasPfx  func(string, string) bool         = strings.HasPrefix
	strct   func(string, string) int          = strings.Count
)

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

/*
RuleExpression contains the value(s) in a Rule condition expression.

In the case of multiple values, such as when parsing  'userdn' or
'extop' conditions, the Style field indicates the following:

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
representation of the receiver instance.
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

/*
Rule contains a keyword, an operator and one (1) or more
values.

Strictly speaking, (Target) Rule instances are MUST ALWAYS be
parenthetical. A sequence of Target Rule instances itself,
however, is NOT parenthetical.

However, (Bind) Rule (singular) instances are never actually
parenthetical per se. A Bind Rule "condition" -- that is, a
single keyword, operator and the expression value(s) -- that
*appears* encapsulated in parenthesis is actually a lone rule
within a parenthetical stack.
*/
type Rule struct {
	// the keyword shall define whether the rule is a
	// target or a bind.
	Keyword string

	// Note target rules ONLY use '=' and '!=', while
	// bind rules use '=', '!=', '>', '<', '>=', '<='.
	Operator string

	// len > 1 always implies delimitation through symbolic
	// OR (||), but only for keywords that allow it.
	Values RuleExpression
}

/*
BindRules is an alias for []any, and represents a simplistic "stack"
type used exclusively for bind rules, which may be nested in complex
fashion.

The structure is as follows:

- A BooleanWord instance indicates 'AND', 'OR' or 'AND NOT'

- A (nested) BindRules instance indicates the given index value will
descend into another stack

- A *Rule instance is a single bind rule condition: [ kw op val(s) ]

- A string value shall be a numbered OPEN_PARENTHESIS_<N> flag, or
a CLOSE_PARENTHESIS_<N> flag. Numbers are relative to the stack in
which they reside, and will appear multiple times in complex stack
structures.
*/
type BindRules []any

/*
TargetRules is an alias for slices of Rule instances ([]*Rule), and
contains up to nine (9) keyword-unique target rules.
*/
type TargetRules []*Rule

/*
Permission defines the rights and disposition of a single Permission
+ Bind Rule pair.
*/
type Permission struct {
	Allow  bool
	Rights []string
}

/*
String is a stringer method that returns the string
representation of the receiver instance.
*/
func (r Permission) String() string {
	if len(r.Rights) == 0 {
		return ``
	}

	switch r.Allow {
	case true:
		return sprintf("allow(%s)", join(r.Rights, `,`))
	}

	return sprintf("deny(%s)", join(r.Rights, `,`))
}

/*
PermissionBindRule contains a single Permission instance, and
a single BindRules instance.
*/
type PermissionBindRule struct {
	Permission
	BindRules
}

/*
String is a stringer method that returns the string
representation of the receiver instance.
*/
func (r PermissionBindRule) String() string {
	if len(r.Permission.Rights) == 0 {
		return ``
	}
	if r.BindRules.Len() == 0 {
		return ``
	}

	return sprintf("%s %s;",
		r.Permission.String(),
		r.BindRules.String())
}

/*
PermissionBindRules is an alias type for slices of PermissionBindRule
instances.
*/
type PermissionBindRules []PermissionBindRule

/*
Len returns the integer length of the receiver instance.
*/
func (r PermissionBindRules) Len() int {
	return len(r)
}

/*
String is a stringer method that returns the string
representation of the receiver instance.
*/
func (r PermissionBindRules) String() string {
	var str []string
	for i := 0; i < r.Len(); i++ {
		str = append(str, r[i].String())
	}

	return join(str, ` `)
}

/*
AccessControlLabel is string alias type used to uniquely
identify a specific Instruction instance within an X.500
DIT.
*/
type AccessControlLabel string

/*
String is a stringer method that returns the string
representation of the receiver instance.
*/
func (r AccessControlLabel) String() string {
	return string(r)
}

/*
BooleanWord describes one (1) of three (3) possible
logical Boolean WORD operators. See the BooleanWord
constants for the specific meaning of values.
*/
type BooleanWord int

const (
	And BooleanWord = iota
	Or
	Not
)

/*
Instruction is a convenience type returned when ParseInstruction
is successfully executed. It will contain all of the central elements
of an Access Control Instruction (ACI) Version 3.0 expression:

- Zero (0) or more Target Rules

- An Access Control Label (unique, descriptive string identifier for an ACI)

- One (1) or more Permission+Bind Rules, each containing one (1)
Permission and one (1) Bind Rules expression.
*/
type Instruction struct {
	TargetRules         // optional
	AccessControlLabel  // required
	PermissionBindRules // >= 1 required
}

/*
String is a stringer method that returns the string
representation of the receiver instance.
*/
func (r Instruction) String() string {
	if r.PermissionBindRules.Len() == 0 {
		return ``
	}

	if len(r.AccessControlLabel) == 0 {
		return ``
	}

	var str string
	if r.TargetRules.Len() > 0 {
		str += r.TargetRules.String()
	}

	str += sprintf("(version 3.0; acl \"%s\"; ", r.AccessControlLabel.String())
	str += r.PermissionBindRules.String()
	str += ` )`

	return str
}

func processTargetRule(itrc ITargetRuleContext) (T *Rule, err error) {
	if itrc == nil {
		err = errorf("%T instance is nil; cannot proceed", itrc)
	}

	ct := itrc.GetChildCount()
	T = new(Rule)

	for k := 0; k < ct; k++ {
		switch tv := itrc.GetChild(k).(type) {

		case *TargetKeywordContext:
			if kw, ok := processRuleKeyword(tv); ok {
				T.Keyword = kw
			}

		case *TargetOperatorContext:
			if op, ok := processRuleOperator(tv); ok {
				T.Operator = op
			}

		case *ExpressionValuesContext:
			if T.Values, err = processRuleExpression(tv); err != nil {
				return
			}
		}
	}

	if !ruleReady(T) {
		err = errorf("%T is missing one of 'Keyword', 'Operator' or 'Values'; cannot proceed", T)
	}

	return
}

/*
Len returns the integer length of the receiver instance.
*/
func (r TargetRules) Len() int {
	return len(r)
}

/*
String is a stringer method that returns the string
representation of the receiver instance.
*/
func (r TargetRules) String() string {
	var str []string
	for i := 0; i < r.Len(); i++ {
		str = append(str, sprintf("( %s )", r[i]))
	}
	return join(str, ``)
}

/*
processTargetRule processes the tokens found within a target rule.
*/
func processTargetRules(itrc ITargetRulesContext) (T TargetRules, err error) {
	if itrc == nil {
		err = errorf("%T instance is nil; cannot proceed", itrc)
		return
	}

	ct := itrc.GetChildCount()
	T = make(TargetRules, 0)

	for k := 0; k < ct; k++ {
		switch tv := itrc.GetChild(k).(type) {

		case *TargetRuleContext:
			var t *Rule
			if t, err = processTargetRule(tv); err != nil {
				return
			}

			if targetRuleUnique(T, t) {
				T = append(T, t)
			} else {
				err = errorf("%T uniqueness violation (%s used more than once)", t.Keyword)
				return
			}
		}
	}

	if ok := targetRulesLengthOK(T); !ok {
		err = errorf("Invalid %T length (%d); must be 1 <= && <= 9",
			T, T.Len())
	}

	return
}

func targetRulesLengthOK(T []*Rule) (ok bool) {
	ok = (1 <= len(T) && len(T) <= 9)
	return
}

func targetRuleUnique(T []*Rule, t *Rule) (unique bool) {
	for i := 0; i < len(T); i++ {
		if t.Keyword == T[i].Keyword {
			return
		}
	}

	unique = true
	return
}

/*
processInstructionAnchor returns the ACL (the Access Control Label) found in an Instruction.
The label is returned as a string, alongside a boolean value indicative of a successful (>2
length) parsing of the label.
*/
func processInstructionAnchor(ins IInstructionAnchorContext) (label AccessControlLabel, err error) {

	// Bail out if the context is nil
	if ins == nil {
		err = errorf("%T is nil; cannot proceed", ins)
		return
	}

	// Be certain the access control label is not nil; we'll
	// grab the value towards the end.
	acl := ins.AccessControlLabel()
	if acl == nil {
		err = errorf("%T is nil; cannot proceed", acl)
		return
	}

	// Just be sure the instruction anchor is terminated
	// properly with a semicolon (ASCII #59); there is no
	// need to preserve this character, however, as it is
	// an ACIv3 "constant" that is predictable.
	if rt := ins.RuleTerminator(); rt == nil {
		err = errorf("Instruction anchor (version 3.0; acl \"...\") is not terminated properly (hint: ';')", ins)
		return
	}

	// make sure that not only did we get a label with
	// a total length greater than two (2) ...
	var raw string
	if raw = acl.GetText(); len(raw) > 2 {

		// ... but that the first and final slice elements
		// are actually double-quotes (ASCII #34).
		if ok := rune(raw[0]) == rune(34) &&
			rune(raw[0]) == rune(raw[len(raw)-1]); !ok {
			err = errorf("%T [%s] is not double-quoted; cannot proceed", label, raw)
			return
		}
	}

	label = AccessControlLabel(raw[1 : len(raw)-1])

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

func processBindPermission(p IPermissionContext) (r Permission, err error) {
	// rights []string, disp *bool) {
	// if nothing exists, this is a bogus context.
	if p == nil {
		err = errorf("%T instance is nil, cannot proceed", p)
		return
	}

	// grab the list of rights (privilege) identifiers.
	var foundRights bool
	if r.Rights, foundRights = processPermissionRights(p.AllPrivilege()); !foundRights {
		err = errorf("No rights found in %T", p)
		return
	}

	// grab the disposition pointer, or bail if nil.
	d := p.Disposition()
	if d == nil {
		err = errorf("%T instance is nil, cannot proceed", d)
		return
	}

	// Determine the permission disposition, or fail
	// if not an anticipated result ...
	if granted := d.GrantedPermission(); granted != nil {
		r.Allow = granted.ALLOW() != nil
	} else if d.WithheldPermission() != nil {
		r.Allow = false // force false for security reasons.
	}

	return
}

func processPermissionRights(r []IPrivilegeContext) (rights []string, ok bool) {
	// obtain slices of rights identifiers, appending
	// each to the return slice type.
	for i := 0; i < len(r); i++ {
		rights = append(rights, r[i].GetText())
	}

	// if length is zero, we're not OK.
	ok = len(rights) > 0
	return
}

func processBooleanWord(word any) (w BooleanWord, ok bool) {
	switch tv := word.(type) {

	case *WordAndContext:
		if tv != nil {
			w = 0
			ok = true
		}

	case *WordOrContext:
		if tv != nil {
			w = 1
			ok = true
		}

	case *WordNotContext:
		if tv != nil {
			w = 2
			ok = true
		}
	}

	return
}

/*
processBindRule processes the tokens found within a bind rule. A
bind rule (in its singular manifestation) is a bind keyword, a
comparison operator and a sequence of one (1) or more expression
values (with symbolic OR (||) delimitation when needed). Therefore
we need only match the contexts of those three (3) components.
*/
func processBindRule(ibrc IBindRuleContext) (r *Rule, err error) {
	if ibrc == nil {
		err = errorf("%T instance is nil; cannot proceed")
		return
	}

	ct := ibrc.GetChildCount()
	r = new(Rule)

	// iterate child components of the bind rule:
	// keyword, operator, value(s).
	for k := 0; k < ct; k++ {
		switch tv := ibrc.GetChild(k).(type) {

		case *BindKeywordContext:
			kw, _ := processRuleKeyword(tv)
			r.Keyword = kw

		case *BindOperatorContext:
			op, _ := processRuleOperator(tv)
			r.Operator = op

		case *ExpressionValuesContext:
			if r.Values, err = processRuleExpression(tv); err != nil {
				return
			}
		}
	}

	// its all or nothing

	if !ruleReady(r) {
		err = errorf("%T is missing one of 'Keyword', 'Operator' or 'Values'; cannot proceed", r)
	}

	return
}

/*
String is a stringer method that returns the string representation
of a Rule, which may be a Target Rule or a Bind Rule.
*/
func (r Rule) String() string {
	if ruleReady(&r) {
		return sprintf("%s %s %s",
			r.Keyword,
			r.Operator,
			r.Values)
	}

	return ``
}

func ruleReady(r *Rule) bool {
	if r == nil {
		return false
	}

	return r.Values.Len() > 0 &&
		len(r.Operator) > 0 &&
		len(r.Keyword) > 0
}

func processBindRules(ctx IBindRulesContext, depth int) (r BindRules, err error) {
	if ctx == nil {
		err = errorf("%T instance is nil; cannot proceed", ctx)
		return
	}

	r = make(BindRules, 0)

	for c := 0; c < ctx.GetChildCount(); c++ {

		// use the child indices to call the
		// slice from the antlr children array.
		child := ctx.GetChild(c)

		switch uv := child.(type) {

		// context type is a nested bind rules
		// (NOTE: plurality)
		case *BindRulesContext:
			if uv.GetChildCount() > 0 {
				var brs BindRules
				if brs, err = processBindRules(uv, depth); err != nil {
					return
				}
				r = append(r, brs)
			}

			// context type is a nested bind rule,
			// (NOTE: singular component of a bind
			// rules context)
		case *BindRuleContext:
			if br, _ := processBindRule(uv); br != nil {
				r = append(r, br)
			}

		// context type(s) are left/opener, meaning
		// this pchild stack is parenthetical. We
		// need not match the closing parenthetical.
		case *OpeningParenthesisContext:
			if uv != nil {
				r = append(r, sprintf("OPEN_PARENTHESIS_%d", depth))
				depth++
			}

		case *ClosingParenthesisContext:
			if uv != nil {
				depth--
				r = append(r, sprintf("CLOSE_PARENTHESIS_%d", depth))
			}

			// context type(s) are Boolean WORD
			// operators that "glue" bind rule(s)
			// together in a logical manner.
		case *WordOrContext,
			*WordAndContext,
			*WordNotContext:
			if uv != nil {
				w, _ := processBooleanWord(uv)
				r = append(r, w)
			}
		}
	}

	if r.Len() == 0 {
		err = errorf("No bind rules were parsed")
		return
	}

	return
}

/*
String is a stringer method that returns the string
representation of the receiver instance.
*/
func (r BooleanWord) String() string {
	switch r {
	case And:
		return `AND`
	case Or:
		return `OR`
	case Not:
		return `AND NOT`
	}

	return ``
}

/*
Len returns the integer length of the receiver instance.
*/
func (r BindRules) Len() int {
	return len(r)
}

/*
String is a stringer method that returns the string
representation of the receiver instance.
*/
func (r BindRules) String() string {
	var str string
	for i := 0; i < r.Len(); i++ {
		switch tv := r[i].(type) {
		case BindRules:
			str += tv.String()
		case *Rule:
			str += sprintf("%s", tv)
		case string:
			if hasPfx(tv, `OPEN`) {
				str += `( `
			} else if hasPfx(tv, `CLOSE`) {
				str += ` )`
			}
		case BooleanWord:
			str += sprintf(" %s ", tv)
		}
	}

	return str
}

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
ParseBindRules returns slices of any alongside a boolean value
indicative of a non-zero length return result.
*/
func ParseBindRules(raw string) (BindRules, error) {
	p, err := initAntlr(raw)
	if err != nil {
		return BindRules{}, err
	}

	// Parse and return our BindRules instance, which
	// (hopefully) contains one (1) or more Bind Rule(s).
	return processBindRules(p.BindRules(), 0)
}

/*
ParsePermission processes a permission statement and returns the disposition,
the rights identifiers and a boolean value indicative of a successful parse.

The disposition boolean instance, when true, indicates 'allow'; false indicates
'deny'.

The slices of rights identifiers must be non-zero in length for a successful
return result.
*/
func ParsePermission(raw string) (perm Permission, err error) {
	var p *ACIParser
	if p, err = initAntlr(raw); err != nil {
		return
	}

	return processBindPermission(p.Permission())
}

/*
ParsePermissionBindRule performs the same operations as ParsePermission
and ParseBindRules, except it processes them together as they normally would
appear in a complete Instruction.
*/
func ParsePermissionBindRule(raw string) (r PermissionBindRule, err error) {
	//(disp bool, rights []string, rules []any, ok bool) {
	p, err := initAntlr(raw)
	if err != nil {
		err = errorf("%T instance is nil; cannot proceed", p)
		return
	}

	// grab whatever PBRs are there, and make sure
	// return is not nil. We'll only be targeting
	// the first index going forward.
	pbr := p.PermissionBindRule()
	if pbr == nil {
		err = errorf("%T instance is nil; cannot proceed", pbr)
		return
	}

	// make sure the input was terminated properly
	if pbr.RuleTerminator() == nil {
		err = errorf("Permission+Bind rule is not terminated properly (hint: ';')")
		return
	}

	if r.Permission, err = processBindPermission(p.Permission()); err != nil {
		return
	}

	// obtain and verify bind rules "stack"
	r.BindRules, err = processBindRules(pbr.BindRules(), 0)

	return
}

/*
ParseTargetRules returns slices of *Rule instances alongside a
boolean value indicative of a non-zero length return result.
*/
func ParseTargetRules(raw string) (T TargetRules, err error) {
	p, err := initAntlr(raw)
	if err != nil {
		return
	}

	return processTargetRules(p.TargetRules())
}

/*
ParseInstruction returns an instance of Instruction alongside an error.

The input string value, raw, should be a legal ACIv3 instruction in its entirely.

See the following package-level functions if you only wish to parse a part of an instruction:

- ParseTargetRules

- ParsePermissionBindRule (or, alternatively, ParsePermission and ParseBindRules individually)

See also the AccessControlLabel type, which can cast a string value for an ACIv3 label.
*/
func ParseInstruction(raw string) (i Instruction, err error) {
	var p *ACIParser
	if p, err = initAntlr(raw); err != nil {
		return
	}

	ipc := p.Parse()
	ins := ipc.Instruction()

	if ins.OpeningParenthesis() == nil ||
		ins.ClosingParenthesis() == nil {
		// anchor/pbr not encapsulated
		// with parenthesis.
		err = errorf("%T is missing outer parenthetical encapsulation for anchor and pbr(s)", i)
		return
	}

	var (
		T  TargetRules
		L  AccessControlLabel
		PB PermissionBindRules = make(PermissionBindRules, 0)
	)

	// obtain our target rules, if any were
	// defined. target rules are optional.
	if T, err = processTargetRules(ins.TargetRules()); err != nil {
		return
	}

	// obtain the access control label
	if L, err = processInstructionAnchor(ins.InstructionAnchor()); err != nil {
		return
	}

	// A single permission+bind rule statement is comprised of one
	// (1) permission and one (1) bind rules expression, terminated
	// by a single semicolon (ASCII #59). An instruction MUST have
	// at least one (1) permission+bind statement that honors this
	// requirement, but may contain more if needed.
	for _, pbr := range ins.AllPermissionBindRule() {

		var p Permission
		var b BindRules

		if p, err = processBindPermission(pbr.Permission()); err != nil {
			return
		}

		// obtain and verify bind rules "stack"
		if b, err = processBindRules(pbr.BindRules(), 0); err != nil {
			return
		}

		PB = append(PB, PermissionBindRule{p, b})
	}

	return Instruction{T, L, PB}, nil
}
