package maze

import (
	"container/list"
	"fmt"
	"math/rand"
)

type point struct {
	r int
	c int
}

func (p *point) String() string {
	return fmt.Sprintf("point{r:%d,c:%d}", p.r, p.c)
}

func (p *point) move(pt point) (rtn point) {
	rtn.r = p.r + pt.r
	rtn.c = p.c + pt.c
	return
}

type pointset struct {
	data list.List
}

func new_pointset() (rtn *pointset) {
	rtn = new(pointset)
	rtn.data.Init()
	return
}

func (set *pointset) String() (rtn string) {
	rtn = ""
	for i := set.data.Front(); i != nil; i = i.Next() {
		if len(rtn) != 0 {
			rtn += ","
		}

		pt := i.Value.(point)
		rtn += pt.String()
	}
	return "pointset{" + rtn + "}"
}

func (set *pointset) len() int {
	return set.data.Len()
}

func (set *pointset) member(pt point) bool {
	for i := set.data.Front(); i != nil; i = i.Next() {
		if i.Value == pt {
			return true
		}
	}
	return false
}

func (set *pointset) add(pt point) bool {
	if set.member(pt) {
		return false
	}
	set.data.PushBack(pt)
	return true
}

func (set *pointset) random() (rtn point) {
	if set.len() == 0 {
		panic("Can't find random point in empty pointset")
	}

	n := rand.Intn(set.data.Len())
	for i := set.data.Front(); i != nil; i = i.Next() {
		if n == 0 {
			rtn = i.Value.(point)
			set.data.Remove(i)
		}
	}
	return
}
