// Code generated from ACIParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package antlraci // ACIParser
import "github.com/antlr4-go/antlr/v4"

// BaseACIParserListener is a complete listener for a parse tree produced by ACIParser.
type BaseACIParserListener struct{}

var _ ACIParserListener = &BaseACIParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseACIParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseACIParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseACIParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseACIParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseACIParserListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseACIParserListener) ExitParse(ctx *ParseContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseACIParserListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseACIParserListener) ExitInstruction(ctx *InstructionContext) {}

// EnterInstructionAnchor is called when production instructionAnchor is entered.
func (s *BaseACIParserListener) EnterInstructionAnchor(ctx *InstructionAnchorContext) {}

// ExitInstructionAnchor is called when production instructionAnchor is exited.
func (s *BaseACIParserListener) ExitInstructionAnchor(ctx *InstructionAnchorContext) {}

// EnterPermissionBindRule is called when production permissionBindRule is entered.
func (s *BaseACIParserListener) EnterPermissionBindRule(ctx *PermissionBindRuleContext) {}

// ExitPermissionBindRule is called when production permissionBindRule is exited.
func (s *BaseACIParserListener) ExitPermissionBindRule(ctx *PermissionBindRuleContext) {}

// EnterPermission is called when production permission is entered.
func (s *BaseACIParserListener) EnterPermission(ctx *PermissionContext) {}

// ExitPermission is called when production permission is exited.
func (s *BaseACIParserListener) ExitPermission(ctx *PermissionContext) {}

// EnterDisposition is called when production disposition is entered.
func (s *BaseACIParserListener) EnterDisposition(ctx *DispositionContext) {}

// ExitDisposition is called when production disposition is exited.
func (s *BaseACIParserListener) ExitDisposition(ctx *DispositionContext) {}

// EnterGrantedPermission is called when production grantedPermission is entered.
func (s *BaseACIParserListener) EnterGrantedPermission(ctx *GrantedPermissionContext) {}

// ExitGrantedPermission is called when production grantedPermission is exited.
func (s *BaseACIParserListener) ExitGrantedPermission(ctx *GrantedPermissionContext) {}

// EnterWithheldPermission is called when production withheldPermission is entered.
func (s *BaseACIParserListener) EnterWithheldPermission(ctx *WithheldPermissionContext) {}

// ExitWithheldPermission is called when production withheldPermission is exited.
func (s *BaseACIParserListener) ExitWithheldPermission(ctx *WithheldPermissionContext) {}

// EnterPrivilege is called when production privilege is entered.
func (s *BaseACIParserListener) EnterPrivilege(ctx *PrivilegeContext) {}

// ExitPrivilege is called when production privilege is exited.
func (s *BaseACIParserListener) ExitPrivilege(ctx *PrivilegeContext) {}

// EnterReadPrivilege is called when production readPrivilege is entered.
func (s *BaseACIParserListener) EnterReadPrivilege(ctx *ReadPrivilegeContext) {}

// ExitReadPrivilege is called when production readPrivilege is exited.
func (s *BaseACIParserListener) ExitReadPrivilege(ctx *ReadPrivilegeContext) {}

// EnterWritePrivilege is called when production writePrivilege is entered.
func (s *BaseACIParserListener) EnterWritePrivilege(ctx *WritePrivilegeContext) {}

// ExitWritePrivilege is called when production writePrivilege is exited.
func (s *BaseACIParserListener) ExitWritePrivilege(ctx *WritePrivilegeContext) {}

// EnterSelfWritePrivilege is called when production selfWritePrivilege is entered.
func (s *BaseACIParserListener) EnterSelfWritePrivilege(ctx *SelfWritePrivilegeContext) {}

// ExitSelfWritePrivilege is called when production selfWritePrivilege is exited.
func (s *BaseACIParserListener) ExitSelfWritePrivilege(ctx *SelfWritePrivilegeContext) {}

