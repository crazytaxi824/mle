package slice

type int64Type struct{}

func Int64() int64Type { return int64Type{} }

func (int64Type) IndexOf(s []int64, n int64) int {
	for k := range s {
		if s[k] == n {
			return k
		}
	}
	return -1
}

func (int64Type) LastIndexOf(s []int64, n int64) int {
	if len(s) < 1 {
		return -1
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == n {
			return i
		}
	}
	return -1
}

func (int64Type) Equal(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}

	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func (int64Type) ContainsAny(a, b []int64) bool {
	for ka := range a {
		for kb := range b {
			if a[ka] == b[kb] {
				return true
			}
		}
	}
	return false
}
