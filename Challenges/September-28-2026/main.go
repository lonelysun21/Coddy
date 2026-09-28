func OrganizeSale(neighbor1Items []string, neighbor2Items []string) []string {
    result := make([]string, 0, 10)
    for _, i := range neighbor1Items {
        result = append(result, i)
    }
    for _, i := range neighbor2Items {
        result = append(result, i)
    }
    return result
}
