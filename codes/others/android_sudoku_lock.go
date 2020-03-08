package others

/*
给定九宫格(参考安桌解锁九宫格), 以及m,n, m表示连接点的最小个数,
	n表示连接连接点的最大个数. 有效路径符合如下条件
1. 不能越过某个点而不连接, 即路径上的点都属于该路径
2. 不同顺序算作两条路径

给定m,n, 求符合条件的有效路径数.

示例: 如下九宫格, m=1,n=1时 有9个路径. m=2,n=2时有56种.
有效路径举例: [1,2,3,5]. 无效路径举例: [1,3,5](经过2但是2不再路径上)
1 2 3
4 5 6
7 8 9
*/

// tag: mobile

//points = [[1,2,3],[4,5,6],[7,8,9]]
//# 用 dx/dy 这种方式表示步长, 是我之前没有的思维
//dx = [1,-1,0,0]
//dy = [0,0,1,-1]
//
//def dfs(mink, maxk, x, y, visitd, ans):
//    if  points[x][y] in visitd:
//        return
//    visitd.append(points[x][y])
//    if len(visitd)>maxk:
//        visitd.remove(points[x][y])
//        return
//    if len(visitd) >= mink:
//        ans+=1
//    for i in range(4):
//        nx = x+dx[i]
//        ny = y+dy[i]
//        dfs(mink,maxk,nx,ny,visitd,ans)
//    visitd.remove(points[x][y])
//
//def do(m,n):
//    ans = 0
//    # 从每个点开始遍历
//    for x in range(3):
//        for y in range(3):
//            dfs(m,n,x,y,[],ans)
//    return ans
