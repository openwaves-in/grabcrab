package extractorpro

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/openwaves-in/grabcrab/csvsaver"
	"golang.org/x/net/html"
)



func Extractor() {
	// Open the local HTML file
	file, err := os.Open("page.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse the HTML content
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	// Extract job data
	var jobs []csvsaver.Job
	extractJobs(doc, &jobs)

	// Save extracted job data to CSV
	err = csvsaver.SaveToCSV(jobs, "jobs.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data saved to jobs.csv successfully.")
}

func extractJobs(n *html.Node, jobs *[]csvsaver.Job) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, attr := range n.Attr {
			if attr.Key == "class" && strings.Contains(attr.Val, "cust-job-tuple") {
				var job csvsaver.Job

				// Extract data fields
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "div" && strings.Contains(getAttributeValue(c, "class"), "row1") {
						for a := c.FirstChild; a != nil; a = a.NextSibling {
							if a.Type == html.ElementNode && a.Data == "a" && strings.Contains(getAttributeValue(a, "class"), "title") {
								job.Title = strings.TrimSpace(a.FirstChild.Data)
								job.URL = getAttributeValue(a, "href")
							}
						}
					} else if c.Type == html.ElementNode && c.Data == "div" && strings.Contains(getAttributeValue(c, "class"), "row2") {
						for span := c.FirstChild; span != nil; span = span.NextSibling {
							if span.Type == html.ElementNode && span.Data == "span" && strings.Contains(getAttributeValue(span, "class"), "comp-dtls-wrap") {
								for a := span.FirstChild; a != nil; a = a.NextSibling {
									if a.Type == html.ElementNode && a.Data == "a" && strings.Contains(getAttributeValue(a, "class"), "comp-name") {
										job.Company = strings.TrimSpace(a.FirstChild.Data)
									}
								}
							}
						}
					} else if c.Type == html.ElementNode && c.Data == "div" && strings.Contains(getAttributeValue(c, "class"), "row3") {
						for child := c.FirstChild; child != nil; child = child.NextSibling {
							if child.Type == html.ElementNode && strings.Contains(getAttributeValue(child, "class"), "job-details") {
								for details := child.FirstChild; details != nil; details = details.NextSibling {
									if details.Type == html.ElementNode && strings.Contains(getAttributeValue(details, "class"), "exp-wrap") {
										for span := details.FirstChild; span != nil; span = span.NextSibling {
											if span.Type == html.ElementNode && strings.Contains(getAttributeValue(span, "class"), "ni-job-tuple-icon ni-job-tuple-icon-srp-experience exp") {
												for span2 := span.FirstChild; span2 != nil; span2 = span2.NextSibling {
													if span2.Type == html.ElementNode && strings.Contains(getAttributeValue(span2, "class"), "expwdth") {
														job.Experience = strings.TrimSpace(span2.FirstChild.Data)
													}
												}
											}
										}
									}
									if details.Type == html.ElementNode && strings.Contains(getAttributeValue(details, "class"), "sal-wrap ver-line") {
										for span := details.FirstChild; span != nil; span = span.NextSibling {
											if span.Type == html.ElementNode && strings.Contains(getAttributeValue(span, "class"), "ni-job-tuple-icon ni-job-tuple-icon-srp-rupee sal") {
												for span2 := span.FirstChild; span2 != nil; span2 = span2.NextSibling {
													if span2.Type == html.ElementNode && span2.Data == "span" {
														// Check if the 'title' attribute is present
														if titleAttr := getAttributeValue(span2, "title"); titleAttr != "" {
															job.Salary = strings.TrimSpace(titleAttr)
															break // exit the loop after extracting salary
														}
													}
												}
											}
										}
									}
									if details.Type == html.ElementNode && strings.Contains(getAttributeValue(details, "class"), "loc-wrap ver-line") {
										for span := details.FirstChild; span != nil; span = span.NextSibling {
											if span.Type == html.ElementNode && strings.Contains(getAttributeValue(span, "class"), "ni-job-tuple-icon ni-job-tuple-icon-srp-location loc") {
												for span2 := span.FirstChild; span2 != nil; span2 = span2.NextSibling {
													if span2.Type == html.ElementNode && strings.Contains(getAttributeValue(span2, "class"), "locWdth") {
														job.Location = strings.TrimSpace(span2.FirstChild.Data)
													}
												}
											}
										}
									}
								}
							}
						}

					}else if c.Type == html.ElementNode && c.Data == "div" && strings.Contains(getAttributeValue(c, "class"), "row5") {
						for ul := c.FirstChild; ul != nil; ul = ul.NextSibling {
							if ul.Type == html.ElementNode && ul.Data == "ul" && strings.Contains(getAttributeValue(ul, "class"), "tags-gt") {
								for li := ul.FirstChild; li != nil; li = li.NextSibling {
									if li.Type == html.ElementNode && li.Data == "li" && strings.Contains(getAttributeValue(li, "class"), "tag-li") {
										job.Tags = append(job.Tags, strings.TrimSpace(li.FirstChild.Data))
									}
								}
							}
						}
					} else if c.Type == html.ElementNode && c.Data == "div" && strings.Contains(getAttributeValue(c, "class"), "row6") {
						for span := c.FirstChild; span != nil; span = span.NextSibling {
							if span.Type == html.ElementNode && span.Data == "span" && strings.Contains(getAttributeValue(span, "class"), "job-post-day") {
								job.Posted = strings.TrimSpace(span.FirstChild.Data)
							}
						}
					}
				}

				// Append job to the list
				*jobs = append(*jobs, job)
			}
		}
	}

	// Recursively traverse the HTML tree
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractJobs(c, jobs)
	}
}

func getAttributeValue(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}
