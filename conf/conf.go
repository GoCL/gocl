package conf

type DB struct {
	Url  string
	Port string
	MongoUrl string
}

type Global struct {
	 Port string
	 SignKey string
}

var DBConf *DB = &DB{
	Url: "127.0.0.1",
	Port: "27017",
}


var GlobalConf *Global = &Global{
	Port: "8080",
	SignKey : "sdfsdfdsfsf",
}
