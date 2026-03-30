package predownload

import (
	"bytes"
	"fmt"
	"strconv"
)

type ParseBencodeResult struct {
	announce     string
	comment      string
	creationDate int64
	info         BencodeInfo
}

type BencodeInfo struct {
	length      int64
	name        string
	pieceLength int64
	pieces      []byte
}

func ParseBencode(bencode []byte) (ParseBencodeResult, error) {
	// Takes in a Bencoded string and extracts fields in ParseBencodeResult.
	data, _, err := parse(bencode, 0)
	if err != nil {
		return ParseBencodeResult{}, err
	}

	result := ParseBencodeResult{}

	switch data.(type) {
	case map[string]any:
		result.announce = data["announce"].(string)
		result.comment = data["comment"].(string)
		result.creationDate = data["creation date"].(int64)
		result.info = data["info"].(BencodeInfo)
	}

	return result, nil
}

func parse(bencode []byte, pos int) (any, int, error) {
	// Parses a portion of bencode and returns the correct data structure. Returns (data, newPos, error)

	switch bencode[pos] {
	case 'i':
		return parseInt(bencode, pos)
	case 'l':
		return parseList(bencode, pos)
	case 'd':
		return parseDict(bencode, pos)
	default:
		return parseStr(bencode, pos)
	}
}

func parseStr(bencode []byte, pos int) (string, int, error) {
	// Bencode strings look like: 4:spam
	endLength := bytes.IndexByte(bencode[pos:], ':')

	if endLength == -1 {
		return "", 0, fmt.Errorf("invalid string")
	}

	length, err := strconv.ParseInt(string(bencode[pos:endLength]), 10, 64)

	if err != nil {
		return "", 0, err
	}

	return string(bencode[endLength+1 : endLength+1+int(length)]), endLength + 1 + int(length), nil
}

func parseInt(bencode []byte, pos int) (int64, int, error) {
	// Bencode ints look like: i123e
	end := bytes.IndexByte(bencode[pos:], 'e')

	if end == -1 {
		return 0, 0, fmt.Errorf("invalid integer")
	}

	value, err := strconv.ParseInt(string(bencode[pos+1:end]), 10, 64)
	return value, end + 1, err
}

func parseList(bencode []byte, pos int) ([]any, int, error) {
	// Bencode lists look like: l4:spam4:eggse
	var list []any

	for pos < len(bencode) && bencode[pos] != 'e' {
		data, newPos, err := parse(bencode, pos)
		if err != nil {
			return list, pos, err
		}

		list = append(list, data)
		pos = newPos
	}

	return list, pos + 1, nil
}

func parseDict(bencode []byte, pos int) (map[string]any, int, error) {
	// Bencode dicts look like: d3:cow3:moo4:spam4:eggse
	dict := make(map[string]any)

	for pos < len(bencode) && bencode[pos] != 'e' {
		key, newPos, err := parseStr(bencode, pos)
		if err != nil {
			return dict, pos, err
		}

		data, newPos, err := parse(bencode, newPos)
		if err != nil {
			return dict, pos, err
		}

		dict[key] = data
		pos = newPos
	}

	return dict, pos + 1, nil
}
