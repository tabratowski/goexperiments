package validation

type ValidationRule interface {
	Validate() (bool, error)
}
