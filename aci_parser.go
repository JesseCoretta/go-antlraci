// Code generated from ACIParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package antlraci // ACIParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ACIParser struct {
	*antlr.BaseParser
}

var ACIParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func aciparserParserInit() {
	staticData := &ACIParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'userdn'", "'groupdn'", "'roledn'", "'userattr'", "'groupattr'",
		"'ssf'", "'ip'", "'dns'", "'authmethod'", "'timeofday'", "'dayofweek'",
		"'target'", "'target_to'", "'target_from'", "'targetscope'", "'targetattr'",
		"'targetfilter'", "'targattrfilters'", "'targetcontrol'", "'extop'",
		"'\"'", "'('", "')'", "','", "';'", "'='", "'!='", "'<'", "'<='", "'>'",
		"'>='", "'&&'", "'||'", "", "'read'", "'write'", "'add'", "'compare'",
		"'search'", "'delete'", "'proxy'", "'export'", "'import'", "'selfwrite'",
		"'all'", "'none'", "'allow'", "'deny'",
	}
	staticData.SymbolicNames = []string{
		"", "BKW_UDN", "BKW_GDN", "BKW_RDN", "BKW_UAT", "BKW_GAT", "BKW_SSF",
		"BKW_IP", "BKW_DNS", "BKW_AM", "BKW_TOD", "BKW_DOW", "TKW_TARGET", "TKW_TO",
		"TKW_FROM", "TKW_SCOPE", "TKW_ATTR", "TKW_FILTER", "TKW_AF", "TKW_CTRL",
		"TKW_EXTOP", "DQUOTE", "LPAREN", "RPAREN", "COMMA", "SEMI", "EQ", "NE",
		"LT", "LE", "GT", "GE", "SYMBOLIC_AND", "SYMBOLIC_OR", "ANCHOR", "READ_PRIV",
		"WRITE_PRIV", "ADD_PRIV", "CMP_PRIV", "SRC_PRIV", "DEL_PRIV", "PRX_PRIV",
		"EXP_PRIV", "IMP_PRIV", "SLF_PRIV", "ALL_PRIV", "NO_PRIV", "ALLOW",
		"DENY", "WORD_AND", "WORD_OR", "WORD_NOT", "NCTF_TO_WHSP", "QUOTED_STRING",
		"WS",
	}
	staticData.RuleNames = []string{
		"parse", "instruction", "instructionAnchor", "permissionBindRule", "permission",
		"disposition", "grantedPermission", "withheldPermission", "privilege",
		"readPrivilege", "writePrivilege", "selfWritePrivilege", "comparePrivilege",
		"searchPrivilege", "proxyPrivilege", "addPrivilege", "deletePrivilege",
		"importPrivilege", "exportPrivilege", "allPrivilege", "noPrivilege",
		"targetRule", "targetOperator", "targetRules", "wordAnd", "wordOr",
		"wordNot", "bindRules", "bindRule", "expressionValues", "symbolicOr",
		"bindOperator", "lessThan", "lessThanOrEqual", "greaterThan", "greaterThanOrEqual",
		"equalTo", "notEqualTo", "bindKeyword", "bindUserDistinguishedName",
		"bindGroupDistinguishedName", "bindRoleDistinguishedName", "bindUserAttributes",
		"bindGroupAttributes", "bindSecurityStrengthFactor", "bindIPAddress",
		"bindDNSName", "bindTimeOfDay", "bindDayOfWeek", "bindAuthenticationMethod",
		"targetKeyword", "targetDistinguishedName", "targetToDistinguishedName",
		"targetFromDistinguishedName", "targetSearchScope", "targetSearchAttribute",
		"targetSearchFilter", "targetAttributeSearchFilters", "targetControlObjectIdentifier",
		"targetExtendedOperationObjectIdentifier", "openingParenthesis", "closingParenthesis",
		"commaDelimiter", "ruleTerminator", "anchor", "accessControlLabel",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 372, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 1, 0, 1, 0, 1, 0, 1, 1, 3, 1, 137,
		8, 1, 1, 1, 1, 1, 1, 1, 4, 1, 142, 8, 1, 11, 1, 12, 1, 143, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 5, 4, 162, 8, 4, 10, 4, 12, 4, 165, 9, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 3, 5, 171, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 189, 8, 8, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 3,
		22, 223, 8, 22, 1, 23, 4, 23, 226, 8, 23, 11, 23, 12, 23, 227, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		3, 27, 242, 8, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 248, 8, 27, 1, 27,
		1, 27, 4, 27, 252, 8, 27, 11, 27, 12, 27, 253, 5, 27, 256, 8, 27, 10, 27,
		12, 27, 259, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1,
		29, 5, 29, 269, 8, 29, 10, 29, 12, 29, 272, 9, 29, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 282, 8, 31, 1, 32, 1, 32, 1,
		33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3,
		38, 307, 8, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1,
		48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50,
		1, 50, 1, 50, 3, 50, 340, 8, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1,
		53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58,
		1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1,
		64, 1, 64, 1, 65, 1, 65, 1, 65, 0, 1, 54, 66, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86,
		88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118,
		120, 122, 124, 126, 128, 130, 0, 0, 351, 0, 132, 1, 0, 0, 0, 2, 136, 1,
		0, 0, 0, 4, 147, 1, 0, 0, 0, 6, 151, 1, 0, 0, 0, 8, 155, 1, 0, 0, 0, 10,
		170, 1, 0, 0, 0, 12, 172, 1, 0, 0, 0, 14, 174, 1, 0, 0, 0, 16, 188, 1,
		0, 0, 0, 18, 190, 1, 0, 0, 0, 20, 192, 1, 0, 0, 0, 22, 194, 1, 0, 0, 0,
		24, 196, 1, 0, 0, 0, 26, 198, 1, 0, 0, 0, 28, 200, 1, 0, 0, 0, 30, 202,
		1, 0, 0, 0, 32, 204, 1, 0, 0, 0, 34, 206, 1, 0, 0, 0, 36, 208, 1, 0, 0,
		0, 38, 210, 1, 0, 0, 0, 40, 212, 1, 0, 0, 0, 42, 214, 1, 0, 0, 0, 44, 222,
		1, 0, 0, 0, 46, 225, 1, 0, 0, 0, 48, 229, 1, 0, 0, 0, 50, 231, 1, 0, 0,
		0, 52, 233, 1, 0, 0, 0, 54, 241, 1, 0, 0, 0, 56, 260, 1, 0, 0, 0, 58, 264,
		1, 0, 0, 0, 60, 273, 1, 0, 0, 0, 62, 281, 1, 0, 0, 0, 64, 283, 1, 0, 0,
		0, 66, 285, 1, 0, 0, 0, 68, 287, 1, 0, 0, 0, 70, 289, 1, 0, 0, 0, 72, 291,
		1, 0, 0, 0, 74, 293, 1, 0, 0, 0, 76, 306, 1, 0, 0, 0, 78, 308, 1, 0, 0,
		0, 80, 310, 1, 0, 0, 0, 82, 312, 1, 0, 0, 0, 84, 314, 1, 0, 0, 0, 86, 316,
		1, 0, 0, 0, 88, 318, 1, 0, 0, 0, 90, 320, 1, 0, 0, 0, 92, 322, 1, 0, 0,
		0, 94, 324, 1, 0, 0, 0, 96, 326, 1, 0, 0, 0, 98, 328, 1, 0, 0, 0, 100,
		339, 1, 0, 0, 0, 102, 341, 1, 0, 0, 0, 104, 343, 1, 0, 0, 0, 106, 345,
		1, 0, 0, 0, 108, 347, 1, 0, 0, 0, 110, 349, 1, 0, 0, 0, 112, 351, 1, 0,
		0, 0, 114, 353, 1, 0, 0, 0, 116, 355, 1, 0, 0, 0, 118, 357, 1, 0, 0, 0,
		120, 359, 1, 0, 0, 0, 122, 361, 1, 0, 0, 0, 124, 363, 1, 0, 0, 0, 126,
		365, 1, 0, 0, 0, 128, 367, 1, 0, 0, 0, 130, 369, 1, 0, 0, 0, 132, 133,
		3, 2, 1, 0, 133, 134, 5, 0, 0, 1, 134, 1, 1, 0, 0, 0, 135, 137, 3, 46,
		23, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0,
		138, 139, 3, 120, 60, 0, 139, 141, 3, 4, 2, 0, 140, 142, 3, 6, 3, 0, 141,
		140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144,
		1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 3, 122, 61, 0, 146, 3, 1, 0,
		0, 0, 147, 148, 3, 128, 64, 0, 148, 149, 3, 130, 65, 0, 149, 150, 3, 126,
		63, 0, 150, 5, 1, 0, 0, 0, 151, 152, 3, 8, 4, 0, 152, 153, 3, 54, 27, 0,
		153, 154, 3, 126, 63, 0, 154, 7, 1, 0, 0, 0, 155, 156, 3, 10, 5, 0, 156,
		157, 3, 120, 60, 0, 157, 163, 3, 16, 8, 0, 158, 159, 3, 124, 62, 0, 159,
		160, 3, 16, 8, 0, 160, 162, 1, 0, 0, 0, 161, 158, 1, 0, 0, 0, 162, 165,
		1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 166, 1, 0,
		0, 0, 165, 163, 1, 0, 0, 0, 166, 167, 3, 122, 61, 0, 167, 9, 1, 0, 0, 0,
		168, 171, 3, 12, 6, 0, 169, 171, 3, 14, 7, 0, 170, 168, 1, 0, 0, 0, 170,
		169, 1, 0, 0, 0, 171, 11, 1, 0, 0, 0, 172, 173, 5, 47, 0, 0, 173, 13, 1,
		0, 0, 0, 174, 175, 5, 48, 0, 0, 175, 15, 1, 0, 0, 0, 176, 189, 3, 18, 9,
		0, 177, 189, 3, 20, 10, 0, 178, 189, 3, 22, 11, 0, 179, 189, 3, 24, 12,
		0, 180, 189, 3, 26, 13, 0, 181, 189, 3, 28, 14, 0, 182, 189, 3, 30, 15,
		0, 183, 189, 3, 32, 16, 0, 184, 189, 3, 34, 17, 0, 185, 189, 3, 36, 18,
		0, 186, 189, 3, 38, 19, 0, 187, 189, 3, 40, 20, 0, 188, 176, 1, 0, 0, 0,
		188, 177, 1, 0, 0, 0, 188, 178, 1, 0, 0, 0, 188, 179, 1, 0, 0, 0, 188,
		180, 1, 0, 0, 0, 188, 181, 1, 0, 0, 0, 188, 182, 1, 0, 0, 0, 188, 183,
		1, 0, 0, 0, 188, 184, 1, 0, 0, 0, 188, 185, 1, 0, 0, 0, 188, 186, 1, 0,
		0, 0, 188, 187, 1, 0, 0, 0, 189, 17, 1, 0, 0, 0, 190, 191, 5, 35, 0, 0,
		191, 19, 1, 0, 0, 0, 192, 193, 5, 36, 0, 0, 193, 21, 1, 0, 0, 0, 194, 195,
		5, 44, 0, 0, 195, 23, 1, 0, 0, 0, 196, 197, 5, 38, 0, 0, 197, 25, 1, 0,
		0, 0, 198, 199, 5, 39, 0, 0, 199, 27, 1, 0, 0, 0, 200, 201, 5, 41, 0, 0,
		201, 29, 1, 0, 0, 0, 202, 203, 5, 37, 0, 0, 203, 31, 1, 0, 0, 0, 204, 205,
		5, 40, 0, 0, 205, 33, 1, 0, 0, 0, 206, 207, 5, 43, 0, 0, 207, 35, 1, 0,
		0, 0, 208, 209, 5, 42, 0, 0, 209, 37, 1, 0, 0, 0, 210, 211, 5, 45, 0, 0,
		211, 39, 1, 0, 0, 0, 212, 213, 5, 46, 0, 0, 213, 41, 1, 0, 0, 0, 214, 215,
		3, 120, 60, 0, 215, 216, 3, 100, 50, 0, 216, 217, 3, 44, 22, 0, 217, 218,
		3, 58, 29, 0, 218, 219, 3, 122, 61, 0, 219, 43, 1, 0, 0, 0, 220, 223, 3,
		72, 36, 0, 221, 223, 3, 74, 37, 0, 222, 220, 1, 0, 0, 0, 222, 221, 1, 0,
		0, 0, 223, 45, 1, 0, 0, 0, 224, 226, 3, 42, 21, 0, 225, 224, 1, 0, 0, 0,
		226, 227, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228,
		47, 1, 0, 0, 0, 229, 230, 5, 49, 0, 0, 230, 49, 1, 0, 0, 0, 231, 232, 5,
		50, 0, 0, 232, 51, 1, 0, 0, 0, 233, 234, 5, 51, 0, 0, 234, 53, 1, 0, 0,
		0, 235, 236, 6, 27, -1, 0, 236, 237, 3, 120, 60, 0, 237, 238, 3, 54, 27,
		0, 238, 239, 3, 122, 61, 0, 239, 242, 1, 0, 0, 0, 240, 242, 3, 56, 28,
		0, 241, 235, 1, 0, 0, 0, 241, 240, 1, 0, 0, 0, 242, 257, 1, 0, 0, 0, 243,
		251, 10, 2, 0, 0, 244, 248, 3, 48, 24, 0, 245, 248, 3, 50, 25, 0, 246,
		248, 3, 52, 26, 0, 247, 244, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 247, 246,
		1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 3, 54, 27, 0, 250, 252, 1,
		0, 0, 0, 251, 247, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 251, 1, 0, 0,
		0, 253, 254, 1, 0, 0, 0, 254, 256, 1, 0, 0, 0, 255, 243, 1, 0, 0, 0, 256,
		259, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 55, 1,
		0, 0, 0, 259, 257, 1, 0, 0, 0, 260, 261, 3, 76, 38, 0, 261, 262, 3, 62,
		31, 0, 262, 263, 3, 58, 29, 0, 263, 57, 1, 0, 0, 0, 264, 270, 5, 53, 0,
		0, 265, 266, 3, 60, 30, 0, 266, 267, 5, 53, 0, 0, 267, 269, 1, 0, 0, 0,
		268, 265, 1, 0, 0, 0, 269, 272, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 270,
		271, 1, 0, 0, 0, 271, 59, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 273, 274, 5,
		33, 0, 0, 274, 61, 1, 0, 0, 0, 275, 282, 3, 72, 36, 0, 276, 282, 3, 74,
		37, 0, 277, 282, 3, 64, 32, 0, 278, 282, 3, 66, 33, 0, 279, 282, 3, 68,
		34, 0, 280, 282, 3, 70, 35, 0, 281, 275, 1, 0, 0, 0, 281, 276, 1, 0, 0,
		0, 281, 277, 1, 0, 0, 0, 281, 278, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 281,
		280, 1, 0, 0, 0, 282, 63, 1, 0, 0, 0, 283, 284, 5, 28, 0, 0, 284, 65, 1,
		0, 0, 0, 285, 286, 5, 29, 0, 0, 286, 67, 1, 0, 0, 0, 287, 288, 5, 30, 0,
		0, 288, 69, 1, 0, 0, 0, 289, 290, 5, 31, 0, 0, 290, 71, 1, 0, 0, 0, 291,
		292, 5, 26, 0, 0, 292, 73, 1, 0, 0, 0, 293, 294, 5, 27, 0, 0, 294, 75,
		1, 0, 0, 0, 295, 307, 3, 78, 39, 0, 296, 307, 3, 80, 40, 0, 297, 307, 3,
		82, 41, 0, 298, 307, 3, 84, 42, 0, 299, 307, 3, 86, 43, 0, 300, 307, 3,
		88, 44, 0, 301, 307, 3, 90, 45, 0, 302, 307, 3, 92, 46, 0, 303, 307, 3,
		94, 47, 0, 304, 307, 3, 96, 48, 0, 305, 307, 3, 98, 49, 0, 306, 295, 1,
		0, 0, 0, 306, 296, 1, 0, 0, 0, 306, 297, 1, 0, 0, 0, 306, 298, 1, 0, 0,
		0, 306, 299, 1, 0, 0, 0, 306, 300, 1, 0, 0, 0, 306, 301, 1, 0, 0, 0, 306,
		302, 1, 0, 0, 0, 306, 303, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306, 305,
		1, 0, 0, 0, 307, 77, 1, 0, 0, 0, 308, 309, 5, 1, 0, 0, 309, 79, 1, 0, 0,
		0, 310, 311, 5, 2, 0, 0, 311, 81, 1, 0, 0, 0, 312, 313, 5, 3, 0, 0, 313,
		83, 1, 0, 0, 0, 314, 315, 5, 4, 0, 0, 315, 85, 1, 0, 0, 0, 316, 317, 5,
		5, 0, 0, 317, 87, 1, 0, 0, 0, 318, 319, 5, 6, 0, 0, 319, 89, 1, 0, 0, 0,
		320, 321, 5, 7, 0, 0, 321, 91, 1, 0, 0, 0, 322, 323, 5, 8, 0, 0, 323, 93,
		1, 0, 0, 0, 324, 325, 5, 10, 0, 0, 325, 95, 1, 0, 0, 0, 326, 327, 5, 11,
		0, 0, 327, 97, 1, 0, 0, 0, 328, 329, 5, 9, 0, 0, 329, 99, 1, 0, 0, 0, 330,
		340, 3, 102, 51, 0, 331, 340, 3, 104, 52, 0, 332, 340, 3, 106, 53, 0, 333,
		340, 3, 108, 54, 0, 334, 340, 3, 110, 55, 0, 335, 340, 3, 112, 56, 0, 336,
		340, 3, 114, 57, 0, 337, 340, 3, 116, 58, 0, 338, 340, 3, 118, 59, 0, 339,
		330, 1, 0, 0, 0, 339, 331, 1, 0, 0, 0, 339, 332, 1, 0, 0, 0, 339, 333,
		1, 0, 0, 0, 339, 334, 1, 0, 0, 0, 339, 335, 1, 0, 0, 0, 339, 336, 1, 0,
		0, 0, 339, 337, 1, 0, 0, 0, 339, 338, 1, 0, 0, 0, 340, 101, 1, 0, 0, 0,
		341, 342, 5, 12, 0, 0, 342, 103, 1, 0, 0, 0, 343, 344, 5, 13, 0, 0, 344,
		105, 1, 0, 0, 0, 345, 346, 5, 14, 0, 0, 346, 107, 1, 0, 0, 0, 347, 348,
		5, 15, 0, 0, 348, 109, 1, 0, 0, 0, 349, 350, 5, 16, 0, 0, 350, 111, 1,
		0, 0, 0, 351, 352, 5, 17, 0, 0, 352, 113, 1, 0, 0, 0, 353, 354, 5, 18,
		0, 0, 354, 115, 1, 0, 0, 0, 355, 356, 5, 19, 0, 0, 356, 117, 1, 0, 0, 0,
		357, 358, 5, 20, 0, 0, 358, 119, 1, 0, 0, 0, 359, 360, 5, 22, 0, 0, 360,
		121, 1, 0, 0, 0, 361, 362, 5, 23, 0, 0, 362, 123, 1, 0, 0, 0, 363, 364,
		5, 24, 0, 0, 364, 125, 1, 0, 0, 0, 365, 366, 5, 25, 0, 0, 366, 127, 1,
		0, 0, 0, 367, 368, 5, 34, 0, 0, 368, 129, 1, 0, 0, 0, 369, 370, 5, 53,
		0, 0, 370, 131, 1, 0, 0, 0, 15, 136, 143, 163, 170, 188, 222, 227, 241,
		247, 253, 257, 270, 281, 306, 339,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ACIParserInit initializes any static state used to implement ACIParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewACIParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ACIParserInit() {
	staticData := &ACIParserParserStaticData
	staticData.once.Do(aciparserParserInit)
}

// NewACIParser produces a new parser instance for the optional input antlr.TokenStream.
func NewACIParser(input antlr.TokenStream) *ACIParser {
	ACIParserInit()
	this := new(ACIParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ACIParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ACIParser.g4"

	return this
}

// ACIParser tokens.
const (
	ACIParserEOF           = antlr.TokenEOF
	ACIParserBKW_UDN       = 1
	ACIParserBKW_GDN       = 2
	ACIParserBKW_RDN       = 3
	ACIParserBKW_UAT       = 4
	ACIParserBKW_GAT       = 5
	ACIParserBKW_SSF       = 6
	ACIParserBKW_IP        = 7
	ACIParserBKW_DNS       = 8
	ACIParserBKW_AM        = 9
	ACIParserBKW_TOD       = 10
	ACIParserBKW_DOW       = 11
	ACIParserTKW_TARGET    = 12
	ACIParserTKW_TO        = 13
	ACIParserTKW_FROM      = 14
	ACIParserTKW_SCOPE     = 15
	ACIParserTKW_ATTR      = 16
	ACIParserTKW_FILTER    = 17
	ACIParserTKW_AF        = 18
	ACIParserTKW_CTRL      = 19
	ACIParserTKW_EXTOP     = 20
	ACIParserDQUOTE        = 21
	ACIParserLPAREN        = 22
	ACIParserRPAREN        = 23
	ACIParserCOMMA         = 24
	ACIParserSEMI          = 25
	ACIParserEQ            = 26
	ACIParserNE            = 27
	ACIParserLT            = 28
	ACIParserLE            = 29
	ACIParserGT            = 30
	ACIParserGE            = 31
	ACIParserSYMBOLIC_AND  = 32
	ACIParserSYMBOLIC_OR   = 33
	ACIParserANCHOR        = 34
	ACIParserREAD_PRIV     = 35
	ACIParserWRITE_PRIV    = 36
	ACIParserADD_PRIV      = 37
	ACIParserCMP_PRIV      = 38
	ACIParserSRC_PRIV      = 39
	ACIParserDEL_PRIV      = 40
	ACIParserPRX_PRIV      = 41
	ACIParserEXP_PRIV      = 42
	ACIParserIMP_PRIV      = 43
	ACIParserSLF_PRIV      = 44
	ACIParserALL_PRIV      = 45
	ACIParserNO_PRIV       = 46
	ACIParserALLOW         = 47
	ACIParserDENY          = 48
	ACIParserWORD_AND      = 49
	ACIParserWORD_OR       = 50
	ACIParserWORD_NOT      = 51
	ACIParserNCTF_TO_WHSP  = 52
	ACIParserQUOTED_STRING = 53
	ACIParserWS            = 54
)

// ACIParser rules.
const (
	ACIParserRULE_parse                                   = 0
	ACIParserRULE_instruction                             = 1
	ACIParserRULE_instructionAnchor                       = 2
	ACIParserRULE_permissionBindRule                      = 3
	ACIParserRULE_permission                              = 4
	ACIParserRULE_disposition                             = 5
	ACIParserRULE_grantedPermission                       = 6
	ACIParserRULE_withheldPermission                      = 7
	ACIParserRULE_privilege                               = 8
	ACIParserRULE_readPrivilege                           = 9
	ACIParserRULE_writePrivilege                          = 10
	ACIParserRULE_selfWritePrivilege                      = 11
	ACIParserRULE_comparePrivilege                        = 12
	ACIParserRULE_searchPrivilege                         = 13
	ACIParserRULE_proxyPrivilege                          = 14
	ACIParserRULE_addPrivilege                            = 15
	ACIParserRULE_deletePrivilege                         = 16
	ACIParserRULE_importPrivilege                         = 17
	ACIParserRULE_exportPrivilege                         = 18
	ACIParserRULE_allPrivilege                            = 19
	ACIParserRULE_noPrivilege                             = 20
	ACIParserRULE_targetRule                              = 21
	ACIParserRULE_targetOperator                          = 22
	ACIParserRULE_targetRules                             = 23
	ACIParserRULE_wordAnd                                 = 24
	ACIParserRULE_wordOr                                  = 25
	ACIParserRULE_wordNot                                 = 26
	ACIParserRULE_bindRules                               = 27
	ACIParserRULE_bindRule                                = 28
	ACIParserRULE_expressionValues                        = 29
	ACIParserRULE_symbolicOr                              = 30
	ACIParserRULE_bindOperator                            = 31
	ACIParserRULE_lessThan                                = 32
	ACIParserRULE_lessThanOrEqual                         = 33
	ACIParserRULE_greaterThan                             = 34
	ACIParserRULE_greaterThanOrEqual                      = 35
	ACIParserRULE_equalTo                                 = 36
	ACIParserRULE_notEqualTo                              = 37
	ACIParserRULE_bindKeyword                             = 38
	ACIParserRULE_bindUserDistinguishedName               = 39
	ACIParserRULE_bindGroupDistinguishedName              = 40
	ACIParserRULE_bindRoleDistinguishedName               = 41
	ACIParserRULE_bindUserAttributes                      = 42
	ACIParserRULE_bindGroupAttributes                     = 43
	ACIParserRULE_bindSecurityStrengthFactor              = 44
	ACIParserRULE_bindIPAddress                           = 45
	ACIParserRULE_bindDNSName                             = 46
	ACIParserRULE_bindTimeOfDay                           = 47
	ACIParserRULE_bindDayOfWeek                           = 48
	ACIParserRULE_bindAuthenticationMethod                = 49
	ACIParserRULE_targetKeyword                           = 50
	ACIParserRULE_targetDistinguishedName                 = 51
	ACIParserRULE_targetToDistinguishedName               = 52
	ACIParserRULE_targetFromDistinguishedName             = 53
	ACIParserRULE_targetSearchScope                       = 54
	ACIParserRULE_targetSearchAttribute                   = 55
	ACIParserRULE_targetSearchFilter                      = 56
	ACIParserRULE_targetAttributeSearchFilters            = 57
	ACIParserRULE_targetControlObjectIdentifier           = 58
	ACIParserRULE_targetExtendedOperationObjectIdentifier = 59
	ACIParserRULE_openingParenthesis                      = 60
	ACIParserRULE_closingParenthesis                      = 61
	ACIParserRULE_commaDelimiter                          = 62
	ACIParserRULE_ruleTerminator                          = 63
	ACIParserRULE_anchor                                  = 64
	ACIParserRULE_accessControlLabel                      = 65
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Instruction() IInstructionContext
	EOF() antlr.TerminalNode

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_parse
	return p
}

func InitEmptyParseContext(p *ParseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_parse
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Instruction() IInstructionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(ACIParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParse(s)
	}
}

func (p *ACIParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ACIParserRULE_parse)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Instruction()
	}
	{
		p.SetState(133)
		p.Match(ACIParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpeningParenthesis() IOpeningParenthesisContext
	InstructionAnchor() IInstructionAnchorContext
	ClosingParenthesis() IClosingParenthesisContext
	TargetRules() ITargetRulesContext
	AllPermissionBindRule() []IPermissionBindRuleContext
	PermissionBindRule(i int) IPermissionBindRuleContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) OpeningParenthesis() IOpeningParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpeningParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpeningParenthesisContext)
}

