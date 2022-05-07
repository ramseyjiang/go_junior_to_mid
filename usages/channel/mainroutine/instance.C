"".main STEXT size=64 args=0x0 locals=0x8 funcid=0x0 align=0x0
	0x0000 00000 (instance.go:8)	TEXT	"".main(SB), ABIInternal, $16-0
	0x0000 00000 (instance.go:8)	MOVD	16(g), R16
	0x0004 00004 (instance.go:8)	PCDATA	$0, $-2
	0x0004 00004 (instance.go:8)	MOVD	RSP, R17
	0x0008 00008 (instance.go:8)	CMP	R16, R17
	0x000c 00012 (instance.go:8)	BLS	48
	0x0010 00016 (instance.go:8)	PCDATA	$0, $-1
	0x0010 00016 (instance.go:8)	MOVD.W	R30, -16(RSP)
	0x0014 00020 (instance.go:8)	MOVD	R29, -8(RSP)
	0x0018 00024 (instance.go:8)	SUB	$8, RSP, R29
	0x001c 00028 (instance.go:8)	FUNCDATA	ZR, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001c 00028 (instance.go:8)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001c 00028 (instance.go:9)	PCDATA	$1, ZR
	0x001c 00028 (instance.go:9)	CALL	"".firstCase(SB)
	0x0020 00032 (instance.go:10)	CALL	"".correctCase(SB)
	0x0024 00036 (instance.go:11)	MOVD	-8(RSP), R29
	0x0028 00040 (instance.go:11)	MOVD.P	16(RSP), R30
	0x002c 00044 (instance.go:11)	RET	(R30)
	0x0030 00048 (instance.go:11)	NOP
	0x0030 00048 (instance.go:8)	PCDATA	$1, $-1
	0x0030 00048 (instance.go:8)	PCDATA	$0, $-2
	0x0030 00048 (instance.go:8)	MOVD	R30, R3
	0x0034 00052 (instance.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x0038 00056 (instance.go:8)	PCDATA	$0, $-1
	0x0038 00056 (instance.go:8)	JMP	0
	0x0000 90 0b 40 f9 f1 03 00 91 3f 02 10 eb 29 01 00 54  ..@.....?...)..T
	0x0010 fe 0f 1f f8 fd 83 1f f8 fd 23 00 d1 00 00 00 94  .........#......
	0x0020 00 00 00 94 fd 83 5f f8 fe 07 41 f8 c0 03 5f d6  ......_...A..._.
	0x0030 e3 03 1e aa 00 00 00 94 f2 ff ff 17 00 00 00 00  ................
	rel 28+4 t=9 "".firstCase+0
	rel 32+4 t=9 "".correctCase+0
	rel 52+4 t=9 runtime.morestack_noctxt+0
