package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Block representa um bloco da blockchain
type Block struct {
	Id           uuid.UUID // Identificador único do nó
	Hash         string    // Hash do nó
	Value        string    // Valor do nó
	Nonce        int       // Número de uso único
	PreviousHash string    // Hash do bloco anterior
	Timestamp    time.Time // Timestamp do bloco
}

// NewBlock cria um novo bloco
func NewBlock(value string, previousNodes []*Block) *Block {
	id, _ := uuid.NewV7()
	block := &Block{
		Id:           id,
		Hash:         "",
		Value:        value,
		Nonce:        0,
		PreviousHash: previousNodes[len(previousNodes)-1].Hash,
		Timestamp:    time.Now(),
	}

	block.Hash = block.CalculateHash()

	return block
}

// CalculateHash calcula o hash do bloco
func (b *Block) CalculateHash() string {
	data := fmt.Sprintf("%s%s%s%d%s",
		b.Id.String(),
		b.Value,
		b.Timestamp.String(),
		b.Nonce,
		b.PreviousHash)

	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// GetId retorna o identificador do bloco
func (b *Block) GetId() uuid.UUID {
	return b.Id
}

// GetHash retorna o hash do bloco
func (b *Block) GetHash() string {
	return b.Hash
}

// GetPreviousHash retorna o hash do bloco anterior
func (b *Block) GetPreviousHash() string {
	return b.PreviousHash
}

// GetValue retorna o valor do bloco
func (b *Block) GetValue() string {
	return b.Value
}

// GetTimestamp retorna o timestamp do bloco
func (b *Block) GetTimestamp() time.Time {
	return b.Timestamp
}

// Blockchain representa a blockchain
type Blockchain struct {
	sync.RWMutex          // Mutex para proteger a blockchain
	Chain        []*Block // Cadeia de blocos
	Difficulty   int      // Dificuldade do consenso
}

// AddBlock adiciona um bloco à blockchain
func (bc *Blockchain) AddBlock(block *Block) {
	bc.Lock()
	defer bc.Unlock()

	if len(bc.Chain) > 0 {
		lastBlock := bc.Chain[len(bc.Chain)-1]
		if block.PreviousHash != lastBlock.Hash {
			return
		}
	}

	bc.MineBlock(block)

	bc.Chain = append(bc.Chain, block)
}

// MineBlock mina um bloco
func (bc *Blockchain) MineBlock(block *Block) {
	target := strings.Repeat("0", bc.Difficulty)

	for {
		hash := block.CalculateHash()
		if strings.HasPrefix(hash, target) {
			block.Hash = hash
			return
		}
		block.Nonce++
	}
}

// GetChain retorna a cadeia de blocos
func (bc *Blockchain) GetChain() []*Block {
	return bc.Chain
}

// GetDifficulty retorna a dificuldade do consenso
func (bc *Blockchain) GetDifficulty() int {
	return bc.Difficulty
}

// GetLastBlock retorna o último bloco da cadeia
func (bc *Blockchain) GetLastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

// GetBlockByHash retorna um bloco pelo hash
func (bc *Blockchain) GetBlockByHash(hash string) *Block {
	for _, block := range bc.Chain {
		if block.Hash == hash {
			return block
		}
	}
	return nil
}

// GetBlockByIndex retorna um bloco pelo índice
func (bc *Blockchain) GetBlockByIndex(index int) *Block {
	return bc.Chain[index]
}

// CreateGenesisBlock cria o bloco genesis
func (bc *Blockchain) CreateGenesisBlock() *Block {
	if len(bc.Chain) > 0 {
		return nil
	}

	genesisBlock := NewBlock("Genesis Block", []*Block{})
	genesisBlock.Hash = genesisBlock.CalculateHash()
	genesisBlock.PreviousHash = ""
	genesisBlock.Timestamp = time.Now()

	bc.AddBlock(genesisBlock)

	return genesisBlock
}
