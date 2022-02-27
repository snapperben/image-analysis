package image_analysis

import "testing"

func TestLoadCSV(t *testing.T) {
	// TODO Test that a CSV gallery can be loaded
}

const CSV1 string = `2020-03-30 14:12:19,40.728808,-73.996106
2020-03-30 14:20:10,40.728656,-73.998790
2020-03-30 14:32:02,40.727160,-73.996044
2020-03-30 14:41:18,40.725468,-73.995701
2020-03-30 14:43:48,40.721983,-74.003248
2020-03-30 14:49:23,40.721332,-74.004535
2020-03-30 15:37:13,40.712115,-74.005764
2020-03-30 15:42:00,40.711887,-74.006150
2020-03-30 15:51:40,40.710944,-74.002933
2020-03-30 17:01:19,40.706225,-74.009415
2020-03-30 18:10:01,40.706485,-74.010359
2020-03-30 18:12:37,40.706063,-74.010230
2020-03-30 18:19:40,40.705882,-74.010093
2020-03-31 11:19:19,40.702006,-74.015375
2020-03-31 11:28:34,40.702347,-74.015847
2020-03-31 11:33:53,40.702347,-74.015847
2020-03-31 11:48:12,40.702591,-74.016426
2020-03-31 11:52:50,40.701420,-74.015531
2020-03-31 12:02:19,40.701257,-74.015016
2020-03-31 14:01:03,40.768538,-73.977133
2020-03-31 14:23:19,40.768559,-73.975680
2020-03-31 14:36:19,40.768559,-73.975680
2020-03-31 14:49:28,40.771647,-73.972676
2020-03-31 15:01:39,40.778417,-73.967884
2020-03-31 15:12:12,40.779487,-73.963623
2020-03-31 15:19:45,40.789947,-73.960867
2020-03-31 16:12:29,40.796404,-73.951393`

