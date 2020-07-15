package add

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/dkaslovsky/MyMint/pkg/category"
	"github.com/dkaslovsky/MyMint/pkg/conf"
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/ledger"
	"github.com/doug-martin/goqu/v9"
	"github.com/spf13/cobra"
)

// Options are command options
type Options struct {
	Db          string
	Date        string
	Amount      float64
	Description string
	Category    string
	Positive    bool
}

// CreateAddCmd generates the configuration for the add subcommand
func CreateAddCmd() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add an entry to the ledger",
		RunE: func(cmd *cobra.Command, args []string) error {

			categories, err := category.LoadCategories(conf.Config.LedgerCategoryFilePath)
			if err != nil {
				return err
			}
			var category interface{}
			if opts.Category != "" {
				if !categories.Contains(opts.Category) {
					return fmt.Errorf("unknown category [%s] must be added before it can be used", opts.Category)
				}
				category = opts.Category
			}

			err = validateDate(opts.Date)
			if err != nil {
				return err
			}

			amount := setAmountSign(opts.Amount, opts.Positive)

			row := goqu.Record{
				"Date":        opts.Date,
				"Amount":      amount,
				"Description": opts.Description,
				"Category":    category,
			}

			db, err := sqlite.NewDb(opts.Db)
			if err != nil {
				return err
			}
			defer db.Close()

			_, id, err := db.InsertRows(ledger.TableName, []interface{}{row})
			if err != nil {
				return err
			}

			log.Printf("Inserted row: %v", &ledger.Row{
				ID:          id,
				Date:        opts.Date,
				Amount:      amount,
				Description: opts.Description,
				Category:    category,
			})
			return nil

		},
	}
	attachOpts(cmd, &opts)
	return cmd
}

func attachOpts(cmd *cobra.Command, opts *Options) {
	flags := cmd.Flags()
	flags.StringVarP(&opts.Db, "database", "", conf.Config.DefaultSqliteDbPath, "Name of database")
	flags.StringVarP(&opts.Date, "date", "d", "", "Entry date")
	flags.Float64VarP(&opts.Amount, "amount", "a", 0, "Entry amount in dollars")
	flags.StringVarP(&opts.Description, "description", "r", "", "Entry description")
	flags.StringVarP(&opts.Category, "category", "c", "", "Entry category")
	flags.BoolVarP(&opts.Positive, "positive", "p", false, "Entry amount is positive dollar value")
	cmd.MarkFlagRequired("date")
	cmd.MarkFlagRequired("amount")
}

func setAmountSign(amount float64, positive bool) (signedAmount float64) {
	if positive {
		return math.Abs(amount)
	}
	return -1 * math.Abs(amount)
}

const rfc3339FullDate = "2006-01-02"

func validateDate(date string) (err error) {
	_, err = time.Parse(rfc3339FullDate, date)
	return err
}
