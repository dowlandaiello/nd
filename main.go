package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/adrg/xdg"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

// Time represents a link source corresponding to a particular time.
type Time struct {
	From string

	Cmd *string
}

// Link represents a conditional symlink.
type Link struct {
	Day   *Time
	Night *Time

	To string
}

// NormalizePath expands the path contained in a link.
func (l *Link) NormalizePath() error {
	var err error

	l.To, err = homedir.Expand(l.To)
	if err != nil {
		return err
	}

	if l.Day != nil {
		l.Day.From, err = homedir.Expand(l.Day.From)
	}

	if l.Night != nil {
		l.Night.From, err = homedir.Expand(l.Night.From)
	}

	return err
}

// Config is a configuration file representation comprised by a slice of links.
type Config map[string]Link

func main() {
	app := &cli.App{
		Name:  "nd",
		Usage: "toggle between night and day configurations for various tools",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "v",
				Value: false,
				Usage: "Displays runtime errors from commands attached to links",
			},
		},
		Action: func(c *cli.Context) error {
			time := c.Args().Get(0)

			linkfile, err := os.Open(fmt.Sprintf("%s/nd/links.yml", xdg.ConfigHome))
			if err != nil {
				return err
			}

			links := Config{}

			dec := yaml.NewDecoder(linkfile)
			err = dec.Decode(&links)

			if err != nil {
				return err
			}

			for i, link := range links {
				err = link.NormalizePath()
				if err != nil {
					return err
				}

				var linkNode *Time

				switch time {
				case "day":
					if link.Day == nil {
						return fmt.Errorf("link %s doesn't have a daytime source", i)
					}

					linkNode = link.Day
				case "night":
					if link.Night == nil {
						return fmt.Errorf("link %s doesn't have a nighttime source", i)
					}

					linkNode = link.Night
				default:
					return fmt.Errorf("unrecognized time %s", time)
				}

				if _, err := os.Stat(link.To); err == nil {
					err = os.Remove(link.To)
					if err != nil {
						return err
					}
				}

				if err := os.Link(linkNode.From, link.To); err != nil {
					return err
				}

				if linkNode.Cmd == nil {
					continue
				}

				raw := strings.Split(*linkNode.Cmd, " ")
				cmd := exec.Command(raw[0], raw[1:]...)

				err = cmd.Run()
				if err != nil && c.Bool("v") {
					log.Printf("encountered an error while executing command for link %s: %s\n", i, err)
				}
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
