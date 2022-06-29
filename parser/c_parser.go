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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 93, 426,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 7, 2, 74, 10, 2, 12, 2, 14, 2,
	77, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 5, 3, 92, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 98, 10,
	4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 107, 10, 5, 3, 6, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5,
	7, 122, 10, 7, 3, 7, 3, 7, 7, 7, 126, 10, 7, 12, 7, 14, 7, 129, 11, 7,
	3, 8, 3, 8, 3, 8, 5, 8, 134, 10, 8, 3, 8, 5, 8, 137, 10, 8, 3, 9, 3, 9,
	7, 9, 141, 10, 9, 12, 9, 14, 9, 144, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 5,
	10, 150, 10, 10, 3, 10, 3, 10, 3, 11, 3, 11, 5, 11, 156, 10, 11, 3, 11,
	3, 11, 3, 11, 5, 11, 161, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 5,
	12, 168, 10, 12, 3, 12, 3, 12, 5, 12, 172, 10, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 185, 10,
	15, 5, 15, 187, 10, 15, 3, 15, 3, 15, 5, 15, 191, 10, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 7, 16, 197, 10, 16, 12, 16, 14, 16, 200, 11, 16, 3, 16, 5,
	16, 203, 10, 16, 5, 16, 205, 10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 7, 17, 213, 10, 17, 12, 17, 14, 17, 216, 11, 17, 3, 17, 5, 17, 219,
	10, 17, 5, 17, 221, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19,
	239, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	5, 19, 260, 10, 19, 5, 19, 262, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 5, 19, 292, 10, 19, 3, 19, 3, 19, 3, 19, 7, 19, 297,
	10, 19, 12, 19, 14, 19, 300, 11, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 5, 24, 313, 10, 24, 3, 25, 3,
	25, 7, 25, 317, 10, 25, 12, 25, 14, 25, 320, 11, 25, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 333, 10,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26,
	344, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 353,
	10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 361, 10, 28, 12,
	28, 14, 28, 364, 11, 28, 3, 28, 5, 28, 367, 10, 28, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 29, 3, 29, 7, 29, 375, 10, 29, 12, 29, 14, 29, 378, 11, 29, 3,
	30, 3, 30, 3, 30, 7, 30, 383, 10, 30, 12, 30, 14, 30, 386, 11, 30, 3, 31,
	3, 31, 3, 31, 3, 31, 5, 31, 392, 10, 31, 3, 31, 3, 31, 5, 31, 396, 10,
	31, 3, 31, 3, 31, 5, 31, 400, 10, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 2, 4,
	12, 36, 37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
	70, 2, 7, 4, 2, 79, 80, 82, 82, 3, 2, 62, 72, 6, 2, 38, 44, 46, 46, 48,
	55, 73, 74, 4, 2, 45, 45, 47, 47, 5, 2, 44, 48, 51, 51, 56, 57, 2, 471,
	2, 75, 3, 2, 2, 2, 4, 91, 3, 2, 2, 2, 6, 93, 3, 2, 2, 2, 8, 102, 3, 2,
	2, 2, 10, 108, 3, 2, 2, 2, 12, 121, 3, 2, 2, 2, 14, 130, 3, 2, 2, 2, 16,
	138, 3, 2, 2, 2, 18, 147, 3, 2, 2, 2, 20, 153, 3, 2, 2, 2, 22, 164, 3,
	2, 2, 2, 24, 173, 3, 2, 2, 2, 26, 177, 3, 2, 2, 2, 28, 180, 3, 2, 2, 2,
	30, 192, 3, 2, 2, 2, 32, 208, 3, 2, 2, 2, 34, 224, 3, 2, 2, 2, 36, 261,
	3, 2, 2, 2, 38, 301, 3, 2, 2, 2, 40, 303, 3, 2, 2, 2, 42, 305, 3, 2, 2,
	2, 44, 307, 3, 2, 2, 2, 46, 309, 3, 2, 2, 2, 48, 314, 3, 2, 2, 2, 50, 343,
	3, 2, 2, 2, 52, 345, 3, 2, 2, 2, 54, 354, 3, 2, 2, 2, 56, 370, 3, 2, 2,
	2, 58, 379, 3, 2, 2, 2, 60, 387, 3, 2, 2, 2, 62, 404, 3, 2, 2, 2, 64, 411,
	3, 2, 2, 2, 66, 417, 3, 2, 2, 2, 68, 420, 3, 2, 2, 2, 70, 423, 3, 2, 2,
	2, 72, 74, 5, 4, 3, 2, 73, 72, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73,
	3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 3, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2,
	78, 92, 5, 6, 4, 2, 79, 80, 5, 14, 8, 2, 80, 81, 7, 60, 2, 2, 81, 92, 3,
	2, 2, 2, 82, 83, 5, 20, 11, 2, 83, 84, 7, 60, 2, 2, 84, 92, 3, 2, 2, 2,
	85, 92, 5, 70, 36, 2, 86, 87, 5, 24, 13, 2, 87, 88, 7, 60, 2, 2, 88, 92,
	3, 2, 2, 2, 89, 92, 7, 92, 2, 2, 90, 92, 7, 93, 2, 2, 91, 78, 3, 2, 2,
	2, 91, 79, 3, 2, 2, 2, 91, 82, 3, 2, 2, 2, 91, 85, 3, 2, 2, 2, 91, 86,
	3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 90, 3, 2, 2, 2, 92, 5, 3, 2, 2, 2,
	93, 94, 5, 12, 7, 2, 94, 95, 7, 78, 2, 2, 95, 97, 7, 32, 2, 2, 96, 98,
	5, 8, 5, 2, 97, 96, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2,
	99, 100, 7, 33, 2, 2, 100, 101, 5, 48, 25, 2, 101, 7, 3, 2, 2, 2, 102,
	103, 5, 12, 7, 2, 103, 106, 7, 78, 2, 2, 104, 105, 7, 61, 2, 2, 105, 107,
	5, 8, 5, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 9, 3, 2, 2,
	2, 108, 109, 7, 20, 2, 2, 109, 110, 5, 36, 19, 2, 110, 11, 3, 2, 2, 2,
	111, 112, 8, 7, 1, 2, 112, 122, 7, 30, 2, 2, 113, 122, 7, 18, 2, 2, 114,
	122, 7, 21, 2, 2, 115, 122, 7, 19, 2, 2, 116, 122, 7, 5, 2, 2, 117, 122,
	7, 14, 2, 2, 118, 122, 7, 10, 2, 2, 119, 122, 5, 14, 8, 2, 120, 122, 7,
	78, 2, 2, 121, 111, 3, 2, 2, 2, 121, 113, 3, 2, 2, 2, 121, 114, 3, 2, 2,
	2, 121, 115, 3, 2, 2, 2, 121, 116, 3, 2, 2, 2, 121, 117, 3, 2, 2, 2, 121,
	118, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121, 120, 3, 2, 2, 2, 122, 127,
	3, 2, 2, 2, 123, 124, 12, 3, 2, 2, 124, 126, 7, 48, 2, 2, 125, 123, 3,
	2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2,
	2, 128, 13, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 130, 136, 7, 25, 2, 2, 131,
	133, 7, 78, 2, 2, 132, 134, 5, 16, 9, 2, 133, 132, 3, 2, 2, 2, 133, 134,
	3, 2, 2, 2, 134, 137, 3, 2, 2, 2, 135, 137, 5, 16, 9, 2, 136, 131, 3, 2,
	2, 2, 136, 135, 3, 2, 2, 2, 137, 15, 3, 2, 2, 2, 138, 142, 7, 36, 2, 2,
	139, 141, 5, 18, 10, 2, 140, 139, 3, 2, 2, 2, 141, 144, 3, 2, 2, 2, 142,
	140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145, 3, 2, 2, 2, 144, 142,
	3, 2, 2, 2, 145, 146, 7, 37, 2, 2, 146, 17, 3, 2, 2, 2, 147, 149, 5, 12,
	7, 2, 148, 150, 7, 78, 2, 2, 149, 148, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2,
	150, 151, 3, 2, 2, 2, 151, 152, 7, 60, 2, 2, 152, 19, 3, 2, 2, 2, 153,
	155, 7, 12, 2, 2, 154, 156, 7, 78, 2, 2, 155, 154, 3, 2, 2, 2, 155, 156,
	3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 158, 7, 36, 2, 2, 158, 160, 5, 22,
	12, 2, 159, 161, 7, 61, 2, 2, 160, 159, 3, 2, 2, 2, 160, 161, 3, 2, 2,
	2, 161, 162, 3, 2, 2, 2, 162, 163, 7, 37, 2, 2, 163, 21, 3, 2, 2, 2, 164,
	167, 7, 78, 2, 2, 165, 166, 7, 62, 2, 2, 166, 168, 5, 36, 19, 2, 167, 165,
	3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 171, 3, 2, 2, 2, 169, 170, 7, 61,
	2, 2, 170, 172, 5, 22, 12, 2, 171, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2,
	2, 172, 23, 3, 2, 2, 2, 173, 174, 7, 27, 2, 2, 174, 175, 5, 12, 7, 2, 175,
	176, 7, 78, 2, 2, 176, 25, 3, 2, 2, 2, 177, 178, 5, 12, 7, 2, 178, 179,
	5, 28, 15, 2, 179, 27, 3, 2, 2, 2, 180, 186, 7, 78, 2, 2, 181, 184, 7,
	62, 2, 2, 182, 185, 5, 36, 19, 2, 183, 185, 5, 30, 16, 2, 184, 182, 3,
	2, 2, 2, 184, 183, 3, 2, 2, 2, 185, 187, 3, 2, 2, 2, 186, 181, 3, 2, 2,
	2, 186, 187, 3, 2, 2, 2, 187, 190, 3, 2, 2, 2, 188, 189, 7, 61, 2, 2, 189,
	191, 5, 28, 15, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 29,
	3, 2, 2, 2, 192, 204, 7, 36, 2, 2, 193, 198, 5, 36, 19, 2, 194, 195, 7,
	61, 2, 2, 195, 197, 5, 36, 19, 2, 196, 194, 3, 2, 2, 2, 197, 200, 3, 2,
	2, 2, 198, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 202, 3, 2, 2, 2,
	200, 198, 3, 2, 2, 2, 201, 203, 7, 61, 2, 2, 202, 201, 3, 2, 2, 2, 202,
	203, 3, 2, 2, 2, 203, 205, 3, 2, 2, 2, 204, 193, 3, 2, 2, 2, 204, 205,
	3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207, 7, 37, 2, 2, 207, 31, 3, 2,
	2, 2, 208, 220, 7, 36, 2, 2, 209, 214, 5, 34, 18, 2, 210, 211, 7, 61, 2,
	2, 211, 213, 5, 34, 18, 2, 212, 210, 3, 2, 2, 2, 213, 216, 3, 2, 2, 2,
	214, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 218, 3, 2, 2, 2, 216,
	214, 3, 2, 2, 2, 217, 219, 7, 61, 2, 2, 218, 217, 3, 2, 2, 2, 218, 219,
	3, 2, 2, 2, 219, 221, 3, 2, 2, 2, 220, 209, 3, 2, 2, 2, 220, 221, 3, 2,
	2, 2, 221, 222, 3, 2, 2, 2, 222, 223, 7, 37, 2, 2, 223, 33, 3, 2, 2, 2,
	224, 225, 7, 76, 2, 2, 225, 226, 7, 78, 2, 2, 226, 227, 7, 62, 2, 2, 227,
	228, 5, 36, 19, 2, 228, 35, 3, 2, 2, 2, 229, 230, 8, 19, 1, 2, 230, 262,
	7, 78, 2, 2, 231, 262, 9, 2, 2, 2, 232, 262, 7, 83, 2, 2, 233, 234, 7,
	32, 2, 2, 234, 235, 5, 12, 7, 2, 235, 238, 7, 33, 2, 2, 236, 239, 5, 30,
	16, 2, 237, 239, 5, 32, 17, 2, 238, 236, 3, 2, 2, 2, 238, 237, 3, 2, 2,
	2, 239, 262, 3, 2, 2, 2, 240, 241, 7, 32, 2, 2, 241, 242, 5, 12, 7, 2,
	242, 243, 7, 33, 2, 2, 243, 244, 5, 36, 19, 10, 244, 262, 3, 2, 2, 2, 245,
	246, 7, 32, 2, 2, 246, 247, 5, 36, 19, 2, 247, 248, 7, 33, 2, 2, 248, 262,
	3, 2, 2, 2, 249, 250, 5, 44, 23, 2, 250, 251, 5, 36, 19, 7, 251, 262, 3,
	2, 2, 2, 252, 259, 7, 23, 2, 2, 253, 254, 7, 32, 2, 2, 254, 255, 5, 12,
	7, 2, 255, 256, 7, 33, 2, 2, 256, 260, 3, 2, 2, 2, 257, 260, 5, 12, 7,
	2, 258, 260, 5, 36, 19, 2, 259, 253, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2,
	259, 258, 3, 2, 2, 2, 260, 262, 3, 2, 2, 2, 261, 229, 3, 2, 2, 2, 261,
	231, 3, 2, 2, 2, 261, 232, 3, 2, 2, 2, 261, 233, 3, 2, 2, 2, 261, 240,
	3, 2, 2, 2, 261, 245, 3, 2, 2, 2, 261, 249, 3, 2, 2, 2, 261, 252, 3, 2,
	2, 2, 262, 298, 3, 2, 2, 2, 263, 264, 12, 6, 2, 2, 264, 265, 5, 38, 20,
	2, 265, 266, 5, 36, 19, 7, 266, 297, 3, 2, 2, 2, 267, 268, 12, 5, 2, 2,
	268, 269, 5, 40, 21, 2, 269, 270, 5, 36, 19, 6, 270, 297, 3, 2, 2, 2, 271,
	272, 12, 3, 2, 2, 272, 273, 7, 58, 2, 2, 273, 274, 5, 36, 19, 2, 274, 275,
	7, 59, 2, 2, 275, 276, 5, 36, 19, 4, 276, 297, 3, 2, 2, 2, 277, 278, 12,
	14, 2, 2, 278, 279, 7, 76, 2, 2, 279, 297, 7, 78, 2, 2, 280, 281, 12, 13,
	2, 2, 281, 282, 7, 75, 2, 2, 282, 297, 7, 78, 2, 2, 283, 284, 12, 12, 2,
	2, 284, 285, 7, 34, 2, 2, 285, 286, 5, 36, 19, 2, 286, 287, 7, 35, 2, 2,
	287, 297, 3, 2, 2, 2, 288, 289, 12, 11, 2, 2, 289, 291, 7, 32, 2, 2, 290,
	292, 5, 46, 24, 2, 291, 290, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 293,
	3, 2, 2, 2, 293, 297, 7, 33, 2, 2, 294, 295, 12, 8, 2, 2, 295, 297, 5,
	42, 22, 2, 296, 263, 3, 2, 2, 2, 296, 267, 3, 2, 2, 2, 296, 271, 3, 2,
	2, 2, 296, 277, 3, 2, 2, 2, 296, 280, 3, 2, 2, 2, 296, 283, 3, 2, 2, 2,
	296, 288, 3, 2, 2, 2, 296, 294, 3, 2, 2, 2, 297, 300, 3, 2, 2, 2, 298,
	296, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 37, 3, 2, 2, 2, 300, 298, 3,
	2, 2, 2, 301, 302, 9, 3, 2, 2, 302, 39, 3, 2, 2, 2, 303, 304, 9, 4, 2,
	2, 304, 41, 3, 2, 2, 2, 305, 306, 9, 5, 2, 2, 306, 43, 3, 2, 2, 2, 307,
	308, 9, 6, 2, 2, 308, 45, 3, 2, 2, 2, 309, 312, 5, 36, 19, 2, 310, 311,
	7, 61, 2, 2, 311, 313, 5, 46, 24, 2, 312, 310, 3, 2, 2, 2, 312, 313, 3,
	2, 2, 2, 313, 47, 3, 2, 2, 2, 314, 318, 7, 36, 2, 2, 315, 317, 5, 50, 26,
	2, 316, 315, 3, 2, 2, 2, 317, 320, 3, 2, 2, 2, 318, 316, 3, 2, 2, 2, 318,
	319, 3, 2, 2, 2, 319, 321, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 321, 322,
	7, 37, 2, 2, 322, 49, 3, 2, 2, 2, 323, 333, 5, 26, 14, 2, 324, 333, 5,
	36, 19, 2, 325, 333, 5, 10, 6, 2, 326, 333, 7, 3, 2, 2, 327, 333, 7, 7,
	2, 2, 328, 333, 5, 14, 8, 2, 329, 333, 5, 20, 11, 2, 330, 333, 5, 66, 34,
	2, 331, 333, 5, 62, 32, 2, 332, 323, 3, 2, 2, 2, 332, 324, 3, 2, 2, 2,
	332, 325, 3, 2, 2, 2, 332, 326, 3, 2, 2, 2, 332, 327, 3, 2, 2, 2, 332,
	328, 3, 2, 2, 2, 332, 329, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2, 332, 331,
	3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 344, 7, 60, 2, 2, 335, 344, 5, 52,
	27, 2, 336, 344, 5, 54, 28, 2, 337, 344, 5, 60, 31, 2, 338, 344, 5, 64,
	33, 2, 339, 344, 5, 48, 25, 2, 340, 344, 5, 68, 35, 2, 341, 344, 7, 92,
	2, 2, 342, 344, 7, 93, 2, 2, 343, 332, 3, 2, 2, 2, 343, 335, 3, 2, 2, 2,
	343, 336, 3, 2, 2, 2, 343, 337, 3, 2, 2, 2, 343, 338, 3, 2, 2, 2, 343,
	339, 3, 2, 2, 2, 343, 340, 3, 2, 2, 2, 343, 341, 3, 2, 2, 2, 343, 342,
	3, 2, 2, 2, 344, 51, 3, 2, 2, 2, 345, 346, 7, 17, 2, 2, 346, 347, 7, 32,
	2, 2, 347, 348, 5, 36, 19, 2, 348, 349, 7, 33, 2, 2, 349, 352, 5, 50, 26,
	2, 350, 351, 7, 11, 2, 2, 351, 353, 5, 50, 26, 2, 352, 350, 3, 2, 2, 2,
	352, 353, 3, 2, 2, 2, 353, 53, 3, 2, 2, 2, 354, 355, 7, 26, 2, 2, 355,
	356, 7, 32, 2, 2, 356, 357, 5, 36, 19, 2, 357, 358, 7, 33, 2, 2, 358, 362,
	7, 36, 2, 2, 359, 361, 5, 56, 29, 2, 360, 359, 3, 2, 2, 2, 361, 364, 3,
	2, 2, 2, 362, 360, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 366, 3, 2, 2,
	2, 364, 362, 3, 2, 2, 2, 365, 367, 5, 58, 30, 2, 366, 365, 3, 2, 2, 2,
	366, 367, 3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368, 369, 7, 37, 2, 2, 369,
	55, 3, 2, 2, 2, 370, 371, 7, 4, 2, 2, 371, 372, 5, 36, 19, 2, 372, 376,
	7, 59, 2, 2, 373, 375, 5, 50, 26, 2, 374, 373, 3, 2, 2, 2, 375, 378, 3,
	2, 2, 2, 376, 374, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 57, 3, 2, 2,
	2, 378, 376, 3, 2, 2, 2, 379, 380, 7, 8, 2, 2, 380, 384, 7, 59, 2, 2, 381,
	383, 5, 50, 26, 2, 382, 381, 3, 2, 2, 2, 383, 386, 3, 2, 2, 2, 384, 382,
	3, 2, 2, 2, 384, 385, 3, 2, 2, 2, 385, 59, 3, 2, 2, 2, 386, 384, 3, 2,
	2, 2, 387, 388, 7, 15, 2, 2, 388, 391, 7, 32, 2, 2, 389, 392, 5, 36, 19,
	2, 390, 392, 5, 26, 14, 2, 391, 389, 3, 2, 2, 2, 391, 390, 3, 2, 2, 2,
	391, 392, 3, 2, 2, 2, 392, 393, 3, 2, 2, 2, 393, 395, 7, 60, 2, 2, 394,
	396, 5, 36, 19, 2, 395, 394, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396, 397,
	3, 2, 2, 2, 397, 399, 7, 60, 2, 2, 398, 400, 5, 36, 19, 2, 399, 398, 3,
	2, 2, 2, 399, 400, 3, 2, 2, 2, 400, 401, 3, 2, 2, 2, 401, 402, 7, 33, 2,
	2, 402, 403, 5, 50, 26, 2, 403, 61, 3, 2, 2, 2, 404, 405, 7, 9, 2, 2, 405,
	406, 5, 50, 26, 2, 406, 407, 7, 31, 2, 2, 407, 408, 7, 32, 2, 2, 408, 409,
	5, 36, 19, 2, 409, 410, 7, 33, 2, 2, 410, 63, 3, 2, 2, 2, 411, 412, 7,
	31, 2, 2, 412, 413, 7, 32, 2, 2, 413, 414, 5, 36, 19, 2, 414, 415, 7, 33,
	2, 2, 415, 416, 5, 50, 26, 2, 416, 65, 3, 2, 2, 2, 417, 418, 7, 16, 2,
	2, 418, 419, 7, 78, 2, 2, 419, 67, 3, 2, 2, 2, 420, 421, 7, 78, 2, 2, 421,
	422, 7, 59, 2, 2, 422, 69, 3, 2, 2, 2, 423, 424, 7, 85, 2, 2, 424, 71,
	3, 2, 2, 2, 43, 75, 91, 97, 106, 121, 127, 133, 136, 142, 149, 155, 160,
	167, 171, 184, 186, 190, 198, 202, 204, 214, 218, 220, 238, 259, 261, 291,
	296, 298, 312, 318, 332, 343, 352, 362, 366, 376, 384, 391, 395, 399,
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
	"Arrow", "Dot", "Ellipsis", "Identifier", "IntegerConstant", "FloatingConstant",
	"DigitSequence", "CharacterConstant", "StringLiteral", "ComplexDefine",
	"IncludeDirective", "AsmBlock", "LineAfterPreprocessing", "LineDirective",
	"PragmaDirective", "Whitespace", "Newline", "BlockComment", "LineComment",
}

