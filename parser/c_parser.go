// Code generated from C.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // C

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 91, 286,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 7, 2, 54, 10, 2, 12, 2,
	14, 2, 57, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3,
	67, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 73, 10, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 5, 5, 82, 10, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 96, 10, 7, 3, 7, 3, 7, 7,
	7, 100, 10, 7, 12, 7, 14, 7, 103, 11, 7, 3, 8, 3, 8, 3, 8, 5, 8, 108, 10,
	8, 3, 8, 5, 8, 111, 10, 8, 3, 9, 3, 9, 7, 9, 115, 10, 9, 12, 9, 14, 9,
	118, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 124, 10, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 5, 11, 130, 10, 11, 3, 11, 3, 11, 3, 11, 5, 11, 135, 10,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 5, 12, 142, 10, 12, 3, 12, 3, 12,
	5, 12, 146, 10, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 5, 14, 154,
	10, 14, 3, 14, 3, 14, 5, 14, 158, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 183, 10,
	15, 5, 15, 185, 10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 203,
	10, 15, 3, 15, 3, 15, 3, 15, 7, 15, 208, 10, 15, 12, 15, 14, 15, 211, 11,
	15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 5, 20, 224, 10, 20, 3, 21, 3, 21, 7, 21, 228, 10, 21, 12, 21, 14,
	21, 231, 11, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	5, 22, 241, 10, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5,
	22, 250, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23,
	259, 10, 23, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 265, 10, 24, 3, 24, 3,
	24, 5, 24, 269, 10, 24, 3, 24, 3, 24, 5, 24, 273, 10, 24, 3, 24, 3, 24,
	3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 2,
	4, 12, 28, 27, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 2, 6, 3, 2, 62, 72, 6, 2, 38, 44,
	46, 46, 48, 55, 73, 74, 4, 2, 45, 45, 47, 47, 5, 2, 44, 48, 51, 51, 56,
	57, 2, 316, 2, 55, 3, 2, 2, 2, 4, 66, 3, 2, 2, 2, 6, 68, 3, 2, 2, 2, 8,
	77, 3, 2, 2, 2, 10, 83, 3, 2, 2, 2, 12, 95, 3, 2, 2, 2, 14, 104, 3, 2,
	2, 2, 16, 112, 3, 2, 2, 2, 18, 121, 3, 2, 2, 2, 20, 127, 3, 2, 2, 2, 22,
	138, 3, 2, 2, 2, 24, 147, 3, 2, 2, 2, 26, 150, 3, 2, 2, 2, 28, 184, 3,
	2, 2, 2, 30, 212, 3, 2, 2, 2, 32, 214, 3, 2, 2, 2, 34, 216, 3, 2, 2, 2,
	36, 218, 3, 2, 2, 2, 38, 220, 3, 2, 2, 2, 40, 225, 3, 2, 2, 2, 42, 249,
	3, 2, 2, 2, 44, 251, 3, 2, 2, 2, 46, 260, 3, 2, 2, 2, 48, 277, 3, 2, 2,
	2, 50, 283, 3, 2, 2, 2, 52, 54, 5, 4, 3, 2, 53, 52, 3, 2, 2, 2, 54, 57,
	3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 3, 3, 2, 2, 2,
	57, 55, 3, 2, 2, 2, 58, 67, 5, 6, 4, 2, 59, 60, 5, 14, 8, 2, 60, 61, 7,
	60, 2, 2, 61, 67, 3, 2, 2, 2, 62, 63, 5, 20, 11, 2, 63, 64, 7, 60, 2, 2,
	64, 67, 3, 2, 2, 2, 65, 67, 5, 50, 26, 2, 66, 58, 3, 2, 2, 2, 66, 59, 3,
	2, 2, 2, 66, 62, 3, 2, 2, 2, 66, 65, 3, 2, 2, 2, 67, 5, 3, 2, 2, 2, 68,
	69, 5, 12, 7, 2, 69, 70, 7, 78, 2, 2, 70, 72, 7, 32, 2, 2, 71, 73, 5, 8,
	5, 2, 72, 71, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75,
	7, 33, 2, 2, 75, 76, 5, 40, 21, 2, 76, 7, 3, 2, 2, 2, 77, 78, 5, 12, 7,
	2, 78, 81, 7, 78, 2, 2, 79, 80, 7, 61, 2, 2, 80, 82, 5, 8, 5, 2, 81, 79,
	3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 9, 3, 2, 2, 2, 83, 84, 7, 20, 2, 2,
	84, 85, 5, 28, 15, 2, 85, 11, 3, 2, 2, 2, 86, 87, 8, 7, 1, 2, 87, 96, 7,
	30, 2, 2, 88, 96, 7, 18, 2, 2, 89, 96, 7, 21, 2, 2, 90, 96, 7, 19, 2, 2,
	91, 96, 7, 5, 2, 2, 92, 96, 7, 14, 2, 2, 93, 96, 7, 10, 2, 2, 94, 96, 5,
	14, 8, 2, 95, 86, 3, 2, 2, 2, 95, 88, 3, 2, 2, 2, 95, 89, 3, 2, 2, 2, 95,
	90, 3, 2, 2, 2, 95, 91, 3, 2, 2, 2, 95, 92, 3, 2, 2, 2, 95, 93, 3, 2, 2,
	2, 95, 94, 3, 2, 2, 2, 96, 101, 3, 2, 2, 2, 97, 98, 12, 3, 2, 2, 98, 100,
	7, 48, 2, 2, 99, 97, 3, 2, 2, 2, 100, 103, 3, 2, 2, 2, 101, 99, 3, 2, 2,
	2, 101, 102, 3, 2, 2, 2, 102, 13, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 104,
	110, 7, 25, 2, 2, 105, 107, 7, 78, 2, 2, 106, 108, 5, 16, 9, 2, 107, 106,
	3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 111, 5, 16,
	9, 2, 110, 105, 3, 2, 2, 2, 110, 109, 3, 2, 2, 2, 111, 15, 3, 2, 2, 2,
	112, 116, 7, 36, 2, 2, 113, 115, 5, 18, 10, 2, 114, 113, 3, 2, 2, 2, 115,
	118, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 119,
	3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 119, 120, 7, 37, 2, 2, 120, 17, 3, 2,
	2, 2, 121, 123, 5, 12, 7, 2, 122, 124, 7, 78, 2, 2, 123, 122, 3, 2, 2,
	2, 123, 124, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 126, 7, 60, 2, 2, 126,
	19, 3, 2, 2, 2, 127, 129, 7, 12, 2, 2, 128, 130, 7, 78, 2, 2, 129, 128,
	3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 7, 36,
	2, 2, 132, 134, 5, 22, 12, 2, 133, 135, 7, 61, 2, 2, 134, 133, 3, 2, 2,
	2, 134, 135, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 137, 7, 37, 2, 2, 137,
	21, 3, 2, 2, 2, 138, 141, 7, 78, 2, 2, 139, 140, 7, 62, 2, 2, 140, 142,
	5, 28, 15, 2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 145, 3,
	2, 2, 2, 143, 144, 7, 61, 2, 2, 144, 146, 5, 22, 12, 2, 145, 143, 3, 2,
	2, 2, 145, 146, 3, 2, 2, 2, 146, 23, 3, 2, 2, 2, 147, 148, 5, 12, 7, 2,
	148, 149, 5, 26, 14, 2, 149, 25, 3, 2, 2, 2, 150, 153, 7, 78, 2, 2, 151,
	152, 7, 62, 2, 2, 152, 154, 5, 28, 15, 2, 153, 151, 3, 2, 2, 2, 153, 154,
	3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 156, 7, 61, 2, 2, 156, 158, 5, 26,
	14, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 27, 3, 2, 2, 2,
	159, 160, 8, 15, 1, 2, 160, 185, 7, 78, 2, 2, 161, 185, 7, 79, 2, 2, 162,
	185, 7, 81, 2, 2, 163, 164, 7, 32, 2, 2, 164, 165, 5, 12, 7, 2, 165, 166,
	7, 33, 2, 2, 166, 167, 5, 28, 15, 9, 167, 185, 3, 2, 2, 2, 168, 169, 7,
	32, 2, 2, 169, 170, 5, 28, 15, 2, 170, 171, 7, 33, 2, 2, 171, 185, 3, 2,
	2, 2, 172, 173, 5, 36, 19, 2, 173, 174, 5, 28, 15, 6, 174, 185, 3, 2, 2,
	2, 175, 182, 7, 23, 2, 2, 176, 177, 7, 32, 2, 2, 177, 178, 5, 12, 7, 2,
	178, 179, 7, 33, 2, 2, 179, 183, 3, 2, 2, 2, 180, 183, 5, 12, 7, 2, 181,
	183, 5, 28, 15, 2, 182, 176, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 181,
	3, 2, 2, 2, 183, 185, 3, 2, 2, 2, 184, 159, 3, 2, 2, 2, 184, 161, 3, 2,
	2, 2, 184, 162, 3, 2, 2, 2, 184, 163, 3, 2, 2, 2, 184, 168, 3, 2, 2, 2,
	184, 172, 3, 2, 2, 2, 184, 175, 3, 2, 2, 2, 185, 209, 3, 2, 2, 2, 186,
	187, 12, 5, 2, 2, 187, 188, 5, 30, 16, 2, 188, 189, 5, 28, 15, 6, 189,
	208, 3, 2, 2, 2, 190, 191, 12, 4, 2, 2, 191, 192, 5, 32, 17, 2, 192, 193,
	5, 28, 15, 5, 193, 208, 3, 2, 2, 2, 194, 195, 12, 11, 2, 2, 195, 196, 7,
	34, 2, 2, 196, 197, 5, 28, 15, 2, 197, 198, 7, 35, 2, 2, 198, 208, 3, 2,
	2, 2, 199, 200, 12, 10, 2, 2, 200, 202, 7, 32, 2, 2, 201, 203, 5, 38, 20,
	2, 202, 201, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204,
	208, 7, 33, 2, 2, 205, 206, 12, 7, 2, 2, 206, 208, 5, 34, 18, 2, 207, 186,
	3, 2, 2, 2, 207, 190, 3, 2, 2, 2, 207, 194, 3, 2, 2, 2, 207, 199, 3, 2,
	2, 2, 207, 205, 3, 2, 2, 2, 208, 211, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2,
	209, 210, 3, 2, 2, 2, 210, 29, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 212, 213,
	9, 2, 2, 2, 213, 31, 3, 2, 2, 2, 214, 215, 9, 3, 2, 2, 215, 33, 3, 2, 2,
	2, 216, 217, 9, 4, 2, 2, 217, 35, 3, 2, 2, 2, 218, 219, 9, 5, 2, 2, 219,
	37, 3, 2, 2, 2, 220, 223, 5, 28, 15, 2, 221, 222, 7, 61, 2, 2, 222, 224,
	5, 38, 20, 2, 223, 221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 39, 3, 2,
	2, 2, 225, 229, 7, 36, 2, 2, 226, 228, 5, 42, 22, 2, 227, 226, 3, 2, 2,
	2, 228, 231, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230,
	232, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 232, 233, 7, 37, 2, 2, 233, 41,
	3, 2, 2, 2, 234, 241, 5, 24, 13, 2, 235, 241, 5, 28, 15, 2, 236, 241, 5,
	10, 6, 2, 237, 241, 7, 3, 2, 2, 238, 241, 5, 14, 8, 2, 239, 241, 5, 20,
	11, 2, 240, 234, 3, 2, 2, 2, 240, 235, 3, 2, 2, 2, 240, 236, 3, 2, 2, 2,
	240, 237, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 240, 239, 3, 2, 2, 2, 241,
	242, 3, 2, 2, 2, 242, 250, 7, 60, 2, 2, 243, 250, 5, 44, 23, 2, 244, 250,
	5, 46, 24, 2, 245, 250, 5, 48, 25, 2, 246, 250, 5, 40, 21, 2, 247, 250,
	7, 90, 2, 2, 248, 250, 7, 91, 2, 2, 249, 240, 3, 2, 2, 2, 249, 243, 3,
	2, 2, 2, 249, 244, 3, 2, 2, 2, 249, 245, 3, 2, 2, 2, 249, 246, 3, 2, 2,
	2, 249, 247, 3, 2, 2, 2, 249, 248, 3, 2, 2, 2, 250, 43, 3, 2, 2, 2, 251,
	252, 7, 17, 2, 2, 252, 253, 7, 32, 2, 2, 253, 254, 5, 28, 15, 2, 254, 255,
	7, 33, 2, 2, 255, 258, 5, 42, 22, 2, 256, 257, 7, 11, 2, 2, 257, 259, 5,
	42, 22, 2, 258, 256, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 45, 3, 2, 2,
	2, 260, 261, 7, 15, 2, 2, 261, 264, 7, 32, 2, 2, 262, 265, 5, 28, 15, 2,
	263, 265, 5, 24, 13, 2, 264, 262, 3, 2, 2, 2, 264, 263, 3, 2, 2, 2, 264,
	265, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 268, 7, 60, 2, 2, 267, 269,
	5, 28, 15, 2, 268, 267, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 270, 3,
	2, 2, 2, 270, 272, 7, 60, 2, 2, 271, 273, 5, 28, 15, 2, 272, 271, 3, 2,
	2, 2, 272, 273, 3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 275, 7, 33, 2, 2,
	275, 276, 5, 42, 22, 2, 276, 47, 3, 2, 2, 2, 277, 278, 7, 31, 2, 2, 278,
	279, 7, 32, 2, 2, 279, 280, 5, 28, 15, 2, 280, 281, 7, 33, 2, 2, 281, 282,
	5, 42, 22, 2, 282, 49, 3, 2, 2, 2, 283, 284, 7, 83, 2, 2, 284, 51, 3, 2,
	2, 2, 31, 55, 66, 72, 81, 95, 101, 107, 110, 116, 123, 129, 134, 141, 145,
	153, 157, 182, 184, 202, 207, 209, 223, 229, 240, 249, 258, 264, 268, 272,
}
var literalNames = []string{
	"", "'break'", "'case'", "'char'", "'const'", "'continue'", "'default'",
	"'do'", "'double'", "'else'", "'enum'", "'extern'", "'float'", "'for'",
	"'goto'", "'if'", "'int'", "'long'", "'return'", "'short'", "'signed'",
	"'sizeof'", "'static'", "'struct'", "'switch'", "'typedef'", "'union'",
	"'unsigned'", "'void'", "'while'", "'('", "')'", "'['", "']'", "'{'", "'}'",
	"'<'", "'<='", "'>'", "'>='", "'<<'", "'>>'", "'+'", "'++'", "'-'", "'--'",
	"'*'", "'/'", "'%'", "'&'", "'|'", "'&&'", "'||'", "'^'", "'!'", "'~'",
	"'?'", "':'", "';'", "','", "'='", "'*='", "'/='", "'%='", "'+='", "'-='",
	"'<<='", "'>>='", "'&='", "'^='", "'|='", "'=='", "'!='", "'->'", "'.'",
	"'...'",
}
var symbolicNames = []string{
	"", "Break", "Case", "Char", "Const", "Continue", "Default", "Do", "Double",
	"Else", "Enum", "Extern", "Float", "For", "Goto", "If", "Int", "Long",
	"Return", "Short", "Signed", "Sizeof", "Static", "Struct", "Switch", "Typedef",
	"Union", "Unsigned", "Void", "While", "LeftParen", "RightParen", "LeftBracket",
	"RightBracket", "LeftBrace", "RightBrace", "Less", "LessEqual", "Greater",
	"GreaterEqual", "LeftShift", "RightShift", "Plus", "PlusPlus", "Minus",
	"MinusMinus", "Star", "Div", "Mod", "And", "Or", "AndAnd", "OrOr", "Caret",
	"Not", "Tilde", "Question", "Colon", "Semi", "Comma", "Assign", "StarAssign",
	"DivAssign", "ModAssign", "PlusAssign", "MinusAssign", "LeftShiftAssign",
	"RightShiftAssign", "AndAssign", "XorAssign", "OrAssign", "Equal", "NotEqual",
	"Arrow", "Dot", "Ellipsis", "Identifier", "Constant", "DigitSequence",
	"StringLiteral", "ComplexDefine", "IncludeDirective", "AsmBlock", "LineAfterPreprocessing",
	"LineDirective", "PragmaDirective", "Whitespace", "Newline", "BlockComment",
	"LineComment",
}

