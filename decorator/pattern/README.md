# Decorator Pattern in Go

This example showcases the decorator pattern in Go, particularly in the context of text formatting. The decorator pattern allows dynamic and flexible behavior extension for objects without altering their structure.

## Files

### `decorator.go`

This file demonstrates the decorator pattern by defining a set of text formatters. It includes:

- `Formatter` interface: Common methods for formatting text.
- Concrete formatters (`PlainText`, `BoldText`, `ItalicText`, `UnderlineText`): Implement the `Formatter` interface.
- Decorators (`BoldText`, `ItalicText`, `UnderlineText`): Wrap other formatters to add specific formatting styles.
- The `main` function showcases different formatting combinations using the decorator pattern.

### `simple.go`

This file provides an alternative approach to text formatting without using the decorator pattern. It illustrates:

- Formatters (`PlainText`, `BoldText`, `ItalicText`, `UnderlineText`): Implement the `Formatter` interface.
- The `main` function demonstrates text formatting without the decorator pattern, relying on method chaining.

## Why Decorator Pattern is Useful

The decorator pattern offers several advantages:

1. **Dynamic Composition:** Decorators allow you to dynamically compose and combine behaviors at runtime. This flexibility contrasts with static inheritance and promotes a more adaptable design.

2. **Open/Closed Principle:** The decorator pattern adheres to the Open/Closed Principle, enabling the extension of behavior without modifying existing code. New decorators can be added without altering the base components.

3. **Single Responsibility Principle:** Each decorator class has a single responsibility, making the code more modular and easier to maintain. This separation of concerns enhances code readability and scalability.

4. **Code Reusability:** Decorators can be reused across different components, providing a modular and reusable way to add functionalities. This reusability minimizes code duplication.

5. **Easy Removal of Decorations:** Decorators can be easily added or removed, allowing for fine-grained control over the behavior of objects. This promotes a more maintainable and adaptable codebase.

By embracing the decorator pattern, developers can create extensible and maintainable systems, where the responsibilities of structs are clearly defined, and new features can be introduced with minimal impact on existing code.