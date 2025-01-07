package config

type StringBool string

const Enable StringBool = "enable"
const Disable StringBool = "disable"

func (s *StringBool) check() bool {
	return *s == Enable || *s == Disable
}

func (s *StringBool) Is(v StringBool) bool {
	if !s.check() {
		panic("bad value")
	}

	return *s == v
}

func (s *StringBool) SetDefault(v StringBool) {
	if !s.check() {
		*s = v
	}
}
