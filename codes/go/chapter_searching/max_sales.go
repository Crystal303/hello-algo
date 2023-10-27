package chapter_searching

// maxSales 求连续区间的最大累计值
func maxSales(sales []int) int {
	return get(sales, 0, len(sales)-1).mSum
}

func get(sales []int, left, right int) status {
	if left == right {
		return status{
			iSum: sales[left],
			lSum: sales[left],
			rSum: sales[left],
			mSum: sales[left],
		}
	}
	m := left + (right-left)>>1
	leftSub := get(sales, left, m)
	rightSub := get(sales, m+1, right)
	return push(leftSub, rightSub)
}

func push(l, r status) status {
	var res status
	res.iSum = l.iSum + r.iSum
	res.lSum = calcMax(l.iSum+r.lSum, l.lSum)
	res.rSum = calcMax(r.iSum+l.rSum, r.rSum)
	res.mSum = calcMax(calcMax(l.mSum, r.mSum), l.rSum+r.lSum)
	return res
}

func calcMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type status struct {
	iSum, lSum, rSum, mSum int
}
