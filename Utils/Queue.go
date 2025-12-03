package Utils

type Queue []interface{}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x)
}

func (q *Queue) Pop() interface{} {
	h := *q
	var el interface{}
	l := len(h)
	el, *q = h[0], h[1:l]
	return el
}

func (q *Queue) Length() int {
	return len(*q)
}

func NewQueue() *Queue {
	return &Queue{}
}
