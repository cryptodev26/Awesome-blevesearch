//  Copyright (c) 2018 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// converted to Go from Lucene's AsciiFoldingFilter
// https://lucene.apache.org/core/4_0_0/analyzers-common/org/apache/lucene/analysis/miscellaneous/ASCIIFoldingFilter.html

package asciifolding

import (
	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
)

const Name = "asciifolding"

type AsciiFoldingFilter struct{}

func New() *AsciiFoldingFilter {
	return &AsciiFoldingFilter{}
}

func (s *AsciiFoldingFilter) Filter(input []byte) []byte {
	if len(input) == 0 {
		return input
	}

	in := []rune(string(input))
	length := len(in)

	// Worst-case length required if all runes fold to 4 runes
	out := make([]rune, length, length*4)

	out = foldToASCII(in, 0, out, 0, length)
	return []byte(string(out))
}

func AsciiFoldingFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.CharFilter, error) {
	return New(), nil
}

func init() {
	registry.RegisterCharFilter(Name, AsciiFoldingFilterConstructor)
}

// Converts characters above ASCII to their ASCII equivalents.
// For example, accents are removed from accented characters.
func foldToASCII(input []rune, inputPos int, output []rune, outputPos int, length int) []rune {
	end := inputPos + length
	for pos := inputPos; pos < end; pos++ {
		c := input[pos]

		// Quick test: if it's not in range then just keep current character
		if c < '\u0080' {
			output[outputPos] = c
			outputPos++
		} else {
			switch c {
			case '\u00C0': // ?? [LATIN CAPITAL LETTER A WITH GRAVE]
				fallthrough
			case '\u00C1': // ?? [LATIN CAPITAL LETTER A WITH ACUTE]
				fallthrough
			case '\u00C2': // ?? [LATIN CAPITAL LETTER A WITH CIRCUMFLEX]
				fallthrough
			case '\u00C3': // ?? [LATIN CAPITAL LETTER A WITH TILDE]
				fallthrough
			case '\u00C4': // ?? [LATIN CAPITAL LETTER A WITH DIAERESIS]
				fallthrough
			case '\u00C5': // ?? [LATIN CAPITAL LETTER A WITH RING ABOVE]
				fallthrough
			case '\u0100': // ?? [LATIN CAPITAL LETTER A WITH MACRON]
				fallthrough
			case '\u0102': // ?? [LATIN CAPITAL LETTER A WITH BREVE]
				fallthrough
			case '\u0104': // ?? [LATIN CAPITAL LETTER A WITH OGONEK]
				fallthrough
			case '\u018F': // ?? http://en.wikipedia.org/wiki/Schwa [LATIN CAPITAL LETTER SCHWA]
				fallthrough
			case '\u01CD': // ?? [LATIN CAPITAL LETTER A WITH CARON]
				fallthrough
			case '\u01DE': // ?? [LATIN CAPITAL LETTER A WITH DIAERESIS AND MACRON]
				fallthrough
			case '\u01E0': // ?? [LATIN CAPITAL LETTER A WITH DOT ABOVE AND MACRON]
				fallthrough
			case '\u01FA': // ?? [LATIN CAPITAL LETTER A WITH RING ABOVE AND ACUTE]
				fallthrough
			case '\u0200': // ?? [LATIN CAPITAL LETTER A WITH DOUBLE GRAVE]
				fallthrough
			case '\u0202': // ?? [LATIN CAPITAL LETTER A WITH INVERTED BREVE]
				fallthrough
			case '\u0226': // ?? [LATIN CAPITAL LETTER A WITH DOT ABOVE]
				fallthrough
			case '\u023A': // ?? [LATIN CAPITAL LETTER A WITH STROKE]
				fallthrough
			case '\u1D00': // ??? [LATIN LETTER SMALL CAPITAL A]
				fallthrough
			case '\u1E00': // ??? [LATIN CAPITAL LETTER A WITH RING BELOW]
				fallthrough
			case '\u1EA0': // ??? [LATIN CAPITAL LETTER A WITH DOT BELOW]
				fallthrough
			case '\u1EA2': // ??? [LATIN CAPITAL LETTER A WITH HOOK ABOVE]
				fallthrough
			case '\u1EA4': // ??? [LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND ACUTE]
				fallthrough
			case '\u1EA6': // ??? [LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND GRAVE]
				fallthrough
			case '\u1EA8': // ??? [LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND HOOK ABOVE]
				fallthrough
			case '\u1EAA': // ??? [LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND TILDE]
				fallthrough
			case '\u1EAC': // ??? [LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND DOT BELOW]
				fallthrough
			case '\u1EAE': // ??? [LATIN CAPITAL LETTER A WITH BREVE AND ACUTE]
				fallthrough
			case '\u1EB0': // ??? [LATIN CAPITAL LETTER A WITH BREVE AND GRAVE]
				fallthrough
			case '\u1EB2': // ??? [LATIN CAPITAL LETTER A WITH BREVE AND HOOK ABOVE]
				fallthrough
			case '\u1EB4': // ??? [LATIN CAPITAL LETTER A WITH BREVE AND TILDE]
				fallthrough
			case '\u24B6': // ??? [CIRCLED LATIN CAPITAL LETTER A]
				fallthrough
			case '\uFF21': // ??? [FULLWIDTH LATIN CAPITAL LETTER A]
				fallthrough
			case '\u1EB6': // ??? [LATIN CAPITAL LETTER A WITH BREVE AND DOT BELOW]
				output[outputPos] = 'A'
				outputPos++

			case '\u00E0': // ?? [LATIN SMALL LETTER A WITH GRAVE]
				fallthrough
			case '\u00E1': // ?? [LATIN SMALL LETTER A WITH ACUTE]
				fallthrough
			case '\u00E2': // ?? [LATIN SMALL LETTER A WITH CIRCUMFLEX]
				fallthrough
			case '\u00E3': // ?? [LATIN SMALL LETTER A WITH TILDE]
				fallthrough
			case '\u00E4': // ?? [LATIN SMALL LETTER A WITH DIAERESIS]
				fallthrough
			case '\u00E5': // ?? [LATIN SMALL LETTER A WITH RING ABOVE]
				fallthrough
			case '\u0101': // ?? [LATIN SMALL LETTER A WITH MACRON]
				fallthrough
			case '\u0103': // ?? [LATIN SMALL LETTER A WITH BREVE]
				fallthrough
			case '\u0105': // ?? [LATIN SMALL LETTER A WITH OGONEK]
				fallthrough
			case '\u01CE': // ?? [LATIN SMALL LETTER A WITH CARON]
				fallthrough
			case '\u01DF': // ?? [LATIN SMALL LETTER A WITH DIAERESIS AND MACRON]
				fallthrough
			case '\u01E1': // ?? [LATIN SMALL LETTER A WITH DOT ABOVE AND MACRON]
				fallthrough
			case '\u01FB': // ?? [LATIN SMALL LETTER A WITH RING ABOVE AND ACUTE]
				fallthrough
			case '\u0201': // ?? [LATIN SMALL LETTER A WITH DOUBLE GRAVE]
				fallthrough
			case '\u0203': // ?? [LATIN SMALL LETTER A WITH INVERTED BREVE]
				fallthrough
			case '\u0227': // ?? [LATIN SMALL LETTER A WITH DOT ABOVE]
				fallthrough
			case '\u0250': // ?? [LATIN SMALL LETTER TURNED A]
				fallthrough
			case '\u0259': // ?? [LATIN SMALL LETTER SCHWA]
				fallthrough
			case '\u025A': // ?? [LATIN SMALL LETTER SCHWA WITH HOOK]
				fallthrough
			case '\u1D8F': // ??? [LATIN SMALL LETTER A WITH RETROFLEX HOOK]
				fallthrough
			case '\u1D95': // ??? [LATIN SMALL LETTER SCHWA WITH RETROFLEX HOOK]
				fallthrough
			case '\u1E01': // ??? [LATIN SMALL LETTER A WITH RING BELOW]
				fallthrough
			case '\u1E9A': // ??? [LATIN SMALL LETTER A WITH RIGHT HALF RING]
				fallthrough
			case '\u1EA1': // ??? [LATIN SMALL LETTER A WITH DOT BELOW]
				fallthrough
			case '\u1EA3': // ??? [LATIN SMALL LETTER A WITH HOOK ABOVE]
				fallthrough
			case '\u1EA5': // ??? [LATIN SMALL LETTER A WITH CIRCUMFLEX AND ACUTE]
				fallthrough
			case '\u1EA7': // ??? [LATIN SMALL LETTER A WITH CIRCUMFLEX AND GRAVE]
				fallthrough
			case '\u1EA9': // ??? [LATIN SMALL LETTER A WITH CIRCUMFLEX AND HOOK ABOVE]
				fallthrough
			case '\u1EAB': // ??? [LATIN SMALL LETTER A WITH CIRCUMFLEX AND TILDE]
				fallthrough
			case '\u1EAD': // ??? [LATIN SMALL LETTER A WITH CIRCUMFLEX AND DOT BELOW]
				fallthrough
			case '\u1EAF': // ??? [LATIN SMALL LETTER A WITH BREVE AND ACUTE]
				fallthrough
			case '\u1EB1': // ??? [LATIN SMALL LETTER A WITH BREVE AND GRAVE]
				fallthrough
			case '\u1EB3': // ??? [LATIN SMALL LETTER A WITH BREVE AND HOOK ABOVE]
				fallthrough
			case '\u1EB5': // ??? [LATIN SMALL LETTER A WITH BREVE AND TILDE]
				fallthrough
			case '\u1EB7': // ??? [LATIN SMALL LETTER A WITH BREVE AND DOT BELOW]
				fallthrough
			case '\u2090': // ??? [LATIN SUBSCRIPT SMALL LETTER A]
				fallthrough
			case '\u2094': // ??? [LATIN SUBSCRIPT SMALL LETTER SCHWA]
				fallthrough
			case '\u24D0': // ??? [CIRCLED LATIN SMALL LETTER A]
				fallthrough
			case '\u2C65': // ??? [LATIN SMALL LETTER A WITH STROKE]
				fallthrough
			case '\u2C6F': // ??? [LATIN CAPITAL LETTER TURNED A]
				fallthrough
			case '\uFF41': // ??? [FULLWIDTH LATIN SMALL LETTER A]
				output[outputPos] = 'a'
				outputPos++

			case '\uA732': // ??? [LATIN CAPITAL LETTER AA]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'A'
				outputPos++
				output[outputPos] = 'A'
				outputPos++

			case '\u00C6': // ?? [LATIN CAPITAL LETTER AE]
				fallthrough
			case '\u01E2': // ?? [LATIN CAPITAL LETTER AE WITH MACRON]
				fallthrough
			case '\u01FC': // ?? [LATIN CAPITAL LETTER AE WITH ACUTE]
				fallthrough
			case '\u1D01': // ??? [LATIN LETTER SMALL CAPITAL AE]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'A'
				outputPos++
				output[outputPos] = 'E'
				outputPos++

			case '\uA734': // ??? [LATIN CAPITAL LETTER AO]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'A'
				outputPos++
				output[outputPos] = 'O'
				outputPos++

			case '\uA736': // ??? [LATIN CAPITAL LETTER AU]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'A'
				outputPos++
				output[outputPos] = 'U'
				outputPos++

			case '\uA738': // ??? [LATIN CAPITAL LETTER AV]
				fallthrough
			case '\uA73A': // ??? [LATIN CAPITAL LETTER AV WITH HORIZONTAL BAR]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'A'
				outputPos++
				output[outputPos] = 'V'
				outputPos++

			case '\uA73C': // ??? [LATIN CAPITAL LETTER AY]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'A'
				outputPos++
				output[outputPos] = 'Y'
				outputPos++

			case '\u249C': // ??? [PARENTHESIZED LATIN SMALL LETTER A]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'a'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\uA733': // ??? [LATIN SMALL LETTER AA]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'a'
				outputPos++
				output[outputPos] = 'a'
				outputPos++

			case '\u00E6': // ?? [LATIN SMALL LETTER AE]
				fallthrough
			case '\u01E3': // ?? [LATIN SMALL LETTER AE WITH MACRON]
				fallthrough
			case '\u01FD': // ?? [LATIN SMALL LETTER AE WITH ACUTE]
				fallthrough
			case '\u1D02': // ??? [LATIN SMALL LETTER TURNED AE]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'a'
				outputPos++
				output[outputPos] = 'e'
				outputPos++

			case '\uA735': // ??? [LATIN SMALL LETTER AO]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'a'
				outputPos++
				output[outputPos] = 'o'
				outputPos++

			case '\uA737': // ??? [LATIN SMALL LETTER AU]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'a'
				outputPos++
				output[outputPos] = 'u'
				outputPos++

			case '\uA739': // ??? [LATIN SMALL LETTER AV]
				fallthrough
			case '\uA73B': // ??? [LATIN SMALL LETTER AV WITH HORIZONTAL BAR]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'a'
				outputPos++
				output[outputPos] = 'v'
				outputPos++

			case '\uA73D': // ??? [LATIN SMALL LETTER AY]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'a'
				outputPos++
				output[outputPos] = 'y'
				outputPos++

			case '\u0181': // ?? [LATIN CAPITAL LETTER B WITH HOOK]
				fallthrough
			case '\u0182': // ?? [LATIN CAPITAL LETTER B WITH TOPBAR]
				fallthrough
			case '\u0243': // ?? [LATIN CAPITAL LETTER B WITH STROKE]
				fallthrough
			case '\u0299': // ?? [LATIN LETTER SMALL CAPITAL B]
				fallthrough
			case '\u1D03': // ??? [LATIN LETTER SMALL CAPITAL BARRED B]
				fallthrough
			case '\u1E02': // ??? [LATIN CAPITAL LETTER B WITH DOT ABOVE]
				fallthrough
			case '\u1E04': // ??? [LATIN CAPITAL LETTER B WITH DOT BELOW]
				fallthrough
			case '\u1E06': // ??? [LATIN CAPITAL LETTER B WITH LINE BELOW]
				fallthrough
			case '\u24B7': // ??? [CIRCLED LATIN CAPITAL LETTER B]
				fallthrough
			case '\uFF22': // ??? [FULLWIDTH LATIN CAPITAL LETTER B]
				output[outputPos] = 'B'
				outputPos++

			case '\u0180': // ?? [LATIN SMALL LETTER B WITH STROKE]
				fallthrough
			case '\u0183': // ?? [LATIN SMALL LETTER B WITH TOPBAR]
				fallthrough
			case '\u0253': // ?? [LATIN SMALL LETTER B WITH HOOK]
				fallthrough
			case '\u1D6C': // ??? [LATIN SMALL LETTER B WITH MIDDLE TILDE]
				fallthrough
			case '\u1D80': // ??? [LATIN SMALL LETTER B WITH PALATAL HOOK]
				fallthrough
			case '\u1E03': // ??? [LATIN SMALL LETTER B WITH DOT ABOVE]
				fallthrough
			case '\u1E05': // ??? [LATIN SMALL LETTER B WITH DOT BELOW]
				fallthrough
			case '\u1E07': // ??? [LATIN SMALL LETTER B WITH LINE BELOW]
				fallthrough
			case '\u24D1': // ??? [CIRCLED LATIN SMALL LETTER B]
				fallthrough
			case '\uFF42': // ??? [FULLWIDTH LATIN SMALL LETTER B]
				output[outputPos] = 'b'
				outputPos++

			case '\u249D': // ??? [PARENTHESIZED LATIN SMALL LETTER B]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'b'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u00C7': // ?? [LATIN CAPITAL LETTER C WITH CEDILLA]
				fallthrough
			case '\u0106': // ?? [LATIN CAPITAL LETTER C WITH ACUTE]
				fallthrough
			case '\u0108': // ?? [LATIN CAPITAL LETTER C WITH CIRCUMFLEX]
				fallthrough
			case '\u010A': // ?? [LATIN CAPITAL LETTER C WITH DOT ABOVE]
				fallthrough
			case '\u010C': // ?? [LATIN CAPITAL LETTER C WITH CARON]
				fallthrough
			case '\u0187': // ?? [LATIN CAPITAL LETTER C WITH HOOK]
				fallthrough
			case '\u023B': // ?? [LATIN CAPITAL LETTER C WITH STROKE]
				fallthrough
			case '\u0297': // ?? [LATIN LETTER STRETCHED C]
				fallthrough
			case '\u1D04': // ??? [LATIN LETTER SMALL CAPITAL C]
				fallthrough
			case '\u1E08': // ??? [LATIN CAPITAL LETTER C WITH CEDILLA AND ACUTE]
				fallthrough
			case '\u24B8': // ??? [CIRCLED LATIN CAPITAL LETTER C]
				fallthrough
			case '\uFF23': // ??? [FULLWIDTH LATIN CAPITAL LETTER C]
				output[outputPos] = 'C'
				outputPos++

			case '\u00E7': // ?? [LATIN SMALL LETTER C WITH CEDILLA]
				fallthrough
			case '\u0107': // ?? [LATIN SMALL LETTER C WITH ACUTE]
				fallthrough
			case '\u0109': // ?? [LATIN SMALL LETTER C WITH CIRCUMFLEX]
				fallthrough
			case '\u010B': // ?? [LATIN SMALL LETTER C WITH DOT ABOVE]
				fallthrough
			case '\u010D': // ?? [LATIN SMALL LETTER C WITH CARON]
				fallthrough
			case '\u0188': // ?? [LATIN SMALL LETTER C WITH HOOK]
				fallthrough
			case '\u023C': // ?? [LATIN SMALL LETTER C WITH STROKE]
				fallthrough
			case '\u0255': // ?? [LATIN SMALL LETTER C WITH CURL]
				fallthrough
			case '\u1E09': // ??? [LATIN SMALL LETTER C WITH CEDILLA AND ACUTE]
				fallthrough
			case '\u2184': // ??? [LATIN SMALL LETTER REVERSED C]
				fallthrough
			case '\u24D2': // ??? [CIRCLED LATIN SMALL LETTER C]
				fallthrough
			case '\uA73E': // ??? [LATIN CAPITAL LETTER REVERSED C WITH DOT]
				fallthrough
			case '\uA73F': // ??? [LATIN SMALL LETTER REVERSED C WITH DOT]
				fallthrough
			case '\uFF43': // ??? [FULLWIDTH LATIN SMALL LETTER C]
				output[outputPos] = 'c'
				outputPos++

			case '\u249E': // ??? [PARENTHESIZED LATIN SMALL LETTER C]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'c'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u00D0': // ?? [LATIN CAPITAL LETTER ETH]
				fallthrough
			case '\u010E': // ?? [LATIN CAPITAL LETTER D WITH CARON]
				fallthrough
			case '\u0110': // ?? [LATIN CAPITAL LETTER D WITH STROKE]
				fallthrough
			case '\u0189': // ?? [LATIN CAPITAL LETTER AFRICAN D]
				fallthrough
			case '\u018A': // ?? [LATIN CAPITAL LETTER D WITH HOOK]
				fallthrough
			case '\u018B': // ?? [LATIN CAPITAL LETTER D WITH TOPBAR]
				fallthrough
			case '\u1D05': // ??? [LATIN LETTER SMALL CAPITAL D]
				fallthrough
			case '\u1D06': // ??? [LATIN LETTER SMALL CAPITAL ETH]
				fallthrough
			case '\u1E0A': // ??? [LATIN CAPITAL LETTER D WITH DOT ABOVE]
				fallthrough
			case '\u1E0C': // ??? [LATIN CAPITAL LETTER D WITH DOT BELOW]
				fallthrough
			case '\u1E0E': // ??? [LATIN CAPITAL LETTER D WITH LINE BELOW]
				fallthrough
			case '\u1E10': // ??? [LATIN CAPITAL LETTER D WITH CEDILLA]
				fallthrough
			case '\u1E12': // ??? [LATIN CAPITAL LETTER D WITH CIRCUMFLEX BELOW]
				fallthrough
			case '\u24B9': // ??? [CIRCLED LATIN CAPITAL LETTER D]
				fallthrough
			case '\uA779': // ??? [LATIN CAPITAL LETTER INSULAR D]
				fallthrough
			case '\uFF24': // ??? [FULLWIDTH LATIN CAPITAL LETTER D]
				output[outputPos] = 'D'
				outputPos++

			case '\u00F0': // ?? [LATIN SMALL LETTER ETH]
				fallthrough
			case '\u010F': // ?? [LATIN SMALL LETTER D WITH CARON]
				fallthrough
			case '\u0111': // ?? [LATIN SMALL LETTER D WITH STROKE]
				fallthrough
			case '\u018C': // ?? [LATIN SMALL LETTER D WITH TOPBAR]
				fallthrough
			case '\u0221': // ?? [LATIN SMALL LETTER D WITH CURL]
				fallthrough
			case '\u0256': // ?? [LATIN SMALL LETTER D WITH TAIL]
				fallthrough
			case '\u0257': // ?? [LATIN SMALL LETTER D WITH HOOK]
				fallthrough
			case '\u1D6D': // ??? [LATIN SMALL LETTER D WITH MIDDLE TILDE]
				fallthrough
			case '\u1D81': // ??? [LATIN SMALL LETTER D WITH PALATAL HOOK]
				fallthrough
			case '\u1D91': // ??? [LATIN SMALL LETTER D WITH HOOK AND TAIL]
				fallthrough
			case '\u1E0B': // ??? [LATIN SMALL LETTER D WITH DOT ABOVE]
				fallthrough
			case '\u1E0D': // ??? [LATIN SMALL LETTER D WITH DOT BELOW]
				fallthrough
			case '\u1E0F': // ??? [LATIN SMALL LETTER D WITH LINE BELOW]
				fallthrough
			case '\u1E11': // ??? [LATIN SMALL LETTER D WITH CEDILLA]
				fallthrough
			case '\u1E13': // ??? [LATIN SMALL LETTER D WITH CIRCUMFLEX BELOW]
				fallthrough
			case '\u24D3': // ??? [CIRCLED LATIN SMALL LETTER D]
				fallthrough
			case '\uA77A': // ??? [LATIN SMALL LETTER INSULAR D]
				fallthrough
			case '\uFF44': // ??? [FULLWIDTH LATIN SMALL LETTER D]
				output[outputPos] = 'd'
				outputPos++

			case '\u01C4': // ?? [LATIN CAPITAL LETTER DZ WITH CARON]
				fallthrough
			case '\u01F1': // ?? [LATIN CAPITAL LETTER DZ]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'D'
				outputPos++
				output[outputPos] = 'Z'
				outputPos++

			case '\u01C5': // ?? [LATIN CAPITAL LETTER D WITH SMALL LETTER Z WITH CARON]
				fallthrough
			case '\u01F2': // ?? [LATIN CAPITAL LETTER D WITH SMALL LETTER Z]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'D'
				outputPos++
				output[outputPos] = 'z'
				outputPos++

			case '\u249F': // ??? [PARENTHESIZED LATIN SMALL LETTER D]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'd'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u0238': // ?? [LATIN SMALL LETTER DB DIGRAPH]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'd'
				outputPos++
				output[outputPos] = 'b'
				outputPos++

			case '\u01C6': // ?? [LATIN SMALL LETTER DZ WITH CARON]
				fallthrough
			case '\u01F3': // ?? [LATIN SMALL LETTER DZ]
				fallthrough
			case '\u02A3': // ?? [LATIN SMALL LETTER DZ DIGRAPH]
				fallthrough
			case '\u02A5': // ?? [LATIN SMALL LETTER DZ DIGRAPH WITH CURL]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'd'
				outputPos++
				output[outputPos] = 'z'
				outputPos++

			case '\u00C8': // ?? [LATIN CAPITAL LETTER E WITH GRAVE]
				fallthrough
			case '\u00C9': // ?? [LATIN CAPITAL LETTER E WITH ACUTE]
				fallthrough
			case '\u00CA': // ?? [LATIN CAPITAL LETTER E WITH CIRCUMFLEX]
				fallthrough
			case '\u00CB': // ?? [LATIN CAPITAL LETTER E WITH DIAERESIS]
				fallthrough
			case '\u0112': // ?? [LATIN CAPITAL LETTER E WITH MACRON]
				fallthrough
			case '\u0114': // ?? [LATIN CAPITAL LETTER E WITH BREVE]
				fallthrough
			case '\u0116': // ?? [LATIN CAPITAL LETTER E WITH DOT ABOVE]
				fallthrough
			case '\u0118': // ?? [LATIN CAPITAL LETTER E WITH OGONEK]
				fallthrough
			case '\u011A': // ?? [LATIN CAPITAL LETTER E WITH CARON]
				fallthrough
			case '\u018E': // ?? [LATIN CAPITAL LETTER REVERSED E]
				fallthrough
			case '\u0190': // ?? [LATIN CAPITAL LETTER OPEN E]
				fallthrough
			case '\u0204': // ?? [LATIN CAPITAL LETTER E WITH DOUBLE GRAVE]
				fallthrough
			case '\u0206': // ?? [LATIN CAPITAL LETTER E WITH INVERTED BREVE]
				fallthrough
			case '\u0228': // ?? [LATIN CAPITAL LETTER E WITH CEDILLA]
				fallthrough
			case '\u0246': // ?? [LATIN CAPITAL LETTER E WITH STROKE]
				fallthrough
			case '\u1D07': // ??? [LATIN LETTER SMALL CAPITAL E]
				fallthrough
			case '\u1E14': // ??? [LATIN CAPITAL LETTER E WITH MACRON AND GRAVE]
				fallthrough
			case '\u1E16': // ??? [LATIN CAPITAL LETTER E WITH MACRON AND ACUTE]
				fallthrough
			case '\u1E18': // ??? [LATIN CAPITAL LETTER E WITH CIRCUMFLEX BELOW]
				fallthrough
			case '\u1E1A': // ??? [LATIN CAPITAL LETTER E WITH TILDE BELOW]
				fallthrough
			case '\u1E1C': // ??? [LATIN CAPITAL LETTER E WITH CEDILLA AND BREVE]
				fallthrough
			case '\u1EB8': // ??? [LATIN CAPITAL LETTER E WITH DOT BELOW]
				fallthrough
			case '\u1EBA': // ??? [LATIN CAPITAL LETTER E WITH HOOK ABOVE]
				fallthrough
			case '\u1EBC': // ??? [LATIN CAPITAL LETTER E WITH TILDE]
				fallthrough
			case '\u1EBE': // ??? [LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND ACUTE]
				fallthrough
			case '\u1EC0': // ??? [LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND GRAVE]
				fallthrough
			case '\u1EC2': // ??? [LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND HOOK ABOVE]
				fallthrough
			case '\u1EC4': // ??? [LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND TILDE]
				fallthrough
			case '\u1EC6': // ??? [LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND DOT BELOW]
				fallthrough
			case '\u24BA': // ??? [CIRCLED LATIN CAPITAL LETTER E]
				fallthrough
			case '\u2C7B': // ??? [LATIN LETTER SMALL CAPITAL TURNED E]
				fallthrough
			case '\uFF25': // ??? [FULLWIDTH LATIN CAPITAL LETTER E]
				output[outputPos] = 'E'
				outputPos++

			case '\u00E8': // ?? [LATIN SMALL LETTER E WITH GRAVE]
				fallthrough
			case '\u00E9': // ?? [LATIN SMALL LETTER E WITH ACUTE]
				fallthrough
			case '\u00EA': // ?? [LATIN SMALL LETTER E WITH CIRCUMFLEX]
				fallthrough
			case '\u00EB': // ?? [LATIN SMALL LETTER E WITH DIAERESIS]
				fallthrough
			case '\u0113': // ?? [LATIN SMALL LETTER E WITH MACRON]
				fallthrough
			case '\u0115': // ?? [LATIN SMALL LETTER E WITH BREVE]
				fallthrough
			case '\u0117': // ?? [LATIN SMALL LETTER E WITH DOT ABOVE]
				fallthrough
			case '\u0119': // ?? [LATIN SMALL LETTER E WITH OGONEK]
				fallthrough
			case '\u011B': // ?? [LATIN SMALL LETTER E WITH CARON]
				fallthrough
			case '\u01DD': // ?? [LATIN SMALL LETTER TURNED E]
				fallthrough
			case '\u0205': // ?? [LATIN SMALL LETTER E WITH DOUBLE GRAVE]
				fallthrough
			case '\u0207': // ?? [LATIN SMALL LETTER E WITH INVERTED BREVE]
				fallthrough
			case '\u0229': // ?? [LATIN SMALL LETTER E WITH CEDILLA]
				fallthrough
			case '\u0247': // ?? [LATIN SMALL LETTER E WITH STROKE]
				fallthrough
			case '\u0258': // ?? [LATIN SMALL LETTER REVERSED E]
				fallthrough
			case '\u025B': // ?? [LATIN SMALL LETTER OPEN E]
				fallthrough
			case '\u025C': // ?? [LATIN SMALL LETTER REVERSED OPEN E]
				fallthrough
			case '\u025D': // ?? [LATIN SMALL LETTER REVERSED OPEN E WITH HOOK]
				fallthrough
			case '\u025E': // ?? [LATIN SMALL LETTER CLOSED REVERSED OPEN E]
				fallthrough
			case '\u029A': // ?? [LATIN SMALL LETTER CLOSED OPEN E]
				fallthrough
			case '\u1D08': // ??? [LATIN SMALL LETTER TURNED OPEN E]
				fallthrough
			case '\u1D92': // ??? [LATIN SMALL LETTER E WITH RETROFLEX HOOK]
				fallthrough
			case '\u1D93': // ??? [LATIN SMALL LETTER OPEN E WITH RETROFLEX HOOK]
				fallthrough
			case '\u1D94': // ??? [LATIN SMALL LETTER REVERSED OPEN E WITH RETROFLEX HOOK]
				fallthrough
			case '\u1E15': // ??? [LATIN SMALL LETTER E WITH MACRON AND GRAVE]
				fallthrough
			case '\u1E17': // ??? [LATIN SMALL LETTER E WITH MACRON AND ACUTE]
				fallthrough
			case '\u1E19': // ??? [LATIN SMALL LETTER E WITH CIRCUMFLEX BELOW]
				fallthrough
			case '\u1E1B': // ??? [LATIN SMALL LETTER E WITH TILDE BELOW]
				fallthrough
			case '\u1E1D': // ??? [LATIN SMALL LETTER E WITH CEDILLA AND BREVE]
				fallthrough
			case '\u1EB9': // ??? [LATIN SMALL LETTER E WITH DOT BELOW]
				fallthrough
			case '\u1EBB': // ??? [LATIN SMALL LETTER E WITH HOOK ABOVE]
				fallthrough
			case '\u1EBD': // ??? [LATIN SMALL LETTER E WITH TILDE]
				fallthrough
			case '\u1EBF': // ??? [LATIN SMALL LETTER E WITH CIRCUMFLEX AND ACUTE]
				fallthrough
			case '\u1EC1': // ??? [LATIN SMALL LETTER E WITH CIRCUMFLEX AND GRAVE]
				fallthrough
			case '\u1EC3': // ??? [LATIN SMALL LETTER E WITH CIRCUMFLEX AND HOOK ABOVE]
				fallthrough
			case '\u1EC5': // ??? [LATIN SMALL LETTER E WITH CIRCUMFLEX AND TILDE]
				fallthrough
			case '\u1EC7': // ??? [LATIN SMALL LETTER E WITH CIRCUMFLEX AND DOT BELOW]
				fallthrough
			case '\u2091': // ??? [LATIN SUBSCRIPT SMALL LETTER E]
				fallthrough
			case '\u24D4': // ??? [CIRCLED LATIN SMALL LETTER E]
				fallthrough
			case '\u2C78': // ??? [LATIN SMALL LETTER E WITH NOTCH]
				fallthrough
			case '\uFF45': // ??? [FULLWIDTH LATIN SMALL LETTER E]
				output[outputPos] = 'e'
				outputPos++

			case '\u24A0': // ??? [PARENTHESIZED LATIN SMALL LETTER E]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'e'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u0191': // ?? [LATIN CAPITAL LETTER F WITH HOOK]
				fallthrough
			case '\u1E1E': // ??? [LATIN CAPITAL LETTER F WITH DOT ABOVE]
				fallthrough
			case '\u24BB': // ??? [CIRCLED LATIN CAPITAL LETTER F]
				fallthrough
			case '\uA730': // ??? [LATIN LETTER SMALL CAPITAL F]
				fallthrough
			case '\uA77B': // ??? [LATIN CAPITAL LETTER INSULAR F]
				fallthrough
			case '\uA7FB': // ??? [LATIN EPIGRAPHIC LETTER REVERSED F]
				fallthrough
			case '\uFF26': // ??? [FULLWIDTH LATIN CAPITAL LETTER F]
				output[outputPos] = 'F'
				outputPos++

			case '\u0192': // ?? [LATIN SMALL LETTER F WITH HOOK]
				fallthrough
			case '\u1D6E': // ??? [LATIN SMALL LETTER F WITH MIDDLE TILDE]
				fallthrough
			case '\u1D82': // ??? [LATIN SMALL LETTER F WITH PALATAL HOOK]
				fallthrough
			case '\u1E1F': // ??? [LATIN SMALL LETTER F WITH DOT ABOVE]
				fallthrough
			case '\u1E9B': // ??? [LATIN SMALL LETTER LONG S WITH DOT ABOVE]
				fallthrough
			case '\u24D5': // ??? [CIRCLED LATIN SMALL LETTER F]
				fallthrough
			case '\uA77C': // ??? [LATIN SMALL LETTER INSULAR F]
				fallthrough
			case '\uFF46': // ??? [FULLWIDTH LATIN SMALL LETTER F]
				output[outputPos] = 'f'
				outputPos++

			case '\u24A1': // ??? [PARENTHESIZED LATIN SMALL LETTER F]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'f'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\uFB00': // ??? [LATIN SMALL LIGATURE FF]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'f'
				outputPos++
				output[outputPos] = 'f'
				outputPos++

			case '\uFB03': // ??? [LATIN SMALL LIGATURE FFI]
				output = output[:(len(output) + 2)]
				output[outputPos] = 'f'
				outputPos++
				output[outputPos] = 'f'
				outputPos++
				output[outputPos] = 'i'
				outputPos++

			case '\uFB04': // ??? [LATIN SMALL LIGATURE FFL]
				output = output[:(len(output) + 2)]
				output[outputPos] = 'f'
				outputPos++
				output[outputPos] = 'f'
				outputPos++
				output[outputPos] = 'l'
				outputPos++

			case '\uFB01': // ??? [LATIN SMALL LIGATURE FI]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'f'
				outputPos++
				output[outputPos] = 'i'
				outputPos++

			case '\uFB02': // ??? [LATIN SMALL LIGATURE FL]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'f'
				outputPos++
				output[outputPos] = 'l'
				outputPos++

			case '\u011C': // ?? [LATIN CAPITAL LETTER G WITH CIRCUMFLEX]
				fallthrough
			case '\u011E': // ?? [LATIN CAPITAL LETTER G WITH BREVE]
				fallthrough
			case '\u0120': // ?? [LATIN CAPITAL LETTER G WITH DOT ABOVE]
				fallthrough
			case '\u0122': // ?? [LATIN CAPITAL LETTER G WITH CEDILLA]
				fallthrough
			case '\u0193': // ?? [LATIN CAPITAL LETTER G WITH HOOK]
				fallthrough
			case '\u01E4': // ?? [LATIN CAPITAL LETTER G WITH STROKE]
				fallthrough
			case '\u01E5': // ?? [LATIN SMALL LETTER G WITH STROKE]
				fallthrough
			case '\u01E6': // ?? [LATIN CAPITAL LETTER G WITH CARON]
				fallthrough
			case '\u01E7': // ?? [LATIN SMALL LETTER G WITH CARON]
				fallthrough
			case '\u01F4': // ?? [LATIN CAPITAL LETTER G WITH ACUTE]
				fallthrough
			case '\u0262': // ?? [LATIN LETTER SMALL CAPITAL G]
				fallthrough
			case '\u029B': // ?? [LATIN LETTER SMALL CAPITAL G WITH HOOK]
				fallthrough
			case '\u1E20': // ??? [LATIN CAPITAL LETTER G WITH MACRON]
				fallthrough
			case '\u24BC': // ??? [CIRCLED LATIN CAPITAL LETTER G]
				fallthrough
			case '\uA77D': // ??? [LATIN CAPITAL LETTER INSULAR G]
				fallthrough
			case '\uA77E': // ??? [LATIN CAPITAL LETTER TURNED INSULAR G]
				fallthrough
			case '\uFF27': // ??? [FULLWIDTH LATIN CAPITAL LETTER G]
				output[outputPos] = 'G'
				outputPos++

			case '\u011D': // ?? [LATIN SMALL LETTER G WITH CIRCUMFLEX]
				fallthrough
			case '\u011F': // ?? [LATIN SMALL LETTER G WITH BREVE]
				fallthrough
			case '\u0121': // ?? [LATIN SMALL LETTER G WITH DOT ABOVE]
				fallthrough
			case '\u0123': // ?? [LATIN SMALL LETTER G WITH CEDILLA]
				fallthrough
			case '\u01F5': // ?? [LATIN SMALL LETTER G WITH ACUTE]
				fallthrough
			case '\u0260': // ?? [LATIN SMALL LETTER G WITH HOOK]
				fallthrough
			case '\u0261': // ?? [LATIN SMALL LETTER SCRIPT G]
				fallthrough
			case '\u1D77': // ??? [LATIN SMALL LETTER TURNED G]
				fallthrough
			case '\u1D79': // ??? [LATIN SMALL LETTER INSULAR G]
				fallthrough
			case '\u1D83': // ??? [LATIN SMALL LETTER G WITH PALATAL HOOK]
				fallthrough
			case '\u1E21': // ??? [LATIN SMALL LETTER G WITH MACRON]
				fallthrough
			case '\u24D6': // ??? [CIRCLED LATIN SMALL LETTER G]
				fallthrough
			case '\uA77F': // ??? [LATIN SMALL LETTER TURNED INSULAR G]
				fallthrough
			case '\uFF47': // ??? [FULLWIDTH LATIN SMALL LETTER G]
				output[outputPos] = 'g'
				outputPos++

			case '\u24A2': // ??? [PARENTHESIZED LATIN SMALL LETTER G]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'g'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u0124': // ?? [LATIN CAPITAL LETTER H WITH CIRCUMFLEX]
				fallthrough
			case '\u0126': // ?? [LATIN CAPITAL LETTER H WITH STROKE]
				fallthrough
			case '\u021E': // ?? [LATIN CAPITAL LETTER H WITH CARON]
				fallthrough
			case '\u029C': // ?? [LATIN LETTER SMALL CAPITAL H]
				fallthrough
			case '\u1E22': // ??? [LATIN CAPITAL LETTER H WITH DOT ABOVE]
				fallthrough
			case '\u1E24': // ??? [LATIN CAPITAL LETTER H WITH DOT BELOW]
				fallthrough
			case '\u1E26': // ??? [LATIN CAPITAL LETTER H WITH DIAERESIS]
				fallthrough
			case '\u1E28': // ??? [LATIN CAPITAL LETTER H WITH CEDILLA]
				fallthrough
			case '\u1E2A': // ??? [LATIN CAPITAL LETTER H WITH BREVE BELOW]
				fallthrough
			case '\u24BD': // ??? [CIRCLED LATIN CAPITAL LETTER H]
				fallthrough
			case '\u2C67': // ??? [LATIN CAPITAL LETTER H WITH DESCENDER]
				fallthrough
			case '\u2C75': // ??? [LATIN CAPITAL LETTER HALF H]
				fallthrough
			case '\uFF28': // ??? [FULLWIDTH LATIN CAPITAL LETTER H]
				output[outputPos] = 'H'
				outputPos++

			case '\u0125': // ?? [LATIN SMALL LETTER H WITH CIRCUMFLEX]
				fallthrough
			case '\u0127': // ?? [LATIN SMALL LETTER H WITH STROKE]
				fallthrough
			case '\u021F': // ?? [LATIN SMALL LETTER H WITH CARON]
				fallthrough
			case '\u0265': // ?? [LATIN SMALL LETTER TURNED H]
				fallthrough
			case '\u0266': // ?? [LATIN SMALL LETTER H WITH HOOK]
				fallthrough
			case '\u02AE': // ?? [LATIN SMALL LETTER TURNED H WITH FISHHOOK]
				fallthrough
			case '\u02AF': // ?? [LATIN SMALL LETTER TURNED H WITH FISHHOOK AND TAIL]
				fallthrough
			case '\u1E23': // ??? [LATIN SMALL LETTER H WITH DOT ABOVE]
				fallthrough
			case '\u1E25': // ??? [LATIN SMALL LETTER H WITH DOT BELOW]
				fallthrough
			case '\u1E27': // ??? [LATIN SMALL LETTER H WITH DIAERESIS]
				fallthrough
			case '\u1E29': // ??? [LATIN SMALL LETTER H WITH CEDILLA]
				fallthrough
			case '\u1E2B': // ??? [LATIN SMALL LETTER H WITH BREVE BELOW]
				fallthrough
			case '\u1E96': // ??? [LATIN SMALL LETTER H WITH LINE BELOW]
				fallthrough
			case '\u24D7': // ??? [CIRCLED LATIN SMALL LETTER H]
				fallthrough
			case '\u2C68': // ??? [LATIN SMALL LETTER H WITH DESCENDER]
				fallthrough
			case '\u2C76': // ??? [LATIN SMALL LETTER HALF H]
				fallthrough
			case '\uFF48': // ??? [FULLWIDTH LATIN SMALL LETTER H]
				output[outputPos] = 'h'
				outputPos++

			case '\u01F6': // ?? http://en.wikipedia.org/wiki/Hwair [LATIN CAPITAL LETTER HWAIR]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'H'
				outputPos++
				output[outputPos] = 'V'
				outputPos++

			case '\u24A3': // ??? [PARENTHESIZED LATIN SMALL LETTER H]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'h'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u0195': // ?? [LATIN SMALL LETTER HV]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'h'
				outputPos++
				output[outputPos] = 'v'
				outputPos++

			case '\u00CC': // ?? [LATIN CAPITAL LETTER I WITH GRAVE]
				fallthrough
			case '\u00CD': // ?? [LATIN CAPITAL LETTER I WITH ACUTE]
				fallthrough
			case '\u00CE': // ?? [LATIN CAPITAL LETTER I WITH CIRCUMFLEX]
				fallthrough
			case '\u00CF': // ?? [LATIN CAPITAL LETTER I WITH DIAERESIS]
				fallthrough
			case '\u0128': // ?? [LATIN CAPITAL LETTER I WITH TILDE]
				fallthrough
			case '\u012A': // ?? [LATIN CAPITAL LETTER I WITH MACRON]
				fallthrough
			case '\u012C': // ?? [LATIN CAPITAL LETTER I WITH BREVE]
				fallthrough
			case '\u012E': // ?? [LATIN CAPITAL LETTER I WITH OGONEK]
				fallthrough
			case '\u0130': // ?? [LATIN CAPITAL LETTER I WITH DOT ABOVE]
				fallthrough
			case '\u0196': // ?? [LATIN CAPITAL LETTER IOTA]
				fallthrough
			case '\u0197': // ?? [LATIN CAPITAL LETTER I WITH STROKE]
				fallthrough
			case '\u01CF': // ?? [LATIN CAPITAL LETTER I WITH CARON]
				fallthrough
			case '\u0208': // ?? [LATIN CAPITAL LETTER I WITH DOUBLE GRAVE]
				fallthrough
			case '\u020A': // ?? [LATIN CAPITAL LETTER I WITH INVERTED BREVE]
				fallthrough
			case '\u026A': // ?? [LATIN LETTER SMALL CAPITAL I]
				fallthrough
			case '\u1D7B': // ??? [LATIN SMALL CAPITAL LETTER I WITH STROKE]
				fallthrough
			case '\u1E2C': // ??? [LATIN CAPITAL LETTER I WITH TILDE BELOW]
				fallthrough
			case '\u1E2E': // ??? [LATIN CAPITAL LETTER I WITH DIAERESIS AND ACUTE]
				fallthrough
			case '\u1EC8': // ??? [LATIN CAPITAL LETTER I WITH HOOK ABOVE]
				fallthrough
			case '\u1ECA': // ??? [LATIN CAPITAL LETTER I WITH DOT BELOW]
				fallthrough
			case '\u24BE': // ??? [CIRCLED LATIN CAPITAL LETTER I]
				fallthrough
			case '\uA7FE': // ??? [LATIN EPIGRAPHIC LETTER I LONGA]
				fallthrough
			case '\uFF29': // ??? [FULLWIDTH LATIN CAPITAL LETTER I]
				output[outputPos] = 'I'
				outputPos++

			case '\u00EC': // ?? [LATIN SMALL LETTER I WITH GRAVE]
				fallthrough
			case '\u00ED': // ?? [LATIN SMALL LETTER I WITH ACUTE]
				fallthrough
			case '\u00EE': // ?? [LATIN SMALL LETTER I WITH CIRCUMFLEX]
				fallthrough
			case '\u00EF': // ?? [LATIN SMALL LETTER I WITH DIAERESIS]
				fallthrough
			case '\u0129': // ?? [LATIN SMALL LETTER I WITH TILDE]
				fallthrough
			case '\u012B': // ?? [LATIN SMALL LETTER I WITH MACRON]
				fallthrough
			case '\u012D': // ?? [LATIN SMALL LETTER I WITH BREVE]
				fallthrough
			case '\u012F': // ?? [LATIN SMALL LETTER I WITH OGONEK]
				fallthrough
			case '\u0131': // ?? [LATIN SMALL LETTER DOTLESS I]
				fallthrough
			case '\u01D0': // ?? [LATIN SMALL LETTER I WITH CARON]
				fallthrough
			case '\u0209': // ?? [LATIN SMALL LETTER I WITH DOUBLE GRAVE]
				fallthrough
			case '\u020B': // ?? [LATIN SMALL LETTER I WITH INVERTED BREVE]
				fallthrough
			case '\u0268': // ?? [LATIN SMALL LETTER I WITH STROKE]
				fallthrough
			case '\u1D09': // ??? [LATIN SMALL LETTER TURNED I]
				fallthrough
			case '\u1D62': // ??? [LATIN SUBSCRIPT SMALL LETTER I]
				fallthrough
			case '\u1D7C': // ??? [LATIN SMALL LETTER IOTA WITH STROKE]
				fallthrough
			case '\u1D96': // ??? [LATIN SMALL LETTER I WITH RETROFLEX HOOK]
				fallthrough
			case '\u1E2D': // ??? [LATIN SMALL LETTER I WITH TILDE BELOW]
				fallthrough
			case '\u1E2F': // ??? [LATIN SMALL LETTER I WITH DIAERESIS AND ACUTE]
				fallthrough
			case '\u1EC9': // ??? [LATIN SMALL LETTER I WITH HOOK ABOVE]
				fallthrough
			case '\u1ECB': // ??? [LATIN SMALL LETTER I WITH DOT BELOW]
				fallthrough
			case '\u2071': // ??? [SUPERSCRIPT LATIN SMALL LETTER I]
				fallthrough
			case '\u24D8': // ??? [CIRCLED LATIN SMALL LETTER I]
				fallthrough
			case '\uFF49': // ??? [FULLWIDTH LATIN SMALL LETTER I]
				output[outputPos] = 'i'
				outputPos++

			case '\u0132': // ?? [LATIN CAPITAL LIGATURE IJ]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'I'
				outputPos++
				output[outputPos] = 'J'
				outputPos++

			case '\u24A4': // ??? [PARENTHESIZED LATIN SMALL LETTER I]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'i'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u0133': // ?? [LATIN SMALL LIGATURE IJ]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'i'
				outputPos++
				output[outputPos] = 'j'
				outputPos++

			case '\u0134': // ?? [LATIN CAPITAL LETTER J WITH CIRCUMFLEX]
				fallthrough
			case '\u0248': // ?? [LATIN CAPITAL LETTER J WITH STROKE]
				fallthrough
			case '\u1D0A': // ??? [LATIN LETTER SMALL CAPITAL J]
				fallthrough
			case '\u24BF': // ??? [CIRCLED LATIN CAPITAL LETTER J]
				fallthrough
			case '\uFF2A': // ??? [FULLWIDTH LATIN CAPITAL LETTER J]
				output[outputPos] = 'J'
				outputPos++

			case '\u0135': // ?? [LATIN SMALL LETTER J WITH CIRCUMFLEX]
				fallthrough
			case '\u01F0': // ?? [LATIN SMALL LETTER J WITH CARON]
				fallthrough
			case '\u0237': // ?? [LATIN SMALL LETTER DOTLESS J]
				fallthrough
			case '\u0249': // ?? [LATIN SMALL LETTER J WITH STROKE]
				fallthrough
			case '\u025F': // ?? [LATIN SMALL LETTER DOTLESS J WITH STROKE]
				fallthrough
			case '\u0284': // ?? [LATIN SMALL LETTER DOTLESS J WITH STROKE AND HOOK]
				fallthrough
			case '\u029D': // ?? [LATIN SMALL LETTER J WITH CROSSED-TAIL]
				fallthrough
			case '\u24D9': // ??? [CIRCLED LATIN SMALL LETTER J]
				fallthrough
			case '\u2C7C': // ??? [LATIN SUBSCRIPT SMALL LETTER J]
				fallthrough
			case '\uFF4A': // ??? [FULLWIDTH LATIN SMALL LETTER J]
				output[outputPos] = 'j'
				outputPos++

			case '\u24A5': // ??? [PARENTHESIZED LATIN SMALL LETTER J]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'j'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u0136': // ?? [LATIN CAPITAL LETTER K WITH CEDILLA]
				fallthrough
			case '\u0198': // ?? [LATIN CAPITAL LETTER K WITH HOOK]
				fallthrough
			case '\u01E8': // ?? [LATIN CAPITAL LETTER K WITH CARON]
				fallthrough
			case '\u1D0B': // ??? [LATIN LETTER SMALL CAPITAL K]
				fallthrough
			case '\u1E30': // ??? [LATIN CAPITAL LETTER K WITH ACUTE]
				fallthrough
			case '\u1E32': // ??? [LATIN CAPITAL LETTER K WITH DOT BELOW]
				fallthrough
			case '\u1E34': // ??? [LATIN CAPITAL LETTER K WITH LINE BELOW]
				fallthrough
			case '\u24C0': // ??? [CIRCLED LATIN CAPITAL LETTER K]
				fallthrough
			case '\u2C69': // ??? [LATIN CAPITAL LETTER K WITH DESCENDER]
				fallthrough
			case '\uA740': // ??? [LATIN CAPITAL LETTER K WITH STROKE]
				fallthrough
			case '\uA742': // ??? [LATIN CAPITAL LETTER K WITH DIAGONAL STROKE]
				fallthrough
			case '\uA744': // ??? [LATIN CAPITAL LETTER K WITH STROKE AND DIAGONAL STROKE]
				fallthrough
			case '\uFF2B': // ??? [FULLWIDTH LATIN CAPITAL LETTER K]
				output[outputPos] = 'K'
				outputPos++

			case '\u0137': // ?? [LATIN SMALL LETTER K WITH CEDILLA]
				fallthrough
			case '\u0199': // ?? [LATIN SMALL LETTER K WITH HOOK]
				fallthrough
			case '\u01E9': // ?? [LATIN SMALL LETTER K WITH CARON]
				fallthrough
			case '\u029E': // ?? [LATIN SMALL LETTER TURNED K]
				fallthrough
			case '\u1D84': // ??? [LATIN SMALL LETTER K WITH PALATAL HOOK]
				fallthrough
			case '\u1E31': // ??? [LATIN SMALL LETTER K WITH ACUTE]
				fallthrough
			case '\u1E33': // ??? [LATIN SMALL LETTER K WITH DOT BELOW]
				fallthrough
			case '\u1E35': // ??? [LATIN SMALL LETTER K WITH LINE BELOW]
				fallthrough
			case '\u24DA': // ??? [CIRCLED LATIN SMALL LETTER K]
				fallthrough
			case '\u2C6A': // ??? [LATIN SMALL LETTER K WITH DESCENDER]
				fallthrough
			case '\uA741': // ??? [LATIN SMALL LETTER K WITH STROKE]
				fallthrough
			case '\uA743': // ??? [LATIN SMALL LETTER K WITH DIAGONAL STROKE]
				fallthrough
			case '\uA745': // ??? [LATIN SMALL LETTER K WITH STROKE AND DIAGONAL STROKE]
				fallthrough
			case '\uFF4B': // ??? [FULLWIDTH LATIN SMALL LETTER K]
				output[outputPos] = 'k'
				outputPos++

			case '\u24A6': // ??? [PARENTHESIZED LATIN SMALL LETTER K]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'k'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u0139': // ?? [LATIN CAPITAL LETTER L WITH ACUTE]
				fallthrough
			case '\u013B': // ?? [LATIN CAPITAL LETTER L WITH CEDILLA]
				fallthrough
			case '\u013D': // ?? [LATIN CAPITAL LETTER L WITH CARON]
				fallthrough
			case '\u013F': // ?? [LATIN CAPITAL LETTER L WITH MIDDLE DOT]
				fallthrough
			case '\u0141': // ?? [LATIN CAPITAL LETTER L WITH STROKE]
				fallthrough
			case '\u023D': // ?? [LATIN CAPITAL LETTER L WITH BAR]
				fallthrough
			case '\u029F': // ?? [LATIN LETTER SMALL CAPITAL L]
				fallthrough
			case '\u1D0C': // ??? [LATIN LETTER SMALL CAPITAL L WITH STROKE]
				fallthrough
			case '\u1E36': // ??? [LATIN CAPITAL LETTER L WITH DOT BELOW]
				fallthrough
			case '\u1E38': // ??? [LATIN CAPITAL LETTER L WITH DOT BELOW AND MACRON]
				fallthrough
			case '\u1E3A': // ??? [LATIN CAPITAL LETTER L WITH LINE BELOW]
				fallthrough
			case '\u1E3C': // ??? [LATIN CAPITAL LETTER L WITH CIRCUMFLEX BELOW]
				fallthrough
			case '\u24C1': // ??? [CIRCLED LATIN CAPITAL LETTER L]
				fallthrough
			case '\u2C60': // ??? [LATIN CAPITAL LETTER L WITH DOUBLE BAR]
				fallthrough
			case '\u2C62': // ??? [LATIN CAPITAL LETTER L WITH MIDDLE TILDE]
				fallthrough
			case '\uA746': // ??? [LATIN CAPITAL LETTER BROKEN L]
				fallthrough
			case '\uA748': // ??? [LATIN CAPITAL LETTER L WITH HIGH STROKE]
				fallthrough
			case '\uA780': // ??? [LATIN CAPITAL LETTER TURNED L]
				fallthrough
			case '\uFF2C': // ??? [FULLWIDTH LATIN CAPITAL LETTER L]
				output[outputPos] = 'L'
				outputPos++

			case '\u013A': // ?? [LATIN SMALL LETTER L WITH ACUTE]
				fallthrough
			case '\u013C': // ?? [LATIN SMALL LETTER L WITH CEDILLA]
				fallthrough
			case '\u013E': // ?? [LATIN SMALL LETTER L WITH CARON]
				fallthrough
			case '\u0140': // ?? [LATIN SMALL LETTER L WITH MIDDLE DOT]
				fallthrough
			case '\u0142': // ?? [LATIN SMALL LETTER L WITH STROKE]
				fallthrough
			case '\u019A': // ?? [LATIN SMALL LETTER L WITH BAR]
				fallthrough
			case '\u0234': // ?? [LATIN SMALL LETTER L WITH CURL]
				fallthrough
			case '\u026B': // ?? [LATIN SMALL LETTER L WITH MIDDLE TILDE]
				fallthrough
			case '\u026C': // ?? [LATIN SMALL LETTER L WITH BELT]
				fallthrough
			case '\u026D': // ?? [LATIN SMALL LETTER L WITH RETROFLEX HOOK]
				fallthrough
			case '\u1D85': // ??? [LATIN SMALL LETTER L WITH PALATAL HOOK]
				fallthrough
			case '\u1E37': // ??? [LATIN SMALL LETTER L WITH DOT BELOW]
				fallthrough
			case '\u1E39': // ??? [LATIN SMALL LETTER L WITH DOT BELOW AND MACRON]
				fallthrough
			case '\u1E3B': // ??? [LATIN SMALL LETTER L WITH LINE BELOW]
				fallthrough
			case '\u1E3D': // ??? [LATIN SMALL LETTER L WITH CIRCUMFLEX BELOW]
				fallthrough
			case '\u24DB': // ??? [CIRCLED LATIN SMALL LETTER L]
				fallthrough
			case '\u2C61': // ??? [LATIN SMALL LETTER L WITH DOUBLE BAR]
				fallthrough
			case '\uA747': // ??? [LATIN SMALL LETTER BROKEN L]
				fallthrough
			case '\uA749': // ??? [LATIN SMALL LETTER L WITH HIGH STROKE]
				fallthrough
			case '\uA781': // ??? [LATIN SMALL LETTER TURNED L]
				fallthrough
			case '\uFF4C': // ??? [FULLWIDTH LATIN SMALL LETTER L]
				output[outputPos] = 'l'
				outputPos++

			case '\u01C7': // ?? [LATIN CAPITAL LETTER LJ]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'L'
				outputPos++
				output[outputPos] = 'J'
				outputPos++

			case '\u1EFA': // ??? [LATIN CAPITAL LETTER MIDDLE-WELSH LL]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'L'
				outputPos++
				output[outputPos] = 'L'
				outputPos++

			case '\u01C8': // ?? [LATIN CAPITAL LETTER L WITH SMALL LETTER J]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'L'
				outputPos++
				output[outputPos] = 'j'
				outputPos++

			case '\u24A7': // ??? [PARENTHESIZED LATIN SMALL LETTER L]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'l'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u01C9': // ?? [LATIN SMALL LETTER LJ]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'l'
				outputPos++
				output[outputPos] = 'j'
				outputPos++

			case '\u1EFB': // ??? [LATIN SMALL LETTER MIDDLE-WELSH LL]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'l'
				outputPos++
				output[outputPos] = 'l'
				outputPos++

			case '\u02AA': // ?? [LATIN SMALL LETTER LS DIGRAPH]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'l'
				outputPos++
				output[outputPos] = 's'
				outputPos++

			case '\u02AB': // ?? [LATIN SMALL LETTER LZ DIGRAPH]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'l'
				outputPos++
				output[outputPos] = 'z'
				outputPos++

			case '\u019C': // ?? [LATIN CAPITAL LETTER TURNED M]
				fallthrough
			case '\u1D0D': // ??? [LATIN LETTER SMALL CAPITAL M]
				fallthrough
			case '\u1E3E': // ??? [LATIN CAPITAL LETTER M WITH ACUTE]
				fallthrough
			case '\u1E40': // ??? [LATIN CAPITAL LETTER M WITH DOT ABOVE]
				fallthrough
			case '\u1E42': // ??? [LATIN CAPITAL LETTER M WITH DOT BELOW]
				fallthrough
			case '\u24C2': // ??? [CIRCLED LATIN CAPITAL LETTER M]
				fallthrough
			case '\u2C6E': // ??? [LATIN CAPITAL LETTER M WITH HOOK]
				fallthrough
			case '\uA7FD': // ??? [LATIN EPIGRAPHIC LETTER INVERTED M]
				fallthrough
			case '\uA7FF': // ??? [LATIN EPIGRAPHIC LETTER ARCHAIC M]
				fallthrough
			case '\uFF2D': // ??? [FULLWIDTH LATIN CAPITAL LETTER M]
				output[outputPos] = 'M'
				outputPos++

			case '\u026F': // ?? [LATIN SMALL LETTER TURNED M]
				fallthrough
			case '\u0270': // ?? [LATIN SMALL LETTER TURNED M WITH LONG LEG]
				fallthrough
			case '\u0271': // ?? [LATIN SMALL LETTER M WITH HOOK]
				fallthrough
			case '\u1D6F': // ??? [LATIN SMALL LETTER M WITH MIDDLE TILDE]
				fallthrough
			case '\u1D86': // ??? [LATIN SMALL LETTER M WITH PALATAL HOOK]
				fallthrough
			case '\u1E3F': // ??? [LATIN SMALL LETTER M WITH ACUTE]
				fallthrough
			case '\u1E41': // ??? [LATIN SMALL LETTER M WITH DOT ABOVE]
				fallthrough
			case '\u1E43': // ??? [LATIN SMALL LETTER M WITH DOT BELOW]
				fallthrough
			case '\u24DC': // ??? [CIRCLED LATIN SMALL LETTER M]
				fallthrough
			case '\uFF4D': // ??? [FULLWIDTH LATIN SMALL LETTER M]
				output[outputPos] = 'm'
				outputPos++

			case '\u24A8': // ??? [PARENTHESIZED LATIN SMALL LETTER M]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'm'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u00D1': // ?? [LATIN CAPITAL LETTER N WITH TILDE]
				fallthrough
			case '\u0143': // ?? [LATIN CAPITAL LETTER N WITH ACUTE]
				fallthrough
			case '\u0145': // ?? [LATIN CAPITAL LETTER N WITH CEDILLA]
				fallthrough
			case '\u0147': // ?? [LATIN CAPITAL LETTER N WITH CARON]
				fallthrough
			case '\u014A': // ?? http://en.wikipedia.org/wiki/Eng_(letter) [LATIN CAPITAL LETTER ENG]
				fallthrough
			case '\u019D': // ?? [LATIN CAPITAL LETTER N WITH LEFT HOOK]
				fallthrough
			case '\u01F8': // ?? [LATIN CAPITAL LETTER N WITH GRAVE]
				fallthrough
			case '\u0220': // ?? [LATIN CAPITAL LETTER N WITH LONG RIGHT LEG]
				fallthrough
			case '\u0274': // ?? [LATIN LETTER SMALL CAPITAL N]
				fallthrough
			case '\u1D0E': // ??? [LATIN LETTER SMALL CAPITAL REVERSED N]
				fallthrough
			case '\u1E44': // ??? [LATIN CAPITAL LETTER N WITH DOT ABOVE]
				fallthrough
			case '\u1E46': // ??? [LATIN CAPITAL LETTER N WITH DOT BELOW]
				fallthrough
			case '\u1E48': // ??? [LATIN CAPITAL LETTER N WITH LINE BELOW]
				fallthrough
			case '\u1E4A': // ??? [LATIN CAPITAL LETTER N WITH CIRCUMFLEX BELOW]
				fallthrough
			case '\u24C3': // ??? [CIRCLED LATIN CAPITAL LETTER N]
				fallthrough
			case '\uFF2E': // ??? [FULLWIDTH LATIN CAPITAL LETTER N]
				output[outputPos] = 'N'
				outputPos++

			case '\u00F1': // ?? [LATIN SMALL LETTER N WITH TILDE]
				fallthrough
			case '\u0144': // ?? [LATIN SMALL LETTER N WITH ACUTE]
				fallthrough
			case '\u0146': // ?? [LATIN SMALL LETTER N WITH CEDILLA]
				fallthrough
			case '\u0148': // ?? [LATIN SMALL LETTER N WITH CARON]
				fallthrough
			case '\u0149': // ?? [LATIN SMALL LETTER N PRECEDED BY APOSTROPHE]
				fallthrough
			case '\u014B': // ?? http://en.wikipedia.org/wiki/Eng_(letter) [LATIN SMALL LETTER ENG]
				fallthrough
			case '\u019E': // ?? [LATIN SMALL LETTER N WITH LONG RIGHT LEG]
				fallthrough
			case '\u01F9': // ?? [LATIN SMALL LETTER N WITH GRAVE]
				fallthrough
			case '\u0235': // ?? [LATIN SMALL LETTER N WITH CURL]
				fallthrough
			case '\u0272': // ?? [LATIN SMALL LETTER N WITH LEFT HOOK]
				fallthrough
			case '\u0273': // ?? [LATIN SMALL LETTER N WITH RETROFLEX HOOK]
				fallthrough
			case '\u1D70': // ??? [LATIN SMALL LETTER N WITH MIDDLE TILDE]
				fallthrough
			case '\u1D87': // ??? [LATIN SMALL LETTER N WITH PALATAL HOOK]
				fallthrough
			case '\u1E45': // ??? [LATIN SMALL LETTER N WITH DOT ABOVE]
				fallthrough
			case '\u1E47': // ??? [LATIN SMALL LETTER N WITH DOT BELOW]
				fallthrough
			case '\u1E49': // ??? [LATIN SMALL LETTER N WITH LINE BELOW]
				fallthrough
			case '\u1E4B': // ??? [LATIN SMALL LETTER N WITH CIRCUMFLEX BELOW]
				fallthrough
			case '\u207F': // ??? [SUPERSCRIPT LATIN SMALL LETTER N]
				fallthrough
			case '\u24DD': // ??? [CIRCLED LATIN SMALL LETTER N]
				fallthrough
			case '\uFF4E': // ??? [FULLWIDTH LATIN SMALL LETTER N]
				output[outputPos] = 'n'
				outputPos++

			case '\u01CA': // ?? [LATIN CAPITAL LETTER NJ]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'N'
				outputPos++
				output[outputPos] = 'J'
				outputPos++

			case '\u01CB': // ?? [LATIN CAPITAL LETTER N WITH SMALL LETTER J]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'N'
				outputPos++
				output[outputPos] = 'j'
				outputPos++

			case '\u24A9': // ??? [PARENTHESIZED LATIN SMALL LETTER N]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'n'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u01CC': // ?? [LATIN SMALL LETTER NJ]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'n'
				outputPos++
				output[outputPos] = 'j'
				outputPos++

			case '\u00D2': // ?? [LATIN CAPITAL LETTER O WITH GRAVE]
				fallthrough
			case '\u00D3': // ?? [LATIN CAPITAL LETTER O WITH ACUTE]
				fallthrough
			case '\u00D4': // ?? [LATIN CAPITAL LETTER O WITH CIRCUMFLEX]
				fallthrough
			case '\u00D5': // ?? [LATIN CAPITAL LETTER O WITH TILDE]
				fallthrough
			case '\u00D6': // ?? [LATIN CAPITAL LETTER O WITH DIAERESIS]
				fallthrough
			case '\u00D8': // ?? [LATIN CAPITAL LETTER O WITH STROKE]
				fallthrough
			case '\u014C': // ?? [LATIN CAPITAL LETTER O WITH MACRON]
				fallthrough
			case '\u014E': // ?? [LATIN CAPITAL LETTER O WITH BREVE]
				fallthrough
			case '\u0150': // ?? [LATIN CAPITAL LETTER O WITH DOUBLE ACUTE]
				fallthrough
			case '\u0186': // ?? [LATIN CAPITAL LETTER OPEN O]
				fallthrough
			case '\u019F': // ?? [LATIN CAPITAL LETTER O WITH MIDDLE TILDE]
				fallthrough
			case '\u01A0': // ?? [LATIN CAPITAL LETTER O WITH HORN]
				fallthrough
			case '\u01D1': // ?? [LATIN CAPITAL LETTER O WITH CARON]
				fallthrough
			case '\u01EA': // ?? [LATIN CAPITAL LETTER O WITH OGONEK]
				fallthrough
			case '\u01EC': // ?? [LATIN CAPITAL LETTER O WITH OGONEK AND MACRON]
				fallthrough
			case '\u01FE': // ?? [LATIN CAPITAL LETTER O WITH STROKE AND ACUTE]
				fallthrough
			case '\u020C': // ?? [LATIN CAPITAL LETTER O WITH DOUBLE GRAVE]
				fallthrough
			case '\u020E': // ?? [LATIN CAPITAL LETTER O WITH INVERTED BREVE]
				fallthrough
			case '\u022A': // ?? [LATIN CAPITAL LETTER O WITH DIAERESIS AND MACRON]
				fallthrough
			case '\u022C': // ?? [LATIN CAPITAL LETTER O WITH TILDE AND MACRON]
				fallthrough
			case '\u022E': // ?? [LATIN CAPITAL LETTER O WITH DOT ABOVE]
				fallthrough
			case '\u0230': // ?? [LATIN CAPITAL LETTER O WITH DOT ABOVE AND MACRON]
				fallthrough
			case '\u1D0F': // ??? [LATIN LETTER SMALL CAPITAL O]
				fallthrough
			case '\u1D10': // ??? [LATIN LETTER SMALL CAPITAL OPEN O]
				fallthrough
			case '\u1E4C': // ??? [LATIN CAPITAL LETTER O WITH TILDE AND ACUTE]
				fallthrough
			case '\u1E4E': // ??? [LATIN CAPITAL LETTER O WITH TILDE AND DIAERESIS]
				fallthrough
			case '\u1E50': // ??? [LATIN CAPITAL LETTER O WITH MACRON AND GRAVE]
				fallthrough
			case '\u1E52': // ??? [LATIN CAPITAL LETTER O WITH MACRON AND ACUTE]
				fallthrough
			case '\u1ECC': // ??? [LATIN CAPITAL LETTER O WITH DOT BELOW]
				fallthrough
			case '\u1ECE': // ??? [LATIN CAPITAL LETTER O WITH HOOK ABOVE]
				fallthrough
			case '\u1ED0': // ??? [LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND ACUTE]
				fallthrough
			case '\u1ED2': // ??? [LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND GRAVE]
				fallthrough
			case '\u1ED4': // ??? [LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND HOOK ABOVE]
				fallthrough
			case '\u1ED6': // ??? [LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND TILDE]
				fallthrough
			case '\u1ED8': // ??? [LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND DOT BELOW]
				fallthrough
			case '\u1EDA': // ??? [LATIN CAPITAL LETTER O WITH HORN AND ACUTE]
				fallthrough
			case '\u1EDC': // ??? [LATIN CAPITAL LETTER O WITH HORN AND GRAVE]
				fallthrough
			case '\u1EDE': // ??? [LATIN CAPITAL LETTER O WITH HORN AND HOOK ABOVE]
				fallthrough
			case '\u1EE0': // ??? [LATIN CAPITAL LETTER O WITH HORN AND TILDE]
				fallthrough
			case '\u1EE2': // ??? [LATIN CAPITAL LETTER O WITH HORN AND DOT BELOW]
				fallthrough
			case '\u24C4': // ??? [CIRCLED LATIN CAPITAL LETTER O]
				fallthrough
			case '\uA74A': // ??? [LATIN CAPITAL LETTER O WITH LONG STROKE OVERLAY]
				fallthrough
			case '\uA74C': // ??? [LATIN CAPITAL LETTER O WITH LOOP]
				fallthrough
			case '\uFF2F': // ??? [FULLWIDTH LATIN CAPITAL LETTER O]
				output[outputPos] = 'O'
				outputPos++

			case '\u00F2': // ?? [LATIN SMALL LETTER O WITH GRAVE]
				fallthrough
			case '\u00F3': // ?? [LATIN SMALL LETTER O WITH ACUTE]
				fallthrough
			case '\u00F4': // ?? [LATIN SMALL LETTER O WITH CIRCUMFLEX]
				fallthrough
			case '\u00F5': // ?? [LATIN SMALL LETTER O WITH TILDE]
				fallthrough
			case '\u00F6': // ?? [LATIN SMALL LETTER O WITH DIAERESIS]
				fallthrough
			case '\u00F8': // ?? [LATIN SMALL LETTER O WITH STROKE]
				fallthrough
			case '\u014D': // ?? [LATIN SMALL LETTER O WITH MACRON]
				fallthrough
			case '\u014F': // ?? [LATIN SMALL LETTER O WITH BREVE]
				fallthrough
			case '\u0151': // ?? [LATIN SMALL LETTER O WITH DOUBLE ACUTE]
				fallthrough
			case '\u01A1': // ?? [LATIN SMALL LETTER O WITH HORN]
				fallthrough
			case '\u01D2': // ?? [LATIN SMALL LETTER O WITH CARON]
				fallthrough
			case '\u01EB': // ?? [LATIN SMALL LETTER O WITH OGONEK]
				fallthrough
			case '\u01ED': // ?? [LATIN SMALL LETTER O WITH OGONEK AND MACRON]
				fallthrough
			case '\u01FF': // ?? [LATIN SMALL LETTER O WITH STROKE AND ACUTE]
				fallthrough
			case '\u020D': // ?? [LATIN SMALL LETTER O WITH DOUBLE GRAVE]
				fallthrough
			case '\u020F': // ?? [LATIN SMALL LETTER O WITH INVERTED BREVE]
				fallthrough
			case '\u022B': // ?? [LATIN SMALL LETTER O WITH DIAERESIS AND MACRON]
				fallthrough
			case '\u022D': // ?? [LATIN SMALL LETTER O WITH TILDE AND MACRON]
				fallthrough
			case '\u022F': // ?? [LATIN SMALL LETTER O WITH DOT ABOVE]
				fallthrough
			case '\u0231': // ?? [LATIN SMALL LETTER O WITH DOT ABOVE AND MACRON]
				fallthrough
			case '\u0254': // ?? [LATIN SMALL LETTER OPEN O]
				fallthrough
			case '\u0275': // ?? [LATIN SMALL LETTER BARRED O]
				fallthrough
			case '\u1D16': // ??? [LATIN SMALL LETTER TOP HALF O]
				fallthrough
			case '\u1D17': // ??? [LATIN SMALL LETTER BOTTOM HALF O]
				fallthrough
			case '\u1D97': // ??? [LATIN SMALL LETTER OPEN O WITH RETROFLEX HOOK]
				fallthrough
			case '\u1E4D': // ??? [LATIN SMALL LETTER O WITH TILDE AND ACUTE]
				fallthrough
			case '\u1E4F': // ??? [LATIN SMALL LETTER O WITH TILDE AND DIAERESIS]
				fallthrough
			case '\u1E51': // ??? [LATIN SMALL LETTER O WITH MACRON AND GRAVE]
				fallthrough
			case '\u1E53': // ??? [LATIN SMALL LETTER O WITH MACRON AND ACUTE]
				fallthrough
			case '\u1ECD': // ??? [LATIN SMALL LETTER O WITH DOT BELOW]
				fallthrough
			case '\u1ECF': // ??? [LATIN SMALL LETTER O WITH HOOK ABOVE]
				fallthrough
			case '\u1ED1': // ??? [LATIN SMALL LETTER O WITH CIRCUMFLEX AND ACUTE]
				fallthrough
			case '\u1ED3': // ??? [LATIN SMALL LETTER O WITH CIRCUMFLEX AND GRAVE]
				fallthrough
			case '\u1ED5': // ??? [LATIN SMALL LETTER O WITH CIRCUMFLEX AND HOOK ABOVE]
				fallthrough
			case '\u1ED7': // ??? [LATIN SMALL LETTER O WITH CIRCUMFLEX AND TILDE]
				fallthrough
			case '\u1ED9': // ??? [LATIN SMALL LETTER O WITH CIRCUMFLEX AND DOT BELOW]
				fallthrough
			case '\u1EDB': // ??? [LATIN SMALL LETTER O WITH HORN AND ACUTE]
				fallthrough
			case '\u1EDD': // ??? [LATIN SMALL LETTER O WITH HORN AND GRAVE]
				fallthrough
			case '\u1EDF': // ??? [LATIN SMALL LETTER O WITH HORN AND HOOK ABOVE]
				fallthrough
			case '\u1EE1': // ??? [LATIN SMALL LETTER O WITH HORN AND TILDE]
				fallthrough
			case '\u1EE3': // ??? [LATIN SMALL LETTER O WITH HORN AND DOT BELOW]
				fallthrough
			case '\u2092': // ??? [LATIN SUBSCRIPT SMALL LETTER O]
				fallthrough
			case '\u24DE': // ??? [CIRCLED LATIN SMALL LETTER O]
				fallthrough
			case '\u2C7A': // ??? [LATIN SMALL LETTER O WITH LOW RING INSIDE]
				fallthrough
			case '\uA74B': // ??? [LATIN SMALL LETTER O WITH LONG STROKE OVERLAY]
				fallthrough
			case '\uA74D': // ??? [LATIN SMALL LETTER O WITH LOOP]
				fallthrough
			case '\uFF4F': // ??? [FULLWIDTH LATIN SMALL LETTER O]
				output[outputPos] = 'o'
				outputPos++

			case '\u0152': // ?? [LATIN CAPITAL LIGATURE OE]
				fallthrough
			case '\u0276': // ?? [LATIN LETTER SMALL CAPITAL OE]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'O'
				outputPos++
				output[outputPos] = 'E'
				outputPos++

			case '\uA74E': // ??? [LATIN CAPITAL LETTER OO]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'O'
				outputPos++
				output[outputPos] = 'O'
				outputPos++

			case '\u0222': // ?? http://en.wikipedia.org/wiki/OU [LATIN CAPITAL LETTER OU]
				fallthrough
			case '\u1D15': // ??? [LATIN LETTER SMALL CAPITAL OU]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'O'
				outputPos++
				output[outputPos] = 'U'
				outputPos++

			case '\u24AA': // ??? [PARENTHESIZED LATIN SMALL LETTER O]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'o'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u0153': // ?? [LATIN SMALL LIGATURE OE]
				fallthrough
			case '\u1D14': // ??? [LATIN SMALL LETTER TURNED OE]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'o'
				outputPos++
				output[outputPos] = 'e'
				outputPos++

			case '\uA74F': // ??? [LATIN SMALL LETTER OO]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'o'
				outputPos++
				output[outputPos] = 'o'
				outputPos++

			case '\u0223': // ?? http://en.wikipedia.org/wiki/OU [LATIN SMALL LETTER OU]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'o'
				outputPos++
				output[outputPos] = 'u'
				outputPos++

			case '\u01A4': // ?? [LATIN CAPITAL LETTER P WITH HOOK]
				fallthrough
			case '\u1D18': // ??? [LATIN LETTER SMALL CAPITAL P]
				fallthrough
			case '\u1E54': // ??? [LATIN CAPITAL LETTER P WITH ACUTE]
				fallthrough
			case '\u1E56': // ??? [LATIN CAPITAL LETTER P WITH DOT ABOVE]
				fallthrough
			case '\u24C5': // ??? [CIRCLED LATIN CAPITAL LETTER P]
				fallthrough
			case '\u2C63': // ??? [LATIN CAPITAL LETTER P WITH STROKE]
				fallthrough
			case '\uA750': // ??? [LATIN CAPITAL LETTER P WITH STROKE THROUGH DESCENDER]
				fallthrough
			case '\uA752': // ??? [LATIN CAPITAL LETTER P WITH FLOURISH]
				fallthrough
			case '\uA754': // ??? [LATIN CAPITAL LETTER P WITH SQUIRREL TAIL]
				fallthrough
			case '\uFF30': // ??? [FULLWIDTH LATIN CAPITAL LETTER P]
				output[outputPos] = 'P'
				outputPos++

			case '\u01A5': // ?? [LATIN SMALL LETTER P WITH HOOK]
				fallthrough
			case '\u1D71': // ??? [LATIN SMALL LETTER P WITH MIDDLE TILDE]
				fallthrough
			case '\u1D7D': // ??? [LATIN SMALL LETTER P WITH STROKE]
				fallthrough
			case '\u1D88': // ??? [LATIN SMALL LETTER P WITH PALATAL HOOK]
				fallthrough
			case '\u1E55': // ??? [LATIN SMALL LETTER P WITH ACUTE]
				fallthrough
			case '\u1E57': // ??? [LATIN SMALL LETTER P WITH DOT ABOVE]
				fallthrough
			case '\u24DF': // ??? [CIRCLED LATIN SMALL LETTER P]
				fallthrough
			case '\uA751': // ??? [LATIN SMALL LETTER P WITH STROKE THROUGH DESCENDER]
				fallthrough
			case '\uA753': // ??? [LATIN SMALL LETTER P WITH FLOURISH]
				fallthrough
			case '\uA755': // ??? [LATIN SMALL LETTER P WITH SQUIRREL TAIL]
				fallthrough
			case '\uA7FC': // ??? [LATIN EPIGRAPHIC LETTER REVERSED P]
				fallthrough
			case '\uFF50': // ??? [FULLWIDTH LATIN SMALL LETTER P]
				output[outputPos] = 'p'
				outputPos++

			case '\u24AB': // ??? [PARENTHESIZED LATIN SMALL LETTER P]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'p'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u024A': // ?? [LATIN CAPITAL LETTER SMALL Q WITH HOOK TAIL]
				fallthrough
			case '\u24C6': // ??? [CIRCLED LATIN CAPITAL LETTER Q]
				fallthrough
			case '\uA756': // ??? [LATIN CAPITAL LETTER Q WITH STROKE THROUGH DESCENDER]
				fallthrough
			case '\uA758': // ??? [LATIN CAPITAL LETTER Q WITH DIAGONAL STROKE]
				fallthrough
			case '\uFF31': // ??? [FULLWIDTH LATIN CAPITAL LETTER Q]
				output[outputPos] = 'Q'
				outputPos++

			case '\u0138': // ?? http://en.wikipedia.org/wiki/Kra_(letter) [LATIN SMALL LETTER KRA]
				fallthrough
			case '\u024B': // ?? [LATIN SMALL LETTER Q WITH HOOK TAIL]
				fallthrough
			case '\u02A0': // ?? [LATIN SMALL LETTER Q WITH HOOK]
				fallthrough
			case '\u24E0': // ??? [CIRCLED LATIN SMALL LETTER Q]
				fallthrough
			case '\uA757': // ??? [LATIN SMALL LETTER Q WITH STROKE THROUGH DESCENDER]
				fallthrough
			case '\uA759': // ??? [LATIN SMALL LETTER Q WITH DIAGONAL STROKE]
				fallthrough
			case '\uFF51': // ??? [FULLWIDTH LATIN SMALL LETTER Q]
				output[outputPos] = 'q'
				outputPos++

			case '\u24AC': // ??? [PARENTHESIZED LATIN SMALL LETTER Q]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'q'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u0239': // ?? [LATIN SMALL LETTER QP DIGRAPH]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'q'
				outputPos++
				output[outputPos] = 'p'
				outputPos++

			case '\u0154': // ?? [LATIN CAPITAL LETTER R WITH ACUTE]
				fallthrough
			case '\u0156': // ?? [LATIN CAPITAL LETTER R WITH CEDILLA]
				fallthrough
			case '\u0158': // ?? [LATIN CAPITAL LETTER R WITH CARON]
				fallthrough
			case '\u0210': // ?? [LATIN CAPITAL LETTER R WITH DOUBLE GRAVE]
				fallthrough
			case '\u0212': // ?? [LATIN CAPITAL LETTER R WITH INVERTED BREVE]
				fallthrough
			case '\u024C': // ?? [LATIN CAPITAL LETTER R WITH STROKE]
				fallthrough
			case '\u0280': // ?? [LATIN LETTER SMALL CAPITAL R]
				fallthrough
			case '\u0281': // ?? [LATIN LETTER SMALL CAPITAL INVERTED R]
				fallthrough
			case '\u1D19': // ??? [LATIN LETTER SMALL CAPITAL REVERSED R]
				fallthrough
			case '\u1D1A': // ??? [LATIN LETTER SMALL CAPITAL TURNED R]
				fallthrough
			case '\u1E58': // ??? [LATIN CAPITAL LETTER R WITH DOT ABOVE]
				fallthrough
			case '\u1E5A': // ??? [LATIN CAPITAL LETTER R WITH DOT BELOW]
				fallthrough
			case '\u1E5C': // ??? [LATIN CAPITAL LETTER R WITH DOT BELOW AND MACRON]
				fallthrough
			case '\u1E5E': // ??? [LATIN CAPITAL LETTER R WITH LINE BELOW]
				fallthrough
			case '\u24C7': // ??? [CIRCLED LATIN CAPITAL LETTER R]
				fallthrough
			case '\u2C64': // ??? [LATIN CAPITAL LETTER R WITH TAIL]
				fallthrough
			case '\uA75A': // ??? [LATIN CAPITAL LETTER R ROTUNDA]
				fallthrough
			case '\uA782': // ??? [LATIN CAPITAL LETTER INSULAR R]
				fallthrough
			case '\uFF32': // ??? [FULLWIDTH LATIN CAPITAL LETTER R]
				output[outputPos] = 'R'
				outputPos++

			case '\u0155': // ?? [LATIN SMALL LETTER R WITH ACUTE]
				fallthrough
			case '\u0157': // ?? [LATIN SMALL LETTER R WITH CEDILLA]
				fallthrough
			case '\u0159': // ?? [LATIN SMALL LETTER R WITH CARON]
				fallthrough
			case '\u0211': // ?? [LATIN SMALL LETTER R WITH DOUBLE GRAVE]
				fallthrough
			case '\u0213': // ?? [LATIN SMALL LETTER R WITH INVERTED BREVE]
				fallthrough
			case '\u024D': // ?? [LATIN SMALL LETTER R WITH STROKE]
				fallthrough
			case '\u027C': // ?? [LATIN SMALL LETTER R WITH LONG LEG]
				fallthrough
			case '\u027D': // ?? [LATIN SMALL LETTER R WITH TAIL]
				fallthrough
			case '\u027E': // ?? [LATIN SMALL LETTER R WITH FISHHOOK]
				fallthrough
			case '\u027F': // ?? [LATIN SMALL LETTER REVERSED R WITH FISHHOOK]
				fallthrough
			case '\u1D63': // ??? [LATIN SUBSCRIPT SMALL LETTER R]
				fallthrough
			case '\u1D72': // ??? [LATIN SMALL LETTER R WITH MIDDLE TILDE]
				fallthrough
			case '\u1D73': // ??? [LATIN SMALL LETTER R WITH FISHHOOK AND MIDDLE TILDE]
				fallthrough
			case '\u1D89': // ??? [LATIN SMALL LETTER R WITH PALATAL HOOK]
				fallthrough
			case '\u1E59': // ??? [LATIN SMALL LETTER R WITH DOT ABOVE]
				fallthrough
			case '\u1E5B': // ??? [LATIN SMALL LETTER R WITH DOT BELOW]
				fallthrough
			case '\u1E5D': // ??? [LATIN SMALL LETTER R WITH DOT BELOW AND MACRON]
				fallthrough
			case '\u1E5F': // ??? [LATIN SMALL LETTER R WITH LINE BELOW]
				fallthrough
			case '\u24E1': // ??? [CIRCLED LATIN SMALL LETTER R]
				fallthrough
			case '\uA75B': // ??? [LATIN SMALL LETTER R ROTUNDA]
				fallthrough
			case '\uA783': // ??? [LATIN SMALL LETTER INSULAR R]
				fallthrough
			case '\uFF52': // ??? [FULLWIDTH LATIN SMALL LETTER R]
				output[outputPos] = 'r'
				outputPos++

			case '\u24AD': // ??? [PARENTHESIZED LATIN SMALL LETTER R]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'r'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u015A': // ?? [LATIN CAPITAL LETTER S WITH ACUTE]
				fallthrough
			case '\u015C': // ?? [LATIN CAPITAL LETTER S WITH CIRCUMFLEX]
				fallthrough
			case '\u015E': // ?? [LATIN CAPITAL LETTER S WITH CEDILLA]
				fallthrough
			case '\u0160': // ?? [LATIN CAPITAL LETTER S WITH CARON]
				fallthrough
			case '\u0218': // ?? [LATIN CAPITAL LETTER S WITH COMMA BELOW]
				fallthrough
			case '\u1E60': // ??? [LATIN CAPITAL LETTER S WITH DOT ABOVE]
				fallthrough
			case '\u1E62': // ??? [LATIN CAPITAL LETTER S WITH DOT BELOW]
				fallthrough
			case '\u1E64': // ??? [LATIN CAPITAL LETTER S WITH ACUTE AND DOT ABOVE]
				fallthrough
			case '\u1E66': // ??? [LATIN CAPITAL LETTER S WITH CARON AND DOT ABOVE]
				fallthrough
			case '\u1E68': // ??? [LATIN CAPITAL LETTER S WITH DOT BELOW AND DOT ABOVE]
				fallthrough
			case '\u24C8': // ??? [CIRCLED LATIN CAPITAL LETTER S]
				fallthrough
			case '\uA731': // ??? [LATIN LETTER SMALL CAPITAL S]
				fallthrough
			case '\uA785': // ??? [LATIN SMALL LETTER INSULAR S]
				fallthrough
			case '\uFF33': // ??? [FULLWIDTH LATIN CAPITAL LETTER S]
				output[outputPos] = 'S'
				outputPos++

			case '\u015B': // ?? [LATIN SMALL LETTER S WITH ACUTE]
				fallthrough
			case '\u015D': // ?? [LATIN SMALL LETTER S WITH CIRCUMFLEX]
				fallthrough
			case '\u015F': // ?? [LATIN SMALL LETTER S WITH CEDILLA]
				fallthrough
			case '\u0161': // ?? [LATIN SMALL LETTER S WITH CARON]
				fallthrough
			case '\u017F': // ?? http://en.wikipedia.org/wiki/Long_S [LATIN SMALL LETTER LONG S]
				fallthrough
			case '\u0219': // ?? [LATIN SMALL LETTER S WITH COMMA BELOW]
				fallthrough
			case '\u023F': // ?? [LATIN SMALL LETTER S WITH SWASH TAIL]
				fallthrough
			case '\u0282': // ?? [LATIN SMALL LETTER S WITH HOOK]
				fallthrough
			case '\u1D74': // ??? [LATIN SMALL LETTER S WITH MIDDLE TILDE]
				fallthrough
			case '\u1D8A': // ??? [LATIN SMALL LETTER S WITH PALATAL HOOK]
				fallthrough
			case '\u1E61': // ??? [LATIN SMALL LETTER S WITH DOT ABOVE]
				fallthrough
			case '\u1E63': // ??? [LATIN SMALL LETTER S WITH DOT BELOW]
				fallthrough
			case '\u1E65': // ??? [LATIN SMALL LETTER S WITH ACUTE AND DOT ABOVE]
				fallthrough
			case '\u1E67': // ??? [LATIN SMALL LETTER S WITH CARON AND DOT ABOVE]
				fallthrough
			case '\u1E69': // ??? [LATIN SMALL LETTER S WITH DOT BELOW AND DOT ABOVE]
				fallthrough
			case '\u1E9C': // ??? [LATIN SMALL LETTER LONG S WITH DIAGONAL STROKE]
				fallthrough
			case '\u1E9D': // ??? [LATIN SMALL LETTER LONG S WITH HIGH STROKE]
				fallthrough
			case '\u24E2': // ??? [CIRCLED LATIN SMALL LETTER S]
				fallthrough
			case '\uA784': // ??? [LATIN CAPITAL LETTER INSULAR S]
				fallthrough
			case '\uFF53': // ??? [FULLWIDTH LATIN SMALL LETTER S]
				output[outputPos] = 's'
				outputPos++

			case '\u1E9E': // ??? [LATIN CAPITAL LETTER SHARP S]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'S'
				outputPos++
				output[outputPos] = 'S'
				outputPos++

			case '\u24AE': // ??? [PARENTHESIZED LATIN SMALL LETTER S]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 's'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u00DF': // ?? [LATIN SMALL LETTER SHARP S]
				output = output[:(len(output) + 1)]
				output[outputPos] = 's'
				outputPos++
				output[outputPos] = 's'
				outputPos++

			case '\uFB06': // ??? [LATIN SMALL LIGATURE ST]
				output = output[:(len(output) + 1)]
				output[outputPos] = 's'
				outputPos++
				output[outputPos] = 't'
				outputPos++

			case '\u0162': // ?? [LATIN CAPITAL LETTER T WITH CEDILLA]
				fallthrough
			case '\u0164': // ?? [LATIN CAPITAL LETTER T WITH CARON]
				fallthrough
			case '\u0166': // ?? [LATIN CAPITAL LETTER T WITH STROKE]
				fallthrough
			case '\u01AC': // ?? [LATIN CAPITAL LETTER T WITH HOOK]
				fallthrough
			case '\u01AE': // ?? [LATIN CAPITAL LETTER T WITH RETROFLEX HOOK]
				fallthrough
			case '\u021A': // ?? [LATIN CAPITAL LETTER T WITH COMMA BELOW]
				fallthrough
			case '\u023E': // ?? [LATIN CAPITAL LETTER T WITH DIAGONAL STROKE]
				fallthrough
			case '\u1D1B': // ??? [LATIN LETTER SMALL CAPITAL T]
				fallthrough
			case '\u1E6A': // ??? [LATIN CAPITAL LETTER T WITH DOT ABOVE]
				fallthrough
			case '\u1E6C': // ??? [LATIN CAPITAL LETTER T WITH DOT BELOW]
				fallthrough
			case '\u1E6E': // ??? [LATIN CAPITAL LETTER T WITH LINE BELOW]
				fallthrough
			case '\u1E70': // ??? [LATIN CAPITAL LETTER T WITH CIRCUMFLEX BELOW]
				fallthrough
			case '\u24C9': // ??? [CIRCLED LATIN CAPITAL LETTER T]
				fallthrough
			case '\uA786': // ??? [LATIN CAPITAL LETTER INSULAR T]
				fallthrough
			case '\uFF34': // ??? [FULLWIDTH LATIN CAPITAL LETTER T]
				output[outputPos] = 'T'
				outputPos++

			case '\u0163': // ?? [LATIN SMALL LETTER T WITH CEDILLA]
				fallthrough
			case '\u0165': // ?? [LATIN SMALL LETTER T WITH CARON]
				fallthrough
			case '\u0167': // ?? [LATIN SMALL LETTER T WITH STROKE]
				fallthrough
			case '\u01AB': // ?? [LATIN SMALL LETTER T WITH PALATAL HOOK]
				fallthrough
			case '\u01AD': // ?? [LATIN SMALL LETTER T WITH HOOK]
				fallthrough
			case '\u021B': // ?? [LATIN SMALL LETTER T WITH COMMA BELOW]
				fallthrough
			case '\u0236': // ?? [LATIN SMALL LETTER T WITH CURL]
				fallthrough
			case '\u0287': // ?? [LATIN SMALL LETTER TURNED T]
				fallthrough
			case '\u0288': // ?? [LATIN SMALL LETTER T WITH RETROFLEX HOOK]
				fallthrough
			case '\u1D75': // ??? [LATIN SMALL LETTER T WITH MIDDLE TILDE]
				fallthrough
			case '\u1E6B': // ??? [LATIN SMALL LETTER T WITH DOT ABOVE]
				fallthrough
			case '\u1E6D': // ??? [LATIN SMALL LETTER T WITH DOT BELOW]
				fallthrough
			case '\u1E6F': // ??? [LATIN SMALL LETTER T WITH LINE BELOW]
				fallthrough
			case '\u1E71': // ??? [LATIN SMALL LETTER T WITH CIRCUMFLEX BELOW]
				fallthrough
			case '\u1E97': // ??? [LATIN SMALL LETTER T WITH DIAERESIS]
				fallthrough
			case '\u24E3': // ??? [CIRCLED LATIN SMALL LETTER T]
				fallthrough
			case '\u2C66': // ??? [LATIN SMALL LETTER T WITH DIAGONAL STROKE]
				fallthrough
			case '\uFF54': // ??? [FULLWIDTH LATIN SMALL LETTER T]
				output[outputPos] = 't'
				outputPos++

			case '\u00DE': // ?? [LATIN CAPITAL LETTER THORN]
				fallthrough
			case '\uA766': // ??? [LATIN CAPITAL LETTER THORN WITH STROKE THROUGH DESCENDER]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'T'
				outputPos++
				output[outputPos] = 'H'
				outputPos++

			case '\uA728': // ??? [LATIN CAPITAL LETTER TZ]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'T'
				outputPos++
				output[outputPos] = 'Z'
				outputPos++

			case '\u24AF': // ??? [PARENTHESIZED LATIN SMALL LETTER T]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 't'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u02A8': // ?? [LATIN SMALL LETTER TC DIGRAPH WITH CURL]
				output = output[:(len(output) + 1)]
				output[outputPos] = 't'
				outputPos++
				output[outputPos] = 'c'
				outputPos++

			case '\u00FE': // ?? [LATIN SMALL LETTER THORN]
				fallthrough
			case '\u1D7A': // ??? [LATIN SMALL LETTER TH WITH STRIKETHROUGH]
				fallthrough
			case '\uA767': // ??? [LATIN SMALL LETTER THORN WITH STROKE THROUGH DESCENDER]
				output = output[:(len(output) + 1)]
				output[outputPos] = 't'
				outputPos++
				output[outputPos] = 'h'
				outputPos++

			case '\u02A6': // ?? [LATIN SMALL LETTER TS DIGRAPH]
				output = output[:(len(output) + 1)]
				output[outputPos] = 't'
				outputPos++
				output[outputPos] = 's'
				outputPos++

			case '\uA729': // ??? [LATIN SMALL LETTER TZ]
				output = output[:(len(output) + 1)]
				output[outputPos] = 't'
				outputPos++
				output[outputPos] = 'z'
				outputPos++

			case '\u00D9': // ?? [LATIN CAPITAL LETTER U WITH GRAVE]
				fallthrough
			case '\u00DA': // ?? [LATIN CAPITAL LETTER U WITH ACUTE]
				fallthrough
			case '\u00DB': // ?? [LATIN CAPITAL LETTER U WITH CIRCUMFLEX]
				fallthrough
			case '\u00DC': // ?? [LATIN CAPITAL LETTER U WITH DIAERESIS]
				fallthrough
			case '\u0168': // ?? [LATIN CAPITAL LETTER U WITH TILDE]
				fallthrough
			case '\u016A': // ?? [LATIN CAPITAL LETTER U WITH MACRON]
				fallthrough
			case '\u016C': // ?? [LATIN CAPITAL LETTER U WITH BREVE]
				fallthrough
			case '\u016E': // ?? [LATIN CAPITAL LETTER U WITH RING ABOVE]
				fallthrough
			case '\u0170': // ?? [LATIN CAPITAL LETTER U WITH DOUBLE ACUTE]
				fallthrough
			case '\u0172': // ?? [LATIN CAPITAL LETTER U WITH OGONEK]
				fallthrough
			case '\u01AF': // ?? [LATIN CAPITAL LETTER U WITH HORN]
				fallthrough
			case '\u01D3': // ?? [LATIN CAPITAL LETTER U WITH CARON]
				fallthrough
			case '\u01D5': // ?? [LATIN CAPITAL LETTER U WITH DIAERESIS AND MACRON]
				fallthrough
			case '\u01D7': // ?? [LATIN CAPITAL LETTER U WITH DIAERESIS AND ACUTE]
				fallthrough
			case '\u01D9': // ?? [LATIN CAPITAL LETTER U WITH DIAERESIS AND CARON]
				fallthrough
			case '\u01DB': // ?? [LATIN CAPITAL LETTER U WITH DIAERESIS AND GRAVE]
				fallthrough
			case '\u0214': // ?? [LATIN CAPITAL LETTER U WITH DOUBLE GRAVE]
				fallthrough
			case '\u0216': // ?? [LATIN CAPITAL LETTER U WITH INVERTED BREVE]
				fallthrough
			case '\u0244': // ?? [LATIN CAPITAL LETTER U BAR]
				fallthrough
			case '\u1D1C': // ??? [LATIN LETTER SMALL CAPITAL U]
				fallthrough
			case '\u1D7E': // ??? [LATIN SMALL CAPITAL LETTER U WITH STROKE]
				fallthrough
			case '\u1E72': // ??? [LATIN CAPITAL LETTER U WITH DIAERESIS BELOW]
				fallthrough
			case '\u1E74': // ??? [LATIN CAPITAL LETTER U WITH TILDE BELOW]
				fallthrough
			case '\u1E76': // ??? [LATIN CAPITAL LETTER U WITH CIRCUMFLEX BELOW]
				fallthrough
			case '\u1E78': // ??? [LATIN CAPITAL LETTER U WITH TILDE AND ACUTE]
				fallthrough
			case '\u1E7A': // ??? [LATIN CAPITAL LETTER U WITH MACRON AND DIAERESIS]
				fallthrough
			case '\u1EE4': // ??? [LATIN CAPITAL LETTER U WITH DOT BELOW]
				fallthrough
			case '\u1EE6': // ??? [LATIN CAPITAL LETTER U WITH HOOK ABOVE]
				fallthrough
			case '\u1EE8': // ??? [LATIN CAPITAL LETTER U WITH HORN AND ACUTE]
				fallthrough
			case '\u1EEA': // ??? [LATIN CAPITAL LETTER U WITH HORN AND GRAVE]
				fallthrough
			case '\u1EEC': // ??? [LATIN CAPITAL LETTER U WITH HORN AND HOOK ABOVE]
				fallthrough
			case '\u1EEE': // ??? [LATIN CAPITAL LETTER U WITH HORN AND TILDE]
				fallthrough
			case '\u1EF0': // ??? [LATIN CAPITAL LETTER U WITH HORN AND DOT BELOW]
				fallthrough
			case '\u24CA': // ??? [CIRCLED LATIN CAPITAL LETTER U]
				fallthrough
			case '\uFF35': // ??? [FULLWIDTH LATIN CAPITAL LETTER U]
				output[outputPos] = 'U'
				outputPos++

			case '\u00F9': // ?? [LATIN SMALL LETTER U WITH GRAVE]
				fallthrough
			case '\u00FA': // ?? [LATIN SMALL LETTER U WITH ACUTE]
				fallthrough
			case '\u00FB': // ?? [LATIN SMALL LETTER U WITH CIRCUMFLEX]
				fallthrough
			case '\u00FC': // ?? [LATIN SMALL LETTER U WITH DIAERESIS]
				fallthrough
			case '\u0169': // ?? [LATIN SMALL LETTER U WITH TILDE]
				fallthrough
			case '\u016B': // ?? [LATIN SMALL LETTER U WITH MACRON]
				fallthrough
			case '\u016D': // ?? [LATIN SMALL LETTER U WITH BREVE]
				fallthrough
			case '\u016F': // ?? [LATIN SMALL LETTER U WITH RING ABOVE]
				fallthrough
			case '\u0171': // ?? [LATIN SMALL LETTER U WITH DOUBLE ACUTE]
				fallthrough
			case '\u0173': // ?? [LATIN SMALL LETTER U WITH OGONEK]
				fallthrough
			case '\u01B0': // ?? [LATIN SMALL LETTER U WITH HORN]
				fallthrough
			case '\u01D4': // ?? [LATIN SMALL LETTER U WITH CARON]
				fallthrough
			case '\u01D6': // ?? [LATIN SMALL LETTER U WITH DIAERESIS AND MACRON]
				fallthrough
			case '\u01D8': // ?? [LATIN SMALL LETTER U WITH DIAERESIS AND ACUTE]
				fallthrough
			case '\u01DA': // ?? [LATIN SMALL LETTER U WITH DIAERESIS AND CARON]
				fallthrough
			case '\u01DC': // ?? [LATIN SMALL LETTER U WITH DIAERESIS AND GRAVE]
				fallthrough
			case '\u0215': // ?? [LATIN SMALL LETTER U WITH DOUBLE GRAVE]
				fallthrough
			case '\u0217': // ?? [LATIN SMALL LETTER U WITH INVERTED BREVE]
				fallthrough
			case '\u0289': // ?? [LATIN SMALL LETTER U BAR]
				fallthrough
			case '\u1D64': // ??? [LATIN SUBSCRIPT SMALL LETTER U]
				fallthrough
			case '\u1D99': // ??? [LATIN SMALL LETTER U WITH RETROFLEX HOOK]
				fallthrough
			case '\u1E73': // ??? [LATIN SMALL LETTER U WITH DIAERESIS BELOW]
				fallthrough
			case '\u1E75': // ??? [LATIN SMALL LETTER U WITH TILDE BELOW]
				fallthrough
			case '\u1E77': // ??? [LATIN SMALL LETTER U WITH CIRCUMFLEX BELOW]
				fallthrough
			case '\u1E79': // ??? [LATIN SMALL LETTER U WITH TILDE AND ACUTE]
				fallthrough
			case '\u1E7B': // ??? [LATIN SMALL LETTER U WITH MACRON AND DIAERESIS]
				fallthrough
			case '\u1EE5': // ??? [LATIN SMALL LETTER U WITH DOT BELOW]
				fallthrough
			case '\u1EE7': // ??? [LATIN SMALL LETTER U WITH HOOK ABOVE]
				fallthrough
			case '\u1EE9': // ??? [LATIN SMALL LETTER U WITH HORN AND ACUTE]
				fallthrough
			case '\u1EEB': // ??? [LATIN SMALL LETTER U WITH HORN AND GRAVE]
				fallthrough
			case '\u1EED': // ??? [LATIN SMALL LETTER U WITH HORN AND HOOK ABOVE]
				fallthrough
			case '\u1EEF': // ??? [LATIN SMALL LETTER U WITH HORN AND TILDE]
				fallthrough
			case '\u1EF1': // ??? [LATIN SMALL LETTER U WITH HORN AND DOT BELOW]
				fallthrough
			case '\u24E4': // ??? [CIRCLED LATIN SMALL LETTER U]
				fallthrough
			case '\uFF55': // ??? [FULLWIDTH LATIN SMALL LETTER U]
				output[outputPos] = 'u'
				outputPos++

			case '\u24B0': // ??? [PARENTHESIZED LATIN SMALL LETTER U]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'u'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u1D6B': // ??? [LATIN SMALL LETTER UE]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'u'
				outputPos++
				output[outputPos] = 'e'
				outputPos++

			case '\u01B2': // ?? [LATIN CAPITAL LETTER V WITH HOOK]
				fallthrough
			case '\u0245': // ?? [LATIN CAPITAL LETTER TURNED V]
				fallthrough
			case '\u1D20': // ??? [LATIN LETTER SMALL CAPITAL V]
				fallthrough
			case '\u1E7C': // ??? [LATIN CAPITAL LETTER V WITH TILDE]
				fallthrough
			case '\u1E7E': // ??? [LATIN CAPITAL LETTER V WITH DOT BELOW]
				fallthrough
			case '\u1EFC': // ??? [LATIN CAPITAL LETTER MIDDLE-WELSH V]
				fallthrough
			case '\u24CB': // ??? [CIRCLED LATIN CAPITAL LETTER V]
				fallthrough
			case '\uA75E': // ??? [LATIN CAPITAL LETTER V WITH DIAGONAL STROKE]
				fallthrough
			case '\uA768': // ??? [LATIN CAPITAL LETTER VEND]
				fallthrough
			case '\uFF36': // ??? [FULLWIDTH LATIN CAPITAL LETTER V]
				output[outputPos] = 'V'
				outputPos++

			case '\u028B': // ?? [LATIN SMALL LETTER V WITH HOOK]
				fallthrough
			case '\u028C': // ?? [LATIN SMALL LETTER TURNED V]
				fallthrough
			case '\u1D65': // ??? [LATIN SUBSCRIPT SMALL LETTER V]
				fallthrough
			case '\u1D8C': // ??? [LATIN SMALL LETTER V WITH PALATAL HOOK]
				fallthrough
			case '\u1E7D': // ??? [LATIN SMALL LETTER V WITH TILDE]
				fallthrough
			case '\u1E7F': // ??? [LATIN SMALL LETTER V WITH DOT BELOW]
				fallthrough
			case '\u24E5': // ??? [CIRCLED LATIN SMALL LETTER V]
				fallthrough
			case '\u2C71': // ??? [LATIN SMALL LETTER V WITH RIGHT HOOK]
				fallthrough
			case '\u2C74': // ??? [LATIN SMALL LETTER V WITH CURL]
				fallthrough
			case '\uA75F': // ??? [LATIN SMALL LETTER V WITH DIAGONAL STROKE]
				fallthrough
			case '\uFF56': // ??? [FULLWIDTH LATIN SMALL LETTER V]
				output[outputPos] = 'v'
				outputPos++

			case '\uA760': // ??? [LATIN CAPITAL LETTER VY]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'V'
				outputPos++
				output[outputPos] = 'Y'
				outputPos++

			case '\u24B1': // ??? [PARENTHESIZED LATIN SMALL LETTER V]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'v'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\uA761': // ??? [LATIN SMALL LETTER VY]
				output = output[:(len(output) + 1)]
				output[outputPos] = 'v'
				outputPos++
				output[outputPos] = 'y'
				outputPos++

			case '\u0174': // ?? [LATIN CAPITAL LETTER W WITH CIRCUMFLEX]
				fallthrough
			case '\u01F7': // ?? http://en.wikipedia.org/wiki/Wynn [LATIN CAPITAL LETTER WYNN]
				fallthrough
			case '\u1D21': // ??? [LATIN LETTER SMALL CAPITAL W]
				fallthrough
			case '\u1E80': // ??? [LATIN CAPITAL LETTER W WITH GRAVE]
				fallthrough
			case '\u1E82': // ??? [LATIN CAPITAL LETTER W WITH ACUTE]
				fallthrough
			case '\u1E84': // ??? [LATIN CAPITAL LETTER W WITH DIAERESIS]
				fallthrough
			case '\u1E86': // ??? [LATIN CAPITAL LETTER W WITH DOT ABOVE]
				fallthrough
			case '\u1E88': // ??? [LATIN CAPITAL LETTER W WITH DOT BELOW]
				fallthrough
			case '\u24CC': // ??? [CIRCLED LATIN CAPITAL LETTER W]
				fallthrough
			case '\u2C72': // ??? [LATIN CAPITAL LETTER W WITH HOOK]
				fallthrough
			case '\uFF37': // ??? [FULLWIDTH LATIN CAPITAL LETTER W]
				output[outputPos] = 'W'
				outputPos++

			case '\u0175': // ?? [LATIN SMALL LETTER W WITH CIRCUMFLEX]
				fallthrough
			case '\u01BF': // ?? http://en.wikipedia.org/wiki/Wynn [LATIN LETTER WYNN]
				fallthrough
			case '\u028D': // ?? [LATIN SMALL LETTER TURNED W]
				fallthrough
			case '\u1E81': // ??? [LATIN SMALL LETTER W WITH GRAVE]
				fallthrough
			case '\u1E83': // ??? [LATIN SMALL LETTER W WITH ACUTE]
				fallthrough
			case '\u1E85': // ??? [LATIN SMALL LETTER W WITH DIAERESIS]
				fallthrough
			case '\u1E87': // ??? [LATIN SMALL LETTER W WITH DOT ABOVE]
				fallthrough
			case '\u1E89': // ??? [LATIN SMALL LETTER W WITH DOT BELOW]
				fallthrough
			case '\u1E98': // ??? [LATIN SMALL LETTER W WITH RING ABOVE]
				fallthrough
			case '\u24E6': // ??? [CIRCLED LATIN SMALL LETTER W]
				fallthrough
			case '\u2C73': // ??? [LATIN SMALL LETTER W WITH HOOK]
				fallthrough
			case '\uFF57': // ??? [FULLWIDTH LATIN SMALL LETTER W]
				output[outputPos] = 'w'
				outputPos++

			case '\u24B2': // ??? [PARENTHESIZED LATIN SMALL LETTER W]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'w'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u1E8A': // ??? [LATIN CAPITAL LETTER X WITH DOT ABOVE]
				fallthrough
			case '\u1E8C': // ??? [LATIN CAPITAL LETTER X WITH DIAERESIS]
				fallthrough
			case '\u24CD': // ??? [CIRCLED LATIN CAPITAL LETTER X]
				fallthrough
			case '\uFF38': // ??? [FULLWIDTH LATIN CAPITAL LETTER X]
				output[outputPos] = 'X'
				outputPos++

			case '\u1D8D': // ??? [LATIN SMALL LETTER X WITH PALATAL HOOK]
				fallthrough
			case '\u1E8B': // ??? [LATIN SMALL LETTER X WITH DOT ABOVE]
				fallthrough
			case '\u1E8D': // ??? [LATIN SMALL LETTER X WITH DIAERESIS]
				fallthrough
			case '\u2093': // ??? [LATIN SUBSCRIPT SMALL LETTER X]
				fallthrough
			case '\u24E7': // ??? [CIRCLED LATIN SMALL LETTER X]
				fallthrough
			case '\uFF58': // ??? [FULLWIDTH LATIN SMALL LETTER X]
				output[outputPos] = 'x'
				outputPos++

			case '\u24B3': // ??? [PARENTHESIZED LATIN SMALL LETTER X]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'x'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u00DD': // ?? [LATIN CAPITAL LETTER Y WITH ACUTE]
				fallthrough
			case '\u0176': // ?? [LATIN CAPITAL LETTER Y WITH CIRCUMFLEX]
				fallthrough
			case '\u0178': // ?? [LATIN CAPITAL LETTER Y WITH DIAERESIS]
				fallthrough
			case '\u01B3': // ?? [LATIN CAPITAL LETTER Y WITH HOOK]
				fallthrough
			case '\u0232': // ?? [LATIN CAPITAL LETTER Y WITH MACRON]
				fallthrough
			case '\u024E': // ?? [LATIN CAPITAL LETTER Y WITH STROKE]
				fallthrough
			case '\u028F': // ?? [LATIN LETTER SMALL CAPITAL Y]
				fallthrough
			case '\u1E8E': // ??? [LATIN CAPITAL LETTER Y WITH DOT ABOVE]
				fallthrough
			case '\u1EF2': // ??? [LATIN CAPITAL LETTER Y WITH GRAVE]
				fallthrough
			case '\u1EF4': // ??? [LATIN CAPITAL LETTER Y WITH DOT BELOW]
				fallthrough
			case '\u1EF6': // ??? [LATIN CAPITAL LETTER Y WITH HOOK ABOVE]
				fallthrough
			case '\u1EF8': // ??? [LATIN CAPITAL LETTER Y WITH TILDE]
				fallthrough
			case '\u1EFE': // ??? [LATIN CAPITAL LETTER Y WITH LOOP]
				fallthrough
			case '\u24CE': // ??? [CIRCLED LATIN CAPITAL LETTER Y]
				fallthrough
			case '\uFF39': // ??? [FULLWIDTH LATIN CAPITAL LETTER Y]
				output[outputPos] = 'Y'
				outputPos++

			case '\u00FD': // ?? [LATIN SMALL LETTER Y WITH ACUTE]
				fallthrough
			case '\u00FF': // ?? [LATIN SMALL LETTER Y WITH DIAERESIS]
				fallthrough
			case '\u0177': // ?? [LATIN SMALL LETTER Y WITH CIRCUMFLEX]
				fallthrough
			case '\u01B4': // ?? [LATIN SMALL LETTER Y WITH HOOK]
				fallthrough
			case '\u0233': // ?? [LATIN SMALL LETTER Y WITH MACRON]
				fallthrough
			case '\u024F': // ?? [LATIN SMALL LETTER Y WITH STROKE]
				fallthrough
			case '\u028E': // ?? [LATIN SMALL LETTER TURNED Y]
				fallthrough
			case '\u1E8F': // ??? [LATIN SMALL LETTER Y WITH DOT ABOVE]
				fallthrough
			case '\u1E99': // ??? [LATIN SMALL LETTER Y WITH RING ABOVE]
				fallthrough
			case '\u1EF3': // ??? [LATIN SMALL LETTER Y WITH GRAVE]
				fallthrough
			case '\u1EF5': // ??? [LATIN SMALL LETTER Y WITH DOT BELOW]
				fallthrough
			case '\u1EF7': // ??? [LATIN SMALL LETTER Y WITH HOOK ABOVE]
				fallthrough
			case '\u1EF9': // ??? [LATIN SMALL LETTER Y WITH TILDE]
				fallthrough
			case '\u1EFF': // ??? [LATIN SMALL LETTER Y WITH LOOP]
				fallthrough
			case '\u24E8': // ??? [CIRCLED LATIN SMALL LETTER Y]
				fallthrough
			case '\uFF59': // ??? [FULLWIDTH LATIN SMALL LETTER Y]
				output[outputPos] = 'y'
				outputPos++

			case '\u24B4': // ??? [PARENTHESIZED LATIN SMALL LETTER Y]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'y'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u0179': // ?? [LATIN CAPITAL LETTER Z WITH ACUTE]
				fallthrough
			case '\u017B': // ?? [LATIN CAPITAL LETTER Z WITH DOT ABOVE]
				fallthrough
			case '\u017D': // ?? [LATIN CAPITAL LETTER Z WITH CARON]
				fallthrough
			case '\u01B5': // ?? [LATIN CAPITAL LETTER Z WITH STROKE]
				fallthrough
			case '\u021C': // ?? http://en.wikipedia.org/wiki/Yogh [LATIN CAPITAL LETTER YOGH]
				fallthrough
			case '\u0224': // ?? [LATIN CAPITAL LETTER Z WITH HOOK]
				fallthrough
			case '\u1D22': // ??? [LATIN LETTER SMALL CAPITAL Z]
				fallthrough
			case '\u1E90': // ??? [LATIN CAPITAL LETTER Z WITH CIRCUMFLEX]
				fallthrough
			case '\u1E92': // ??? [LATIN CAPITAL LETTER Z WITH DOT BELOW]
				fallthrough
			case '\u1E94': // ??? [LATIN CAPITAL LETTER Z WITH LINE BELOW]
				fallthrough
			case '\u24CF': // ??? [CIRCLED LATIN CAPITAL LETTER Z]
				fallthrough
			case '\u2C6B': // ??? [LATIN CAPITAL LETTER Z WITH DESCENDER]
				fallthrough
			case '\uA762': // ??? [LATIN CAPITAL LETTER VISIGOTHIC Z]
				fallthrough
			case '\uFF3A': // ??? [FULLWIDTH LATIN CAPITAL LETTER Z]
				output[outputPos] = 'Z'
				outputPos++

			case '\u017A': // ?? [LATIN SMALL LETTER Z WITH ACUTE]
				fallthrough
			case '\u017C': // ?? [LATIN SMALL LETTER Z WITH DOT ABOVE]
				fallthrough
			case '\u017E': // ?? [LATIN SMALL LETTER Z WITH CARON]
				fallthrough
			case '\u01B6': // ?? [LATIN SMALL LETTER Z WITH STROKE]
				fallthrough
			case '\u021D': // ?? http://en.wikipedia.org/wiki/Yogh [LATIN SMALL LETTER YOGH]
				fallthrough
			case '\u0225': // ?? [LATIN SMALL LETTER Z WITH HOOK]
				fallthrough
			case '\u0240': // ?? [LATIN SMALL LETTER Z WITH SWASH TAIL]
				fallthrough
			case '\u0290': // ?? [LATIN SMALL LETTER Z WITH RETROFLEX HOOK]
				fallthrough
			case '\u0291': // ?? [LATIN SMALL LETTER Z WITH CURL]
				fallthrough
			case '\u1D76': // ??? [LATIN SMALL LETTER Z WITH MIDDLE TILDE]
				fallthrough
			case '\u1D8E': // ??? [LATIN SMALL LETTER Z WITH PALATAL HOOK]
				fallthrough
			case '\u1E91': // ??? [LATIN SMALL LETTER Z WITH CIRCUMFLEX]
				fallthrough
			case '\u1E93': // ??? [LATIN SMALL LETTER Z WITH DOT BELOW]
				fallthrough
			case '\u1E95': // ??? [LATIN SMALL LETTER Z WITH LINE BELOW]
				fallthrough
			case '\u24E9': // ??? [CIRCLED LATIN SMALL LETTER Z]
				fallthrough
			case '\u2C6C': // ??? [LATIN SMALL LETTER Z WITH DESCENDER]
				fallthrough
			case '\uA763': // ??? [LATIN SMALL LETTER VISIGOTHIC Z]
				fallthrough
			case '\uFF5A': // ??? [FULLWIDTH LATIN SMALL LETTER Z]
				output[outputPos] = 'z'
				outputPos++

			case '\u24B5': // ??? [PARENTHESIZED LATIN SMALL LETTER Z]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = 'z'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u2070': // ??? [SUPERSCRIPT ZERO]
				fallthrough
			case '\u2080': // ??? [SUBSCRIPT ZERO]
				fallthrough
			case '\u24EA': // ??? [CIRCLED DIGIT ZERO]
				fallthrough
			case '\u24FF': // ??? [NEGATIVE CIRCLED DIGIT ZERO]
				fallthrough
			case '\uFF10': // ??? [FULLWIDTH DIGIT ZERO]
				output[outputPos] = '0'
				outputPos++

			case '\u00B9': // ?? [SUPERSCRIPT ONE]
				fallthrough
			case '\u2081': // ??? [SUBSCRIPT ONE]
				fallthrough
			case '\u2460': // ??? [CIRCLED DIGIT ONE]
				fallthrough
			case '\u24F5': // ??? [DOUBLE CIRCLED DIGIT ONE]
				fallthrough
			case '\u2776': // ??? [DINGBAT NEGATIVE CIRCLED DIGIT ONE]
				fallthrough
			case '\u2780': // ??? [DINGBAT CIRCLED SANS-SERIF DIGIT ONE]
				fallthrough
			case '\u278A': // ??? [DINGBAT NEGATIVE CIRCLED SANS-SERIF DIGIT ONE]
				fallthrough
			case '\uFF11': // ??? [FULLWIDTH DIGIT ONE]
				output[outputPos] = '1'
				outputPos++

			case '\u2488': // ??? [DIGIT ONE FULL STOP]
				output = output[:(len(output) + 1)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2474': // ??? [PARENTHESIZED DIGIT ONE]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u00B2': // ?? [SUPERSCRIPT TWO]
				fallthrough
			case '\u2082': // ??? [SUBSCRIPT TWO]
				fallthrough
			case '\u2461': // ??? [CIRCLED DIGIT TWO]
				fallthrough
			case '\u24F6': // ??? [DOUBLE CIRCLED DIGIT TWO]
				fallthrough
			case '\u2777': // ??? [DINGBAT NEGATIVE CIRCLED DIGIT TWO]
				fallthrough
			case '\u2781': // ??? [DINGBAT CIRCLED SANS-SERIF DIGIT TWO]
				fallthrough
			case '\u278B': // ??? [DINGBAT NEGATIVE CIRCLED SANS-SERIF DIGIT TWO]
				fallthrough
			case '\uFF12': // ??? [FULLWIDTH DIGIT TWO]
				output[outputPos] = '2'
				outputPos++

			case '\u2489': // ??? [DIGIT TWO FULL STOP]
				output = output[:(len(output) + 1)]
				output[outputPos] = '2'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2475': // ??? [PARENTHESIZED DIGIT TWO]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '2'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u00B3': // ?? [SUPERSCRIPT THREE]
				fallthrough
			case '\u2083': // ??? [SUBSCRIPT THREE]
				fallthrough
			case '\u2462': // ??? [CIRCLED DIGIT THREE]
				fallthrough
			case '\u24F7': // ??? [DOUBLE CIRCLED DIGIT THREE]
				fallthrough
			case '\u2778': // ??? [DINGBAT NEGATIVE CIRCLED DIGIT THREE]
				fallthrough
			case '\u2782': // ??? [DINGBAT CIRCLED SANS-SERIF DIGIT THREE]
				fallthrough
			case '\u278C': // ??? [DINGBAT NEGATIVE CIRCLED SANS-SERIF DIGIT THREE]
				fallthrough
			case '\uFF13': // ??? [FULLWIDTH DIGIT THREE]
				output[outputPos] = '3'
				outputPos++

			case '\u248A': // ??? [DIGIT THREE FULL STOP]
				output = output[:(len(output) + 1)]
				output[outputPos] = '3'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2476': // ??? [PARENTHESIZED DIGIT THREE]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '3'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u2074': // ??? [SUPERSCRIPT FOUR]
				fallthrough
			case '\u2084': // ??? [SUBSCRIPT FOUR]
				fallthrough
			case '\u2463': // ??? [CIRCLED DIGIT FOUR]
				fallthrough
			case '\u24F8': // ??? [DOUBLE CIRCLED DIGIT FOUR]
				fallthrough
			case '\u2779': // ??? [DINGBAT NEGATIVE CIRCLED DIGIT FOUR]
				fallthrough
			case '\u2783': // ??? [DINGBAT CIRCLED SANS-SERIF DIGIT FOUR]
				fallthrough
			case '\u278D': // ??? [DINGBAT NEGATIVE CIRCLED SANS-SERIF DIGIT FOUR]
				fallthrough
			case '\uFF14': // ??? [FULLWIDTH DIGIT FOUR]
				output[outputPos] = '4'
				outputPos++

			case '\u248B': // ??? [DIGIT FOUR FULL STOP]
				output = output[:(len(output) + 1)]
				output[outputPos] = '4'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2477': // ??? [PARENTHESIZED DIGIT FOUR]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '4'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u2075': // ??? [SUPERSCRIPT FIVE]
				fallthrough
			case '\u2085': // ??? [SUBSCRIPT FIVE]
				fallthrough
			case '\u2464': // ??? [CIRCLED DIGIT FIVE]
				fallthrough
			case '\u24F9': // ??? [DOUBLE CIRCLED DIGIT FIVE]
				fallthrough
			case '\u277A': // ??? [DINGBAT NEGATIVE CIRCLED DIGIT FIVE]
				fallthrough
			case '\u2784': // ??? [DINGBAT CIRCLED SANS-SERIF DIGIT FIVE]
				fallthrough
			case '\u278E': // ??? [DINGBAT NEGATIVE CIRCLED SANS-SERIF DIGIT FIVE]
				fallthrough
			case '\uFF15': // ??? [FULLWIDTH DIGIT FIVE]
				output[outputPos] = '5'
				outputPos++

			case '\u248C': // ??? [DIGIT FIVE FULL STOP]
				output = output[:(len(output) + 1)]
				output[outputPos] = '5'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2478': // ??? [PARENTHESIZED DIGIT FIVE]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '5'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u2076': // ??? [SUPERSCRIPT SIX]
				fallthrough
			case '\u2086': // ??? [SUBSCRIPT SIX]
				fallthrough
			case '\u2465': // ??? [CIRCLED DIGIT SIX]
				fallthrough
			case '\u24FA': // ??? [DOUBLE CIRCLED DIGIT SIX]
				fallthrough
			case '\u277B': // ??? [DINGBAT NEGATIVE CIRCLED DIGIT SIX]
				fallthrough
			case '\u2785': // ??? [DINGBAT CIRCLED SANS-SERIF DIGIT SIX]
				fallthrough
			case '\u278F': // ??? [DINGBAT NEGATIVE CIRCLED SANS-SERIF DIGIT SIX]
				fallthrough
			case '\uFF16': // ??? [FULLWIDTH DIGIT SIX]
				output[outputPos] = '6'
				outputPos++

			case '\u248D': // ??? [DIGIT SIX FULL STOP]
				output = output[:(len(output) + 1)]
				output[outputPos] = '6'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2479': // ??? [PARENTHESIZED DIGIT SIX]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '6'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u2077': // ??? [SUPERSCRIPT SEVEN]
				fallthrough
			case '\u2087': // ??? [SUBSCRIPT SEVEN]
				fallthrough
			case '\u2466': // ??? [CIRCLED DIGIT SEVEN]
				fallthrough
			case '\u24FB': // ??? [DOUBLE CIRCLED DIGIT SEVEN]
				fallthrough
			case '\u277C': // ??? [DINGBAT NEGATIVE CIRCLED DIGIT SEVEN]
				fallthrough
			case '\u2786': // ??? [DINGBAT CIRCLED SANS-SERIF DIGIT SEVEN]
				fallthrough
			case '\u2790': // ??? [DINGBAT NEGATIVE CIRCLED SANS-SERIF DIGIT SEVEN]
				fallthrough
			case '\uFF17': // ??? [FULLWIDTH DIGIT SEVEN]
				output[outputPos] = '7'
				outputPos++

			case '\u248E': // ??? [DIGIT SEVEN FULL STOP]
				output = output[:(len(output) + 1)]
				output[outputPos] = '7'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u247A': // ??? [PARENTHESIZED DIGIT SEVEN]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '7'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u2078': // ??? [SUPERSCRIPT EIGHT]
				fallthrough
			case '\u2088': // ??? [SUBSCRIPT EIGHT]
				fallthrough
			case '\u2467': // ??? [CIRCLED DIGIT EIGHT]
				fallthrough
			case '\u24FC': // ??? [DOUBLE CIRCLED DIGIT EIGHT]
				fallthrough
			case '\u277D': // ??? [DINGBAT NEGATIVE CIRCLED DIGIT EIGHT]
				fallthrough
			case '\u2787': // ??? [DINGBAT CIRCLED SANS-SERIF DIGIT EIGHT]
				fallthrough
			case '\u2791': // ??? [DINGBAT NEGATIVE CIRCLED SANS-SERIF DIGIT EIGHT]
				fallthrough
			case '\uFF18': // ??? [FULLWIDTH DIGIT EIGHT]
				output[outputPos] = '8'
				outputPos++

			case '\u248F': // ??? [DIGIT EIGHT FULL STOP]
				output = output[:(len(output) + 1)]
				output[outputPos] = '8'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u247B': // ??? [PARENTHESIZED DIGIT EIGHT]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '8'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u2079': // ??? [SUPERSCRIPT NINE]
				fallthrough
			case '\u2089': // ??? [SUBSCRIPT NINE]
				fallthrough
			case '\u2468': // ??? [CIRCLED DIGIT NINE]
				fallthrough
			case '\u24FD': // ??? [DOUBLE CIRCLED DIGIT NINE]
				fallthrough
			case '\u277E': // ??? [DINGBAT NEGATIVE CIRCLED DIGIT NINE]
				fallthrough
			case '\u2788': // ??? [DINGBAT CIRCLED SANS-SERIF DIGIT NINE]
				fallthrough
			case '\u2792': // ??? [DINGBAT NEGATIVE CIRCLED SANS-SERIF DIGIT NINE]
				fallthrough
			case '\uFF19': // ??? [FULLWIDTH DIGIT NINE]
				output[outputPos] = '9'
				outputPos++

			case '\u2490': // ??? [DIGIT NINE FULL STOP]
				output = output[:(len(output) + 1)]
				output[outputPos] = '9'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u247C': // ??? [PARENTHESIZED DIGIT NINE]
				output = output[:(len(output) + 2)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '9'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u2469': // ??? [CIRCLED NUMBER TEN]
				fallthrough
			case '\u24FE': // ??? [DOUBLE CIRCLED NUMBER TEN]
				fallthrough
			case '\u277F': // ??? [DINGBAT NEGATIVE CIRCLED NUMBER TEN]
				fallthrough
			case '\u2789': // ??? [DINGBAT CIRCLED SANS-SERIF NUMBER TEN]
				fallthrough
			case '\u2793': // ??? [DINGBAT NEGATIVE CIRCLED SANS-SERIF NUMBER TEN]
				output = output[:(len(output) + 1)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '0'
				outputPos++

			case '\u2491': // ??? [NUMBER TEN FULL STOP]
				output = output[:(len(output) + 2)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '0'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u247D': // ??? [PARENTHESIZED NUMBER TEN]
				output = output[:(len(output) + 3)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '0'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u246A': // ??? [CIRCLED NUMBER ELEVEN]
				fallthrough
			case '\u24EB': // ??? [NEGATIVE CIRCLED NUMBER ELEVEN]
				output = output[:(len(output) + 1)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '1'
				outputPos++

			case '\u2492': // ??? [NUMBER ELEVEN FULL STOP]
				output = output[:(len(output) + 2)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u247E': // ??? [PARENTHESIZED NUMBER ELEVEN]
				output = output[:(len(output) + 3)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u246B': // ??? [CIRCLED NUMBER TWELVE]
				fallthrough
			case '\u24EC': // ??? [NEGATIVE CIRCLED NUMBER TWELVE]
				output = output[:(len(output) + 1)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '2'
				outputPos++

			case '\u2493': // ??? [NUMBER TWELVE FULL STOP]
				output = output[:(len(output) + 2)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '2'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u247F': // ??? [PARENTHESIZED NUMBER TWELVE]
				output = output[:(len(output) + 3)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '2'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u246C': // ??? [CIRCLED NUMBER THIRTEEN]
				fallthrough
			case '\u24ED': // ??? [NEGATIVE CIRCLED NUMBER THIRTEEN]
				output = output[:(len(output) + 1)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '3'
				outputPos++

			case '\u2494': // ??? [NUMBER THIRTEEN FULL STOP]
				output = output[:(len(output) + 2)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '3'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2480': // ??? [PARENTHESIZED NUMBER THIRTEEN]
				output = output[:(len(output) + 3)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '3'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u246D': // ??? [CIRCLED NUMBER FOURTEEN]
				fallthrough
			case '\u24EE': // ??? [NEGATIVE CIRCLED NUMBER FOURTEEN]
				output = output[:(len(output) + 1)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '4'
				outputPos++

			case '\u2495': // ??? [NUMBER FOURTEEN FULL STOP]
				output = output[:(len(output) + 2)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '4'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2481': // ??? [PARENTHESIZED NUMBER FOURTEEN]
				output = output[:(len(output) + 3)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '4'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u246E': // ??? [CIRCLED NUMBER FIFTEEN]
				fallthrough
			case '\u24EF': // ??? [NEGATIVE CIRCLED NUMBER FIFTEEN]
				output = output[:(len(output) + 1)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '5'
				outputPos++

			case '\u2496': // ??? [NUMBER FIFTEEN FULL STOP]
				output = output[:(len(output) + 2)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '5'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2482': // ??? [PARENTHESIZED NUMBER FIFTEEN]
				output = output[:(len(output) + 3)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '5'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u246F': // ??? [CIRCLED NUMBER SIXTEEN]
				fallthrough
			case '\u24F0': // ??? [NEGATIVE CIRCLED NUMBER SIXTEEN]
				output = output[:(len(output) + 1)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '6'
				outputPos++

			case '\u2497': // ??? [NUMBER SIXTEEN FULL STOP]
				output = output[:(len(output) + 2)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '6'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2483': // ??? [PARENTHESIZED NUMBER SIXTEEN]
				output = output[:(len(output) + 3)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '6'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u2470': // ??? [CIRCLED NUMBER SEVENTEEN]
				fallthrough
			case '\u24F1': // ??? [NEGATIVE CIRCLED NUMBER SEVENTEEN]
				output = output[:(len(output) + 1)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '7'
				outputPos++

			case '\u2498': // ??? [NUMBER SEVENTEEN FULL STOP]
				output = output[:(len(output) + 2)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '7'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2484': // ??? [PARENTHESIZED NUMBER SEVENTEEN]
				output = output[:(len(output) + 3)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '7'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u2471': // ??? [CIRCLED NUMBER EIGHTEEN]
				fallthrough
			case '\u24F2': // ??? [NEGATIVE CIRCLED NUMBER EIGHTEEN]
				output = output[:(len(output) + 1)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '8'
				outputPos++

			case '\u2499': // ??? [NUMBER EIGHTEEN FULL STOP]
				output = output[:(len(output) + 2)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '8'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2485': // ??? [PARENTHESIZED NUMBER EIGHTEEN]
				output = output[:(len(output) + 3)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '8'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u2472': // ??? [CIRCLED NUMBER NINETEEN]
				fallthrough
			case '\u24F3': // ??? [NEGATIVE CIRCLED NUMBER NINETEEN]
				output = output[:(len(output) + 1)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '9'
				outputPos++

			case '\u249A': // ??? [NUMBER NINETEEN FULL STOP]
				output = output[:(len(output) + 2)]
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '9'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2486': // ??? [PARENTHESIZED NUMBER NINETEEN]
				output = output[:(len(output) + 3)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '1'
				outputPos++
				output[outputPos] = '9'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u2473': // ??? [CIRCLED NUMBER TWENTY]
				fallthrough
			case '\u24F4': // ??? [NEGATIVE CIRCLED NUMBER TWENTY]
				output = output[:(len(output) + 1)]
				output[outputPos] = '2'
				outputPos++
				output[outputPos] = '0'
				outputPos++

			case '\u249B': // ??? [NUMBER TWENTY FULL STOP]
				output = output[:(len(output) + 2)]
				output[outputPos] = '2'
				outputPos++
				output[outputPos] = '0'
				outputPos++
				output[outputPos] = '.'
				outputPos++

			case '\u2487': // ??? [PARENTHESIZED NUMBER TWENTY]
				output = output[:(len(output) + 3)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '2'
				outputPos++
				output[outputPos] = '0'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u00AB': // ?? [LEFT-POINTING DOUBLE ANGLE QUOTATION MARK]
				fallthrough
			case '\u00BB': // ?? [RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK]
				fallthrough
			case '\u201C': // ??? [LEFT DOUBLE QUOTATION MARK]
				fallthrough
			case '\u201D': // ??? [RIGHT DOUBLE QUOTATION MARK]
				fallthrough
			case '\u201E': // ??? [DOUBLE LOW-9 QUOTATION MARK]
				fallthrough
			case '\u2033': // ??? [DOUBLE PRIME]
				fallthrough
			case '\u2036': // ??? [REVERSED DOUBLE PRIME]
				fallthrough
			case '\u275D': // ??? [HEAVY DOUBLE TURNED COMMA QUOTATION MARK ORNAMENT]
				fallthrough
			case '\u275E': // ??? [HEAVY DOUBLE COMMA QUOTATION MARK ORNAMENT]
				fallthrough
			case '\u276E': // ??? [HEAVY LEFT-POINTING ANGLE QUOTATION MARK ORNAMENT]
				fallthrough
			case '\u276F': // ??? [HEAVY RIGHT-POINTING ANGLE QUOTATION MARK ORNAMENT]
				fallthrough
			case '\uFF02': // ??? [FULLWIDTH QUOTATION MARK]
				output[outputPos] = '"'
				outputPos++

			case '\u2018': // ??? [LEFT SINGLE QUOTATION MARK]
				fallthrough
			case '\u2019': // ??? [RIGHT SINGLE QUOTATION MARK]
				fallthrough
			case '\u201A': // ??? [SINGLE LOW-9 QUOTATION MARK]
				fallthrough
			case '\u201B': // ??? [SINGLE HIGH-REVERSED-9 QUOTATION MARK]
				fallthrough
			case '\u2032': // ??? [PRIME]
				fallthrough
			case '\u2035': // ??? [REVERSED PRIME]
				fallthrough
			case '\u2039': // ??? [SINGLE LEFT-POINTING ANGLE QUOTATION MARK]
				fallthrough
			case '\u203A': // ??? [SINGLE RIGHT-POINTING ANGLE QUOTATION MARK]
				fallthrough
			case '\u275B': // ??? [HEAVY SINGLE TURNED COMMA QUOTATION MARK ORNAMENT]
				fallthrough
			case '\u275C': // ??? [HEAVY SINGLE COMMA QUOTATION MARK ORNAMENT]
				fallthrough
			case '\uFF07': // ??? [FULLWIDTH APOSTROPHE]
				output[outputPos] = '\''
				outputPos++

			case '\u2010': // ??? [HYPHEN]
				fallthrough
			case '\u2011': // ??? [NON-BREAKING HYPHEN]
				fallthrough
			case '\u2012': // ??? [FIGURE DASH]
				fallthrough
			case '\u2013': // ??? [EN DASH]
				fallthrough
			case '\u2014': // ??? [EM DASH]
				fallthrough
			case '\u207B': // ??? [SUPERSCRIPT MINUS]
				fallthrough
			case '\u208B': // ??? [SUBSCRIPT MINUS]
				fallthrough
			case '\uFF0D': // ??? [FULLWIDTH HYPHEN-MINUS]
				output[outputPos] = '-'
				outputPos++

			case '\u2045': // ??? [LEFT SQUARE BRACKET WITH QUILL]
				fallthrough
			case '\u2772': // ??? [LIGHT LEFT TORTOISE SHELL BRACKET ORNAMENT]
				fallthrough
			case '\uFF3B': // ??? [FULLWIDTH LEFT SQUARE BRACKET]
				output[outputPos] = '['
				outputPos++

			case '\u2046': // ??? [RIGHT SQUARE BRACKET WITH QUILL]
				fallthrough
			case '\u2773': // ??? [LIGHT RIGHT TORTOISE SHELL BRACKET ORNAMENT]
				fallthrough
			case '\uFF3D': // ??? [FULLWIDTH RIGHT SQUARE BRACKET]
				output[outputPos] = ']'
				outputPos++

			case '\u207D': // ??? [SUPERSCRIPT LEFT PARENTHESIS]
				fallthrough
			case '\u208D': // ??? [SUBSCRIPT LEFT PARENTHESIS]
				fallthrough
			case '\u2768': // ??? [MEDIUM LEFT PARENTHESIS ORNAMENT]
				fallthrough
			case '\u276A': // ??? [MEDIUM FLATTENED LEFT PARENTHESIS ORNAMENT]
				fallthrough
			case '\uFF08': // ??? [FULLWIDTH LEFT PARENTHESIS]
				output[outputPos] = '('
				outputPos++

			case '\u2E28': // ??? [LEFT DOUBLE PARENTHESIS]
				output = output[:(len(output) + 1)]
				output[outputPos] = '('
				outputPos++
				output[outputPos] = '('
				outputPos++

			case '\u207E': // ??? [SUPERSCRIPT RIGHT PARENTHESIS]
				fallthrough
			case '\u208E': // ??? [SUBSCRIPT RIGHT PARENTHESIS]
				fallthrough
			case '\u2769': // ??? [MEDIUM RIGHT PARENTHESIS ORNAMENT]
				fallthrough
			case '\u276B': // ??? [MEDIUM FLATTENED RIGHT PARENTHESIS ORNAMENT]
				fallthrough
			case '\uFF09': // ??? [FULLWIDTH RIGHT PARENTHESIS]
				output[outputPos] = ')'
				outputPos++

			case '\u2E29': // ??? [RIGHT DOUBLE PARENTHESIS]
				output = output[:(len(output) + 1)]
				output[outputPos] = ')'
				outputPos++
				output[outputPos] = ')'
				outputPos++

			case '\u276C': // ??? [MEDIUM LEFT-POINTING ANGLE BRACKET ORNAMENT]
				fallthrough
			case '\u2770': // ??? [HEAVY LEFT-POINTING ANGLE BRACKET ORNAMENT]
				fallthrough
			case '\uFF1C': // ??? [FULLWIDTH LESS-THAN SIGN]
				output[outputPos] = '<'
				outputPos++

			case '\u276D': // ??? [MEDIUM RIGHT-POINTING ANGLE BRACKET ORNAMENT]
				fallthrough
			case '\u2771': // ??? [HEAVY RIGHT-POINTING ANGLE BRACKET ORNAMENT]
				fallthrough
			case '\uFF1E': // ??? [FULLWIDTH GREATER-THAN SIGN]
				output[outputPos] = '>'
				outputPos++

			case '\u2774': // ??? [MEDIUM LEFT CURLY BRACKET ORNAMENT]
				fallthrough
			case '\uFF5B': // ??? [FULLWIDTH LEFT CURLY BRACKET]
				output[outputPos] = '{'
				outputPos++

			case '\u2775': // ??? [MEDIUM RIGHT CURLY BRACKET ORNAMENT]
				fallthrough
			case '\uFF5D': // ??? [FULLWIDTH RIGHT CURLY BRACKET]
				output[outputPos] = '}'
				outputPos++

			case '\u207A': // ??? [SUPERSCRIPT PLUS SIGN]
				fallthrough
			case '\u208A': // ??? [SUBSCRIPT PLUS SIGN]
				fallthrough
			case '\uFF0B': // ??? [FULLWIDTH PLUS SIGN]
				output[outputPos] = '+'
				outputPos++

			case '\u207C': // ??? [SUPERSCRIPT EQUALS SIGN]
				fallthrough
			case '\u208C': // ??? [SUBSCRIPT EQUALS SIGN]
				fallthrough
			case '\uFF1D': // ??? [FULLWIDTH EQUALS SIGN]
				output[outputPos] = '='
				outputPos++

			case '\uFF01': // ??? [FULLWIDTH EXCLAMATION MARK]
				output[outputPos] = '!'
				outputPos++

			case '\u203C': // ??? [DOUBLE EXCLAMATION MARK]
				output = output[:(len(output) + 1)]
				output[outputPos] = '!'
				outputPos++
				output[outputPos] = '!'
				outputPos++

			case '\u2049': // ??? [EXCLAMATION QUESTION MARK]
				output = output[:(len(output) + 1)]
				output[outputPos] = '!'
				outputPos++
				output[outputPos] = '?'
				outputPos++

			case '\uFF03': // ??? [FULLWIDTH NUMBER SIGN]
				output[outputPos] = '#'
				outputPos++

			case '\uFF04': // ??? [FULLWIDTH DOLLAR SIGN]
				output[outputPos] = '$'
				outputPos++

			case '\u2052': // ??? [COMMERCIAL MINUS SIGN]
				fallthrough
			case '\uFF05': // ??? [FULLWIDTH PERCENT SIGN]
				output[outputPos] = '%'
				outputPos++

			case '\uFF06': // ??? [FULLWIDTH AMPERSAND]
				output[outputPos] = '&'
				outputPos++

			case '\u204E': // ??? [LOW ASTERISK]
				fallthrough
			case '\uFF0A': // ??? [FULLWIDTH ASTERISK]
				output[outputPos] = '*'
				outputPos++

			case '\uFF0C': // ??? [FULLWIDTH COMMA]
				output[outputPos] = ','
				outputPos++

			case '\uFF0E': // ??? [FULLWIDTH FULL STOP]
				output[outputPos] = '.'
				outputPos++

			case '\u2044': // ??? [FRACTION SLASH]
				fallthrough
			case '\uFF0F': // ??? [FULLWIDTH SOLIDUS]
				output[outputPos] = '/'
				outputPos++

			case '\uFF1A': // ??? [FULLWIDTH COLON]
				output[outputPos] = ':'
				outputPos++

			case '\u204F': // ??? [REVERSED SEMICOLON]
				fallthrough
			case '\uFF1B': // ??? [FULLWIDTH SEMICOLON]
				output[outputPos] = ';'
				outputPos++

			case '\uFF1F': // ??? [FULLWIDTH QUESTION MARK]
				output[outputPos] = '?'
				outputPos++

			case '\u2047': // ??? [DOUBLE QUESTION MARK]
				output = output[:(len(output) + 1)]
				output[outputPos] = '?'
				outputPos++
				output[outputPos] = '?'
				outputPos++

			case '\u2048': // ??? [QUESTION EXCLAMATION MARK]
				output = output[:(len(output) + 1)]
				output[outputPos] = '?'
				outputPos++
				output[outputPos] = '!'
				outputPos++

			case '\uFF20': // ??? [FULLWIDTH COMMERCIAL AT]
				output[outputPos] = '@'
				outputPos++

			case '\uFF3C': // ??? [FULLWIDTH REVERSE SOLIDUS]
				output[outputPos] = '\\'
				outputPos++

			case '\u2038': // ??? [CARET]
				fallthrough
			case '\uFF3E': // ??? [FULLWIDTH CIRCUMFLEX ACCENT]
				output[outputPos] = '^'
				outputPos++

			case '\uFF3F': // ??? [FULLWIDTH LOW LINE]
				output[outputPos] = '_'
				outputPos++

			case '\u2053': // ??? [SWUNG DASH]
				fallthrough
			case '\uFF5E': // ??? [FULLWIDTH TILDE]
				output[outputPos] = '~'
				outputPos++
				break

			default:
				output[outputPos] = c
				outputPos++
			}
		}
	}
	return output
}