var ruleNames = []string{
	"translation", "declaration", "functionDeclaration", "functionArguments",
	"functionReturn", "typeSpecifier", "structDeclaration", "structDeclarationBody",
	"structProperty", "enumDeclaration", "enumProperties", "typedefDeclaration",
	"variableDeclaration", "variableDeclarationList", "listInitialization",
	"namedListInitialization", "namedListInitializationItem", "expression",
	"assignementOperator", "binaryOperator", "unaryOperatorPost", "unaryOperatorPre",
	"functionCallArguments", "block", "statement", "ifStatement", "switchStatement",
	"caseStatement", "defaultStatement", "forStatement", "doWhileStatement",
	"whileStatement", "gotoStatement", "labelStatement", "includePreprocessor",
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
	CParserIntegerConstant        = 77
	CParserFloatingConstant       = 78
	CParserDigitSequence          = 79
	CParserCharacterConstant      = 80
	CParserStringLiteral          = 81
	CParserComplexDefine          = 82
	CParserIncludeDirective       = 83
	CParserAsmBlock               = 84
	CParserLineAfterPreprocessing = 85
	CParserLineDirective          = 86
	CParserPragmaDirective        = 87
	CParserWhitespace             = 88
	CParserNewline                = 89
	CParserBlockComment           = 90
	CParserLineComment            = 91
)

