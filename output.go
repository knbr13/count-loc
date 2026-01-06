package main

import (
	"fmt"
	"sort"
	"strings"
)

const (
	// Table formatting constants
	colLanguage = 20
	colFiles    = 10
	colBlank    = 12
	colComment  = 12
	colCode     = 12
	colTotal    = 12
)

// PrintResults prints the results in a formatted table
func PrintResults(langStats map[string]*LanguageStats, total *LanguageStats, processedFiles, skippedFiles, errorCount int) {
	// Print header
	printHeader()

	// Sort languages by code lines (descending)
	sortedLangs := sortLanguagesByCode(langStats)

	// Print each language row
	for _, lang := range sortedLangs {
		stats := langStats[lang]
		printRow(stats.Language, stats.FileCount, stats.BlankLines, stats.CommentLines, stats.CodeLines, stats.TotalLines)
	}

	// Print separator
	printSeparator()

	// Print total row
	printRow("Total", total.FileCount, total.BlankLines, total.CommentLines, total.CodeLines, total.TotalLines)

	// Print footer with summary
	printFooter(processedFiles, skippedFiles, errorCount)
}

// printHeader prints the table header
func printHeader() {
	fmt.Println()
	printSeparator()
	fmt.Printf("%-*s %*s %*s %*s %*s %*s\n",
		colLanguage, "Language",
		colFiles, "Files",
		colBlank, "Blank",
		colComment, "Comment",
		colCode, "Code",
		colTotal, "Total")
	printSeparator()
}

// printSeparator prints a separator line
func printSeparator() {
	totalWidth := colLanguage + colFiles + colBlank + colComment + colCode + colTotal + 5 // 5 spaces between columns
	fmt.Println(strings.Repeat("-", totalWidth))
}

// printRow prints a single row of the table
func printRow(language string, files, blank, comment, code, total int) {
	// Truncate language name if too long
	if len(language) > colLanguage {
		language = language[:colLanguage-3] + "..."
	}

	fmt.Printf("%-*s %*d %*d %*d %*d %*d\n",
		colLanguage, language,
		colFiles, files,
		colBlank, blank,
		colComment, comment,
		colCode, code,
		colTotal, total)
}

// printFooter prints the summary footer
func printFooter(processedFiles, skippedFiles, errorCount int) {
	printSeparator()
	fmt.Println()
	fmt.Printf("Summary:\n")
	fmt.Printf("  Files processed: %d\n", processedFiles)
	fmt.Printf("  Files skipped:   %d\n", skippedFiles)
	if errorCount > 0 {
		fmt.Printf("  Errors:          %d\n", errorCount)
	}
	fmt.Println()
}

// sortLanguagesByCode sorts languages by code lines in descending order
func sortLanguagesByCode(langStats map[string]*LanguageStats) []string {
	langs := make([]string, 0, len(langStats))
	for lang := range langStats {
		langs = append(langs, lang)
	}

	sort.Slice(langs, func(i, j int) bool {
		return langStats[langs[i]].CodeLines > langStats[langs[j]].CodeLines
	})

	return langs
}

// PrintErrors prints the list of errors encountered
func PrintErrors(errors []error) {
	if len(errors) == 0 {
		return
	}

	fmt.Println("\nErrors encountered:")
	for i, err := range errors {
		if i >= 10 {
			fmt.Printf("  ... and %d more errors\n", len(errors)-10)
			break
		}
		fmt.Printf("  - %v\n", err)
	}
	fmt.Println()
}

// PrintCompact prints a compact summary
func PrintCompact(total *LanguageStats) {
	fmt.Printf("Files: %d | Blank: %d | Comment: %d | Code: %d | Total: %d\n",
		total.FileCount, total.BlankLines, total.CommentLines, total.CodeLines, total.TotalLines)
}

