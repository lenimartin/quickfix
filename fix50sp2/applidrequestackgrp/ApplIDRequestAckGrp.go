package applidrequestackgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties"
)

//NoApplIDs is a repeating group in ApplIDRequestAckGrp
type NoApplIDs struct {
	//RefApplID is a non-required field for NoApplIDs.
	RefApplID *string `fix:"1355"`
	//ApplBegSeqNum is a non-required field for NoApplIDs.
	ApplBegSeqNum *int `fix:"1182"`
	//ApplEndSeqNum is a non-required field for NoApplIDs.
	ApplEndSeqNum *int `fix:"1183"`
	//RefApplLastSeqNum is a non-required field for NoApplIDs.
	RefApplLastSeqNum *int `fix:"1357"`
	//ApplResponseError is a non-required field for NoApplIDs.
	ApplResponseError *int `fix:"1354"`
	//NestedParties Component
	NestedParties nestedparties.Component
	//RefApplReqID is a non-required field for NoApplIDs.
	RefApplReqID *string `fix:"1433"`
}

//Component is a fix50sp2 ApplIDRequestAckGrp Component
type Component struct {
	//NoApplIDs is a non-required field for ApplIDRequestAckGrp.
	NoApplIDs []NoApplIDs `fix:"1351,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoApplIDs(v []NoApplIDs) { m.NoApplIDs = v }
