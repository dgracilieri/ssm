package resource

import (
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/encoding"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client := ssm.New(req.Session)
	tags := make([]*ssm.Tag, 0)
	for _, t := range currentModel.Tags {
		tags = append(tags, &ssm.Tag{
			Key:   t.Key.Value(),
			Value: t.Value.Value(),
		})
	}

	input := &ssm.PutParameterInput{
		AllowedPattern: currentModel.AllowedPattern.Value(),
		Description:    currentModel.Description.Value(),
		KeyId:          currentModel.KeyId.Value(),
		Name:           currentModel.Name.Value(),
		Overwrite:      aws.Bool(false),
		Policies:       currentModel.Policies.Value(),
		Tags:           tags,
		Tier:           currentModel.Tier.Value(),
		Type:           aws.String(ssm.ParameterTypeSecureString),
		Value:          currentModel.Value.Value(),
	}

	_, err := client.PutParameter(input)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	// Construct a new handler.ProgressEvent and return it
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client := ssm.New(req.Session)
	input := &ssm.GetParameterInput{
		Name:           currentModel.Name.Value(),
		WithDecryption: aws.Bool(true),
	}
	parameter, err := client.GetParameter(input)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	// Assign the value
	currentModel.Value = encoding.NewString(*parameter.Parameter.Value)

	// Construct a new handler.ProgressEvent and return it
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client := ssm.New(req.Session)
	tags := make([]*ssm.Tag, 0)
	for _, t := range currentModel.Tags {
		tags = append(tags, &ssm.Tag{
			Key:   t.Key.Value(),
			Value: t.Value.Value(),
		})
	}

	input := &ssm.PutParameterInput{
		AllowedPattern: currentModel.AllowedPattern.Value(),
		Description:    currentModel.Description.Value(),
		KeyId:          currentModel.KeyId.Value(),
		Name:           currentModel.Name.Value(),
		Overwrite:      aws.Bool(true),
		Policies:       currentModel.Policies.Value(),
		Tags:           tags,
		Tier:           currentModel.Tier.Value(),
		Type:           aws.String(ssm.ParameterTypeSecureString),
		Value:          currentModel.Value.Value(),
	}

	_, err := client.PutParameter(input)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	// Construct a new handler.ProgressEvent and return it
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client := ssm.New(req.Session)

	input := &ssm.DeleteParameterInput{
		Name:           currentModel.Name.Value(),
	}
	_, err := client.DeleteParameter(input)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message: "Delete complete",
		ResourceModel: currentModel,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client := ssm.New(req.Session)
	input := &ssm.GetParameterInput{
		Name:           currentModel.Name.Value(),
		WithDecryption: aws.Bool(true),
	}
	parameter, err := client.GetParameter(input)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	// Assign the value
	currentModel.Value = encoding.NewString(*parameter.Parameter.Value)

	// Construct a new handler.ProgressEvent and return it
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read complete",
		ResourceModel:   currentModel,
	}, nil

}