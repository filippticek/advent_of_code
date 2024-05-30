package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type Map struct {
	Destination int64
	Source      int64
	Length      int64
}

type Range struct {
	Start  int64
	Length int64
}

func main() {
	maxProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(maxProcs)
	part2(test2)
}

func part2(input string) {
	lines := strings.Split(input, "\n")

	seedsStr := strings.Split(strings.Split(lines[0], ":")[1], " ")[1:]
	seeds := make([]Range, 0)
	for i := 0; i < len(seedsStr); i += 2 {
		seeds = append(seeds, Range{Start: getInt(seedsStr[i]), Length: getInt(seedsStr[i+1])})
	}

	lineNum := 3
	slice := seeds
	for i := 0; i < 7; i++ {
		// fmt.Println(i, "#####################")
		mapSlice := getMapSlice(lines, &lineNum)
		slice = mapToRangeSlice(slice, mapSlice)
		// fmt.Println(slice)

		lineNum += 2
	}

	min := slice[0].Start
	for _, num := range slice {
		if num.Start < min {
			min = num.Start
		}
	}
	fmt.Println(min)
}

func mapToRangeSlice(init []Range, mapSlice []Map) []Range {
	dest := make([]Range, 0)
	// fmt.Println("Init", init)
	for _, rang := range init {
		notMapped := []Range{rang}
		for _, m := range mapSlice {
			csTmp := []Range{}
			for _, nm := range notMapped {
				r1, r2, r3 := getCrossSection(m, nm)
				// fmt.Println(r1, r2, r3)
				if r1 != nil {
					csTmp = append(csTmp, *r1)
				}
				if r2 != nil {
					dest = append(dest, *r2)
				}
				if r3 != nil {
					csTmp = append(csTmp, *r3)
				}
			}
			notMapped = csTmp
			// fmt.Println("Not mapped", notMapped)
		}
		dest = append(dest, notMapped...)
	}
	return dest
}

func getCrossSection(m Map, rang Range) (*Range, *Range, *Range) {
	rFinish := rang.Start + rang.Length - 1
	mFinish := m.Source + m.Length - 1
	// fmt.Println("Getting", m, rang)
	switch {
	case rang.Start > mFinish || rFinish < m.Source:
		return &rang, nil, nil
	case rang.Start < m.Source && rFinish > mFinish:
		r1 := Range{Start: rang.Start, Length: m.Source - rang.Start}
		r2 := Range{Start: m.Destination, Length: m.Length}
		r3 := Range{Start: mFinish + 1, Length: rFinish - mFinish}
		return &r1, &r2, &r3
	case rang.Start < m.Source && rFinish >= m.Source:
		r1 := Range{Start: rang.Start, Length: m.Source - rang.Start}
		r2 := Range{Start: m.Destination, Length: rFinish - m.Source + 1}
		return &r1, &r2, nil
	case rang.Start >= m.Source && rFinish <= mFinish:
		r1 := Range{Start: m.Destination + (rang.Start - m.Source), Length: rang.Length}
		return nil, &r1, nil
	case rang.Start <= mFinish && rFinish > mFinish:
		r1 := Range{Start: m.Destination + (rang.Start - m.Source), Length: mFinish - rang.Start + 1}
		r2 := Range{Start: mFinish + 1, Length: rFinish - mFinish}
		return nil, &r1, &r2
	}
	return nil, nil, nil

}

func part1(input string) {
	lines := strings.Split(input, "\n")

	seedsStr := strings.Split(strings.Split(lines[0], ":")[1], " ")
	seeds := make([]int64, len(seedsStr)-1)
	for i, seed := range seedsStr[1:] {
		seeds[i] = getInt(seed)
	}
	lineNum := 3
	slice := seeds
	for i := 0; i < 7; i++ {
		mapSlice := getMapSlice(lines, &lineNum)
		slice = mapToSlice(slice, mapSlice)
		fmt.Println(slice)

		lineNum += 2
	}

	min := slice[0]
	for _, num := range slice {
		if num < min {
			min = num
		}
	}
	fmt.Println(min)
}

