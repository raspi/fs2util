# fs2util

![GitHub All Releases](https://img.shields.io/github/downloads/raspi/fs2util/total?style=for-the-badge)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/raspi/fs2util?style=for-the-badge)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/raspi/fs2util?style=for-the-badge)

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

...

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

...

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

...

source:Root_fs2.vp file:data/players/squads/volsquad.pcx size:7347 bytes, offset:6385100 date:1999-09-22 15:32:30 +0000 UTC
source:Root_fs2.vp file:data/players/squads/vssver.scc size:384 bytes, offset:6392447 date:1999-11-17 20:55:44 +0000 UTC
source:Root_fs2.vp file:data/pxohelp.txt size:4095 bytes, offset:6392831 date:1999-09-22 15:17:30 +0000 UTC
```