const CSV2 string = `2019-12-01T05:16:45Z,36.102825,-115.173813
2019-12-01T03:45:33Z,36.096181,-115.175278
2019-12-01T03:45:25Z,36.096183,-115.175270
2019-11-30T19:30:02Z,36.103417,-115.173888
2019-11-30T19:28:48Z,36.103425,-115.173997
2019-11-30T08:20:10Z,36.142500,-115.160500
2019-11-30T07:48:45Z,36.142461,-115.160630
2019-11-30T07:48:32Z,36.142445,-115.160645
2019-11-30T07:37:33Z,36.142500,-115.161100
2019-11-30T06:21:23Z,36.142600,-115.160500
2019-11-30T06:20:13Z,36.142600,-115.160500
2019-11-30T06:19:49Z,36.140800,-115.159400
2019-11-30T06:18:00Z,36.142175,-115.160645
2019-11-30T05:36:34Z,36.142000,-115.160800
2019-11-30T05:22:54Z,36.142400,-115.161811
2019-11-30T05:22:47Z,36.142508,-115.161750
2019-11-30T05:22:43Z,36.142422,-115.161820
2019-11-30T05:22:02Z,36.142342,-115.161842
2019-11-30T05:22:00Z,36.142364,-115.161842
2019-11-30T05:21:48Z,36.142400,-115.161789
2019-11-30T05:21:46Z,36.142525,-115.161728
2019-11-30T05:21:44Z,36.143397,-115.163217
2019-11-30T05:19:10Z,36.142753,-115.162025
2019-11-30T05:11:37Z,36.142600,-115.160800
2019-11-30T05:11:29Z,36.142581,-115.160836
2019-11-30T05:11:24Z,36.142505,-115.161050
2019-11-30T05:11:23Z,36.142505,-115.161050
2019-11-30T05:10:34Z,36.142472,-115.161103
2019-11-30T05:08:18Z,36.142119,-115.161056
2019-11-30T05:08:11Z,36.141525,-115.161086
2019-11-30T05:08:10Z,36.140942,-115.160661
2019-11-30T05:06:01Z,36.142000,-115.160800
2019-11-30T05:05:41Z,36.142025,-115.160775
2019-11-30T05:05:37Z,36.142039,-115.160781
2019-11-30T05:05:36Z,36.142047,-115.160781
2019-11-30T05:05:33Z,36.141997,-115.160805
2019-11-30T05:02:12Z,36.141933,-115.160256
2019-11-29T20:59:47Z,36.122192,-115.167938
2019-11-29T20:52:55Z,36.121670,-115.167708
2019-11-29T20:51:47Z,36.122192,-115.167938
2019-11-29T20:49:20Z,36.122192,-115.167938
2019-11-29T18:48:14Z,36.105930,-115.179558
2019-11-29T18:48:03Z,36.105930,-115.179567
2019-11-29T18:47:59Z,36.105914,-115.179572
2019-11-29T18:47:38Z,36.105975,-115.179550
2019-11-29T18:44:21Z,36.105995,-115.179572
2019-11-29T05:53:01Z,36.096030,-115.174653
2019-11-28T03:13:55Z,36.122550,-115.169328
2019-11-28T03:13:51Z,36.122528,-115.169220
2019-11-28T03:13:48Z,36.122528,-115.169220
2019-11-28T02:59:56Z,36.122528,-115.169220
2019-11-28T02:59:45Z,36.122528,-115.169220
2019-11-27T03:08:10Z,36.122830,-115.166169
2019-11-27T03:08:07Z,36.122830,-115.166169
2019-11-27T02:00:11Z,36.122711,-115.166420
2019-11-26T21:03:04Z,36.122192,-115.167938
2019-11-26T21:02:51Z,36.122192,-115.167938
2019-11-26T01:22:23Z,36.125095,-115.170280
2019-11-26T01:22:15Z,36.125003,-115.170197
2019-11-25T23:07:49Z,36.126700,-115.166862
2019-11-25T23:06:52Z,36.126489,-115.166550
2019-11-25T23:06:38Z,36.126403,-115.166450
2019-11-25T22:37:48Z,36.125908,-115.168647
2019-11-25T22:37:47Z,36.125897,-115.168647
2019-11-25T22:37:46Z,36.125892,-115.168655
2019-11-25T22:37:45Z,36.125900,-115.168655
2019-11-25T22:37:44Z,36.125900,-115.168655
2019-11-25T22:37:44Z,36.125900,-115.168655
2019-11-25T22:37:21Z,36.126200,-115.169000
2019-11-25T22:36:47Z,36.125592,-115.168747
2019-11-25T22:36:44Z,36.125706,-115.168738
2019-11-25T22:04:39Z,36.122895,-115.165620
2019-11-25T21:44:37Z,36.122928,-115.165680
2019-11-25T21:44:35Z,36.122928,-115.165680
2019-11-25T21:44:21Z,36.122875,-115.165992
2019-11-25T21:44:11Z,36.122875,-115.165992
2019-11-25T21:43:58Z,36.122875,-115.165992
2019-11-25T21:39:17Z,36.122222,-115.170517
2019-11-25T20:59:52Z,36.118942,-115.176047
2019-11-25T19:49:19Z,36.102986,-115.173347
2019-11-25T19:49:18Z,36.103000,-115.173247
2019-11-25T19:02:50Z,36.095481,-115.175653
2019-11-25T18:57:38Z,36.095858,-115.175287`

