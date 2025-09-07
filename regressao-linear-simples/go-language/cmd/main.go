package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

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

	data, err := filemanager.GetFileData(*fileName)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range data{
		fmt.Println(record)
	}

	fmt.Println("end...")
}

func recordConverter(records [][]string) ([][]float64,error){
	result := make([][]float64,len(records))
	for i, row := range records {
		floatRow := make([]float64, len(row))
		for j, val := range row {
			f, err := strconv.ParseFloat(val,64)
			if err != nil {
				return nil, fmt.Errorf("erro convertendo linha %d, coluna %d: %w",i,j,err)
			}
			floatRow[j] = f
		}
		result[i] = floatRow
	}
	return result,nil
}