"".firstCase STEXT size=240 args=0x0 locals=0x38 funcid=0x0 align=0x0
	0x0000 00000 (instance.go:13)	TEXT	"".firstCase(SB), ABIInternal, $64-0
	0x0000 00000 (instance.go:13)	MOVD	16(g), R16
	0x0004 00004 (instance.go:13)	PCDATA	$0, $-2
	0x0004 00004 (instance.go:13)	MOVD	RSP, R17
	0x0008 00008 (instance.go:13)	CMP	R16, R17
	0x000c 00012 (instance.go:13)	BLS	220
	0x0010 00016 (instance.go:13)	PCDATA	$0, $-1
	0x0010 00016 (instance.go:13)	MOVD.W	R30, -64(RSP)
	0x0014 00020 (instance.go:13)	MOVD	R29, -8(RSP)
	0x0018 00024 (instance.go:13)	SUB	$8, RSP, R29
	0x001c 00028 (instance.go:13)	FUNCDATA	ZR, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x001c 00028 (instance.go:13)	FUNCDATA	$1, gclocals·bfec7e55b3f043d1941c093912808913(SB)
	0x001c 00028 (instance.go:14)	MOVD	$type.sync.WaitGroup(SB), R0
	0x0024 00036 (instance.go:14)	PCDATA	$1, ZR
	0x0024 00036 (instance.go:14)	CALL	runtime.newobject(SB)
	0x0028 00040 (instance.go:14)	MOVD	R0, "".wg-16(SP)
	0x002c 00044 (instance.go:16)	MOVD	$type.int(SB), R0
	0x0034 00052 (instance.go:16)	PCDATA	$1, $1
	0x0034 00052 (instance.go:16)	CALL	runtime.newobject(SB)
	0x0038 00056 (instance.go:16)	MOVD	R0, "".&i-8(SP)
	0x003c 00060 (instance.go:16)	JMP	84
	0x0040 00064 (instance.go:19)	PCDATA	$1, $2
	0x0040 00064 (instance.go:19)	CALL	runtime.newproc(SB)
	0x0044 00068 (instance.go:16)	MOVD	"".&i-8(SP), R0
	0x0048 00072 (instance.go:16)	MOVD	(R0), R1
	0x004c 00076 (instance.go:16)	ADD	$1, R1, R1
	0x0050 00080 (instance.go:16)	MOVD	R1, (R0)
	0x0054 00084 (instance.go:16)	MOVD	(R0), R2
	0x0058 00088 (instance.go:16)	CMP	$10, R2
	0x005c 00092 (instance.go:16)	BGE	200
	0x0060 00096 (instance.go:17)	MOVD	"".wg-16(SP), R0
	0x0064 00100 (instance.go:17)	MOVD	$1, R1
	0x0068 00104 (instance.go:17)	CALL	sync.(*WaitGroup).Add(SB)
	0x006c 00108 (instance.go:19)	MOVD	$type.noalg.struct { F uintptr; "".i *int; "".wg *sync.WaitGroup }(SB), R0
	0x0074 00116 (instance.go:19)	CALL	runtime.newobject(SB)
	0x0078 00120 (instance.go:19)	MOVD	$"".firstCase.func1(SB), R2
	0x0080 00128 (instance.go:19)	MOVD	R2, (R0)
	0x0084 00132 (instance.go:19)	PCDATA	ZR, $-2
	0x0084 00132 (instance.go:19)	MOVWU	runtime.writeBarrier(SB), R3
	0x0090 00144 (instance.go:19)	CBNZW	R3, 168
	0x0094 00148 (instance.go:19)	MOVD	"".&i-8(SP), R1
	0x0098 00152 (instance.go:19)	MOVD	R1, 8(R0)
	0x009c 00156 (instance.go:19)	MOVD	"".wg-16(SP), R3
	0x00a0 00160 (instance.go:19)	MOVD	R3, 16(R0)
	0x00a4 00164 (instance.go:19)	JMP	64
	0x00a8 00168 (instance.go:19)	ADD	$8, R0, R1
	0x00ac 00172 (instance.go:19)	MOVD	"".&i-8(SP), R3
	0x00b0 00176 (instance.go:19)	MOVD	R1, R2
	0x00b4 00180 (instance.go:19)	CALL	runtime.gcWriteBarrier(SB)
	0x00b8 00184 (instance.go:19)	ADD	$16, R0, R2
	0x00bc 00188 (instance.go:19)	MOVD	"".wg-16(SP), R3
	0x00c0 00192 (instance.go:19)	CALL	runtime.gcWriteBarrier(SB)
	0x00c4 00196 (instance.go:19)	JMP	64
	0x00c8 00200 (instance.go:27)	PCDATA	ZR, $-1
	0x00c8 00200 (instance.go:27)	MOVD	"".wg-16(SP), R0
	0x00cc 00204 (instance.go:27)	PCDATA	$1, ZR
	0x00cc 00204 (instance.go:27)	CALL	sync.(*WaitGroup).Wait(SB)
	0x00d0 00208 (instance.go:28)	MOVD	-8(RSP), R29
	0x00d4 00212 (instance.go:28)	MOVD.P	64(RSP), R30
	0x00d8 00216 (instance.go:28)	RET	(R30)
	0x00dc 00220 (instance.go:28)	NOP
	0x00dc 00220 (instance.go:13)	PCDATA	$1, $-1
	0x00dc 00220 (instance.go:13)	PCDATA	$0, $-2
	0x00dc 00220 (instance.go:13)	MOVD	R30, R3
	0x00e0 00224 (instance.go:13)	CALL	runtime.morestack_noctxt(SB)
	0x00e4 00228 (instance.go:13)	PCDATA	$0, $-1
	0x00e4 00228 (instance.go:13)	JMP	0
	0x0000 90 0b 40 f9 f1 03 00 91 3f 02 10 eb 89 06 00 54  ..@.....?......T
	0x0010 fe 0f 1c f8 fd 83 1f f8 fd 23 00 d1 00 00 00 90  .........#......
	0x0020 00 00 00 91 00 00 00 94 e0 17 00 f9 00 00 00 90  ................
	0x0030 00 00 00 91 00 00 00 94 e0 1b 00 f9 06 00 00 14  ................
	0x0040 00 00 00 94 e0 1b 40 f9 01 00 40 f9 21 04 00 91  ......@...@.!...
	0x0050 01 00 00 f9 02 00 40 f9 5f 28 00 f1 6a 03 00 54  ......@._(..j..T
	0x0060 e0 17 40 f9 e1 03 40 b2 00 00 00 94 00 00 00 90  ..@...@.........
	0x0070 00 00 00 91 00 00 00 94 02 00 00 90 42 00 00 91  ............B...
	0x0080 02 00 00 f9 1b 00 00 90 7b 03 00 91 63 03 40 b9  ........{...c.@.
	0x0090 c3 00 00 35 e1 1b 40 f9 01 04 00 f9 e3 17 40 f9  ...5..@.......@.
	0x00a0 03 08 00 f9 e7 ff ff 17 01 20 00 91 e3 1b 40 f9  ......... ....@.
	0x00b0 e2 03 01 aa 00 00 00 94 02 40 00 91 e3 17 40 f9  .........@....@.
	0x00c0 00 00 00 94 df ff ff 17 e0 17 40 f9 00 00 00 94  ..........@.....
	0x00d0 fd 83 5f f8 fe 07 44 f8 c0 03 5f d6 e3 03 1e aa  .._...D..._.....
	0x00e0 00 00 00 94 c7 ff ff 17 00 00 00 00 00 00 00 00  ................
	rel 28+8 t=3 type.sync.WaitGroup+0
	rel 36+4 t=9 runtime.newobject+0
	rel 44+8 t=3 type.int+0
	rel 52+4 t=9 runtime.newobject+0
	rel 64+4 t=9 runtime.newproc+0
	rel 104+4 t=9 sync.(*WaitGroup).Add+0
	rel 108+8 t=3 type.noalg.struct { F uintptr; "".i *int; "".wg *sync.WaitGroup }+0
	rel 116+4 t=9 runtime.newobject+0
	rel 120+8 t=3 "".firstCase.func1+0
	rel 132+8 t=3 runtime.writeBarrier+0
	rel 180+4 t=9 runtime.gcWriteBarrier+0
	rel 192+4 t=9 runtime.gcWriteBarrier+0
	rel 204+4 t=9 sync.(*WaitGroup).Wait+0
	rel 224+4 t=9 runtime.morestack_noctxt+0
"".firstCase.func1 STEXT size=144 args=0x0 locals=0x48 funcid=0x0 align=0x0
	0x0000 00000 (instance.go:19)	TEXT	"".firstCase.func1(SB), NEEDCTXT|ABIInternal, $80-0
	0x0000 00000 (instance.go:19)	MOVD	16(g), R16
	0x0004 00004 (instance.go:19)	PCDATA	$0, $-2
	0x0004 00004 (instance.go:19)	MOVD	RSP, R17
	0x0008 00008 (instance.go:19)	CMP	R16, R17
	0x000c 00012 (instance.go:19)	BLS	132
	0x0010 00016 (instance.go:19)	PCDATA	$0, $-1
	0x0010 00016 (instance.go:19)	MOVD.W	R30, -80(RSP)
	0x0014 00020 (instance.go:19)	MOVD	R29, -8(RSP)
	0x0018 00024 (instance.go:19)	SUB	$8, RSP, R29
	0x001c 00028 (instance.go:19)	FUNCDATA	ZR, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x001c 00028 (instance.go:19)	FUNCDATA	$1, gclocals·ab5b6194452d266823944709d4cd8167(SB)
	0x001c 00028 (instance.go:19)	FUNCDATA	$2, "".firstCase.func1.stkobj(SB)
	0x001c 00028 (instance.go:19)	MOVD	16(R26), R1
	0x0020 00032 (instance.go:19)	MOVD	R1, "".wg-24(SP)
	0x0024 00036 (instance.go:19)	MOVD	8(R26), R2
	0x0028 00040 (instance.go:20)	STP	(ZR, ZR), ""..autotmp_9-16(SP)
	0x002c 00044 (instance.go:20)	MOVD	(R2), R0
	0x0030 00048 (instance.go:20)	PCDATA	$1, $1
	0x0030 00048 (instance.go:20)	CALL	runtime.convT64(SB)
	0x0034 00052 (instance.go:20)	MOVD	$type.int(SB), R1
	0x003c 00060 (instance.go:20)	MOVD	R1, ""..autotmp_9-16(SP)
	0x0040 00064 (instance.go:20)	MOVD	R0, ""..autotmp_9-8(SP)
	0x0044 00068 (<unknown line number>)	NOP
	0x0044 00068 (<unknown line number>)	PCDATA	$0, $-3
	0x0044 00068 ($GOROOT/src/fmt/print.go:274)	MOVD	os.Stdout(SB), R1
	0x0050 00080 ($GOROOT/src/fmt/print.go:274)	PCDATA	$0, $-1
	0x0050 00080 ($GOROOT/src/fmt/print.go:274)	MOVD	$go.itab.*os.File,io.Writer(SB), R0
	0x0058 00088 ($GOROOT/src/fmt/print.go:274)	MOVD	$""..autotmp_9-16(SP), R2
	0x005c 00092 ($GOROOT/src/fmt/print.go:274)	MOVD	$1, R3
	0x0060 00096 ($GOROOT/src/fmt/print.go:274)	MOVD	R3, R4
	0x0064 00100 ($GOROOT/src/fmt/print.go:274)	PCDATA	$1, $2
	0x0064 00100 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0068 00104 (instance.go:22)	HINT	ZR
	0x006c 00108 ($GOROOT/src/sync/waitgroup.go:105)	MOVD	"".wg-24(SP), R0
	0x0070 00112 ($GOROOT/src/sync/waitgroup.go:105)	MOVD	$-1, R1
	0x0074 00116 ($GOROOT/src/sync/waitgroup.go:105)	PCDATA	$1, ZR
	0x0074 00116 ($GOROOT/src/sync/waitgroup.go:105)	CALL	sync.(*WaitGroup).Add(SB)
	0x0078 00120 (instance.go:23)	MOVD	-8(RSP), R29
	0x007c 00124 (instance.go:23)	MOVD.P	80(RSP), R30
	0x0080 00128 (instance.go:23)	RET	(R30)
	0x0084 00132 (instance.go:23)	NOP
	0x0084 00132 (instance.go:19)	PCDATA	$1, $-1
	0x0084 00132 (instance.go:19)	PCDATA	$0, $-2
	0x0084 00132 (instance.go:19)	MOVD	R30, R3
	0x0088 00136 (instance.go:19)	CALL	runtime.morestack(SB)
	0x008c 00140 (instance.go:19)	PCDATA	$0, $-1
	0x008c 00140 (instance.go:19)	JMP	0
	0x0000 90 0b 40 f9 f1 03 00 91 3f 02 10 eb c9 03 00 54  ..@.....?......T
	0x0010 fe 0f 1b f8 fd 83 1f f8 fd 23 00 d1 41 0b 40 f9  .........#..A.@.
	0x0020 e1 1b 00 f9 42 07 40 f9 ff ff 03 a9 40 00 40 f9  ....B.@.....@.@.
	0x0030 00 00 00 94 01 00 00 90 21 00 00 91 e1 1f 00 f9  ........!.......
	0x0040 e0 23 00 f9 1b 00 00 90 7b 03 00 91 61 03 40 f9  .#......{...a.@.
	0x0050 00 00 00 90 00 00 00 91 e2 e3 00 91 e3 03 40 b2  ..............@.
	0x0060 e4 03 03 aa 00 00 00 94 1f 20 03 d5 e0 1b 40 f9  ......... ....@.
	0x0070 01 00 80 92 00 00 00 94 fd 83 5f f8 fe 07 45 f8  .........._...E.
	0x0080 c0 03 5f d6 e3 03 1e aa 00 00 00 94 dd ff ff 17  .._.............
	rel 0+0 t=23 type.int+0
	rel 0+0 t=23 type.*os.File+0
	rel 48+4 t=9 runtime.convT64+0
	rel 52+8 t=3 type.int+0
	rel 68+8 t=3 os.Stdout+0
	rel 80+8 t=3 go.itab.*os.File,io.Writer+0
	rel 100+4 t=9 fmt.Fprintln+0
	rel 116+4 t=9 sync.(*WaitGroup).Add+0
	rel 136+4 t=9 runtime.morestack+0
"".correctCase STEXT size=400 args=0x0 locals=0x68 funcid=0x0 align=0x0
	0x0000 00000 (instance.go:30)	TEXT	"".correctCase(SB), ABIInternal, $112-0
	0x0000 00000 (instance.go:30)	MOVD	16(g), R16
	0x0004 00004 (instance.go:30)	PCDATA	$0, $-2
	0x0004 00004 (instance.go:30)	MOVD	RSP, R17
	0x0008 00008 (instance.go:30)	CMP	R16, R17
	0x000c 00012 (instance.go:30)	BLS	388
	0x0010 00016 (instance.go:30)	PCDATA	$0, $-1
	0x0010 00016 (instance.go:30)	MOVD.W	R30, -112(RSP)
	0x0014 00020 (instance.go:30)	MOVD	R29, -8(RSP)
	0x0018 00024 (instance.go:30)	SUB	$8, RSP, R29
	0x001c 00028 (instance.go:30)	FUNCDATA	ZR, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x001c 00028 (instance.go:30)	FUNCDATA	$1, gclocals·06278046951d8fddcb532a8e07e1db1a(SB)
	0x001c 00028 (instance.go:30)	FUNCDATA	$2, "".correctCase.stkobj(SB)
	0x001c 00028 (instance.go:31)	MOVD	$type.sync.WaitGroup(SB), R0
	0x0024 00036 (instance.go:31)	PCDATA	$1, ZR
	0x0024 00036 (instance.go:31)	CALL	runtime.newobject(SB)
	0x0028 00040 (instance.go:31)	MOVD	R0, "".wg-48(SP)
	0x002c 00044 (instance.go:32)	STP	(ZR, ZR), ""..autotmp_18-16(SP)
	0x0030 00048 (instance.go:32)	MOVD	$type.string(SB), R1
	0x0038 00056 (instance.go:32)	MOVD	R1, ""..autotmp_18-16(SP)
	0x003c 00060 (instance.go:32)	MOVD	$""..stmp_0(SB), R2
	0x0044 00068 (instance.go:32)	MOVD	R2, ""..autotmp_18-8(SP)
	0x0048 00072 (<unknown line number>)	NOP
	0x0048 00072 (<unknown line number>)	PCDATA	$0, $-3
	0x0048 00072 ($GOROOT/src/fmt/print.go:274)	MOVD	os.Stdout(SB), R2
	0x0054 00084 ($GOROOT/src/fmt/print.go:274)	PCDATA	$0, $-1
	0x0054 00084 ($GOROOT/src/fmt/print.go:274)	MOVD	$1, R3
	0x0058 00088 ($GOROOT/src/fmt/print.go:274)	MOVD	R3, R4
	0x005c 00092 ($GOROOT/src/fmt/print.go:274)	MOVD	$go.itab.*os.File,io.Writer(SB), R0
	0x0064 00100 ($GOROOT/src/fmt/print.go:274)	MOVD	R2, R1
	0x0068 00104 ($GOROOT/src/fmt/print.go:274)	MOVD	$""..autotmp_18-16(SP), R2
	0x006c 00108 ($GOROOT/src/fmt/print.go:274)	PCDATA	$1, $1
	0x006c 00108 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0070 00112 ($GOROOT/src/fmt/print.go:274)	MOVD	ZR, R0
	0x0074 00116 (instance.go:32)	JMP	140
	0x0078 00120 (instance.go:36)	MOVD	"".i-56(SP), R1
	0x007c 00124 (instance.go:36)	MOVD	R1, 16(R0)
	0x0080 00128 (instance.go:36)	CALL	runtime.newproc(SB)
	0x0084 00132 (instance.go:34)	MOVD	"".i-56(SP), R1
	0x0088 00136 (instance.go:34)	ADD	$1, R1, R0
	0x008c 00140 (instance.go:34)	CMP	$10, R0
	0x0090 00144 (instance.go:34)	BGE	304
	0x0094 00148 (instance.go:34)	MOVD	R0, "".i-56(SP)
	0x0098 00152 (instance.go:35)	MOVD	"".wg-48(SP), R0
	0x009c 00156 (instance.go:35)	MOVD	$1, R1
	0x00a0 00160 (instance.go:35)	CALL	sync.(*WaitGroup).Add(SB)
	0x00a4 00164 (instance.go:36)	MOVD	$type.noalg.struct { F uintptr; "".wg *sync.WaitGroup }(SB), R0
	0x00ac 00172 (instance.go:36)	CALL	runtime.newobject(SB)
	0x00b0 00176 (instance.go:36)	MOVD	R0, ""..autotmp_29-40(SP)
	0x00b4 00180 (instance.go:36)	MOVD	$"".correctCase.func1(SB), R2
	0x00bc 00188 (instance.go:36)	MOVD	R2, (R0)
	0x00c0 00192 (instance.go:36)	PCDATA	ZR, $-2
	0x00c0 00192 (instance.go:36)	MOVWU	runtime.writeBarrier(SB), R3
	0x00cc 00204 (instance.go:36)	CBNZW	R3, 220
	0x00d0 00208 (instance.go:36)	MOVD	"".wg-48(SP), R1
	0x00d4 00212 (instance.go:36)	MOVD	R1, 8(R0)
	0x00d8 00216 (instance.go:36)	JMP	236
	0x00dc 00220 (instance.go:36)	ADD	$8, R0, R1
	0x00e0 00224 (instance.go:36)	MOVD	"".wg-48(SP), R3
	0x00e4 00228 (instance.go:36)	MOVD	R1, R2
	0x00e8 00232 (instance.go:36)	CALL	runtime.gcWriteBarrier(SB)
	0x00ec 00236 (instance.go:36)	PCDATA	ZR, $-1
	0x00ec 00236 (instance.go:36)	MOVD	$type.noalg.struct { F uintptr; ""..autotmp_12 func(int); ""..autotmp_13 int }(SB), R0
	0x00f4 00244 (instance.go:36)	PCDATA	$1, $2
	0x00f4 00244 (instance.go:36)	CALL	runtime.newobject(SB)
	0x00f8 00248 (instance.go:36)	MOVD	$"".correctCase.func2(SB), R1
	0x0100 00256 (instance.go:36)	MOVD	R1, (R0)
	0x0104 00260 (instance.go:36)	PCDATA	ZR, $-2
	0x0104 00260 (instance.go:36)	MOVWU	runtime.writeBarrier(SB), R2
	0x0110 00272 (instance.go:36)	CBNZW	R2, 288
	0x0114 00276 (instance.go:36)	MOVD	""..autotmp_29-40(SP), R2
	0x0118 00280 (instance.go:36)	MOVD	R2, 8(R0)
	0x011c 00284 (instance.go:36)	JMP	120
	0x0120 00288 (instance.go:36)	ADD	$8, R0, R2
	0x0124 00292 (instance.go:36)	MOVD	""..autotmp_29-40(SP), R3
	0x0128 00296 (instance.go:36)	CALL	runtime.gcWriteBarrier(SB)
	0x012c 00300 (instance.go:36)	JMP	120
	0x0130 00304 (instance.go:42)	PCDATA	ZR, $-1
	0x0130 00304 (instance.go:42)	MOVD	"".wg-48(SP), R0
	0x0134 00308 (instance.go:42)	PCDATA	$1, ZR
	0x0134 00308 (instance.go:42)	CALL	sync.(*WaitGroup).Wait(SB)
	0x0138 00312 (instance.go:43)	STP	(ZR, ZR), ""..autotmp_22-32(SP)
	0x013c 00316 (instance.go:43)	MOVD	$type.string(SB), R1
	0x0144 00324 (instance.go:43)	MOVD	R1, ""..autotmp_22-32(SP)
	0x0148 00328 (instance.go:43)	MOVD	$""..stmp_1(SB), R1
	0x0150 00336 (instance.go:43)	MOVD	R1, ""..autotmp_22-24(SP)
	0x0154 00340 (<unknown line number>)	NOP
	0x0154 00340 (<unknown line number>)	PCDATA	$0, $-4
	0x0154 00340 ($GOROOT/src/fmt/print.go:274)	MOVD	os.Stdout(SB), R1
	0x0160 00352 ($GOROOT/src/fmt/print.go:274)	PCDATA	$0, $-1
	0x0160 00352 ($GOROOT/src/fmt/print.go:274)	MOVD	$go.itab.*os.File,io.Writer(SB), R0
	0x0168 00360 ($GOROOT/src/fmt/print.go:274)	MOVD	$""..autotmp_22-32(SP), R2
	0x016c 00364 ($GOROOT/src/fmt/print.go:274)	MOVD	$1, R3
	0x0170 00368 ($GOROOT/src/fmt/print.go:274)	MOVD	R3, R4
	0x0174 00372 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0178 00376 (instance.go:44)	MOVD	-8(RSP), R29
	0x017c 00380 (instance.go:44)	MOVD.P	112(RSP), R30
	0x0180 00384 (instance.go:44)	RET	(R30)
	0x0184 00388 (instance.go:44)	NOP
	0x0184 00388 (instance.go:30)	PCDATA	$1, $-1
	0x0184 00388 (instance.go:30)	PCDATA	$0, $-2
	0x0184 00388 (instance.go:30)	MOVD	R30, R3
	0x0188 00392 (instance.go:30)	CALL	runtime.morestack_noctxt(SB)
	0x018c 00396 (instance.go:30)	PCDATA	$0, $-1
	0x018c 00396 (instance.go:30)	JMP	0
	0x0000 90 0b 40 f9 f1 03 00 91 3f 02 10 eb c9 0b 00 54  ..@.....?......T
	0x0010 fe 0f 19 f8 fd 83 1f f8 fd 23 00 d1 00 00 00 90  .........#......
	0x0020 00 00 00 91 00 00 00 94 e0 1f 00 f9 ff ff 05 a9  ................
	0x0030 01 00 00 90 21 00 00 91 e1 2f 00 f9 02 00 00 90  ....!..../......
	0x0040 42 00 00 91 e2 33 00 f9 1b 00 00 90 7b 03 00 91  B....3......{...
	0x0050 62 03 40 f9 e3 03 40 b2 e4 03 03 aa 00 00 00 90  b.@...@.........
	0x0060 00 00 00 91 e1 03 02 aa e2 63 01 91 00 00 00 94  .........c......
	0x0070 00 00 80 d2 06 00 00 14 e1 1b 40 f9 01 08 00 f9  ..........@.....
	0x0080 00 00 00 94 e1 1b 40 f9 20 04 00 91 1f 28 00 f1  ......@. ....(..
	0x0090 0a 05 00 54 e0 1b 00 f9 e0 1f 40 f9 e1 03 40 b2  ...T......@...@.
	0x00a0 00 00 00 94 00 00 00 90 00 00 00 91 00 00 00 94  ................
	0x00b0 e0 23 00 f9 02 00 00 90 42 00 00 91 02 00 00 f9  .#......B.......
	0x00c0 1b 00 00 90 7b 03 00 91 63 03 40 b9 83 00 00 35  ....{...c.@....5
	0x00d0 e1 1f 40 f9 01 04 00 f9 05 00 00 14 01 20 00 91  ..@.......... ..
	0x00e0 e3 1f 40 f9 e2 03 01 aa 00 00 00 94 00 00 00 90  ..@.............
	0x00f0 00 00 00 91 00 00 00 94 01 00 00 90 21 00 00 91  ............!...
	0x0100 01 00 00 f9 1b 00 00 90 7b 03 00 91 62 03 40 b9  ........{...b.@.
	0x0110 82 00 00 35 e2 23 40 f9 02 04 00 f9 d7 ff ff 17  ...5.#@.........
	0x0120 02 20 00 91 e3 23 40 f9 00 00 00 94 d3 ff ff 17  . ...#@.........
	0x0130 e0 1f 40 f9 00 00 00 94 ff ff 04 a9 01 00 00 90  ..@.............
	0x0140 21 00 00 91 e1 27 00 f9 01 00 00 90 21 00 00 91  !....'......!...
	0x0150 e1 2b 00 f9 1b 00 00 90 7b 03 00 91 61 03 40 f9  .+......{...a.@.
	0x0160 00 00 00 90 00 00 00 91 e2 23 01 91 e3 03 40 b2  .........#....@.
	0x0170 e4 03 03 aa 00 00 00 94 fd 83 5f f8 fe 07 47 f8  .........._...G.
	0x0180 c0 03 5f d6 e3 03 1e aa 00 00 00 94 9d ff ff 17  .._.............
	rel 0+0 t=23 type.string+0
	rel 0+0 t=23 type.*os.File+0
	rel 0+0 t=23 type.string+0
	rel 0+0 t=23 type.*os.File+0
	rel 28+8 t=3 type.sync.WaitGroup+0
	rel 36+4 t=9 runtime.newobject+0
	rel 48+8 t=3 type.string+0
	rel 60+8 t=3 ""..stmp_0+0
	rel 72+8 t=3 os.Stdout+0
	rel 92+8 t=3 go.itab.*os.File,io.Writer+0
	rel 108+4 t=9 fmt.Fprintln+0
	rel 128+4 t=9 runtime.newproc+0
	rel 160+4 t=9 sync.(*WaitGroup).Add+0
	rel 164+8 t=3 type.noalg.struct { F uintptr; "".wg *sync.WaitGroup }+0
	rel 172+4 t=9 runtime.newobject+0
	rel 180+8 t=3 "".correctCase.func1+0
	rel 192+8 t=3 runtime.writeBarrier+0
	rel 232+4 t=9 runtime.gcWriteBarrier+0
	rel 236+8 t=3 type.noalg.struct { F uintptr; ""..autotmp_12 func(int); ""..autotmp_13 int }+0
	rel 244+4 t=9 runtime.newobject+0
	rel 248+8 t=3 "".correctCase.func2+0
	rel 260+8 t=3 runtime.writeBarrier+0
	rel 296+4 t=9 runtime.gcWriteBarrier+0
	rel 308+4 t=9 sync.(*WaitGroup).Wait+0
	rel 316+8 t=3 type.string+0
	rel 328+8 t=3 ""..stmp_1+0
	rel 340+8 t=3 os.Stdout+0
	rel 352+8 t=3 go.itab.*os.File,io.Writer+0
	rel 372+4 t=9 fmt.Fprintln+0
	rel 392+4 t=9 runtime.morestack_noctxt+0
"".correctCase.func2 STEXT size=112 args=0x0 locals=0x18 funcid=0x15 align=0x0
	0x0000 00000 (instance.go:36)	TEXT	"".correctCase.func2(SB), WRAPPER|NEEDCTXT|ABIInternal, $32-0
	0x0000 00000 (instance.go:36)	MOVD	16(g), R16
	0x0004 00004 (instance.go:36)	PCDATA	$0, $-2
	0x0004 00004 (instance.go:36)	MOVD	RSP, R17
	0x0008 00008 (instance.go:36)	CMP	R16, R17
	0x000c 00012 (instance.go:36)	BLS	64
	0x0010 00016 (instance.go:36)	PCDATA	$0, $-1
	0x0010 00016 (instance.go:36)	MOVD.W	R30, -32(RSP)
	0x0014 00020 (instance.go:36)	MOVD	R29, -8(RSP)
	0x0018 00024 (instance.go:36)	SUB	$8, RSP, R29
	0x001c 00028 (instance.go:36)	MOVD	32(g), R16
	0x0020 00032 (instance.go:36)	CBNZ	R16, 76
	0x0024 00036 (instance.go:36)	NOP
	0x0024 00036 (instance.go:36)	FUNCDATA	ZR, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0024 00036 (instance.go:36)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0024 00036 (instance.go:36)	MOVD	16(R26), R0
	0x0028 00040 (instance.go:36)	MOVD	8(R26), R26
	0x002c 00044 (instance.go:39)	MOVD	(R26), R1
	0x0030 00048 (instance.go:39)	PCDATA	$1, ZR
	0x0030 00048 (instance.go:39)	CALL	(R1)
	0x0034 00052 (instance.go:39)	MOVD	-8(RSP), R29
	0x0038 00056 (instance.go:39)	MOVD.P	32(RSP), R30
	0x003c 00060 (instance.go:39)	RET	(R30)
	0x0040 00064 (instance.go:39)	NOP
	0x0040 00064 (instance.go:36)	PCDATA	$1, $-1
	0x0040 00064 (instance.go:36)	PCDATA	$0, $-2
	0x0040 00064 (instance.go:36)	MOVD	R30, R3
	0x0044 00068 (instance.go:36)	CALL	runtime.morestack(SB)
	0x0048 00072 (instance.go:36)	PCDATA	$0, $-1
	0x0048 00072 (instance.go:36)	JMP	0
	0x004c 00076 (instance.go:36)	MOVD	(R16), R17
	0x0050 00080 (instance.go:36)	ADD	$40, RSP, R20
	0x0054 00084 (instance.go:36)	CMP	R17, R20
	0x0058 00088 (instance.go:36)	BNE	36
	0x005c 00092 (instance.go:36)	ADD	$8, RSP, R20
	0x0060 00096 (instance.go:36)	MOVD	R20, (R16)
	0x0064 00100 (instance.go:36)	JMP	36
	0x0000 90 0b 40 f9 f1 03 00 91 3f 02 10 eb a9 01 00 54  ..@.....?......T
	0x0010 fe 0f 1e f8 fd 83 1f f8 fd 23 00 d1 90 13 40 f9  .........#....@.
	0x0020 70 01 00 b5 40 0b 40 f9 5a 07 40 f9 41 03 40 f9  p...@.@.Z.@.A.@.
	0x0030 20 00 3f d6 fd 83 5f f8 fe 07 42 f8 c0 03 5f d6   .?..._...B..._.
	0x0040 e3 03 1e aa 00 00 00 94 ee ff ff 17 11 02 40 f9  ..............@.
	0x0050 f4 a3 00 91 9f 02 11 eb 61 fe ff 54 f4 23 00 91  ........a..T.#..
	0x0060 14 02 00 f9 f0 ff ff 17 00 00 00 00 00 00 00 00  ................
	rel 48+0 t=10 +0
	rel 68+4 t=9 runtime.morestack+0
"".correctCase.func1 STEXT size=144 args=0x8 locals=0x48 funcid=0x0 align=0x0
	0x0000 00000 (instance.go:36)	TEXT	"".correctCase.func1(SB), NEEDCTXT|ABIInternal, $80-8
	0x0000 00000 (instance.go:36)	MOVD	16(g), R16
	0x0004 00004 (instance.go:36)	PCDATA	$0, $-2
	0x0004 00004 (instance.go:36)	MOVD	RSP, R17
	0x0008 00008 (instance.go:36)	CMP	R16, R17
	0x000c 00012 (instance.go:36)	BLS	124
	0x0010 00016 (instance.go:36)	PCDATA	$0, $-1
	0x0010 00016 (instance.go:36)	MOVD.W	R30, -80(RSP)
	0x0014 00020 (instance.go:36)	MOVD	R29, -8(RSP)
	0x0018 00024 (instance.go:36)	SUB	$8, RSP, R29
	0x001c 00028 (instance.go:36)	FUNCDATA	ZR, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x001c 00028 (instance.go:36)	FUNCDATA	$1, gclocals·ab5b6194452d266823944709d4cd8167(SB)
	0x001c 00028 (instance.go:36)	FUNCDATA	$2, "".correctCase.func1.stkobj(SB)
	0x001c 00028 (instance.go:36)	FUNCDATA	$5, "".correctCase.func1.arginfo1(SB)
	0x001c 00028 (instance.go:36)	FUNCDATA	$6, "".correctCase.func1.argliveinfo(SB)
	0x001c 00028 (instance.go:36)	PCDATA	$3, $1
	0x001c 00028 (instance.go:36)	MOVD	8(R26), R1
	0x0020 00032 (instance.go:36)	MOVD	R1, "".wg-24(SP)
	0x0024 00036 (instance.go:37)	STP	(ZR, ZR), ""..autotmp_10-16(SP)
	0x0028 00040 (instance.go:37)	PCDATA	$1, $1
	0x0028 00040 (instance.go:37)	CALL	runtime.convT64(SB)
	0x002c 00044 (instance.go:37)	MOVD	$type.int(SB), R1
	0x0034 00052 (instance.go:37)	MOVD	R1, ""..autotmp_10-16(SP)
	0x0038 00056 (instance.go:37)	MOVD	R0, ""..autotmp_10-8(SP)
	0x003c 00060 (<unknown line number>)	NOP
	0x003c 00060 (<unknown line number>)	PCDATA	$0, $-3
	0x003c 00060 ($GOROOT/src/fmt/print.go:274)	MOVD	os.Stdout(SB), R1
	0x0048 00072 ($GOROOT/src/fmt/print.go:274)	PCDATA	$0, $-1
	0x0048 00072 ($GOROOT/src/fmt/print.go:274)	MOVD	$go.itab.*os.File,io.Writer(SB), R0
	0x0050 00080 ($GOROOT/src/fmt/print.go:274)	MOVD	$""..autotmp_10-16(SP), R2
	0x0054 00084 ($GOROOT/src/fmt/print.go:274)	MOVD	$1, R3
	0x0058 00088 ($GOROOT/src/fmt/print.go:274)	MOVD	R3, R4
	0x005c 00092 ($GOROOT/src/fmt/print.go:274)	PCDATA	$1, $2
	0x005c 00092 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x0060 00096 (instance.go:38)	HINT	ZR
	0x0064 00100 ($GOROOT/src/sync/waitgroup.go:105)	MOVD	"".wg-24(SP), R0
	0x0068 00104 ($GOROOT/src/sync/waitgroup.go:105)	MOVD	$-1, R1
	0x006c 00108 ($GOROOT/src/sync/waitgroup.go:105)	PCDATA	$1, ZR
	0x006c 00108 ($GOROOT/src/sync/waitgroup.go:105)	CALL	sync.(*WaitGroup).Add(SB)
	0x0070 00112 (instance.go:39)	MOVD	-8(RSP), R29
	0x0074 00116 (instance.go:39)	MOVD.P	80(RSP), R30
	0x0078 00120 (instance.go:39)	RET	(R30)
	0x007c 00124 (instance.go:39)	NOP
	0x007c 00124 (instance.go:36)	PCDATA	$1, $-1
	0x007c 00124 (instance.go:36)	PCDATA	$0, $-2
	0x007c 00124 (instance.go:36)	MOVD	R0, 8(RSP)
	0x0080 00128 (instance.go:36)	MOVD	R30, R3
	0x0084 00132 (instance.go:36)	CALL	runtime.morestack(SB)
	0x0088 00136 (instance.go:36)	MOVD	8(RSP), R0
	0x008c 00140 (instance.go:36)	PCDATA	$0, $-1
	0x008c 00140 (instance.go:36)	JMP	0
	0x0000 90 0b 40 f9 f1 03 00 91 3f 02 10 eb 89 03 00 54  ..@.....?......T
	0x0010 fe 0f 1b f8 fd 83 1f f8 fd 23 00 d1 41 07 40 f9  .........#..A.@.
	0x0020 e1 1b 00 f9 ff ff 03 a9 00 00 00 94 01 00 00 90  ................
	0x0030 21 00 00 91 e1 1f 00 f9 e0 23 00 f9 1b 00 00 90  !........#......
	0x0040 7b 03 00 91 61 03 40 f9 00 00 00 90 00 00 00 91  {...a.@.........
	0x0050 e2 e3 00 91 e3 03 40 b2 e4 03 03 aa 00 00 00 94  ......@.........
	0x0060 1f 20 03 d5 e0 1b 40 f9 01 00 80 92 00 00 00 94  . ....@.........
	0x0070 fd 83 5f f8 fe 07 45 f8 c0 03 5f d6 e0 07 00 f9  .._...E..._.....
	0x0080 e3 03 1e aa 00 00 00 94 e0 07 40 f9 dd ff ff 17  ..........@.....
	rel 0+0 t=23 type.int+0
	rel 0+0 t=23 type.*os.File+0
	rel 40+4 t=9 runtime.convT64+0
	rel 44+8 t=3 type.int+0
	rel 60+8 t=3 os.Stdout+0
	rel 72+8 t=3 go.itab.*os.File,io.Writer+0
	rel 92+4 t=9 fmt.Fprintln+0
	rel 108+4 t=9 sync.(*WaitGroup).Add+0
	rel 132+4 t=9 runtime.morestack+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info.fmt.Println$abstract SDWARFABSFCN dupok size=42
	0x0000 05 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 13  .fmt.Println....
	0x0010 61 00 00 00 00 00 00 13 6e 00 01 00 00 00 00 13  a.......n.......
	0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
	rel 0+0 t=22 type.[]interface {}+0
	rel 0+0 t=22 type.error+0
	rel 0+0 t=22 type.int+0
	rel 19+4 t=31 go.info.[]interface {}+0
	rel 27+4 t=31 go.info.int+0
	rel 37+4 t=31 go.info.error+0
go.info.sync.(*WaitGroup).Done$abstract SDWARFABSFCN dupok size=36
	0x0000 05 73 79 6e 63 2e 28 2a 57 61 69 74 47 72 6f 75  .sync.(*WaitGrou
	0x0010 70 29 2e 44 6f 6e 65 00 01 01 13 77 67 00 00 00  p).Done....wg...
	0x0020 00 00 00 00                                      ....
	rel 0+0 t=22 type.*sync.WaitGroup+0
	rel 31+4 t=31 go.info.*sync.WaitGroup+0
""..inittask SNOPTRDATA size=40
	0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 fmt..inittask+0
	rel 32+8 t=1 sync..inittask+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=-32767 os.(*File).Write+0
go.string."-----correctCase start-----" SRODATA dupok size=27
	0x0000 2d 2d 2d 2d 2d 63 6f 72 72 65 63 74 43 61 73 65  -----correctCase
	0x0010 20 73 74 61 72 74 2d 2d 2d 2d 2d                  start-----
go.string."-----correctCase end-----" SRODATA dupok size=25
	0x0000 2d 2d 2d 2d 2d 63 6f 72 72 65 63 74 43 61 73 65  -----correctCase
	0x0010 20 65 6e 64 2d 2d 2d 2d 2d                        end-----
""..stmp_0 SRODATA static size=16
	0x0000 00 00 00 00 00 00 00 00 1b 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."-----correctCase start-----"+0
""..stmp_1 SRODATA static size=16
	0x0000 00 00 00 00 00 00 00 00 19 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."-----correctCase end-----"+0
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=15
	0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=-32763 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=17
	0x0000 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20 7b  ..*[]interface {
	0x0010 7d                                               }
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=-32763 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*func(int)- SRODATA dupok size=12
	0x0000 00 0a 2a 66 75 6e 63 28 69 6e 74 29              ..*func(int)
type.*func(int) SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 2b 1f 12 8c 08 08 08 36 00 00 00 00 00 00 00 00  +......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int)-+0
	rel 48+8 t=1 type.func(int)+0
type.func(int) SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 84 e6 f1 18 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int)-+0
	rel 44+4 t=-32763 type.*func(int)+0
	rel 56+8 t=1 type.int+0
type..namedata.*struct { F uintptr; .autotmp_12 func(int); .autotmp_13 int }- SRODATA dupok size=63
	0x0000 00 3d 2a 73 74 72 75 63 74 20 7b 20 46 20 75 69  .=*struct { F ui
	0x0010 6e 74 70 74 72 3b 20 2e 61 75 74 6f 74 6d 70 5f  ntptr; .autotmp_
	0x0020 31 32 20 66 75 6e 63 28 69 6e 74 29 3b 20 2e 61  12 func(int); .a
	0x0030 75 74 6f 74 6d 70 5f 31 33 20 69 6e 74 20 7d     utotmp_13 int }
type.*struct { F uintptr; ""..autotmp_12 func(int); ""..autotmp_13 int } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 2f c6 7d 47 08 08 08 36 00 00 00 00 00 00 00 00  /.}G...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; .autotmp_12 func(int); .autotmp_13 int }-+0
	rel 48+8 t=1 type.noalg.struct { F uintptr; ""..autotmp_12 func(int); ""..autotmp_13 int }+0
type..namedata..F- SRODATA dupok size=4
	0x0000 00 02 2e 46                                      ...F
type..namedata..autotmp_12- SRODATA dupok size=13
	0x0000 00 0b 2e 61 75 74 6f 74 6d 70 5f 31 32           ...autotmp_12
type..namedata..autotmp_13- SRODATA dupok size=13
	0x0000 00 0b 2e 61 75 74 6f 74 6d 70 5f 31 33           ...autotmp_13
type.noalg.struct { F uintptr; ""..autotmp_12 func(int); ""..autotmp_13 int } SRODATA dupok size=152
	0x0000 18 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 97 9b fa 20 02 08 08 19 00 00 00 00 00 00 00 00  ... ............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 03 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 20 00 00 00 00 00 00 00                           .......
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; .autotmp_12 func(int); .autotmp_13 int }-+0
	rel 44+4 t=-32763 type.*struct { F uintptr; ""..autotmp_12 func(int); ""..autotmp_13 int }+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type.noalg.struct { F uintptr; ""..autotmp_12 func(int); ""..autotmp_13 int }+80
	rel 80+8 t=1 type..namedata..F-+0
	rel 88+8 t=1 type.uintptr+0
	rel 104+8 t=1 type..namedata..autotmp_12-+0
	rel 112+8 t=1 type.func(int)+0
	rel 128+8 t=1 type..namedata..autotmp_13-+0
	rel 136+8 t=1 type.int+0
type..namedata.*struct { F uintptr; i *int; wg *sync.WaitGroup }- SRODATA dupok size=51
	0x0000 00 31 2a 73 74 72 75 63 74 20 7b 20 46 20 75 69  .1*struct { F ui
	0x0010 6e 74 70 74 72 3b 20 69 20 2a 69 6e 74 3b 20 77  ntptr; i *int; w
	0x0020 67 20 2a 73 79 6e 63 2e 57 61 69 74 47 72 6f 75  g *sync.WaitGrou
	0x0030 70 20 7d                                         p }
type.*struct { F uintptr; "".i *int; "".wg *sync.WaitGroup } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 15 96 e9 64 08 08 08 36 00 00 00 00 00 00 00 00  ...d...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; i *int; wg *sync.WaitGroup }-+0
	rel 48+8 t=1 type.noalg.struct { F uintptr; "".i *int; "".wg *sync.WaitGroup }+0
runtime.gcbits.06 SRODATA dupok size=1
	0x0000 06                                               .
type..namedata.i- SRODATA dupok size=3
	0x0000 00 01 69                                         ..i
type..namedata.wg- SRODATA dupok size=4
	0x0000 00 02 77 67                                      ..wg
type.noalg.struct { F uintptr; "".i *int; "".wg *sync.WaitGroup } SRODATA dupok size=152
	0x0000 18 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
	0x0010 7f 99 6e 3d 02 08 08 19 00 00 00 00 00 00 00 00  ..n=............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 03 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 20 00 00 00 00 00 00 00                           .......
	rel 32+8 t=1 runtime.gcbits.06+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; i *int; wg *sync.WaitGroup }-+0
	rel 44+4 t=-32763 type.*struct { F uintptr; "".i *int; "".wg *sync.WaitGroup }+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type.noalg.struct { F uintptr; "".i *int; "".wg *sync.WaitGroup }+80
	rel 80+8 t=1 type..namedata..F-+0
	rel 88+8 t=1 type.uintptr+0
	rel 104+8 t=1 type..namedata.i-+0
	rel 112+8 t=1 type.*int+0
	rel 128+8 t=1 type..namedata.wg-+0
	rel 136+8 t=1 type.*sync.WaitGroup+0
type..namedata.*struct { F uintptr; wg *sync.WaitGroup }- SRODATA dupok size=43
	0x0000 00 29 2a 73 74 72 75 63 74 20 7b 20 46 20 75 69  .)*struct { F ui
	0x0010 6e 74 70 74 72 3b 20 77 67 20 2a 73 79 6e 63 2e  ntptr; wg *sync.
	0x0020 57 61 69 74 47 72 6f 75 70 20 7d                 WaitGroup }
