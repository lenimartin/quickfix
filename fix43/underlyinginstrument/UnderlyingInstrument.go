package underlyinginstrument

//NoUnderlyingSecurityAltID is a repeating group in UnderlyingInstrument
type NoUnderlyingSecurityAltID struct {
	//UnderlyingSecurityAltID is a non-required field for NoUnderlyingSecurityAltID.
	UnderlyingSecurityAltID *string `fix:"458"`
	//UnderlyingSecurityAltIDSource is a non-required field for NoUnderlyingSecurityAltID.
	UnderlyingSecurityAltIDSource *string `fix:"459"`
}

//Component is a fix43 UnderlyingInstrument Component
type Component struct {
	//UnderlyingSymbol is a non-required field for UnderlyingInstrument.
	UnderlyingSymbol *string `fix:"311"`
	//UnderlyingSymbolSfx is a non-required field for UnderlyingInstrument.
	UnderlyingSymbolSfx *string `fix:"312"`
	//UnderlyingSecurityID is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityID *string `fix:"309"`
	//UnderlyingSecurityIDSource is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityIDSource *string `fix:"305"`
	//NoUnderlyingSecurityAltID is a non-required field for UnderlyingInstrument.
	NoUnderlyingSecurityAltID []NoUnderlyingSecurityAltID `fix:"457,omitempty"`
	//UnderlyingProduct is a non-required field for UnderlyingInstrument.
	UnderlyingProduct *int `fix:"462"`
	//UnderlyingCFICode is a non-required field for UnderlyingInstrument.
	UnderlyingCFICode *string `fix:"463"`
	//UnderlyingSecurityType is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityType *string `fix:"310"`
	//UnderlyingMaturityMonthYear is a non-required field for UnderlyingInstrument.
	UnderlyingMaturityMonthYear *string `fix:"313"`
	//UnderlyingMaturityDate is a non-required field for UnderlyingInstrument.
	UnderlyingMaturityDate *string `fix:"542"`
	//UnderlyingCouponPaymentDate is a non-required field for UnderlyingInstrument.
	UnderlyingCouponPaymentDate *string `fix:"241"`
	//UnderlyingIssueDate is a non-required field for UnderlyingInstrument.
	UnderlyingIssueDate *string `fix:"242"`
	//UnderlyingRepoCollateralSecurityType is a non-required field for UnderlyingInstrument.
	UnderlyingRepoCollateralSecurityType *int `fix:"243"`
	//UnderlyingRepurchaseTerm is a non-required field for UnderlyingInstrument.
	UnderlyingRepurchaseTerm *int `fix:"244"`
	//UnderlyingRepurchaseRate is a non-required field for UnderlyingInstrument.
	UnderlyingRepurchaseRate *float64 `fix:"245"`
	//UnderlyingFactor is a non-required field for UnderlyingInstrument.
	UnderlyingFactor *float64 `fix:"246"`
	//UnderlyingCreditRating is a non-required field for UnderlyingInstrument.
	UnderlyingCreditRating *string `fix:"256"`
	//UnderlyingInstrRegistry is a non-required field for UnderlyingInstrument.
	UnderlyingInstrRegistry *string `fix:"595"`
	//UnderlyingCountryOfIssue is a non-required field for UnderlyingInstrument.
	UnderlyingCountryOfIssue *string `fix:"592"`
	//UnderlyingStateOrProvinceOfIssue is a non-required field for UnderlyingInstrument.
	UnderlyingStateOrProvinceOfIssue *string `fix:"593"`
	//UnderlyingLocaleOfIssue is a non-required field for UnderlyingInstrument.
	UnderlyingLocaleOfIssue *string `fix:"594"`
	//UnderlyingRedemptionDate is a non-required field for UnderlyingInstrument.
	UnderlyingRedemptionDate *string `fix:"247"`
	//UnderlyingStrikePrice is a non-required field for UnderlyingInstrument.
	UnderlyingStrikePrice *float64 `fix:"316"`
	//UnderlyingOptAttribute is a non-required field for UnderlyingInstrument.
	UnderlyingOptAttribute *string `fix:"317"`
	//UnderlyingContractMultiplier is a non-required field for UnderlyingInstrument.
	UnderlyingContractMultiplier *float64 `fix:"436"`
	//UnderlyingCouponRate is a non-required field for UnderlyingInstrument.
	UnderlyingCouponRate *float64 `fix:"435"`
	//UnderlyingSecurityExchange is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityExchange *string `fix:"308"`
	//UnderlyingIssuer is a non-required field for UnderlyingInstrument.
	UnderlyingIssuer *string `fix:"306"`
	//EncodedUnderlyingIssuerLen is a non-required field for UnderlyingInstrument.
	EncodedUnderlyingIssuerLen *int `fix:"362"`
	//EncodedUnderlyingIssuer is a non-required field for UnderlyingInstrument.
	EncodedUnderlyingIssuer *string `fix:"363"`
	//UnderlyingSecurityDesc is a non-required field for UnderlyingInstrument.
	UnderlyingSecurityDesc *string `fix:"307"`
	//EncodedUnderlyingSecurityDescLen is a non-required field for UnderlyingInstrument.
	EncodedUnderlyingSecurityDescLen *int `fix:"364"`
	//EncodedUnderlyingSecurityDesc is a non-required field for UnderlyingInstrument.
	EncodedUnderlyingSecurityDesc *string `fix:"365"`
}

