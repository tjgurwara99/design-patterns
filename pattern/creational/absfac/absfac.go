package absfac

// AbstractFactory is the actual abstract factory that we use to create the
// concrete services.
type AbstractFactory func(db *DatabaseDriverSub) ProductCreator
