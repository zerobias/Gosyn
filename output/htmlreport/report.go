package htmlreport

import (
	"gosyn/controller"
	"gosyn/models"
	//"bufio"
	// "io"
	// "io/ioutil"
	//"bytes"
	"fmt"
	//"gosyn/output"
	"gosyn/utils"
	//"strconv"
	"strings"
)

func SetFilename(filename string) {
	filenameReport = filename
}
func MakeReport(tree *controller.DataFacade) {
	newTree := controller.DataFacade{}
	treeRebuilding(tree, &newTree)
	fmt.Println("-----Rebuilded:")
	fmt.Println(*tree)
	storeResult(htmlHead + makeList(&newTree, true) + htmlEnd)
	fmt.Println(&newTree)
}

func rebuildTree(node *controller.DataFacade, newTree *controller.DataFacade) {
	if (*node).HasChilds() {
		treeRebuilding(node, newTree)
	} else {
		(*newTree).Add(*node)
	}
}

func treeRebuilding(tree *controller.DataFacade, newTree *controller.DataFacade) {
	rebuildChilds := func(parent *controller.DataFacade) {
		/*for _, child := range *tree.Childs() {
			rebuildTree(child, parent)
		}*/
	}

	switch (*tree).StepType() {
	case models.ST_SEQ:
		rebuildChilds(newTree)
	default:
		result := controller.NewFacadeText(tree.GetStep(), tree.GetLex())
		rebuildChilds(result)
		(*newTree).Add(*result)
	}
	return
}

func makeList(tree *controller.DataFacade, isFirst bool) (result string) {
	for _, e := range *tree.Childs() {
		var resLi string
		if !e.HasChilds() {
			resLi = EscapedLink(e.String())
		} else {
			resLi = string(e.StepType()) + " " + strings.TrimSpace(e.String())
			if len(strings.TrimSpace(resLi)) == 0 {
				resLi = "EMPTY"
			}
			resLi = EscapedLink(resLi)
			//resLi += makeList(e, false)
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