var ruleNames = []string{
	"translation", "declaration", "functionDeclaration", "functionArguments",
	"functionReturn", "typeSpecifier", "structDeclaration", "structDeclarationBody",
	"structProperty", "enumDeclaration", "enumProperties", "variableDeclaration",
	"variableDeclarationList", "expression", "assignementOperator", "binaryOperator",
	"unaryOperatorPost", "unaryOperatorPre", "functionCallArguments", "block",
	"statement", "ifStatement", "forStatement", "whileStatement", "includePreprocessor",
}

type CParser struct {
	*antlr.BaseParser
}

// NewCParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *CParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCParser(input antlr.TokenStream) *CParser {
	this := new(CParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "C.g4"

	return this
}

// CParser tokens.
const (
	CParserEOF                    = antlr.TokenEOF
	CParserBreak                  = 1
	CParserCase                   = 2
	CParserChar                   = 3
	CParserConst                  = 4
	CParserContinue               = 5
	CParserDefault                = 6
	CParserDo                     = 7
	CParserDouble                 = 8
	CParserElse                   = 9
	CParserEnum                   = 10
	CParserExtern                 = 11
	CParserFloat                  = 12
	CParserFor                    = 13
	CParserGoto                   = 14
	CParserIf                     = 15
	CParserInt                    = 16
	CParserLong                   = 17
	CParserReturn                 = 18
	CParserShort                  = 19
	CParserSigned                 = 20
	CParserSizeof                 = 21
	CParserStatic                 = 22
	CParserStruct                 = 23
	CParserSwitch                 = 24
	CParserTypedef                = 25
	CParserUnion                  = 26
	CParserUnsigned               = 27
	CParserVoid                   = 28
	CParserWhile                  = 29
	CParserLeftParen              = 30
	CParserRightParen             = 31
	CParserLeftBracket            = 32
	CParserRightBracket           = 33
	CParserLeftBrace              = 34
	CParserRightBrace             = 35
	CParserLess                   = 36
	CParserLessEqual              = 37
	CParserGreater                = 38
	CParserGreaterEqual           = 39
	CParserLeftShift              = 40
	CParserRightShift             = 41
	CParserPlus                   = 42
	CParserPlusPlus               = 43
	CParserMinus                  = 44
	CParserMinusMinus             = 45
	CParserStar                   = 46
	CParserDiv                    = 47
	CParserMod                    = 48
	CParserAnd                    = 49
	CParserOr                     = 50
	CParserAndAnd                 = 51
	CParserOrOr                   = 52
	CParserCaret                  = 53
	CParserNot                    = 54
	CParserTilde                  = 55
	CParserQuestion               = 56
	CParserColon                  = 57
	CParserSemi                   = 58
	CParserComma                  = 59
	CParserAssign                 = 60
	CParserStarAssign             = 61
	CParserDivAssign              = 62
	CParserModAssign              = 63
	CParserPlusAssign             = 64
	CParserMinusAssign            = 65
	CParserLeftShiftAssign        = 66
	CParserRightShiftAssign       = 67
	CParserAndAssign              = 68
	CParserXorAssign              = 69
	CParserOrAssign               = 70
	CParserEqual                  = 71
	CParserNotEqual               = 72
	CParserArrow                  = 73
	CParserDot                    = 74
	CParserEllipsis               = 75
	CParserIdentifier             = 76
	CParserConstant               = 77
	CParserDigitSequence          = 78
	CParserStringLiteral          = 79
	CParserComplexDefine          = 80
	CParserIncludeDirective       = 81
	CParserAsmBlock               = 82
	CParserLineAfterPreprocessing = 83
	CParserLineDirective          = 84
	CParserPragmaDirective        = 85
	CParserWhitespace             = 86
	CParserNewline                = 87
	CParserBlockComment           = 88
	CParserLineComment            = 89
)

