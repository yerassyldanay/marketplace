package utils

/*
	this function has been created to simplify the if-else condition
 */
var If_condition_then_first_else_second = func(condition bool, first interface{}, second interface{}) (interface {}) {
	if condition {
		return first
	}
	return second
}