func (s *InstructionContext) InstructionAnchor() IInstructionAnchorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionAnchorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionAnchorContext)
}

func (s *InstructionContext) ClosingParenthesis() IClosingParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosingParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosingParenthesisContext)
}

func (s *InstructionContext) TargetRules() ITargetRulesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetRulesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetRulesContext)
}

func (s *InstructionContext) AllPermissionBindRule() []IPermissionBindRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPermissionBindRuleContext); ok {
			len++
		}
	}

	tst := make([]IPermissionBindRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPermissionBindRuleContext); ok {
			tst[i] = t.(IPermissionBindRuleContext)
			i++
		}
	}

	return tst
}

func (s *InstructionContext) PermissionBindRule(i int) IPermissionBindRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPermissionBindRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPermissionBindRuleContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *ACIParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ACIParserRULE_instruction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(135)
			p.TargetRules()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(138)
		p.OpeningParenthesis()
	}
	{
		p.SetState(139)
		p.InstructionAnchor()
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ACIParserALLOW || _la == ACIParserDENY {
		{
			p.SetState(140)
			p.PermissionBindRule()
		}

		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(145)
		p.ClosingParenthesis()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionAnchorContext is an interface to support dynamic dispatch.
type IInstructionAnchorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Anchor() IAnchorContext
	AccessControlLabel() IAccessControlLabelContext
	RuleTerminator() IRuleTerminatorContext

	// IsInstructionAnchorContext differentiates from other interfaces.
	IsInstructionAnchorContext()
}

type InstructionAnchorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionAnchorContext() *InstructionAnchorContext {
	var p = new(InstructionAnchorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_instructionAnchor
	return p
}

func InitEmptyInstructionAnchorContext(p *InstructionAnchorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_instructionAnchor
}

func (*InstructionAnchorContext) IsInstructionAnchorContext() {}

func NewInstructionAnchorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionAnchorContext {
	var p = new(InstructionAnchorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_instructionAnchor

	return p
}

func (s *InstructionAnchorContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionAnchorContext) Anchor() IAnchorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnchorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnchorContext)
}

func (s *InstructionAnchorContext) AccessControlLabel() IAccessControlLabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessControlLabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessControlLabelContext)
}

func (s *InstructionAnchorContext) RuleTerminator() IRuleTerminatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleTerminatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleTerminatorContext)
}

func (s *InstructionAnchorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionAnchorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionAnchorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterInstructionAnchor(s)
	}
}

func (s *InstructionAnchorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitInstructionAnchor(s)
	}
}

func (p *ACIParser) InstructionAnchor() (localctx IInstructionAnchorContext) {
	localctx = NewInstructionAnchorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ACIParserRULE_instructionAnchor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Anchor()
	}
	{
		p.SetState(148)
		p.AccessControlLabel()
	}
	{
		p.SetState(149)
		p.RuleTerminator()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPermissionBindRuleContext is an interface to support dynamic dispatch.
type IPermissionBindRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Permission() IPermissionContext
	BindRules() IBindRulesContext
	RuleTerminator() IRuleTerminatorContext

	// IsPermissionBindRuleContext differentiates from other interfaces.
	IsPermissionBindRuleContext()
}

type PermissionBindRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPermissionBindRuleContext() *PermissionBindRuleContext {
	var p = new(PermissionBindRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_permissionBindRule
	return p
}

func InitEmptyPermissionBindRuleContext(p *PermissionBindRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_permissionBindRule
}

func (*PermissionBindRuleContext) IsPermissionBindRuleContext() {}

func NewPermissionBindRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PermissionBindRuleContext {
	var p = new(PermissionBindRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_permissionBindRule

	return p
}

func (s *PermissionBindRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *PermissionBindRuleContext) Permission() IPermissionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPermissionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPermissionContext)
}

func (s *PermissionBindRuleContext) BindRules() IBindRulesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRulesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRulesContext)
}

