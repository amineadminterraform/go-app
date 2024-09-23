package utils

// constants for all supported template types
const (
	terraform = "terraform"
	opentofu = "opentofu"
	harbor = "harbor"
	shell = "shell"
	pulumi = "pulumi"
	crossplane = "crossplane"
	cloudfront = "cloudfront"
)
func IsSupportedTemplateType(templateType string) bool{
	if templateType == terraform || templateType == opentofu || templateType == harbor || templateType == shell || templateType == pulumi || templateType == crossplane || templateType == cloudfront {
		return true
	}
	return false
}

	