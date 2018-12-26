package blockchain
import (
	"crypto/sha256")
import "strings"	
import b64 "encoding/base64"

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	zeros := []byte {}
	for i := 0 ; i < 33  ; i++ {
		zeros[i] = byte(0)
	}
	genesis_blk := Block {PrevHash: zeros, Generation :0 , Difficulty : difficulty }
	return genesis_blk
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	next_blk := Block{PrevHash : prev_block.Hash , Generation : prev_block.Generation + 1 , Difficulty : prev_block.Difficulty , Data : data }
    return next_blk
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	// TODO
	h_string := b64.StdEncoding.EncodeToString(blk.PrevHash) +":" +  string(blk.Generation) +  ":" +  string(blk.Difficulty) + ":" +  blk.Data + ":" + string(blk.Proof)
	hash := sha256.Sum256([]byte(h_string))
	var hash_res []byte = hash[:] // convert [size]byte to  []byte 
	return hash_res
}


// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	// TODO
	
	hash := string(blk.CalcHash())
	suffix := strings.Repeat("0", int(blk.Difficulty))

	
	return strings.HasSuffix(hash , suffix)


}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
