package GameLogic

type Event interface {
	RunEvent()
}
type Choice struct {
	Active     bool
	Repeatable bool
	ChoiceText string
	choices    []*Option
}

type Option struct {
	id          int
	text        string
	consequence Event
}

type Combat struct {
	Active    bool
	EnemyList []Enemy
}
type BossFight struct {
	Combat
}
type RandomEncounter struct {
	Active        bool
	EncounterText string
	Choices       Choice
}

var NavChoices []Choice = []Choice{
	{Active: false, ChoiceText: "A corridor splits in half up ahead. Do you go left or right?", choices: []*Option{{id: 1, text: "Take the left door (You hear men screaming)"}, {id: 2, text: "Take the right door (You hear... music?)"}}},
	{Active: false, ChoiceText: "A lamp lights up the room you are in. There are three closed doors...", choices: []*Option{}},
}
var EncounterChoices []Choice = []Choice{
	{Active: false, ChoiceText: `Ahead of you lies two chalices. The left chalice seems to grant great strength at the 
	cost of some of your vitality while the right one gives you greater vitality at the cost of health..`, choices: []*Option{}},
}

var VictoryChoice Choice = Choice{Active: false, ChoiceText: "You won the fight!", choices: []*Option{}}
var BossChoice Choice = Choice{Active: false, ChoiceText: "You sense that a dastardly foe awaits ahead", choices: []*Option{{id: 0, text: "No turning back now... (Press any key)"}}}
