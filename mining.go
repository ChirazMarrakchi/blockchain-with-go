package blockchain
import "work_queue"





type miningWorker struct {
	// TODO. Should implement work_queue.Worker
	blc Block   // use this block to pass info
	number uint64
	
	


}


func (obj miningWorker) Run() MiningResult {   // dummy block
 obj.blc.SetProof(obj.number)
    x := MiningResult {}
 if obj.blc.HashValid() {
	 x.Proof = obj.number
	 x.Found = true 

 } 
return x
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	maxJobs := uint(end - start)
	i := start
	x := MiningResult{}
	// type assertion of Run function in work_queue
	
	m := miningWorker{}
	
   for(i <= end ) {    // two conditions to stop : i =< end && Mining results were found  
	for j := i  ; j < chunks ; j += workers {// give nworkers small chuncks of work
		
	m = miningWorker {blc : blk , number : uint64(j)} //  .Run() from work_queue works as .Run()from mining                                      
	
	work_queue.Run() = m.Run()				//i.e blk and pow infos are needed                                      
											  
	q := Create(uint(workers) , maxJobs )        // in every j loop , every worker performs nworkers jobs 
		
	                                              //thus we increment counter by number of jobs done = nWokers
     // check if we get mining results ? if yes we stop and break from the loop 
	
	
	
	m :=  <-q.Results
	 
	x = m.(MiningResult)
	if x.Found == true {
		blk.SetProof(x.Proof)
		blk.Hash = blk.CalcHash()

	}

}
i += chunks   // first or 2nd chunck is done ? .. we move to the next chunck 
}	
return x
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << blk.Difficulty) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}

