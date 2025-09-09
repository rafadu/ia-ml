package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ia-ml/regressao-linear-simples/go-language/internal/filemanager"
	"github.com/ia-ml/regressao-linear-simples/go-language/internal/linearregression"
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
	reader := bufio.NewReader(os.Stdin)
	fileName := flag.String("arquivo","", "Passe o caminho do arquivo csv")
	showTraining := flag.Bool("show",true, "informe se deseja ver o resultado do treinamento")

	flag.Parse()

	if fileName == nil || *fileName == ""{
		log.Fatal("fileName empty...")
	}

	data, err := filemanager.GetFileData(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	convertedData, err  := recordConverter(data)

	if err != nil {
		log.Fatal(err)
	}
	
	linearModel := linearregression.TrainModel(convertedData,showTraining)

	fmt.Println(linearModel.ShowModel())
	fmt.Println(linearModel.ShowRSquare())

	for {
		fmt.Printf("Escreva um valor para ser processado: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "q"{
			fmt.Println("Exiting...")
			break
		}
		
		value, err := strconv.ParseFloat(input,64)

		if err == nil{
			fmt.Printf("Resultado: %.2f\n",linearModel.InferData(value))
		}

	}

}

func recordConverter(records *[][]string) (*[][]float64,error){
	result := make([][]float64,len(*records))
	for i, row := range *records {
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
	return &result,nil
}

