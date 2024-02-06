VK-Finance

<< Build exe with icon >>

go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo

Copy testdata/resource/versioninfo.json into your working directory and then modify the file with your own settings.

make go.mod file

1. go mod init github.com/VkHyperNova/VK-FINANCE
2. Run Go mod tidy

copy manifest file to resource folder

ADD this at the start of go file
//go:generate goversioninfo -icon=testdata/resource/icon.ico -manifest=testdata/resource/goversioninfo.exe.manifest

1. Generate resource.syso file
2. go build (not go build vkf.go. this will not work. just type go build!)

