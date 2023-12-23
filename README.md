# KSP2 Save File Parser
Some little project to extract the finished science reports in KSP2 over command line.
Save files can be found under: 
```
example path: C:\Users\[user]\AppData\LocalLow\Intercept Games\Kerbal Space Program 2\Saves\SinglePlayer\[Campaign Name]

usage: ksp_save_file_parser.exe save [pathToSaveFile]
to save the output to a file: ksp_save_file_parser.exe save -s [pathToKSP2SaveFile]
```

To build the binary for yourself, you need go version 1.21.* installed, execute following command in the root folder:
```
GOOS=windows GOARCH=amd64 go build -o ksp_save_file_parser.exe main.go
```

The end result looks something like this:
```
Science data for Your Agency
SciencePointCapacity: 615
AdditionalSciencePoints: 0

===== Science Reports =====
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_HighOrbit | ResearchReportType: DataType | FinalScienceValue: 48.000000
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_Landed_KerbinMountain | ResearchReportType: DataType | FinalScienceValue: 4.000000
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_Landed_KerbinGrasslands | ResearchReportType: DataType | FinalScienceValue: 4.000000
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_Landed_KerbinKSC | ResearchReportType: DataType | FinalScienceValue: 4.000000
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_Splashed_KerbinWater | ResearchReportType: DataType | FinalScienceValue: 4.000000
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_Atmosphere | ResearchReportType: DataType | FinalScienceValue: 4.000000
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_LowOrbit | ResearchReportType: DataType | FinalScienceValue: 24.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Atmosphere | ResearchReportType: SampleType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Landed_KerbinKSC | ResearchReportType: SampleType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_LowOrbit | ResearchReportType: DataType | FinalScienceValue: 30.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_HighOrbit | ResearchReportType: DataType | FinalScienceValue: 60.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_HighOrbit | ResearchReportType: SampleType | FinalScienceValue: 60.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Atmosphere | ResearchReportType: DataType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Landed_KerbinKSC | ResearchReportType: DataType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_LowOrbit | ResearchReportType: SampleType | FinalScienceValue: 30.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Landed_KerbinMountain | ResearchReportType: SampleType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Landed_KerbinGrasslands | ResearchReportType: DataType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Landed_KerbinGrasslands | ResearchReportType: SampleType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Landed_KerbinMountain | ResearchReportType: DataType | FinalScienceValue: 5.000000
ExperimentId: SurfaceSurvey | ResearchLocationId: Kerbin_Landed_KerbinBeach | ResearchReportType: DataType | FinalScienceValue: 6.000000
ExperimentId: SurfaceSurvey | ResearchLocationId: Kerbin_Landed_KerbinBeach | ResearchReportType: SampleType | FinalScienceValue: 18.000000
ExperimentId: SurfaceSurvey | ResearchLocationId: Kerbin_Landed_KerbinMountain | ResearchReportType: DataType | FinalScienceValue: 6.000000
ExperimentId: SurfaceSurvey | ResearchLocationId: Kerbin_Landed_KerbinMountain | ResearchReportType: SampleType | FinalScienceValue: 18.000000
```