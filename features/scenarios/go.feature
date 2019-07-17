@all-2
Feature: GOlang - Learn Godog

  @test-a
  Scenario: Open website and search
    Given user has url
    When user access the url
    Then user must directed to correct url

  @test-b
  Scenario: Eat 5 out of 12
    Given there are 12 godogs
    When I eat 5
    Then there should be 7 remaining

  @test-c
  Scenario: Open localhost valid welcome message
    Given user with name "mpermperpisang"
    When user access url with given name
    Then user must get window title Welcome "mpermperpisang"

  @test-d
  Scenario: Open localhost invalid welcome message
    Given user with name "mpermperpisang"
    When user access url with given name
    Then user must get window title Welcome "ferawati"
