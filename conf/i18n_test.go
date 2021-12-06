package conf

import (
	"testing"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestIL(t *testing.T) {
	msgKey := "inbox.title"

	message.SetString(language.SimplifiedChinese, msgKey, "%d 年快乐\n")
	message.SetString(language.German, msgKey, "%d Frohes neues jahr\n")
	message.SetString(language.AmericanEnglish, msgKey, "happy %d\n")

	scp := message.NewPrinter(language.SimplifiedChinese)
	scp.Printf(msgKey, 2019)
	gp := message.NewPrinter(language.German)
	gp.Printf(msgKey, 2019)
	aep := message.NewPrinter(language.AmericanEnglish)
	aep.Printf("happy new Year", 2019)
}
