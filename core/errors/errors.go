package errors

type NotFound struct {
	description string
}

func NewNotFound(description string) NotFound {
	return NotFound{description: description}
}

func (notFound NotFound) Error() string {
	return notFound.description
}

type Invalid struct {
	description string
}

func NewInvalid(description string) Invalid {
	return Invalid{description: description}
}

func (invalid Invalid) Error() string {
	return invalid.description
}

// TODO: need a better name

type Domain struct {
	description string
}

func NewDomain(description string) Domain {
	return Domain{description: description}
}

func (domain Domain) Error() string {
	return domain.description
}