func (s *PermissionBindRuleContext) RuleTerminator() IRuleTerminatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleTerminatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuleTerminatorContext)
}

func (s *PermissionBindRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PermissionBindRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PermissionBindRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterPermissionBindRule(s)
	}
}

func (s *PermissionBindRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitPermissionBindRule(s)
	}
}

func (p *ACIParser) PermissionBindRule() (localctx IPermissionBindRuleContext) {
	localctx = NewPermissionBindRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ACIParserRULE_permissionBindRule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Permission()
	}
	{
		p.SetState(152)
		p.bindRules(0)
	}
	{
		p.SetState(153)
		p.RuleTerminator()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPermissionContext is an interface to support dynamic dispatch.
type IPermissionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Disposition() IDispositionContext
	OpeningParenthesis() IOpeningParenthesisContext
	AllPrivilege() []IPrivilegeContext
	Privilege(i int) IPrivilegeContext
	ClosingParenthesis() IClosingParenthesisContext
	AllCommaDelimiter() []ICommaDelimiterContext
	CommaDelimiter(i int) ICommaDelimiterContext

	// IsPermissionContext differentiates from other interfaces.
	IsPermissionContext()
}

type PermissionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPermissionContext() *PermissionContext {
	var p = new(PermissionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_permission
	return p
}

func InitEmptyPermissionContext(p *PermissionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_permission
}

func (*PermissionContext) IsPermissionContext() {}

func NewPermissionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PermissionContext {
	var p = new(PermissionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_permission

	return p
}

func (s *PermissionContext) GetParser() antlr.Parser { return s.parser }

func (s *PermissionContext) Disposition() IDispositionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDispositionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDispositionContext)
}

func (s *PermissionContext) OpeningParenthesis() IOpeningParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpeningParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpeningParenthesisContext)
}

func (s *PermissionContext) AllPrivilege() []IPrivilegeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrivilegeContext); ok {
			len++
		}
	}

	tst := make([]IPrivilegeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrivilegeContext); ok {
			tst[i] = t.(IPrivilegeContext)
			i++
		}
	}

	return tst
}

func (s *PermissionContext) Privilege(i int) IPrivilegeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrivilegeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrivilegeContext)
}

func (s *PermissionContext) ClosingParenthesis() IClosingParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosingParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosingParenthesisContext)
}

func (s *PermissionContext) AllCommaDelimiter() []ICommaDelimiterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommaDelimiterContext); ok {
			len++
		}
	}

	tst := make([]ICommaDelimiterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommaDelimiterContext); ok {
			tst[i] = t.(ICommaDelimiterContext)
			i++
		}
	}

	return tst
}

func (s *PermissionContext) CommaDelimiter(i int) ICommaDelimiterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommaDelimiterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommaDelimiterContext)
}

func (s *PermissionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PermissionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PermissionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterPermission(s)
	}
}

func (s *PermissionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitPermission(s)
	}
}

func (p *ACIParser) Permission() (localctx IPermissionContext) {
	localctx = NewPermissionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ACIParserRULE_permission)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Disposition()
	}
	{
		p.SetState(156)
		p.OpeningParenthesis()
	}
	{
		p.SetState(157)
		p.Privilege()
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ACIParserCOMMA {
		{
			p.SetState(158)
			p.CommaDelimiter()
		}
		{
			p.SetState(159)
			p.Privilege()
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(166)
		p.ClosingParenthesis()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDispositionContext is an interface to support dynamic dispatch.
type IDispositionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GrantedPermission() IGrantedPermissionContext
	WithheldPermission() IWithheldPermissionContext

	// IsDispositionContext differentiates from other interfaces.
	IsDispositionContext()
}

type DispositionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDispositionContext() *DispositionContext {
	var p = new(DispositionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_disposition
	return p
}

func InitEmptyDispositionContext(p *DispositionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_disposition
}

func (*DispositionContext) IsDispositionContext() {}

func NewDispositionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DispositionContext {
	var p = new(DispositionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_disposition

	return p
}

func (s *DispositionContext) GetParser() antlr.Parser { return s.parser }

func (s *DispositionContext) GrantedPermission() IGrantedPermissionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGrantedPermissionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGrantedPermissionContext)
}

func (s *DispositionContext) WithheldPermission() IWithheldPermissionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWithheldPermissionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWithheldPermissionContext)
}

func (s *DispositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DispositionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DispositionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterDisposition(s)
	}
}

func (s *DispositionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitDisposition(s)
	}
}

func (p *ACIParser) Disposition() (localctx IDispositionContext) {
	localctx = NewDispositionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ACIParserRULE_disposition)
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserALLOW:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.GrantedPermission()
		}

	case ACIParserDENY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(169)
			p.WithheldPermission()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGrantedPermissionContext is an interface to support dynamic dispatch.
type IGrantedPermissionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALLOW() antlr.TerminalNode

	// IsGrantedPermissionContext differentiates from other interfaces.
	IsGrantedPermissionContext()
}

type GrantedPermissionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGrantedPermissionContext() *GrantedPermissionContext {
	var p = new(GrantedPermissionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_grantedPermission
	return p
}

func InitEmptyGrantedPermissionContext(p *GrantedPermissionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_grantedPermission
}

func (*GrantedPermissionContext) IsGrantedPermissionContext() {}

func NewGrantedPermissionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GrantedPermissionContext {
	var p = new(GrantedPermissionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_grantedPermission

	return p
}

func (s *GrantedPermissionContext) GetParser() antlr.Parser { return s.parser }

func (s *GrantedPermissionContext) ALLOW() antlr.TerminalNode {
	return s.GetToken(ACIParserALLOW, 0)
}

func (s *GrantedPermissionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GrantedPermissionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GrantedPermissionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterGrantedPermission(s)
	}
}

func (s *GrantedPermissionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitGrantedPermission(s)
	}
}

func (p *ACIParser) GrantedPermission() (localctx IGrantedPermissionContext) {
	localctx = NewGrantedPermissionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ACIParserRULE_grantedPermission)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(ACIParserALLOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWithheldPermissionContext is an interface to support dynamic dispatch.
type IWithheldPermissionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DENY() antlr.TerminalNode

	// IsWithheldPermissionContext differentiates from other interfaces.
	IsWithheldPermissionContext()
}

type WithheldPermissionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWithheldPermissionContext() *WithheldPermissionContext {
	var p = new(WithheldPermissionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_withheldPermission
	return p
}

func InitEmptyWithheldPermissionContext(p *WithheldPermissionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_withheldPermission
}

func (*WithheldPermissionContext) IsWithheldPermissionContext() {}

func NewWithheldPermissionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithheldPermissionContext {
	var p = new(WithheldPermissionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_withheldPermission

	return p
}

func (s *WithheldPermissionContext) GetParser() antlr.Parser { return s.parser }

func (s *WithheldPermissionContext) DENY() antlr.TerminalNode {
	return s.GetToken(ACIParserDENY, 0)
}

func (s *WithheldPermissionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithheldPermissionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WithheldPermissionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterWithheldPermission(s)
	}
}

func (s *WithheldPermissionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitWithheldPermission(s)
	}
}

func (p *ACIParser) WithheldPermission() (localctx IWithheldPermissionContext) {
	localctx = NewWithheldPermissionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ACIParserRULE_withheldPermission)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(ACIParserDENY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrivilegeContext is an interface to support dynamic dispatch.
type IPrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ReadPrivilege() IReadPrivilegeContext
	WritePrivilege() IWritePrivilegeContext
	SelfWritePrivilege() ISelfWritePrivilegeContext
	ComparePrivilege() IComparePrivilegeContext
	SearchPrivilege() ISearchPrivilegeContext
	ProxyPrivilege() IProxyPrivilegeContext
	AddPrivilege() IAddPrivilegeContext
	DeletePrivilege() IDeletePrivilegeContext
	ImportPrivilege() IImportPrivilegeContext
	ExportPrivilege() IExportPrivilegeContext
	AllPrivilege() IAllPrivilegeContext
	NoPrivilege() INoPrivilegeContext

	// IsPrivilegeContext differentiates from other interfaces.
	IsPrivilegeContext()
}

type PrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrivilegeContext() *PrivilegeContext {
	var p = new(PrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_privilege
	return p
}

func InitEmptyPrivilegeContext(p *PrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_privilege
}

func (*PrivilegeContext) IsPrivilegeContext() {}

func NewPrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrivilegeContext {
	var p = new(PrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_privilege

	return p
}

func (s *PrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrivilegeContext) ReadPrivilege() IReadPrivilegeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReadPrivilegeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReadPrivilegeContext)
}

func (s *PrivilegeContext) WritePrivilege() IWritePrivilegeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWritePrivilegeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWritePrivilegeContext)
}

func (s *PrivilegeContext) SelfWritePrivilege() ISelfWritePrivilegeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelfWritePrivilegeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelfWritePrivilegeContext)
}

func (s *PrivilegeContext) ComparePrivilege() IComparePrivilegeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparePrivilegeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparePrivilegeContext)
}

func (s *PrivilegeContext) SearchPrivilege() ISearchPrivilegeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchPrivilegeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchPrivilegeContext)
}

func (s *PrivilegeContext) ProxyPrivilege() IProxyPrivilegeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProxyPrivilegeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProxyPrivilegeContext)
}

func (s *PrivilegeContext) AddPrivilege() IAddPrivilegeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddPrivilegeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddPrivilegeContext)
}

func (s *PrivilegeContext) DeletePrivilege() IDeletePrivilegeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeletePrivilegeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeletePrivilegeContext)
}

func (s *PrivilegeContext) ImportPrivilege() IImportPrivilegeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportPrivilegeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportPrivilegeContext)
}

func (s *PrivilegeContext) ExportPrivilege() IExportPrivilegeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExportPrivilegeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExportPrivilegeContext)
}

func (s *PrivilegeContext) AllPrivilege() IAllPrivilegeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllPrivilegeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllPrivilegeContext)
}

func (s *PrivilegeContext) NoPrivilege() INoPrivilegeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoPrivilegeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoPrivilegeContext)
}

func (s *PrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterPrivilege(s)
	}
}

func (s *PrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitPrivilege(s)
	}
}

func (p *ACIParser) Privilege() (localctx IPrivilegeContext) {
	localctx = NewPrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ACIParserRULE_privilege)
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserREAD_PRIV:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.ReadPrivilege()
		}

	case ACIParserWRITE_PRIV:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.WritePrivilege()
		}

	case ACIParserSLF_PRIV:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(178)
			p.SelfWritePrivilege()
		}

	case ACIParserCMP_PRIV:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(179)
			p.ComparePrivilege()
		}

	case ACIParserSRC_PRIV:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(180)
			p.SearchPrivilege()
		}

	case ACIParserPRX_PRIV:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(181)
			p.ProxyPrivilege()
		}

	case ACIParserADD_PRIV:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(182)
			p.AddPrivilege()
		}

	case ACIParserDEL_PRIV:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(183)
			p.DeletePrivilege()
		}

	case ACIParserIMP_PRIV:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(184)
			p.ImportPrivilege()
		}

	case ACIParserEXP_PRIV:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(185)
			p.ExportPrivilege()
		}

	case ACIParserALL_PRIV:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(186)
			p.AllPrivilege()
		}

	case ACIParserNO_PRIV:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(187)
			p.NoPrivilege()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReadPrivilegeContext is an interface to support dynamic dispatch.
type IReadPrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	READ_PRIV() antlr.TerminalNode

	// IsReadPrivilegeContext differentiates from other interfaces.
	IsReadPrivilegeContext()
}

type ReadPrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadPrivilegeContext() *ReadPrivilegeContext {
	var p = new(ReadPrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_readPrivilege
	return p
}

func InitEmptyReadPrivilegeContext(p *ReadPrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_readPrivilege
}

func (*ReadPrivilegeContext) IsReadPrivilegeContext() {}

func NewReadPrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadPrivilegeContext {
	var p = new(ReadPrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_readPrivilege

	return p
}

func (s *ReadPrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadPrivilegeContext) READ_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserREAD_PRIV, 0)
}

func (s *ReadPrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadPrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadPrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterReadPrivilege(s)
	}
}

func (s *ReadPrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitReadPrivilege(s)
	}
}

func (p *ACIParser) ReadPrivilege() (localctx IReadPrivilegeContext) {
	localctx = NewReadPrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ACIParserRULE_readPrivilege)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(ACIParserREAD_PRIV)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWritePrivilegeContext is an interface to support dynamic dispatch.
type IWritePrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WRITE_PRIV() antlr.TerminalNode

	// IsWritePrivilegeContext differentiates from other interfaces.
	IsWritePrivilegeContext()
}

type WritePrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWritePrivilegeContext() *WritePrivilegeContext {
	var p = new(WritePrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_writePrivilege
	return p
}

func InitEmptyWritePrivilegeContext(p *WritePrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_writePrivilege
}

func (*WritePrivilegeContext) IsWritePrivilegeContext() {}

func NewWritePrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WritePrivilegeContext {
	var p = new(WritePrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_writePrivilege

	return p
}

func (s *WritePrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *WritePrivilegeContext) WRITE_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserWRITE_PRIV, 0)
}

func (s *WritePrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WritePrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WritePrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterWritePrivilege(s)
	}
}

func (s *WritePrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitWritePrivilege(s)
	}
}

func (p *ACIParser) WritePrivilege() (localctx IWritePrivilegeContext) {
	localctx = NewWritePrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ACIParserRULE_writePrivilege)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(ACIParserWRITE_PRIV)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelfWritePrivilegeContext is an interface to support dynamic dispatch.
type ISelfWritePrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SLF_PRIV() antlr.TerminalNode

	// IsSelfWritePrivilegeContext differentiates from other interfaces.
	IsSelfWritePrivilegeContext()
}

type SelfWritePrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelfWritePrivilegeContext() *SelfWritePrivilegeContext {
	var p = new(SelfWritePrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_selfWritePrivilege
	return p
}

func InitEmptySelfWritePrivilegeContext(p *SelfWritePrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_selfWritePrivilege
}

func (*SelfWritePrivilegeContext) IsSelfWritePrivilegeContext() {}

func NewSelfWritePrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelfWritePrivilegeContext {
	var p = new(SelfWritePrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_selfWritePrivilege

	return p
}

func (s *SelfWritePrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *SelfWritePrivilegeContext) SLF_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserSLF_PRIV, 0)
}

func (s *SelfWritePrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfWritePrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelfWritePrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterSelfWritePrivilege(s)
	}
}

func (s *SelfWritePrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitSelfWritePrivilege(s)
	}
}

func (p *ACIParser) SelfWritePrivilege() (localctx ISelfWritePrivilegeContext) {
	localctx = NewSelfWritePrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ACIParserRULE_selfWritePrivilege)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(ACIParserSLF_PRIV)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparePrivilegeContext is an interface to support dynamic dispatch.
type IComparePrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CMP_PRIV() antlr.TerminalNode

	// IsComparePrivilegeContext differentiates from other interfaces.
	IsComparePrivilegeContext()
}

type ComparePrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparePrivilegeContext() *ComparePrivilegeContext {
	var p = new(ComparePrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_comparePrivilege
	return p
}

func InitEmptyComparePrivilegeContext(p *ComparePrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_comparePrivilege
}

func (*ComparePrivilegeContext) IsComparePrivilegeContext() {}

func NewComparePrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparePrivilegeContext {
	var p = new(ComparePrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_comparePrivilege

	return p
}

func (s *ComparePrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparePrivilegeContext) CMP_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserCMP_PRIV, 0)
}

