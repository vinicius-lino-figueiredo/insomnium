package insomnium

import "testing"

func TestMain(t *testing.T) {
	i := NewInsomnium("/home/vinicius.figueiredo/.config/Insomnium")
	i.Load()
	println("teste")
}