// CParser rules.
const (
	CParserRULE_translation                 = 0
	CParserRULE_declaration                 = 1
	CParserRULE_functionDeclaration         = 2
	CParserRULE_functionArguments           = 3
	CParserRULE_functionReturn              = 4
	CParserRULE_typeSpecifier               = 5
	CParserRULE_structDeclaration           = 6
	CParserRULE_structDeclarationBody       = 7
	CParserRULE_structProperty              = 8
	CParserRULE_enumDeclaration             = 9
	CParserRULE_enumProperties              = 10
	CParserRULE_typedefDeclaration          = 11
	CParserRULE_variableDeclaration         = 12
	CParserRULE_variableDeclarationList     = 13
	CParserRULE_listInitialization          = 14
	CParserRULE_namedListInitialization     = 15
	CParserRULE_namedListInitializationItem = 16
	CParserRULE_expression                  = 17
	CParserRULE_assignementOperator         = 18
	CParserRULE_binaryOperator              = 19
	CParserRULE_unaryOperatorPost           = 20
	CParserRULE_unaryOperatorPre            = 21
	CParserRULE_functionCallArguments       = 22
	CParserRULE_block                       = 23
	CParserRULE_statement                   = 24
	CParserRULE_ifStatement                 = 25
	CParserRULE_switchStatement             = 26
	CParserRULE_caseStatement               = 27
	CParserRULE_defaultStatement            = 28
	CParserRULE_forStatement                = 29
	CParserRULE_doWhileStatement            = 30
	CParserRULE_whileStatement              = 31
	CParserRULE_gotoStatement               = 32
	CParserRULE_labelStatement              = 33
	CParserRULE_includePreprocessor         = 34
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
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserEnum)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserStruct)|(1<<CParserTypedef)|(1<<CParserVoid))) != 0) || (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(CParserIdentifier-76))|(1<<(CParserIncludeDirective-76))|(1<<(CParserBlockComment-76))|(1<<(CParserLineComment-76)))) != 0) {
		{
			p.SetState(70)
			p.Declaration()
		}

		p.SetState(75)
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

func (s *DeclarationContext) TypedefDeclaration() ITypedefDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypedefDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypedefDeclarationContext)
}

