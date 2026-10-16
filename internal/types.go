package internal


type URL struct {
	Key string `json:"key"`
	Redirect string `json:"redirect"`
}


type Data struct {
	URLS []URL  `json:"urls"`
}
//http:localhost:8080/short-key shd redirect user to website stored for short-key, if the short-key doesnt exist we just show /create-route post request details