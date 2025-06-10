package main

import (
	"log"

	filesvsdatabase "github.com/nt2311-vn/build_myown_database/01_Files_VS_Database"
)

func main() {
	err := filesvsdatabase.SaveData2("file.txt", []byte("Write to tmp and rename file"))
	if err != nil {
		log.Fatalln("cannot save data ", err.Error())
	}
}
