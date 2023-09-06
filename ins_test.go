package antlraci

import (
	"testing"
)

/*
TestParseInstruction iterates the testInstructionManifest, parses each member value
and compares the return result with the original.
*/
func TestParseInstruction(t *testing.T) {
	ct := len(testInstructionManifest)

	for idx, want := range testInstructionManifest {
		/*
					// TODO: add bad instruction testers
					// similar to the other tests ...
					want, found := testInstructionManifest[idx]
			                if !found {
			                        t.Errorf("%s [VALID] failed [idx[%d]::%d/%d]: MISSING MAP ENTRY FOR INDEX %d?",
			                                t.Name(), idx, idx+1, ct, idx)
			                        return
			                }
		*/

		r, err := ParseInstruction(want)
		if err != nil {
			// There was an error
			if idx%2 == 0 {
				// Valid test should have worked, but did not.
				t.Errorf("%s [VALID] failed [idx[%d]::%d/%d]: %v\nwant: '%s'\ngot:  '%s'",
					t.Name(), idx, idx+1, ct, err, want, r)
				return
			}
			continue
		} else {
			// There was no error
			if idx%2 != 0 {
				// Invalid test should have failed, but did not.
				t.Errorf("%s [INVALID] failed [idx[%d]::%d/%d]: %T (%s) parse attempt returned no error",
					t.Name(), idx, idx+1, ct, r, r.String())
				return
			}
		}

		if got := r.String(); want != got {
			// There was an (unexpected) result during strcmp.
			t.Errorf("%s [VALID] failed [idx:%d/%d]:\nwant '%s'\ngot  '%s'",
				t.Name(), idx, ct, want, got)
			return
		}

	}

}
