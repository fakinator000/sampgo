package sampgo

import (
	"fmt"
	"time"
)

type ObjectLike interface {
	GetID() int
	Destroy()
	IsValid() bool
	Move(x, y, z, speed, rotX, rotY, rotZ float32) time.Duration
	IsMoving() bool
	Stop()
	SetPos(x, y, z float32)
	SetRot(rx, ry, rz float32)
	GetPos() (x, y, z float32, err error)
	GetRot() (rx, ry, rz float32, err error)
}

type PlayerObjectLike interface {
	ObjectLike
	HasPlayerLike
}

type Object struct {
	ID int
}

func (o *Object) GetID() int {
	return o.ID
}

func NewObject(modelid int, x, y, z, rX, rY, rZ, drawDistance float32) (o *Object) {
	o = new(Object)
	o.ID = CreateObject(modelid, x, y, z, rX, rY, rZ, drawDistance)
	return
}

func (o *Object) Destroy() {
	DestroyObject(o.ID)
}

func (o *Object) IsValid() bool {
	return IsValidObject(o.ID)
}

// Returns the time it will take for the object to move
func (o *Object) Move(x, y, z, speed, rotX, rotY, rotZ float32) time.Duration {
	return time.Duration(MoveObject(o.ID, x, y, z, speed, rotX, rotY, rotZ)) * time.Millisecond
}

func (o *Object) IsMoving() bool {
	return IsObjectMoving(o.ID)
}

func (o *Object) Stop() {
	StopObject(o.ID)
}

func (o *Object) SetPos(x, y, z float32) {
	SetObjectPos(o.ID, x, y, z)
}
func (o *Object) SetRot(rx, ry, rz float32) {
	SetObjectRot(o.ID, rx, ry, rz)
}

func (o *Object) GetPos() (x, y, z float32, err error) {
	if !GetObjectPos(o.ID, &x, &y, &z) {
		err = fmt.Errorf("invalid object")
	}
	return
}

func (o *Object) GetRot() (rx, ry, rz float32, err error) {
	if !GetObjectRot(o.ID, &rx, &ry, &rz) {
		err = fmt.Errorf("invalid object")
	}
	return
}

type PlayerObject struct {
	ID     int
	player PlayerLike
}

func (o *PlayerObject) GetID() int {
	return o.ID
}

func (o *PlayerObject) GetPlayer() PlayerLike {
	return o.player
}

func NewPlayerObject(p PlayerLike, modelid int, x, y, z, rX, rY, rZ, drawDistance float32) (o *PlayerObject) {
	o = new(PlayerObject)
	o.player = p
	o.ID = CreatePlayerObject(o.player.GetID(), modelid, x, y, z, rX, rY, rZ, drawDistance)
	return
}

func (o *PlayerObject) Destroy() {
	DestroyPlayerObject(o.player.GetID(), o.ID)
}

func (o *PlayerObject) IsValid() bool {
	return IsValidPlayerObject(o.player.GetID(), o.ID)
}

// Returns the time it will take for the object to move
func (o *PlayerObject) Move(x, y, z, speed, rotX, rotY, rotZ float32) time.Duration {
	return time.Duration(MovePlayerObject(o.player.GetID(), o.ID, x, y, z, speed, rotX, rotY, rotZ)) * time.Millisecond
}

func (o *PlayerObject) IsMoving() bool {
	return IsPlayerObjectMoving(o.player.GetID(), o.ID)
}

func (o *PlayerObject) Stop() {
	StopPlayerObject(o.player.GetID(), o.ID)
}

func (o *PlayerObject) SetPos(x, y, z float32) {
	SetPlayerObjectPos(o.player.GetID(), o.ID, x, y, z)
}
func (o *PlayerObject) SetRot(rx, ry, rz float32) {
	SetPlayerObjectRot(o.player.GetID(), o.ID, rx, ry, rz)
}

func (o *PlayerObject) GetPos() (x, y, z float32, err error) {
	if !GetPlayerObjectPos(o.player.GetID(), o.ID, &x, &y, &z) {
		err = fmt.Errorf("invalid object")
	}
	return
}

func (o *PlayerObject) GetRot() (rx, ry, rz float32, err error) {
	if !GetPlayerObjectRot(o.player.GetID(), o.ID, &rx, &ry, &rz) {
		err = fmt.Errorf("invalid object")
	}
	return
}
