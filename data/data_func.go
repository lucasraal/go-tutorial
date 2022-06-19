package data

func Team() (Teams []team) {

	t1 := team{
		Id:		1,
		Name:	"Flamengo",
		Power:	73,
		Played:	false,
	}

	t2 := team{
		Id:		2,
		Name:	"Palmeiras",
		Power:	88,
		Played:	false,
	}

	Teams = []team{t1, t2}

	return Teams
}

