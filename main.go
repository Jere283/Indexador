package main

import (
	prof "Indexador/profiling"
	zinc "Indexador/zincsearch"
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

// We create a struct that contains the structure of the JSON we will send to zincsearch
type Email struct {
	MessageID               string `json:"Message_id"`
	Date                    string `json:"Date"`
	From                    string `json:"From"`
	To                      string `json:"To"`
	Subject                 string `json:"Subject"`
	MimeVersion             string `json:"Mime_version"`
	ContentType             string `json:"Content_type"`
	ContentTransferEncoding string `json:"Content_transfer_encoding"`
	X_from                  string `json:"X-from"`
	X_to                    string `json:"X-to"`
	X_CC                    string `json:"X-cc"`
	X_BCC                   string `json:"X-bcc"`
	X_folder                string `json:"X-folder"`
	X_origin                string `json:"X-origin"`
	X_fileName              string `json:"X-file_name"`
	Body                    string `json:"body"`
}

// This function will list all the folders and files insider the path we send
// This will be used to list all the email files and all the subfolders inside the dataset
func listFolder(folder_path string) []string {
	var files_list []string

	//using the os.ReadDir we get the information of the path we send
	files, err := os.ReadDir(folder_path)
	if err != nil {
		log.Fatal(err)
	}

	//We save all the files that were found in an array
	for _, file := range files {
		files_list = append(files_list, file.Name()) //we access the filenames using file.Name()
	}

	// it returns array with all the names
	return files_list
}

// This function will return a Email object which is the structure of the Json we will send
func StructureTheData(key string, value string, emailStruct Email) Email {
	switch key {
	case "message-id":
		emailStruct.MessageID = value
	case "date":
		emailStruct.Date = value
	case "from":
		emailStruct.From = value
	case "to":
		emailStruct.To = value
	case "subject":
		emailStruct.Subject = value
	case "mime-version":
		emailStruct.MimeVersion = value
	case "content-type":
		emailStruct.ContentType = value
	case "content-transfer-encoding":
		emailStruct.ContentTransferEncoding = value
	case "x-from":
		emailStruct.X_from = value
	case "x-to":
		emailStruct.X_to = value
	case "x-cc":
		emailStruct.X_CC = value
	case "x-bcc":
		emailStruct.X_BCC = value
	case "x-folder":
		emailStruct.X_folder = value
	case "x-origin":
		emailStruct.X_origin = value
	case "x-filename":
		emailStruct.X_fileName = value
	}

	return emailStruct
}

func ConvertEmailFileToJson(filePath string) []byte {
	var bodyLines []string
	var emailStructure Email
	var bodyStarted bool

	// We read the email file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// We read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// If the line is empty, it indicates the start of the email body
		if line == "" {
			bodyStarted = true
			continue
		}

		if bodyStarted {
			// Store body lines in a slice
			bodyLines = append(bodyLines, line)
		} else {
			// Parse email headers
			parts := strings.SplitN(line, ": ", 2)
			if len(parts) == 2 {
				key := strings.ToLower(parts[0])
				value := parts[1]
				// Create an Email object
				emailStructure = StructureTheData(key, value, emailStructure)
			}
		}
	}

	// Concatenate body lines with line breaks
	emailStructure.Body = strings.Join(bodyLines, "\n")

	// Convert the struct to JSON using json.MarshalIndent
	jsonDocument, err := json.MarshalIndent(emailStructure, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return jsonDocument
}

// Check if the file that was found is a directory
func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func main() {
	config := zinc.Config{
		BaseURL:  "http://localhost:4080",
		Index:    "EnronEmailDataset1.0",
		Username: "admin",
		Password: "Complexpass#123",
	}

	//CPU  and memory Profiling
	cpuProfile := prof.StartCPUProfile()
	defer prof.StopCpuProfile(cpuProfile)

	memoryProfile := prof.StartMemoryProfile()
	defer prof.StopMemoryProfile(memoryProfile)
	//maildir path
	var path string = "C:/Users/jerem/OneDrive/Escritorio/proyecto/enron_mail_20110402/maildir"
	employees := listFolder(path) // list the folders which have the people's names

	//TODO: Improve this repetitive code
	for i := 0; i < len(employees); i++ {
		mailPath := path + "/" + employees[i]
		mailFolders := listFolder(mailPath) // list the subfolder of each employee

		for i := 0; i < len(mailFolders); i++ {
			filesPath := mailPath + "/" + mailFolders[i] // the path of the email files
			//check if the filespath is a directory to avoid issues
			if isDirectory(filesPath) {
				files := listFolder(filesPath)

				for i := 0; i < len(files); i++ {
					filePath := filesPath + "/" + files[i]
					bodyQuery := ConvertEmailFileToJson(filePath)
					zinc.CreateDocument(bodyQuery, config)
				}

			} else {
				bodyQuery := ConvertEmailFileToJson(filesPath) //if the file is not a directory it will read the email file
				zinc.CreateDocument(bodyQuery, config)
			}
		}
	}
}
