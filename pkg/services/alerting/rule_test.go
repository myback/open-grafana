package alerting

import (
	"testing"

	"github.com/myback/open-grafana/pkg/components/simplejson"
	"github.com/myback/open-grafana/pkg/models"
	"github.com/myback/open-grafana/pkg/services/sqlstore"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type FakeCondition struct{}

func (f *FakeCondition) Eval(context *EvalContext) (*ConditionResult, error) {
	return &ConditionResult{}, nil
}

func TestAlertRuleFrequencyParsing(t *testing.T) {
	tcs := []struct {
		input  string
		err    error
		result int64
	}{
		{input: "10s", result: 10},
		{input: "10m", result: 600},
		{input: "1h", result: 3600},
		{input: "1d", result: 86400},
		{input: "1o", result: 1},
		{input: "0s", err: ErrFrequencyCannotBeZeroOrLess},
		{input: "0m", err: ErrFrequencyCannotBeZeroOrLess},
		{input: "0h", err: ErrFrequencyCannotBeZeroOrLess},
		{input: "0", err: ErrFrequencyCannotBeZeroOrLess},
		{input: "-1s", err: ErrFrequencyCouldNotBeParsed},
	}

	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			r, err := getTimeDurationStringToSeconds(tc.input)
			if tc.err == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.err.Error())
			}
			assert.Equal(t, tc.result, r)
		})
	}
}

func TestAlertRuleModel(t *testing.T) {
	sqlstore.InitTestDB(t)
	Convey("Testing alert rule", t, func() {
		RegisterCondition("test", func(model *simplejson.Json, index int) (Condition, error) {
			return &FakeCondition{}, nil
		})

		Convey("should return err for empty string", func() {
			_, err := getTimeDurationStringToSeconds("")
			So(err, ShouldNotBeNil)
		})

		Convey("can construct alert rule model", func() {
			firstNotification := models.CreateAlertNotificationCommand{Uid: "notifier1", OrgId: 1, Name: "1"}
			err := sqlstore.CreateAlertNotificationCommand(&firstNotification)
			So(err, ShouldBeNil)
			secondNotification := models.CreateAlertNotificationCommand{Uid: "notifier2", OrgId: 1, Name: "2"}
			err = sqlstore.CreateAlertNotificationCommand(&secondNotification)
			So(err, ShouldBeNil)

			Convey("with notification id and uid", func() {
				json := `
				{
					"name": "name2",
					"description": "desc2",
					"handler": 0,
					"noDataMode": "critical",
					"enabled": true,
					"frequency": "60s",
					"conditions": [
						{
							"type": "test",
							"prop": 123
						}
					],
					"notifications": [
						{"id": 1},
						{"uid": "notifier2"}
					]
				}
				`

				alertJSON, jsonErr := simplejson.NewJson([]byte(json))
				So(jsonErr, ShouldBeNil)

				alert := &models.Alert{
					Id:          1,
					OrgId:       1,
					DashboardId: 1,
					PanelId:     1,

					Settings: alertJSON,
				}

				alertRule, err := NewRuleFromDBAlert(alert, false)
				So(err, ShouldBeNil)

				So(len(alertRule.Conditions), ShouldEqual, 1)
				So(len(alertRule.Notifications), ShouldEqual, 2)

				Convey("Can read Id and Uid notifications (translate Id to Uid)", func() {
					So(alertRule.Notifications, ShouldContain, "notifier2")
					So(alertRule.Notifications, ShouldContain, "notifier1")
				})
			})
		})

		Convey("with non existing notification id", func() {
			json := `
				{
					"name": "name3",
					"description": "desc3",
					"handler": 0,
					"noDataMode": "critical",
					"enabled": true,
					"frequency": "60s",
					"conditions": [{"type": "test", "prop": 123 }],
					"notifications": [
						{"id": 999},
						{"uid": "notifier2"}
					]
				}
				`

			alertJSON, jsonErr := simplejson.NewJson([]byte(json))
			So(jsonErr, ShouldBeNil)

			alert := &models.Alert{
				Id:          1,
				OrgId:       1,
				DashboardId: 1,
				PanelId:     1,

				Settings: alertJSON,
			}

			alertRule, err := NewRuleFromDBAlert(alert, false)
			Convey("swallows the error", func() {
				So(err, ShouldBeNil)
				So(alertRule.Notifications, ShouldNotContain, "999")
				So(alertRule.Notifications, ShouldContain, "notifier2")
			})
		})

		Convey("can construct alert rule model with invalid frequency", func() {
			json := `
			{
				"name": "name2",
				"description": "desc2",
				"noDataMode": "critical",
				"enabled": true,
				"frequency": "0s",
				"conditions": [ { "type": "test", "prop": 123 } ],
				"notifications": []
			}`

			alertJSON, jsonErr := simplejson.NewJson([]byte(json))
			So(jsonErr, ShouldBeNil)

			alert := &models.Alert{
				Id:          1,
				OrgId:       1,
				DashboardId: 1,
				PanelId:     1,
				Frequency:   0,

				Settings: alertJSON,
			}

			alertRule, err := NewRuleFromDBAlert(alert, false)
			So(err, ShouldBeNil)
			So(alertRule.Frequency, ShouldEqual, 60)
		})

		Convey("raise error in case of missing notification id and uid", func() {
			json := `
			{
				"name": "name2",
				"description": "desc2",
				"noDataMode": "critical",
				"enabled": true,
				"frequency": "60s",
				"conditions": [
					{
						"type": "test",
						"prop": 123
					}
				],
				"notifications": [
					{"not_id_uid": "1134"}
				]
			}
			`

			alertJSON, jsonErr := simplejson.NewJson([]byte(json))
			So(jsonErr, ShouldBeNil)

			alert := &models.Alert{
				Id:          1,
				OrgId:       1,
				DashboardId: 1,
				PanelId:     1,
				Frequency:   0,

				Settings: alertJSON,
			}

			_, err := NewRuleFromDBAlert(alert, false)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "alert validation error: Neither id nor uid is specified in 'notifications' block, type assertion to string failed AlertId: 1 PanelId: 1 DashboardId: 1")
		})
	})
}
