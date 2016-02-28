//Package sequencereset msg type = 4.
package sequencereset

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
)

//Message is a SequenceReset FIX Message
type Message struct {
	FIXMsgType string `fix:"4"`
	Header     fix40.Header
	//GapFillFlag is a non-required field for SequenceReset.
	GapFillFlag *string `fix:"123"`
	//NewSeqNo is a required field for SequenceReset.
	NewSeqNo int `fix:"36"`
	Trailer  fix40.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetGapFillFlag(v string) { m.GapFillFlag = &v }
func (m *Message) SetNewSeqNo(v int)       { m.NewSeqNo = v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX40, "4", r
}
