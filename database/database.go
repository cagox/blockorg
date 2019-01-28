package database


import (
  "time"
  //Database Stuff
  "github.com/globalsign/mgo"

  "github.com/cagox/blockorg/config"

)

//DialMongoSession starts the main mongo session.
func DialMongoSession() {
  addresses := make([]string, 1)
  addresses[0] = config.Config.MongoServerURL
  info := mgo.DialInfo{
    Addrs: addresses,
    Timeout: 60 * time.Second,
    Database: "blockorg",

    Username: config.Config.MongoUserName,
    Password: config.Config.MongoPassword,
    AppName: "blockorg"}                        //Not sure if AppName is needed.


  session, err := mgo.DialWithInfo(&info)
  if err != nil {
    panic(err)
  }
  session.SetMode(mgo.Monotonic, true)

  config.Config.MongoSession = session
}
