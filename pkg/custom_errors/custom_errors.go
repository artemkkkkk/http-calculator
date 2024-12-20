package custom_errors

import "errors"

var (
	InvalidExpression = errors.New("invalid expression")
	UnknownOperator   = errors.New("unknown operator")
)
