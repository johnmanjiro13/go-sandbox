package model

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"net/url"
	"strconv"
)

type MyURL struct {
	url.URL
}

// MarshalGQL implements the graphql.Marshaler interface
func (u MyURL) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(u.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *MyURL) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		if result, err := url.Parse(v); err != nil {
			return err
		} else {
			u = &MyURL{*result}
		}
		return nil
	case []byte:
		result := &url.URL{}
		if err := result.UnmarshalBinary(v); err != nil {
			return err
		}
		u = &MyURL{*result}
		return nil
	default:
		return fmt.Errorf("%T is not a string", v)
	}
}

func MarshalURI(u url.URL) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(u.String()))
	})
}

func UnmarshalURI(v interface{}) (url.URL, error) {
	switch v := v.(type) {
	case string:
		if result, err := url.Parse(v); err != nil {
			return url.URL{}, err
		} else {
			return *result, nil
		}
	case []byte:
		result := &url.URL{}
		if err := result.UnmarshalBinary(v); err != nil {
			return url.URL{}, err
		}
		return *result, nil
	default:
		return url.URL{}, fmt.Errorf("%T is not a string", v)
	}
}
