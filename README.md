# fs2util
FreeSpace 2 util for command line

* Extract files inside `.VP` file
* Get file checksums inside `.VP` file
* Show file information inside `.VP` file

# Requirements
* Operating system
  * GNU/Linux: x64 arm arm64 ppc64 ppc64le
  * Microsoft Windows: x64
  * Darwin (Apple Mac): x64
  * FreeBSD: x64 arm
  * NetBSD: x64 arm
  * OpenBSD: x64 arm arm64
  * Other OSes supported by Go, for full list, see: https://golang.org/doc/install/source#environment

# Usage

```
fs2util - Inspect FreeSpace 2 .VP file
Version v0.0.3 a9bd85ce00df0215e17775c5cd9ecaa10ca38edd 2021-01-24T14:21:04+02:00
(c) Pekka Järvinen 2021- [ https://github.com/raspi/fs2util ]

Parameters:
  -algo   Checksum algorithm: md5, sha1, sha256, sha512   default: "sha256"
  -c      Calculate checksums inside the file   default: "false"
  -edir   Extract directory name   default: "extracted"
  -i      Display information about files inside the .vp file   default: "false"
  -x      Extract files   default: "false"

List file checksums:
  ./fs2util -c <file.vp>
Extract file(s):
  ./fs2util -x <file.vp>

Examples:
  Show checksums:
    ./fs2util -c root_fs2.vp
  Extract files:
    ./fs2util -x root_fs2.vp
```

# Example

## Extract files from `.vp`:

