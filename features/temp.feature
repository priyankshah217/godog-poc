Feature: Calculator

  Scenario Outline: Arithmetic Operations
    When I perform below <Operations> on below numbers
      | Op1 | Op2 |
      | 6   | 2   |
    Then the <Result> should be below
    Examples:
      | Operations | Result |
      | "Add"      | 8      |
      | "Min"      | 4      |
      | "Mul"      | 12     |
      | "Div"      | 3      |