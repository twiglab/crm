package gqlx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func MarshalInt64(i int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		fmt.Fprintf(w, "\"%d\"", i)
	})
}

func UnmarshalInt64(v any) (int64, error) {
	/*
		var (
			val string
			ok  bool
		)

		if val, ok = v.(string); !ok {
			return 0, fmt.Errorf("%T is not a int", v)
		}

		return strconv.ParseInt(val, 10, 64)
	*/
	switch v := v.(type) {
	case string:
		return strconv.ParseInt(v, 10, 64)
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case json.Number:
		return v.Int64()
		//return strconv.ParseInt(string(v), 10, 64)
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}

func MarshalUUID(id uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		bs, _ := id.MarshalBinary()
		w.Write(bs)
	})
}

func UnmarshalUUID(v any) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		return uuid.Parse(v)
	case []byte:
		return uuid.ParseBytes(v)
	default:
		return uuid.Nil, fmt.Errorf("%T is not a uuid", v)
	}
}

func MarshalDateTime(t time.Time) graphql.Marshaler {
	if t.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(time.DateTime)))
	})
}

func UnmarshalDateTime(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse(time.DateTime, tmpStr)
	}
	return time.Time{}, errors.New("time should be RFC3339Nano formatted string")
}
