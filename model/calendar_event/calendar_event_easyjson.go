// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package calendar_event

import (
	json "encoding/json"
	attachment "github.com/deadline-team/dtalks-bot-api/model/attachment"
	authorization "github.com/deadline-team/dtalks-bot-api/model/authorization"
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

func easyjsonFfeb30e3DecodeGithubComDeadlineTeamDtalksBotApiModelCalendarEvent(in *jlexer.Lexer, out *CalendarEvent) {
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
		case "createDate":
			if in.IsNull() {
				in.Skip()
				out.CreateDate = nil
			} else {
				if out.CreateDate == nil {
					out.CreateDate = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.CreateDate).UnmarshalJSON(data))
				}
			}
		case "updateDate":
			if in.IsNull() {
				in.Skip()
				out.UpdateDate = nil
			} else {
				if out.UpdateDate == nil {
					out.UpdateDate = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.UpdateDate).UnmarshalJSON(data))
				}
			}
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "location":
			out.Location = string(in.String())
		case "organizer":
			if in.IsNull() {
				in.Skip()
				out.Organizer = nil
			} else {
				if out.Organizer == nil {
					out.Organizer = new(authorization.User)
				}
				(*out.Organizer).UnmarshalEasyJSON(in)
			}
		case "startDate":
			if in.IsNull() {
				in.Skip()
				out.StartDate = nil
			} else {
				if out.StartDate == nil {
					out.StartDate = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.StartDate).UnmarshalJSON(data))
				}
			}
		case "endDate":
			if in.IsNull() {
				in.Skip()
				out.EndDate = nil
			} else {
				if out.EndDate == nil {
					out.EndDate = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.EndDate).UnmarshalJSON(data))
				}
			}
		case "transparency":
			out.Transparency = TransparencyType(in.String())
		case "visibility":
			out.Visibility = VisibilityType(in.String())
		case "members":
			if in.IsNull() {
				in.Skip()
				out.Members = nil
			} else {
				in.Delim('[')
				if out.Members == nil {
					if !in.IsDelim(']') {
						out.Members = make([]*authorization.User, 0, 8)
					} else {
						out.Members = []*authorization.User{}
					}
				} else {
					out.Members = (out.Members)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *authorization.User
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(authorization.User)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Members = append(out.Members, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "externalMembers":
			if in.IsNull() {
				in.Skip()
				out.ExternalMembers = nil
			} else {
				in.Delim('[')
				if out.ExternalMembers == nil {
					if !in.IsDelim(']') {
						out.ExternalMembers = make([]string, 0, 4)
					} else {
						out.ExternalMembers = []string{}
					}
				} else {
					out.ExternalMembers = (out.ExternalMembers)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.ExternalMembers = append(out.ExternalMembers, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "membersStatuses":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.MembersStatuses = make(MembersStatusesMap)
				} else {
					out.MembersStatuses = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 MemberStatusType
					v3 = MemberStatusType(in.String())
					(out.MembersStatuses)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "attachments":
			if in.IsNull() {
				in.Skip()
				out.Attachments = nil
			} else {
				in.Delim('[')
				if out.Attachments == nil {
					if !in.IsDelim(']') {
						out.Attachments = make([]*attachment.Attachment, 0, 8)
					} else {
						out.Attachments = []*attachment.Attachment{}
					}
				} else {
					out.Attachments = (out.Attachments)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *attachment.Attachment
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(attachment.Attachment)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Attachments = append(out.Attachments, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonFfeb30e3EncodeGithubComDeadlineTeamDtalksBotApiModelCalendarEvent(out *jwriter.Writer, in CalendarEvent) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.CreateDate != nil {
		const prefix string = ",\"createDate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.CreateDate).MarshalJSON())
	}
	if in.UpdateDate != nil {
		const prefix string = ",\"updateDate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.UpdateDate).MarshalJSON())
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.Location != "" {
		const prefix string = ",\"location\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Location))
	}
	if in.Organizer != nil {
		const prefix string = ",\"organizer\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Organizer).MarshalEasyJSON(out)
	}
	if in.StartDate != nil {
		const prefix string = ",\"startDate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.StartDate).MarshalJSON())
	}
	if in.EndDate != nil {
		const prefix string = ",\"endDate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.EndDate).MarshalJSON())
	}
	if in.Transparency != "" {
		const prefix string = ",\"transparency\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Transparency))
	}
	if in.Visibility != "" {
		const prefix string = ",\"visibility\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Visibility))
	}
	if len(in.Members) != 0 {
		const prefix string = ",\"members\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Members {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.ExternalMembers) != 0 {
		const prefix string = ",\"externalMembers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v7, v8 := range in.ExternalMembers {
				if v7 > 0 {
					out.RawByte(',')
				}
				out.String(string(v8))
			}
			out.RawByte(']')
		}
	}
	if len(in.MembersStatuses) != 0 {
		const prefix string = ",\"membersStatuses\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v9First := true
			for v9Name, v9Value := range in.MembersStatuses {
				if v9First {
					v9First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v9Name))
				out.RawByte(':')
				out.String(string(v9Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.Attachments) != 0 {
		const prefix string = ",\"attachments\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v10, v11 := range in.Attachments {
				if v10 > 0 {
					out.RawByte(',')
				}
				if v11 == nil {
					out.RawString("null")
				} else {
					(*v11).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CalendarEvent) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFfeb30e3EncodeGithubComDeadlineTeamDtalksBotApiModelCalendarEvent(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CalendarEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFfeb30e3EncodeGithubComDeadlineTeamDtalksBotApiModelCalendarEvent(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CalendarEvent) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFfeb30e3DecodeGithubComDeadlineTeamDtalksBotApiModelCalendarEvent(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CalendarEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFfeb30e3DecodeGithubComDeadlineTeamDtalksBotApiModelCalendarEvent(l, v)
}