type.*struct { F uintptr; "".wg *sync.WaitGroup } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 44 ed ce 0a 08 08 08 36 00 00 00 00 00 00 00 00  D......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; wg *sync.WaitGroup }-+0
	rel 48+8 t=1 type.noalg.struct { F uintptr; "".wg *sync.WaitGroup }+0
type.noalg.struct { F uintptr; "".wg *sync.WaitGroup } SRODATA dupok size=128
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 44 c1 e8 50 02 08 08 19 00 00 00 00 00 00 00 00  D..P............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; wg *sync.WaitGroup }-+0
	rel 44+4 t=-32763 type.*struct { F uintptr; "".wg *sync.WaitGroup }+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type.noalg.struct { F uintptr; "".wg *sync.WaitGroup }+80
	rel 80+8 t=1 type..namedata..F-+0
	rel 88+8 t=1 type.uintptr+0
	rel 104+8 t=1 type..namedata.wg-+0
	rel 112+8 t=1 type.*sync.WaitGroup+0
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
type..importpath.sync. SRODATA dupok size=6
	0x0000 00 04 73 79 6e 63                                ..sync
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
	0x0000 03 00 00 00 00 00 00 00                          ........
gclocals·bfec7e55b3f043d1941c093912808913 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 01 03                 ...........
gclocals·ab5b6194452d266823944709d4cd8167 SRODATA dupok size=11
	0x0000 03 00 00 00 03 00 00 00 00 05 01                 ...........
"".firstCase.func1.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
gclocals·06278046951d8fddcb532a8e07e1db1a SRODATA dupok size=11
	0x0000 03 00 00 00 06 00 00 00 00 01 03                 ...........
"".correctCase.stkobj SRODATA static size=40
	0x0000 02 00 00 00 00 00 00 00 e0 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	0x0020 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
	rel 36+4 t=5 runtime.gcbits.02+0
"".correctCase.func1.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
"".correctCase.func1.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
"".correctCase.func1.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
