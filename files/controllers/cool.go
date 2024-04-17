package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Uaa(ctx *gin.Context) {

	fmt.Println("sent")

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.Error(err)
		return
	}

	//Open received file
	csvFileToImport, err := fileHeader.Open()
	if err != nil {
		ctx.Error(err)
		return
	}
	defer csvFileToImport.Close()

	//Create temp file
	tempFile, err := os.CreateTemp("", fileHeader.Filename)
	if err != nil {
		ctx.Error(err)
		return
	}
	defer tempFile.Close()

	// fmt.Println(tempFile.Name())
	//Delete temp file after importing
	defer os.Remove(tempFile.Name())

	// //Write data from received file to temp file
	fileBytes, err := io.ReadAll(csvFileToImport)
	if err != nil {
		ctx.Error(err)
		return
	}
	os.Stdout.Write(fileBytes)

	err = os.WriteFile("text.jpg", fileBytes, 0644)

	bytesString := string(fileBytes)

	// bytesString = bytesString[1 : len(bytesString)-1]

	// imgBase64Str := make([]byte, base64.StdEncoding.EncodeToString(bytesString))
	fmt.Println(bytesString)
	// _, err = tempFile.Write(fileBytes)
	if err != nil {
		ctx.Error(err)
		return
	}

	// ctx.JSON(http.StatusOK, string(fileBytes))
	ctx.String(http.StatusOK, fmt.Sprintf("files uploaded!"))

}
