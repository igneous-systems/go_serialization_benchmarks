package goserbench

//go:generate msgp -o msgp_gen.go -io=false -tests=false

type Primitive struct {
	String  string
	Int64   int64
	Bool    bool
	Float64 float64
}

type Data struct {
	Metadata string
	Data     []byte
}

type Map struct {
	Map map[string]string
}
