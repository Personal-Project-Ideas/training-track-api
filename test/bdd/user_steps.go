package bdd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/PratesJr/training-track-api/internal/bootstrap"
	"github.com/cucumber/godog"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	FullName string `gorm:"not null"`
	Age      int    `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

var (
	db             *gorm.DB
	currentUser    User
	responseStatus int
	expectedStatus int // Track expected status for outline
	app            *fiber.App
)

func resetDB() error {
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open in-memory sqlite: %w", err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}
	return nil
}

func iHaveAValidUserPayload() error {
	currentUser = User{
		FullName: "Alice Smith",
		Age:      30,
		Email:    "alice@example.com",
		Password: "pass1234",
	}
	return nil
}

func iHaveAUserPayloadMissingTheEmail() error {
	currentUser = User{
		FullName: "Bob",
		Age:      25,
		Email:    "",
		Password: "pass5678",
	}
	return nil
}

func iHaveAUserPayloadWithAnEmailThatAlreadyExists() error {
	// First, create a user
	user := User{
		FullName: "Carol Jones",
		Age:      25,
		Email:    "carol@example.com",
		Password: "pass9999",
	}
	db.Create(&user)
	// Now, try to create again with same email
	currentUser = user
	return nil
}

func iHaveAUserPayloadWith(fullName string, ageStr string, email string, password string) error {
	age := 0
	if strings.TrimSpace(ageStr) != "" {
		_, err := fmt.Sscanf(ageStr, "%d", &age)
		if err != nil {
			return fmt.Errorf("invalid age: %w", err)
		}
	}
	currentUser = User{
		FullName: fullName,
		Age:      age,
		Email:    email,
		Password: password,
	}
	// Pre-insert for duplicate email scenario
	if expectedStatus == 409 {
		preUser := User{
			FullName: "PreExisting",
			Age:      99,
			Email:    email,
			Password: "prepass",
		}
		db.Create(&preUser)
	}
	return nil
}

func iHaveAUserPayloadWithFullNameAgeEmailPassword(fullName string, age int, email, password string) error {
	currentUser = User{
		FullName: fullName,
		Age:      age,
		Email:    email,
		Password: password,
	}
	return nil
}

func iHaveAUserPayloadWithFullNameAgeEmailPasswordMissingAge(fullName, email, password string) error {
	currentUser = User{
		FullName: fullName,
		Age:      0, // missing age
		Email:    email,
		Password: password,
	}
	return nil
}

func iSendARequestToCreateTheUser() error {
	// Prepare JSON payload
	payload := map[string]interface{}{
		"full_name": currentUser.FullName,
		"age":       currentUser.Age,
		"email":     currentUser.Email,
		"password":  currentUser.Password,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	// Simulate HTTP POST request
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(body)))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)
	if err != nil {
		return err
	}
	responseStatus = resp.StatusCode
	return nil
}

func theResponseStatusShouldBe(expected int) error {
	expectedStatus = expected
	if responseStatus != expected {
		return fmt.Errorf("expected status %d, got %d", expected, responseStatus)
	}
	return nil
}

func aUserWithEmailAlreadyExists(email string) error {
	preUser := User{
		FullName: "PreExisting",
		Age:      99,
		Email:    email,
		Password: "prepass",
	}
	result := db.Create(&preUser)
	if result.Error != nil {
		return fmt.Errorf("failed to pre-insert user: %w", result.Error)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		if err := resetDB(); err != nil {
			return ctx, err
		}
		currentUser = User{}
		responseStatus = 0
		appContainer := bootstrap.Bootstrap(db)
		app = appContainer.App
		return ctx, nil
	})
	ctx.Step("I have a valid user payload", iHaveAValidUserPayload)
	ctx.Step("I have a user payload missing the email", iHaveAUserPayloadMissingTheEmail)
	ctx.Step("I have a user payload with an email that already exists", iHaveAUserPayloadWithAnEmailThatAlreadyExists)
	ctx.Step(`^a user with email "([^"]*)" already exists$`, aUserWithEmailAlreadyExists)
	ctx.Step(`^I have a user payload with fullName "([^"]*)", age (\d+), email "([^"]*)", password "([^"]*)"$`, iHaveAUserPayloadWithFullNameAgeEmailPassword)
	ctx.Step(`^I have a user payload with fullName "([^"]*)", age , email "([^"]*)", password "([^"]*)"$`, iHaveAUserPayloadWithFullNameAgeEmailPasswordMissingAge)
	ctx.Step("I have a user payload with fullName \"<fullName>\", age <age>, email \"<email>\", password \"<password>\"", iHaveAUserPayloadWith)
	ctx.Step("I send a request to create the user", iSendARequestToCreateTheUser)
	ctx.Step("the response status should be <status>", theResponseStatusShouldBe)
	ctx.Step("the response status should be {int}", theResponseStatusShouldBe)
	ctx.Step(`^the response status should be (\d+)$`, theResponseStatusShouldBe)
}
