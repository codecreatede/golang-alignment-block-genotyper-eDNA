package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-27

given an alignment either from the whole genome alignment or the gene alignment and
given an start and the end estimates, it will extract the upstream and the downstream
of that given block to work on the genotyper tags. This is especially when you are dealing
with the HybSeq data for the eDNA tags.


*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	alignment  string
	start      int
	end        int
	upstream   int
	downstream int
)

var rootCmd = &cobra.Command{
	Use:  "flags",
	Long: "This estimates the site proportion in your whole genome or gene specific alignment",
	Run:  flagsFunc,
}

func init() {
	rootCmd.Flags().StringVarP(&alignment, "alignmentfile", "a", "align", "a alignment file")
	rootCmd.Flags().IntVarP(&start, "startcoordinate", "s", 1, "start of the alignment block")
	rootCmd.Flags().IntVarP(&end, "endcoordinate", "e", 40, "end of the alignment block")
	rootCmd.Flags().IntVarP(&upstream, "upstream-alignment", "u", 4, "upstream of the alignment")
	rootCmd.Flags().
		IntVarP(&downstream, "downstream-alignment", "d", 4, "downstream of the alignment")
}

func flagsFunc(cmd *cobra.Command, args []string) {
	type alignmentID struct {
		id string
	}

	type alignmentSeq struct {
		seq string
	}

	type alignBlock struct {
		id  string
		seq string
	}

	type updownStream struct {
		id  string
		seq string
	}

	type upstreamStart struct {
		id  string
		seq string
	}

	type downstreamStart struct {
		id  string
		seq string
	}

	fOpen, err := os.Open(alignment)
	if err != nil {
		log.Fatal(err)
	}

	alignIDcapture := []alignmentID{}
	alignSeqcapture := []alignmentSeq{}
	sequenceCap := []string{}
	sequenceID := []string{}
	alignmentBlock := []alignBlock{}
	upstreamBlock := []updownStream{}
	upstreamfinal := start - upstream
	downstreamfinal := end + downstream
	upstreamS := []upstreamStart{}
	downstreamS := []downstreamStart{}

	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignIDcapture = append(alignIDcapture, alignmentID{
				id: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignSeqcapture = append(alignSeqcapture, alignmentSeq{
				seq: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			sequenceCap = append(sequenceCap, string(line))
		}
		if strings.HasPrefix(string(line), ">") {
			sequenceID = append(sequenceID, string(line))
		}
	}

	for i := 0; i < len(sequenceID); i++ {
		alignmentBlock = append(alignmentBlock, alignBlock{
			id:  string(sequenceID[i]),
			seq: string(sequenceCap[i][start:end]),
		})
	}

	for i := range alignmentBlock {
		fmt.Println("This is the alignment block that has been extracted")
		fmt.Println(alignmentBlock[i].id, "\t", alignmentBlock[i].seq)
	}

	for i := 0; i < len(sequenceID); i++ {
		upstreamBlock = append(upstreamBlock, updownStream{
			id:  string(sequenceID[i]),
			seq: string(sequenceCap[i][upstreamfinal:downstreamfinal]),
		})
	}
	fmt.Println(
		"These are the upstream and the downstream blocks for the chosen block including the block",
	)
	for i := range upstreamBlock {
		fmt.Println(upstreamBlock[i].id, "\t", upstreamBlock[i].seq)
	}

	for i := 0; i < len(sequenceID); i++ {
		upstreamS = append(upstreamS, upstreamStart{
			id:  string(sequenceID[i]),
			seq: string(sequenceCap[i][upstreamfinal:start]),
		})
	}
	for i := 0; i < len(sequenceID); i++ {
		downstreamS = append(downstreamS, downstreamStart{
			id:  string(sequenceID[i]),
			seq: string(sequenceCap[i][end:downstreamfinal]),
		})
	}

	fmt.Println("The upstream from the given position till the start is given below:")
	for i := range upstreamS {
		fmt.Println(upstreamS[i].id, "\t", upstreamS[i].seq)
	}
	fmt.Println(
		"The downstream from the end to the given downstream coordinate is given below:",
	)
	for i := range downstreamS {
		fmt.Println(downstreamS[i].id, "\t", downstreamS[i].seq)
	}
}
