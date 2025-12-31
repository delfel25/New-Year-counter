package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var baseDigits = [10][5][3]bool{
	{{true, true, true}, {true, false, true}, {true, false, true}, {true, false, true}, {true, true, true}},
	{{false, false, true}, {false, false, true}, {false, false, true}, {false, false, true}, {false, false, true}},
	{{true, true, true}, {false, false, true}, {true, true, true}, {true, false, false}, {true, true, true}},
	{{true, true, true}, {false, false, true}, {true, true, true}, {false, false, true}, {true, true, true}},
	{{true, false, true}, {true, false, true}, {true, true, true}, {false, false, true}, {false, false, true}},
	{{true, true, true}, {true, false, false}, {true, true, true}, {false, false, true}, {true, true, true}},
	{{true, true, true}, {true, false, false}, {true, true, true}, {true, false, true}, {true, true, true}},
	{{true, true, true}, {false, false, true}, {false, false, true}, {false, false, true}, {false, false, true}},
	{{true, true, true}, {true, false, true}, {true, true, true}, {true, false, true}, {true, true, true}},
	{{true, true, true}, {true, false, true}, {true, true, true}, {false, false, true}, {true, true, true}},
}

func getTerminalSize() (int, int) {
	w, h := 80, 24
	if runtime.GOOS != "windows" {
		cmd := exec.Command("stty", "size")
		cmd.Stdin = os.Stdin
		out, _ := cmd.Output()
		parts := strings.Fields(string(out))
		if len(parts) == 2 {
			h, _ = strconv.Atoi(parts[0])
			w, _ = strconv.Atoi(parts[1])
		}
	}
	return w, h
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J")
	}
}

func fireworks(w, h int) {
	colors := []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m"}
	for i := 0; i < 30; i++ {
		clearScreen()
		for j := 0; j < 20; j++ {
			x, y := rand.Intn(w), rand.Intn(h)
			if y == 0 {
				y = 1
			}
			fmt.Printf("\033[%d;%dH%s%s\033[0m", y, x, colors[rand.Intn(len(colors))], "★")
		}
		time.Sleep(150 * time.Millisecond)
	}
}

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		clearScreen()
		os.Exit(0)
	}()

	for {
		w, h := getTerminalSize()
		now := time.Now()
		nextYear := time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, now.Location())
		diff := nextYear.Sub(now)

		if diff <= 0 {
			fireworks(w, h)
			continue
		}

		clearScreen()

		total := int(diff.Seconds())
		hours, mins, secs := (total%86400)/3600, (total%3600)/60, total%60
		tStr := fmt.Sprintf("%02d%02d%02d", hours, mins, secs)

		scaleW := w / 50
		scaleH := h / 10
		if scaleW < 1 {
			scaleW = 1
		}
		if scaleH < 1 {
			scaleH = 1
		}

		fullContentWidth := (6 * 3 * scaleW) + (2 * scaleW) + (5 * scaleW)
		leftPadding := (w - fullContentWidth) / 2
		if leftPadding < 0 {
			leftPadding = 0
		}

		vPadding := (h - (5 * scaleH) - 4) / 2
		if vPadding < 0 {
			vPadding = 0
		}
		fmt.Print(strings.Repeat("\n", vPadding))

		for row := 0; row < 5; row++ {
			for sh := 0; sh < scaleH; sh++ {
				var line strings.Builder
				line.WriteString(strings.Repeat(" ", leftPadding))

				for i, char := range tStr {
					val, _ := strconv.Atoi(string(char))
					for col := 0; col < 3; col++ {
						if baseDigits[val][row][col] {
							line.WriteString(strings.Repeat("▇", scaleW))
						} else {
							line.WriteString(strings.Repeat(" ", scaleW))
						}
					}
					line.WriteString(strings.Repeat(" ", scaleW))

					if i == 1 || i == 3 {
						if row == 1 || row == 3 {
							line.WriteString(strings.Repeat("▇", scaleW))
						} else {
							line.WriteString(strings.Repeat(" ", scaleW))
						}
						line.WriteString(strings.Repeat(" ", scaleW))
					}
				}
				fmt.Println(line.String())
			}
		}

		fmt.Println("\n")
		footer := fmt.Sprintf("DAYS UNTIL %d: %d | CURRENT TIME: %s", now.Year()+1, total/86400, now.Format("15:04:05"))
		fmt.Println(strings.Repeat(" ", (w-len(footer))/2) + footer)

		time.Sleep(time.Second)
	}
}
