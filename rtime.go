// Package rtime provides functionality for formatting dates in russian language with new placeholders.
//
// For example, this code formats date according to the actual GOST 2006 standart:
//
//	t := rtime.Now()
//	t.Format(rtime.GOST2016Word) // 1 марта 2023 г.
//
// List of the new placeholders: Январь, январь, Янв, янв, Января, января, Понедельник, понедельник, ПН, пн.
// Every placeholder existing in the original time package works as usual:
//
//	t := rtime.RTime{time.Now()}
//	t.Format("ПН/Mon, 2 Янв/Jan 2006") // "СР/Wed, 1 Мар/Mar 2023
package rtime

import (
	"errors"
	"fmt"
	"time"
)

const (
	Layout      = time.Layout
	ANSIC       = time.ANSIC
	UnixDate    = time.UnixDate
	RubyDate    = time.RubyDate
	RFC822      = time.RFC822
	RFC822Z     = time.RFC822Z
	RFC850      = time.RFC850
	RFC1123     = time.RFC1123
	RFC1123Z    = time.RFC1123Z
	RFC3339     = time.RFC3339
	RFC3339Nano = time.RFC3339Nano
	Kitchen     = time.Kitchen
	Stamp       = time.Stamp
	StampMilli  = time.StampMilli
	StampMicro  = time.StampMicro
	StampNano   = time.StampNano
)

// Additional
const (
	GOST2016Numeric        = "02.01.2006"
	GOST2016Word           = "2 января 2006 г."
	GOST2003Word           = "02 января 2006 г."
	GOST2003NumericReverse = "2006.01.02"
)

const (
	codeLongMonth              = iota + 1 // Январь
	codeLongMonthLower                    // январь
	codeMonth                             // Янв
	codeMonthLower                        // янв
	codeLongMonthGenitive                 // Января
	codeLongMonthGenitiveLower            // января
	codeLongWeekDay                       // Понедельник
	codeLongWeekDayLower                  // понедельник
	codeWeekDay                           // ПН
	codeWeekDayLower                      // пн
)

var longMonthNames = []string{
	"Январь",
	"Февраль",
	"Март",
	"Апрель",
	"Май",
	"Июнь",
	"Июль",
	"Август",
	"Сентябрь",
	"Октябрь",
	"Ноябрь",
	"Декабрь",
}

var longMonthLowerNames = []string{
	"январь",
	"февраль",
	"март",
	"апрель",
	"май",
	"июнь",
	"июль",
	"август",
	"сентябрь",
	"октябрь",
	"ноябрь",
	"декабрь",
}

var monthNames = []string{
	"Янв",
	"Фев",
	"Мар",
	"Апр",
	"Май",
	"Июнь",
	"Июль",
	"Авг",
	"Сен",
	"Окт",
	"Ноя",
	"Дек",
}

var monthLowerNames = []string{
	"янв",
	"фев",
	"мар",
	"апр",
	"май",
	"июнь",
	"июль",
	"авг",
	"сен",
	"окт",
	"ноя",
	"дек",
}

var longMonthGenitiveNames = []string{
	"Января",
	"Февраля",
	"Марта",
	"Апреля",
	"Мая",
	"Июня",
	"Июля",
	"Августа",
	"Сентября",
	"Октября",
	"Ноября",
	"Декабря",
}

var longMonthGenitiveLowerNames = []string{
	"января",
	"февраля",
	"марта",
	"апреля",
	"мая",
	"июня",
	"июля",
	"августа",
	"сентября",
	"октября",
	"ноября",
	"декабря",
}

var longWeekDayNames = []string{
	"Понедельник",
	"Вторник",
	"Среда",
	"Четверг",
	"Пятница",
	"Суббота",
	"Воскресенье",
}

var longWeekDayLowerNames = []string{
	"понедельник",
	"вторник",
	"среда",
	"четверг",
	"пятница",
	"суббота",
	"воскресенье",
}

var weekDayNames = []string{
	"ПН",
	"ВТ",
	"СР",
	"ЧТ",
	"ПТ",
	"СБ",
	"ВС",
}

var weekDayLowerNames = []string{
	"пн",
	"вт",
	"ср",
	"чт",
	"пт",
	"сб",
	"вс",
}

