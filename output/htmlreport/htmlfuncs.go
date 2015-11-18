package htmlreport

import (
	//"gosyn/controller"
	"html"
	//"gosyn/models"
	//"fmt"
	//"gosyn/output"
	//"strconv"
	//"strings"
)

const reportExt = ".html"

var (
	filenameReport   = "report"
	filenameTemplate = "report_template.html"
)

const (
	htmlHead = `<html>
<head>
    <link rel='stylesheet' type='text/css' href='tree.css'>
</head>
<body><div class="content">`
	htmlEnd = `</div></body></html>`
)

func EscapedLink(str string) string {
	return TAG_LINK.Wrap(html.EscapeString(str))
}

/*type TemplateData struct {
	Title string
	Tree  *models.TreeNode
}

func GetSE(se models.StringElement) string {
	return se.XMLName.Local
}
func GetNode(args *models.TreeNode) string {
	return string(args.Type()) + strconv.Itoa(len(args.ChildList)) + strconv.Itoa(int(args.ChildList[0].ID)) + strconv.Itoa(int(args.ChildList[0].Type))
}
func loadTemplate(filename string) *template.Template {
	templateFuncs := template.FuncMap{"SElement": GetSE, "Node": GetNode} //"TreeNode": RangeStructer}
	t := template.New("Report template").Funcs(templateFuncs)
	var err error
	t, err = t.Parse(utils.LoadfileDefault(filename))
	utils.Check(err)
	return t
}*/
