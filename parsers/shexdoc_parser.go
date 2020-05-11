// Code generated from ../../ShExDoc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ShExDoc

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 74, 710, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9, 
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44, 
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4, 
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55, 
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9, 
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65, 
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 70, 4, 
	71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75, 4, 76, 
	9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 3, 2, 7, 2, 158, 10, 2, 12, 2, 14, 2, 
	161, 11, 2, 3, 2, 3, 2, 5, 2, 165, 10, 2, 3, 2, 7, 2, 168, 10, 2, 12, 2, 
	14, 2, 171, 11, 2, 5, 2, 173, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 
	180, 10, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 
	3, 7, 3, 7, 5, 7, 194, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 6, 9, 201, 
	10, 9, 13, 9, 14, 9, 202, 3, 10, 3, 10, 5, 10, 207, 10, 10, 3, 11, 5, 11, 
	210, 10, 11, 3, 11, 3, 11, 7, 11, 214, 10, 11, 12, 11, 14, 11, 217, 11, 
	11, 3, 11, 3, 11, 5, 11, 221, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 
	3, 14, 3, 14, 7, 14, 230, 10, 14, 12, 14, 14, 14, 233, 11, 14, 3, 15, 3, 
	15, 3, 15, 7, 15, 238, 10, 15, 12, 15, 14, 15, 241, 11, 15, 3, 16, 3, 16, 
	3, 16, 7, 16, 246, 10, 16, 12, 16, 14, 16, 249, 11, 16, 3, 17, 3, 17, 3, 
	17, 7, 17, 254, 10, 17, 12, 17, 14, 17, 257, 11, 17, 3, 18, 5, 18, 260, 
	10, 18, 3, 18, 3, 18, 3, 19, 5, 19, 265, 10, 19, 3, 19, 3, 19, 3, 20, 3, 
	20, 5, 20, 271, 10, 20, 3, 20, 3, 20, 3, 20, 5, 20, 276, 10, 20, 3, 20, 
	3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 283, 10, 20, 3, 21, 3, 21, 5, 21, 287, 
	10, 21, 3, 21, 3, 21, 3, 21, 5, 21, 292, 10, 21, 3, 21, 3, 21, 3, 21, 3, 
	21, 3, 21, 5, 21, 299, 10, 21, 3, 22, 3, 22, 5, 22, 303, 10, 22, 3, 23, 
	3, 23, 5, 23, 307, 10, 23, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 313, 10, 
	24, 3, 25, 3, 25, 7, 25, 317, 10, 25, 12, 25, 14, 25, 320, 11, 25, 3, 25, 
	3, 25, 7, 25, 324, 10, 25, 12, 25, 14, 25, 327, 11, 25, 3, 25, 3, 25, 7, 
	25, 331, 10, 25, 12, 25, 14, 25, 334, 11, 25, 3, 25, 3, 25, 7, 25, 338, 
	10, 25, 12, 25, 14, 25, 341, 11, 25, 3, 25, 6, 25, 344, 10, 25, 13, 25, 
	14, 25, 345, 5, 25, 348, 10, 25, 3, 26, 3, 26, 7, 26, 352, 10, 26, 12, 
	26, 14, 26, 355, 11, 26, 3, 26, 7, 26, 358, 10, 26, 12, 26, 14, 26, 361, 
	11, 26, 3, 27, 3, 27, 7, 27, 365, 10, 27, 12, 27, 14, 27, 368, 11, 27, 
	3, 27, 6, 27, 371, 10, 27, 13, 27, 14, 27, 372, 5, 27, 375, 10, 27, 3, 
	28, 3, 28, 7, 28, 379, 10, 28, 12, 28, 14, 28, 382, 11, 28, 3, 28, 7, 28, 
	385, 10, 28, 12, 28, 14, 28, 388, 11, 28, 3, 29, 3, 29, 3, 30, 3, 30, 5, 
	30, 394, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 401, 10, 31, 
	5, 31, 403, 10, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 
	33, 5, 33, 413, 10, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 
	3, 37, 7, 37, 423, 10, 37, 12, 37, 14, 37, 426, 11, 37, 3, 37, 7, 37, 429, 
	10, 37, 12, 37, 14, 37, 432, 11, 37, 3, 38, 7, 38, 435, 10, 38, 12, 38, 
	14, 38, 438, 11, 38, 3, 38, 3, 38, 5, 38, 442, 10, 38, 3, 38, 3, 38, 3, 
	39, 3, 39, 3, 39, 5, 39, 449, 10, 39, 3, 40, 3, 40, 6, 40, 453, 10, 40, 
	13, 40, 14, 40, 454, 3, 41, 3, 41, 3, 42, 3, 42, 5, 42, 461, 10, 42, 3, 
	43, 3, 43, 3, 43, 6, 43, 466, 10, 43, 13, 43, 14, 43, 467, 3, 44, 3, 44, 
	5, 44, 472, 10, 44, 3, 45, 3, 45, 5, 45, 476, 10, 45, 3, 46, 3, 46, 3, 
	46, 6, 46, 481, 10, 46, 13, 46, 14, 46, 482, 3, 46, 5, 46, 486, 10, 46, 
	3, 47, 3, 47, 5, 47, 490, 10, 47, 3, 47, 3, 47, 5, 47, 494, 10, 47, 3, 
	47, 5, 47, 497, 10, 47, 3, 48, 3, 48, 3, 48, 3, 48, 5, 48, 503, 10, 48, 
	3, 48, 7, 48, 506, 10, 48, 12, 48, 14, 48, 509, 11, 48, 3, 48, 7, 48, 512, 
	10, 48, 12, 48, 14, 48, 515, 11, 48, 3, 49, 5, 49, 518, 10, 49, 3, 49, 
	3, 49, 3, 49, 5, 49, 523, 10, 49, 3, 49, 7, 49, 526, 10, 49, 12, 49, 14, 
	49, 529, 11, 49, 3, 49, 7, 49, 532, 10, 49, 12, 49, 14, 49, 535, 11, 49, 
	3, 50, 3, 50, 3, 50, 3, 50, 5, 50, 541, 10, 50, 3, 51, 3, 51, 3, 51, 3, 
	51, 3, 51, 3, 51, 3, 51, 5, 51, 550, 10, 51, 3, 51, 5, 51, 553, 10, 51, 
	3, 52, 3, 52, 3, 53, 3, 53, 7, 53, 559, 10, 53, 12, 53, 14, 53, 562, 11, 
	53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 6, 54, 571, 10, 54, 
	13, 54, 14, 54, 572, 3, 54, 6, 54, 576, 10, 54, 13, 54, 14, 54, 577, 3, 
	54, 6, 54, 581, 10, 54, 13, 54, 14, 54, 582, 5, 54, 585, 10, 54, 5, 54, 
	587, 10, 54, 3, 55, 3, 55, 3, 55, 7, 55, 592, 10, 55, 12, 55, 14, 55, 595, 
	11, 55, 5, 55, 597, 10, 55, 3, 56, 3, 56, 3, 56, 5, 56, 602, 10, 56, 3, 
	57, 3, 57, 3, 57, 7, 57, 607, 10, 57, 12, 57, 14, 57, 610, 11, 57, 5, 57, 
	612, 10, 57, 3, 58, 3, 58, 3, 58, 5, 58, 617, 10, 58, 3, 59, 3, 59, 3, 
	59, 7, 59, 622, 10, 59, 12, 59, 14, 59, 625, 11, 59, 5, 59, 627, 10, 59, 
	3, 59, 3, 59, 3, 59, 7, 59, 632, 10, 59, 12, 59, 14, 59, 635, 11, 59, 5, 
	59, 637, 10, 59, 3, 60, 3, 60, 3, 60, 5, 60, 642, 10, 60, 3, 61, 3, 61, 
	3, 61, 3, 62, 3, 62, 3, 62, 3, 62, 5, 62, 651, 10, 62, 3, 63, 3, 63, 3, 
	63, 3, 63, 3, 64, 3, 64, 3, 64, 5, 64, 660, 10, 64, 3, 65, 3, 65, 5, 65, 
	664, 10, 65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 5, 68, 672, 10, 
	68, 3, 69, 3, 69, 5, 69, 676, 10, 69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 71, 
	3, 71, 5, 71, 684, 10, 71, 3, 72, 3, 72, 3, 73, 3, 73, 3, 74, 3, 74, 5, 
	74, 692, 10, 74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 77, 3, 77, 3, 77, 3, 77, 
	5, 77, 702, 10, 77, 3, 78, 3, 78, 3, 78, 3, 78, 5, 78, 708, 10, 78, 3, 
	78, 2, 2, 79, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 
	70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 
	106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 
	136, 138, 140, 142, 144, 146, 148, 150, 152, 154, 2, 12, 3, 2, 36, 38, 
	3, 2, 45, 47, 3, 2, 41, 44, 3, 2, 48, 49, 3, 2, 66, 68, 4, 2, 66, 66, 70, 
	70, 4, 2, 22, 22, 55, 55, 3, 2, 51, 52, 3, 2, 71, 74, 3, 2, 58, 59, 2, 
	745, 2, 159, 3, 2, 2, 2, 4, 179, 3, 2, 2, 2, 6, 181, 3, 2, 2, 2, 8, 184, 
	3, 2, 2, 2, 10, 188, 3, 2, 2, 2, 12, 193, 3, 2, 2, 2, 14, 195, 3, 2, 2, 
	2, 16, 200, 3, 2, 2, 2, 18, 206, 3, 2, 2, 2, 20, 209, 3, 2, 2, 2, 22, 222, 
	3, 2, 2, 2, 24, 224, 3, 2, 2, 2, 26, 226, 3, 2, 2, 2, 28, 234, 3, 2, 2, 
	2, 30, 242, 3, 2, 2, 2, 32, 250, 3, 2, 2, 2, 34, 259, 3, 2, 2, 2, 36, 264, 
	3, 2, 2, 2, 38, 282, 3, 2, 2, 2, 40, 298, 3, 2, 2, 2, 42, 302, 3, 2, 2, 
	2, 44, 306, 3, 2, 2, 2, 46, 312, 3, 2, 2, 2, 48, 347, 3, 2, 2, 2, 50, 349, 
	3, 2, 2, 2, 52, 374, 3, 2, 2, 2, 54, 376, 3, 2, 2, 2, 56, 389, 3, 2, 2, 
	2, 58, 393, 3, 2, 2, 2, 60, 402, 3, 2, 2, 2, 62, 404, 3, 2, 2, 2, 64, 412, 
	3, 2, 2, 2, 66, 414, 3, 2, 2, 2, 68, 416, 3, 2, 2, 2, 70, 418, 3, 2, 2, 
	2, 72, 420, 3, 2, 2, 2, 74, 436, 3, 2, 2, 2, 76, 448, 3, 2, 2, 2, 78, 450, 
	3, 2, 2, 2, 80, 456, 3, 2, 2, 2, 82, 460, 3, 2, 2, 2, 84, 462, 3, 2, 2, 
	2, 86, 471, 3, 2, 2, 2, 88, 473, 3, 2, 2, 2, 90, 477, 3, 2, 2, 2, 92, 496, 
	3, 2, 2, 2, 94, 498, 3, 2, 2, 2, 96, 517, 3, 2, 2, 2, 98, 540, 3, 2, 2, 
	2, 100, 552, 3, 2, 2, 2, 102, 554, 3, 2, 2, 2, 104, 556, 3, 2, 2, 2, 106, 
	586, 3, 2, 2, 2, 108, 588, 3, 2, 2, 2, 110, 598, 3, 2, 2, 2, 112, 603, 
	3, 2, 2, 2, 114, 613, 3, 2, 2, 2, 116, 636, 3, 2, 2, 2, 118, 638, 3, 2, 
	2, 2, 120, 643, 3, 2, 2, 2, 122, 646, 3, 2, 2, 2, 124, 652, 3, 2, 2, 2, 
	126, 659, 3, 2, 2, 2, 128, 663, 3, 2, 2, 2, 130, 665, 3, 2, 2, 2, 132, 
	667, 3, 2, 2, 2, 134, 671, 3, 2, 2, 2, 136, 675, 3, 2, 2, 2, 138, 677, 
	3, 2, 2, 2, 140, 679, 3, 2, 2, 2, 142, 685, 3, 2, 2, 2, 144, 687, 3, 2, 
	2, 2, 146, 691, 3, 2, 2, 2, 148, 693, 3, 2, 2, 2, 150, 695, 3, 2, 2, 2, 
	152, 701, 3, 2, 2, 2, 154, 707, 3, 2, 2, 2, 156, 158, 5, 4, 3, 2, 157, 
	156, 3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 159, 160, 
	3, 2, 2, 2, 160, 172, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 162, 165, 5, 12, 
	7, 2, 163, 165, 5, 16, 9, 2, 164, 162, 3, 2, 2, 2, 164, 163, 3, 2, 2, 2, 
	165, 169, 3, 2, 2, 2, 166, 168, 5, 18, 10, 2, 167, 166, 3, 2, 2, 2, 168, 
	171, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 173, 
	3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 172, 164, 3, 2, 2, 2, 172, 173, 3, 2, 
	2, 2, 173, 174, 3, 2, 2, 2, 174, 175, 7, 2, 2, 3, 175, 3, 3, 2, 2, 2, 176, 
	180, 5, 6, 4, 2, 177, 180, 5, 8, 5, 2, 178, 180, 5, 10, 6, 2, 179, 176, 
	3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 179, 178, 3, 2, 2, 2, 180, 5, 3, 2, 2, 
	2, 181, 182, 7, 25, 2, 2, 182, 183, 7, 57, 2, 2, 183, 7, 3, 2, 2, 2, 184, 
	185, 7, 30, 2, 2, 185, 186, 7, 58, 2, 2, 186, 187, 7, 57, 2, 2, 187, 9, 
	3, 2, 2, 2, 188, 189, 7, 27, 2, 2, 189, 190, 7, 57, 2, 2, 190, 11, 3, 2, 
	2, 2, 191, 194, 5, 14, 8, 2, 192, 194, 5, 20, 11, 2, 193, 191, 3, 2, 2, 
	2, 193, 192, 3, 2, 2, 2, 194, 13, 3, 2, 2, 2, 195, 196, 7, 31, 2, 2, 196, 
	197, 7, 3, 2, 2, 197, 198, 5, 22, 12, 2, 198, 15, 3, 2, 2, 2, 199, 201, 
	5, 124, 63, 2, 200, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 200, 3, 
	2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 17, 3, 2, 2, 2, 204, 207, 5, 4, 3, 
	2, 205, 207, 5, 12, 7, 2, 206, 204, 3, 2, 2, 2, 206, 205, 3, 2, 2, 2, 207, 
	19, 3, 2, 2, 2, 208, 210, 7, 24, 2, 2, 209, 208, 3, 2, 2, 2, 209, 210, 
	3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 215, 5, 134, 68, 2, 212, 214, 5, 
	154, 78, 2, 213, 212, 3, 2, 2, 2, 214, 217, 3, 2, 2, 2, 215, 213, 3, 2, 
	2, 2, 215, 216, 3, 2, 2, 2, 216, 220, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 
	218, 221, 5, 22, 12, 2, 219, 221, 7, 29, 2, 2, 220, 218, 3, 2, 2, 2, 220, 
	219, 3, 2, 2, 2, 221, 21, 3, 2, 2, 2, 222, 223, 5, 26, 14, 2, 223, 23, 
	3, 2, 2, 2, 224, 225, 5, 28, 15, 2, 225, 25, 3, 2, 2, 2, 226, 231, 5, 30, 
	16, 2, 227, 228, 7, 40, 2, 2, 228, 230, 5, 30, 16, 2, 229, 227, 3, 2, 2, 
	2, 230, 233, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 
	27, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 234, 239, 5, 32, 17, 2, 235, 236, 
	7, 40, 2, 2, 236, 238, 5, 32, 17, 2, 237, 235, 3, 2, 2, 2, 238, 241, 3, 
	2, 2, 2, 239, 237, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240, 29, 3, 2, 2, 
	2, 241, 239, 3, 2, 2, 2, 242, 247, 5, 34, 18, 2, 243, 244, 7, 39, 2, 2, 
	244, 246, 5, 34, 18, 2, 245, 243, 3, 2, 2, 2, 246, 249, 3, 2, 2, 2, 247, 
	245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 31, 3, 2, 2, 2, 249, 247, 3, 
	2, 2, 2, 250, 255, 5, 36, 19, 2, 251, 252, 7, 39, 2, 2, 252, 254, 5, 36, 
	19, 2, 253, 251, 3, 2, 2, 2, 254, 257, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 
	255, 256, 3, 2, 2, 2, 256, 33, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 258, 260, 
	7, 50, 2, 2, 259, 258, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 261, 3, 2, 
	2, 2, 261, 262, 5, 38, 20, 2, 262, 35, 3, 2, 2, 2, 263, 265, 7, 50, 2, 
	2, 264, 263, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 
	267, 5, 40, 21, 2, 267, 37, 3, 2, 2, 2, 268, 270, 5, 54, 28, 2, 269, 271, 
	5, 42, 22, 2, 270, 269, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 283, 3, 
	2, 2, 2, 272, 283, 5, 50, 26, 2, 273, 275, 5, 42, 22, 2, 274, 276, 5, 54, 
	28, 2, 275, 274, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 283, 3, 2, 2, 2, 
	277, 278, 7, 4, 2, 2, 278, 279, 5, 22, 12, 2, 279, 280, 7, 5, 2, 2, 280, 
	283, 3, 2, 2, 2, 281, 283, 7, 6, 2, 2, 282, 268, 3, 2, 2, 2, 282, 272, 
	3, 2, 2, 2, 282, 273, 3, 2, 2, 2, 282, 277, 3, 2, 2, 2, 282, 281, 3, 2, 
	2, 2, 283, 39, 3, 2, 2, 2, 284, 286, 5, 52, 27, 2, 285, 287, 5, 44, 23, 
	2, 286, 285, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 299, 3, 2, 2, 2, 288, 
	299, 5, 48, 25, 2, 289, 291, 5, 44, 23, 2, 290, 292, 5, 52, 27, 2, 291, 
	290, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 299, 3, 2, 2, 2, 293, 294, 
	7, 4, 2, 2, 294, 295, 5, 22, 12, 2, 295, 296, 7, 5, 2, 2, 296, 299, 3, 
	2, 2, 2, 297, 299, 7, 6, 2, 2, 298, 284, 3, 2, 2, 2, 298, 288, 3, 2, 2, 
	2, 298, 289, 3, 2, 2, 2, 298, 293, 3, 2, 2, 2, 298, 297, 3, 2, 2, 2, 299, 
	41, 3, 2, 2, 2, 300, 303, 5, 72, 37, 2, 301, 303, 5, 46, 24, 2, 302, 300, 
	3, 2, 2, 2, 302, 301, 3, 2, 2, 2, 303, 43, 3, 2, 2, 2, 304, 307, 5, 74, 
	38, 2, 305, 307, 5, 46, 24, 2, 306, 304, 3, 2, 2, 2, 306, 305, 3, 2, 2, 
	2, 307, 45, 3, 2, 2, 2, 308, 313, 7, 61, 2, 2, 309, 313, 7, 60, 2, 2, 310, 
	311, 7, 7, 2, 2, 311, 313, 5, 134, 68, 2, 312, 308, 3, 2, 2, 2, 312, 309, 
	3, 2, 2, 2, 312, 310, 3, 2, 2, 2, 313, 47, 3, 2, 2, 2, 314, 318, 7, 35, 
	2, 2, 315, 317, 5, 58, 30, 2, 316, 315, 3, 2, 2, 2, 317, 320, 3, 2, 2, 
	2, 318, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 348, 3, 2, 2, 2, 320, 
	318, 3, 2, 2, 2, 321, 325, 5, 56, 29, 2, 322, 324, 5, 60, 31, 2, 323, 322, 
	3, 2, 2, 2, 324, 327, 3, 2, 2, 2, 325, 323, 3, 2, 2, 2, 325, 326, 3, 2, 
	2, 2, 326, 348, 3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 328, 332, 5, 132, 67, 
	2, 329, 331, 5, 58, 30, 2, 330, 329, 3, 2, 2, 2, 331, 334, 3, 2, 2, 2, 
	332, 330, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 348, 3, 2, 2, 2, 334, 
	332, 3, 2, 2, 2, 335, 339, 5, 104, 53, 2, 336, 338, 5, 58, 30, 2, 337, 
	336, 3, 2, 2, 2, 338, 341, 3, 2, 2, 2, 339, 337, 3, 2, 2, 2, 339, 340, 
	3, 2, 2, 2, 340, 348, 3, 2, 2, 2, 341, 339, 3, 2, 2, 2, 342, 344, 5, 64, 
	33, 2, 343, 342, 3, 2, 2, 2, 344, 345, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 
	345, 346, 3, 2, 2, 2, 346, 348, 3, 2, 2, 2, 347, 314, 3, 2, 2, 2, 347, 
	321, 3, 2, 2, 2, 347, 328, 3, 2, 2, 2, 347, 335, 3, 2, 2, 2, 347, 343, 
	3, 2, 2, 2, 348, 49, 3, 2, 2, 2, 349, 353, 5, 48, 25, 2, 350, 352, 5, 122, 
	62, 2, 351, 350, 3, 2, 2, 2, 352, 355, 3, 2, 2, 2, 353, 351, 3, 2, 2, 2, 
	353, 354, 3, 2, 2, 2, 354, 359, 3, 2, 2, 2, 355, 353, 3, 2, 2, 2, 356, 
	358, 5, 124, 63, 2, 357, 356, 3, 2, 2, 2, 358, 361, 3, 2, 2, 2, 359, 357, 
	3, 2, 2, 2, 359, 360, 3, 2, 2, 2, 360, 51, 3, 2, 2, 2, 361, 359, 3, 2, 
	2, 2, 362, 366, 5, 56, 29, 2, 363, 365, 5, 60, 31, 2, 364, 363, 3, 2, 2, 
	2, 365, 368, 3, 2, 2, 2, 366, 364, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 
	375, 3, 2, 2, 2, 368, 366, 3, 2, 2, 2, 369, 371, 5, 60, 31, 2, 370, 369, 
	3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 370, 3, 2, 2, 2, 372, 373, 3, 2, 
	2, 2, 373, 375, 3, 2, 2, 2, 374, 362, 3, 2, 2, 2, 374, 370, 3, 2, 2, 2, 
	375, 53, 3, 2, 2, 2, 376, 380, 5, 52, 27, 2, 377, 379, 5, 122, 62, 2, 378, 
	377, 3, 2, 2, 2, 379, 382, 3, 2, 2, 2, 380, 378, 3, 2, 2, 2, 380, 381, 
	3, 2, 2, 2, 381, 386, 3, 2, 2, 2, 382, 380, 3, 2, 2, 2, 383, 385, 5, 124, 
	63, 2, 384, 383, 3, 2, 2, 2, 385, 388, 3, 2, 2, 2, 386, 384, 3, 2, 2, 2, 
	386, 387, 3, 2, 2, 2, 387, 55, 3, 2, 2, 2, 388, 386, 3, 2, 2, 2, 389, 390, 
	9, 2, 2, 2, 390, 57, 3, 2, 2, 2, 391, 394, 5, 60, 31, 2, 392, 394, 5, 64, 
	33, 2, 393, 391, 3, 2, 2, 2, 393, 392, 3, 2, 2, 2, 394, 59, 3, 2, 2, 2, 
	395, 396, 5, 62, 32, 2, 396, 397, 7, 66, 2, 2, 397, 403, 3, 2, 2, 2, 398, 
	400, 7, 62, 2, 2, 399, 401, 7, 63, 2, 2, 400, 399, 3, 2, 2, 2, 400, 401, 
	3, 2, 2, 2, 401, 403, 3, 2, 2, 2, 402, 395, 3, 2, 2, 2, 402, 398, 3, 2, 
	2, 2, 403, 61, 3, 2, 2, 2, 404, 405, 9, 3, 2, 2, 405, 63, 3, 2, 2, 2, 406, 
	407, 5, 66, 34, 2, 407, 408, 5, 70, 36, 2, 408, 413, 3, 2, 2, 2, 409, 410, 
	5, 68, 35, 2, 410, 411, 7, 66, 2, 2, 411, 413, 3, 2, 2, 2, 412, 406, 3, 
	2, 2, 2, 412, 409, 3, 2, 2, 2, 413, 65, 3, 2, 2, 2, 414, 415, 9, 4, 2, 
	2, 415, 67, 3, 2, 2, 2, 416, 417, 9, 5, 2, 2, 417, 69, 3, 2, 2, 2, 418, 
	419, 9, 6, 2, 2, 419, 71, 3, 2, 2, 2, 420, 424, 5, 74, 38, 2, 421, 423, 
	5, 122, 62, 2, 422, 421, 3, 2, 2, 2, 423, 426, 3, 2, 2, 2, 424, 422, 3, 
	2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 430, 3, 2, 2, 2, 426, 424, 3, 2, 2, 
	2, 427, 429, 5, 124, 63, 2, 428, 427, 3, 2, 2, 2, 429, 432, 3, 2, 2, 2, 
	430, 428, 3, 2, 2, 2, 430, 431, 3, 2, 2, 2, 431, 73, 3, 2, 2, 2, 432, 430, 
	3, 2, 2, 2, 433, 435, 5, 76, 39, 2, 434, 433, 3, 2, 2, 2, 435, 438, 3, 
	2, 2, 2, 436, 434, 3, 2, 2, 2, 436, 437, 3, 2, 2, 2, 437, 439, 3, 2, 2, 
	2, 438, 436, 3, 2, 2, 2, 439, 441, 7, 8, 2, 2, 440, 442, 5, 80, 41, 2, 
	441, 440, 3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 443, 3, 2, 2, 2, 443, 
	444, 7, 9, 2, 2, 444, 75, 3, 2, 2, 2, 445, 449, 5, 152, 77, 2, 446, 449, 
	5, 78, 40, 2, 447, 449, 7, 33, 2, 2, 448, 445, 3, 2, 2, 2, 448, 446, 3, 
	2, 2, 2, 448, 447, 3, 2, 2, 2, 449, 77, 3, 2, 2, 2, 450, 452, 7, 34, 2, 
	2, 451, 453, 5, 128, 65, 2, 452, 451, 3, 2, 2, 2, 453, 454, 3, 2, 2, 2, 
	454, 452, 3, 2, 2, 2, 454, 455, 3, 2, 2, 2, 455, 79, 3, 2, 2, 2, 456, 457, 
	5, 82, 42, 2, 457, 81, 3, 2, 2, 2, 458, 461, 5, 86, 44, 2, 459, 461, 5, 
	84, 43, 2, 460, 458, 3, 2, 2, 2, 460, 459, 3, 2, 2, 2, 461, 83, 3, 2, 2, 
	2, 462, 465, 5, 86, 44, 2, 463, 464, 7, 10, 2, 2, 464, 466, 5, 86, 44, 
	2, 465, 463, 3, 2, 2, 2, 466, 467, 3, 2, 2, 2, 467, 465, 3, 2, 2, 2, 467, 
	468, 3, 2, 2, 2, 468, 85, 3, 2, 2, 2, 469, 472, 5, 88, 45, 2, 470, 472, 
	5, 90, 46, 2, 471, 469, 3, 2, 2, 2, 471, 470, 3, 2, 2, 2, 472, 87, 3, 2, 
	2, 2, 473, 475, 5, 92, 47, 2, 474, 476, 7, 11, 2, 2, 475, 474, 3, 2, 2, 
	2, 475, 476, 3, 2, 2, 2, 476, 89, 3, 2, 2, 2, 477, 480, 5, 92, 47, 2, 478, 
	479, 7, 11, 2, 2, 479, 481, 5, 92, 47, 2, 480, 478, 3, 2, 2, 2, 481, 482, 
	3, 2, 2, 2, 482, 480, 3, 2, 2, 2, 482, 483, 3, 2, 2, 2, 483, 485, 3, 2, 
	2, 2, 484, 486, 7, 11, 2, 2, 485, 484, 3, 2, 2, 2, 485, 486, 3, 2, 2, 2, 
	486, 91, 3, 2, 2, 2, 487, 488, 7, 12, 2, 2, 488, 490, 5, 136, 69, 2, 489, 
	487, 3, 2, 2, 2, 489, 490, 3, 2, 2, 2, 490, 493, 3, 2, 2, 2, 491, 494, 
	5, 96, 49, 2, 492, 494, 5, 94, 48, 2, 493, 491, 3, 2, 2, 2, 493, 492, 3, 
	2, 2, 2, 494, 497, 3, 2, 2, 2, 495, 497, 5, 120, 61, 2, 496, 489, 3, 2, 
	2, 2, 496, 495, 3, 2, 2, 2, 497, 93, 3, 2, 2, 2, 498, 499, 7, 4, 2, 2, 
	499, 500, 5, 80, 41, 2, 500, 502, 7, 5, 2, 2, 501, 503, 5, 98, 50, 2, 502, 
	501, 3, 2, 2, 2, 502, 503, 3, 2, 2, 2, 503, 507, 3, 2, 2, 2, 504, 506, 
	5, 122, 62, 2, 505, 504, 3, 2, 2, 2, 506, 509, 3, 2, 2, 2, 507, 505, 3, 
	2, 2, 2, 507, 508, 3, 2, 2, 2, 508, 513, 3, 2, 2, 2, 509, 507, 3, 2, 2, 
	2, 510, 512, 5, 124, 63, 2, 511, 510, 3, 2, 2, 2, 512, 515, 3, 2, 2, 2, 
	513, 511, 3, 2, 2, 2, 513, 514, 3, 2, 2, 2, 514, 95, 3, 2, 2, 2, 515, 513, 
	3, 2, 2, 2, 516, 518, 5, 102, 52, 2, 517, 516, 3, 2, 2, 2, 517, 518, 3, 
	2, 2, 2, 518, 519, 3, 2, 2, 2, 519, 520, 5, 128, 65, 2, 520, 522, 5, 24, 
	13, 2, 521, 523, 5, 98, 50, 2, 522, 521, 3, 2, 2, 2, 522, 523, 3, 2, 2, 
	2, 523, 527, 3, 2, 2, 2, 524, 526, 5, 122, 62, 2, 525, 524, 3, 2, 2, 2, 
	526, 529, 3, 2, 2, 2, 527, 525, 3, 2, 2, 2, 527, 528, 3, 2, 2, 2, 528, 
	533, 3, 2, 2, 2, 529, 527, 3, 2, 2, 2, 530, 532, 5, 124, 63, 2, 531, 530, 
	3, 2, 2, 2, 532, 535, 3, 2, 2, 2, 533, 531, 3, 2, 2, 2, 533, 534, 3, 2, 
	2, 2, 534, 97, 3, 2, 2, 2, 535, 533, 3, 2, 2, 2, 536, 541, 7, 70, 2, 2, 
	537, 541, 7, 13, 2, 2, 538, 541, 7, 14, 2, 2, 539, 541, 5, 100, 51, 2, 
	540, 536, 3, 2, 2, 2, 540, 537, 3, 2, 2, 2, 540, 538, 3, 2, 2, 2, 540, 
	539, 3, 2, 2, 2, 541, 99, 3, 2, 2, 2, 542, 543, 7, 8, 2, 2, 543, 544, 7, 
	66, 2, 2, 544, 553, 7, 9, 2, 2, 545, 546, 7, 8, 2, 2, 546, 547, 7, 66, 
	2, 2, 547, 549, 7, 15, 2, 2, 548, 550, 9, 7, 2, 2, 549, 548, 3, 2, 2, 2, 
	549, 550, 3, 2, 2, 2, 550, 551, 3, 2, 2, 2, 551, 553, 7, 9, 2, 2, 552, 
	542, 3, 2, 2, 2, 552, 545, 3, 2, 2, 2, 553, 101, 3, 2, 2, 2, 554, 555, 
	7, 16, 2, 2, 555, 103, 3, 2, 2, 2, 556, 560, 7, 17, 2, 2, 557, 559, 5, 
	106, 54, 2, 558, 557, 3, 2, 2, 2, 559, 562, 3, 2, 2, 2, 560, 558, 3, 2, 
	2, 2, 560, 561, 3, 2, 2, 2, 561, 563, 3, 2, 2, 2, 562, 560, 3, 2, 2, 2, 
	563, 564, 7, 18, 2, 2, 564, 105, 3, 2, 2, 2, 565, 587, 5, 108, 55, 2, 566, 
	587, 5, 112, 57, 2, 567, 587, 5, 116, 59, 2, 568, 584, 7, 6, 2, 2, 569, 
	571, 5, 110, 56, 2, 570, 569, 3, 2, 2, 2, 571, 572, 3, 2, 2, 2, 572, 570, 
	3, 2, 2, 2, 572, 573, 3, 2, 2, 2, 573, 585, 3, 2, 2, 2, 574, 576, 5, 114, 
	58, 2, 575, 574, 3, 2, 2, 2, 576, 577, 3, 2, 2, 2, 577, 575, 3, 2, 2, 2, 
	577, 578, 3, 2, 2, 2, 578, 585, 3, 2, 2, 2, 579, 581, 5, 118, 60, 2, 580, 
	579, 3, 2, 2, 2, 581, 582, 3, 2, 2, 2, 582, 580, 3, 2, 2, 2, 582, 583, 
	3, 2, 2, 2, 583, 585, 3, 2, 2, 2, 584, 570, 3, 2, 2, 2, 584, 575, 3, 2, 
	2, 2, 584, 580, 3, 2, 2, 2, 585, 587, 3, 2, 2, 2, 586, 565, 3, 2, 2, 2, 
	586, 566, 3, 2, 2, 2, 586, 567, 3, 2, 2, 2, 586, 568, 3, 2, 2, 2, 587, 
	107, 3, 2, 2, 2, 588, 596, 5, 146, 74, 2, 589, 593, 7, 69, 2, 2, 590, 592, 
	5, 110, 56, 2, 591, 590, 3, 2, 2, 2, 592, 595, 3, 2, 2, 2, 593, 591, 3, 
	2, 2, 2, 593, 594, 3, 2, 2, 2, 594, 597, 3, 2, 2, 2, 595, 593, 3, 2, 2, 
	2, 596, 589, 3, 2, 2, 2, 596, 597, 3, 2, 2, 2, 597, 109, 3, 2, 2, 2, 598, 
	599, 7, 19, 2, 2, 599, 601, 5, 146, 74, 2, 600, 602, 7, 69, 2, 2, 601, 
	600, 3, 2, 2, 2, 601, 602, 3, 2, 2, 2, 602, 111, 3, 2, 2, 2, 603, 611, 
	5, 126, 64, 2, 604, 608, 7, 69, 2, 2, 605, 607, 5, 114, 58, 2, 606, 605, 
	3, 2, 2, 2, 607, 610, 3, 2, 2, 2, 608, 606, 3, 2, 2, 2, 608, 609, 3, 2, 
	2, 2, 609, 612, 3, 2, 2, 2, 610, 608, 3, 2, 2, 2, 611, 604, 3, 2, 2, 2, 
	611, 612, 3, 2, 2, 2, 612, 113, 3, 2, 2, 2, 613, 614, 7, 19, 2, 2, 614, 
	616, 5, 126, 64, 2, 615, 617, 7, 69, 2, 2, 616, 615, 3, 2, 2, 2, 616, 617, 
	3, 2, 2, 2, 617, 115, 3, 2, 2, 2, 618, 626, 7, 65, 2, 2, 619, 623, 7, 69, 
	2, 2, 620, 622, 5, 118, 60, 2, 621, 620, 3, 2, 2, 2, 622, 625, 3, 2, 2, 
	2, 623, 621, 3, 2, 2, 2, 623, 624, 3, 2, 2, 2, 624, 627, 3, 2, 2, 2, 625, 
	623, 3, 2, 2, 2, 626, 619, 3, 2, 2, 2, 626, 627, 3, 2, 2, 2, 627, 637, 
	3, 2, 2, 2, 628, 629, 7, 7, 2, 2, 629, 633, 7, 69, 2, 2, 630, 632, 5, 118, 
	60, 2, 631, 630, 3, 2, 2, 2, 632, 635, 3, 2, 2, 2, 633, 631, 3, 2, 2, 2, 
	633, 634, 3, 2, 2, 2, 634, 637, 3, 2, 2, 2, 635, 633, 3, 2, 2, 2, 636, 
	618, 3, 2, 2, 2, 636, 628, 3, 2, 2, 2, 637, 117, 3, 2, 2, 2, 638, 639, 
	7, 19, 2, 2, 639, 641, 7, 65, 2, 2, 640, 642, 7, 69, 2, 2, 641, 640, 3, 
	2, 2, 2, 641, 642, 3, 2, 2, 2, 642, 119, 3, 2, 2, 2, 643, 644, 7, 20, 2, 
	2, 644, 645, 5, 136, 69, 2, 645, 121, 3, 2, 2, 2, 646, 647, 7, 21, 2, 2, 
	647, 650, 5, 128, 65, 2, 648, 651, 5, 146, 74, 2, 649, 651, 5, 126, 64, 
	2, 650, 648, 3, 2, 2, 2, 650, 649, 3, 2, 2, 2, 651, 123, 3, 2, 2, 2, 652, 
	653, 7, 22, 2, 2, 653, 654, 5, 146, 74, 2, 654, 655, 9, 8, 2, 2, 655, 125, 
	3, 2, 2, 2, 656, 660, 5, 140, 71, 2, 657, 660, 5, 138, 70, 2, 658, 660, 
	5, 142, 72, 2, 659, 656, 3, 2, 2, 2, 659, 657, 3, 2, 2, 2, 659, 658, 3, 
	2, 2, 2, 660, 127, 3, 2, 2, 2, 661, 664, 5, 146, 74, 2, 662, 664, 5, 130, 
	66, 2, 663, 661, 3, 2, 2, 2, 663, 662, 3, 2, 2, 2, 664, 129, 3, 2, 2, 2, 
	665, 666, 7, 56, 2, 2, 666, 131, 3, 2, 2, 2, 667, 668, 5, 146, 74, 2, 668, 
	133, 3, 2, 2, 2, 669, 672, 5, 146, 74, 2, 670, 672, 5, 150, 76, 2, 671, 
	669, 3, 2, 2, 2, 671, 670, 3, 2, 2, 2, 672, 135, 3, 2, 2, 2, 673, 676, 
	5, 146, 74, 2, 674, 676, 5, 150, 76, 2, 675, 673, 3, 2, 2, 2, 675, 674, 
	3, 2, 2, 2, 676, 137, 3, 2, 2, 2, 677, 678, 9, 6, 2, 2, 678, 139, 3, 2, 
	2, 2, 679, 683, 5, 144, 73, 2, 680, 684, 7, 65, 2, 2, 681, 682, 7, 23, 
	2, 2, 682, 684, 5, 132, 67, 2, 683, 680, 3, 2, 2, 2, 683, 681, 3, 2, 2, 
	2, 683, 684, 3, 2, 2, 2, 684, 141, 3, 2, 2, 2, 685, 686, 9, 9, 2, 2, 686, 
	143, 3, 2, 2, 2, 687, 688, 9, 10, 2, 2, 688, 145, 3, 2, 2, 2, 689, 692, 
	7, 57, 2, 2, 690, 692, 5, 148, 75, 2, 691, 689, 3, 2, 2, 2, 691, 690, 3, 
	2, 2, 2, 692, 147, 3, 2, 2, 2, 693, 694, 9, 11, 2, 2, 694, 149, 3, 2, 2, 
	2, 695, 696, 7, 64, 2, 2, 696, 151, 3, 2, 2, 2, 697, 698, 7, 26, 2, 2, 
	698, 702, 5, 134, 68, 2, 699, 700, 7, 20, 2, 2, 700, 702, 5, 134, 68, 2, 
	701, 697, 3, 2, 2, 2, 701, 699, 3, 2, 2, 2, 702, 153, 3, 2, 2, 2, 703, 
	704, 7, 28, 2, 2, 704, 708, 5, 134, 68, 2, 705, 706, 7, 19, 2, 2, 706, 
	708, 5, 134, 68, 2, 707, 703, 3, 2, 2, 2, 707, 705, 3, 2, 2, 2, 708, 155, 
	3, 2, 2, 2, 96, 159, 164, 169, 172, 179, 193, 202, 206, 209, 215, 220, 
	231, 239, 247, 255, 259, 264, 270, 275, 282, 286, 291, 298, 302, 306, 312, 
	318, 325, 332, 339, 345, 347, 353, 359, 366, 372, 374, 380, 386, 393, 400, 
	402, 412, 424, 430, 436, 441, 448, 454, 460, 467, 471, 475, 482, 485, 489, 
	493, 496, 502, 507, 513, 517, 522, 527, 533, 540, 549, 552, 560, 572, 577, 
	582, 584, 586, 593, 596, 601, 608, 611, 616, 623, 626, 633, 636, 641, 650, 
	659, 663, 671, 675, 683, 691, 701, 707,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'('", "')'", "'.'", "'@'", "'{'", "'}'", "'|'", "';'", "'$'", 
	"'+'", "'?'", "','", "'^'", "'['", "']'", "'-'", "'&'", "'//'", "'%'", 
	"'^^'", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "", "", "", "", "", "", "", "", "", "", "'true'", "'false'", "", "", 
	"", "'a'", "", "", "", "", "", "", "", "", "", "", "", "", "'~'", "'*'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "", "", "", "KW_ABSTRACT", "KW_BASE", "KW_EXTENDS", "KW_IMPORT", "KW_RESTRICTS", 
	"KW_EXTERNAL", "KW_PREFIX", "KW_START", "KW_VIRTUAL", "KW_CLOSED", "KW_EXTRA", 
	"KW_LITERAL", "KW_IRI", "KW_NONLITERAL", "KW_BNODE", "KW_AND", "KW_OR", 
	"KW_MININCLUSIVE", "KW_MINEXCLUSIVE", "KW_MAXINCLUSIVE", "KW_MAXEXCLUSIVE", 
	"KW_LENGTH", "KW_MINLENGTH", "KW_MAXLENGTH", "KW_TOTALDIGITS", "KW_FRACTIONDIGITS", 
	"KW_NOT", "KW_TRUE", "KW_FALSE", "PASS", "COMMENT", "CODE", "RDF_TYPE", 
	"IRIREF", "PNAME_NS", "PNAME_LN", "ATPNAME_NS", "ATPNAME_LN", "REGEXP", 
	"REGEXP_FLAGS", "BLANK_NODE_LABEL", "LANGTAG", "INTEGER", "DECIMAL", "DOUBLE", 
	"STEM_MARK", "UNBOUNDED", "STRING_LITERAL1", "STRING_LITERAL2", "STRING_LITERAL_LONG1", 
	"STRING_LITERAL_LONG2",
}

