// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package conversation

import (
	json "encoding/json"
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

func easyjson16eb09bcDecodeGithubComDeadlineTeamDtalksBotApiModelConversation(in *jlexer.Lexer, out *LinkFilter) {
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
		case "Value":
			out.Value = string(in.String())
		case "Title":
			out.Title = string(in.String())
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
func easyjson16eb09bcEncodeGithubComDeadlineTeamDtalksBotApiModelConversation(out *jwriter.Writer, in LinkFilter) {
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
		const prefix string = ",\"Value\":"
		out.RawString(prefix)
		out.String(string(in.Value))
	}
	{
		const prefix string = ",\"Title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LinkFilter) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson16eb09bcEncodeGithubComDeadlineTeamDtalksBotApiModelConversation(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LinkFilter) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson16eb09bcEncodeGithubComDeadlineTeamDtalksBotApiModelConversation(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LinkFilter) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson16eb09bcDecodeGithubComDeadlineTeamDtalksBotApiModelConversation(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LinkFilter) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson16eb09bcDecodeGithubComDeadlineTeamDtalksBotApiModelConversation(l, v)
}
func easyjson16eb09bcDecodeGithubComDeadlineTeamDtalksBotApiModelConversation1(in *jlexer.Lexer, out *Link) {
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
		case "value":
			out.Value = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "imageUrl":
			out.ImageUrl = string(in.String())
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
func easyjson16eb09bcEncodeGithubComDeadlineTeamDtalksBotApiModelConversation1(out *jwriter.Writer, in Link) {
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
	if in.Value != "" {
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Value))
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
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
	if in.ImageUrl != "" {
		const prefix string = ",\"imageUrl\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ImageUrl))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Link) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson16eb09bcEncodeGithubComDeadlineTeamDtalksBotApiModelConversation1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Link) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson16eb09bcEncodeGithubComDeadlineTeamDtalksBotApiModelConversation1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Link) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson16eb09bcDecodeGithubComDeadlineTeamDtalksBotApiModelConversation1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Link) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson16eb09bcDecodeGithubComDeadlineTeamDtalksBotApiModelConversation1(l, v)
}
