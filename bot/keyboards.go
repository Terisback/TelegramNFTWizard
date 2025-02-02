package bot

import tele "gopkg.in/telebot.v3"

// Keyboards
var (
	menu      = &tele.ReplyMarkup{}
	btnCreate = menu.Data("Create collection", "create_collection")

	completeFiles    = &tele.ReplyMarkup{}
	btnCompleteFiles = completeFiles.Data("That's all files", "complete_files")

	skip    = &tele.ReplyMarkup{}
	btnSkip = skip.Data("Skip", "skip")

	minted    = &tele.ReplyMarkup{}
	btnMinted = skip.Data("Check mint status", "check_status")
)

// Keyboards init
func init() {
	menu.Inline(
		menu.Row(btnCreate),
	)

	completeFiles.Inline(
		completeFiles.Row(btnCompleteFiles),
	)

	skip.Inline(
		skip.Row(btnSkip),
	)

	minted.Inline(
		minted.Row(btnMinted),
	)
}