var ruleNames = []string{
	"shExDoc", "directive", "baseDecl", "prefixDecl", "importDecl", "notStartAction", 
	"start", "startActions", "statement", "shapeExprDecl", "shapeExpression", 
	"inlineShapeExpression", "shapeOr", "inlineShapeOr", "shapeAnd", "inlineShapeAnd", 
	"shapeNot", "inlineShapeNot", "shapeAtom", "inlineShapeAtom", "shapeOrRef", 
	"inlineShapeOrRef", "shapeRef", "inlineLitNodeConstraint", "litNodeConstraint", 
	"inlineNonLitNodeConstraint", "nonLitNodeConstraint", "nonLiteralKind", 
	"xsFacet", "stringFacet", "stringLength", "numericFacet", "numericRange", 
	"numericLength", "rawNumeric", "shapeDefinition", "inlineShapeDefinition", 
	"qualifier", "extraPropertySet", "tripleExpression", "oneOfTripleExpr", 
	"multiElementOneOf", "groupTripleExpr", "singleElementGroup", "multiElementGroup", 
	"unaryTripleExpr", "bracketedTripleExpr", "tripleConstraint", "cardinality", 
	"repeatRange", "senseFlags", "valueSet", "valueSetValue", "iriRange", "iriExclusion", 
	"literalRange", "literalExclusion", "languageRange", "languageExclusion", 
	"include", "annotation", "semanticAction", "literal", "predicate", "rdfType", 
	"datatype", "shapeExprLabel", "tripleExprLabel", "numericLiteral", "rdfLiteral", 
	"booleanLiteral", "rdfString", "iri", "prefixedName", "blankNode", "extension", 
	"restrictions",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ShExDocParser struct {
	*antlr.BaseParser
}

func NewShExDocParser(input antlr.TokenStream) *ShExDocParser {
	this := new(ShExDocParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ShExDoc.g4"

	return this
}

// ShExDocParser tokens.
const (
	ShExDocParserEOF = antlr.TokenEOF
	ShExDocParserT__0 = 1
	ShExDocParserT__1 = 2
	ShExDocParserT__2 = 3
	ShExDocParserT__3 = 4
	ShExDocParserT__4 = 5
	ShExDocParserT__5 = 6
	ShExDocParserT__6 = 7
	ShExDocParserT__7 = 8
	ShExDocParserT__8 = 9
	ShExDocParserT__9 = 10
	ShExDocParserT__10 = 11
	ShExDocParserT__11 = 12
	ShExDocParserT__12 = 13
	ShExDocParserT__13 = 14
	ShExDocParserT__14 = 15
	ShExDocParserT__15 = 16
	ShExDocParserT__16 = 17
	ShExDocParserT__17 = 18
	ShExDocParserT__18 = 19
	ShExDocParserT__19 = 20
	ShExDocParserT__20 = 21
	ShExDocParserKW_ABSTRACT = 22
	ShExDocParserKW_BASE = 23
	ShExDocParserKW_EXTENDS = 24
	ShExDocParserKW_IMPORT = 25
	ShExDocParserKW_RESTRICTS = 26
	ShExDocParserKW_EXTERNAL = 27
	ShExDocParserKW_PREFIX = 28
	ShExDocParserKW_START = 29
	ShExDocParserKW_VIRTUAL = 30
	ShExDocParserKW_CLOSED = 31
	ShExDocParserKW_EXTRA = 32
	ShExDocParserKW_LITERAL = 33
	ShExDocParserKW_IRI = 34
	ShExDocParserKW_NONLITERAL = 35
	ShExDocParserKW_BNODE = 36
	ShExDocParserKW_AND = 37
	ShExDocParserKW_OR = 38
	ShExDocParserKW_MININCLUSIVE = 39
	ShExDocParserKW_MINEXCLUSIVE = 40
	ShExDocParserKW_MAXINCLUSIVE = 41
	ShExDocParserKW_MAXEXCLUSIVE = 42
	ShExDocParserKW_LENGTH = 43
	ShExDocParserKW_MINLENGTH = 44
	ShExDocParserKW_MAXLENGTH = 45
	ShExDocParserKW_TOTALDIGITS = 46
	ShExDocParserKW_FRACTIONDIGITS = 47
	ShExDocParserKW_NOT = 48
	ShExDocParserKW_TRUE = 49
	ShExDocParserKW_FALSE = 50
	ShExDocParserPASS = 51
	ShExDocParserCOMMENT = 52
	ShExDocParserCODE = 53
	ShExDocParserRDF_TYPE = 54
	ShExDocParserIRIREF = 55
	ShExDocParserPNAME_NS = 56
	ShExDocParserPNAME_LN = 57
	ShExDocParserATPNAME_NS = 58
	ShExDocParserATPNAME_LN = 59
	ShExDocParserREGEXP = 60
	ShExDocParserREGEXP_FLAGS = 61
	ShExDocParserBLANK_NODE_LABEL = 62
	ShExDocParserLANGTAG = 63
	ShExDocParserINTEGER = 64
	ShExDocParserDECIMAL = 65
	ShExDocParserDOUBLE = 66
	ShExDocParserSTEM_MARK = 67
	ShExDocParserUNBOUNDED = 68
	ShExDocParserSTRING_LITERAL1 = 69
	ShExDocParserSTRING_LITERAL2 = 70
	ShExDocParserSTRING_LITERAL_LONG1 = 71
	ShExDocParserSTRING_LITERAL_LONG2 = 72
)

// ShExDocParser rules.
const (
	ShExDocParserRULE_shExDoc = 0
	ShExDocParserRULE_directive = 1
	ShExDocParserRULE_baseDecl = 2
	ShExDocParserRULE_prefixDecl = 3
	ShExDocParserRULE_importDecl = 4
	ShExDocParserRULE_notStartAction = 5
	ShExDocParserRULE_start = 6
	ShExDocParserRULE_startActions = 7
	ShExDocParserRULE_statement = 8
	ShExDocParserRULE_shapeExprDecl = 9
	ShExDocParserRULE_shapeExpression = 10
	ShExDocParserRULE_inlineShapeExpression = 11
	ShExDocParserRULE_shapeOr = 12
	ShExDocParserRULE_inlineShapeOr = 13
	ShExDocParserRULE_shapeAnd = 14
	ShExDocParserRULE_inlineShapeAnd = 15
	ShExDocParserRULE_shapeNot = 16
	ShExDocParserRULE_inlineShapeNot = 17
	ShExDocParserRULE_shapeAtom = 18
	ShExDocParserRULE_inlineShapeAtom = 19
	ShExDocParserRULE_shapeOrRef = 20
	ShExDocParserRULE_inlineShapeOrRef = 21
	ShExDocParserRULE_shapeRef = 22
	ShExDocParserRULE_inlineLitNodeConstraint = 23
	ShExDocParserRULE_litNodeConstraint = 24
	ShExDocParserRULE_inlineNonLitNodeConstraint = 25
	ShExDocParserRULE_nonLitNodeConstraint = 26
	ShExDocParserRULE_nonLiteralKind = 27
	ShExDocParserRULE_xsFacet = 28
	ShExDocParserRULE_stringFacet = 29
	ShExDocParserRULE_stringLength = 30
	ShExDocParserRULE_numericFacet = 31
	ShExDocParserRULE_numericRange = 32
	ShExDocParserRULE_numericLength = 33
	ShExDocParserRULE_rawNumeric = 34
	ShExDocParserRULE_shapeDefinition = 35
	ShExDocParserRULE_inlineShapeDefinition = 36
	ShExDocParserRULE_qualifier = 37
	ShExDocParserRULE_extraPropertySet = 38
	ShExDocParserRULE_tripleExpression = 39
	ShExDocParserRULE_oneOfTripleExpr = 40
	ShExDocParserRULE_multiElementOneOf = 41
	ShExDocParserRULE_groupTripleExpr = 42
	ShExDocParserRULE_singleElementGroup = 43
	ShExDocParserRULE_multiElementGroup = 44
	ShExDocParserRULE_unaryTripleExpr = 45
	ShExDocParserRULE_bracketedTripleExpr = 46
	ShExDocParserRULE_tripleConstraint = 47
	ShExDocParserRULE_cardinality = 48
	ShExDocParserRULE_repeatRange = 49
	ShExDocParserRULE_senseFlags = 50
	ShExDocParserRULE_valueSet = 51
	ShExDocParserRULE_valueSetValue = 52
	ShExDocParserRULE_iriRange = 53
	ShExDocParserRULE_iriExclusion = 54
	ShExDocParserRULE_literalRange = 55
	ShExDocParserRULE_literalExclusion = 56
	ShExDocParserRULE_languageRange = 57
	ShExDocParserRULE_languageExclusion = 58
	ShExDocParserRULE_include = 59
	ShExDocParserRULE_annotation = 60
	ShExDocParserRULE_semanticAction = 61
	ShExDocParserRULE_literal = 62
	ShExDocParserRULE_predicate = 63
	ShExDocParserRULE_rdfType = 64
	ShExDocParserRULE_datatype = 65
	ShExDocParserRULE_shapeExprLabel = 66
	ShExDocParserRULE_tripleExprLabel = 67
	ShExDocParserRULE_numericLiteral = 68
	ShExDocParserRULE_rdfLiteral = 69
	ShExDocParserRULE_booleanLiteral = 70
	ShExDocParserRULE_rdfString = 71
	ShExDocParserRULE_iri = 72
	ShExDocParserRULE_prefixedName = 73
	ShExDocParserRULE_blankNode = 74
	ShExDocParserRULE_extension = 75
	ShExDocParserRULE_restrictions = 76
)

// IShExDocContext is an interface to support dynamic dispatch.
type IShExDocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShExDocContext differentiates from other interfaces.
	IsShExDocContext()
}

type ShExDocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShExDocContext() *ShExDocContext {
	var p = new(ShExDocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_shExDoc
	return p
}

func (*ShExDocContext) IsShExDocContext() {}

func NewShExDocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShExDocContext {
	var p = new(ShExDocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_shExDoc

	return p
}

func (s *ShExDocContext) GetParser() antlr.Parser { return s.parser }

func (s *ShExDocContext) EOF() antlr.TerminalNode {
	return s.GetToken(ShExDocParserEOF, 0)
}

func (s *ShExDocContext) AllDirective() []IDirectiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDirectiveContext)(nil)).Elem())
	var tst = make([]IDirectiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDirectiveContext)
		}
	}

	return tst
}

