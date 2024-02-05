package docker_commands

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/L0G1C06/godocker/basic"
	"github.com/L0G1C06/godocker/schemas"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ImagePush(dockerClient *client.Client, dockerRegistryUserID string, nameTag string, dockerPassword string) error{
	var authConfig = schemas.AuthConfig{
		Username: dockerRegistryUserID,
		Password: dockerPassword,
		ServerAddress: "https://index.docker.io/v1/",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()
	
	authConfigBytes, _ := json.Marshal(authConfig)
	authConfigEncoded := base64.URLEncoding.EncodeToString(authConfigBytes)

	tag := dockerRegistryUserID + "/" + nameTag
	opts := types.ImagePushOptions{RegistryAuth: authConfigEncoded}
	rd, err := dockerClient.ImagePush(ctx, tag, opts)
	if err != nil{
		return err
	}

	defer rd.Close()

	err = basic.Print(rd)
	if err != nil{
		return err
	}

	return nil 
}