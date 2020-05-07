堆：

    是一种的特殊的树。应用场景非常多，最经典的用法是堆排序。
    堆排序是一种原地的，时间复杂度为O(nlogn)的排序算法。
    
思考：
      堆排序的时间复杂度和快速排序的平均时间复杂度都是O(nlogn)，
      但是实际软件开发中，快速排序的性能比堆排序好，这是为什么？

堆的定义：
      堆是一种特殊的树，特殊在哪里？
      1。堆是一个完全二叉树
      2。堆中的每一个节点的值都必须大于等于（或小于等于）其子树中每个节点的值。（分别为大顶堆和小顶堆）
      
实现堆：
    1。存储：完全二叉树，用数组存储最为合适。 不需要额外的存储指针的空间。通过下标可以定位到孩子，以及父亲节点
    2。操作：
            1。插入元素：将元素插入数组末尾，与父亲结点比较大小，调整。直到满足堆定义的第二条
            2。删除元素：删除堆顶数据，将数组末尾数据，放进堆顶，调整，直到满足堆定义的第二条
            
堆排序：
    1。建堆
        从非叶子结点开始，从上到下堆化
    2。排序
        大顶堆建立好之后，把堆顶元素和数组最后一个元素交换。重新堆化，此次堆化，不计算数组最后一个结点。
        相对于删除堆顶，但是把堆顶放到数组的最后一个位置。
        
解答思考：
    1。堆排序数据访问是按照树的节奏，是下标跳跃的，而快速排序是下标连续的。所以堆排序的cpu缓存没有快速排序友好
    2。堆排序过程中，数据交换次数比较多
    