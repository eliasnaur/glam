// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package noise

import (
	"fmt"
	"testing"
)

import (
	"github.com/drakmaniso/glm"
)

//------------------------------------------------------------------------------

func ExamplePerlin3DAt() {
	for z := 0; z < 4; z++ {
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				v := Perlin3DAt(glm.Vec3{float32(x), float32(y), float32(z * 10)}.Slash(10.0))
				iv := uint8(v * 50.0 + 50.0)
				fmt.Printf("%2d", iv)
				if x < 9 {
					fmt.Printf(" ")
				}
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
	// Output:
	// 50 50 50 50 50 50 50 50 50 50
	// 45 45 45 46 48 49 51 53 54 54
	// 40 41 42 44 46 49 52 54 56 56
	// 37 38 40 42 45 49 52 54 56 56
	// 36 38 40 43 46 49 51 53 53 54
	// 37 40 42 45 48 50 50 50 50 50
	// 40 43 47 49 51 51 50 48 47 46
	// 44 48 52 54 55 54 51 48 45 43
	// 47 52 56 58 59 57 53 49 45 43
	// 49 54 58 61 61 59 56 51 47 45
	// 
	// 50 54 59 62 63 62 59 55 52 50
	// 55 60 64 67 68 67 64 60 57 55
	// 61 66 71 74 75 73 70 66 62 60
	// 68 73 77 80 81 79 76 71 66 63
	// 73 78 82 85 86 84 80 74 69 65
	// 75 80 84 88 89 87 82 76 70 65
	// 73 78 83 87 88 87 82 76 69 63
	// 68 73 78 83 85 84 80 73 66 60
	// 61 66 72 78 81 81 77 71 64 57
	// 55 60 66 72 76 77 75 69 62 55
	// 
	// 50 49 47 44 40 37 36 37 40 45
	// 44 44 42 40 38 37 38 40 45 49
	// 38 38 37 36 36 37 40 44 49 54
	// 31 31 31 32 35 38 43 49 54 58
	// 26 27 28 30 35 40 46 53 58 61
	// 25 25 27 31 36 43 50 56 60 62
	// 26 27 29 34 40 47 54 58 61 61
	// 31 32 34 39 45 52 57 60 60 58
	// 38 38 41 45 51 56 59 61 59 56
	// 44 45 47 51 55 59 61 61 59 55
	// 
	// 50 45 40 37 36 37 40 44 47 49
	// 54 50 45 42 41 42 45 49 52 54
	// 59 55 50 47 46 47 50 53 56 58
	// 62 59 55 53 51 52 53 56 59 61
	// 63 61 59 57 55 55 55 57 59 61
	// 62 62 61 59 57 56 55 56 57 60
	// 59 61 61 60 58 55 53 53 53 56
	// 55 58 60 59 56 53 50 49 49 51
	// 52 56 58 58 55 51 47 45 45 47
	// 50 54 57 57 54 50 46 43 43 45
}

//------------------------------------------------------------------------------

var outputPerlin3DAt [10][10][10]float32
var resultsPerlin3DAt [10][10][10]float32 = [10][10][10]float32{
	[10][10]float32{
		[10]float32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[10]float32{-0.099144004, 0.106848, -0.106848, 0.099144004, 0.099144004, -0.106848, 0, 0, 0, 0.106848},
		[10]float32{-0.188416, 0.234752, -0.234752, 0.188416, 0.188416, -0.234752, 0, 0, 0, 0.234752},
		[10]float32{-0.251076, 0.36523202, -0.36523202, 0.251076, 0.251076, -0.36523202, 0, 0, 0, 0.36523202},
		[10]float32{-0.27302402, 0.463488, -0.463488, 0.27302402, 0.27302402, -0.463488, 0, 0, 0, 0.463488},
		[10]float32{-0.25, 0.5, -0.5, 0.25, 0.25, -0.5, 0, 0, 0, 0.5},
		[10]float32{-0.19046399, 0.46348798, -0.46348798, 0.19046399, 0.19046399, -0.46348798, 0, 0, 0, 0.46348798},
		[10]float32{-0.114155985, 0.365232, -0.365232, 0.114155985, 0.114155985, -0.365232, 0, 0, 0, 0.365232},
		[10]float32{-0.046335887, 0.23475191, -0.23475191, 0.046335887, 0.046335887, -0.23475191, 0, 0, 0, 0.23475191},
		[10]float32{-0.007704323, 0.106848314, -0.106848314, 0.007704323, 0.007704323, -0.106848314, 0, 0, 0, 0.106848314}},
	[10][10]float32{
		[10]float32{0, 0.099144004, -0.0077040005, -0.099144004, 0.106848, 0.09144001, -0.106848, 0.106848, 0.099144004, -0.099144004},
		[10]float32{-0.09666394, 0.205992, -0.11265682, 0.0016313994, 0.20246543, -0.015144216, -0.10776262, 0.107564785, 0.09653205, 0.005940709},
		[10]float32{-0.17984454, 0.33394557, -0.23794127, 0.10031064, 0.27971062, -0.14575589, -0.10885748, 0.10756842, 0.08560011, 0.13203527},
		[10]float32{-0.23158634, 0.46465522, -0.36514768, 0.18301246, 0.31905523, -0.28334412, -0.10997439, 0.10648448, 0.06340161, 0.2610916},
		[10]float32{-0.23850796, 0.56344724, -0.45999634, 0.234379, 0.30817738, -0.39303133, -0.110815465, 0.10429397, 0.031416822, 0.35885465},
		[10]float32{-0.198288, 0.600856, -0.49357998, 0.24614802, 0.247432, -0.444008, -0.111128, 0.101284005, -0.005991999, 0.39614803},
		[10]float32{-0.121868595, 0.56555337, -0.4550773, 0.221405, 0.15111226, -0.4228983, -0.110815465, 0.09796149, -0.043088272, 0.3616675},
		[10]float32{-0.031375237, 0.46867442, -0.35593718, 0.17451555, 0.044501595, -0.33859673, -0.10997439, 0.09492992, -0.074232005, 0.26628286},
		[10]float32{0.046246134, 0.33954084, -0.22553422, 0.1267373, -0.043283198, -0.21857563, -0.108857475, 0.09272907, -0.09531361, 0.13884674},
		[10]float32{0.0898742, 0.2127817, -0.098295644, 0.09751286, -0.090656884, -0.096664205, -0.10776263, 0.09163786, -0.10515059, 0.013513118}},
	[10][10]float32{
		[10]float32{0, 0.188416, -0.046336003, -0.188416, 0.234752, 0.14208001, -0.234752, 0.234752, 0.188416, -0.188416},
		[10]float32{-0.08649254, 0.2952144, -0.14045967, -0.08644296, 0.31834266, 0.041047283, -0.24094065, 0.23970115, 0.17905125, -0.09354864},
		[10]float32{-0.15836063, 0.42316803, -0.25121057, 0.01914233, 0.37414825, -0.08193688, -0.24834886, 0.24029751, 0.15299307, 0.02182611},
		[10]float32{-0.19787638, 0.55459255, -0.36075848, 0.11655729, 0.3811212, -0.2109998, -0.25590622, 0.23418137, 0.10486342, 0.14206402},
		[10]float32{-0.19261783, 0.65558124, -0.43774742, 0.18952066, 0.33078682, -0.31392407, -0.26159725, 0.22114778, 0.038272016, 0.23609047},
		[10]float32{-0.141312, 0.697104, -0.45656002, 0.226832, 0.22972801, -0.362352, -0.263712, 0.202896, -0.037648, 0.276832},
		[10]float32{-0.055608943, 0.6677176, -0.4086929, 0.22763136, 0.09850153, -0.34410024, -0.26159725, 0.18252948, -0.11145324, 0.25300863},
		[10]float32{0.04221464, 0.5778852, -0.30624282, 0.20233871, -0.033015937, -0.26758552, -0.25590625, 0.16380487, -0.17235368, 0.17328705},
		[10]float32{0.12562157, 0.45590708, -0.17750284, 0.16927361, -0.13385089, -0.15836054, -0.24834885, 0.15013133, -0.21292594, 0.06279447},
		[10]float32{0.17424881, 0.3354612, -0.05467073, 0.14695516, -0.1837241, -0.047761634, -0.24094066, 0.14331955, -0.23157577, -0.04800565}},
	[10][10]float32{
		[10]float32{0, 0.251076, -0.114156015, -0.251076, 0.36523202, 0.13692, -0.36523202, 0.36523202, 0.251076, -0.251076},
		[10]float32{-0.06591435, 0.35764483, -0.18545648, -0.14861077, 0.42550722, 0.048558276, -0.3826568, 0.3794461, 0.229632, -0.17810038},
		[10]float32{-0.11997639, 0.4848834, -0.26667395, -0.040187277, 0.4489412, -0.05331067, -0.4035154, 0.38273507, 0.18465257, -0.08627881},
		[10]float32{-0.14685613, 0.61630803, -0.3416474, 0.0632744, 0.41529295, -0.15384577, -0.42479405, 0.36894435, 0.109623, 0.013648488},
		[10]float32{-0.1353338, 0.71974087, -0.38505828, 0.14511345, 0.32250485, -0.22714773, -0.44081765, 0.33728144, 0.0109104365, 0.09747843},
		[10]float32{-0.083691984, 0.767384, -0.37768996, 0.192922, 0.18476799, -0.25215197, -0.44677204, 0.29184604, -0.09784801, 0.14292198},
		[10]float32{-0.0014925748, 0.7479576, -0.31516084, 0.20421857, 0.028382286, -0.2219954, -0.44081765, 0.24045624, -0.2006521, 0.13915905},
		[10]float32{0.09226219, 0.670902, -0.21013051, 0.1878016, -0.11459112, -0.14685613, -0.42479405, 0.19276968, -0.28334108, 0.09057136},
		[10]float32{0.17458387, 0.56264514, -0.08797988, 0.16078322, -0.21488345, -0.050267175, -0.40351534, 0.15770027, -0.33709198, 0.014653326},
		[10]float32{0.22756682, 0.45493472, 0.02403452, 0.14130294, -0.2567771, 0.0410958, -0.38265684, 0.14013086, -0.3612126, -0.06589854}},
	[10][10]float32{
		[10]float32{0, 0.27302402, -0.19046399, -0.27302402, 0.463488, 0.08256003, -0.463488, 0.463488, 0.27302402, -0.27302402},
		[10]float32{-0.036307935, 0.37905684, -0.22866115, -0.17083621, 0.49093735, 0.012999071, -0.49740583, 0.49169952, 0.2352472, -0.2323813},
		[10]float32{-0.06768981, 0.50409883, -0.2688222, -0.06401259, 0.4749148, -0.05671361, -0.5380077, 0.5012354, 0.17055446, -0.1762797},
		[10]float32{-0.0833858, 0.63307923, -0.29793346, 0.03604064, 0.3988128, -0.11289809, -0.57942724, 0.4810678, 0.07321206, -0.1086096},
		[10]float32{-0.073478684, 0.73651206, -0.2992319, 0.112876594, 0.26891533, -0.13908444, -0.61061764, 0.42923498, -0.047443084, -0.04333437},
		[10]float32{-0.034127995, 0.78889596, -0.26192003, 0.154768, 0.107152, -0.12540801, -0.622208, 0.35238403, -0.17459199, 0.0047680065},
		[10]float32{0.030144334, 0.7798463, -0.18635526, 0.16014743, -0.056352183, -0.073478684, -0.61061764, 0.2639427, -0.29015052, 0.026207846},
		[10]float32{0.107117064, 0.7179576, -0.08471283, 0.13872738, -0.19093458, 0.0032758266, -0.57942724, 0.18091944, -0.37961534, 0.019732641},
		[10]float32{0.18048157, 0.62739754, 0.02287729, 0.108300544, -0.27325797, 0.08379273, -0.5380076, 0.119332194, -0.4355382, -0.007879039},
		[10]float32{0.23640157, 0.5372336, 0.11671838, 0.08722042, -0.29537848, 0.14808217, -0.49740595, 0.08826657, -0.45962888, -0.045178007}},
	[10][10]float32{
		[10]float32{0, 0.25, -0.25, -0.25, 0.5, 0, -0.5, 0.5, 0.25, -0.25},
		[10]float32{-0.0017120001, 0.35513604, -0.249572, -0.148716, 0.489728, -0.047431998, -0.553424, 0.545292, 0.19400801, -0.24786},
		[10]float32{-0.008688001, 0.476064, -0.24420801, -0.047104005, 0.43339202, -0.07972799, -0.617376, 0.565248, 0.112352, -0.23552},
		[10]float32{-0.016308, 0.59892404, -0.22553799, 0.041846007, 0.32061195, -0.084768, -0.682616, 0.54399794, 0.002151981, -0.20922999},
		[10]float32{-0.015872002, 0.697616, -0.18651201, 0.10238401, 0.16668802, -0.057152018, -0.73174405, 0.47779202, -0.12459199, -0.17064},
		[10]float32{0, 0.75, -0.125, 0.125, 0, 0, -0.75, 0.375, -0.25, -0.125},
		[10]float32{0.03412801, 0.74761605, -0.045231983, 0.111104, -0.14843202, 0.07540801, -0.73174405, 0.253952, -0.35715201, -0.07935999},
		[10]float32{0.083692, 0.698924, 0.042922005, 0.07338599, -0.25322804, 0.15215202, -0.682616, 0.13861798, -0.43476802, -0.040769994},
		[10]float32{0.14131203, 0.626064, 0.12683207, 0.031855922, -0.3007681, 0.2123521, -0.617376, 0.05212787, -0.47972804, -0.014479965},
		[10]float32{0.19828792, 0.55513626, 0.19614783, 0.0055642333, -0.29315174, 0.24400775, -0.5534242, 0.008132341, -0.4974319, -0.0021400899}},
	[10][10]float32{
		[10]float32{0, 0.19046399, -0.273024, -0.19046399, 0.46348798, -0.08256, -0.46348798, 0.46348798, 0.19046399, -0.19046399},
		[10]float32{0.03257139, 0.29439062, -0.23459592, -0.09039633, 0.41643226, -0.1081756, -0.5364182, 0.52648556, 0.11719443, -0.22713925},
		[10]float32{0.048199035, 0.40940246, -0.18731134, 0.004201812, 0.32518953, -0.10485715, -0.62372035, 0.5604661, 0.023981858, -0.2603631},
		[10]float32{0.04481542, 0.52230245, -0.12853928, 0.078209, 0.18725027, -0.06259227, -0.7127808, 0.5458129, -0.08755696, -0.27929276},
		[10]float32{0.030144319, 0.61061764, -0.060460884, 0.11681306, 0.02620782, 0.013190076, -0.7798463, 0.4765058, -0.2034818, -0.27302402},
		[10]float32{0.015872002, 0.656336, 0.01192002, 0.113488, -0.12540802, 0.107152015, -0.80476797, 0.36110398, -0.307152, -0.23651202},
		[10]float32{0.013190065, 0.653952, 0.08256003, 0.073650956, -0.23877096, 0.19937307, -0.7798463, 0.22078057, -0.3859006, -0.17333749},
		[10]float32{0.029709315, 0.6128208, 0.14595357, 0.013999008, -0.29687262, 0.2704706, -0.71278083, 0.08440781, -0.43475983, -0.09531828},
		[10]float32{0.06774526, 0.55382115, 0.19850436, -0.042473897, -0.2981106, 0.30651423, -0.6237203, -0.019305952, -0.4572382, -0.018966135},
		[10]float32{0.12397483, 0.5003274, 0.23969036, -0.07577938, -0.25535065, 0.30373386, -0.5364183, -0.07262698, -0.4631486, 0.041210353}},
	[10][10]float32{
		[10]float32{0, 0.114155985, -0.251076, -0.114155985, 0.365232, -0.13692003, -0.365232, 0.365232, 0.114155985, -0.114155985},
		[10]float32{0.061336737, 0.21670559, -0.18122676, -0.015206837, 0.2878736, -0.14457592, -0.45465523, 0.44390917, 0.02707682, -0.18400523},
		[10]float32{0.09479463, 0.32467082, -0.1025856, 0.072941534, 0.17172407, -0.11395113, -0.5617007, 0.49383652, -0.0712993, -0.25779897},
		[10]float32{0.09226217, 0.42479402, -0.018616544, 0.13320765, 0.022328876, -0.037668206, -0.67090195, 0.49347156, -0.17415315, -0.31931844},
		[10]float32{0.060809046, 0.49794242, 0.06124076, 0.15164182, -0.13032255, 0.07006295, -0.7531344, 0.43432817, -0.2665202, -0.34677118},
		[10]float32{0.016307995, 0.530464, 0.12769002, 0.12446199, -0.25215203, 0.18476802, -0.78369206, 0.323386, -0.334768, -0.32553798},
		[10]float32{-0.022238687, 0.52051914, 0.17549041, 0.060770176, -0.3188206, 0.28082418, -0.75313437, 0.18188623, -0.37245822, -0.25509828},
		[10]float32{-0.03766825, 0.47938797, 0.20516248, -0.0190517, -0.3230308, 0.33837014, -0.670902, 0.04051034, -0.38259283, -0.15013337},
		[10]float32{-0.018922076, 0.4277525, 0.22248754, -0.08926566, -0.27530202, 0.34800914, -0.5617006, -0.069056034, -0.37624532, -0.035807487},
		[10]float32{0.03539466, 0.3869554, 0.23580047, -0.12902084, -0.19821931, 0.31330538, -0.45465547, -0.12617354, -0.36757615, 0.06277206}},
	[10][10]float32{
		[10]float32{0, 0.046335887, -0.18841602, -0.046335887, 0.23475191, -0.14208013, -0.23475191, 0.23475191, 0.046335887, -0.046335887},
		[10]float32{0.08079804, 0.14753896, -0.09797736, 0.051988542, 0.13742882, -0.13818192, -0.33541128, 0.3249278, -0.04947182, -0.1391938},
		[10]float32{0.1256215, 0.24834873, -0.0026837345, 0.1365346, 0.008229125, -0.092882365, -0.4559071, 0.39042875, -0.14744769, -0.24298126},
		[10]float32{0.12200379, 0.33475068, 0.08841729, 0.18912615, -0.14062345, -0.0017928928, -0.5788297, 0.40983182, -0.23603764, -0.33853263},
		[10]float32{0.07667351, 0.39020234, 0.16157088, 0.19629508, -0.27530593, 0.115419745, -0.6713948, 0.3723409, -0.30010313, -0.39632282},
		[10]float32{0.008687973, 0.40502387, 0.2065601, 0.155792, -0.3623521, 0.22972804, -0.70579195, 0.2818559, -0.32972807, -0.394208},
		[10]float32{-0.05718276, 0.38121864, 0.22138166, 0.07877691, -0.38271856, 0.31386864, -0.67139477, 0.15697378, -0.32495576, -0.32752833},
		[10]float32{-0.09682213, 0.3327233, 0.21335214, -0.01231017, -0.33796203, 0.34989828, -0.5788297, 0.02691771, -0.29645622, -0.21157038},
		[10]float32{-0.092882514, 0.2810878, 0.19664529, -0.090198755, -0.24852666, 0.3331799, -0.45590702, -0.07660195, -0.26212347, -0.07639121},
		[10]float32{-0.04065049, 0.24658573, 0.18625964, -0.13355619, -0.14414312, 0.2727998, -0.33541158, -0.13159621, -0.23960379, 0.045996185}},
	[10][10]float32{
		[10]float32{0, 0.007704323, -0.09914399, -0.007704323, 0.106848314, -0.091439664, -0.106848314, 0.106848314, 0.007704323, -0.007704323},
		[10]float32{0.08987449, 0.10776295, -6.600667e-05, 0.0907229, 0.0007831049, -0.083084874, -0.21278164, 0.20344624, -0.09157154, -0.11187434},
		[10]float32{0.13969709, 0.20225942, 0.0971344, 0.17586173, -0.13091612, -0.036471277, -0.33959073, 0.28216654, -0.18573368, -0.23269479},
		[10]float32{0.13485503, 0.27592596, 0.178993, 0.22971612, -0.27178913, 0.049693003, -0.4689538, 0.3234382, -0.25990364, -0.35051495},
		[10]float32{0.08196175, 0.31383103, 0.23051976, 0.23873875, -0.38669848, 0.15392515, -0.5663687, 0.3145906, -0.29934612, -0.4317848},
		[10]float32{0.0017120689, 0.3094164, 0.24357973, 0.20042801, -0.44400775, 0.24743192, -0.60256815, 0.25556424, -0.29743192, -0.44957197},
		[10]float32{-0.07822503, 0.26817724, 0.22106537, 0.12560529, -0.4292307, 0.3053643, -0.5663687, 0.16033843, -0.2593183, -0.39558536},
		[10]float32{-0.13027728, 0.2069852, 0.17685944, 0.03637188, -0.35015136, 0.3138638, -0.4689538, 0.054075897, -0.20134583, -0.2837077},
		[10]float32{-0.13400245, 0.1490547, 0.13158886, -0.040253837, -0.23341514, 0.27289897, -0.33959064, -0.03401562, -0.14615268, -0.14503554},
		[10]float32{-0.08308449, 0.11455264, 0.104170285, -0.083018534, -0.11259141, 0.19489314, -0.21278195, -0.08210388, -0.11350607, -0.014428034},
	},
}

func TestPerlin3DAt(t *testing.T) {
	for z := 0; z < 10; z++ {
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				outputPerlin3DAt[x][y][z] = Perlin3DAt(glm.Vec3{float32(x), float32(y), float32(z * 10)}.Slash(10.0))
			}
		}
	}

	for z := 0; z < 10; z++ {
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				if outputPerlin3DAt[x][y][z] != resultsPerlin3DAt[x][y][z] {
					t.Errorf("Wrong result at (%d,%d,%d): %f instead of %f", x, y, z, outputPerlin3DAt[x][y][z], resultsPerlin3DAt[x][y][z])
					return
				}
			}
		}
	}
}

//------------------------------------------------------------------------------
