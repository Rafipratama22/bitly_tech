package config

import "github.com/gomodul/envy"

type server struct {
	PORTHTTP  string
}

var Server = server{
	PORTHTTP: ":" + envy.Get("PORT_HTTP", "8081"),
}