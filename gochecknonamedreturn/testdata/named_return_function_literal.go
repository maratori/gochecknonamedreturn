package testdata

func LiteralSingleNamedReturn() {
	_ = func() (x string) { // want `don't use named return values`
		return ""
	}
}

func LiteralNestedSingleNamedReturn() {
	_ = func() (x string) { // want `don't use named return values`
		_ = func() (y int) { // want `don't use named return values`
			return 0
		}
		return ""
	}
}

func LiteralSingleNamedReturnUnderscore() {
	_ = func() (_ error) { // want `don't use named return values`
		return nil
	}
}

func LiteralSingleTypeTwoNamedReturns() {
	_ = func() (err, ignore error) { // want `don't use named return values`
		return nil, nil
	}
}

func LiteralSingleTypeTwoNamedReturnsUnderscore() {
	_ = func() (_, y *int) { // want `don't use named return values`
		return nil, nil
	}
}