var ErrInvalidNamesList = errors.New("invalid new names list")

// SetMonthNames set short month names (Янв, Фев, etc.)
func SetMonthNames(newNames []string) error {
	if len(newNames) != len(monthNames) {
		return ErrInvalidNamesList
	}
	monthNames = newNames
	return nil
}

// SetMonthLowerNames set short month lower names (янв, фев, etc.)
func SetMonthLowerNames(newNames []string) error {
	if len(newNames) != len(monthLowerNames) {
		return ErrInvalidNamesList
	}
	monthLowerNames = newNames
	return nil
}

// SetWeekDayNames set short week day names (ПН, ВТ, etc.)
func SetWeekDayNames(newNames []string) error {
	if len(newNames) != len(weekDayNames) {
		return ErrInvalidNamesList
	}
	weekDayNames = newNames
	return nil
}

// SetWeekDayLowerNames set short week day lower names (пн, вт, etc.)
func SetWeekDayLowerNames(newNames []string) error {
	if len(newNames) != len(weekDayLowerNames) {
		return ErrInvalidNamesList
	}
	weekDayLowerNames = newNames
	return nil
}

type month int

const (
	January month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (m month) string(code int) string {
	if January <= m && m <= December {
		switch code {
		case codeLongMonth:
			return longMonthNames[m-1]
		case codeLongMonthLower:
			return longMonthLowerNames[m-1]
		case codeMonth:
			return monthNames[m-1]
		case codeMonthLower:
			return monthLowerNames[m-1]
		case codeLongMonthGenitive:
			return longMonthGenitiveNames[m-1]
		case codeLongMonthGenitiveLower:
			return longMonthGenitiveLowerNames[m-1]
		default:
			return fmt.Sprintf("Month(%d)", m)
		}

	}
	return fmt.Sprintf("Month(%d)", m)
}

type day int

const (
	Monday day = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (d day) string(code int) string {
	if Monday <= d && d <= Sunday {
		switch code {
		case codeWeekDay:
			return weekDayNames[d-1]
		case codeWeekDayLower:
			return weekDayLowerNames[d-1]
		case codeLongWeekDay:
			return longWeekDayNames[d-1]
		case codeLongWeekDayLower:
			return longWeekDayLowerNames[d-1]
		default:
			return fmt.Sprintf("Day(%d)", d)
		}

	}
	return fmt.Sprintf("Day(%d)", d)
}

// Now returns the current local time.
func Now() RTime {
	return RTime{time.Now()}
}

// Date is the envelope for the time.Date function and creates RTime for the passed date params.
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) RTime {
	t := time.Date(year, month, day, hour, min, sec, nsec, loc)
	return RTime{t}
}

// Unix is the envelope for the time.Unix function and creates RTime corresponding to the given Unix time,
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
func Unix(sec int64, nsec int64) RTime {
	return RTime{time.Unix(sec, nsec)}
}

// UnixMilli is the envelope for the time.UnixMilli and returns the local Time corresponding to the given Unix time,
// msec milliseconds since January 1, 1970 UTC.
func UnixMilli(msec int64) RTime {
	return RTime{time.UnixMilli(msec)}
}

// UnixMicro is the envelope for the time.UnixMicro and returns the local Time corresponding to the given Unix time,
// usec microseconds since January 1, 1970 UTC.
func UnixMicro(usec int64) RTime {
	return RTime{time.UnixMicro(usec)}
}

type RTime struct {
	time.Time
}

// Format returns a textual representation of the time value formatted according
// to the layout defined by the argument. See the documentation for the
// constant called Layout to see how to represent the layout format.
func (t RTime) Format(layout string) string {
	var b []byte
	max := len(layout) + 10
	b = make([]byte, 0, max)

	month := month(t.Month())
	day := day(t.Weekday())
	if day == 0 {
		day = Sunday
	}

	rLayout := layout
	for rLayout != "" {
		prefix, std, suffix := nextChunk(layout)

		if prefix != "" {
			b = append(b, prefix...)
		}
		if std == 0 {
			break
		}
		layout = suffix

		switch std {
		case codeLongMonth, codeLongMonthLower, codeMonth, codeMonthLower, codeLongMonthGenitive, codeLongMonthGenitiveLower:
			m := month.string(std)
			b = append(b, m...)
		case codeWeekDay, codeWeekDayLower, codeLongWeekDay, codeLongWeekDayLower:
			d := day.string(std)
			b = append(b, d...)
		}
	}

	layout = string(b)

	return t.Time.Format(layout)
}

func nextChunk(layout string) (prefix string, std int, suffix string) {
	rLayout := []rune(layout)
	for i, r := range rLayout {
		switch c := r; c {
		case 'Я': // Январь, Янв, Января
			if len(rLayout) >= i+3 && string(rLayout[i:i+3]) == "Янв" {
				if len(rLayout) >= i+6 && string(rLayout[i:i+6]) == "Январь" {
					return string(rLayout[0:i]), codeLongMonth, string(rLayout[i+6:])
				}
				if len(rLayout) >= i+6 && string(rLayout[i:i+6]) == "Января" {
					return string(rLayout[0:i]), codeLongMonthGenitive, string(rLayout[i+6:])
				}
				return string(rLayout[0:i]), codeMonth, string(rLayout[i+3:])
			}
		case 'я': // январь, янв, января
			if len(rLayout) >= i+3 && string(rLayout[i:i+3]) == "янв" {
				if len(rLayout) >= i+6 && string(rLayout[i:i+6]) == "январь" {
					return string(rLayout[0:i]), codeLongMonthLower, string(rLayout[i+6:])
				}
				if len(rLayout) >= i+6 && string(rLayout[i:i+6]) == "января" {
					return string(rLayout[0:i]), codeLongMonthGenitiveLower, string(rLayout[i+6:])
				}
				return string(rLayout[0:i]), codeMonthLower, string(rLayout[i+3:])
			}
		case 'П': // Понедельник, ПН
			if len(rLayout) >= i+2 && string(rLayout[i:i+2]) == "ПН" {
				return string(rLayout[0:i]), codeWeekDay, string(rLayout[i+2:])
			}
			if len(rLayout) >= i+11 && string(rLayout[i:i+11]) == "Понедельник" {
				return string(rLayout[0:i]), codeLongWeekDay, string(rLayout[i+11:])
			}
		case 'п': // понедельник, пн
			if len(rLayout) >= i+2 && string(rLayout[i:i+2]) == "пн" {
				return string(rLayout[0:i]), codeWeekDayLower, string(rLayout[i+2:])
			}
			if len(rLayout) >= i+11 && string(rLayout[i:i+11]) == "понедельник" {
				return string(rLayout[0:i]), codeLongWeekDayLower, string(rLayout[i+11:])
			}
		}

	}
	return layout, 0, ""
}

// Add is the envelope for the Time.Add and returns the time t+d.
func (t RTime) Add(d time.Duration) RTime {
	return RTime{t.Time.Add(d)}
}

// AddDate is the envelope for the Time.AddDate and returns the time corresponding to adding the
// given number of years, months, and days to t.
func (t RTime) AddDate(years int, months int, days int) RTime {
	return RTime{t.Time.AddDate(years, months, days)}
}

// UTC is the envelope for the Time.UTC and returns t with the location set to UTC.
func (t RTime) UTC() RTime {
	return RTime{t.Time.UTC()}
}

// Local is the envelope for the Time.Local and returns t with the location set to local time.
func (t RTime) Local() RTime {
	return RTime{t.Time.Local()}
}

// In is the envelope for the Time.In returns a copy of t representing the same time instant, but
// with the copy's location information set to loc for display purposes.
func (t RTime) In(loc *time.Location) RTime {
	return RTime{t.Time.In(loc)}
}

// Truncate is the envelope for the Time.Truncate
// and returns the result of rounding t down to a multiple of d (since the zero time).
func (t RTime) Truncate(d time.Duration) RTime {
	return RTime{t.Time.Truncate(d)}
}

// Round is the envelope for the Time.Round
// and returns the result of rounding t to the nearest multiple of d (since the zero time).
func (t RTime) Round(d time.Duration) RTime {
	return RTime{t.Time.Round(d)}
}
