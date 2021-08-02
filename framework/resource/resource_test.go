package resource

import (
	"reflect"
	"testing"

	"github.com/hashicorp/waypoint-plugin-sdk/internal/testproto"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	sdkpb "github.com/hashicorp/waypoint-plugin-sdk/proto/gen"
)

func TestResourceCreate_state(t *testing.T) {
	require := require.New(t)

	var destroyVal int
	r := NewResource(
		WithName("test"),
		WithState(&testState{}),
		WithCreate(func(state *testState, v int) error {
			state.Value = v
			return nil
		}),
		WithDestroy(func(state *testState) error {
			destroyVal = state.Value
			return nil
		}),
	)

	// Create
	require.NoError(r.Create(int(42)))

	// Ensure we were called with the proper value
	state := r.State().(*testState)
	require.NotNil(state)
	require.Equal(state.Value, 42)

	// Destroy
	require.NoError(r.Destroy())
	require.Equal(destroyVal, 42)
	require.Nil(r.State())
	require.Nil(r.State().(*testState))
}

func TestResourceCreate_stateNoDestroy(t *testing.T) {
	require := require.New(t)

	r := NewResource(
		WithName("test"),
		WithState(&testState{}),
		WithCreate(func(state *testState, v int) error {
			state.Value = v
			return nil
		}),
	)

	// Create
	require.NoError(r.Create(int(42)))

	// Ensure we were called with the proper value
	state := r.State().(*testState)
	require.NotNil(state)
	require.Equal(state.Value, 42)

	// Destroy
	require.NoError(r.Destroy())
	require.Nil(r.State())
	require.Nil(r.State().(*testState))
}

func TestResourceCreate_noStateNoDestroy(t *testing.T) {
	require := require.New(t)

	r := NewResource(
		WithName("test"),
		WithCreate(func(v int) error {
			return nil
		}),
	)

	// Create
	require.NoError(r.Create(int(42)))

	// Destroy
	require.NoError(r.Destroy())
	require.Nil(r.State())
}

func TestResourceCreate_noState(t *testing.T) {
	require := require.New(t)

	var actual int
	var destroyCalled bool
	r := NewResource(
		WithName("test"),
		WithCreate(func(v int) error {
			actual = v
			return nil
		}),
		WithDestroy(func() error {
			destroyCalled = true
			return nil
		}),
	)

	// Create
	require.NoError(r.Create(int(42)))

	// Ensure we were called with the proper value
	require.Equal(actual, int(42))

	// Destroy
	require.NoError(r.Destroy())
	require.True(destroyCalled)
}

type testState struct {
	Value int
}

type testState2 testState

func TestResource_DeclaredResource(t *testing.T) {
	require := require.New(t)

	testResource := &Resource{
		name:      "test resource",
		stateType: reflect.TypeOf(&testproto.Data{}),
		stateValue: &testproto.Data{
			Value:  "val",
			Number: 1,
		},
		createFunc:          func() {},
		destroyFunc:         func() {},
		platform:            "test-platform",
		categoryDisplayHint: sdkpb.ResourceCategoryDisplayHint_OTHER,
	}

	dr, err := testResource.DeclaredResource()
	require.Nil(err)

	require.Equal(dr.Name, testResource.name)
	require.Equal(dr.Platform, testResource.platform)
	require.Equal(dr.CategoryDisplayHint, testResource.categoryDisplayHint)
	require.NotEmpty(dr.StateJson)
	require.True(dr.State.MessageIs(testResource.State().(proto.Message)))
}