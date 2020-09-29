package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var factoryFunctions = make(map[string]commandHandlerFactory)

func Init() {
	factoryFunctions["help"] = newHelpFactory("help")
	factoryFunctions["emails"] = newEmailsFactory("emails")
	factoryFunctions["grupos"] = newGruposFactory("grupos")
	factoryFunctions["readme"] = newReadmeFactory("readme")
	factoryFunctions["email"] = newEmailFactory("email")
	factoryFunctions["temas"] = newTemasFactory("temas")
}

func NewCommandHandler(bot *tgbotapi.BotAPI, message *tgbotapi.Message) (CommandHandler, error) {
	commandName := message.Command()
	factory, ok := factoryFunctions[commandName]
	if !ok {
		return nil, fmt.Errorf("Command name \"%s\" is invalid, or command was not implemented yet.", commandName)
	}
	return factory(bot, message)
}
