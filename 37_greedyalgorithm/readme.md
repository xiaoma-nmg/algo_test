贪心算法：
   
一、经典应用
    霍夫曼编码、 Prim和Kruskal最小生成树、 DijKstra单源最短路径

二、解题思路
    1。首先要联想到贪心算法
        针对一组数据，定义了限制值和期望值，希望从中选出几个数据，在满足限制值的情况下，期望值最大。
    2。尝试问题是否可以用贪心算法解决
        选择当前情况下，相同限制值中对期望值贡献最大的数据
    3。验证贪心的结果是否是最优的
    