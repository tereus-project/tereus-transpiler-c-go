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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 93, 471,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 3, 2, 7, 2, 82, 10, 2, 12, 2, 14, 2, 85, 11, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	5, 3, 100, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 106, 10, 4, 3, 4, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 115, 10, 5, 3, 6, 3, 6, 5, 6, 119,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 126, 10, 7, 12, 7, 14, 7, 129,
	11, 7, 3, 8, 5, 8, 132, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 139,
	10, 8, 5, 8, 141, 10, 8, 3, 9, 3, 9, 3, 9, 5, 9, 146, 10, 9, 3, 9, 5, 9,
	149, 10, 9, 3, 10, 3, 10, 7, 10, 153, 10, 10, 12, 10, 14, 10, 156, 11,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 5, 11, 162, 10, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 5, 12, 168, 10, 12, 3, 12, 3, 12, 3, 12, 5, 12, 173, 10, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 5, 13, 180, 10, 13, 3, 13, 3, 13, 5, 13,
	184, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 7,
	16, 194, 10, 16, 12, 16, 14, 16, 197, 11, 16, 3, 16, 3, 16, 7, 16, 201,
	10, 16, 12, 16, 14, 16, 204, 11, 16, 3, 16, 3, 16, 3, 16, 5, 16, 209, 10,
	16, 5, 16, 211, 10, 16, 3, 16, 3, 16, 5, 16, 215, 10, 16, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 225, 10, 18, 12, 18, 14,
	18, 228, 11, 18, 3, 18, 5, 18, 231, 10, 18, 5, 18, 233, 10, 18, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 241, 10, 19, 12, 19, 14, 19,
	244, 11, 19, 3, 19, 5, 19, 247, 10, 19, 5, 19, 249, 10, 19, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 5, 21, 267, 10, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 288, 10, 21, 5, 21, 290, 10, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 328, 10, 21, 3, 21, 3, 21, 3, 21,
	7, 21, 333, 10, 21, 12, 21, 14, 21, 336, 11, 21, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 347, 10, 24, 3, 25, 3, 25,
	3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 5, 28, 358, 10, 28, 3,
	29, 3, 29, 7, 29, 362, 10, 29, 12, 29, 14, 29, 365, 11, 29, 3, 29, 3, 29,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 378,
	10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	5, 30, 389, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5,
	31, 398, 10, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32, 406,
	10, 32, 12, 32, 14, 32, 409, 11, 32, 3, 32, 5, 32, 412, 10, 32, 3, 32,
	3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 420, 10, 33, 12, 33, 14, 33,
	423, 11, 33, 3, 34, 3, 34, 3, 34, 7, 34, 428, 10, 34, 12, 34, 14, 34, 431,
	11, 34, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 437, 10, 35, 3, 35, 3, 35, 5,
	35, 441, 10, 35, 3, 35, 3, 35, 5, 35, 445, 10, 35, 3, 35, 3, 35, 3, 35,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40,
	3, 40, 2, 4, 12, 40, 41, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
	64, 66, 68, 70, 72, 74, 76, 78, 2, 10, 4, 2, 22, 22, 29, 29, 6, 2, 5, 5,
	18, 19, 21, 21, 78, 78, 4, 2, 79, 80, 82, 82, 3, 2, 62, 72, 3, 2, 53, 54,
	8, 2, 38, 38, 40, 40, 42, 44, 46, 46, 48, 52, 55, 55, 4, 2, 45, 45, 47,
	47, 5, 2, 44, 48, 51, 51, 56, 57, 2, 518, 2, 83, 3, 2, 2, 2, 4, 99, 3,
	2, 2, 2, 6, 101, 3, 2, 2, 2, 8, 110, 3, 2, 2, 2, 10, 116, 3, 2, 2, 2, 12,
	120, 3, 2, 2, 2, 14, 140, 3, 2, 2, 2, 16, 142, 3, 2, 2, 2, 18, 150, 3,
	2, 2, 2, 20, 159, 3, 2, 2, 2, 22, 165, 3, 2, 2, 2, 24, 176, 3, 2, 2, 2,
	26, 185, 3, 2, 2, 2, 28, 189, 3, 2, 2, 2, 30, 195, 3, 2, 2, 2, 32, 216,
	3, 2, 2, 2, 34, 220, 3, 2, 2, 2, 36, 236, 3, 2, 2, 2, 38, 252, 3, 2, 2,
	2, 40, 289, 3, 2, 2, 2, 42, 337, 3, 2, 2, 2, 44, 339, 3, 2, 2, 2, 46, 346,
	3, 2, 2, 2, 48, 348, 3, 2, 2, 2, 50, 350, 3, 2, 2, 2, 52, 352, 3, 2, 2,
	2, 54, 354, 3, 2, 2, 2, 56, 359, 3, 2, 2, 2, 58, 388, 3, 2, 2, 2, 60, 390,
	3, 2, 2, 2, 62, 399, 3, 2, 2, 2, 64, 415, 3, 2, 2, 2, 66, 424, 3, 2, 2,
	2, 68, 432, 3, 2, 2, 2, 70, 449, 3, 2, 2, 2, 72, 456, 3, 2, 2, 2, 74, 462,
	3, 2, 2, 2, 76, 465, 3, 2, 2, 2, 78, 468, 3, 2, 2, 2, 80, 82, 5, 4, 3,
	2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84,
	3, 2, 2, 2, 84, 3, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 100, 5, 6, 4, 2,
	87, 88, 5, 16, 9, 2, 88, 89, 7, 60, 2, 2, 89, 100, 3, 2, 2, 2, 90, 91,
	5, 22, 12, 2, 91, 92, 7, 60, 2, 2, 92, 100, 3, 2, 2, 2, 93, 100, 5, 78,
	40, 2, 94, 95, 5, 26, 14, 2, 95, 96, 7, 60, 2, 2, 96, 100, 3, 2, 2, 2,
	97, 100, 7, 92, 2, 2, 98, 100, 7, 93, 2, 2, 99, 86, 3, 2, 2, 2, 99, 87,
	3, 2, 2, 2, 99, 90, 3, 2, 2, 2, 99, 93, 3, 2, 2, 2, 99, 94, 3, 2, 2, 2,
	99, 97, 3, 2, 2, 2, 99, 98, 3, 2, 2, 2, 100, 5, 3, 2, 2, 2, 101, 102, 5,
	12, 7, 2, 102, 103, 7, 78, 2, 2, 103, 105, 7, 32, 2, 2, 104, 106, 5, 8,
	5, 2, 105, 104, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2,
	107, 108, 7, 33, 2, 2, 108, 109, 5, 56, 29, 2, 109, 7, 3, 2, 2, 2, 110,
	111, 5, 12, 7, 2, 111, 114, 7, 78, 2, 2, 112, 113, 7, 61, 2, 2, 113, 115,
	5, 8, 5, 2, 114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 9, 3, 2, 2,
	2, 116, 118, 7, 20, 2, 2, 117, 119, 5, 40, 21, 2, 118, 117, 3, 2, 2, 2,
	118, 119, 3, 2, 2, 2, 119, 11, 3, 2, 2, 2, 120, 121, 8, 7, 1, 2, 121, 122,
	5, 14, 8, 2, 122, 127, 3, 2, 2, 2, 123, 124, 12, 3, 2, 2, 124, 126, 7,
	48, 2, 2, 125, 123, 3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2,
	2, 127, 128, 3, 2, 2, 2, 128, 13, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 130,
	132, 9, 2, 2, 2, 131, 130, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 133,
	3, 2, 2, 2, 133, 141, 9, 3, 2, 2, 134, 139, 7, 30, 2, 2, 135, 139, 7, 14,
	2, 2, 136, 139, 7, 10, 2, 2, 137, 139, 5, 16, 9, 2, 138, 134, 3, 2, 2,
	2, 138, 135, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 138, 137, 3, 2, 2, 2, 139,
	141, 3, 2, 2, 2, 140, 131, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 141, 15, 3,
	2, 2, 2, 142, 148, 7, 25, 2, 2, 143, 145, 7, 78, 2, 2, 144, 146, 5, 18,
	10, 2, 145, 144, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 149, 3, 2, 2, 2,
	147, 149, 5, 18, 10, 2, 148, 143, 3, 2, 2, 2, 148, 147, 3, 2, 2, 2, 149,
	17, 3, 2, 2, 2, 150, 154, 7, 36, 2, 2, 151, 153, 5, 20, 11, 2, 152, 151,
	3, 2, 2, 2, 153, 156, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 154, 155, 3, 2,
	2, 2, 155, 157, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 157, 158, 7, 37, 2, 2,
	158, 19, 3, 2, 2, 2, 159, 161, 5, 12, 7, 2, 160, 162, 7, 78, 2, 2, 161,
	160, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 164,
	7, 60, 2, 2, 164, 21, 3, 2, 2, 2, 165, 167, 7, 12, 2, 2, 166, 168, 7, 78,
	2, 2, 167, 166, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2,
	169, 170, 7, 36, 2, 2, 170, 172, 5, 24, 13, 2, 171, 173, 7, 61, 2, 2, 172,
	171, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 175,
	7, 37, 2, 2, 175, 23, 3, 2, 2, 2, 176, 179, 7, 78, 2, 2, 177, 178, 7, 62,
	2, 2, 178, 180, 5, 40, 21, 2, 179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2,
	2, 180, 183, 3, 2, 2, 2, 181, 182, 7, 61, 2, 2, 182, 184, 5, 24, 13, 2,
	183, 181, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 25, 3, 2, 2, 2, 185, 186,
	7, 27, 2, 2, 186, 187, 5, 12, 7, 2, 187, 188, 7, 78, 2, 2, 188, 27, 3,
	2, 2, 2, 189, 190, 5, 14, 8, 2, 190, 191, 5, 30, 16, 2, 191, 29, 3, 2,
	2, 2, 192, 194, 7, 48, 2, 2, 193, 192, 3, 2, 2, 2, 194, 197, 3, 2, 2, 2,
	195, 193, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 198, 3, 2, 2, 2, 197,
	195, 3, 2, 2, 2, 198, 202, 7, 78, 2, 2, 199, 201, 5, 32, 17, 2, 200, 199,
	3, 2, 2, 2, 201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 202, 203, 3, 2,
	2, 2, 203, 210, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 205, 208, 7, 62, 2, 2,
	206, 209, 5, 40, 21, 2, 207, 209, 5, 34, 18, 2, 208, 206, 3, 2, 2, 2, 208,
	207, 3, 2, 2, 2, 209, 211, 3, 2, 2, 2, 210, 205, 3, 2, 2, 2, 210, 211,
	3, 2, 2, 2, 211, 214, 3, 2, 2, 2, 212, 213, 7, 61, 2, 2, 213, 215, 5, 30,
	16, 2, 214, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 31, 3, 2, 2, 2,
	216, 217, 7, 34, 2, 2, 217, 218, 5, 40, 21, 2, 218, 219, 7, 35, 2, 2, 219,
	33, 3, 2, 2, 2, 220, 232, 7, 36, 2, 2, 221, 226, 5, 40, 21, 2, 222, 223,
	7, 61, 2, 2, 223, 225, 5, 40, 21, 2, 224, 222, 3, 2, 2, 2, 225, 228, 3,
	2, 2, 2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 230, 3, 2, 2,
	2, 228, 226, 3, 2, 2, 2, 229, 231, 7, 61, 2, 2, 230, 229, 3, 2, 2, 2, 230,
	231, 3, 2, 2, 2, 231, 233, 3, 2, 2, 2, 232, 221, 3, 2, 2, 2, 232, 233,
	3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 235, 7, 37, 2, 2, 235, 35, 3, 2,
	2, 2, 236, 248, 7, 36, 2, 2, 237, 242, 5, 38, 20, 2, 238, 239, 7, 61, 2,
	2, 239, 241, 5, 38, 20, 2, 240, 238, 3, 2, 2, 2, 241, 244, 3, 2, 2, 2,
	242, 240, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 246, 3, 2, 2, 2, 244,
	242, 3, 2, 2, 2, 245, 247, 7, 61, 2, 2, 246, 245, 3, 2, 2, 2, 246, 247,
	3, 2, 2, 2, 247, 249, 3, 2, 2, 2, 248, 237, 3, 2, 2, 2, 248, 249, 3, 2,
	2, 2, 249, 250, 3, 2, 2, 2, 250, 251, 7, 37, 2, 2, 251, 37, 3, 2, 2, 2,
	252, 253, 7, 76, 2, 2, 253, 254, 7, 78, 2, 2, 254, 255, 7, 62, 2, 2, 255,
	256, 5, 40, 21, 2, 256, 39, 3, 2, 2, 2, 257, 258, 8, 21, 1, 2, 258, 290,
	7, 78, 2, 2, 259, 290, 9, 4, 2, 2, 260, 290, 7, 83, 2, 2, 261, 262, 7,
	32, 2, 2, 262, 263, 5, 12, 7, 2, 263, 266, 7, 33, 2, 2, 264, 267, 5, 34,
	18, 2, 265, 267, 5, 36, 19, 2, 266, 264, 3, 2, 2, 2, 266, 265, 3, 2, 2,
	2, 267, 290, 3, 2, 2, 2, 268, 269, 7, 32, 2, 2, 269, 270, 5, 12, 7, 2,
	270, 271, 7, 33, 2, 2, 271, 272, 5, 40, 21, 12, 272, 290, 3, 2, 2, 2, 273,
	274, 7, 32, 2, 2, 274, 275, 5, 40, 21, 2, 275, 276, 7, 33, 2, 2, 276, 290,
	3, 2, 2, 2, 277, 278, 5, 52, 27, 2, 278, 279, 5, 40, 21, 9, 279, 290, 3,
	2, 2, 2, 280, 287, 7, 23, 2, 2, 281, 282, 7, 32, 2, 2, 282, 283, 5, 12,
	7, 2, 283, 284, 7, 33, 2, 2, 284, 288, 3, 2, 2, 2, 285, 288, 5, 12, 7,
	2, 286, 288, 5, 40, 21, 2, 287, 281, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2,
	287, 286, 3, 2, 2, 2, 288, 290, 3, 2, 2, 2, 289, 257, 3, 2, 2, 2, 289,
	259, 3, 2, 2, 2, 289, 260, 3, 2, 2, 2, 289, 261, 3, 2, 2, 2, 289, 268,
	3, 2, 2, 2, 289, 273, 3, 2, 2, 2, 289, 277, 3, 2, 2, 2, 289, 280, 3, 2,
	2, 2, 290, 334, 3, 2, 2, 2, 291, 292, 12, 8, 2, 2, 292, 293, 5, 42, 22,
	2, 293, 294, 5, 40, 21, 9, 294, 333, 3, 2, 2, 2, 295, 296, 12, 7, 2, 2,
	296, 297, 5, 48, 25, 2, 297, 298, 5, 40, 21, 8, 298, 333, 3, 2, 2, 2, 299,
	300, 12, 6, 2, 2, 300, 301, 5, 46, 24, 2, 301, 302, 5, 40, 21, 7, 302,
	333, 3, 2, 2, 2, 303, 304, 12, 5, 2, 2, 304, 305, 5, 44, 23, 2, 305, 306,
	5, 40, 21, 6, 306, 333, 3, 2, 2, 2, 307, 308, 12, 3, 2, 2, 308, 309, 7,
	58, 2, 2, 309, 310, 5, 40, 21, 2, 310, 311, 7, 59, 2, 2, 311, 312, 5, 40,
	21, 4, 312, 333, 3, 2, 2, 2, 313, 314, 12, 16, 2, 2, 314, 315, 7, 76, 2,
	2, 315, 333, 7, 78, 2, 2, 316, 317, 12, 15, 2, 2, 317, 318, 7, 75, 2, 2,
	318, 333, 7, 78, 2, 2, 319, 320, 12, 14, 2, 2, 320, 321, 7, 34, 2, 2, 321,
	322, 5, 40, 21, 2, 322, 323, 7, 35, 2, 2, 323, 333, 3, 2, 2, 2, 324, 325,
	12, 13, 2, 2, 325, 327, 7, 32, 2, 2, 326, 328, 5, 54, 28, 2, 327, 326,
	3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 333, 7, 33,
	2, 2, 330, 331, 12, 10, 2, 2, 331, 333, 5, 50, 26, 2, 332, 291, 3, 2, 2,
	2, 332, 295, 3, 2, 2, 2, 332, 299, 3, 2, 2, 2, 332, 303, 3, 2, 2, 2, 332,
	307, 3, 2, 2, 2, 332, 313, 3, 2, 2, 2, 332, 316, 3, 2, 2, 2, 332, 319,
	3, 2, 2, 2, 332, 324, 3, 2, 2, 2, 332, 330, 3, 2, 2, 2, 333, 336, 3, 2,
	2, 2, 334, 332, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 41, 3, 2, 2, 2,
	336, 334, 3, 2, 2, 2, 337, 338, 9, 5, 2, 2, 338, 43, 3, 2, 2, 2, 339, 340,
	9, 6, 2, 2, 340, 45, 3, 2, 2, 2, 341, 347, 7, 73, 2, 2, 342, 347, 7, 74,
	2, 2, 343, 347, 3, 2, 2, 2, 344, 347, 7, 39, 2, 2, 345, 347, 7, 41, 2,
	2, 346, 341, 3, 2, 2, 2, 346, 342, 3, 2, 2, 2, 346, 343, 3, 2, 2, 2, 346,
	344, 3, 2, 2, 2, 346, 345, 3, 2, 2, 2, 347, 47, 3, 2, 2, 2, 348, 349, 9,
	7, 2, 2, 349, 49, 3, 2, 2, 2, 350, 351, 9, 8, 2, 2, 351, 51, 3, 2, 2, 2,
	352, 353, 9, 9, 2, 2, 353, 53, 3, 2, 2, 2, 354, 357, 5, 40, 21, 2, 355,
	356, 7, 61, 2, 2, 356, 358, 5, 54, 28, 2, 357, 355, 3, 2, 2, 2, 357, 358,
	3, 2, 2, 2, 358, 55, 3, 2, 2, 2, 359, 363, 7, 36, 2, 2, 360, 362, 5, 58,
	30, 2, 361, 360, 3, 2, 2, 2, 362, 365, 3, 2, 2, 2, 363, 361, 3, 2, 2, 2,
	363, 364, 3, 2, 2, 2, 364, 366, 3, 2, 2, 2, 365, 363, 3, 2, 2, 2, 366,
	367, 7, 37, 2, 2, 367, 57, 3, 2, 2, 2, 368, 378, 5, 28, 15, 2, 369, 378,
	5, 40, 21, 2, 370, 378, 5, 10, 6, 2, 371, 378, 7, 3, 2, 2, 372, 378, 7,
	7, 2, 2, 373, 378, 5, 16, 9, 2, 374, 378, 5, 22, 12, 2, 375, 378, 5, 74,
	38, 2, 376, 378, 5, 70, 36, 2, 377, 368, 3, 2, 2, 2, 377, 369, 3, 2, 2,
	2, 377, 370, 3, 2, 2, 2, 377, 371, 3, 2, 2, 2, 377, 372, 3, 2, 2, 2, 377,
	373, 3, 2, 2, 2, 377, 374, 3, 2, 2, 2, 377, 375, 3, 2, 2, 2, 377, 376,
	3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379, 389, 7, 60, 2, 2, 380, 389, 5, 60,
	31, 2, 381, 389, 5, 62, 32, 2, 382, 389, 5, 68, 35, 2, 383, 389, 5, 72,
	37, 2, 384, 389, 5, 56, 29, 2, 385, 389, 5, 76, 39, 2, 386, 389, 7, 92,
	2, 2, 387, 389, 7, 93, 2, 2, 388, 377, 3, 2, 2, 2, 388, 380, 3, 2, 2, 2,
	388, 381, 3, 2, 2, 2, 388, 382, 3, 2, 2, 2, 388, 383, 3, 2, 2, 2, 388,
	384, 3, 2, 2, 2, 388, 385, 3, 2, 2, 2, 388, 386, 3, 2, 2, 2, 388, 387,
	3, 2, 2, 2, 389, 59, 3, 2, 2, 2, 390, 391, 7, 17, 2, 2, 391, 392, 7, 32,
	2, 2, 392, 393, 5, 40, 21, 2, 393, 394, 7, 33, 2, 2, 394, 397, 5, 58, 30,
	2, 395, 396, 7, 11, 2, 2, 396, 398, 5, 58, 30, 2, 397, 395, 3, 2, 2, 2,
	397, 398, 3, 2, 2, 2, 398, 61, 3, 2, 2, 2, 399, 400, 7, 26, 2, 2, 400,
	401, 7, 32, 2, 2, 401, 402, 5, 40, 21, 2, 402, 403, 7, 33, 2, 2, 403, 407,
	7, 36, 2, 2, 404, 406, 5, 64, 33, 2, 405, 404, 3, 2, 2, 2, 406, 409, 3,
	2, 2, 2, 407, 405, 3, 2, 2, 2, 407, 408, 3, 2, 2, 2, 408, 411, 3, 2, 2,
	2, 409, 407, 3, 2, 2, 2, 410, 412, 5, 66, 34, 2, 411, 410, 3, 2, 2, 2,
	411, 412, 3, 2, 2, 2, 412, 413, 3, 2, 2, 2, 413, 414, 7, 37, 2, 2, 414,
	63, 3, 2, 2, 2, 415, 416, 7, 4, 2, 2, 416, 417, 5, 40, 21, 2, 417, 421,
	7, 59, 2, 2, 418, 420, 5, 58, 30, 2, 419, 418, 3, 2, 2, 2, 420, 423, 3,
	2, 2, 2, 421, 419, 3, 2, 2, 2, 421, 422, 3, 2, 2, 2, 422, 65, 3, 2, 2,
	2, 423, 421, 3, 2, 2, 2, 424, 425, 7, 8, 2, 2, 425, 429, 7, 59, 2, 2, 426,
	428, 5, 58, 30, 2, 427, 426, 3, 2, 2, 2, 428, 431, 3, 2, 2, 2, 429, 427,
	3, 2, 2, 2, 429, 430, 3, 2, 2, 2, 430, 67, 3, 2, 2, 2, 431, 429, 3, 2,
	2, 2, 432, 433, 7, 15, 2, 2, 433, 436, 7, 32, 2, 2, 434, 437, 5, 40, 21,
	2, 435, 437, 5, 28, 15, 2, 436, 434, 3, 2, 2, 2, 436, 435, 3, 2, 2, 2,
	436, 437, 3, 2, 2, 2, 437, 438, 3, 2, 2, 2, 438, 440, 7, 60, 2, 2, 439,
	441, 5, 40, 21, 2, 440, 439, 3, 2, 2, 2, 440, 441, 3, 2, 2, 2, 441, 442,
	3, 2, 2, 2, 442, 444, 7, 60, 2, 2, 443, 445, 5, 40, 21, 2, 444, 443, 3,
	2, 2, 2, 444, 445, 3, 2, 2, 2, 445, 446, 3, 2, 2, 2, 446, 447, 7, 33, 2,
	2, 447, 448, 5, 58, 30, 2, 448, 69, 3, 2, 2, 2, 449, 450, 7, 9, 2, 2, 450,
	451, 5, 58, 30, 2, 451, 452, 7, 31, 2, 2, 452, 453, 7, 32, 2, 2, 453, 454,
	5, 40, 21, 2, 454, 455, 7, 33, 2, 2, 455, 71, 3, 2, 2, 2, 456, 457, 7,
	31, 2, 2, 457, 458, 7, 32, 2, 2, 458, 459, 5, 40, 21, 2, 459, 460, 7, 33,
	2, 2, 460, 461, 5, 58, 30, 2, 461, 73, 3, 2, 2, 2, 462, 463, 7, 16, 2,
	2, 463, 464, 7, 78, 2, 2, 464, 75, 3, 2, 2, 2, 465, 466, 7, 78, 2, 2, 466,
	467, 7, 59, 2, 2, 467, 77, 3, 2, 2, 2, 468, 469, 7, 85, 2, 2, 469, 79,
	3, 2, 2, 2, 49, 83, 99, 105, 114, 118, 127, 131, 138, 140, 145, 148, 154,
	161, 167, 172, 179, 183, 195, 202, 208, 210, 214, 226, 230, 232, 242, 246,
	248, 266, 287, 289, 327, 332, 334, 346, 357, 363, 377, 388, 397, 407, 411,
	421, 429, 436, 440, 444,
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
	"functionReturn", "typeSpecifier", "typeSpecifierNoPointer", "structDeclaration",
	"structDeclarationBody", "structProperty", "enumDeclaration", "enumProperties",
	"typedefDeclaration", "variableDeclaration", "variableDeclarationList",
	"sizedArrayModifier", "listInitialization", "namedListInitialization",
	"namedListInitializationItem", "expression", "assignementOperator", "conditionalBinaryOperator",
	"equalityOperator", "binaryOperator", "unaryOperatorPost", "unaryOperatorPre",
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
	CParserRULE_typeSpecifierNoPointer      = 6
	CParserRULE_structDeclaration           = 7
	CParserRULE_structDeclarationBody       = 8
	CParserRULE_structProperty              = 9
	CParserRULE_enumDeclaration             = 10
	CParserRULE_enumProperties              = 11
	CParserRULE_typedefDeclaration          = 12
	CParserRULE_variableDeclaration         = 13
	CParserRULE_variableDeclarationList     = 14
	CParserRULE_sizedArrayModifier          = 15
	CParserRULE_listInitialization          = 16
	CParserRULE_namedListInitialization     = 17
	CParserRULE_namedListInitializationItem = 18
	CParserRULE_expression                  = 19
	CParserRULE_assignementOperator         = 20
	CParserRULE_conditionalBinaryOperator   = 21
	CParserRULE_equalityOperator            = 22
	CParserRULE_binaryOperator              = 23
	CParserRULE_unaryOperatorPost           = 24
	CParserRULE_unaryOperatorPre            = 25
	CParserRULE_functionCallArguments       = 26
	CParserRULE_block                       = 27
	CParserRULE_statement                   = 28
	CParserRULE_ifStatement                 = 29
	CParserRULE_switchStatement             = 30
	CParserRULE_caseStatement               = 31
	CParserRULE_defaultStatement            = 32
	CParserRULE_forStatement                = 33
	CParserRULE_doWhileStatement            = 34
	CParserRULE_whileStatement              = 35
	CParserRULE_gotoStatement               = 36
	CParserRULE_labelStatement              = 37
	CParserRULE_includePreprocessor         = 38
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
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserEnum)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserSigned)|(1<<CParserStruct)|(1<<CParserTypedef)|(1<<CParserUnsigned)|(1<<CParserVoid))) != 0) || (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(CParserIdentifier-76))|(1<<(CParserIncludeDirective-76))|(1<<(CParserBlockComment-76))|(1<<(CParserLineComment-76)))) != 0) {
		{
			p.SetState(78)
			p.Declaration()
		}

		p.SetState(83)
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

	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.FunctionDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.StructDeclaration()
		}
		{
			p.SetState(86)
			p.Match(CParserSemi)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.EnumDeclaration()
		}
		{
			p.SetState(89)
			p.Match(CParserSemi)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.IncludePreprocessor()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(92)
			p.TypedefDeclaration()
		}
		{
			p.SetState(93)
			p.Match(CParserSemi)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(95)
			p.Match(CParserBlockComment)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(96)
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
		p.SetState(99)
		p.typeSpecifier(0)
	}
	{
		p.SetState(100)
		p.Match(CParserIdentifier)
	}
	{
		p.SetState(101)
		p.Match(CParserLeftParen)
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserSigned)|(1<<CParserStruct)|(1<<CParserUnsigned)|(1<<CParserVoid))) != 0) || _la == CParserIdentifier {
		{
			p.SetState(102)
			p.FunctionArguments()
		}

	}
	{
		p.SetState(105)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(106)
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
		p.SetState(108)
		p.typeSpecifier(0)
	}
	{
		p.SetState(109)
		p.Match(CParserIdentifier)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(110)
			p.Match(CParserComma)
		}
		{
			p.SetState(111)
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
		p.SetState(114)
		p.Match(CParserReturn)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CParserSizeof-21))|(1<<(CParserLeftParen-21))|(1<<(CParserPlus-21))|(1<<(CParserPlusPlus-21))|(1<<(CParserMinus-21))|(1<<(CParserMinusMinus-21))|(1<<(CParserStar-21))|(1<<(CParserAnd-21)))) != 0) || (((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(CParserNot-54))|(1<<(CParserTilde-54))|(1<<(CParserIdentifier-54))|(1<<(CParserIntegerConstant-54))|(1<<(CParserFloatingConstant-54))|(1<<(CParserCharacterConstant-54))|(1<<(CParserStringLiteral-54)))) != 0) {
		{
			p.SetState(115)
			p.expression(0)
		}

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

func (s *TypeSpecifierContext) CopyFrom(ctx *TypeSpecifierContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeSpecifierPointerContext struct {
	*TypeSpecifierContext
}

func NewTypeSpecifierPointerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecifierPointerContext {
	var p = new(TypeSpecifierPointerContext)

	p.TypeSpecifierContext = NewEmptyTypeSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierContext))

	return p
}

