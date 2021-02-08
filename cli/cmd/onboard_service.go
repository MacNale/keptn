package cmd

import (
	"bytes"
	"errors"
	"os"
	"text/template"

	keptncommon "github.com/keptn/go-utils/pkg/lib/keptn"

	"github.com/keptn/keptn/cli/pkg/validator"
	keptnutils "github.com/keptn/kubernetes-utils/pkg"
	"github.com/spf13/cobra"
)

type onboardServiceCmdParams struct {
	Project       *string
	ChartFilePath *string
}

var onboardServiceParams *onboardServiceCmdParams

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service SERVICENAME --project=PROJECTNAME --chart=FILEPATH",
	Short: "Onboards a new service and its Helm chart to a project",
	Long: `Onboards a new service and its Helm chart to the provided project. 
Therefore, this command takes a folder to a Helm chart or an already packed Helm chart as .tgz.
`,
	Deprecated: "please use \"create service\" and 'add-resource' instead.",
	Example: `keptn onboard service SERVICENAME --project=PROJECTNAME --chart=FILEPATH

keptn onboard service SERVICENAME --project=PROJECTNAME --chart=HELM_CHART.tgz
`,
	SilenceUsage: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			cmd.SilenceUsage = false
			return errors.New("required argument SERVICENAME not set")
		}
		if !keptncommon.ValidateKeptnEntityName(args[0]) {
			errorMsg := "Service name contains upper case letter(s) or special character(s).\n"
			return errors.New(errorMsg)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := doOnboardServicePreRunChecks(args); err != nil {
			return err
		}

		data := struct {
			CliPath string
			Project string
			Service string
			HelmChartPath string
		}{
			os.Args[0],
			*onboardServiceParams.Project,
			args[0],
			*onboardServiceParams.ChartFilePath,
		}

		msg := `This command is deprecated and no longer supported. Instead, please use: 

{{.CliPath}} create service {{.Service}} --project={{.Project}}
{{.CliPath}} add-resource --project={{.Project}} --all-stages --service={{.Service}} --resource={{.HelmChartPath}} --resourceUri=helm/{{.Service}}

Please execute these commands sequentially.`

		tmpl, err := template.New("msg").Parse(msg)

		if err != nil {
			return err
		}

		// store template in a Bytes Buffer
		var tpl bytes.Buffer
		err = tmpl.Execute(&tpl, data)

		if err != nil {
			return err
		}

		// return above message as an error - this will lead the CLI to respond with an exit code != 0
		return errors.New(tpl.String())
	},
}

func doOnboardServicePreRunChecks(args []string) error {
	// validate chart flag
	*onboardServiceParams.ChartFilePath = keptnutils.ExpandTilde(*onboardServiceParams.ChartFilePath)

	if _, err := os.Stat(*onboardServiceParams.ChartFilePath); os.IsNotExist(err) {
		return errors.New("Provided Helm chart does not exist")
	}

	ch, err := keptnutils.LoadChartFromPath(*onboardServiceParams.ChartFilePath)
	if err != nil {
		return err
	}

	res, err := validator.ValidateHelmChart(ch, args[0])
	if err != nil {
		return err
	}

	if !res {
		return errors.New("The provided Helm chart is invalid. Please checkout the requirements")
	}

	return nil
}

func init() {
	onboardCmd.AddCommand(serviceCmd)
	onboardServiceParams = &onboardServiceCmdParams{}
	onboardServiceParams.Project = serviceCmd.Flags().StringP("project", "p", "", "The name of the project")
	serviceCmd.MarkFlagRequired("project")

	onboardServiceParams.ChartFilePath = serviceCmd.Flags().StringP("chart", "", "", "A path to a Helm chart folder or an already archived Helm chart")
	serviceCmd.MarkFlagRequired("chart")
}
