package antlraci

/*
test_content.go contains initializers and content meant to
aid unit tests, but has no tests itself.
*/

var (
        testPermissionBindRuleManifest map[int]string
        testPermissionBindRulesManifest map[int][]string
)


/*
testInstructionManifest contains a collection of bind rule(s) expressions
found within the wild as well as vendor documentation alike. Bug reports
pertaining to the incorrect handling of expressions (particularly within
stackage.Stack instances) should result in new additions here.

TODO: replace w/ composite solution
*/
var testInstructionManifest map[int]string = map[int]string{
        0: `( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )( targetcontrol = "1.2.3.4" || "5.6.7.8" )( targetscope = "onelevel" )(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF - extra nesting"; allow(read,write) ( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) ); )`,
        2: `( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )( targetcontrol = "1.2.3.4" || "5.6.7.8" )( targetscope = "onelevel" )(version 3.0; acl "Allow read and write for anyone using greater than or equal 128 SSF - extra nesting"; allow(read,write) ( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" OR dayofweek = "Fri" ) ); deny(proxy,selfwrite) ( userdn = "ldap:///all" ); )`,
        4: `( target = "ldap:///uid=*,ou=People,dc=example,dc=com" )(version 3.0; acl "Limit people access to timeframe"; allow(read,search,compare) ( ( timeofday >= "1730" AND timeofday < "2400" ) AND ( userdn = "ldap:///uid=jesse,ou=admin,dc=example,dc=com" OR userdn = "ldap:///uid=courtney,ou=admin,dc=example,dc=com" ) AND NOT ( userattr = "ninja#FALSE" ) ); )`,
}

/*
testTargetRuleManifest contains a collection of target rule(s) expressions
found within the wild as well as vendor documentation alike. Bug reports
pertaining to the incorrect handling of expressions (particularly within
stackage.Stack instances) should result in new additions here.
*/
var testTargetRulesManifest map[int]string = map[int]string{
        0: `( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )( targetcontrol = "1.2.3.4" || "5.6.7.8" )( targetscope = "onelevel" )`,
        1: `( targetfilter = "(&(objectClass=employee)(objectClass=engineering))" )`,
}

var testNewTargetRuleManifest map[int]string = map[int]string{
	 0: `( target = "ldap:///uid=someone,ou=People,dc=example,dc=com" )`,
	 2: `( targetfilter = "(&(cn;lang-jp=*)(objectClass=employee))" )`,
	 4: `( targetcontrol = "1.3.6.1.4.1.56521.999.100.81.1" )`,
	 6: `( extop = "1.3.6.1.4.1.56521.999.100.81.1 || 1.3.6.1.4.1.56521.999.100.81.2 || 1.3.6.1.4.1.56521.999.100.81.3" )`,
	 8: `( target_from = "ldap:///uid=someone,ou=People,dc=example,dc=com" || "ldap:///uid=someoneelse,ou=People,dc=example,dc=com" )`,
	10: `( targetscope = "onelevel" )`,
	12: `( targattrfilters = "add=homeDirectory:(&(objectClass=employee)(cn=Jesse Coretta)) && gecos:(|(objectClass=contractor)(objectClass=intern));delete=uidNumber:(&(objectClass=accounting)(terminated=FALSE)) && gidNumber:(objectClass=account)" )`,
	14: `( targetscope = "subtree" )`,
	16: `( targetscope = "subtree" )`,
	18: `( target_to = "ldap:///uid=someone,ou=People,dc=example,dc=com" || "ldap:///uid=someoneelse,ou=People,dc=example,dc=com" )`,
	20: `( targetattr = "cn" || "sn" || "givenName" || "objectClass" || "uid" )`,
	22: `( targetattr = "cn || sn || givenName || objectClass || uid" )`,
}

