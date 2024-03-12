package regular_expression_matching01

func isMatch(s string, p string) bool {
	return dp(s, p, 0, 0)
}

func dp(s, p string, i, j int) bool {
	if j == len(p) {
		return len(s) == i
	}

	if len(s) == i {
		if (len(p)-j)%2 != 0 {
			return false
		}

		for j = j + 1; j <= len(p)-1; j += 2 {
			if p[j] != '*' {
				return false
			}
		}

		return true
	}

	if s[i] == p[j] || p[j] == '.' {
		if j < len(p)-1 && p[j+1] == '*' {
			return dp(s, p, i, j+2) || dp(s, p, i+1, j)
		} else {
			return dp(s, p, i+1, j+1)
		}

	} else {
		if j < len(p)-1 && p[j+1] == '*' {
			return dp(s, p, i, j+2)
		} else {
			return false
		}
	}

	return false
}
