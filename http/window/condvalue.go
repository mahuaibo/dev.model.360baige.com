package window
type CondValue struct {
	Type  string // And AndNot Or OrNot
	Exprs  string
	Args   interface{}
}