package Awsfuzzer

import (
	"fmt"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/alecthomas/kingpin.v2"
)

type Options struct {
	Region       string
	UsePrivateIp bool
	Template     string
	Filters      []string
}

func ParseOptions() (Options, error) {
	options := Options{
		Region:       "us-east-1",
		UsePrivateIp: false,
		Template:     `{{index .Tags "Name"}}`,
	}

	path, err := homedir.Expand("~/.config/aws-fuzzer")
	if err != nil {
		return Options{}, err
	}
	toml.DecodeFile(path, &options)

	region := kingpin.Flag("region", "The AWS region").Default(options.Region).String()
	usePrivateIp := kingpin.Flag("private", "return the private IP address of the instance rather than the public dns").Default(strconv.FormatBool(options.UsePrivateIp)).Bool()
	template := kingpin.Flag("template", "Template").Default(options.Template).String()
	version := kingpin.Flag("version", "Show the version of aws-fuzzer").Default("false").Bool()
	filters := kingpin.Flag("filters", "Ec2 describe-instance filters").Strings()

	kingpin.Parse()

	if *version {
		fmt.Printf("Ec2-fzf version %s\n", "1.0")
		os.Exit(1)
	}

	return Options{
		Region:       *region,
		UsePrivateIp: *usePrivateIp,
		Template:     *template,
		Filters:      *filters,
	}, nil
}
