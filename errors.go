package greenspun

type LifoError int

func (s LifoError) Error() string {
	return stackErrText[s]
}

const(
	STACK_EMPTY = LifoError(iota)
	STACK_TOO_SHALLOW
	STACK_UNINITIALIZED
	STACK_REQUIRED
)

var stackErrText = map[LifoError] string {
	STACK_EMPTY:					"stack empty",
	STACK_TOO_SHALLOW:		"stack contains too few items",
	STACK_UNINITIALIZED:	"stack needs to be initialised",
	STACK_REQUIRED:				"a stack is required",
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
