package InformationProcessing

import (
	"strings"

	"xyz.nyan/ShionBot/src/Plugin"
	"xyz.nyan/ShionBot/src/Struct"
	"xyz.nyan/ShionBot/src/utils"
	"xyz.nyan/ShionBot/src/utils/Language"
)

//Wiki链接错误返回
func Error(SNSName string, UserID string, WikiLink string) string {
	return utils.StringVariable(Language.Message(SNSName, UserID).WikiLinkError, []string{WikiLink})
}

//命令处理，判断命令是否匹配，匹配则输出命令和命令参数
func CommandExtraction(SNSName string, json Struct.WebHookJson, text string) (bool, string, string) {
	if find := strings.Contains(text, ":"); find {
		if Plugin.WikiNameExist(text) {
			countSplit := strings.SplitN(text, ":", 2)
			Command := countSplit[0]
			Text := countSplit[1]
			return true, Text, Command
		}
	} else if find := strings.Contains(text, "[["); find {
		if find := strings.Contains(text, "]]"); find {
			//获取主Wiki名字
			MainWikiName := Plugin.GeiMainWikiName()

			trimStr := strings.Trim(text, "[")
			Text := strings.Trim(trimStr, "]")
			return true, Text, MainWikiName
		}
	} else if find := strings.Contains(text, "/"); find {
		switch SNSName {
		case "QQ":
			QQSettingsMessageProcessing(json)
		case "Telegram":
			TelegramSettingsMessageProcessing(json)
		case "Line":
			LineSettingsMessageProcessing(json)
		case "KaiHeiLa":
			KaiHeiLaSettingsMessageProcessing(json)
		}
		return false, "", "/"
	}

	return false, "", ""
}
