package sexp

var rootList []interface{}

func parseLang(buf []byte, offset, end int) (interface{}, int, error) {
	isRoot := false
	if rootList == nil {
		rootList = make([]interface{}, 0)
		isRoot = true
	}

	var headf interface{}
	var err error
	headf, offset, err = parseReturningOffset(buf, offset)
	if err != nil {
		return nil, offset, err
	}
	rootList = append(rootList, headf)

	for offset < end {
		_, offset, err = parseReturningOffset(buf, offset)
		if err != nil {
			return nil, offset, err
		}
	}

	var rtn interface{}
	if isRoot {
		rtn = rootList
		rootList = nil
	}
	return rtn, offset, nil
}
