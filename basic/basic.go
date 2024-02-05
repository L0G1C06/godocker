package basic

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/L0G1C06/godocker/schemas"
)

func Print(rd io.Reader) error {
	var lastLine string

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		lastLine = scanner.Text()
		fmt.Println(scanner.Text())
	}

	errLine := &schemas.ErrorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		return errors.New(errLine.Error)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func WriteDockerfileInfo(outputPath string, dockerfile_type string){
	file, err := os.Create(outputPath + "/" + "Dockerfile")
	if err != nil{
		panic(err)
	}
	
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(dockerfile_type)
	if err != nil{
		panic(err)
	}

	writer.Flush()
	
}