package main

import (
	"reflect"
	"sync"
	"testing"
)

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "test1",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "test2",
			args: args{n: 5},
			want: 5,
			//1, 1, 2, 3, 5, 8, 13, 21, 34, 55
		},
		{
			name: "test3",
			args: args{n: 10},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
			t.Log(1 & 0x07)
			t.Log(2 & 0x07)
		})
	}
}

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			want: 6,
		},
		{
			name: "test1",
			args: args{[]int{4, 2, 0, 3, 2, 5}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
			t.Log(1 & 0x07)
			t.Log(2 & 0x07)
		})
	}
}

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		// TODO: Add test cases.
		{
			name: "Test",
			args: args{
				p: "abc",
				s: "cbaebabacd",
			},
			wantAns: []int{0, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findAnagrams() = %v, want %v", gotAns, tt.wantAns)
			}
			var mu1, mu2 sync.Mutex
			mu2 = mu1
			t.Logf("%p, %p", &mu1, &mu2)

		})
	}
}

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},
		{
			name: "test",
			args: args{
				nums: []int{1, -1, 0},
				k:    0,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
