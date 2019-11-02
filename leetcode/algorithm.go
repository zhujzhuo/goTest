package main


//BubbleSort(冒泡排序)
/*
比较相邻的元素。如果第一个比第二个大，就交换它们两个；
对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
针对所有的元素重复以上的步骤，除了最后一个；
*/
func Sort(list []int, left, right int)  {
    if right == 0 {
        return
    }
    for index,num := range list {
        if index < right && num > list[index + 1] {
            utils.SwapGo(list, index, index + 1)
        }
    }
    Sort(list, left, right - 1)
}


//BucketSort(桶排序)

func Sort(list []int)  []int{
    max := max(list)
    min := min(list)
    base := 0
    if min < 0 {
        base = -min
    } else {
        base = min
    }
    max = (max + base)/10
    min = (min + base)/10
    bucket := make([][]int, max - min + 1)
    var result []int
    for _,value := range list {
        i := (int)((value+base)/10)
        bucket[i] = append(bucket[i], value)
    }

    for _,value := range bucket {
        if len(value) > 0 {
            quicksort.Sort(value, 0, len(value)-1)
        }
    }

    for _,value := range bucket {
        if len(value) > 0 {
            for _,v := range value {
                result = append(result,v)
            }
        }
    }
    return result
}

func max(list []int) int  {
    max := list[0]
    for _,value := range list {
        if value > max {
            max = value
        }
    }
    return max
}

func min(list []int) int  {
    min := list[0]
    for _,value := range list {
        if value < min {
            min = value
        }
    }
    return min
}
//CountSort (计数排序)

func Sort(list []int) []int{
    max := max(list)
    min := min(list)
    base := -min
    max = max - base
    min = min - base
    numbers := make([]int, max - min + 1)
    for _,value := range list{
        numbers[value + base] = numbers[value + base] + 1
    }
    var result []int
    for i,value := range numbers {
        for j:=value;j>0 && value > 0;j-- {
            result = append(result, i - base)
        }
    }
    return result

}

func max(list []int) int  {
    max := list[0]
    for _,value := range list {
        if value > max {
            max = value
        }
    }
    return max
}

func min(list []int) int  {
    min := list[0]
    for _,value := range list {
        if value < min {
            min = value
        }
    }
    return min
}

//HeapSort(堆排序)

func Sort(list []int)  {
    length := len(list)
    for {
        if length < 1 {
            break
        }
        index := length/2 -1
        for ;index>=0;index-- {
            swap(list, index, length-1)
        }
        tmp := list[0]
        list[0] = list[length - 1]
        list[length - 1] = tmp
        length--
    }
}

func swap(list []int, index int, length int)  {

    left := 2*index + 1
    right := 2*index + 2
    if left <= length && list[left] > list[index] {
        tmp := list[index]
        list[index] = list[left]
        list[left] = tmp
    }
    if right <= length && list[right] > list[index] {
        tmp := list[index]
        list[index] = list[right]
        list[right] = tmp
    }
}

//InsertSort(插入排序)
/*
插入排序（Insertion-Sort）的算法描述是一种简单直观的排序算法。它的工作原理是通过构建有序序列，
对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。插入排序在实现上，通常采用in-place排序
（即只需用到O(1)的额外空间的排序），因而在从后向前扫描过程中，需要反复把已排序元素逐步向后挪位，为最新元素提供插入空间。
3.1 算法描述
一般来说，插入排序都采用in-place在数组上实现。具体算法描述如下：
从第一个元素开始，该元素可以认为已经被排序；
取出下一个元素，在已经排序的元素序列中从后向前扫描；
如果该元素（已排序）大于新元素，将该元素移到下一位置；
重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
将新元素插入到该位置后；
重复步骤2~5。
*/
func Sort(list []int, left, right int)  {
    for index := left;index <= right;index++ {
        if index > 0 {
            for i:=index;i>0;i-- {
                current := list[i]
                pre := list[i-1]
                if current <= pre {
                    utils.SwapGo(list, i, i-1)
                } else {
                    break
                }
            }
        }
    }
}

//MergeSort(合并排序)
/*
和选择排序一样，归并排序的性能不受输入数据的影响，但表现比选择排序好的多，因为始终都是O(n log n）
	的时间复杂度。代价是需要额外的内存空间。
归并排序是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一
个非常典型的应用。归并排序是一种稳定的排序方法。将已有序的子序列合并，得到完全有序的序列；即先使每
个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。
5.1 算法描述
把长度为n的输入序列分成两个长度为n/2的子序列；
对这两个子序列分别采用归并排序；
将两个排序好的子序列合并成一个最终的排序序列。
*/
func Sort(list []int, left, right int) []int{
    return mergeSort(list[left:right-left + 1])
}

func mergeSort(list []int) []int {
    if len(list) < 2 {
        return list
    } else {
        return merge(mergeSort(list[:len(list)/2]), mergeSort(list[len(list)/2:]))

    }
}

func merge(list0, list1 []int) []int{
    var result []int
    index0 := 0
    index1 := 0
    for {
        if index0 < len(list0) && index1 < len(list1) {
            if list0[index0] < list1[index1] {
                result = append(result, list0[index0])
                index0 = index0 + 1
            } else {
                result = append(result, list1[index1])
                index1 = index1 + 1
            }
        } else {
            break
        }
    }
    if index0 < len(list0) {
        result = append(result, list0[index0:]...)
    }
    if index1 < len(list1) {
        result = append(result, list1[index1:]...)
    }
    return result
}