func (s *ShExDocContext) Directive(i int) IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *ShExDocContext) NotStartAction() INotStartActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotStartActionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotStartActionContext)
}

func (s *ShExDocContext) StartActions() IStartActionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStartActionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStartActionsContext)
}

func (s *ShExDocContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ShExDocContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ShExDocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShExDocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ShExDocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShExDoc(s)
	}
}

func (s *ShExDocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShExDoc(s)
	}
}




func (p *ShExDocParser) ShExDoc() (localctx IShExDocContext) {
	localctx = NewShExDocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ShExDocParserRULE_shExDoc)
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
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ShExDocParserKW_BASE) | (1 << ShExDocParserKW_IMPORT) | (1 << ShExDocParserKW_PREFIX))) != 0) {
		{
			p.SetState(154)
			p.Directive()
		}


		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ShExDocParserT__19) | (1 << ShExDocParserKW_ABSTRACT) | (1 << ShExDocParserKW_START))) != 0) || ((((_la - 55)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 55))) & ((1 << (ShExDocParserIRIREF - 55)) | (1 << (ShExDocParserPNAME_NS - 55)) | (1 << (ShExDocParserPNAME_LN - 55)) | (1 << (ShExDocParserBLANK_NODE_LABEL - 55)))) != 0) {
		p.SetState(162)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ShExDocParserKW_ABSTRACT, ShExDocParserKW_START, ShExDocParserIRIREF, ShExDocParserPNAME_NS, ShExDocParserPNAME_LN, ShExDocParserBLANK_NODE_LABEL:
			{
				p.SetState(160)
				p.NotStartAction()
			}


		case ShExDocParserT__19:
			{
				p.SetState(161)
				p.StartActions()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ShExDocParserKW_ABSTRACT) | (1 << ShExDocParserKW_BASE) | (1 << ShExDocParserKW_IMPORT) | (1 << ShExDocParserKW_PREFIX) | (1 << ShExDocParserKW_START))) != 0) || ((((_la - 55)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 55))) & ((1 << (ShExDocParserIRIREF - 55)) | (1 << (ShExDocParserPNAME_NS - 55)) | (1 << (ShExDocParserPNAME_LN - 55)) | (1 << (ShExDocParserBLANK_NODE_LABEL - 55)))) != 0) {
			{
				p.SetState(164)
				p.Statement()
			}


			p.SetState(169)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(172)
		p.Match(ShExDocParserEOF)
	}



	return localctx
}


// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) BaseDecl() IBaseDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseDeclContext)
}

func (s *DirectiveContext) PrefixDecl() IPrefixDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixDeclContext)
}

func (s *DirectiveContext) ImportDecl() IImportDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportDeclContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitDirective(s)
	}
}




func (p *ShExDocParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ShExDocParserRULE_directive)

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

	p.SetState(177)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserKW_BASE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(174)
			p.BaseDecl()
		}


	case ShExDocParserKW_PREFIX:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.PrefixDecl()
		}


	case ShExDocParserKW_IMPORT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(176)
			p.ImportDecl()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IBaseDeclContext is an interface to support dynamic dispatch.
type IBaseDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseDeclContext differentiates from other interfaces.
	IsBaseDeclContext()
}

type BaseDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseDeclContext() *BaseDeclContext {
	var p = new(BaseDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_baseDecl
	return p
}

func (*BaseDeclContext) IsBaseDeclContext() {}

func NewBaseDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseDeclContext {
	var p = new(BaseDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_baseDecl

	return p
}

func (s *BaseDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseDeclContext) KW_BASE() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_BASE, 0)
}

func (s *BaseDeclContext) IRIREF() antlr.TerminalNode {
	return s.GetToken(ShExDocParserIRIREF, 0)
}

func (s *BaseDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BaseDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterBaseDecl(s)
	}
}

func (s *BaseDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitBaseDecl(s)
	}
}




func (p *ShExDocParser) BaseDecl() (localctx IBaseDeclContext) {
	localctx = NewBaseDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ShExDocParserRULE_baseDecl)

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
		p.SetState(179)
		p.Match(ShExDocParserKW_BASE)
	}
	{
		p.SetState(180)
		p.Match(ShExDocParserIRIREF)
	}



	return localctx
}


// IPrefixDeclContext is an interface to support dynamic dispatch.
type IPrefixDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixDeclContext differentiates from other interfaces.
	IsPrefixDeclContext()
}

type PrefixDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixDeclContext() *PrefixDeclContext {
	var p = new(PrefixDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_prefixDecl
	return p
}

func (*PrefixDeclContext) IsPrefixDeclContext() {}

func NewPrefixDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixDeclContext {
	var p = new(PrefixDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_prefixDecl

	return p
}

func (s *PrefixDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixDeclContext) KW_PREFIX() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_PREFIX, 0)
}

func (s *PrefixDeclContext) PNAME_NS() antlr.TerminalNode {
	return s.GetToken(ShExDocParserPNAME_NS, 0)
}

func (s *PrefixDeclContext) IRIREF() antlr.TerminalNode {
	return s.GetToken(ShExDocParserIRIREF, 0)
}

func (s *PrefixDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PrefixDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterPrefixDecl(s)
	}
}

func (s *PrefixDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitPrefixDecl(s)
	}
}




func (p *ShExDocParser) PrefixDecl() (localctx IPrefixDeclContext) {
	localctx = NewPrefixDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ShExDocParserRULE_prefixDecl)

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
		p.SetState(182)
		p.Match(ShExDocParserKW_PREFIX)
	}
	{
		p.SetState(183)
		p.Match(ShExDocParserPNAME_NS)
	}
	{
		p.SetState(184)
		p.Match(ShExDocParserIRIREF)
	}



	return localctx
}


// IImportDeclContext is an interface to support dynamic dispatch.
type IImportDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclContext differentiates from other interfaces.
	IsImportDeclContext()
}

type ImportDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclContext() *ImportDeclContext {
	var p = new(ImportDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_importDecl
	return p
}

func (*ImportDeclContext) IsImportDeclContext() {}

func NewImportDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclContext {
	var p = new(ImportDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_importDecl

	return p
}

func (s *ImportDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclContext) KW_IMPORT() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_IMPORT, 0)
}

func (s *ImportDeclContext) IRIREF() antlr.TerminalNode {
	return s.GetToken(ShExDocParserIRIREF, 0)
}

func (s *ImportDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ImportDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterImportDecl(s)
	}
}

func (s *ImportDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitImportDecl(s)
	}
}




func (p *ShExDocParser) ImportDecl() (localctx IImportDeclContext) {
	localctx = NewImportDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ShExDocParserRULE_importDecl)

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
		p.SetState(186)
		p.Match(ShExDocParserKW_IMPORT)
	}
	{
		p.SetState(187)
		p.Match(ShExDocParserIRIREF)
	}



	return localctx
}


// INotStartActionContext is an interface to support dynamic dispatch.
type INotStartActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotStartActionContext differentiates from other interfaces.
	IsNotStartActionContext()
}

type NotStartActionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotStartActionContext() *NotStartActionContext {
	var p = new(NotStartActionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_notStartAction
	return p
}

func (*NotStartActionContext) IsNotStartActionContext() {}

func NewNotStartActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotStartActionContext {
	var p = new(NotStartActionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_notStartAction

	return p
}

func (s *NotStartActionContext) GetParser() antlr.Parser { return s.parser }

func (s *NotStartActionContext) Start() IStartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStartContext)
}

func (s *NotStartActionContext) ShapeExprDecl() IShapeExprDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeExprDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeExprDeclContext)
}

func (s *NotStartActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotStartActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NotStartActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterNotStartAction(s)
	}
}

func (s *NotStartActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitNotStartAction(s)
	}
}




func (p *ShExDocParser) NotStartAction() (localctx INotStartActionContext) {
	localctx = NewNotStartActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ShExDocParserRULE_notStartAction)

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

	p.SetState(191)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserKW_START:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Start()
		}


	case ShExDocParserKW_ABSTRACT, ShExDocParserIRIREF, ShExDocParserPNAME_NS, ShExDocParserPNAME_LN, ShExDocParserBLANK_NODE_LABEL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.ShapeExprDecl()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) KW_START() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_START, 0)
}

func (s *StartContext) ShapeExpression() IShapeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeExpressionContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitStart(s)
	}
}




func (p *ShExDocParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ShExDocParserRULE_start)

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
		p.SetState(193)
		p.Match(ShExDocParserKW_START)
	}
	{
		p.SetState(194)
		p.Match(ShExDocParserT__0)
	}
	{
		p.SetState(195)
		p.ShapeExpression()
	}



	return localctx
}


// IStartActionsContext is an interface to support dynamic dispatch.
type IStartActionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartActionsContext differentiates from other interfaces.
	IsStartActionsContext()
}

type StartActionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartActionsContext() *StartActionsContext {
	var p = new(StartActionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_startActions
	return p
}

func (*StartActionsContext) IsStartActionsContext() {}

func NewStartActionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartActionsContext {
	var p = new(StartActionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_startActions

	return p
}

func (s *StartActionsContext) GetParser() antlr.Parser { return s.parser }

func (s *StartActionsContext) AllSemanticAction() []ISemanticActionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISemanticActionContext)(nil)).Elem())
	var tst = make([]ISemanticActionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISemanticActionContext)
		}
	}

	return tst
}

func (s *StartActionsContext) SemanticAction(i int) ISemanticActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISemanticActionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISemanticActionContext)
}

func (s *StartActionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartActionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StartActionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterStartActions(s)
	}
}

func (s *StartActionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitStartActions(s)
	}
}




func (p *ShExDocParser) StartActions() (localctx IStartActionsContext) {
	localctx = NewStartActionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ShExDocParserRULE_startActions)
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
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == ShExDocParserT__19 {
		{
			p.SetState(197)
			p.SemanticAction()
		}


		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = ShExDocParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Directive() IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *StatementContext) NotStartAction() INotStartActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotStartActionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotStartActionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitStatement(s)
	}
}




func (p *ShExDocParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ShExDocParserRULE_statement)

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

	p.SetState(204)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserKW_BASE, ShExDocParserKW_IMPORT, ShExDocParserKW_PREFIX:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.Directive()
		}


	case ShExDocParserKW_ABSTRACT, ShExDocParserKW_START, ShExDocParserIRIREF, ShExDocParserPNAME_NS, ShExDocParserPNAME_LN, ShExDocParserBLANK_NODE_LABEL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
			p.NotStartAction()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IShapeExprDeclContext is an interface to support dynamic dispatch.
type IShapeExprDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShapeExprDeclContext differentiates from other interfaces.
	IsShapeExprDeclContext()
}

type ShapeExprDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShapeExprDeclContext() *ShapeExprDeclContext {
	var p = new(ShapeExprDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_shapeExprDecl
	return p
}

func (*ShapeExprDeclContext) IsShapeExprDeclContext() {}

func NewShapeExprDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShapeExprDeclContext {
	var p = new(ShapeExprDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_shapeExprDecl

	return p
}

func (s *ShapeExprDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ShapeExprDeclContext) ShapeExprLabel() IShapeExprLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeExprLabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeExprLabelContext)
}

func (s *ShapeExprDeclContext) ShapeExpression() IShapeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeExpressionContext)
}

func (s *ShapeExprDeclContext) KW_EXTERNAL() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_EXTERNAL, 0)
}

func (s *ShapeExprDeclContext) KW_ABSTRACT() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_ABSTRACT, 0)
}

func (s *ShapeExprDeclContext) AllRestrictions() []IRestrictionsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRestrictionsContext)(nil)).Elem())
	var tst = make([]IRestrictionsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRestrictionsContext)
		}
	}

	return tst
}

func (s *ShapeExprDeclContext) Restrictions(i int) IRestrictionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRestrictionsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRestrictionsContext)
}

func (s *ShapeExprDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeExprDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ShapeExprDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeExprDecl(s)
	}
}

func (s *ShapeExprDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeExprDecl(s)
	}
}




func (p *ShExDocParser) ShapeExprDecl() (localctx IShapeExprDeclContext) {
	localctx = NewShapeExprDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ShExDocParserRULE_shapeExprDecl)
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
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ShExDocParserKW_ABSTRACT {
		{
			p.SetState(206)
			p.Match(ShExDocParserKW_ABSTRACT)
		}

	}
	{
		p.SetState(209)
		p.ShapeExprLabel()
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserT__16 || _la == ShExDocParserKW_RESTRICTS {
		{
			p.SetState(210)
			p.Restrictions()
		}


		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserT__1, ShExDocParserT__3, ShExDocParserT__4, ShExDocParserT__5, ShExDocParserT__14, ShExDocParserT__17, ShExDocParserKW_EXTENDS, ShExDocParserKW_CLOSED, ShExDocParserKW_EXTRA, ShExDocParserKW_LITERAL, ShExDocParserKW_IRI, ShExDocParserKW_NONLITERAL, ShExDocParserKW_BNODE, ShExDocParserKW_MININCLUSIVE, ShExDocParserKW_MINEXCLUSIVE, ShExDocParserKW_MAXINCLUSIVE, ShExDocParserKW_MAXEXCLUSIVE, ShExDocParserKW_LENGTH, ShExDocParserKW_MINLENGTH, ShExDocParserKW_MAXLENGTH, ShExDocParserKW_TOTALDIGITS, ShExDocParserKW_FRACTIONDIGITS, ShExDocParserKW_NOT, ShExDocParserIRIREF, ShExDocParserPNAME_NS, ShExDocParserPNAME_LN, ShExDocParserATPNAME_NS, ShExDocParserATPNAME_LN, ShExDocParserREGEXP:
		{
			p.SetState(216)
			p.ShapeExpression()
		}


	case ShExDocParserKW_EXTERNAL:
		{
			p.SetState(217)
			p.Match(ShExDocParserKW_EXTERNAL)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IShapeExpressionContext is an interface to support dynamic dispatch.
type IShapeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShapeExpressionContext differentiates from other interfaces.
	IsShapeExpressionContext()
}

type ShapeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShapeExpressionContext() *ShapeExpressionContext {
	var p = new(ShapeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_shapeExpression
	return p
}

func (*ShapeExpressionContext) IsShapeExpressionContext() {}

func NewShapeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShapeExpressionContext {
	var p = new(ShapeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_shapeExpression

	return p
}

func (s *ShapeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ShapeExpressionContext) ShapeOr() IShapeOrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeOrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeOrContext)
}

func (s *ShapeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ShapeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeExpression(s)
	}
}

func (s *ShapeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeExpression(s)
	}
}




func (p *ShExDocParser) ShapeExpression() (localctx IShapeExpressionContext) {
	localctx = NewShapeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ShExDocParserRULE_shapeExpression)

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
		p.SetState(220)
		p.ShapeOr()
	}



	return localctx
}


// IInlineShapeExpressionContext is an interface to support dynamic dispatch.
type IInlineShapeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineShapeExpressionContext differentiates from other interfaces.
	IsInlineShapeExpressionContext()
}

type InlineShapeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineShapeExpressionContext() *InlineShapeExpressionContext {
	var p = new(InlineShapeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_inlineShapeExpression
	return p
}

func (*InlineShapeExpressionContext) IsInlineShapeExpressionContext() {}

func NewInlineShapeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineShapeExpressionContext {
	var p = new(InlineShapeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_inlineShapeExpression

	return p
}

func (s *InlineShapeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineShapeExpressionContext) InlineShapeOr() IInlineShapeOrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineShapeOrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineShapeOrContext)
}

func (s *InlineShapeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineShapeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InlineShapeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterInlineShapeExpression(s)
	}
}

func (s *InlineShapeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitInlineShapeExpression(s)
	}
}




func (p *ShExDocParser) InlineShapeExpression() (localctx IInlineShapeExpressionContext) {
	localctx = NewInlineShapeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ShExDocParserRULE_inlineShapeExpression)

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
		p.InlineShapeOr()
	}



	return localctx
}


// IShapeOrContext is an interface to support dynamic dispatch.
type IShapeOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShapeOrContext differentiates from other interfaces.
	IsShapeOrContext()
}

type ShapeOrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShapeOrContext() *ShapeOrContext {
	var p = new(ShapeOrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_shapeOr
	return p
}

func (*ShapeOrContext) IsShapeOrContext() {}

func NewShapeOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShapeOrContext {
	var p = new(ShapeOrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_shapeOr

	return p
}

func (s *ShapeOrContext) GetParser() antlr.Parser { return s.parser }

func (s *ShapeOrContext) AllShapeAnd() []IShapeAndContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IShapeAndContext)(nil)).Elem())
	var tst = make([]IShapeAndContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IShapeAndContext)
		}
	}

	return tst
}

func (s *ShapeOrContext) ShapeAnd(i int) IShapeAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeAndContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IShapeAndContext)
}

func (s *ShapeOrContext) AllKW_OR() []antlr.TerminalNode {
	return s.GetTokens(ShExDocParserKW_OR)
}

func (s *ShapeOrContext) KW_OR(i int) antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_OR, i)
}

func (s *ShapeOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ShapeOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeOr(s)
	}
}

func (s *ShapeOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeOr(s)
	}
}




func (p *ShExDocParser) ShapeOr() (localctx IShapeOrContext) {
	localctx = NewShapeOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ShExDocParserRULE_shapeOr)
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
		p.SetState(224)
		p.ShapeAnd()
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserKW_OR {
		{
			p.SetState(225)
			p.Match(ShExDocParserKW_OR)
		}
		{
			p.SetState(226)
			p.ShapeAnd()
		}


		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IInlineShapeOrContext is an interface to support dynamic dispatch.
type IInlineShapeOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineShapeOrContext differentiates from other interfaces.
	IsInlineShapeOrContext()
}

type InlineShapeOrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineShapeOrContext() *InlineShapeOrContext {
	var p = new(InlineShapeOrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_inlineShapeOr
	return p
}

func (*InlineShapeOrContext) IsInlineShapeOrContext() {}

func NewInlineShapeOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineShapeOrContext {
	var p = new(InlineShapeOrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_inlineShapeOr

	return p
}

func (s *InlineShapeOrContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineShapeOrContext) AllInlineShapeAnd() []IInlineShapeAndContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInlineShapeAndContext)(nil)).Elem())
	var tst = make([]IInlineShapeAndContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInlineShapeAndContext)
		}
	}

	return tst
}

func (s *InlineShapeOrContext) InlineShapeAnd(i int) IInlineShapeAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineShapeAndContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInlineShapeAndContext)
}

func (s *InlineShapeOrContext) AllKW_OR() []antlr.TerminalNode {
	return s.GetTokens(ShExDocParserKW_OR)
}

func (s *InlineShapeOrContext) KW_OR(i int) antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_OR, i)
}

func (s *InlineShapeOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineShapeOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InlineShapeOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterInlineShapeOr(s)
	}
}

func (s *InlineShapeOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitInlineShapeOr(s)
	}
}




func (p *ShExDocParser) InlineShapeOr() (localctx IInlineShapeOrContext) {
	localctx = NewInlineShapeOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ShExDocParserRULE_inlineShapeOr)
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
		p.SetState(232)
		p.InlineShapeAnd()
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserKW_OR {
		{
			p.SetState(233)
			p.Match(ShExDocParserKW_OR)
		}
		{
			p.SetState(234)
			p.InlineShapeAnd()
		}


		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IShapeAndContext is an interface to support dynamic dispatch.
type IShapeAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShapeAndContext differentiates from other interfaces.
	IsShapeAndContext()
}

type ShapeAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShapeAndContext() *ShapeAndContext {
	var p = new(ShapeAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_shapeAnd
	return p
}

func (*ShapeAndContext) IsShapeAndContext() {}

func NewShapeAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShapeAndContext {
	var p = new(ShapeAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_shapeAnd

	return p
}

func (s *ShapeAndContext) GetParser() antlr.Parser { return s.parser }

func (s *ShapeAndContext) AllShapeNot() []IShapeNotContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IShapeNotContext)(nil)).Elem())
	var tst = make([]IShapeNotContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IShapeNotContext)
		}
	}

	return tst
}

func (s *ShapeAndContext) ShapeNot(i int) IShapeNotContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeNotContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IShapeNotContext)
}

func (s *ShapeAndContext) AllKW_AND() []antlr.TerminalNode {
	return s.GetTokens(ShExDocParserKW_AND)
}

func (s *ShapeAndContext) KW_AND(i int) antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_AND, i)
}

func (s *ShapeAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ShapeAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeAnd(s)
	}
}

func (s *ShapeAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeAnd(s)
	}
}




func (p *ShExDocParser) ShapeAnd() (localctx IShapeAndContext) {
	localctx = NewShapeAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ShExDocParserRULE_shapeAnd)
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
		p.SetState(240)
		p.ShapeNot()
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserKW_AND {
		{
			p.SetState(241)
			p.Match(ShExDocParserKW_AND)
		}
		{
			p.SetState(242)
			p.ShapeNot()
		}


		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IInlineShapeAndContext is an interface to support dynamic dispatch.
type IInlineShapeAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineShapeAndContext differentiates from other interfaces.
	IsInlineShapeAndContext()
}

type InlineShapeAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineShapeAndContext() *InlineShapeAndContext {
	var p = new(InlineShapeAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_inlineShapeAnd
	return p
}

func (*InlineShapeAndContext) IsInlineShapeAndContext() {}

func NewInlineShapeAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineShapeAndContext {
	var p = new(InlineShapeAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_inlineShapeAnd

	return p
}

func (s *InlineShapeAndContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineShapeAndContext) AllInlineShapeNot() []IInlineShapeNotContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInlineShapeNotContext)(nil)).Elem())
	var tst = make([]IInlineShapeNotContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInlineShapeNotContext)
		}
	}

	return tst
}

func (s *InlineShapeAndContext) InlineShapeNot(i int) IInlineShapeNotContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineShapeNotContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInlineShapeNotContext)
}

func (s *InlineShapeAndContext) AllKW_AND() []antlr.TerminalNode {
	return s.GetTokens(ShExDocParserKW_AND)
}

func (s *InlineShapeAndContext) KW_AND(i int) antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_AND, i)
}

func (s *InlineShapeAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineShapeAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InlineShapeAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterInlineShapeAnd(s)
	}
}

func (s *InlineShapeAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitInlineShapeAnd(s)
	}
}




func (p *ShExDocParser) InlineShapeAnd() (localctx IInlineShapeAndContext) {
	localctx = NewInlineShapeAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ShExDocParserRULE_inlineShapeAnd)
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
		p.SetState(248)
		p.InlineShapeNot()
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserKW_AND {
		{
			p.SetState(249)
			p.Match(ShExDocParserKW_AND)
		}
		{
			p.SetState(250)
			p.InlineShapeNot()
		}


		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IShapeNotContext is an interface to support dynamic dispatch.
type IShapeNotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShapeNotContext differentiates from other interfaces.
	IsShapeNotContext()
}

type ShapeNotContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShapeNotContext() *ShapeNotContext {
	var p = new(ShapeNotContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_shapeNot
	return p
}

func (*ShapeNotContext) IsShapeNotContext() {}

func NewShapeNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShapeNotContext {
	var p = new(ShapeNotContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_shapeNot

	return p
}

func (s *ShapeNotContext) GetParser() antlr.Parser { return s.parser }

func (s *ShapeNotContext) ShapeAtom() IShapeAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeAtomContext)
}

func (s *ShapeNotContext) KW_NOT() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_NOT, 0)
}

func (s *ShapeNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeNotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ShapeNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeNot(s)
	}
}

func (s *ShapeNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeNot(s)
	}
}




func (p *ShExDocParser) ShapeNot() (localctx IShapeNotContext) {
	localctx = NewShapeNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ShExDocParserRULE_shapeNot)
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
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ShExDocParserKW_NOT {
		{
			p.SetState(256)
			p.Match(ShExDocParserKW_NOT)
		}

	}
	{
		p.SetState(259)
		p.ShapeAtom()
	}



	return localctx
}


// IInlineShapeNotContext is an interface to support dynamic dispatch.
type IInlineShapeNotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineShapeNotContext differentiates from other interfaces.
	IsInlineShapeNotContext()
}

type InlineShapeNotContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineShapeNotContext() *InlineShapeNotContext {
	var p = new(InlineShapeNotContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_inlineShapeNot
	return p
}

func (*InlineShapeNotContext) IsInlineShapeNotContext() {}

func NewInlineShapeNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineShapeNotContext {
	var p = new(InlineShapeNotContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_inlineShapeNot

	return p
}

func (s *InlineShapeNotContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineShapeNotContext) InlineShapeAtom() IInlineShapeAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineShapeAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineShapeAtomContext)
}

func (s *InlineShapeNotContext) KW_NOT() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_NOT, 0)
}

func (s *InlineShapeNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineShapeNotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InlineShapeNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterInlineShapeNot(s)
	}
}

func (s *InlineShapeNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitInlineShapeNot(s)
	}
}




func (p *ShExDocParser) InlineShapeNot() (localctx IInlineShapeNotContext) {
	localctx = NewInlineShapeNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ShExDocParserRULE_inlineShapeNot)
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
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ShExDocParserKW_NOT {
		{
			p.SetState(261)
			p.Match(ShExDocParserKW_NOT)
		}

	}
	{
		p.SetState(264)
		p.InlineShapeAtom()
	}



	return localctx
}


// IShapeAtomContext is an interface to support dynamic dispatch.
type IShapeAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShapeAtomContext differentiates from other interfaces.
	IsShapeAtomContext()
}

type ShapeAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShapeAtomContext() *ShapeAtomContext {
	var p = new(ShapeAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_shapeAtom
	return p
}

func (*ShapeAtomContext) IsShapeAtomContext() {}

func NewShapeAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShapeAtomContext {
	var p = new(ShapeAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_shapeAtom

	return p
}

func (s *ShapeAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *ShapeAtomContext) CopyFrom(ctx *ShapeAtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ShapeAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type ShapeAtomShapeOrRefContext struct {
	*ShapeAtomContext
}

func NewShapeAtomShapeOrRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShapeAtomShapeOrRefContext {
	var p = new(ShapeAtomShapeOrRefContext)

	p.ShapeAtomContext = NewEmptyShapeAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ShapeAtomContext))

	return p
}

func (s *ShapeAtomShapeOrRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeAtomShapeOrRefContext) ShapeOrRef() IShapeOrRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeOrRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeOrRefContext)
}

func (s *ShapeAtomShapeOrRefContext) NonLitNodeConstraint() INonLitNodeConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonLitNodeConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INonLitNodeConstraintContext)
}


func (s *ShapeAtomShapeOrRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeAtomShapeOrRef(s)
	}
}

func (s *ShapeAtomShapeOrRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeAtomShapeOrRef(s)
	}
}


type ShapeAtomNonLitNodeConstraintContext struct {
	*ShapeAtomContext
}

func NewShapeAtomNonLitNodeConstraintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShapeAtomNonLitNodeConstraintContext {
	var p = new(ShapeAtomNonLitNodeConstraintContext)

	p.ShapeAtomContext = NewEmptyShapeAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ShapeAtomContext))

	return p
}

func (s *ShapeAtomNonLitNodeConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeAtomNonLitNodeConstraintContext) NonLitNodeConstraint() INonLitNodeConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonLitNodeConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INonLitNodeConstraintContext)
}

func (s *ShapeAtomNonLitNodeConstraintContext) ShapeOrRef() IShapeOrRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeOrRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeOrRefContext)
}


func (s *ShapeAtomNonLitNodeConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeAtomNonLitNodeConstraint(s)
	}
}

func (s *ShapeAtomNonLitNodeConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeAtomNonLitNodeConstraint(s)
	}
}


type ShapeAtomLitNodeConstraintContext struct {
	*ShapeAtomContext
}

func NewShapeAtomLitNodeConstraintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShapeAtomLitNodeConstraintContext {
	var p = new(ShapeAtomLitNodeConstraintContext)

	p.ShapeAtomContext = NewEmptyShapeAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ShapeAtomContext))

	return p
}

func (s *ShapeAtomLitNodeConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeAtomLitNodeConstraintContext) LitNodeConstraint() ILitNodeConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILitNodeConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILitNodeConstraintContext)
}


func (s *ShapeAtomLitNodeConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeAtomLitNodeConstraint(s)
	}
}

func (s *ShapeAtomLitNodeConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeAtomLitNodeConstraint(s)
	}
}


type ShapeAtomShapeExpressionContext struct {
	*ShapeAtomContext
}

func NewShapeAtomShapeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShapeAtomShapeExpressionContext {
	var p = new(ShapeAtomShapeExpressionContext)

	p.ShapeAtomContext = NewEmptyShapeAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ShapeAtomContext))

	return p
}

func (s *ShapeAtomShapeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeAtomShapeExpressionContext) ShapeExpression() IShapeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeExpressionContext)
}


func (s *ShapeAtomShapeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeAtomShapeExpression(s)
	}
}

func (s *ShapeAtomShapeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeAtomShapeExpression(s)
	}
}


type ShapeAtomAnyContext struct {
	*ShapeAtomContext
}

func NewShapeAtomAnyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShapeAtomAnyContext {
	var p = new(ShapeAtomAnyContext)

	p.ShapeAtomContext = NewEmptyShapeAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ShapeAtomContext))

	return p
}

func (s *ShapeAtomAnyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeAtomAnyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeAtomAny(s)
	}
}

func (s *ShapeAtomAnyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeAtomAny(s)
	}
}



func (p *ShExDocParser) ShapeAtom() (localctx IShapeAtomContext) {
	localctx = NewShapeAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ShExDocParserRULE_shapeAtom)
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

	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewShapeAtomNonLitNodeConstraintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.NonLitNodeConstraint()
		}
		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ShExDocParserT__4) | (1 << ShExDocParserT__5) | (1 << ShExDocParserT__17) | (1 << ShExDocParserKW_EXTENDS) | (1 << ShExDocParserKW_CLOSED))) != 0) || ((((_la - 32)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 32))) & ((1 << (ShExDocParserKW_EXTRA - 32)) | (1 << (ShExDocParserATPNAME_NS - 32)) | (1 << (ShExDocParserATPNAME_LN - 32)))) != 0) {
			{
				p.SetState(267)
				p.ShapeOrRef()
			}

		}


	case 2:
		localctx = NewShapeAtomLitNodeConstraintContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)
			p.LitNodeConstraint()
		}


	case 3:
		localctx = NewShapeAtomShapeOrRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(271)
			p.ShapeOrRef()
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if ((((_la - 34)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 34))) & ((1 << (ShExDocParserKW_IRI - 34)) | (1 << (ShExDocParserKW_NONLITERAL - 34)) | (1 << (ShExDocParserKW_BNODE - 34)) | (1 << (ShExDocParserKW_LENGTH - 34)) | (1 << (ShExDocParserKW_MINLENGTH - 34)) | (1 << (ShExDocParserKW_MAXLENGTH - 34)) | (1 << (ShExDocParserREGEXP - 34)))) != 0) {
			{
				p.SetState(272)
				p.NonLitNodeConstraint()
			}

		}


	case 4:
		localctx = NewShapeAtomShapeExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(275)
			p.Match(ShExDocParserT__1)
		}
		{
			p.SetState(276)
			p.ShapeExpression()
		}
		{
			p.SetState(277)
			p.Match(ShExDocParserT__2)
		}


	case 5:
		localctx = NewShapeAtomAnyContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(279)
			p.Match(ShExDocParserT__3)
		}

	}


	return localctx
}


// IInlineShapeAtomContext is an interface to support dynamic dispatch.
type IInlineShapeAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineShapeAtomContext differentiates from other interfaces.
	IsInlineShapeAtomContext()
}

type InlineShapeAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineShapeAtomContext() *InlineShapeAtomContext {
	var p = new(InlineShapeAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_inlineShapeAtom
	return p
}

func (*InlineShapeAtomContext) IsInlineShapeAtomContext() {}

func NewInlineShapeAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineShapeAtomContext {
	var p = new(InlineShapeAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_inlineShapeAtom

	return p
}

func (s *InlineShapeAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineShapeAtomContext) CopyFrom(ctx *InlineShapeAtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *InlineShapeAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineShapeAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type InlineShapeAtomShapeExpressionContext struct {
	*InlineShapeAtomContext
}

func NewInlineShapeAtomShapeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InlineShapeAtomShapeExpressionContext {
	var p = new(InlineShapeAtomShapeExpressionContext)

	p.InlineShapeAtomContext = NewEmptyInlineShapeAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InlineShapeAtomContext))

	return p
}

func (s *InlineShapeAtomShapeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineShapeAtomShapeExpressionContext) ShapeExpression() IShapeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeExpressionContext)
}


func (s *InlineShapeAtomShapeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterInlineShapeAtomShapeExpression(s)
	}
}

func (s *InlineShapeAtomShapeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitInlineShapeAtomShapeExpression(s)
	}
}


type InlineShapeAtomLitNodeConstraintContext struct {
	*InlineShapeAtomContext
}

func NewInlineShapeAtomLitNodeConstraintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InlineShapeAtomLitNodeConstraintContext {
	var p = new(InlineShapeAtomLitNodeConstraintContext)

	p.InlineShapeAtomContext = NewEmptyInlineShapeAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InlineShapeAtomContext))

	return p
}

func (s *InlineShapeAtomLitNodeConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineShapeAtomLitNodeConstraintContext) InlineLitNodeConstraint() IInlineLitNodeConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineLitNodeConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineLitNodeConstraintContext)
}


func (s *InlineShapeAtomLitNodeConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterInlineShapeAtomLitNodeConstraint(s)
	}
}

func (s *InlineShapeAtomLitNodeConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitInlineShapeAtomLitNodeConstraint(s)
	}
}


type InlineShapeAtomShapeOrRefContext struct {
	*InlineShapeAtomContext
}

func NewInlineShapeAtomShapeOrRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InlineShapeAtomShapeOrRefContext {
	var p = new(InlineShapeAtomShapeOrRefContext)

	p.InlineShapeAtomContext = NewEmptyInlineShapeAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InlineShapeAtomContext))

	return p
}

func (s *InlineShapeAtomShapeOrRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineShapeAtomShapeOrRefContext) InlineShapeOrRef() IInlineShapeOrRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineShapeOrRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineShapeOrRefContext)
}

func (s *InlineShapeAtomShapeOrRefContext) InlineNonLitNodeConstraint() IInlineNonLitNodeConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineNonLitNodeConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineNonLitNodeConstraintContext)
}


func (s *InlineShapeAtomShapeOrRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterInlineShapeAtomShapeOrRef(s)
	}
}

func (s *InlineShapeAtomShapeOrRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitInlineShapeAtomShapeOrRef(s)
	}
}


type InlineShapeAtomAnyContext struct {
	*InlineShapeAtomContext
}

func NewInlineShapeAtomAnyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InlineShapeAtomAnyContext {
	var p = new(InlineShapeAtomAnyContext)

	p.InlineShapeAtomContext = NewEmptyInlineShapeAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InlineShapeAtomContext))

	return p
}

func (s *InlineShapeAtomAnyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineShapeAtomAnyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterInlineShapeAtomAny(s)
	}
}

func (s *InlineShapeAtomAnyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitInlineShapeAtomAny(s)
	}
}


type InlineShapeAtomNonLitNodeConstraintContext struct {
	*InlineShapeAtomContext
}

func NewInlineShapeAtomNonLitNodeConstraintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InlineShapeAtomNonLitNodeConstraintContext {
	var p = new(InlineShapeAtomNonLitNodeConstraintContext)

	p.InlineShapeAtomContext = NewEmptyInlineShapeAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InlineShapeAtomContext))

	return p
}

func (s *InlineShapeAtomNonLitNodeConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineShapeAtomNonLitNodeConstraintContext) InlineNonLitNodeConstraint() IInlineNonLitNodeConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineNonLitNodeConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineNonLitNodeConstraintContext)
}

func (s *InlineShapeAtomNonLitNodeConstraintContext) InlineShapeOrRef() IInlineShapeOrRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineShapeOrRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineShapeOrRefContext)
}


func (s *InlineShapeAtomNonLitNodeConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterInlineShapeAtomNonLitNodeConstraint(s)
	}
}

func (s *InlineShapeAtomNonLitNodeConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitInlineShapeAtomNonLitNodeConstraint(s)
	}
}



func (p *ShExDocParser) InlineShapeAtom() (localctx IInlineShapeAtomContext) {
	localctx = NewInlineShapeAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ShExDocParserRULE_inlineShapeAtom)
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

	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInlineShapeAtomNonLitNodeConstraintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(282)
			p.InlineNonLitNodeConstraint()
		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(283)
				p.InlineShapeOrRef()
			}


		}


	case 2:
		localctx = NewInlineShapeAtomLitNodeConstraintContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.InlineLitNodeConstraint()
		}


	case 3:
		localctx = NewInlineShapeAtomShapeOrRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(287)
			p.InlineShapeOrRef()
		}
		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if ((((_la - 34)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 34))) & ((1 << (ShExDocParserKW_IRI - 34)) | (1 << (ShExDocParserKW_NONLITERAL - 34)) | (1 << (ShExDocParserKW_BNODE - 34)) | (1 << (ShExDocParserKW_LENGTH - 34)) | (1 << (ShExDocParserKW_MINLENGTH - 34)) | (1 << (ShExDocParserKW_MAXLENGTH - 34)) | (1 << (ShExDocParserREGEXP - 34)))) != 0) {
			{
				p.SetState(288)
				p.InlineNonLitNodeConstraint()
			}

		}


	case 4:
		localctx = NewInlineShapeAtomShapeExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(291)
			p.Match(ShExDocParserT__1)
		}
		{
			p.SetState(292)
			p.ShapeExpression()
		}
		{
			p.SetState(293)
			p.Match(ShExDocParserT__2)
		}


	case 5:
		localctx = NewInlineShapeAtomAnyContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(295)
			p.Match(ShExDocParserT__3)
		}

	}


	return localctx
}


// IShapeOrRefContext is an interface to support dynamic dispatch.
type IShapeOrRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShapeOrRefContext differentiates from other interfaces.
	IsShapeOrRefContext()
}

type ShapeOrRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShapeOrRefContext() *ShapeOrRefContext {
	var p = new(ShapeOrRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_shapeOrRef
	return p
}

func (*ShapeOrRefContext) IsShapeOrRefContext() {}

func NewShapeOrRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShapeOrRefContext {
	var p = new(ShapeOrRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_shapeOrRef

	return p
}

func (s *ShapeOrRefContext) GetParser() antlr.Parser { return s.parser }

func (s *ShapeOrRefContext) ShapeDefinition() IShapeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeDefinitionContext)
}

func (s *ShapeOrRefContext) ShapeRef() IShapeRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeRefContext)
}

func (s *ShapeOrRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeOrRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ShapeOrRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeOrRef(s)
	}
}

func (s *ShapeOrRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeOrRef(s)
	}
}




func (p *ShExDocParser) ShapeOrRef() (localctx IShapeOrRefContext) {
	localctx = NewShapeOrRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ShExDocParserRULE_shapeOrRef)

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

	p.SetState(300)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserT__5, ShExDocParserT__17, ShExDocParserKW_EXTENDS, ShExDocParserKW_CLOSED, ShExDocParserKW_EXTRA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)
			p.ShapeDefinition()
		}


	case ShExDocParserT__4, ShExDocParserATPNAME_NS, ShExDocParserATPNAME_LN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(299)
			p.ShapeRef()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IInlineShapeOrRefContext is an interface to support dynamic dispatch.
type IInlineShapeOrRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineShapeOrRefContext differentiates from other interfaces.
	IsInlineShapeOrRefContext()
}

type InlineShapeOrRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineShapeOrRefContext() *InlineShapeOrRefContext {
	var p = new(InlineShapeOrRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_inlineShapeOrRef
	return p
}

func (*InlineShapeOrRefContext) IsInlineShapeOrRefContext() {}

func NewInlineShapeOrRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineShapeOrRefContext {
	var p = new(InlineShapeOrRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_inlineShapeOrRef

	return p
}

func (s *InlineShapeOrRefContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineShapeOrRefContext) InlineShapeDefinition() IInlineShapeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineShapeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineShapeDefinitionContext)
}

func (s *InlineShapeOrRefContext) ShapeRef() IShapeRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeRefContext)
}

func (s *InlineShapeOrRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineShapeOrRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InlineShapeOrRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterInlineShapeOrRef(s)
	}
}

func (s *InlineShapeOrRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitInlineShapeOrRef(s)
	}
}




func (p *ShExDocParser) InlineShapeOrRef() (localctx IInlineShapeOrRefContext) {
	localctx = NewInlineShapeOrRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ShExDocParserRULE_inlineShapeOrRef)

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

	p.SetState(304)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserT__5, ShExDocParserT__17, ShExDocParserKW_EXTENDS, ShExDocParserKW_CLOSED, ShExDocParserKW_EXTRA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(302)
			p.InlineShapeDefinition()
		}


	case ShExDocParserT__4, ShExDocParserATPNAME_NS, ShExDocParserATPNAME_LN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(303)
			p.ShapeRef()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IShapeRefContext is an interface to support dynamic dispatch.
type IShapeRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShapeRefContext differentiates from other interfaces.
	IsShapeRefContext()
}

type ShapeRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShapeRefContext() *ShapeRefContext {
	var p = new(ShapeRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_shapeRef
	return p
}

func (*ShapeRefContext) IsShapeRefContext() {}

func NewShapeRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShapeRefContext {
	var p = new(ShapeRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_shapeRef

	return p
}

func (s *ShapeRefContext) GetParser() antlr.Parser { return s.parser }

func (s *ShapeRefContext) ATPNAME_LN() antlr.TerminalNode {
	return s.GetToken(ShExDocParserATPNAME_LN, 0)
}

func (s *ShapeRefContext) ATPNAME_NS() antlr.TerminalNode {
	return s.GetToken(ShExDocParserATPNAME_NS, 0)
}

func (s *ShapeRefContext) ShapeExprLabel() IShapeExprLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeExprLabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeExprLabelContext)
}

func (s *ShapeRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ShapeRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeRef(s)
	}
}

func (s *ShapeRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeRef(s)
	}
}




func (p *ShExDocParser) ShapeRef() (localctx IShapeRefContext) {
	localctx = NewShapeRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ShExDocParserRULE_shapeRef)

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

	p.SetState(310)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserATPNAME_LN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(306)
			p.Match(ShExDocParserATPNAME_LN)
		}


	case ShExDocParserATPNAME_NS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(307)
			p.Match(ShExDocParserATPNAME_NS)
		}


	case ShExDocParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(308)
			p.Match(ShExDocParserT__4)
		}
		{
			p.SetState(309)
			p.ShapeExprLabel()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IInlineLitNodeConstraintContext is an interface to support dynamic dispatch.
type IInlineLitNodeConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineLitNodeConstraintContext differentiates from other interfaces.
	IsInlineLitNodeConstraintContext()
}

type InlineLitNodeConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineLitNodeConstraintContext() *InlineLitNodeConstraintContext {
	var p = new(InlineLitNodeConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_inlineLitNodeConstraint
	return p
}

func (*InlineLitNodeConstraintContext) IsInlineLitNodeConstraintContext() {}

func NewInlineLitNodeConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineLitNodeConstraintContext {
	var p = new(InlineLitNodeConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_inlineLitNodeConstraint

	return p
}

func (s *InlineLitNodeConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineLitNodeConstraintContext) CopyFrom(ctx *InlineLitNodeConstraintContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *InlineLitNodeConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineLitNodeConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type NodeConstraintNumericFacetContext struct {
	*InlineLitNodeConstraintContext
}

func NewNodeConstraintNumericFacetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NodeConstraintNumericFacetContext {
	var p = new(NodeConstraintNumericFacetContext)

	p.InlineLitNodeConstraintContext = NewEmptyInlineLitNodeConstraintContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InlineLitNodeConstraintContext))

	return p
}

func (s *NodeConstraintNumericFacetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeConstraintNumericFacetContext) AllNumericFacet() []INumericFacetContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumericFacetContext)(nil)).Elem())
	var tst = make([]INumericFacetContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumericFacetContext)
		}
	}

	return tst
}

