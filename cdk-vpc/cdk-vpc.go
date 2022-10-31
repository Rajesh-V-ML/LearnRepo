package main

import (
	"cdk-vpc/vpc"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
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

	return &awscdk.Environment{
		Account: jsii.String("469875231790"),
		Region:  jsii.String("ap-south-1"),
	}

}
