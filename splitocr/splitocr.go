package splitocr

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

const (
	dockerImageName = "jerrywang1981/text-detection-tf:splitocr0.0.1"
	imageName       = "docker.io/" + dockerImageName
)

func SplitOcr(folderName, resultFolderName string) error {
	fullFolderName, err := filepath.Abs(folderName)
	if err != nil {
		return err
	}
	fullResultFolderName := filepath.Join(fullFolderName, resultFolderName)
	_, err = os.Stat(fullResultFolderName)
	if !os.IsNotExist(err) {
		return err
	}

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	reader, err := cli.ImagePull(ctx, dockerImageName, types.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   []string{resultFolderName},
		Tty:   true,
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: fullFolderName,
				Target: "/data",
			},
		},
	}, nil, "")

	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	_, err = cli.ContainerWait(ctx, resp.ID)
	if err != nil {
		return err
	}
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return err
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	return nil
}
