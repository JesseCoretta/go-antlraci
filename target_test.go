package antlraci

import (
	"testing"
)

/*
TestTargetRules iterates the testTargetRulesManifest, parses each
member value and compares the return result with the original.
*/
func TestTargetRule(t *testing.T) {

	ct := len(testNewTargetRuleManifest)

	for idx, want := range testNewTargetRuleManifest {
		//for idx := 0; idx < ct; idx++ {
		//want, found := testNewTargetRuleManifest[idx]
                //if !found {
                //       t.Errorf("%s failed [idx:%d/%d]: MISSING MAP ENTRY FOR INDEX %d?",
                //              t.Name(), idx, ct, idx)
                //       return
                //}

		r, err := ParseTargetRules(want)
		if err != nil {
			// There was an error ...
			if idx % 2 == 0 || idx == 0 {
				// Valid test should have worked, but did not ...
                                t.Errorf("%s [VALID] failed [idx:%d/%d]: err:%v\nwant: '%s'\ngot:  '%s'",
                                        t.Name(), idx, ct, err, want, r)
                                return
			}
			continue
		} else {
			// There was no error ...
			if idx % 2 != 0 {
                                // Invalid test should have failed, but did not ...
                                t.Errorf("%s [INVALID] failed [idx:%d/%d]: %T (%s) parse attempt returned no error",
                                        t.Name(), idx, ct, r, want)
                                return
                        }
		}

		r.NoPadding(true)
		if got := r.String(); want != got {
			t.Errorf("%s failed [idx:%d/%d]:\nwant '%s'\ngot  '%s'",
				t.Name(), idx, ct, want, got)
		}
	}

}