// EnterComparePrivilege is called when production comparePrivilege is entered.
func (s *BaseACIParserListener) EnterComparePrivilege(ctx *ComparePrivilegeContext) {}

// ExitComparePrivilege is called when production comparePrivilege is exited.
func (s *BaseACIParserListener) ExitComparePrivilege(ctx *ComparePrivilegeContext) {}

// EnterSearchPrivilege is called when production searchPrivilege is entered.
func (s *BaseACIParserListener) EnterSearchPrivilege(ctx *SearchPrivilegeContext) {}

// ExitSearchPrivilege is called when production searchPrivilege is exited.
func (s *BaseACIParserListener) ExitSearchPrivilege(ctx *SearchPrivilegeContext) {}

// EnterProxyPrivilege is called when production proxyPrivilege is entered.
func (s *BaseACIParserListener) EnterProxyPrivilege(ctx *ProxyPrivilegeContext) {}

// ExitProxyPrivilege is called when production proxyPrivilege is exited.
func (s *BaseACIParserListener) ExitProxyPrivilege(ctx *ProxyPrivilegeContext) {}

// EnterAddPrivilege is called when production addPrivilege is entered.
func (s *BaseACIParserListener) EnterAddPrivilege(ctx *AddPrivilegeContext) {}

// ExitAddPrivilege is called when production addPrivilege is exited.
func (s *BaseACIParserListener) ExitAddPrivilege(ctx *AddPrivilegeContext) {}

// EnterDeletePrivilege is called when production deletePrivilege is entered.
func (s *BaseACIParserListener) EnterDeletePrivilege(ctx *DeletePrivilegeContext) {}

// ExitDeletePrivilege is called when production deletePrivilege is exited.
func (s *BaseACIParserListener) ExitDeletePrivilege(ctx *DeletePrivilegeContext) {}

// EnterImportPrivilege is called when production importPrivilege is entered.
func (s *BaseACIParserListener) EnterImportPrivilege(ctx *ImportPrivilegeContext) {}

// ExitImportPrivilege is called when production importPrivilege is exited.
func (s *BaseACIParserListener) ExitImportPrivilege(ctx *ImportPrivilegeContext) {}

// EnterExportPrivilege is called when production exportPrivilege is entered.
func (s *BaseACIParserListener) EnterExportPrivilege(ctx *ExportPrivilegeContext) {}

// ExitExportPrivilege is called when production exportPrivilege is exited.
func (s *BaseACIParserListener) ExitExportPrivilege(ctx *ExportPrivilegeContext) {}

// EnterAllPrivilege is called when production allPrivilege is entered.
func (s *BaseACIParserListener) EnterAllPrivilege(ctx *AllPrivilegeContext) {}

// ExitAllPrivilege is called when production allPrivilege is exited.
func (s *BaseACIParserListener) ExitAllPrivilege(ctx *AllPrivilegeContext) {}

// EnterNoPrivilege is called when production noPrivilege is entered.
func (s *BaseACIParserListener) EnterNoPrivilege(ctx *NoPrivilegeContext) {}

// ExitNoPrivilege is called when production noPrivilege is exited.
func (s *BaseACIParserListener) ExitNoPrivilege(ctx *NoPrivilegeContext) {}

// EnterTargetRule is called when production targetRule is entered.
func (s *BaseACIParserListener) EnterTargetRule(ctx *TargetRuleContext) {}

// ExitTargetRule is called when production targetRule is exited.
func (s *BaseACIParserListener) ExitTargetRule(ctx *TargetRuleContext) {}

// EnterTargetOperator is called when production targetOperator is entered.
func (s *BaseACIParserListener) EnterTargetOperator(ctx *TargetOperatorContext) {}

// ExitTargetOperator is called when production targetOperator is exited.
func (s *BaseACIParserListener) ExitTargetOperator(ctx *TargetOperatorContext) {}

// EnterTargetRules is called when production targetRules is entered.
func (s *BaseACIParserListener) EnterTargetRules(ctx *TargetRulesContext) {}