```
% ./fs2util -x Root_fs2.vp
Extracted file data/tables/ai.tbl size:1874 bytes, offset:16 (1999-10-28 07:14:14 +0000 UTC) to extracted/Root_fs2.vp/data/tables/ai.tbl
Extracted file data/tables/asteroid.tbl size:4309 bytes, offset:1890 (1999-10-28 07:14:14 +0000 UTC) to extracted/Root_fs2.vp/data/tables/asteroid.tbl
Extracted file data/tables/credits.tbl size:8542 bytes, offset:6199 (1999-11-01 20:09:40 +0000 UTC) to extracted/Root_fs2.vp/data/tables/credits.tbl
Extracted file data/tables/cutscenes.tbl size:2402 bytes, offset:14741 (1999-10-28 07:14:14 +0000 UTC) to extracted/Root_fs2.vp/data/tables/cutscenes.tbl
Extracted file data/tables/fireball.tbl size:531 bytes, offset:17143 (1999-10-28 07:14:14 +0000 UTC) to extracted/Root_fs2.vp/data/tables/fireball.tbl
Extracted file data/tables/help.tbl size:14099 bytes, offset:17674 (1999-10-28 07:14:14 +0000 UTC) to extracted/Root_fs2.vp/data/tables/help.tbl
Extracted file data/tables/hud.tbl size:800 bytes, offset:31773 (1999-10-28 07:14:14 +0000 UTC) to extracted/Root_fs2.vp/data/tables/hud.tbl
Extracted file data/tables/icons.tbl size:14261 bytes, offset:32573 (1999-10-28 07:14:14 +0000 UTC) to extracted/Root_fs2.vp/data/tables/icons.tbl
Extracted file data/tables/launchhelp.tbl size:4374 bytes, offset:46834 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/launchhelp.tbl
Extracted file data/tables/lightning.tbl size:5788 bytes, offset:51208 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/lightning.tbl
Extracted file data/tables/mainhall.tbl size:18121 bytes, offset:56996 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/mainhall.tbl
Extracted file data/tables/medals.tbl size:1947 bytes, offset:75117 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/medals.tbl
Extracted file data/tables/menu.tbl size:1962 bytes, offset:77064 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/menu.tbl
Extracted file data/tables/messages.tbl size:34028 bytes, offset:79026 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/messages.tbl
Extracted file data/tables/mflash.tbl size:597 bytes, offset:113054 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/mflash.tbl
Extracted file data/tables/music.tbl size:7287 bytes, offset:113651 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/music.tbl
Extracted file data/tables/nebula.tbl size:440 bytes, offset:120938 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/nebula.tbl
Extracted file data/tables/pixels.tbl size:81 bytes, offset:121378 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/pixels.tbl
Extracted file data/tables/rank.tbl size:3397 bytes, offset:121459 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/rank.tbl
Extracted file data/tables/ships.tbl size:383719 bytes, offset:124856 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/ships.tbl
Extracted file data/tables/sounds.tbl size:20735 bytes, offset:508575 (1999-11-17 20:13:26 +0000 UTC) to extracted/Root_fs2.vp/data/tables/sounds.tbl
Extracted file data/tables/Species.tbl size:16708 bytes, offset:529310 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/Species.tbl
Extracted file data/tables/ssm.tbl size:45 bytes, offset:546018 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/ssm.tbl
Extracted file data/tables/stars.tbl size:1910 bytes, offset:546063 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/stars.tbl
Extracted file data/tables/strings.tbl size:235241 bytes, offset:547973 (1999-11-02 16:27:22 +0000 UTC) to extracted/Root_fs2.vp/data/tables/strings.tbl
Extracted file data/tables/tips.tbl size:4390 bytes, offset:783214 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/tips.tbl
Extracted file data/tables/traitor.tbl size:665 bytes, offset:787604 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/traitor.tbl
Extracted file data/tables/tstrings.tbl size:984199 bytes, offset:788269 (1999-11-11 21:07:26 +0000 UTC) to extracted/Root_fs2.vp/data/tables/tstrings.tbl
Extracted file data/tables/weapon_expl.tbl size:228 bytes, offset:1772468 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/weapon_expl.tbl
Extracted file data/tables/weapons.tbl size:259791 bytes, offset:1772696 (1999-10-28 07:14:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/weapons.tbl
Extracted file data/tables/vssver.scc size:512 bytes, offset:2032487 (1999-11-02 03:55:16 +0000 UTC) to extracted/Root_fs2.vp/data/tables/vssver.scc
Extracted file data/missions/G-Shi.fs2 size:43642 bytes, offset:2032999 (1999-11-02 03:57:28 +0000 UTC) to extracted/Root_fs2.vp/data/missions/G-Shi.fs2
Extracted file data/missions/G-Ter.fs2 size:43671 bytes, offset:2076641 (1999-11-02 03:57:28 +0000 UTC) to extracted/Root_fs2.vp/data/missions/G-Ter.fs2
Extracted file data/missions/G-Vas.fs2 size:42306 bytes, offset:2120312 (1999-11-02 03:57:28 +0000 UTC) to extracted/Root_fs2.vp/data/missions/G-Vas.fs2
Extracted file data/missions/loop1-1.fs2 size:59582 bytes, offset:2162618 (1999-11-02 03:57:28 +0000 UTC) to extracted/Root_fs2.vp/data/missions/loop1-1.fs2
Extracted file data/missions/loop1-2.fs2 size:114569 bytes, offset:2222200 (1999-11-02 03:57:28 +0000 UTC) to extracted/Root_fs2.vp/data/missions/loop1-2.fs2
Extracted file data/missions/loop1-3.fs2 size:58172 bytes, offset:2336769 (1999-11-02 03:57:28 +0000 UTC) to extracted/Root_fs2.vp/data/missions/loop1-3.fs2
Extracted file data/missions/loop2-1.fs2 size:61415 bytes, offset:2394941 (1999-11-02 03:57:28 +0000 UTC) to extracted/Root_fs2.vp/data/missions/loop2-1.fs2
Extracted file data/missions/loop2-2.fs2 size:84727 bytes, offset:2456356 (1999-11-02 03:57:28 +0000 UTC) to extracted/Root_fs2.vp/data/missions/loop2-2.fs2
Extracted file data/missions/M-01.fs2 size:31309 bytes, offset:2541083 (1999-11-02 03:57:28 +0000 UTC) to extracted/Root_fs2.vp/data/missions/M-01.fs2
Extracted file data/missions/M-02.fs2 size:41218 bytes, offset:2572392 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/M-02.fs2
Extracted file data/missions/M-03.fs2 size:59154 bytes, offset:2613610 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/M-03.fs2
Extracted file data/missions/M-04.fs2 size:38628 bytes, offset:2672764 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/M-04.fs2
Extracted file data/missions/MDH-01.fs2 size:17922 bytes, offset:2711392 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDH-01.fs2
Extracted file data/missions/MDH-02.fs2 size:17138 bytes, offset:2729314 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDH-02.fs2
Extracted file data/missions/MDH-03.fs2 size:17765 bytes, offset:2746452 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDH-03.fs2
Extracted file data/missions/MDH-04.fs2 size:17627 bytes, offset:2764217 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDH-04.fs2
Extracted file data/missions/MDH-05.fs2 size:19905 bytes, offset:2781844 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDH-05.fs2
Extracted file data/missions/MDH-06.fs2 size:17749 bytes, offset:2801749 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDH-06.fs2
Extracted file data/missions/MDH-07.fs2 size:13682 bytes, offset:2819498 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDH-07.fs2
Extracted file data/missions/mdH-08.fs2 size:15115 bytes, offset:2833180 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/mdH-08.fs2
Extracted file data/missions/mdH-09.fs2 size:17542 bytes, offset:2848295 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/mdH-09.fs2
Extracted file data/missions/MDL-01.fs2 size:17791 bytes, offset:2865837 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDL-01.fs2
Extracted file data/missions/MDL-02.fs2 size:16993 bytes, offset:2883628 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDL-02.fs2
Extracted file data/missions/MDL-03.fs2 size:17867 bytes, offset:2900621 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDL-03.fs2
Extracted file data/missions/MDL-04.fs2 size:17745 bytes, offset:2918488 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDL-04.fs2
Extracted file data/missions/MDL-05.fs2 size:20006 bytes, offset:2936233 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDL-05.fs2
Extracted file data/missions/MDL-06.fs2 size:17868 bytes, offset:2956239 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDL-06.fs2
Extracted file data/missions/MDL-07.fs2 size:13789 bytes, offset:2974107 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDL-07.fs2
Extracted file data/missions/mdL-08.fs2 size:15205 bytes, offset:2987896 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/mdL-08.fs2
Extracted file data/missions/mdL-09.fs2 size:17661 bytes, offset:3003101 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/mdL-09.fs2
Extracted file data/missions/MDM-01.fs2 size:18169 bytes, offset:3020762 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDM-01.fs2
Extracted file data/missions/MDM-02.fs2 size:17372 bytes, offset:3038931 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDM-02.fs2
Extracted file data/missions/MDM-03.fs2 size:18248 bytes, offset:3056303 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDM-03.fs2
Extracted file data/missions/MDM-04.fs2 size:18119 bytes, offset:3074551 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDM-04.fs2
Extracted file data/missions/MDM-05.fs2 size:20393 bytes, offset:3092670 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDM-05.fs2
Extracted file data/missions/MDM-06.fs2 size:18243 bytes, offset:3113063 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDM-06.fs2
Extracted file data/missions/MDM-07.fs2 size:14165 bytes, offset:3131306 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MDM-07.fs2
Extracted file data/missions/mdM-08.fs2 size:15599 bytes, offset:3145471 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/mdM-08.fs2
Extracted file data/missions/mdM-09.fs2 size:18051 bytes, offset:3161070 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/mdM-09.fs2
Extracted file data/missions/MT-01.fs2 size:27140 bytes, offset:3179121 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MT-01.fs2
Extracted file data/missions/MT-02.fs2 size:26610 bytes, offset:3206261 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MT-02.fs2
Extracted file data/missions/MT-03.fs2 size:28498 bytes, offset:3232871 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MT-03.fs2
Extracted file data/missions/MT-04.fs2 size:17882 bytes, offset:3261369 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MT-04.fs2
Extracted file data/missions/MT-05.fs2 size:37392 bytes, offset:3279251 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MT-05.fs2
Extracted file data/missions/MT-06.fs2 size:15934 bytes, offset:3316643 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MT-06.fs2
Extracted file data/missions/MT-07.fs2 size:21556 bytes, offset:3332577 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MT-07.fs2
Extracted file data/missions/MT-08.fs2 size:17620 bytes, offset:3354133 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MT-08.fs2
Extracted file data/missions/MT-09.fs2 size:14388 bytes, offset:3371753 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MT-09.fs2
Extracted file data/missions/MT-10.fs2 size:19286 bytes, offset:3386141 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/MT-10.fs2
Extracted file data/missions/OSDog.fs2 size:19388 bytes, offset:3405427 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/OSDog.fs2
Extracted file data/missions/SM1-01.fs2 size:68242 bytes, offset:3424815 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM1-01.fs2
Extracted file data/missions/SM1-02.fs2 size:72067 bytes, offset:3493057 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM1-02.fs2
Extracted file data/missions/SM1-03.fs2 size:57855 bytes, offset:3565124 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM1-03.fs2
Extracted file data/missions/SM1-04.fs2 size:75736 bytes, offset:3622979 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM1-04.fs2
Extracted file data/missions/SM1-05.fs2 size:82941 bytes, offset:3698715 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM1-05.fs2
Extracted file data/missions/SM1-06.fs2 size:67908 bytes, offset:3781656 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM1-06.fs2
Extracted file data/missions/SM1-07.fs2 size:51290 bytes, offset:3849564 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM1-07.fs2
Extracted file data/missions/SM1-08.fs2 size:67506 bytes, offset:3900854 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM1-08.fs2
Extracted file data/missions/SM1-09.fs2 size:86239 bytes, offset:3968360 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM1-09.fs2
Extracted file data/missions/SM1-10.fs2 size:67658 bytes, offset:4054599 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM1-10.fs2
Extracted file data/missions/SM2-01.fs2 size:62708 bytes, offset:4122257 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM2-01.fs2
Extracted file data/missions/SM2-02.fs2 size:56674 bytes, offset:4184965 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM2-02.fs2
Extracted file data/missions/SM2-03.fs2 size:61816 bytes, offset:4241639 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM2-03.fs2
Extracted file data/missions/SM2-04.fs2 size:72553 bytes, offset:4303455 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM2-04.fs2
Extracted file data/missions/SM2-05.fs2 size:71511 bytes, offset:4376008 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM2-05.fs2
Extracted file data/missions/SM2-06.fs2 size:71655 bytes, offset:4447519 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM2-06.fs2
Extracted file data/missions/SM2-07.fs2 size:70704 bytes, offset:4519174 (1999-11-04 20:52:28 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM2-07.fs2
Extracted file data/missions/SM2-08.fs2 size:53983 bytes, offset:4589878 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM2-08.fs2
Extracted file data/missions/SM2-09.fs2 size:53565 bytes, offset:4643861 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM2-09.fs2
Extracted file data/missions/sm2-10.fs2 size:85760 bytes, offset:4697426 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/sm2-10.fs2
Extracted file data/missions/SM3-01.fs2 size:77854 bytes, offset:4783186 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM3-01.fs2
Extracted file data/missions/sm3-02.fs2 size:42113 bytes, offset:4861040 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/sm3-02.fs2
Extracted file data/missions/SM3-03.fs2 size:79484 bytes, offset:4903153 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM3-03.fs2
Extracted file data/missions/SM3-04.fs2 size:64871 bytes, offset:4982637 (1999-11-02 03:57:30 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM3-04.fs2
Extracted file data/missions/SM3-05.fs2 size:87235 bytes, offset:5047508 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM3-05.fs2
Extracted file data/missions/SM3-06.fs2 size:79601 bytes, offset:5134743 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM3-06.fs2
Extracted file data/missions/SM3-07.fs2 size:97275 bytes, offset:5214344 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM3-07.fs2
Extracted file data/missions/SM3-08.fs2 size:84092 bytes, offset:5311619 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/SM3-08.fs2
Extracted file data/missions/sm3-09.fs2 size:78558 bytes, offset:5395711 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/sm3-09.fs2
Extracted file data/missions/sm3-10.fs2 size:101883 bytes, offset:5474269 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/sm3-10.fs2
Extracted file data/missions/Templar-01.fs2 size:49062 bytes, offset:5576152 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/Templar-01.fs2
Extracted file data/missions/Templar-02.fs2 size:48699 bytes, offset:5625214 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/Templar-02.fs2
Extracted file data/missions/Templar-03.fs2 size:65106 bytes, offset:5673913 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/Templar-03.fs2
Extracted file data/missions/Templar-04.fs2 size:56363 bytes, offset:5739019 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/Templar-04.fs2
Extracted file data/missions/Training-1.fs2 size:35107 bytes, offset:5795382 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/Training-1.fs2
Extracted file data/missions/Training-2.fs2 size:37616 bytes, offset:5830489 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/Training-2.fs2
Extracted file data/missions/Training-3.fs2 size:38863 bytes, offset:5868105 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/Training-3.fs2
Extracted file data/missions/TSM-104.fs2 size:51541 bytes, offset:5906968 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/TSM-104.fs2
Extracted file data/missions/TSM-105.fs2 size:28534 bytes, offset:5958509 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/TSM-105.fs2
Extracted file data/missions/TSM-106.fs2 size:46482 bytes, offset:5987043 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/TSM-106.fs2
Extracted file data/missions/FreeSpace2.fc2 size:12617 bytes, offset:6033525 (1999-11-02 03:57:28 +0000 UTC) to extracted/Root_fs2.vp/data/missions/FreeSpace2.fc2
Extracted file data/missions/Templar.fc2 size:2374 bytes, offset:6046142 (1999-11-02 03:57:32 +0000 UTC) to extracted/Root_fs2.vp/data/missions/Templar.fc2
Extracted file data/players/images/Ter0001.pcx size:11138 bytes, offset:6048516 (1999-09-22 15:32:26 +0000 UTC) to extracted/Root_fs2.vp/data/players/images/Ter0001.pcx
Extracted file data/players/images/Ter0002.pcx size:9224 bytes, offset:6059654 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/images/Ter0002.pcx
Extracted file data/players/images/Ter0003.pcx size:10501 bytes, offset:6068878 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/images/Ter0003.pcx
Extracted file data/players/images/Ter0004.pcx size:11287 bytes, offset:6079379 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/images/Ter0004.pcx
Extracted file data/players/images/Ter0005.pcx size:11120 bytes, offset:6090666 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/images/Ter0005.pcx
Extracted file data/players/images/Ter0006.pcx size:12220 bytes, offset:6101786 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/images/Ter0006.pcx
Extracted file data/players/images/Ter0007.pcx size:10999 bytes, offset:6114006 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/images/Ter0007.pcx
Extracted file data/players/images/Ter0008.pcx size:11592 bytes, offset:6125005 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/images/Ter0008.pcx
Extracted file data/players/images/Ter0009.pcx size:10046 bytes, offset:6136597 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/images/Ter0009.pcx
Extracted file data/players/images/Vas0001.pcx size:12438 bytes, offset:6146643 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/images/Vas0001.pcx
Extracted file data/players/images/Vas0002.pcx size:12124 bytes, offset:6159081 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/images/Vas0002.pcx
Extracted file data/players/images/Vas0003.pcx size:12147 bytes, offset:6171205 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/images/Vas0003.pcx
Extracted file data/players/images/vssver.scc size:224 bytes, offset:6183352 (1999-11-17 20:55:44 +0000 UTC) to extracted/Root_fs2.vp/data/players/images/vssver.scc
Extracted file data/players/squads/242suicide.pcx size:8772 bytes, offset:6183576 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/242suicide.pcx
Extracted file data/players/squads/77th.pcx size:6931 bytes, offset:6192348 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/77th.pcx
Extracted file data/players/squads/barracudas.pcx size:12018 bytes, offset:6199279 (1999-09-22 15:32:28 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/barracudas.pcx
Extracted file data/players/squads/DaBombInsig.pcx size:16783 bytes, offset:6211297 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/DaBombInsig.pcx
Extracted file data/players/squads/deathsquad.pcx size:12973 bytes, offset:6228080 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/deathsquad.pcx
Extracted file data/players/squads/desquad.pcx size:10287 bytes, offset:6241053 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/desquad.pcx
Extracted file data/players/squads/hammer1.pcx size:6269 bytes, offset:6251340 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/hammer1.pcx
Extracted file data/players/squads/jollyroger.pcx size:8266 bytes, offset:6257609 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/jollyroger.pcx
Extracted file data/players/squads/lion1.pcx size:10232 bytes, offset:6265875 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/lion1.pcx
Extracted file data/players/squads/raptors.pcx size:11142 bytes, offset:6276107 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/raptors.pcx
Extracted file data/players/squads/ravens.pcx size:10219 bytes, offset:6287249 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/ravens.pcx
Extracted file data/players/squads/rgrdgr1.pcx size:8970 bytes, offset:6297468 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/rgrdgr1.pcx
Extracted file data/players/squads/rgrdgr2.pcx size:7850 bytes, offset:6306438 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/rgrdgr2.pcx
Extracted file data/players/squads/scorpsquada.pcx size:7810 bytes, offset:6314288 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/scorpsquada.pcx
Extracted file data/players/squads/scorpsquadb.pcx size:10641 bytes, offset:6322098 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/scorpsquadb.pcx
Extracted file data/players/squads/scorpsquadc.pcx size:9471 bytes, offset:6332739 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/scorpsquadc.pcx
Extracted file data/players/squads/scorpsquadd.pcx size:7117 bytes, offset:6342210 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/scorpsquadd.pcx
Extracted file data/players/squads/sheepsquad.pcx size:9669 bytes, offset:6349327 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/sheepsquad.pcx
Extracted file data/players/squads/skullsquada.pcx size:7678 bytes, offset:6358996 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/skullsquada.pcx
Extracted file data/players/squads/tsusquad.pcx size:11214 bytes, offset:6366674 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/tsusquad.pcx
Extracted file data/players/squads/vasudan1.pcx size:7212 bytes, offset:6377888 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/vasudan1.pcx
Extracted file data/players/squads/volsquad.pcx size:7347 bytes, offset:6385100 (1999-09-22 15:32:30 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/volsquad.pcx
Extracted file data/players/squads/vssver.scc size:384 bytes, offset:6392447 (1999-11-17 20:55:44 +0000 UTC) to extracted/Root_fs2.vp/data/players/squads/vssver.scc
Extracted file data/pxohelp.txt size:4095 bytes, offset:6392831 (1999-09-22 15:17:30 +0000 UTC) to extracted/Root_fs2.vp/data/pxohelp.txt
```

