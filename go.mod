module github.com/Win-Man/tidb-cmd

replace github.com/Win-Man/tidb-cmd/command => ./tidb-cmd/command

go 1.15

require (
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
)
