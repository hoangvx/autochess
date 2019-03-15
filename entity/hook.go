package entity

type Sender struct {
	Id				string		`json:"id"`
}

type Recipient struct {
	Id				string		`json:"id"`
}

type Message struct {
	Mid					string 		`json:"mid"`
	Text				string 		`json:"text"`
}

type Messaging struct {
	Sender			Sender 		`json:"sender"`
	Recipient		Recipient `json:"recipient"`
	TimeStamp		int 			`json:"timestamp"`
	Message			Message		`json:"message"`
}

type Entry struct {
	Id				string				`json:"id"`
	Time			int						`json:"time"`
	Messages	[]Messaging		`json:"messaging"`
}

// Hook is facebook webhook models property
// {"object":"page","entry":[{"id":"2166338150070224","time":1552621073035,"messaging":[{"sender":{"id":"1844512938986859"},"recipient":{"id":"2166338150070224"},"timestamp":1552621071721,"message":{"mid":"t3jDSwg5zoeZ6bmXYrE9LJWFqcw6nrjMeQQWKtdovzYkZWVcX4YW4qJmjzHD9sP70gjF5n38gsT3ntDu57A3Lg","seq":774355,"text":"japan"}}]}]}
type Hook struct {
  Object    string   `json:"object"`
	Entries  	[]Entry  `json:"entry"`
}
