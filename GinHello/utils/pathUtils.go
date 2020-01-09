package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func RootPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Panicln("发生错误",err.Error())
	}
	log.Println("s:",s)
	i := strings.LastIndex(s, "\\")
	path := s[0 : i+1]
	log.Println("path:",path)
	return path
}