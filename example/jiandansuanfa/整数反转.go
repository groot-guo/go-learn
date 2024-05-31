package main

func reverseInt(x int) int {
	res := int32(0)
	for x != 0 {
		t := x % 10
		newRes := res*10 + int32(t)
		//如果数字溢出，直接返回0
		if (newRes-int32(t))/10 != res {
			return 0
		}
		res = newRes
		x = x / 10
	}
	return int(res)
}

func reverse2(x int) int {
	k := 0
	for x != 0 {
		if i := int32(k); (10*i)/10 != i {
			return 0
		}
		k = k*10 + x%10
		x = x / 10
	}
	return k
}
