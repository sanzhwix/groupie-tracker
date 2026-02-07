groupie-tracker
Description

Groupie Tracker is a web application written in Go that consumes a public REST API and displays information about music artists and bands.
The project focuses on data retrieval, manipulation, and visualization using a client–server architecture.

The application fetches and combines data from multiple API endpoints to present artist details, concert locations, and dates through a user-friendly web interface.

Objectives

Consume and process data from a RESTful API

Manipulate and link related datasets

Build a reliable Go HTTP server

Render dynamic HTML pages

Implement client-triggered events (request–response)

Handle errors safely to prevent server crashes

API Overview

The provided API consists of four parts:

artists
Contains artist and band information:

Name

Image

Year of activity start

First album date

Members

locations
Contains last and/or upcoming concert locations

dates
Contains last and/or upcoming concert dates

relation
Links artists with their corresponding locations and dates

Features

Fetches and parses JSON data from all API endpoints

Combines related data into structured Go models

Displays artist information using multiple visual layouts:

Cards

Lists

Pages

Implements client–server interaction through user actions

Ensures application stability with proper error handling

Technologies Used

Language: Go

Backend: net/http, JSON decoding

Frontend: HTML templates

Architecture: Client–Server

Packages: Standard Go packages only

Usage
Run the application
go run .


Then open your browser and navigate to:

http://localhost:8080


(Port may vary depending on implementation)

Project Structure
.
├── main.go
├── handlers/
├── templates/
├── static/ (if applicable)
├── go.mod
└── README.md

Error Handling

The server is designed to never crash

All pages handle missing or invalid data gracefully

API and template errors are checked and handled properly
