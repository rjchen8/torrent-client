package predownload

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

	return ParseBencodeResult{}, nil
}

func parse(bencode []byte, pos int) (any, error) {
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

func parseInt() {

}

func parseList() {

}

func parseDict() {

}

func parseStr() {

}
