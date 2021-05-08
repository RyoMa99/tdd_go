package main

type Dictionary map[string]string

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrorDoesExists
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

// map は参照型なのでポインタでなくてよい
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}
	return nil
}

const (
	ErrorNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrorWordExists = DictionaryErr("cannot add word becaouse it already exists")
	ErrorDoesExists = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}
