package filemanager

import (
	"encoding/csv"
	"fmt"
	"os"
)

func GetFileData(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil,err	
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil,err
	}

	//records is a matrix
	if len(records) <= 1 {
		return nil,fmt.Errorf("arquivo %s sem valores",fileName)
	}

	return records[1:],nil
		
}
