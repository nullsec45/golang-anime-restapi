package domain

import "errors"

var AnimeNotFound=errors.New("Anime not found!.")
var AnimeEpisodeNotFound=errors.New("Anime episode not found!.")
var AnimeGenreNotFound=errors.New("Anime genre not found!.")
var AnimeGenresNotFound=errors.New("Anime genres not found!.")
var AuthFail=errors.New("Authentication Failed!.")
var EmailRegister=errors.New("Email already register!.")
var PasswordNotMatch=errors.New("Password and Confirm Password not matching!.")
var AnimeGenresAlready=errors.New("Anime Genres already in database!")
