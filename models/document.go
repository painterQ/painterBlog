package models

import "bytes"

type DocumentComparer struct{}

func getPrefixLevel(in []byte) int {
	if bytes.HasPrefix(in, matePreFix) {
		return 1
	}
	if bytes.HasPrefix(in, tagPreFix) {
		return 2
	}
	return 0
}

func (DocumentComparer) Compare(a, b []byte) int {
	i, j := getPrefixLevel(a), getPrefixLevel(b)
	if i == j {
		return bytes.Compare(a, b)
	}
	if i > j {
		return 1
	}

	return -1

}

func (DocumentComparer) Name() string {
	return "documentComparator"
}

func (DocumentComparer) Separator(dst, a, b []byte) []byte {
	i, n := 0, len(a)
	if n > len(b) {
		n = len(b)
	}
	for ; i < n && a[i] == b[i]; i++ {
	}
	if i >= n {
		// Do not shorten if one string is a prefix of the other
	} else if c := a[i]; c < 0xff && c+1 < b[i] {
		dst = append(dst, a[:i+1]...)
		dst[len(dst)-1]++
		return dst
	}
	return nil
}

func (DocumentComparer) Successor(dst, b []byte) []byte {
	for i, c := range b {
		if c != 0xff {
			dst = append(dst, b[:i+1]...)
			dst[len(dst)-1]++
			return dst
		}
	}
	return nil
}
