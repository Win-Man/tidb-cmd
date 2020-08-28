/*
 * Created: 2020/8/25 21:03
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
	"database/sql"
	"errors"
	"fmt"
	"github.com/Win-Man/tidb-cmd/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

type DBConfig struct {
	user     string
	password string
	port     int
	host     string
}

func newAnalyzeCmd() *cobra.Command {
	var cnf DBConfig
	var dbList string
	var tableList string
	logger.InitLogger()
	cmd := &cobra.Command{
		Use:   "analyze -u -p -P",
		Short: "analyze table",
		RunE: func(cmd *cobra.Command, args []string) error {
			//if len(args) < 1 {
			//	log.Infof("args less than 1")
			//	return cmd.Help()
			//}
			err := checkListArgs(dbList, tableList)
			if err != nil {
				return err
			}
			analyzeDBs(cnf, dbList, tableList)
			return nil
		},
	}
	cmd.Flags().StringVar(&cnf.host, "host", "", "The host ip")
	cmd.Flags().StringVar(&cnf.user, "user", "", "The user name")
	cmd.Flags().StringVar(&cnf.password, "password", "", "The user password")
	cmd.Flags().IntVar(&cnf.port, "port", 0, "The port")
	cmd.Flags().StringVar(&dbList, "db-list", "", "The db list")
	cmd.Flags().StringVar(&tableList, "table-list", "", "The table list")
	return cmd
}

func analyzeDBs(cnf DBConfig, dbList string, tableList string) {
	log.Debug("analyzeDBs")
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=%s", cnf.user, cnf.password, cnf.host, cnf.port, "utf8"))
	if err != nil {
		panic(err)
	}
	dbSlice := strings.Split(dbList, ",")
	tableSlice := strings.Split(tableList, ",")
	log.Debugf("dbSlice:%+v tableSlice:%+v", dbSlice, tableSlice)
	log.Debugf("len(dbSlice):%d len(tableSlice):%d", len(dbSlice), len(tableSlice))
	if tableList == "" && len(dbSlice) > 0 {
		// analyze db

		for _, db := range dbSlice {
			log.Infof(fmt.Sprintf("Start Analyze DB:%s", db))
			rows, err := conn.Query(fmt.Sprintf("show tables in %s", db))
			log.Debugf("rows:%+v", rows)
			if err != nil {
				panic(err)
			}
			for rows.Next() {
				e := rows.Scan()
				fmt.Println(e)
			}
			log.Infof(fmt.Sprintf("Finish Analyze db:%s", db))
		}
	} else if len(tableSlice) > 0 && len(dbSlice) == 0 {
		// analyze table

	} else if len(tableSlice) > 1 && len(dbSlice) == 0 {
		//

	}

}

func checkListArgs(dbList string, tableList string) error {
	log.Debug("checkListArgs")
	dbSlice := strings.Split(dbList, ",")
	tableSlice := strings.Split(tableList, ",")
	if len(tableSlice) > 1 && len(dbSlice) > 1 {
		log.Errorf("db-list should be configed no more than one database when congfiged table-list option")
		return errors.New("db-list should be configed no more than one database when congfiged table-list option")
	}
	return nil
}
