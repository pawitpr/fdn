package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Upload a file")
		fmt.Println("2. Download a file")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice (1/2/3): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter the Upload url: ")
	uploadURL  , _ := reader.ReadString('\n')
	uploadURL   = cleanInput(uploadURL  ) 

	fmt.Print("Enter the File Path: ")
	fileToUpload, _ := reader.ReadString('\n')
	fileToUpload = cleanInput(fileToUpload ) // Clean the input by removing newlines and extra whitespaces

	// Now you have the user-provided values in the variables uploadFileUrl and fileToUpload 
	fmt.Printf("Uploaded File URL: %s\n", uploadURL)
	fmt.Printf("Uploaded File Name: %s\n", fileToUpload )
	err := uploadFile(fileToUpload, uploadURL)
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}

	
		case "2":
			downloadFile()
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}

func cleanInput(input string) string {
	return strings.TrimSpace(strings.Replace(input, "\n", "", -1))
}
func uploadFile(filename string, targetURL string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	writer.Close()

	request, err := http.NewRequest("POST", targetURL, body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Server returned non-200 status: %v", response.Status)
	}

	fmt.Println("File uploaded successfully!")
	return nil
}


func downloadFile(fileURL, destFileName string) error {
	response, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Server returned non-200 status: %v", response.Status)
	}

	file, err := os.Create(destFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Println("File downloaded successfully!")
	return nil
}