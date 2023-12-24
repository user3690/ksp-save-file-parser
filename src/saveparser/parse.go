package saveparser

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	atmosphereSurvey  = "AtmosphereSurvey"
	crewObservation   = "CrewObservation"
	environmentSurvey = "EnvironmentSurvey"
	surfaceSurvey     = "SurfaceSurvey"
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

func Parse(filePath string, saveToFile bool) error {
	var (
		save       kerbalSaveFile
		data       []byte
		file       *os.File
		workingDir string
		err        error
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

	save.Agencies[0].ScienceReports, err = sortExperiments(save.Agencies[0].ScienceReports)
	if err != nil {
		return err
	}

	if saveToFile {
		workingDir, err = os.Getwd()
		if err != nil {
			return err
		}

		pathToReportsFile := workingDir + string(os.PathSeparator) + "reports.txt"

		file, err = os.Create(pathToReportsFile)
		if err != nil {
			return err
		}

		for _, curAgency := range save.Agencies {
			output := createOutput(curAgency)

			_, err = fmt.Fprint(file, output)
			if err != nil {
				return err
			}
		}

		file.Close()

		workingDir, err = os.Getwd()
		if err != nil {
			return err
		}

		fmt.Printf("saved to %s\n", pathToReportsFile)

		return err
	}

	for _, curAgency := range save.Agencies {
		output := createOutput(curAgency)
		fmt.Print(output)
	}

	return err
}

func sortExperiments(unsorted []scienceReport) (sorted []scienceReport, err error) {
	var (
		atmosphereSurveyReports  []scienceReport
		crewObservationReports   []scienceReport
		environmentSurveyReports []scienceReport
		surfaceSurveyReports     []scienceReport
		sortFunc                 = func(a, b scienceReport) int {
			return strings.Compare(a.ResearchLocationId, b.ResearchLocationId)
		}
	)

	// sort by experiment id
	for _, report := range unsorted {
		switch report.ExperimentId {
		case atmosphereSurvey:
			atmosphereSurveyReports = append(atmosphereSurveyReports, report)
		case crewObservation:
			crewObservationReports = append(crewObservationReports, report)
		case environmentSurvey:
			environmentSurveyReports = append(environmentSurveyReports, report)
		case surfaceSurvey:
			surfaceSurveyReports = append(surfaceSurveyReports, report)
		default:
			return nil, errors.New(fmt.Sprintf("unkown experiment found: %s", report.ExperimentId))
		}
	}

	slices.SortFunc(atmosphereSurveyReports, sortFunc)
	slices.SortFunc(crewObservationReports, sortFunc)
	slices.SortFunc(environmentSurveyReports, sortFunc)
	slices.SortFunc(surfaceSurveyReports, sortFunc)

	return concatMultipleSlices([][]scienceReport{
		atmosphereSurveyReports,
		crewObservationReports,
		environmentSurveyReports,
		surfaceSurveyReports,
	}), nil
}

func createOutput(curAgency agency) (output string) {
	output += fmt.Sprintf("Science data for %s\n", curAgency.Name)
	output += fmt.Sprintf("SciencePointCapacity: %d\n", curAgency.SciencePointCapacity)
	output += fmt.Sprintf("AdditionalSciencePoints: %d\n", curAgency.AdditionalSciencePoints)
	output += fmt.Sprintln("")
	output += fmt.Sprintln("===== Science Reports =====")

	for _, report := range curAgency.ScienceReports {
		output += fmt.Sprintf(
			"ExperimentId: %s | ResearchLocationId: %s | ResearchReportType: %s | FinalScienceValue: %f\n",
			report.ExperimentId,
			report.ResearchLocationId,
			report.ResearchReportType,
			report.FinalScienceValue,
		)
	}

	return output
}

func concatMultipleSlices[T any](slices [][]T) []T {
	var (
		totalLength int
		i           int
	)

	for _, s := range slices {
		totalLength += len(s)
	}

	result := make([]T, totalLength)

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}
