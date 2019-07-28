package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/schafer14/chess/board"
)

type divideErrorType uint

const (
	LEGAL_MOVE_NOT_FOUND divideErrorType = iota
	ILLEGAL_MOVE_FOUND
	MOVE_COUNT_MISMATCH
	MISMATCH_MOVE_NUM
)

// The start position and the moves to make from the start position
type position struct {
	fen   string
	moves []string
}

type divideError struct {
	errorType divideErrorType
	move      string
	position  position
}

func roceCmp(position position, depth int) []divideError {
	// Setup the board
	b := board.FromFen(position.fen)

	// Make moves on board
	for _, moveStr := range position.moves {
		move, _ := b.MoveFromSrcDestNotation(moveStr)
		b.Move(move)
	}

	// Actual testing
	for i := depth; i > 0; i++ {
		bDivide := divide(b, depth)
		rDivide := roceDivide(b.String(), depth)

		errors := cmpDivideOutput(bDivide, rDivide, position)

		if len(errors) == 0 {
			return errors
		}

		for _, err := range errors {
			if err.errorType == MISMATCH_MOVE_NUM || err.errorType == ILLEGAL_MOVE_FOUND || err.errorType == LEGAL_MOVE_NOT_FOUND {
				return errors
			}
		}

		position.moves = append(position.moves, errors[0].move)
		return roceCmp(position, depth-1)

	}

	return []divideError{}
}

func (de divideError) String() string {
	var moves string
	if len(de.position.moves) > 0 {
		moves = "after moves " + strings.Join(de.position.moves, ", ")
	}

	if de.errorType == LEGAL_MOVE_NOT_FOUND {
		return fmt.Sprintf(
			"Legal move %v not available in position '%v' %v",
			de.move,
			de.position.fen,
			moves,
		)
	}
	if de.errorType == ILLEGAL_MOVE_FOUND {
		return fmt.Sprintf(
			"Illegal move %v available in position '%v' %v",
			de.move,
			de.position.fen,
			moves,
		)
	}
	if de.errorType == MOVE_COUNT_MISMATCH {
		return fmt.Sprintf(
			"Move count mismatch for move %v after position '%v' %v",
			de.move,
			de.position.fen,
			moves,
		)
	}
	if de.errorType == MISMATCH_MOVE_NUM {
		return fmt.Sprintf(
			"Move count mismatch after position '%v' %v",
			de.position.fen,
			moves,
		)
	}

	return "Unknown Divide Error"
}

func cmpDivideOutput(output divideOutput, groundTruth divideOutput, position position) (errs []divideError) {
	if output.moves != groundTruth.moves {
		errs = append(errs, divideError{
			errorType: MISMATCH_MOVE_NUM,
			position:  position,
		})
	}

	for _, trueMoveCount := range groundTruth.moveStr {
		hasMove := false
		foundCorrespondingMove := false
		for _, outputMoveCount := range output.moveStr {
			if outputMoveCount.move == trueMoveCount.move {
				hasMove = true
			}
			if outputMoveCount.move == trueMoveCount.move && outputMoveCount.count == trueMoveCount.count {
				foundCorrespondingMove = true
			}
		}

		if !hasMove {
			errs = append(errs, divideError{
				errorType: LEGAL_MOVE_NOT_FOUND,
				move:      trueMoveCount.move,
				position:  position,
			})
		}
		if !foundCorrespondingMove {
			errs = append(errs, divideError{
				errorType: MOVE_COUNT_MISMATCH,
				move:      trueMoveCount.move,
				position:  position,
			})
		}
	}

	for _, outputMoveCount := range output.moveStr {
		hasMove := false
		for _, trueMoveCount := range groundTruth.moveStr {
			if outputMoveCount.move == trueMoveCount.move {
				hasMove = true
			}
		}

		if !hasMove {
			errs = append(errs, divideError{
				errorType: ILLEGAL_MOVE_FOUND,
				move:      outputMoveCount.move,
				position:  position,
			})
		}
	}

	return errs
}

func roceDivide(fen string, d int) divideOutput {
	if fen == "startpos" {
		fen = STARTPOS
	}

	cmd := exec.Command("roce38")
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdin: %s", err.Error())
	}
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdout: %s", err.Error())
	}

	if err := cmd.Start(); nil != err {
		log.Fatalf("Error starting program: %s, %s", cmd.Path, err.Error())
	}

	reader := bufio.NewReader(stdout)
	waitForPrompt(reader)

	stdin.Write([]byte(fmt.Sprintf("st 1\n")))
	waitForPrompt(reader)

	stdin.Write([]byte(fmt.Sprintf("setboard %s\n", fen)))
	waitForPrompt(reader)

	stdin.Write([]byte(fmt.Sprintf("divide %v\n", d)))
	divide := waitForPrompt(reader)

	return parseDivide(divide)
}

func parseDivide(lines []string) divideOutput {
	var moves, nodes int
	var moveStr []moveStringCount

	length := len(lines)

	for i := 0; i < length; i++ {
		if strings.TrimSpace(lines[i]) == "" {
			continue
		}
		parts := strings.Split(strings.TrimSpace(lines[i]), " ")

		if parts[0] == "Moves" {
			moveString := strings.Split(lines[length-2], " ")[1]
			moveInt, _ := strconv.ParseInt(moveString, 10, 32)
			moves = int(moveInt)
		} else if parts[0] == "Nodes" {
			nodeString := strings.Split(lines[length-1], " ")[1]
			nodesInt, _ := strconv.ParseInt(nodeString, 10, 32)
			nodes = int(nodesInt)
		} else {
			numInt, _ := strconv.ParseInt(parts[1], 10, 32)
			moveStr = append(moveStr, moveStringCount{move: parts[0], count: int(numInt)})
		}
	}

	return divideOutput{
		moveStr: moveStr,
		moves:   moves,
		nodes:   nodes,
	}
}

func waitForPrompt(scanner *bufio.Reader) []string {
	var curLine string
	var outputs []string

loop:
	for {
		ch, _ := scanner.ReadByte()

		switch ch {
		case '\n':
			if curLine != "" {
				outputs = append(outputs, curLine)
				curLine = ""
			}
		case ':':
			if curLine == "roce" {
				break loop
			}
		default:
			curLine = curLine + string(ch)
		}

	}

	return outputs
}
