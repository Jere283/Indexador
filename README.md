# ZincSearch Indexer and Data Visualizer

## Project Overview

The project is a combination of an Indexer and an API that allows you to search over the indexed files and hosts a Vue.js application that works as a User Interface to search over the indexed files. The Indexer is responsible for crawling through a specified directory structure, extracting information from the email files, and indexing them using the ZincSearch API. My API, built with Go and Chi, allows users to search for emails based on specific keywords.

### Technologies used in this project
* [Go 1.21.5 for the backend](https://go.dev)
* [Go-Chi/V5 for the API](https://github.com/go-chi/chi)
* [Vue 3 for the frontend](https://vuejs.org)
* [ZincSearch as the database](https://github.com/zincsearch/zincsearch)
* [ZincSearch API to interact with the database](https://zincsearch-docs.zinc.dev)

### Project Structure

- ZincSearch-Indexer-WebSearchTool
  - api
    - main.go     ## Api main source code, search endpoint and static serve of the dist folder (the build of the vue app)
    - dist ## vue app minimized folder
    - go.mod
    - go.sum
    - api.exe     ## Api executable file, hosted in port 3000
  - frontend ## vue 3 source code
  - profiling 
    - proftests     ## folder with the profiling tests of indexer V1 and V2 
    - go.mod
    - profiling.go     # go package with function to control the profiling profiles
  - zincsearch
    - go.mod
    - zinc.go     ## go package with functions to interact with the zincsearch API (createDocument, BulkCreateDocuments, Search)
  - go.mod
  - go.work
  - go.work.sum
  - improvementsV2.md     ## Document with information about the improvements in v1 and v2
  - Indexer.exe     ## Indexer executable file 
  - main.go     ## Indexer main source code
  - README.md
### Configuration

The Config struct holds the configuration details, including the ZincSearch base URL, index name, username, and password.
 
### Installation 
1. Clone the git repository
```bash 
  Git clone github.com/Jere283/ZincSearch-Indexer-WebSearchTool
```
2. Download [ZincSearch](https://github.com/zincsearch/zincsearch) and follow the [ZincSearch Quick Start](https://zincsearch-docs.zinc.dev/quickstart/)
3. Download the Enron-Email-Dataset ( you will need the path of this folder later)
```bash
  curl -L http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz -o enron_mail_20110402.tgz && tar -xf enron_mail_20110402.tgz
```
4. Copy the path of the maildir folder inside the enron_mail_20110402 and add it as the value of the path variable in line 186 inside the https://github.com/Jere283/ZincSearch-Indexer-WebSearchTool/blob/main/main.go file.
```go
var path string = "your_path" //here
```

With these steps  you will be able to index the files inside of the maildir folder.

**IMPORTANT NOTE**: to serve the vue js app you need to copy the folderpath and paste it in the fs variable in line 49 of your https://github.com/Jere283/ZincSearch-Indexer-WebSearchTool/blob/main/api/main.go file. **This will be fixed in future versions**
```go
  fs := http.FileServer(http.Dir("dist folder directory")) //Here
```


## Indexer (Go Application)

The Indexer is a Go application that performs the following tasks:

1. Folder Listing: Recursively lists all files and subfolders in the specified directory.
2. Email Parsing: Reads email files, extracts relevant information (headers and body), and structures the data into a JSON format.
3. Bulk Indexing: Utilizes the ZincSearch API to bulk index the parsed email data.

### Usage

- The main function configures the ZincSearch connection (-You need to set the Index name you want to use in the config structure in line 172.), performs CPU and memory profiling, and processes the email files in the specified directory.
- The ConvertEmailFileToJson function parses individual email files.
- The ProcessFiles function handles the parallel processing of files and subfolders.

### Dependencies

- ZincSearch: A search and analytics engine for Elasticsearch.

## API (Go application)

The API is a Go application built with the Chi router. It provides basic CORS support and exposes endpoints for retrieving information from the ZincSearch index.

### Endpoints

- /api/v1: Welcome message.
- /api/v1/search/{word}: Search endpoint to retrieve emails jsons based on a keyword.

### Serving a Vue.js App

The API also serves a Vue.js dist folder to allow users to interact with the indexed data visually.

![image](https://github.com/Jere283/ZincSearch-Indexer-WebSearchTool/assets/111548280/e6147dc3-a62f-40b3-bf28-d336900c076f)


### Dependencies

Chi: A lightweight, idiomatic web framework.

## ZincSearch Functions (Go package)

The Zinc package contains functions for interacting with the ZincSearch API. It includes functions for creating documents, bulk indexing, and searching.

### Functions

- CreateDocument: Creates a single document in the ZincSearch index.
- BulkCreateDocument: Performs bulk indexing of multiple documents.
- SearchDocument: Searches for documents in the ZincSearch index based on a specified word.

#### Additional Notes

- CPU and memory profiling have been implemented in the Indexer for performance analysis.
- The project assumes a specific directory structure for email files.

Feel free to customize the code based on your specific requirements and directory structure. If you encounter any issues or have suggestions for improvement, please open an issue in the project repository.
