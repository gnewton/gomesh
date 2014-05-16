jianGoMeSHi
===========

Go library to read [MEDLINE/PubMed Medical Subject Headings (MeSH)](http://www.nlm.nih.gov/mesh/) XML

MeSH Record Types
===
* Supports the three types of MeSH [record types](http://www.nlm.nih.gov/mesh/intro_record_types.html), Descriptor, Supplemental, Qualifier.
* Also supports the [pharmacological action terms](http://www.nlm.nih.gov/bsd/disted/meshtutorial/pharmacologicalactionterms/)
* Produces the MeSH hierarchy from the descriptor records

Data
====
MeSH XML download is [here](http://www.nlm.nih.gov/mesh/filelist.html). 
"_Download of any of the full data files requires the completion of an online [Memorandum of Understanding](http://www.nlm.nih.gov/mesh/2014/download/termscon.html)._"
A tiny sample of this data is available in https://github.com/gnewton/jianGoMeSHi/tree/master/testData

Example
===
Example usage: see [jsonGoMeSHi](https://github.com/gnewton/jsonGoMeSHi) or a running [example](http://s2.semanticscience.org:8080/mesh) of jsonGoMeSHi running at [Dumontier Lab](http://dumontierlab.com/)


Naming
=============
As many Golang project are names Go_Foo_ I started with _GoMesh_ but then saw a similarity with _Ghomeshi_ and decided to name the project (partially) after (Jian Ghomeshi)[https://en.wikipedia.org/wiki/Jian_Ghomeshi] and settled on jianGoMeSHi (apologies Jian, as this mis-spells your last name, but matches the MeSH basis of this !). Yes, the connection with Jian is only (partially) phonetic. 

TODO
=============
- Document core methods in README.md
- Allow for loading then writing to SQLite then using SQLite so full data is not in memory (once done allow using SQLite as data source)


Acknowledgement
=============
This work is a by-product of my graduate work at Carleton Univerity at [Dumontier Lab](http://dumontierlab.com/)
