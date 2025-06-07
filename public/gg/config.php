<?php
// 登录配置
$username = 'foyeseo';
$password_hash = password_hash('foyeseo', PASSWORD_DEFAULT); // 密码加密存储

// 会话配置
session_start();
