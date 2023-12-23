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