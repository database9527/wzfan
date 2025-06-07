<?php
echo "****************************************";
echo "<br>";
echo "*";
echo "<br>";
echo "*";
echo "<br>";
echo "*";
echo "<br>";
echo "*";
echo "<br>";
echo "*       一键清空所有蜘蛛统计！";
echo "<br>";
echo "*";
echo "<br>";
echo "*";
echo "<br>";
echo "*";
echo "<br>";
echo "*";
echo "****************************************";
echo "<br>";
//----------删除../tongji/下全部文件
//设置需要删除的文件夹
$path = "../kkmp326/tongji/";
//清空文件夹函数和清空文件夹后删除空文件夹函数的处理
function deldir($path) {
    //如果是目录则继续
    if (is_dir($path)) {
        //扫描一个文件夹内的所有文件夹和文件并返回数组
        $p = scandir($path);
        foreach ($p as $val) {
            //排除目录中的.和..
            if ($val != "." && $val != "..") {
                //如果是目录则递归子目录，继续操作
                if (is_dir($path . $val)) {
                    //子目录中操作删除文件夹和文件
                    deldir($path . $val . '/');
                    //目录清空后删除空文件夹
                    @rmdir($path . $val . '/');
                } else {
                    //如果是文件直接删除
                    unlink($path . $val);
                }
            }
        }
    }
}
//调用函数，传入路径
deldir($path);
//----------新建文件夹
echo "删除/kkmp326/tongji/下全部文件成功";
echo "<br>";
mkdir("../kkmp326/tongji/hour/", 0777);
echo "新建文件夹：../kkmp326/tongji/hour/";
echo "<br>";
mkdir("../kkmp326/tongji/360Spider/", 0777);
echo "新建文件夹：../kkmp326/tongji/360Spider/";
echo "<br>";
mkdir("../kkmp326/tongji/Baiduspider/", 0777);
echo "新建文件夹：../kkmp326/tongji/Baiduspider/";
echo "<br>";
mkdir("../kkmp326/tongji/Bytespider/", 0777);
echo "新建文件夹：../kkmp326/tongji/Bytespider/";
echo "<br>";
mkdir("../kkmp326/tongji/Googlebot/", 0777);
echo "新建文件夹：../kkmp326/tongji/Googlebot/";
echo "<br>";
mkdir("../kkmp326/tongji/Sogou/", 0777);
echo "新建文件夹：../kkmp326/tongji/Sogou/";
echo "<br>";
mkdir("../kkmp326/tongji/Yisouspider/", 0777);
echo "新建文件夹：../kkmp326/tongji/Yisouspider/";
echo "<br>";