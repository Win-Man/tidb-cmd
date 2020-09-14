package command

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func newTestCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "test <name>",
		Short: "test toolkit",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return cmd.Help()
			}
			switch args[0] {
			case "log":
				logTest()
			case "tidb":
				tidbTest()
			default:
				return cmd.Help()
			}
			return nil
		},
	}
	return cmd
}

func logTest(){

}

func tidbTest() {
	rownum := 1000000
	batchSize := 5000
	totalCost := time.Now()
	//tidbUri := "root:@tcp(127.0.0.1:4000)/gangshen?charset=utf8"
	mysqlUri := "root:letsg0@tcp(127.0.0.1:3308)/gangshen?charset=utf8"
	conn, err := sql.Open("mysql", mysqlUri)
	if err != nil{
		panic(err)
	}
	if err = conn.Ping();err != nil{
		fmt.Println("open database fail")
	}
	count := 0
	for {
		txnCost := time.Now()
		//开启事务
		tx, err := conn.Begin()
		if err != nil {
			fmt.Println("tx fail")
		}
		//准备sql语句
		stmt, err := tx.Prepare("INSERT INTO tid (`id`) VALUES (?)")
		if err != nil {
			fmt.Println("Prepare fail")
		}
		for i := 0; i < batchSize; i++ {

			//将参数传递到sql语句中并且执行
			_, err := stmt.Exec(count+i)
			if err != nil {
				fmt.Println("Exec fail")
			}
		}

		//将事务提交
		tx.Commit()
		count += batchSize
		txnTime := time.Since(txnCost)
		fmt.Printf("Total Inserted %d rows, Insert %d rows Cost %s\n",count,batchSize,txnTime)
		if count >= rownum{
			break
		}
	}
	totalTime := time.Since(totalCost)
	fmt.Printf("Total cost time:%s\n",totalTime)
}
