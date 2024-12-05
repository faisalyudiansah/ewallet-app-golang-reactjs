# assignment-e-wallet-rest-api

## note
- create database on your local
- set up your env in .env file. You can see example env in .env.examples file
- command `go mod tidy` if necessary

## command app
- `run` : to start app without hot reload
- `start-nodemon`: to start app with
- `test-cover`: to testing the app
- `create-mock`: to generate mockery
- `db-generate`: to migrate and seed data (but first, make your own database on your local)

## assumptions game
- If you want to start the game, the user must try the opportunity to play first by uploading Rp. 10000000. 
- if there is, then visit = /game/start (login required)
- After that a list box will appear The user must choose 1 box from the 9 boxes. 
- by sending body box_index to url /game/choose

## API DOCS
https://documenter.getpostman.com/view/29632965/2sA3e5c7eX