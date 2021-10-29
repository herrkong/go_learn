	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 11, 0	sdk_version 11, 3
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:
	sub	sp, sp, #48                     ; =48
	stp	x29, x30, [sp, #32]             ; 16-byte Folded Spill
	add	x29, sp, #32                    ; =32
	.cfi_def_cfa w29, 16
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	mov	w8, #0
	stur	wzr, [x29, #-4]
	ldur	w9, [x29, #-8]
                                        ; implicit-def: $x0
	mov	x0, x9
	adrp	x10, l_.str@PAGE
	add	x10, x10, l_.str@PAGEOFF
	str	x0, [sp, #16]                   ; 8-byte Folded Spill
	mov	x0, x10
	mov	x10, sp
	ldr	x11, [sp, #16]                  ; 8-byte Folded Reload
	str	x11, [x10]
	str	w8, [sp, #12]                   ; 4-byte Folded Spill
	bl	_printf
	ldr	w8, [sp, #12]                   ; 4-byte Folded Reload
	mov	x0, x8
	ldp	x29, x30, [sp, #32]             ; 16-byte Folded Reload
	add	sp, sp, #48                     ; =48
	ret
	.cfi_endproc
                                        ; -- End function
	.section	__TEXT,__cstring,cstring_literals
l_.str:                                 ; @.str
	.asciz	"\n%d\n"

.subsections_via_symbols
