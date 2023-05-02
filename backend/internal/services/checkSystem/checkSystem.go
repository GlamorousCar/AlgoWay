package checkSystem

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"os"
	"strings"
)

type CheckSystem struct {
	config config
}

func NewCheckSystem(lang string) (*CheckSystem, error) {
	config, err := newConfig(lang)
	if err != nil {
		return nil, err
	}
	return &CheckSystem{config: config}, nil
}

func (s *CheckSystem) WriteCodeToFile(code string) error {
	fileName := fmt.Sprintf("./tests/code.%s", s.config.getExtension())
	outfile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = outfile.WriteString(code)
	if err != nil {
		return err
	}

	return nil
}

func (s *CheckSystem) RunTests(testData models.TestData) error {
	cmd := s.config.getCMD()

	if len(testData.InputData) != len(testData.OutputData) {
		return errors.New("len(testData.InputData) != len(testData.OutputData)")
	}

	testCnt := len(testData.InputData)
	for i := 0; i < testCnt; i++ {
		var stderr bytes.Buffer
		var codeResult bytes.Buffer

		cmd.Stdout = &codeResult
		cmd.Stderr = &stderr
		cmd.Stdin = strings.NewReader(testData.InputData[i])

		err := cmd.Run()
		if err != nil {
			return err
		}

		isRight := s.check(codeResult.String(), testData.OutputData[i])
		fmt.Println(isRight)
	}

	return nil
}

func (s *CheckSystem) check(codeResult string, output string) bool {
	codeResult = strings.Map(deleteShitSymbols, codeResult)
	output = strings.Map(deleteShitSymbols, output)

	return codeResult == output
}

func deleteShitSymbols(r rune) rune {
	if r >= 0x20 && r <= 0x7e {
		return r
	}
	return -1
}