func (s *ComparePrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparePrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparePrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterComparePrivilege(s)
	}
}

func (s *ComparePrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitComparePrivilege(s)
	}
}

func (p *ACIParser) ComparePrivilege() (localctx IComparePrivilegeContext) {
	localctx = NewComparePrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ACIParserRULE_comparePrivilege)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(ACIParserCMP_PRIV)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISearchPrivilegeContext is an interface to support dynamic dispatch.
type ISearchPrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SRC_PRIV() antlr.TerminalNode

	// IsSearchPrivilegeContext differentiates from other interfaces.
	IsSearchPrivilegeContext()
}

type SearchPrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchPrivilegeContext() *SearchPrivilegeContext {
	var p = new(SearchPrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_searchPrivilege
	return p
}

func InitEmptySearchPrivilegeContext(p *SearchPrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_searchPrivilege
}

func (*SearchPrivilegeContext) IsSearchPrivilegeContext() {}

func NewSearchPrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchPrivilegeContext {
	var p = new(SearchPrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_searchPrivilege

	return p
}

func (s *SearchPrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchPrivilegeContext) SRC_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserSRC_PRIV, 0)
}

func (s *SearchPrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchPrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchPrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterSearchPrivilege(s)
	}
}

func (s *SearchPrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitSearchPrivilege(s)
	}
}

func (p *ACIParser) SearchPrivilege() (localctx ISearchPrivilegeContext) {
	localctx = NewSearchPrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ACIParserRULE_searchPrivilege)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(ACIParserSRC_PRIV)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProxyPrivilegeContext is an interface to support dynamic dispatch.
type IProxyPrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRX_PRIV() antlr.TerminalNode

	// IsProxyPrivilegeContext differentiates from other interfaces.
	IsProxyPrivilegeContext()
}

type ProxyPrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProxyPrivilegeContext() *ProxyPrivilegeContext {
	var p = new(ProxyPrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_proxyPrivilege
	return p
}

func InitEmptyProxyPrivilegeContext(p *ProxyPrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_proxyPrivilege
}

func (*ProxyPrivilegeContext) IsProxyPrivilegeContext() {}

func NewProxyPrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProxyPrivilegeContext {
	var p = new(ProxyPrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_proxyPrivilege

	return p
}

func (s *ProxyPrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *ProxyPrivilegeContext) PRX_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserPRX_PRIV, 0)
}

func (s *ProxyPrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProxyPrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProxyPrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterProxyPrivilege(s)
	}
}

func (s *ProxyPrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitProxyPrivilege(s)
	}
}

func (p *ACIParser) ProxyPrivilege() (localctx IProxyPrivilegeContext) {
	localctx = NewProxyPrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ACIParserRULE_proxyPrivilege)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(ACIParserPRX_PRIV)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddPrivilegeContext is an interface to support dynamic dispatch.
type IAddPrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD_PRIV() antlr.TerminalNode

	// IsAddPrivilegeContext differentiates from other interfaces.
	IsAddPrivilegeContext()
}

type AddPrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddPrivilegeContext() *AddPrivilegeContext {
	var p = new(AddPrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_addPrivilege
	return p
}

func InitEmptyAddPrivilegeContext(p *AddPrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_addPrivilege
}

func (*AddPrivilegeContext) IsAddPrivilegeContext() {}

func NewAddPrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddPrivilegeContext {
	var p = new(AddPrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_addPrivilege

	return p
}

func (s *AddPrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *AddPrivilegeContext) ADD_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserADD_PRIV, 0)
}

func (s *AddPrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddPrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddPrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAddPrivilege(s)
	}
}

func (s *AddPrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAddPrivilege(s)
	}
}

func (p *ACIParser) AddPrivilege() (localctx IAddPrivilegeContext) {
	localctx = NewAddPrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ACIParserRULE_addPrivilege)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(ACIParserADD_PRIV)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeletePrivilegeContext is an interface to support dynamic dispatch.
type IDeletePrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEL_PRIV() antlr.TerminalNode

	// IsDeletePrivilegeContext differentiates from other interfaces.
	IsDeletePrivilegeContext()
}

type DeletePrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeletePrivilegeContext() *DeletePrivilegeContext {
	var p = new(DeletePrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_deletePrivilege
	return p
}

func InitEmptyDeletePrivilegeContext(p *DeletePrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_deletePrivilege
}

func (*DeletePrivilegeContext) IsDeletePrivilegeContext() {}

func NewDeletePrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeletePrivilegeContext {
	var p = new(DeletePrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_deletePrivilege

	return p
}

func (s *DeletePrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *DeletePrivilegeContext) DEL_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserDEL_PRIV, 0)
}

func (s *DeletePrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeletePrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeletePrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterDeletePrivilege(s)
	}
}

func (s *DeletePrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitDeletePrivilege(s)
	}
}

func (p *ACIParser) DeletePrivilege() (localctx IDeletePrivilegeContext) {
	localctx = NewDeletePrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ACIParserRULE_deletePrivilege)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(ACIParserDEL_PRIV)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportPrivilegeContext is an interface to support dynamic dispatch.
type IImportPrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMP_PRIV() antlr.TerminalNode

	// IsImportPrivilegeContext differentiates from other interfaces.
	IsImportPrivilegeContext()
}

type ImportPrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportPrivilegeContext() *ImportPrivilegeContext {
	var p = new(ImportPrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_importPrivilege
	return p
}

func InitEmptyImportPrivilegeContext(p *ImportPrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_importPrivilege
}

func (*ImportPrivilegeContext) IsImportPrivilegeContext() {}

func NewImportPrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportPrivilegeContext {
	var p = new(ImportPrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_importPrivilege

	return p
}

func (s *ImportPrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportPrivilegeContext) IMP_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserIMP_PRIV, 0)
}

func (s *ImportPrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportPrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportPrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterImportPrivilege(s)
	}
}

func (s *ImportPrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitImportPrivilege(s)
	}
}

func (p *ACIParser) ImportPrivilege() (localctx IImportPrivilegeContext) {
	localctx = NewImportPrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ACIParserRULE_importPrivilege)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(ACIParserIMP_PRIV)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExportPrivilegeContext is an interface to support dynamic dispatch.
type IExportPrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXP_PRIV() antlr.TerminalNode

	// IsExportPrivilegeContext differentiates from other interfaces.
	IsExportPrivilegeContext()
}

type ExportPrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExportPrivilegeContext() *ExportPrivilegeContext {
	var p = new(ExportPrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_exportPrivilege
	return p
}

func InitEmptyExportPrivilegeContext(p *ExportPrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_exportPrivilege
}

func (*ExportPrivilegeContext) IsExportPrivilegeContext() {}

func NewExportPrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportPrivilegeContext {
	var p = new(ExportPrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_exportPrivilege

	return p
}

func (s *ExportPrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportPrivilegeContext) EXP_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserEXP_PRIV, 0)
}

func (s *ExportPrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportPrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExportPrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterExportPrivilege(s)
	}
}

func (s *ExportPrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitExportPrivilege(s)
	}
}

func (p *ACIParser) ExportPrivilege() (localctx IExportPrivilegeContext) {
	localctx = NewExportPrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ACIParserRULE_exportPrivilege)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(ACIParserEXP_PRIV)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAllPrivilegeContext is an interface to support dynamic dispatch.
type IAllPrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALL_PRIV() antlr.TerminalNode

	// IsAllPrivilegeContext differentiates from other interfaces.
	IsAllPrivilegeContext()
}

type AllPrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllPrivilegeContext() *AllPrivilegeContext {
	var p = new(AllPrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_allPrivilege
	return p
}

func InitEmptyAllPrivilegeContext(p *AllPrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_allPrivilege
}

func (*AllPrivilegeContext) IsAllPrivilegeContext() {}

func NewAllPrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllPrivilegeContext {
	var p = new(AllPrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_allPrivilege

	return p
}

func (s *AllPrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *AllPrivilegeContext) ALL_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserALL_PRIV, 0)
}

func (s *AllPrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllPrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllPrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAllPrivilege(s)
	}
}

func (s *AllPrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAllPrivilege(s)
	}
}

func (p *ACIParser) AllPrivilege() (localctx IAllPrivilegeContext) {
	localctx = NewAllPrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ACIParserRULE_allPrivilege)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(ACIParserALL_PRIV)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INoPrivilegeContext is an interface to support dynamic dispatch.
type INoPrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NO_PRIV() antlr.TerminalNode

	// IsNoPrivilegeContext differentiates from other interfaces.
	IsNoPrivilegeContext()
}

type NoPrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoPrivilegeContext() *NoPrivilegeContext {
	var p = new(NoPrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_noPrivilege
	return p
}

func InitEmptyNoPrivilegeContext(p *NoPrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_noPrivilege
}

func (*NoPrivilegeContext) IsNoPrivilegeContext() {}

func NewNoPrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoPrivilegeContext {
	var p = new(NoPrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_noPrivilege

	return p
}

func (s *NoPrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *NoPrivilegeContext) NO_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserNO_PRIV, 0)
}

func (s *NoPrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoPrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoPrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterNoPrivilege(s)
	}
}

func (s *NoPrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitNoPrivilege(s)
	}
}

func (p *ACIParser) NoPrivilege() (localctx INoPrivilegeContext) {
	localctx = NewNoPrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ACIParserRULE_noPrivilege)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(ACIParserNO_PRIV)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetRuleContext is an interface to support dynamic dispatch.
type ITargetRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpeningParenthesis() IOpeningParenthesisContext
	TargetKeyword() ITargetKeywordContext
	TargetOperator() ITargetOperatorContext
	ExpressionValues() IExpressionValuesContext
	ClosingParenthesis() IClosingParenthesisContext

	// IsTargetRuleContext differentiates from other interfaces.
	IsTargetRuleContext()
}

type TargetRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetRuleContext() *TargetRuleContext {
	var p = new(TargetRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetRule
	return p
}

func InitEmptyTargetRuleContext(p *TargetRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetRule
}

func (*TargetRuleContext) IsTargetRuleContext() {}

func NewTargetRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetRuleContext {
	var p = new(TargetRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetRule

	return p
}

func (s *TargetRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetRuleContext) OpeningParenthesis() IOpeningParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpeningParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpeningParenthesisContext)
}

func (s *TargetRuleContext) TargetKeyword() ITargetKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetKeywordContext)
}

func (s *TargetRuleContext) TargetOperator() ITargetOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetOperatorContext)
}

func (s *TargetRuleContext) ExpressionValues() IExpressionValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionValuesContext)
}

func (s *TargetRuleContext) ClosingParenthesis() IClosingParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosingParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosingParenthesisContext)
}

func (s *TargetRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetRule(s)
	}
}

func (s *TargetRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetRule(s)
	}
}

func (p *ACIParser) TargetRule() (localctx ITargetRuleContext) {
	localctx = NewTargetRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ACIParserRULE_targetRule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.OpeningParenthesis()
	}
	{
		p.SetState(215)
		p.TargetKeyword()
	}
	{
		p.SetState(216)
		p.TargetOperator()
	}
	{
		p.SetState(217)
		p.ExpressionValues()
	}
	{
		p.SetState(218)
		p.ClosingParenthesis()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetOperatorContext is an interface to support dynamic dispatch.
type ITargetOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EqualTo() IEqualToContext
	NotEqualTo() INotEqualToContext

	// IsTargetOperatorContext differentiates from other interfaces.
	IsTargetOperatorContext()
}

type TargetOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetOperatorContext() *TargetOperatorContext {
	var p = new(TargetOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetOperator
	return p
}

func InitEmptyTargetOperatorContext(p *TargetOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetOperator
}

func (*TargetOperatorContext) IsTargetOperatorContext() {}

func NewTargetOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetOperatorContext {
	var p = new(TargetOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetOperator

	return p
}

func (s *TargetOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetOperatorContext) EqualTo() IEqualToContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualToContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualToContext)
}

func (s *TargetOperatorContext) NotEqualTo() INotEqualToContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotEqualToContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotEqualToContext)
}

func (s *TargetOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetOperator(s)
	}
}

func (s *TargetOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetOperator(s)
	}
}

func (p *ACIParser) TargetOperator() (localctx ITargetOperatorContext) {
	localctx = NewTargetOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ACIParserRULE_targetOperator)
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserEQ:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)
			p.EqualTo()
		}

	case ACIParserNE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(221)
			p.NotEqualTo()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetRulesContext is an interface to support dynamic dispatch.
type ITargetRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTargetRule() []ITargetRuleContext
	TargetRule(i int) ITargetRuleContext

	// IsTargetRulesContext differentiates from other interfaces.
	IsTargetRulesContext()
}

type TargetRulesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetRulesContext() *TargetRulesContext {
	var p = new(TargetRulesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetRules
	return p
}

func InitEmptyTargetRulesContext(p *TargetRulesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetRules
}

func (*TargetRulesContext) IsTargetRulesContext() {}

func NewTargetRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetRulesContext {
	var p = new(TargetRulesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetRules

	return p
}

func (s *TargetRulesContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetRulesContext) AllTargetRule() []ITargetRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITargetRuleContext); ok {
			len++
		}
	}

	tst := make([]ITargetRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITargetRuleContext); ok {
			tst[i] = t.(ITargetRuleContext)
			i++
		}
	}

	return tst
}

func (s *TargetRulesContext) TargetRule(i int) ITargetRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetRuleContext)
}

func (s *TargetRulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetRulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetRulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetRules(s)
	}
}

func (s *TargetRulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetRules(s)
	}
}

func (p *ACIParser) TargetRules() (localctx ITargetRulesContext) {
	localctx = NewTargetRulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ACIParserRULE_targetRules)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(224)
				p.TargetRule()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWordAndContext is an interface to support dynamic dispatch.
type IWordAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD_AND() antlr.TerminalNode

	// IsWordAndContext differentiates from other interfaces.
	IsWordAndContext()
}

type WordAndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordAndContext() *WordAndContext {
	var p = new(WordAndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_wordAnd
	return p
}

func InitEmptyWordAndContext(p *WordAndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_wordAnd
}

func (*WordAndContext) IsWordAndContext() {}

func NewWordAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordAndContext {
	var p = new(WordAndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_wordAnd

	return p
}

func (s *WordAndContext) GetParser() antlr.Parser { return s.parser }

func (s *WordAndContext) WORD_AND() antlr.TerminalNode {
	return s.GetToken(ACIParserWORD_AND, 0)
}

func (s *WordAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterWordAnd(s)
	}
}

func (s *WordAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitWordAnd(s)
	}
}

func (p *ACIParser) WordAnd() (localctx IWordAndContext) {
	localctx = NewWordAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ACIParserRULE_wordAnd)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(ACIParserWORD_AND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWordOrContext is an interface to support dynamic dispatch.
type IWordOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD_OR() antlr.TerminalNode

	// IsWordOrContext differentiates from other interfaces.
	IsWordOrContext()
}

type WordOrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordOrContext() *WordOrContext {
	var p = new(WordOrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_wordOr
	return p
}

func InitEmptyWordOrContext(p *WordOrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_wordOr
}

func (*WordOrContext) IsWordOrContext() {}

func NewWordOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordOrContext {
	var p = new(WordOrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_wordOr

	return p
}

func (s *WordOrContext) GetParser() antlr.Parser { return s.parser }

func (s *WordOrContext) WORD_OR() antlr.TerminalNode {
	return s.GetToken(ACIParserWORD_OR, 0)
}

func (s *WordOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterWordOr(s)
	}
}

func (s *WordOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitWordOr(s)
	}
}