// ExitTargetRules is called when production targetRules is exited.
func (s *BaseACIParserListener) ExitTargetRules(ctx *TargetRulesContext) {}

// EnterWordAnd is called when production wordAnd is entered.
func (s *BaseACIParserListener) EnterWordAnd(ctx *WordAndContext) {}

// ExitWordAnd is called when production wordAnd is exited.
func (s *BaseACIParserListener) ExitWordAnd(ctx *WordAndContext) {}

// EnterWordOr is called when production wordOr is entered.
func (s *BaseACIParserListener) EnterWordOr(ctx *WordOrContext) {}

// ExitWordOr is called when production wordOr is exited.
func (s *BaseACIParserListener) ExitWordOr(ctx *WordOrContext) {}

// EnterWordNot is called when production wordNot is entered.
func (s *BaseACIParserListener) EnterWordNot(ctx *WordNotContext) {}

// ExitWordNot is called when production wordNot is exited.
func (s *BaseACIParserListener) ExitWordNot(ctx *WordNotContext) {}

// EnterBindRules is called when production bindRules is entered.
func (s *BaseACIParserListener) EnterBindRules(ctx *BindRulesContext) {}

// ExitBindRules is called when production bindRules is exited.
func (s *BaseACIParserListener) ExitBindRules(ctx *BindRulesContext) {}

// EnterBindRule is called when production bindRule is entered.
func (s *BaseACIParserListener) EnterBindRule(ctx *BindRuleContext) {}

// ExitBindRule is called when production bindRule is exited.
func (s *BaseACIParserListener) ExitBindRule(ctx *BindRuleContext) {}

// EnterExpressionValues is called when production expressionValues is entered.
func (s *BaseACIParserListener) EnterExpressionValues(ctx *ExpressionValuesContext) {}

// ExitExpressionValues is called when production expressionValues is exited.
func (s *BaseACIParserListener) ExitExpressionValues(ctx *ExpressionValuesContext) {}

// EnterSymbolicOr is called when production symbolicOr is entered.
func (s *BaseACIParserListener) EnterSymbolicOr(ctx *SymbolicOrContext) {}

// ExitSymbolicOr is called when production symbolicOr is exited.
func (s *BaseACIParserListener) ExitSymbolicOr(ctx *SymbolicOrContext) {}

// EnterBindOperator is called when production bindOperator is entered.
func (s *BaseACIParserListener) EnterBindOperator(ctx *BindOperatorContext) {}

// ExitBindOperator is called when production bindOperator is exited.
func (s *BaseACIParserListener) ExitBindOperator(ctx *BindOperatorContext) {}

// EnterLessThan is called when production lessThan is entered.
func (s *BaseACIParserListener) EnterLessThan(ctx *LessThanContext) {}

// ExitLessThan is called when production lessThan is exited.
func (s *BaseACIParserListener) ExitLessThan(ctx *LessThanContext) {}

// EnterLessThanOrEqual is called when production lessThanOrEqual is entered.
func (s *BaseACIParserListener) EnterLessThanOrEqual(ctx *LessThanOrEqualContext) {}

// ExitLessThanOrEqual is called when production lessThanOrEqual is exited.
func (s *BaseACIParserListener) ExitLessThanOrEqual(ctx *LessThanOrEqualContext) {}

// EnterGreaterThan is called when production greaterThan is entered.
func (s *BaseACIParserListener) EnterGreaterThan(ctx *GreaterThanContext) {}

// ExitGreaterThan is called when production greaterThan is exited.
func (s *BaseACIParserListener) ExitGreaterThan(ctx *GreaterThanContext) {}

// EnterGreaterThanOrEqual is called when production greaterThanOrEqual is entered.
func (s *BaseACIParserListener) EnterGreaterThanOrEqual(ctx *GreaterThanOrEqualContext) {}

// ExitGreaterThanOrEqual is called when production greaterThanOrEqual is exited.
func (s *BaseACIParserListener) ExitGreaterThanOrEqual(ctx *GreaterThanOrEqualContext) {}

