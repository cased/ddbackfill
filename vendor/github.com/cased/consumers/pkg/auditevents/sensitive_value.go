package auditevents

type SensitiveValue struct {
	Value  string `json:"-"`
	Ranges []SensitiveRange
}

func NewSensitiveValue(value string, label string) SensitiveValue {
	return SensitiveValue{
		Value: value,
		Ranges: []SensitiveRange{
			{Begin: 0, End: len(value), Label: label},
		},
	}
}
