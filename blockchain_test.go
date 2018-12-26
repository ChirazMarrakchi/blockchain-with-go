package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"strings"
	"encoding/hex"
)


func TestCalcHash(t *testing.T) {
	zeros := []byte {}
	for i := 0 ; i < 33  ; i++ {
		zeros[i] = byte(0)
	}
	
	s1 := Block {PrevHash : zeros , Generation :0 , Difficulty : 16 , Proof : 56231 }
	res := hex.EncodeToString(s1.CalcHash())
	acres := "6c71ff02a08a22309b7dbbcee45d291d4ce955caa32031c50d941e3e9dbd0000"
	if res != acres {
		t.Error("Wrong CalcHash result:", acres)
	}

	
}

func TestValidHashes(t *testing.T){

	zeros := []byte {}
	for i := 0 ; i < 33  ; i++ {
		zeros[i] = byte(0)
	}

	blc := Block {PrevHash: zeros ,  Generation :0 , Difficulty : 16 , Proof : 56231 }
	res := blc.ValidHash()

	if ( res == false    ){
		t.Error( "WRONG ValidHash .. should return :" , true)
	}

	blc1 := Block {PrevHash: zeros ,  Generation :0 , Difficulty : 16 , Proof : 56230 }
	res1 := blc1.ValidHash()

	if ( res1 == true    ){
		t.Error( "WRONG ValidHash .. should return :" , false)
	}

}

func TestMine( t *testing.T){
	zeros := []byte {}
	for i := 0 ; i < 33  ; i++ {
		zeros[i] = byte(0)
	}

	blc := Block {PrevHash: zeros ,  Generation :0 , Difficulty : 7  }
	res := blc.Mine(1)
	if blc.Proof != uint64(385) {
		t.Error("Wrong Mine .. should return :" , 385 )
	}

}
