package handlers

import (
	"MyTGbot/internal/storage"

	"gopkg.in/telebot.v3"
)

type NotesHandler struct {
	db storage.Database
}

func (h *NotesHandler) HandleSaveNote(c telebot.Context) error {
	// Логика сохранения заметки
	return c.Send("Заметка сохранена!")
}
