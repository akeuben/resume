pdf:	
	pdflatex letter.tex
	pdflatex a4.tex
	pdflatex legal.tex
	latexmk -c *.tex
	mv letter.pdf resume_letter.pdf
	mv a4.pdf resume_a4.pdf
	mv legal.pdf resume_legal.pdf
	rm pdfa.xmpi
