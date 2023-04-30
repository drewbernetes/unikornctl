package create

import (
	"context"
	"eckctl/pkg/auth"
	"eckctl/pkg/generated"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var createControlPlaneCmd = &cobra.Command{
	Use:   "controlplane",
	Short: "Create a control plane",
	Run: func(cmd *cobra.Command, args []string) {
		url, u, p, project = cmd.Flag("url").Value.String(), cmd.Flag("username").Value.String(),
			cmd.Flag("password").Value.String(), cmd.Flag("project").Value.String()
		token = auth.GetToken(url, u, p, project)
		createControlPlane()
	},
}

func createControlPlane() {
	client := auth.InitClient(url)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cp := generated.ControlPlane{}

	cp.Name = controlPlaneName
	cp.ApplicationBundle.Name = "control-plane-" + controlPlaneVersion
	cp.ApplicationBundle.Version = controlPlaneVersion

	// Create the Unikorn Project if it doesn't already exist, 409s are OK
	resp, err := client.PostApiV1Project(ctx, auth.SetAuthorizationHeader(token))
	if (resp.StatusCode != http.StatusConflict) && (resp.StatusCode != http.StatusAccepted) {
		fmt.Println(resp.StatusCode)
		log.Fatal(err)
	}

	resp, err = client.PostApiV1Controlplanes(ctx, cp, auth.SetAuthorizationHeader(token))
	if resp.StatusCode != http.StatusAccepted {
		fmt.Println(resp.StatusCode)
		log.Fatal(err)
	}

	fmt.Printf("Control Plane %s is being created\n", cp.Name)

}
