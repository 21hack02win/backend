package ws

import (
	"github.com/21hack02win/nascalay-backend/model"
	"github.com/21hack02win/nascalay-backend/usecases/repository"
)

type Hub struct {
	repo           repository.Repository
	userIdToClient map[model.UserId]*Client
	registerCh     chan *Client
	unregisterCh   chan *Client
}

func NewHub(repo repository.Repository) *Hub {
	return &Hub{
		repo:           repo,
		userIdToClient: make(map[model.UserId]*Client),
		registerCh:     make(chan *Client),
		unregisterCh:   make(chan *Client),
	}
}

func (h *Hub) Register(client *Client) {
	h.registerCh <- client
}
func (h *Hub) Unregister(client *Client) {
	h.unregisterCh <- client
}

func (h *Hub) Run() {
	for {
		select {
		case cli := <-h.registerCh:
			h.register(cli)
		case cli := <-h.unregisterCh:
			h.unregister(cli)
		}
	}
}

func (h *Hub) register(cli *Client) {
	h.userIdToClient[cli.userId] = cli
}

func (h *Hub) unregister(cli *Client) {
	close(cli.send)
	delete(h.userIdToClient, cli.userId)
}
