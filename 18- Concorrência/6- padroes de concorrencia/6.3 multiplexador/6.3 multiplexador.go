package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá, mundo"), escrever("programando em go"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}
func multiplexar(canaldeentrada1, canaldeentrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canaldeentrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canaldeentrada2:
				canalDeSaida <- mensagem
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return canal
}
