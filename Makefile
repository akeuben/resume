pdf:	
	pdflatex main.tex
	latexmk -c *.tex
	mv main.pdf resume.pdf
