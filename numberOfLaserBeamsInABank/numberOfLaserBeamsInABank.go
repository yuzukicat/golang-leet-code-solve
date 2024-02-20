package numberOfLaserBeamsInABank

func numberOfBeams(bank []string) int {
        nCellLast := 0
        total := 0
        for _, row := range bank {
                nCell := 0
                for _, cell := range row {
                        if cell == '1' {
                                nCell++
                        }
                }
                if nCell != 0 {
                        total = nCellLast * nCell +total
                        nCellLast = nCell
                }
        }
        return total
}

