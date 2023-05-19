package checkSystem

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"os"
	"strings"
)

type CheckSystem struct {
	config config
}

const verdictOKId = 1
const verdictWAId = 2

func NewCheckSystem(lang string) (*CheckSystem, error) {
	config, err := newConfig(lang)
	if err != nil {
		return nil, err
	}
	return &CheckSystem{config: config}, nil
}

func (s *CheckSystem) WriteCodeToFile(code string) (*os.File, error) {
	fileName := fmt.Sprintf("%s/code.%s", filePath, s.config.getExtension())
	codeFile, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	_, err = codeFile.WriteString(code)
	if err != nil {
		return nil, err
	}

	return codeFile, nil
}

func (s *CheckSystem) RunTests(testData models.TestData) (*models.Verdict, error) {
	if len(testData.InputData) != len(testData.OutputData) {
		return nil, errors.New("len(testData.InputData) != len(testData.OutputData)")
	}

	testCnt := len(testData.InputData)
	for i := 0; i < testCnt; i++ {
		var stderr bytes.Buffer
		var codeResult bytes.Buffer

		cmd := s.config.getCMD()
		cmd.Stdout = &codeResult
		cmd.Stderr = &stderr
		cmd.Stdin = strings.NewReader(testData.InputData[i])

		err := cmd.Run()

		if err != nil {
			return nil, err
		}

		isRight := s.check(codeResult.String(), testData.OutputData[i])

		if !isRight {
			return &models.Verdict{
				ID:    verdictWAId,
				Abbr:  "WA",
				Title: "Wrong Answer",
			}, nil
		}
		helpers.InfoLogger.Printf("RunTests: %t ", isRight)
	}

	return &models.Verdict{ID: verdictOKId, Abbr: "OK", Title: "All test passed"}, nil
}

func (s *CheckSystem) check(codeResult string, output string) bool {
	codeResult = strings.Map(filterBadSymbol, codeResult)
	output = strings.Map(filterBadSymbol, output)

	return codeResult == output
}

// Checks if symbol utf-8 code between 20 and 7E - english letters numbers and some special symbols
func filterBadSymbol(r rune) rune {
	if r >= 0x20 && r <= 0x7e {
		return r
	}
	return -1
}
