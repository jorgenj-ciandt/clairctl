package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/coreos/clair/api/v1"
	"github.com/jorgenj-ciandt/clairctl/clair"
	"github.com/spf13/cobra"
)

var jsonPath string
var htmlPath string

//ImageAnalysis Full image analysis
type ImageAnalysis struct {
	Registry, ImageName, Tag string
	Layers                   []v1.LayerEnvelope
}

var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "Generate HTML report from JSON",
	Long:  `Generate Docker Image vulnerabilities report as HTML`,
	Run: func(cmd *cobra.Command, args []string) {
		var analysis clair.ImageAnalysis
		err := json.Unmarshal([]byte(getAnalysisFromFile(jsonPath)), &analysis)
		if err != nil {
			fmt.Printf("clairctl: \"html\" could not get analysis")
			os.Exit(1)
		}

		html, err := clair.ReportAsHTML(analysis)
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile(htmlPath, []byte(html), 0700)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func getAnalysisFromFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorf("File error: %v\n", err)
	}

	return file
}

func init() {
	RootCmd.AddCommand(htmlCmd)
	htmlCmd.Flags().StringVarP(&jsonPath, "vulnerability", "v", "/tmp/clair_vulnerabilities.json", "Clair vulnerabilities JSON path")
	htmlCmd.Flags().StringVarP(&htmlPath, "output", "o", "/tmp/clair_report.html", "HTML report path output")
}
