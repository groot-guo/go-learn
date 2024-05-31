package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	if isBadVersion(n) && !isBadVersion(n-1) {
		return n
	}
	return firstBadVersion(n - 1)
}

func check(start, end int) int {
	mid := (start + end) / 2

	if isBadVersion(mid) && !isBadVersion(mid-1) {
		return mid
	}
	res1 := check(start, mid)
	if res1 != 0 {
		return res1
	}
	res2 := check(mid, start)
	if res2 != 0 {
		return res2
	}
	return 0
}
