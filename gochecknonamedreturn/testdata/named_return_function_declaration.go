package testdata

func SingleNamedReturn() (x string) { // want `don't use named return values`
	return ""
}

func SingleNamedReturnUnderscore() (_ error) { // want `don't use named return values`
	return nil
}

func SingleTypeTwoNamedReturns() (err, ignore error) { // want `don't use named return values`
	return nil, nil
}

func SingleTypeTwoNamedReturnsUnderscore() (_, y *int) { // want `don't use named return values`
	return nil, nil
}
