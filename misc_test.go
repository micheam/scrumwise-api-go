package wiseman

import "testing"

func TestTimeUnit_String(t *testing.T) {
	type fields struct {
		Value float64
		Typ   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"0.00 Point", fields{0, "Point"}, "0.00 Point"},
		{"1.05 Hour", fields{1.05, "Hour"}, "1.05 Hour"},

		{"none", fields{-1, "Point"}, "none"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tu := TimeUnit{
				Value: tt.fields.Value,
				Typ:   tt.fields.Typ,
			}
			if got := tu.String(); got != tt.want {
				t.Errorf("TimeUnit.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
