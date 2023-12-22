package saveparser

import (
	"encoding/json"
	"fmt"
	"os"
)

type kerbalSaveFile struct {
	Agencies []agency `json:"Agencies"`
}

type agency struct {
	Id                      uint8           `json:"AgencyId"`
	Name                    string          `json:"AgencyName"`
	ScienceReports          []scienceReport `json:"SubmittedResearchReports"`
	SciencePointCapacity    int32           `json:"SciencePointCapacity"`
	AdditionalSciencePoints int32           `json:"AdditionalSciencePoints"`
}

type scienceReport struct {
	ExperimentId       string  `json:"ExperimentID"`
	ResearchLocationId string  `json:"ResearchLocationID"`
	ResearchReportType string  `json:"ResearchReportType"`
	FinalScienceValue  float64 `json:"FinalScienceValue"`
}

func Parse(filePath string) error {
	var (
		save kerbalSaveFile
		data []byte
		err  error
	)

	if _, err = os.Stat(filePath); err != nil {
		return err
	}

	data, err = os.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &save); err != nil {
		return err
	}

	for _, curAgency := range save.Agencies {
		fmt.Printf("Science data for %s\n", curAgency.Name)
		fmt.Printf("SciencePointCapacity: %d\n", curAgency.SciencePointCapacity)
		fmt.Printf("AdditionalSciencePoints: %d\n", curAgency.AdditionalSciencePoints)
		fmt.Println("")
		fmt.Println("===== Science Reports =====")
		for _, report := range curAgency.ScienceReports {
			fmt.Printf("ExperimentId: %s\n", report.ExperimentId)
			fmt.Printf("ResearchLocationId: %s\n", report.ResearchLocationId)
			fmt.Printf("ResearchReportType: %s\n", report.ResearchReportType)
			fmt.Printf("FinalScienceValue: %f\n", report.FinalScienceValue)
			fmt.Println("=====")
		}
	}

	return err
}
