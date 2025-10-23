package domain

import "errors"

var AnimeNotFound=errors.New("Anime not found!.")
var AnimeEpisodeNotFound=errors.New("Anime Episode not found!.")
var AnimeGenreNotFound=errors.New("Anime Genre not found!.")
var AnimeGenresNotFound=errors.New("Anime Genres not found!.")
var AuthFail=errors.New("Unauthenticated, Email or Password is wrong.")
var EmailRegister=errors.New("Email already register!.")
var AnimeGenresAlready=errors.New("Anime Genres already in database!")
var AnimeTagNotFound=errors.New("Anime Tag not found!.")
var AnimeTagsAlready=errors.New("Anime Tags already in database!.")
var AnimeTagsNotFound=errors.New("Anime Tags not found!.")
var AnimeMediaNotFound=errors.New("Anime Media not found!.")
var AnimeMediaOutsideDir=errors.New("Invalid media path: outside of base dir")
var AnimeStudioNotFound=errors.New("Anime Studio not found!.")
var AnimeStudiosNotFound=errors.New("Anime Studios not found!.")
var AnimeStudiosAlready=errors.New("Anime Studios already in database!.")
var CurrentPasswordWrong=errors.New("Current Password is Wrong!.")
var PeopleNotFound=errors.New("People Not Found!")