func (p *ACIParser) WordOr() (localctx IWordOrContext) {
	localctx = NewWordOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ACIParserRULE_wordOr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(ACIParserWORD_OR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWordNotContext is an interface to support dynamic dispatch.
type IWordNotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD_NOT() antlr.TerminalNode

	// IsWordNotContext differentiates from other interfaces.
	IsWordNotContext()
}

type WordNotContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordNotContext() *WordNotContext {
	var p = new(WordNotContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_wordNot
	return p
}

func InitEmptyWordNotContext(p *WordNotContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_wordNot
}

func (*WordNotContext) IsWordNotContext() {}

func NewWordNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordNotContext {
	var p = new(WordNotContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_wordNot

	return p
}

func (s *WordNotContext) GetParser() antlr.Parser { return s.parser }

func (s *WordNotContext) WORD_NOT() antlr.TerminalNode {
	return s.GetToken(ACIParserWORD_NOT, 0)
}

func (s *WordNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordNotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterWordNot(s)
	}
}

func (s *WordNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitWordNot(s)
	}
}

func (p *ACIParser) WordNot() (localctx IWordNotContext) {
	localctx = NewWordNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ACIParserRULE_wordNot)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Match(ACIParserWORD_NOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindRulesContext is an interface to support dynamic dispatch.
type IBindRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpeningParenthesis() IOpeningParenthesisContext
	AllBindRules() []IBindRulesContext
	BindRules(i int) IBindRulesContext
	ClosingParenthesis() IClosingParenthesisContext
	BindRule() IBindRuleContext
	AllWordAnd() []IWordAndContext
	WordAnd(i int) IWordAndContext
	AllWordOr() []IWordOrContext
	WordOr(i int) IWordOrContext
	AllWordNot() []IWordNotContext
	WordNot(i int) IWordNotContext

	// IsBindRulesContext differentiates from other interfaces.
	IsBindRulesContext()
}

type BindRulesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRulesContext() *BindRulesContext {
	var p = new(BindRulesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRules
	return p
}

func InitEmptyBindRulesContext(p *BindRulesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRules
}

func (*BindRulesContext) IsBindRulesContext() {}

func NewBindRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRulesContext {
	var p = new(BindRulesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRules

	return p
}

func (s *BindRulesContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRulesContext) OpeningParenthesis() IOpeningParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpeningParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpeningParenthesisContext)
}

func (s *BindRulesContext) AllBindRules() []IBindRulesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBindRulesContext); ok {
			len++
		}
	}

	tst := make([]IBindRulesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBindRulesContext); ok {
			tst[i] = t.(IBindRulesContext)
			i++
		}
	}

	return tst
}

func (s *BindRulesContext) BindRules(i int) IBindRulesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRulesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRulesContext)
}

func (s *BindRulesContext) ClosingParenthesis() IClosingParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosingParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosingParenthesisContext)
}

func (s *BindRulesContext) BindRule() IBindRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleContext)
}

func (s *BindRulesContext) AllWordAnd() []IWordAndContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordAndContext); ok {
			len++
		}
	}

	tst := make([]IWordAndContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordAndContext); ok {
			tst[i] = t.(IWordAndContext)
			i++
		}
	}

	return tst
}

func (s *BindRulesContext) WordAnd(i int) IWordAndContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordAndContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordAndContext)
}

func (s *BindRulesContext) AllWordOr() []IWordOrContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordOrContext); ok {
			len++
		}
	}

	tst := make([]IWordOrContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordOrContext); ok {
			tst[i] = t.(IWordOrContext)
			i++
		}
	}

	return tst
}

func (s *BindRulesContext) WordOr(i int) IWordOrContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordOrContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordOrContext)
}

func (s *BindRulesContext) AllWordNot() []IWordNotContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordNotContext); ok {
			len++
		}
	}

	tst := make([]IWordNotContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordNotContext); ok {
			tst[i] = t.(IWordNotContext)
			i++
		}
	}

	return tst
}

func (s *BindRulesContext) WordNot(i int) IWordNotContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordNotContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordNotContext)
}

func (s *BindRulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindRulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindRules(s)
	}
}

func (s *BindRulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindRules(s)
	}
}

func (p *ACIParser) BindRules() (localctx IBindRulesContext) {
	return p.bindRules(0)
}

func (p *ACIParser) bindRules(_p int) (localctx IBindRulesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewBindRulesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBindRulesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, ACIParserRULE_bindRules, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserLPAREN:
		{
			p.SetState(236)
			p.OpeningParenthesis()
		}
		{
			p.SetState(237)
			p.bindRules(0)
		}
		{
			p.SetState(238)
			p.ClosingParenthesis()
		}

	case ACIParserBKW_UDN, ACIParserBKW_GDN, ACIParserBKW_RDN, ACIParserBKW_UAT, ACIParserBKW_GAT, ACIParserBKW_SSF, ACIParserBKW_IP, ACIParserBKW_DNS, ACIParserBKW_AM, ACIParserBKW_TOD, ACIParserBKW_DOW:
		{
			p.SetState(240)
			p.BindRule()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBindRulesContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ACIParserRULE_bindRules)
			p.SetState(243)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			p.SetState(251)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					p.SetState(247)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}

					switch p.GetTokenStream().LA(1) {
					case ACIParserWORD_AND:
						{
							p.SetState(244)
							p.WordAnd()
						}

					case ACIParserWORD_OR:
						{
							p.SetState(245)
							p.WordOr()
						}

					case ACIParserWORD_NOT:
						{
							p.SetState(246)
							p.WordNot()
						}

					default:
						p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
						goto errorExit
					}
					{
						p.SetState(249)
						p.bindRules(0)
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

				p.SetState(253)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		}
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindRuleContext is an interface to support dynamic dispatch.
type IBindRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BindKeyword() IBindKeywordContext
	BindOperator() IBindOperatorContext
	ExpressionValues() IExpressionValuesContext

	// IsBindRuleContext differentiates from other interfaces.
	IsBindRuleContext()
}

type BindRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleContext() *BindRuleContext {
	var p = new(BindRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRule
	return p
}

func InitEmptyBindRuleContext(p *BindRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRule
}

func (*BindRuleContext) IsBindRuleContext() {}

func NewBindRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleContext {
	var p = new(BindRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRule

	return p
}

func (s *BindRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleContext) BindKeyword() IBindKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindKeywordContext)
}

func (s *BindRuleContext) BindOperator() IBindOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindOperatorContext)
}

func (s *BindRuleContext) ExpressionValues() IExpressionValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionValuesContext)
}

func (s *BindRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindRule(s)
	}
}

func (s *BindRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindRule(s)
	}
}

func (p *ACIParser) BindRule() (localctx IBindRuleContext) {
	localctx = NewBindRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ACIParserRULE_bindRule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.BindKeyword()
	}
	{
		p.SetState(261)
		p.BindOperator()
	}
	{
		p.SetState(262)
		p.ExpressionValues()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionValuesContext is an interface to support dynamic dispatch.
type IExpressionValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTED_STRING() []antlr.TerminalNode
	QUOTED_STRING(i int) antlr.TerminalNode
	AllSymbolicOr() []ISymbolicOrContext
	SymbolicOr(i int) ISymbolicOrContext

	// IsExpressionValuesContext differentiates from other interfaces.
	IsExpressionValuesContext()
}

type ExpressionValuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionValuesContext() *ExpressionValuesContext {
	var p = new(ExpressionValuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_expressionValues
	return p
}

func InitEmptyExpressionValuesContext(p *ExpressionValuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_expressionValues
}

func (*ExpressionValuesContext) IsExpressionValuesContext() {}

func NewExpressionValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionValuesContext {
	var p = new(ExpressionValuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_expressionValues

	return p
}

func (s *ExpressionValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionValuesContext) AllQUOTED_STRING() []antlr.TerminalNode {
	return s.GetTokens(ACIParserQUOTED_STRING)
}

func (s *ExpressionValuesContext) QUOTED_STRING(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserQUOTED_STRING, i)
}

func (s *ExpressionValuesContext) AllSymbolicOr() []ISymbolicOrContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISymbolicOrContext); ok {
			len++
		}
	}

	tst := make([]ISymbolicOrContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISymbolicOrContext); ok {
			tst[i] = t.(ISymbolicOrContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionValuesContext) SymbolicOr(i int) ISymbolicOrContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISymbolicOrContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISymbolicOrContext)
}

func (s *ExpressionValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterExpressionValues(s)
	}
}

func (s *ExpressionValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitExpressionValues(s)
	}
}

func (p *ACIParser) ExpressionValues() (localctx IExpressionValuesContext) {
	localctx = NewExpressionValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ACIParserRULE_expressionValues)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Match(ACIParserQUOTED_STRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(265)
				p.SymbolicOr()
			}
			{
				p.SetState(266)
				p.Match(ACIParserQUOTED_STRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISymbolicOrContext is an interface to support dynamic dispatch.
type ISymbolicOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYMBOLIC_OR() antlr.TerminalNode

	// IsSymbolicOrContext differentiates from other interfaces.
	IsSymbolicOrContext()
}

type SymbolicOrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolicOrContext() *SymbolicOrContext {
	var p = new(SymbolicOrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_symbolicOr
	return p
}

func InitEmptySymbolicOrContext(p *SymbolicOrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_symbolicOr
}

func (*SymbolicOrContext) IsSymbolicOrContext() {}

func NewSymbolicOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolicOrContext {
	var p = new(SymbolicOrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_symbolicOr

	return p
}

func (s *SymbolicOrContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolicOrContext) SYMBOLIC_OR() antlr.TerminalNode {
	return s.GetToken(ACIParserSYMBOLIC_OR, 0)
}

func (s *SymbolicOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolicOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolicOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterSymbolicOr(s)
	}
}

func (s *SymbolicOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitSymbolicOr(s)
	}
}

func (p *ACIParser) SymbolicOr() (localctx ISymbolicOrContext) {
	localctx = NewSymbolicOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ACIParserRULE_symbolicOr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(ACIParserSYMBOLIC_OR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindOperatorContext is an interface to support dynamic dispatch.
type IBindOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EqualTo() IEqualToContext
	NotEqualTo() INotEqualToContext
	LessThan() ILessThanContext
	LessThanOrEqual() ILessThanOrEqualContext
	GreaterThan() IGreaterThanContext
	GreaterThanOrEqual() IGreaterThanOrEqualContext

	// IsBindOperatorContext differentiates from other interfaces.
	IsBindOperatorContext()
}

type BindOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindOperatorContext() *BindOperatorContext {
	var p = new(BindOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindOperator
	return p
}

func InitEmptyBindOperatorContext(p *BindOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindOperator
}

func (*BindOperatorContext) IsBindOperatorContext() {}

func NewBindOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindOperatorContext {
	var p = new(BindOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindOperator

	return p
}

func (s *BindOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BindOperatorContext) EqualTo() IEqualToContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualToContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualToContext)
}

func (s *BindOperatorContext) NotEqualTo() INotEqualToContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotEqualToContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotEqualToContext)
}

func (s *BindOperatorContext) LessThan() ILessThanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILessThanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILessThanContext)
}

func (s *BindOperatorContext) LessThanOrEqual() ILessThanOrEqualContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILessThanOrEqualContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILessThanOrEqualContext)
}

func (s *BindOperatorContext) GreaterThan() IGreaterThanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGreaterThanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGreaterThanContext)
}

func (s *BindOperatorContext) GreaterThanOrEqual() IGreaterThanOrEqualContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGreaterThanOrEqualContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGreaterThanOrEqualContext)
}

func (s *BindOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindOperator(s)
	}
}

func (s *BindOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindOperator(s)
	}
}

func (p *ACIParser) BindOperator() (localctx IBindOperatorContext) {
	localctx = NewBindOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ACIParserRULE_bindOperator)
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserEQ:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.EqualTo()
		}

	case ACIParserNE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)
			p.NotEqualTo()
		}

	case ACIParserLT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(277)
			p.LessThan()
		}

	case ACIParserLE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(278)
			p.LessThanOrEqual()
		}

	case ACIParserGT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(279)
			p.GreaterThan()
		}

	case ACIParserGE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(280)
			p.GreaterThanOrEqual()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILessThanContext is an interface to support dynamic dispatch.
type ILessThanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LT() antlr.TerminalNode

	// IsLessThanContext differentiates from other interfaces.
	IsLessThanContext()
}

type LessThanContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLessThanContext() *LessThanContext {
	var p = new(LessThanContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lessThan
	return p
}

func InitEmptyLessThanContext(p *LessThanContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lessThan
}

func (*LessThanContext) IsLessThanContext() {}

func NewLessThanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LessThanContext {
	var p = new(LessThanContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_lessThan

	return p
}

func (s *LessThanContext) GetParser() antlr.Parser { return s.parser }

func (s *LessThanContext) LT() antlr.TerminalNode {
	return s.GetToken(ACIParserLT, 0)
}

func (s *LessThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LessThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterLessThan(s)
	}
}

func (s *LessThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitLessThan(s)
	}
}

func (p *ACIParser) LessThan() (localctx ILessThanContext) {
	localctx = NewLessThanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ACIParserRULE_lessThan)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Match(ACIParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILessThanOrEqualContext is an interface to support dynamic dispatch.
type ILessThanOrEqualContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LE() antlr.TerminalNode

	// IsLessThanOrEqualContext differentiates from other interfaces.
	IsLessThanOrEqualContext()
}

type LessThanOrEqualContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLessThanOrEqualContext() *LessThanOrEqualContext {
	var p = new(LessThanOrEqualContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lessThanOrEqual
	return p
}

func InitEmptyLessThanOrEqualContext(p *LessThanOrEqualContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lessThanOrEqual
}

func (*LessThanOrEqualContext) IsLessThanOrEqualContext() {}

func NewLessThanOrEqualContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LessThanOrEqualContext {
	var p = new(LessThanOrEqualContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_lessThanOrEqual

	return p
}

func (s *LessThanOrEqualContext) GetParser() antlr.Parser { return s.parser }

func (s *LessThanOrEqualContext) LE() antlr.TerminalNode {
	return s.GetToken(ACIParserLE, 0)
}

func (s *LessThanOrEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanOrEqualContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LessThanOrEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterLessThanOrEqual(s)
	}
}

func (s *LessThanOrEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitLessThanOrEqual(s)
	}
}

func (p *ACIParser) LessThanOrEqual() (localctx ILessThanOrEqualContext) {
	localctx = NewLessThanOrEqualContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ACIParserRULE_lessThanOrEqual)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Match(ACIParserLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGreaterThanContext is an interface to support dynamic dispatch.
type IGreaterThanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GT() antlr.TerminalNode

	// IsGreaterThanContext differentiates from other interfaces.
	IsGreaterThanContext()
}

type GreaterThanContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGreaterThanContext() *GreaterThanContext {
	var p = new(GreaterThanContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_greaterThan
	return p
}

func InitEmptyGreaterThanContext(p *GreaterThanContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_greaterThan
}

func (*GreaterThanContext) IsGreaterThanContext() {}

func NewGreaterThanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GreaterThanContext {
	var p = new(GreaterThanContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_greaterThan

	return p
}

func (s *GreaterThanContext) GetParser() antlr.Parser { return s.parser }

func (s *GreaterThanContext) GT() antlr.TerminalNode {
	return s.GetToken(ACIParserGT, 0)
}

func (s *GreaterThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GreaterThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterGreaterThan(s)
	}
}

func (s *GreaterThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitGreaterThan(s)
	}
}

