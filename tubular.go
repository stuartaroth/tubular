package tubular

import (
	"archive/zip"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Page is the struct used to build either a separated values file or an xlsx sheet
//
// Name is the filename or sheet name depending on what you build
//
// Rows are the tabular data that will be written in the file or sheet
type Page struct {
	Name string
	Rows [][]interface{}
}

// WriteXLSXFile allows you to write an XLSX file
//
// filenameWithouEnding is the desired filename.
//
// pages is a slice of Page. Each Page will create a sheet in the XLSX file
//
// WriteXLSXFile will create an XLSX file and return the created filename and an error
func WriteXLSXFile(filenameWithoutEnding string, pages []Page) (string, error) {
	xlsxFilename := filenameWithoutEnding + ".xlsx"

	xlsxFile := xlsx.NewFile()
	for _, page := range pages {
		xlsxSheet, err := xlsxFile.AddSheet(page.Name)
		if err != nil {
			return xlsxFilename, err
		}

		for _, row := range page.Rows {
			xlsxRow := xlsxSheet.AddRow()
			for _, column := range row {
				xlsxCell := xlsxRow.AddCell()
				switch column.(type) {
				case bool:
					xlsxCell.SetBool(column.(bool))
				default:
					xlsxCell.SetValue(column)
				}
			}
		}
	}

	err := xlsxFile.Save(xlsxFilename)
	if err != nil {
		return xlsxFilename, err
	}

	return xlsxFilename, nil
}

// WriteSeparatedValuesFile allows you to write multiple separated values files
//
// filenameWithouEnding is the desired filename of the created zipfile.
//
// pages is a slice of Page. Each Page will create a separate file
//
// WriteSeparatedValuesFile will create an zip file that will include a file per passed page and return the created filename and an error
func WriteSeparatedValuesFile(filenameWithoutEnding string, fileEnding string, newLineString string, separatorString string, pages []Page) (string, error) {
	zippedFilename := filenameWithoutEnding + ".zip"

	filenames := []string{}
	for _, page := range pages {
		filename, err := writeSVFile(fileEnding, newLineString, separatorString, page)
		if err != nil {
			return zippedFilename, err
		}

		filenames = append(filenames, filename)
	}

	err := createZippedFile(zippedFilename, filenames)
	if err != nil {
		return zippedFilename, err
	}

	err = deleteSVFiles(filenames)
	if err != nil {
		return zippedFilename, err
	}

	return zippedFilename, nil
}

func writeSVFile(fileEnding string, newLineString string, separatorString string, page Page) (string, error) {
	filename := page.Name + "." + fileEnding

	stringRows := getStringRows(page.Rows)

	separatorRows := make([]string, len(stringRows))
	for i := range stringRows {
		separatorRows[i] = strings.Join(stringRows[i], separatorString)
	}

	fileContents := strings.Join(separatorRows, newLineString) + newLineString

	err := ioutil.WriteFile(filename, []byte(fileContents), 0666)
	if err != nil {
		return filename, err
	}

	return filename, nil
}

func getStringRows(rows [][]interface{}) [][]string {
	stringRows := make([][]string, len(rows))
	for i := range rows {
		stringRows[i] = make([]string, len(rows[i]))
		for j := range rows[i] {
			stringRows[i][j] = fmt.Sprintf("%v", rows[i][j])
		}
	}

	return stringRows
}

func createZippedFile(filename string, files []string) error {
	zippedFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer zippedFile.Close()

	zippedFileWriter := zip.NewWriter(zippedFile)
	defer zippedFileWriter.Close()

	for _, file := range files {
		fileToZip, err := os.Open(file)
		if err != nil {
			return err
		}

		fileToZipInfo, err := fileToZip.Stat()
		if err != nil {
			return err
		}

		fileToZipHeader, err := zip.FileInfoHeader(fileToZipInfo)
		if err != nil {
			return err
		}

		fileToZipHeader.Method = zip.Deflate

		writer, err := zippedFileWriter.CreateHeader(fileToZipHeader)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, fileToZip)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteSVFiles(filenames []string) error {
	for _, filename := range filenames {
		err := os.Remove(filename)
		if err != nil {
			return err
		}
	}

	return nil
}
