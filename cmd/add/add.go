package add

import (
	"encoding/json"
	"fmt"
	"github.com/esteam85/interviews-tracker/process/domain"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	postulationType string
	platform        string
	company         string
	position        string
	jobType         string
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add an interview",
	Run: func(cmd *cobra.Command, args []string) {
		process, err := domain.NewProcess(uuid.New().String(), postulationType, platform, company, position, jobType)
		if err != nil {
			fmt.Println(err.Error())
		}
		b, err := json.Marshal(*process)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("The Process output is: %s\n", string(b))
	},
}

func init() {

	AddCmd.Flags().StringVarP(&postulationType, "postulation-type", "t", "", "the postulation type")
	err := AddCmd.MarkFlagRequired("postulation-type")
	if err != nil {
		fmt.Println(err.Error())
	}
	AddCmd.Flags().StringVarP(&platform, "platform", "", "", "the recruiter platform")
	err = AddCmd.MarkFlagRequired("platform")
	if err != nil {
		fmt.Println(err.Error())
	}
	AddCmd.Flags().StringVarP(&company, "company", "", "", "the recruiter company")
	err = AddCmd.MarkFlagRequired("company")
	if err != nil {
		fmt.Println(err.Error())
	}
	AddCmd.Flags().StringVarP(&position, "position", "", "", "The position name")
	err = AddCmd.MarkFlagRequired("position")
	if err != nil {
		fmt.Println(err.Error())
	}
	AddCmd.Flags().StringVarP(&jobType, "job-type", "j", "", "the job type")
	err = AddCmd.MarkFlagRequired("job-type")
	if err != nil {
		fmt.Println(err.Error())
	}
}
