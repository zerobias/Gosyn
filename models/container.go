package models

/*func (ls *lexer.LexemeSequence) Transform() *ParsedSentence {

	var result ParsedSentence
	for i := range []lexer.Lexeme(*ls) {
		el := new(ParsedSentence)
		el.Assocciated = &[]lexer.Lexeme(*ls)[i]
		result.childs = append(result.childs, el)
	}
	return &result
}*/

/*
func Container() {
	var set struct {
		Set
		Add, Rem func(float64)
	}
	makeSet(&set)
	set.Add(3.14)
	set.Add(42.)
	//set.Rem(42.)
	fmt.Println(set.Set)
	var wtf struct {
		Set
		Add, Rem func(string)
	}
	makeSet(&wtf)
	wtf.Add("!!!!!!")
	wtf.Add("lol")
	wtf.Add("AAA")
	wtf.Rem("AAA")
	fmt.Println(wtf.Set)
}

type Set map[interface{}]struct{}

// Input:
//	*struct{Set; Add, Remove func(int) T}
//	where T is an arbitrary type
func makeSet(stptr interface{}) {
	// get struct fields and types
	st := reflect.ValueOf(stptr).Elem()
	st_set := st.Field(0)
	st_add := st.Field(1)
	st_rem := st.Field(2)

	// make Set
	st_set.Set(reflect.ValueOf(make(Set)))

	// make Add
	adder := func(in []reflect.Value) []reflect.Value {
		k := in[0]
		v := reflect.ValueOf(struct{}{})
		st_set.SetMapIndex(k, v)
		return []reflect.Value{}
	}
	st_add.Set(reflect.MakeFunc(st_add.Type(), adder))

	// make Rem
	remover := func(in []reflect.Value) []reflect.Value {
		k := in[0]
		var v reflect.Value
		st_set.SetMapIndex(k, v)
		return []reflect.Value{}
	}
	st_rem.Set(reflect.MakeFunc(st_rem.Type(), remover))
}
*/
