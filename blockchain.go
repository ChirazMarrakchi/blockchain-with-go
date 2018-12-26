package blockchain
import("strings")
import "reflect"
type Blockchain struct {
	Chain []Block
}


// this function aims to add a block to a chain 
func (chain *Blockchain) Add(blk Block) {
	// You can remove the panic() here if you wish.
	
	
	
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}

	if (chain == nil ) {
		blc := Block {}
		*chain = Blockchain {blc.Initial(blk.Difficulty) }
	} else {
		
		
		*chain.Next(blk.Data)
	}

	
}

func (chain Blockchain) IsValid() bool {
	blc := [] byte {}
// set a flag for each condition
// The initial block has previous hash all null bytes and is generation zero.
     zeros := []byte {}
     for i := 0 ; i < 33  ; i++ {
	 zeros[i] = byte(0)
     }
	res1 := (chain.Chain[0].Generation == 0) && (chain.Chain[0].PrevHash == zeros)
	res2 := true
	res3 := true
	res4 := true
	res5 := true 
	gen := uint64(0)
	//Each block has the same difficulty value.
	for pos , val := range (chain.Chain) {
		if (val.Difficulty != chain.Chain[0].Difficulty){
			res2 = false
			
		}
		//Each block has a generation value that is one more than the previous block.
		if (val.Generation != gen) {
			res3 = false
		}
		// gen = gen +1 
		// Each block's hash value ends in difficulty null bits.
		gen = gen +1 
		if (!val.ValidHash()){
			res4 = false
		}
		

	}
	




//Each block's previous hash matches the previous block's hash.
	
 prev_blc := chain.Chain[0]
	n := len(chain.Chain)
	for i := 1 ; i < n ; i++  {
		if !reflect.DeepEqual(prev_blc.Hash, chain.Chain[i].PrevHash){
			res5 = false
		}
		prev_blc = chain.Chain[i]

	}
//Each block's hash value actually matches its contents.
	res6 := true
	j := 0 
	for pos , val := range (chain.Chain){
		if ! reflect.DeepEqual(chain.Chain[j].Hash , val){
			res6 = false
		}
		j +=1

	}
	
	
	res := res1 && res2 && res3 && res4 && res5 && res6
	return res
}

