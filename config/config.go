//Package config handles the configuraion for the BlockOrg node
package config



import (
  "os"
  "encoding/json"

  "github.com/mitchellh/go-homedir"
  "github.com/gorilla/mux"
  "github.com/globalsign/mgo"

)

//Config holds the system configuration.
var Config *BlockOrgConfiguration


//GodGameConfiguration holds the configuration information for the GGE program.
type BlockOrgConfiguration struct {
  StaticPath        string
  TemplateRoot      string
  AdminToken        string
  ManifestoPath     string
  MongoUserName     string
  MongoPassword     string
  MongoServerURL    string
  EncKey            string
  AuthKey           string

  //The items bellow are set by the application itself.
  MongoSession      *mgo.Session
  Router            *mux.Router
}

func init() {
  LoadConfiguration()
}



//LoadConfiguration Loads the configuration and sets the global variable.
func LoadConfiguration() {
  //TODO: Add error handling.
  Config = GetConfigs()
  Config.Router = mux.NewRouter()
}

//GetConfigs returns a configuration struct.
func GetConfigs() *BlockOrgConfiguration {
  configpath := getConfigPath()
  ensureDirectory(configpath)

  configuration := BlockOrgConfiguration{}

  file, err := os.Open(configpath+"/blockorg.json")

  //If there is an error (if the file doesn't exist or is malformed) we will created a default  configuration.
  if err != nil {
    //return defaultConfiguration(configpath)
    panic("Configuration File Did Not Open")
  }

  decoder := json.NewDecoder(file)
  err = decoder.Decode(&configuration)

  if err != nil {
    //return defaultConfiguration(configpath)
    panic("Configuration File Malformed.")
  }

  return &configuration
}

//ensureDirectory() makes sure the given directory exists or return an error.
func ensureDirectory(base string) error {
  _, statErr := os.Stat(base)

  if statErr != nil {
    createErr := os.MkdirAll(base, 0755)
    if createErr != nil {
      panic("Could not ensure the configuration path.")
    }
  }

  return nil
}

//getConfigPath() Gets the config path from the environmental variable or uses a default.
func getConfigPath() string {
  //First we figure out where the configuration files should be.
  configpath, isEnv := os.LookupEnv("BLOCKORGROOT")
  if !isEnv {
    home, _ := homedir.Dir() //TODO: Add error checking.
    configpath = home+"/.config/blockorg"
  }
  return configpath
}
