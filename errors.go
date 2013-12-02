package greenspun

type ListError int

func (s ListError) Error() string {
	return listErrText[s]
}

const(
	LIST_EMPTY = ListError(iota)
	LIST_TOO_SHALLOW
	LIST_UNINITIALIZED
	LIST_REQUIRED
)

var listErrText = map[ListError] string {
	LIST_EMPTY:					"list empty",
	LIST_TOO_SHALLOW:		"list contains too few items",
	LIST_UNINITIALIZED:	"list needs to be initialised",
	LIST_REQUIRED:				"a list is required",
}


type PairError int

func (s PairError) Error() string {
	return pairErrText[s]
}

const(
	PAIR_EMPTY = PairError(iota)
	PAIR_LIST_TOO_SHALLOW
	PAIR_UNINITIALIZED
	PAIR_REQUIRED
)

var pairErrText = map[PairError] string {
	PAIR_EMPTY:							"pair empty",
	PAIR_LIST_TOO_SHALLOW:	"pair list contains too few items",
	PAIR_REQUIRED:					"a pair is required",
}


type ArgumentError int

func (s ArgumentError) Error() string {
	return argErrText[s]
}

const(
	ARGUMENT_NEGATIVE_INDEX = ArgumentError(iota)
)

var argErrText = map[ArgumentError] string {
	ARGUMENT_NEGATIVE_INDEX:	"positive index required",
}
