package server

import (
	"fmt"
	"os"
	"time"
)

func Logg(info string){
	file, err := os.Open("logs/logs.txt")
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().GoString() + " \n\n" + info)
}
func Logerr(err error){
	file, errs := os.Create("logs/" + "ERR" + time.Now().String() + ".txt")
	if errs != nil {
		fmt.Println(errs)
	}

	file.WriteString(err.Error())
}