package docker_commands

import (
	"log"
	"os"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/L0G1C06/godocker/basic"
	"github.com/L0G1C06/godocker/schemas"
)

func Commands(){
	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Fatal("usage error: Command is required.\n Type go run main.go help to see the docs")
	}
	if (os.Args[1] != "buildImage" && os.Args[1] != "push"){
		log.Fatalf("usage error: Command %s does not exist. Type go run main.go help to see the docs", os.Args[1])
	}
	if os.Args[1] == "help"{
		fmt.Println(`
Usage:
	
	go run main.go <command>
	
The commands are:

	buildImage	builds a Docker image based on a Dockerfile
	deploy 		upload the image to DockerHub
		`)
		os.Exit(0)
	}
	if os.Args[1] == "buildImage"{
		inputInformation := schemas.Variables{}

		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// Dockerfile create configuration
		fmt.Print("Type Dockerfile main language (Python || Go): ")
		fmt.Scan(&inputInformation.DockerMain)
		if inputInformation.DockerMain == "Python"{
			inputInformation.DockerFileMainLanguage = schemas.IsPython("main.py")
		}
		if inputInformation.DockerMain == "Go"{
			inputInformation.DockerFileMainLanguage = schemas.IsGo()
		}
		// Docker image build configuration
		fmt.Print("Type Docker Hub userID: ")
		fmt.Scan(&inputInformation.DockerRegistryUserID)
		fmt.Print("Type projectDirectory path: ")
		fmt.Scan(&inputInformation.ProjectDirectory)
		fmt.Print("Type nameTag for the image name: ")
		fmt.Scan(&inputInformation.NameTag)
		fmt.Printf("\n Docker image being created with the name: %s/%s\n", inputInformation.DockerRegistryUserID, inputInformation.NameTag)
		basic.WriteDockerfileInfo(inputInformation.ProjectDirectory, inputInformation.DockerFileMainLanguage)
		err = ImageBuild(cli, inputInformation.ProjectDirectory, inputInformation.DockerRegistryUserID, inputInformation.NameTag)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	if os.Args[1] == "push"{
		inputInformation := schemas.Push{}
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// Docker push auth configuration
		fmt.Print("Type your DockerHub username ID: ")
		fmt.Scan(&inputInformation.Username)
		fmt.Print("Type your DockerHub password: ")
		fmt.Println("\033[8m")
		fmt.Scan(&inputInformation.Password)
		fmt.Println("\033[28m")
		fmt.Print("Type the tag of the image to be pushed: ")
		fmt.Scan(&inputInformation.NameTag)
		// Pushing the docker image to DockerHub
		ImagePush(cli, inputInformation.Username, inputInformation.NameTag, inputInformation.Password)
	}
}