// CParser rules.
const (
	CParserRULE_translation             = 0
	CParserRULE_declaration             = 1
	CParserRULE_functionDeclaration     = 2
	CParserRULE_functionArguments       = 3
	CParserRULE_functionReturn          = 4
	CParserRULE_typeSpecifier           = 5
	CParserRULE_structDeclaration       = 6
	CParserRULE_structDeclarationBody   = 7
	CParserRULE_structProperty          = 8
	CParserRULE_enumDeclaration         = 9
	CParserRULE_enumProperties          = 10
	CParserRULE_variableDeclaration     = 11
	CParserRULE_variableDeclarationList = 12
	CParserRULE_expression              = 13
	CParserRULE_assignementOperator     = 14
	CParserRULE_binaryOperator          = 15
	CParserRULE_unaryOperatorPost       = 16
	CParserRULE_unaryOperatorPre        = 17
	CParserRULE_functionCallArguments   = 18
	CParserRULE_block                   = 19
	CParserRULE_statement               = 20
	CParserRULE_ifStatement             = 21
	CParserRULE_forStatement            = 22
	CParserRULE_whileStatement          = 23
	CParserRULE_includePreprocessor     = 24
)

// ITranslationContext is an interface to support dynamic dispatch.
type ITranslationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTranslationContext differentiates from other interfaces.
	IsTranslationContext()
}

type TranslationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTranslationContext() *TranslationContext {
	var p = new(TranslationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_translation
	return p
}

func (*TranslationContext) IsTranslationContext() {}

func NewTranslationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TranslationContext {
	var p = new(TranslationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_translation

	return p
}

func (s *TranslationContext) GetParser() antlr.Parser { return s.parser }

func (s *TranslationContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *TranslationContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *TranslationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TranslationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TranslationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitTranslation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Translation() (localctx ITranslationContext) {
	localctx = NewTranslationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CParserRULE_translation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserEnum)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserStruct)|(1<<CParserVoid))) != 0) || _la == CParserIncludeDirective {
		{
			p.SetState(50)
			p.Declaration()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *DeclarationContext) StructDeclaration() IStructDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructDeclarationContext)
}

func (s *DeclarationContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *DeclarationContext) EnumDeclaration() IEnumDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDeclarationContext)
}

func (s *DeclarationContext) IncludePreprocessor() IIncludePreprocessorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncludePreprocessorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncludePreprocessorContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.FunctionDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.StructDeclaration()
		}
		{
			p.SetState(58)
			p.Match(CParserSemi)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
			p.EnumDeclaration()
		}
		{
			p.SetState(61)
			p.Match(CParserSemi)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(63)
			p.IncludePreprocessor()
		}

	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *FunctionDeclarationContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *FunctionDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclarationContext) FunctionArguments() IFunctionArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgumentsContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CParserRULE_functionDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.typeSpecifier(0)
	}
	{
		p.SetState(67)
		p.Match(CParserIdentifier)
	}
	{
		p.SetState(68)
		p.Match(CParserLeftParen)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserStruct)|(1<<CParserVoid))) != 0 {
		{
			p.SetState(69)
			p.FunctionArguments()
		}

	}
	{
		p.SetState(72)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(73)
		p.Block()
	}

	return localctx
}

// IFunctionArgumentsContext is an interface to support dynamic dispatch.
type IFunctionArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgumentsContext differentiates from other interfaces.
	IsFunctionArgumentsContext()
}

type FunctionArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgumentsContext() *FunctionArgumentsContext {
	var p = new(FunctionArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_functionArguments
	return p
}

func (*FunctionArgumentsContext) IsFunctionArgumentsContext() {}

func NewFunctionArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgumentsContext {
	var p = new(FunctionArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_functionArguments

	return p
}

func (s *FunctionArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgumentsContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *FunctionArgumentsContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *FunctionArgumentsContext) Comma() antlr.TerminalNode {
	return s.GetToken(CParserComma, 0)
}

func (s *FunctionArgumentsContext) FunctionArguments() IFunctionArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgumentsContext)
}

func (s *FunctionArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitFunctionArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) FunctionArguments() (localctx IFunctionArgumentsContext) {
	localctx = NewFunctionArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CParserRULE_functionArguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.typeSpecifier(0)
	}
	{
		p.SetState(76)
		p.Match(CParserIdentifier)
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(77)
			p.Match(CParserComma)
		}
		{
			p.SetState(78)
			p.FunctionArguments()
		}

	}

	return localctx
}

// IFunctionReturnContext is an interface to support dynamic dispatch.
type IFunctionReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionReturnContext differentiates from other interfaces.
	IsFunctionReturnContext()
}

type FunctionReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionReturnContext() *FunctionReturnContext {
	var p = new(FunctionReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_functionReturn
	return p
}

func (*FunctionReturnContext) IsFunctionReturnContext() {}

func NewFunctionReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionReturnContext {
	var p = new(FunctionReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_functionReturn

	return p
}

func (s *FunctionReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionReturnContext) Return() antlr.TerminalNode {
	return s.GetToken(CParserReturn, 0)
}

func (s *FunctionReturnContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitFunctionReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) FunctionReturn() (localctx IFunctionReturnContext) {
	localctx = NewFunctionReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CParserRULE_functionReturn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(CParserReturn)
	}
	{
		p.SetState(82)
		p.expression(0)
	}

	return localctx
}

// ITypeSpecifierContext is an interface to support dynamic dispatch.
type ITypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecifierContext differentiates from other interfaces.
	IsTypeSpecifierContext()
}

type TypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierContext() *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_typeSpecifier
	return p
}

func (*TypeSpecifierContext) IsTypeSpecifierContext() {}

func NewTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_typeSpecifier

	return p
}

func (s *TypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierContext) Void() antlr.TerminalNode {
	return s.GetToken(CParserVoid, 0)
}

func (s *TypeSpecifierContext) Int() antlr.TerminalNode {
	return s.GetToken(CParserInt, 0)
}

func (s *TypeSpecifierContext) Short() antlr.TerminalNode {
	return s.GetToken(CParserShort, 0)
}

func (s *TypeSpecifierContext) Long() antlr.TerminalNode {
	return s.GetToken(CParserLong, 0)
}

func (s *TypeSpecifierContext) Char() antlr.TerminalNode {
	return s.GetToken(CParserChar, 0)
}

func (s *TypeSpecifierContext) Float() antlr.TerminalNode {
	return s.GetToken(CParserFloat, 0)
}

func (s *TypeSpecifierContext) Double() antlr.TerminalNode {
	return s.GetToken(CParserDouble, 0)
}

func (s *TypeSpecifierContext) StructDeclaration() IStructDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructDeclarationContext)
}

func (s *TypeSpecifierContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *TypeSpecifierContext) Star() antlr.TerminalNode {
	return s.GetToken(CParserStar, 0)
}

func (s *TypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitTypeSpecifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) TypeSpecifier() (localctx ITypeSpecifierContext) {
	return p.typeSpecifier(0)
}

func (p *CParser) typeSpecifier(_p int) (localctx ITypeSpecifierContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTypeSpecifierContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITypeSpecifierContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, CParserRULE_typeSpecifier, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserVoid:
		{
			p.SetState(85)
			p.Match(CParserVoid)
		}

	case CParserInt:
		{
			p.SetState(86)
			p.Match(CParserInt)
		}

	case CParserShort:
		{
			p.SetState(87)
			p.Match(CParserShort)
		}

	case CParserLong:
		{
			p.SetState(88)
			p.Match(CParserLong)
		}

	case CParserChar:
		{
			p.SetState(89)
			p.Match(CParserChar)
		}

	case CParserFloat:
		{
			p.SetState(90)
			p.Match(CParserFloat)
		}

	case CParserDouble:
		{
			p.SetState(91)
			p.Match(CParserDouble)
		}

	case CParserStruct:
		{
			p.SetState(92)
			p.StructDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTypeSpecifierContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CParserRULE_typeSpecifier)
			p.SetState(95)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(96)
				p.Match(CParserStar)
			}

		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IStructDeclarationContext is an interface to support dynamic dispatch.
type IStructDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructDeclarationContext differentiates from other interfaces.
	IsStructDeclarationContext()
}

type StructDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclarationContext() *StructDeclarationContext {
	var p = new(StructDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_structDeclaration
	return p
}

func (*StructDeclarationContext) IsStructDeclarationContext() {}

func NewStructDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclarationContext {
	var p = new(StructDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_structDeclaration

	return p
}

func (s *StructDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclarationContext) Struct() antlr.TerminalNode {
	return s.GetToken(CParserStruct, 0)
}

func (s *StructDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *StructDeclarationContext) StructDeclarationBody() IStructDeclarationBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructDeclarationBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructDeclarationBodyContext)
}

func (s *StructDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitStructDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) StructDeclaration() (localctx IStructDeclarationContext) {
	localctx = NewStructDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CParserRULE_structDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(CParserStruct)
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserIdentifier:
		{
			p.SetState(103)
			p.Match(CParserIdentifier)
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(104)
				p.StructDeclarationBody()
			}

		}

	case CParserLeftBrace:
		{
			p.SetState(107)
			p.StructDeclarationBody()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStructDeclarationBodyContext is an interface to support dynamic dispatch.
type IStructDeclarationBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructDeclarationBodyContext differentiates from other interfaces.
	IsStructDeclarationBodyContext()
}

type StructDeclarationBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclarationBodyContext() *StructDeclarationBodyContext {
	var p = new(StructDeclarationBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_structDeclarationBody
	return p
}

func (*StructDeclarationBodyContext) IsStructDeclarationBodyContext() {}

func NewStructDeclarationBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclarationBodyContext {
	var p = new(StructDeclarationBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_structDeclarationBody

	return p
}

func (s *StructDeclarationBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclarationBodyContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(CParserLeftBrace, 0)
}

func (s *StructDeclarationBodyContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(CParserRightBrace, 0)
}

func (s *StructDeclarationBodyContext) AllStructProperty() []IStructPropertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructPropertyContext)(nil)).Elem())
	var tst = make([]IStructPropertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructPropertyContext)
		}
	}

	return tst
}

func (s *StructDeclarationBodyContext) StructProperty(i int) IStructPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructPropertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructPropertyContext)
}

func (s *StructDeclarationBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclarationBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclarationBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitStructDeclarationBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) StructDeclarationBody() (localctx IStructDeclarationBodyContext) {
	localctx = NewStructDeclarationBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CParserRULE_structDeclarationBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(CParserLeftBrace)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserStruct)|(1<<CParserVoid))) != 0 {
		{
			p.SetState(111)
			p.StructProperty()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(CParserRightBrace)
	}

	return localctx
}

// IStructPropertyContext is an interface to support dynamic dispatch.
type IStructPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructPropertyContext differentiates from other interfaces.
	IsStructPropertyContext()
}

type StructPropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructPropertyContext() *StructPropertyContext {
	var p = new(StructPropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_structProperty
	return p
}

func (*StructPropertyContext) IsStructPropertyContext() {}

func NewStructPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructPropertyContext {
	var p = new(StructPropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_structProperty

	return p
}

func (s *StructPropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *StructPropertyContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *StructPropertyContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *StructPropertyContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *StructPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructPropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructPropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitStructProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) StructProperty() (localctx IStructPropertyContext) {
	localctx = NewStructPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CParserRULE_structProperty)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.typeSpecifier(0)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserIdentifier {
		{
			p.SetState(120)
			p.Match(CParserIdentifier)
		}

	}
	{
		p.SetState(123)
		p.Match(CParserSemi)
	}

	return localctx
}

// IEnumDeclarationContext is an interface to support dynamic dispatch.
type IEnumDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumDeclarationContext differentiates from other interfaces.
	IsEnumDeclarationContext()
}

type EnumDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDeclarationContext() *EnumDeclarationContext {
	var p = new(EnumDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_enumDeclaration
	return p
}

func (*EnumDeclarationContext) IsEnumDeclarationContext() {}

func NewEnumDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDeclarationContext {
	var p = new(EnumDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_enumDeclaration

	return p
}

func (s *EnumDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDeclarationContext) Enum() antlr.TerminalNode {
	return s.GetToken(CParserEnum, 0)
}

func (s *EnumDeclarationContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(CParserLeftBrace, 0)
}

func (s *EnumDeclarationContext) EnumProperties() IEnumPropertiesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumPropertiesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumPropertiesContext)
}

func (s *EnumDeclarationContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(CParserRightBrace, 0)
}

func (s *EnumDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *EnumDeclarationContext) Comma() antlr.TerminalNode {
	return s.GetToken(CParserComma, 0)
}

func (s *EnumDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitEnumDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) EnumDeclaration() (localctx IEnumDeclarationContext) {
	localctx = NewEnumDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CParserRULE_enumDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(CParserEnum)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserIdentifier {
		{
			p.SetState(126)
			p.Match(CParserIdentifier)
		}

	}
	{
		p.SetState(129)
		p.Match(CParserLeftBrace)
	}
	{
		p.SetState(130)
		p.EnumProperties()
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(131)
			p.Match(CParserComma)
		}

	}
	{
		p.SetState(134)
		p.Match(CParserRightBrace)
	}

	return localctx
}

// IEnumPropertiesContext is an interface to support dynamic dispatch.
type IEnumPropertiesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumPropertiesContext differentiates from other interfaces.
	IsEnumPropertiesContext()
}

type EnumPropertiesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumPropertiesContext() *EnumPropertiesContext {
	var p = new(EnumPropertiesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_enumProperties
	return p
}

func (*EnumPropertiesContext) IsEnumPropertiesContext() {}

func NewEnumPropertiesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumPropertiesContext {
	var p = new(EnumPropertiesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_enumProperties

	return p
}

func (s *EnumPropertiesContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumPropertiesContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *EnumPropertiesContext) Assign() antlr.TerminalNode {
	return s.GetToken(CParserAssign, 0)
}

func (s *EnumPropertiesContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EnumPropertiesContext) Comma() antlr.TerminalNode {
	return s.GetToken(CParserComma, 0)
}

func (s *EnumPropertiesContext) EnumProperties() IEnumPropertiesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumPropertiesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumPropertiesContext)
}

func (s *EnumPropertiesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumPropertiesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumPropertiesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitEnumProperties(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) EnumProperties() (localctx IEnumPropertiesContext) {
	localctx = NewEnumPropertiesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CParserRULE_enumProperties)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(CParserIdentifier)
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserAssign {
		{
			p.SetState(137)
			p.Match(CParserAssign)
		}
		{
			p.SetState(138)
			p.expression(0)
		}

	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(141)
			p.Match(CParserComma)
		}
		{
			p.SetState(142)
			p.EnumProperties()
		}

	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *VariableDeclarationContext) VariableDeclarationList() IVariableDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationListContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CParserRULE_variableDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.typeSpecifier(0)
	}
	{
		p.SetState(146)
		p.VariableDeclarationList()
	}

	return localctx
}

// IVariableDeclarationListContext is an interface to support dynamic dispatch.
type IVariableDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationListContext differentiates from other interfaces.
	IsVariableDeclarationListContext()
}

type VariableDeclarationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationListContext() *VariableDeclarationListContext {
	var p = new(VariableDeclarationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_variableDeclarationList
	return p
}

func (*VariableDeclarationListContext) IsVariableDeclarationListContext() {}

func NewVariableDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationListContext {
	var p = new(VariableDeclarationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_variableDeclarationList

	return p
}

func (s *VariableDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationListContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *VariableDeclarationListContext) Assign() antlr.TerminalNode {
	return s.GetToken(CParserAssign, 0)
}

func (s *VariableDeclarationListContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationListContext) Comma() antlr.TerminalNode {
	return s.GetToken(CParserComma, 0)
}

func (s *VariableDeclarationListContext) VariableDeclarationList() IVariableDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationListContext)
}

func (s *VariableDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitVariableDeclarationList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) VariableDeclarationList() (localctx IVariableDeclarationListContext) {
	localctx = NewVariableDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CParserRULE_variableDeclarationList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(CParserIdentifier)
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserAssign {
		{
			p.SetState(149)
			p.Match(CParserAssign)
		}
		{
			p.SetState(150)
			p.expression(0)
		}

	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(153)
			p.Match(CParserComma)
		}
		{
			p.SetState(154)
			p.VariableDeclarationList()
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParenthesizedExpressionContext struct {
	*ExpressionContext
}

func NewParenthesizedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesizedExpressionContext {
	var p = new(ParenthesizedExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesizedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *ParenthesizedExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesizedExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *ParenthesizedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitParenthesizedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type SizeofExpressionContext struct {
	*ExpressionContext
}

func NewSizeofExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SizeofExpressionContext {
	var p = new(SizeofExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SizeofExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SizeofExpressionContext) Sizeof() antlr.TerminalNode {
	return s.GetToken(CParserSizeof, 0)
}

func (s *SizeofExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *SizeofExpressionContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *SizeofExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *SizeofExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SizeofExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitSizeofExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayIndexExpressionContext struct {
	*ExpressionContext
}

func NewArrayIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayIndexExpressionContext {
	var p = new(ArrayIndexExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ArrayIndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayIndexExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArrayIndexExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayIndexExpressionContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(CParserLeftBracket, 0)
}

func (s *ArrayIndexExpressionContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(CParserRightBracket, 0)
}

func (s *ArrayIndexExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitArrayIndexExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExpressionPostContext struct {
	*ExpressionContext
}

func NewUnaryExpressionPostContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionPostContext {
	var p = new(UnaryExpressionPostContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExpressionPostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionPostContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionPostContext) UnaryOperatorPost() IUnaryOperatorPostContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOperatorPostContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorPostContext)
}

func (s *UnaryExpressionPostContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitUnaryExpressionPost(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentExpressionContext struct {
	*ExpressionContext
}

func NewAssignmentExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AssignmentExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentExpressionContext) AssignementOperator() IAssignementOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignementOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignementOperatorContext)
}

func (s *AssignmentExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAssignmentExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryExpressionContext struct {
	*ExpressionContext
}

func NewBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BinaryExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryExpressionContext) BinaryOperator() IBinaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryOperatorContext)
}

func (s *BinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitBinaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstantExpressionContext struct {
	*ExpressionContext
}

func NewConstantExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantExpressionContext {
	var p = new(ConstantExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ConstantExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantExpressionContext) Constant() antlr.TerminalNode {
	return s.GetToken(CParserConstant, 0)
}

func (s *ConstantExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitConstantExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExpressionPreContext struct {
	*ExpressionContext
}

func NewUnaryExpressionPreContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionPreContext {
	var p = new(UnaryExpressionPreContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExpressionPreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionPreContext) UnaryOperatorPre() IUnaryOperatorPreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOperatorPreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorPreContext)
}

func (s *UnaryExpressionPreContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionPreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitUnaryExpressionPre(s)

	default:
		return t.VisitChildren(s)
	}
}

type CastExpressionContext struct {
	*ExpressionContext
}

func NewCastExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CastExpressionContext {
	var p = new(CastExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CastExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *CastExpressionContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *CastExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *CastExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CastExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitCastExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstantStringExpressionContext struct {
	*ExpressionContext
}

func NewConstantStringExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantStringExpressionContext {
	var p = new(ConstantStringExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ConstantStringExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantStringExpressionContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(CParserStringLiteral, 0)
}

func (s *ConstantStringExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitConstantStringExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallExpressionContext struct {
	*ExpressionContext
}

func NewFunctionCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpressionContext {
	var p = new(FunctionCallExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionCallExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *FunctionCallExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *FunctionCallExpressionContext) FunctionCallArguments() IFunctionCallArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallArgumentsContext)
}

func (s *FunctionCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitFunctionCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	*ExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *CParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, CParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(158)
			p.Match(CParserIdentifier)
		}

	case 2:
		localctx = NewConstantExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(159)
			p.Match(CParserConstant)
		}

	case 3:
		localctx = NewConstantStringExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(160)
			p.Match(CParserStringLiteral)
		}

	case 4:
		localctx = NewCastExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(161)
			p.Match(CParserLeftParen)
		}
		{
			p.SetState(162)
			p.typeSpecifier(0)
		}
		{
			p.SetState(163)
			p.Match(CParserRightParen)
		}
		{
			p.SetState(164)
			p.expression(7)
		}

	case 5:
		localctx = NewParenthesizedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(166)
			p.Match(CParserLeftParen)
		}
		{
			p.SetState(167)
			p.expression(0)
		}
		{
			p.SetState(168)
			p.Match(CParserRightParen)
		}

	case 6:
		localctx = NewUnaryExpressionPreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(170)
			p.UnaryOperatorPre()
		}
		{
			p.SetState(171)
			p.expression(4)
		}

	case 7:
		localctx = NewSizeofExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(173)
			p.Match(CParserSizeof)
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(174)
				p.Match(CParserLeftParen)
			}
			{
				p.SetState(175)
				p.typeSpecifier(0)
			}
			{
				p.SetState(176)
				p.Match(CParserRightParen)
			}

		case 2:
			{
				p.SetState(178)
				p.typeSpecifier(0)
			}

		case 3:
			{
				p.SetState(179)
				p.expression(0)
			}

		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAssignmentExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(184)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(185)
					p.AssignementOperator()
				}
				{
					p.SetState(186)
					p.expression(4)
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(188)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(189)
					p.BinaryOperator()
				}
				{
					p.SetState(190)
					p.expression(3)
				}

			case 3:
				localctx = NewArrayIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(192)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(193)
					p.Match(CParserLeftBracket)
				}
				{
					p.SetState(194)
					p.expression(0)
				}
				{
					p.SetState(195)
					p.Match(CParserRightBracket)
				}

			case 4:
				localctx = NewFunctionCallExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(197)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(198)
					p.Match(CParserLeftParen)
				}
				p.SetState(200)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CParserSizeof-21))|(1<<(CParserLeftParen-21))|(1<<(CParserPlus-21))|(1<<(CParserPlusPlus-21))|(1<<(CParserMinus-21))|(1<<(CParserMinusMinus-21))|(1<<(CParserStar-21))|(1<<(CParserAnd-21)))) != 0) || (((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(CParserNot-54))|(1<<(CParserTilde-54))|(1<<(CParserIdentifier-54))|(1<<(CParserConstant-54))|(1<<(CParserStringLiteral-54)))) != 0) {
					{
						p.SetState(199)
						p.FunctionCallArguments()
					}

				}
				{
					p.SetState(202)
					p.Match(CParserRightParen)
				}

			case 5:
				localctx = NewUnaryExpressionPostContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(203)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(204)
					p.UnaryOperatorPost()
				}

			}

		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IAssignementOperatorContext is an interface to support dynamic dispatch.
type IAssignementOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignementOperatorContext differentiates from other interfaces.
	IsAssignementOperatorContext()
}

type AssignementOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignementOperatorContext() *AssignementOperatorContext {
	var p = new(AssignementOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_assignementOperator
	return p
}

func (*AssignementOperatorContext) IsAssignementOperatorContext() {}

func NewAssignementOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignementOperatorContext {
	var p = new(AssignementOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_assignementOperator

	return p
}

func (s *AssignementOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignementOperatorContext) Assign() antlr.TerminalNode {
	return s.GetToken(CParserAssign, 0)
}

func (s *AssignementOperatorContext) StarAssign() antlr.TerminalNode {
	return s.GetToken(CParserStarAssign, 0)
}

func (s *AssignementOperatorContext) DivAssign() antlr.TerminalNode {
	return s.GetToken(CParserDivAssign, 0)
}

func (s *AssignementOperatorContext) ModAssign() antlr.TerminalNode {
	return s.GetToken(CParserModAssign, 0)
}

func (s *AssignementOperatorContext) PlusAssign() antlr.TerminalNode {
	return s.GetToken(CParserPlusAssign, 0)
}

func (s *AssignementOperatorContext) MinusAssign() antlr.TerminalNode {
	return s.GetToken(CParserMinusAssign, 0)
}

func (s *AssignementOperatorContext) LeftShiftAssign() antlr.TerminalNode {
	return s.GetToken(CParserLeftShiftAssign, 0)
}

func (s *AssignementOperatorContext) RightShiftAssign() antlr.TerminalNode {
	return s.GetToken(CParserRightShiftAssign, 0)
}

func (s *AssignementOperatorContext) AndAssign() antlr.TerminalNode {
	return s.GetToken(CParserAndAssign, 0)
}

func (s *AssignementOperatorContext) XorAssign() antlr.TerminalNode {
	return s.GetToken(CParserXorAssign, 0)
}

func (s *AssignementOperatorContext) OrAssign() antlr.TerminalNode {
	return s.GetToken(CParserOrAssign, 0)
}

func (s *AssignementOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignementOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignementOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitAssignementOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) AssignementOperator() (localctx IAssignementOperatorContext) {
	localctx = NewAssignementOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CParserRULE_assignementOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-60)&-(0x1f+1)) == 0 && ((1<<uint((_la-60)))&((1<<(CParserAssign-60))|(1<<(CParserStarAssign-60))|(1<<(CParserDivAssign-60))|(1<<(CParserModAssign-60))|(1<<(CParserPlusAssign-60))|(1<<(CParserMinusAssign-60))|(1<<(CParserLeftShiftAssign-60))|(1<<(CParserRightShiftAssign-60))|(1<<(CParserAndAssign-60))|(1<<(CParserXorAssign-60))|(1<<(CParserOrAssign-60)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBinaryOperatorContext is an interface to support dynamic dispatch.
type IBinaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryOperatorContext differentiates from other interfaces.
	IsBinaryOperatorContext()
}

type BinaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperatorContext() *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_binaryOperator
	return p
}

func (*BinaryOperatorContext) IsBinaryOperatorContext() {}

func NewBinaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_binaryOperator

	return p
}

func (s *BinaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperatorContext) Star() antlr.TerminalNode {
	return s.GetToken(CParserStar, 0)
}

func (s *BinaryOperatorContext) Div() antlr.TerminalNode {
	return s.GetToken(CParserDiv, 0)
}

func (s *BinaryOperatorContext) Mod() antlr.TerminalNode {
	return s.GetToken(CParserMod, 0)
}

func (s *BinaryOperatorContext) Plus() antlr.TerminalNode {
	return s.GetToken(CParserPlus, 0)
}

func (s *BinaryOperatorContext) Minus() antlr.TerminalNode {
	return s.GetToken(CParserMinus, 0)
}

func (s *BinaryOperatorContext) LeftShift() antlr.TerminalNode {
	return s.GetToken(CParserLeftShift, 0)
}

func (s *BinaryOperatorContext) RightShift() antlr.TerminalNode {
	return s.GetToken(CParserRightShift, 0)
}

func (s *BinaryOperatorContext) And() antlr.TerminalNode {
	return s.GetToken(CParserAnd, 0)
}

func (s *BinaryOperatorContext) Caret() antlr.TerminalNode {
	return s.GetToken(CParserCaret, 0)
}

func (s *BinaryOperatorContext) Or() antlr.TerminalNode {
	return s.GetToken(CParserOr, 0)
}

func (s *BinaryOperatorContext) Less() antlr.TerminalNode {
	return s.GetToken(CParserLess, 0)
}

func (s *BinaryOperatorContext) Greater() antlr.TerminalNode {
	return s.GetToken(CParserGreater, 0)
}

func (s *BinaryOperatorContext) LessEqual() antlr.TerminalNode {
	return s.GetToken(CParserLessEqual, 0)
}

func (s *BinaryOperatorContext) GreaterEqual() antlr.TerminalNode {
	return s.GetToken(CParserGreaterEqual, 0)
}

func (s *BinaryOperatorContext) Equal() antlr.TerminalNode {
	return s.GetToken(CParserEqual, 0)
}

func (s *BinaryOperatorContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(CParserNotEqual, 0)
}

func (s *BinaryOperatorContext) AndAnd() antlr.TerminalNode {
	return s.GetToken(CParserAndAnd, 0)
}

func (s *BinaryOperatorContext) OrOr() antlr.TerminalNode {
	return s.GetToken(CParserOrOr, 0)
}

func (s *BinaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitBinaryOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) BinaryOperator() (localctx IBinaryOperatorContext) {
	localctx = NewBinaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CParserRULE_binaryOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(CParserLess-36))|(1<<(CParserLessEqual-36))|(1<<(CParserGreater-36))|(1<<(CParserGreaterEqual-36))|(1<<(CParserLeftShift-36))|(1<<(CParserRightShift-36))|(1<<(CParserPlus-36))|(1<<(CParserMinus-36))|(1<<(CParserStar-36))|(1<<(CParserDiv-36))|(1<<(CParserMod-36))|(1<<(CParserAnd-36))|(1<<(CParserOr-36))|(1<<(CParserAndAnd-36))|(1<<(CParserOrOr-36))|(1<<(CParserCaret-36)))) != 0) || _la == CParserEqual || _la == CParserNotEqual) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnaryOperatorPostContext is an interface to support dynamic dispatch.
type IUnaryOperatorPostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperatorPostContext differentiates from other interfaces.
	IsUnaryOperatorPostContext()
}

type UnaryOperatorPostContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorPostContext() *UnaryOperatorPostContext {
	var p = new(UnaryOperatorPostContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_unaryOperatorPost
	return p
}

func (*UnaryOperatorPostContext) IsUnaryOperatorPostContext() {}

func NewUnaryOperatorPostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorPostContext {
	var p = new(UnaryOperatorPostContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_unaryOperatorPost

	return p
}

func (s *UnaryOperatorPostContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperatorPostContext) PlusPlus() antlr.TerminalNode {
	return s.GetToken(CParserPlusPlus, 0)
}

func (s *UnaryOperatorPostContext) MinusMinus() antlr.TerminalNode {
	return s.GetToken(CParserMinusMinus, 0)
}

func (s *UnaryOperatorPostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorPostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorPostContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitUnaryOperatorPost(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) UnaryOperatorPost() (localctx IUnaryOperatorPostContext) {
	localctx = NewUnaryOperatorPostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CParserRULE_unaryOperatorPost)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserPlusPlus || _la == CParserMinusMinus) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnaryOperatorPreContext is an interface to support dynamic dispatch.
type IUnaryOperatorPreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperatorPreContext differentiates from other interfaces.
	IsUnaryOperatorPreContext()
}

type UnaryOperatorPreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorPreContext() *UnaryOperatorPreContext {
	var p = new(UnaryOperatorPreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_unaryOperatorPre
	return p
}

func (*UnaryOperatorPreContext) IsUnaryOperatorPreContext() {}

func NewUnaryOperatorPreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorPreContext {
	var p = new(UnaryOperatorPreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_unaryOperatorPre

	return p
}

func (s *UnaryOperatorPreContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperatorPreContext) Plus() antlr.TerminalNode {
	return s.GetToken(CParserPlus, 0)
}

func (s *UnaryOperatorPreContext) PlusPlus() antlr.TerminalNode {
	return s.GetToken(CParserPlusPlus, 0)
}

func (s *UnaryOperatorPreContext) Minus() antlr.TerminalNode {
	return s.GetToken(CParserMinus, 0)
}

func (s *UnaryOperatorPreContext) MinusMinus() antlr.TerminalNode {
	return s.GetToken(CParserMinusMinus, 0)
}

func (s *UnaryOperatorPreContext) Tilde() antlr.TerminalNode {
	return s.GetToken(CParserTilde, 0)
}

func (s *UnaryOperatorPreContext) Not() antlr.TerminalNode {
	return s.GetToken(CParserNot, 0)
}

func (s *UnaryOperatorPreContext) And() antlr.TerminalNode {
	return s.GetToken(CParserAnd, 0)
}

func (s *UnaryOperatorPreContext) Star() antlr.TerminalNode {
	return s.GetToken(CParserStar, 0)
}

func (s *UnaryOperatorPreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorPreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorPreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitUnaryOperatorPre(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) UnaryOperatorPre() (localctx IUnaryOperatorPreContext) {
	localctx = NewUnaryOperatorPreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CParserRULE_unaryOperatorPre)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(CParserPlus-42))|(1<<(CParserPlusPlus-42))|(1<<(CParserMinus-42))|(1<<(CParserMinusMinus-42))|(1<<(CParserStar-42))|(1<<(CParserAnd-42))|(1<<(CParserNot-42))|(1<<(CParserTilde-42)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunctionCallArgumentsContext is an interface to support dynamic dispatch.
type IFunctionCallArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallArgumentsContext differentiates from other interfaces.
	IsFunctionCallArgumentsContext()
}

type FunctionCallArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallArgumentsContext() *FunctionCallArgumentsContext {
	var p = new(FunctionCallArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_functionCallArguments
	return p
}

func (*FunctionCallArgumentsContext) IsFunctionCallArgumentsContext() {}

func NewFunctionCallArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallArgumentsContext {
	var p = new(FunctionCallArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_functionCallArguments

	return p
}

func (s *FunctionCallArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallArgumentsContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionCallArgumentsContext) Comma() antlr.TerminalNode {
	return s.GetToken(CParserComma, 0)
}

func (s *FunctionCallArgumentsContext) FunctionCallArguments() IFunctionCallArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallArgumentsContext)
}

func (s *FunctionCallArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitFunctionCallArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) FunctionCallArguments() (localctx IFunctionCallArgumentsContext) {
	localctx = NewFunctionCallArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CParserRULE_functionCallArguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.expression(0)
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(219)
			p.Match(CParserComma)
		}
		{
			p.SetState(220)
			p.FunctionCallArguments()
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(CParserLeftBrace, 0)
}

func (s *BlockContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(CParserRightBrace, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(CParserLeftBrace)
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserBreak)|(1<<CParserChar)|(1<<CParserDouble)|(1<<CParserEnum)|(1<<CParserFloat)|(1<<CParserFor)|(1<<CParserIf)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserReturn)|(1<<CParserShort)|(1<<CParserSizeof)|(1<<CParserStruct)|(1<<CParserVoid)|(1<<CParserWhile)|(1<<CParserLeftParen))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(CParserLeftBrace-34))|(1<<(CParserPlus-34))|(1<<(CParserPlusPlus-34))|(1<<(CParserMinus-34))|(1<<(CParserMinusMinus-34))|(1<<(CParserStar-34))|(1<<(CParserAnd-34))|(1<<(CParserNot-34))|(1<<(CParserTilde-34)))) != 0) || (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(CParserIdentifier-76))|(1<<(CParserConstant-76))|(1<<(CParserStringLiteral-76))|(1<<(CParserBlockComment-76))|(1<<(CParserLineComment-76)))) != 0) {
		{
			p.SetState(224)
			p.Statement()
		}

		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(230)
		p.Match(CParserRightBrace)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Semi() antlr.TerminalNode {
	return s.GetToken(CParserSemi, 0)
}

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) FunctionReturn() IFunctionReturnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionReturnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionReturnContext)
}

