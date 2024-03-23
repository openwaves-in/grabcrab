package csvsaver

import (
	"encoding/csv"
	"io"
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

func ReadFromCSV(filename string) ([]Job, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip header
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var jobs []Job
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		job := Job{
			Title:       row[0],
			Company:     row[1],
			Experience:  row[2],
			Salary:      row[3],
			Location:    row[4],
			Description: row[5],
			Tags:        strings.Split(row[6], ", "),
			Posted:      row[7],
			URL:         row[8],
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func SaveToCSV(jobs []Job, filename string) error {
	var file *os.File
	var writer *csv.Writer

	// Read existing data if file exists
	var existingJobs []Job
	if _, err := os.Stat(filename); err == nil {
		existingJobs, err = ReadFromCSV(filename)
		if err != nil {
			return err
		}
	}

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
		if !jobExists(existingJobs, job) {
			row := []string{job.Title, job.Company, job.Experience, job.Salary, job.Location, job.Description, strings.Join(job.Tags, ", "), job.Posted, job.URL}
			if err := writer.Write(row); err != nil {
				return err
			}
		}
	}

	return nil
}

func jobExists(jobs []Job, job Job) bool {
	for _, j := range jobs {
		if j.Title == job.Title && j.Company == job.Company && j.Experience == job.Experience &&
			j.Salary == job.Salary && j.Location == job.Location && j.Description == job.Description &&
			j.Posted == job.Posted && j.URL == job.URL {
			return true
		}
	}
	return false
}
