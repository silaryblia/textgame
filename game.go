package main

import "strings"

type User struct {
	CurrentRoom string
	BackPack    bool
	Inventory   []string
}

type Room struct {
	Description string
	Items       []string
	Exits       []string
}

var (
	user     User
	rooms    map[string]*Room
	doorOpen bool
)

func initGame() {
	doorOpen = false
	user = User{
		CurrentRoom: "кухня",
		BackPack:    false,
		Inventory:   []string{},
	}

	rooms = map[string]*Room{
		"кухня": {
			Description: "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ",
			Items:       []string{"чай"},
			Exits:       []string{"коридор"},
		},
		"комната": {
			Description: "ты в своей комнате",
			Items:       []string{"ключи", "конспекты", "рюкзак"},
			Exits:       []string{"коридор"},
		},
		"коридор": {
			Description: "ничего интересного",
			Items:       []string{""},
			Exits:       []string{"кухня", "комната", "улица"},
		},
		"улица": {
			Description: "на улице весна",
			Items:       []string{""},
			Exits:       []string{"домой"},
		},
	}
}

func (u *User) LookAround() string {
	room := rooms[u.CurrentRoom]

	switch u.CurrentRoom {
	case "кухня":
		if !u.BackPack {
			return "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
		}
		return "ты находишься на кухне, на столе: чай, надо идти в универ. можно пройти - коридор"
	case "комната":
		if len(room.Items) == 0 {
			return "пустая комната. можно пройти - коридор"
		}
		if contains(room.Items, "рюкзак") {
			return "на столе: ключи, конспекты, на стуле: рюкзак. можно пройти - коридор"
		}
		return "на столе: " + strings.Join(room.Items, ", ") + ". можно пройти - коридор"
	case "коридор":
		return room.Description + ". можно пройти - кухня, комната, улица"
	case "улица":
		return room.Description + ". можно пройти - домой"
	}
	return "ничего не видно"
}

func (u *User) Move(target string) string {
	room := rooms[u.CurrentRoom]
	found := false
	for _, exit := range room.Exits {
		if exit == target {
			found = true
			break
		}
	}
	if !found {
		return "нет пути в " + target
	}
	// проверка на закрытую дверь
	if target == "улица" && !doorOpen {
		return "дверь закрыта"
	}
	u.CurrentRoom = target

	switch target {
	case "комната":
		return "ты в своей комнате. можно пройти - коридор"
	case "кухня":
		if !u.BackPack {
			return "ты находишься на кухне, на столе: чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
		}
		return "кухня, ничего интересного. можно пройти - коридор"
	case "коридор":
		return "ничего интересного. можно пройти - кухня, комната, улица"
	case "улица":
		return "на улице весна. можно пройти - домой"
	default:
		return rooms[target].Description + ". можно пройти - " + strings.Join(room.Exits, ", ")
	}
}

func (u *User) Take(item string) string {
	room := rooms[u.CurrentRoom]
	// ищем предмет
	found := false
	for _, v := range room.Items {
		if v == item {
			found = true
			break
		}
	}
	if !found {
		return "нет такого"
	}
	// рюкзак
	if item == "рюкзак" {
		if u.BackPack {
			return "уже надет"
		}
		u.BackPack = true
		room.Items = removeItem(room.Items, item)
		return "вы надели: рюкзак"
	}

	if !u.BackPack {
		return "некуда класть"
	}

	u.Inventory = append(u.Inventory, item)
	room.Items = removeItem(room.Items, item)
	return "предмет добавлен в инвентарь: " + item
}

func (u *User) Use(item, target string) string {
	if !hasItem(item) {
		return "нет предмета в инвентаре - " + item
	}
	if item == "ключи" && target == "дверь" {
		doorOpen = true
		return "дверь открыта"
	}
	return "не к чему применить"
}

func hasItem(name string) bool {
	for _, v := range user.Inventory {
		if v == name {
			return true
		}
	}
	return false
}

func removeItem(items []string, name string) []string {
	newItems := []string{}
	for _, v := range items {
		if v != name {
			newItems = append(newItems, v)
		}
	}
	return newItems
}

// Диспетчер команд
func handleCommand(cmd string) string {
	cmd = strings.ToLower(cmd)
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return "неизвестная команда"
	}

	switch parts[0] {
	case "осмотреться":
		return user.LookAround()
	case "идти":
		if len(parts) < 2 {
			return "куда идти?"
		}
		return user.Move(parts[1])
	case "взять":
		if len(parts) < 2 {
			return "что взять?"
		}
		return user.Take(parts[1])
	case "надеть":
		if len(parts) < 2 {
			return "что надеть?"
		}
		return user.Take(parts[1])
	case "применить":
		if len(parts) < 3 {
			return "не к чему применит"
		}
		return user.Use(parts[1], parts[2])
	}
	return "неизвестная команда"
}

func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
