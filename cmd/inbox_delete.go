package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sakib0hasan/yogo/inbox"
)

// inboxDeleteCmd delete an email in inbox
var inboxDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete email at given position in inbox",
	Run: func(cmd *cobra.Command, args []string) {

		identifier, offset := parseMailAndOffsetArgs(args)

		in, err := inbox.ParseInboxPages(identifier, offset)

		if err != nil {
			perror(err)

			errorExit()
		}

		checkOffset(in.Count(), offset)

		in.Delete(offset - 1)
		success(fmt.Sprintf(`Email "%d" successfully deleted`, offset))
	},
}

func init() {
	inboxCmd.AddCommand(inboxDeleteCmd)
}
