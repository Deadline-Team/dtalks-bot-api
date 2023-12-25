// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package user

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

func easyjson9e1087fdDecodeGithubComDeadlineTeamDtalksBotApiModelUser(in *jlexer.Lexer, out *UserFilter) {
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
		case "IDs":
			if in.IsNull() {
				in.Skip()
				out.IDs = nil
			} else {
				in.Delim('[')
				if out.IDs == nil {
					if !in.IsDelim(']') {
						out.IDs = make([]string, 0, 4)
					} else {
						out.IDs = []string{}
					}
				} else {
					out.IDs = (out.IDs)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.IDs = append(out.IDs, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Username":
			out.Username = string(in.String())
		case "FirstName":
			out.FirstName = string(in.String())
		case "LastName":
			out.LastName = string(in.String())
		case "Email":
			out.Email = string(in.String())
		case "Search":
			out.Search = string(in.String())
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
func easyjson9e1087fdEncodeGithubComDeadlineTeamDtalksBotApiModelUser(out *jwriter.Writer, in UserFilter) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"IDs\":"
		out.RawString(prefix[1:])
		if in.IDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.IDs {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"FirstName\":"
		out.RawString(prefix)
		out.String(string(in.FirstName))
	}
	{
		const prefix string = ",\"LastName\":"
		out.RawString(prefix)
		out.String(string(in.LastName))
	}
	{
		const prefix string = ",\"Email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"Search\":"
		out.RawString(prefix)
		out.String(string(in.Search))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserFilter) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComDeadlineTeamDtalksBotApiModelUser(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserFilter) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComDeadlineTeamDtalksBotApiModelUser(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserFilter) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComDeadlineTeamDtalksBotApiModelUser(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserFilter) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComDeadlineTeamDtalksBotApiModelUser(l, v)
}
func easyjson9e1087fdDecodeGithubComDeadlineTeamDtalksBotApiModelUser1(in *jlexer.Lexer, out *User) {
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
		case "source":
			out.Source = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "firstName":
			out.FirstName = string(in.String())
		case "lastName":
			out.LastName = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "position":
			out.Position = string(in.String())
		case "avatar":
			if in.IsNull() {
				in.Skip()
				out.Avatar = nil
			} else {
				if out.Avatar == nil {
					out.Avatar = new(model.Avatar)
				}
				(*out.Avatar).UnmarshalEasyJSON(in)
			}
		case "birthday":
			if in.IsNull() {
				in.Skip()
				out.Birthday = nil
			} else {
				if out.Birthday == nil {
					out.Birthday = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Birthday).UnmarshalJSON(data))
				}
			}
		case "phoneNumber":
			out.PhoneNumber = int64(in.Int64())
		case "city":
			out.City = string(in.String())
		case "company":
			out.Company = string(in.String())
		case "department":
			out.Department = string(in.String())
		case "chief":
			if in.IsNull() {
				in.Skip()
				out.Chief = nil
			} else {
				if out.Chief == nil {
					out.Chief = new(User)
				}
				(*out.Chief).UnmarshalEasyJSON(in)
			}
		case "lastActivity":
			if in.IsNull() {
				in.Skip()
				out.LastActivity = nil
			} else {
				if out.LastActivity == nil {
					out.LastActivity = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LastActivity).UnmarshalJSON(data))
				}
			}
		case "blocked":
			out.Blocked = bool(in.Bool())
		case "timeZone":
			out.TimeZone = int64(in.Int64())
		case "canChangePassword":
			out.CanChangePassword = bool(in.Bool())
		case "canChangeAvatar":
			out.CanChangeAvatar = bool(in.Bool())
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
func easyjson9e1087fdEncodeGithubComDeadlineTeamDtalksBotApiModelUser1(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Source != "" {
		const prefix string = ",\"source\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Source))
	}
	if in.Username != "" {
		const prefix string = ",\"username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	if in.FirstName != "" {
		const prefix string = ",\"firstName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FirstName))
	}
	if in.LastName != "" {
		const prefix string = ",\"lastName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LastName))
	}
	if in.Email != "" {
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	if in.Position != "" {
		const prefix string = ",\"position\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Position))
	}
	if in.Avatar != nil {
		const prefix string = ",\"avatar\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Avatar).MarshalEasyJSON(out)
	}
	if in.Birthday != nil {
		const prefix string = ",\"birthday\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.Birthday).MarshalJSON())
	}
	if in.PhoneNumber != 0 {
		const prefix string = ",\"phoneNumber\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PhoneNumber))
	}
	if in.City != "" {
		const prefix string = ",\"city\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.City))
	}
	if in.Company != "" {
		const prefix string = ",\"company\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Company))
	}
	if in.Department != "" {
		const prefix string = ",\"department\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Department))
	}
	if in.Chief != nil {
		const prefix string = ",\"chief\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Chief).MarshalEasyJSON(out)
	}
	if in.LastActivity != nil {
		const prefix string = ",\"lastActivity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.LastActivity).MarshalJSON())
	}
	if in.Blocked {
		const prefix string = ",\"blocked\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Blocked))
	}
	if in.TimeZone != 0 {
		const prefix string = ",\"timeZone\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.TimeZone))
	}
	if in.CanChangePassword {
		const prefix string = ",\"canChangePassword\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.CanChangePassword))
	}
	if in.CanChangeAvatar {
		const prefix string = ",\"canChangeAvatar\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.CanChangeAvatar))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComDeadlineTeamDtalksBotApiModelUser1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComDeadlineTeamDtalksBotApiModelUser1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComDeadlineTeamDtalksBotApiModelUser1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComDeadlineTeamDtalksBotApiModelUser1(l, v)
}
