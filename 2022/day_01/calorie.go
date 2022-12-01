package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt");

	if err != nil {
		return;
	}

	defer f.Close();

	scanner := bufio.NewScanner(f);

	max_total := 0;
	total := 0;
	sums := make([]int, 15)
	for scanner.Scan() {
		text := scanner.Text();

		if len(text) == 0 {
			sums = append(sums, total)
			if total > max_total {
				max_total = total;
			}
			total = 0;
		} else {
			value, _ := strconv.Atoi(text);
			total += value;
		}
	}
	
	biggest_sum := get_largest_n(sums, 3);
	
	sort.Sort(sort.Reverse(sort.IntSlice(sums)));

	fmt.Println(biggest_sum)
}

func get_largest_n(sums []int, n int) int {
	sort.Ints(sums)
	biggest := sums[len(sums)-n:]
	sum := 0
	for i := 0; i < len(biggest); i++ {
		sum += biggest[i];
	}

	return sum;
}
