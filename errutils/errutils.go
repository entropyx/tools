package errutils

func IsNilFunc(f func() error) bool { return f() == nil }
