package htmlreport

import (
	//"gosyn/controller"
	"github.com/zerobias/gosyn/models"
	//"bufio"
	// "io"
	// "io/ioutil"
	//"bytes"
	"fmt"
	//"gosyn/output"
	"github.com/zerobias/gosyn/utils"
	//"strconv"
	"strings"
)

func SetFilename(filename string) {
	filenameReport = filename
}
func MakeReport(tree *models.TreeNode) {
	newTree := models.NewTreeNode()
	treeRebuilding(tree, newTree)
	fmt.Println("-----Rebuilded:")
	fmt.Println(*tree)
	storeResult(htmlHead + makeList(newTree, true) + htmlEnd)
	fmt.Println(*newTree)
}

func rebuildTree(node models.TreeLink, newTree models.PTreeNode) {
	switch node.Type {
	case models.TREE_NODE:
		treeRebuilding(node.AsNode(), newTree)
	case models.TREE_VALUE:
		(*newTree).AddChild(node)
	default:
		fmt.Println("[rebuildTree] UNEXCEPTION TreeElementType! ID ", node.ID)
	}
}

func treeRebuilding(tree models.PTreeNode, newTree models.PTreeNode) {
	rebuildChilds := func(parent models.PTreeNode) {
		for _, child := range tree.ChildList {
			rebuildTree(*child, parent)
		}
	}

	switch (*tree).Type() {
	case models.ST_SEQ:
		rebuildChilds(newTree)
	default:
		result := models.NewTreeNode()
		(*result).StringElement = tree.StringElement
		rebuildChilds(result)
		(*newTree).AddTreeNode(result)
	}
	return
}

func makeList(tree models.PTreeNode, isFirst bool) (result string) {
	var e *models.TreeLink
	for _, e = range tree.ChildList {
		var resLi string
		if e.Type == models.TREE_VALUE {
			resLi = EscapedLink((*(e.AsValue().Value.Value)).String())
		} else {
			node := e.AsNode()
			resLi = string(node.Type()) + " " + strings.TrimSpace(node.Value)
			if len(strings.TrimSpace(resLi)) == 0 {
				resLi = "EMPTY"
			}
			resLi = EscapedLink(resLi)
			resLi += makeList(node, false)
		}
		result += TAG_ELEM.Wrap(resLi)
	}
	if isFirst {
		result = TAG_LIST.AddAttr(ATTR_CLASS, "tree").Wrap(result)
	} else {
		result = TAG_LIST.Wrap(result)
	}
	return
}

func storeResult(htmlCode string) {
	utils.WriteFileDefault(filenameReport+reportExt, htmlCode)
}
