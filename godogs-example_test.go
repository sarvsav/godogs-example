package main

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

// godogsCtxKey is the key used to store the available godogs in the context.Context
type godogsCtxKey struct{}

func iAmABankCustomer(ctx context.Context, isBankCustomer string) (context.Context, error) {
	// isBankCustomer is a boolean value that verifies if the customer is a bank customer or not
	return context.WithValue(ctx, godogsCtxKey{}, isBankCustomer), nil
}

func iBookAMeeting(ctx context.Context) (context.Context, error) {
	// Here we check for conditions the day and time of the meeting
	// and availablity of the slots
	weekday := time.Now().Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return ctx, errors.New("Weekend meetings are not allowed")
	}

	// Just for demo, we are setting the available slots to 10
	availableSlots := 10
	if availableSlots <= 0 {
		return ctx, errors.New("No slots available")
	}

	return ctx, nil
}

func iShouldBeTold(ctx context.Context) error {
	userIsBankCustomer := ctx.Value(godogsCtxKey{}).(string)
	if userIsBankCustomer == "customer" {
		return nil
	}
	return errors.New("User is not a bank customer")
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Step(`^I am a bank (\w+)$`, iAmABankCustomer)
	sc.Step(`^I book a meeting$`, iBookAMeeting)
	sc.Step(`^I should be told "([^"]*)"$`, iShouldBeTold)
}
