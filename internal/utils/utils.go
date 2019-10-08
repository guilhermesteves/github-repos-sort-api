package utils

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var DefaultTimeMask = "2006-01-02T15:04:05-0700"

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func Check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func CheckIf(condition bool, ifTrue interface{}, ifFalse interface{}) interface{} {
	if condition {
		switch ifTrue.(type) {
		case func() interface{}:
			return ifTrue.(func() interface{})()
		}
		return ifTrue
	} else {
		switch ifFalse.(type) {
		case func() interface{}:
			return ifFalse.(func() interface{})()
		}
		return ifFalse
	}
}

func ToFormattedTime(d interface{}) string {
	switch d.(type) {
	case primitive.DateTime:
		t := time.Unix(int64(d.(primitive.DateTime))/1000, int64(d.(primitive.DateTime))%1000*1000000)
		return t.Format(DefaultTimeMask)
	case time.Time:
		return d.(time.Time).Format(DefaultTimeMask)
	}

	return ""
}

func ToStringMapString(m interface{}) map[string]string {
	if m != nil {
		switch m.(type) {
		case primitive.M:
			return cast.ToStringMapString(map[string]interface{}(m.(primitive.M)))
		case map[string]interface{}:
			return cast.ToStringMapString(m)
		case map[string]string:
			return cast.ToStringMapString(m)
		}
	}
	return map[string]string{}
}

func ToStringSlice(a interface{}) []string {
	if a != nil {
		switch a.(type) {
		case primitive.A:
			result := []string{}
			for _, element := range a.(primitive.A) {
				result = append(result, element.(string))
			}
			return result
		case []string:
			return cast.ToStringSlice(a)
		}
	}
	return []string{}
}