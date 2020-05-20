package splitocr

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/jerrywang1981/handy-tools/util"
)

const (
	dockerImageName = "jerrywang1981/text-detection-tf:splitocr0.0.3"
)

func SplitOcr(folderName, resultFolderName string) error {
	fullFolderName, err := filepath.Abs(folderName)
	if err != nil {
		return err
	}
	fullResultFolderName, err := filepath.Abs(resultFolderName)
	if err != nil {
		return err
	}
	_, err = os.Stat(fullResultFolderName)
	if os.IsNotExist(err) {
		return err
	}

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return err
	}

	imageList := []string{}
	for _, image := range images {
		imageList = append(imageList, image.RepoTags...)
	}

	if !util.Contains(imageList, dockerImageName) {
		return fmt.Errorf("Please pull image: %s", dockerImageName)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: dockerImageName,
		Tty:   true,
	}, &container.HostConfig{
		AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: fullFolderName,
				Target: "/data",
			},
			{
				Type:   mount.TypeBind,
				Source: fullResultFolderName,
				Target: "/result",
			},
		},
	}, nil, "")

	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println(err)
		return err
	}

	_, err = cli.ContainerWait(ctx, resp.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return err
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	return nil
}
