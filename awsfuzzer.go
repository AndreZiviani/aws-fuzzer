package Awsfuzzer

import (
	"fmt"
	"os"
	"text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Awsfuzzer struct {
	ec2      *ec2.EC2
	options  Options
	template *template.Template
}

func New() (*Awsfuzzer, error) {
	options, err := ParseOptions()
	if err != nil {
		return nil, err
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(options.Region),
		},
	})
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("Instance").Parse(options.Template)

	return &Awsfuzzer{
		ec2:      ec2.New(sess),
		options:  options,
		template: tmpl,
	}, nil
}

func (e *Awsfuzzer) Run() {
	instances, err := e.ListInstances()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, i := range instances {
		fmt.Println(i)
		// write cache
	}

	// run fzf
}
