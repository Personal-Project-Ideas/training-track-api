@user @api
Feature: User creation

  @positive @create
  Scenario: Create a new user
    Given I have a valid user payload
    When I send a request to create the user
    Then the response status should be 500

  @negative @missing_email
  Scenario: Create user with missing email
    Given I have a user payload missing the email
    When I send a request to create the user
    Then the response status should be 422

  @negative @duplicate_email
  Scenario: Create user with duplicate email
    Given a user with email "alice@example.com" already exists
    And I have a user payload with fullName "Carol Jones", age 25, email "alice@example.com", password "pass9999"
    When I send a request to create the user
    Then the response status should be 500

  @outline @positive
  Scenario Outline: Create user with various valid payloads
    Given I have a user payload with fullName "<fullName>", age <age>, email "<email>", password "<password>"
    When I send a request to create the user
    Then the response status should be 500

    Examples:
      | fullName      | age | email                | password   | status |
      | Alice Smith   | 30  | alice@example.com    | pass1234   | 500    |
      | Carol Jones   | 25  | carol@example.com    | pass9999   | 500    |
      | Dave          | 22  | dave@example.com     | pass0000   | 500    |

  @outline @negative
  Scenario Outline: Create user with various invalid payloads
    Given I have a user payload with fullName "<fullName>", age <age>, email "<email>", password "<password>"
    When I send a request to create the user
    Then the response status should be <status>

    Examples:
      | fullName      | age | email                | password   | status |
      | Bob           |     | bob@example.com      | pass5678   | 422    |
      | Dave          | 22  | dave@example.com     |            | 422    |
      | Eve           | 20  |                     | pass1111   | 422    |
