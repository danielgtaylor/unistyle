package unistyle

// StrikeStyle defines a style for strikethrough text.
type StrikeStyle rune

// Strikethrough styles
const (
	StrikeTilde StrikeStyle = '\u0334' + iota
	StrikeStrokeShort
	StrikeStrokeLong
	StrikeSolidusShort
	StrikeSolidusLong
)

// UnderlineStyle defines the style for underlined text.
type UnderlineStyle rune

// Underline styles
const (
	UnderlineInvertedDoubleArch UnderlineStyle = '\u032b' + iota
	UnderlineCaron
	UnderlineCircumflexAccent
	UnderlineBreve
	UnderlineInvertedBreve
	UnderlineTilde
	UnderlineMacron
	UnderlineLine
	UnderlineDoubleLine
)

// OverlineStyle defines the style for overlined text.
type OverlineStyle rune

//  Overline styles
const (
	OverlineCircumflex OverlineStyle = '\u0302' + iota
	OverlineTilde
	OverlineMacron
	OverlineLine
	OverlineBreve
	OverlineDot
	OverlineDiaeresis
)

// diacritic modifies each character of the input string to append the given
// character code in `r`, except for new lines.
func diacritic(text string, r rune) string {
	out := ""
	for _, c := range text {
		out += string(c)
		if c != '\n' {
			out += string(r)
		}
	}
	return out
}

// Strikethrough modifies the input text with a Unicode diacritic to create
// a line through each character like S̶t̶r̶i̶k̶e̶t̶h̶r̶o̶u̶g̶h̶ or S̷t̷r̷i̷k̷e̷t̷h̷r̷o̷u̷g̷h̷.
func Strikethrough(text string, style StrikeStyle) string {
	return diacritic(text, rune(style))
}

// Underline modifies the input text with a Unicode diacritic to create
// a line underneach each character like U̲n̲d̲e̲r̲l̲i̲n̲e̲ or U̳n̳d̳e̳r̳l̳i̳n̳e̳.
func Underline(text string, style UnderlineStyle) string {
	return diacritic(text, rune(style))
}

// Overline modifies the input text with a Unicode diacritic to create
// a line over each character like .
func Overline(text string, style OverlineStyle) string {
	return diacritic(text, rune(style))
}

// offsetTranslate takes an input string and list of offset range definitions
// to translate from one set of characters to another. Offset definitions are
// made of a starting character, ending character, and replacement for the
// starting character.
func offsetTranslate(text string, offsets [][]rune) string {
	out := ""
	for _, r := range text {
		found := false

		for _, definition := range offsets {
			start := definition[0]
			end := definition[1]
			offset := definition[2] - start

			if r >= start && r <= end {
				out += string(r + offset)
				found = true
				break
			}
		}

		if !found {
			out += string(r)
		}
	}

	return out
}

// BoldSans attempts to convert alphanumeric characters to their bolded
// equivalents. Characters without a bolded equivalent are left untouched.
func BoldSans(text string) string {
	offsets := [][]rune{
		[]rune{'A', 'Z', '𝗔'},
		[]rune{'a', 'z', '𝗮'},
		[]rune{'0', '9', '𝟬'},
	}

	return offsetTranslate(text, offsets)
}

// BoldSerif attempts to convert alphanumeric characters to their bolded
// equivalents. Characters without a bolded equivalent are left untouched.
func BoldSerif(text string) string {
	offsets := [][]rune{
		[]rune{'A', 'Z', '𝐀'},
		[]rune{'a', 'z', '𝐚'},
		[]rune{'0', '9', '𝟎'},
	}

	return offsetTranslate(text, offsets)
}

// ItalicSans attempts to convert alphanumeric characters to their italicized
// equivalents. Characters without an italicized equivalent are left
// untouched.
func ItalicSans(text string) string {
	offsets := [][]rune{
		[]rune{'A', 'Z', '𝘈'},
		[]rune{'a', 'z', '𝘢'},
	}

	return offsetTranslate(text, offsets)
}

// ItalicSerif attempts to convert alphanumeric characters to their italicized
// equivalents. Characters without an italicized equivalent are left
// untouched.
func ItalicSerif(text string) string {
	offsets := [][]rune{
		[]rune{'A', 'Z', '𝐴'},
		[]rune{'a', 'z', '𝑎'},
	}

	return offsetTranslate(text, offsets)
}

// BoldItalicSans attempts to convert alphanumeric characters to their bolded
// and italicized equivalents. Characters without a bolded and italicized
// equivalent are given their bolded equivalent if that exists, otherwise are
// left untouched.
func BoldItalicSans(text string) string {
	offsets := [][]rune{
		[]rune{'A', 'Z', '𝘼'},
		[]rune{'a', 'z', '𝙖'},
		[]rune{'0', '9', '𝟬'},
	}

	return offsetTranslate(text, offsets)
}

// BoldItalicSerif attempts to convert alphanumeric characters to their bolded
// and italicized equivalents. Characters without a bolded and italicized
// equivalent are given their bolded equivalent if that exists, otherwise are
// left untouched.
func BoldItalicSerif(text string) string {
	offsets := [][]rune{
		[]rune{'A', 'Z', '𝑨'},
		[]rune{'a', 'z', '𝒂'},
		[]rune{'0', '9', '𝟎'},
	}

	return offsetTranslate(text, offsets)
}

// Cursive attempts to convert alphanumeric characters to their cursive
// equivalents. Characters without a cursive equivalent are left untouched.
func Cursive(text string) string {
	// There is no single group of these unfortunately, so we need a separate
	// set for each non-contiguous grouping.
	offsets := [][]rune{
		[]rune{'A', 'A', '𝒜'},
		[]rune{'B', 'B', 'ℬ'},
		[]rune{'C', 'D', '𝒞'},
		[]rune{'E', 'F', 'ℰ'},
		[]rune{'G', 'G', '𝒢'},
		[]rune{'H', 'H', 'ℋ'},
		[]rune{'I', 'I', 'ℐ'},
		[]rune{'J', 'K', '𝒥'},
		[]rune{'L', 'L', 'ℒ'},
		[]rune{'M', 'M', 'ℳ'},
		[]rune{'N', 'Q', '𝒩'},
		[]rune{'R', 'R', 'ℛ'},
		[]rune{'S', 'Z', '𝒮'},
		[]rune{'a', 'd', '𝒶'},
		[]rune{'e', 'e', 'ℯ'},
		[]rune{'f', 'f', '𝒻'},
		[]rune{'g', 'g', 'ℊ'},
		[]rune{'h', 'n', '𝒽'},
		[]rune{'o', 'o', 'ℴ'},
		[]rune{'p', 'z', '𝓅'},
	}

	return offsetTranslate(text, offsets)
}

// Fraktur attempts to convert latin characters to their gothic equivalents.
// Characters without a gothic equivalent are left untouched.
func Fraktur(text string) string {
	// There is no single group of these unfortunately, so we need a separate
	// set for each non-contiguous grouping.
	offset := [][]rune{
		[]rune{'A', 'B', '𝔄'},
		[]rune{'C', 'C', 'ℭ'},
		[]rune{'D', 'G', '𝔇'},
		[]rune{'H', 'H', 'ℌ'},
		[]rune{'I', 'I', 'ℑ'},
		[]rune{'J', 'Q', '𝔍'},
		[]rune{'R', 'R', 'ℜ'},
		[]rune{'S', 'Y', '𝔖'},
		[]rune{'Z', 'Z', 'ℨ'},
		[]rune{'a', 'z', '𝔞'},
	}

	return offsetTranslate(text, offset)
}
