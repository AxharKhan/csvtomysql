package logic

//"bufio"

func ReadCSV(filename string) (result string) {
	// Open the file
	// csvfile, err := os.Open(filename)
	// if err != nil {
	// 	log.Fatalln("Couldn't open the csv file", err)
	// }

	// // Parse the file
	// r := csv.NewReader(csvfile)
	// //r := csv.NewReader(bufio.NewReader(csvfile))

	// // Iterate through the records
	// for {
	// 	// Read each record from csv
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	o := orm.NewOrm()
	// 	var person models.Person
	// 	var i int64

	// 	person.Firstname = record[0]
	// 	person.LastName = record[1]
	// 	i, err = strconv.ParseInt(record[2], 10, 64)
	// 	if err != nil {
	// 		panic(err.Error)
	// 	}
	// 	person.Age = i
	// 	person.BloodGroup = record[3]

	// 	id, err := o.Insert(&person)

	// 	if err != nil {
	// 		panic(err.Error)
	// 	}

	// 	log.Printf("Question: %s Answer %s\n", id, record[0], record[1])
	// }
	return "Success"
}
