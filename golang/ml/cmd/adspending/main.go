package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/anton-kapralov/experimental/golang/ml/gradientdescent"
)

func main() {
	filepath := flag.String("data", "", "Path to a CSV file with the training data")
	spendingX := flag.Float64("spending", 0, "Value of spending to predict for")
	flag.Parse()
	if *filepath == "" {
		flag.Usage()
		return
	}

	f, err := os.Open(*filepath)
	if err != nil {
		log.Fatal(err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()

	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var spendings []float64
	var sales []float64
	for _, record := range records {
		spending, _ := strconv.ParseFloat(record[2], 64)
		spendings = append(spendings, spending)
		sale, _ := strconv.ParseFloat(record[4], 64)
		sales = append(sales, sale)
	}

	model := gradientdescent.NewModel(0, 0)
	model.Fit(spendings, sales, 0.001, 15000)

	fmt.Println(model)
	fmt.Println(model.Predict(*spendingX))
}
