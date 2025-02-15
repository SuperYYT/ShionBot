package InformationProcessing

import (
	"strings"

	"xyz.nyan/ShionBot/src/MediaWikiAPI"
	"xyz.nyan/ShionBot/src/MessagePushAPI"
	"xyz.nyan/ShionBot/src/Plugin/GetWikiInfo"
	"xyz.nyan/ShionBot/src/Plugin/Command"
	"xyz.nyan/ShionBot/src/Struct"
)

var sns_name_line string = "Line"

func LineMessageProcessing(json Struct.WebHookJson) {
	text := json.Events[0].Message.Text
	find, QueryText, Command := CommandExtraction(sns_name_line, json, text)
	if find {
		UserID := json.Events[0].Source.UserId
		ChatType := json.Events[0].Source.Type
		Log(sns_name_line, ChatType, UserID, text)
		switch ChatType {
		case "user":
			WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_line, json, UserID, Command, QueryText, "")
			if err != nil {
				WikiLink := MediaWikiAPI.GetWikiLink(sns_name_line, json, Command)
				MessagePushAPI.SendMessage(sns_name_line, "Default", UserID, UserID, Error(sns_name_line, UserID, WikiLink), false, "", "", 0)
				return
			}
			MessagePushAPI.SendMessage(sns_name_line, "Default", UserID, UserID, WikiInfo, false, "", "", 0)
		case "group":
			GroupId := json.Events[0].Source.GroupId
			QuoteID := json.Events[0].ReplyToken
			WikiInfo, err := GetWikiInfo.GetWikiInfo(sns_name_line, json, UserID, Command, QueryText, "")
			if err != nil {
				WikiLink := MediaWikiAPI.GetWikiLink(sns_name_line, json, Command)
				MessagePushAPI.SendMessage(sns_name_line, "Group", UserID, GroupId, Error(sns_name_line, UserID, WikiLink), true, QuoteID, "", 0)
				return
			}
			MessagePushAPI.SendMessage(sns_name_line, "Group", UserID, GroupId, WikiInfo, true, QuoteID, "", 0)
		}
	}
}

//设置消息返回
func LineSettingsMessageProcessing(json Struct.WebHookJson) {
	text := json.Events[0].Message.Text
	countSplit := strings.SplitN(text, "/", 2)
	Text := countSplit[1]
	Message, Bool := Command.Command(sns_name_line, json, Text)
	if Bool {
		ChatType := json.Events[0].Source.Type
		UserID := json.Events[0].Source.UserId
		Log(sns_name_line, ChatType, UserID, text)
		switch ChatType {
		case "user":
			MessagePushAPI.SendMessage(sns_name_line, "Default", UserID, UserID, Message, false, "", "", 0)
		case "group":
			GroupId := json.Events[0].Source.GroupId
			QuoteID := json.Events[0].ReplyToken
			MessagePushAPI.SendMessage(sns_name_line, "Group", UserID, GroupId, Message, true, QuoteID, "", 0)
		}
	}
}
