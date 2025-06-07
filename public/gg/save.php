<?php
include 'config.php';
if (!isset($_SESSION['loggedin'])) {
    http_response_code(403);
    exit('无权访问');
}

$type = $_GET['type'] ?? '';
$content = $_POST['content'] ?? '';

if ($type === 'stats') {
    // 更新 PC 模板统计内容
    $pcFile = '../../config/pc.html';
    $pcContent = file_get_contents($pcFile);
    $newPcContent = preg_replace(
        '/<!--统计开始-->.*?<!--统计结束-->/s',
        "<!--统计开始-->\n".trim($content)."\n<!--统计结束-->",
        $pcContent
    );
    file_put_contents($pcFile, $newPcContent);

    // 更新移动端模板统计内容
    $mobFile = '../../config/mob.html';
    $mobContent = file_get_contents($mobFile);
    $newMobContent = preg_replace(
        '/<!--统计开始-->.*?<!--统计结束-->/s',
        "<!--统计开始-->\n".trim($content)."\n<!--统计结束-->",
        $mobContent
    );
    file_put_contents($mobFile, $newMobContent);
} elseif ($type === 'tz1') {
    // 更新 PC 跳转文件
    file_put_contents('../../public/tz1.js', $content);
} elseif ($type === 'tz2') {
    // 更新移动端跳转文件
    file_put_contents('../../public/tz2.js', $content);
}

header('Location: adminxxx.php');