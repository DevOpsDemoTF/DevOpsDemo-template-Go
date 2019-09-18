package handlers

import "github.com/butzist/DevOpsDemo-template-Go/state"

type Handler struct {
	state *state.State
}

func Create(state *state.State) *Handler {
	return &Handler{
		state,
	}
}