func (p *ACIParser) GreaterThan() (localctx IGreaterThanContext) {
	localctx = NewGreaterThanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ACIParserRULE_greaterThan)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.Match(ACIParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGreaterThanOrEqualContext is an interface to support dynamic dispatch.
type IGreaterThanOrEqualContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GE() antlr.TerminalNode

	// IsGreaterThanOrEqualContext differentiates from other interfaces.
	IsGreaterThanOrEqualContext()
}

type GreaterThanOrEqualContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGreaterThanOrEqualContext() *GreaterThanOrEqualContext {
	var p = new(GreaterThanOrEqualContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_greaterThanOrEqual
	return p
}

func InitEmptyGreaterThanOrEqualContext(p *GreaterThanOrEqualContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_greaterThanOrEqual
}

func (*GreaterThanOrEqualContext) IsGreaterThanOrEqualContext() {}

func NewGreaterThanOrEqualContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GreaterThanOrEqualContext {
	var p = new(GreaterThanOrEqualContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_greaterThanOrEqual

	return p
}

func (s *GreaterThanOrEqualContext) GetParser() antlr.Parser { return s.parser }

func (s *GreaterThanOrEqualContext) GE() antlr.TerminalNode {
	return s.GetToken(ACIParserGE, 0)
}

func (s *GreaterThanOrEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanOrEqualContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GreaterThanOrEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterGreaterThanOrEqual(s)
	}
}

func (s *GreaterThanOrEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitGreaterThanOrEqual(s)
	}
}

func (p *ACIParser) GreaterThanOrEqual() (localctx IGreaterThanOrEqualContext) {
	localctx = NewGreaterThanOrEqualContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ACIParserRULE_greaterThanOrEqual)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Match(ACIParserGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualToContext is an interface to support dynamic dispatch.
type IEqualToContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode

	// IsEqualToContext differentiates from other interfaces.
	IsEqualToContext()
}

type EqualToContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualToContext() *EqualToContext {
	var p = new(EqualToContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_equalTo
	return p
}

func InitEmptyEqualToContext(p *EqualToContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_equalTo
}

func (*EqualToContext) IsEqualToContext() {}

func NewEqualToContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualToContext {
	var p = new(EqualToContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_equalTo

	return p
}

func (s *EqualToContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualToContext) EQ() antlr.TerminalNode {
	return s.GetToken(ACIParserEQ, 0)
}

func (s *EqualToContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualToContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualToContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterEqualTo(s)
	}
}

func (s *EqualToContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitEqualTo(s)
	}
}

func (p *ACIParser) EqualTo() (localctx IEqualToContext) {
	localctx = NewEqualToContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ACIParserRULE_equalTo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(ACIParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INotEqualToContext is an interface to support dynamic dispatch.
type INotEqualToContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NE() antlr.TerminalNode

	// IsNotEqualToContext differentiates from other interfaces.
	IsNotEqualToContext()
}

type NotEqualToContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotEqualToContext() *NotEqualToContext {
	var p = new(NotEqualToContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_notEqualTo
	return p
}

func InitEmptyNotEqualToContext(p *NotEqualToContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_notEqualTo
}

func (*NotEqualToContext) IsNotEqualToContext() {}

func NewNotEqualToContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotEqualToContext {
	var p = new(NotEqualToContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_notEqualTo

	return p
}

func (s *NotEqualToContext) GetParser() antlr.Parser { return s.parser }

func (s *NotEqualToContext) NE() antlr.TerminalNode {
	return s.GetToken(ACIParserNE, 0)
}

func (s *NotEqualToContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualToContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotEqualToContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterNotEqualTo(s)
	}
}

func (s *NotEqualToContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitNotEqualTo(s)
	}
}

func (p *ACIParser) NotEqualTo() (localctx INotEqualToContext) {
	localctx = NewNotEqualToContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ACIParserRULE_notEqualTo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(ACIParserNE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindKeywordContext is an interface to support dynamic dispatch.
type IBindKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BindUserDistinguishedName() IBindUserDistinguishedNameContext
	BindGroupDistinguishedName() IBindGroupDistinguishedNameContext
	BindRoleDistinguishedName() IBindRoleDistinguishedNameContext
	BindUserAttributes() IBindUserAttributesContext
	BindGroupAttributes() IBindGroupAttributesContext
	BindSecurityStrengthFactor() IBindSecurityStrengthFactorContext
	BindIPAddress() IBindIPAddressContext
	BindDNSName() IBindDNSNameContext
	BindTimeOfDay() IBindTimeOfDayContext
	BindDayOfWeek() IBindDayOfWeekContext
	BindAuthenticationMethod() IBindAuthenticationMethodContext

	// IsBindKeywordContext differentiates from other interfaces.
	IsBindKeywordContext()
}

type BindKeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindKeywordContext() *BindKeywordContext {
	var p = new(BindKeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindKeyword
	return p
}

func InitEmptyBindKeywordContext(p *BindKeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindKeyword
}

func (*BindKeywordContext) IsBindKeywordContext() {}

func NewBindKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindKeywordContext {
	var p = new(BindKeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindKeyword

	return p
}

func (s *BindKeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *BindKeywordContext) BindUserDistinguishedName() IBindUserDistinguishedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindUserDistinguishedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindUserDistinguishedNameContext)
}

func (s *BindKeywordContext) BindGroupDistinguishedName() IBindGroupDistinguishedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindGroupDistinguishedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindGroupDistinguishedNameContext)
}

func (s *BindKeywordContext) BindRoleDistinguishedName() IBindRoleDistinguishedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRoleDistinguishedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRoleDistinguishedNameContext)
}

func (s *BindKeywordContext) BindUserAttributes() IBindUserAttributesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindUserAttributesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindUserAttributesContext)
}

func (s *BindKeywordContext) BindGroupAttributes() IBindGroupAttributesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindGroupAttributesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindGroupAttributesContext)
}

func (s *BindKeywordContext) BindSecurityStrengthFactor() IBindSecurityStrengthFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindSecurityStrengthFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindSecurityStrengthFactorContext)
}

func (s *BindKeywordContext) BindIPAddress() IBindIPAddressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindIPAddressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindIPAddressContext)
}

func (s *BindKeywordContext) BindDNSName() IBindDNSNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindDNSNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindDNSNameContext)
}

func (s *BindKeywordContext) BindTimeOfDay() IBindTimeOfDayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindTimeOfDayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindTimeOfDayContext)
}

func (s *BindKeywordContext) BindDayOfWeek() IBindDayOfWeekContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindDayOfWeekContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindDayOfWeekContext)
}

func (s *BindKeywordContext) BindAuthenticationMethod() IBindAuthenticationMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindAuthenticationMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindAuthenticationMethodContext)
}

func (s *BindKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindKeyword(s)
	}
}

func (s *BindKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindKeyword(s)
	}
}

func (p *ACIParser) BindKeyword() (localctx IBindKeywordContext) {
	localctx = NewBindKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ACIParserRULE_bindKeyword)
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserBKW_UDN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(295)
			p.BindUserDistinguishedName()
		}

	case ACIParserBKW_GDN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(296)
			p.BindGroupDistinguishedName()
		}

	case ACIParserBKW_RDN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(297)
			p.BindRoleDistinguishedName()
		}

	case ACIParserBKW_UAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(298)
			p.BindUserAttributes()
		}

	case ACIParserBKW_GAT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(299)
			p.BindGroupAttributes()
		}

	case ACIParserBKW_SSF:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(300)
			p.BindSecurityStrengthFactor()
		}

	case ACIParserBKW_IP:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(301)
			p.BindIPAddress()
		}

	case ACIParserBKW_DNS:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(302)
			p.BindDNSName()
		}

	case ACIParserBKW_TOD:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(303)
			p.BindTimeOfDay()
		}

	case ACIParserBKW_DOW:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(304)
			p.BindDayOfWeek()
		}

	case ACIParserBKW_AM:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(305)
			p.BindAuthenticationMethod()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindUserDistinguishedNameContext is an interface to support dynamic dispatch.
type IBindUserDistinguishedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BKW_UDN() antlr.TerminalNode

	// IsBindUserDistinguishedNameContext differentiates from other interfaces.
	IsBindUserDistinguishedNameContext()
}

type BindUserDistinguishedNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindUserDistinguishedNameContext() *BindUserDistinguishedNameContext {
	var p = new(BindUserDistinguishedNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindUserDistinguishedName
	return p
}

func InitEmptyBindUserDistinguishedNameContext(p *BindUserDistinguishedNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindUserDistinguishedName
}

func (*BindUserDistinguishedNameContext) IsBindUserDistinguishedNameContext() {}

func NewBindUserDistinguishedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindUserDistinguishedNameContext {
	var p = new(BindUserDistinguishedNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindUserDistinguishedName

	return p
}

func (s *BindUserDistinguishedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *BindUserDistinguishedNameContext) BKW_UDN() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_UDN, 0)
}

func (s *BindUserDistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindUserDistinguishedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindUserDistinguishedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindUserDistinguishedName(s)
	}
}

func (s *BindUserDistinguishedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindUserDistinguishedName(s)
	}
}

func (p *ACIParser) BindUserDistinguishedName() (localctx IBindUserDistinguishedNameContext) {
	localctx = NewBindUserDistinguishedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ACIParserRULE_bindUserDistinguishedName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(ACIParserBKW_UDN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindGroupDistinguishedNameContext is an interface to support dynamic dispatch.
type IBindGroupDistinguishedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BKW_GDN() antlr.TerminalNode

	// IsBindGroupDistinguishedNameContext differentiates from other interfaces.
	IsBindGroupDistinguishedNameContext()
}

type BindGroupDistinguishedNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindGroupDistinguishedNameContext() *BindGroupDistinguishedNameContext {
	var p = new(BindGroupDistinguishedNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindGroupDistinguishedName
	return p
}

func InitEmptyBindGroupDistinguishedNameContext(p *BindGroupDistinguishedNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindGroupDistinguishedName
}

func (*BindGroupDistinguishedNameContext) IsBindGroupDistinguishedNameContext() {}

func NewBindGroupDistinguishedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindGroupDistinguishedNameContext {
	var p = new(BindGroupDistinguishedNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindGroupDistinguishedName

	return p
}

func (s *BindGroupDistinguishedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *BindGroupDistinguishedNameContext) BKW_GDN() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_GDN, 0)
}

func (s *BindGroupDistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindGroupDistinguishedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindGroupDistinguishedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindGroupDistinguishedName(s)
	}
}

func (s *BindGroupDistinguishedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindGroupDistinguishedName(s)
	}
}

func (p *ACIParser) BindGroupDistinguishedName() (localctx IBindGroupDistinguishedNameContext) {
	localctx = NewBindGroupDistinguishedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ACIParserRULE_bindGroupDistinguishedName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(ACIParserBKW_GDN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindRoleDistinguishedNameContext is an interface to support dynamic dispatch.
type IBindRoleDistinguishedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BKW_RDN() antlr.TerminalNode

	// IsBindRoleDistinguishedNameContext differentiates from other interfaces.
	IsBindRoleDistinguishedNameContext()
}

type BindRoleDistinguishedNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRoleDistinguishedNameContext() *BindRoleDistinguishedNameContext {
	var p = new(BindRoleDistinguishedNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRoleDistinguishedName
	return p
}

func InitEmptyBindRoleDistinguishedNameContext(p *BindRoleDistinguishedNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRoleDistinguishedName
}

func (*BindRoleDistinguishedNameContext) IsBindRoleDistinguishedNameContext() {}

func NewBindRoleDistinguishedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRoleDistinguishedNameContext {
	var p = new(BindRoleDistinguishedNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRoleDistinguishedName

	return p
}

func (s *BindRoleDistinguishedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRoleDistinguishedNameContext) BKW_RDN() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_RDN, 0)
}

func (s *BindRoleDistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRoleDistinguishedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindRoleDistinguishedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindRoleDistinguishedName(s)
	}
}

func (s *BindRoleDistinguishedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindRoleDistinguishedName(s)
	}
}

func (p *ACIParser) BindRoleDistinguishedName() (localctx IBindRoleDistinguishedNameContext) {
	localctx = NewBindRoleDistinguishedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ACIParserRULE_bindRoleDistinguishedName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(ACIParserBKW_RDN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindUserAttributesContext is an interface to support dynamic dispatch.
type IBindUserAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BKW_UAT() antlr.TerminalNode

	// IsBindUserAttributesContext differentiates from other interfaces.
	IsBindUserAttributesContext()
}

type BindUserAttributesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindUserAttributesContext() *BindUserAttributesContext {
	var p = new(BindUserAttributesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindUserAttributes
	return p
}

func InitEmptyBindUserAttributesContext(p *BindUserAttributesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindUserAttributes
}

func (*BindUserAttributesContext) IsBindUserAttributesContext() {}

func NewBindUserAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindUserAttributesContext {
	var p = new(BindUserAttributesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindUserAttributes

	return p
}

func (s *BindUserAttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *BindUserAttributesContext) BKW_UAT() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_UAT, 0)
}

func (s *BindUserAttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindUserAttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindUserAttributesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindUserAttributes(s)
	}
}

func (s *BindUserAttributesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindUserAttributes(s)
	}
}

func (p *ACIParser) BindUserAttributes() (localctx IBindUserAttributesContext) {
	localctx = NewBindUserAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ACIParserRULE_bindUserAttributes)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(ACIParserBKW_UAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindGroupAttributesContext is an interface to support dynamic dispatch.
type IBindGroupAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BKW_GAT() antlr.TerminalNode

	// IsBindGroupAttributesContext differentiates from other interfaces.
	IsBindGroupAttributesContext()
}

type BindGroupAttributesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindGroupAttributesContext() *BindGroupAttributesContext {
	var p = new(BindGroupAttributesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindGroupAttributes
	return p
}

func InitEmptyBindGroupAttributesContext(p *BindGroupAttributesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindGroupAttributes
}

func (*BindGroupAttributesContext) IsBindGroupAttributesContext() {}

func NewBindGroupAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindGroupAttributesContext {
	var p = new(BindGroupAttributesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindGroupAttributes

	return p
}

func (s *BindGroupAttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *BindGroupAttributesContext) BKW_GAT() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_GAT, 0)
}

func (s *BindGroupAttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindGroupAttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindGroupAttributesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindGroupAttributes(s)
	}
}

func (s *BindGroupAttributesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindGroupAttributes(s)
	}
}

func (p *ACIParser) BindGroupAttributes() (localctx IBindGroupAttributesContext) {
	localctx = NewBindGroupAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ACIParserRULE_bindGroupAttributes)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.Match(ACIParserBKW_GAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindSecurityStrengthFactorContext is an interface to support dynamic dispatch.
type IBindSecurityStrengthFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BKW_SSF() antlr.TerminalNode

	// IsBindSecurityStrengthFactorContext differentiates from other interfaces.
	IsBindSecurityStrengthFactorContext()
}

type BindSecurityStrengthFactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindSecurityStrengthFactorContext() *BindSecurityStrengthFactorContext {
	var p = new(BindSecurityStrengthFactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindSecurityStrengthFactor
	return p
}

func InitEmptyBindSecurityStrengthFactorContext(p *BindSecurityStrengthFactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindSecurityStrengthFactor
}

func (*BindSecurityStrengthFactorContext) IsBindSecurityStrengthFactorContext() {}

func NewBindSecurityStrengthFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindSecurityStrengthFactorContext {
	var p = new(BindSecurityStrengthFactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindSecurityStrengthFactor

	return p
}

func (s *BindSecurityStrengthFactorContext) GetParser() antlr.Parser { return s.parser }

func (s *BindSecurityStrengthFactorContext) BKW_SSF() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_SSF, 0)
}

func (s *BindSecurityStrengthFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindSecurityStrengthFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindSecurityStrengthFactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindSecurityStrengthFactor(s)
	}
}

func (s *BindSecurityStrengthFactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindSecurityStrengthFactor(s)
	}
}

func (p *ACIParser) BindSecurityStrengthFactor() (localctx IBindSecurityStrengthFactorContext) {
	localctx = NewBindSecurityStrengthFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ACIParserRULE_bindSecurityStrengthFactor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.Match(ACIParserBKW_SSF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindIPAddressContext is an interface to support dynamic dispatch.
type IBindIPAddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BKW_IP() antlr.TerminalNode

	// IsBindIPAddressContext differentiates from other interfaces.
	IsBindIPAddressContext()
}

type BindIPAddressContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindIPAddressContext() *BindIPAddressContext {
	var p = new(BindIPAddressContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindIPAddress
	return p
}

func InitEmptyBindIPAddressContext(p *BindIPAddressContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindIPAddress
}

func (*BindIPAddressContext) IsBindIPAddressContext() {}

func NewBindIPAddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindIPAddressContext {
	var p = new(BindIPAddressContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindIPAddress

	return p
}

func (s *BindIPAddressContext) GetParser() antlr.Parser { return s.parser }

func (s *BindIPAddressContext) BKW_IP() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_IP, 0)
}

func (s *BindIPAddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindIPAddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindIPAddressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindIPAddress(s)
	}
}

func (s *BindIPAddressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindIPAddress(s)
	}
}

func (p *ACIParser) BindIPAddress() (localctx IBindIPAddressContext) {
	localctx = NewBindIPAddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ACIParserRULE_bindIPAddress)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(ACIParserBKW_IP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindDNSNameContext is an interface to support dynamic dispatch.
type IBindDNSNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BKW_DNS() antlr.TerminalNode

	// IsBindDNSNameContext differentiates from other interfaces.
	IsBindDNSNameContext()
}

type BindDNSNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindDNSNameContext() *BindDNSNameContext {
	var p = new(BindDNSNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindDNSName
	return p
}

func InitEmptyBindDNSNameContext(p *BindDNSNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindDNSName
}

func (*BindDNSNameContext) IsBindDNSNameContext() {}

func NewBindDNSNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindDNSNameContext {
	var p = new(BindDNSNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindDNSName

	return p
}

func (s *BindDNSNameContext) GetParser() antlr.Parser { return s.parser }

func (s *BindDNSNameContext) BKW_DNS() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_DNS, 0)
}

func (s *BindDNSNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindDNSNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindDNSNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindDNSName(s)
	}
}

func (s *BindDNSNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindDNSName(s)
	}
}

func (p *ACIParser) BindDNSName() (localctx IBindDNSNameContext) {
	localctx = NewBindDNSNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ACIParserRULE_bindDNSName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(ACIParserBKW_DNS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindTimeOfDayContext is an interface to support dynamic dispatch.
type IBindTimeOfDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BKW_TOD() antlr.TerminalNode

	// IsBindTimeOfDayContext differentiates from other interfaces.
	IsBindTimeOfDayContext()
}

type BindTimeOfDayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindTimeOfDayContext() *BindTimeOfDayContext {
	var p = new(BindTimeOfDayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindTimeOfDay
	return p
}

func InitEmptyBindTimeOfDayContext(p *BindTimeOfDayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindTimeOfDay
}

func (*BindTimeOfDayContext) IsBindTimeOfDayContext() {}

func NewBindTimeOfDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindTimeOfDayContext {
	var p = new(BindTimeOfDayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindTimeOfDay

	return p
}

func (s *BindTimeOfDayContext) GetParser() antlr.Parser { return s.parser }

func (s *BindTimeOfDayContext) BKW_TOD() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_TOD, 0)
}

func (s *BindTimeOfDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindTimeOfDayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindTimeOfDayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindTimeOfDay(s)
	}
}

func (s *BindTimeOfDayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindTimeOfDay(s)
	}
}

func (p *ACIParser) BindTimeOfDay() (localctx IBindTimeOfDayContext) {
	localctx = NewBindTimeOfDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ACIParserRULE_bindTimeOfDay)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.Match(ACIParserBKW_TOD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindDayOfWeekContext is an interface to support dynamic dispatch.
type IBindDayOfWeekContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BKW_DOW() antlr.TerminalNode

	// IsBindDayOfWeekContext differentiates from other interfaces.
	IsBindDayOfWeekContext()
}

type BindDayOfWeekContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindDayOfWeekContext() *BindDayOfWeekContext {
	var p = new(BindDayOfWeekContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindDayOfWeek
	return p
}

func InitEmptyBindDayOfWeekContext(p *BindDayOfWeekContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindDayOfWeek
}

func (*BindDayOfWeekContext) IsBindDayOfWeekContext() {}

func NewBindDayOfWeekContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindDayOfWeekContext {
	var p = new(BindDayOfWeekContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindDayOfWeek

	return p
}

func (s *BindDayOfWeekContext) GetParser() antlr.Parser { return s.parser }

func (s *BindDayOfWeekContext) BKW_DOW() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_DOW, 0)
}

func (s *BindDayOfWeekContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindDayOfWeekContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindDayOfWeekContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindDayOfWeek(s)
	}
}

func (s *BindDayOfWeekContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindDayOfWeek(s)
	}
}

func (p *ACIParser) BindDayOfWeek() (localctx IBindDayOfWeekContext) {
	localctx = NewBindDayOfWeekContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, ACIParserRULE_bindDayOfWeek)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Match(ACIParserBKW_DOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindAuthenticationMethodContext is an interface to support dynamic dispatch.
type IBindAuthenticationMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BKW_AM() antlr.TerminalNode

	// IsBindAuthenticationMethodContext differentiates from other interfaces.
	IsBindAuthenticationMethodContext()
}

type BindAuthenticationMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindAuthenticationMethodContext() *BindAuthenticationMethodContext {
	var p = new(BindAuthenticationMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindAuthenticationMethod
	return p
}

func InitEmptyBindAuthenticationMethodContext(p *BindAuthenticationMethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindAuthenticationMethod
}

func (*BindAuthenticationMethodContext) IsBindAuthenticationMethodContext() {}

func NewBindAuthenticationMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindAuthenticationMethodContext {
	var p = new(BindAuthenticationMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindAuthenticationMethod

	return p
}

func (s *BindAuthenticationMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *BindAuthenticationMethodContext) BKW_AM() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_AM, 0)
}

func (s *BindAuthenticationMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindAuthenticationMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindAuthenticationMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindAuthenticationMethod(s)
	}
}

func (s *BindAuthenticationMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindAuthenticationMethod(s)
	}
}

func (p *ACIParser) BindAuthenticationMethod() (localctx IBindAuthenticationMethodContext) {
	localctx = NewBindAuthenticationMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, ACIParserRULE_bindAuthenticationMethod)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(ACIParserBKW_AM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetKeywordContext is an interface to support dynamic dispatch.
type ITargetKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TargetDistinguishedName() ITargetDistinguishedNameContext
	TargetToDistinguishedName() ITargetToDistinguishedNameContext
	TargetFromDistinguishedName() ITargetFromDistinguishedNameContext
	TargetSearchScope() ITargetSearchScopeContext
	TargetSearchAttribute() ITargetSearchAttributeContext
	TargetSearchFilter() ITargetSearchFilterContext
	TargetAttributeSearchFilters() ITargetAttributeSearchFiltersContext
	TargetControlObjectIdentifier() ITargetControlObjectIdentifierContext
	TargetExtendedOperationObjectIdentifier() ITargetExtendedOperationObjectIdentifierContext

	// IsTargetKeywordContext differentiates from other interfaces.
	IsTargetKeywordContext()
}

type TargetKeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetKeywordContext() *TargetKeywordContext {
	var p = new(TargetKeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetKeyword
	return p
}

func InitEmptyTargetKeywordContext(p *TargetKeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetKeyword
}

func (*TargetKeywordContext) IsTargetKeywordContext() {}

func NewTargetKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetKeywordContext {
	var p = new(TargetKeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetKeyword

	return p
}

func (s *TargetKeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetKeywordContext) TargetDistinguishedName() ITargetDistinguishedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetDistinguishedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetDistinguishedNameContext)
}

func (s *TargetKeywordContext) TargetToDistinguishedName() ITargetToDistinguishedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetToDistinguishedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetToDistinguishedNameContext)
}

func (s *TargetKeywordContext) TargetFromDistinguishedName() ITargetFromDistinguishedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetFromDistinguishedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetFromDistinguishedNameContext)
}

func (s *TargetKeywordContext) TargetSearchScope() ITargetSearchScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetSearchScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetSearchScopeContext)
}

func (s *TargetKeywordContext) TargetSearchAttribute() ITargetSearchAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetSearchAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetSearchAttributeContext)
}

func (s *TargetKeywordContext) TargetSearchFilter() ITargetSearchFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetSearchFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetSearchFilterContext)
}

func (s *TargetKeywordContext) TargetAttributeSearchFilters() ITargetAttributeSearchFiltersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetAttributeSearchFiltersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetAttributeSearchFiltersContext)
}

func (s *TargetKeywordContext) TargetControlObjectIdentifier() ITargetControlObjectIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetControlObjectIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetControlObjectIdentifierContext)
}

func (s *TargetKeywordContext) TargetExtendedOperationObjectIdentifier() ITargetExtendedOperationObjectIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetExtendedOperationObjectIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetExtendedOperationObjectIdentifierContext)
}

func (s *TargetKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetKeyword(s)
	}
}

func (s *TargetKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetKeyword(s)
	}
}

func (p *ACIParser) TargetKeyword() (localctx ITargetKeywordContext) {
	localctx = NewTargetKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, ACIParserRULE_targetKeyword)
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserTKW_TARGET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.TargetDistinguishedName()
		}

	case ACIParserTKW_TO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)
			p.TargetToDistinguishedName()
		}

	case ACIParserTKW_FROM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(332)
			p.TargetFromDistinguishedName()
		}

	case ACIParserTKW_SCOPE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(333)
			p.TargetSearchScope()
		}

	case ACIParserTKW_ATTR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(334)
			p.TargetSearchAttribute()
		}

	case ACIParserTKW_FILTER:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(335)
			p.TargetSearchFilter()
		}

	case ACIParserTKW_AF:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(336)
			p.TargetAttributeSearchFilters()
		}

	case ACIParserTKW_CTRL:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(337)
			p.TargetControlObjectIdentifier()
		}

	case ACIParserTKW_EXTOP:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(338)
			p.TargetExtendedOperationObjectIdentifier()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetDistinguishedNameContext is an interface to support dynamic dispatch.
type ITargetDistinguishedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TKW_TARGET() antlr.TerminalNode

	// IsTargetDistinguishedNameContext differentiates from other interfaces.
	IsTargetDistinguishedNameContext()
}

type TargetDistinguishedNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetDistinguishedNameContext() *TargetDistinguishedNameContext {
	var p = new(TargetDistinguishedNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetDistinguishedName
	return p
}

func InitEmptyTargetDistinguishedNameContext(p *TargetDistinguishedNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetDistinguishedName
}

func (*TargetDistinguishedNameContext) IsTargetDistinguishedNameContext() {}

func NewTargetDistinguishedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetDistinguishedNameContext {
	var p = new(TargetDistinguishedNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetDistinguishedName

	return p
}

func (s *TargetDistinguishedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetDistinguishedNameContext) TKW_TARGET() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_TARGET, 0)
}

func (s *TargetDistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetDistinguishedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetDistinguishedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetDistinguishedName(s)
	}
}

func (s *TargetDistinguishedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetDistinguishedName(s)
	}
}

func (p *ACIParser) TargetDistinguishedName() (localctx ITargetDistinguishedNameContext) {
	localctx = NewTargetDistinguishedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, ACIParserRULE_targetDistinguishedName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Match(ACIParserTKW_TARGET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetToDistinguishedNameContext is an interface to support dynamic dispatch.
type ITargetToDistinguishedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TKW_TO() antlr.TerminalNode

	// IsTargetToDistinguishedNameContext differentiates from other interfaces.
	IsTargetToDistinguishedNameContext()
}

type TargetToDistinguishedNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetToDistinguishedNameContext() *TargetToDistinguishedNameContext {
	var p = new(TargetToDistinguishedNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetToDistinguishedName
	return p
}

func InitEmptyTargetToDistinguishedNameContext(p *TargetToDistinguishedNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetToDistinguishedName
}

func (*TargetToDistinguishedNameContext) IsTargetToDistinguishedNameContext() {}

func NewTargetToDistinguishedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetToDistinguishedNameContext {
	var p = new(TargetToDistinguishedNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetToDistinguishedName

	return p
}

func (s *TargetToDistinguishedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetToDistinguishedNameContext) TKW_TO() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_TO, 0)
}

func (s *TargetToDistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetToDistinguishedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetToDistinguishedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetToDistinguishedName(s)
	}
}

func (s *TargetToDistinguishedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetToDistinguishedName(s)
	}
}

func (p *ACIParser) TargetToDistinguishedName() (localctx ITargetToDistinguishedNameContext) {
	localctx = NewTargetToDistinguishedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, ACIParserRULE_targetToDistinguishedName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		p.Match(ACIParserTKW_TO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetFromDistinguishedNameContext is an interface to support dynamic dispatch.
type ITargetFromDistinguishedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TKW_FROM() antlr.TerminalNode

	// IsTargetFromDistinguishedNameContext differentiates from other interfaces.
	IsTargetFromDistinguishedNameContext()
}

type TargetFromDistinguishedNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetFromDistinguishedNameContext() *TargetFromDistinguishedNameContext {
	var p = new(TargetFromDistinguishedNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetFromDistinguishedName
	return p
}

func InitEmptyTargetFromDistinguishedNameContext(p *TargetFromDistinguishedNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetFromDistinguishedName
}

func (*TargetFromDistinguishedNameContext) IsTargetFromDistinguishedNameContext() {}

func NewTargetFromDistinguishedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetFromDistinguishedNameContext {
	var p = new(TargetFromDistinguishedNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetFromDistinguishedName

	return p
}

func (s *TargetFromDistinguishedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetFromDistinguishedNameContext) TKW_FROM() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_FROM, 0)
}

func (s *TargetFromDistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetFromDistinguishedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetFromDistinguishedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetFromDistinguishedName(s)
	}
}

func (s *TargetFromDistinguishedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetFromDistinguishedName(s)
	}
}

