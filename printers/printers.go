package printers

import (
	"bufio"
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"

	"github.com/qwertypomy/printers/dao/factory"
	"github.com/qwertypomy/printers/models"
)

var brandN, technologyN, functionTypeN, printSizeN, resolutionXN, resolutionYN, nameN, descriptionN,
	additionalInfoN, amountN, weightN, pagePerMinuteN, priceN, sizeN int

func Populate(config models.Config) (err error) {
	factoryDao := factory.FactoryDao{Engine: config.Engine}
	printerDao := factoryDao.GetPrinterDaoInterface()

	arr, err := ReadCSV()
	if err != nil {
		return
	}

	brands, technologies, functionTypes, sizes, resolutions, err := ArrToObjects(arr)

	for _, brand := range brands {
		err = printerDao.CreateBrand(&brand)
		if err != nil {
			return
		}
	}
	for _, technology := range technologies {
		err = printerDao.CreatePrintingTechnology(&technology)
		if err != nil {
			return
		}
	}
	for _, functionType := range functionTypes {
		err = printerDao.CreateFunctionType(&functionType)
		if err != nil {
			return
		}
	}
	for _, size := range sizes {
		err = printerDao.CreatePrintSize(&size)
		if err != nil {
			return
		}
	}
	for _, resolution := range resolutions {
		err = printerDao.CreatePrintResolution(&resolution)
		if err != nil {
			return
		}
	}

	printers, err := ArrToPrinters(arr, config)

	for _, printer := range printers {
		err = printerDao.CreatePrinter(&printer)
		if err != nil {
			return
		}
	}

	return
}

func ReadCSV() (arr [][]string, err error) {
	path, _ := filepath.Abs("./printers/printers.csv")
	f, err := os.Open(path)
	if err != nil {
		return
	}
	r := csv.NewReader(bufio.NewReader(f))
	arr, err = r.ReadAll()
	if err != nil {
		return
	}

	for i, name := range arr[0] {
		switch name {
		case "Brand":
			brandN = i
		case "PrintingTechnology":
			technologyN = i
		case "FunctionType":
			functionTypeN = i
		case "PrintSize":
			printSizeN = i
		case "PrintResolutionX":
			resolutionXN = i
		case "PrintResolutionY":
			resolutionYN = i
		case "Name":
			nameN = i
		case "Description":
			descriptionN = i
		case "AdditionalInfo":
			additionalInfoN = i
		case "Amount":
			amountN = i
		case "Weight":
			weightN = i
		case "PagePerMinute":
			pagePerMinuteN = i
		case "Price":
			priceN = i
		case "Size":
			sizeN = i
		}
	}

	return
}

func ArrToObjects(arr [][]string) (brands []models.Brand,
	technologies []models.PrintingTechnology,
	functionTypes []models.FunctionType,
	sizes []models.PrintSize,
	resolutions []models.PrintResolution, err error) {

	// Fill brands, technologies, functionTypes, sizes, resolutions
	for _, printer := range arr[1:] {
		unique := true

		temp := printer[brandN]
		for _, v := range brands {
			if v.Name == temp {
				unique = false
				break
			}
		}
		if unique {
			brands = append(brands, models.Brand{Name: temp})
		}
		unique = true

		temp = printer[technologyN]
		for _, v := range technologies {
			if v.Name == temp {
				unique = false
				break
			}
		}
		if unique {
			technologies = append(technologies, models.PrintingTechnology{Name: temp})
		}
		unique = true

		temp = printer[functionTypeN]
		for _, v := range functionTypes {
			if v.Name == temp {
				unique = false
				break
			}
		}
		if unique {
			functionTypes = append(functionTypes, models.FunctionType{Name: temp})
		}
		unique = true

		temp = printer[printSizeN]
		for _, v := range sizes {
			if v.Name == temp {
				unique = false
				break
			}
		}
		if unique {
			sizes = append(sizes, models.PrintSize{Name: temp})
		}
		unique = true
		var tempx int
		tempx, err = strconv.Atoi(printer[resolutionXN])
		if err != nil {
			return
		}
		tempX := uint(tempx)
		var tempy int
		tempy, err = strconv.Atoi(printer[resolutionYN])
		if err != nil {
			return
		}
		tempY := uint(tempy)
		for _, v := range resolutions {
			if v.X == tempX && v.Y == tempY {
				unique = false
				break
			}
		}
		if unique {
			resolutions = append(resolutions, models.PrintResolution{X: tempX, Y: tempY})
		}
		unique = true
	}
	return
}

func ArrToPrinters(arr [][]string, config models.Config) (printers []models.Printer, err error) {
	factoryDao := factory.FactoryDao{Engine: config.Engine}
	printerDao := factoryDao.GetPrinterDaoInterface()

	brands, err := printerDao.BrandList()
	if err != nil {
		return
	}
	technologies, err := printerDao.PrintingTechnologyList()
	if err != nil {
		return
	}
	functionTypes, err := printerDao.FunctionTypeList()
	if err != nil {
		return
	}
	printSizes, err := printerDao.PrintSizeList()
	if err != nil {
		return
	}
	resolutions, err := printerDao.PrintResolutionList()
	if err != nil {
		return
	}
	for _, printer := range arr[1:] {
		var amount int
		amount, err = strconv.Atoi(printer[amountN])
		if err != nil {
			return
		}
		var weight float64
		weight, err = strconv.ParseFloat(printer[weightN], 32)
		if err != nil {
			return
		}
		var ppm float64
		ppm, err = strconv.ParseFloat(printer[pagePerMinuteN], 32)
		if err != nil {
			return
		}
		var price int
		price, err = strconv.Atoi(printer[priceN])
		if err != nil {
			return
		}

		var brand, technology, functionType, printSize string
		var resolutionX, resolutionY uint

		for _, v := range brands {
			if v.Name == printer[brandN] {
				brand = v.ID
			}
		}
		for _, v := range technologies {
			if v.Name == printer[technologyN] {
				technology = v.ID
			}
		}
		for _, v := range functionTypes {
			if v.Name == printer[functionTypeN] {
				functionType = v.ID
			}
		}
		for _, v := range printSizes {
			if v.Name == printer[printSizeN] {
				printSize = v.ID
			}
		}
		for _, v := range resolutions {
			var tempx int
			tempx, err = strconv.Atoi(printer[resolutionXN])
			if err != nil {
				return
			}
			tempX := uint(tempx)
			var tempy int
			tempy, err = strconv.Atoi(printer[resolutionYN])
			if err != nil {
				return
			}
			tempY := uint(tempy)
			if v.X == tempX && v.Y == tempY {
				resolutionX = v.X
				resolutionY = v.Y
			}
		}

		printers = append(printers, models.Printer{
			Name:                 printer[nameN],
			Description:          printer[descriptionN],
			AdditionalInfo:       printer[additionalInfoN],
			Amount:               uint(amount),
			Weight:               float32(weight),
			PagePerMinute:        float32(ppm),
			Price:                uint(price),
			Size:                 printer[sizeN],
			BrandID:              brand,
			PrintingTechnologyID: technology,
			FunctionTypeID:       functionType,
			PrintSizeID:          printSize,
			PrintResolutionX:     resolutionX,
			PrintResolutionY:     resolutionY,
		})
	}
	return
}