const CSV3 string = `2019-10-31T18:52:59Z,40.627883,14.366858
2019-10-31T18:52:52Z,40.627883,14.366858
2019-10-31T18:35:23Z,40.627808,14.364933
2019-10-31T18:26:12Z,40.628197,14.367075
2019-10-31T12:17:17Z,40.634303,14.602580
2019-10-31T12:08:10Z,40.635808,14.601962
2019-10-31T12:07:51Z,40.635708,14.601955
2019-10-31T12:07:41Z,40.635583,14.602013
2019-10-31T11:27:13Z,40.632400,14.603013
2019-10-31T11:26:50Z,40.632347,14.603062
2019-10-31T11:26:42Z,40.632350,14.603067
2019-10-31T11:26:35Z,40.632338,14.603063
2019-10-31T11:26:15Z,40.632330,14.603028
2019-10-31T11:25:18Z,40.632358,14.603080
2019-10-31T11:25:05Z,40.632325,14.603017
2019-10-31T11:24:57Z,40.632337,14.603038
2019-10-31T11:24:55Z,40.632328,14.603042
2019-10-31T11:24:54Z,40.632337,14.603033
2019-10-31T11:24:50Z,40.632347,14.603025
2019-10-31T11:24:49Z,40.632347,14.603017
2019-10-31T11:24:47Z,40.632347,14.603012
2019-10-31T11:16:12Z,40.631775,14.595912
2019-10-31T11:13:15Z,40.626672,14.584980
2019-10-31T11:10:47Z,40.621478,14.579305
2019-10-31T10:50:16Z,40.612258,14.522980
2019-10-31T10:47:56Z,40.615363,14.523095
2019-10-31T10:47:51Z,40.615525,14.523222
2019-10-31T10:47:11Z,40.616317,14.521728
2019-10-31T10:47:04Z,40.616128,14.520967
2019-10-31T10:47:03Z,40.616037,14.520895
2019-10-31T10:28:56Z,40.630163,14.487425
2019-10-31T10:28:52Z,40.630220,14.487412
2019-10-31T10:28:45Z,40.630583,14.487155
2019-10-31T10:28:35Z,40.630875,14.487178
2019-10-31T10:28:13Z,40.631053,14.486542
2019-10-31T10:28:09Z,40.630938,14.486470
2019-10-31T10:28:02Z,40.630170,14.485288
2019-10-31T10:26:18Z,40.633113,14.483203
2019-10-31T10:26:18Z,40.633113,14.483203
2019-10-31T10:26:04Z,40.632645,14.483420
2019-10-31T10:25:36Z,40.631663,14.482542
2019-10-31T10:25:33Z,40.631683,14.482247
2019-10-31T10:25:32Z,40.631683,14.482247
2019-10-31T10:25:32Z,40.631553,14.482178
2019-10-31T10:25:04Z,40.630517,14.481880
2019-10-31T10:24:15Z,40.629528,14.482303
2019-10-31T10:24:10Z,40.629475,14.482275
2019-10-31T10:23:52Z,40.629538,14.481422
2019-10-31T10:21:40Z,40.629862,14.477763
2019-10-31T10:21:35Z,40.629638,14.477472
2019-10-31T10:21:15Z,40.628680,14.478628
2019-10-31T10:21:06Z,40.628087,14.479083
2019-10-31T10:21:04Z,40.627983,14.479192
2019-10-31T10:21:01Z,40.627750,14.479238
2019-10-31T10:20:59Z,40.627597,14.479225
2019-10-31T10:19:06Z,40.624070,14.471063
2019-10-31T10:19:06Z,40.624070,14.471063
2019-10-31T10:18:54Z,40.623475,14.471497
2019-10-31T10:18:50Z,40.623287,14.471275
2019-10-31T10:18:49Z,40.623178,14.471262
2019-10-31T10:18:45Z,40.622958,14.471005
2019-10-31T10:18:21Z,40.621963,14.469878
2019-10-31T10:18:19Z,40.621925,14.469797
2019-10-31T10:18:07Z,40.621762,14.468628
2019-10-31T10:17:30Z,40.621637,14.465778
2019-10-31T10:15:45Z,40.619737,14.459608
2019-10-31T10:15:43Z,40.619853,14.459633
2019-10-31T10:15:41Z,40.619970,14.459542
2019-10-31T10:15:40Z,40.620712,14.457122
2019-10-31T10:13:32Z,40.619870,14.449938
2019-10-31T10:13:22Z,40.619672,14.449180
2019-10-31T10:12:08Z,40.618062,14.440650
2019-10-31T10:12:07Z,40.618062,14.440650
2019-10-31T10:11:21Z,40.617533,14.440333
2019-10-31T10:11:20Z,40.617455,14.440303
2019-10-31T10:11:14Z,40.617408,14.439872
2019-10-31T10:11:12Z,40.617520,14.439530
2019-10-31T10:11:02Z,40.620425,14.459792
2019-10-31T10:11:00Z,40.620425,14.459792
2019-10-31T10:08:26Z,40.616600,14.433362
2019-10-31T10:08:23Z,40.616455,14.433200
2019-10-31T10:08:10Z,40.616225,14.432038
2019-10-31T10:08:04Z,40.616230,14.431575
2019-10-30T19:30:28Z,0.000000,0.000000
2019-10-30T19:25:32Z,0.000000,0.000000
2019-10-30T13:21:49Z,40.659725,14.424322
2019-10-30T13:21:49Z,40.659725,14.424322
2019-10-30T12:08:11Z,40.661950,14.427728
2019-10-30T12:08:07Z,40.662022,14.427753
2019-10-30T12:08:04Z,40.662033,14.427722
2019-10-30T12:08:02Z,40.662033,14.427722
2019-10-30T11:54:22Z,40.662012,14.427867
2019-10-30T11:54:10Z,40.662003,14.427553
2019-10-30T11:53:56Z,40.661983,14.427775
2019-10-30T11:53:33Z,40.661970,14.427733
2019-10-30T11:36:17Z,40.661828,14.424942
2019-10-30T11:29:37Z,40.663692,14.423875
2019-10-30T11:29:21Z,40.663708,14.423825
2019-10-30T11:29:12Z,40.663637,14.423758
2019-10-30T11:21:42Z,40.663772,14.428003
2019-10-30T11:21:34Z,40.663778,14.427988
2019-10-30T11:21:24Z,40.663778,14.427992
2019-10-30T11:21:23Z,40.663778,14.427987
2019-10-30T11:21:11Z,40.663795,14.428000
2019-10-30T11:21:06Z,40.663792,14.427997
2019-10-30T11:19:40Z,40.663345,14.428725
2019-10-29T15:45:13Z,40.624700,14.376800
2019-10-29T13:48:52Z,40.624700,14.377000
2019-10-29T13:48:37Z,40.624700,14.376900
2019-10-29T13:44:21Z,40.624600,14.377000
2019-10-29T11:15:05Z,40.627758,14.375808
2019-10-29T11:15:04Z,40.627762,14.375813
2019-10-29T11:15:03Z,40.627755,14.375817
2019-10-29T11:11:59Z,40.628075,14.375383
2019-10-29T11:11:57Z,40.628063,14.375378
2019-10-29T11:09:06Z,40.627013,14.375920
2019-10-29T11:08:57Z,40.626913,14.375895
2019-10-29T11:08:47Z,40.626870,14.375925
2019-10-29T11:08:41Z,40.626878,14.375947
2019-10-28T21:41:04Z,40.626133,14.375692
2019-10-28T21:40:56Z,40.626112,14.375658
2019-10-28T21:40:36Z,40.626100,14.375687
2019-10-28T21:40:34Z,40.625983,14.375795
2019-10-28T21:37:42Z,40.626053,14.375650
2019-10-28T21:21:12Z,40.625503,14.372938
2019-10-28T21:21:09Z,40.625512,14.372937
2019-10-28T21:20:31Z,40.625428,14.372837
2019-10-28T21:20:31Z,40.625428,14.372837
2019-10-28T21:20:14Z,40.625492,14.372795
2019-10-28T21:17:45Z,40.624787,14.371995
2019-10-28T17:01:55Z,40.627980,14.366945
2019-10-28T17:01:24Z,40.627937,14.366975
2019-10-28T17:01:19Z,40.627962,14.366883
2019-10-28T16:21:54Z,40.627922,14.366963
2019-10-28T16:21:00Z,40.627955,14.366917
2019-10-28T16:20:56Z,40.627933,14.366928
2019-10-28T16:18:08Z,40.627955,14.367005
2019-10-28T13:36:17Z,40.627500,14.373275
2019-10-28T13:36:16Z,40.627500,14.373275
2019-10-28T13:10:23Z,40.627445,14.373328
2019-10-28T13:10:19Z,40.627475,14.373280
2019-10-28T13:10:15Z,40.627453,14.373262
2019-10-27T19:35:02Z,40.624813,14.376455
2019-10-27T19:34:57Z,40.624803,14.376475
2019-10-27T19:34:52Z,40.624788,14.376450
2019-10-27T13:30:29Z,40.627987,14.367008
2019-10-27T13:30:28Z,40.628012,14.366950
2019-10-27T13:27:58Z,40.627863,14.366958
2019-10-27T13:17:24Z,40.628005,14.367228
2019-10-27T13:13:09Z,40.627462,14.369295
2019-10-27T13:03:34Z,40.627725,14.373042
2019-10-27T13:03:32Z,40.627717,14.373030
2019-10-27T12:58:35Z,40.627272,14.373583
2019-10-27T12:58:26Z,40.627392,14.373733
2019-10-27T12:33:51Z,40.625988,14.378670
2019-10-27T12:33:44Z,40.626075,14.378703
2019-10-27T12:33:40Z,40.626045,14.378597`
