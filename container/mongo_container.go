package container

import (
	"os/exec"

	"github.com/pkg/errors"
)

const docker = "docker"
const portMap = "27017:27017"
const imageName = "mongo"

// MongoContainer is a testing container of https://hub.docker.com/_/mongo/
type MongoContainer struct {
	ContainerName string
}

// Start creates and runs a new mongo container.
func (m MongoContainer) Start() error {
	if err := createContainer(m.ContainerName); err != nil {
		return err
	}

	return startContainer(m.ContainerName)
}

// Stop halts and delete the created container. Must be called to release resources.
func (m MongoContainer) Stop() error {
	err := stopContainer(m.ContainerName)

	if err != nil {
		return err
	}

	return removeContainer(m.ContainerName)
}

func createContainer(containerName string) error {
	args := []string{
		"create",
		"--name",
		containerName,
		"-p",
		portMap,
		imageName,
	}

	output, err := exec.Command(docker, args...).CombinedOutput()

	if err != nil {
		err = errors.WithMessage(err, string(output))
	}

	return err
}

func startContainer(containerName string) error {
	args := []string{
		"start",
		containerName,
	}

	output, err := exec.Command(docker, args...).CombinedOutput()

	if err != nil {
		err = errors.WithMessage(err, string(output))
	}

	return err
}

func stopContainer(containerName string) error {
	args := []string{
		"container",
		"stop",
		containerName,
	}
	return exec.Command(docker, args...).Run()
}

func removeContainer(containerName string) error {
	args := []string{
		"container",
		"rm",
		containerName,
	}

	return exec.Command(docker, args...).Run()
}
