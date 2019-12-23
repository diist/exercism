package binarysearchtree

type SearchTreeData struct {
	data        int
	left, right *SearchTreeData
}

func Bst(data int) *SearchTreeData {
	return &SearchTreeData{data: data}
}

func (bst *SearchTreeData) Insert(data int) {
	if data <= bst.data {
		if bst.left == nil {
			bst.left = &SearchTreeData{data: data}
		} else {
			bst.left.Insert(data)
		}
	} else {
		if bst.right == nil {
			bst.right = &SearchTreeData{data: data}
		} else {
			bst.right.Insert(data)
		}
	}
}

func (bst *SearchTreeData) MapString(convert func(int) string) []string {
	m := []string{convert(bst.data)}

	if bst.left != nil {
		m = append(bst.left.MapString(convert), m...)
	}

	if bst.right != nil {
		m = append(m, bst.right.MapString(convert)...)
	}

	return m
}

func (bst *SearchTreeData) MapInt(convert func(int) int) []int {
	m := []int{convert(bst.data)}

	if bst.left != nil {
		m = append(bst.left.MapInt(convert), m...)
	}

	if bst.right != nil {
		m = append(m, bst.right.MapInt(convert)...)
	}

	return m
}