func (s *TypeSpecifierPointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierPointerContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *TypeSpecifierPointerContext) Star() antlr.TerminalNode {
	return s.GetToken(CParserStar, 0)
}

func (s *TypeSpecifierPointerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitTypeSpecifierPointer(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeSpecifierClassicContext struct {
	*TypeSpecifierContext
}

func NewTypeSpecifierClassicContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecifierClassicContext {
	var p = new(TypeSpecifierClassicContext)

	p.TypeSpecifierContext = NewEmptyTypeSpecifierContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierContext))

	return p
}

func (s *TypeSpecifierClassicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierClassicContext) TypeSpecifierNoPointer() ITypeSpecifierNoPointerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierNoPointerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierNoPointerContext)
}

func (s *TypeSpecifierClassicContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitTypeSpecifierClassic(s)

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
	localctx = NewTypeSpecifierClassicContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(119)
		p.TypeSpecifierNoPointer()
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
			localctx = NewTypeSpecifierPointerContext(p, NewTypeSpecifierContext(p, _parentctx, _parentState))
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

// ITypeSpecifierNoPointerContext is an interface to support dynamic dispatch.
type ITypeSpecifierNoPointerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecifierNoPointerContext differentiates from other interfaces.
	IsTypeSpecifierNoPointerContext()
}

type TypeSpecifierNoPointerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierNoPointerContext() *TypeSpecifierNoPointerContext {
	var p = new(TypeSpecifierNoPointerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_typeSpecifierNoPointer
	return p
}

func (*TypeSpecifierNoPointerContext) IsTypeSpecifierNoPointerContext() {}

func NewTypeSpecifierNoPointerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierNoPointerContext {
	var p = new(TypeSpecifierNoPointerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_typeSpecifierNoPointer

	return p
}

func (s *TypeSpecifierNoPointerContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierNoPointerContext) CopyFrom(ctx *TypeSpecifierNoPointerContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeSpecifierNoPointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierNoPointerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeSpecifierWithModifierContext struct {
	*TypeSpecifierNoPointerContext
}

func NewTypeSpecifierWithModifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecifierWithModifierContext {
	var p = new(TypeSpecifierWithModifierContext)

	p.TypeSpecifierNoPointerContext = NewEmptyTypeSpecifierNoPointerContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierNoPointerContext))

	return p
}

func (s *TypeSpecifierWithModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierWithModifierContext) Int() antlr.TerminalNode {
	return s.GetToken(CParserInt, 0)
}

func (s *TypeSpecifierWithModifierContext) Short() antlr.TerminalNode {
	return s.GetToken(CParserShort, 0)
}

func (s *TypeSpecifierWithModifierContext) Long() antlr.TerminalNode {
	return s.GetToken(CParserLong, 0)
}

func (s *TypeSpecifierWithModifierContext) Char() antlr.TerminalNode {
	return s.GetToken(CParserChar, 0)
}

func (s *TypeSpecifierWithModifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CParserIdentifier, 0)
}

