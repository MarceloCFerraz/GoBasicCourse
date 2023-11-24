package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)


const sites_file_name string = "sites.txt"
const logs_file_name string = "logs.txt"
var logs string


func main(){
	clear_file(logs_file_name)
	// test := 1;
	// fmt.Println("Test",test)
	// fmt.Scan(&test) // reading something from console

	// m := return_dict()
	// m[1] = "Not Something" // updates key 1 with string
	// fmt.Println(m)

	// number, str := return_two_values()
	// fmt.Printf("%d: %s\n", number, str)

	sites := read_file(sites_file_name)

	for i:= 0; i < 5; i++ {
		for _, site := range sites{
		// for i := 0; i <= len(sites); i++ {
			// site := sites[i]
			watch_url(site)
		}
	}

	check_logs(logs_file_name)
}


func clear_file(logs_file_name string) {
	file, err := os.Create(logs_file_name)
	if err != nil{
		print_error(err)
	}
	file.Close()
}


func check_logs(logs_file_name string) {
	content := read_whole_file(logs_file_name)

	if len(content) > 0{
		if content == logs{
			println("Logs were all saved correctly")
		} else {
			println("Logs are not equal. Check them below:")
			println("\nLOGS:")
			println(logs)
			println("==================")
			println("\nLOGS FILE:")
			println(content)
			println("==================")
		}
	} else {
		println("The file '",logs_file_name,"' is empty")
	}
}


func watch_url(site_url string){
	resp, err := http.Get(site_url)

	if err == nil{ //all good
		log := fmt.Sprintf("%s -> %d", site_url, resp.StatusCode)
		// please, refer to this to format dates: https://gosamples.dev/date-time-format-cheatsheet/
		// go's date formatting is a complete bullshit
		log = time.Now().UTC().Format("[2006-01-02T15:04:05 -07:00]") + fmt.Sprintf(" %s",log)		
		println(log)

		logs += log + "\n"
		write_logs(log)
	} else {
		print_error(err)
	}
}


func write_logs(log string) {
	if !file_exists(logs_file_name) {
		create_file(logs_file_name)
	} 
	file, err := os.OpenFile(logs_file_name, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	
	if err != nil {
		print_error(err)
	} else {
		file.WriteString(log + "\n")
	}
}

func file_exists(logs_file_name string) bool{
	_, err := os.Stat(logs_file_name); 
	return os.IsExist(err)
}

func create_file(logs_file_name string) {
	file, err := os.OpenFile(logs_file_name, os.O_CREATE, 0644)

	if err != nil {
		print_error(err)
	} 
	file.Close()
}

func print_error(err error) {
	print("Something went wrong:\n> ", err)
}


func read_file(file_name string) []string{
	// read_whole_file(file_name)
	// read_file_in_byte_chunks(file_name)
	
	return read_file_in_lines(file_name)
}


func read_whole_file(file_name string) string {
	// gets the whole file content (bad for large files)
	content, err := os.ReadFile(file_name) // file content (array of bytes) and error
	// fmt.Println("Printing whole file content:")
	
	if err == nil {
		return string(content) // prints string converted content
	} else {
		fmt.Println("Something went wrong", err)
	}
	fmt.Println()
	return ""
}


func read_file_in_lines(file_name string) []string{
	file, err := os.Open(file_name) // loads a file pointer
	// fmt.Println("Printing file line by line:")
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
	// fmt.Print("Printing file in chunks of ")
	
	if err == nil{
		defer file.Close() // making sure the file gets closed at the end of the function

		const max_size = 4
		// fmt.Println(max_size,"bytes:")
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

// just some testing functions


func return_dict() map[int]string{
	return map[int]string{
		1: "Something",
	}
	// python dictionary equivalent
}


func return_two_values() (int, string){
	return 1, "Something"
	// Go accetpts returning more than 1 thing
}