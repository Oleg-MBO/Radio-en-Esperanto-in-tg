package tgApiHelper

import (
	"github.com/Oleg-MBO/Radio-en-Esperanto/botdb"
	"github.com/Oleg-MBO/Radio-en-Esperanto/tgApiHelper/callbackQuery"
	"github.com/Oleg-MBO/Radio-en-Esperanto/tgApiHelper/message"
	"github.com/azer/logger"
	"github.com/olebedev/go-tgbot"
	"time"
)

var logTg = logger.New("tgApiHelper")

var Tgapi *tgbot.Router
var ChannelId string

const MaxLengthFile = 50000000 //limit file in tg api

var TimeDelayPodcastParse = 10 * time.Minute

func InitApi(token string, channelId string) *tgbot.Router {
	Tgapi = tgbot.New(&tgbot.Options{
		Context: nil,
		Token:   token,
	})
	Message.Tgapi = Tgapi
	callbackQuery.Tgapi = Tgapi

	ChannelId = channelId

	//Bind handlers
	Tgapi.Message("^/start", Message.StartMessage)
	Tgapi.Message("^/ek", Message.StartMessage)
	Tgapi.Message("^/help", Message.StartMessage)
	Tgapi.Message("^/helpo", Message.StartMessage)
	Tgapi.Message("^/add", Message.AddListPodkasts)
	Tgapi.Message("^/aboni", Message.AddListPodkasts)
	Tgapi.Message("^/rm", Message.RmListPodkasts)
	Tgapi.Message("^/malaboni", Message.RmListPodkasts)
	Tgapi.Message("^/list", Message.ListAllPodkasts)
	Tgapi.Message("^/listo", Message.ListAllPodkasts)
	Tgapi.Message("^.*", Message.AllText)

	Tgapi.CallbackQuery(`^add_.+$`, callbackQuery.AddPodkastId)
	Tgapi.CallbackQuery("^rm_.+$", callbackQuery.RmPodkastId)
	Tgapi.CallbackQuery("^0$", callbackQuery.UpdateKB)

	return Tgapi
}

func PodkastsProcessing() {
	defer func() {
		if r := recover(); r != nil {
			logTg.Error("recovered PodkastsProcessing: %#v", r)
			go PodkastsProcessing()
		}
	}()
	for {
		logTg.Info("start parse podcasts")
		for _, p := range botdb.GetNewPodcasts() {
			logTg.Info("parse ")
			if !p.IsUnique() {
				logTg.Info("is not unique")
				continue
			}
			logTg.Info("try parsed to send")
			if cs := SendPodcastToChannelAndSave(&p); cs {
				logTg.Info("SendPodcastToChannelAndSave OK")
				ForwardToAllChatsPodcast(p)
			}
			time.Sleep(time.Second * 3)
		}
		logTg.Info("finish parse podcasts, sleep")
		time.Sleep(TimeDelayPodcastParse)
	}
}
