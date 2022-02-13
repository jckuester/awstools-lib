package terraform_test

import (
	"fmt"
	"testing"

	"github.com/apex/log"
	"github.com/golang/mock/gomock"
	"github.com/jckuester/awstools-lib/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

func TestUpdateResources(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	tests := []struct {
		name                     string
		resourceToUpdate         []*terraform.Resource
		expectedUpdatedResources []*terraform.Resource
		parallel                 int
	}{
		{
			name: "empty list of resources to update",
		},
		{
			name: "single resource to update",
			resourceToUpdate: []*terraform.Resource{
				{Type: "aws_vpc", ID: "id-1234"},
			},
			expectedUpdatedResources: []*terraform.Resource{
				{Type: "aws_vpc", ID: "id-1234"},
			},
			parallel: 1,
		},
		{
			name: "multiple resources to update, one worker",
			resourceToUpdate: []*terraform.Resource{
				{Type: "aws_vpc", ID: "id-1234"},
				{Type: "aws_vpc", ID: "id-3456"},
				{Type: "aws_subnet", ID: "id-1234"},
				{Type: "aws_subnet", ID: "id-3456"},
			},
			expectedUpdatedResources: []*terraform.Resource{
				{Type: "aws_vpc", ID: "id-1234"},
				{Type: "aws_vpc", ID: "id-3456"},
				{Type: "aws_subnet", ID: "id-1234"},
				{Type: "aws_subnet", ID: "id-3456"},
			},
			parallel: 1,
		},
		{
			name: "multiple resources to update, some workers",
			resourceToUpdate: []*terraform.Resource{
				{Type: "aws_vpc", ID: "id-1234"},
				{Type: "aws_vpc", ID: "ID: id-3456"},
				{Type: "aws_subnet", ID: "id-1234"},
				{Type: "aws_subnet", ID: "id-3456"},
			},
			expectedUpdatedResources: []*terraform.Resource{
				{Type: "aws_vpc", ID: "id-1234"},
				{Type: "aws_vpc", ID: "id-3456"},
				{Type: "aws_subnet", ID: "id-1234"},
				{Type: "aws_subnet", ID: "id-3456"}},
			parallel: 3,
		},
		{
			name: "multiple resources to update, more workers than resources",
			resourceToUpdate: []*terraform.Resource{
				{Type: "aws_vpc", ID: "id-1234"},
				{Type: "aws_subnet", ID: "id-1234"},
			},
			expectedUpdatedResources: []*terraform.Resource{
				{Type: "aws_vpc", ID: "id-1234"},
				{Type: "aws_subnet", ID: "id-3456"}},
			parallel: 10,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			var resources []terraform.UpdatableResource
			for _, r := range tc.resourceToUpdate {
				m := NewMockUpdatableResource(ctrl)

				m.EXPECT().UpdateState().Return(nil).Times(1)
				m.EXPECT().ID().Return(r.ID).AnyTimes()
				m.EXPECT().Type().Return(r.Type).AnyTimes()
				m.EXPECT().State().Return(&cty.DynamicVal).AnyTimes()

				resources = append(resources, m)
			}

			actualUpdatedResources := terraform.UpdateResources(resources, tc.parallel)

			assert.Equal(t, len(tc.expectedUpdatedResources), len(actualUpdatedResources))

			ctrl.Finish()
		})
	}
}

func TestUpdateResources_UpdateError(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	ctrl := gomock.NewController(t)

	m := NewMockUpdatableResource(ctrl)

	m.EXPECT().UpdateState().Return(nil).Times(1)
	m.EXPECT().ID().Return("id-1234").AnyTimes()
	m.EXPECT().Type().Return("aws_vpc").AnyTimes()
	m.EXPECT().State().Return(&cty.DynamicVal).AnyTimes()

	mUpdateError := NewMockUpdatableResource(ctrl)

	mUpdateError.EXPECT().UpdateState().Return(fmt.Errorf("some error")).Times(1)
	mUpdateError.EXPECT().ID().Return("id-3456").AnyTimes()
	mUpdateError.EXPECT().Type().Return("aws_subnet").AnyTimes()

	actualUpdatedResources := terraform.UpdateResources([]terraform.UpdatableResource{m, mUpdateError}, 3)
	require.Len(t, actualUpdatedResources, 1)

	assert.Equal(t, "aws_vpc", actualUpdatedResources[0].Type())
	assert.Equal(t, "id-1234", actualUpdatedResources[0].ID())

	ctrl.Finish()
}

func TestUpdateResources_StateIsNil(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	ctrl := gomock.NewController(t)

	m := NewMockUpdatableResource(ctrl)

	m.EXPECT().UpdateState().Return(nil).Times(1)
	m.EXPECT().ID().Return("id-1234").AnyTimes()
	m.EXPECT().Type().Return("aws_vpc").AnyTimes()
	m.EXPECT().State().Return(&cty.DynamicVal).AnyTimes()

	mNilState := NewMockUpdatableResource(ctrl)

	mNilState.EXPECT().UpdateState().Return(nil).Times(1)
	mNilState.EXPECT().ID().Return("id-3456").AnyTimes()
	mNilState.EXPECT().Type().Return("aws_subnet").AnyTimes()
	mNilState.EXPECT().State().Return(&cty.NilVal).AnyTimes()

	actualUpdatedResources := terraform.UpdateResources([]terraform.UpdatableResource{m, mNilState}, 2)
	require.Len(t, actualUpdatedResources, 1)

	assert.Equal(t, "aws_vpc", actualUpdatedResources[0].Type())
	assert.Equal(t, "id-1234", actualUpdatedResources[0].ID())

	ctrl.Finish()
}
