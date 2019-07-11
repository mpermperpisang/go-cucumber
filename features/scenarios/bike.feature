@all-1
Feature: Bike - Landing Page

  @test-1
  Scenario: Fill bike location suggestion form valid message
    Given user access bike url
    When user fill suggestion form
    Then user must not see window "Usulkan Lokasi BukaBike Selanjutnya"
