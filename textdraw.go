package sampgo

import "fmt"

const (
	FontSanAndreas = iota
	FontClear
	FontCapitalClear
	FontGTA
	FontSprite
)

type CommonTextDrawLike interface {
	Destroy()
	SetString(text string)
	Font(font int)
	UseBox(use bool)
	SetAlignment(align int)
	SetTextSize(x, y float32)
	SetColor(color int)
	SetBoxColor(color int)
	SetBackgroundColor(color int)
	SetSelectable(selectable bool)
	SetPreviewModel(modelindex int) error
	SetPreviewRot(rotX, rotY, rotZ, zoom float32) error
	SetPreviewVehCol(color1, color2 int) error
}

type TextDrawLike interface {
	CommonTextDrawLike
	ShowForPlayer(p PlayerLike)
	ShowForAll()
	HideForPlayer(p PlayerLike)
	HideForAll()
}

type PlayerTextDrawLike interface {
	CommonTextDrawLike
	HasPlayerLike
	Show()
	Hide()
}

type TextDraw struct {
	textDraw int
	align    int
}

func NewTextDraw(x, y float32, text string) (TextDraw, error) {
	td := TextDraw{textDraw: TextDrawCreate(x, y, text)}
	if td.textDraw == InvalidTextDraw {
		return td, fmt.Errorf("invalid textdraw")
	}
	return td, nil
}

func (td *TextDraw) Destroy() {
	TextDrawDestroy(td.textDraw)
}

func (td *TextDraw) SetString(text string) {
	TextDrawSetString(td.textDraw, text)
}

func (td *TextDraw) ShowForPlayer(p PlayerLike) {
	TextDrawShowForPlayer(p.GetID(), td.textDraw)
}
func (td *TextDraw) ShowForAll() {
	TextDrawShowForAll(td.textDraw)
}

func (td *TextDraw) HideForPlayer(p PlayerLike) {
	TextDrawHideForPlayer(p.GetID(), td.textDraw)
}

func (td *TextDraw) HideForAll() {
	TextDrawHideForAll(td.textDraw)
}

func (td *TextDraw) Font(font int) {
	TextDrawFont(td.textDraw, font)
}

func (td *TextDraw) UseBox(use bool) {
	TextDrawUseBox(td.textDraw, use)
}

func (td *TextDraw) SetAlignment(align int) {
	td.align = align
	TextDrawAlignment(td.textDraw, td.align)
}

func (td *TextDraw) SetTextSize(x, y float32) {
	if td.align == 2 {
		x, y = y, x
	}
	TextDrawTextSize(td.textDraw, x, y)
}

func (td *TextDraw) SetColor(color int) {
	TextDrawColor(td.textDraw, color)
}

func (td *TextDraw) SetBoxColor(color int) {
	TextDrawBoxColor(td.textDraw, color)
}

func (td *TextDraw) SetBackgroundColor(color int) {
	TextDrawBackgroundColor(td.textDraw, color)
}

func (td *TextDraw) SetSelectable(selectable bool) {
	TextDrawSetSelectable(td.textDraw, selectable)
}

func (td *TextDraw) SetPreviewModel(modelindex int) error {
	if !TextDrawSetPreviewModel(td.textDraw, modelindex) {
		return fmt.Errorf("invalid td.ayer or textdraw")
	}
	return nil
}

func (td *TextDraw) SetPreviewRot(rotX, rotY, rotZ, zoom float32) error {
	if !TextDrawSetPreviewRot(td.textDraw, rotX, rotY, rotZ, zoom) {
		return fmt.Errorf("invalid td.ayer or textdraw")
	}
	return nil
}

func (td *TextDraw) SetPreviewVehCol(color1, color2 int) error {
	if !TextDrawSetPreviewVehCol(td.textDraw, color1, color2) {
		return fmt.Errorf("invalid td.ayer or textdraw")
	}
	return nil
}

type PlayerTextDraw struct {
	player   *Player
	textDraw int
	align    int
}

func (p *Player) NewPlayerTextDraw(x, y float32, text string) (PlayerTextDraw, error) {
	td := PlayerTextDraw{player: p, textDraw: CreatePlayerTextDraw(p.ID, x, y, text)}
	if td.textDraw == InvalidTextDraw {
		return td, fmt.Errorf("invalid playertextdraw")
	}
	return td, nil
}

func (p *PlayerTextDraw) Destroy() {
	PlayerTextDrawDestroy(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) SetString(text string) {
	PlayerTextDrawSetString(p.player.ID, p.textDraw, text)
}

func (p *PlayerTextDraw) Show() {
	PlayerTextDrawShow(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) Hide() {
	PlayerTextDrawHide(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) Font(font int) {
	PlayerTextDrawFont(p.player.ID, p.textDraw, font)
}

func (p *PlayerTextDraw) UseBox(use bool) {
	PlayerTextDrawUseBox(p.player.ID, p.textDraw, use)
}

func (p *PlayerTextDraw) SetAlignment(align int) {
	p.align = align
	PlayerTextDrawAlignment(p.player.ID, p.textDraw, p.align)
}

func (p *PlayerTextDraw) SetTextSize(x, y float32) {
	if p.align == 2 {
		x, y = y, x
	}
	PlayerTextDrawTextSize(p.player.ID, p.textDraw, x, y)
}

func (p *PlayerTextDraw) SetColor(color int) {
	PlayerTextDrawColor(p.player.ID, p.textDraw, color)
}

var SetColour = (*PlayerTextDraw).SetColor

func (p *PlayerTextDraw) SetBoxColor(color int) {
	PlayerTextDrawBoxColor(p.player.ID, p.textDraw, color)
}

var SetBoxColour = (*PlayerTextDraw).SetBoxColor

func (p *PlayerTextDraw) SetBackgroundColor(color int) {
	PlayerTextDrawBackgroundColor(p.player.ID, p.textDraw, color)
}

var SetBackgroundColour = (*PlayerTextDraw).SetBackgroundColor

func (p *PlayerTextDraw) SetSelectable(selectable bool) {
	PlayerTextDrawSetSelectable(p.player.ID, p.textDraw, selectable)
}

func (p *PlayerTextDraw) SetPreviewModel(modelindex int) error {
	if !PlayerTextDrawSetPreviewModel(p.player.ID, p.textDraw, modelindex) {
		return fmt.Errorf("invalid player or textdraw")
	}
	return nil
}

func (p *PlayerTextDraw) SetPreviewRot(rotX, rotY, rotZ, zoom float32) error {
	if !PlayerTextDrawSetPreviewRot(p.player.ID, p.textDraw, rotX, rotY, rotZ, zoom) {
		return fmt.Errorf("invalid player or textdraw")
	}
	return nil
}

func (p *PlayerTextDraw) SetPreviewVehCol(color1, color2 int) error {
	if !PlayerTextDrawSetPreviewVehCol(p.player.ID, p.textDraw, color1, color2) {
		return fmt.Errorf("invalid player or textdraw")
	}
	return nil
}
