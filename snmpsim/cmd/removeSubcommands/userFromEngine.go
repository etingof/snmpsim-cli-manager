package removesubcommands

import (
	"fmt"
	"os"

	snmpsimclient "github.com/inexio/snmpsim-restapi-go-client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// UserFromEngineCmd represents the userFromEngine command
var UserFromEngineCmd = &cobra.Command{
	Use:   "user-from-engine",
	Args:  cobra.ExactArgs(0),
	Short: "Removes an user from an engine",
	Long:  `Removes the user with the given user-id from the engine with the give engine-id`,
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

		//Read in the user-id
		userId, err := cmd.Flags().GetInt("user")
		if err != nil {
			log.Error().
				Msg("Error while retrieving userId")
			os.Exit(1)
		}

		//Read in the engine-id
		engineId, err := cmd.Flags().GetInt("engine")
		if err != nil {
			log.Error().
				Msg("Error while retrieving engineId")
			os.Exit(1)
		}

		//Remove the user from the engine
		err = client.RemoveUserFromEngine(engineId, userId)
		if err != nil {
			log.Error().
				Msg("Error during removal of the user from the engine")
			os.Exit(1)
		}
		fmt.Println("User", userId, "has been removed from engine", engineId)
	},
}

func init() {
	//Set user flag
	UserFromEngineCmd.Flags().Int("user", 0, "Id of the user that is to be removed from the engine")
	err := UserFromEngineCmd.MarkFlagRequired("user")
	if err != nil {
		log.Error().
			Msg("Could not mark 'user' flag required")
		os.Exit(1)
	}

	//Set engine flag
	UserFromEngineCmd.Flags().Int("engine", 0, "Id of the engine from that the user will be removed")
	err = UserFromEngineCmd.MarkFlagRequired("engine")
	if err != nil {
		log.Error().
			Msg("Could not mark 'engine' flag required")
		os.Exit(1)
	}
}
