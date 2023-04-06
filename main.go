/*
Example output :

== Summary ==
Number of occurrences  if Go: 10
Number of lines: 7
Lines: [ 1 - 8 - 15 - 17 - 19 - 23 - 28 ]
== End of Summary ==
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ProcessLine searches for old_pattern in line to replace bu new_pattern
// It returns found=true if pattern was found, res with result string and occ occurence number of old
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line
	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)

		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)

	}
	return found, res, occ
}

/*
@src_file string : source file to replace
@old_pattern string : pattern to find
@new_pattern string : pattern to replace

@occ int : Number of occurences of old
@lines []int : every lines where old appear
@err error : errors
*/
func FindReplaceFile(src, dst, old, new string) (occ int, lines []int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return occ, lines, err
	}
	defer dstFile.Close()

	old = old + " "
	new = new + " "

	lineIdx := 1
	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()
	for scanner.Scan() {
		found, res, o := ProcessLine(scanner.Text(), old, new)
		if found {
			occ += o
			lines = append(lines, lineIdx)
		}
		fmt.Fprintf(writer, res)
		lineIdx++
	}
	return occ, lines, nil
}

func main() {

	old := "Go"
	new := "Python"
	occ, lines, err := FindReplaceFile("wikigo.txt", "wikidest.txt", old, new)
	if err != nil {
		fmt.Printf("Error while executing find replace: %v\n", err)
	}

	fmt.Println("=== SUMMARY ===")
	defer fmt.Println("=== END OF SUMMARY ===")
	fmt.Printf("Number of occurence of %v : %v\n", old, occ)
	fmt.Printf("Number of lines : %d\n", len(lines))
	fmt.Print("Lines: [")
	len := len(lines)
	for i, l := range lines {
		fmt.Printf("%v", l)
		if i < len-1 {
			fmt.Printf(" - ")
		}
	}
	fmt.Println(" ]")
}
