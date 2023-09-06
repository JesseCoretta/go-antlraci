// Code generated from ACIParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package antlraci // ACIParser
import "github.com/antlr4-go/antlr/v4"

// ACIParserListener is a complete listener for a parse tree produced by ACIParser.
type ACIParserListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterInstructionAnchor is called when entering the instructionAnchor production.
	EnterInstructionAnchor(c *InstructionAnchorContext)

	// EnterPermissionBindRules is called when entering the permissionBindRules production.
	EnterPermissionBindRules(c *PermissionBindRulesContext)

	// EnterPermissionBindRule is called when entering the permissionBindRule production.
	EnterPermissionBindRule(c *PermissionBindRuleContext)

	// EnterPermission is called when entering the permission production.
	EnterPermission(c *PermissionContext)

	// EnterDisposition is called when entering the disposition production.
	EnterDisposition(c *DispositionContext)

	// EnterGrantedPermission is called when entering the grantedPermission production.
	EnterGrantedPermission(c *GrantedPermissionContext)

	// EnterWithheldPermission is called when entering the withheldPermission production.
	EnterWithheldPermission(c *WithheldPermissionContext)

	// EnterPrivilege is called when entering the privilege production.
	EnterPrivilege(c *PrivilegeContext)

	// EnterReadPrivilege is called when entering the readPrivilege production.
	EnterReadPrivilege(c *ReadPrivilegeContext)

	// EnterWritePrivilege is called when entering the writePrivilege production.
	EnterWritePrivilege(c *WritePrivilegeContext)

	// EnterSelfWritePrivilege is called when entering the selfWritePrivilege production.
	EnterSelfWritePrivilege(c *SelfWritePrivilegeContext)

	// EnterComparePrivilege is called when entering the comparePrivilege production.
	EnterComparePrivilege(c *ComparePrivilegeContext)

	// EnterSearchPrivilege is called when entering the searchPrivilege production.
	EnterSearchPrivilege(c *SearchPrivilegeContext)

	// EnterProxyPrivilege is called when entering the proxyPrivilege production.
	EnterProxyPrivilege(c *ProxyPrivilegeContext)

	// EnterAddPrivilege is called when entering the addPrivilege production.
	EnterAddPrivilege(c *AddPrivilegeContext)

	// EnterDeletePrivilege is called when entering the deletePrivilege production.
	EnterDeletePrivilege(c *DeletePrivilegeContext)

	// EnterImportPrivilege is called when entering the importPrivilege production.
	EnterImportPrivilege(c *ImportPrivilegeContext)

	// EnterExportPrivilege is called when entering the exportPrivilege production.
	EnterExportPrivilege(c *ExportPrivilegeContext)

	// EnterAllPrivilege is called when entering the allPrivilege production.
	EnterAllPrivilege(c *AllPrivilegeContext)

	// EnterNoPrivilege is called when entering the noPrivilege production.
	EnterNoPrivilege(c *NoPrivilegeContext)

	// EnterTargetRule is called when entering the targetRule production.
	EnterTargetRule(c *TargetRuleContext)

	// EnterTargetOperator is called when entering the targetOperator production.
	EnterTargetOperator(c *TargetOperatorContext)

	// EnterTargetRules is called when entering the targetRules production.
	EnterTargetRules(c *TargetRulesContext)

	// EnterWordAnd is called when entering the wordAnd production.
	EnterWordAnd(c *WordAndContext)

	// EnterWordOr is called when entering the wordOr production.
	EnterWordOr(c *WordOrContext)

	// EnterWordNot is called when entering the wordNot production.
	EnterWordNot(c *WordNotContext)

	// EnterBindRules is called when entering the bindRules production.
	EnterBindRules(c *BindRulesContext)

	// EnterBindRule is called when entering the bindRule production.
	EnterBindRule(c *BindRuleContext)

	// EnterExpressionValues is called when entering the expressionValues production.
	EnterExpressionValues(c *ExpressionValuesContext)

	// EnterSymbolicOr is called when entering the symbolicOr production.
	EnterSymbolicOr(c *SymbolicOrContext)

	// EnterBindOperator is called when entering the bindOperator production.
	EnterBindOperator(c *BindOperatorContext)

	// EnterLessThan is called when entering the lessThan production.
	EnterLessThan(c *LessThanContext)

	// EnterLessThanOrEqual is called when entering the lessThanOrEqual production.
	EnterLessThanOrEqual(c *LessThanOrEqualContext)

	// EnterGreaterThan is called when entering the greaterThan production.
	EnterGreaterThan(c *GreaterThanContext)

	// EnterGreaterThanOrEqual is called when entering the greaterThanOrEqual production.
	EnterGreaterThanOrEqual(c *GreaterThanOrEqualContext)

	// EnterEqualTo is called when entering the equalTo production.
	EnterEqualTo(c *EqualToContext)

	// EnterNotEqualTo is called when entering the notEqualTo production.
	EnterNotEqualTo(c *NotEqualToContext)

	// EnterBindKeyword is called when entering the bindKeyword production.
	EnterBindKeyword(c *BindKeywordContext)

	// EnterBindUserDistinguishedName is called when entering the bindUserDistinguishedName production.
	EnterBindUserDistinguishedName(c *BindUserDistinguishedNameContext)

	// EnterBindGroupDistinguishedName is called when entering the bindGroupDistinguishedName production.
	EnterBindGroupDistinguishedName(c *BindGroupDistinguishedNameContext)

	// EnterBindRoleDistinguishedName is called when entering the bindRoleDistinguishedName production.
	EnterBindRoleDistinguishedName(c *BindRoleDistinguishedNameContext)

	// EnterBindUserAttributes is called when entering the bindUserAttributes production.
	EnterBindUserAttributes(c *BindUserAttributesContext)

	// EnterBindGroupAttributes is called when entering the bindGroupAttributes production.
	EnterBindGroupAttributes(c *BindGroupAttributesContext)

	// EnterBindSecurityStrengthFactor is called when entering the bindSecurityStrengthFactor production.
	EnterBindSecurityStrengthFactor(c *BindSecurityStrengthFactorContext)

	// EnterBindIPAddress is called when entering the bindIPAddress production.
	EnterBindIPAddress(c *BindIPAddressContext)

	// EnterBindDNSName is called when entering the bindDNSName production.
	EnterBindDNSName(c *BindDNSNameContext)

	// EnterBindTimeOfDay is called when entering the bindTimeOfDay production.
	EnterBindTimeOfDay(c *BindTimeOfDayContext)

	// EnterBindDayOfWeek is called when entering the bindDayOfWeek production.
	EnterBindDayOfWeek(c *BindDayOfWeekContext)

	// EnterBindAuthenticationMethod is called when entering the bindAuthenticationMethod production.
	EnterBindAuthenticationMethod(c *BindAuthenticationMethodContext)

	// EnterTargetKeyword is called when entering the targetKeyword production.
	EnterTargetKeyword(c *TargetKeywordContext)

	// EnterTargetDistinguishedName is called when entering the targetDistinguishedName production.
	EnterTargetDistinguishedName(c *TargetDistinguishedNameContext)

	// EnterTargetToDistinguishedName is called when entering the targetToDistinguishedName production.
	EnterTargetToDistinguishedName(c *TargetToDistinguishedNameContext)

	// EnterTargetFromDistinguishedName is called when entering the targetFromDistinguishedName production.
	EnterTargetFromDistinguishedName(c *TargetFromDistinguishedNameContext)

	// EnterTargetSearchScope is called when entering the targetSearchScope production.
	EnterTargetSearchScope(c *TargetSearchScopeContext)

	// EnterTargetSearchAttribute is called when entering the targetSearchAttribute production.
	EnterTargetSearchAttribute(c *TargetSearchAttributeContext)

	// EnterTargetSearchFilter is called when entering the targetSearchFilter production.
	EnterTargetSearchFilter(c *TargetSearchFilterContext)

	// EnterTargetAttributeSearchFilters is called when entering the targetAttributeSearchFilters production.
	EnterTargetAttributeSearchFilters(c *TargetAttributeSearchFiltersContext)

	// EnterTargetControlObjectIdentifier is called when entering the targetControlObjectIdentifier production.
	EnterTargetControlObjectIdentifier(c *TargetControlObjectIdentifierContext)

	// EnterTargetExtendedOperationObjectIdentifier is called when entering the targetExtendedOperationObjectIdentifier production.
	EnterTargetExtendedOperationObjectIdentifier(c *TargetExtendedOperationObjectIdentifierContext)

	// EnterOpeningParenthesis is called when entering the openingParenthesis production.
	EnterOpeningParenthesis(c *OpeningParenthesisContext)

	// EnterClosingParenthesis is called when entering the closingParenthesis production.
	EnterClosingParenthesis(c *ClosingParenthesisContext)

	// EnterCommaDelimiter is called when entering the commaDelimiter production.
	EnterCommaDelimiter(c *CommaDelimiterContext)

	// EnterRuleTerminator is called when entering the ruleTerminator production.
	EnterRuleTerminator(c *RuleTerminatorContext)

	// EnterAnchor is called when entering the anchor production.
	EnterAnchor(c *AnchorContext)

	// EnterAccessControlLabel is called when entering the accessControlLabel production.
	EnterAccessControlLabel(c *AccessControlLabelContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitInstructionAnchor is called when exiting the instructionAnchor production.
	ExitInstructionAnchor(c *InstructionAnchorContext)

	// ExitPermissionBindRules is called when exiting the permissionBindRules production.
	ExitPermissionBindRules(c *PermissionBindRulesContext)

	// ExitPermissionBindRule is called when exiting the permissionBindRule production.
	ExitPermissionBindRule(c *PermissionBindRuleContext)

	// ExitPermission is called when exiting the permission production.
	ExitPermission(c *PermissionContext)

	// ExitDisposition is called when exiting the disposition production.
	ExitDisposition(c *DispositionContext)

	// ExitGrantedPermission is called when exiting the grantedPermission production.
	ExitGrantedPermission(c *GrantedPermissionContext)

	// ExitWithheldPermission is called when exiting the withheldPermission production.
	ExitWithheldPermission(c *WithheldPermissionContext)

	// ExitPrivilege is called when exiting the privilege production.
	ExitPrivilege(c *PrivilegeContext)

	// ExitReadPrivilege is called when exiting the readPrivilege production.
	ExitReadPrivilege(c *ReadPrivilegeContext)

	// ExitWritePrivilege is called when exiting the writePrivilege production.
	ExitWritePrivilege(c *WritePrivilegeContext)

	// ExitSelfWritePrivilege is called when exiting the selfWritePrivilege production.
	ExitSelfWritePrivilege(c *SelfWritePrivilegeContext)

	// ExitComparePrivilege is called when exiting the comparePrivilege production.
	ExitComparePrivilege(c *ComparePrivilegeContext)

	// ExitSearchPrivilege is called when exiting the searchPrivilege production.
	ExitSearchPrivilege(c *SearchPrivilegeContext)

	// ExitProxyPrivilege is called when exiting the proxyPrivilege production.
	ExitProxyPrivilege(c *ProxyPrivilegeContext)

	// ExitAddPrivilege is called when exiting the addPrivilege production.
	ExitAddPrivilege(c *AddPrivilegeContext)

	// ExitDeletePrivilege is called when exiting the deletePrivilege production.
	ExitDeletePrivilege(c *DeletePrivilegeContext)

	// ExitImportPrivilege is called when exiting the importPrivilege production.
	ExitImportPrivilege(c *ImportPrivilegeContext)

	// ExitExportPrivilege is called when exiting the exportPrivilege production.
	ExitExportPrivilege(c *ExportPrivilegeContext)

	// ExitAllPrivilege is called when exiting the allPrivilege production.
	ExitAllPrivilege(c *AllPrivilegeContext)

	// ExitNoPrivilege is called when exiting the noPrivilege production.
	ExitNoPrivilege(c *NoPrivilegeContext)

	// ExitTargetRule is called when exiting the targetRule production.
	ExitTargetRule(c *TargetRuleContext)

	// ExitTargetOperator is called when exiting the targetOperator production.
	ExitTargetOperator(c *TargetOperatorContext)

	// ExitTargetRules is called when exiting the targetRules production.
	ExitTargetRules(c *TargetRulesContext)

	// ExitWordAnd is called when exiting the wordAnd production.
	ExitWordAnd(c *WordAndContext)

	// ExitWordOr is called when exiting the wordOr production.
	ExitWordOr(c *WordOrContext)

	// ExitWordNot is called when exiting the wordNot production.
	ExitWordNot(c *WordNotContext)

	// ExitBindRules is called when exiting the bindRules production.
	ExitBindRules(c *BindRulesContext)

	// ExitBindRule is called when exiting the bindRule production.
	ExitBindRule(c *BindRuleContext)

	// ExitExpressionValues is called when exiting the expressionValues production.
	ExitExpressionValues(c *ExpressionValuesContext)

	// ExitSymbolicOr is called when exiting the symbolicOr production.
	ExitSymbolicOr(c *SymbolicOrContext)

	// ExitBindOperator is called when exiting the bindOperator production.
	ExitBindOperator(c *BindOperatorContext)

	// ExitLessThan is called when exiting the lessThan production.
	ExitLessThan(c *LessThanContext)

	// ExitLessThanOrEqual is called when exiting the lessThanOrEqual production.
	ExitLessThanOrEqual(c *LessThanOrEqualContext)

	// ExitGreaterThan is called when exiting the greaterThan production.
	ExitGreaterThan(c *GreaterThanContext)

	// ExitGreaterThanOrEqual is called when exiting the greaterThanOrEqual production.
	ExitGreaterThanOrEqual(c *GreaterThanOrEqualContext)

	// ExitEqualTo is called when exiting the equalTo production.
	ExitEqualTo(c *EqualToContext)

	// ExitNotEqualTo is called when exiting the notEqualTo production.
	ExitNotEqualTo(c *NotEqualToContext)

	// ExitBindKeyword is called when exiting the bindKeyword production.
	ExitBindKeyword(c *BindKeywordContext)

	// ExitBindUserDistinguishedName is called when exiting the bindUserDistinguishedName production.
	ExitBindUserDistinguishedName(c *BindUserDistinguishedNameContext)

	// ExitBindGroupDistinguishedName is called when exiting the bindGroupDistinguishedName production.
	ExitBindGroupDistinguishedName(c *BindGroupDistinguishedNameContext)

	// ExitBindRoleDistinguishedName is called when exiting the bindRoleDistinguishedName production.
	ExitBindRoleDistinguishedName(c *BindRoleDistinguishedNameContext)

	// ExitBindUserAttributes is called when exiting the bindUserAttributes production.
	ExitBindUserAttributes(c *BindUserAttributesContext)

	// ExitBindGroupAttributes is called when exiting the bindGroupAttributes production.
	ExitBindGroupAttributes(c *BindGroupAttributesContext)

	// ExitBindSecurityStrengthFactor is called when exiting the bindSecurityStrengthFactor production.
	ExitBindSecurityStrengthFactor(c *BindSecurityStrengthFactorContext)

	// ExitBindIPAddress is called when exiting the bindIPAddress production.
	ExitBindIPAddress(c *BindIPAddressContext)

	// ExitBindDNSName is called when exiting the bindDNSName production.
	ExitBindDNSName(c *BindDNSNameContext)

	// ExitBindTimeOfDay is called when exiting the bindTimeOfDay production.
	ExitBindTimeOfDay(c *BindTimeOfDayContext)

	// ExitBindDayOfWeek is called when exiting the bindDayOfWeek production.
	ExitBindDayOfWeek(c *BindDayOfWeekContext)

	// ExitBindAuthenticationMethod is called when exiting the bindAuthenticationMethod production.
	ExitBindAuthenticationMethod(c *BindAuthenticationMethodContext)

	// ExitTargetKeyword is called when exiting the targetKeyword production.
	ExitTargetKeyword(c *TargetKeywordContext)

	// ExitTargetDistinguishedName is called when exiting the targetDistinguishedName production.
	ExitTargetDistinguishedName(c *TargetDistinguishedNameContext)

	// ExitTargetToDistinguishedName is called when exiting the targetToDistinguishedName production.
	ExitTargetToDistinguishedName(c *TargetToDistinguishedNameContext)

	// ExitTargetFromDistinguishedName is called when exiting the targetFromDistinguishedName production.
	ExitTargetFromDistinguishedName(c *TargetFromDistinguishedNameContext)

	// ExitTargetSearchScope is called when exiting the targetSearchScope production.
	ExitTargetSearchScope(c *TargetSearchScopeContext)

	// ExitTargetSearchAttribute is called when exiting the targetSearchAttribute production.
	ExitTargetSearchAttribute(c *TargetSearchAttributeContext)

	// ExitTargetSearchFilter is called when exiting the targetSearchFilter production.
	ExitTargetSearchFilter(c *TargetSearchFilterContext)

	// ExitTargetAttributeSearchFilters is called when exiting the targetAttributeSearchFilters production.
	ExitTargetAttributeSearchFilters(c *TargetAttributeSearchFiltersContext)

	// ExitTargetControlObjectIdentifier is called when exiting the targetControlObjectIdentifier production.
	ExitTargetControlObjectIdentifier(c *TargetControlObjectIdentifierContext)

	// ExitTargetExtendedOperationObjectIdentifier is called when exiting the targetExtendedOperationObjectIdentifier production.
	ExitTargetExtendedOperationObjectIdentifier(c *TargetExtendedOperationObjectIdentifierContext)

	// ExitOpeningParenthesis is called when exiting the openingParenthesis production.
	ExitOpeningParenthesis(c *OpeningParenthesisContext)

	// ExitClosingParenthesis is called when exiting the closingParenthesis production.
	ExitClosingParenthesis(c *ClosingParenthesisContext)

	// ExitCommaDelimiter is called when exiting the commaDelimiter production.
	ExitCommaDelimiter(c *CommaDelimiterContext)

	// ExitRuleTerminator is called when exiting the ruleTerminator production.
	ExitRuleTerminator(c *RuleTerminatorContext)

	// ExitAnchor is called when exiting the anchor production.
	ExitAnchor(c *AnchorContext)

	// ExitAccessControlLabel is called when exiting the accessControlLabel production.
	ExitAccessControlLabel(c *AccessControlLabelContext)
}
