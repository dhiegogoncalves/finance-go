package db

import (
	"context"
	"testing"

	"github.com/dhiegogoncalves/gofinance/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	user := createRandomUser(t)
	arg := CreateCategoryParams{
		UserID:      user.ID,
		Title:       util.RandomString(12),
		Type:        "debit",
		Description: util.RandomString(20),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.UserID, category.UserID)
	require.Equal(t, arg.Title, category.Title)
	require.Equal(t, arg.Type, category.Type)
	require.Equal(t, arg.Description, category.Description)
	require.NotEmpty(t, category.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategoryById(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategoryById(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.UserID, category2.UserID)
	require.Equal(t, category1.Title, category2.Title)
	require.Equal(t, category1.Type, category2.Type)
	require.Equal(t, category1.Description, category2.Description)
	require.NotEmpty(t, category2.CreatedAt)
}

func TestListCategories(t *testing.T) {
	category := createRandomCategory(t)

	arg := GetCategoriesParams{
		UserID:      category.UserID,
		Type:        category.Type,
		Title:       category.Title,
		Description: category.Description,
	}

	categories, err := testQueries.GetCategories(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, categories)

	for _, c := range categories {
		require.Equal(t, c.ID, category.ID)
		require.Equal(t, c.UserID, category.UserID)
		require.Equal(t, c.Title, category.Title)
		require.Equal(t, c.Type, category.Type)
		require.Equal(t, c.Description, category.Description)
		require.NotEmpty(t, c.CreatedAt)
	}

}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)

	arg := UpdateCategoryParams{
		ID:          category1.ID,
		Title:       util.RandomString(12),
		Description: util.RandomString(20),
	}

	category2, err := testQueries.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.UserID, category2.UserID)
	require.Equal(t, arg.Title, category2.Title)
	require.Equal(t, category1.Type, category2.Type)
	require.Equal(t, arg.Description, category2.Description)
	require.NotEmpty(t, category2.CreatedAt)
}

func TestDeleteCategoryById(t *testing.T) {
	category := createRandomCategory(t)

	err := testQueries.DeleteCategoryById(context.Background(), category.ID)
	require.NoError(t, err)

	_, err = testQueries.GetCategoryById(context.Background(), category.ID)
	require.Error(t, err)
}
