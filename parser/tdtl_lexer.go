// Generated from TDTL.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 49, 489,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 18, 3, 18, 5, 18, 220, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 233, 10, 20, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 241, 10, 21, 3, 22, 3, 22, 3,
	22, 3, 22, 5, 22, 247, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	5, 23, 255, 10, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5,
	24, 264, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 271, 10, 25,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32,
	3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 40, 3, 40, 7, 40, 338, 10, 40, 12, 40, 14, 40, 341, 11, 40, 3, 41, 3,
	41, 3, 41, 7, 41, 346, 10, 41, 12, 41, 14, 41, 349, 11, 41, 5, 41, 351,
	10, 41, 3, 42, 5, 42, 354, 10, 42, 3, 42, 3, 42, 3, 43, 5, 43, 359, 10,
	43, 3, 43, 6, 43, 362, 10, 43, 13, 43, 14, 43, 363, 3, 43, 3, 43, 6, 43,
	368, 10, 43, 13, 43, 14, 43, 369, 3, 43, 6, 43, 373, 10, 43, 13, 43, 14,
	43, 374, 3, 43, 3, 43, 3, 43, 3, 43, 6, 43, 381, 10, 43, 13, 43, 14, 43,
	382, 5, 43, 385, 10, 43, 3, 44, 6, 44, 388, 10, 44, 13, 44, 14, 44, 389,
	3, 45, 3, 45, 5, 45, 394, 10, 45, 3, 45, 3, 45, 3, 45, 5, 45, 399, 10,
	45, 7, 45, 401, 10, 45, 12, 45, 14, 45, 404, 11, 45, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 46, 5, 46, 413, 10, 46, 3, 47, 3, 47, 3, 47, 3,
	47, 7, 47, 419, 10, 47, 12, 47, 14, 47, 422, 11, 47, 3, 47, 3, 47, 3, 48,
	6, 48, 427, 10, 48, 13, 48, 14, 48, 428, 3, 48, 3, 48, 3, 49, 3, 49, 3,
	50, 3, 50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 54, 3, 54, 3, 55,
	3, 55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 60, 3,
	60, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3, 65,
	3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 70, 3, 70, 3,
	71, 3, 71, 3, 72, 3, 72, 3, 73, 3, 73, 3, 74, 3, 74, 3, 75, 6, 75, 486,
	10, 75, 13, 75, 14, 75, 487, 2, 2, 76, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44,
	87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 2, 99, 2, 101, 2, 103, 2, 105,
	2, 107, 2, 109, 2, 111, 2, 113, 2, 115, 2, 117, 2, 119, 2, 121, 2, 123,
	2, 125, 2, 127, 2, 129, 2, 131, 2, 133, 2, 135, 2, 137, 2, 139, 2, 141,
	2, 143, 2, 145, 2, 147, 2, 149, 2, 3, 2, 36, 6, 2, 37, 37, 67, 92, 97,
	97, 99, 124, 8, 2, 37, 38, 47, 47, 50, 59, 66, 92, 97, 97, 99, 124, 3,
	2, 51, 59, 3, 2, 50, 59, 4, 2, 45, 45, 47, 47, 8, 2, 37, 38, 47, 47, 49,
	59, 66, 92, 97, 97, 99, 124, 3, 2, 41, 41, 5, 2, 11, 12, 15, 15, 34, 34,
	4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69, 101, 101, 4,
	2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 104, 104, 4,
	2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4,
	2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4,
	2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4,
	2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4,
	2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4,
	2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4,
	2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124, 2, 489, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2,
	2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3,
	2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43,
	3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2,
	51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2,
	2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2,
	2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2,
	2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3,
	2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89,
	3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 3,
	151, 3, 2, 2, 2, 5, 153, 3, 2, 2, 2, 7, 155, 3, 2, 2, 2, 9, 157, 3, 2,
	2, 2, 11, 159, 3, 2, 2, 2, 13, 161, 3, 2, 2, 2, 15, 163, 3, 2, 2, 2, 17,
	165, 3, 2, 2, 2, 19, 168, 3, 2, 2, 2, 21, 172, 3, 2, 2, 2, 23, 179, 3,
	2, 2, 2, 25, 184, 3, 2, 2, 2, 27, 189, 3, 2, 2, 2, 29, 195, 3, 2, 2, 2,
	31, 202, 3, 2, 2, 2, 33, 209, 3, 2, 2, 2, 35, 219, 3, 2, 2, 2, 37, 221,
	3, 2, 2, 2, 39, 232, 3, 2, 2, 2, 41, 240, 3, 2, 2, 2, 43, 246, 3, 2, 2,
	2, 45, 254, 3, 2, 2, 2, 47, 263, 3, 2, 2, 2, 49, 270, 3, 2, 2, 2, 51, 272,
	3, 2, 2, 2, 53, 277, 3, 2, 2, 2, 55, 282, 3, 2, 2, 2, 57, 290, 3, 2, 2,
	2, 59, 297, 3, 2, 2, 2, 61, 305, 3, 2, 2, 2, 63, 312, 3, 2, 2, 2, 65, 314,
	3, 2, 2, 2, 67, 316, 3, 2, 2, 2, 69, 318, 3, 2, 2, 2, 71, 320, 3, 2, 2,
	2, 73, 322, 3, 2, 2, 2, 75, 324, 3, 2, 2, 2, 77, 329, 3, 2, 2, 2, 79, 335,
	3, 2, 2, 2, 81, 350, 3, 2, 2, 2, 83, 353, 3, 2, 2, 2, 85, 358, 3, 2, 2,
	2, 87, 387, 3, 2, 2, 2, 89, 391, 3, 2, 2, 2, 91, 412, 3, 2, 2, 2, 93, 414,
	3, 2, 2, 2, 95, 426, 3, 2, 2, 2, 97, 432, 3, 2, 2, 2, 99, 434, 3, 2, 2,
	2, 101, 436, 3, 2, 2, 2, 103, 438, 3, 2, 2, 2, 105, 440, 3, 2, 2, 2, 107,
	442, 3, 2, 2, 2, 109, 444, 3, 2, 2, 2, 111, 446, 3, 2, 2, 2, 113, 448,
	3, 2, 2, 2, 115, 450, 3, 2, 2, 2, 117, 452, 3, 2, 2, 2, 119, 454, 3, 2,
	2, 2, 121, 456, 3, 2, 2, 2, 123, 458, 3, 2, 2, 2, 125, 460, 3, 2, 2, 2,
	127, 462, 3, 2, 2, 2, 129, 464, 3, 2, 2, 2, 131, 466, 3, 2, 2, 2, 133,
	468, 3, 2, 2, 2, 135, 470, 3, 2, 2, 2, 137, 472, 3, 2, 2, 2, 139, 474,
	3, 2, 2, 2, 141, 476, 3, 2, 2, 2, 143, 478, 3, 2, 2, 2, 145, 480, 3, 2,
	2, 2, 147, 482, 3, 2, 2, 2, 149, 485, 3, 2, 2, 2, 151, 152, 7, 46, 2, 2,
	152, 4, 3, 2, 2, 2, 153, 154, 7, 42, 2, 2, 154, 6, 3, 2, 2, 2, 155, 156,
	7, 43, 2, 2, 156, 8, 3, 2, 2, 2, 157, 158, 7, 36, 2, 2, 158, 10, 3, 2,
	2, 2, 159, 160, 7, 93, 2, 2, 160, 12, 3, 2, 2, 2, 161, 162, 7, 95, 2, 2,
	162, 14, 3, 2, 2, 2, 163, 164, 7, 37, 2, 2, 164, 16, 3, 2, 2, 2, 165, 166,
	7, 93, 2, 2, 166, 167, 7, 95, 2, 2, 167, 18, 3, 2, 2, 2, 168, 169, 7, 93,
	2, 2, 169, 170, 7, 37, 2, 2, 170, 171, 7, 95, 2, 2, 171, 20, 3, 2, 2, 2,
	172, 173, 5, 113, 57, 2, 173, 174, 5, 123, 62, 2, 174, 175, 5, 133, 67,
	2, 175, 176, 5, 105, 53, 2, 176, 177, 5, 131, 66, 2, 177, 178, 5, 135,
	68, 2, 178, 22, 3, 2, 2, 2, 179, 180, 5, 113, 57, 2, 180, 181, 5, 123,
	62, 2, 181, 182, 5, 135, 68, 2, 182, 183, 5, 125, 63, 2, 183, 24, 3, 2,
	2, 2, 184, 185, 5, 149, 75, 2, 185, 186, 5, 97, 49, 2, 186, 187, 5, 133,
	67, 2, 187, 188, 5, 149, 75, 2, 188, 26, 3, 2, 2, 2, 189, 190, 5, 149,
	75, 2, 190, 191, 5, 97, 49, 2, 191, 192, 5, 123, 62, 2, 192, 193, 5, 103,
	52, 2, 193, 194, 5, 149, 75, 2, 194, 28, 3, 2, 2, 2, 195, 196, 5, 149,
	75, 2, 196, 197, 5, 101, 51, 2, 197, 198, 5, 97, 49, 2, 198, 199, 5, 133,
	67, 2, 199, 200, 5, 105, 53, 2, 200, 201, 5, 149, 75, 2, 201, 30, 3, 2,
	2, 2, 202, 203, 5, 149, 75, 2, 203, 204, 5, 105, 53, 2, 204, 205, 5, 119,
	60, 2, 205, 206, 5, 133, 67, 2, 206, 207, 5, 105, 53, 2, 207, 208, 5, 149,
	75, 2, 208, 32, 3, 2, 2, 2, 209, 210, 5, 149, 75, 2, 210, 211, 5, 105,
	53, 2, 211, 212, 5, 123, 62, 2, 212, 213, 5, 103, 52, 2, 213, 214, 5, 149,
	75, 2, 214, 34, 3, 2, 2, 2, 215, 216, 5, 105, 53, 2, 216, 217, 5, 129,
	65, 2, 217, 220, 3, 2, 2, 2, 218, 220, 7, 63, 2, 2, 219, 215, 3, 2, 2,
	2, 219, 218, 3, 2, 2, 2, 220, 36, 3, 2, 2, 2, 221, 222, 5, 149, 75, 2,
	222, 223, 5, 107, 54, 2, 223, 224, 5, 131, 66, 2, 224, 225, 5, 125, 63,
	2, 225, 226, 5, 121, 61, 2, 226, 227, 5, 149, 75, 2, 227, 38, 3, 2, 2,
	2, 228, 229, 5, 109, 55, 2, 229, 230, 5, 135, 68, 2, 230, 233, 3, 2, 2,
	2, 231, 233, 7, 64, 2, 2, 232, 228, 3, 2, 2, 2, 232, 231, 3, 2, 2, 2, 233,
	40, 3, 2, 2, 2, 234, 235, 5, 109, 55, 2, 235, 236, 5, 135, 68, 2, 236,
	237, 5, 105, 53, 2, 237, 241, 3, 2, 2, 2, 238, 239, 7, 64, 2, 2, 239, 241,
	7, 63, 2, 2, 240, 234, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 241, 42, 3, 2,
	2, 2, 242, 243, 5, 119, 60, 2, 243, 244, 5, 135, 68, 2, 244, 247, 3, 2,
	2, 2, 245, 247, 7, 62, 2, 2, 246, 242, 3, 2, 2, 2, 246, 245, 3, 2, 2, 2,
	247, 44, 3, 2, 2, 2, 248, 249, 5, 119, 60, 2, 249, 250, 5, 135, 68, 2,
	250, 251, 5, 105, 53, 2, 251, 255, 3, 2, 2, 2, 252, 253, 7, 62, 2, 2, 253,
	255, 7, 63, 2, 2, 254, 248, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 255, 46,
	3, 2, 2, 2, 256, 257, 5, 123, 62, 2, 257, 258, 5, 105, 53, 2, 258, 264,
	3, 2, 2, 2, 259, 260, 7, 35, 2, 2, 260, 264, 7, 63, 2, 2, 261, 262, 7,
	62, 2, 2, 262, 264, 7, 64, 2, 2, 263, 256, 3, 2, 2, 2, 263, 259, 3, 2,
	2, 2, 263, 261, 3, 2, 2, 2, 264, 48, 3, 2, 2, 2, 265, 266, 5, 123, 62,
	2, 266, 267, 5, 125, 63, 2, 267, 268, 5, 135, 68, 2, 268, 271, 3, 2, 2,
	2, 269, 271, 7, 35, 2, 2, 270, 265, 3, 2, 2, 2, 270, 269, 3, 2, 2, 2, 271,
	50, 3, 2, 2, 2, 272, 273, 5, 123, 62, 2, 273, 274, 5, 137, 69, 2, 274,
	275, 5, 119, 60, 2, 275, 276, 5, 119, 60, 2, 276, 52, 3, 2, 2, 2, 277,
	278, 5, 149, 75, 2, 278, 279, 5, 125, 63, 2, 279, 280, 5, 131, 66, 2, 280,
	281, 5, 149, 75, 2, 281, 54, 3, 2, 2, 2, 282, 283, 5, 133, 67, 2, 283,
	284, 5, 105, 53, 2, 284, 285, 5, 119, 60, 2, 285, 286, 5, 105, 53, 2, 286,
	287, 5, 101, 51, 2, 287, 288, 5, 135, 68, 2, 288, 289, 5, 149, 75, 2, 289,
	56, 3, 2, 2, 2, 290, 291, 5, 149, 75, 2, 291, 292, 5, 135, 68, 2, 292,
	293, 5, 111, 56, 2, 293, 294, 5, 105, 53, 2, 294, 295, 5, 123, 62, 2, 295,
	296, 5, 149, 75, 2, 296, 58, 3, 2, 2, 2, 297, 298, 5, 149, 75, 2, 298,
	299, 5, 141, 71, 2, 299, 300, 5, 111, 56, 2, 300, 301, 5, 105, 53, 2, 301,
	302, 5, 131, 66, 2, 302, 303, 5, 105, 53, 2, 303, 304, 5, 149, 75, 2, 304,
	60, 3, 2, 2, 2, 305, 306, 5, 149, 75, 2, 306, 307, 5, 141, 71, 2, 307,
	308, 5, 111, 56, 2, 308, 309, 5, 105, 53, 2, 309, 310, 5, 123, 62, 2, 310,
	311, 5, 149, 75, 2, 311, 62, 3, 2, 2, 2, 312, 313, 7, 44, 2, 2, 313, 64,
	3, 2, 2, 2, 314, 315, 7, 49, 2, 2, 315, 66, 3, 2, 2, 2, 316, 317, 7, 39,
	2, 2, 317, 68, 3, 2, 2, 2, 318, 319, 7, 45, 2, 2, 319, 70, 3, 2, 2, 2,
	320, 321, 7, 47, 2, 2, 321, 72, 3, 2, 2, 2, 322, 323, 7, 48, 2, 2, 323,
	74, 3, 2, 2, 2, 324, 325, 5, 135, 68, 2, 325, 326, 5, 131, 66, 2, 326,
	327, 5, 137, 69, 2, 327, 328, 5, 105, 53, 2, 328, 76, 3, 2, 2, 2, 329,
	330, 5, 107, 54, 2, 330, 331, 5, 97, 49, 2, 331, 332, 5, 119, 60, 2, 332,
	333, 5, 133, 67, 2, 333, 334, 5, 105, 53, 2, 334, 78, 3, 2, 2, 2, 335,
	339, 9, 2, 2, 2, 336, 338, 9, 3, 2, 2, 337, 336, 3, 2, 2, 2, 338, 341,
	3, 2, 2, 2, 339, 337, 3, 2, 2, 2, 339, 340, 3, 2, 2, 2, 340, 80, 3, 2,
	2, 2, 341, 339, 3, 2, 2, 2, 342, 351, 7, 50, 2, 2, 343, 347, 9, 4, 2, 2,
	344, 346, 9, 5, 2, 2, 345, 344, 3, 2, 2, 2, 346, 349, 3, 2, 2, 2, 347,
	345, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 351, 3, 2, 2, 2, 349, 347,
	3, 2, 2, 2, 350, 342, 3, 2, 2, 2, 350, 343, 3, 2, 2, 2, 351, 82, 3, 2,
	2, 2, 352, 354, 9, 6, 2, 2, 353, 352, 3, 2, 2, 2, 353, 354, 3, 2, 2, 2,
	354, 355, 3, 2, 2, 2, 355, 356, 5, 81, 41, 2, 356, 84, 3, 2, 2, 2, 357,
	359, 9, 6, 2, 2, 358, 357, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 384,
	3, 2, 2, 2, 360, 362, 5, 81, 41, 2, 361, 360, 3, 2, 2, 2, 362, 363, 3,
	2, 2, 2, 363, 361, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 365, 3, 2, 2,
	2, 365, 367, 5, 73, 37, 2, 366, 368, 5, 81, 41, 2, 367, 366, 3, 2, 2, 2,
	368, 369, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370,
	385, 3, 2, 2, 2, 371, 373, 5, 81, 41, 2, 372, 371, 3, 2, 2, 2, 373, 374,
	3, 2, 2, 2, 374, 372, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375, 376, 3, 2,
	2, 2, 376, 377, 5, 73, 37, 2, 377, 385, 3, 2, 2, 2, 378, 380, 5, 73, 37,
	2, 379, 381, 5, 81, 41, 2, 380, 379, 3, 2, 2, 2, 381, 382, 3, 2, 2, 2,
	382, 380, 3, 2, 2, 2, 382, 383, 3, 2, 2, 2, 383, 385, 3, 2, 2, 2, 384,
	361, 3, 2, 2, 2, 384, 372, 3, 2, 2, 2, 384, 378, 3, 2, 2, 2, 385, 86, 3,
	2, 2, 2, 386, 388, 9, 7, 2, 2, 387, 386, 3, 2, 2, 2, 388, 389, 3, 2, 2,
	2, 389, 387, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 88, 3, 2, 2, 2, 391,
	393, 5, 87, 44, 2, 392, 394, 5, 91, 46, 2, 393, 392, 3, 2, 2, 2, 393, 394,
	3, 2, 2, 2, 394, 402, 3, 2, 2, 2, 395, 396, 5, 73, 37, 2, 396, 398, 5,
	87, 44, 2, 397, 399, 5, 91, 46, 2, 398, 397, 3, 2, 2, 2, 398, 399, 3, 2,
	2, 2, 399, 401, 3, 2, 2, 2, 400, 395, 3, 2, 2, 2, 401, 404, 3, 2, 2, 2,
	402, 400, 3, 2, 2, 2, 402, 403, 3, 2, 2, 2, 403, 90, 3, 2, 2, 2, 404, 402,
	3, 2, 2, 2, 405, 406, 7, 93, 2, 2, 406, 407, 5, 81, 41, 2, 407, 408, 7,
	95, 2, 2, 408, 413, 3, 2, 2, 2, 409, 410, 7, 93, 2, 2, 410, 411, 7, 37,
	2, 2, 411, 413, 7, 95, 2, 2, 412, 405, 3, 2, 2, 2, 412, 409, 3, 2, 2, 2,
	413, 92, 3, 2, 2, 2, 414, 420, 7, 41, 2, 2, 415, 419, 10, 8, 2, 2, 416,
	417, 7, 41, 2, 2, 417, 419, 7, 41, 2, 2, 418, 415, 3, 2, 2, 2, 418, 416,
	3, 2, 2, 2, 419, 422, 3, 2, 2, 2, 420, 418, 3, 2, 2, 2, 420, 421, 3, 2,
	2, 2, 421, 423, 3, 2, 2, 2, 422, 420, 3, 2, 2, 2, 423, 424, 7, 41, 2, 2,
	424, 94, 3, 2, 2, 2, 425, 427, 9, 9, 2, 2, 426, 425, 3, 2, 2, 2, 427, 428,
	3, 2, 2, 2, 428, 426, 3, 2, 2, 2, 428, 429, 3, 2, 2, 2, 429, 430, 3, 2,
	2, 2, 430, 431, 8, 48, 2, 2, 431, 96, 3, 2, 2, 2, 432, 433, 9, 10, 2, 2,
	433, 98, 3, 2, 2, 2, 434, 435, 9, 11, 2, 2, 435, 100, 3, 2, 2, 2, 436,
	437, 9, 12, 2, 2, 437, 102, 3, 2, 2, 2, 438, 439, 9, 13, 2, 2, 439, 104,
	3, 2, 2, 2, 440, 441, 9, 14, 2, 2, 441, 106, 3, 2, 2, 2, 442, 443, 9, 15,
	2, 2, 443, 108, 3, 2, 2, 2, 444, 445, 9, 16, 2, 2, 445, 110, 3, 2, 2, 2,
	446, 447, 9, 17, 2, 2, 447, 112, 3, 2, 2, 2, 448, 449, 9, 18, 2, 2, 449,
	114, 3, 2, 2, 2, 450, 451, 9, 19, 2, 2, 451, 116, 3, 2, 2, 2, 452, 453,
	9, 20, 2, 2, 453, 118, 3, 2, 2, 2, 454, 455, 9, 21, 2, 2, 455, 120, 3,
	2, 2, 2, 456, 457, 9, 22, 2, 2, 457, 122, 3, 2, 2, 2, 458, 459, 9, 23,
	2, 2, 459, 124, 3, 2, 2, 2, 460, 461, 9, 24, 2, 2, 461, 126, 3, 2, 2, 2,
	462, 463, 9, 25, 2, 2, 463, 128, 3, 2, 2, 2, 464, 465, 9, 26, 2, 2, 465,
	130, 3, 2, 2, 2, 466, 467, 9, 27, 2, 2, 467, 132, 3, 2, 2, 2, 468, 469,
	9, 28, 2, 2, 469, 134, 3, 2, 2, 2, 470, 471, 9, 29, 2, 2, 471, 136, 3,
	2, 2, 2, 472, 473, 9, 30, 2, 2, 473, 138, 3, 2, 2, 2, 474, 475, 9, 31,
	2, 2, 475, 140, 3, 2, 2, 2, 476, 477, 9, 32, 2, 2, 477, 142, 3, 2, 2, 2,
	478, 479, 9, 33, 2, 2, 479, 144, 3, 2, 2, 2, 480, 481, 9, 34, 2, 2, 481,
	146, 3, 2, 2, 2, 482, 483, 9, 35, 2, 2, 483, 148, 3, 2, 2, 2, 484, 486,
	9, 9, 2, 2, 485, 484, 3, 2, 2, 2, 486, 487, 3, 2, 2, 2, 487, 485, 3, 2,
	2, 2, 487, 488, 3, 2, 2, 2, 488, 150, 3, 2, 2, 2, 29, 2, 219, 232, 240,
	246, 254, 263, 270, 339, 347, 350, 353, 358, 363, 369, 374, 382, 384, 389,
	393, 398, 402, 412, 418, 420, 428, 487, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'('", "')'", "'\"'", "'['", "']'", "'#'", "'[]'", "'[#]'",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "'*'", "'/'", "'%'", "'+'", "'-'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "INSERT", "INTO", "AS", "AND",
	"CASE", "ELSE", "END", "EQ", "FROM", "GT", "GTE", "LT", "LTE", "NE", "NOT",
	"NULL", "OR", "SELECT", "THEN", "WHERE", "WHEN", "MUL", "DIV", "MOD", "ADD",
	"SUB", "DOT", "TRUE", "FALSE", "INDENTIFIER", "NUMBER", "INTEGER", "FLOAT",
	"TOPICITEM", "PATHITEM", "ARRAYITEM", "STRING", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"INSERT", "INTO", "AS", "AND", "CASE", "ELSE", "END", "EQ", "FROM", "GT",
	"GTE", "LT", "LTE", "NE", "NOT", "NULL", "OR", "SELECT", "THEN", "WHERE",
	"WHEN", "MUL", "DIV", "MOD", "ADD", "SUB", "DOT", "TRUE", "FALSE", "INDENTIFIER",
	"NUMBER", "INTEGER", "FLOAT", "TOPICITEM", "PATHITEM", "ARRAYITEM", "STRING",
	"WHITESPACE", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
	"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "STUFF",
}

type TDTLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewTDTLLexer(input antlr.CharStream) *TDTLLexer {

	l := new(TDTLLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "TDTL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TDTLLexer tokens.
const (
	TDTLLexerT__0        = 1
	TDTLLexerT__1        = 2
	TDTLLexerT__2        = 3
	TDTLLexerT__3        = 4
	TDTLLexerT__4        = 5
	TDTLLexerT__5        = 6
	TDTLLexerT__6        = 7
	TDTLLexerT__7        = 8
	TDTLLexerT__8        = 9
	TDTLLexerINSERT      = 10
	TDTLLexerINTO        = 11
	TDTLLexerAS          = 12
	TDTLLexerAND         = 13
	TDTLLexerCASE        = 14
	TDTLLexerELSE        = 15
	TDTLLexerEND         = 16
	TDTLLexerEQ          = 17
	TDTLLexerFROM        = 18
	TDTLLexerGT          = 19
	TDTLLexerGTE         = 20
	TDTLLexerLT          = 21
	TDTLLexerLTE         = 22
	TDTLLexerNE          = 23
	TDTLLexerNOT         = 24
	TDTLLexerNULL        = 25
	TDTLLexerOR          = 26
	TDTLLexerSELECT      = 27
	TDTLLexerTHEN        = 28
	TDTLLexerWHERE       = 29
	TDTLLexerWHEN        = 30
	TDTLLexerMUL         = 31
	TDTLLexerDIV         = 32
	TDTLLexerMOD         = 33
	TDTLLexerADD         = 34
	TDTLLexerSUB         = 35
	TDTLLexerDOT         = 36
	TDTLLexerTRUE        = 37
	TDTLLexerFALSE       = 38
	TDTLLexerINDENTIFIER = 39
	TDTLLexerNUMBER      = 40
	TDTLLexerINTEGER     = 41
	TDTLLexerFLOAT       = 42
	TDTLLexerTOPICITEM   = 43
	TDTLLexerPATHITEM    = 44
	TDTLLexerARRAYITEM   = 45
	TDTLLexerSTRING      = 46
	TDTLLexerWHITESPACE  = 47
)
