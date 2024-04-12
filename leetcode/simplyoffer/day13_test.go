package simplyoffer

import (
	"reflect"
	"testing"
)

func Test_exchange(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exchange(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exchange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s: "1 ",
			},
			want: "1",
		},
		{
			name: "test",
			args: args{
				s: "  hello world!  ",
			},
			want: "world! hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
