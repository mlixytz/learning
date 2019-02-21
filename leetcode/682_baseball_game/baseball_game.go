package leettcode682

import (
	"container/list"
	"strconv"
)

func calPoints(ops []string) int {
	history := list.New()
	var sum int
	for _, op := range ops {
		switch op {
		case "C":
			sum -= history.Remove(history.Front()).(int)
		case "D":
			history.PushFront(history.Front().Value.(int) * 2)
			sum += history.Front().Value.(int)
		case "+":
			history.PushFront(history.Front().Value.(int) + history.Front().Next().Value.(int))
			sum += history.Front().Value.(int)
		default:
			num, _ := strconv.Atoi(op)
			history.PushFront(num)
			sum += history.Front().Value.(int)
		}
	}
	return sum
}
