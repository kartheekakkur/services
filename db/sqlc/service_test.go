package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateService(t *testing.T) {
	arg := CreateServiceParams{
		Name:        "konnect10",
		Description: "I am Konnect10",
		Versions:    "v1.1",
	}
	service, err := testQueries.CreateService(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, service)

	require.Equal(t, arg.Name, service.Name)
	require.Equal(t, arg.Description, service.Description)
	require.Equal(t, arg.Versions, service.Versions)

	require.NotZero(t, service.ID)
	require.NotZero(t, service.CreatedAt)

}

//similarly we run test cases for get,list,delete and update operations
