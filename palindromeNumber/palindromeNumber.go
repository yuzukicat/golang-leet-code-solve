package palindromeNumber

func isPalindrome(x int) bool {
        if x < 0 || (x%10 == 0 && x != 0) {
                return false
        }

        revertedNumber := 0

        for revertedNumber > 0 {
                revertedNumber = revertedNumber * 10 + x%10
                x /= 10
        }
        return x==revertedNumber || x==revertedNumber/10
}
