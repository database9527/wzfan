<?php
session_start();
if (!isset($_SESSION['loggedin']) || $_SESSION['loggedin'] !== true) {
    header('Location: login.php');
    exit;
}

include 'config.php';

$error = '';
$success = '';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $old_password = $_POST['old_password'];
    $new_password = $_POST['new_password'];
    $confirm_password = $_POST['confirm_password'];

    // 验证旧密码
    if (!password_verify($old_password, $password_hash)) {
        $error = '旧密码错误';
    } 
    // 检查新密码是否匹配
    elseif ($new_password !== $confirm_password) {
        $error = '新密码和确认密码不一致';
    } 
    // 检查新密码复杂度
    elseif (strlen($new_password) < 8) {
        $error = '密码长度至少为8个字符';
    } else {
        // 生成新哈希
        $new_hash = password_hash($new_password, PASSWORD_DEFAULT);

        // 构建新的配置文件内容
        $config_content = "<?php\n";
        $config_content .= "session_start();\n\n";
        $config_content .= "\$username = '" . addslashes($username) . "';\n";
        $config_content .= "\$password_hash = '" . addslashes($new_hash) . "';\n";
        $config_content .= "?>";

        // 写入配置文件
        if (file_put_contents('config.php', $config_content) !== false) {
            $success = '密码修改成功！';
        } else {
            $error = '保存密码失败，请检查文件权限';
        }
    }
}
?>
<!DOCTYPE html>
<html>
<head>
    <title>修改密码</title>
    <link href="/gg/css/bootstrap.min.css" rel="stylesheet">
</head>
<body class="bg-light">
    <div class="container mt-5" style="max-width: 400px;">
        <div class="card shadow">
            <div class="card-body">
                <h4 class="mb-4">修改密码</h4>
                <?php if ($error): ?>
                    <div class="alert alert-danger"><?= $error ?></div>
                <?php endif; ?>
                <?php if ($success): ?>
                    <div class="alert alert-success"><?= $success ?></div>
                <?php endif; ?>
                <form method="post">
                    <div class="mb-3">
                        <input type="password" name="old_password" class="form-control" placeholder="旧密码" required>
                    </div>
                    <div class="mb-3">
                        <input type="password" name="new_password" class="form-control" placeholder="新密码" required>
                    </div>
                    <div class="mb-3">
                        <input type="password" name="confirm_password" class="form-control" placeholder="确认新密码" required>
                    </div>
                    <button type="submit" class="btn btn-primary w-100">修改密码</button>
                </form>
                <div class="mt-3 text-center">
                    <a href="adminxxx.php" class="text-secondary">返回后台</a>
                </div>
            </div>
        </div>
    </div>
</body>
</html>