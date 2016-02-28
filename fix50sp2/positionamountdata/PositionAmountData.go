package positionamountdata

//NoPosAmt is a repeating group in PositionAmountData
type NoPosAmt struct {
	//PosAmtType is a non-required field for NoPosAmt.
	PosAmtType *string `fix:"707"`
	//PosAmt is a non-required field for NoPosAmt.
	PosAmt *float64 `fix:"708"`
	//PositionCurrency is a non-required field for NoPosAmt.
	PositionCurrency *string `fix:"1055"`
}

//Component is a fix50sp2 PositionAmountData Component
type Component struct {
	//NoPosAmt is a non-required field for PositionAmountData.
	NoPosAmt []NoPosAmt `fix:"753,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoPosAmt(v []NoPosAmt) { m.NoPosAmt = v }
