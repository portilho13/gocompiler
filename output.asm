section .text
	global main

main:
	push ebp
	mov ebp, esp
	mov dword [ebp - 4], 4
	mov eax, 0
	leave
	ret
