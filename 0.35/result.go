package main

type Result struct {
	WinCount       int
	LoseCount      int
	ContinuousLose int
}

func NewResult(win, lose, avgLose int) Result {
	return Result{
		WinCount:       win,
		LoseCount:      lose,
		ContinuousLose: avgLose,
	}
}

type Results struct {
	results []Result
}

func NewResults(r []Result) Results {
	return Results{
		results: r,
	}
}

func (r Results) calcTotalWin() int {
	sum := 0
	for _, s := range r.results {
		sum += s.WinCount
	}
	return sum
}

func (r Results) calcTotalLose() int {
	sum := 0
	for _, s := range r.results {
		sum += s.LoseCount
	}
	return sum
}

func (r Results) calcAvgContinuLose() int {
	total := 0
	for _, s := range r.results {
		total += s.ContinuousLose
	}

	return int(total / len(r.results))
}
