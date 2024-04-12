package mid_suanfa

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() (_ Codec) {
	return
}

func (Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

func (Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

type Codec2 struct {
}

func Constructor2() (_ Codec2) {
	return
}

func (c Codec2) serialize(root *TreeNode) string {
	if root == nil {
		return "X"
	}
	left := "(" + c.serialize(root.Left) + ")"
	right := "(" + c.serialize(root.Right) + ")"
	return left + strconv.Itoa(root.Val) + right
}

func (c Codec2) deserialize(data string) *TreeNode {
	var parse func() *TreeNode
	parse = func() *TreeNode {
		if data[0] == 'X' {
			data = data[1:]
			return nil
		}
		node := &TreeNode{}
		data = data[1:] // 跳过左括号
		node.Left = parse()
		data = data[1:] // 跳过右括号
		i := 0
		for data[i] == '-' || '0' <= data[i] && data[i] <= '9' {
			i++
		}
		node.Val, _ = strconv.Atoi(data[:i])
		data = data[i:]
		data = data[1:] // 跳过左括号
		node.Right = parse()
		data = data[1:] // 跳过右括号
		return node
	}

	return parse()
}

// 速度最快
/*
type Codec struct {
	Sb       *strings.Builder
	Null     string
	Sep      string
	DataList []string
}

func Constructor() Codec {
	return Codec{
		Sb:       &strings.Builder{},
		Null:     "null",
		Sep:      ",",
		DataList: make([]string, 0),
	}
}

func (this *Codec) serializeHelper(root *TreeNode) {
	if root == nil {
		this.Sb.WriteString(this.Null)
		this.Sb.WriteString(this.Sep)
		return
	}

	this.Sb.WriteString(strconv.Itoa(root.Val))
	this.Sb.WriteString(this.Sep)

	this.serialize(root.Left)
	this.serialize(root.Right)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.serializeHelper(root)
	return this.Sb.String()
}

func (this *Codec) deserializeHelper() *TreeNode {
	if len(this.DataList) <= 0 {
		return nil
	}

	rootVal := this.DataList[0]
	this.DataList = this.DataList[1:]

	//log.Printf("rootVal %v dataList %v", rootVal, this.DataList)

	if rootVal == this.Null {
		return nil
	}

	val, _ := strconv.Atoi(rootVal)
	root := &TreeNode{Val: val}

	root.Left = this.deserializeHelper()
	root.Right = this.deserializeHelper()

	return root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) <= 0 {
		return nil
	}

	this.DataList = strings.Split(data, this.Sep)
	return this.deserializeHelper()
}
*/
