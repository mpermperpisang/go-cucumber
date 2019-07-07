@all
Feature: Bike - Landing Page

  @team
  Scenario: Fill bike location suggestion form
    Given user access bike url
    When user fill suggestion form
    Then user must get success message "Usul kamu sudah diterima!"