## Get checksums inside `.vp` file:

```
% ./fs2util -c Root_fs2.vp                                   
415739ffc79b5fbff7a234460b814174dab4ec9874ce52da3df1fee05fa1e3ff  Root_fs2.vp/data/tables/ai.tbl
755b3804ae65b92b3a1485574ecb2593f029627d3c56f99a6402e73827636e81  Root_fs2.vp/data/tables/asteroid.tbl
d67a949c70519f295d9c4c9079b035f0e4869cd15b6f7834956668e857fc03b9  Root_fs2.vp/data/tables/credits.tbl
2edb757552ba1036de0d4467908c6aa5ef381e77702c196b4def5f6eaf0061fa  Root_fs2.vp/data/tables/cutscenes.tbl
83fd81da4d25bdcc56f14024bf867650056d75c04cce29d989b87135995d9917  Root_fs2.vp/data/tables/fireball.tbl
6521042ec179897f5ae5ccd83677818a90c6177f7f8794d60987c225734757d3  Root_fs2.vp/data/tables/help.tbl
7703084109104eae56cba606b2f396ed043ca4fbc9949722ad122e29d5748aec  Root_fs2.vp/data/tables/hud.tbl
30bf5e6745088ef1afccd32db84da5e5c7eb82e42a23a9991c96fb703d3b346b  Root_fs2.vp/data/tables/icons.tbl
d6f08ad15462809da514d178d12b01961c2d1a7ba89418884e70eb251e5b43e0  Root_fs2.vp/data/tables/launchhelp.tbl
27ad6b6a40da74086f5856c9b7f73e001ff8665138c02aecd204a2dfd02e3e9e  Root_fs2.vp/data/tables/lightning.tbl
0be24bbdbd2c2f0699d2e156f07271bd882a520850f03236fb59a378269b8b91  Root_fs2.vp/data/tables/mainhall.tbl
0e965d476927452251a9da0aabf97bbaa47c93a9be1fc10a560503e1154595aa  Root_fs2.vp/data/tables/medals.tbl
b5ea838d1322b33c1fca5d82067d7df2fa711dc05dffb2120fb667a720173c75  Root_fs2.vp/data/tables/menu.tbl
391a5fa6cf0d564b230f3d9a0c71d57c5dd22aa5dc604aae400f5f94c7bd9d0a  Root_fs2.vp/data/tables/messages.tbl
918d32fb372300f3d6db88715149cb0fe3608ac6923adb29a2fd552ac35dae72  Root_fs2.vp/data/tables/mflash.tbl
a6002ab5e0c4d55b2b08ad1852bc0b8fe038ecda9fb68745a4a60b1993f4839c  Root_fs2.vp/data/tables/music.tbl
77ab1a6851bfa789005b3cf0c2a1bd2945380b2709a0806204080e4c1922e054  Root_fs2.vp/data/tables/nebula.tbl
b0905a6e4bd71190be526c9c35808035b0ba7c65786ba14a80ec91fb907848c0  Root_fs2.vp/data/tables/pixels.tbl
286808e8cab51c723ea93cb799c56ee804e2c5e1cfbf40a3eb0374506b374ab5  Root_fs2.vp/data/tables/rank.tbl
60266906ff4bc25932ecda1aa8070c92de0e7ab6cde42c2d5a02169c7e6315ca  Root_fs2.vp/data/tables/ships.tbl
c4e18155b499a9962768ea175e30c037baf199fddfeb7900d6283d71c004b23b  Root_fs2.vp/data/tables/sounds.tbl
44034c1db75250ff76324e24cea8cd703bd3c0afd16bbdf64de8bdde60e0b547  Root_fs2.vp/data/tables/Species.tbl
6f02b9cd810ae22f69cdb768bc5a59f3280bc44bce25a62d70ac38b6e2b07c0d  Root_fs2.vp/data/tables/ssm.tbl
a49766f566ea0f477779df7bd4907f5ea03a6ce4158fde55ddcc27363cd50c26  Root_fs2.vp/data/tables/stars.tbl
4440c2b3b49669ae5f9c243fb6c5d1b6f31f4e33ae0538a89e5f2c72240c2b3e  Root_fs2.vp/data/tables/strings.tbl
76df663ea62a777794d322c28684d9adae0cf2b8d36571d417679337aa47d6a3  Root_fs2.vp/data/tables/tips.tbl
73f4db09981fa8d15db7a3da6dbdb3d91757248c8f7dfe0c0f8eb58f7d877f2b  Root_fs2.vp/data/tables/traitor.tbl
d6c4c0e2fa969b67ccc2c3d99852dcef64a3417dccd44ed5223ffc741d57add1  Root_fs2.vp/data/tables/tstrings.tbl
4e33030540bdac725c14672b1fce7bf087629ddbab87dc0e121726cfb87253db  Root_fs2.vp/data/tables/weapon_expl.tbl
9e94ba10b7cd5acafefb22dcaada2545ca40d713fae112597a4dca1f63db6f9a  Root_fs2.vp/data/tables/weapons.tbl
e04f2e42a2f0a34e5ea88107e2fabb3b46c7042f8be038488daf18ab5ccb3c0f  Root_fs2.vp/data/tables/vssver.scc
64e38222160ac6e72c458a40fe74d3908c75685b5d7ce6a8a327ad27d00f2583  Root_fs2.vp/data/missions/G-Shi.fs2
7f34bfdf98d5aa239c3b530f627d99ca3e6d07eaea583583f9038cd033e14866  Root_fs2.vp/data/missions/G-Ter.fs2
cf4a73e61f5474e3eb42c53cf8b93f8ca314140d74129d3ed6876be68096a95e  Root_fs2.vp/data/missions/G-Vas.fs2
6575f3d496e14f80269e5e8174bec5cdf5adbbcaf9d2ebbe8b3c4b8a8dfc0e67  Root_fs2.vp/data/missions/loop1-1.fs2
1313cbab4c0c397b8147d111241cae34a37f995f6e30a491e937fa3d3cbe77c9  Root_fs2.vp/data/missions/loop1-2.fs2
8f7c7a8232efb227f123d32175cb16cf86dacde633d57980f6ce1f11cf86cd8b  Root_fs2.vp/data/missions/loop1-3.fs2
a4c4970c2a7dc9d3568748d65a05626a19c87c832ad358d68ac91fb470bdf318  Root_fs2.vp/data/missions/loop2-1.fs2
c4ba664e756e8be446133daa7d38778f468415cfb95bea07023328855da94eb7  Root_fs2.vp/data/missions/loop2-2.fs2
4f00ddc588ab7329e2776169f3236a2a70a9ea3008f41727f6feeb2675d18904  Root_fs2.vp/data/missions/M-01.fs2
8167043384691767be1de11321db94c1781b2d6031f656185b90258db66db2c5  Root_fs2.vp/data/missions/M-02.fs2
3a9984d3134ebc8d4b0b1e8d91934024d04f28988d2b20f41853845c1bcb9bed  Root_fs2.vp/data/missions/M-03.fs2
e88ddde9abee675ec9bf7f6f58c79e9ed67b827c7e723e8389c837d9734cb702  Root_fs2.vp/data/missions/M-04.fs2
0c87f893aa6518daf3e3bd8e645c4a188ab4b5eaa15eb4bca1aa7c3f318236bb  Root_fs2.vp/data/missions/MDH-01.fs2
5a4d309a57b4d68bc571b4a0cc6fb3b14ce4d9c42927bba2f86cf2b55801f852  Root_fs2.vp/data/missions/MDH-02.fs2
42118d248b9f2325eb503683b68da819385b5c8204a5002aa1843a503963802c  Root_fs2.vp/data/missions/MDH-03.fs2
144fdae46411cddfcf5938de9494c98182443f77f2fa3e65d76f2999ec8a3fd6  Root_fs2.vp/data/missions/MDH-04.fs2
e725304d7c3e926d15178f6eae0f01f682ad1a618ef06998c0dd583cf794f6c4  Root_fs2.vp/data/missions/MDH-05.fs2
75d1645107be3caf97505d8df3f57fb8faf39683032f1eab321d4ef1d40fd430  Root_fs2.vp/data/missions/MDH-06.fs2
2f9f231885a13aae3651b69e9eb0c5c0350cd38ecece4d27bf370ebac1c5558a  Root_fs2.vp/data/missions/MDH-07.fs2
ae0fb14b473c63d8795f4a065904832e9511a376918681b697c68c4e9cb24d0f  Root_fs2.vp/data/missions/mdH-08.fs2
0bc926bbc0730c9e43c4536239137773d93ced24650a552bca3cb386e4e32149  Root_fs2.vp/data/missions/mdH-09.fs2
f5c8a1833044dc79c7b79f30d1c711dc6b2fe3ea0f3b91e39c483d52046fe1da  Root_fs2.vp/data/missions/MDL-01.fs2
c98eed8f82a74d36042f15aabc9701e17b2f455611636dbfddcbfbc600c81fc6  Root_fs2.vp/data/missions/MDL-02.fs2
55322681469804ffc0630554bed7ef581a7caea3592ae77dce841b269314c2fb  Root_fs2.vp/data/missions/MDL-03.fs2
98a3e6f5768b44774e93040f6dcce4cc1f5548eb118ecc4e08516a41df6fd5e5  Root_fs2.vp/data/missions/MDL-04.fs2
cf379b45180efa27f8a8de4acf105ad8190092e9ccf6b6283f687da6f6729179  Root_fs2.vp/data/missions/MDL-05.fs2
e9e8a91240b863c1c0e5d286f50e3735f4b1061bf60f295703c6469d91751b6a  Root_fs2.vp/data/missions/MDL-06.fs2
fd8be40e57b21b48d359b0a6446678cb55d73b2d869ef797cd0edf0355645442  Root_fs2.vp/data/missions/MDL-07.fs2
32990977fb5323e6cc1323eadcbbc6935c9b4e69ceaecb7d74fa9446a6d855ed  Root_fs2.vp/data/missions/mdL-08.fs2
74c4ab33f35bb11e66cd5e9f2efd1c2651dbe6f3b3e92f1f7bace48175957c5e  Root_fs2.vp/data/missions/mdL-09.fs2
4494e372e17b56da8afcef0878c9d96ce1d7687ec9c26679e148bc139f23c18f  Root_fs2.vp/data/missions/MDM-01.fs2
850d28b4adeaa6d38fbca836fc9602a2dc91f3031634bfb38b25641eedd80f82  Root_fs2.vp/data/missions/MDM-02.fs2
6345bcda0419f1125a55dc60044affecdb951d8c1d557b5e7115fdc357ea2e89  Root_fs2.vp/data/missions/MDM-03.fs2
c3607ce330fd9ec3a13c0f2a75c9daf6d4a4a505de6ff944f2346603b1c3b866  Root_fs2.vp/data/missions/MDM-04.fs2
c5acb2331829e7af5253f46235a09897d1412a947ffbd4fe2b9a405b3fe7bbd7  Root_fs2.vp/data/missions/MDM-05.fs2
6ef1ae3cbd3839bf34bc599d9b7570943b479c62782ab31841636488e8b9c0fe  Root_fs2.vp/data/missions/MDM-06.fs2
f11c1066c6adcd1dadaf4204b70b63198009c0e5e0f57698d634453b15e0da88  Root_fs2.vp/data/missions/MDM-07.fs2
57f13518e95ccd0f0dbfb920e5de74c1d8e30dd39f838ee9ba230f5535613e71  Root_fs2.vp/data/missions/mdM-08.fs2
7d40bce625d26b52547b535e6e80d1eca58a8e526cd62b9ed86d8323d4621b40  Root_fs2.vp/data/missions/mdM-09.fs2
a2b039ed1c1d486813ca4170b8266d234c3fda5dfc41ee677cd314463816d052  Root_fs2.vp/data/missions/MT-01.fs2
992a512fe531895cbf6e3fd9d325b7ec43aeaeae28029578ba866a9b6bad2479  Root_fs2.vp/data/missions/MT-02.fs2
3ce8a411884e19b44646a847ae3578c3547cbef4ffcdb503bf819346475dd058  Root_fs2.vp/data/missions/MT-03.fs2
6bbdda7aafefd512320c86fa641f478754975586b3fdeb28a49a21d2139b8b1f  Root_fs2.vp/data/missions/MT-04.fs2
1ce027989a644f7b0ee6088785eaad9a52a2f25f4c1a65a84e29940e444f45b4  Root_fs2.vp/data/missions/MT-05.fs2
a63d95c8a7586547b0aab7e32b0f1d5c3fe0e67c6ad313bf6599df3526a13c67  Root_fs2.vp/data/missions/MT-06.fs2
aba3db84eed6c82990c407d5cd390bc8b512b685554590a22bf64a53846b0327  Root_fs2.vp/data/missions/MT-07.fs2
bdeeb4b7909c44b8c90ed2073139f922de51af68ba744565b89b495ceae7f7f6  Root_fs2.vp/data/missions/MT-08.fs2
408049bf2fef4c0de0afe919ba35fbbc2317d36c4e6e55bb1c3783558c4ee795  Root_fs2.vp/data/missions/MT-09.fs2
b78a846adbd63e40c6b59ccb8b4a5b39a240297e8b547e434e6bfc8a13d9ca87  Root_fs2.vp/data/missions/MT-10.fs2
b55c9d3c30b538b8edc10fd3cc02169a666f7e03f6a24c6f2445aa8b5ade9216  Root_fs2.vp/data/missions/OSDog.fs2
17d7fc1d707bcd102631b70fb44c688f78a0790bda8097642118a089887c09c0  Root_fs2.vp/data/missions/SM1-01.fs2
91172b4e70b31e7aca5a0a44ce0e7a788139699ddfa289b8337d532e775e2735  Root_fs2.vp/data/missions/SM1-02.fs2
a92b893ae1c98e16e1d8525b5d90cd4679db5ed6297ecce257ad8a81f30420f7  Root_fs2.vp/data/missions/SM1-03.fs2
5efbd217616fa1cfc22ef7bb0a779a371ca75de0ce8f0b8c8375f2e2a0062c6e  Root_fs2.vp/data/missions/SM1-04.fs2
85c153568cff98302e1e5142c2f0097b497a3282a9974713beff0d9c0443bea6  Root_fs2.vp/data/missions/SM1-05.fs2
ed1a26dde89496dfdc4ecfad4915ff306d86245a2f1500d8d9fbe00d6debdef3  Root_fs2.vp/data/missions/SM1-06.fs2
9d382a9dbeb6975742498a867482e189cf72edd51a7dd955e62bd68a768e8374  Root_fs2.vp/data/missions/SM1-07.fs2
ffa6b10d7d353f3beb2f03cc076c417b4ac7e2be510e0f5f03a222c8b1dcafff  Root_fs2.vp/data/missions/SM1-08.fs2
6a6b452ed24cf4fa89f68d212790b7cdd4941683dac0df41028f38d311c80d4d  Root_fs2.vp/data/missions/SM1-09.fs2
7c43ec9fb6b61a6ff8eab1337413c35cefe1b92257e3c69f2eb4b2ca3f4eb01d  Root_fs2.vp/data/missions/SM1-10.fs2
5180d69f055cf9d1ecac444d6c94afdecd4e0d45f49ff9a38c284c1af5283846  Root_fs2.vp/data/missions/SM2-01.fs2
32c5a288f704f39b26dbd2116f04a00ace8275a52723846bd7a6a9081871998a  Root_fs2.vp/data/missions/SM2-02.fs2
d5d383dbd530413cf5cceb46a3cc4b7e23ac24e27f254f2db4e2da070b5165eb  Root_fs2.vp/data/missions/SM2-03.fs2
9085a61b78271ae130dd716eb57d33a8289a298db4968c2447ccad3ad29fdabf  Root_fs2.vp/data/missions/SM2-04.fs2
f9d592e013e60adc44b33876faabd27d81e3afa249f51694267d917afe12a531  Root_fs2.vp/data/missions/SM2-05.fs2
899a75296d2b0a2b4989178d8a2dc6448cf3330efdbf62f28995eaeab87d87b7  Root_fs2.vp/data/missions/SM2-06.fs2
9ba00c7868c054a5560f9cec46f8771f3203dbb38df868f26d9b7338ebbbed45  Root_fs2.vp/data/missions/SM2-07.fs2
b39dec05b9a8ddf0d993c7b0b8e73add18ac0e871bc46279948de30b0cf334a2  Root_fs2.vp/data/missions/SM2-08.fs2
076a8b556a839c0ced6321ddfcaf23bc568bcca7c2aa953be9b062f4734333c8  Root_fs2.vp/data/missions/SM2-09.fs2
562204701a6e7767b1ed30aa3b78390777551b3de75ea37092c41c972cf7774d  Root_fs2.vp/data/missions/sm2-10.fs2
e4eba99e54460c1049416725fbe4e9b9107a3550e5df8f8d29f2967238abd2b5  Root_fs2.vp/data/missions/SM3-01.fs2
6c6b84b09df957c29b9a1195ca580bd8255411bea428d1a737bfc8075d552042  Root_fs2.vp/data/missions/sm3-02.fs2
f498313c1d7f4ab143676fb71e52494a6d4f5a8f159c77231f33cf4f7d20894a  Root_fs2.vp/data/missions/SM3-03.fs2
ccab18fbe1247791d7c0fef63743290b609657e24a218b75f2224837f857b0ea  Root_fs2.vp/data/missions/SM3-04.fs2
5ff349718da6c733e187e3d4dbf972f6a86f77285af9b14a5bacd4cfc186ee6b  Root_fs2.vp/data/missions/SM3-05.fs2
e86fe0816197c5ac8417c1d2799e2a297c6f0d12448c1209d518f8c7789dfcb5  Root_fs2.vp/data/missions/SM3-06.fs2
5b3ba8bdc8a91a6b3f038a474903995fba82115b7666c3a4411354265bbe3a32  Root_fs2.vp/data/missions/SM3-07.fs2
252761e12777afb62fadf4ba5e1dcbd446d2092a3981cceb94f4639267702509  Root_fs2.vp/data/missions/SM3-08.fs2
84d2566f35c3deca93fac9cb1cc25db566d76bae6fa6c49435eab49a3f3ab67f  Root_fs2.vp/data/missions/sm3-09.fs2
7499be3b5f34eee317c8c971f2c6f97a8ee0faecc15942fd0ca769b278640b99  Root_fs2.vp/data/missions/sm3-10.fs2
1c6e230dcbc70a7340c2a5fb2fd30927629b5684948b19bb6bb17814e67218b1  Root_fs2.vp/data/missions/Templar-01.fs2
f79cd17214a1c74f1a6f5ba9339ec37089fee57d3d05c802e2d992ba9bce8d29  Root_fs2.vp/data/missions/Templar-02.fs2
8b8d46de14a0e62693295ed7ef8560dbbcf4c3c7e1de5ce291c4284b49a8a665  Root_fs2.vp/data/missions/Templar-03.fs2
85062529ae16f45acccbbb6617694fde0adbf09dd1567b2934814f08643401db  Root_fs2.vp/data/missions/Templar-04.fs2
76653fcd4f60a3889103b117e3b172e8dae8492af500d0ad201d88187f45cba4  Root_fs2.vp/data/missions/Training-1.fs2
ca00b7627161a03ad149cf370d13bcda7bc9586c3dbccb4291b90f5eefbd16d0  Root_fs2.vp/data/missions/Training-2.fs2
835b0deb55ad2c419e1d2891693989c2dfb8c23973e73edd2df7f42c656260a2  Root_fs2.vp/data/missions/Training-3.fs2
a162983593455d5898d2c080770f85a4066e8866a1cf123af26c9471af831063  Root_fs2.vp/data/missions/TSM-104.fs2
834a8b9150a6bb9370039d416855d7b96f9ee7fbb721bc1d0cd15647f2d00314  Root_fs2.vp/data/missions/TSM-105.fs2
74b761c253ad74a0ed926f0429ad30670ed76ada17862fe607a9c5ceaf0e09af  Root_fs2.vp/data/missions/TSM-106.fs2
7baf341b0b95df7ca48b5d6101ed5c68ad040c3cb6362db6411a10267828315d  Root_fs2.vp/data/missions/FreeSpace2.fc2
8cd786aadd4d2495199dcfdce2302e300bf5b6a4339ef06fbeec4b2b44d0b478  Root_fs2.vp/data/missions/Templar.fc2
b9fec91aa1adc795868d7f9d34ef17711396fc7c1b6fc7c133afbfc290211f2b  Root_fs2.vp/data/players/images/Ter0001.pcx
c5d43c50744aa908564511d041345ec800d6dcb34f1d4a1045d1dbb82a9229e3  Root_fs2.vp/data/players/images/Ter0002.pcx
10bae6acfd6a58618ec6538cba21322f07cbf9a3d01cda268720d071f153366c  Root_fs2.vp/data/players/images/Ter0003.pcx
206284993eb7cf32bfe3fdc4236e985806fe62543a72341c54520af877d3520d  Root_fs2.vp/data/players/images/Ter0004.pcx
2bab513ae1de4296b45ffa81c74067d372a373ac8dadbcc5e4d4d6c08bb286e8  Root_fs2.vp/data/players/images/Ter0005.pcx
1de629489633006e2bf64517ba457dc0f6259e0a95bd158c42f6e2ebbeb89c13  Root_fs2.vp/data/players/images/Ter0006.pcx
4dc2c23f30a1bfc768ac1a7f1a9b4e881ec11c3211237466012177dc3227a0d6  Root_fs2.vp/data/players/images/Ter0007.pcx
65568fe76c0f1e1d658f6c8a958223cc51e9243c33fa08cfa266a18b3c0ce047  Root_fs2.vp/data/players/images/Ter0008.pcx
54e4ac7af25641d03b52a314436c1b44b9ae90e1acfdb4de57e58fa1248920fb  Root_fs2.vp/data/players/images/Ter0009.pcx
36f6d28b5e2410df4e5f8672feeb9418cd5bce8a6380d53be41475955d9d49e0  Root_fs2.vp/data/players/images/Vas0001.pcx
9835b8580a70e545db613f87551fda50f7fe79393409200956b1e391e2d7ec85  Root_fs2.vp/data/players/images/Vas0002.pcx
4ef940ed135731f11a192dd16e125dac539088aaf8a76df5e8e8f0cb2c46bf63  Root_fs2.vp/data/players/images/Vas0003.pcx
7a0dec4281b82fe65bb290f78c01f62288b49d77295082eeaa0f9544479fe214  Root_fs2.vp/data/players/images/vssver.scc
57864c6275100f6fc530715878bcc1e5910ff2fd366deec965bed02dc47eafc5  Root_fs2.vp/data/players/squads/242suicide.pcx
545b61196ab2f272d94c1a44ef3036f4529f769716edb0a78d60e0710d487b9c  Root_fs2.vp/data/players/squads/77th.pcx
8d52f8da684f3f4c47c6d1dbe9055a39a51148e5d92a54801823aaf4074c9bdd  Root_fs2.vp/data/players/squads/barracudas.pcx
83f0688834a88dae40a3f8152b17df51be8605696a9e1db5b152564041a5084c  Root_fs2.vp/data/players/squads/DaBombInsig.pcx
0bc5490d83f08d4f8f48a0b994041d556057ed2dfbe9b1d173caef8f4f08a00a  Root_fs2.vp/data/players/squads/deathsquad.pcx
2480eafd29dbf56296334dc11eabd7dcc4fbc078f7453508cbbdb410161f8096  Root_fs2.vp/data/players/squads/desquad.pcx
9ae2ac403457f7cd32202e6bff396292b043c3d2f3843ce9a910315f8d9b1004  Root_fs2.vp/data/players/squads/hammer1.pcx
cf7057bdf7c18c26d75221f792cc355c2a503b5c197f0f5654d0186256b09e49  Root_fs2.vp/data/players/squads/jollyroger.pcx
084d8e310f9e32ced9b32f4f970516ec1c6786823001a0e728ddc960967c55e9  Root_fs2.vp/data/players/squads/lion1.pcx
8b734526deb2d4cab3dda9586af1a3b74cc19c496de824f34b8b68f19f04b99f  Root_fs2.vp/data/players/squads/raptors.pcx
42cb4db24d7a73745f8a2a8bbf79605450d678c4fbf72ddef6d4d52c7b1240d0  Root_fs2.vp/data/players/squads/ravens.pcx
f98fa2c63d3019dd9fe15ac6a0d0562ee1582cdd3cf700f3a0d2fa93e30ef612  Root_fs2.vp/data/players/squads/rgrdgr1.pcx
d6e5a9a5ce8953587f08c7c4e30695a0bee56c74b54b4045f87e9e4b267d0762  Root_fs2.vp/data/players/squads/rgrdgr2.pcx
f426daeea03b19c0106b3f9481de367c9d08e010cb70422d5d4d125c7d8f48aa  Root_fs2.vp/data/players/squads/scorpsquada.pcx
74f91945f0a66fc68a6743882c040dabaa486b62d566bb0659945d48085eb362  Root_fs2.vp/data/players/squads/scorpsquadb.pcx
ef29c5c7f37e7993123bbae7f599ee85f7e3d89d02dc61f1d744cac6c80a886c  Root_fs2.vp/data/players/squads/scorpsquadc.pcx
5e5db3dd83d8d54e1e4f735a692152966f41cd45ba42a1d6ac508abb6ce678ea  Root_fs2.vp/data/players/squads/scorpsquadd.pcx
84519be203fb19922b07343410f46e2d1df2c8fa31e54dc07a56f7dafabb79d7  Root_fs2.vp/data/players/squads/sheepsquad.pcx
a0b6ad3600680d227b9dfdfb3db821bf1b2f8720ded93791b08f691d38315bbe  Root_fs2.vp/data/players/squads/skullsquada.pcx
6d7dc9175152cb9cf8e2c491af54229248884a3db21d82ce2bad35f5680db732  Root_fs2.vp/data/players/squads/tsusquad.pcx
cee15c0ae5f9284e18446b6b1caa03335a5f6ba333ccac0d8b0e812ce14452a4  Root_fs2.vp/data/players/squads/vasudan1.pcx
411f29bd389e57c521ec14f229120467bab744984bec306590e8e7924bffea33  Root_fs2.vp/data/players/squads/volsquad.pcx
ab0f42fb8c44c685d506dfbb1efc83ab183c307cd19e4d0127972bf521e98ffd  Root_fs2.vp/data/players/squads/vssver.scc
cf73ca95128cff5569f88faedf731e5c434778b2874144f82aedf1206ebd4ef0  Root_fs2.vp/data/pxohelp.txt
```

