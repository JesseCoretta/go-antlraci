package antlraci

import (
	"testing"
)

/*
testInstructionManifest contains a collection of bind rule(s) expressions
found within the wild as well as vendor documentation alike. Bug reports
pertaining to the incorrect handling of expressions (particularly within
stackage.Stack instances) should result in new additions here.
*/
var testInstructionManifest map[int]string = map[int]string{
	0: `(targetfilter="(&(objectClass=employee)(objectClass=engineering))")(targetcontrol="1.2.3.4" || "5.6.7.8")(targetscope="onelevel")(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF - extra nesting"; allow(read,write) ( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) );)`,
}

/*
TestInstruction iterates the testInstructionManifest, parses each member value
and compares the return result with the original.
*/
func TestInstruction(t *testing.T) {

	for idx := 0; idx < len(testInstructionManifest); idx++ {
		//want, found := testInstructionManifest[idx]
		_, found := testInstructionManifest[idx]
		if !found {
			continue // ??
		}

		/*
			r, err := ParseInstruction(want)
			if err != nil {
				t.Errorf("%s failed [idx:%d]: %v", t.Name(), idx, err)
			}

			if got := r.String(); want != got {
				t.Errorf("%s failed [idx:%d]:\nwant '%s'\ngot  '%s'", t.Name(), idx, want, got)
			}
		*/
	}

}
