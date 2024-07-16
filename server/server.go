package main


import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main(){
	http.HandleFunc("/",getRoot)
	http.HandleFunc("/hello",getHello)

	err := http.ListenAndServe(":3333",nil)

	if errors.Is(err, http.ErrServerClosed){
		fmt.Printf("server closed\n")
	}else if err != nil {
		fmt.Printf("error starting the server:  %s\n",err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "got / request\n")
	io.WriteString(w,"This is my Website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "got /hello request\n")
	io.WriteString(w,"Hello, HTTP!\n")
}

