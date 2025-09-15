package domain

import "errors"

var AnimeNotFound=errors.New("Anime not found!.")
var AnimeEpisodeNotFound=errors.New("Anime episode not found!.")
var AnimeGenreNotFound=errors.New("Anime genre episode not found!.")
var AuthFail=errors.New("Authentication Failed!.")
var EmailRegister=errors.New("Email already register!.")
var PasswordNotMatch=errors.New("Password and Confirm Password not matching!.")