func (p *ACIParser) TargetFromDistinguishedName() (localctx ITargetFromDistinguishedNameContext) {
	localctx = NewTargetFromDistinguishedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, ACIParserRULE_targetFromDistinguishedName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Match(ACIParserTKW_FROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetSearchScopeContext is an interface to support dynamic dispatch.
type ITargetSearchScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TKW_SCOPE() antlr.TerminalNode

	// IsTargetSearchScopeContext differentiates from other interfaces.
	IsTargetSearchScopeContext()
}

type TargetSearchScopeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetSearchScopeContext() *TargetSearchScopeContext {
	var p = new(TargetSearchScopeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetSearchScope
	return p
}

func InitEmptyTargetSearchScopeContext(p *TargetSearchScopeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetSearchScope
}

func (*TargetSearchScopeContext) IsTargetSearchScopeContext() {}

func NewTargetSearchScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetSearchScopeContext {
	var p = new(TargetSearchScopeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetSearchScope

	return p
}

func (s *TargetSearchScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetSearchScopeContext) TKW_SCOPE() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_SCOPE, 0)
}

func (s *TargetSearchScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetSearchScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetSearchScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetSearchScope(s)
	}
}

func (s *TargetSearchScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetSearchScope(s)
	}
}

func (p *ACIParser) TargetSearchScope() (localctx ITargetSearchScopeContext) {
	localctx = NewTargetSearchScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, ACIParserRULE_targetSearchScope)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(ACIParserTKW_SCOPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetSearchAttributeContext is an interface to support dynamic dispatch.
type ITargetSearchAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TKW_ATTR() antlr.TerminalNode

	// IsTargetSearchAttributeContext differentiates from other interfaces.
	IsTargetSearchAttributeContext()
}

type TargetSearchAttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetSearchAttributeContext() *TargetSearchAttributeContext {
	var p = new(TargetSearchAttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetSearchAttribute
	return p
}

func InitEmptyTargetSearchAttributeContext(p *TargetSearchAttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetSearchAttribute
}

func (*TargetSearchAttributeContext) IsTargetSearchAttributeContext() {}

func NewTargetSearchAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetSearchAttributeContext {
	var p = new(TargetSearchAttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetSearchAttribute

	return p
}

func (s *TargetSearchAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetSearchAttributeContext) TKW_ATTR() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_ATTR, 0)
}

func (s *TargetSearchAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetSearchAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetSearchAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetSearchAttribute(s)
	}
}

func (s *TargetSearchAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetSearchAttribute(s)
	}
}

func (p *ACIParser) TargetSearchAttribute() (localctx ITargetSearchAttributeContext) {
	localctx = NewTargetSearchAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, ACIParserRULE_targetSearchAttribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Match(ACIParserTKW_ATTR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetSearchFilterContext is an interface to support dynamic dispatch.
type ITargetSearchFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TKW_FILTER() antlr.TerminalNode

	// IsTargetSearchFilterContext differentiates from other interfaces.
	IsTargetSearchFilterContext()
}

type TargetSearchFilterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetSearchFilterContext() *TargetSearchFilterContext {
	var p = new(TargetSearchFilterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetSearchFilter
	return p
}

func InitEmptyTargetSearchFilterContext(p *TargetSearchFilterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetSearchFilter
}

func (*TargetSearchFilterContext) IsTargetSearchFilterContext() {}

func NewTargetSearchFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetSearchFilterContext {
	var p = new(TargetSearchFilterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetSearchFilter

	return p
}

func (s *TargetSearchFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetSearchFilterContext) TKW_FILTER() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_FILTER, 0)
}

func (s *TargetSearchFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetSearchFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetSearchFilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetSearchFilter(s)
	}
}

func (s *TargetSearchFilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetSearchFilter(s)
	}
}

func (p *ACIParser) TargetSearchFilter() (localctx ITargetSearchFilterContext) {
	localctx = NewTargetSearchFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, ACIParserRULE_targetSearchFilter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.Match(ACIParserTKW_FILTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetAttributeSearchFiltersContext is an interface to support dynamic dispatch.
type ITargetAttributeSearchFiltersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TKW_AF() antlr.TerminalNode

	// IsTargetAttributeSearchFiltersContext differentiates from other interfaces.
	IsTargetAttributeSearchFiltersContext()
}

type TargetAttributeSearchFiltersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetAttributeSearchFiltersContext() *TargetAttributeSearchFiltersContext {
	var p = new(TargetAttributeSearchFiltersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetAttributeSearchFilters
	return p
}

func InitEmptyTargetAttributeSearchFiltersContext(p *TargetAttributeSearchFiltersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetAttributeSearchFilters
}

func (*TargetAttributeSearchFiltersContext) IsTargetAttributeSearchFiltersContext() {}

func NewTargetAttributeSearchFiltersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetAttributeSearchFiltersContext {
	var p = new(TargetAttributeSearchFiltersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetAttributeSearchFilters

	return p
}

func (s *TargetAttributeSearchFiltersContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetAttributeSearchFiltersContext) TKW_AF() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_AF, 0)
}

func (s *TargetAttributeSearchFiltersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetAttributeSearchFiltersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetAttributeSearchFiltersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetAttributeSearchFilters(s)
	}
}

func (s *TargetAttributeSearchFiltersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetAttributeSearchFilters(s)
	}
}

func (p *ACIParser) TargetAttributeSearchFilters() (localctx ITargetAttributeSearchFiltersContext) {
	localctx = NewTargetAttributeSearchFiltersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, ACIParserRULE_targetAttributeSearchFilters)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.Match(ACIParserTKW_AF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetControlObjectIdentifierContext is an interface to support dynamic dispatch.
type ITargetControlObjectIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TKW_CTRL() antlr.TerminalNode

	// IsTargetControlObjectIdentifierContext differentiates from other interfaces.
	IsTargetControlObjectIdentifierContext()
}

type TargetControlObjectIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetControlObjectIdentifierContext() *TargetControlObjectIdentifierContext {
	var p = new(TargetControlObjectIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetControlObjectIdentifier
	return p
}

func InitEmptyTargetControlObjectIdentifierContext(p *TargetControlObjectIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetControlObjectIdentifier
}

func (*TargetControlObjectIdentifierContext) IsTargetControlObjectIdentifierContext() {}

func NewTargetControlObjectIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetControlObjectIdentifierContext {
	var p = new(TargetControlObjectIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetControlObjectIdentifier

	return p
}

func (s *TargetControlObjectIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetControlObjectIdentifierContext) TKW_CTRL() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_CTRL, 0)
}

func (s *TargetControlObjectIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetControlObjectIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetControlObjectIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetControlObjectIdentifier(s)
	}
}

func (s *TargetControlObjectIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetControlObjectIdentifier(s)
	}
}

func (p *ACIParser) TargetControlObjectIdentifier() (localctx ITargetControlObjectIdentifierContext) {
	localctx = NewTargetControlObjectIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, ACIParserRULE_targetControlObjectIdentifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(ACIParserTKW_CTRL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITargetExtendedOperationObjectIdentifierContext is an interface to support dynamic dispatch.
type ITargetExtendedOperationObjectIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TKW_EXTOP() antlr.TerminalNode

	// IsTargetExtendedOperationObjectIdentifierContext differentiates from other interfaces.
	IsTargetExtendedOperationObjectIdentifierContext()
}

type TargetExtendedOperationObjectIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetExtendedOperationObjectIdentifierContext() *TargetExtendedOperationObjectIdentifierContext {
	var p = new(TargetExtendedOperationObjectIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetExtendedOperationObjectIdentifier
	return p
}

func InitEmptyTargetExtendedOperationObjectIdentifierContext(p *TargetExtendedOperationObjectIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetExtendedOperationObjectIdentifier
}

func (*TargetExtendedOperationObjectIdentifierContext) IsTargetExtendedOperationObjectIdentifierContext() {
}

func NewTargetExtendedOperationObjectIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetExtendedOperationObjectIdentifierContext {
	var p = new(TargetExtendedOperationObjectIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetExtendedOperationObjectIdentifier

	return p
}

func (s *TargetExtendedOperationObjectIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetExtendedOperationObjectIdentifierContext) TKW_EXTOP() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_EXTOP, 0)
}

func (s *TargetExtendedOperationObjectIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetExtendedOperationObjectIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetExtendedOperationObjectIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetExtendedOperationObjectIdentifier(s)
	}
}

func (s *TargetExtendedOperationObjectIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetExtendedOperationObjectIdentifier(s)
	}
}

func (p *ACIParser) TargetExtendedOperationObjectIdentifier() (localctx ITargetExtendedOperationObjectIdentifierContext) {
	localctx = NewTargetExtendedOperationObjectIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, ACIParserRULE_targetExtendedOperationObjectIdentifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(ACIParserTKW_EXTOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOpeningParenthesisContext is an interface to support dynamic dispatch.
type IOpeningParenthesisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode

	// IsOpeningParenthesisContext differentiates from other interfaces.
	IsOpeningParenthesisContext()
}

type OpeningParenthesisContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpeningParenthesisContext() *OpeningParenthesisContext {
	var p = new(OpeningParenthesisContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_openingParenthesis
	return p
}

func InitEmptyOpeningParenthesisContext(p *OpeningParenthesisContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_openingParenthesis
}

func (*OpeningParenthesisContext) IsOpeningParenthesisContext() {}

func NewOpeningParenthesisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpeningParenthesisContext {
	var p = new(OpeningParenthesisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_openingParenthesis

	return p
}

func (s *OpeningParenthesisContext) GetParser() antlr.Parser { return s.parser }

func (s *OpeningParenthesisContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *OpeningParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpeningParenthesisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpeningParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterOpeningParenthesis(s)
	}
}

func (s *OpeningParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitOpeningParenthesis(s)
	}
}

func (p *ACIParser) OpeningParenthesis() (localctx IOpeningParenthesisContext) {
	localctx = NewOpeningParenthesisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, ACIParserRULE_openingParenthesis)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Match(ACIParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClosingParenthesisContext is an interface to support dynamic dispatch.
type IClosingParenthesisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RPAREN() antlr.TerminalNode

	// IsClosingParenthesisContext differentiates from other interfaces.
	IsClosingParenthesisContext()
}

type ClosingParenthesisContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClosingParenthesisContext() *ClosingParenthesisContext {
	var p = new(ClosingParenthesisContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_closingParenthesis
	return p
}

func InitEmptyClosingParenthesisContext(p *ClosingParenthesisContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_closingParenthesis
}

func (*ClosingParenthesisContext) IsClosingParenthesisContext() {}

func NewClosingParenthesisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClosingParenthesisContext {
	var p = new(ClosingParenthesisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_closingParenthesis

	return p
}

func (s *ClosingParenthesisContext) GetParser() antlr.Parser { return s.parser }

func (s *ClosingParenthesisContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ClosingParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClosingParenthesisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClosingParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterClosingParenthesis(s)
	}
}

func (s *ClosingParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitClosingParenthesis(s)
	}
}

func (p *ACIParser) ClosingParenthesis() (localctx IClosingParenthesisContext) {
	localctx = NewClosingParenthesisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, ACIParserRULE_closingParenthesis)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.Match(ACIParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommaDelimiterContext is an interface to support dynamic dispatch.
type ICommaDelimiterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode

	// IsCommaDelimiterContext differentiates from other interfaces.
	IsCommaDelimiterContext()
}

type CommaDelimiterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommaDelimiterContext() *CommaDelimiterContext {
	var p = new(CommaDelimiterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_commaDelimiter
	return p
}

func InitEmptyCommaDelimiterContext(p *CommaDelimiterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_commaDelimiter
}

func (*CommaDelimiterContext) IsCommaDelimiterContext() {}

func NewCommaDelimiterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommaDelimiterContext {
	var p = new(CommaDelimiterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_commaDelimiter

	return p
}

func (s *CommaDelimiterContext) GetParser() antlr.Parser { return s.parser }

func (s *CommaDelimiterContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ACIParserCOMMA, 0)
}

func (s *CommaDelimiterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommaDelimiterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommaDelimiterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterCommaDelimiter(s)
	}
}

func (s *CommaDelimiterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitCommaDelimiter(s)
	}
}

func (p *ACIParser) CommaDelimiter() (localctx ICommaDelimiterContext) {
	localctx = NewCommaDelimiterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, ACIParserRULE_commaDelimiter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(ACIParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRuleTerminatorContext is an interface to support dynamic dispatch.
type IRuleTerminatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMI() antlr.TerminalNode

	// IsRuleTerminatorContext differentiates from other interfaces.
	IsRuleTerminatorContext()
}

type RuleTerminatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleTerminatorContext() *RuleTerminatorContext {
	var p = new(RuleTerminatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_ruleTerminator
	return p
}

func InitEmptyRuleTerminatorContext(p *RuleTerminatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_ruleTerminator
}

func (*RuleTerminatorContext) IsRuleTerminatorContext() {}

func NewRuleTerminatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleTerminatorContext {
	var p = new(RuleTerminatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_ruleTerminator

	return p
}

func (s *RuleTerminatorContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleTerminatorContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ACIParserSEMI, 0)
}

func (s *RuleTerminatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleTerminatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleTerminatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterRuleTerminator(s)
	}
}

func (s *RuleTerminatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitRuleTerminator(s)
	}
}

func (p *ACIParser) RuleTerminator() (localctx IRuleTerminatorContext) {
	localctx = NewRuleTerminatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, ACIParserRULE_ruleTerminator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(ACIParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnchorContext is an interface to support dynamic dispatch.
type IAnchorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ANCHOR() antlr.TerminalNode

	// IsAnchorContext differentiates from other interfaces.
	IsAnchorContext()
}

type AnchorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnchorContext() *AnchorContext {
	var p = new(AnchorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_anchor
	return p
}

func InitEmptyAnchorContext(p *AnchorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_anchor
}

func (*AnchorContext) IsAnchorContext() {}

func NewAnchorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnchorContext {
	var p = new(AnchorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_anchor

	return p
}

func (s *AnchorContext) GetParser() antlr.Parser { return s.parser }

func (s *AnchorContext) ANCHOR() antlr.TerminalNode {
	return s.GetToken(ACIParserANCHOR, 0)
}

func (s *AnchorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnchorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnchorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAnchor(s)
	}
}

func (s *AnchorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAnchor(s)
	}
}

func (p *ACIParser) Anchor() (localctx IAnchorContext) {
	localctx = NewAnchorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, ACIParserRULE_anchor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)
		p.Match(ACIParserANCHOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccessControlLabelContext is an interface to support dynamic dispatch.
type IAccessControlLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUOTED_STRING() antlr.TerminalNode

	// IsAccessControlLabelContext differentiates from other interfaces.
	IsAccessControlLabelContext()
}

type AccessControlLabelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessControlLabelContext() *AccessControlLabelContext {
	var p = new(AccessControlLabelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_accessControlLabel
	return p
}

func InitEmptyAccessControlLabelContext(p *AccessControlLabelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_accessControlLabel
}

func (*AccessControlLabelContext) IsAccessControlLabelContext() {}

func NewAccessControlLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessControlLabelContext {
	var p = new(AccessControlLabelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_accessControlLabel

	return p
}

func (s *AccessControlLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessControlLabelContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(ACIParserQUOTED_STRING, 0)
}

func (s *AccessControlLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessControlLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessControlLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAccessControlLabel(s)
	}
}

func (s *AccessControlLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAccessControlLabel(s)
	}
}

func (p *ACIParser) AccessControlLabel() (localctx IAccessControlLabelContext) {
	localctx = NewAccessControlLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, ACIParserRULE_accessControlLabel)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Match(ACIParserQUOTED_STRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ACIParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 27:
		var t *BindRulesContext = nil
		if localctx != nil {
			t = localctx.(*BindRulesContext)
		}
		return p.BindRules_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ACIParser) BindRules_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
