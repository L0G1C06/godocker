package schemas

import "fmt"

type Variables struct{
	DockerRegistryUserID string 
	ProjectDirectory string 
	NameTag string 
	DockerMain string 
	DockerFileMainLanguage string 
}

func IsGo() string{
	imageBlueprint := `
	FROM golang:latest
	
	WORKDIR /go/src/app
	
	COPY . .
	
	ENV PATH="${PATH}:(go env GOPATH)/bin"
	ENV PATH="${PATH}:(go env GOROOT)/bin"
	
	RUN go mod download
	
	EXPOSE 8080
	
	RUN go build -o main .
	CMD ["./main"]
	`

	return imageBlueprint
}

func IsPython(filename string) string{
	imageBlueprint :=fmt.Sprintf( `
	FROM python:3.9-slim

	WORKDIR /app 

	COPY . /app 

	RUN pip install --no-cache-dir -r requirements.txt

	EXPOSE 8080

	CMD ["python", "%s"]
	`, filename)

	return imageBlueprint
}