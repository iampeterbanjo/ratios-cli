package main

import (
  "os"
  "fmt"
  "strconv"
  "github.com/codegangsta/cli"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// aka greatest common divisor
func HighestCommonFactor(a int, b int) int {
  if b == 0 {
    return a
  }

  return HighestCommonFactor(b, a % b)
}

func main() {
  app := cli.NewApp()
	app.Name = "Ratios Cli"
	app.Version = "0.0.1"
	app.Usage = `Calculate ratios given some variables
              by solving the expression a : b = x : y`

	app.Flags = []cli.Flag {
		cli.StringFlag {
			Name: "a",
			Usage: "Variable a",
		},
		cli.StringFlag {
			Name: "b",
			Usage: "Variable b",
		},
		cli.StringFlag {
			Name: "x",
			Usage: "Variable x",
		},
		cli.StringFlag {
			Name: "y",
			Usage: "Variable y",
		},
  }

  app.Action = func(c *cli.Context) {
    a := c.String("a")
    b := c.String("b")

    if a != "" && a != "" {
      x, err := strconv.Atoi(a)
      check(err)

      y, err := strconv.Atoi(b)
      check(err)

      factor := HighestCommonFactor(x, y)

      fmt.Printf("%s : %s = %s : %s \n", a, b, strconv.Itoa(x/factor), strconv.Itoa(y/factor))
    }
  }

  app.Run(os.Args)
}
