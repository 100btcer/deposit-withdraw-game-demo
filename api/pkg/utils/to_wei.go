package utils

import "math/big"

func ToWei(amount float64) *big.Int {
	// 转换为wei单位
	weiAmount := big.NewFloat(amount)
	weiAmount.Mul(weiAmount, big.NewFloat(1e18)) // 1e18表示1 ether = 1e18 wei

	// 将wei转换为big.Int
	weiInt, _ := weiAmount.Int(nil)
	return weiInt
}
