package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

type VariantFunc func(io.Reader, io.Writer)

func Test_AllVariants(t *testing.T) {
	//variantFuncs := []VariantFunc{
	//	get1,
	//}

	//for range variantFuncs {
	t.Run("example1", func(t *testing.T) {
		in := prepareData(`
				5
				0 1 4 9 0
			`)

		out := bytes.NewBuffer(nil)
		get3(in, out)
		assert.EqualValues(t, "0 1 2 1 0", string(out.Bytes()))
	})

	//t.Run("example2", func(t *testing.T) {
	//	in := prepareData(`
	//		6
	//		0 7 9 4 8 20
	//	`)
	//
	//	assert.EqualValues(t, []string{"0", "1", "2", "3", "4", "5"}, variantFunc(in))
	//})

	//t.Run("example3", func(t *testing.T) {
	//	in := prepareData(`
	//		1000000
	//		0 341835781 698351622 448790534 713031692 220200975 578813969 853540884 671088661 373293080 557842457 92274713 304087068 180355101 704643117 962592813 664797230 627048495 404750395 39845950 25165887 989855807 427819071 551551044 65011781 893386823 543162439 627048519 520093772 320864333 939524172 387973197 312475727 295698513 513802322 859832402 230686802 383778901 700448849 513802321 402653276 874512477 77594720 190840931 987758693 117440613 337641580 436207725 346030198 446693495 90177654 922747001 761266301 509608062 476053633 27263105 46137475 31457414 857735307 912261260 358613134 532676752 482345106 761266324 830472340 266338453 499122328 897581210 387973277 929038495 994050209 429916322 138412195 943718562 654311586 587202726 322961574 643825834 683671722 136315054 771752112 234881201 27263154 792723636 895484087 847249593 341835962 186646716 316670143 176160962 356516035 952107207 784335048 606077129 54526154 815792327 511705292 16777417 782237904 232784085 775946455 874512599 482345178 719323
	//	`)
	//
	//	assert.EqualValues(t, 0, variantFunc(in))
	//})
	//
	//t.Run("example4", func(t *testing.T) {
	//	in := prepareData(`
	//		4
	//		2941
	//		.287
	//		3798
	//		.221
	//	`)
	//
	//	assert.EqualValues(t, 7, variantFunc(in))
	//})
	//}
}

// prepareData убирает пробелы в начале и конце и пустые строки
// помещает все данные в буфер
func prepareData(str string) io.Reader {
	lines := strings.Split(str, "\n")
	filteredLines := lines[:0]
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		filteredLines = append(filteredLines, line)
	}
	return bytes.NewBufferString(strings.Join(filteredLines, "\n"))
}
