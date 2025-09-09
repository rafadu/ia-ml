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
	yArray *[]float64
	xArray *[]float64
	trainingResults *[]float64
}

type SimpleLinearRegression struct {
	a float64
	b float64
	rSquare float64
}

func (slr *SimpleLinearRegression) InferData(x float64) float64{
	return slr.a + (slr.b * x)
}


func (slr *SimpleLinearRegression) ShowModel() string{
	return fmt.Sprintf("y = %.1f + %.1fx",slr.a,slr.b)
}

func (slr *SimpleLinearRegression) ShowRSquare() string {
	return fmt.Sprintf("R² = %.2f\n",slr.rSquare)
}

func (slm *SimpleLinearModel) executeModel(x float64) float64{
	return slm.a + (slm.b * x)
}

func (slm *SimpleLinearModel) getRSquare() float64{
	ssr := slm.getSSR()
	sst := slm.getSST()
	return ssr/sst
}

func (slm *SimpleLinearModel) getSSR() float64{
	//somatorio de (alvo - media)²
	somatorium := make([]float64,len(*slm.trainingResults))
	for i, result := range *slm.trainingResults{
		value := result - slm.yMed
		somatorium[i] = value*value
	}
	return vectorSomatorium(&somatorium)
}

func (slm *SimpleLinearModel) getSST() float64{
	//somatorio de (y - media)²
	somatorium := make([]float64,len(*slm.yArray))
	for i, record := range *slm.yArray{
		value := record - slm.yMed
		somatorium[i] = value*value
	}

	return vectorSomatorium(&somatorium)
}

func TrainModel(matrix *[][]float64, showTraining *bool) SimpleLinearRegression{
	var matrixLen = float64(len(*matrix))
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
	
	model.insertTrainingResults(showTraining)

	return SimpleLinearRegression{a: a,b: b,rSquare: model.getRSquare()}
}

func (slm *SimpleLinearModel) insertTrainingResults(show *bool){
	results := make([]float64,len(*slm.xArray))
	for i,record := range *slm.xArray{
		result := slm.executeModel(record)
		results[i] = result
		if(*show){
			fmt.Printf("prediction of %.2f : %.2f\n",record,result)
		}
	}
	slm.trainingResults = &results
	fmt.Println()
}


func getMatrixVector(matrix *[][]float64,index int) *[]float64{
	vector := make([]float64, len(*matrix))

	for i, row := range *matrix{
		for j, val := range row {
			if j == index{
				vector[i] = val
			}	
		}
	}

	return &vector
}


func vectorSquareSomatorium(vector *[]float64) float64{
	var sum float64 = 0
	for _,value := range *vector{
		sum += value*value
	}
	return sum
}

func vectorSomatorium(vector *[]float64) float64 {
	var sum float64 = 0
	for _,value := range *vector{
		sum += value 
	}
	return sum
}

func matrixSomatorium(matrix *[][]float64) float64{
	var sum float64 = 0
	for _, row := range *matrix{
		sum += row[0] * row[1]
	}
	return sum
}
