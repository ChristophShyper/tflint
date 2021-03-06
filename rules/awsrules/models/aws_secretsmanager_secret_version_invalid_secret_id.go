// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSecretsmanagerSecretVersionInvalidSecretIDRule checks the pattern is valid
type AwsSecretsmanagerSecretVersionInvalidSecretIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSecretsmanagerSecretVersionInvalidSecretIDRule returns new rule with default attributes
func NewAwsSecretsmanagerSecretVersionInvalidSecretIDRule() *AwsSecretsmanagerSecretVersionInvalidSecretIDRule {
	return &AwsSecretsmanagerSecretVersionInvalidSecretIDRule{
		resourceType:  "aws_secretsmanager_secret_version",
		attributeName: "secret_id",
		max:           2048,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSecretsmanagerSecretVersionInvalidSecretIDRule) Name() string {
	return "aws_secretsmanager_secret_version_invalid_secret_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecretsmanagerSecretVersionInvalidSecretIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecretsmanagerSecretVersionInvalidSecretIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecretsmanagerSecretVersionInvalidSecretIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecretsmanagerSecretVersionInvalidSecretIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"secret_id must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"secret_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
