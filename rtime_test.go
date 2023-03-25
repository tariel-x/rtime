package rtime_test

import (
	"testing"
	"time"

	"github.com/tariel-x/rtime"
)

var t1 = time.Date(2023, 03, 1, 02, 48, 05, 0, time.UTC)

var layouts1 = map[string]string{
	rtime.GOST2016Word:              "1 марта 2023 г.",
	rtime.GOST2016Numeric:           "01.03.2023",
	rtime.GOST2003Word:              "01 марта 2023 г.",
	rtime.GOST2003NumericReverse:    "2023.03.01",
	"2 января 2006 г., понедельник": "1 марта 2023 г., среда",
	"2 января 2006 г., Понедельник": "1 марта 2023 г., Среда",
	"2 января 2006 г., пн":          "1 марта 2023 г., ср",
	"2 января 2006 г., ПН":          "1 марта 2023 г., СР",
	"Понедельник, 2 янв 2006 г. в 15 часов 04 минут": "Среда, 1 мар 2023 г. в 02 часов 48 минут",
	"ПН, 2 Янв 2006 г.":      "СР, 1 Мар 2023 г.",
	"ПН/Mon, 2 Янв/Jan 2006": "СР/Wed, 1 Мар/Mar 2023",
	"Дата 2 Январь январь Янв янв Января января Понедельник понедельник ПН пн 06 г.": "Дата 1 Март март Мар мар Марта марта Среда среда СР ср 23 г.",
	"Дата 2 😊 янв 2006 г.": "Дата 1 😊 мар 2023 г.",
	time.RFC3339:           "2023-03-01T02:48:05Z",
}

var loc, _ = time.LoadLocation("Europe/Moscow")
var t2 = time.Date(1993, 03, 20, 15, 21, 31, 0, loc)

var layouts2 = map[string]string{
	rtime.GOST2016Word:              "20 марта 1993 г.",
	rtime.GOST2016Numeric:           "20.03.1993",
	rtime.GOST2003Word:              "20 марта 1993 г.",
	rtime.GOST2003NumericReverse:    "1993.03.20",
	"2 января 2006 г., понедельник": "20 марта 1993 г., суббота",
	"2 января 2006 г., Понедельник": "20 марта 1993 г., Суббота",
	"2 января 2006 г., пн":          "20 марта 1993 г., сб",
	"2 января 2006 г., ПН":          "20 марта 1993 г., СБ",
	"Понедельник, 2 янв 2006 г. в 15 часов 04 минут": "Суббота, 20 мар 1993 г. в 15 часов 21 минут",
	"ПН, 2 Янв 2006 г.":      "СБ, 20 Мар 1993 г.",
	"ПН/Mon, 2 Янв/Jan 2006": "СБ/Sat, 20 Мар/Mar 1993",
	"Дата 2 Январь январь Янв янв Января января Понедельник понедельник ПН пн 06 г.": "Дата 20 Март март Мар мар Марта марта Суббота суббота СБ сб 93 г.",
	"Дата 2 😊 янв 2006 г.": "Дата 20 😊 мар 1993 г.",
	time.RFC3339:           "1993-03-20T15:21:31+03:00",
}

func TestFormat(t *testing.T) {
	rt1 := rtime.RTime{Time: t1}
	for layout, expected := range layouts1 {
		result := rt1.Format(layout)
		if result != expected {
			t.Errorf("Expected %q, result %q, layoyt %q", expected, result, layout)
		}
	}

	rt2 := rtime.RTime{Time: t2}
	for layout, expected := range layouts2 {
		result := rt2.Format(layout)
		if result != expected {
			t.Errorf("Expected %q, result %q, layoyt %q", expected, result, layout)
		}
	}
}

func TestDate(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Moscow")
	tt := time.Date(2023, 03, 1, 02, 48, 05, 0, loc)
	rt := rtime.Date(2023, 03, 1, 02, 48, 05, 0, loc)
	ttFormatted := tt.Format(time.RFC3339)
	rtFormatted := rt.Format(time.RFC3339)
	if ttFormatted != rtFormatted {
		t.Errorf("time %s and rtime %s are not equal", ttFormatted, rtFormatted)
	}
}

func TestAdd(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Moscow")
	rt := rtime.Date(2023, 03, 1, 02, 48, 05, 0, loc)
	rt = rt.Add(time.Hour * 24)
	if rt.Format(rtime.GOST2016Numeric) != "02.03.2023" {
		t.Errorf("Expected 02.03.2023, actual %s", rt.Format(rtime.GOST2016Numeric))
	}
}
