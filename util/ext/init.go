package ext

//
///
func Must(src interface{}, err error) interface{} {
	rst := src
	if err != nil { panic(err) }
	return rst
}

func init() {
	//
}