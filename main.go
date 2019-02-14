package main

import (
	"encoding/base64"
	"io"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 4 {
		log.Fatalln("参数不完整[input] [output] [output-type](1. file to base64,2. base64 to file)")
	}
	inputPath := os.Args[1]
	outputPath := os.Args[2]
	outputType := os.Args[3]
	var err error
	var inputFile *os.File
	inputFile, err = os.Open(inputPath)
	if err != nil {
		log.Fatalf("读取：%s，文件错误：%s \n", inputPath, err)
	}
	defer inputFile.Close()

	var outputFile *os.File
	outputFile, err = os.Create(outputPath)
	if err != nil {
		log.Fatalf("读取：%s，文件错误：%s \n", outputPath, err)
	}
	defer outputFile.Close()

	switch outputType {
	case "1":
		err := fileToBase64(inputFile, outputFile)
		if err != nil {
			log.Fatalf("转换 %s -> %s 错误：%s \n", inputPath, outputPath, err)
		}
	case "2":
		err := base64ToFile(inputFile, outputFile)
		if err != nil {
			log.Fatalf("转换 %s -> %s 错误：%s \n", inputPath, outputPath, err)
		}
	default:
		log.Fatalf("转换类型[%s]错误，1或2 (1. file to base64,2. base64 to file)", outputType)
	}
	log.Printf("转换 %s -> %s 完成\n", inputPath, outputPath)
}

func fileToBase64(inputFile, outputFile *os.File) error {
	encode := base64.NewEncoder(base64.StdEncoding, outputFile)
	_, err := io.Copy(encode, inputFile)
	return err
}

func base64ToFile(inputFile, outputFile *os.File) error {
	decode := base64.NewDecoder(base64.StdEncoding, inputFile)
	_, err := io.Copy(outputFile, decode)
	return err
}
