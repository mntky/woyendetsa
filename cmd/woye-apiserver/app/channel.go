package app

import (
)

func NewChannel() <-chan {
	ch := make(chan struct{})
	return ch
}
