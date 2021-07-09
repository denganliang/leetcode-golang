/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func firstBadVersion(n int) int {
    m := n/2
    e := n
    for m>=1 && isBadVersion(m) {
        e = m
        m = m/2
    }
    for m < e {
        if isBadVersion(m) {
            return m
        }
        m++
    }
    return m
}
