clean:
	rm -rf dist

doc: clean
	bee run -gendoc=true -downdoc=true
