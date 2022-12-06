package day5

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day5() {

	var stacks, stacks1, stacks2 [][]string
	var lines []string

	data, err := os.Open("day5/inputDay5.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	idk := []string{"VRHBGDW", "FRCGNJ", "JNDHFSL", "VSDJ", "VNWQRDHS", "MCHGP", "CHZLGBJF", "RJS", "MVNBRSGL"} //welp i dunno other way to read that file lmao
	//idk := []string{"NZ", "DCM", "P"}

	for i := 1; i <= len(idk); i++ {
		stacks = append(stacks, strings.Split(idk[i-1], ""))
	}
	stacks1 = dupe(stacks)
	stacks2 = dupe(stacks)
	crane9000(lines, stacks1)
	crane9001(lines, stacks2)

	fmt.Printf("Answer 5_1: %s\n----------------------\nAnswer 5_2: %s\n----------------------", res(stacks1), res(stacks2))
}
func crane9000(beep []string, stacks [][]string) {
	re := regexp.MustCompile(`\d+`)
	for _, v := range beep {
		beep := re.FindAllString(v, -1)
		inst := inst(strToInt(beep))
		stk := stacks[inst.From]
		stacks[inst.From] = stk[inst.Amt:] //remaining shit
		things := stk[:inst.Amt]           //picked up shit
		sgniht := reverseShit(things)
		stacks[inst.To] = append(sgniht, stacks[inst.To]...) //Dropping picked up shit
	}
}

func crane9001(beep []string, stacks [][]string) {
	re := regexp.MustCompile(`\d+`)
	for _, v := range beep {
		beep := re.FindAllString(v, -1)
		inst := inst(strToInt(beep))
		stk := stacks[inst.From]
		stacks[inst.From] = stk[inst.Amt:] //remaining shit
		things := stk[:inst.Amt]           //picked up shit
		sgniht := reverseShit(things)
		things = reverseShit(sgniht)                         //fuk u code
		stacks[inst.To] = append(things, stacks[inst.To]...) //Dropping picked up shit
	}
}

func res(stacks [][]string) (res string) {
	for _, c := range stacks {
		res = res + c[0]
	}
	return res
}
func reverseShit(things []string) (out []string) {
	for i := range things {
		out = append(out, things[len(things)-i-1])
	}
	return out
}
func strToInt(beep []string) (out []int) {
	for _, i := range beep {
		n, _ := strconv.Atoi(i)
		out = append(out, n)
	}
	return out
}
func inst(in []int) (out Inst) {
	out.Amt = in[0]
	out.From = in[1] - 1
	out.To = in[2] - 1
	return out
}
func dupe(inp [][]string) [][]string {
	duplicate := make([][]string, len(inp))
	for i := range inp {
		duplicate[i] = make([]string, len(inp[i]))
		copy(duplicate[i], inp[i])
	}
	return duplicate
}

type Inst struct {
	Amt  int
	From int
	To   int
}