func (s *DeclarationContext) BlockComment() antlr.TerminalNode {
	return s.GetToken(CParserBlockComment, 0)
}

func (s *DeclarationContext) LineComment() antlr.TerminalNode {
	return s.GetToken(CParserLineComment, 0)
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

	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.FunctionDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.StructDeclaration()
		}
		{
			p.SetState(78)
			p.Match(CParserSemi)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
			p.EnumDeclaration()
		}
		{
			p.SetState(81)
			p.Match(CParserSemi)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.IncludePreprocessor()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(84)
			p.TypedefDeclaration()
		}
		{
			p.SetState(85)
			p.Match(CParserSemi)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(87)
			p.Match(CParserBlockComment)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(88)
			p.Match(CParserLineComment)
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
		p.SetState(91)
		p.typeSpecifier(0)
	}
	{
		p.SetState(92)
		p.Match(CParserIdentifier)
	}
	{
		p.SetState(93)
		p.Match(CParserLeftParen)
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserStruct)|(1<<CParserVoid))) != 0) || _la == CParserIdentifier {
		{
			p.SetState(94)
			p.FunctionArguments()
		}

	}
	{
		p.SetState(97)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(98)
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
		p.SetState(100)
		p.typeSpecifier(0)
	}
	{
		p.SetState(101)
		p.Match(CParserIdentifier)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(102)
			p.Match(CParserComma)
		}
		{
			p.SetState(103)
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
		p.SetState(106)
		p.Match(CParserReturn)
	}
	{
		p.SetState(107)
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

func (s *TypeSpecifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
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
	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserVoid:
		{
			p.SetState(110)
			p.Match(CParserVoid)
		}

	case CParserInt:
		{
			p.SetState(111)
			p.Match(CParserInt)
		}

	case CParserShort:
		{
			p.SetState(112)
			p.Match(CParserShort)
		}

	case CParserLong:
		{
			p.SetState(113)
			p.Match(CParserLong)
		}

	case CParserChar:
		{
			p.SetState(114)
			p.Match(CParserChar)
		}

	case CParserFloat:
		{
			p.SetState(115)
			p.Match(CParserFloat)
		}

	case CParserDouble:
		{
			p.SetState(116)
			p.Match(CParserDouble)
		}

	case CParserStruct:
		{
			p.SetState(117)
			p.StructDeclaration()
		}

	case CParserIdentifier:
		{
			p.SetState(118)
			p.Match(CParserIdentifier)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(125)
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
			p.SetState(121)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(122)
				p.Match(CParserStar)
			}

		}
		p.SetState(127)
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
		p.SetState(128)
		p.Match(CParserStruct)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserIdentifier:
		{
			p.SetState(129)
			p.Match(CParserIdentifier)
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(130)
				p.StructDeclarationBody()
			}

		}

	case CParserLeftBrace:
		{
			p.SetState(133)
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
		p.SetState(136)
		p.Match(CParserLeftBrace)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserStruct)|(1<<CParserVoid))) != 0) || _la == CParserIdentifier {
		{
			p.SetState(137)
			p.StructProperty()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
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
		p.SetState(145)
		p.typeSpecifier(0)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserIdentifier {
		{
			p.SetState(146)
			p.Match(CParserIdentifier)
		}

	}
	{
		p.SetState(149)
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
		p.SetState(151)
		p.Match(CParserEnum)
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserIdentifier {
		{
			p.SetState(152)
			p.Match(CParserIdentifier)
		}

	}
	{
		p.SetState(155)
		p.Match(CParserLeftBrace)
	}
	{
		p.SetState(156)
		p.EnumProperties()
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(157)
			p.Match(CParserComma)
		}

	}
	{
		p.SetState(160)
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
		p.SetState(162)
		p.Match(CParserIdentifier)
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserAssign {
		{
			p.SetState(163)
			p.Match(CParserAssign)
		}
		{
			p.SetState(164)
			p.expression(0)
		}

	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(167)
			p.Match(CParserComma)
		}
		{
			p.SetState(168)
			p.EnumProperties()
		}

	}

	return localctx
}

// ITypedefDeclarationContext is an interface to support dynamic dispatch.
type ITypedefDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypedefDeclarationContext differentiates from other interfaces.
	IsTypedefDeclarationContext()
}

type TypedefDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedefDeclarationContext() *TypedefDeclarationContext {
	var p = new(TypedefDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_typedefDeclaration
	return p
}

func (*TypedefDeclarationContext) IsTypedefDeclarationContext() {}

func NewTypedefDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedefDeclarationContext {
	var p = new(TypedefDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_typedefDeclaration

	return p
}

func (s *TypedefDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedefDeclarationContext) Typedef() antlr.TerminalNode {
	return s.GetToken(CParserTypedef, 0)
}

func (s *TypedefDeclarationContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *TypedefDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *TypedefDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedefDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedefDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitTypedefDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) TypedefDeclaration() (localctx ITypedefDeclarationContext) {
	localctx = NewTypedefDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CParserRULE_typedefDeclaration)

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
		p.SetState(171)
		p.Match(CParserTypedef)
	}
	{
		p.SetState(172)
		p.typeSpecifier(0)
	}
	{
		p.SetState(173)
		p.Match(CParserIdentifier)
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
	p.EnterRule(localctx, 24, CParserRULE_variableDeclaration)

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
		p.SetState(175)
		p.typeSpecifier(0)
	}
	{
		p.SetState(176)
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

func (s *VariableDeclarationListContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationListContext) ListInitialization() IListInitializationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListInitializationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListInitializationContext)
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
	p.EnterRule(localctx, 26, CParserRULE_variableDeclarationList)
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
		p.SetState(178)
		p.Match(CParserIdentifier)
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserAssign {
		{
			p.SetState(179)
			p.Match(CParserAssign)
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CParserSizeof, CParserLeftParen, CParserPlus, CParserPlusPlus, CParserMinus, CParserMinusMinus, CParserStar, CParserAnd, CParserNot, CParserTilde, CParserIdentifier, CParserIntegerConstant, CParserFloatingConstant, CParserCharacterConstant, CParserStringLiteral:
			{
				p.SetState(180)
				p.expression(0)
			}

		case CParserLeftBrace:
			{
				p.SetState(181)
				p.ListInitialization()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(186)
			p.Match(CParserComma)
		}
		{
			p.SetState(187)
			p.VariableDeclarationList()
		}

	}

	return localctx
}

// IListInitializationContext is an interface to support dynamic dispatch.
type IListInitializationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListInitializationContext differentiates from other interfaces.
	IsListInitializationContext()
}

type ListInitializationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListInitializationContext() *ListInitializationContext {
	var p = new(ListInitializationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_listInitialization
	return p
}

func (*ListInitializationContext) IsListInitializationContext() {}

func NewListInitializationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListInitializationContext {
	var p = new(ListInitializationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_listInitialization

	return p
}

func (s *ListInitializationContext) GetParser() antlr.Parser { return s.parser }

func (s *ListInitializationContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(CParserLeftBrace, 0)
}

func (s *ListInitializationContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(CParserRightBrace, 0)
}

func (s *ListInitializationContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ListInitializationContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListInitializationContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(CParserComma)
}

func (s *ListInitializationContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(CParserComma, i)
}

func (s *ListInitializationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListInitializationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListInitializationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitListInitialization(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) ListInitialization() (localctx IListInitializationContext) {
	localctx = NewListInitializationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CParserRULE_listInitialization)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(CParserLeftBrace)
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CParserSizeof-21))|(1<<(CParserLeftParen-21))|(1<<(CParserPlus-21))|(1<<(CParserPlusPlus-21))|(1<<(CParserMinus-21))|(1<<(CParserMinusMinus-21))|(1<<(CParserStar-21))|(1<<(CParserAnd-21)))) != 0) || (((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(CParserNot-54))|(1<<(CParserTilde-54))|(1<<(CParserIdentifier-54))|(1<<(CParserIntegerConstant-54))|(1<<(CParserFloatingConstant-54))|(1<<(CParserCharacterConstant-54))|(1<<(CParserStringLiteral-54)))) != 0) {
		{
			p.SetState(191)
			p.expression(0)
		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(192)
					p.Match(CParserComma)
				}
				{
					p.SetState(193)
					p.expression(0)
				}

			}
			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CParserComma {
			{
				p.SetState(199)
				p.Match(CParserComma)
			}

		}

	}
	{
		p.SetState(204)
		p.Match(CParserRightBrace)
	}

	return localctx
}

// INamedListInitializationContext is an interface to support dynamic dispatch.
type INamedListInitializationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedListInitializationContext differentiates from other interfaces.
	IsNamedListInitializationContext()
}

type NamedListInitializationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedListInitializationContext() *NamedListInitializationContext {
	var p = new(NamedListInitializationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_namedListInitialization
	return p
}

func (*NamedListInitializationContext) IsNamedListInitializationContext() {}

func NewNamedListInitializationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedListInitializationContext {
	var p = new(NamedListInitializationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_namedListInitialization

	return p
}

func (s *NamedListInitializationContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedListInitializationContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(CParserLeftBrace, 0)
}

func (s *NamedListInitializationContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(CParserRightBrace, 0)
}

func (s *NamedListInitializationContext) AllNamedListInitializationItem() []INamedListInitializationItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INamedListInitializationItemContext)(nil)).Elem())
	var tst = make([]INamedListInitializationItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INamedListInitializationItemContext)
		}
	}

	return tst
}

func (s *NamedListInitializationContext) NamedListInitializationItem(i int) INamedListInitializationItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedListInitializationItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INamedListInitializationItemContext)
}

func (s *NamedListInitializationContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(CParserComma)
}

func (s *NamedListInitializationContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(CParserComma, i)
}

func (s *NamedListInitializationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedListInitializationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedListInitializationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitNamedListInitialization(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) NamedListInitialization() (localctx INamedListInitializationContext) {
	localctx = NewNamedListInitializationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CParserRULE_namedListInitialization)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(CParserLeftBrace)
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserDot {
		{
			p.SetState(207)
			p.NamedListInitializationItem()
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(208)
					p.Match(CParserComma)
				}
				{
					p.SetState(209)
					p.NamedListInitializationItem()
				}

			}
			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CParserComma {
			{
				p.SetState(215)
				p.Match(CParserComma)
			}

		}

	}
	{
		p.SetState(220)
		p.Match(CParserRightBrace)
	}

	return localctx
}

// INamedListInitializationItemContext is an interface to support dynamic dispatch.
type INamedListInitializationItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedListInitializationItemContext differentiates from other interfaces.
	IsNamedListInitializationItemContext()
}

type NamedListInitializationItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedListInitializationItemContext() *NamedListInitializationItemContext {
	var p = new(NamedListInitializationItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_namedListInitializationItem
	return p
}

func (*NamedListInitializationItemContext) IsNamedListInitializationItemContext() {}

func NewNamedListInitializationItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedListInitializationItemContext {
	var p = new(NamedListInitializationItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_namedListInitializationItem

	return p
}

func (s *NamedListInitializationItemContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedListInitializationItemContext) Dot() antlr.TerminalNode {
	return s.GetToken(CParserDot, 0)
}

func (s *NamedListInitializationItemContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *NamedListInitializationItemContext) Assign() antlr.TerminalNode {
	return s.GetToken(CParserAssign, 0)
}

func (s *NamedListInitializationItemContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NamedListInitializationItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedListInitializationItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedListInitializationItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitNamedListInitializationItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) NamedListInitializationItem() (localctx INamedListInitializationItemContext) {
	localctx = NewNamedListInitializationItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CParserRULE_namedListInitializationItem)

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
		p.SetState(222)
		p.Match(CParserDot)
	}
	{
		p.SetState(223)
		p.Match(CParserIdentifier)
	}
	{
		p.SetState(224)
		p.Match(CParserAssign)
	}
	{
		p.SetState(225)
		p.expression(0)
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

type StructInitializationExpressionContext struct {
	*ExpressionContext
}

func NewStructInitializationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructInitializationExpressionContext {
	var p = new(StructInitializationExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *StructInitializationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInitializationExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *StructInitializationExpressionContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *StructInitializationExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *StructInitializationExpressionContext) ListInitialization() IListInitializationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListInitializationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListInitializationContext)
}

func (s *StructInitializationExpressionContext) NamedListInitialization() INamedListInitializationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedListInitializationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedListInitializationContext)
}

