package event

import (
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/bwmarrin/discordgo"
)

func Listener(s *discordgo.Session) {
	s.AddHandler(Message)
}

var Champ []*Car
var Need = map[int]*Car{}
var Discovered bool
var V player

type (
	cham struct {
		Name      string
		Gender    string
		Position  string
		Species   string
		Ressource string
		Range     string
	}
	player struct {
		Champions map[string]cham
	}
)

type Car struct {
	Name      string
	Gender    string
	Position  string
	Species   string
	Ressource string
	Range     string
}

func Init() {
	file, err := ioutil.ReadFile("./ressources/champs.toml")
	if err != nil {
		println(err.Error())
	}

	err = toml.Unmarshal(file, &V)
	if err != nil {
		println(err.Error())
	}

	println(len(V.Champions))
	b := 0
	t := new(Car)

	for i := 0; i < (len(V.Champions)); i++ {
		o := strconv.Itoa(i)

		t.Name = V.Champions[""+o+""].Name
		t.Gender = V.Champions[""+o+""].Gender
		t.Position = V.Champions[""+o+""].Position
		t.Range = V.Champions[""+o+""].Range
		t.Ressource = V.Champions[""+o+""].Ressource
		t.Species = V.Champions[""+o+""].Species
		// println(t.Name)

		Champ = append(Champ, &Car{
			Name:      t.Name,
			Gender:    t.Gender,
			Position:  t.Position,
			Species:   t.Species,
			Ressource: t.Ressource,
			Range:     t.Range,
		})
		print(Champ[b].Name + "\n")
		b = +1
	}
}

func Message(s *discordgo.Session, message *discordgo.MessageCreate) {
	println(len(Champ))

	if message.Author.ID == s.State.User.ID {
		return
	}

	switch strings.Split(message.Content, " ")[0] {
	case "!ping":
		s.ChannelMessageSend(message.ChannelID, "pong!")
	case "test":
		for i := range Champ {
			println(Champ[i].Name)
		}
		if !Discovered {
			for i := range Champ {
				var Gender string
				var Position string
				var Species string
				var Ressource string
				var Range string
				// println(c.Name)
				// println(Champ[i].Name)
				if strings.Split(message.Content, " ")[1] == Champ[i].Name {
					if Need[0].Name == strings.Split(message.Content, " ")[1] {
						s.ChannelMessageSend(message.ChannelID, "You found the champion, it was "+Need[0].Name)
						Discovered = true
					} else {
						println("Passed")
						if Need[0].Gender != Champ[i].Gender {
							Gender = "false"
						} else {
							Gender = "True"
						}

						if Need[0].Position != Champ[i].Position {
							Position = "false"
						} else {
							Position = "True"
						}

						if Need[0].Species != Champ[i].Species {
							Species = "false"
						} else {
							Species = "True"
						}

						if Need[0].Ressource != Champ[i].Ressource {
							Ressource = "false"
						} else {
							Ressource = "True"
						}

						if Need[0].Range != Champ[i].Range {
							Range = "false"
						} else {
							Range = "True"
						}
						s.ChannelMessageSend(message.ChannelID, ("Name: " + Champ[i].Name +
							"false \nGender: " + Champ[i].Gender + " " + Gender +
							"\nPosition: " + Champ[i].Position + " " + Position +
							"\nSpecies: " + Champ[i].Species + " " + Species +
							"\nRessource: " + Champ[i].Ressource + " " + Ressource +
							"\nRange: " + Champ[i].Range + " " + Range))
					}
				}
			}
			println(Need[0].Name)
		}
		if Discovered {
			Discovered = false
			RandomIntegerwithinRange := rand.Intn(len(V.Champions))

			integer := strconv.Itoa(RandomIntegerwithinRange)

			// println(integer)
			// println(V.Champions[""+integer+""].Name)

			sauce := new(Car)
			sauce.Name = V.Champions[""+integer+""].Name
			sauce.Gender = V.Champions[""+integer+""].Gender
			sauce.Position = V.Champions[""+integer+""].Position
			sauce.Range = V.Champions[""+integer+""].Range
			sauce.Ressource = V.Champions[""+integer+""].Ressource
			sauce.Species = V.Champions[""+integer+""].Species

			Need[0] = sauce

			println(Need[0].Name + "NEW §!!")
		}

	}

}
