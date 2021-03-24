package main

type NestedIterator struct {
    vals []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    var vals []int
    var dfs func([]*NestedInteger)
    dfs = func(nestedList []*NestedInteger) {
        for _, nest := range nestedList {
            if nest.IsInteger() {
                vals = append(vals, nest.GetInteger())
            } else {
                dfs(nest.GetList())
            }
        }
    }
    dfs(nestedList)
    return &NestedIterator{vals}
}

func (it *NestedIterator) Next() int {
    val := it.vals[0]
    it.vals = it.vals[1:]
    return val
}

func (it *NestedIterator) HasNext() bool {
    return len(it.vals) > 0
}

func main() {

}
