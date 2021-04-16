package pkg

import (
	"testing"
	"time"
)

func TestDiff(t *testing.T) {
	type args struct {
		a time.Time
		b time.Time
	}

	startDate, _ := time.Parse(`2006-01-02`, `2021-01-30`)
	endDate, _ := time.Parse(`2006-01-02`, `2021-03-30`)
	endFebDate, _ := time.Parse(`2006-01-02`, `2021-02-28`)
	endNextYearFebDate, _ := time.Parse(`2006-01-02`, `2022-02-28`)

	tests := []struct {
		name           string
		args           args
		wantYears      int
		wantMonths     int
		wantDays       int
		wantWeeks      int
		wantFortNights int
		wantQuarters   int
	}{
		{
			name: "jan_30_to_march_30",
			args: args{
				a: startDate,
				b: endDate,
			},
			wantYears:      1,
			wantMonths:     2,
			wantQuarters:   1,
			wantDays:       60,
			wantWeeks:      9,
			wantFortNights: 5,
		},
		{
			name: "jan_30_to_feb_28",
			args: args{
				a: startDate,
				b: endFebDate,
			},
			wantYears:      1,
			wantMonths:     1,
			wantQuarters:   1,
			wantDays:       30,
			wantWeeks:      5,
			wantFortNights: 3,
		},
		{
			name: "jan_30_to_next_year_feb_28",
			args: args{
				a: startDate,
				b: endNextYearFebDate,
			},
			wantYears:      2,
			wantMonths:     13,
			wantQuarters:   4,
			wantDays:       395,
			wantWeeks:      57,
			wantFortNights: 29,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			gotYears, gotMonths, gotDays, gotWeeks, gotFortNights, gotQuartes := Diff(tt.args.a, tt.args.b)
			if gotYears != tt.wantYears {
				t.Errorf("Diff() gotYears = %v, want %v", gotYears, tt.wantYears)
			}
			if gotMonths != tt.wantMonths {
				t.Errorf("Diff() gotMonths = %v, want %v", gotMonths, tt.wantMonths)
			}
			if gotDays != tt.wantDays {
				t.Errorf("Diff() gotDays = %v, want %v", gotDays, tt.wantDays)
			}
			if gotWeeks != tt.wantWeeks {
				t.Errorf("Diff() gotWeeks = %v, want %v", gotWeeks, tt.wantWeeks)
			}
			if gotFortNights != tt.wantFortNights {
				t.Errorf("Diff() gotFortNights = %v, want %v", gotFortNights, tt.wantFortNights)
			}
			if gotQuartes != tt.wantQuarters {
				t.Errorf("Diff() gotQuarters = %v, want %v", gotQuartes, tt.wantQuarters)
			}
		})
	}
}

func TestRestructuredDate(t *testing.T) {
	type args struct {
		t time.Time
	}

	date, _ := time.Parse(`2006-01-02`, `2021-01-30`)
	expected := 20210130

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "format_date",
			args: args{t: date},
			want: expected,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			if got := RestructuredDate(tt.args.t); got != tt.want {
				t.Errorf("RestructuredDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
