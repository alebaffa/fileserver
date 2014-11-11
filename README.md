#HelpTheSheriff webapp
Small and simple web application written in Go. Martini manages the routing, the frontend is made with Bootstrap and there's a MongoDB database behind the scene. It is deployed on Heroku.
https://helpthesheriff.herokuapp.com/.

I am sorry for this crappy name (fileserver), but it was a different app at the beginning, and I was too lazy to create a new one.

## Prerequisites
If you want to modify it, these are the technical requirements:
* [Go latest version](https://golang.org/doc/install): the backend is written in Go. I think it's compatible also with older versions, but there's no reason to not use the latest one.
* [Heroku toolbelt](https://toolbelt.heroku.com/): to push to heroku server using Git.
* [Account on Heroku](http://heroku.com/): to create a new app.

## Install locally
* [install and configure Go](http://golang.org/doc/install) and the [Heroku toolbelt](https://toolbelt.heroku.com/) on your machine
* create an account on Heroku
* fork and clone this repository on your local machine

## Run locally
In order to run it locally you just need to have Go installed and configured ($GOPATH!!).
* *cd /fileserver*
* *go build*
* *./fileserver* : at this point you should be able to run this webapp locally at [http://localhost:8080](http://localhost:3000).

## Deploy on Heroku
* follow the official procedure [here](https://toolbelt.heroku.com/). It's easy as hell.