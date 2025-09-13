package linearregression

import "fmt"

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

