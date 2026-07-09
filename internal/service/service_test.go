package service

import (
	"reflect"
	"testing"

	"github.com/jabrail059/weather-dashboard/internal/models"
)

func TestNormalizeCityName(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "simple test",
			in:   "moScOw",
			want: "Moscow",
		},
		{
			name: "trim spaces",
			in:   " MoscoW     ",
			want: "Moscow",
		},
		{
			name: "two words",
			in:   "saint petersburg",
			want: "Saint Petersburg",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := normalizeCityName(test.in)
			if got != test.want {
				t.Errorf("normalizeCityName(%q) = %q, want = %q", test.in, got, test.want)
			}
		})
	}
}

func TestBuildDayWeather(t *testing.T) {
	t.Run("builds day weather with valid forecast", func(t *testing.T) {
		forecast := models.Daily{
			Time:           []string{"2026-05-12"},
			TemperatureMax: []float64{22.3},
			TemperatureMin: []float64{10.6},
			WeatherCode:    []int{0},
			Sunrise:        []string{"2026-05-12T05:03"},
			Sunset:         []string{"2026-05-12T19:38"},
		}

		got, err := BuildDayWeather(&forecast)
		if err != nil {
			t.Fatalf("BuildDayWeather returned error: %v", err)
		}

		if len(got) != 1 {
			t.Fatalf("len(got) = %v, want = 1", len(got))
		}

		day := got[0]

		if day.Time != "May 12, 2026" {
			t.Errorf("Time = %v, want = May 12, 2026", day.Time)
		}

		if day.TemperatureMax != 22 {
			t.Errorf("TemperatureMax = %v, want = 22", day.TemperatureMax)
		}

		if day.TemperatureMin != 11 {
			t.Errorf("TemperatureMin = %v, want = 11", day.TemperatureMin)
		}

		if day.WeatherDescription != "Ясно" {
			t.Errorf("WeatherDescription = %v, want = Ясно", day.WeatherDescription)
		}

		if day.ImageSource != "/static/icons/clear-day.svg" {
			t.Errorf("ImageSource = %v, want = /static/icons/clear-day.svg", day.ImageSource)
		}

		if day.Sunrise != "05:03" {
			t.Errorf("Sunrise = %v, want = 05:03", day.Sunrise)
		}

		if day.Sunset != "19:38" {
			t.Errorf("Sunset = %v, want = 19:38", day.Sunset)
		}
	})
	t.Run("returns error for invalid date", func(t *testing.T) {
		forecast := models.Daily{
			Time:           []string{"12-05-2026"},
			TemperatureMax: []float64{22.3},
			TemperatureMin: []float64{10.6},
			WeatherCode:    []int{0},
			Sunrise:        []string{"2026-05-12T05:03"},
			Sunset:         []string{"2026-05-12T19:38"},
		}

		_, err := BuildDayWeather(&forecast)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestBuildOneDayForecast(t *testing.T) {
	t.Run("filters one day forecast by date", func(t *testing.T) {
		forecast := models.Hourly{
			Time: []string{
				"2026-05-12T00:00",
				"2026-05-12T01:00",
				"2026-05-13T00:00",
			},
			Temperature: []float64{12.3, 12.6, 13.1},
			WeatherCode: []int{0, 0, 1},
		}

		got, err := BuildOneDayForecast(&forecast, "2026-05-12")

		if err != nil {
			t.Fatalf("BuildOneDayForecast returned error: %v", err)
		}

		if len(got.Time) != 2 {
			t.Fatalf("len(got.Time) = %v, want = 2", len(got.Time))
		}

		want := &models.Hourly{
			Time: []string{
				"2026-05-12T00:00",
				"2026-05-12T01:00",
			},
			Temperature: []float64{12.3, 12.6},
			WeatherCode: []int{0, 0},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("returns error when date is not found", func(t *testing.T) {
		forecast := models.Hourly{
			Time: []string{
				"2026-05-12T00:00",
				"2026-05-12T01:00",
			},
			Temperature: []float64{12.3, 12.6},
			WeatherCode: []int{0, 1},
		}

		_, err := BuildOneDayForecast(&forecast, "2026-05-13")
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

func TestBuildHourlyWeather(t *testing.T) {
	t.Run("builds hourly weather with valid forecast", func(t *testing.T) {
		forecast := models.Hourly{
			Time: []string{
				"2026-05-12T05:00",
				"2026-05-12T06:00",
				"2026-05-12T19:00",
				"2026-05-12T20:00",
			},
			Temperature: []float64{12.3, 12.6, 19.1, 18.7},
			WeatherCode: []int{0, 0, 1, 2},
		}

		got, err := BuildHourlyWeather(&forecast, "05:03", "19:38")

		if err != nil {
			t.Fatalf("BuildHourlyWeather returned error: %v", err)
		}

		if len(got) != 4 {
			t.Fatalf("len(got) = %v, want = 4", len(got))
		}

		if got[0].Time != "05:00" ||
			got[0].Temperature != 12 ||
			got[0].ImageSource != "/static/icons/clear-night.svg" {
			t.Errorf("[0]Time = %v, want = 05:00, \n[0]Temperature = %v, want = 12, \n[0]ImageSource = %v, want = /static/icons/clear-night.svg",
				got[0].Time, got[0].Temperature, got[0].ImageSource)
		}

		if got[1].Time != "06:00" ||
			got[1].Temperature != 13 ||
			got[1].ImageSource != "/static/icons/clear-day.svg" {
			t.Errorf("[1]Time = %v, want = 06:00, \n[1]Temperature = %v, want = 13, \n[1]ImageSource = %v, want = /static/icons/clear-day.svg",
				got[1].Time, got[1].Temperature, got[1].ImageSource)
		}

		if got[2].Time != "19:00" ||
			got[2].Temperature != 19 ||
			got[2].ImageSource != "/static/icons/mostly-clear-day.svg" {
			t.Errorf("[2]Time = %v, want = 19:00, \n[2]Temperature = %v, want = 19, \n[2]ImageSource = %v, want = /static/icons/mostly-clear-day.svg",
				got[2].Time, got[2].Temperature, got[2].ImageSource)
		}

		if got[3].Time != "20:00" ||
			got[3].Temperature != 19 ||
			got[3].ImageSource != "/static/icons/partly-cloudy-night.svg" {
			t.Errorf("[3]Time = %v, want = 20:00, \n[3]Temperature = %v, want = 19, \n[3]ImageSource = %v, want = /static/icons/partly-cloudy-night.svg",
				got[3].Time, got[3].Temperature, got[3].ImageSource)
		}
	})
	t.Run("returns error when hourly time format is invalid", func(t *testing.T) {
		forecast := models.Hourly{
			Time: []string{
				"2026-05-12",
				"2026-05-12",
			},
			Temperature: []float64{12.3, 12.6},
			WeatherCode: []int{0, 1},
		}

		_, err := BuildHourlyWeather(&forecast, "05:03", "19:38")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestBuildCurrentWeather(t *testing.T) {
	t.Run("builds current weather with valid forecast", func(t *testing.T) {
		forecast := models.Current{
			Temperature:  19.2,
			ApparentTemp: 20.6,
			WindSpeed:    13.6,
			WeatherCode:  0,
			IsDay:        1,
		}

		got := BuildCurrentWeather(&forecast)

		if got.Temperature != 19 {
			t.Errorf("Temperature = %v, want = 19", got.Temperature)
		}

		if got.ApparentTemp != 21 {
			t.Errorf("ApparentTemp = %v, want = 21", got.ApparentTemp)
		}

		if got.ImageSource != "/static/icons/clear-day.svg" {
			t.Errorf("ImageSource = %v, want = /static/icons/clear-day.svg", got.ImageSource)
		}
	})
}