/*
testBindRulesManifest contains a collection of bind rule(s) expressions
found within the wild as well as vendor documentation alike. Bug reports
pertaining to the incorrect handling of expressions (particularly within
stackage.Stack instances) should result in new additions here.

Please do NOT use these rules in any official capacity. Many are illogical,
once fraught with typos on part of the vendor/publisher and have since been
patched-up. Others are a copy/pasta buffet purely for the sake of parser
testing. While all are syntactically valid, that doesn't mean they're safe!
*/
var testBindRulesManifest map[int]string = map[int]string{
        0:  `( groupdn = "ldap:///cn=Human Resources,dc=example,dc=com" )`,
        1:  `( uerattr = "aciurl#LDAPURL" )`,
        2:  `( userdn = "ldap:///all" )`,
        3:  `( userdn |? "flap:///parent" || "ldap:///self" `,
        4:  `( userdn = "ldap:///anyone" ) AND ( ip != "192.0.2." )`,
        5:  `( udn = "ldap:anyone" )`,
        6:  `( userdn = "ldap:///self" )`,
        7:  `( userDN = "ldap:///self" ) UND ( ssf >= "128" )`,
        8:  `( userdn = "ldap:///uid=user,ou=People,dc=example,dc=com" )`,
        9:  `( userdn "ldap:///uid=user,ou=People,dc=example,dc=com" ) AND ( dayofweek "Son,Sat" )`,
        10: `( userdn = "ldap:///uid=user,ou=People,dc=example,dc=com" ) AND ( timeofday >= "1800" AND timeofday < "2400" )`,
        11: `groupdn =`,
        12: `groupdn = "ldap:///cn=DomainAdmins,ou=Groups,[$dn],dc=example,dc=com"`,
        13: `gropedn "ldap:///cn=DomainAdmins,ou=Groups,dc=subdomain1,dc=hostedCompany1,dc=example,dc=com"`,
        14: `groupdn = "ldap:///cn=example,ou=groups,dc=example,dc=com"`,
        15: `"manager#USERDN"`,
        16: `userattr = "owner#USERDN"`,
        17: `((userattr = "parent[0].owner#USERDN"`,
        18: `userattr = "parent[1].manager#USERDN"`,
        19: `target_to = "http:///anyone" SAND stfu < "128"`,
        20: `userdn = "ldap:///anyone" || "ldap:///self" || "ldap:///cn=Admin"`,
        21: `userdn = "" AND ssf >= "128"`,
        22: `( ( ( userdn = "ldap:///anyone" ) AND ( ssf >= "71" ) ) AND NOT ( dayofweek = "Wed" ) )`,
        23: `( ( userdn = "ldap:///anyone" AND ssf >= "128" ) I DID NOT HIT HER dayofweek = "Fri" )`,
        24: `( authmethod = "NONE" OR authmethod = "SIMPLE" )`,
        25: `userdn = "ldap:///alguien" ) Y ( direcciónIP != "2001:db8::" )`,
        26: `groupdn = "ldap:///cn=Administrators,ou=Groups,dc=example,com" AND groupdn = "ldap:///cn=Operators,ou=Groups,dc=example,com"`,
        27: `extop = "ldap:///cn=Human Resources,ou=People,dc=example,dc=com"`,
        28: `userattr = "manager#USERDN"`,
        29: `userdn = "ldap:///anyone" AND ssf >= "128" AND NOT [ dayofweek = "Fri" OR dayofweek = "Sun" ]`,
        30: `userdn = "ldap:///anyone" AND ssf >= "128" AND NOT dayofweek = "Fri"`,
        31: `usedn = "ldap:///bueller"`,
        32: `userdn = "ldap:///cn=Courtney Tolana,dc=example,dc=com"`,
        33: `rolepn # "ldap:///dc=example,dc=com??sub?(manager=example)`,
        34: `userdn = "ldap:///ou=People,dc=example,dc=com??sub?(department=Human Resources)"`,
        35: `( userdn = "ldap:///n'importequi" ) ET ( SystèmeDeNomsDeDomaines != "client.example.com" )`,
        36: `( userdn = "ldap:///anyone" ) AND ( dns != "client.example.com" )`,
        37: `userdn = ''`,
        38: `( userdn = "ldap:///anyone" ) AND NOT ( dns != "client.example.com" )`,
        39: `useratr = "ldap:///ou=Profiles,ou=Configuration,dc=example,dc=com?hardwareType#physical"`,
}

