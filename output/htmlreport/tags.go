package htmlreport

import (
//"gosyn/controller"
//	"html"
//"gosyn/models"
//"fmt"
//"gosyn/output"
//"strconv"
//"strings"
)

type htmlTag interface {
	String() string
	Open() string
	Close() string
	Wrap(str string) string
}

const (
	SYM_SPACE = " "
	SYM_EQUAL = "="
)

const (
	TAG_LIST   simpleTag = "ul"
	TAG_ELEM   simpleTag = "li"
	TAG_LINK   simpleTag = "a"
	ATTR_CLASS string    = "class"
)

type simpleTag string

type varTag struct {
	simpleTag
	attr map[string]string
}

func (st simpleTag) String() string { return string(st) }
func (st simpleTag) Open() string   { return brackets(string(st)) }
func (st simpleTag) Close() string  { return brackets("/" + string(st)) }
func (st simpleTag) Wrap(str string) string {
	return wrapper(str, st.Open(), st.Close())
}

func (st simpleTag) AddAttr(key, value string) varTag {
	result := varTag{st, make(map[string]string)}
	result.AddAttr(key, value)
	return result
}

func (vt varTag) String() string { return string(vt.simpleTag) }
func (vt varTag) Open() string {
	result := vt.String()
	for attr, val := range vt.attr {
		result += SYM_SPACE + attr + SYM_EQUAL + quotes(val)
	}
	return brackets(result)
}

//func (vt varTag) Close() string { return brackets(vt.simpleTag.Close()) }
func (vt varTag) Wrap(str string) string {
	return wrapper(str, vt.Open(), vt.Close())
}

func (vt *varTag) AddAttr(key, value string) {
	vt.attr[string(key)] = string(value)
}

func quotes(str string) string {
	return wrapper(str, "\"", "\"")
}
func brackets(str string) string {
	return wrapper(str, "<", ">")
}
func wrapper(body, first, last string) string {
	return string(first + body + last)
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
