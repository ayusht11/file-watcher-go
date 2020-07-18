# Go1.13 parameters

TOOLS = golang.org/x/tools/cmd/goimports 

tools: ; $(info $(M) building tools…)
	go1.13 get -v $(TOOLS)

format: 
	goimports -w $$(find . -type f -name '*.go' -not -path "./vendor/*")

watcher:
	go1.13 run main.go

server:
	go1.13 run server.go
