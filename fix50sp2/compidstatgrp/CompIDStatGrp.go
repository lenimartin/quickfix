package compidstatgrp

//NoCompIDs is a repeating group in CompIDStatGrp
type NoCompIDs struct {
	//RefCompID is a required field for NoCompIDs.
	RefCompID string `fix:"930"`
	//RefSubID is a non-required field for NoCompIDs.
	RefSubID *string `fix:"931"`
	//LocationID is a non-required field for NoCompIDs.
	LocationID *string `fix:"283"`
	//DeskID is a non-required field for NoCompIDs.
	DeskID *string `fix:"284"`
	//StatusValue is a required field for NoCompIDs.
	StatusValue int `fix:"928"`
	//StatusText is a non-required field for NoCompIDs.
	StatusText *string `fix:"929"`
}

//Component is a fix50sp2 CompIDStatGrp Component
type Component struct {
	//NoCompIDs is a required field for CompIDStatGrp.
	NoCompIDs []NoCompIDs `fix:"936"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoCompIDs(v []NoCompIDs) { m.NoCompIDs = v }
