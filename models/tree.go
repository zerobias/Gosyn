package models

import (
	//"fmt"
	"gosyn/lexer"
)

type PTreeNode *TreeNode
type PTreeValue *TreeValue

type Typer interface {
	AsNode() PTreeNode
	AsValue() PTreeValue
}
type Stored func() Typer

/*type Node interface {
	Init(link TreeLink, tree PTreeNode)
	Get() Typer
	Clone(original PTreeNode)
	MoveTo(newParent PTreeNode)
}*/

func (tn TreeNode) AsNode() PTreeNode    { return &tn }
func (tn TreeNode) AsValue() PTreeValue  { return new(TreeValue) }
func (tv TreeValue) AsNode() PTreeNode   { return new(TreeNode) }
func (tv TreeValue) AsValue() PTreeValue { return &tv }

type TreeElementType int
type TreeLinkID int

const (
	TREE_NODE TreeElementType = iota
	TREE_VALUE
)

type TreeLink struct {
	ID     TreeLinkID
	Type   TreeElementType
	Parent PTreeNode
}

func (ptl *TreeLink) AsNode() PTreeNode   { return (*(*ptl).Parent).GetTreeNode(*ptl) }
func (ptl *TreeLink) AsValue() PTreeValue { return (*(*ptl).Parent).GetTreeValue(*ptl) }

type TreeValue struct {
	StringElement
	Value *lexer.Lexeme //TODO !!!remove conflict with StrEl.Value
}

type TreeNode struct {
	StringElement
	ChildList  []*TreeLink
	valuesList map[TreeLinkID]PTreeValue
	nodesList  map[TreeLinkID]PTreeNode
}

const c_CHILD_MAX = 255

func NewTreeNode() PTreeNode {
	node := *new(TreeNode)
	node.ChildList = make([]*TreeLink, 0, c_CHILD_MAX)
	node.valuesList = make(map[TreeLinkID]PTreeValue)
	node.nodesList = make(map[TreeLinkID]PTreeNode)
	node.StringElement = *new(StringElement)
	return &node
}

func (tn *TreeNode) GetTreeValue(link TreeLink) PTreeValue {
	return tn.valuesList[link.ID]
}

func (tn *TreeNode) GetTreeNode(link TreeLink) PTreeNode {
	return tn.nodesList[link.ID]
}

func (tn *TreeNode) GetStringElement(link TreeLink) StringElement {
	if link.Type == TREE_NODE {
		return tn.GetTreeNode(link).StringElement
	} else {
		return tn.GetTreeNode(link).StringElement
	}
}

func (tn *TreeNode) addTreeLink(elType TreeElementType) TreeLinkID {
	link := TreeLink{TreeLinkID(len(tn.ChildList)), elType, tn}
	tn.ChildList = append(tn.ChildList, &link)
	return link.ID
}

func (tn *TreeNode) AddTreeValue(strEl *StringElement, val *lexer.Lexeme) { //TODO remove possible duplication StringEl/Lexeme
	//tn.valuesList
	tn.valuesList[tn.addTreeLink(TREE_VALUE)] = &TreeValue{*strEl, val}
}

func (tn *TreeNode) AddTreeNode(node *TreeNode) {
	tn.nodesList[tn.addTreeLink(TREE_NODE)] = node
}

func (tn *TreeNode) AddChild(link TreeLink) {
	switch link.Type {
	case TREE_NODE:
		tn.AddTreeNode(link.AsNode())
	case TREE_VALUE:
		tn.valuesList[tn.addTreeLink(TREE_VALUE)] = link.AsValue()
		/*default:
		fmt.Println("[AddChild] UNEXCEPTION TreeElementType! ID ", link.ID)*/
	}
}
