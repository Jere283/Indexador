package main

import (
	prof "Indexador/profiling"
	zinc "Indexador/zincsearch"
	"bufio"
	"encoding/json"
	"log"
	"os"
	"path"
	"strings"
	"sync"
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
	case "Message-ID":
		emailStruct.MessageID = value
	case "Date":
		emailStruct.Date = value
	case "From":
		emailStruct.From = value
	case "To":
		emailStruct.To = value
	case "Subject":
		emailStruct.Subject = value
	case "Mime-Version":
		emailStruct.MimeVersion = value
	case "Content-Type":
		emailStruct.ContentType = value
	case "Content-Transfer-Encoding":
		emailStruct.ContentTransferEncoding = value
	case "X-From":
		emailStruct.X_from = value
	case "X-To":
		emailStruct.X_to = value
	case "X-cc":
		emailStruct.X_CC = value
	case "X-bcc":
		emailStruct.X_BCC = value
	case "X-Folder":
		emailStruct.X_folder = value
	case "X-Origin":
		emailStruct.X_origin = value
	case "X-FileName":
		emailStruct.X_fileName = value
	}

	return emailStruct
}

func ConvertEmailFileToJson(filePath string) []byte {
	var bodyLines strings.Builder
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
			bodyLines.WriteString(line)
			bodyLines.WriteString("\n")
		} else {
			// Parse email headers
			parts := strings.SplitN(line, ": ", 2)
			if len(parts) == 2 {
				key := parts[0]
				value := parts[1]
				// Create an Email object
				emailStructure = StructureTheData(key, value, emailStructure)
			}
		}
	}

	// Add the body to the email structure
	emailStructure.Body = bodyLines.String()

	// Convert the struct to JSON using json.MarshalIndent
	jsonDocument, err := json.MarshalIndent(emailStructure, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return jsonDocument
}

func ProcessFiles(filePaths []string, dir string) [][]byte {
	var wg sync.WaitGroup
	var emailJsons [][]byte
	var m sync.Mutex

	for _, filePath := range filePaths {
		wg.Add(1)

		go func(fp string) {
			defer wg.Done()
			fulldir := path.Join(dir, fp)
			emailJson := ConvertEmailFileToJson(fulldir)

			m.Lock()
			emailJsons = append(emailJsons, emailJson)
			m.Unlock()
		}(filePath)
	}

	wg.Wait()

	return emailJsons
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
		Index:    "testV2.6",
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
		var allEmailJsons [][]byte
		mailPath := path + "/" + employees[i]
		mailFolders := listFolder(mailPath) // list the subfolder of each employee

		for j := 0; j < len(mailFolders); j++ {
			filesPath := mailPath + "/" + mailFolders[j] // the path of the email files
			//check if the filespath is a directory to avoid issues
			if isDirectory(filesPath) {
				files := listFolder(filesPath)

				emailJsons := ProcessFiles(files, filesPath)
				allEmailJsons = append(allEmailJsons, emailJsons...)

			}

		}
		for _, json := range allEmailJsons {
			zinc.CreateDocument(json, config)
		}
	}
}
