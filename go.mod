module github.com/Win-Man/tidb-cmd

replace (
	github.com/Win-Man/tidb-cmd/command => ./tidb-cmd/command
	github.com/Win-Man/tidb-cmd/pkg/logger => ./tidb-cmd/pkg/logger
)

go 1.15

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/pingcap/dm v1.0.6
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
)
