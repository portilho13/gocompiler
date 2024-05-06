section .text
	global main

main:
	push ebp
	mov ebp, esp
	mov dword [ebp - 4], 4
	push ebp
	mov ebp, esp
	mov dword [ebp - 8], 2
	push ebp
	mov ebp, esp
	mov dword [ebp - 12], 6
	push ebp
	mov ebp, esp
	mov dword [ebp - 16], 8
	mov eax, 0
	leave
	ret