func mapToSlice(init []int64, mapSlice []Map) []int64 {
	dest := make([]int64, len(init))
	copy(dest, init)
	var wg sync.WaitGroup
	for i := range mapSlice {
		wg.Add(1)
		go func(m Map) {
			for i, num := range init {
				if num >= m.Source && num < m.Source+m.Length {
					dest[i] = (num - m.Source) + m.Destination
				}
			}
			wg.Done()
		}(mapSlice[i])
	}
	wg.Wait()
	return dest
}

func getMapSlice(lines []string, lineNum *int) []Map {
	mapSlice := []Map{}

	for ; *lineNum < len(lines) && lines[*lineNum] != ""; (*lineNum)++ {
		nums := strings.Split(lines[*lineNum], " ")
		// fmt.Println(nums)
		mapSlice = append(mapSlice, Map{
			Destination: getInt(nums[0]),
			Source:      getInt(nums[1]),
			Length:      getInt(nums[2]),
		})

	}
	return mapSlice
}

func getInt(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return n
	}
	fmt.Println(s)

	panic(err)
}

var test1 = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

var test2 = `seeds: 1310704671 312415190 1034820096 106131293 682397438 30365957 2858337556 1183890307 665754577 13162298 2687187253 74991378 1782124901 3190497 208902075 226221606 4116455504 87808390 2403629707 66592398

seed-to-soil map:
2879792625 0 201678008
2425309256 1035790247 296756276
2722065532 1759457739 157727093
400354950 1917184832 1164285801
0 201678008 400354950
1564640751 602032958 433757289
1998398040 1332546523 426911216

soil-to-fertilizer map:
3434127746 3670736129 29685965
1809924203 1168707872 308179
2108903682 1437989162 44479258
237181023 2915565442 27901445
1173998623 2434447796 13633544
75539025 740516241 29278225
41104738 706081954 34434287
3279397405 3488165796 12149874
3463813711 3827946213 157129363
1810232382 769794466 15695437
877824710 677909236 28172718
2215709448 1746651561 307558709
1825927819 1692597620 54053941
104817250 420198730 132363773
2916210208 392942051 27256679
1022591555 2448081340 151407068
3925105941 3985075576 182313682
1897186025 2212065968 211717657
2198981202 1304666789 16728246
850656807 2054210270 27167903
3766599721 3500315670 158506220
3419071398 3279397405 15056348
7830088 2126976435 33274650
3620943074 3658821890 11914239
1264213180 2599488408 138420934
811586355 2160251085 12020898
3632857313 3354423388 133742408
1612763314 1169016051 108601184
1721364498 2172271983 39793985
1187632167 601328223 76581013
823607253 1277617235 27049554
728944387 2737909342 82641968
0 2426617708 7830088
3291547279 3700422094 127524119
1402634114 1482468420 210129200
905997428 1321395035 107714902
4107419623 3294453753 59969635
1879981760 785489903 17204265
2153382940 2081378173 45598262
277361019 802694168 366013704
1761158483 552562503 48765720
646208806 2832829861 82735581
2523268157 0 392942051
1013712330 1429109937 8879225
643374723 2423783625 2834083
265082468 2820551310 12278551

fertilizer-to-water map:
4253122607 1473424614 41844689
3040447798 2659805568 46237011
0 146022665 42081460
55436822 188104125 65067713
42081460 132667303 13355362
2429043181 3587614447 54605699
888256662 672288214 24436041
4064969883 1978094070 95324589
3086684809 977403736 339965972
120504535 253171838 93494065
2810558403 2603914183 55891385
3898695123 2901215107 166274760
2483648880 4002918707 103777141
1300545784 2848997109 52217998
2418717938 1463099371 10325243
1022681665 808998429 30429585
2866449788 1411682577 4750813
1181605510 4172708724 118940274
2078503930 2466708865 42530000
1105548530 1545561518 76056980
978705579 2573458117 30456066
2324405069 1317369708 94312869
1991848966 3429793336 22435712
4190586687 2706042579 43180396
1352763782 1416433390 46665981
3760606255 1683093685 138088868
1399429763 3452229048 135385399
2121033930 839428014 137975722
2940673664 2749222975 99774134
1053111250 2073418659 52437280
3426650781 1821182553 152991287
1534815162 2195329002 252024339
730962658 3067489867 157294004
3579642068 710244275 98754154
1786839501 3224783871 205009465
2259009652 1974173840 3920230
2587426021 370264097 223132382
2871200601 2125855939 69473063
213998600 44701447 87965856
4233767083 2447353341 19355524
2262929882 1621618498 61475187
1009161645 696724255 13520020
3678396222 593396479 78891735
912692703 4106695848 66012876
3757287957 4291648998 3318298
301964456 0 44701447
2014284678 2509238865 64219252
370264097 3642220146 360698561
4160294472 1515269303 30292215

water-to-light map:
4066036887 2992193346 95912236
531075515 493316918 162009008
3260565192 854248031 437396028
1341316194 4205924684 89042612
1879858967 2058162578 692895326
452475911 655325926 78599604
2997176790 1690328655 208783332
2731804884 3324847814 265371906
355611136 0 96864775
2572754293 1899111987 159050591
1081338600 3590219720 138271571
1430358806 2779435417 212757929
3234337635 4179697127 26227557
854248031 3728491291 227090569
4161949123 3955581860 102409244
3205960122 2751057904 28377513
50952557 147817332 304658579
1219610171 4057991104 121706023
4264358367 1291644059 30608929
3697961220 1322252988 368075667
1643116735 3088105582 236742232
693084523 452475911 40841007
0 96864775 50952557

light-to-temperature map:
2756401132 2384899493 13749631
1163093625 0 117407544
3603435593 3599927411 262731037
2081436411 2089913126 119300659
693703633 117407544 395383894
1672621164 1405157690 24997208
3873714258 2780774148 107551276
3355072403 2593861641 186912507
1953100586 3862658448 62069331
143286272 672639421 194814248
1562062673 1010739941 110558491
2869050867 2888325424 31673634
3159859886 2398649124 195212517
2900724501 3298674599 34708838
2243940568 4059045429 56605170
691405879 1193483066 2297754
2300545738 2005749676 25248062
3541984910 3924727779 61450683
2200737070 3986178462 43203498
3981265534 2030997738 58915388
2530829166 4276276595 18690701
621411866 641250212 31389209
1784026205 4037549491 21495938
1519774068 1362869085 42288605
3866166630 3584674072 7547628
652801075 1430154898 38604804
2015169917 4029381960 8167531
2770150763 2936555750 98900104
1813227854 2258880377 123316040
3032290681 1784026205 127569205
0 867453669 143286272
1805522143 3592221700 7705711
4040180922 3043888225 254786374
2023337448 3035455854 8432371
3029587605 2382196417 2703076
392553196 1468759702 228858670
2710145863 3538418803 46255269
1089087527 567244114 74006098
2325793800 3333383437 205035366
2549519867 4115650599 160625996
338100520 512791438 54452676
2935433339 1911595410 94154266
1280501169 1121298432 72184634
1352685803 1195780820 167088265
2031769819 2209213785 49666592
1936543894 2919999058 16556692

temperature-to-humidity map:
1606220966 2958863752 268926464
2994413958 1467440292 348583188
1347324773 3453966662 171497865
3342997146 3227790216 188948930
0 211826810 113744983
1875147430 1816023480 774831860
699941162 0 211826810
443679044 325571793 256262118
3531946076 1280528675 186911617
1280528675 4228171198 66796098
113744983 581833911 329934061
1518822638 2590855340 50170812
1568993450 3416739146 37227516
2967816890 4201574130 26597068
3718857693 3625464527 576109603
2649979290 2641026152 317837600

humidity-to-location map:
3244927 955737016 9389705
380524056 2531586403 38604778
3713586211 965126721 158937945
3122843287 1406574654 236795236
776685423 1643369890 534268825
2053493196 0 55930434
582662115 695344450 194023308
3885666529 3855399097 320692779
88096722 283368340 98672354
1901561222 3703467123 151931974
1317500428 2570191181 151780331
3872524156 3690324750 13142373
2109423630 249685414 30437999
1310954248 4199813128 6546180
1751790747 382040694 149770475
3056474029 889367758 66369258
2139861629 4176091876 23721252
12634632 2721971512 75462090
186769076 55930434 193754980
419128834 531811169 163533281
3359638523 2177638715 353947688
2163582881 2797433602 892891148
1469280759 1124064666 282509988
0 280123413 3244927`
