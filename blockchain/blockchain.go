package blockchain


import (
  "encoding/gob"
)


/*
BlockChain is the BlockChain.

It will hold the following fields:
  Count      uint   //How many blocks are in the chain.
  First      string //Hash of the first block on the chain. This will often be "Genesis" or "AccountGenesis"
  Head       string //Hash of the current head of the chain. That is the most recent block.
  Collection string //Which "collection" is this chain concerned with.  This is database stuff.

  Collection, as stated above, will deal with which collection it is in in the database.
  If we continue with mongodb, this will be the collection. If we move to BoltDB, this will be
  the bucket that the data is in. In most cases, Collection will be "root" or "accounts" but may
  be something else if a later application needs to use a different bucket.  For instance, the
  chat chains that we eventually want to implement may create a new named collection chat_identifier
  where identifier is whatever unique identifier the app decides to user.

*/
type BlockChain struct {
  Count      uint
  First      string
  Head       string
  Collection string
}

/*
ChainManager will primarily be a list of BlockChains with functions attached
to carry out various activities.
*/
type ChainManager struct {
  BlockChains map[string]BlockChain
}



func init() {
  gob.Register(BlockChain{})
  gob.Register(ChainManager{})
}
