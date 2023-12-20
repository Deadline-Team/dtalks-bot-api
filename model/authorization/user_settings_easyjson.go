// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package authorization

import (
	json "encoding/json"
	model "github.com/deadline-team/dtalks-bot-api/model"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson87cedb45DecodeGithubComDeadlineTeamDtalksBotApiModelAuthorization(in *jlexer.Lexer, out *UserSettings) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "status":
			out.Status = string(in.String())
		case "baseStatus":
			if in.IsNull() {
				in.Skip()
				out.BaseStatus = nil
			} else {
				if out.BaseStatus == nil {
					out.BaseStatus = new(BaseUserStatus)
				}
				(*out.BaseStatus).UnmarshalEasyJSON(in)
			}
		case "baseStatusMeta":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.BaseStatusMeta = make(model.Meta)
				} else {
					out.BaseStatusMeta = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 interface{}
					if m, ok := v1.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v1.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v1 = in.Interface()
					}
					(out.BaseStatusMeta)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "baseStatusEndDate":
			if in.IsNull() {
				in.Skip()
				out.BaseStatusEndDate = nil
			} else {
				if out.BaseStatusEndDate == nil {
					out.BaseStatusEndDate = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.BaseStatusEndDate).UnmarshalJSON(data))
				}
			}
		case "iamLeftHanded":
			out.IamLeftHanded = bool(in.Bool())
		case "hideMyPhoneNumber":
			out.HideMyPhoneNumber = bool(in.Bool())
		case "hideBackgroundImage":
			out.HideBackgroundImage = bool(in.Bool())
		case "disableAnimation":
			out.DisableAnimation = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson87cedb45EncodeGithubComDeadlineTeamDtalksBotApiModelAuthorization(out *jwriter.Writer, in UserSettings) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Status != "" {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	if in.BaseStatus != nil {
		const prefix string = ",\"baseStatus\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.BaseStatus).MarshalEasyJSON(out)
	}
	if len(in.BaseStatusMeta) != 0 {
		const prefix string = ",\"baseStatusMeta\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.BaseStatusMeta {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				if m, ok := v2Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v2Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v2Value))
				}
			}
			out.RawByte('}')
		}
	}
	if in.BaseStatusEndDate != nil {
		const prefix string = ",\"baseStatusEndDate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.BaseStatusEndDate).MarshalJSON())
	}
	if in.IamLeftHanded {
		const prefix string = ",\"iamLeftHanded\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IamLeftHanded))
	}
	if in.HideMyPhoneNumber {
		const prefix string = ",\"hideMyPhoneNumber\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.HideMyPhoneNumber))
	}
	if in.HideBackgroundImage {
		const prefix string = ",\"hideBackgroundImage\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.HideBackgroundImage))
	}
	if in.DisableAnimation {
		const prefix string = ",\"disableAnimation\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.DisableAnimation))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserSettings) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson87cedb45EncodeGithubComDeadlineTeamDtalksBotApiModelAuthorization(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserSettings) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson87cedb45EncodeGithubComDeadlineTeamDtalksBotApiModelAuthorization(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserSettings) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson87cedb45DecodeGithubComDeadlineTeamDtalksBotApiModelAuthorization(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserSettings) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson87cedb45DecodeGithubComDeadlineTeamDtalksBotApiModelAuthorization(l, v)
}