// EnterEqualTo is called when production equalTo is entered.
func (s *BaseACIParserListener) EnterEqualTo(ctx *EqualToContext) {}

// ExitEqualTo is called when production equalTo is exited.
func (s *BaseACIParserListener) ExitEqualTo(ctx *EqualToContext) {}

// EnterNotEqualTo is called when production notEqualTo is entered.
func (s *BaseACIParserListener) EnterNotEqualTo(ctx *NotEqualToContext) {}

// ExitNotEqualTo is called when production notEqualTo is exited.
func (s *BaseACIParserListener) ExitNotEqualTo(ctx *NotEqualToContext) {}

// EnterBindKeyword is called when production bindKeyword is entered.
func (s *BaseACIParserListener) EnterBindKeyword(ctx *BindKeywordContext) {}

// ExitBindKeyword is called when production bindKeyword is exited.
func (s *BaseACIParserListener) ExitBindKeyword(ctx *BindKeywordContext) {}

// EnterBindUserDistinguishedName is called when production bindUserDistinguishedName is entered.
func (s *BaseACIParserListener) EnterBindUserDistinguishedName(ctx *BindUserDistinguishedNameContext) {
}

// ExitBindUserDistinguishedName is called when production bindUserDistinguishedName is exited.
func (s *BaseACIParserListener) ExitBindUserDistinguishedName(ctx *BindUserDistinguishedNameContext) {
}

// EnterBindGroupDistinguishedName is called when production bindGroupDistinguishedName is entered.
func (s *BaseACIParserListener) EnterBindGroupDistinguishedName(ctx *BindGroupDistinguishedNameContext) {
}

// ExitBindGroupDistinguishedName is called when production bindGroupDistinguishedName is exited.
func (s *BaseACIParserListener) ExitBindGroupDistinguishedName(ctx *BindGroupDistinguishedNameContext) {
}

// EnterBindRoleDistinguishedName is called when production bindRoleDistinguishedName is entered.
func (s *BaseACIParserListener) EnterBindRoleDistinguishedName(ctx *BindRoleDistinguishedNameContext) {
}

// ExitBindRoleDistinguishedName is called when production bindRoleDistinguishedName is exited.
func (s *BaseACIParserListener) ExitBindRoleDistinguishedName(ctx *BindRoleDistinguishedNameContext) {
}

// EnterBindUserAttributes is called when production bindUserAttributes is entered.
func (s *BaseACIParserListener) EnterBindUserAttributes(ctx *BindUserAttributesContext) {}

// ExitBindUserAttributes is called when production bindUserAttributes is exited.
func (s *BaseACIParserListener) ExitBindUserAttributes(ctx *BindUserAttributesContext) {}

// EnterBindGroupAttributes is called when production bindGroupAttributes is entered.
func (s *BaseACIParserListener) EnterBindGroupAttributes(ctx *BindGroupAttributesContext) {}

// ExitBindGroupAttributes is called when production bindGroupAttributes is exited.
func (s *BaseACIParserListener) ExitBindGroupAttributes(ctx *BindGroupAttributesContext) {}

// EnterBindSecurityStrengthFactor is called when production bindSecurityStrengthFactor is entered.
func (s *BaseACIParserListener) EnterBindSecurityStrengthFactor(ctx *BindSecurityStrengthFactorContext) {
}

// ExitBindSecurityStrengthFactor is called when production bindSecurityStrengthFactor is exited.
func (s *BaseACIParserListener) ExitBindSecurityStrengthFactor(ctx *BindSecurityStrengthFactorContext) {
}

// EnterBindIPAddress is called when production bindIPAddress is entered.
func (s *BaseACIParserListener) EnterBindIPAddress(ctx *BindIPAddressContext) {}

// ExitBindIPAddress is called when production bindIPAddress is exited.
func (s *BaseACIParserListener) ExitBindIPAddress(ctx *BindIPAddressContext) {}

