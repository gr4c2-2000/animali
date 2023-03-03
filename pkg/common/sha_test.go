package common

import "testing"

func TestAsSha256(t *testing.T) {
	type args struct {
		o interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"string", args{"1"}, "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b"},
		{"int", args{1}, "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsSha256(tt.args.o); got != tt.want {
				t.Errorf("AsSha256() = %v, want %v on test %v", got, tt.want, tt.name)
			}
		})
	}
}
