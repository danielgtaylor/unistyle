package unistyle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrikethrough(t *testing.T) {
	out := Strikethrough("Strikethrough", StrikeSolidusLong)
	assert.Equal(t, "S̸t̸r̸i̸k̸e̸t̸h̸r̸o̸u̸g̸h̸", out)

	out = Strikethrough("Strikethrough", StrikeSolidusShort)
	assert.Equal(t, "S̷t̷r̷i̷k̷e̷t̷h̷r̷o̷u̷g̷h̷", out)

	out = Strikethrough("Strikethrough", StrikeStrokeLong)
	assert.Equal(t, "S̶t̶r̶i̶k̶e̶t̶h̶r̶o̶u̶g̶h̶", out)

	out = Strikethrough("Strikethrough", StrikeStrokeShort)
	assert.Equal(t, "S̵t̵r̵i̵k̵e̵t̵h̵r̵o̵u̵g̵h̵", out)

	out = Strikethrough("Strikethrough", StrikeTilde)
	assert.Equal(t, "S̴t̴r̴i̴k̴e̴t̴h̴r̴o̴u̴g̴h̴", out)
}

func TestUnderline(t *testing.T) {
	out := Underline("Underline", UnderlineBreve)
	assert.Equal(t, "U̮n̮d̮e̮r̮l̮i̮n̮e̮", out)

	out = Underline("Underline", UnderlineCaron)
	assert.Equal(t, "U̬n̬d̬e̬r̬l̬i̬n̬e̬", out)

	out = Underline("Underline", UnderlineCircumflexAccent)
	assert.Equal(t, "Ṷṋḓḙr̭ḽi̭ṋḙ", out)

	out = Underline("Underline", UnderlineDoubleLine)
	assert.Equal(t, "U̳n̳d̳e̳r̳l̳i̳n̳e̳", out)

	out = Underline("Underline", UnderlineInvertedBreve)
	assert.Equal(t, "U̯n̯d̯e̯r̯l̯i̯n̯e̯", out)

	out = Underline("Underline", UnderlineInvertedDoubleArch)
	assert.Equal(t, "U̫n̫d̫e̫r̫l̫i̫n̫e̫", out)

	out = Underline("Underline", UnderlineLine)
	assert.Equal(t, "U̲n̲d̲e̲r̲l̲i̲n̲e̲", out)

	out = Underline("Underline", UnderlineMacron)
	assert.Equal(t, "U̱ṉḏe̱ṟḻi̱ṉe̱", out)

	out = Underline("Underline", UnderlineTilde)
	assert.Equal(t, "Ṵn̰d̰ḛr̰l̰ḭn̰ḛ", out)
}

func TestOverline(t *testing.T) {
	out := Overline("Overline", OverlineBreve)
	assert.Equal(t, "Ŏv̆ĕr̆l̆ĭn̆ĕ", out)

	out = Overline("Overline", OverlineCircumflex)
	assert.Equal(t, "Ôv̂êr̂l̂în̂ê", out)

	out = Overline("Overline", OverlineDiaeresis)
	assert.Equal(t, "Öv̈ër̈l̈ïn̈ë", out)

	out = Overline("Overline", OverlineDot)
	assert.Equal(t, "Ȯv̇ėṙl̇i̇ṅė", out)

	out = Overline("Overline", OverlineLine)
	assert.Equal(t, "O̅v̅e̅r̅l̅i̅n̅e̅", out)

	out = Overline("Overline", OverlineMacron)
	assert.Equal(t, "Ōv̄ēr̄l̄īn̄ē", out)

	out = Overline("Overline", OverlineTilde)
	assert.Equal(t, "Õṽẽr̃l̃ĩñẽ", out)
}

func TestBoldSans(t *testing.T) {
	out := BoldSans("Bold")
	assert.Equal(t, "𝗕𝗼𝗹𝗱", out)

	out = BoldSans("1.2%?!ä")
	assert.Equal(t, "𝟭.𝟮%?!ä", out)
}

func TestBoldSerif(t *testing.T) {
	out := BoldSerif("Bold")
	assert.Equal(t, "𝐁𝐨𝐥𝐝", out)

	out = BoldSerif("1.2%?!ä")
	assert.Equal(t, "𝟏.𝟐%?!ä", out)
}

func TestItalicSans(t *testing.T) {
	out := ItalicSans("Italic")
	assert.Equal(t, "𝘐𝘵𝘢𝘭𝘪𝘤", out)

	out = ItalicSans("1.2%?!ä")
	assert.Equal(t, "1.2%?!ä", out)
}

func TestItalicSerif(t *testing.T) {
	out := ItalicSerif("Italic")
	assert.Equal(t, "𝐼𝑡𝑎𝑙𝑖𝑐", out)

	out = ItalicSerif("1.2%?!ä")
	assert.Equal(t, "1.2%?!ä", out)
}

func TestBoldItalicSans(t *testing.T) {
	out := BoldItalicSans("Italic")
	assert.Equal(t, "𝙄𝙩𝙖𝙡𝙞𝙘", out)

	out = BoldItalicSans("1.2%?!ä")
	assert.Equal(t, "𝟭.𝟮%?!ä", out)
}

func TestBoldItalicSerif(t *testing.T) {
	out := BoldItalicSerif("Italic")
	assert.Equal(t, "𝑰𝒕𝒂𝒍𝒊𝒄", out)

	out = BoldItalicSerif("1.2%?!ä")
	assert.Equal(t, "𝟏.𝟐%?!ä", out)
}

func TestCursive(t *testing.T) {
	out := Cursive("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	assert.Equal(t, "𝒜ℬ𝒞𝒟ℰℱ𝒢ℋℐ𝒥𝒦ℒℳ𝒩𝒪𝒫𝒬ℛ𝒮𝒯𝒰𝒱𝒲𝒳𝒴𝒵", out)

	out = Cursive("abcdefghijklmnopqrstuvwxyz")
	assert.Equal(t, "𝒶𝒷𝒸𝒹ℯ𝒻ℊ𝒽𝒾𝒿𝓀𝓁𝓂𝓃ℴ𝓅𝓆𝓇𝓈𝓉𝓊𝓋𝓌𝓍𝓎𝓏", out)
}

func TestFraktur(t *testing.T) {
	out := Fraktur("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	assert.Equal(t, "𝔄𝔅ℭ𝔇𝔈𝔉𝔊ℌℑ𝔍𝔎𝔏𝔐𝔑𝔒𝔓𝔔ℜ𝔖𝔗𝔘𝔙𝔚𝔛𝔜ℨ", out)

	out = Fraktur("abcdefghijklmnopqrstuvwxyz")
	assert.Equal(t, "𝔞𝔟𝔠𝔡𝔢𝔣𝔤𝔥𝔦𝔧𝔨𝔩𝔪𝔫𝔬𝔭𝔮𝔯𝔰𝔱𝔲𝔳𝔴𝔵𝔶𝔷", out)
}
