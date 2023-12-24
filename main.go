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

// We create a struct that contains the structore of the JSON we will send to zincsearch
type Email struct {
	MessageID               string `json:"message_id"`
	Date                    string `json:"date"`
	From                    string `json:"from"`
	To                      string `json:"to"`
	Subject                 string `json:"subject"`
	MimeVersion             string `json:"mime_version"`
	ContentType             string `json:"content_type"`
	ContentTransferEncoding string `json:"content_transfer_encoding"`
	ToName                  string `json:"to_name"`
	CC                      string `json:"cc"`
	BCC                     string `json:"bcc"`
	Folder                  string `json:"folder"`
	Origin                  string `json:"origin"`
	FileName                string `json:"file_name"`
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

	for _, file := range files {
		files_list = append(files_list, file.Name()) //we access the filenames using file.Name()
	}

	// it returns array with all the names
	return files_list
}

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
		emailStruct.From = value
	case "x-to":
		emailStruct.ToName = value
	case "x-cc":
		emailStruct.CC = value
	case "x-bcc":
		emailStruct.BCC = value
	case "x-folder":
		emailStruct.Folder = value
	case "x-origin":
		emailStruct.Origin = value
	case "x-filename":
		emailStruct.FileName = value
	}

	return emailStruct
}

func ConvertEmailFileToJson(filePath string) []byte {
	var body string
	var emailStructure Email

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			key := strings.ToLower(parts[0])
			value := parts[1]

			emailStructure = StructureTheData(key, value, emailStructure)
		}

	}
	for scanner.Scan() {
		line := scanner.Text()
		body += line
	}
	body = strings.Replace(body, "  ", " ", -1)
	emailStructure.Body = body

	// Convert the struct to a json using the json.MarshalIndent

	jsonDocument, err := json.MarshalIndent(emailStructure, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return jsonDocument
}

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func main() {
	cpuProfile := prof.StartCPUProfile()
	defer prof.StopCpuProfile(cpuProfile)

	memoryProfile := prof.StartMemoryProfile()
	defer prof.StopMemoryProfile(memoryProfile)

	var path string = "C:/Users/jerem/OneDrive/Escritorio/proyecto/enron_mail_20110402/maildir2"
	employees := listFolder(path)

	for i := 0; i < len(employees); i++ {
		mailPath := path + "/" + employees[i]
		mailFolders := listFolder(mailPath)

		for i := 0; i < len(mailFolders); i++ {
			filesPath := mailPath + "/" + mailFolders[i]

			if isDirectory(filesPath) {
				files := listFolder(filesPath)

				for i := 0; i < len(files); i++ {
					filePath := filesPath + "/" + files[i]
					bodyQuery := ConvertEmailFileToJson(filePath)
					zinc.CreateDocument(bodyQuery, "finalIndex0.1")
				}
			}

		}

	}

}
