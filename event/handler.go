package event

import (
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"

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
		Name       string
		Gender     string
		Position   string
		Species    string
		Ressources string
		Range      string
		Region     string
	}
	player struct {
		Champions map[string]cham
	}
)

type Car struct {
	Name       string
	Gender     string
	Position   string
	Species    string
	Ressources string
	Range      string
	Region     string
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
	t := new(Car)

	for i := 0; i < (len(V.Champions)); i++ {
		o := strconv.Itoa(i)

		t.Name = V.Champions[""+o+""].Name
		t.Gender = V.Champions[""+o+""].Gender
		t.Position = V.Champions[""+o+""].Position
		t.Range = V.Champions[""+o+""].Range
		t.Ressources = V.Champions[""+o+""].Ressources
		t.Species = V.Champions[""+o+""].Species
		t.Region = V.Champions[""+o+""].Region
		// println(t.Name)

		Champ = append(Champ, &Car{
			Name:       t.Name,
			Gender:     t.Gender,
			Position:   t.Position,
			Species:    t.Species,
			Ressources: t.Ressources,
			Range:      t.Range,
			Region:     t.Region,
		})
	}
	println("List Loaded")
}

func Message(s *discordgo.Session, message *discordgo.MessageCreate) {
	println(len(Champ))

	if message.Author.ID == s.State.User.ID {
		return
	}

	switch strings.Split(message.Content, " ")[0] {
	case "!ping":
		s.ChannelMessageSend(message.ChannelID, "pong!")
	case "!guess":
		if !Discovered {
			discov(s, message)
			return
		}
		if Discovered {
			Discovered = false

			rand.Seed(time.Now().UTC().UnixNano())
			RandomIntegerwithinRange := rand.Intn(len(Champ))

			integer := strconv.Itoa(RandomIntegerwithinRange)

			// println(integer)
			// println(V.Champions[""+integer+""].Name)

			sauce := new(Car)
			sauce.Name = V.Champions[""+integer+""].Name
			sauce.Gender = V.Champions[""+integer+""].Gender
			sauce.Position = V.Champions[""+integer+""].Position
			sauce.Range = V.Champions[""+integer+""].Range
			sauce.Ressources = V.Champions[""+integer+""].Ressources
			sauce.Species = V.Champions[""+integer+""].Species
			sauce.Region = V.Champions[""+integer+""].Region

			Need[0] = sauce

			println(Need[0].Name + "NEW §!!")

			discov(s, message)
		}

	}

}

func isValid(c bool) string {
	if !c {
		return " ❌"
	}
	return " ✅"
}

func discov(s *discordgo.Session, message *discordgo.MessageCreate) {
	for i := range Champ {
		var Gender string
		var Position string
		var Species string
		var Ressource string
		var Range string
		var Region string
		// println(c.Name)
		// println(Champ[i].Name)
		//var msg []string
		//var c string

		msg := strings.ToUpper(string(strings.Split(message.Content, "!guess ")[1]))

		if msg == strings.ToUpper(Champ[i].Name) {
			if strings.ToUpper(Need[0].Name) == msg {
				s.ChannelMessageSend(message.ChannelID, "You found the champion, it was "+Need[0].Name)
				Discovered = true
			} else {
				println("Passed")
				Gender = isValid(Need[0].Gender != Champ[i].Gender) 
				Position = isValid(Need[0].Position != Champ[i].Position)
				Species = isValid(Need[0].Species != Champ[i].Species)
				Ressource = isValid(Need[0].Ressources != Champ[i].Ressources)
				Range = isValid(Need[0].Range != Champ[i].Range)
				Region = isValid(Need[0].Region != Champ[i].Region)

				s.ChannelMessageSendEmbeds(message.ChannelID, []*discordgo.MessageEmbed{
					{
						Title: Champ[i].Name,
						Color: 16711680,
						Fields: []*discordgo.MessageEmbedField{
							{
								Name: Champ[i].Gender + Gender,
							},
							{
								Name: Champ[i].Position + Position,
							},
							{
								Name: Champ[i].Species + Species,
							},
							{
								Name: Champ[i].Ressources + Ressource,
							},
							{
								Name: Champ[i].Range + Range,
							},
							{
								Name: Champ[i].Region + Region,
							},
						},
					},
				})
			}
		}
	}
}
