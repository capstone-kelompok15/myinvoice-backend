package fileutils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"text/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type EvaluateValueFunc func(value interface{}) interface{}

func EvaluateValue(value interface{}) interface{} {
	typeV := reflect.ValueOf(value)

	if typeV.Kind() == reflect.Ptr {
		if typeV.IsNil() {
			return " "
		} else {
			return typeV.Elem()
		}
	}

	return value
}

func ReadableDate(value interface{}) interface{} {
	evaluatedValue := EvaluateValue(value)
	if evaluatedValue == " " {
		return " "
	}

	stringDate := fmt.Sprintf("%v", evaluatedValue)

	splittedDate := strings.Split(stringDate, "T")

	return splittedDate[0]
}

func SubTotalValue(quantity interface{}, price interface{}) interface{} {
	evaluatedQtyValue := EvaluateValue(quantity)
	if evaluatedQtyValue == " " {
		return " "
	}

	evaluatedPriceValue := EvaluateValue(price)
	if evaluatedPriceValue == " " {
		return " "
	}

	intQty := evaluatedQtyValue.(int)
	intPrice := evaluatedPriceValue.(int64)
	subTotal := int64(intQty) * intPrice
	subTotalValue := ReadableIdr(strconv.FormatInt(int64(subTotal), 10))

	return subTotalValue
}

func InvoiceString(quantity interface{}) interface{} {
	evaluatedValue := EvaluateValue(quantity)
	return fmt.Sprintf("%06v", evaluatedValue)
}

func ReadableIdr(value interface{}) interface{} {
	str := fmt.Sprintf("%v", value)
	strLen := len(str)

	idr := ""

	for strLen >= 3 {
		idr = str[strLen-3:strLen] + idr
		strLen -= 3
		if strLen != 0 {
			idr = "." + idr
		}
	}

	if strLen < 3 {
		idr = str[:strLen] + idr
	}

	return idr
}

func CreatePDFFromHTMLFile(templatePath, resFileName string, data interface{}) error {
	templateFileName := resFileName + ".html"

	htmlFile, err := os.Create(templateFileName)
	if err != nil {
		log.Println(err)
		return err
	}
	defer htmlFile.Close()

	templ := template.New(path.Base(templatePath)).Funcs(template.FuncMap{
		"EvaluateValue": EvaluateValue,
		"ReadableDate":  ReadableDate,
		"InvoiceString": InvoiceString,
		"SubTotalValue": SubTotalValue,
		"ReadableIdr":   ReadableIdr,
	})

	templ, err = templ.ParseFiles(templatePath)
	if err != nil {
		log.Println(err)
		return err
	}

	err = templ.Execute(htmlFile, data)
	if err != nil {
		log.Println(err)
		return err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println(err)
		return err
	}

	byteHtml, err := ioutil.ReadFile(templateFileName)
	if err != nil {
		log.Println(err)
		return err
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewReader(byteHtml))
	page.PageOptions = wkhtmltopdf.NewPageOptions()
	page.PageOptions.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	err = pdfg.Create()
	if err != nil {
		log.Println(err)
		return err
	}

	err = pdfg.WriteFile(resFileName)
	if err != nil {
		log.Println(err)
		return err
	}

	os.Remove(templateFileName)
	return nil
}
