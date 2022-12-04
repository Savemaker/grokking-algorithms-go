package main

const (
	NOVOKOSINO stationName = iota
	NOVOGIREEVO
	PEROVO
	AVIAMOTORNAYA
	MARKSISTSKAYA
	TRETYAKOVSKAYA
	TAGANSKAYA
	KITAY_GOROD
)

type stationName int

type moscowMetro struct {
	graph map[stationName][]station
}

type station struct {
	name          stationName
	timeCostInMin int
}

func NewMoscowMetro() *moscowMetro {
	moscowMetro := moscowMetro{
		graph: map[stationName][]station{},
	}

	graph := moscowMetro.graph

	graph[NOVOKOSINO] = []station{
		{name: NOVOGIREEVO, timeCostInMin: 3}}

	graph[NOVOGIREEVO] = []station{
		{name: PEROVO, timeCostInMin: 3},
		{name: NOVOKOSINO, timeCostInMin: 3}}

	graph[PEROVO] = []station{
		{name: NOVOGIREEVO, timeCostInMin: 3},
		{name: AVIAMOTORNAYA, timeCostInMin: 3},
	}

	graph[AVIAMOTORNAYA] = []station{
		{name: PEROVO, timeCostInMin: 3},
		{name: MARKSISTSKAYA, timeCostInMin: 3},
	}

	graph[MARKSISTSKAYA] = []station{
		{name: AVIAMOTORNAYA, timeCostInMin: 3},
		{name: TAGANSKAYA, timeCostInMin: 5},
		{name: TRETYAKOVSKAYA, timeCostInMin: 3},
	}

	graph[TAGANSKAYA] = []station{
		{name: MARKSISTSKAYA, timeCostInMin: 5},
		{name: KITAY_GOROD, timeCostInMin: 3},
	}

	graph[KITAY_GOROD] = []station{
		{name: TAGANSKAYA, timeCostInMin: 3},
		{name: TRETYAKOVSKAYA, timeCostInMin: 3},
	}

	graph[TRETYAKOVSKAYA] = []station{
		{name: MARKSISTSKAYA, timeCostInMin: 3},
		{name: KITAY_GOROD, timeCostInMin: 3},
	}

	return &moscowMetro
}

func (m moscowMetro) CalculateRoute(origin stationName, destination stationName) bool {
	queue := m.graph[origin]
	searchMap := map[stationName]bool{}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if _, found := searchMap[top.name]; !found {
			searchMap[top.name] = true
			if top.name == destination {
				return true
			} else {
				queue = append(queue, m.graph[top.name]...)
			}
		}
	}
	return false
}
