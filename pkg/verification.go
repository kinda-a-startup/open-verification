package pkg

import (
	"crypto/sha256"
	"time"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Identity struct {
	age       int
	address   string
	name      string
	dob       time.Time
	expiresAt time.Time
	issuedAt  time.Time
	country   string
	gender    Gender
	id        string
	image     []byte
}

func (i *Identity) isId(id string) bool {
	return i.id == id
}

func (i *Identity) isName(name string, strict bool) bool {
	// need to account for middle name not being included in input
	return i.name == name
}

func (i *Identity) isOver(age int) bool {
	return i.age >= age
}

func (i *Identity) isGender(gender Gender) bool {
	return i.gender == gender
}

func (i *Identity) isAddress(address string) bool {
	// there needs to be some kind of fuzzy matching here
	return i.address == address
}

func (i *Identity) isCountry(country string) bool {
	return i.country == country
}

func (i *Identity) isDobMonth(month string) bool {
	parsedMonth, err := time.Parse("January", month)
	if err != nil {
		return false
	}
	return i.dob.Month() == parsedMonth.Month()
}

func (i *Identity) isDobYear(year string) bool {
	parsedYear, err := time.Parse("2006", year)
	if err != nil {
		return false
	}
	return i.dob.Year() == parsedYear.Year()
}

func (i *Identity) isMclovin() bool {
	return false
}

func (i *Identity) isFace(image []byte) bool {
	imageHash := sha256.Sum256(image)
	identityHash := sha256.Sum256(i.image)
	return identityHash == imageHash
}

// dod date comparisons

func (i *Identity) isDob(date string) bool {
	parsedDate, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return false
	}

	return i.dob.Equal(parsedDate)
}

func (i *Identity) isDobBefore(date string) bool {
	parsedDate, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return false
	}
	return i.dob.Before(parsedDate)
}

func (i *Identity) isDobAfter(date string) bool {
	parsedDate, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return false
	}
	return i.dob.After(parsedDate)
}

func (i *Identity) isExpired() bool {
	return i.expiresAt.Before(time.Now())
}

func (i *Identity) isExpiresAt(expires_at string) bool {
	parsedDate, err := time.Parse(time.DateOnly, expires_at)
	if err != nil {
		return false
	}
	return i.expiresAt.Equal(parsedDate)
}

func (i *Identity) isExpiresAtBefore(expires_at string) bool {
	parsedDate, err := time.Parse(time.DateOnly, expires_at)
	if err != nil {
		return false
	}
	return i.expiresAt.Before(parsedDate)
}

func (i *Identity) isExpiresAtAfter(expires_at string) bool {
	parsedDate, err := time.Parse(time.DateOnly, expires_at)
	if err != nil {
		return false
	}
	return i.expiresAt.After(parsedDate)
}

func (i *Identity) isIssuedAtBefore(issued_at string) bool {
	parsedDate, err := time.Parse(time.DateOnly, issued_at)
	if err != nil {
		return false
	}
	return i.issuedAt.Before(parsedDate)
}

func (i *Identity) isIssuedAtAfter(issued_at string) bool {
	parsedDate, err := time.Parse(time.DateOnly, issued_at)
	if err != nil {
		return false
	}
	return i.issuedAt.After(parsedDate)
}
