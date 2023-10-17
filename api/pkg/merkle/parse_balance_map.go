package merkle

import (
	"github.com/ethereum/go-ethereum/common"
)

type Claim struct {
	Address string        `json:"address"`
	Amount  string        `json:"amount"`
	Id      int           `json:"id"`
	Proof   []common.Hash `json:"proof"`
}

type MerkleDistributorInfo struct {
	MerkleRoot common.Hash `json:"merkleRoot"`
	TokenTotal string      `json:"tokenTotal"`
	Claims     []Claim     `json:"claims"`
}

func ParseBalanceMap(balances []Balance) (MerkleDistributorInfo, error) {
	info := MerkleDistributorInfo{
		Claims: make([]Claim, 0),
	}

	tree, err := NewBalanceTree(balances)
	if err != nil {
		return info, err
	}

	//tokenTotal := big.NewInt(0)
	for _, balance := range balances {
		proof, err := tree.GetProof(balance.Account, balance.TokenContract, balance.Amount, balance.Id)
		if err != nil {
			return info, err
		}
		//tokenTotal = big.NewInt(0).Add(tokenTotal, balance.Amount)

		info.Claims = append(info.Claims, Claim{
			Address: balance.Account.String(),
			Amount:  "0x" + balance.Amount.Text(16),
			Id:      balance.Id,
			Proof:   proof,
		})
	}

	info.MerkleRoot = tree.GetRoot()
	//info.TokenTotal = "0x" + tokenTotal.Text(16)

	return info, nil
}
