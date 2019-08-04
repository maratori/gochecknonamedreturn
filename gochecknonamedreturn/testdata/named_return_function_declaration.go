package testdata

func DeclarationSingleNamedReturn() (x string) { // want `don't use named return values`
	return ""
}

func DeclarationSingleNamedReturnUnderscore() (_ error) { // want `don't use named return values`
	return nil
}

func DeclarationSingleTypeTwoNamedReturns() (err, ignore error) { // want `don't use named return values`
	return nil, nil
}

func DeclarationSingleTypeTwoNamedReturnsUnderscore() (_, y *int) { // want `don't use named return values`
	return nil, nil
}
