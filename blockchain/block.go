package blockchain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hjkimGithub/nomadcoin/db"
	"github.com/hjkimGithub/nomadcoin/utils"
)

type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevHash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"dificulty"`
	Nonce      int    `json:"nounce"`
	Timestamp  int    `json:"timestamp"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToByetes(b))
}

var ErrNotFound = errors.New("BLOCK NOT FOUND")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		fmt.Printf("\n\n\nTarget:%s\nHash:%s\nNonce:%d\n\n\n", target, hash, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}

	}
}

func createBlock(data string, prevHash string, height int) *Block {
	block := Block{
		Data:       data,
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: BlockChain().difficulty(),
		Nonce:      0,
	}
	block.mine()
	block.persist()
	return &block
}
