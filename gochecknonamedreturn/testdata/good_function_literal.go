package testdata

func LiteralNoReturn() {
	_ = func() {
	}
}

func LiteralEmptyParentheses() () {
	_ = func() () {
	}
}

func LiteralSingleReturn() {
	_ = func() int {
		return 0
	}
}

func LiteralSingleReturnParentheses() {
	_ = func() (bool) {
		return false
	}
}
