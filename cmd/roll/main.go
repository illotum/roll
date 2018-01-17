package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/illotum/roll"
	"github.com/spf13/cobra"
)

var version = "nightly"

func main() {
	log.SetFlags(0)

	var seed int64
	root := &cobra.Command{
		Use:   "roll FILE",
		Short: "Roll on a table",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := roll.ReadFile(args[0])
			if err != nil {
				log.Fatal(err)
			}
			err = r.Text.ExecuteTemplate(os.Stdout, "", r.Data)
			if err != nil {
				log.Fatal(err)
			}
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			rand.Seed(seed)
		},
	}
	root.PersistentFlags().Int64VarP(&seed, "seed", "s", time.Now().Unix(), "random number generator `SEED` (default current timestamp)")

	root.Execute()
}
