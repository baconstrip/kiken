// Package message defines all messages sent between the client and server.
// Types in the this package should rely on primitive types and types in this
// directory only.
package message

// BoardOverview defines a message that the server sends to clients
// representing the board without exposing question information.
type BoardOverview struct {
    Round string
    Categories []*CategoryOverview
}

// CategoryOverview defines a message that the server sends to clients
// representing a category without exposing question information.
type CategoryOverview struct {
    Name string
    Questions []*QuestionHidden
}

// QuestionHidden defines a message that the server sends to clients
// representing a question without exposing the prompt or answer.
type QuestionHidden struct {
    Value int
    Played bool
}