func (s *TypeSpecifierWithModifierContext) Signed() antlr.TerminalNode {
	return s.GetToken(CParserSigned, 0)
}

func (s *TypeSpecifierWithModifierContext) Unsigned() antlr.TerminalNode {
	return s.GetToken(CParserUnsigned, 0)
}

func (s *TypeSpecifierWithModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitTypeSpecifierWithModifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeSpecifierGenericContext struct {
	*TypeSpecifierNoPointerContext
}

func NewTypeSpecifierGenericContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSpecifierGenericContext {
	var p = new(TypeSpecifierGenericContext)

	p.TypeSpecifierNoPointerContext = NewEmptyTypeSpecifierNoPointerContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecifierNoPointerContext))

	return p
}

func (s *TypeSpecifierGenericContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierGenericContext) Void() antlr.TerminalNode {
	return s.GetToken(CParserVoid, 0)
}

func (s *TypeSpecifierGenericContext) Float() antlr.TerminalNode {
	return s.GetToken(CParserFloat, 0)
}

func (s *TypeSpecifierGenericContext) Double() antlr.TerminalNode {
	return s.GetToken(CParserDouble, 0)
}

func (s *TypeSpecifierGenericContext) StructDeclaration() IStructDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructDeclarationContext)
}

