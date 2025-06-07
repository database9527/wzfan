<?php 
include 'config.php';
if (!isset($_SESSION['loggedin'])) {
    header('Location: login.php');
    exit;
}

// 读取内容
$pcstatsContent = getContentBetweenMarkers('../../config/pc.html', '统计开始', '统计结束');
$tz1Content = file_get_contents('../../public/tz1.js');
$tz2Content = file_get_contents('../../public/tz2.js');

function getContentBetweenMarkers($file, $startMarker, $endMarker) {
    $content = file_get_contents($file);
    preg_match("/<!--$startMarker-->(.*?)<!--$endMarker-->/s", $content, $matches);
    return $matches[1] ?? '';
}
?>
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>后台管理</title>
    <link href="/gg/css/bootstrap.min.css" rel="stylesheet">
    <style>
        body { background: #f8f9fa; min-height: 100vh; }
        .editor-box { border: 1px solid #dee2e6; border-radius: 4px; }
        textarea { 
            width: 100%; 
            padding: 15px;
            font-family: Monaco, Consolas, monospace;
            border: none;
            background: #f8f9fa;
        }
        nav { background: #fff; box-shadow: 0 2px 4px rgba(0,0,0,.1); }
    </style>
</head>
<body>
    <nav class="navbar mb-4">
        <div class="container">
            <span class="navbar-brand">管理后台</span>
            <p><a href="editadmin.php" class="btn btn-secondary">修改密码</a>
            <a href="logout.php" class="btn btn-sm btn-outline-secondary">退出</a>
        </div>
    </nav>

    <div class="container">
        <ul class="nav nav-tabs mb-3">
            <li class="nav-item">
                <a class="nav-link active" href="#stats" data-bs-toggle="tab">统计代码</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="#redirect1" data-bs-toggle="tab">PC端</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="#redirect2" data-bs-toggle="tab">移动端</a>
            </li>
        </ul>

        <div class="tab-content">
            <div class="tab-pane show active" id="stats">
                <form action="save.php?type=stats" method="post">
                    <div class="editor-box">
                        <textarea name="content" rows="15"><?= htmlspecialchars($pcstatsContent) ?></textarea>
                    </div>
                    <button class="btn btn-primary mt-3">保存</button>
                </form>
            </div>

            <div class="tab-pane" id="redirect1">
                <form action="save.php?type=tz1" method="post">
                    <div class="editor-box">
                        <textarea name="content" rows="15"><?= htmlspecialchars($tz1Content) ?></textarea>
                    </div>
                    <button class="btn btn-primary mt-3">保存</button>
                </form>
            </div>

            <div class="tab-pane" id="redirect2">
                <form action="save.php?type=tz2" method="post">
                    <div class="editor-box">
                        <textarea name="content" rows="15"><?= htmlspecialchars($tz2Content)?></textarea>
                    </div>
                    <button class="btn btn-primary mt-3">保存</button>
                </form>
            </div>
        </div>
    </div>

    <script src="/gg/js/bootstrap.bundle.min.js"></script>
</body>
</html>
