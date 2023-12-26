# KSP2 Save File Parser
Some little project to extract the completed science reports in KSP2 over command line.
It shows you what, where and how much science you have done.
A for Windows precompiled binary can be found in the `bin` folder.

## Save files
Save files can be found under: `C:\Users\[user]\AppData\LocalLow\Intercept Games\Kerbal Space Program 2\Saves\SinglePlayer\[Campaign Name]`

## Usage of Command on CLI
Flags:
```
-s, --saveToFile   save output to reports.txt
```
Command Examples:
```
ksp_save_file_parser.exe save [pathToSaveFile]
ksp_save_file_parser.exe save -s [pathToSaveFile]
```

To build the binary for yourself, you need go version 1.21.* installed, execute following command in the root folder:
```
GOOS=windows GOARCH=amd64 go build -o ksp_save_file_parser.exe main.go
```

The end result looks something like this:
```
Science data for : Your Agency
SciencePointCapacity: 2213
AdditionalSciencePoints: 0

===== Science Reports =====
ExperimentId: AtmosphereSurvey | ResearchLocationId: Kerbin_Atmosphere_KerbinBeach | ResearchReportType: DataType | FinalScienceValue: 12.000000
ExperimentId: AtmosphereSurvey | ResearchLocationId: Kerbin_Atmosphere_KerbinGrasslands | ResearchReportType: DataType | FinalScienceValue: 12.000000
ExperimentId: AtmosphereSurvey | ResearchLocationId: Kerbin_Atmosphere_KerbinHighlands | ResearchReportType: DataType | FinalScienceValue: 12.000000
ExperimentId: AtmosphereSurvey | ResearchLocationId: Kerbin_Atmosphere_KerbinMountain | ResearchReportType: DataType | FinalScienceValue: 12.000000
ExperimentId: AtmosphereSurvey | ResearchLocationId: Kerbin_Atmosphere_KerbinWater | ResearchReportType: DataType | FinalScienceValue: 12.000000
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_Atmosphere | ResearchReportType: DataType | FinalScienceValue: 4.000000
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_HighOrbit | ResearchReportType: DataType | FinalScienceValue: 48.000000
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_Landed_KerbinGrasslands | ResearchReportType: DataType | FinalScienceValue: 4.000000
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_Landed_KerbinKSC | ResearchReportType: DataType | FinalScienceValue: 4.000000
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_Landed_KerbinMountain | ResearchReportType: DataType | FinalScienceValue: 4.000000
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_LowOrbit | ResearchReportType: DataType | FinalScienceValue: 24.000000
ExperimentId: CrewObservation | ResearchLocationId: Kerbin_Splashed_KerbinWater | ResearchReportType: DataType | FinalScienceValue: 4.000000
ExperimentId: CrewObservation | ResearchLocationId: Mun_HighOrbit | ResearchReportType: DataType | FinalScienceValue: 40.000000
ExperimentId: CrewObservation | ResearchLocationId: Mun_Landed_MunMare | ResearchReportType: DataType | FinalScienceValue: 24.000000
ExperimentId: CrewObservation | ResearchLocationId: Mun_LowOrbit | ResearchReportType: DataType | FinalScienceValue: 56.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Atmosphere | ResearchReportType: SampleType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Atmosphere | ResearchReportType: DataType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_HighOrbit | ResearchReportType: DataType | FinalScienceValue: 60.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_HighOrbit | ResearchReportType: SampleType | FinalScienceValue: 60.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Landed_KerbinGrasslands | ResearchReportType: DataType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Landed_KerbinGrasslands | ResearchReportType: SampleType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Landed_KerbinKSC | ResearchReportType: DataType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Landed_KerbinKSC | ResearchReportType: SampleType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Landed_KerbinMountain | ResearchReportType: DataType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_Landed_KerbinMountain | ResearchReportType: SampleType | FinalScienceValue: 5.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_LowOrbit | ResearchReportType: SampleType | FinalScienceValue: 30.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Kerbin_LowOrbit | ResearchReportType: DataType | FinalScienceValue: 30.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Minmus_HighOrbit | ResearchReportType: DataType | FinalScienceValue: 70.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Minmus_LowOrbit | ResearchReportType: DataType | FinalScienceValue: 80.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Mun_HighOrbit | ResearchReportType: DataType | FinalScienceValue: 50.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Mun_HighOrbit | ResearchReportType: SampleType | FinalScienceValue: 50.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Mun_Landed_MunMare | ResearchReportType: DataType | FinalScienceValue: 30.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Mun_Landed_MunMare | ResearchReportType: SampleType | FinalScienceValue: 30.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Mun_LowOrbit | ResearchReportType: DataType | FinalScienceValue: 70.000000
ExperimentId: EnvironmentSurvey | ResearchLocationId: Mun_LowOrbit | ResearchReportType: SampleType | FinalScienceValue: 70.000000
ExperimentId: SurfaceSurvey | ResearchLocationId: Kerbin_Landed_KerbinBeach | ResearchReportType: DataType | FinalScienceValue: 6.000000
ExperimentId: SurfaceSurvey | ResearchLocationId: Kerbin_Landed_KerbinBeach | ResearchReportType: SampleType | FinalScienceValue: 18.000000
ExperimentId: SurfaceSurvey | ResearchLocationId: Kerbin_Landed_KerbinKSC | ResearchReportType: DataType | FinalScienceValue: 6.000000
ExperimentId: SurfaceSurvey | ResearchLocationId: Kerbin_Landed_KerbinKSC | ResearchReportType: SampleType | FinalScienceValue: 18.000000
ExperimentId: SurfaceSurvey | ResearchLocationId: Kerbin_Landed_KerbinMountain | ResearchReportType: DataType | FinalScienceValue: 6.000000
ExperimentId: SurfaceSurvey | ResearchLocationId: Kerbin_Landed_KerbinMountain | ResearchReportType: SampleType | FinalScienceValue: 18.000000
ExperimentId: SurfaceSurvey | ResearchLocationId: Mun_Landed_MunMare | ResearchReportType: DataType | FinalScienceValue: 36.000000
ExperimentId: SurfaceSurvey | ResearchLocationId: Mun_Landed_MunMare | ResearchReportType: SampleType | FinalScienceValue: 108.000000
```
