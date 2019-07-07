@all
Feature: eat godogs

  @logic
  Scenario: Eat 5 out of 12
    Given there are 12 godogs
    When I eat 5
    Then there should be 7 remaining

  @localhost
  Scenario: Open localhost valid welcome message
    Given user open browser
    When user access url with name "mpermperpisang"
    Then user must get window title Welcome "mpermperpisang"

  @localhost
  Scenario: Open localhost invalid welcome message
    Given user open browser
    When user access url with name "ferawati"
    Then user must get window title Welcome "mpermperpisang"
