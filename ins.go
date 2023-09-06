package antlraci

import (
	"github.com/JesseCoretta/go-stackage"
)

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
Instruction is a convenience type returned when ParseInstruction
is successfully executed. It will contain all of the central elements
of an Access Control Instruction (ACI) Version 3.0 expression:

- Zero (0) or more Target Rules

- An Access Control Label (unique, descriptive string identifier for an ACI)

- One (1) or more Permission+Bind Rules, each containing one (1)
Permission and one (1) Bind Rules expression.
*/
type Instruction struct {
	T  stackage.Stack     // optional
	L  AccessControlLabel // required
	PB stackage.Stack     // >= 1 required
}

/*
String is a stringer method that returns the string
representation of the receiver instance.
*/
func (r Instruction) String() string {
	if r.PB.Len() == 0 {
		return ``
	}

	if len(r.L) == 0 {
		return ``
	}

	var str string
	if r.T.Len() > 0 {
		str += r.T.String()
	}

	str += sprintf("(version 3.0; acl \"%s\"; ", r.L.String())
	str += r.PB.String()
	str += ` )`

	return str
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
	} else if hasPfx(rt.GetText(), `<missing`) {
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
		T  stackage.Stack	// target rules
		L  AccessControlLabel	// access control label
		PB stackage.Stack	// permission bind rules
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
	if PB, err = processPermissionBindRules(ins.AllPermissionBindRule()); err != nil {
		return
	}

	/*
	for _, pbr := range ins.AllPermissionBindRule() {

		var p Permission
		var b stackage.Stack

		if p, err = processBindPermission(pbr.Permission()); err != nil {
			return
		}

		// obtain and verify bind rules "stack"
		if b, err = processBindRules(pbr.BindRules(), 0, false); err != nil {
			return
		}

		PB.Push(PermissionBindRule{p, b})
	}
	*/

	return Instruction{T, L, PB}, nil
}
