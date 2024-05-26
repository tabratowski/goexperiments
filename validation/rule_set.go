package validation

type ruleSet struct {
	rules []ValidationRule
}

func RuleSet(fields ...ruleField) *ruleSet {
	rs := &ruleSet{}
	for _, f := range fields {
		rs.rules = append(rs.rules, f.RulesSet()...)
	}
	return rs
}

func (r *ruleSet) Validate() []error {
	var errs []error
	for _, rule := range r.rules {
		valid, err := rule.Validate()
		if !valid {
			errs = append(errs, err)
		}
	}
	return errs
}
