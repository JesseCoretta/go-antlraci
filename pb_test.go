package antlraci

import (
	"testing"
)

/*
TestParseBindRules iterates the testBindRulesManifest and directly executes
the ParseBindRules package-level function using the current iteration value
and comparing the return result with the original.

- Even numbered map entries are VALID tests which SHOULD NOT FAIL for any reason

- Odd numbered map entries are INVALID tests which SHOULD FAIL w/ an error
*/
func TestParseBindRules(t *testing.T) {

	ct := len(testBindRulesManifest)

	var ect int
	for idx := 0; idx < ct; idx++ {
		want, found := testBindRulesManifest[idx]
		if !found {
			t.Errorf("%s failed [idx:%d/%d]: MISSING MAP ENTRY FOR INDEX %d?",
				t.Name(), idx, ct, idx)
			return
		}

		r, err := ParseBindRules(want)
		if err != nil {
			// There was an error
			if idx%2 == 0 {
				// Valid test should have worked, but did not.
				t.Errorf("%s failed [idx:%d/%d]: %v\nwant: '%s'\ngot:  '%s'",
					t.Name(), idx, ct, err, want, r)
				return
			}
			continue
		} else {
			// There was no error
			if idx%2 != 0 {
				// Invalid test should have failed, but did not.
				t.Errorf("%s failed [idx:%d/%d]: invalid %T (%s) parse attempt returned no error",
					t.Name(), idx, ct, r, r.String())
				return
			}
		}

		if got := r.String(); want != got {

			// There was an (unexpected) result during strcmp.
			ect++

			// execute rudimentary stack dumper for
			// troubleshooting of failed strcmp ...
			printf("\n###################################\n")
			printStack(r, 0)
			printf("###################################\n")

			t.Errorf("%s failed [idx:%d/%d]:\nwant '%s' [%d]\ngot  '%s' [%d]",
				t.Name(), idx, ct, want, r.Len(), got, r.Len())
		}
	}

	if ect != 0 {
		t.Logf("%s suffered %d total errors", t.Name(), ect)
	}
}

/*
TestParseBindRule tests the direct execution of the *singular* ParseBindRule
package-level function.

A BindRule is a statement of the following syntax:

	                                           Available expression values/contexts
	                                      ----------------------------------------------
	   One (1) Bind Keyword        +----- Allows LDAP DistinguishedName/URI abstracts,
	+---------------------------+ /       possibly in a multi-valued context
	| userdn,  groupdn, roledn, |+
	+---------------------------+         Allows network/mechanism/confidential
	| ip, dns, ssf, authmethod, |+------- abstract contexts
	+---------------------------+
	| userattr, groupattr,      |+------- Allows composite attr/val value-matching
	+---------------------------+         statements, possibly enveloped as a URI
	| timeofday or dayofweek    |+        or basic attributeBindTypeOrValue context
	+---------------------------+ \
	    |                          +----- Allows temporal contexts
	    |    <comparison
	<keyword> operator> <expression>                 ^
	       ______|______        \                    |
	      /  |  | |  |  \        See expression descriptions
	     EQ LE LT GT GE NE

- Even numbered map entries are VALID tests which SHOULD NOT FAIL for any reason

- Odd numbered map entries are INVALID tests which SHOULD FAIL w/ an error
*/
func TestParseBindRule(t *testing.T) {
	ct := len(testBindRuleIndices)

	for i := 0; i < ct; i++ {
		idx := testBindRuleIndices[i]
		want, found := testBindRulesManifest[idx]
		if !found {
			continue // don't error, just skip ahead.
		}

		r, err := ParseBindRule(want)
		if err != nil {
			// There was an error ...
			if idx%2 == 0 {
				// Valid test should have worked, but did not ...
				t.Errorf("%s [VALID] failed [idx[%d]::%d/%d]: err:%v\nwant: '%s'\ngot:  '%s'",
					t.Name(), idx, i, ct, err, want, r)
				return
			}
			continue
		} else {
			// There was no error ...
			if idx%2 != 0 {
				// Invalid test should have failed, but did not ...
				t.Errorf("%s [INVALID] failed [idx[%d]::%d/%d]: %T (%s) parse attempt returned no error",
					t.Name(), idx, i, ct, r, want)
				return
			}
		}

		if got := r.String(); want != got {
			// There was an (unexpected) result during strcmp.
			t.Errorf("%s [VALID] failed [idx[%d]::%d/%d]:\nwant '%s'\ngot  '%s'",
				t.Name(), idx, i, ct, want, got)
			return
		}
	}
}