func (s *TypeSpecifierGenericContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitTypeSpecifierGeneric(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) TypeSpecifierNoPointer() (localctx ITypeSpecifierNoPointerContext) {
	localctx = NewTypeSpecifierNoPointerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CParserRULE_typeSpecifierNoPointer)
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

	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserChar, CParserInt, CParserLong, CParserShort, CParserSigned, CParserUnsigned, CParserIdentifier:
		localctx = NewTypeSpecifierWithModifierContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CParserSigned || _la == CParserUnsigned {
			{
				p.SetState(128)
				_la = p.GetTokenStream().LA(1)

				if !(_la == CParserSigned || _la == CParserUnsigned) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(131)
			_la = p.GetTokenStream().LA(1)

			if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort))) != 0) || _la == CParserIdentifier) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case CParserDouble, CParserFloat, CParserStruct, CParserVoid:
		localctx = NewTypeSpecifierGenericContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(136)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CParserVoid:
			{
				p.SetState(132)
				p.Match(CParserVoid)
			}

		case CParserFloat:
			{
				p.SetState(133)
				p.Match(CParserFloat)
			}

		case CParserDouble:
			{
				p.SetState(134)
				p.Match(CParserDouble)
			}

		case CParserStruct:
			{
				p.SetState(135)
				p.StructDeclaration()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 14, CParserRULE_structDeclaration)

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
		p.SetState(140)
		p.Match(CParserStruct)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserIdentifier:
		{
			p.SetState(141)
			p.Match(CParserIdentifier)
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(142)
				p.StructDeclarationBody()
			}

		}

	case CParserLeftBrace:
		{
			p.SetState(145)
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
	p.EnterRule(localctx, 16, CParserRULE_structDeclarationBody)
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
		p.Match(CParserLeftBrace)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserChar)|(1<<CParserDouble)|(1<<CParserFloat)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserShort)|(1<<CParserSigned)|(1<<CParserStruct)|(1<<CParserUnsigned)|(1<<CParserVoid))) != 0) || _la == CParserIdentifier {
		{
			p.SetState(149)
			p.StructProperty()
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(155)
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
	p.EnterRule(localctx, 18, CParserRULE_structProperty)
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
		p.SetState(157)
		p.typeSpecifier(0)
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserIdentifier {
		{
			p.SetState(158)
			p.Match(CParserIdentifier)
		}

	}
	{
		p.SetState(161)
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
	p.EnterRule(localctx, 20, CParserRULE_enumDeclaration)
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
		p.SetState(163)
		p.Match(CParserEnum)
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserIdentifier {
		{
			p.SetState(164)
			p.Match(CParserIdentifier)
		}

	}
	{
		p.SetState(167)
		p.Match(CParserLeftBrace)
	}
	{
		p.SetState(168)
		p.EnumProperties()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(169)
			p.Match(CParserComma)
		}

	}
	{
		p.SetState(172)
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
	p.EnterRule(localctx, 22, CParserRULE_enumProperties)
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
		p.SetState(174)
		p.Match(CParserIdentifier)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserAssign {
		{
			p.SetState(175)
			p.Match(CParserAssign)
		}
		{
			p.SetState(176)
			p.expression(0)
		}

	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(179)
			p.Match(CParserComma)
		}
		{
			p.SetState(180)
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
	p.EnterRule(localctx, 24, CParserRULE_typedefDeclaration)

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
		p.SetState(183)
		p.Match(CParserTypedef)
	}
	{
		p.SetState(184)
		p.typeSpecifier(0)
	}
	{
		p.SetState(185)
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

func (s *VariableDeclarationContext) TypeSpecifierNoPointer() ITypeSpecifierNoPointerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierNoPointerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierNoPointerContext)
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
	p.EnterRule(localctx, 26, CParserRULE_variableDeclaration)

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
		p.SetState(187)
		p.TypeSpecifierNoPointer()
	}
	{
		p.SetState(188)
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

func (s *VariableDeclarationListContext) AllStar() []antlr.TerminalNode {
	return s.GetTokens(CParserStar)
}

func (s *VariableDeclarationListContext) Star(i int) antlr.TerminalNode {
	return s.GetToken(CParserStar, i)
}

func (s *VariableDeclarationListContext) AllSizedArrayModifier() []ISizedArrayModifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISizedArrayModifierContext)(nil)).Elem())
	var tst = make([]ISizedArrayModifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISizedArrayModifierContext)
		}
	}

	return tst
}

