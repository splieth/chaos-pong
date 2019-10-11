package chaos

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type mockEC2Client struct {
	ec2iface.EC2API
}

var counter = 0

func (m mockEC2Client) DescribeInstances(in *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	var nextToken *string = nil
	if counter < 2 {
		s := "token"
		nextToken = &s
	}
	counter += 2

	output := ec2.DescribeInstancesOutput{
		NextToken: nextToken,
		Reservations: []*ec2.Reservation{
			{
				Instances: []*ec2.Instance{
					{InstanceId: generateInstanceId()},
					{InstanceId: generateInstanceId()},
				},
			},
		},
	}
	return &output, nil
}

func (m mockEC2Client) TerminateInstances(in *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	output := ec2.TerminateInstancesOutput{
		TerminatingInstances: []*ec2.InstanceStateChange{{
			InstanceId: in.InstanceIds[0],
		}},
	}
	return &output, nil
}

func generateInstanceId() *string {
	id := "i-" + string(rand.Int())
	return &id
}

func TestListInstances(t *testing.T) {
	client := mockEC2Client{}
	c := listInstances(client, []ec2.Instance{}, nil)
	assert.Equal(t, 4, len(c))
}

func TestTerminateInstance(t *testing.T) {
	client := mockEC2Client{}
	instanceId := "stirb"
	res, _ := terminateInstance(client, []*string{&instanceId})
	assert.Equal(t, []*string{&instanceId}, res)
}
