package main

import (
	"collatz/command"
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	Num    = "num"
	Output = "output"
	End    = "end"
)

func main() {
	app := &cli.App{
		Name:  "collatz",
		Usage: "collatz calculator",
		Flags: []cli.Flag{
			&cli.Uint64Flag{
				Name:    Num,
				Aliases: []string{Num[0:1]},
				Value:   0,
				Usage:   "num",
			},
			&cli.StringFlag{
				Name:    Output,
				Aliases: []string{Output[0:1]},
				Value:   "",
				Usage:   "output file path",
			},
			&cli.Uint64Flag{
				Name:    End,
				Aliases: []string{End[0:1]},
				Value:   0,
				Usage:   "start num",
			},
		},
		Action: func(c *cli.Context) error {
			num := c.Uint64(Num)
			output := c.String(Output)
			end := c.Uint64(End)
			if num < 1 {
				return errors.New("num must be 1 or more")
			}
			err := command.Collatz(num, command.Output(output), command.End(end))
			if err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success")
}