func (s *VariableDeclarationListContext) SizedArrayModifier(i int) ISizedArrayModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISizedArrayModifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISizedArrayModifierContext)
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
	p.EnterRule(localctx, 28, CParserRULE_variableDeclarationList)
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
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserStar {
		{
			p.SetState(190)
			p.Match(CParserStar)
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(196)
		p.Match(CParserIdentifier)
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserLeftBracket {
		{
			p.SetState(197)
			p.SizedArrayModifier()
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserAssign {
		{
			p.SetState(203)
			p.Match(CParserAssign)
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case CParserSizeof, CParserLeftParen, CParserPlus, CParserPlusPlus, CParserMinus, CParserMinusMinus, CParserStar, CParserAnd, CParserNot, CParserTilde, CParserIdentifier, CParserIntegerConstant, CParserFloatingConstant, CParserCharacterConstant, CParserStringLiteral:
			{
				p.SetState(204)
				p.expression(0)
			}

		case CParserLeftBrace:
			{
				p.SetState(205)
				p.ListInitialization()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(210)
			p.Match(CParserComma)
		}
		{
			p.SetState(211)
			p.VariableDeclarationList()
		}

	}

	return localctx
}

// ISizedArrayModifierContext is an interface to support dynamic dispatch.
type ISizedArrayModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSizedArrayModifierContext differentiates from other interfaces.
	IsSizedArrayModifierContext()
}

type SizedArrayModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySizedArrayModifierContext() *SizedArrayModifierContext {
	var p = new(SizedArrayModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_sizedArrayModifier
	return p
}

func (*SizedArrayModifierContext) IsSizedArrayModifierContext() {}

func NewSizedArrayModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SizedArrayModifierContext {
	var p = new(SizedArrayModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_sizedArrayModifier

	return p
}

func (s *SizedArrayModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *SizedArrayModifierContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(CParserLeftBracket, 0)
}

func (s *SizedArrayModifierContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SizedArrayModifierContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(CParserRightBracket, 0)
}

func (s *SizedArrayModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SizedArrayModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SizedArrayModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitSizedArrayModifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) SizedArrayModifier() (localctx ISizedArrayModifierContext) {
	localctx = NewSizedArrayModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CParserRULE_sizedArrayModifier)

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
		p.Match(CParserLeftBracket)
	}
	{
		p.SetState(215)
		p.expression(0)
	}
	{
		p.SetState(216)
		p.Match(CParserRightBracket)
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
	p.EnterRule(localctx, 32, CParserRULE_listInitialization)
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
		p.SetState(218)
		p.Match(CParserLeftBrace)
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CParserSizeof-21))|(1<<(CParserLeftParen-21))|(1<<(CParserPlus-21))|(1<<(CParserPlusPlus-21))|(1<<(CParserMinus-21))|(1<<(CParserMinusMinus-21))|(1<<(CParserStar-21))|(1<<(CParserAnd-21)))) != 0) || (((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(CParserNot-54))|(1<<(CParserTilde-54))|(1<<(CParserIdentifier-54))|(1<<(CParserIntegerConstant-54))|(1<<(CParserFloatingConstant-54))|(1<<(CParserCharacterConstant-54))|(1<<(CParserStringLiteral-54)))) != 0) {
		{
			p.SetState(219)
			p.expression(0)
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(220)
					p.Match(CParserComma)
				}
				{
					p.SetState(221)
					p.expression(0)
				}

			}
			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CParserComma {
			{
				p.SetState(227)
				p.Match(CParserComma)
			}

		}

	}
	{
		p.SetState(232)
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
	p.EnterRule(localctx, 34, CParserRULE_namedListInitialization)
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
		p.SetState(234)
		p.Match(CParserLeftBrace)
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserDot {
		{
			p.SetState(235)
			p.NamedListInitializationItem()
		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(236)
					p.Match(CParserComma)
				}
				{
					p.SetState(237)
					p.NamedListInitializationItem()
				}

			}
			p.SetState(242)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
		}
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CParserComma {
			{
				p.SetState(243)
				p.Match(CParserComma)
			}

		}

	}
	{
		p.SetState(248)
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
	p.EnterRule(localctx, 36, CParserRULE_namedListInitializationItem)

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
		p.SetState(250)
		p.Match(CParserDot)
	}
	{
		p.SetState(251)
		p.Match(CParserIdentifier)
	}
	{
		p.SetState(252)
		p.Match(CParserAssign)
	}
	{
		p.SetState(253)
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

type ConditionalBinaryExpressionContext struct {
	*ExpressionContext
}

func NewConditionalBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionalBinaryExpressionContext {
	var p = new(ConditionalBinaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ConditionalBinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalBinaryExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionalBinaryExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalBinaryExpressionContext) ConditionalBinaryOperator() IConditionalBinaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalBinaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalBinaryOperatorContext)
}

func (s *ConditionalBinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitConditionalBinaryExpression(s)

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

type EqualityBinaryExpressionContext struct {
	*ExpressionContext
}

func NewEqualityBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityBinaryExpressionContext {
	var p = new(EqualityBinaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqualityBinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityBinaryExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqualityBinaryExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqualityBinaryExpressionContext) EqualityOperator() IEqualityOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityOperatorContext)
}

func (s *EqualityBinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitEqualityBinaryExpression(s)

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
	_startState := 38
	p.EnterRecursionRule(localctx, 38, CParserRULE_expression, _p)
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
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(256)
			p.Match(CParserIdentifier)
		}

	case 2:
		localctx = NewConstantExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(257)
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
			p.SetState(258)
			p.Match(CParserStringLiteral)
		}

	case 4:
		localctx = NewStructInitializationExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(259)
			p.Match(CParserLeftParen)
		}
		{
			p.SetState(260)
			p.typeSpecifier(0)
		}
		{
			p.SetState(261)
			p.Match(CParserRightParen)
		}
		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(262)
				p.ListInitialization()
			}

		case 2:
			{
				p.SetState(263)
				p.NamedListInitialization()
			}

		}

	case 5:
		localctx = NewCastExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(266)
			p.Match(CParserLeftParen)
		}
		{
			p.SetState(267)
			p.typeSpecifier(0)
		}
		{
			p.SetState(268)
			p.Match(CParserRightParen)
		}
		{
			p.SetState(269)
			p.expression(10)
		}

	case 6:
		localctx = NewParenthesizedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(271)
			p.Match(CParserLeftParen)
		}
		{
			p.SetState(272)
			p.expression(0)
		}
		{
			p.SetState(273)
			p.Match(CParserRightParen)
		}

	case 7:
		localctx = NewUnaryExpressionPreContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(275)
			p.UnaryOperatorPre()
		}
		{
			p.SetState(276)
			p.expression(7)
		}

	case 8:
		localctx = NewSizeofExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(278)
			p.Match(CParserSizeof)
		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(279)
				p.Match(CParserLeftParen)
			}
			{
				p.SetState(280)
				p.typeSpecifier(0)
			}
			{
				p.SetState(281)
				p.Match(CParserRightParen)
			}

		case 2:
			{
				p.SetState(283)
				p.typeSpecifier(0)
			}

		case 3:
			{
				p.SetState(284)
				p.expression(0)
			}

		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(330)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAssignmentExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(290)
					p.AssignementOperator()
				}
				{
					p.SetState(291)
					p.expression(7)
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(294)
					p.BinaryOperator()
				}
				{
					p.SetState(295)
					p.expression(6)
				}

			case 3:
				localctx = NewEqualityBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(297)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(298)
					p.EqualityOperator()
				}
				{
					p.SetState(299)
					p.expression(5)
				}

			case 4:
				localctx = NewConditionalBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(301)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(302)
					p.ConditionalBinaryOperator()
				}
				{
					p.SetState(303)
					p.expression(4)
				}

			case 5:
				localctx = NewTernaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(306)
					p.Match(CParserQuestion)
				}
				{
					p.SetState(307)
					p.expression(0)
				}
				{
					p.SetState(308)
					p.Match(CParserColon)
				}
				{
					p.SetState(309)
					p.expression(2)
				}

			case 6:
				localctx = NewPropertyAccessExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(311)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(312)
					p.Match(CParserDot)
				}
				{
					p.SetState(313)
					p.Match(CParserIdentifier)
				}

			case 7:
				localctx = NewPointerPropertyAccessExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(314)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(315)
					p.Match(CParserArrow)
				}
				{
					p.SetState(316)
					p.Match(CParserIdentifier)
				}

			case 8:
				localctx = NewArrayIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(317)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(318)
					p.Match(CParserLeftBracket)
				}
				{
					p.SetState(319)
					p.expression(0)
				}
				{
					p.SetState(320)
					p.Match(CParserRightBracket)
				}

			case 9:
				localctx = NewFunctionCallExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(322)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(323)
					p.Match(CParserLeftParen)
				}
				p.SetState(325)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CParserSizeof-21))|(1<<(CParserLeftParen-21))|(1<<(CParserPlus-21))|(1<<(CParserPlusPlus-21))|(1<<(CParserMinus-21))|(1<<(CParserMinusMinus-21))|(1<<(CParserStar-21))|(1<<(CParserAnd-21)))) != 0) || (((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(CParserNot-54))|(1<<(CParserTilde-54))|(1<<(CParserIdentifier-54))|(1<<(CParserIntegerConstant-54))|(1<<(CParserFloatingConstant-54))|(1<<(CParserCharacterConstant-54))|(1<<(CParserStringLiteral-54)))) != 0) {
					{
						p.SetState(324)
						p.FunctionCallArguments()
					}

				}
				{
					p.SetState(327)
					p.Match(CParserRightParen)
				}

			case 10:
				localctx = NewUnaryExpressionPostContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CParserRULE_expression)
				p.SetState(328)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(329)
					p.UnaryOperatorPost()
				}

			}

		}
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 40, CParserRULE_assignementOperator)
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
		p.SetState(335)
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

