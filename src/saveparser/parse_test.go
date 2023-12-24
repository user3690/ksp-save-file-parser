package saveparser

import "testing"

func TestSortExperiments(t *testing.T) {
	var unsorted = []scienceReport{
		{
			ExperimentId:       "Test",
			ResearchLocationId: "Kerbin_Atmosphere_KerbinMountain",
			ResearchReportType: "DataType",
			FinalScienceValue:  5,
		},
		{
			ExperimentId:       "AtmosphereSurvey",
			ResearchLocationId: "Kerbin_Atmosphere_KerbinMountain",
			ResearchReportType: "DataType",
			FinalScienceValue:  1,
		},
		{
			ExperimentId:       "AtmosphereSurvey",
			ResearchLocationId: "Kerbin_Atmosphere_KerbinHighlands",
			ResearchReportType: "DataType",
			FinalScienceValue:  1,
		},
		{
			ExperimentId:       "Test",
			ResearchLocationId: "Kerbin_Atmosphere_KerbinHighlands",
			ResearchReportType: "DataType",
			FinalScienceValue:  5,
		},
		{
			ExperimentId:       "CrewObservation",
			ResearchLocationId: "Kerbin_Atmosphere_KerbinMountain",
			ResearchReportType: "DataType",
			FinalScienceValue:  2,
		},
		{
			ExperimentId:       "EnvironmentSurvey",
			ResearchLocationId: "Kerbin_Atmosphere_KerbinMountain",
			ResearchReportType: "DataType",
			FinalScienceValue:  3,
		},
		{
			ExperimentId:       "Unknown",
			ResearchLocationId: "Kerbin_Atmosphere_KerbinMountain",
			ResearchReportType: "DataType",
			FinalScienceValue:  6,
		},
		{
			ExperimentId:       "SurfaceSurvey",
			ResearchLocationId: "Kerbin_Atmosphere_KerbinMountain",
			ResearchReportType: "DataType",
			FinalScienceValue:  4,
		},
	}

	sorted := sortExperiments(unsorted)

	if len(sorted) != 8 {
		t.Fatalf("expected length %d but is %d", 8, len(sorted))
	}

	if sorted[1].ResearchLocationId != "Kerbin_Atmosphere_KerbinMountain" {
		t.Fatalf("wrong research location id at position 2: %s", sorted[1].ResearchLocationId)
	}

	if sorted[5].ExperimentId != "Test" {
		t.Fatalf("wrong experiment id at position 6: %s", sorted[5].ExperimentId)
	}

	if sorted[5].ResearchLocationId != "Kerbin_Atmosphere_KerbinHighlands" {
		t.Fatalf("wrong research location id at position 6: %s", sorted[5].ResearchLocationId)
	}

	if sorted[6].ExperimentId != "Test" {
		t.Fatalf("wrong experiment id at position 7: %s", sorted[6].ExperimentId)
	}

	if sorted[6].ResearchLocationId != "Kerbin_Atmosphere_KerbinMountain" {
		t.Fatalf("wrong research location id at position 7: %s", sorted[6].ResearchLocationId)
	}

	if sorted[7].ExperimentId != "Unknown" {
		t.Fatalf("wrong experiment id at position 7: %s", sorted[7].ExperimentId)
	}
}
