package main

import (
	"os"
	"os/signal"
	"syscall"

	"fr.akika.lodlebot/event"
	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Args[1])
	if err != nil {
		panic(err)
	}

	err = discord.Open()
	if err != nil {
		println(err.Error())
	}
	println("Bot was succefully started")
	event.Discovered = true

	event.Listener(discord)
	event.Init()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}
