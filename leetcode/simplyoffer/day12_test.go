package simplyoffer

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}

	tmp := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test",
			args: args{
				headA: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  1,
							Next: tmp,
						},
					},
				},
				headB: &ListNode{
					Val:  3,
					Next: tmp,
				},
			},
			want: tmp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