func (s *StructInitializationExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitStructInitializationExpression(s)

	default:
		return t.VisitChildren(s)
	}
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

type TernaryExpressionContext struct {
	*ExpressionContext
}

func NewTernaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExpressionContext {
	var p = new(TernaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TernaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *TernaryExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TernaryExpressionContext) Question() antlr.TerminalNode {
	return s.GetToken(CParserQuestion, 0)
}

func (s *TernaryExpressionContext) Colon() antlr.TerminalNode {
	return s.GetToken(CParserColon, 0)
}

func (s *TernaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitTernaryExpression(s)

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

func (s *ConstantExpressionContext) IntegerConstant() antlr.TerminalNode {
	return s.GetToken(CParserIntegerConstant, 0)
}

func (s *ConstantExpressionContext) FloatingConstant() antlr.TerminalNode {
	return s.GetToken(CParserFloatingConstant, 0)
}

func (s *ConstantExpressionContext) CharacterConstant() antlr.TerminalNode {
	return s.GetToken(CParserCharacterConstant, 0)
}

func (s *ConstantExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitConstantExpression(s)

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

type PropertyAccessExpressionContext struct {
	*ExpressionContext
}

func NewPropertyAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyAccessExpressionContext {
	var p = new(PropertyAccessExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PropertyAccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyAccessExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PropertyAccessExpressionContext) Dot() antlr.TerminalNode {
	return s.GetToken(CParserDot, 0)
}

func (s *PropertyAccessExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *PropertyAccessExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitPropertyAccessExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type PointerPropertyAccessExpressionContext struct {
	*ExpressionContext
}

func NewPointerPropertyAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PointerPropertyAccessExpressionContext {
	var p = new(PointerPropertyAccessExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PointerPropertyAccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerPropertyAccessExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PointerPropertyAccessExpressionContext) Arrow() antlr.TerminalNode {
	return s.GetToken(CParserArrow, 0)
}

func (s *PointerPropertyAccessExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *PointerPropertyAccessExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitPointerPropertyAccessExpression(s)

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

func (p *CParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *CParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, CParserRULE_expression, _p)
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
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(228)
			p.Match(CParserIdentifier)
		}

	case 2:
		localctx = NewConstantExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(229)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-77)&-(0x1f+1)) == 0 && ((1<<uint((_la-77)))&((1<<(CParserIntegerConstant-77))|(1<<(CParserFloatingConstant-77))|(1<<(CParserCharacterConstant-77)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 3:
		localctx = NewConstantStringExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(230)
			p.Match(CParserStringLiteral)
		}

	case 4:
		localctx = NewStructInitializationExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(231)
			p.Match(CParserLeftParen)
		}
		{
			p.SetState(232)
			p.typeSpecifier(0)
		}
		{
			p.SetState(233)
			p.Match(CParserRightParen)
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(234)
				p.ListInitialization()
			}

		case 2:
			{
				p.SetState(235)
				p.NamedListInitialization()
			}

		}

	case 5:
		localctx = NewCastExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(238)
			p.Match(CParserLeftParen)
		}
		{
			p.SetState(239)
			p.typeSpecifier(0)
		}
		{
			p.SetState(240)
			p.Match(CParserRightParen)
		}
		{
			p.SetState(241)
			p.expression(8)
		}

	case 6:
		localctx = NewParenthesizedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(243)
			p.Match(CParserLeftParen)
		}
		{
			p.SetState(244)
			p.expression(0)
		}
		{
			p.SetState(245)
			p.Match(CParserRightParen)
		}

	case 7:
		localctx = NewUnaryExpressionPreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(247)
			p.UnaryOperatorPre()
		}
		{
			p.SetState(248)
			p.expression(5)
		}

	case 8:
		localctx = NewSizeofExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(250)
			p.Match(CParserSizeof)
		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(251)
				p.Match(CParserLeftParen)
			}
			{
				p.SetState(252)
				p.typeSpecifier(0)
			}
			{
				p.SetState(253)
				p.Match(CParserRightParen)
			}

		case 2:
			{
				p.SetState(255)
				p.typeSpecifier(0)
			}

		case 3:
			{
				p.SetState(256)
				p.expression(0)
			}

		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(294)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAssignmentExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(261)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(262)
					p.AssignementOperator()
				}
				{
					p.SetState(263)
					p.expression(5)
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(265)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(266)
					p.BinaryOperator()
				}
				{
					p.SetState(267)
					p.expression(4)
				}

			case 3:
				localctx = NewTernaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(270)
					p.Match(CParserQuestion)
				}
				{
					p.SetState(271)
					p.expression(0)
				}
				{
					p.SetState(272)
					p.Match(CParserColon)
				}
				{
					p.SetState(273)
					p.expression(2)
				}

			case 4:
				localctx = NewPropertyAccessExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(276)
					p.Match(CParserDot)
				}
				{
					p.SetState(277)
					p.Match(CParserIdentifier)
				}

			case 5:
				localctx = NewPointerPropertyAccessExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(278)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(279)
					p.Match(CParserArrow)
				}
				{
					p.SetState(280)
					p.Match(CParserIdentifier)
				}

			case 6:
				localctx = NewArrayIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(282)
					p.Match(CParserLeftBracket)
				}
				{
					p.SetState(283)
					p.expression(0)
				}
				{
					p.SetState(284)
					p.Match(CParserRightBracket)
				}

			case 7:
				localctx = NewFunctionCallExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(286)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(287)
					p.Match(CParserLeftParen)
				}
				p.SetState(289)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CParserSizeof-21))|(1<<(CParserLeftParen-21))|(1<<(CParserPlus-21))|(1<<(CParserPlusPlus-21))|(1<<(CParserMinus-21))|(1<<(CParserMinusMinus-21))|(1<<(CParserStar-21))|(1<<(CParserAnd-21)))) != 0) || (((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(CParserNot-54))|(1<<(CParserTilde-54))|(1<<(CParserIdentifier-54))|(1<<(CParserIntegerConstant-54))|(1<<(CParserFloatingConstant-54))|(1<<(CParserCharacterConstant-54))|(1<<(CParserStringLiteral-54)))) != 0) {
					{
						p.SetState(288)
						p.FunctionCallArguments()
					}

				}
				{
					p.SetState(291)
					p.Match(CParserRightParen)
				}

			case 8:
				localctx = NewUnaryExpressionPostContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(293)
					p.UnaryOperatorPost()
				}

			}

		}
		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 36, CParserRULE_assignementOperator)
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
		p.SetState(299)
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
	p.EnterRule(localctx, 38, CParserRULE_binaryOperator)
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
		p.SetState(301)
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
	p.EnterRule(localctx, 40, CParserRULE_unaryOperatorPost)
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
		p.SetState(303)
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
	p.EnterRule(localctx, 42, CParserRULE_unaryOperatorPre)
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
		p.SetState(305)
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
	p.EnterRule(localctx, 44, CParserRULE_functionCallArguments)
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
		p.SetState(307)
		p.expression(0)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(308)
			p.Match(CParserComma)
		}
		{
			p.SetState(309)
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
	p.EnterRule(localctx, 46, CParserRULE_block)
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
		p.SetState(312)
		p.Match(CParserLeftBrace)
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserBreak)|(1<<CParserChar)|(1<<CParserContinue)|(1<<CParserDo)|(1<<CParserDouble)|(1<<CParserEnum)|(1<<CParserFloat)|(1<<CParserFor)|(1<<CParserGoto)|(1<<CParserIf)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserReturn)|(1<<CParserShort)|(1<<CParserSizeof)|(1<<CParserStruct)|(1<<CParserSwitch)|(1<<CParserVoid)|(1<<CParserWhile)|(1<<CParserLeftParen))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(CParserLeftBrace-34))|(1<<(CParserPlus-34))|(1<<(CParserPlusPlus-34))|(1<<(CParserMinus-34))|(1<<(CParserMinusMinus-34))|(1<<(CParserStar-34))|(1<<(CParserAnd-34))|(1<<(CParserNot-34))|(1<<(CParserTilde-34)))) != 0) || (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(CParserIdentifier-76))|(1<<(CParserIntegerConstant-76))|(1<<(CParserFloatingConstant-76))|(1<<(CParserCharacterConstant-76))|(1<<(CParserStringLiteral-76))|(1<<(CParserBlockComment-76))|(1<<(CParserLineComment-76)))) != 0) {
		{
			p.SetState(313)
			p.Statement()
		}

		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(319)
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

func (s *StatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(CParserContinue, 0)
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

func (s *StatementContext) GotoStatement() IGotoStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGotoStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGotoStatementContext)
}

