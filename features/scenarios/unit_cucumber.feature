@all-5
Feature: GO Language - Unit Testing From Cucumber Point of View

  @hello-user
  Scenario: Hello User
    Given user has a name "Mper"
    When neighbor meet user
    Then neighbor say "Hello Mper!"