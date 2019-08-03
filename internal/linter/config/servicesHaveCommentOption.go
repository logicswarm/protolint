package config

// ServicesHaveCommentOption represents the option for the SERVICES_HAVE_COMMENT rule.
type ServicesHaveCommentOption struct {
	ShouldFollowGolangStyle bool `yaml:"should_follow_golang_style"`
}
