package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	urls, result := "C://Users/Стажер/Desktop/task/3/adres.txt", "C://Users/Стажер/Desktop/task/3/result/"
	file, err := os.Open(urls)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	/*	dir, er := os.Open(result) //хз правильно ли
		if err != nil {
			fmt.Println(er)
			os.Exit(1)
		}
		defer dir.Close()*/

	var dataUrl []string
	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		//dataUrl = append(dataUrl, string(data[:n]))
		dataUrl = strings.Fields(string(data[:n]))
	}

	for i := 0; i < len(dataUrl); i++ {
		httpRequest := "GET / HTTP/1.1\n" + "Host: " + dataUrl[i] + "\n\n"
		//httpRequest := "GET / HTTP/1.1\n" + dataUrl[i]

		conn, err := net.Dial("tcp", dataUrl[i]+":http")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		if _, err = conn.Write([]byte(httpRequest)); err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println(conn)
		tmp, errr := os.Create(result + dataUrl[i] + ".txt")
		if errr != nil {
			fmt.Println("Unable to create file:", errr)
			os.Exit(1)
		}
		defer file.Close()
		//запись результата в новый созданный файл в папке
		//tmp.WriteString("")
		io.Copy(tmp, conn) //это в файл, maybe
	}
}

// util.exe -url path/urls.txt -result=path/result/
// go run task -url path/urls.txt -result=path/result/
