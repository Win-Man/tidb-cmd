/*
 * Created: 2020/8/27 19:30
 * Author : Win-Man
 * Email : gang.shen0423@gmail.com
 * -----
 * Last Modified:
 * Modified By:
 * -----
 * Description:
 */

package command

import (
	"fmt"
	"github.com/Win-Man/tidb-cmd/pkg/logger"
	"github.com/pingcap/dm/pkg/utils"
	"github.com/spf13/cobra"
)

func newDecryptCmd() *cobra.Command {
	var str string
	logger.InitLogger()
	cmd := &cobra.Command{
		Use:   "decrypt --string={base64 string}",
		Short: "decrypt string",
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := decryptString(str)
			if err != nil {
				panic(err)
			} else {
				fmt.Printf("Plaintext: %s\n", res)
				return nil
			}

			return cmd.Help()
		},
	}
	cmd.Flags().StringVar(&str, "string", "", "The base64 encoded ciphertext")
	return cmd
}

// Decrypt tries to decrypt base64 encoded ciphertext to plaintext
func decryptString(str string) (string, error) {
	return utils.Decrypt(str)

}
