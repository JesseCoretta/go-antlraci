package antlraci

import (
	"testing"
)

/*
testBindRulesManifest contains a collection of bind rule(s) expressions
found within the wild as well as vendor documentation alike. Bug reports
pertaining to the incorrect handling of expressions (particularly within
stackage.Stack instances) should result in new additions here.
*/
var testBindRulesManifest map[int]string = map[int]string{
	0:  `( groupdn = "ldap:///cn=Human Resources,dc=example,dc.com" )`,
	1:  `( userattr = "aciurl#LDAPURL" )`,
	2:  `( userdn = "ldap:///all" )`,
	3:  `( userdn = "ldap:///anyone || ldap:///parent || ldap:///self" )`,
	4:  `( userdn = "ldap:///anyone" ) AND ( ip != "192.0.2." )`,
	5:  `( userdn = "ldap:///anyone" )`,
	6:  `( userdn = "ldap:///self" ) AND ( ssf >= "128" )`,
	7:  `( userdn = "ldap:///self" )`,
	8:  `( userdn = "ldap:///uid=user,ou=People,dc=example,dc.com" )`,
	9:  `( userdn = "ldap:///uid=user,ou=People,dc=example,dc=com" ) AND ( dayofweek = "Sun,Sat" )`,
	10: `( userdn = "ldap:///uid=user,ou=People,dc=example,dc=com" ) AND ( timeofday >= "1800" AND timeofday < "2400" )`,
	11: `groupdn = "ldap:///cn=DomainAdmins,ou=Groups,($dn),dc=example,dc=com"`,
	12: `groupdn = "ldap:///cn=DomainAdmins,ou=Groups,[$dn],dc=example,dc=com"`,
	13: `groupdn = "ldap:///cn=DomainAdmins,ou=Groups,dc=subdomain1,dc=hostedCompany1,dc=example,dc=com"`,
	14: `groupdn = "ldap:///cn=example,ou=groups,dc=example,dc=com"`,
	15: `userattr = "manager#USERDN"`,
	16: `userattr = "owner#USERDN"`,
	17: `userattr = "parent[0,1].owner#USERDN"`,
	18: `userattr = "parent[1].manager#USERDN"`,
	19: `userdn = "ldap:///anyone" AND ssf < "128"`,
	20: `userdn = "ldap:///anyone"`,
	21: `userdn = "ldap:///self" AND ssf >= "128"`,
	22: `( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" ) )`,
	23: `( ( userdn = "ldap:///anyone" AND ssf >= "128" ) AND NOT dayofweek = "Fri" )`,
	24: `( authmethod = "NONE" OR authmethod = "SIMPLE" )`,
	25: `( userdn = "ldap:///anyone" ) AND ( ip != "2001:db8::" )`,
	26: `groupdn = "ldap:///cn=Administrators,ou=Groups,dc=example,com" AND groupdn = "ldap:///cn=Operators,ou=Groups,dc=example,com"`,
	27: `roledn = "ldap:///cn=Human Resources,ou=People,dc=example,dc=com"`,
	28: `userattr = "manager#USERDN"`,
	29: `userdn = "ldap:///anyone" AND ssf >= "128" AND NOT ( dayofweek = "Fri" OR dayofweek = "Sun" )`,
	30: `userdn = "ldap:///anyone" AND ssf >= "128" AND NOT dayofweek = "Fri"`,
	31: `userdn = "ldap:///anyone"`,
	32: `userdn = "ldap:///cn=Courtney Tolana,dc=example,dc=com"`,
	33: `userdn = "ldap:///dc=example,dc=com??sub?(manager=example)"`,
	34: `userdn = "ldap:///ou=People,dc=example,dc=com??sub?(department=Human Resources)"`,
	35: `userdn = "ldap:///parent"`,
	//36: `( userdn = "ldap:///anyone" ) AND NOT ( dns != "client.example.com" )`,	// FIX ME :P
}

/*
TestBindRules iterates the testBindRulesManifest, parses each member value
and compares the return result with the original.
*/
func TestBindRules(t *testing.T) {

	ct := len(testBindRulesManifest)
	var ect int
	for idx := 0; idx < ct; idx++ {
		want, found := testBindRulesManifest[idx]
		if !found {
			continue // ??
		}

		r, err := ParseBindRules(want)
		if err != nil {
			t.Errorf("%s failed [idx:%d/%d]: %v", t.Name(), idx, ct, err)
		}

		if got := r.String(); want != got {
			ect++
			printf("\n###################################\n")
			printStack(r, 0)
			printf("###################################\n")
			t.Errorf("%s failed [idx:%d/%d]:\nwant '%s' [%d]\ngot  '%s' [%d]", t.Name(), idx, ct, want, r.Len(), got, r.Len())
		}
	}

	if ect != 0 {
		t.Logf("%s suffered %d total errors", t.Name(), ect)
	}

}

//want := `allow(read,write) ( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) );`
//r, err := parser.ParsePermissionBindRule(want)

/*
testPermissionManifest contains a collection of bind rule(s) expressions
found within the wild as well as vendor documentation alike. Bug reports
pertaining to the incorrect handling of expressions (particularly within
stackage.Stack instances) should result in new additions here.
*/
var testPermissionManifest map[int]string = map[int]string{
	0: `allow(read,write)`,
	1: `allow(read,write,compare,search,delete)`,
	2: `deny(all)`,
	3: `deny(proxy)`,
}

/*
TestPermission iterates the testPermissionManifest, parses each member value
and compares the return result with the original.
*/
func TestPermission(t *testing.T) {

	for idx := 0; idx < len(testPermissionManifest); idx++ {
		want, found := testPermissionManifest[idx]
		if !found {
			continue // ??
		}

		r, err := ParsePermission(want)
		if err != nil {
			t.Errorf("%s failed [idx:%d]: %v", t.Name(), idx, err)
		}

		if got := r.String(); want != got {
			t.Errorf("%s failed [idx:%d]:\nwant '%s'\ngot  '%s'", t.Name(), idx, want, got)
		}
	}

}
