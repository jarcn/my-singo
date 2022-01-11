package test

import (
	"testing"
	"time"

	af "github.com/carpenterscode/appsflyer-go"
)

func TestAf(t *testing.T) {
	tracker := af.NewTracker()
	tracker.SetConfig("appsflyer.json")
	startTrial(tracker)
	subscribe(tracker)
	cancelSubscription(tracker)
}

func startTrial(tracker af.Tracker) {

	// User starts a trial
	evt := af.NewEvent("1111111111111-1111111", af.IOS).
		SetName(af.StartTrial).
		SetAdvertisingID("AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA").
		SetDeviceIP("1.2.3.4").
		SetPrice(59.99, "USD").
		SetDateValue("expiry", <-time.After(time.Minute*30)).
		SetEventTime(time.Now())

	if err := tracker.Send(evt); err != nil {
		panic(err)
	}
}

func subscribe(tracker af.Tracker) {
	// User ends trial and pays for first subscription period
	evt := af.NewEvent("1111111111111-1111111", af.IOS).
		SetName(af.Subscribe).
		SetAdvertisingID("AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA").
		SetDeviceIP("1.2.3.4").
		SetRevenue(59.99, "USD").
		SetDateValue("expiry", <-time.After(time.Minute*30))
	if err := tracker.Send(evt); err != nil {
		panic(err)
	}
}

func cancelSubscription(tracker af.Tracker) {
	// User cancels a subscription
	evt := af.NewEvent("1111111111111-1111111", af.Android)
	evt.SetName(af.EventName("cancel_subscription"))
	evt.SetRevenue(-59.99, "USD")
	if err := tracker.Send(evt); err != nil {
		panic(err)
	}
}