//QuickSort(快速排序)

import "github.com/go-algorithm/utils"

func Sort(list []int, left, right int)  {
    if right < left {
        return
    }
    flag := list[left]
    start := left
    end := right
    for {
        if start == end {
            break
        }
        for list[end] >= flag && end > start {
            end--
        }
        for list[start] <= flag && end > start {
            start++
        }
        if end > start {
            utils.SwapGo(list, start, end)
        }
    }
    utils.SwapGo(list, left, start)
    Sort(list, left, start - 1)
    Sort(list, start + 1, right)
}

//RadixSort(基数排序)

func Sort(list []int)  {
    baseList := make([][]int, 10)
    maxDigist := maxDigist(list)
    for i:=0;i<maxDigist;i++ {
        for _,value := range list {
            baseList[getDigist(value, i)] = append(baseList[getDigist(value, i)], value)
        }

        j := 0
        for index,value :=range baseList {
            if len(value) > 0 {
                for _,v := range value {
                    list[j] = v
                    j++
                }
            }
            baseList[index] = nil
        }
    }
}

func maxDigist(list []int) int {
    maxDigist := 1
    for _,value := range list {
        if len(strconv.Itoa(value)) > maxDigist {
            maxDigist = len(strconv.Itoa(value))
        }
    }
    return maxDigist
}

func getDigist(number int, index int) int  {
    strNum := strconv.Itoa(number)
    if index > len(strNum) - 1 {
        return 0
    }
    index = len(strNum) - 1 - index
    //fmt.Println("index = ", index)
    result,error := strconv.Atoi(string(strNum[index]))
    if error != nil {

        return -1
    } else {
        return result
    }
}

//SelectSort(选择排序)
/*
选择排序(Selection-sort)是一种简单直观的排序算法。它的工作原理：
首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，然后，
再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
以此类推，直到所有元素均排序完毕。
*/

func Sort(list []int, left, right int)  {
    if left == right {
        return
    }
    minIndex := left
    for i:=left;i<=right;i++ {
        if list[i] <= list[minIndex] {
            minIndex = i
        }
    }
    utils.SwapGo(list, left, minIndex)
    Sort(list, left + 1, right)
}

//shellsort(希尔排序)
/*
我们来看下希尔排序的基本步骤，在此我们选择增量gap=length/2，缩小增量继续以gap = gap/2的方式，
这种增量选择我们可以用一个序列来表示，{n/2,(n/2)/2...1}，称为增量序列。希尔排序的增量序列的选
择与证明是个数学难题，我们选择的这个增量序列是比较常用的，也是希尔建议的增量，称为希尔增量，但其
实这个增量序列不是最优的。此处我们做示例使用希尔增量。
先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，具体算法描述：
选择一个增量序列t1，t2，…，tk，其中ti>tj，tk=1；
按增量序列个数k，对序列进行k 趟排序；
每趟排序，根据对应的增量ti，将待排序列分割成若干长度为m 的子序列，分别对各子表进行直接插入排序。
仅增量因子为1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
*/
func Sort(list []int, left, right int)  {
    increment := len(list)/3 + 1
    for {
        if increment < 1 {
            break
        }
        for i:=left;i<increment;i++ {
            for j:=i+increment;j<=right;j++ {
                if list[j] < list[j-increment] {
                    tmp := list[j]
                    list[j] = list[j-increment]
                    list[j-increment] = tmp
                }
            }
        }
        increment--
    }
}


var data = []int{8, 3, 6, 9, 11, 2, 7, 23, 65, 13, 9}

func main() {
    datePrintln("桶排序")
    datePrintln("计数排序")
    datePrintln("冒泡排序")
    datePrintln("快速排序")
    datePrintln("选择排序")
    datePrintln("插入排序")
    datePrintln("希尔排序")
    datePrintln("合并排序")
    datePrintln("基数排序")
    datePrintln("堆排序")

}
func datePrintln(name string) {
    var result []int
    fmt.Println("初始数据:", data)
    switch name {
    case "桶排序":
        result = bucketsort.Sort(data)
        break
    case "计数排序":
        result = countsort.Sort(data)
        break
    case "合并排序":
        result = mergesort.Sort(data, 0, len(data)-1)
        break
    case "冒泡排序":
        bubblesort.Sort(data, 0, len(data)-1)
        result = data
        break
    case "快速排序":
        quicksort.Sort(data, 0, len(data)-1)
        result = data
        break
    case "选择排序":
        selectsort.Sort(data, 0, len(data)-1)
        result = data
        break
    case "插入排序":
        insertsort.Sort(data, 0, len(data)-1)
        result = data
        break
    case "希尔排序":
        shellsort.Sort(data, 0, len(data)-1)
        result = data
        break
    case "基数排序":
        radixsort.Sort(data)
        result = data
        break
    case "堆排序":
        heapsort.Sort(data)
        result = data
        break
    }
    fmt.Println(name+":", result)
    data = []int{8, 3, 6, 9, 11, 2, 7, 23, 65, 13, 9}
    fmt.Println("原始数据:", data, "\n")
}