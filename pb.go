package antlraci

/*
pb.go contains all Permission, Bind Rules and Permission+Bind Rules
types and methods.
*/

import (
	"github.com/JesseCoretta/go-stackage"
)

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
	P Permission
	B stackage.Stack
}

/*
String is a stringer method that returns the string
representation of the receiver instance.
*/
func (r PermissionBindRule) String() string {
	if len(r.P.Rights) == 0 {
		return ``
	}
	if r.B.Len() == 0 {
		return ``
	}

	return sprintf("%s %s;",
		r.P.String(),
		r.B.String())
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

	// grab the disposition, and bail if zero.
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

/*
ParseBindRule parses a single bind rule expression, e.g.: 'userdn = "ldap:///anyone"',
and returns a stackage.Condition instance alongside an error.
*/
func ParseBindRule(raw string) (r stackage.Condition, err error) {
	var p *ACIParser
	if p, err = initAntlr(raw); err != nil {
		return
	}

	// Parse and return our BindRule instance
	return processBindRule(p.BindRule())
}

/*
processBindRule processes the tokens found within a bind rule. A
bind rule (in its singular manifestation) is a bind keyword, a
comparison operator and a sequence of one (1) or more expression
values (with symbolic OR (||) delimitation when needed). Therefore
we need only match the contexts of those three (3) components.

A stackage.Condition, bearing the bind rule components, is returned
alongside an error.
*/
func processBindRule(ibrc IBindRuleContext) (r stackage.Condition, err error) {
	if ibrc == nil {
		err = errorf("%T instance is nil; cannot proceed")
		return
	}

	var ct int = ibrc.GetChildCount()

	// iterate child components of the bind rule:
	// keyword, operator, value(s).
	for k := 0; k < ct; k++ {
		var ok bool

		switch tv := ibrc.GetChild(k).(type) {

		case *BindKeywordContext:
			var kw string
			if kw, ok = processRuleKeyword(tv); !ok {
				err = errorf("Failed to process %T for Bind Rule keyword", tv)
				return
			}
			r.SetKeyword(kw)

		case *BindOperatorContext:
			var op string
			if op, ok = processRuleOperator(tv); !ok {
				err = errorf("Failed to process %T for Bind Rule operator", tv)
				return
			}

			cop, found := matchComparisonOperator(op)
			if !found {
				err = errorf("Unknown comparison operator '%s'", op)
				return
			}
			r.SetOperator(cop)

		case *ExpressionValuesContext:
			var ex RuleExpression
			if ex, err = processRuleExpression(tv); err != nil {
				return
			}
			r.SetExpression(ex)
		}
	}

	// Stamp the return as known to be a valid BIND rule
	r.SetCategory(`bind`)

	err = r.Valid()

	return
}

/*
ParseBindRules returns a stackage.Stack containing a Bind Rules
expression (which may or may not involving nesting). An instance
of error is returned alongisde the resulting stack.
*/
func ParseBindRules(raw string) (stackage.Stack, error) {
	p, err := initAntlr(raw)
	if err != nil {
		return stackage.Stack{}, err
	}

	// Parse and return our BindRules instance, which
	// (hopefully) contains one (1) or more Bind Rule(s).
	var X stackage.Stack
	brs := p.BindRules()
	w, _ := booleanContextFromBindRules(brs)
	X, err = processBindRules(brs, 0, isParentheticalContext(brs), w)
	return X, err
}

func initIOStack(ctx IBindRulesContext, word ...string) (cc, in stackage.Stack) {
	cc = stackageBasic() // contiguous condition stack, temporary use
	in = newStack()      // uninitialized, we dont know the bool word yet.

	// out repre
	if len(word) > 0 {
		in, _ = stackByBoolOp(word[0])
	} else {
		if nxw, found := booleanContextFromBindRules(ctx); found {
			in, _ = stackByBoolOp(nxw)
			in.Paren(isParentheticalContext(ctx))
		} else {
			in = stackageBasic()
		}
	}

	return
}

/*
processBindRules converts the Bind Rules token stream into a stackage.Stack
instance, which is returned alongside an error. Nested stacks and any opener
or closer encapsulating parentheticals will be honored wherever possible.
*/
func processBindRules(ctx IBindRulesContext, depth int, oparen bool, boolword ...string) (outer stackage.Stack, err error) {
	if ctx == nil {
		err = errorf("%T instance is nil; cannot proceed", ctx)
		return
	}

	var iparen bool = isParentheticalContext(ctx)
	var nxw string
	var cct int = ctx.GetChildCount()
	var cc, inner stackage.Stack
	cc, inner = initIOStack(ctx, boolword...)

	// iterate on all child objects, performing
	// recursive calls wherever needed.
	for c := 0; c < cct; c++ {

		// use the child indices to call the
		// slice from the antlr children array.
		child := ctx.GetChild(c)

		// is the current token (child, #c) within this
		// context sequence (ctx) parenthetical itself?
		//var bparen bool = currentTokenIsParenthetical(c, ctx)

		// perform type switch on current child instance
		// iteration ...
		switch {

		// context type is a nested bind rules (NOTE:
		// mind the plurality!)
		case isBindRulesContext(child):

			if cc.Len() > 0 {
				cc.Transfer(inner)
				cc.Reset()
			}

			if inner, err = handleBindRulesContext(
				child.(*BindRulesContext),
				inner,
				nxw,
				iparen || currentTokenIsParenthetical(c, ctx),
				depth); err != nil {
				return
			}

			if !hasSfx(nxw, `NOT`) && !hasSfx(trimS(inner.Kind()), `NOT`) {
				x, _ := stackByBoolOp(nxw)
				inner.Transfer(x)
				inner = x
			}

			// wipe out nxw, since it survives loops and influences
			// what kind of stacks are (or will be) created.
			nxw = ``

		// context type is a nested bind rule,
		// (NOTE: singular component of a bind
		// rules context)
		case isBindRuleContext(child):
			var br stackage.Condition
			if br, err = processBindRule(child.(*BindRuleContext)); err != nil {
				return
			}

			if !cc.IsParen() {
				cc.Push(br)
			}

		// context type is a parenthetical
		// child, either an opener (left)
		// or a closer (right).
		case isParentheticalContext(child):
			_, depth = handleParentheticalContext(c, depth, oparen, child)

		// context type is a Boolean logical
		// WORD operator, one of the following:
		// 'AND', 'OR', 'AND NOT'.
		case isBooleanWordContext(child):
			_, nxw = handleBooleanWordContext(child, inner)
		}
	}

	// We've exited the child context loop(s)
	// and have arrived at the final assembly
	// stage. Before we begin iterating the
	// inner stack, let's xfer the stack of
	// contiguous conditions, if applicable.
	if cc.Len() > 0 {
		cc.Transfer(inner)
	}

	// scan for any pointless enveloping
	// that might have occurred, and then
	// add inner to the return stack outer.
	//outer, err = cleanupBindStack(outer, inner, oparen)
	if inner.Len() == 0 {
		err = errorf("inner stack (%T) contains no slices", inner)
		return
	}
	outer = inner

	return
}

func booleanContextFromBindRules(ctx IBindRulesContext) (w string, found bool) {
	ct := ctx.GetChildCount()
	if ct != 3 {
		return
	}

	// get "middle slice, which ought to be a
	// Boolean WORD
	switch tv := ctx.GetChild(1).(type) {
	case *WordAndContext:
		if tv != nil {
			w = `AND`
		}
	case *WordOrContext:
		if tv != nil {
			w = `OR`
		}
	case *WordNotContext:
		if tv != nil {
			w = `AND NOT`
		}
	}

	found = len(w) > 0
	return
}

func isBooleanWordContext(ctx any) (is bool) {
	switch tv := ctx.(type) {
	case *WordAndContext:
		is = tv != nil
	case *WordOrContext:
		is = tv != nil
	case *WordNotContext:
		is = tv != nil
	}

	return
}

/*
get rid of this
*/
func handleBooleanWordContext(ctx any, child stackage.Stack) (stackage.Stack, string) {
	// nxw contains the (probable) Boolean
	// logical WORD operator for the next
	// stack.
	var nxw string

	switch tv := ctx.(type) {
	case *WordOrContext:
		nxw = tv.GetText()

	case *WordAndContext:
		nxw = tv.GetText()

	case *WordNotContext:
		nxw = tv.GetText()
	}

	return child, nxw
}

func isBindRuleContext(ctx any) bool {
	if assert, is := ctx.(*BindRuleContext); is {
		return assert != nil
	}

	return false
}

func isBindRulesContext(ctx any) bool {
	if assert, is := ctx.(*BindRulesContext); is {
		return assert != nil
	}

	return false
}

/*
handleBindRulesContext handles the recursive operations associated with a nested stack (BindRulesContext)
encountered during processBindRules execution. A populated incarnation of dest is returned, hopefully
containing one (1) or more elements obtained from the ctx input instance. An error is returned alongside
the stack instance should processing fail for any reason.
*/
func handleBindRulesContext(ctx *BindRulesContext, dest stackage.Stack, nxw string, paren bool, depth int) (stackage.Stack, error) {
	var (
		err error
		key string         // type of stack (e.g.: and, list, etc)
		rec stackage.Stack // stack returned via recursive call
	)

	if ctx.GetChildCount() == 0 {
		err = errorf("%T contains no children to process", ctx)
		return dest, err
	}

	// Spawn a new recursion of the CALLER of this function.
	if rec, err = processBindRules(ctx, depth, paren, nxw); err != nil {
		return dest, err
	}

	if rec.Len() == 0 {
		err = errorf("%T recursion returned empty sub-stack", ctx)
		return dest, err
	}

	key = trimS(rec.Kind())

	if eq(key, `not`) {
		if !eq(trimS(rec.Kind()), `NOT`) {
			rec, _ = disenvelopStack(rec)
			dest.Push(rec)
		} else {
			dest.Push(rec)
		}

		rec.Paren(paren && contextIsParenthetical(ctx)) // fixes #30

	} else {
		rec.Paren(paren)
		dest.Push(rec)
	}

	if dest.Len() == 0 {
		err = errorf("dest stack (%T) contains no slices", dest)
	}

	return dest, err
}

/*
ParsePermissionBindRule performs the same operations as ParsePermission
and ParseBindRules, except it processes them together as they normally would
appear in a complete Instruction.
*/
func ParsePermissionBindRule(raw string) (r PermissionBindRule, err error) {
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

	// Obtain the permission portion of the PB
	if r.P, err = processBindPermission(pbr.Permission()); err != nil {
		return
	}

	// obtain and verify bind rules stackage.Stack ...
	r.B, err = processBindRules(pbr.BindRules(), 0, false)

	return
}

/*
isNextTokenBoolOp returns a boolean WORD operator string value (w)
alongside a presence-indicative boolean value (is). This function
scans the BindRulesContext ahead by exactly one (1) token. If the
next token is AND, OR or (AND) NOT, the string representation is
returned.

If a CLOSING parenthesis (right) is encountered, this function will
recurse into another call of itself, scanning ahead by one (1) more
index in an attempt to find the "next" logical WORD operator.
*/
func isNextTokenBoolOp(i int, ctx IBindRulesContext) (w string, is bool) {
	if i+1 >= ctx.GetChildCount() {
		return
	}

	switch ctx.GetChild(i + 1).(type) {
	case *WordAndContext:
		w = `AND`
		is = true
	case *WordOrContext:
		w = `OR`
		is = true
	case *WordNotContext:
		w = `NOT` // this might need to be AND NOT :S
		is = true
		//case *ClosingParenthesisContext:
		// let's retry +1
		//w, is = isNextTokenBoolOp(i+1, ctx)
	}

	return
}

func handleParentheticalContext(idx, depth int, oparen bool, ctx any) (paren bool, d int) {
	// perform context type switch
	switch tv := ctx.(type) {

	// context type(s) are left/opener, meaning
	// this stack is parenthetical.
	case *OpeningParenthesisContext:
		if idx == 0 {
			if idx == 0 {
				paren = oparen && depth == 0
			}
			depth++
		}

	// context type(s) are right/closer.
	case *ClosingParenthesisContext:
		if tv != nil {
			depth--
		}
	}

	d = depth

	return
}

/*
currentTokenIsParenthetical is a convenient wrapper for getLN,
but only reports a boolean value indicative of whether this
index (i) resides between an opening parenthesis (left) and a
closing parenthesis (right). Both of these conditions must be
true for a true return value.

This function is distinct from contextIsParenthetical, as that
function looks INSIDE a context to discern its parenthetical
state. This function, on the other hand, analyzes a token within
a stream that is currently being iterated, using index (i) for
placement.
*/
func currentTokenIsParenthetical(i int, ctx IBindRulesContext) bool {
	var res [2]bool

	// get the previous and upcoming tokens from the
	// context using the i as the current positional
	// index. Both l and r must be non nil.
	if l, n := getLN(i, ctx); l != nil && n != nil {
		for i, val := range []any{
			l, // last (previous)
			n, // next (upcoming)
		} {
			switch val.(type) {
			case *OpeningParenthesisContext:
				res[i] = i == 0
			case *ClosingParenthesisContext:
				res[i] = i == 1
			}
		}
	}

	// regardless of the result above, evaluate
	// the boolean pair using &&.
	return res[0] && res[1]
}

/*
contextIsParenthetical reviews the first (0) and final (len-1)
context child values in order to determine whether the former
is an *OpeningParenthesisContext and the latter is its partner
*ClosingParenthesisContext. Note that a return of true doesn't
necessarily indicate the expression itself is parenthetical,
meaning there could be two (2) parenthetical sets, bound by a
Boolean WORD operator, which might still return true because
the first and last are openers and closers respectively, even
if they're not "mates". This is merely one (1) check of several
that might allow for logical deduction regarding the effective
parenthetical state of a context.
*/
func contextIsParenthetical(ctx IBindRulesContext) (ok bool) {
	count := ctx.GetChildCount()

	// In order for this evaluation to have
	// any meaning, the child count MUST be
	// three (3). Any parenthetical context
	// shall appear (logically) as:
	//
	//       0          1          2
	//   OPEN_PAREN SOMETHING CLOSE_PAREN
	//
	// The 'something' (slice #1) should be
	// a Bind Rule -OR- Bind RuleS context.
	if count != 3 {
		return
	}

	f := ctx.GetChild(0)         // first (left) child
	l := ctx.GetChild(count - 1) // final (right) child

	// Analyze first (f) and last (l) tokens acquired
	// through GetChild index calls above ...
	var x any
	if x, ok = f.(*OpeningParenthesisContext); ok && x != nil {
		// The first token was confirmed to be a
		// non-nil opening parenthetical context.

		if x, ok = l.(*ClosingParenthesisContext); ok && x != nil {
			// The final token was confirmed to be a
			// non-nil closing parenthetical context.

			// Before we finish, look at the "middle"
			// token. If it is either a Bind Rule or
			// Bind RuleS context, return true IF non
			// nil.
			if x, ok = ctx.GetChild(1).(*BindRulesContext); ok {
				ok = x != nil
			} else if x, ok = ctx.GetChild(1).(*BindRuleContext); ok {
				ok = x != nil
			}
		}
	}

	return
}

/*
isParentheticalContext returns a boolean value alongside a depth-indicative
integer. The boolean value, when true, indicates the context (ctx) input by
the caller is parenthetical in nature (e.g.: either a left/opening parenthesis
character (ASCII #40) or a right/closing parenthesis (ASCII #41) character).

This is generalized for use in switch/case matching.
*/
func isParentheticalContext(ctx any) (is bool) {
	// perform context type switch
	switch tv := ctx.(type) {
	case *OpeningParenthesisContext:
		if tv != nil {
			is = true
		}
	case *ClosingParenthesisContext:
		if tv != nil {
			is = true
		}
	}

	return
}

/*
parentheticalContextCount will scan the IBindRulesContext (ctx) instance for
parenthetical opener and closer characters (ASCII #40 and #41 respectively).
A complete pair, that is: one (1) opener and one (1) closer, shall result in
a +1 to the return value. Placement of these parenthetical characters is not
significant. This function is only interested in the number of occurrences,
and not in their distribution.

This function is useful for analyzing contexts which have nested parenthetical
stacks in an unusual -- but legal -- manner.
*/
func parentheticalContextCount(ctx IBindRulesContext) int {
	count := ctx.GetChildCount()
	if count <= 1 {
		return 0
	}

	var pair int
	var fs int

	for i := 0; i < count; i++ {
		switch tv := ctx.GetChild(i).(type) {
		case *OpeningParenthesisContext:
			if tv != nil {
				fs++
			}
		case *ClosingParenthesisContext:
			if fs == 1 && tv != nil {
				pair++
				fs--
				continue
			}
			return 0
		}
	}

	return pair
}

/*
getLN returns the previous and upcoming tokens within an IBindRulesContext
instance (ctx). If either last and/or next are nil, it means the request
is outside of the index boundary (e.g.: i > max token idx).
*/
func getLN(i int, ctx IBindRulesContext) (last any, next any) {
	count := ctx.GetChildCount()
	if count <= 1 {
		return
	}

	if i == 0 {
		return nil, ctx.GetChild(i + 1)
	} else if i == count-1 {
		return ctx.GetChild(i - 1), nil
	}

	return ctx.GetChild(i - 1), ctx.GetChild(i + 1)
}
