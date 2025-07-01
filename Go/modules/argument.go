package modules

import (
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
	argp "github.com/spf13/pflag"
)

type Arguments struct {
	Bubble      string
	Selection   string
	Insertion   string
	Mergesort   string
	Quicksort   string
	ShowVersion bool
}

func GetArgument() (args Arguments) {
	argp.StringVarP(&args.Bubble, "bubble", "b", "", "Bubble sort algorithm option")
	argp.StringVarP(&args.Selection, "selection", "s", "", "Selection sort algorithm option")
	argp.StringVarP(&args.Insertion, "insertion", "i", "", "Insertion sort algorithm option")
	argp.StringVarP(&args.Mergesort, "mergesort", "m", "", "Mergesort sort algorithm option")
	argp.StringVarP(&args.Quicksort, "quicksort", "q", "", "Quicksort sort algorithm option")
	argp.BoolVarP(&args.ShowVersion, "version", "v", false, "Show version")
	argp.Usage = func() {
		Banner()
		fmt.Fprintf(os.Stderr, "[%v] Usage: %s [-h] [-b </archive?>] [-s </archive?>] [-i </archive?>] [-m </archive?>] [-q </archive?>] [-r] [-v]\n", color.Green.Render("+"), os.Args[0])
		fmt.Fprintln(os.Stderr, "\nApplication to analyze the behavior and compare the efficiency of common\nsorting algorithms such as Bubble Sort, Selection Sort and Insertion Sort.\n")
		fmt.Fprintln(os.Stderr, `Options:
  -b, --bubble    </file?>   Pass data file for the BubbleSort algorithm
  -s, --selection </file?>   Pass data file for the SelectionSort algorithm
  -i, --insertion </file?>   Pass data file for the InsertionSort algorithm
  -m, --mergesort </file?>   Pass data file for the MergeSort algorithm
  -q, --quicksort </file?>   Pass data file for the QuickSort algorithm
  -v, --version              Show program version
  -h, --help                 Show this help message and exit
`)
		// argp.PrintDefaults()
	}
	argp.Parse()
	if len(os.Args) < 2 {
		argp.Usage()
		os.Exit(1)
	}
	return
}

func (self Arguments) GetSelectedAlgorithm() (name, path string, err error) {
	selected := make(map[string]string)

	if self.Bubble != "" {
		selected["bubble"] = self.Bubble
	}
	if self.Selection != "" {
		selected["selection"] = self.Selection
	}
	if self.Insertion != "" {
		selected["insertion"] = self.Insertion
	}
	if self.Mergesort != "" {
		selected["mergesort"] = self.Mergesort
	}
	if self.Quicksort != "" {
		selected["quicksort"] = self.Quicksort
	}

	if len(selected) > 1 {
		keys := make([]string, 0, len(selected))
		for k := range selected {
			keys = append(keys, k)
		}
		return "", "", fmt.Errorf("You can select only one algorithm at a time (received: %s).", strings.Join(keys, ", "))
	}

	if len(selected) == 1 && self.ShowVersion {
		return "", "", fmt.Errorf("You cannot use -v together with an algorithm.")
	}

	if len(selected) == 0 && self.ShowVersion {
		return "", "", nil
	}

	for name, path := range selected {
		return name, path, nil
	}

	return "", "", fmt.Errorf("No algorithm selected and no flags specified")
}
