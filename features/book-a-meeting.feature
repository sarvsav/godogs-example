Feature: Book a meeting
    In order to be able to book a meeting
    As a bank customer is able to book a meeting with a bank employee

    Scenario: Book a successful meeting
        Given I am a bank customer
        When I book a meeting
        Then I should be told "booking is successful"
