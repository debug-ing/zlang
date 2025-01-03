" Vim syntax file
" Language:	    Zlang
" Maintainer:   Mohammad Mahdi Mohammadi


if exists("b:current_syntax")
  finish
endif

" Keywords
syn keyword	pType var 
syn keyword	pStatement func print if else for && ||

syn match   pComment '//.*$'    display
syn match   pInteger '\d\+'     display
syn match   pFloat '\d\+\.\d*'  display
syn region  pString start='"' end='"' display 

hi def link pType        Type
hi def link pComment     Comment
hi def link pStatement   Statement
hi def link pInteger     Number
hi def link pFloat       Number
hi def link pString      String 