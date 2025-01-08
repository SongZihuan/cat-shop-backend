package utils

type StringBool string

const enable StringBool = "enable"
const disable StringBool = "disable"

func (s *StringBool) check() bool {
	return *s == enable || *s == disable
}

func (s *StringBool) is(v StringBool) bool {
	if !s.check() {
		panic("bad value")
	}

	return *s == v
}

func (s *StringBool) IsEnable() bool {
	return s.is(enable)
}

func (s *StringBool) IsDisable() bool {
	return s.is(disable)
}

func (s *StringBool) setDefault(v StringBool) {
	if !s.check() {
		*s = v
	}
}

func (s *StringBool) SetDefaultEanble() {
	s.setDefault(enable)
}

func (s *StringBool) SetDefaultDisable() {
	s.setDefault(disable)
}

func (s *StringBool) ToString() string {
	return s.ToStringDefaultEnable()
}

func (s *StringBool) ToStringDefaultEnable() string {
	if s.check() {
		return string(*s)
	}

	return string(enable)
}

func (s *StringBool) ToStringDefaultDisable() string {
	if s.check() {
		return string(*s)
	}

	return string(disable)
}
