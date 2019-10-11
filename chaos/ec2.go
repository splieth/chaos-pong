package chaos

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type EC2InstanceTerminateChaos struct {
	Chaos
}

func (e EC2InstanceTerminateChaos) Terminate() Result {
	client := ec2Client()
	instances := listInstances(client, []ec2.Instance{}, nil)
	instance := instances[rand.Intn(len(instances))]
	_, err := terminateInstance(client, []*string{instance.InstanceId})
	if err != nil {
		log.Fatal(err)
	}
	return Result{Success: true, Message: "Terminated stuff"}

}

func ec2Client() *ec2.EC2 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return ec2.New(sess)
}

func listInstances(svc ec2iface.EC2API, instances []ec2.Instance, nextToken *string) []ec2.Instance {
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
		return listInstances(svc, instances, res.NextToken)
	}

	return instances
}

func terminateInstance(svc ec2iface.EC2API, instanceIds []*string) ([]*string, error) {
	input := ec2.TerminateInstancesInput{
		InstanceIds: instanceIds,
	}
	out, err := svc.TerminateInstances(&input)
	if err != nil {
		return nil, err
	}
	var res []*string
	for _, instanceId := range out.TerminatingInstances {
		res = append(res, instanceId.InstanceId)
	}
	return res, nil
}
