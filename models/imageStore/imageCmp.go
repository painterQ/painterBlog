package imageStore

import (
	"math/big"
)

type ImageCmp struct{}

func (ImageCmp) Compare(a, b []byte) int {
	aa,bb := new(big.Int).SetBytes(a),new(big.Int).SetBytes(b)
	return aa.Sub(aa,bb).Sign()
}

func (ImageCmp) Name() string {
	return "imageCmp"
}

func (ImageCmp) Separator(dst, a, b []byte) []byte {
	return nil
}

func (ImageCmp) Successor(dst, b []byte) []byte {
	return nil
}
