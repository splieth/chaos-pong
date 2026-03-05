package aws

import (
	"slices"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/splieth/chaos-pong/chaos"
	"log"
	"math/rand"
)

func init() {
	chaos.RegisterProvider("aws", newAWSChaos)
}

func newAWSChaos(cfg chaos.ProviderConfig) []chaos.Chaos {
	client := ec2ClientWithConfig(cfg.Options["region"], cfg.Options["profile"])

	var fns []chaos.Chaos
	if len(cfg.Actions) == 0 || slices.Contains(cfg.Actions, "ec2-instance-terminate") {
		fns = append(fns, EC2InstanceTerminateChaos{client: client})
	}
	if len(cfg.Actions) == 0 || slices.Contains(cfg.Actions, "ebs-destroy") {
		fns = append(fns, EBSDestroyChaos{client: client})
	}
	return fns
}

type EC2InstanceTerminateChaos struct {
	client *ec2.EC2
}

type EBSDestroyChaos struct {
	client *ec2.EC2
}

func (e EBSDestroyChaos) Terminate() chaos.Result {
	instances := listInstances(e.client, []ec2.Instance{}, nil)
	instance := instances[rand.Intn(len(instances))]
	_, err := terminateEbsVolume(e.client, instance.BlockDeviceMappings[0].Ebs.VolumeId)
	if err != nil {
		log.Fatal(err)
	}
	return chaos.Result{Success: true, Message: "Terminated EBS volume"}
}

func terminateEbsVolume(client *ec2.EC2, volumeId *string) (*string, error) {
	detachInput := ec2.DetachVolumeInput{
		Force:    aws.Bool(true),
		VolumeId: volumeId,
	}
	_, err := client.DetachVolume(&detachInput)
	if err != nil {
		return aws.String(""), err
	}

	input := ec2.DeleteVolumeInput{
		VolumeId: volumeId,
	}
	_, err = client.DeleteVolume(&input)
	if err != nil {
		return aws.String(""), err
	}
	return volumeId, nil
}

func (e EC2InstanceTerminateChaos) Terminate() chaos.Result {
	instances := listInstances(e.client, []ec2.Instance{}, nil)
	instance := instances[rand.Intn(len(instances))]
	_, err := terminateInstance(e.client, []*string{instance.InstanceId})
	if err != nil {
		log.Fatal(err)
	}
	return chaos.Result{Success: true, Message: "Terminated EC2 instance"}
}

func ec2ClientWithConfig(region, profile string) *ec2.EC2 {
	opts := session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}
	if profile != "" {
		opts.Profile = profile
	}
	awsCfg := aws.Config{}
	if region != "" {
		awsCfg.Region = aws.String(region)
	}
	opts.Config = awsCfg

	sess := session.Must(session.NewSessionWithOptions(opts))
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
