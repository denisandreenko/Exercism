package strain

type Ints []int
type Lists [][]int
type Strings []string

func (c *Ints)Keep(f func(int) bool) (res Ints) {
	for _, v := range *c {
		if f(v) {
			res = append(res, v)
		}
	}
	return
}
func (c *Ints)Discard(f func(int) bool) (res Ints) {
	for _, v := range *c {
		if !f(v) {
			res = append(res, v)
		}
	}
	return
}
func (c *Lists)Keep(f func([]int) bool) (res Lists) {
	for _, v := range *c {
		if f(v) {
			res = append(res, v)
		}
	}
	return
}
func (c *Strings)Keep(f func(string) bool) (res Strings) {
	for _, v := range *c {
		if f(v) {
			res = append(res, v)
		}
	}
	return
}
