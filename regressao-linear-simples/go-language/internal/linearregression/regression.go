package linearregression

import "fmt"

/*
Regressão linear simples
y = a + bx

a = (media)y - b(media)x

b = n Σxy - (Σx)(Σy)
   ------------------
    n Σx² - (Σx)²

(media)z = Σz/n
*/

type SimpleLinearModel struct {
	a float64
	b float64
	yMed float64
	yArray []float64
	xArray []float64
	trainingResults []float64
}

func (slm *SimpleLinearModel) ExecuteModel(x float64) float64{
	return slm.a + (slm.b * x)
}

func (slm *SimpleLinearModel) ShowModel() string{
	return fmt.Sprintf("y = %.1f + %.1fx",slm.a,slm.b)
}

func (slm *SimpleLinearModel) GetRSquare() string{
	ssr := slm.getSSR(slm.trainingResults)
	sst := slm.getSST()
	rSquare := ssr/sst
	return fmt.Sprintf("R² = %.2f",rSquare)
}

func (slm *SimpleLinearModel) getSSR(results []float64) float64{
	//somatorio de (alvo - media)²
	somatorium := make([]float64,len(results))
	for i, result := range results{
		value := result - slm.yMed
		somatorium[i] = value*value
	}
	return vectorSomatorium(somatorium)
}

func (slm *SimpleLinearModel) getSST() float64{
	//somatorio de (y - media)²
	somatorium := make([]float64,len(slm.yArray))
	for i, record := range slm.yArray{
		value := record - slm.yMed
		somatorium[i] = value*value
	}

	return vectorSomatorium(somatorium)
}

func TrainModel(matrix [][]float64) SimpleLinearModel{
	var matrixLen = float64(len(matrix))
	//get x array
	xArray := getMatrixVector(matrix,0)
	//get y array
	yArray := getMatrixVector(matrix,1)
	//get x somatorium 
	xSum := vectorSomatorium(xArray)
	//get y somatorium 
	ySum := vectorSomatorium(yArray)
	//get x² somatorium 
	xSqrSum := vectorSquareSomatorium(xArray)
	//get xy somatorium
	xySum := matrixSomatorium(matrix)
	//get x median
	var xMed = xSum/matrixLen
	//get y median 
	var yMed = ySum/matrixLen
	
	b := ((matrixLen*xySum) - (xSum*ySum))/((matrixLen*xSqrSum)-(xSum*xSum))
	
	a := yMed - (b*xMed)
	
	model := SimpleLinearModel{a: a,b: b,yMed: yMed,yArray: yArray, xArray: xArray}
	
	model.insertTrainingResults()

	return model
}

func (slm *SimpleLinearModel) insertTrainingResults(){
	results := make([]float64,len(slm.xArray))
	for i,record := range slm.xArray{
		result := slm.ExecuteModel(record)
		results[i] = result
		fmt.Printf("prediction of %.2f : %.2f\n",record,result)
	}
	slm.trainingResults = results
}


func getMatrixVector(matrix [][]float64,index int) []float64{
	vector := make([]float64, len(matrix))

	for i, row := range matrix{
		for j, val := range row {
			if j == index{
				vector[i] = val
			}	
		}
	}

	return  vector
}


func vectorSquareSomatorium(vector []float64) float64{
	var sum float64 = 0
	for _,value := range vector{
		sum += value*value
	}
	return sum
}

func vectorSomatorium(vector []float64) float64 {
	var sum float64 = 0
	for _,value := range vector{
		sum += value 
	}
	return sum
}

func matrixSomatorium(matrix [][]float64) float64{
	var sum float64 = 0
	for _, row := range matrix{
		sum += row[0] * row[1]
	}
	return sum
}
