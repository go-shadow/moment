package moment

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMomentParser(t *testing.T) {
	parser := new(MomentParser)

	Convey("Given moment month formats", t, func() {
		Convey("It should generate the correct golang formats for months", func() {
			So(parser.Convert("M"), ShouldEqual, "1")
			So(parser.Convert("Mo"), ShouldEqual, "1<stdOrdinal>")
			So(parser.Convert("MM"), ShouldEqual, "01")
			So(parser.Convert("MMM"), ShouldEqual, "Jan")
			So(parser.Convert("MMMM"), ShouldEqual, "January")
		})
	})

	Convey("Given moment day of month formats", t, func() {
		Convey("It should generate the correct golang formats for days", func() {
			So(parser.Convert("D"), ShouldEqual, "2")
			So(parser.Convert("Do"), ShouldEqual, "2<stdOrdinal>")
			So(parser.Convert("DD"), ShouldEqual, "02")
		})
	})

	Convey("Given moment day of year formats", t, func() {
		Convey("It should generate the correct golang formats for days", func() {
			So(parser.Convert("DDD"), ShouldEqual, "<stdDayOfYear>")
			So(parser.Convert("DDDo"), ShouldEqual, "<stdDayOfYear><stdOrdinal>")
			So(parser.Convert("DDDD"), ShouldEqual, "<stdDayOfYearZero>")
		})
	})

	Convey("Given moment day of week formats", t, func() {
		Convey("It should generate the correct golang formats for days", func() {
			So(parser.Convert("d"), ShouldEqual, "<stdDayOfWeek>")
			So(parser.Convert("do"), ShouldEqual, "<stdDayOfWeek><stdOrdinal>")
			// So(parser.Convert("dd"), ShouldEqual, "Mo")
			So(parser.Convert("ddd"), ShouldEqual, "Mon")
			So(parser.Convert("dddd"), ShouldEqual, "Monday")

			// Day of week locale
			So(parser.Convert("e"), ShouldEqual, "<stdDayOfWeek>")
			So(parser.Convert("E"), ShouldEqual, "<stdDayOfWeekISO>")
		})
	})

	Convey("Given moment week of year formats", t, func() {
		Convey("It should generate the correct golang formats for week of year", func() {
			So(parser.Convert("w"), ShouldEqual, "<stdWeekOfYear>")
			So(parser.Convert("wo"), ShouldEqual, "<stdWeekOfYear><stdOrdinal>")
			// So(parser.Convert("ww"), ShouldEqual, "@todo")
			So(parser.Convert("W"), ShouldEqual, "<stdWeekOfYear>")
			So(parser.Convert("Wo"), ShouldEqual, "<stdWeekOfYear><stdOrdinal>")
			// So(parser.Convert("WW"), ShouldEqual, "@todo")
		})
	})

	Convey("Given moment year formats", t, func() {
		Convey("It should generate the correct golang formats for years", func() {
			So(parser.Convert("YY"), ShouldEqual, "06")
			So(parser.Convert("YYYY"), ShouldEqual, "2006")
		})
	})

	Convey("Given moment week year formats", t, func() {
		Convey("It should generate the correct golang formats for week years", func() {
			// So(parser.Convert("gg"), ShouldEqual, "@todo")
			// So(parser.Convert("gggg"), ShouldEqual, "@todo")

			// ISO
			// So(parser.Convert("GG"), ShouldEqual, "@todo")
			// So(parser.Convert("GGGG"), ShouldEqual, "@todo")
		})
	})

	Convey("Given moment hour formats", t, func() {
		Convey("It should generate the correct golang formats for hours", func() {
			So(parser.Convert("H"), ShouldEqual, "<stdHourNoZero>")
			So(parser.Convert("HH"), ShouldEqual, "15")

			So(parser.Convert("h"), ShouldEqual, "3")
			So(parser.Convert("hh"), ShouldEqual, "03")
		})
	})

	Convey("Given moment minute formats", t, func() {
		Convey("It should generate the correct golang formats for minutes", func() {
			So(parser.Convert("m"), ShouldEqual, "4")
			So(parser.Convert("mm"), ShouldEqual, "04")
		})
	})

	Convey("Given moment second formats", t, func() {
		Convey("It should generate the correct golang formats for seconds", func() {
			So(parser.Convert("s"), ShouldEqual, "5")
			So(parser.Convert("ss"), ShouldEqual, "05")
		})
	})

	Convey("Given moment localized formats", t, func() {
		Convey("It should generate the correct golang formats", func() {
			So(parser.Convert("LT"), ShouldEqual, "3:04 PM")
			So(parser.Convert("L"), ShouldEqual, "01/02/2006")
			So(parser.Convert("l"), ShouldEqual, "1/2/2006")
			So(parser.Convert("LL"), ShouldEqual, "January 2 2006")
			So(parser.Convert("ll"), ShouldEqual, "Jan 2 2006")
			So(parser.Convert("LLL"), ShouldEqual, "January 2 2006 3:04 PM")
			So(parser.Convert("lll"), ShouldEqual, "Jan 2 2006 3:04 PM")
			So(parser.Convert("LLLL"), ShouldEqual, "Monday, January 2 2006 3:04 PM")
			So(parser.Convert("llll"), ShouldEqual, "Mon, Jan 2 2006 3:04 PM")
		})
	})

	Convey("Given moment timezone formats", t, func() {
		Convey("It should generate the correct golang timezone formats", func() {
			So(parser.Convert("z"), ShouldEqual, "MST")
			So(parser.Convert("zz"), ShouldEqual, "MST")
			So(parser.Convert("Z"), ShouldEqual, "Z07:00")
			So(parser.Convert("ZZ"), ShouldEqual, "-0700")
		})
	})

	Convey("Given moment misc formats", t, func() {
		Convey("It should generate the correct golang formats", func() {
			So(parser.Convert("Q"), ShouldEqual, "<stdQuarter>")
			So(parser.Convert("A"), ShouldEqual, "PM")
			So(parser.Convert("a"), ShouldEqual, "pm")

			So(parser.Convert("X"), ShouldEqual, "<stdUnix>")
		})
	})

	Convey("Given a moment format", t, func() {
		Convey("It should generate the correct golang time format", func() {
			So(parser.Convert("YYYY-MM-DD"), ShouldEqual, "2006-01-02")
			So(parser.Convert("YYYY/MM/DD"), ShouldEqual, "2006/01/02")
		})
	})
}
