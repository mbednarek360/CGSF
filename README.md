# CGSF

### Text compression algorithm based on converting groups of charectars into graphical formats.

#### ⚠️CGSF is still in alpha, it may have bugs. Updates will come. ⚠️

##### CGSF is a command line based program, no gui is currently available. It tends to work much better compressing large files. It is also faster than most other compression algorithms.

---
To install, download and run install.sh.

`sudo sh install.sh`


To uninstall, run install.sh again located in the CGSF directory.

`sudo sh ~/.local/share/CGSF/install.sh`

---

Encoding takes a file input and a file output arguement.

Ex:
`cgsf-encode input.txt output.cgsf`


Decoding a file requires a similar process.

Ex:
`cgsf-decode input.cgsf output.txt`
