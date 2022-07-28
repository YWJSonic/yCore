package queue

type Queue struct {
	value []interface{}
}

func (q *Queue) Count() int {
	return len(q.value)
}

func (q *Queue) Push(data interface{}) {
	q.value = append(q.value, data)
}

// is not fix data race!!
func (q *Queue) Pop() interface{} {
	var res interface{}
	count := len(q.value)
	if count <= 0 {
		return nil
	}
	res, q.value = q.value[count-1], q.value[:count-1]
	return res
}