// PrintJSON prints results in JSON format
func PrintJSON(langStats map[string]*LanguageStats, total *LanguageStats) {
	fmt.Println("{")
	fmt.Println("  \"languages\": {")

	sortedLangs := sortLanguagesByCode(langStats)
	for i, lang := range sortedLangs {
		stats := langStats[lang]
		comma := ","
		if i == len(sortedLangs)-1 {
			comma = ""
		}
		fmt.Printf("    \"%s\": {\"files\": %d, \"blank\": %d, \"comment\": %d, \"code\": %d, \"total\": %d}%s\n",
			stats.Language, stats.FileCount, stats.BlankLines, stats.CommentLines, stats.CodeLines, stats.TotalLines, comma)
	}

	fmt.Println("  },")
	fmt.Printf("  \"total\": {\"files\": %d, \"blank\": %d, \"comment\": %d, \"code\": %d, \"total\": %d}\n",
		total.FileCount, total.BlankLines, total.CommentLines, total.CodeLines, total.TotalLines)
	fmt.Println("}")
}

// PrintByFiles prints results sorted by file count
func PrintByFiles(langStats map[string]*LanguageStats, total *LanguageStats, processedFiles, skippedFiles, errorCount int) {
	// Print header
	printHeader()

	// Sort languages by file count (descending)
	langs := make([]string, 0, len(langStats))
	for lang := range langStats {
		langs = append(langs, lang)
	}
	sort.Slice(langs, func(i, j int) bool {
		return langStats[langs[i]].FileCount > langStats[langs[j]].FileCount
	})

	// Print each language row
	for _, lang := range langs {
		stats := langStats[lang]
		printRow(stats.Language, stats.FileCount, stats.BlankLines, stats.CommentLines, stats.CodeLines, stats.TotalLines)
	}

	// Print separator
	printSeparator()

	// Print total row
	printRow("Total", total.FileCount, total.BlankLines, total.CommentLines, total.CodeLines, total.TotalLines)

	// Print footer with summary
	printFooter(processedFiles, skippedFiles, errorCount)
}

// FormatNumber formats a number with thousand separators
func FormatNumber(n int) string {
	if n < 1000 {
		return fmt.Sprintf("%d", n)
	}

	str := fmt.Sprintf("%d", n)
	var result strings.Builder
	length := len(str)

	for i, char := range str {
		if i > 0 && (length-i)%3 == 0 {
			result.WriteRune(',')
		}
		result.WriteRune(char)
	}

	return result.String()
}

// PrintResultsFormatted prints results with formatted numbers
func PrintResultsFormatted(langStats map[string]*LanguageStats, total *LanguageStats, processedFiles, skippedFiles, errorCount int) {
	fmt.Println()
	printSeparator()
	fmt.Printf("%-*s %*s %*s %*s %*s %*s\n",
		colLanguage, "Language",
		colFiles, "Files",
		colBlank, "Blank",
		colComment, "Comment",
		colCode, "Code",
		colTotal, "Total")
	printSeparator()

	// Sort languages by code lines (descending)
	sortedLangs := sortLanguagesByCode(langStats)

	// Print each language row with formatted numbers
	for _, lang := range sortedLangs {
		stats := langStats[lang]
		language := stats.Language
		if len(language) > colLanguage {
			language = language[:colLanguage-3] + "..."
		}
		fmt.Printf("%-*s %*s %*s %*s %*s %*s\n",
			colLanguage, language,
			colFiles, FormatNumber(stats.FileCount),
			colBlank, FormatNumber(stats.BlankLines),
			colComment, FormatNumber(stats.CommentLines),
			colCode, FormatNumber(stats.CodeLines),
			colTotal, FormatNumber(stats.TotalLines))
	}

	printSeparator()

	// Print total row with formatted numbers
	fmt.Printf("%-*s %*s %*s %*s %*s %*s\n",
		colLanguage, "Total",
		colFiles, FormatNumber(total.FileCount),
		colBlank, FormatNumber(total.BlankLines),
		colComment, FormatNumber(total.CommentLines),
		colCode, FormatNumber(total.CodeLines),
		colTotal, FormatNumber(total.TotalLines))

	printFooter(processedFiles, skippedFiles, errorCount)
}
