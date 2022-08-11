package transport

import (
	"bitly_backend/app/interface/container"
)

type Transport struct {
	Transporter *transport
}

func SetupTransport(container *container.Container) *Transport {
	tp := NewTransport(container.Usecase)
	return &Transport{
		tp,
	}
}