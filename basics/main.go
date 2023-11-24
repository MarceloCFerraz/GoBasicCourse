package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)


func main(){
	// teste := 1;
	// fmt.Println("Test",teste)
	// fmt.Scan(&teste) // reading something from console

	// m := retornaDict()
	// m[1] = "Not algo" // updates key 1 with string
	// fmt.Println(m)

	// numero, str := retornaDoisValores()
	// fmt.Printf("%d: %s\n", numero, str)

	array := read_file("sites.txt")
	for i, site := range array{
		println(i, "site: ", site) // apparently we can do this without "fmt." at the beggining
	}
}


func retornaDict() map[int]string{
	return map[int]string{
		1: "Algo",
	}
	// esse é o equivalente do dictionary do python
}


func retornaDoisValores() (int, string){
	return 1, "Algo"
	// Go aceita retornar mais de um valor
}


func watch_urls(){
	urls := read_file("sites")

	for _, url_string := range urls {
	// for i := 0; i <= len(urls); i++ {
	// 	url_string := urls[i]
		resp, err := http.Get(url_string)

		if err != nil{
			if resp.StatusCode < 400{
				fmt.Println(resp.StatusCode, resp.Status)
				fmt.Println(resp.Body)				
			}
		}
	}
}


func read_file(file_name string) []string{
	// read_whole_file(file_name)
	// read_file_in_byte_chunks(file_name)
	
	return read_file_in_lines(file_name)
}


func read_whole_file(file_name string) {
	// gets the whole file content (bad for large files)
	content, err := os.ReadFile(file_name) // file content (array of bytes) and error
	fmt.Println("Printing whole file content:")
	
	if err == nil {
		fmt.Println(string(content)) // prints string converted content
	} else {
		fmt.Println("Something went wrong", err)
	}
	fmt.Println()
}


func read_file_in_lines(file_name string) []string{
	file, err := os.Open(file_name) // loads a file pointer
	fmt.Println("Printing file line by line:")
	sites := []string{}
	
	if err == nil{
		defer file.Close() // making sure the file gets closed at the end of the function
		
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines) 
		
		for scanner.Scan(){
			sites = append(sites, scanner.Text())
		}
	} else {
		fmt.Println("Something went wrong", err)
	}
	return sites
}


func read_file_in_byte_chunks(file_name string) {
	file, err := os.Open(file_name) // loads a file pointer
	fmt.Print("Printing file in chunks of ")
	
	if err == nil{
		defer file.Close() // making sure the file gets closed at the end of the function

		const max_size = 4
		fmt.Println(max_size,"bytes:")
		bytes_list := make([]byte, max_size)

		for {
			read_total, err := file.Read(bytes_list) // automatically reads from where the last iteration finished
			if err == nil {
				chunk_str := strings.TrimSpace(string(bytes_list[:read_total]))
				fmt.Println(">", chunk_str)
			} else { // if something went wrong
				if err != io.EOF { // and it wasn't because finished the file
					fmt.Println(err)
				}
				break // either way, finish looping
			}
		}
	} else {
		fmt.Println("Something went wrong", err)
	}
	fmt.Println()
}