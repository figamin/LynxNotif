package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	//"log"
	"net/http"
	"os"
	"time"
	//"github.com/bwmarrin/discordgo"
)

type Config struct {
    Website            string `json:"website"`
	DiscordToken       string `json:"discordToken"`
	QuickReply         bool   `json:"quickReply"`
	UpdateIntervalMins int    `json:"updateIntervalMins"`
	OverboardMode      bool   `json:"overboardMode"`
}
type LynxJson struct {
	LatestPosts []struct {
		BoardURI    string `json:"boardUri"`
		ThreadID    int    `json:"threadId"`
		PostID      int    `json:"postId"`
		PreviewText string `json:"previewText"`
	} `json:"latestPosts"`
	LatestImages []struct {
		BoardURI string `json:"boardUri"`
		ThreadID int    `json:"threadId"`
		PostID   int    `json:"postId"`
		Thumb    string `json:"thumb"`
	} `json:"latestImages"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}
func LoadConfiguration(file string) Config {
    var config Config
    configFile, err := os.Open(file)
	if err != nil {
        fmt.Println(err.Error())
    }
    defer configFile.Close()
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config
}
// d
func NewPostCheckOverboard(Website string, QuickReply bool, previousBoard *string, previousPost *int) string {
	var newPosts LynxJson
	r, err := myClient.Get(Website + "/index.json")
    if err != nil {
        fmt.Println(err)
    }
    defer r.Body.Close()
    err = json.NewDecoder(r.Body).Decode(&newPosts)
	if err != nil {
        fmt.Println(err)
    }
	if newPosts.LatestPosts[0].PostID != *previousPost || newPosts.LatestPosts[0].BoardURI != *previousBoard {
		*previousPost = newPosts.LatestPosts[0].PostID
		*previousBoard = newPosts.LatestPosts[0].BoardURI
		newPostTemplate := "New /" + newPosts.LatestPosts[0].BoardURI + "/ Post: " + Website + "/" + newPosts.LatestPosts[0].BoardURI + "/res/" + fmt.Sprintf("%v", newPosts.LatestPosts[0].PostID) + ".html#"
		if QuickReply {
			return newPostTemplate + "q" + fmt.Sprintf("%v", newPosts.LatestPosts[0].PostID)
		} else {
			return newPostTemplate + fmt.Sprintf("%v", newPosts.LatestPosts[0].PostID)
		}
	} else {
		return "NONE"
	}
	if newPosts.LatestImages[len(newPosts.LatestImages) - 1£]
}
func main() {
	fmt.Println("Bot Invite Link: " + "https://discord.com/api/oauth2/authorize?client_id=1048986212683223133&permissions=8&scope=bot")
	/*dg, err := discordgo.New("Bot " + config.)
    if err != nil {
        fmt.Println("error creating Discord session,", err)
        return
    }*/
	
	var configFile = LoadConfiguration("config.json")
	var previousPost = 0
	var previousBoard = ""
	var postReply string
	for range time.Tick(time.Second * 5) {
		go func() {
			if(configFile.OverboardMode) {
				postReply = NewPostCheckOverboard(configFile.Website, configFile.QuickReply, &previousBoard, &previousPost)
			} else {
				postReply = "Individual Board Mode TODO"
			}
			if(postReply != "NONE") {
				fmt.Println(postReply)
			}
		}()
	}
}