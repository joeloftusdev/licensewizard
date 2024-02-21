package main

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

type TemplateData struct {
	Year    int
	Author  string
	Project string
}

//go:embed templates/*
var templateFiles embed.FS

var LicenseTemplates = map[string]string{
	"MIT":         "mit.tmpl",
	"BSD_2.0":     "bsd2.0.tmpl",
	"BSD_3.0":     "bsd3.0.tmpl",
	"APACHE_2.0":  "apache2.0.tmpl",
	"AGPL_3.0":    "agpl3.0.tmpl",
	"GPL_3.0":     "gpl3.0.tmpl",
	"GPL_2.0":     "gpl2.0.tmpl",
	"MPL_2.0":     "mpl2.0.tmpl",
	"CCZERO_1.0":  "ccszero1.0.tmpl",
	"ECLIPSE_2.0": "eclipse2.0.tmpl",
	"LGPL_2.1":    "lgpl2.1.tmpl",
	"UNLICENSE":   "unlicense.tmpl",
	"BOOST":       "boost.tmpl",
}

func getUserInput(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func generateLicense(templateText, author, project string) (string, error) {
	data := TemplateData{
		Year:    time.Now().Year(),
		Author:  author,
		Project: project,
	}

	tmpl, err := template.New("license").Parse(templateText)
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Unsupported LICENSE type please use supported LICENSES:\nMIT\nBSD_2.0\nBSD_3.0\nAPACHE_2.0\nAGPL_3.0\nGPL_3.0\nGPL_2.0\nMPL_2.0\nCCZERO_1.0\nECLIPSE_2.0\nLGPL_2.1\nUNLICENSE\nBOOST")
		os.Exit(1)
	}

	licenseType := strings.ToUpper(os.Args[1])

	templateFile, ok := LicenseTemplates[licenseType]
	if !ok {
		fmt.Println("Unsupported LICENSE type please use supported LICENSES:\nMIT\nBSD_2.0\nBSD_3.0\nAPACHE_2.0\nAGPL_3.0\nGPL_3.0\nGPL_2.0\nMPL_2.0\nCCZERO_1.0\nECLIPSE_2.0\nLGPL_2.1\nUNLICENSE\nBOOST")
		os.Exit(1)
	}

	author := getUserInput("Author: ")
	project := getUserInput("Project name: ")
	outputDir := getUserInput("Output Directory: ")

	templateText, err := templateFiles.ReadFile("templates/" + templateFile)
	if err != nil {
		fmt.Printf("Error reading LICENSE file: %v\n", err)
		os.Exit(1)
	}

	licenseText, err := generateLicense(string(templateText), author, project)
	if err != nil {
		fmt.Printf("Error generating LICENSE: %v\n", err)
		os.Exit(1)
	}

	outputPath := fmt.Sprintf("%s\\LICENSE", outputDir)
	err = os.WriteFile(outputPath, []byte(licenseText), 0644)
	if err != nil {
		fmt.Printf("Error writing to LICENSE file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("LICENSE generated and saved to %s.\n", outputPath)
}
