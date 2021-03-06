package removesubcommands

import (
	"fmt"
	"os"

	snmpsimclient "github.com/inexio/snmpsim-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AgentFromLabCmd represents the agentFromLab command
var AgentFromLabCmd = &cobra.Command{
	Use:   "agent-from-lab",
	Args:  cobra.ExactArgs(0),
	Short: "Removes an agent from a lab",
	Long:  `Removes the agent with the given agent-id from the lab with the given lab-id`,
	Run: func(cmd *cobra.Command, args []string) {
		//Load the client data from the config
		baseUrl := viper.GetString("mgmt.http.baseUrl")
		username := viper.GetString("mgmt.http.authUsername")
		password := viper.GetString("mgmt.http.authPassword")

		//Create a new client
		client, err := snmpsimclient.NewManagementClient(baseUrl)
		if err != nil {
			log.Error().
				Msg("Error while creating management client")
			os.Exit(1)
		}
		err = client.SetUsernameAndPassword(username, password)
		if err != nil {
			log.Error().
				Msg("Error while setting username and password")
			os.Exit(1)
		}

		//Read in the agent-id
		agentId, err := cmd.Flags().GetInt("agent")
		if err != nil {
			log.Error().
				Msg("Error while retrieving agentId")
			os.Exit(1)
		}

		//Read in the lab-id
		labId, err := cmd.Flags().GetInt("lab")
		if err != nil {
			log.Error().
				Msg("Error while retrieving labId")
			os.Exit(1)
		}

		//Remove the agent from the lab
		err = client.RemoveAgentFromLab(labId, agentId)
		if err != nil {
			log.Error().
				Msg("Error while removing the agent from the lab")
			os.Exit(1)
		}
		fmt.Println("Agent", agentId, "has been removed from lab", labId)
	},
}

func init() {
	//Set agent flag
	AgentFromLabCmd.Flags().Int("agent", 0, "Id of the agent that is to be removed from the lab")
	err := AgentFromLabCmd.MarkFlagRequired("agent")
	if err != nil {
		log.Error().
			Msg("Could not mark 'agent' flag required")
		os.Exit(1)
	}

	//Set lab flag
	AgentFromLabCmd.Flags().Int("lab", 0, "Id of the lab the agent will be removed from")
	err = AgentFromLabCmd.MarkFlagRequired("lab")
	if err != nil {
		log.Error().
			Msg("Could not mark 'lab' flag required")
		os.Exit(1)
	}
}