func (s *StatementContext) DoWhileStatement() IDoWhileStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoWhileStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoWhileStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) SwitchStatement() ISwitchStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
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

func (s *StatementContext) LabelStatement() ILabelStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelStatementContext)
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
	p.EnterRule(localctx, 48, CParserRULE_statement)

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

	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(321)
				p.VariableDeclaration()
			}

		case 2:
			{
				p.SetState(322)
				p.expression(0)
			}

		case 3:
			{
				p.SetState(323)
				p.FunctionReturn()
			}

		case 4:
			{
				p.SetState(324)
				p.Match(CParserBreak)
			}

		case 5:
			{
				p.SetState(325)
				p.Match(CParserContinue)
			}

		case 6:
			{
				p.SetState(326)
				p.StructDeclaration()
			}

		case 7:
			{
				p.SetState(327)
				p.EnumDeclaration()
			}

		case 8:
			{
				p.SetState(328)
				p.GotoStatement()
			}

		case 9:
			{
				p.SetState(329)
				p.DoWhileStatement()
			}

		}
		{
			p.SetState(332)
			p.Match(CParserSemi)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(333)
			p.IfStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(334)
			p.SwitchStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(335)
			p.ForStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(336)
			p.WhileStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(337)
			p.Block()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(338)
			p.LabelStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(339)
			p.Match(CParserBlockComment)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(340)
			p.Match(CParserLineComment)
		}

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
	p.EnterRule(localctx, 50, CParserRULE_ifStatement)

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
		p.SetState(343)
		p.Match(CParserIf)
	}
	{
		p.SetState(344)
		p.Match(CParserLeftParen)
	}
	{
		p.SetState(345)
		p.expression(0)
	}
	{
		p.SetState(346)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(347)
		p.Statement()
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(348)
			p.Match(CParserElse)
		}
		{
			p.SetState(349)
			p.Statement()
		}

	}

	return localctx
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_switchStatement
	return p
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) Switch() antlr.TerminalNode {
	return s.GetToken(CParserSwitch, 0)
}

func (s *SwitchStatementContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *SwitchStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchStatementContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *SwitchStatementContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(CParserLeftBrace, 0)
}

func (s *SwitchStatementContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(CParserRightBrace, 0)
}

func (s *SwitchStatementContext) AllCaseStatement() []ICaseStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseStatementContext)(nil)).Elem())
	var tst = make([]ICaseStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseStatementContext)
		}
	}

	return tst
}

func (s *SwitchStatementContext) CaseStatement(i int) ICaseStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseStatementContext)
}

func (s *SwitchStatementContext) DefaultStatement() IDefaultStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultStatementContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitSwitchStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CParserRULE_switchStatement)
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
		p.SetState(352)
		p.Match(CParserSwitch)
	}
	{
		p.SetState(353)
		p.Match(CParserLeftParen)
	}
	{
		p.SetState(354)
		p.expression(0)
	}
	{
		p.SetState(355)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(356)
		p.Match(CParserLeftBrace)
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserCase {
		{
			p.SetState(357)
			p.CaseStatement()
		}

		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserDefault {
		{
			p.SetState(363)
			p.DefaultStatement()
		}

	}
	{
		p.SetState(366)
		p.Match(CParserRightBrace)
	}

	return localctx
}

// ICaseStatementContext is an interface to support dynamic dispatch.
type ICaseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseStatementContext differentiates from other interfaces.
	IsCaseStatementContext()
}

type CaseStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseStatementContext() *CaseStatementContext {
	var p = new(CaseStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_caseStatement
	return p
}

func (*CaseStatementContext) IsCaseStatementContext() {}

func NewCaseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseStatementContext {
	var p = new(CaseStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_caseStatement

	return p
}

func (s *CaseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseStatementContext) Case() antlr.TerminalNode {
	return s.GetToken(CParserCase, 0)
}

func (s *CaseStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CaseStatementContext) Colon() antlr.TerminalNode {
	return s.GetToken(CParserColon, 0)
}

func (s *CaseStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *CaseStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CaseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitCaseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) CaseStatement() (localctx ICaseStatementContext) {
	localctx = NewCaseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CParserRULE_caseStatement)
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
		p.SetState(368)
		p.Match(CParserCase)
	}
	{
		p.SetState(369)
		p.expression(0)
	}
	{
		p.SetState(370)
		p.Match(CParserColon)
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserBreak)|(1<<CParserChar)|(1<<CParserContinue)|(1<<CParserDo)|(1<<CParserDouble)|(1<<CParserEnum)|(1<<CParserFloat)|(1<<CParserFor)|(1<<CParserGoto)|(1<<CParserIf)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserReturn)|(1<<CParserShort)|(1<<CParserSizeof)|(1<<CParserStruct)|(1<<CParserSwitch)|(1<<CParserVoid)|(1<<CParserWhile)|(1<<CParserLeftParen))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(CParserLeftBrace-34))|(1<<(CParserPlus-34))|(1<<(CParserPlusPlus-34))|(1<<(CParserMinus-34))|(1<<(CParserMinusMinus-34))|(1<<(CParserStar-34))|(1<<(CParserAnd-34))|(1<<(CParserNot-34))|(1<<(CParserTilde-34)))) != 0) || (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(CParserIdentifier-76))|(1<<(CParserIntegerConstant-76))|(1<<(CParserFloatingConstant-76))|(1<<(CParserCharacterConstant-76))|(1<<(CParserStringLiteral-76))|(1<<(CParserBlockComment-76))|(1<<(CParserLineComment-76)))) != 0) {
		{
			p.SetState(371)
			p.Statement()
		}

		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDefaultStatementContext is an interface to support dynamic dispatch.
type IDefaultStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultStatementContext differentiates from other interfaces.
	IsDefaultStatementContext()
}

type DefaultStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultStatementContext() *DefaultStatementContext {
	var p = new(DefaultStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_defaultStatement
	return p
}

func (*DefaultStatementContext) IsDefaultStatementContext() {}

func NewDefaultStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultStatementContext {
	var p = new(DefaultStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_defaultStatement

	return p
}

func (s *DefaultStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultStatementContext) Default() antlr.TerminalNode {
	return s.GetToken(CParserDefault, 0)
}

func (s *DefaultStatementContext) Colon() antlr.TerminalNode {
	return s.GetToken(CParserColon, 0)
}

func (s *DefaultStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *DefaultStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DefaultStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitDefaultStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) DefaultStatement() (localctx IDefaultStatementContext) {
	localctx = NewDefaultStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CParserRULE_defaultStatement)
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
		p.SetState(377)
		p.Match(CParserDefault)
	}
	{
		p.SetState(378)
		p.Match(CParserColon)
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserBreak)|(1<<CParserChar)|(1<<CParserContinue)|(1<<CParserDo)|(1<<CParserDouble)|(1<<CParserEnum)|(1<<CParserFloat)|(1<<CParserFor)|(1<<CParserGoto)|(1<<CParserIf)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserReturn)|(1<<CParserShort)|(1<<CParserSizeof)|(1<<CParserStruct)|(1<<CParserSwitch)|(1<<CParserVoid)|(1<<CParserWhile)|(1<<CParserLeftParen))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(CParserLeftBrace-34))|(1<<(CParserPlus-34))|(1<<(CParserPlusPlus-34))|(1<<(CParserMinus-34))|(1<<(CParserMinusMinus-34))|(1<<(CParserStar-34))|(1<<(CParserAnd-34))|(1<<(CParserNot-34))|(1<<(CParserTilde-34)))) != 0) || (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(CParserIdentifier-76))|(1<<(CParserIntegerConstant-76))|(1<<(CParserFloatingConstant-76))|(1<<(CParserCharacterConstant-76))|(1<<(CParserStringLiteral-76))|(1<<(CParserBlockComment-76))|(1<<(CParserLineComment-76)))) != 0) {
		{
			p.SetState(379)
			p.Statement()
		}

		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 58, CParserRULE_forStatement)
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
		p.SetState(385)
		p.Match(CParserFor)
	}
	{
		p.SetState(386)
		p.Match(CParserLeftParen)
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(387)

			var _x = p.expression(0)

			localctx.(*ForStatementContext).init = _x
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(388)
			p.VariableDeclaration()
		}

	}
	{
		p.SetState(391)
		p.Match(CParserSemi)
	}
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CParserSizeof-21))|(1<<(CParserLeftParen-21))|(1<<(CParserPlus-21))|(1<<(CParserPlusPlus-21))|(1<<(CParserMinus-21))|(1<<(CParserMinusMinus-21))|(1<<(CParserStar-21))|(1<<(CParserAnd-21)))) != 0) || (((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(CParserNot-54))|(1<<(CParserTilde-54))|(1<<(CParserIdentifier-54))|(1<<(CParserIntegerConstant-54))|(1<<(CParserFloatingConstant-54))|(1<<(CParserCharacterConstant-54))|(1<<(CParserStringLiteral-54)))) != 0) {
		{
			p.SetState(392)

			var _x = p.expression(0)

			localctx.(*ForStatementContext).condition = _x
		}

	}
	{
		p.SetState(395)
		p.Match(CParserSemi)
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CParserSizeof-21))|(1<<(CParserLeftParen-21))|(1<<(CParserPlus-21))|(1<<(CParserPlusPlus-21))|(1<<(CParserMinus-21))|(1<<(CParserMinusMinus-21))|(1<<(CParserStar-21))|(1<<(CParserAnd-21)))) != 0) || (((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(CParserNot-54))|(1<<(CParserTilde-54))|(1<<(CParserIdentifier-54))|(1<<(CParserIntegerConstant-54))|(1<<(CParserFloatingConstant-54))|(1<<(CParserCharacterConstant-54))|(1<<(CParserStringLiteral-54)))) != 0) {
		{
			p.SetState(396)

			var _x = p.expression(0)

			localctx.(*ForStatementContext).post = _x
		}

	}
	{
		p.SetState(399)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(400)
		p.Statement()
	}

	return localctx
}

// IDoWhileStatementContext is an interface to support dynamic dispatch.
type IDoWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoWhileStatementContext differentiates from other interfaces.
	IsDoWhileStatementContext()
}

type DoWhileStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoWhileStatementContext() *DoWhileStatementContext {
	var p = new(DoWhileStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_doWhileStatement
	return p
}

func (*DoWhileStatementContext) IsDoWhileStatementContext() {}

func NewDoWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoWhileStatementContext {
	var p = new(DoWhileStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_doWhileStatement

	return p
}

func (s *DoWhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DoWhileStatementContext) Do() antlr.TerminalNode {
	return s.GetToken(CParserDo, 0)
}

func (s *DoWhileStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DoWhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(CParserWhile, 0)
}

func (s *DoWhileStatementContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(CParserLeftParen, 0)
}

func (s *DoWhileStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DoWhileStatementContext) RightParen() antlr.TerminalNode {
	return s.GetToken(CParserRightParen, 0)
}

func (s *DoWhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoWhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoWhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitDoWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) DoWhileStatement() (localctx IDoWhileStatementContext) {
	localctx = NewDoWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CParserRULE_doWhileStatement)

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
		p.SetState(402)
		p.Match(CParserDo)
	}
	{
		p.SetState(403)
		p.Statement()
	}
	{
		p.SetState(404)
		p.Match(CParserWhile)
	}
	{
		p.SetState(405)
		p.Match(CParserLeftParen)
	}
	{
		p.SetState(406)
		p.expression(0)
	}
	{
		p.SetState(407)
		p.Match(CParserRightParen)
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
	p.EnterRule(localctx, 62, CParserRULE_whileStatement)

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
		p.SetState(409)
		p.Match(CParserWhile)
	}
	{
		p.SetState(410)
		p.Match(CParserLeftParen)
	}
	{
		p.SetState(411)
		p.expression(0)
	}
	{
		p.SetState(412)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(413)
		p.Statement()
	}

	return localctx
}

// IGotoStatementContext is an interface to support dynamic dispatch.
type IGotoStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGotoStatementContext differentiates from other interfaces.
	IsGotoStatementContext()
}

type GotoStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGotoStatementContext() *GotoStatementContext {
	var p = new(GotoStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_gotoStatement
	return p
}

func (*GotoStatementContext) IsGotoStatementContext() {}

func NewGotoStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GotoStatementContext {
	var p = new(GotoStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_gotoStatement

	return p
}

func (s *GotoStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *GotoStatementContext) Goto() antlr.TerminalNode {
	return s.GetToken(CParserGoto, 0)
}

func (s *GotoStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *GotoStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GotoStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GotoStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitGotoStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) GotoStatement() (localctx IGotoStatementContext) {
	localctx = NewGotoStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CParserRULE_gotoStatement)

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
		p.SetState(415)
		p.Match(CParserGoto)
	}
	{
		p.SetState(416)
		p.Match(CParserIdentifier)
	}

	return localctx
}

// ILabelStatementContext is an interface to support dynamic dispatch.
type ILabelStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelStatementContext differentiates from other interfaces.
	IsLabelStatementContext()
}

type LabelStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelStatementContext() *LabelStatementContext {
	var p = new(LabelStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_labelStatement
	return p
}

func (*LabelStatementContext) IsLabelStatementContext() {}

func NewLabelStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelStatementContext {
	var p = new(LabelStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_labelStatement

	return p
}

func (s *LabelStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *LabelStatementContext) Colon() antlr.TerminalNode {
	return s.GetToken(CParserColon, 0)
}

func (s *LabelStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitLabelStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) LabelStatement() (localctx ILabelStatementContext) {
	localctx = NewLabelStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CParserRULE_labelStatement)

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
		p.SetState(418)
		p.Match(CParserIdentifier)
	}
	{
		p.SetState(419)
		p.Match(CParserColon)
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
	p.EnterRule(localctx, 68, CParserRULE_includePreprocessor)

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
		p.SetState(421)
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

	case 17:
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
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
