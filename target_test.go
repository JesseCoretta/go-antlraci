package antlraci

import (
	"testing"
)

/*
testTargetRulesManifest contains a collection of target rule(s) expressions
found within the wild as well as vendor documentation alike. Bug reports
pertaining to the incorrect handling of expressions (particularly within
stackage.Stack instances) should result in new additions here.
*/
var testTargetRulesManifest map[int]string = map[int]string{
	0: `( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )( targetcontrol = "1.2.3.4" || "5.6.7.8" )( targetscope = "onelevel" )`,
	1: `( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )`,
}

/*
TestTargetRules iterates the testTargetRulesManifest, parses each member value
and compares the return result with the original.
*/
func TestTargetRules(t *testing.T) {

	for idx := 0; idx < len(testTargetRulesManifest); idx++ {
		want, found := testTargetRulesManifest[idx]
		if !found {
			continue // ??
		}

		r, err := ParseTargetRules(want)
		if err != nil {
			t.Errorf("%s failed [idx:%d]: %v", t.Name(), idx, err)
		}

		if got := r.String(); want != got {
			t.Errorf("%s failed [idx:%d]:\nwant '%s'\ngot  '%s'", t.Name(), idx, want, got)
		}
	}

}
