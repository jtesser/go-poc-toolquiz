package quiz

type Tool struct {
	QuestionNum int
	ImageURL    string
}

func Random() (tool Tool, err error) {
	return getQuizDAO().getRandom()
}

func Hint(questionNum int) (hint string, err error) {
	return getQuizDAO().getHint(questionNum)
}

func Answer(questionNum int, answer string) (correct bool, err error) {
	return getQuizDAO().getAnswer(questionNum, answer)
}
