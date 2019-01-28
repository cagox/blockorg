package blocks


import (
  "time"
  "encoding/gob"

  "github.com/globalsign/mgo/bson"

)

/*
Block is the basic block stucture.
*/
type Block struct {
  ID                  bson.ObjectId `bson:"_id,omitempty"` //This is DataBase nonsense.
  TimeStamp           time.Time //Time that the block was created.
  PreviousHash        string    //Hash of the previous entry
  //Target                      //This will be either added later, or part of extended types.
  BlockType           string    //What type of block is this?
  Payload             bson.M
  Originator          string    //It is yet to be determined if this will be a username or a hash.
  OriginatorSignature string    //Originator's cryptographic signature on the payload.
  Hash                string    //A Hash of everything above, except for the ID.
}






func init() {
  gob.Register(Block{})
}