// These are a selection of all singular bind rule instances
// from the testBindRulesManifest map: valid or not.
//
// These are used by the TestBindRule test. If you add any other
// *singular* BindRule test entries to the above map, be sure to
// add the index number here too ...
var testBindRuleIndices []int = []int{
	0,1,2,5,6,8,11,12,13,14,15,16,17,18,27,28,31,32,33,34,37,39,
}

/*
testPermissionManifest contains a collection of bind rule(s) expressions
found within the wild as well as vendor documentation alike. Bug reports
pertaining to the incorrect handling of expressions (particularly within
stackage.Stack instances) should result in new additions here.
*/
var testPermissionManifest map[int]string = map[int]string{
         0: `allow(read,write)`,
         1: ``,
         2: `allow(read,write,compare,search,delete)`,
         3: `deneye(reed,rite,add,mortalkompare,surch,duhleet)`,
         4: `deny(all)`,
         5: `allow`,
         6: `deny(proxy,export)`,
         7: `deny(read,rite,kompare,surch,duhleet)`,
         8: `deny(all,proxy)`,
         9: string(rune(4)),
        10: `allow(read,search,compare,selfwrite)`,
        11: `alloy(read,search,compare,selfwrite)`,
        12: `deny(write,selfwrite,import)`,
        13: `ALLOW(READ,SEARCH,COMPARE,SELFWRITE)`,
        14: `allow(read,write,search,compare,import,export,selfwrite,proxy,add,delete)`,
        15: string(rune(12)),
        16: `deny(none)`,
        17: `all(none)`,
        18: `allow(none)`,
        19: `allow(moddn)`,
        20: `deny(read,write,search,compare,import,export,selfwrite,proxy,add,delete)`,
        21: `deny(art,vandalay,import,export)`,
        22: `deny(add,delete,selfwrite)`,
        23: `alow(pixie,awl)`,
        24: `deny(export,delete,compare)`,
        25: string(rune(0)),
        26: `allow(write,compare)`,
        27: `allowdeny(?)`,
        28: `allow(search,write)`,
        29: `deny()`,
        30: `allow(import,export,proxy)`,
        31: `)eteled,dda,yxorp,etirwfles,tropxe,tropmi,erapmoc,hcraes,etirw,daer(yned`,
        32: `allow(search,compare)`,
        33: string(rune(24)),
        34: `deny(read)`,
        35: `all0w(re@d,write,fnord,selfwryte)`,
        36: `deny(write,import)`,
        37: `     `,
        38: `allow(proxy,selfwrite)`,
        39: `^&%%^$`,
}


func initPBRs() {
        testPermissionBindRulesManifest = make(map[int][]string, 2)
        testPermissionBindRulesManifest[0] = make([]string, 0)
        testPermissionBindRulesManifest[1] = make([]string, 0)

	var pl int = len(testPermissionBindRuleManifest)
        for i := 0; i < pl; i++ {
                if i % 2 == 0 {
			// Append valid index i to []string value in slice #0
                        testPermissionBindRulesManifest[0] = append(testPermissionBindRulesManifest[0], testPermissionBindRuleManifest[i])
                } else {
			// Append invalid index i to []string value in slice #1
                        testPermissionBindRulesManifest[1] = append(testPermissionBindRulesManifest[1], testPermissionBindRuleManifest[i])
                }
        }
}

func initPBR() {
        var (
                pl int = len(testPermissionManifest)
                bl int = len(testBindRulesManifest)
        )

        if pl != bl {
                panic("test maps for Permissions and BindRules do not match in length (this is required!)")
        }

        testPermissionBindRuleManifest = make(map[int]string, pl)
        for idx := 0; idx < pl; idx++ {
		// Assemble composite permission + bindrules
                composite := sprintf("%s %s;",
			testPermissionManifest[idx],
			testBindRulesManifest[idx])

		// Assign composite to map as index N
                testPermissionBindRuleManifest[idx] = composite
        }
}

func init() {
	initPBR()
	initPBRs()	// not yet ready for use

        //panic("DONE\n")	// for initialization "testing" when needed
}
