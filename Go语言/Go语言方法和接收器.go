package Go语言

import "fmt"

// 面向过程实现方法
//type Bag struct {
//	items []int
//}
//
//// 将一个物品放入背包的过程
//func Insert(b *Bag, itemid int) {
//	b.items = append(b.items, itemid)
//}
//
//func main(){
//	bag := new(Bag)
//	Insert(bag, 1001)
//}

////Go语言的结构体方法
//type Bag struct{
//	items []int
//}
//
//func (b *Bag) Insert(itemid int) {
//	b.items = append(b.items, itemid)
//}
//
//func main() {
//	b := new(Bag)
//	b.Insert(1001)
//}

// 定义属性结构
type Property struct {
	value int // 属性值
}

// 设置属性值
func (p *Property) SetValue(v int) {
	// 修改p的成员变量
	p.value = v
}

// 取属性值
func (p *Property) Value() int {
	return p.value
}

func main() {

	// 实例化属性
	p := new(Property)
	// 设置属性值
	p.SetValue(100)
	//打印值
	fmt.Println(p.Value())
}

func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}