// IConditionalBinaryOperatorContext is an interface to support dynamic dispatch.
type IConditionalBinaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalBinaryOperatorContext differentiates from other interfaces.
	IsConditionalBinaryOperatorContext()
}

type ConditionalBinaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalBinaryOperatorContext() *ConditionalBinaryOperatorContext {
	var p = new(ConditionalBinaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_conditionalBinaryOperator
	return p
}

func (*ConditionalBinaryOperatorContext) IsConditionalBinaryOperatorContext() {}

func NewConditionalBinaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalBinaryOperatorContext {
	var p = new(ConditionalBinaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_conditionalBinaryOperator

	return p
}

func (s *ConditionalBinaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalBinaryOperatorContext) AndAnd() antlr.TerminalNode {
	return s.GetToken(CParserAndAnd, 0)
}

func (s *ConditionalBinaryOperatorContext) OrOr() antlr.TerminalNode {
	return s.GetToken(CParserOrOr, 0)
}

func (s *ConditionalBinaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalBinaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalBinaryOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitConditionalBinaryOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) ConditionalBinaryOperator() (localctx IConditionalBinaryOperatorContext) {
	localctx = NewConditionalBinaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CParserRULE_conditionalBinaryOperator)
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
		p.SetState(337)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CParserAndAnd || _la == CParserOrOr) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IEqualityOperatorContext is an interface to support dynamic dispatch.
type IEqualityOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualityOperatorContext differentiates from other interfaces.
	IsEqualityOperatorContext()
}

type EqualityOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityOperatorContext() *EqualityOperatorContext {
	var p = new(EqualityOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CParserRULE_equalityOperator
	return p
}

func (*EqualityOperatorContext) IsEqualityOperatorContext() {}

func NewEqualityOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityOperatorContext {
	var p = new(EqualityOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CParserRULE_equalityOperator

	return p
}

func (s *EqualityOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityOperatorContext) Equal() antlr.TerminalNode {
	return s.GetToken(CParserEqual, 0)
}

func (s *EqualityOperatorContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(CParserNotEqual, 0)
}

func (s *EqualityOperatorContext) LessEqual() antlr.TerminalNode {
	return s.GetToken(CParserLessEqual, 0)
}

func (s *EqualityOperatorContext) GreaterEqual() antlr.TerminalNode {
	return s.GetToken(CParserGreaterEqual, 0)
}

func (s *EqualityOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CVisitor:
		return t.VisitEqualityOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CParser) EqualityOperator() (localctx IEqualityOperatorContext) {
	localctx = NewEqualityOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CParserRULE_equalityOperator)

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

	p.SetState(344)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CParserEqual:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(339)
			p.Match(CParserEqual)
		}

	case CParserNotEqual:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(340)
			p.Match(CParserNotEqual)
		}

	case CParserSizeof, CParserLeftParen, CParserPlus, CParserPlusPlus, CParserMinus, CParserMinusMinus, CParserStar, CParserAnd, CParserNot, CParserTilde, CParserIdentifier, CParserIntegerConstant, CParserFloatingConstant, CParserCharacterConstant, CParserStringLiteral:
		p.EnterOuterAlt(localctx, 3)

	case CParserLessEqual:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(342)
			p.Match(CParserLessEqual)
		}

	case CParserGreaterEqual:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(343)
			p.Match(CParserGreaterEqual)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 46, CParserRULE_binaryOperator)
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
		p.SetState(346)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(CParserLess-36))|(1<<(CParserGreater-36))|(1<<(CParserLeftShift-36))|(1<<(CParserRightShift-36))|(1<<(CParserPlus-36))|(1<<(CParserMinus-36))|(1<<(CParserStar-36))|(1<<(CParserDiv-36))|(1<<(CParserMod-36))|(1<<(CParserAnd-36))|(1<<(CParserOr-36))|(1<<(CParserCaret-36)))) != 0) {
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
	p.EnterRule(localctx, 48, CParserRULE_unaryOperatorPost)
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
		p.SetState(348)
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
	p.EnterRule(localctx, 50, CParserRULE_unaryOperatorPre)
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
		p.SetState(350)
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
	p.EnterRule(localctx, 52, CParserRULE_functionCallArguments)
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
		p.expression(0)
	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserComma {
		{
			p.SetState(353)
			p.Match(CParserComma)
		}
		{
			p.SetState(354)
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
	p.EnterRule(localctx, 54, CParserRULE_block)
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
		p.SetState(357)
		p.Match(CParserLeftBrace)
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserBreak)|(1<<CParserChar)|(1<<CParserContinue)|(1<<CParserDo)|(1<<CParserDouble)|(1<<CParserEnum)|(1<<CParserFloat)|(1<<CParserFor)|(1<<CParserGoto)|(1<<CParserIf)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserReturn)|(1<<CParserShort)|(1<<CParserSigned)|(1<<CParserSizeof)|(1<<CParserStruct)|(1<<CParserSwitch)|(1<<CParserUnsigned)|(1<<CParserVoid)|(1<<CParserWhile)|(1<<CParserLeftParen))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(CParserLeftBrace-34))|(1<<(CParserPlus-34))|(1<<(CParserPlusPlus-34))|(1<<(CParserMinus-34))|(1<<(CParserMinusMinus-34))|(1<<(CParserStar-34))|(1<<(CParserAnd-34))|(1<<(CParserNot-34))|(1<<(CParserTilde-34)))) != 0) || (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(CParserIdentifier-76))|(1<<(CParserIntegerConstant-76))|(1<<(CParserFloatingConstant-76))|(1<<(CParserCharacterConstant-76))|(1<<(CParserStringLiteral-76))|(1<<(CParserBlockComment-76))|(1<<(CParserLineComment-76)))) != 0) {
		{
			p.SetState(358)
			p.Statement()
		}

		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(364)
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
	p.EnterRule(localctx, 56, CParserRULE_statement)

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

	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(366)
				p.VariableDeclaration()
			}

		case 2:
			{
				p.SetState(367)
				p.expression(0)
			}

		case 3:
			{
				p.SetState(368)
				p.FunctionReturn()
			}

		case 4:
			{
				p.SetState(369)
				p.Match(CParserBreak)
			}

		case 5:
			{
				p.SetState(370)
				p.Match(CParserContinue)
			}

		case 6:
			{
				p.SetState(371)
				p.StructDeclaration()
			}

		case 7:
			{
				p.SetState(372)
				p.EnumDeclaration()
			}

		case 8:
			{
				p.SetState(373)
				p.GotoStatement()
			}

		case 9:
			{
				p.SetState(374)
				p.DoWhileStatement()
			}

		}
		{
			p.SetState(377)
			p.Match(CParserSemi)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(378)
			p.IfStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(379)
			p.SwitchStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(380)
			p.ForStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(381)
			p.WhileStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(382)
			p.Block()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(383)
			p.LabelStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(384)
			p.Match(CParserBlockComment)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(385)
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
	p.EnterRule(localctx, 58, CParserRULE_ifStatement)

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
		p.SetState(388)
		p.Match(CParserIf)
	}
	{
		p.SetState(389)
		p.Match(CParserLeftParen)
	}
	{
		p.SetState(390)
		p.expression(0)
	}
	{
		p.SetState(391)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(392)
		p.Statement()
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(393)
			p.Match(CParserElse)
		}
		{
			p.SetState(394)
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
	p.EnterRule(localctx, 60, CParserRULE_switchStatement)
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
		p.SetState(397)
		p.Match(CParserSwitch)
	}
	{
		p.SetState(398)
		p.Match(CParserLeftParen)
	}
	{
		p.SetState(399)
		p.expression(0)
	}
	{
		p.SetState(400)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(401)
		p.Match(CParserLeftBrace)
	}
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CParserCase {
		{
			p.SetState(402)
			p.CaseStatement()
		}

		p.SetState(407)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CParserDefault {
		{
			p.SetState(408)
			p.DefaultStatement()
		}

	}
	{
		p.SetState(411)
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
	p.EnterRule(localctx, 62, CParserRULE_caseStatement)
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
		p.SetState(413)
		p.Match(CParserCase)
	}
	{
		p.SetState(414)
		p.expression(0)
	}
	{
		p.SetState(415)
		p.Match(CParserColon)
	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserBreak)|(1<<CParserChar)|(1<<CParserContinue)|(1<<CParserDo)|(1<<CParserDouble)|(1<<CParserEnum)|(1<<CParserFloat)|(1<<CParserFor)|(1<<CParserGoto)|(1<<CParserIf)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserReturn)|(1<<CParserShort)|(1<<CParserSigned)|(1<<CParserSizeof)|(1<<CParserStruct)|(1<<CParserSwitch)|(1<<CParserUnsigned)|(1<<CParserVoid)|(1<<CParserWhile)|(1<<CParserLeftParen))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(CParserLeftBrace-34))|(1<<(CParserPlus-34))|(1<<(CParserPlusPlus-34))|(1<<(CParserMinus-34))|(1<<(CParserMinusMinus-34))|(1<<(CParserStar-34))|(1<<(CParserAnd-34))|(1<<(CParserNot-34))|(1<<(CParserTilde-34)))) != 0) || (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(CParserIdentifier-76))|(1<<(CParserIntegerConstant-76))|(1<<(CParserFloatingConstant-76))|(1<<(CParserCharacterConstant-76))|(1<<(CParserStringLiteral-76))|(1<<(CParserBlockComment-76))|(1<<(CParserLineComment-76)))) != 0) {
		{
			p.SetState(416)
			p.Statement()
		}

		p.SetState(421)
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
	p.EnterRule(localctx, 64, CParserRULE_defaultStatement)
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
		p.SetState(422)
		p.Match(CParserDefault)
	}
	{
		p.SetState(423)
		p.Match(CParserColon)
	}
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CParserBreak)|(1<<CParserChar)|(1<<CParserContinue)|(1<<CParserDo)|(1<<CParserDouble)|(1<<CParserEnum)|(1<<CParserFloat)|(1<<CParserFor)|(1<<CParserGoto)|(1<<CParserIf)|(1<<CParserInt)|(1<<CParserLong)|(1<<CParserReturn)|(1<<CParserShort)|(1<<CParserSigned)|(1<<CParserSizeof)|(1<<CParserStruct)|(1<<CParserSwitch)|(1<<CParserUnsigned)|(1<<CParserVoid)|(1<<CParserWhile)|(1<<CParserLeftParen))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(CParserLeftBrace-34))|(1<<(CParserPlus-34))|(1<<(CParserPlusPlus-34))|(1<<(CParserMinus-34))|(1<<(CParserMinusMinus-34))|(1<<(CParserStar-34))|(1<<(CParserAnd-34))|(1<<(CParserNot-34))|(1<<(CParserTilde-34)))) != 0) || (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(CParserIdentifier-76))|(1<<(CParserIntegerConstant-76))|(1<<(CParserFloatingConstant-76))|(1<<(CParserCharacterConstant-76))|(1<<(CParserStringLiteral-76))|(1<<(CParserBlockComment-76))|(1<<(CParserLineComment-76)))) != 0) {
		{
			p.SetState(424)
			p.Statement()
		}

		p.SetState(429)
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
	p.EnterRule(localctx, 66, CParserRULE_forStatement)
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
		p.SetState(430)
		p.Match(CParserFor)
	}
	{
		p.SetState(431)
		p.Match(CParserLeftParen)
	}
	p.SetState(434)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(432)

			var _x = p.expression(0)

			localctx.(*ForStatementContext).init = _x
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(433)
			p.VariableDeclaration()
		}

	}
	{
		p.SetState(436)
		p.Match(CParserSemi)
	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CParserSizeof-21))|(1<<(CParserLeftParen-21))|(1<<(CParserPlus-21))|(1<<(CParserPlusPlus-21))|(1<<(CParserMinus-21))|(1<<(CParserMinusMinus-21))|(1<<(CParserStar-21))|(1<<(CParserAnd-21)))) != 0) || (((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(CParserNot-54))|(1<<(CParserTilde-54))|(1<<(CParserIdentifier-54))|(1<<(CParserIntegerConstant-54))|(1<<(CParserFloatingConstant-54))|(1<<(CParserCharacterConstant-54))|(1<<(CParserStringLiteral-54)))) != 0) {
		{
			p.SetState(437)

			var _x = p.expression(0)

			localctx.(*ForStatementContext).condition = _x
		}

	}
	{
		p.SetState(440)
		p.Match(CParserSemi)
	}
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(CParserSizeof-21))|(1<<(CParserLeftParen-21))|(1<<(CParserPlus-21))|(1<<(CParserPlusPlus-21))|(1<<(CParserMinus-21))|(1<<(CParserMinusMinus-21))|(1<<(CParserStar-21))|(1<<(CParserAnd-21)))) != 0) || (((_la-54)&-(0x1f+1)) == 0 && ((1<<uint((_la-54)))&((1<<(CParserNot-54))|(1<<(CParserTilde-54))|(1<<(CParserIdentifier-54))|(1<<(CParserIntegerConstant-54))|(1<<(CParserFloatingConstant-54))|(1<<(CParserCharacterConstant-54))|(1<<(CParserStringLiteral-54)))) != 0) {
		{
			p.SetState(441)

			var _x = p.expression(0)

			localctx.(*ForStatementContext).post = _x
		}

	}
	{
		p.SetState(444)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(445)
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
	p.EnterRule(localctx, 68, CParserRULE_doWhileStatement)

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
		p.SetState(447)
		p.Match(CParserDo)
	}
	{
		p.SetState(448)
		p.Statement()
	}
	{
		p.SetState(449)
		p.Match(CParserWhile)
	}
	{
		p.SetState(450)
		p.Match(CParserLeftParen)
	}
	{
		p.SetState(451)
		p.expression(0)
	}
	{
		p.SetState(452)
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
	p.EnterRule(localctx, 70, CParserRULE_whileStatement)

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
		p.SetState(454)
		p.Match(CParserWhile)
	}
	{
		p.SetState(455)
		p.Match(CParserLeftParen)
	}
	{
		p.SetState(456)
		p.expression(0)
	}
	{
		p.SetState(457)
		p.Match(CParserRightParen)
	}
	{
		p.SetState(458)
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
	p.EnterRule(localctx, 72, CParserRULE_gotoStatement)

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
		p.SetState(460)
		p.Match(CParserGoto)
	}
	{
		p.SetState(461)
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
	p.EnterRule(localctx, 74, CParserRULE_labelStatement)

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
		p.SetState(463)
		p.Match(CParserIdentifier)
	}
	{
		p.SetState(464)
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
	p.EnterRule(localctx, 76, CParserRULE_includePreprocessor)

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
		p.SetState(466)
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

	case 19:
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
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