func (s *NodeConstraintNumericFacetContext) NumericFacet(i int) INumericFacetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericFacetContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumericFacetContext)
}


func (s *NodeConstraintNumericFacetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterNodeConstraintNumericFacet(s)
	}
}

func (s *NodeConstraintNumericFacetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitNodeConstraintNumericFacet(s)
	}
}


type NodeConstraintLiteralContext struct {
	*InlineLitNodeConstraintContext
}

func NewNodeConstraintLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NodeConstraintLiteralContext {
	var p = new(NodeConstraintLiteralContext)

	p.InlineLitNodeConstraintContext = NewEmptyInlineLitNodeConstraintContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InlineLitNodeConstraintContext))

	return p
}

func (s *NodeConstraintLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeConstraintLiteralContext) KW_LITERAL() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_LITERAL, 0)
}

func (s *NodeConstraintLiteralContext) AllXsFacet() []IXsFacetContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IXsFacetContext)(nil)).Elem())
	var tst = make([]IXsFacetContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IXsFacetContext)
		}
	}

	return tst
}

func (s *NodeConstraintLiteralContext) XsFacet(i int) IXsFacetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXsFacetContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IXsFacetContext)
}


func (s *NodeConstraintLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterNodeConstraintLiteral(s)
	}
}

func (s *NodeConstraintLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitNodeConstraintLiteral(s)
	}
}


type NodeConstraintNonLiteralContext struct {
	*InlineLitNodeConstraintContext
}

func NewNodeConstraintNonLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NodeConstraintNonLiteralContext {
	var p = new(NodeConstraintNonLiteralContext)

	p.InlineLitNodeConstraintContext = NewEmptyInlineLitNodeConstraintContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InlineLitNodeConstraintContext))

	return p
}

func (s *NodeConstraintNonLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeConstraintNonLiteralContext) NonLiteralKind() INonLiteralKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonLiteralKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INonLiteralKindContext)
}

func (s *NodeConstraintNonLiteralContext) AllStringFacet() []IStringFacetContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringFacetContext)(nil)).Elem())
	var tst = make([]IStringFacetContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringFacetContext)
		}
	}

	return tst
}

func (s *NodeConstraintNonLiteralContext) StringFacet(i int) IStringFacetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringFacetContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringFacetContext)
}


func (s *NodeConstraintNonLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterNodeConstraintNonLiteral(s)
	}
}

func (s *NodeConstraintNonLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitNodeConstraintNonLiteral(s)
	}
}


type NodeConstraintDatatypeContext struct {
	*InlineLitNodeConstraintContext
}

func NewNodeConstraintDatatypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NodeConstraintDatatypeContext {
	var p = new(NodeConstraintDatatypeContext)

	p.InlineLitNodeConstraintContext = NewEmptyInlineLitNodeConstraintContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InlineLitNodeConstraintContext))

	return p
}

func (s *NodeConstraintDatatypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeConstraintDatatypeContext) Datatype() IDatatypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatatypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatatypeContext)
}

func (s *NodeConstraintDatatypeContext) AllXsFacet() []IXsFacetContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IXsFacetContext)(nil)).Elem())
	var tst = make([]IXsFacetContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IXsFacetContext)
		}
	}

	return tst
}

func (s *NodeConstraintDatatypeContext) XsFacet(i int) IXsFacetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXsFacetContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IXsFacetContext)
}


func (s *NodeConstraintDatatypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterNodeConstraintDatatype(s)
	}
}

func (s *NodeConstraintDatatypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitNodeConstraintDatatype(s)
	}
}


type NodeConstraintValueSetContext struct {
	*InlineLitNodeConstraintContext
}

func NewNodeConstraintValueSetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NodeConstraintValueSetContext {
	var p = new(NodeConstraintValueSetContext)

	p.InlineLitNodeConstraintContext = NewEmptyInlineLitNodeConstraintContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InlineLitNodeConstraintContext))

	return p
}

func (s *NodeConstraintValueSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeConstraintValueSetContext) ValueSet() IValueSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueSetContext)
}

func (s *NodeConstraintValueSetContext) AllXsFacet() []IXsFacetContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IXsFacetContext)(nil)).Elem())
	var tst = make([]IXsFacetContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IXsFacetContext)
		}
	}

	return tst
}

func (s *NodeConstraintValueSetContext) XsFacet(i int) IXsFacetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXsFacetContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IXsFacetContext)
}


func (s *NodeConstraintValueSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterNodeConstraintValueSet(s)
	}
}

func (s *NodeConstraintValueSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitNodeConstraintValueSet(s)
	}
}



func (p *ShExDocParser) InlineLitNodeConstraint() (localctx IInlineLitNodeConstraintContext) {
	localctx = NewInlineLitNodeConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ShExDocParserRULE_inlineLitNodeConstraint)
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

	p.SetState(345)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserKW_LITERAL:
		localctx = NewNodeConstraintLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(312)
			p.Match(ShExDocParserKW_LITERAL)
		}
		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ((((_la - 39)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 39))) & ((1 << (ShExDocParserKW_MININCLUSIVE - 39)) | (1 << (ShExDocParserKW_MINEXCLUSIVE - 39)) | (1 << (ShExDocParserKW_MAXINCLUSIVE - 39)) | (1 << (ShExDocParserKW_MAXEXCLUSIVE - 39)) | (1 << (ShExDocParserKW_LENGTH - 39)) | (1 << (ShExDocParserKW_MINLENGTH - 39)) | (1 << (ShExDocParserKW_MAXLENGTH - 39)) | (1 << (ShExDocParserKW_TOTALDIGITS - 39)) | (1 << (ShExDocParserKW_FRACTIONDIGITS - 39)) | (1 << (ShExDocParserREGEXP - 39)))) != 0) {
			{
				p.SetState(313)
				p.XsFacet()
			}


			p.SetState(318)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}


	case ShExDocParserKW_IRI, ShExDocParserKW_NONLITERAL, ShExDocParserKW_BNODE:
		localctx = NewNodeConstraintNonLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(319)
			p.NonLiteralKind()
		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ((((_la - 43)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 43))) & ((1 << (ShExDocParserKW_LENGTH - 43)) | (1 << (ShExDocParserKW_MINLENGTH - 43)) | (1 << (ShExDocParserKW_MAXLENGTH - 43)) | (1 << (ShExDocParserREGEXP - 43)))) != 0) {
			{
				p.SetState(320)
				p.StringFacet()
			}


			p.SetState(325)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}


	case ShExDocParserIRIREF, ShExDocParserPNAME_NS, ShExDocParserPNAME_LN:
		localctx = NewNodeConstraintDatatypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(326)
			p.Datatype()
		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ((((_la - 39)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 39))) & ((1 << (ShExDocParserKW_MININCLUSIVE - 39)) | (1 << (ShExDocParserKW_MINEXCLUSIVE - 39)) | (1 << (ShExDocParserKW_MAXINCLUSIVE - 39)) | (1 << (ShExDocParserKW_MAXEXCLUSIVE - 39)) | (1 << (ShExDocParserKW_LENGTH - 39)) | (1 << (ShExDocParserKW_MINLENGTH - 39)) | (1 << (ShExDocParserKW_MAXLENGTH - 39)) | (1 << (ShExDocParserKW_TOTALDIGITS - 39)) | (1 << (ShExDocParserKW_FRACTIONDIGITS - 39)) | (1 << (ShExDocParserREGEXP - 39)))) != 0) {
			{
				p.SetState(327)
				p.XsFacet()
			}


			p.SetState(332)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}


	case ShExDocParserT__14:
		localctx = NewNodeConstraintValueSetContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(333)
			p.ValueSet()
		}
		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ((((_la - 39)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 39))) & ((1 << (ShExDocParserKW_MININCLUSIVE - 39)) | (1 << (ShExDocParserKW_MINEXCLUSIVE - 39)) | (1 << (ShExDocParserKW_MAXINCLUSIVE - 39)) | (1 << (ShExDocParserKW_MAXEXCLUSIVE - 39)) | (1 << (ShExDocParserKW_LENGTH - 39)) | (1 << (ShExDocParserKW_MINLENGTH - 39)) | (1 << (ShExDocParserKW_MAXLENGTH - 39)) | (1 << (ShExDocParserKW_TOTALDIGITS - 39)) | (1 << (ShExDocParserKW_FRACTIONDIGITS - 39)) | (1 << (ShExDocParserREGEXP - 39)))) != 0) {
			{
				p.SetState(334)
				p.XsFacet()
			}


			p.SetState(339)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}


	case ShExDocParserKW_MININCLUSIVE, ShExDocParserKW_MINEXCLUSIVE, ShExDocParserKW_MAXINCLUSIVE, ShExDocParserKW_MAXEXCLUSIVE, ShExDocParserKW_TOTALDIGITS, ShExDocParserKW_FRACTIONDIGITS:
		localctx = NewNodeConstraintNumericFacetContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((((_la - 39)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 39))) & ((1 << (ShExDocParserKW_MININCLUSIVE - 39)) | (1 << (ShExDocParserKW_MINEXCLUSIVE - 39)) | (1 << (ShExDocParserKW_MAXINCLUSIVE - 39)) | (1 << (ShExDocParserKW_MAXEXCLUSIVE - 39)) | (1 << (ShExDocParserKW_TOTALDIGITS - 39)) | (1 << (ShExDocParserKW_FRACTIONDIGITS - 39)))) != 0) {
			{
				p.SetState(340)
				p.NumericFacet()
			}


			p.SetState(343)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ILitNodeConstraintContext is an interface to support dynamic dispatch.
type ILitNodeConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLitNodeConstraintContext differentiates from other interfaces.
	IsLitNodeConstraintContext()
}

type LitNodeConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLitNodeConstraintContext() *LitNodeConstraintContext {
	var p = new(LitNodeConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_litNodeConstraint
	return p
}

func (*LitNodeConstraintContext) IsLitNodeConstraintContext() {}

func NewLitNodeConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LitNodeConstraintContext {
	var p = new(LitNodeConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_litNodeConstraint

	return p
}

func (s *LitNodeConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *LitNodeConstraintContext) InlineLitNodeConstraint() IInlineLitNodeConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineLitNodeConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineLitNodeConstraintContext)
}

func (s *LitNodeConstraintContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *LitNodeConstraintContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *LitNodeConstraintContext) AllSemanticAction() []ISemanticActionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISemanticActionContext)(nil)).Elem())
	var tst = make([]ISemanticActionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISemanticActionContext)
		}
	}

	return tst
}

func (s *LitNodeConstraintContext) SemanticAction(i int) ISemanticActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISemanticActionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISemanticActionContext)
}

func (s *LitNodeConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitNodeConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LitNodeConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterLitNodeConstraint(s)
	}
}

func (s *LitNodeConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitLitNodeConstraint(s)
	}
}




func (p *ShExDocParser) LitNodeConstraint() (localctx ILitNodeConstraintContext) {
	localctx = NewLitNodeConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ShExDocParserRULE_litNodeConstraint)
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
		p.SetState(347)
		p.InlineLitNodeConstraint()
	}
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserT__18 {
		{
			p.SetState(348)
			p.Annotation()
		}


		p.SetState(353)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserT__19 {
		{
			p.SetState(354)
			p.SemanticAction()
		}


		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IInlineNonLitNodeConstraintContext is an interface to support dynamic dispatch.
type IInlineNonLitNodeConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineNonLitNodeConstraintContext differentiates from other interfaces.
	IsInlineNonLitNodeConstraintContext()
}

type InlineNonLitNodeConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineNonLitNodeConstraintContext() *InlineNonLitNodeConstraintContext {
	var p = new(InlineNonLitNodeConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_inlineNonLitNodeConstraint
	return p
}

func (*InlineNonLitNodeConstraintContext) IsInlineNonLitNodeConstraintContext() {}

func NewInlineNonLitNodeConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineNonLitNodeConstraintContext {
	var p = new(InlineNonLitNodeConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_inlineNonLitNodeConstraint

	return p
}

func (s *InlineNonLitNodeConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineNonLitNodeConstraintContext) CopyFrom(ctx *InlineNonLitNodeConstraintContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *InlineNonLitNodeConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineNonLitNodeConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type LitNodeConstraintStringFacetContext struct {
	*InlineNonLitNodeConstraintContext
}

func NewLitNodeConstraintStringFacetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitNodeConstraintStringFacetContext {
	var p = new(LitNodeConstraintStringFacetContext)

	p.InlineNonLitNodeConstraintContext = NewEmptyInlineNonLitNodeConstraintContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InlineNonLitNodeConstraintContext))

	return p
}

func (s *LitNodeConstraintStringFacetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitNodeConstraintStringFacetContext) AllStringFacet() []IStringFacetContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringFacetContext)(nil)).Elem())
	var tst = make([]IStringFacetContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringFacetContext)
		}
	}

	return tst
}

func (s *LitNodeConstraintStringFacetContext) StringFacet(i int) IStringFacetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringFacetContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringFacetContext)
}


func (s *LitNodeConstraintStringFacetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterLitNodeConstraintStringFacet(s)
	}
}

func (s *LitNodeConstraintStringFacetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitLitNodeConstraintStringFacet(s)
	}
}


type LitNodeConstraintLiteralContext struct {
	*InlineNonLitNodeConstraintContext
}

func NewLitNodeConstraintLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitNodeConstraintLiteralContext {
	var p = new(LitNodeConstraintLiteralContext)

	p.InlineNonLitNodeConstraintContext = NewEmptyInlineNonLitNodeConstraintContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InlineNonLitNodeConstraintContext))

	return p
}

func (s *LitNodeConstraintLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitNodeConstraintLiteralContext) NonLiteralKind() INonLiteralKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonLiteralKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INonLiteralKindContext)
}

func (s *LitNodeConstraintLiteralContext) AllStringFacet() []IStringFacetContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringFacetContext)(nil)).Elem())
	var tst = make([]IStringFacetContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringFacetContext)
		}
	}

	return tst
}

func (s *LitNodeConstraintLiteralContext) StringFacet(i int) IStringFacetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringFacetContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringFacetContext)
}


func (s *LitNodeConstraintLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterLitNodeConstraintLiteral(s)
	}
}

func (s *LitNodeConstraintLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitLitNodeConstraintLiteral(s)
	}
}



func (p *ShExDocParser) InlineNonLitNodeConstraint() (localctx IInlineNonLitNodeConstraintContext) {
	localctx = NewInlineNonLitNodeConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ShExDocParserRULE_inlineNonLitNodeConstraint)
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

	p.SetState(372)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserKW_IRI, ShExDocParserKW_NONLITERAL, ShExDocParserKW_BNODE:
		localctx = NewLitNodeConstraintLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(360)
			p.NonLiteralKind()
		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ((((_la - 43)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 43))) & ((1 << (ShExDocParserKW_LENGTH - 43)) | (1 << (ShExDocParserKW_MINLENGTH - 43)) | (1 << (ShExDocParserKW_MAXLENGTH - 43)) | (1 << (ShExDocParserREGEXP - 43)))) != 0) {
			{
				p.SetState(361)
				p.StringFacet()
			}


			p.SetState(366)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}


	case ShExDocParserKW_LENGTH, ShExDocParserKW_MINLENGTH, ShExDocParserKW_MAXLENGTH, ShExDocParserREGEXP:
		localctx = NewLitNodeConstraintStringFacetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((((_la - 43)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 43))) & ((1 << (ShExDocParserKW_LENGTH - 43)) | (1 << (ShExDocParserKW_MINLENGTH - 43)) | (1 << (ShExDocParserKW_MAXLENGTH - 43)) | (1 << (ShExDocParserREGEXP - 43)))) != 0) {
			{
				p.SetState(367)
				p.StringFacet()
			}


			p.SetState(370)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// INonLitNodeConstraintContext is an interface to support dynamic dispatch.
type INonLitNodeConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNonLitNodeConstraintContext differentiates from other interfaces.
	IsNonLitNodeConstraintContext()
}

type NonLitNodeConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonLitNodeConstraintContext() *NonLitNodeConstraintContext {
	var p = new(NonLitNodeConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_nonLitNodeConstraint
	return p
}

func (*NonLitNodeConstraintContext) IsNonLitNodeConstraintContext() {}

func NewNonLitNodeConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonLitNodeConstraintContext {
	var p = new(NonLitNodeConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_nonLitNodeConstraint

	return p
}

func (s *NonLitNodeConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *NonLitNodeConstraintContext) InlineNonLitNodeConstraint() IInlineNonLitNodeConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineNonLitNodeConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineNonLitNodeConstraintContext)
}

func (s *NonLitNodeConstraintContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *NonLitNodeConstraintContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *NonLitNodeConstraintContext) AllSemanticAction() []ISemanticActionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISemanticActionContext)(nil)).Elem())
	var tst = make([]ISemanticActionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISemanticActionContext)
		}
	}

	return tst
}

func (s *NonLitNodeConstraintContext) SemanticAction(i int) ISemanticActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISemanticActionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISemanticActionContext)
}

func (s *NonLitNodeConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonLitNodeConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NonLitNodeConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterNonLitNodeConstraint(s)
	}
}

func (s *NonLitNodeConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitNonLitNodeConstraint(s)
	}
}




func (p *ShExDocParser) NonLitNodeConstraint() (localctx INonLitNodeConstraintContext) {
	localctx = NewNonLitNodeConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ShExDocParserRULE_nonLitNodeConstraint)
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
		p.SetState(374)
		p.InlineNonLitNodeConstraint()
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserT__18 {
		{
			p.SetState(375)
			p.Annotation()
		}


		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserT__19 {
		{
			p.SetState(381)
			p.SemanticAction()
		}


		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// INonLiteralKindContext is an interface to support dynamic dispatch.
type INonLiteralKindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNonLiteralKindContext differentiates from other interfaces.
	IsNonLiteralKindContext()
}

type NonLiteralKindContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonLiteralKindContext() *NonLiteralKindContext {
	var p = new(NonLiteralKindContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_nonLiteralKind
	return p
}

func (*NonLiteralKindContext) IsNonLiteralKindContext() {}

func NewNonLiteralKindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonLiteralKindContext {
	var p = new(NonLiteralKindContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_nonLiteralKind

	return p
}

func (s *NonLiteralKindContext) GetParser() antlr.Parser { return s.parser }

func (s *NonLiteralKindContext) KW_IRI() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_IRI, 0)
}

func (s *NonLiteralKindContext) KW_BNODE() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_BNODE, 0)
}

func (s *NonLiteralKindContext) KW_NONLITERAL() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_NONLITERAL, 0)
}

func (s *NonLiteralKindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonLiteralKindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NonLiteralKindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterNonLiteralKind(s)
	}
}

func (s *NonLiteralKindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitNonLiteralKind(s)
	}
}




func (p *ShExDocParser) NonLiteralKind() (localctx INonLiteralKindContext) {
	localctx = NewNonLiteralKindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ShExDocParserRULE_nonLiteralKind)
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
		p.SetState(387)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 34)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 34))) & ((1 << (ShExDocParserKW_IRI - 34)) | (1 << (ShExDocParserKW_NONLITERAL - 34)) | (1 << (ShExDocParserKW_BNODE - 34)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IXsFacetContext is an interface to support dynamic dispatch.
type IXsFacetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXsFacetContext differentiates from other interfaces.
	IsXsFacetContext()
}

type XsFacetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXsFacetContext() *XsFacetContext {
	var p = new(XsFacetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_xsFacet
	return p
}

func (*XsFacetContext) IsXsFacetContext() {}

func NewXsFacetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XsFacetContext {
	var p = new(XsFacetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_xsFacet

	return p
}

func (s *XsFacetContext) GetParser() antlr.Parser { return s.parser }

func (s *XsFacetContext) StringFacet() IStringFacetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringFacetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringFacetContext)
}

func (s *XsFacetContext) NumericFacet() INumericFacetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericFacetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericFacetContext)
}

func (s *XsFacetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XsFacetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *XsFacetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterXsFacet(s)
	}
}

func (s *XsFacetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitXsFacet(s)
	}
}




func (p *ShExDocParser) XsFacet() (localctx IXsFacetContext) {
	localctx = NewXsFacetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ShExDocParserRULE_xsFacet)

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

	p.SetState(391)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserKW_LENGTH, ShExDocParserKW_MINLENGTH, ShExDocParserKW_MAXLENGTH, ShExDocParserREGEXP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(389)
			p.StringFacet()
		}


	case ShExDocParserKW_MININCLUSIVE, ShExDocParserKW_MINEXCLUSIVE, ShExDocParserKW_MAXINCLUSIVE, ShExDocParserKW_MAXEXCLUSIVE, ShExDocParserKW_TOTALDIGITS, ShExDocParserKW_FRACTIONDIGITS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(390)
			p.NumericFacet()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IStringFacetContext is an interface to support dynamic dispatch.
type IStringFacetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringFacetContext differentiates from other interfaces.
	IsStringFacetContext()
}

type StringFacetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringFacetContext() *StringFacetContext {
	var p = new(StringFacetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_stringFacet
	return p
}

func (*StringFacetContext) IsStringFacetContext() {}

func NewStringFacetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringFacetContext {
	var p = new(StringFacetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_stringFacet

	return p
}

func (s *StringFacetContext) GetParser() antlr.Parser { return s.parser }

func (s *StringFacetContext) StringLength() IStringLengthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLengthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLengthContext)
}

func (s *StringFacetContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ShExDocParserINTEGER, 0)
}

func (s *StringFacetContext) REGEXP() antlr.TerminalNode {
	return s.GetToken(ShExDocParserREGEXP, 0)
}

func (s *StringFacetContext) REGEXP_FLAGS() antlr.TerminalNode {
	return s.GetToken(ShExDocParserREGEXP_FLAGS, 0)
}

func (s *StringFacetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringFacetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StringFacetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterStringFacet(s)
	}
}

func (s *StringFacetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitStringFacet(s)
	}
}




func (p *ShExDocParser) StringFacet() (localctx IStringFacetContext) {
	localctx = NewStringFacetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ShExDocParserRULE_stringFacet)
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

	p.SetState(400)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserKW_LENGTH, ShExDocParserKW_MINLENGTH, ShExDocParserKW_MAXLENGTH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(393)
			p.StringLength()
		}
		{
			p.SetState(394)
			p.Match(ShExDocParserINTEGER)
		}


	case ShExDocParserREGEXP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(396)
			p.Match(ShExDocParserREGEXP)
		}
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == ShExDocParserREGEXP_FLAGS {
			{
				p.SetState(397)
				p.Match(ShExDocParserREGEXP_FLAGS)
			}

		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IStringLengthContext is an interface to support dynamic dispatch.
type IStringLengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLengthContext differentiates from other interfaces.
	IsStringLengthContext()
}

type StringLengthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLengthContext() *StringLengthContext {
	var p = new(StringLengthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_stringLength
	return p
}

func (*StringLengthContext) IsStringLengthContext() {}

func NewStringLengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLengthContext {
	var p = new(StringLengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_stringLength

	return p
}

func (s *StringLengthContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLengthContext) KW_LENGTH() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_LENGTH, 0)
}

func (s *StringLengthContext) KW_MINLENGTH() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_MINLENGTH, 0)
}

func (s *StringLengthContext) KW_MAXLENGTH() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_MAXLENGTH, 0)
}

func (s *StringLengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StringLengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterStringLength(s)
	}
}

func (s *StringLengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitStringLength(s)
	}
}




