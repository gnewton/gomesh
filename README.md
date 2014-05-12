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
"Download of any of the full data files requires the completion of an online [Memorandum of Understanding](http://www.nlm.nih.gov/mesh/2014/download/termscon.html)."
A tiny sample of this data is available in https://github.com/gnewton/jianGoMeSHi/tree/master/testData

Example
===
Example usage: see [jsonGoMeSHi](https://github.com/gnewton/jsonGoMeSHi) or a running [example](http://s2.semanticscience.org:8080/mesh) of jsonGoMeSHi running at [Dumontier Lab](http://dumontierlab.com/)


Acknowledgement
=============
This work is a by-product of my graduate work at Carleton Univerity at [Dumontier Lab](http://dumontierlab.com/)