/*
TestPermissions iterates the testPermissionManifest, and execting the ParsePermission
package-level function using each member as input, and then comparing the return result
with the original.

- Even numbered map entries are VALID tests which SHOULD NOT FAIL for any reason

- Odd numbered map entries are INVALID tests which SHOULD FAIL w/ an error
*/
func TestParsePermissions(t *testing.T) {

	ct := len(testPermissionManifest)

	for idx := 0; idx < ct; idx++ {
		want, found := testPermissionManifest[idx]
		if !found {
			t.Errorf("%s failed [idx:%d/%d]: MISSING MAP ENTRY FOR INDEX %d?",
				t.Name(), idx, ct, idx)
			return
		}

		r, err := ParsePermission(want)
		if err != nil {
			// There was an error ...
			if idx%2 == 0 || idx == 0 {
				// Valid test should have worked, but did not ...
				t.Errorf("%s failed [idx:%d/%d]: err:%v",
					t.Name(), idx, ct, err)
				return
			}
			continue // if it was supposed to fail, skip ahead
		} else {
			// There was no error ...
			if idx%2 != 0 {
				// Invalid test should have failed, but did not ...
				t.Errorf("%s failed [idx:%d/%d: no error returned for bogus value [%s]",
					t.Name(), idx, ct, want)
				return
			}
		}

		if got := r.String(); want != got {
			// There was an (unexpected) result during strcmp.
			t.Errorf("%s failed [idx:%d/%d]:\nwant '%s'\ngot  '%s'",
				t.Name(), idx, ct, want, got)
			return
		}
	}

}

/*
TestPermissionBindRule tests the direct execution of the ParsePermissionBindRule
package level function. A PermissionBindRule is a statement of the following syntax:

	               delimited  ASCII           ASCII
	               privilege   #32              #59
	   allow/deny    names      |                |
	       |           |        |                |
	[ disposition (right,...)  WHSP [BindRules]  ; ]
	 ------------------------   |   -----------   \
	        Permission          |    BindRules     \
	                            +                   +- terminator
				     \
	                              \
	                          seperator

- Even numbered map entries are VALID tests which SHOULD NOT FAIL for any reason

- Odd numbered map entries are INVALID tests which SHOULD FAIL w/ an error
*/
func TestParsePermissionBindRule(t *testing.T) {
	ct := len(testPermissionBindRuleManifest)

	var err error
	for idx := 0; idx < ct; idx++ {
		want, found := testPermissionBindRuleManifest[idx]
		if !found {
			t.Errorf("%s [idx:%d/%d]: MISSING MAP ENTRY FOR INDEX %d?",
				t.Name(), idx, ct, idx)
			return
		}

		var got PermissionBindRule
		if got, err = ParsePermissionBindRule(want); err != nil {
			// There was an error ...
			if idx%2 == 0 || idx == 0 {
				// Valid test should have worked, but did not ...
				t.Errorf("%s [VALID] failed [idx:%d/%d]: err:%v",
					t.Name(), idx, ct, err)
				return
			}

			continue // if it was supposed to fail, skip ahead
		}

		// There was no error ...
		if idx%2 != 0 {
			// Invalid test should have failed, but did not ...
			t.Errorf("%s [INVALID] failed [idx:%d/%d]: %T parse returned no error",
				t.Name(), idx, ct, got)
			return
		}

		if got.String() != want {
			// There was an (unexpected) result during strcmp.
			t.Errorf("%s [VALID] failed [idx:%d/%d]: unexpected result;\nwant '%s'\ngot  '%s'",
				t.Name(), idx, ct, want, got)
			return
		}
	}
}

/*
- Even numbered map entries are VALID tests which SHOULD NOT FAIL for any reason

- Odd numbered map entries are INVALID tests which SHOULD FAIL w/ an error
*/
func TestParsePermissionBindRules(t *testing.T) {
	ct := len(testPermissionBindRulesManifest)

	var err error
	for idx := 0; idx < ct; idx++ {
		wanted, found := testPermissionBindRulesManifest[idx]
		if !found {
			t.Errorf("%s [idx[%d]::%d/%d]: MISSING MAP ENTRY FOR INDEX %d?",
				t.Name(), idx, idx+1, ct, idx)
			return
		}

		want := join(wanted, string(rune(32)))

		got := newStack()
		if got, err = ParsePermissionBindRules(want); err != nil {
			// There was an error ...
			if idx%2 == 0 {
				// Valid test should have worked, but did not ...
				t.Errorf("%s [VALID] failed [idx[%d]::%d/%d]: err:%v",
					t.Name(), idx, idx+1, ct, err)
				return
			}

			continue // if it was supposed to fail, skip ahead
		}

		// There was no error ...
		if idx%2 != 0 {
			// Invalid test should have failed, but did not ...
			t.Errorf("%s [INVALID] failed [idx[%d]::%d/%d]: invalid %T parse returned no error",
				t.Name(), idx, idx+1, ct, got)
			return
		}

		if got.String() != want {
			// There was an (unexpected) result during strcmp.
			t.Errorf("%s [VALID] failed [idx[%d]::%d/%d]: unexpected result;\nwant '%s'\ngot  '%s'",
				t.Name(), idx, idx+1, ct, want, got)
			return
		}
	}
}
