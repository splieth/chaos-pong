package chaos

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"log"
)

type EC2Client struct {
	ec2Client *ec2.EC2
}

func (c *EC2Client) CreateClient() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	c.ec2Client = ec2.New(sess)
}

func hugo(svc ec2iface.EC2API) bool {
	return false
}
func ListInstances(svc ec2iface.EC2API, instances []ec2.Instance, nextToken *string) []ec2.Instance {
	maxResults := int64(100)
	input := ec2.DescribeInstancesInput{
		NextToken:  nextToken,
		MaxResults: &maxResults,
	}
	res, err := svc.DescribeInstances(&input)
	if err != nil {
		log.Fatal(err)
	}

	for _, reservation := range res.Reservations {
		for _, instance := range reservation.Instances {
			instances = append(instances, *instance)
		}
	}

	if res.NextToken != nil {
		return ListInstances(svc, instances, res.NextToken)
	}

	return instances
}

func Terminate(svc ec2iface.EC2API) {
	instances := ListInstances(svc, []ec2.Instance{}, nil)
	fmt.Println(len(instances))
}
