package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func AccesKeyLastUsed(client *iam.Client) (*iam.GetAccessKeyLastUsedOutput, error) {
	input := &iam.GetAccessKeyLastUsedInput{
		AccessKeyId: aws.String("Access_Key_Id"),
	}

	output, err := client.GetAccessKeyLastUsed(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func AttachUserPolicy(client *iam.Client) (*iam.AttachUserPolicyOutput, error) {
	input := &iam.AttachUserPolicyInput{
		PolicyArn: aws.String(""),
		UserName:  aws.String(""),
	}

	output, err := client.AttachUserPolicy(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil

}

func CreateAccessKey(client *iam.Client) (*iam.CreateAccessKeyOutput, error) {
	input := &iam.CreateAccessKeyInput{
		UserName: aws.String(""),
	}

	output, err := client.CreateAccessKey(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func CreateUser(client *iam.Client) (*iam.CreateUserOutput, error) {

	input := &iam.CreateUserInput{
		UserName:            aws.String("NewUser"),
		Path:                aws.String(""),
		PermissionsBoundary: aws.String(""),
		Tags:                []types.Tag{},
	}

	output, err := client.CreateUser(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func CreateAccountAllias(client *iam.Client) (*iam.CreateAccountAliasOutput, error) {
	input := &iam.CreateAccountAliasInput{
		AccountAlias: aws.String(""),
	}

	output, err := client.CreateAccountAlias(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func CreatePolicy(client *iam.Client) (*iam.CreatePolicyOutput, error) {

	input := &iam.CreatePolicyInput{
		PolicyDocument: aws.String(""),
		PolicyName:     aws.String(""),
		Description:    aws.String(""),
		Path:           aws.String(""),
		Tags:           []types.Tag{},
	}

	output, err := client.CreatePolicy(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func UploadCert(client *iam.Client) (*iam.UploadServerCertificateOutput, error) {
	input := &iam.UploadServerCertificateInput{
		CertificateBody:       aws.String(""),
		PrivateKey:            aws.String(""),
		ServerCertificateName: aws.String(""),
		CertificateChain:      aws.String(""),
		Path:                  aws.String(""),
		Tags:                  []types.Tag{},
	}

	output, err := client.UploadServerCertificate(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func DeleteUser(client *iam.Client) (*iam.DeleteUserOutput, error) {

	input := &iam.DeleteUserInput{
		UserName: aws.String("NewUser"),
	}

	output, err := client.DeleteUser(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func DeleteAcceskey(client *iam.Client) (*iam.DeleteAccessKeyOutput, error) {
	input := &iam.DeleteAccessKeyInput{
		AccessKeyId: aws.String(""),
		UserName:    aws.String(""),
	}

	output, err := client.DeleteAccessKey(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func DeleteAccountAllias(client *iam.Client) (*iam.DeleteAccountAliasOutput, error) {
	input := &iam.DeleteAccountAliasInput{
		AccountAlias: aws.String(""),
	}

	output, err := client.DeleteAccountAlias(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func DeleteServerCert(client *iam.Client) (*iam.DeleteServerCertificateOutput, error) {
	input := &iam.DeleteServerCertificateInput{
		ServerCertificateName: aws.String(""),
	}

	output, err := client.DeleteServerCertificate(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func DeleteUserPolicy(client *iam.Client) (*iam.DeleteUserPolicyOutput, error) {
	input := &iam.DeleteUserPolicyInput{
		PolicyName: aws.String(""),
		UserName:   aws.String(""),
	}

	output, err := client.DeleteUserPolicy(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func DeleteLogin(client *iam.Client) (*iam.DeleteLoginProfileOutput, error) {
	input := &iam.DeleteLoginProfileInput{
		UserName: aws.String(""),
	}

	output, err := client.DeleteLoginProfile(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func DetachUserPolicy(client *iam.Client) (*iam.DetachUserPolicyOutput, error) {
	input := &iam.DetachUserPolicyInput{
		UserName:  aws.String(""),
		PolicyArn: aws.String(""),
	}

	output, err := client.DetachUserPolicy(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func DeleteServerCerts(client *iam.Client) (*iam.DeleteServerCertificateOutput, error) {
	input := &iam.DeleteServerCertificateInput{
		ServerCertificateName: aws.String(""),
	}

	output, err := client.DeleteServerCertificate(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func GetPolicy(client *iam.Client) (*iam.GetPolicyOutput, error) {
	input := &iam.GetPolicyInput{
		PolicyArn: aws.String(""),
	}

	output, err := client.GetPolicy(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func GetServerCert(client *iam.Client) (*iam.GetServerCertificateOutput, error) {
	input := &iam.GetServerCertificateInput{
		ServerCertificateName: aws.String(""),
	}

	output, err := client.GetServerCertificate(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func GetLogin(client *iam.Client) (*iam.GetLoginProfileOutput, error) {
	input := &iam.GetLoginProfileInput{
		UserName: aws.String(""),
	}

	output, err := client.GetLoginProfile(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func CreateLogin(client *iam.Client) (*iam.CreateLoginProfileOutput, error) {
	input := &iam.CreateLoginProfileInput{
		Password:              aws.String(""),
		UserName:              aws.String(""),
		PasswordResetRequired: *aws.Bool(false),
	}

	output, err := client.CreateLoginProfile(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func ListAccessKeys(client *iam.Client) (*iam.ListAccessKeysOutput, error) {
	input := &iam.ListAccessKeysInput{
		Marker:   aws.String(""),
		MaxItems: aws.Int32(10),
		UserName: aws.String(""),
	}

	output, err := client.ListAccessKeys(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func ListAccountAllias(client *iam.Client) (*iam.ListAccountAliasesOutput, error) {
	input := &iam.ListAccountAliasesInput{
		Marker:   aws.String(""),
		MaxItems: aws.Int32(10),
	}

	output, err := client.ListAccountAliases(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// users with extra previlages
func ListAdmins(client *iam.Client) (*iam.ListEntitiesForPolicyOutput, error) {
	input := &iam.ListEntitiesForPolicyInput{
		PolicyArn:         aws.String(""),
		EntityFilter:      types.EntityTypeAWSManagedPolicy,
		Marker:            aws.String(""),
		MaxItems:          aws.Int32(10),
		PathPrefix:        aws.String(""),
		PolicyUsageFilter: types.PolicyUsageTypePermissionsPolicy,
	}

	output, err := client.ListEntitiesForPolicy(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func ListServerCerts(client *iam.Client) (*iam.ListServerCertificatesOutput, error) {
	input := &iam.ListServerCertificatesInput{
		Marker:     aws.String(""),
		MaxItems:   aws.Int32(10),
		PathPrefix: aws.String(""),
	}

	output, err := client.ListServerCertificates(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func ListUsers(client *iam.Client) (*iam.ListUsersOutput, error) {
	input := &iam.ListUsersInput{
		Marker:     aws.String(""),
		MaxItems:   aws.Int32(10),
		PathPrefix: aws.String(""),
	}

	output, err := client.ListUsers(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func UpdateAccessKey(client *iam.Client) (*iam.UpdateAccessKeyOutput, error) {
	input := &iam.UpdateAccessKeyInput{
		AccessKeyId: aws.String(""),
		Status:      types.StatusTypeActive,
		UserName:    aws.String(""),
	}

	output, err := client.UpdateAccessKey(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func UpdateUser(client *iam.Client) (*iam.UpdateUserOutput, error) {
	input := &iam.UpdateUserInput{
		UserName:    aws.String(""),
		NewPath:     aws.String(""),
		NewUserName: aws.String(""),
	}

	output, err := client.UpdateUser(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func UpdateServerCerts(client *iam.Client) (*iam.UpdateServerCertificateOutput, error) {
	input := &iam.UpdateServerCertificateInput{
		ServerCertificateName:    aws.String(""),
		NewPath:                  aws.String(""),
		NewServerCertificateName: aws.String(""),
	}

	output, err := client.UpdateServerCertificate(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}
