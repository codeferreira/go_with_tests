package dictionary

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, definition string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
