#!/bin/bash

export GOOS=linux
go build Server.go

ssh e1on@5.180.138.37 "sudo systemctl stop game-server"
scp ./Server e1on@5.180.138.37:/home/e1on/Game
ssh e1on@5.180.138.37 "chmod +x /home/e1on/Game"
ssh e1on@5.180.138.37 "sudo systemctl start game-server"