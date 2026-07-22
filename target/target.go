package target

// Target interface defines the methods of target
type Target interface {
	// Connect connects to the target databse
	Connect() error
}