func New() *Component { return new(Component) }

func (m *Component) SetUnderlyingSymbol(v string)           { m.UnderlyingSymbol = &v }
func (m *Component) SetUnderlyingSymbolSfx(v string)        { m.UnderlyingSymbolSfx = &v }
func (m *Component) SetUnderlyingSecurityID(v string)       { m.UnderlyingSecurityID = &v }
func (m *Component) SetUnderlyingSecurityIDSource(v string) { m.UnderlyingSecurityIDSource = &v }
func (m *Component) SetNoUnderlyingSecurityAltID(v []NoUnderlyingSecurityAltID) {
	m.NoUnderlyingSecurityAltID = v
}
func (m *Component) SetUnderlyingProduct(v int)              { m.UnderlyingProduct = &v }
func (m *Component) SetUnderlyingCFICode(v string)           { m.UnderlyingCFICode = &v }
func (m *Component) SetUnderlyingSecurityType(v string)      { m.UnderlyingSecurityType = &v }
func (m *Component) SetUnderlyingMaturityMonthYear(v string) { m.UnderlyingMaturityMonthYear = &v }
func (m *Component) SetUnderlyingMaturityDate(v string)      { m.UnderlyingMaturityDate = &v }
func (m *Component) SetUnderlyingCouponPaymentDate(v string) { m.UnderlyingCouponPaymentDate = &v }
func (m *Component) SetUnderlyingIssueDate(v string)         { m.UnderlyingIssueDate = &v }
func (m *Component) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.UnderlyingRepoCollateralSecurityType = &v
}
func (m *Component) SetUnderlyingRepurchaseTerm(v int)     { m.UnderlyingRepurchaseTerm = &v }
func (m *Component) SetUnderlyingRepurchaseRate(v float64) { m.UnderlyingRepurchaseRate = &v }
func (m *Component) SetUnderlyingFactor(v float64)         { m.UnderlyingFactor = &v }
func (m *Component) SetUnderlyingCreditRating(v string)    { m.UnderlyingCreditRating = &v }
func (m *Component) SetUnderlyingInstrRegistry(v string)   { m.UnderlyingInstrRegistry = &v }
func (m *Component) SetUnderlyingCountryOfIssue(v string)  { m.UnderlyingCountryOfIssue = &v }
func (m *Component) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.UnderlyingStateOrProvinceOfIssue = &v
}
func (m *Component) SetUnderlyingLocaleOfIssue(v string)       { m.UnderlyingLocaleOfIssue = &v }
func (m *Component) SetUnderlyingRedemptionDate(v string)      { m.UnderlyingRedemptionDate = &v }
func (m *Component) SetUnderlyingStrikePrice(v float64)        { m.UnderlyingStrikePrice = &v }
func (m *Component) SetUnderlyingOptAttribute(v string)        { m.UnderlyingOptAttribute = &v }
func (m *Component) SetUnderlyingContractMultiplier(v float64) { m.UnderlyingContractMultiplier = &v }
func (m *Component) SetUnderlyingCouponRate(v float64)         { m.UnderlyingCouponRate = &v }
func (m *Component) SetUnderlyingSecurityExchange(v string)    { m.UnderlyingSecurityExchange = &v }
func (m *Component) SetUnderlyingIssuer(v string)              { m.UnderlyingIssuer = &v }
func (m *Component) SetEncodedUnderlyingIssuerLen(v int)       { m.EncodedUnderlyingIssuerLen = &v }
func (m *Component) SetEncodedUnderlyingIssuer(v string)       { m.EncodedUnderlyingIssuer = &v }
func (m *Component) SetUnderlyingSecurityDesc(v string)        { m.UnderlyingSecurityDesc = &v }
func (m *Component) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.EncodedUnderlyingSecurityDescLen = &v
}
func (m *Component) SetEncodedUnderlyingSecurityDesc(v string) { m.EncodedUnderlyingSecurityDesc = &v }