// EnterBindDNSName is called when production bindDNSName is entered.
func (s *BaseACIParserListener) EnterBindDNSName(ctx *BindDNSNameContext) {}

// ExitBindDNSName is called when production bindDNSName is exited.
func (s *BaseACIParserListener) ExitBindDNSName(ctx *BindDNSNameContext) {}

// EnterBindTimeOfDay is called when production bindTimeOfDay is entered.
func (s *BaseACIParserListener) EnterBindTimeOfDay(ctx *BindTimeOfDayContext) {}

// ExitBindTimeOfDay is called when production bindTimeOfDay is exited.
func (s *BaseACIParserListener) ExitBindTimeOfDay(ctx *BindTimeOfDayContext) {}

// EnterBindDayOfWeek is called when production bindDayOfWeek is entered.
func (s *BaseACIParserListener) EnterBindDayOfWeek(ctx *BindDayOfWeekContext) {}

// ExitBindDayOfWeek is called when production bindDayOfWeek is exited.
func (s *BaseACIParserListener) ExitBindDayOfWeek(ctx *BindDayOfWeekContext) {}

// EnterBindAuthenticationMethod is called when production bindAuthenticationMethod is entered.
func (s *BaseACIParserListener) EnterBindAuthenticationMethod(ctx *BindAuthenticationMethodContext) {}

// ExitBindAuthenticationMethod is called when production bindAuthenticationMethod is exited.
func (s *BaseACIParserListener) ExitBindAuthenticationMethod(ctx *BindAuthenticationMethodContext) {}

// EnterTargetKeyword is called when production targetKeyword is entered.
func (s *BaseACIParserListener) EnterTargetKeyword(ctx *TargetKeywordContext) {}

// ExitTargetKeyword is called when production targetKeyword is exited.
func (s *BaseACIParserListener) ExitTargetKeyword(ctx *TargetKeywordContext) {}

// EnterTargetDistinguishedName is called when production targetDistinguishedName is entered.
func (s *BaseACIParserListener) EnterTargetDistinguishedName(ctx *TargetDistinguishedNameContext) {}

// ExitTargetDistinguishedName is called when production targetDistinguishedName is exited.
func (s *BaseACIParserListener) ExitTargetDistinguishedName(ctx *TargetDistinguishedNameContext) {}

// EnterTargetToDistinguishedName is called when production targetToDistinguishedName is entered.
func (s *BaseACIParserListener) EnterTargetToDistinguishedName(ctx *TargetToDistinguishedNameContext) {
}

// ExitTargetToDistinguishedName is called when production targetToDistinguishedName is exited.
func (s *BaseACIParserListener) ExitTargetToDistinguishedName(ctx *TargetToDistinguishedNameContext) {
}

// EnterTargetFromDistinguishedName is called when production targetFromDistinguishedName is entered.
func (s *BaseACIParserListener) EnterTargetFromDistinguishedName(ctx *TargetFromDistinguishedNameContext) {
}

// ExitTargetFromDistinguishedName is called when production targetFromDistinguishedName is exited.
func (s *BaseACIParserListener) ExitTargetFromDistinguishedName(ctx *TargetFromDistinguishedNameContext) {
}

// EnterTargetSearchScope is called when production targetSearchScope is entered.
func (s *BaseACIParserListener) EnterTargetSearchScope(ctx *TargetSearchScopeContext) {}

// ExitTargetSearchScope is called when production targetSearchScope is exited.
func (s *BaseACIParserListener) ExitTargetSearchScope(ctx *TargetSearchScopeContext) {}

// EnterTargetSearchAttribute is called when production targetSearchAttribute is entered.
func (s *BaseACIParserListener) EnterTargetSearchAttribute(ctx *TargetSearchAttributeContext) {}

// ExitTargetSearchAttribute is called when production targetSearchAttribute is exited.
func (s *BaseACIParserListener) ExitTargetSearchAttribute(ctx *TargetSearchAttributeContext) {}

