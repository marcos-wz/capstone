package repository

//
//import (
//	pb "github.com/marcos-wz/capstone/internal/fruit"
//	"log"
//	"strconv"
//)
//
//// Parser Fruit Record function, guarantee csv data integrity. It parse from csv records to fruit instance
//// This is an input data method. Always returns a fruit instance.
//// If an error occurs, the default type value is set.
//// Parameters: 1) A record string array from csv file, 2) The number of fruit entity struct fields
//// Returns a fruit entity instance and parse validation errors
//func (*fruitRepo) parseFruitCSV(record []string, numFields int) (*pb.Fruit, []error) {
//	// Values initialization
//	fruit := &pb.Fruit{}
//	values := make([]string, numFields)
//	copy(values, record)
//	errors := make([]error, numFields)
//
//	// Fruit Instance
//	// 0 - ID
//	// fruit.Id, _ = strconv.Atoi(values[0])
//	ID, err := strconv.ParseUint(values[0], 10, 32)
//	if err != nil {
//		errors[0] = err
//		log.Printf("ERROR-Parser: parsing ID %q: %v", ID, err)
//	}
//	fruit.Id = uint32(ID)
//
//	// 1 - NAME
//	fruit.Name = values[1]
//	if fruit.Name == "" {
//
//	}
//	// 2 - DESCRIPTION
//	fruit.Description = values[2]
//	// 3 - COLOR
//	fruit.Color = values[3]
//	// 4 - UNIT
//	//fruit.Unit = values[4]
//	// 5 - PRICE
//	//fruit.Price, _ = strconv.ParseFloat(values[5], 64)
//	// 6 - STOCK
//	//fruit.Stock, _ = strconv.Atoi(values[6])
//	// 7 - CADUCATE
//	//fruit.CaducateDays, _ = strconv.Atoi(values[7])
//	// 8 - COUNTRY
//	fruit.Country = values[8]
//	// 9 - CREATED AT
//	//fruit.CreatedAt, _ = time.Parse(time.RFC3339, values[9])
//
//	return fruit, errors
//}
