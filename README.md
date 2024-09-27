# golang-alignment-block-genotyper

- etxracting the upstream and the downstream of the block alignment. 
- extract specific regions upstream and downstream of the block alignment for genotyper tags. 
- given an alignment either from the whole genome alignment or the gene alignment and given an start and the end estimates, it will extract the upstream and the downstream of that given block to work on the genotyper tags. 
- This is especially when you are dealing with the HybSeq data for the eDNA environmenttal DNA sequencing tags.

```
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang/golang-alignment-block-genotyper% \
go run main.go -a ./samplefile/samplealignment.fasta -s 6 -e 10 -u 4 -d 4
This is the alignment block that has been extracted
>ENA|OX291461|OX291461.1         TC--
This is the alignment block that has been extracted
>ENA|OX291509|OX291509.1         --TC
These are the upstream and the downstream blocks for the chosen block including the block
>ENA|OX291461|OX291461.1         ACTATC----TC
These are the upstream and the downstream blocks for the chosen block including the block
>ENA|OX291509|OX291509.1         TC----TC----
The upstream from the given position till the start is given below:
>ENA|OX291461|OX291461.1         ACTA
The upstream from the given position till the start is given below:
>ENA|OX291509|OX291509.1         TC--
The downstream from the end to the given downstream coordinate is given below:
>ENA|OX291461|OX291461.1         --TC
The downstream from the end to the given downstream coordinate is given below:
>ENA|OX291509|OX291509.1         ----

```

- binary version 

```
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang/golang-alignment-block-genotyper% \
./alignment-block-gentoyper -a ./samplefile/samplealignment.fasta -s 6 -e 10 -u 4 -d 4
This is the alignment block that has been extracted
>ENA|OX291461|OX291461.1         TC--
This is the alignment block that has been extracted
>ENA|OX291509|OX291509.1         --TC
These are the upstream and the downstream blocks for the chosen block including the block
>ENA|OX291461|OX291461.1         ACTATC----TC
These are the upstream and the downstream blocks for the chosen block including the block
>ENA|OX291509|OX291509.1         TC----TC----
The upstream from the given position till the start is given below:
>ENA|OX291461|OX291461.1         ACTA
The upstream from the given position till the start is given below:
>ENA|OX291509|OX291509.1         TC--
The downstream from the end to the given downstream coordinate is given below:
>ENA|OX291461|OX291461.1         --TC
The downstream from the end to the given downstream coordinate is given below:
>ENA|OX291509|OX291509.1         ----
```
Gaurav Sablok
