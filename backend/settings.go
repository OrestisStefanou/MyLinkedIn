package main

//Media settings
const mediaDir = "/home/orestis/Desktop/GitHubRep/MyLinkedIn/backend/media"
const mediaURL = "http://localhost:8080/v1/LinkedIn/media/"

//Database settings
const databaseUser = "<Username>"
const userPassword = "<Password>"
const databaseName = "LinkedIn"

//Variable to communicate with database
var dbclient DBClient

var validImageExtensions = [...]string{".jpeg", ".jpg", ".png", ".gif"}
var validAttachedFileExtensions = [...]string{".jpeg", ".jpg", ".png", ".gif", ".mp4", ".mp3", ".pdf"}

//Admin directories
const adminDir = "/home/orestis/Desktop/GitHubRep/MyLinkedIn/backend/media/admin/"
