all: letter a4 legal cleanup

letter:		
	pdflatex letter.tex
	mv letter.pdf resume_letter.pdf

a4:		
	pdflatex a4.tex
	mv a4.pdf resume_a4.pdf

legal: 
	pdflatex legal.tex
	mv legal.pdf resume_legal.pdf

cleanup:
	latexmk -c *.tex
	rm -f pdfa.xmpi
