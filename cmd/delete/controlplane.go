package delete

import (
	"context"
	"eckctl/pkg/auth"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

func deleteControlPlaneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "controlplane",
		Short: "Delete an ECK control plane",
		Run: func(cmd *cobra.Command, args []string) {
			url, u, p, project = cmd.Flag("url").Value.String(), cmd.Flag("username").Value.String(),
				cmd.Flag("password").Value.String(), cmd.Flag("project").Value.String()
			token = auth.GetToken(url, u, p, project)
			err := deleteControlPlane()
			if err != nil {
				log.Fatalf("Error deleting Control Plane: %s", err)
			}
		},
	}
	cmd.Flags().StringVar(&controlPlaneName, "name", "", "The name of the control plane to be deleted")
	err := cmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatalln(err)
	}
	return cmd
}

func deleteControlPlane() (err error) {

	client := auth.InitClient(url)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.DeleteApiV1ControlplanesControlPlaneName(ctx, controlPlaneName, auth.SetAuthorizationHeader(token))
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusAccepted {
		err = fmt.Errorf("Unexpected response code: %v", resp.StatusCode)
		return
	}

	return
}
