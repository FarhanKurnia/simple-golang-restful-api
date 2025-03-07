package helper

import (
	"FarhanKurnia/simple-golang-restful-api/model/domain"
	"FarhanKurnia/simple-golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Name: category.Name,
		Uuid: category.Uuid,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