// EnterTargetSearchFilter is called when production targetSearchFilter is entered.
func (s *BaseACIParserListener) EnterTargetSearchFilter(ctx *TargetSearchFilterContext) {}

// ExitTargetSearchFilter is called when production targetSearchFilter is exited.
func (s *BaseACIParserListener) ExitTargetSearchFilter(ctx *TargetSearchFilterContext) {}

// EnterTargetAttributeSearchFilters is called when production targetAttributeSearchFilters is entered.
func (s *BaseACIParserListener) EnterTargetAttributeSearchFilters(ctx *TargetAttributeSearchFiltersContext) {
}

// ExitTargetAttributeSearchFilters is called when production targetAttributeSearchFilters is exited.
func (s *BaseACIParserListener) ExitTargetAttributeSearchFilters(ctx *TargetAttributeSearchFiltersContext) {
}

// EnterTargetControlObjectIdentifier is called when production targetControlObjectIdentifier is entered.
func (s *BaseACIParserListener) EnterTargetControlObjectIdentifier(ctx *TargetControlObjectIdentifierContext) {
}

// ExitTargetControlObjectIdentifier is called when production targetControlObjectIdentifier is exited.
func (s *BaseACIParserListener) ExitTargetControlObjectIdentifier(ctx *TargetControlObjectIdentifierContext) {
}

// EnterTargetExtendedOperationObjectIdentifier is called when production targetExtendedOperationObjectIdentifier is entered.
func (s *BaseACIParserListener) EnterTargetExtendedOperationObjectIdentifier(ctx *TargetExtendedOperationObjectIdentifierContext) {
}

// ExitTargetExtendedOperationObjectIdentifier is called when production targetExtendedOperationObjectIdentifier is exited.
func (s *BaseACIParserListener) ExitTargetExtendedOperationObjectIdentifier(ctx *TargetExtendedOperationObjectIdentifierContext) {
}

// EnterOpeningParenthesis is called when production openingParenthesis is entered.
func (s *BaseACIParserListener) EnterOpeningParenthesis(ctx *OpeningParenthesisContext) {}

// ExitOpeningParenthesis is called when production openingParenthesis is exited.
func (s *BaseACIParserListener) ExitOpeningParenthesis(ctx *OpeningParenthesisContext) {}

// EnterClosingParenthesis is called when production closingParenthesis is entered.
func (s *BaseACIParserListener) EnterClosingParenthesis(ctx *ClosingParenthesisContext) {}

// ExitClosingParenthesis is called when production closingParenthesis is exited.
func (s *BaseACIParserListener) ExitClosingParenthesis(ctx *ClosingParenthesisContext) {}

// EnterCommaDelimiter is called when production commaDelimiter is entered.
func (s *BaseACIParserListener) EnterCommaDelimiter(ctx *CommaDelimiterContext) {}

// ExitCommaDelimiter is called when production commaDelimiter is exited.
func (s *BaseACIParserListener) ExitCommaDelimiter(ctx *CommaDelimiterContext) {}

// EnterRuleTerminator is called when production ruleTerminator is entered.
func (s *BaseACIParserListener) EnterRuleTerminator(ctx *RuleTerminatorContext) {}

// ExitRuleTerminator is called when production ruleTerminator is exited.
func (s *BaseACIParserListener) ExitRuleTerminator(ctx *RuleTerminatorContext) {}

// EnterAnchor is called when production anchor is entered.
func (s *BaseACIParserListener) EnterAnchor(ctx *AnchorContext) {}

// ExitAnchor is called when production anchor is exited.
func (s *BaseACIParserListener) ExitAnchor(ctx *AnchorContext) {}

// EnterAccessControlLabel is called when production accessControlLabel is entered.
func (s *BaseACIParserListener) EnterAccessControlLabel(ctx *AccessControlLabelContext) {}

// ExitAccessControlLabel is called when production accessControlLabel is exited.
func (s *BaseACIParserListener) ExitAccessControlLabel(ctx *AccessControlLabelContext) {}