func (s *StatementContext) Break() antlr.TerminalNode {
	return s.GetToken(CParserBreak, 0)
}

func (s *StatementContext) StructDeclaration() IStructDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructDeclarationContext)
}

func (s *StatementContext) EnumDeclaration() IEnumDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDeclarationContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) BlockComment() antlr.TerminalNode {
	return s.GetToken(CParserBlockComment, 0)
}

func (s *StatementContext) LineComment() antlr.TerminalNode {
	return s.GetToken(CParserLineComment, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserBreak, CParserChar, CParserDouble, CParserEnum, CParserFloat, CParserInt, CParserLong, CParserReturn, CParserShort, CParserSizeof, CParserStruct, CParserVoid, CParserLeftParen, CParserPlus, CParserPlusPlus, CParserMinus, CParserMinusMinus, CParserStar, CParserAnd, CParserNot, CParserTilde, CParserIdentifier, CParserConstant, CParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(232)
				p.VariableDeclaration()
			}

		case 2:
			{
				p.SetState(233)
				p.expression(0)
			}

		case 3:
			{
				p.SetState(234)
				p.FunctionReturn()
			}

		case 4:
			{
				p.SetState(235)
				p.Match(CParserBreak)
			}

		case 5:
			{
				p.SetState(236)
				p.StructDeclaration()
			}

		case 6:
			{
				p.SetState(237)
				p.EnumDeclaration()
			}

		}
		{
			p.SetState(240)
			p.Match(CParserSemi)
		}

	case CParserIf:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)
			p.IfStatement()
		}

	case CParserFor:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(242)
			p.ForStatement()
		}

	case CParserWhile:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(243)
			p.WhileStatement()
		}

	case CParserLeftBrace:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(244)
			p.Block()
		}

	case CParserBlockComment:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(245)
			p.Match(CParserBlockComment)
		}

	case CParserLineComment:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(246)
			p.Match(CParserLineComment)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) If() antlr.TerminalNode {
	return s.GetToken(CParserIf, 0)
}

func (s *IfStatementContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(CParserElse, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CParserRULE_ifStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(CParserIf)
	}
	{
		p.SetState(250)
		p.Match(CParserLeftParen)
	}
	{
		p.SetState(251)
		p.expression(0)
	}
	{
		p.SetState(252)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(253)
		p.Statement()
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(254)
			p.Match(CParserElse)
		}
		{
			p.SetState(255)
			p.Statement()
		}

	}

	return localctx
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInit returns the init rule contexts.
	GetInit() IExpressionContext

	// GetCondition returns the condition rule contexts.
	GetCondition() IExpressionContext

	// GetPost returns the post rule contexts.
	GetPost() IExpressionContext

	// SetInit sets the init rule contexts.
	SetInit(IExpressionContext)

	// SetCondition sets the condition rule contexts.
	SetCondition(IExpressionContext)

	// SetPost sets the post rule contexts.
	SetPost(IExpressionContext)

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	init      IExpressionContext
	condition IExpressionContext
	post      IExpressionContext
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_forStatement
	return p
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) GetInit() IExpressionContext { return s.init }

func (s *ForStatementContext) GetCondition() IExpressionContext { return s.condition }

func (s *ForStatementContext) GetPost() IExpressionContext { return s.post }

func (s *ForStatementContext) SetInit(v IExpressionContext) { s.init = v }

func (s *ForStatementContext) SetCondition(v IExpressionContext) { s.condition = v }

func (s *ForStatementContext) SetPost(v IExpressionContext) { s.post = v }

func (s *ForStatementContext) For() antlr.TerminalNode {
	return s.GetToken(CParserFor, 0)
}

func (s *ForStatementContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *ForStatementContext) AllSemi() []antlr.TerminalNode {
	return s.GetTokens(CParserSemi)
}

func (s *ForStatementContext) Semi(i int) antlr.TerminalNode {
	return s.GetToken(CParserSemi, i)
}

func (s *ForStatementContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *ForStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForStatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ForStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ForStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CParserRULE_forStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Match(CParserFor)
	}
	{
		p.SetState(259)
		p.Match(CParserLeftParen)
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserSizeof, CParserLeftParen, CParserPlus, CParserPlusPlus, CParserMinus, CParserMinusMinus, CParserStar, CParserAnd, CParserNot, CParserTilde, CParserIdentifier, CParserConstant, CParserStringLiteral:
		{
			p.SetState(260)

			var _x = p.expression(0)

			localctx.(*ForStatementContext).init = _x
		}

	case CParserChar, CParserDouble, CParserFloat, CParserInt, CParserLong, CParserShort, CParserStruct, CParserVoid:
		{
			p.SetState(261)
			p.VariableDeclaration()
		}

	case CParserSemi:

	default:
	}
	{
		p.SetState(264)
		p.Match(CParserSemi)
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CParserSizeof-21))|(1<<(CParserLeftParen-21))|(1<<(CParserPlus-21))|(1<<(CParserPlusPlus-21))|(1<<(CParserMinus-21))|(1<<(CParserMinusMinus-21))|(1<<(CParserStar-21))|(1<<(CParserAnd-21)))) != 0) || (((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(CParserNot-54))|(1<<(CParserTilde-54))|(1<<(CParserIdentifier-54))|(1<<(CParserConstant-54))|(1<<(CParserStringLiteral-54)))) != 0) {
		{
			p.SetState(265)

			var _x = p.expression(0)

			localctx.(*ForStatementContext).condition = _x
		}

	}
	{
		p.SetState(268)
		p.Match(CParserSemi)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CParserSizeof-21))|(1<<(CParserLeftParen-21))|(1<<(CParserPlus-21))|(1<<(CParserPlusPlus-21))|(1<<(CParserMinus-21))|(1<<(CParserMinusMinus-21))|(1<<(CParserStar-21))|(1<<(CParserAnd-21)))) != 0) || (((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(CParserNot-54))|(1<<(CParserTilde-54))|(1<<(CParserIdentifier-54))|(1<<(CParserConstant-54))|(1<<(CParserStringLiteral-54)))) != 0) {
		{
			p.SetState(269)

			var _x = p.expression(0)

			localctx.(*ForStatementContext).post = _x
		}

	}
	{
		p.SetState(272)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(273)
		p.Statement()
	}

	return localctx
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_whileStatement
	return p
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(CParserWhile, 0)
}

func (s *WhileStatementContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *WhileStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CParserRULE_whileStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.Match(CParserWhile)
	}
	{
		p.SetState(276)
		p.Match(CParserLeftParen)
	}
	{
		p.SetState(277)
		p.expression(0)
	}
	{
		p.SetState(278)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(279)
		p.Statement()
	}

	return localctx
}

// IIncludePreprocessorContext is an interface to support dynamic dispatch.
type IIncludePreprocessorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncludePreprocessorContext differentiates from other interfaces.
	IsIncludePreprocessorContext()
}

type IncludePreprocessorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludePreprocessorContext() *IncludePreprocessorContext {
	var p = new(IncludePreprocessorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_includePreprocessor
	return p
}

func (*IncludePreprocessorContext) IsIncludePreprocessorContext() {}

func NewIncludePreprocessorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludePreprocessorContext {
	var p = new(IncludePreprocessorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_includePreprocessor

	return p
}

func (s *IncludePreprocessorContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludePreprocessorContext) IncludeDirective() antlr.TerminalNode {
	return s.GetToken(CParserIncludeDirective, 0)
}

func (s *IncludePreprocessorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludePreprocessorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludePreprocessorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitIncludePreprocessor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) IncludePreprocessor() (localctx IIncludePreprocessorContext) {
	localctx = NewIncludePreprocessorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CParserRULE_includePreprocessor)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(CParserIncludeDirective)
	}

	return localctx
}

func (p *CParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *TypeSpecifierContext = nil
		if localctx != nil {
			t = localctx.(*TypeSpecifierContext)
		}
		return p.TypeSpecifier_Sempred(t, predIndex)

	case 13:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CParser) TypeSpecifier_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
