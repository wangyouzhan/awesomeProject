package five

type matrix struct {
	wang int
}

func NewMatrix(wang int) *matrix  {
	m := new(matrix)
	m.wang = wang
	return m
}