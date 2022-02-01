package repository

import (
	"github.com/marcos-wz/capstone/internal/entity"
)

// DOMAIN ****************************************************

type iWriterFruitsRepo interface {
	// Write new fruits records to the csv file
	WriteFruits(fruits []entity.Fruit) error
}

type writerFruitsRepo struct{}

func NewWriterFruitsRepo() iWriterFruitsRepo {
	return &writerFruitsRepo{}
}

// IMPLEMENTATION ********************************************

func (writerFruitsRepo) WriteFruits(fruits []entity.Fruit) error {
	// f, err := os.OpenFile(csvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Println("ERROR:", err)
	// 	return err
	// }
	// defer f.Close()
	// // Parse fruit struct to string array
	// numFields := 10
	// // NOTE: check if dynamic legth make sense
	// // t := reflect.TypeOf(*fruit)
	// // numFields := t.NumField()

	// fruitRecord := make([]string, numFields)
	// fruitRecord[0] = strconv.Itoa(fruit.ID)
	// fruitRecord[1] = fruit.Name
	// fruitRecord[2] = fruit.Description
	// fruitRecord[3] = fruit.Color
	// fruitRecord[4] = fruit.Unit
	// // fruitRecord[5] = fmt.Sprintf("%f", fruit.Price)
	// fruitRecord[5] = fmt.Sprintf("%.2f", fruit.Price)
	// fruitRecord[6] = strconv.Itoa(fruit.Stock)
	// fruitRecord[7] = strconv.Itoa(fruit.Caducate)
	// fruitRecord[8] = fruit.Country
	// fruitRecord[9] = fruit.CreateAt.Format(time.RFC3339)

	// log.Printf("Parsed fruit record: %+v", fruitRecord)
	// // Write record
	// csvWriter := csv.NewWriter(f)
	// if err := csvWriter.Write(fruitRecord); err != nil {
	// 	log.Println("ERROR:", err)
	// 	return err
	// }
	// csvWriter.Flush()
	return nil
}
