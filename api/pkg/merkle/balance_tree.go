package merkle

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func ToNode(account common.Address, token common.Address, amount *big.Int, id int) common.Hash {
	var h common.Hash
	sha3 := solsha3.SoliditySHA3(
		[]string{"address", "address", "uint256", "int256"},
		[]interface{}{account, token, amount, id},
	)
	copy(h[:], sha3)
	return h
}

func VerifyProof(account common.Address, token common.Address, amount *big.Int, id int, proof Elements, root common.Hash) bool {
	pair := ToNode(account, token, amount, id)
	for _, item := range proof {
		pair = combinedHash(pair, item)
	}

	return pair == root
}

type Balance struct {
	Account       common.Address
	TokenContract common.Address
	Amount        *big.Int
	Id            int
}

type BalanceTree struct {
	tree *MerkleTree
}

func NewBalanceTree(balances []Balance) (*BalanceTree, error) {
	elements := make(Elements, 0, len(balances))
	for _, balance := range balances {
		elements = append(elements, ToNode(balance.Account, balance.TokenContract, balance.Amount, balance.Id))
	}

	tree, err := NewMerkleTree(elements)
	if err != nil {
		return nil, err
	}

	return &BalanceTree{tree: tree}, nil
}

func (b *BalanceTree) GetRoot() common.Hash {
	return b.tree.GetRoot()
}

func (b *BalanceTree) GetProof(account common.Address, token common.Address, amount *big.Int, id int) ([]common.Hash, error) {
	return b.tree.GetProof(ToNode(account, token, amount, id))
}
