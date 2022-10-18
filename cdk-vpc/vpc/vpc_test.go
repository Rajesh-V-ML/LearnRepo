package vpc_test

import (
	"encoding/json"
	"testing"

	"cdk-vpc/vpc"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestVpcStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := vpc.NewVpcStack(app, "vpc-stack", nil)

	// THEN
	bytes, err := json.Marshal(app.Synth(nil).GetStackArtifact(stack.ArtifactId()).Template())
	if err != nil {
		t.Error(err)
	}

	template := gjson.ParseBytes(bytes)
	cidr := template.Get("Resources.rajevenuvpc8AFD4B1C.Properties.CidrBlock").String()
	assert.Equal(t, "10.0.0.0/16", cidr)
}
