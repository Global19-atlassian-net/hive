// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	ec2 "github.com/aws/aws-sdk-go/service/ec2"
	elb "github.com/aws/aws-sdk-go/service/elb"
	iam "github.com/aws/aws-sdk-go/service/iam"
	resourcegroupstaggingapi "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	route53 "github.com/aws/aws-sdk-go/service/route53"
	s3 "github.com/aws/aws-sdk-go/service/s3"
	s3iface "github.com/aws/aws-sdk-go/service/s3/s3iface"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// DescribeAvailabilityZones mocks base method
func (m *MockClient) DescribeAvailabilityZones(arg0 *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeAvailabilityZones", arg0)
	ret0, _ := ret[0].(*ec2.DescribeAvailabilityZonesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAvailabilityZones indicates an expected call of DescribeAvailabilityZones
func (mr *MockClientMockRecorder) DescribeAvailabilityZones(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAvailabilityZones", reflect.TypeOf((*MockClient)(nil).DescribeAvailabilityZones), arg0)
}

// DescribeImages mocks base method
func (m *MockClient) DescribeImages(arg0 *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeImages", arg0)
	ret0, _ := ret[0].(*ec2.DescribeImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeImages indicates an expected call of DescribeImages
func (mr *MockClientMockRecorder) DescribeImages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImages", reflect.TypeOf((*MockClient)(nil).DescribeImages), arg0)
}

// DescribeVpcs mocks base method
func (m *MockClient) DescribeVpcs(arg0 *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeVpcs", arg0)
	ret0, _ := ret[0].(*ec2.DescribeVpcsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVpcs indicates an expected call of DescribeVpcs
func (mr *MockClientMockRecorder) DescribeVpcs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcs", reflect.TypeOf((*MockClient)(nil).DescribeVpcs), arg0)
}

// DescribeSubnets mocks base method
func (m *MockClient) DescribeSubnets(arg0 *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSubnets", arg0)
	ret0, _ := ret[0].(*ec2.DescribeSubnetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSubnets indicates an expected call of DescribeSubnets
func (mr *MockClientMockRecorder) DescribeSubnets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSubnets", reflect.TypeOf((*MockClient)(nil).DescribeSubnets), arg0)
}

// DescribeSecurityGroups mocks base method
func (m *MockClient) DescribeSecurityGroups(arg0 *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSecurityGroups", arg0)
	ret0, _ := ret[0].(*ec2.DescribeSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSecurityGroups indicates an expected call of DescribeSecurityGroups
func (mr *MockClientMockRecorder) DescribeSecurityGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecurityGroups", reflect.TypeOf((*MockClient)(nil).DescribeSecurityGroups), arg0)
}

// RunInstances mocks base method
func (m *MockClient) RunInstances(arg0 *ec2.RunInstancesInput) (*ec2.Reservation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunInstances", arg0)
	ret0, _ := ret[0].(*ec2.Reservation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunInstances indicates an expected call of RunInstances
func (mr *MockClientMockRecorder) RunInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInstances", reflect.TypeOf((*MockClient)(nil).RunInstances), arg0)
}

// DescribeInstances mocks base method
func (m *MockClient) DescribeInstances(arg0 *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeInstances", arg0)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstances indicates an expected call of DescribeInstances
func (mr *MockClientMockRecorder) DescribeInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstances", reflect.TypeOf((*MockClient)(nil).DescribeInstances), arg0)
}

// TerminateInstances mocks base method
func (m *MockClient) TerminateInstances(arg0 *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateInstances", arg0)
	ret0, _ := ret[0].(*ec2.TerminateInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TerminateInstances indicates an expected call of TerminateInstances
func (mr *MockClientMockRecorder) TerminateInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateInstances", reflect.TypeOf((*MockClient)(nil).TerminateInstances), arg0)
}

// StopInstances mocks base method
func (m *MockClient) StopInstances(arg0 *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopInstances", arg0)
	ret0, _ := ret[0].(*ec2.StopInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopInstances indicates an expected call of StopInstances
func (mr *MockClientMockRecorder) StopInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopInstances", reflect.TypeOf((*MockClient)(nil).StopInstances), arg0)
}

// StartInstances mocks base method
func (m *MockClient) StartInstances(arg0 *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartInstances", arg0)
	ret0, _ := ret[0].(*ec2.StartInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartInstances indicates an expected call of StartInstances
func (mr *MockClientMockRecorder) StartInstances(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartInstances", reflect.TypeOf((*MockClient)(nil).StartInstances), arg0)
}

// RegisterInstancesWithLoadBalancer mocks base method
func (m *MockClient) RegisterInstancesWithLoadBalancer(arg0 *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterInstancesWithLoadBalancer", arg0)
	ret0, _ := ret[0].(*elb.RegisterInstancesWithLoadBalancerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterInstancesWithLoadBalancer indicates an expected call of RegisterInstancesWithLoadBalancer
func (mr *MockClientMockRecorder) RegisterInstancesWithLoadBalancer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInstancesWithLoadBalancer", reflect.TypeOf((*MockClient)(nil).RegisterInstancesWithLoadBalancer), arg0)
}

// CreateAccessKey mocks base method
func (m *MockClient) CreateAccessKey(arg0 *iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessKey", arg0)
	ret0, _ := ret[0].(*iam.CreateAccessKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessKey indicates an expected call of CreateAccessKey
func (mr *MockClientMockRecorder) CreateAccessKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessKey", reflect.TypeOf((*MockClient)(nil).CreateAccessKey), arg0)
}

// CreateUser mocks base method
func (m *MockClient) CreateUser(arg0 *iam.CreateUserInput) (*iam.CreateUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*iam.CreateUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockClientMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockClient)(nil).CreateUser), arg0)
}

// DeleteAccessKey mocks base method
func (m *MockClient) DeleteAccessKey(arg0 *iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessKey", arg0)
	ret0, _ := ret[0].(*iam.DeleteAccessKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccessKey indicates an expected call of DeleteAccessKey
func (mr *MockClientMockRecorder) DeleteAccessKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessKey", reflect.TypeOf((*MockClient)(nil).DeleteAccessKey), arg0)
}

// DeleteUser mocks base method
func (m *MockClient) DeleteUser(arg0 *iam.DeleteUserInput) (*iam.DeleteUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(*iam.DeleteUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockClientMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockClient)(nil).DeleteUser), arg0)
}

// DeleteUserPolicy mocks base method
func (m *MockClient) DeleteUserPolicy(arg0 *iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserPolicy", arg0)
	ret0, _ := ret[0].(*iam.DeleteUserPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserPolicy indicates an expected call of DeleteUserPolicy
func (mr *MockClientMockRecorder) DeleteUserPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserPolicy", reflect.TypeOf((*MockClient)(nil).DeleteUserPolicy), arg0)
}

// GetUser mocks base method
func (m *MockClient) GetUser(arg0 *iam.GetUserInput) (*iam.GetUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*iam.GetUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockClientMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockClient)(nil).GetUser), arg0)
}

// ListAccessKeys mocks base method
func (m *MockClient) ListAccessKeys(arg0 *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessKeys", arg0)
	ret0, _ := ret[0].(*iam.ListAccessKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessKeys indicates an expected call of ListAccessKeys
func (mr *MockClientMockRecorder) ListAccessKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessKeys", reflect.TypeOf((*MockClient)(nil).ListAccessKeys), arg0)
}

// ListUserPolicies mocks base method
func (m *MockClient) ListUserPolicies(arg0 *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserPolicies", arg0)
	ret0, _ := ret[0].(*iam.ListUserPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserPolicies indicates an expected call of ListUserPolicies
func (mr *MockClientMockRecorder) ListUserPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserPolicies", reflect.TypeOf((*MockClient)(nil).ListUserPolicies), arg0)
}

// PutUserPolicy mocks base method
func (m *MockClient) PutUserPolicy(arg0 *iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutUserPolicy", arg0)
	ret0, _ := ret[0].(*iam.PutUserPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutUserPolicy indicates an expected call of PutUserPolicy
func (mr *MockClientMockRecorder) PutUserPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUserPolicy", reflect.TypeOf((*MockClient)(nil).PutUserPolicy), arg0)
}

// CreateBucket mocks base method
func (m *MockClient) CreateBucket(arg0 *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", arg0)
	ret0, _ := ret[0].(*s3.CreateBucketOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBucket indicates an expected call of CreateBucket
func (mr *MockClientMockRecorder) CreateBucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockClient)(nil).CreateBucket), arg0)
}

// DeleteBucket mocks base method
func (m *MockClient) DeleteBucket(arg0 *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket", arg0)
	ret0, _ := ret[0].(*s3.DeleteBucketOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBucket indicates an expected call of DeleteBucket
func (mr *MockClientMockRecorder) DeleteBucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockClient)(nil).DeleteBucket), arg0)
}

// ListBuckets mocks base method
func (m *MockClient) ListBuckets(arg0 *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBuckets", arg0)
	ret0, _ := ret[0].(*s3.ListBucketsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuckets indicates an expected call of ListBuckets
func (mr *MockClientMockRecorder) ListBuckets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuckets", reflect.TypeOf((*MockClient)(nil).ListBuckets), arg0)
}

// GetS3API mocks base method
func (m *MockClient) GetS3API() s3iface.S3API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetS3API")
	ret0, _ := ret[0].(s3iface.S3API)
	return ret0
}

// GetS3API indicates an expected call of GetS3API
func (mr *MockClientMockRecorder) GetS3API() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetS3API", reflect.TypeOf((*MockClient)(nil).GetS3API))
}

// CreateHostedZone mocks base method
func (m *MockClient) CreateHostedZone(input *route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHostedZone", input)
	ret0, _ := ret[0].(*route53.CreateHostedZoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHostedZone indicates an expected call of CreateHostedZone
func (mr *MockClientMockRecorder) CreateHostedZone(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHostedZone", reflect.TypeOf((*MockClient)(nil).CreateHostedZone), input)
}

// GetHostedZone mocks base method
func (m *MockClient) GetHostedZone(arg0 *route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostedZone", arg0)
	ret0, _ := ret[0].(*route53.GetHostedZoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostedZone indicates an expected call of GetHostedZone
func (mr *MockClientMockRecorder) GetHostedZone(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostedZone", reflect.TypeOf((*MockClient)(nil).GetHostedZone), arg0)
}

// ListTagsForResource mocks base method
func (m *MockClient) ListTagsForResource(arg0 *route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForResource", arg0)
	ret0, _ := ret[0].(*route53.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource
func (mr *MockClientMockRecorder) ListTagsForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockClient)(nil).ListTagsForResource), arg0)
}

// ChangeTagsForResource mocks base method
func (m *MockClient) ChangeTagsForResource(input *route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeTagsForResource", input)
	ret0, _ := ret[0].(*route53.ChangeTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeTagsForResource indicates an expected call of ChangeTagsForResource
func (mr *MockClientMockRecorder) ChangeTagsForResource(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeTagsForResource", reflect.TypeOf((*MockClient)(nil).ChangeTagsForResource), input)
}

// DeleteHostedZone mocks base method
func (m *MockClient) DeleteHostedZone(input *route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHostedZone", input)
	ret0, _ := ret[0].(*route53.DeleteHostedZoneOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteHostedZone indicates an expected call of DeleteHostedZone
func (mr *MockClientMockRecorder) DeleteHostedZone(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHostedZone", reflect.TypeOf((*MockClient)(nil).DeleteHostedZone), input)
}

// ListHostedZones mocks base method
func (m *MockClient) ListHostedZones(input *route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHostedZones", input)
	ret0, _ := ret[0].(*route53.ListHostedZonesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostedZones indicates an expected call of ListHostedZones
func (mr *MockClientMockRecorder) ListHostedZones(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostedZones", reflect.TypeOf((*MockClient)(nil).ListHostedZones), input)
}

// ListResourceRecordSets mocks base method
func (m *MockClient) ListResourceRecordSets(input *route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceRecordSets", input)
	ret0, _ := ret[0].(*route53.ListResourceRecordSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceRecordSets indicates an expected call of ListResourceRecordSets
func (mr *MockClientMockRecorder) ListResourceRecordSets(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceRecordSets", reflect.TypeOf((*MockClient)(nil).ListResourceRecordSets), input)
}

// ListHostedZonesByName mocks base method
func (m *MockClient) ListHostedZonesByName(input *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHostedZonesByName", input)
	ret0, _ := ret[0].(*route53.ListHostedZonesByNameOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostedZonesByName indicates an expected call of ListHostedZonesByName
func (mr *MockClientMockRecorder) ListHostedZonesByName(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostedZonesByName", reflect.TypeOf((*MockClient)(nil).ListHostedZonesByName), input)
}

// ChangeResourceRecordSets mocks base method
func (m *MockClient) ChangeResourceRecordSets(arg0 *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeResourceRecordSets", arg0)
	ret0, _ := ret[0].(*route53.ChangeResourceRecordSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeResourceRecordSets indicates an expected call of ChangeResourceRecordSets
func (mr *MockClientMockRecorder) ChangeResourceRecordSets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeResourceRecordSets", reflect.TypeOf((*MockClient)(nil).ChangeResourceRecordSets), arg0)
}

// GetResourcesPages mocks base method
func (m *MockClient) GetResourcesPages(input *resourcegroupstaggingapi.GetResourcesInput, fn func(*resourcegroupstaggingapi.GetResourcesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourcesPages", input, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetResourcesPages indicates an expected call of GetResourcesPages
func (mr *MockClientMockRecorder) GetResourcesPages(input, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcesPages", reflect.TypeOf((*MockClient)(nil).GetResourcesPages), input, fn)
}
