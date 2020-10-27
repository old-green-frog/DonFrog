package lexing

//base types
type Integer int32
type Float float32
type String string

type Utils interface {
	Parse(name string)
	PLusAdd()
	Equal()
}

var MainName string

//containers for variables
var IntVars = make(map[string]Integer)
var FlVars = make(map[string]Float)
var StrVars = make(map[string]String)

//methods
func (i Integer) Parse(name string) { IntVars[name] = i }
func (f Float) Parse(name string)   { FlVars[name] = f }
func (s String) Parse(name string)  { StrVars[name] = s }

//
//
//

func (i *Integer) PlusAdd(val interface{}) {

	if v, t := val.(Integer); t {
		*i += v
	} else {
		panic("Wow! This is a type_error!")
	}

}
func (f *Float) PlusAdd(val interface{}) {

	if v, t := val.(Float); t {
		*f += v
	} else {
		panic("Wow! This is a type_error!")
	}

}
func (s *String) PlusAdd(val interface{}) {

	if v, t := val.(String); t {
		*s += v
	} else {
		panic("Wow! This is a type_error!")
	}
}

//
//
//

//init function for vars
func New(val interface{}) {

	switch v := val.(type) {
	case Integer:
	case Float:
	case String:
		v.Parse(MainName)
	default:
		panic("Wow! This is a type_error!")
	}
}
