package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ia-ml/regressao-linear-simples/go-language/internal/filemanager"
)

/*
Regressão linear simples
y = a + bx

a = (media)y - b(media)x

b = n Σxy - (Σx)(Σy)
   ------------------
    n Σx² - (Σx)²

(media)z = Σz/n
*/

func main(){
	fileName := flag.String("arquivo","", "Passe o caminho do arquivo csv")

	flag.Parse()

	if fileName == nil || *fileName == ""{
		log.Fatal("fileName empty...")
	}

	filemanager.GetFileData(*fileName)

	fmt.Println("end...")
}
