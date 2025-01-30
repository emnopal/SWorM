# SWorM
Simple Workflow Manager is a simple cli based automation tools

> [!NOTE]
> this project is still at its very early stage of prototype
> expect everything to break! or even straight up not working

## run
to run the project,

1. configure the `openapi.yml` file and the `workflow.json` file.

2. then you can run the workflow with the following command
```sh
go run .\main.go

# or

go build -o sworm .\main.go 
./sworm
```
