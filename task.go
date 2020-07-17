package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var url *string
var result *string

func init() {
	url = flag.String("url", "C://Users/Стажер/Desktop/task/3/adres.txt", "a string")
	result = flag.String("result", "C://Users/Стажер/Desktop/task/3/result/", "a string")
}
func main() {
	flag.Parse()
	file, err := os.Open(*url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var dataUrl []string
	data := make([]byte, 128)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		dataUrl = strings.Fields(string(data[:n]))
	}

	for i := 0; i < len(dataUrl); i++ {
		resp, err := http.Get("" + dataUrl[i] + "")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		tmp, errr := os.Create(*result + strconv.Itoa(i) + ".html")
		if errr != nil {
			fmt.Println("Unable to create file:", errr)
			os.Exit(1)
		}
		defer tmp.Close()
		io.Copy(tmp, resp.Body)
	}
}
