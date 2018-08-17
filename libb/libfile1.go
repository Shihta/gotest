package libb

func internalfunc0() string {
	return "this is internal string"
}

func GetFun1_1() string {
	return internalfunc0()
	// return "This is libfile1111111, the fun1_1"
}

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
