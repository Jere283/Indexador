# ZincSearch Indexer and Data Visualizer

## Project Overview

The project is a combination of an Indexer and an API that hosts a Vue.js dist app for searching through indexed emails. The Indexer is responsible for crawling through a specified directory structure, extracting information from email files, and indexing them using the ZincSearch API. The API, built with Go and Chi, allows users to search for emails based on specific keywords.

## Indexer (Go Aplication)

The Indexer is a Go application that performs the following tasks:

1. Folder Listing: Recursively lists all files and subfolders in the specified directory.
2. Email Parsing: Reads email files, extracts relevant information (headers and body), and structures the data into a JSON format.
3. Bulk Indexing: Utilizes the ZincSearch API to bulk index the parsed email data.

### Usage

Usage

- The main function configures the ZincSearch connection, performs CPU and memory profiling, and processes the email files in the specified directory.
- The ConvertEmailFileToJson function parses individual email files.
- The ProcessFiles function handles the parallel processing of files and subfolders.

### Dependencies

- ZincSearch: A search and analytics engine for Elasticsearch.

## API (Go application)

The API is a Go application built with the Chi router. It provides basic CORS support and exposes endpoints for retrieving information from the ZincSearch index.

### Endpoints

- /api/v1: Welcome message.
- /api/v1/search/{word}: Search endpoint to retrieve emails based on a keyword.

### Serving a Vue.js App

The API also serves a Vue.js dist folder to allow users to interact with the indexed data visually.

### Dependencies

Chi: A lightweight, idiomatic web framework.

## ZincSearch Functions (Go package)

The ZincSearch package contains functions for interacting with the ZincSearch API. It includes functions for creating documents, bulk indexing, and searching.

### Functions

- CreateDocument: Creates a single document in the ZincSearch index.
- BulkCreateDocument: Performs bulk indexing of multiple documents.
- SearchDocument: Searches for documents in the ZincSearch index based on a specified word.

### Configuration

The Config struct holds the configuration details, including the ZincSearch base URL, index name, username, and password.

## Getting Started

1. Ensure Go is installed on your system.
2. Set up a ZincSearch instance and configure the Config struct with the appropriate details.
3. Run the Indexer and API applications.

#### Additional Notes

- CPU and memory profiling have been implemented in the Indexer for performance analysis.
- The project assumes a specific directory structure for email files.

Feel free to customize the code based on your specific requirements and directory structure. If you encounter any issues or have suggestions for improvement, please open an issue in the project repository.
