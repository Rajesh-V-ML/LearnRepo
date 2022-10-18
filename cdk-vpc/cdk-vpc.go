package main

import (
	"cdk-vpc/vpc"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func main() {
	app := awscdk.NewApp(nil)

	vpc.NewVpcStack(app, "vpc-stack", &vpc.VpcStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)

}

func env() *awscdk.Environment {
	return nil

}
