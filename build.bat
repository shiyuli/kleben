rem build assets
cd src\templates
go-assets-builder -o assets.go -p templates index.tmpl
cd ..\..\

rem build go
go build -o bin\kleben.exe src\main.go
