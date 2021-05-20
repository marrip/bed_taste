# bed_taste

Simple CLI to combine probes and genetic regions into `.bed` file

![Test and Build](https://github.com/marrip/bed_taste/actions/workflows/main.yaml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## :speech_balloon: Introduction

**bed_taste** is a CLI which takes a list of probes specified by genetic coordinates
and a list of annotated genetic regions (genes, transcripts, exons) and merges them
into a `.bed` file which can be used for clinical variant exploration. The tool
is implemented in `golang` and distributed as a `docker` container image.

## :heavy_exclamation_mark: Dependencies

The tool requires:

[![docker](https://img.shields.io/badge/docker-20.10.0-blue)](https://docs.docker.com/)

Alternatively, it can run with `Singularity`.

## :school_satchel: Preparations

Input `.tsv` files need to be formatted like so:

```bash
<ENSEMBL GENE ID>	<ENSEMBL TRANSCRIPT ID>	<GENE NAME>	<EXON/OTHER IDENTIFIER>	<CHROMOSOME>	<START>	<STOP>
```

`test_data/exons.tsv` may serve as an example. If `ENSEMBL` IDs are not available
these may be omittedi leaving field 1 and 2 empty. 

## :rocket: Usage

`bed_taste` has several options:

Flag | Required | Default | Definition
--- | --- | --- | ---
-probe | yes | - | path to file containing MLPA probes
-genreg | yes | - | path to file containing genetic regions
-out | no | `out.bed` | path to output file to be generated
-hg | no | GRCh38 | version of human genome
-padding | no | 250bp | padding that should be applied to MLPA probe regions

Run it like so:

```bash
docker run -it --rm bed_taste:latest -genreg path/to/genreg.tsv -probe path/to/probes.tsv
```
