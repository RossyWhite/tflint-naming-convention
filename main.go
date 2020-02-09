package main

import (
	"github.com/RossyWhite/tflint-ruleset-onename/rule"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)


func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: tflint.RuleSet{
			Name:    "onename",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rule.NewOneNameRule(),
			},
		},
	})
}
