package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"ppd/t2/internal/admin"
)

func main() {

	/*
		Para que seja possivel acessar os metodos do objeto
		eh necessario registra-lo utilizando a biblioteca rpc.
		O registro gera um erro, sendo nil o caso em que o registro
		foi um sucesso.
	*/
	bank := new(admin.Bank)

	err := rpc.Register(bank)
	if err != nil {
		log.Fatal("Error registering Bank ", err)
	}
	//Permite que a biblioteca utilize http para comunicacao
	rpc.HandleHTTP()

	/*
		Inicializa um processo que escuta toda comunicacao em
		determinada porta, seguindo o protocolo tcp
	*/
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error ", err)
	}
	log.Printf("Serving rpc on port: %d", 4040)

	/*
		Ativa o servidor na porta e com o protocolo definido
		pelo listener
	*/
	http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
