<?php include 'config.php';
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    if ($_POST['username'] === $username && 
        password_verify($_POST['password'], $password_hash)) {
        $_SESSION['loggedin'] = true;
        header('Location: adminxxx.php');
        exit;
    }
    $error = "用户名或密码错误";
}
?>
<!DOCTYPE html>
<html>
<head>
    <title>登录</title>
    <link href="/gg/css/bootstrap.min.css" rel="stylesheet">
</head>
<body class="bg-light">
    <div class="container mt-5" style="max-width: 400px;">
        <div class="card shadow">
            <div class="card-body">
                <h4 class="mb-4">后台登录</h4>
                <?php if(isset($error)): ?>
                <div class="alert alert-danger"><?= $error ?></div>
                <?php endif; ?>
                <form method="post">
                    <div class="mb-3">
                        <input type="text" name="username" class="form-control" placeholder="用户名" required>
                    </div>
                    <div class="mb-3">
                        <input type="password" name="password" class="form-control" placeholder="密码" required>
                    </div>
                    <button class="btn btn-primary w-100">登录</button>
                </form>
            </div>
        </div>
    </div>
</body>
</html>
