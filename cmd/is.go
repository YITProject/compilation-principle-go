package cmd

func IsIdentifier(b byte) bool {
	return IsLetter(b) || b == '_'
}

// IsUpper A-Z
func IsUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'

}

// IsLower a-Z
func IsLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

// IsNumber 0-9
func IsNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

// IsLetter a-zA-Z
func IsLetter(b byte) bool {
	return IsUpper(b) || IsLower(b)
}

// IsSymbol return not id, number, space
func IsSymbol(b byte) bool {
	return !IsIdentifier(b) && !IsNumber(b) && b != ' '
}
