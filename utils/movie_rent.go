package utils

import "github/jorgemvv01/go-api/models"

func CalculateTotalRent(movies []models.MovieSummary, days int) float64 {
	var total float64
	for _, movie := range movies {
		var moviePrice = movie.Price
		var totalMoviePrice float64
		totalMoviePrice = (moviePrice) * float64(days)
		switch movie.TypeID {
		case 2:
			if days > 3 {
				totalMoviePrice = (moviePrice) * (3)
				totalMoviePrice += (moviePrice + (moviePrice * 0.15)) * float64(days-3)
			}
		case 3:
			if days > 5 {
				totalMoviePrice = (moviePrice) * (5)
				totalMoviePrice += (moviePrice + (moviePrice * 0.10)) * float64(days-5)
			}
		}
		total += totalMoviePrice
	}
	return total
}
