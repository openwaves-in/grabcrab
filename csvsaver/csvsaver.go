package csvsaver

import (
	"encoding/csv"
	"os"
	"strings"

	
)
type Job struct {
	Title       string
	Company     string
	Experience  string
	Salary      string
	Location    string
	Description string
	Tags        []string
	Posted      string
	URL         string
}
func SaveToCSV(jobs []Job, filename string) error {
	var file *os.File
	var writer *csv.Writer

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// File doesn't exist, create it
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		writer = csv.NewWriter(file)
		defer writer.Flush()

		header := []string{"Title", "Company", "Experience", "Salary", "Location", "Description", "Tags", "Posted", "URL"}
		if err := writer.Write(header); err != nil {
			return err
		}
	} else {
		// File exists, open in append mode
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		writer = csv.NewWriter(file)
		defer writer.Flush()
	}

	for _, job := range jobs {
		row := []string{job.Title, job.Company, job.Experience, job.Salary, job.Location, job.Description, strings.Join(job.Tags, ", "), job.Posted, job.URL}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