## List files inside `.vp` file:

```
% ./fs2util -i Root_fs2.vp
source:Root_fs2.vp file:data/tables/ai.tbl size:1874 bytes, offset:16 date:1999-10-28 07:14:14 +0000 UTC
source:Root_fs2.vp file:data/tables/asteroid.tbl size:4309 bytes, offset:1890 date:1999-10-28 07:14:14 +0000 UTC
source:Root_fs2.vp file:data/tables/credits.tbl size:8542 bytes, offset:6199 date:1999-11-01 20:09:40 +0000 UTC
source:Root_fs2.vp file:data/tables/cutscenes.tbl size:2402 bytes, offset:14741 date:1999-10-28 07:14:14 +0000 UTC
source:Root_fs2.vp file:data/tables/fireball.tbl size:531 bytes, offset:17143 date:1999-10-28 07:14:14 +0000 UTC
source:Root_fs2.vp file:data/tables/help.tbl size:14099 bytes, offset:17674 date:1999-10-28 07:14:14 +0000 UTC
source:Root_fs2.vp file:data/tables/hud.tbl size:800 bytes, offset:31773 date:1999-10-28 07:14:14 +0000 UTC
source:Root_fs2.vp file:data/tables/icons.tbl size:14261 bytes, offset:32573 date:1999-10-28 07:14:14 +0000 UTC
source:Root_fs2.vp file:data/tables/launchhelp.tbl size:4374 bytes, offset:46834 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/lightning.tbl size:5788 bytes, offset:51208 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/mainhall.tbl size:18121 bytes, offset:56996 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/medals.tbl size:1947 bytes, offset:75117 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/menu.tbl size:1962 bytes, offset:77064 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/messages.tbl size:34028 bytes, offset:79026 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/mflash.tbl size:597 bytes, offset:113054 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/music.tbl size:7287 bytes, offset:113651 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/nebula.tbl size:440 bytes, offset:120938 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/pixels.tbl size:81 bytes, offset:121378 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/rank.tbl size:3397 bytes, offset:121459 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/ships.tbl size:383719 bytes, offset:124856 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/sounds.tbl size:20735 bytes, offset:508575 date:1999-11-17 20:13:26 +0000 UTC
source:Root_fs2.vp file:data/tables/Species.tbl size:16708 bytes, offset:529310 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/ssm.tbl size:45 bytes, offset:546018 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/stars.tbl size:1910 bytes, offset:546063 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/strings.tbl size:235241 bytes, offset:547973 date:1999-11-02 16:27:22 +0000 UTC
source:Root_fs2.vp file:data/tables/tips.tbl size:4390 bytes, offset:783214 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/traitor.tbl size:665 bytes, offset:787604 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/tstrings.tbl size:984199 bytes, offset:788269 date:1999-11-11 21:07:26 +0000 UTC
source:Root_fs2.vp file:data/tables/weapon_expl.tbl size:228 bytes, offset:1772468 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/weapons.tbl size:259791 bytes, offset:1772696 date:1999-10-28 07:14:16 +0000 UTC
source:Root_fs2.vp file:data/tables/vssver.scc size:512 bytes, offset:2032487 date:1999-11-02 03:55:16 +0000 UTC
source:Root_fs2.vp file:data/missions/G-Shi.fs2 size:43642 bytes, offset:2032999 date:1999-11-02 03:57:28 +0000 UTC
source:Root_fs2.vp file:data/missions/G-Ter.fs2 size:43671 bytes, offset:2076641 date:1999-11-02 03:57:28 +0000 UTC
source:Root_fs2.vp file:data/missions/G-Vas.fs2 size:42306 bytes, offset:2120312 date:1999-11-02 03:57:28 +0000 UTC
source:Root_fs2.vp file:data/missions/loop1-1.fs2 size:59582 bytes, offset:2162618 date:1999-11-02 03:57:28 +0000 UTC
source:Root_fs2.vp file:data/missions/loop1-2.fs2 size:114569 bytes, offset:2222200 date:1999-11-02 03:57:28 +0000 UTC
source:Root_fs2.vp file:data/missions/loop1-3.fs2 size:58172 bytes, offset:2336769 date:1999-11-02 03:57:28 +0000 UTC
source:Root_fs2.vp file:data/missions/loop2-1.fs2 size:61415 bytes, offset:2394941 date:1999-11-02 03:57:28 +0000 UTC
source:Root_fs2.vp file:data/missions/loop2-2.fs2 size:84727 bytes, offset:2456356 date:1999-11-02 03:57:28 +0000 UTC
source:Root_fs2.vp file:data/missions/M-01.fs2 size:31309 bytes, offset:2541083 date:1999-11-02 03:57:28 +0000 UTC
source:Root_fs2.vp file:data/missions/M-02.fs2 size:41218 bytes, offset:2572392 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/M-03.fs2 size:59154 bytes, offset:2613610 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/M-04.fs2 size:38628 bytes, offset:2672764 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDH-01.fs2 size:17922 bytes, offset:2711392 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDH-02.fs2 size:17138 bytes, offset:2729314 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDH-03.fs2 size:17765 bytes, offset:2746452 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDH-04.fs2 size:17627 bytes, offset:2764217 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDH-05.fs2 size:19905 bytes, offset:2781844 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDH-06.fs2 size:17749 bytes, offset:2801749 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDH-07.fs2 size:13682 bytes, offset:2819498 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/mdH-08.fs2 size:15115 bytes, offset:2833180 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/mdH-09.fs2 size:17542 bytes, offset:2848295 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDL-01.fs2 size:17791 bytes, offset:2865837 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDL-02.fs2 size:16993 bytes, offset:2883628 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDL-03.fs2 size:17867 bytes, offset:2900621 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDL-04.fs2 size:17745 bytes, offset:2918488 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDL-05.fs2 size:20006 bytes, offset:2936233 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDL-06.fs2 size:17868 bytes, offset:2956239 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDL-07.fs2 size:13789 bytes, offset:2974107 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/mdL-08.fs2 size:15205 bytes, offset:2987896 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/mdL-09.fs2 size:17661 bytes, offset:3003101 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDM-01.fs2 size:18169 bytes, offset:3020762 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDM-02.fs2 size:17372 bytes, offset:3038931 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDM-03.fs2 size:18248 bytes, offset:3056303 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDM-04.fs2 size:18119 bytes, offset:3074551 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDM-05.fs2 size:20393 bytes, offset:3092670 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDM-06.fs2 size:18243 bytes, offset:3113063 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MDM-07.fs2 size:14165 bytes, offset:3131306 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/mdM-08.fs2 size:15599 bytes, offset:3145471 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/mdM-09.fs2 size:18051 bytes, offset:3161070 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MT-01.fs2 size:27140 bytes, offset:3179121 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MT-02.fs2 size:26610 bytes, offset:3206261 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MT-03.fs2 size:28498 bytes, offset:3232871 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MT-04.fs2 size:17882 bytes, offset:3261369 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MT-05.fs2 size:37392 bytes, offset:3279251 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MT-06.fs2 size:15934 bytes, offset:3316643 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MT-07.fs2 size:21556 bytes, offset:3332577 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MT-08.fs2 size:17620 bytes, offset:3354133 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MT-09.fs2 size:14388 bytes, offset:3371753 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/MT-10.fs2 size:19286 bytes, offset:3386141 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/OSDog.fs2 size:19388 bytes, offset:3405427 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM1-01.fs2 size:68242 bytes, offset:3424815 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM1-02.fs2 size:72067 bytes, offset:3493057 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM1-03.fs2 size:57855 bytes, offset:3565124 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM1-04.fs2 size:75736 bytes, offset:3622979 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM1-05.fs2 size:82941 bytes, offset:3698715 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM1-06.fs2 size:67908 bytes, offset:3781656 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM1-07.fs2 size:51290 bytes, offset:3849564 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM1-08.fs2 size:67506 bytes, offset:3900854 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM1-09.fs2 size:86239 bytes, offset:3968360 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM1-10.fs2 size:67658 bytes, offset:4054599 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM2-01.fs2 size:62708 bytes, offset:4122257 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM2-02.fs2 size:56674 bytes, offset:4184965 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM2-03.fs2 size:61816 bytes, offset:4241639 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM2-04.fs2 size:72553 bytes, offset:4303455 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM2-05.fs2 size:71511 bytes, offset:4376008 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM2-06.fs2 size:71655 bytes, offset:4447519 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM2-07.fs2 size:70704 bytes, offset:4519174 date:1999-11-04 20:52:28 +0000 UTC
source:Root_fs2.vp file:data/missions/SM2-08.fs2 size:53983 bytes, offset:4589878 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM2-09.fs2 size:53565 bytes, offset:4643861 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/sm2-10.fs2 size:85760 bytes, offset:4697426 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM3-01.fs2 size:77854 bytes, offset:4783186 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/sm3-02.fs2 size:42113 bytes, offset:4861040 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM3-03.fs2 size:79484 bytes, offset:4903153 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM3-04.fs2 size:64871 bytes, offset:4982637 date:1999-11-02 03:57:30 +0000 UTC
source:Root_fs2.vp file:data/missions/SM3-05.fs2 size:87235 bytes, offset:5047508 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/SM3-06.fs2 size:79601 bytes, offset:5134743 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/SM3-07.fs2 size:97275 bytes, offset:5214344 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/SM3-08.fs2 size:84092 bytes, offset:5311619 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/sm3-09.fs2 size:78558 bytes, offset:5395711 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/sm3-10.fs2 size:101883 bytes, offset:5474269 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/Templar-01.fs2 size:49062 bytes, offset:5576152 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/Templar-02.fs2 size:48699 bytes, offset:5625214 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/Templar-03.fs2 size:65106 bytes, offset:5673913 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/Templar-04.fs2 size:56363 bytes, offset:5739019 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/Training-1.fs2 size:35107 bytes, offset:5795382 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/Training-2.fs2 size:37616 bytes, offset:5830489 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/Training-3.fs2 size:38863 bytes, offset:5868105 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/TSM-104.fs2 size:51541 bytes, offset:5906968 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/TSM-105.fs2 size:28534 bytes, offset:5958509 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/TSM-106.fs2 size:46482 bytes, offset:5987043 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/missions/FreeSpace2.fc2 size:12617 bytes, offset:6033525 date:1999-11-02 03:57:28 +0000 UTC
source:Root_fs2.vp file:data/missions/Templar.fc2 size:2374 bytes, offset:6046142 date:1999-11-02 03:57:32 +0000 UTC
source:Root_fs2.vp file:data/players/images/Ter0001.pcx size:11138 bytes, offset:6048516 date:1999-09-22 15:32:26 +0000 UTC
source:Root_fs2.vp file:data/players/images/Ter0002.pcx size:9224 bytes, offset:6059654 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/images/Ter0003.pcx size:10501 bytes, offset:6068878 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/images/Ter0004.pcx size:11287 bytes, offset:6079379 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/images/Ter0005.pcx size:11120 bytes, offset:6090666 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/images/Ter0006.pcx size:12220 bytes, offset:6101786 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/images/Ter0007.pcx size:10999 bytes, offset:6114006 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/images/Ter0008.pcx size:11592 bytes, offset:6125005 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/images/Ter0009.pcx size:10046 bytes, offset:6136597 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/images/Vas0001.pcx size:12438 bytes, offset:6146643 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/images/Vas0002.pcx size:12124 bytes, offset:6159081 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/images/Vas0003.pcx size:12147 bytes, offset:6171205 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/images/vssver.scc size:224 bytes, offset:6183352 date:1999-11-17 20:55:44 +0000 UTC
source:Root_fs2.vp file:data/players/squads/242suicide.pcx size:8772 bytes, offset:6183576 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/squads/77th.pcx size:6931 bytes, offset:6192348 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/squads/barracudas.pcx size:12018 bytes, offset:6199279 date:1999-09-22 15:32:28 +0000 UTC
source:Root_fs2.vp file:data/players/squads/DaBombInsig.pcx size:16783 bytes, offset:6211297 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/deathsquad.pcx size:12973 bytes, offset:6228080 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/desquad.pcx size:10287 bytes, offset:6241053 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/hammer1.pcx size:6269 bytes, offset:6251340 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/jollyroger.pcx size:8266 bytes, offset:6257609 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/lion1.pcx size:10232 bytes, offset:6265875 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/raptors.pcx size:11142 bytes, offset:6276107 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/ravens.pcx size:10219 bytes, offset:6287249 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/rgrdgr1.pcx size:8970 bytes, offset:6297468 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/rgrdgr2.pcx size:7850 bytes, offset:6306438 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/scorpsquada.pcx size:7810 bytes, offset:6314288 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/scorpsquadb.pcx size:10641 bytes, offset:6322098 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/scorpsquadc.pcx size:9471 bytes, offset:6332739 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/scorpsquadd.pcx size:7117 bytes, offset:6342210 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/sheepsquad.pcx size:9669 bytes, offset:6349327 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/skullsquada.pcx size:7678 bytes, offset:6358996 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/tsusquad.pcx size:11214 bytes, offset:6366674 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/vasudan1.pcx size:7212 bytes, offset:6377888 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/volsquad.pcx size:7347 bytes, offset:6385100 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/vssver.scc size:384 bytes, offset:6392447 date:1999-11-17 20:55:44 +0000 UTC
source:Root_fs2.vp file:data/pxohelp.txt size:4095 bytes, offset:6392831 date:1999-09-22 15:17:30 +0000 UTC
```