func (p *ShExDocParser) StringLength() (localctx IStringLengthContext) {
	localctx = NewStringLengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ShExDocParserRULE_stringLength)
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
		p.SetState(402)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 43)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 43))) & ((1 << (ShExDocParserKW_LENGTH - 43)) | (1 << (ShExDocParserKW_MINLENGTH - 43)) | (1 << (ShExDocParserKW_MAXLENGTH - 43)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// INumericFacetContext is an interface to support dynamic dispatch.
type INumericFacetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericFacetContext differentiates from other interfaces.
	IsNumericFacetContext()
}

type NumericFacetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericFacetContext() *NumericFacetContext {
	var p = new(NumericFacetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_numericFacet
	return p
}

func (*NumericFacetContext) IsNumericFacetContext() {}

func NewNumericFacetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericFacetContext {
	var p = new(NumericFacetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_numericFacet

	return p
}

func (s *NumericFacetContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericFacetContext) NumericRange() INumericRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericRangeContext)
}

func (s *NumericFacetContext) RawNumeric() IRawNumericContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRawNumericContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRawNumericContext)
}

func (s *NumericFacetContext) NumericLength() INumericLengthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLengthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLengthContext)
}

func (s *NumericFacetContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ShExDocParserINTEGER, 0)
}

func (s *NumericFacetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericFacetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NumericFacetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterNumericFacet(s)
	}
}

func (s *NumericFacetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitNumericFacet(s)
	}
}




func (p *ShExDocParser) NumericFacet() (localctx INumericFacetContext) {
	localctx = NewNumericFacetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ShExDocParserRULE_numericFacet)

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

	p.SetState(410)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserKW_MININCLUSIVE, ShExDocParserKW_MINEXCLUSIVE, ShExDocParserKW_MAXINCLUSIVE, ShExDocParserKW_MAXEXCLUSIVE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(404)
			p.NumericRange()
		}
		{
			p.SetState(405)
			p.RawNumeric()
		}


	case ShExDocParserKW_TOTALDIGITS, ShExDocParserKW_FRACTIONDIGITS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(407)
			p.NumericLength()
		}
		{
			p.SetState(408)
			p.Match(ShExDocParserINTEGER)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// INumericRangeContext is an interface to support dynamic dispatch.
type INumericRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericRangeContext differentiates from other interfaces.
	IsNumericRangeContext()
}

type NumericRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericRangeContext() *NumericRangeContext {
	var p = new(NumericRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_numericRange
	return p
}

func (*NumericRangeContext) IsNumericRangeContext() {}

func NewNumericRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericRangeContext {
	var p = new(NumericRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_numericRange

	return p
}

func (s *NumericRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericRangeContext) KW_MININCLUSIVE() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_MININCLUSIVE, 0)
}

func (s *NumericRangeContext) KW_MINEXCLUSIVE() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_MINEXCLUSIVE, 0)
}

func (s *NumericRangeContext) KW_MAXINCLUSIVE() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_MAXINCLUSIVE, 0)
}

func (s *NumericRangeContext) KW_MAXEXCLUSIVE() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_MAXEXCLUSIVE, 0)
}

func (s *NumericRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NumericRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterNumericRange(s)
	}
}

func (s *NumericRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitNumericRange(s)
	}
}




func (p *ShExDocParser) NumericRange() (localctx INumericRangeContext) {
	localctx = NewNumericRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ShExDocParserRULE_numericRange)
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
		p.SetState(412)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 39)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 39))) & ((1 << (ShExDocParserKW_MININCLUSIVE - 39)) | (1 << (ShExDocParserKW_MINEXCLUSIVE - 39)) | (1 << (ShExDocParserKW_MAXINCLUSIVE - 39)) | (1 << (ShExDocParserKW_MAXEXCLUSIVE - 39)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// INumericLengthContext is an interface to support dynamic dispatch.
type INumericLengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericLengthContext differentiates from other interfaces.
	IsNumericLengthContext()
}

type NumericLengthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLengthContext() *NumericLengthContext {
	var p = new(NumericLengthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_numericLength
	return p
}

func (*NumericLengthContext) IsNumericLengthContext() {}

func NewNumericLengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLengthContext {
	var p = new(NumericLengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_numericLength

	return p
}

func (s *NumericLengthContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLengthContext) KW_TOTALDIGITS() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_TOTALDIGITS, 0)
}

func (s *NumericLengthContext) KW_FRACTIONDIGITS() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_FRACTIONDIGITS, 0)
}

func (s *NumericLengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NumericLengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterNumericLength(s)
	}
}

func (s *NumericLengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitNumericLength(s)
	}
}




func (p *ShExDocParser) NumericLength() (localctx INumericLengthContext) {
	localctx = NewNumericLengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ShExDocParserRULE_numericLength)
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
		p.SetState(414)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ShExDocParserKW_TOTALDIGITS || _la == ShExDocParserKW_FRACTIONDIGITS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IRawNumericContext is an interface to support dynamic dispatch.
type IRawNumericContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRawNumericContext differentiates from other interfaces.
	IsRawNumericContext()
}

type RawNumericContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRawNumericContext() *RawNumericContext {
	var p = new(RawNumericContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_rawNumeric
	return p
}

func (*RawNumericContext) IsRawNumericContext() {}

func NewRawNumericContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RawNumericContext {
	var p = new(RawNumericContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_rawNumeric

	return p
}

func (s *RawNumericContext) GetParser() antlr.Parser { return s.parser }

func (s *RawNumericContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ShExDocParserINTEGER, 0)
}

func (s *RawNumericContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(ShExDocParserDECIMAL, 0)
}

func (s *RawNumericContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ShExDocParserDOUBLE, 0)
}

func (s *RawNumericContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawNumericContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RawNumericContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterRawNumeric(s)
	}
}

func (s *RawNumericContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitRawNumeric(s)
	}
}




func (p *ShExDocParser) RawNumeric() (localctx IRawNumericContext) {
	localctx = NewRawNumericContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ShExDocParserRULE_rawNumeric)
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
		p.SetState(416)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 64)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 64))) & ((1 << (ShExDocParserINTEGER - 64)) | (1 << (ShExDocParserDECIMAL - 64)) | (1 << (ShExDocParserDOUBLE - 64)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IShapeDefinitionContext is an interface to support dynamic dispatch.
type IShapeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShapeDefinitionContext differentiates from other interfaces.
	IsShapeDefinitionContext()
}

type ShapeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShapeDefinitionContext() *ShapeDefinitionContext {
	var p = new(ShapeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_shapeDefinition
	return p
}

func (*ShapeDefinitionContext) IsShapeDefinitionContext() {}

func NewShapeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShapeDefinitionContext {
	var p = new(ShapeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_shapeDefinition

	return p
}

func (s *ShapeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ShapeDefinitionContext) InlineShapeDefinition() IInlineShapeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineShapeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineShapeDefinitionContext)
}

func (s *ShapeDefinitionContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *ShapeDefinitionContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *ShapeDefinitionContext) AllSemanticAction() []ISemanticActionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISemanticActionContext)(nil)).Elem())
	var tst = make([]ISemanticActionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISemanticActionContext)
		}
	}

	return tst
}

func (s *ShapeDefinitionContext) SemanticAction(i int) ISemanticActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISemanticActionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISemanticActionContext)
}

func (s *ShapeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ShapeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeDefinition(s)
	}
}

func (s *ShapeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeDefinition(s)
	}
}




func (p *ShExDocParser) ShapeDefinition() (localctx IShapeDefinitionContext) {
	localctx = NewShapeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ShExDocParserRULE_shapeDefinition)
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
		p.SetState(418)
		p.InlineShapeDefinition()
	}
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserT__18 {
		{
			p.SetState(419)
			p.Annotation()
		}


		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserT__19 {
		{
			p.SetState(425)
			p.SemanticAction()
		}


		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IInlineShapeDefinitionContext is an interface to support dynamic dispatch.
type IInlineShapeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineShapeDefinitionContext differentiates from other interfaces.
	IsInlineShapeDefinitionContext()
}

type InlineShapeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineShapeDefinitionContext() *InlineShapeDefinitionContext {
	var p = new(InlineShapeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_inlineShapeDefinition
	return p
}

func (*InlineShapeDefinitionContext) IsInlineShapeDefinitionContext() {}

func NewInlineShapeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineShapeDefinitionContext {
	var p = new(InlineShapeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_inlineShapeDefinition

	return p
}

func (s *InlineShapeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineShapeDefinitionContext) AllQualifier() []IQualifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQualifierContext)(nil)).Elem())
	var tst = make([]IQualifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQualifierContext)
		}
	}

	return tst
}

func (s *InlineShapeDefinitionContext) Qualifier(i int) IQualifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQualifierContext)
}

func (s *InlineShapeDefinitionContext) TripleExpression() ITripleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITripleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITripleExpressionContext)
}

func (s *InlineShapeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineShapeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InlineShapeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterInlineShapeDefinition(s)
	}
}

func (s *InlineShapeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitInlineShapeDefinition(s)
	}
}




func (p *ShExDocParser) InlineShapeDefinition() (localctx IInlineShapeDefinitionContext) {
	localctx = NewInlineShapeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ShExDocParserRULE_inlineShapeDefinition)
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
	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ((((_la - 18)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 18))) & ((1 << (ShExDocParserT__17 - 18)) | (1 << (ShExDocParserKW_EXTENDS - 18)) | (1 << (ShExDocParserKW_CLOSED - 18)) | (1 << (ShExDocParserKW_EXTRA - 18)))) != 0) {
		{
			p.SetState(431)
			p.Qualifier()
		}


		p.SetState(436)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(437)
		p.Match(ShExDocParserT__5)
	}
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ShExDocParserT__1) | (1 << ShExDocParserT__9) | (1 << ShExDocParserT__13) | (1 << ShExDocParserT__17))) != 0) || ((((_la - 54)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 54))) & ((1 << (ShExDocParserRDF_TYPE - 54)) | (1 << (ShExDocParserIRIREF - 54)) | (1 << (ShExDocParserPNAME_NS - 54)) | (1 << (ShExDocParserPNAME_LN - 54)))) != 0) {
		{
			p.SetState(438)
			p.TripleExpression()
		}

	}
	{
		p.SetState(441)
		p.Match(ShExDocParserT__6)
	}



	return localctx
}


// IQualifierContext is an interface to support dynamic dispatch.
type IQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifierContext differentiates from other interfaces.
	IsQualifierContext()
}

type QualifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifierContext() *QualifierContext {
	var p = new(QualifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_qualifier
	return p
}

func (*QualifierContext) IsQualifierContext() {}

func NewQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifierContext {
	var p = new(QualifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_qualifier

	return p
}

func (s *QualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifierContext) Extension() IExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtensionContext)
}

func (s *QualifierContext) ExtraPropertySet() IExtraPropertySetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtraPropertySetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtraPropertySetContext)
}

func (s *QualifierContext) KW_CLOSED() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_CLOSED, 0)
}

func (s *QualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *QualifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterQualifier(s)
	}
}

func (s *QualifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitQualifier(s)
	}
}




func (p *ShExDocParser) Qualifier() (localctx IQualifierContext) {
	localctx = NewQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ShExDocParserRULE_qualifier)

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

	p.SetState(446)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserT__17, ShExDocParserKW_EXTENDS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(443)
			p.Extension()
		}


	case ShExDocParserKW_EXTRA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(444)
			p.ExtraPropertySet()
		}


	case ShExDocParserKW_CLOSED:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(445)
			p.Match(ShExDocParserKW_CLOSED)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IExtraPropertySetContext is an interface to support dynamic dispatch.
type IExtraPropertySetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtraPropertySetContext differentiates from other interfaces.
	IsExtraPropertySetContext()
}

type ExtraPropertySetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtraPropertySetContext() *ExtraPropertySetContext {
	var p = new(ExtraPropertySetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_extraPropertySet
	return p
}

func (*ExtraPropertySetContext) IsExtraPropertySetContext() {}

func NewExtraPropertySetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtraPropertySetContext {
	var p = new(ExtraPropertySetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_extraPropertySet

	return p
}

func (s *ExtraPropertySetContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtraPropertySetContext) KW_EXTRA() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_EXTRA, 0)
}

func (s *ExtraPropertySetContext) AllPredicate() []IPredicateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPredicateContext)(nil)).Elem())
	var tst = make([]IPredicateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPredicateContext)
		}
	}

	return tst
}

func (s *ExtraPropertySetContext) Predicate(i int) IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *ExtraPropertySetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtraPropertySetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExtraPropertySetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterExtraPropertySet(s)
	}
}

func (s *ExtraPropertySetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitExtraPropertySet(s)
	}
}




func (p *ShExDocParser) ExtraPropertySet() (localctx IExtraPropertySetContext) {
	localctx = NewExtraPropertySetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ShExDocParserRULE_extraPropertySet)
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
		p.SetState(448)
		p.Match(ShExDocParserKW_EXTRA)
	}
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((((_la - 54)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 54))) & ((1 << (ShExDocParserRDF_TYPE - 54)) | (1 << (ShExDocParserIRIREF - 54)) | (1 << (ShExDocParserPNAME_NS - 54)) | (1 << (ShExDocParserPNAME_LN - 54)))) != 0) {
		{
			p.SetState(449)
			p.Predicate()
		}


		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ITripleExpressionContext is an interface to support dynamic dispatch.
type ITripleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTripleExpressionContext differentiates from other interfaces.
	IsTripleExpressionContext()
}

type TripleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTripleExpressionContext() *TripleExpressionContext {
	var p = new(TripleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_tripleExpression
	return p
}

func (*TripleExpressionContext) IsTripleExpressionContext() {}

func NewTripleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TripleExpressionContext {
	var p = new(TripleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_tripleExpression

	return p
}

func (s *TripleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *TripleExpressionContext) OneOfTripleExpr() IOneOfTripleExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOneOfTripleExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOneOfTripleExprContext)
}

func (s *TripleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TripleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TripleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterTripleExpression(s)
	}
}

func (s *TripleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitTripleExpression(s)
	}
}




func (p *ShExDocParser) TripleExpression() (localctx ITripleExpressionContext) {
	localctx = NewTripleExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ShExDocParserRULE_tripleExpression)

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
		p.OneOfTripleExpr()
	}



	return localctx
}


// IOneOfTripleExprContext is an interface to support dynamic dispatch.
type IOneOfTripleExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOneOfTripleExprContext differentiates from other interfaces.
	IsOneOfTripleExprContext()
}

type OneOfTripleExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneOfTripleExprContext() *OneOfTripleExprContext {
	var p = new(OneOfTripleExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_oneOfTripleExpr
	return p
}

func (*OneOfTripleExprContext) IsOneOfTripleExprContext() {}

func NewOneOfTripleExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneOfTripleExprContext {
	var p = new(OneOfTripleExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_oneOfTripleExpr

	return p
}

func (s *OneOfTripleExprContext) GetParser() antlr.Parser { return s.parser }

func (s *OneOfTripleExprContext) GroupTripleExpr() IGroupTripleExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupTripleExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupTripleExprContext)
}

func (s *OneOfTripleExprContext) MultiElementOneOf() IMultiElementOneOfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiElementOneOfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiElementOneOfContext)
}

func (s *OneOfTripleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneOfTripleExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OneOfTripleExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterOneOfTripleExpr(s)
	}
}

func (s *OneOfTripleExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitOneOfTripleExpr(s)
	}
}




func (p *ShExDocParser) OneOfTripleExpr() (localctx IOneOfTripleExprContext) {
	localctx = NewOneOfTripleExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ShExDocParserRULE_oneOfTripleExpr)

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

	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(456)
			p.GroupTripleExpr()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(457)
			p.MultiElementOneOf()
		}

	}


	return localctx
}


// IMultiElementOneOfContext is an interface to support dynamic dispatch.
type IMultiElementOneOfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiElementOneOfContext differentiates from other interfaces.
	IsMultiElementOneOfContext()
}

type MultiElementOneOfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiElementOneOfContext() *MultiElementOneOfContext {
	var p = new(MultiElementOneOfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_multiElementOneOf
	return p
}

func (*MultiElementOneOfContext) IsMultiElementOneOfContext() {}

func NewMultiElementOneOfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiElementOneOfContext {
	var p = new(MultiElementOneOfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_multiElementOneOf

	return p
}

func (s *MultiElementOneOfContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiElementOneOfContext) AllGroupTripleExpr() []IGroupTripleExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGroupTripleExprContext)(nil)).Elem())
	var tst = make([]IGroupTripleExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGroupTripleExprContext)
		}
	}

	return tst
}

func (s *MultiElementOneOfContext) GroupTripleExpr(i int) IGroupTripleExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupTripleExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGroupTripleExprContext)
}

func (s *MultiElementOneOfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiElementOneOfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MultiElementOneOfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterMultiElementOneOf(s)
	}
}

func (s *MultiElementOneOfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitMultiElementOneOf(s)
	}
}




func (p *ShExDocParser) MultiElementOneOf() (localctx IMultiElementOneOfContext) {
	localctx = NewMultiElementOneOfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ShExDocParserRULE_multiElementOneOf)
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
		p.SetState(460)
		p.GroupTripleExpr()
	}
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == ShExDocParserT__7 {
		{
			p.SetState(461)
			p.Match(ShExDocParserT__7)
		}
		{
			p.SetState(462)
			p.GroupTripleExpr()
		}


		p.SetState(465)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IGroupTripleExprContext is an interface to support dynamic dispatch.
type IGroupTripleExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupTripleExprContext differentiates from other interfaces.
	IsGroupTripleExprContext()
}

type GroupTripleExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupTripleExprContext() *GroupTripleExprContext {
	var p = new(GroupTripleExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_groupTripleExpr
	return p
}

func (*GroupTripleExprContext) IsGroupTripleExprContext() {}

func NewGroupTripleExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupTripleExprContext {
	var p = new(GroupTripleExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_groupTripleExpr

	return p
}

func (s *GroupTripleExprContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupTripleExprContext) SingleElementGroup() ISingleElementGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleElementGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleElementGroupContext)
}

func (s *GroupTripleExprContext) MultiElementGroup() IMultiElementGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiElementGroupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiElementGroupContext)
}

func (s *GroupTripleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupTripleExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *GroupTripleExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterGroupTripleExpr(s)
	}
}

func (s *GroupTripleExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitGroupTripleExpr(s)
	}
}




func (p *ShExDocParser) GroupTripleExpr() (localctx IGroupTripleExprContext) {
	localctx = NewGroupTripleExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ShExDocParserRULE_groupTripleExpr)

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

	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(467)
			p.SingleElementGroup()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(468)
			p.MultiElementGroup()
		}

	}


	return localctx
}


// ISingleElementGroupContext is an interface to support dynamic dispatch.
type ISingleElementGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleElementGroupContext differentiates from other interfaces.
	IsSingleElementGroupContext()
}

type SingleElementGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleElementGroupContext() *SingleElementGroupContext {
	var p = new(SingleElementGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_singleElementGroup
	return p
}

func (*SingleElementGroupContext) IsSingleElementGroupContext() {}

func NewSingleElementGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleElementGroupContext {
	var p = new(SingleElementGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_singleElementGroup

	return p
}

func (s *SingleElementGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleElementGroupContext) UnaryTripleExpr() IUnaryTripleExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryTripleExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryTripleExprContext)
}

func (s *SingleElementGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleElementGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SingleElementGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterSingleElementGroup(s)
	}
}

func (s *SingleElementGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitSingleElementGroup(s)
	}
}




func (p *ShExDocParser) SingleElementGroup() (localctx ISingleElementGroupContext) {
	localctx = NewSingleElementGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ShExDocParserRULE_singleElementGroup)
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
		p.SetState(471)
		p.UnaryTripleExpr()
	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ShExDocParserT__8 {
		{
			p.SetState(472)
			p.Match(ShExDocParserT__8)
		}

	}



	return localctx
}


// IMultiElementGroupContext is an interface to support dynamic dispatch.
type IMultiElementGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiElementGroupContext differentiates from other interfaces.
	IsMultiElementGroupContext()
}

type MultiElementGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiElementGroupContext() *MultiElementGroupContext {
	var p = new(MultiElementGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_multiElementGroup
	return p
}

func (*MultiElementGroupContext) IsMultiElementGroupContext() {}

func NewMultiElementGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiElementGroupContext {
	var p = new(MultiElementGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_multiElementGroup

	return p
}

func (s *MultiElementGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiElementGroupContext) AllUnaryTripleExpr() []IUnaryTripleExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnaryTripleExprContext)(nil)).Elem())
	var tst = make([]IUnaryTripleExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnaryTripleExprContext)
		}
	}

	return tst
}

func (s *MultiElementGroupContext) UnaryTripleExpr(i int) IUnaryTripleExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryTripleExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnaryTripleExprContext)
}

func (s *MultiElementGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiElementGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MultiElementGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterMultiElementGroup(s)
	}
}

func (s *MultiElementGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitMultiElementGroup(s)
	}
}




func (p *ShExDocParser) MultiElementGroup() (localctx IMultiElementGroupContext) {
	localctx = NewMultiElementGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ShExDocParserRULE_multiElementGroup)
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
		p.SetState(475)
		p.UnaryTripleExpr()
	}
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(476)
					p.Match(ShExDocParserT__8)
				}
				{
					p.SetState(477)
					p.UnaryTripleExpr()
				}




		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(480)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext())
	}
	p.SetState(483)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ShExDocParserT__8 {
		{
			p.SetState(482)
			p.Match(ShExDocParserT__8)
		}

	}



	return localctx
}


// IUnaryTripleExprContext is an interface to support dynamic dispatch.
type IUnaryTripleExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryTripleExprContext differentiates from other interfaces.
	IsUnaryTripleExprContext()
}

type UnaryTripleExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryTripleExprContext() *UnaryTripleExprContext {
	var p = new(UnaryTripleExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_unaryTripleExpr
	return p
}

func (*UnaryTripleExprContext) IsUnaryTripleExprContext() {}

func NewUnaryTripleExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryTripleExprContext {
	var p = new(UnaryTripleExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_unaryTripleExpr

	return p
}

func (s *UnaryTripleExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryTripleExprContext) TripleConstraint() ITripleConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITripleConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITripleConstraintContext)
}

func (s *UnaryTripleExprContext) BracketedTripleExpr() IBracketedTripleExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBracketedTripleExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBracketedTripleExprContext)
}

func (s *UnaryTripleExprContext) TripleExprLabel() ITripleExprLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITripleExprLabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITripleExprLabelContext)
}

func (s *UnaryTripleExprContext) Include() IIncludeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncludeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncludeContext)
}

func (s *UnaryTripleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryTripleExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UnaryTripleExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterUnaryTripleExpr(s)
	}
}

func (s *UnaryTripleExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitUnaryTripleExpr(s)
	}
}




func (p *ShExDocParser) UnaryTripleExpr() (localctx IUnaryTripleExprContext) {
	localctx = NewUnaryTripleExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ShExDocParserRULE_unaryTripleExpr)
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

	p.SetState(494)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserT__1, ShExDocParserT__9, ShExDocParserT__13, ShExDocParserRDF_TYPE, ShExDocParserIRIREF, ShExDocParserPNAME_NS, ShExDocParserPNAME_LN:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(487)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == ShExDocParserT__9 {
			{
				p.SetState(485)
				p.Match(ShExDocParserT__9)
			}
			{
				p.SetState(486)
				p.TripleExprLabel()
			}

		}
		p.SetState(491)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ShExDocParserT__13, ShExDocParserRDF_TYPE, ShExDocParserIRIREF, ShExDocParserPNAME_NS, ShExDocParserPNAME_LN:
			{
				p.SetState(489)
				p.TripleConstraint()
			}


		case ShExDocParserT__1:
			{
				p.SetState(490)
				p.BracketedTripleExpr()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}


	case ShExDocParserT__17:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(493)
			p.Include()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IBracketedTripleExprContext is an interface to support dynamic dispatch.
type IBracketedTripleExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBracketedTripleExprContext differentiates from other interfaces.
	IsBracketedTripleExprContext()
}

type BracketedTripleExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracketedTripleExprContext() *BracketedTripleExprContext {
	var p = new(BracketedTripleExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_bracketedTripleExpr
	return p
}

func (*BracketedTripleExprContext) IsBracketedTripleExprContext() {}

func NewBracketedTripleExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracketedTripleExprContext {
	var p = new(BracketedTripleExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_bracketedTripleExpr

	return p
}

func (s *BracketedTripleExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BracketedTripleExprContext) TripleExpression() ITripleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITripleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITripleExpressionContext)
}

func (s *BracketedTripleExprContext) Cardinality() ICardinalityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICardinalityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICardinalityContext)
}

func (s *BracketedTripleExprContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *BracketedTripleExprContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *BracketedTripleExprContext) AllSemanticAction() []ISemanticActionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISemanticActionContext)(nil)).Elem())
	var tst = make([]ISemanticActionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISemanticActionContext)
		}
	}

	return tst
}

func (s *BracketedTripleExprContext) SemanticAction(i int) ISemanticActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISemanticActionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISemanticActionContext)
}

func (s *BracketedTripleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketedTripleExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BracketedTripleExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterBracketedTripleExpr(s)
	}
}

func (s *BracketedTripleExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitBracketedTripleExpr(s)
	}
}




func (p *ShExDocParser) BracketedTripleExpr() (localctx IBracketedTripleExprContext) {
	localctx = NewBracketedTripleExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ShExDocParserRULE_bracketedTripleExpr)
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
		p.SetState(496)
		p.Match(ShExDocParserT__1)
	}
	{
		p.SetState(497)
		p.TripleExpression()
	}
	{
		p.SetState(498)
		p.Match(ShExDocParserT__2)
	}
	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ShExDocParserT__5) | (1 << ShExDocParserT__10) | (1 << ShExDocParserT__11))) != 0) || _la == ShExDocParserUNBOUNDED {
		{
			p.SetState(499)
			p.Cardinality()
		}

	}
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserT__18 {
		{
			p.SetState(502)
			p.Annotation()
		}


		p.SetState(507)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserT__19 {
		{
			p.SetState(508)
			p.SemanticAction()
		}


		p.SetState(513)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ITripleConstraintContext is an interface to support dynamic dispatch.
type ITripleConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTripleConstraintContext differentiates from other interfaces.
	IsTripleConstraintContext()
}

type TripleConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTripleConstraintContext() *TripleConstraintContext {
	var p = new(TripleConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_tripleConstraint
	return p
}

func (*TripleConstraintContext) IsTripleConstraintContext() {}

func NewTripleConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TripleConstraintContext {
	var p = new(TripleConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_tripleConstraint

	return p
}

func (s *TripleConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *TripleConstraintContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *TripleConstraintContext) InlineShapeExpression() IInlineShapeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineShapeExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineShapeExpressionContext)
}

func (s *TripleConstraintContext) SenseFlags() ISenseFlagsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISenseFlagsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISenseFlagsContext)
}

func (s *TripleConstraintContext) Cardinality() ICardinalityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICardinalityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICardinalityContext)
}

func (s *TripleConstraintContext) AllAnnotation() []IAnnotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *TripleConstraintContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *TripleConstraintContext) AllSemanticAction() []ISemanticActionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISemanticActionContext)(nil)).Elem())
	var tst = make([]ISemanticActionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISemanticActionContext)
		}
	}

	return tst
}

func (s *TripleConstraintContext) SemanticAction(i int) ISemanticActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISemanticActionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISemanticActionContext)
}

func (s *TripleConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TripleConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TripleConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterTripleConstraint(s)
	}
}

func (s *TripleConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitTripleConstraint(s)
	}
}




func (p *ShExDocParser) TripleConstraint() (localctx ITripleConstraintContext) {
	localctx = NewTripleConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ShExDocParserRULE_tripleConstraint)
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
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ShExDocParserT__13 {
		{
			p.SetState(514)
			p.SenseFlags()
		}

	}
	{
		p.SetState(517)
		p.Predicate()
	}
	{
		p.SetState(518)
		p.InlineShapeExpression()
	}
	p.SetState(520)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << ShExDocParserT__5) | (1 << ShExDocParserT__10) | (1 << ShExDocParserT__11))) != 0) || _la == ShExDocParserUNBOUNDED {
		{
			p.SetState(519)
			p.Cardinality()
		}

	}
	p.SetState(525)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserT__18 {
		{
			p.SetState(522)
			p.Annotation()
		}


		p.SetState(527)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserT__19 {
		{
			p.SetState(528)
			p.SemanticAction()
		}


		p.SetState(533)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ICardinalityContext is an interface to support dynamic dispatch.
type ICardinalityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCardinalityContext differentiates from other interfaces.
	IsCardinalityContext()
}

type CardinalityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCardinalityContext() *CardinalityContext {
	var p = new(CardinalityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_cardinality
	return p
}

func (*CardinalityContext) IsCardinalityContext() {}

func NewCardinalityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CardinalityContext {
	var p = new(CardinalityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_cardinality

	return p
}

func (s *CardinalityContext) GetParser() antlr.Parser { return s.parser }

func (s *CardinalityContext) CopyFrom(ctx *CardinalityContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CardinalityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CardinalityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type StarCardinalityContext struct {
	*CardinalityContext
}

func NewStarCardinalityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StarCardinalityContext {
	var p = new(StarCardinalityContext)

	p.CardinalityContext = NewEmptyCardinalityContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CardinalityContext))

	return p
}

func (s *StarCardinalityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarCardinalityContext) UNBOUNDED() antlr.TerminalNode {
	return s.GetToken(ShExDocParserUNBOUNDED, 0)
}


func (s *StarCardinalityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterStarCardinality(s)
	}
}

func (s *StarCardinalityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitStarCardinality(s)
	}
}


type RepeatCardinalityContext struct {
	*CardinalityContext
}

func NewRepeatCardinalityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RepeatCardinalityContext {
	var p = new(RepeatCardinalityContext)

	p.CardinalityContext = NewEmptyCardinalityContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CardinalityContext))

	return p
}

func (s *RepeatCardinalityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatCardinalityContext) RepeatRange() IRepeatRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRepeatRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRepeatRangeContext)
}


func (s *RepeatCardinalityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterRepeatCardinality(s)
	}
}

func (s *RepeatCardinalityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitRepeatCardinality(s)
	}
}


type PlusCardinalityContext struct {
	*CardinalityContext
}

func NewPlusCardinalityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusCardinalityContext {
	var p = new(PlusCardinalityContext)

	p.CardinalityContext = NewEmptyCardinalityContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CardinalityContext))

	return p
}

func (s *PlusCardinalityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusCardinalityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterPlusCardinality(s)
	}
}

func (s *PlusCardinalityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitPlusCardinality(s)
	}
}


type OptionalCardinalityContext struct {
	*CardinalityContext
}

func NewOptionalCardinalityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OptionalCardinalityContext {
	var p = new(OptionalCardinalityContext)

	p.CardinalityContext = NewEmptyCardinalityContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CardinalityContext))

	return p
}

func (s *OptionalCardinalityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalCardinalityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterOptionalCardinality(s)
	}
}

func (s *OptionalCardinalityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitOptionalCardinality(s)
	}
}



func (p *ShExDocParser) Cardinality() (localctx ICardinalityContext) {
	localctx = NewCardinalityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, ShExDocParserRULE_cardinality)

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

	p.SetState(538)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserUNBOUNDED:
		localctx = NewStarCardinalityContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(534)
			p.Match(ShExDocParserUNBOUNDED)
		}


	case ShExDocParserT__10:
		localctx = NewPlusCardinalityContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(535)
			p.Match(ShExDocParserT__10)
		}


	case ShExDocParserT__11:
		localctx = NewOptionalCardinalityContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(536)
			p.Match(ShExDocParserT__11)
		}


	case ShExDocParserT__5:
		localctx = NewRepeatCardinalityContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(537)
			p.RepeatRange()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IRepeatRangeContext is an interface to support dynamic dispatch.
type IRepeatRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRepeatRangeContext differentiates from other interfaces.
	IsRepeatRangeContext()
}

type RepeatRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatRangeContext() *RepeatRangeContext {
	var p = new(RepeatRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_repeatRange
	return p
}

func (*RepeatRangeContext) IsRepeatRangeContext() {}

func NewRepeatRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatRangeContext {
	var p = new(RepeatRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_repeatRange

	return p
}

func (s *RepeatRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatRangeContext) CopyFrom(ctx *RepeatRangeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *RepeatRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type ExactRangeContext struct {
	*RepeatRangeContext
}

func NewExactRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExactRangeContext {
	var p = new(ExactRangeContext)

	p.RepeatRangeContext = NewEmptyRepeatRangeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RepeatRangeContext))

	return p
}

func (s *ExactRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExactRangeContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ShExDocParserINTEGER, 0)
}


func (s *ExactRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterExactRange(s)
	}
}

func (s *ExactRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitExactRange(s)
	}
}


type MinMaxRangeContext struct {
	*RepeatRangeContext
}

func NewMinMaxRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinMaxRangeContext {
	var p = new(MinMaxRangeContext)

	p.RepeatRangeContext = NewEmptyRepeatRangeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RepeatRangeContext))

	return p
}

func (s *MinMaxRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinMaxRangeContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(ShExDocParserINTEGER)
}

func (s *MinMaxRangeContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(ShExDocParserINTEGER, i)
}

func (s *MinMaxRangeContext) UNBOUNDED() antlr.TerminalNode {
	return s.GetToken(ShExDocParserUNBOUNDED, 0)
}


func (s *MinMaxRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterMinMaxRange(s)
	}
}

func (s *MinMaxRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitMinMaxRange(s)
	}
}



func (p *ShExDocParser) RepeatRange() (localctx IRepeatRangeContext) {
	localctx = NewRepeatRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, ShExDocParserRULE_repeatRange)
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

	p.SetState(550)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 67, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExactRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(540)
			p.Match(ShExDocParserT__5)
		}
		{
			p.SetState(541)
			p.Match(ShExDocParserINTEGER)
		}
		{
			p.SetState(542)
			p.Match(ShExDocParserT__6)
		}


	case 2:
		localctx = NewMinMaxRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(543)
			p.Match(ShExDocParserT__5)
		}
		{
			p.SetState(544)
			p.Match(ShExDocParserINTEGER)
		}
		{
			p.SetState(545)
			p.Match(ShExDocParserT__12)
		}
		p.SetState(547)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == ShExDocParserINTEGER || _la == ShExDocParserUNBOUNDED {
			{
				p.SetState(546)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ShExDocParserINTEGER || _la == ShExDocParserUNBOUNDED) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(549)
			p.Match(ShExDocParserT__6)
		}

	}


	return localctx
}


// ISenseFlagsContext is an interface to support dynamic dispatch.
type ISenseFlagsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSenseFlagsContext differentiates from other interfaces.
	IsSenseFlagsContext()
}

type SenseFlagsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySenseFlagsContext() *SenseFlagsContext {
	var p = new(SenseFlagsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_senseFlags
	return p
}

func (*SenseFlagsContext) IsSenseFlagsContext() {}

func NewSenseFlagsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SenseFlagsContext {
	var p = new(SenseFlagsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_senseFlags

	return p
}

func (s *SenseFlagsContext) GetParser() antlr.Parser { return s.parser }
func (s *SenseFlagsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SenseFlagsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SenseFlagsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterSenseFlags(s)
	}
}

func (s *SenseFlagsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitSenseFlags(s)
	}
}




func (p *ShExDocParser) SenseFlags() (localctx ISenseFlagsContext) {
	localctx = NewSenseFlagsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, ShExDocParserRULE_senseFlags)

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
		p.SetState(552)
		p.Match(ShExDocParserT__13)
	}



	return localctx
}


// IValueSetContext is an interface to support dynamic dispatch.
type IValueSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueSetContext differentiates from other interfaces.
	IsValueSetContext()
}

type ValueSetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueSetContext() *ValueSetContext {
	var p = new(ValueSetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_valueSet
	return p
}

func (*ValueSetContext) IsValueSetContext() {}

func NewValueSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueSetContext {
	var p = new(ValueSetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_valueSet

	return p
}

func (s *ValueSetContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueSetContext) AllValueSetValue() []IValueSetValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueSetValueContext)(nil)).Elem())
	var tst = make([]IValueSetValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueSetValueContext)
		}
	}

	return tst
}

func (s *ValueSetContext) ValueSetValue(i int) IValueSetValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueSetValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueSetValueContext)
}

func (s *ValueSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ValueSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterValueSet(s)
	}
}

func (s *ValueSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitValueSet(s)
	}
}




func (p *ShExDocParser) ValueSet() (localctx IValueSetContext) {
	localctx = NewValueSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, ShExDocParserRULE_valueSet)
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
		p.SetState(554)
		p.Match(ShExDocParserT__14)
	}
	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == ShExDocParserT__3 || _la == ShExDocParserT__4 || ((((_la - 49)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 49))) & ((1 << (ShExDocParserKW_TRUE - 49)) | (1 << (ShExDocParserKW_FALSE - 49)) | (1 << (ShExDocParserIRIREF - 49)) | (1 << (ShExDocParserPNAME_NS - 49)) | (1 << (ShExDocParserPNAME_LN - 49)) | (1 << (ShExDocParserLANGTAG - 49)) | (1 << (ShExDocParserINTEGER - 49)) | (1 << (ShExDocParserDECIMAL - 49)) | (1 << (ShExDocParserDOUBLE - 49)) | (1 << (ShExDocParserSTRING_LITERAL1 - 49)) | (1 << (ShExDocParserSTRING_LITERAL2 - 49)) | (1 << (ShExDocParserSTRING_LITERAL_LONG1 - 49)) | (1 << (ShExDocParserSTRING_LITERAL_LONG2 - 49)))) != 0) {
		{
			p.SetState(555)
			p.ValueSetValue()
		}


		p.SetState(560)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(561)
		p.Match(ShExDocParserT__15)
	}



	return localctx
}


// IValueSetValueContext is an interface to support dynamic dispatch.
type IValueSetValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueSetValueContext differentiates from other interfaces.
	IsValueSetValueContext()
}

type ValueSetValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueSetValueContext() *ValueSetValueContext {
	var p = new(ValueSetValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_valueSetValue
	return p
}

func (*ValueSetValueContext) IsValueSetValueContext() {}

func NewValueSetValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueSetValueContext {
	var p = new(ValueSetValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_valueSetValue

	return p
}

func (s *ValueSetValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueSetValueContext) IriRange() IIriRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriRangeContext)
}

func (s *ValueSetValueContext) LiteralRange() ILiteralRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralRangeContext)
}

func (s *ValueSetValueContext) LanguageRange() ILanguageRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILanguageRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILanguageRangeContext)
}

func (s *ValueSetValueContext) AllIriExclusion() []IIriExclusionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIriExclusionContext)(nil)).Elem())
	var tst = make([]IIriExclusionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIriExclusionContext)
		}
	}

	return tst
}

func (s *ValueSetValueContext) IriExclusion(i int) IIriExclusionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriExclusionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIriExclusionContext)
}

func (s *ValueSetValueContext) AllLiteralExclusion() []ILiteralExclusionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralExclusionContext)(nil)).Elem())
	var tst = make([]ILiteralExclusionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralExclusionContext)
		}
	}

	return tst
}

func (s *ValueSetValueContext) LiteralExclusion(i int) ILiteralExclusionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralExclusionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralExclusionContext)
}

func (s *ValueSetValueContext) AllLanguageExclusion() []ILanguageExclusionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILanguageExclusionContext)(nil)).Elem())
	var tst = make([]ILanguageExclusionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILanguageExclusionContext)
		}
	}

	return tst
}

func (s *ValueSetValueContext) LanguageExclusion(i int) ILanguageExclusionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILanguageExclusionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILanguageExclusionContext)
}

func (s *ValueSetValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueSetValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ValueSetValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterValueSetValue(s)
	}
}

func (s *ValueSetValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitValueSetValue(s)
	}
}




func (p *ShExDocParser) ValueSetValue() (localctx IValueSetValueContext) {
	localctx = NewValueSetValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, ShExDocParserRULE_valueSetValue)
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

	p.SetState(584)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserIRIREF, ShExDocParserPNAME_NS, ShExDocParserPNAME_LN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(563)
			p.IriRange()
		}


	case ShExDocParserKW_TRUE, ShExDocParserKW_FALSE, ShExDocParserINTEGER, ShExDocParserDECIMAL, ShExDocParserDOUBLE, ShExDocParserSTRING_LITERAL1, ShExDocParserSTRING_LITERAL2, ShExDocParserSTRING_LITERAL_LONG1, ShExDocParserSTRING_LITERAL_LONG2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(564)
			p.LiteralRange()
		}


	case ShExDocParserT__4, ShExDocParserLANGTAG:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(565)
			p.LanguageRange()
		}


	case ShExDocParserT__3:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(566)
			p.Match(ShExDocParserT__3)
		}
		p.SetState(582)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext()) {
		case 1:
			p.SetState(568)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for ok := true; ok; ok = _la == ShExDocParserT__16 {
				{
					p.SetState(567)
					p.IriExclusion()
				}


				p.SetState(570)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}


		case 2:
			p.SetState(573)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for ok := true; ok; ok = _la == ShExDocParserT__16 {
				{
					p.SetState(572)
					p.LiteralExclusion()
				}


				p.SetState(575)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}


		case 3:
			p.SetState(578)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for ok := true; ok; ok = _la == ShExDocParserT__16 {
				{
					p.SetState(577)
					p.LanguageExclusion()
				}


				p.SetState(580)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IIriRangeContext is an interface to support dynamic dispatch.
type IIriRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIriRangeContext differentiates from other interfaces.
	IsIriRangeContext()
}

type IriRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIriRangeContext() *IriRangeContext {
	var p = new(IriRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_iriRange
	return p
}

func (*IriRangeContext) IsIriRangeContext() {}

func NewIriRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IriRangeContext {
	var p = new(IriRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_iriRange

	return p
}

func (s *IriRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *IriRangeContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *IriRangeContext) STEM_MARK() antlr.TerminalNode {
	return s.GetToken(ShExDocParserSTEM_MARK, 0)
}

func (s *IriRangeContext) AllIriExclusion() []IIriExclusionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIriExclusionContext)(nil)).Elem())
	var tst = make([]IIriExclusionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIriExclusionContext)
		}
	}

	return tst
}

func (s *IriRangeContext) IriExclusion(i int) IIriExclusionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriExclusionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIriExclusionContext)
}

func (s *IriRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IriRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IriRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterIriRange(s)
	}
}

func (s *IriRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitIriRange(s)
	}
}




func (p *ShExDocParser) IriRange() (localctx IIriRangeContext) {
	localctx = NewIriRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, ShExDocParserRULE_iriRange)
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
		p.SetState(586)
		p.Iri()
	}
	p.SetState(594)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ShExDocParserSTEM_MARK {
		{
			p.SetState(587)
			p.Match(ShExDocParserSTEM_MARK)
		}
		p.SetState(591)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == ShExDocParserT__16 {
			{
				p.SetState(588)
				p.IriExclusion()
			}


			p.SetState(593)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}



	return localctx
}


// IIriExclusionContext is an interface to support dynamic dispatch.
type IIriExclusionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIriExclusionContext differentiates from other interfaces.
	IsIriExclusionContext()
}

type IriExclusionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIriExclusionContext() *IriExclusionContext {
	var p = new(IriExclusionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_iriExclusion
	return p
}

func (*IriExclusionContext) IsIriExclusionContext() {}

func NewIriExclusionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IriExclusionContext {
	var p = new(IriExclusionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_iriExclusion

	return p
}

func (s *IriExclusionContext) GetParser() antlr.Parser { return s.parser }

func (s *IriExclusionContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *IriExclusionContext) STEM_MARK() antlr.TerminalNode {
	return s.GetToken(ShExDocParserSTEM_MARK, 0)
}

func (s *IriExclusionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IriExclusionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IriExclusionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterIriExclusion(s)
	}
}

func (s *IriExclusionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitIriExclusion(s)
	}
}




func (p *ShExDocParser) IriExclusion() (localctx IIriExclusionContext) {
	localctx = NewIriExclusionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, ShExDocParserRULE_iriExclusion)
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
		p.SetState(596)
		p.Match(ShExDocParserT__16)
	}
	{
		p.SetState(597)
		p.Iri()
	}
	p.SetState(599)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ShExDocParserSTEM_MARK {
		{
			p.SetState(598)
			p.Match(ShExDocParserSTEM_MARK)
		}

	}



	return localctx
}


// ILiteralRangeContext is an interface to support dynamic dispatch.
type ILiteralRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralRangeContext differentiates from other interfaces.
	IsLiteralRangeContext()
}

type LiteralRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralRangeContext() *LiteralRangeContext {
	var p = new(LiteralRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_literalRange
	return p
}

func (*LiteralRangeContext) IsLiteralRangeContext() {}

func NewLiteralRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralRangeContext {
	var p = new(LiteralRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_literalRange

	return p
}

func (s *LiteralRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralRangeContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralRangeContext) STEM_MARK() antlr.TerminalNode {
	return s.GetToken(ShExDocParserSTEM_MARK, 0)
}

func (s *LiteralRangeContext) AllLiteralExclusion() []ILiteralExclusionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralExclusionContext)(nil)).Elem())
	var tst = make([]ILiteralExclusionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralExclusionContext)
		}
	}

	return tst
}

func (s *LiteralRangeContext) LiteralExclusion(i int) ILiteralExclusionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralExclusionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralExclusionContext)
}

func (s *LiteralRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LiteralRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterLiteralRange(s)
	}
}

func (s *LiteralRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitLiteralRange(s)
	}
}




func (p *ShExDocParser) LiteralRange() (localctx ILiteralRangeContext) {
	localctx = NewLiteralRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, ShExDocParserRULE_literalRange)
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
		p.SetState(601)
		p.Literal()
	}
	p.SetState(609)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ShExDocParserSTEM_MARK {
		{
			p.SetState(602)
			p.Match(ShExDocParserSTEM_MARK)
		}
		p.SetState(606)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == ShExDocParserT__16 {
			{
				p.SetState(603)
				p.LiteralExclusion()
			}


			p.SetState(608)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}



	return localctx
}


// ILiteralExclusionContext is an interface to support dynamic dispatch.
type ILiteralExclusionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralExclusionContext differentiates from other interfaces.
	IsLiteralExclusionContext()
}

type LiteralExclusionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralExclusionContext() *LiteralExclusionContext {
	var p = new(LiteralExclusionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_literalExclusion
	return p
}

func (*LiteralExclusionContext) IsLiteralExclusionContext() {}

func NewLiteralExclusionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralExclusionContext {
	var p = new(LiteralExclusionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_literalExclusion

	return p
}

func (s *LiteralExclusionContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralExclusionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExclusionContext) STEM_MARK() antlr.TerminalNode {
	return s.GetToken(ShExDocParserSTEM_MARK, 0)
}

func (s *LiteralExclusionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExclusionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LiteralExclusionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterLiteralExclusion(s)
	}
}

func (s *LiteralExclusionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitLiteralExclusion(s)
	}
}




func (p *ShExDocParser) LiteralExclusion() (localctx ILiteralExclusionContext) {
	localctx = NewLiteralExclusionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, ShExDocParserRULE_literalExclusion)
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
		p.SetState(611)
		p.Match(ShExDocParserT__16)
	}
	{
		p.SetState(612)
		p.Literal()
	}
	p.SetState(614)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ShExDocParserSTEM_MARK {
		{
			p.SetState(613)
			p.Match(ShExDocParserSTEM_MARK)
		}

	}



	return localctx
}


// ILanguageRangeContext is an interface to support dynamic dispatch.
type ILanguageRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLanguageRangeContext differentiates from other interfaces.
	IsLanguageRangeContext()
}

type LanguageRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLanguageRangeContext() *LanguageRangeContext {
	var p = new(LanguageRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_languageRange
	return p
}

func (*LanguageRangeContext) IsLanguageRangeContext() {}

func NewLanguageRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LanguageRangeContext {
	var p = new(LanguageRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_languageRange

	return p
}

func (s *LanguageRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *LanguageRangeContext) CopyFrom(ctx *LanguageRangeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LanguageRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LanguageRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type LanguageRangeFullContext struct {
	*LanguageRangeContext
}

func NewLanguageRangeFullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LanguageRangeFullContext {
	var p = new(LanguageRangeFullContext)

	p.LanguageRangeContext = NewEmptyLanguageRangeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LanguageRangeContext))

	return p
}

func (s *LanguageRangeFullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LanguageRangeFullContext) LANGTAG() antlr.TerminalNode {
	return s.GetToken(ShExDocParserLANGTAG, 0)
}

func (s *LanguageRangeFullContext) STEM_MARK() antlr.TerminalNode {
	return s.GetToken(ShExDocParserSTEM_MARK, 0)
}

func (s *LanguageRangeFullContext) AllLanguageExclusion() []ILanguageExclusionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILanguageExclusionContext)(nil)).Elem())
	var tst = make([]ILanguageExclusionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILanguageExclusionContext)
		}
	}

	return tst
}

func (s *LanguageRangeFullContext) LanguageExclusion(i int) ILanguageExclusionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILanguageExclusionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILanguageExclusionContext)
}


func (s *LanguageRangeFullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterLanguageRangeFull(s)
	}
}

func (s *LanguageRangeFullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitLanguageRangeFull(s)
	}
}


type LanguageRangeAtContext struct {
	*LanguageRangeContext
}

func NewLanguageRangeAtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LanguageRangeAtContext {
	var p = new(LanguageRangeAtContext)

	p.LanguageRangeContext = NewEmptyLanguageRangeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LanguageRangeContext))

	return p
}

func (s *LanguageRangeAtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LanguageRangeAtContext) STEM_MARK() antlr.TerminalNode {
	return s.GetToken(ShExDocParserSTEM_MARK, 0)
}

func (s *LanguageRangeAtContext) AllLanguageExclusion() []ILanguageExclusionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILanguageExclusionContext)(nil)).Elem())
	var tst = make([]ILanguageExclusionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILanguageExclusionContext)
		}
	}

	return tst
}

func (s *LanguageRangeAtContext) LanguageExclusion(i int) ILanguageExclusionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILanguageExclusionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILanguageExclusionContext)
}


func (s *LanguageRangeAtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterLanguageRangeAt(s)
	}
}

func (s *LanguageRangeAtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitLanguageRangeAt(s)
	}
}



func (p *ShExDocParser) LanguageRange() (localctx ILanguageRangeContext) {
	localctx = NewLanguageRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, ShExDocParserRULE_languageRange)
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

	p.SetState(634)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserLANGTAG:
		localctx = NewLanguageRangeFullContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(616)
			p.Match(ShExDocParserLANGTAG)
		}
		p.SetState(624)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == ShExDocParserSTEM_MARK {
			{
				p.SetState(617)
				p.Match(ShExDocParserSTEM_MARK)
			}
			p.SetState(621)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for _la == ShExDocParserT__16 {
				{
					p.SetState(618)
					p.LanguageExclusion()
				}


				p.SetState(623)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}


	case ShExDocParserT__4:
		localctx = NewLanguageRangeAtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(626)
			p.Match(ShExDocParserT__4)
		}
		{
			p.SetState(627)
			p.Match(ShExDocParserSTEM_MARK)
		}
		p.SetState(631)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == ShExDocParserT__16 {
			{
				p.SetState(628)
				p.LanguageExclusion()
			}


			p.SetState(633)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ILanguageExclusionContext is an interface to support dynamic dispatch.
type ILanguageExclusionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLanguageExclusionContext differentiates from other interfaces.
	IsLanguageExclusionContext()
}

type LanguageExclusionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLanguageExclusionContext() *LanguageExclusionContext {
	var p = new(LanguageExclusionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_languageExclusion
	return p
}

func (*LanguageExclusionContext) IsLanguageExclusionContext() {}

func NewLanguageExclusionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LanguageExclusionContext {
	var p = new(LanguageExclusionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_languageExclusion

	return p
}

func (s *LanguageExclusionContext) GetParser() antlr.Parser { return s.parser }

func (s *LanguageExclusionContext) LANGTAG() antlr.TerminalNode {
	return s.GetToken(ShExDocParserLANGTAG, 0)
}

func (s *LanguageExclusionContext) STEM_MARK() antlr.TerminalNode {
	return s.GetToken(ShExDocParserSTEM_MARK, 0)
}

func (s *LanguageExclusionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LanguageExclusionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LanguageExclusionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterLanguageExclusion(s)
	}
}

func (s *LanguageExclusionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitLanguageExclusion(s)
	}
}




func (p *ShExDocParser) LanguageExclusion() (localctx ILanguageExclusionContext) {
	localctx = NewLanguageExclusionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, ShExDocParserRULE_languageExclusion)
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
		p.SetState(636)
		p.Match(ShExDocParserT__16)
	}
	{
		p.SetState(637)
		p.Match(ShExDocParserLANGTAG)
	}
	p.SetState(639)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ShExDocParserSTEM_MARK {
		{
			p.SetState(638)
			p.Match(ShExDocParserSTEM_MARK)
		}

	}



	return localctx
}


// IIncludeContext is an interface to support dynamic dispatch.
type IIncludeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncludeContext differentiates from other interfaces.
	IsIncludeContext()
}

type IncludeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeContext() *IncludeContext {
	var p = new(IncludeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_include
	return p
}

func (*IncludeContext) IsIncludeContext() {}

func NewIncludeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeContext {
	var p = new(IncludeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_include

	return p
}

func (s *IncludeContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeContext) TripleExprLabel() ITripleExprLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITripleExprLabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITripleExprLabelContext)
}

func (s *IncludeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IncludeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterInclude(s)
	}
}

func (s *IncludeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitInclude(s)
	}
}




func (p *ShExDocParser) Include() (localctx IIncludeContext) {
	localctx = NewIncludeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, ShExDocParserRULE_include)

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
		p.SetState(641)
		p.Match(ShExDocParserT__17)
	}
	{
		p.SetState(642)
		p.TripleExprLabel()
	}



	return localctx
}


// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *AnnotationContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *AnnotationContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitAnnotation(s)
	}
}




func (p *ShExDocParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, ShExDocParserRULE_annotation)

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
		p.SetState(644)
		p.Match(ShExDocParserT__18)
	}
	{
		p.SetState(645)
		p.Predicate()
	}
	p.SetState(648)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserIRIREF, ShExDocParserPNAME_NS, ShExDocParserPNAME_LN:
		{
			p.SetState(646)
			p.Iri()
		}


	case ShExDocParserKW_TRUE, ShExDocParserKW_FALSE, ShExDocParserINTEGER, ShExDocParserDECIMAL, ShExDocParserDOUBLE, ShExDocParserSTRING_LITERAL1, ShExDocParserSTRING_LITERAL2, ShExDocParserSTRING_LITERAL_LONG1, ShExDocParserSTRING_LITERAL_LONG2:
		{
			p.SetState(647)
			p.Literal()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// ISemanticActionContext is an interface to support dynamic dispatch.
type ISemanticActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSemanticActionContext differentiates from other interfaces.
	IsSemanticActionContext()
}

type SemanticActionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySemanticActionContext() *SemanticActionContext {
	var p = new(SemanticActionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_semanticAction
	return p
}

func (*SemanticActionContext) IsSemanticActionContext() {}

func NewSemanticActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SemanticActionContext {
	var p = new(SemanticActionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_semanticAction

	return p
}

func (s *SemanticActionContext) GetParser() antlr.Parser { return s.parser }

func (s *SemanticActionContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *SemanticActionContext) CODE() antlr.TerminalNode {
	return s.GetToken(ShExDocParserCODE, 0)
}

func (s *SemanticActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SemanticActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SemanticActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterSemanticAction(s)
	}
}

func (s *SemanticActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitSemanticAction(s)
	}
}




func (p *ShExDocParser) SemanticAction() (localctx ISemanticActionContext) {
	localctx = NewSemanticActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, ShExDocParserRULE_semanticAction)
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
		p.SetState(650)
		p.Match(ShExDocParserT__19)
	}
	{
		p.SetState(651)
		p.Iri()
	}
	{
		p.SetState(652)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ShExDocParserT__19 || _la == ShExDocParserCODE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) RdfLiteral() IRdfLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRdfLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRdfLiteralContext)
}

func (s *LiteralContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *LiteralContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitLiteral(s)
	}
}




func (p *ShExDocParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, ShExDocParserRULE_literal)

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

	p.SetState(657)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserSTRING_LITERAL1, ShExDocParserSTRING_LITERAL2, ShExDocParserSTRING_LITERAL_LONG1, ShExDocParserSTRING_LITERAL_LONG2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(654)
			p.RdfLiteral()
		}


	case ShExDocParserINTEGER, ShExDocParserDECIMAL, ShExDocParserDOUBLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(655)
			p.NumericLiteral()
		}


	case ShExDocParserKW_TRUE, ShExDocParserKW_FALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(656)
			p.BooleanLiteral()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *PredicateContext) RdfType() IRdfTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRdfTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRdfTypeContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitPredicate(s)
	}
}




func (p *ShExDocParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, ShExDocParserRULE_predicate)

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

	p.SetState(661)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserIRIREF, ShExDocParserPNAME_NS, ShExDocParserPNAME_LN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(659)
			p.Iri()
		}


	case ShExDocParserRDF_TYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(660)
			p.RdfType()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IRdfTypeContext is an interface to support dynamic dispatch.
type IRdfTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRdfTypeContext differentiates from other interfaces.
	IsRdfTypeContext()
}

type RdfTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRdfTypeContext() *RdfTypeContext {
	var p = new(RdfTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_rdfType
	return p
}

func (*RdfTypeContext) IsRdfTypeContext() {}

func NewRdfTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RdfTypeContext {
	var p = new(RdfTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_rdfType

	return p
}

func (s *RdfTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *RdfTypeContext) RDF_TYPE() antlr.TerminalNode {
	return s.GetToken(ShExDocParserRDF_TYPE, 0)
}

func (s *RdfTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RdfTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RdfTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterRdfType(s)
	}
}

func (s *RdfTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitRdfType(s)
	}
}




func (p *ShExDocParser) RdfType() (localctx IRdfTypeContext) {
	localctx = NewRdfTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, ShExDocParserRULE_rdfType)

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
		p.SetState(663)
		p.Match(ShExDocParserRDF_TYPE)
	}



	return localctx
}


// IDatatypeContext is an interface to support dynamic dispatch.
type IDatatypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatatypeContext differentiates from other interfaces.
	IsDatatypeContext()
}

type DatatypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatatypeContext() *DatatypeContext {
	var p = new(DatatypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_datatype
	return p
}

func (*DatatypeContext) IsDatatypeContext() {}

func NewDatatypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatatypeContext {
	var p = new(DatatypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_datatype

	return p
}

func (s *DatatypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DatatypeContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *DatatypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatatypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DatatypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterDatatype(s)
	}
}

func (s *DatatypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitDatatype(s)
	}
}




func (p *ShExDocParser) Datatype() (localctx IDatatypeContext) {
	localctx = NewDatatypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, ShExDocParserRULE_datatype)

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
		p.SetState(665)
		p.Iri()
	}



	return localctx
}


// IShapeExprLabelContext is an interface to support dynamic dispatch.
type IShapeExprLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShapeExprLabelContext differentiates from other interfaces.
	IsShapeExprLabelContext()
}

type ShapeExprLabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShapeExprLabelContext() *ShapeExprLabelContext {
	var p = new(ShapeExprLabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_shapeExprLabel
	return p
}

func (*ShapeExprLabelContext) IsShapeExprLabelContext() {}

func NewShapeExprLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShapeExprLabelContext {
	var p = new(ShapeExprLabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_shapeExprLabel

	return p
}

func (s *ShapeExprLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *ShapeExprLabelContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *ShapeExprLabelContext) BlankNode() IBlankNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlankNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlankNodeContext)
}

func (s *ShapeExprLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeExprLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ShapeExprLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterShapeExprLabel(s)
	}
}

func (s *ShapeExprLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitShapeExprLabel(s)
	}
}




func (p *ShExDocParser) ShapeExprLabel() (localctx IShapeExprLabelContext) {
	localctx = NewShapeExprLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, ShExDocParserRULE_shapeExprLabel)

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

	p.SetState(669)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserIRIREF, ShExDocParserPNAME_NS, ShExDocParserPNAME_LN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(667)
			p.Iri()
		}


	case ShExDocParserBLANK_NODE_LABEL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(668)
			p.BlankNode()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ITripleExprLabelContext is an interface to support dynamic dispatch.
type ITripleExprLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTripleExprLabelContext differentiates from other interfaces.
	IsTripleExprLabelContext()
}

type TripleExprLabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTripleExprLabelContext() *TripleExprLabelContext {
	var p = new(TripleExprLabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_tripleExprLabel
	return p
}

func (*TripleExprLabelContext) IsTripleExprLabelContext() {}

func NewTripleExprLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TripleExprLabelContext {
	var p = new(TripleExprLabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_tripleExprLabel

	return p
}

func (s *TripleExprLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *TripleExprLabelContext) Iri() IIriContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriContext)
}

func (s *TripleExprLabelContext) BlankNode() IBlankNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlankNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlankNodeContext)
}

func (s *TripleExprLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TripleExprLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TripleExprLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterTripleExprLabel(s)
	}
}

func (s *TripleExprLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitTripleExprLabel(s)
	}
}




func (p *ShExDocParser) TripleExprLabel() (localctx ITripleExprLabelContext) {
	localctx = NewTripleExprLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, ShExDocParserRULE_tripleExprLabel)

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

	p.SetState(673)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserIRIREF, ShExDocParserPNAME_NS, ShExDocParserPNAME_LN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(671)
			p.Iri()
		}


	case ShExDocParserBLANK_NODE_LABEL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(672)
			p.BlankNode()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_numericLiteral
	return p
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ShExDocParserINTEGER, 0)
}

func (s *NumericLiteralContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(ShExDocParserDECIMAL, 0)
}

func (s *NumericLiteralContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ShExDocParserDOUBLE, 0)
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}




func (p *ShExDocParser) NumericLiteral() (localctx INumericLiteralContext) {
	localctx = NewNumericLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, ShExDocParserRULE_numericLiteral)
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
		p.SetState(675)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 64)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 64))) & ((1 << (ShExDocParserINTEGER - 64)) | (1 << (ShExDocParserDECIMAL - 64)) | (1 << (ShExDocParserDOUBLE - 64)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IRdfLiteralContext is an interface to support dynamic dispatch.
type IRdfLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRdfLiteralContext differentiates from other interfaces.
	IsRdfLiteralContext()
}

type RdfLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRdfLiteralContext() *RdfLiteralContext {
	var p = new(RdfLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_rdfLiteral
	return p
}

func (*RdfLiteralContext) IsRdfLiteralContext() {}

func NewRdfLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RdfLiteralContext {
	var p = new(RdfLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_rdfLiteral

	return p
}

func (s *RdfLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RdfLiteralContext) RdfString() IRdfStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRdfStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRdfStringContext)
}

func (s *RdfLiteralContext) LANGTAG() antlr.TerminalNode {
	return s.GetToken(ShExDocParserLANGTAG, 0)
}

func (s *RdfLiteralContext) Datatype() IDatatypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatatypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatatypeContext)
}

func (s *RdfLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RdfLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RdfLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterRdfLiteral(s)
	}
}

func (s *RdfLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitRdfLiteral(s)
	}
}




func (p *ShExDocParser) RdfLiteral() (localctx IRdfLiteralContext) {
	localctx = NewRdfLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, ShExDocParserRULE_rdfLiteral)

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
		p.SetState(677)
		p.RdfString()
	}
	p.SetState(681)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 90, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(678)
			p.Match(ShExDocParserLANGTAG)
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 90, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(679)
			p.Match(ShExDocParserT__20)
		}
		{
			p.SetState(680)
			p.Datatype()
		}


	}



	return localctx
}


// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) KW_TRUE() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_TRUE, 0)
}

func (s *BooleanLiteralContext) KW_FALSE() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_FALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}




func (p *ShExDocParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, ShExDocParserRULE_booleanLiteral)
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
		p.SetState(683)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ShExDocParserKW_TRUE || _la == ShExDocParserKW_FALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IRdfStringContext is an interface to support dynamic dispatch.
type IRdfStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRdfStringContext differentiates from other interfaces.
	IsRdfStringContext()
}

type RdfStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRdfStringContext() *RdfStringContext {
	var p = new(RdfStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_rdfString
	return p
}

func (*RdfStringContext) IsRdfStringContext() {}

func NewRdfStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RdfStringContext {
	var p = new(RdfStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_rdfString

	return p
}

func (s *RdfStringContext) GetParser() antlr.Parser { return s.parser }

func (s *RdfStringContext) STRING_LITERAL_LONG1() antlr.TerminalNode {
	return s.GetToken(ShExDocParserSTRING_LITERAL_LONG1, 0)
}

func (s *RdfStringContext) STRING_LITERAL_LONG2() antlr.TerminalNode {
	return s.GetToken(ShExDocParserSTRING_LITERAL_LONG2, 0)
}

func (s *RdfStringContext) STRING_LITERAL1() antlr.TerminalNode {
	return s.GetToken(ShExDocParserSTRING_LITERAL1, 0)
}

func (s *RdfStringContext) STRING_LITERAL2() antlr.TerminalNode {
	return s.GetToken(ShExDocParserSTRING_LITERAL2, 0)
}

func (s *RdfStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RdfStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RdfStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterRdfString(s)
	}
}

func (s *RdfStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitRdfString(s)
	}
}




func (p *ShExDocParser) RdfString() (localctx IRdfStringContext) {
	localctx = NewRdfStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, ShExDocParserRULE_rdfString)
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
		p.SetState(685)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 69)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 69))) & ((1 << (ShExDocParserSTRING_LITERAL1 - 69)) | (1 << (ShExDocParserSTRING_LITERAL2 - 69)) | (1 << (ShExDocParserSTRING_LITERAL_LONG1 - 69)) | (1 << (ShExDocParserSTRING_LITERAL_LONG2 - 69)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IIriContext is an interface to support dynamic dispatch.
type IIriContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIriContext differentiates from other interfaces.
	IsIriContext()
}

type IriContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIriContext() *IriContext {
	var p = new(IriContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_iri
	return p
}

func (*IriContext) IsIriContext() {}

func NewIriContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IriContext {
	var p = new(IriContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_iri

	return p
}

func (s *IriContext) GetParser() antlr.Parser { return s.parser }

func (s *IriContext) IRIREF() antlr.TerminalNode {
	return s.GetToken(ShExDocParserIRIREF, 0)
}

func (s *IriContext) PrefixedName() IPrefixedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixedNameContext)
}

func (s *IriContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IriContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IriContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterIri(s)
	}
}

func (s *IriContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitIri(s)
	}
}




func (p *ShExDocParser) Iri() (localctx IIriContext) {
	localctx = NewIriContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, ShExDocParserRULE_iri)

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

	p.SetState(689)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserIRIREF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(687)
			p.Match(ShExDocParserIRIREF)
		}


	case ShExDocParserPNAME_NS, ShExDocParserPNAME_LN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(688)
			p.PrefixedName()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IPrefixedNameContext is an interface to support dynamic dispatch.
type IPrefixedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixedNameContext differentiates from other interfaces.
	IsPrefixedNameContext()
}

type PrefixedNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixedNameContext() *PrefixedNameContext {
	var p = new(PrefixedNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_prefixedName
	return p
}

func (*PrefixedNameContext) IsPrefixedNameContext() {}

func NewPrefixedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixedNameContext {
	var p = new(PrefixedNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_prefixedName

	return p
}

func (s *PrefixedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixedNameContext) PNAME_LN() antlr.TerminalNode {
	return s.GetToken(ShExDocParserPNAME_LN, 0)
}

func (s *PrefixedNameContext) PNAME_NS() antlr.TerminalNode {
	return s.GetToken(ShExDocParserPNAME_NS, 0)
}

func (s *PrefixedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PrefixedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterPrefixedName(s)
	}
}

func (s *PrefixedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitPrefixedName(s)
	}
}




func (p *ShExDocParser) PrefixedName() (localctx IPrefixedNameContext) {
	localctx = NewPrefixedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, ShExDocParserRULE_prefixedName)
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
		p.SetState(691)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ShExDocParserPNAME_NS || _la == ShExDocParserPNAME_LN) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IBlankNodeContext is an interface to support dynamic dispatch.
type IBlankNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlankNodeContext differentiates from other interfaces.
	IsBlankNodeContext()
}

type BlankNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlankNodeContext() *BlankNodeContext {
	var p = new(BlankNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_blankNode
	return p
}

func (*BlankNodeContext) IsBlankNodeContext() {}

func NewBlankNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlankNodeContext {
	var p = new(BlankNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_blankNode

	return p
}

func (s *BlankNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *BlankNodeContext) BLANK_NODE_LABEL() antlr.TerminalNode {
	return s.GetToken(ShExDocParserBLANK_NODE_LABEL, 0)
}

func (s *BlankNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BlankNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterBlankNode(s)
	}
}

func (s *BlankNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitBlankNode(s)
	}
}




func (p *ShExDocParser) BlankNode() (localctx IBlankNodeContext) {
	localctx = NewBlankNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, ShExDocParserRULE_blankNode)

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
		p.SetState(693)
		p.Match(ShExDocParserBLANK_NODE_LABEL)
	}



	return localctx
}


// IExtensionContext is an interface to support dynamic dispatch.
type IExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtensionContext differentiates from other interfaces.
	IsExtensionContext()
}

type ExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionContext() *ExtensionContext {
	var p = new(ExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_extension
	return p
}

func (*ExtensionContext) IsExtensionContext() {}

func NewExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionContext {
	var p = new(ExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_extension

	return p
}

func (s *ExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionContext) KW_EXTENDS() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_EXTENDS, 0)
}

func (s *ExtensionContext) ShapeExprLabel() IShapeExprLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeExprLabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeExprLabelContext)
}

func (s *ExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterExtension(s)
	}
}

func (s *ExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitExtension(s)
	}
}




func (p *ShExDocParser) Extension() (localctx IExtensionContext) {
	localctx = NewExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, ShExDocParserRULE_extension)

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

	p.SetState(699)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserKW_EXTENDS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(695)
			p.Match(ShExDocParserKW_EXTENDS)
		}
		{
			p.SetState(696)
			p.ShapeExprLabel()
		}


	case ShExDocParserT__17:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(697)
			p.Match(ShExDocParserT__17)
		}
		{
			p.SetState(698)
			p.ShapeExprLabel()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IRestrictionsContext is an interface to support dynamic dispatch.
type IRestrictionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRestrictionsContext differentiates from other interfaces.
	IsRestrictionsContext()
}

type RestrictionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRestrictionsContext() *RestrictionsContext {
	var p = new(RestrictionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShExDocParserRULE_restrictions
	return p
}

func (*RestrictionsContext) IsRestrictionsContext() {}

func NewRestrictionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RestrictionsContext {
	var p = new(RestrictionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShExDocParserRULE_restrictions

	return p
}

func (s *RestrictionsContext) GetParser() antlr.Parser { return s.parser }

func (s *RestrictionsContext) KW_RESTRICTS() antlr.TerminalNode {
	return s.GetToken(ShExDocParserKW_RESTRICTS, 0)
}

func (s *RestrictionsContext) ShapeExprLabel() IShapeExprLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IShapeExprLabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IShapeExprLabelContext)
}

func (s *RestrictionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RestrictionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RestrictionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.EnterRestrictions(s)
	}
}

func (s *RestrictionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShExDocListener); ok {
		listenerT.ExitRestrictions(s)
	}
}




func (p *ShExDocParser) Restrictions() (localctx IRestrictionsContext) {
	localctx = NewRestrictionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 152, ShExDocParserRULE_restrictions)

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

	p.SetState(705)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ShExDocParserKW_RESTRICTS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(701)
			p.Match(ShExDocParserKW_RESTRICTS)
		}
		{
			p.SetState(702)
			p.ShapeExprLabel()
		}


	case ShExDocParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(703)
			p.Match(ShExDocParserT__16)
		}
		{
			p.SetState(704)
			p.ShapeExprLabel()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


