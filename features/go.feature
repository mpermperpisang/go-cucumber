Feature: eat godogs
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

  @logic
  Scenario: Eat 5 out of 12
    Given there are 12 godogs
    When I eat 5
    Then there should be 7 remaining

  @localhost
  Scenario: Open Bukalapak
    Given user open browser
    When user access url with name "ferawati"
    Then user get window title of his name

  @localhost
  Scenario: Open Bukalapak
    Given user open browser
    When user access url with name "mpermperpisang"
    Then user get window title of his name