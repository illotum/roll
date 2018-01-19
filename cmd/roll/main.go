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
			fn := args[0]
			r, err := roll.ReadFile(fn)
			if err != nil {
				log.Fatalf("Error encountered while parsing %q: %s", fn, err)
			}
			err = r.Template.ExecuteTemplate(os.Stdout, "", r.Tables)
			if err != nil {
				log.Fatalf("Error encountered while executing %q: %s", fn, err)
			}
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			rand.Seed(seed)
		},
	}
	root.PersistentFlags().Int64VarP(&seed, "seed", "s", time.Now().Unix(), "random number generator `SEED`")

	root.Execute()
}
