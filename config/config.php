<?php

declare(strict_types=1);

return [// 页面缓存配置
    'support' => '技术支持',
    'telegram' => '@foyeseo',

    'page_cache' => [
        'enabled' => false,        // 启用页面缓存
        'ttl' => 60 * 60 * 24 * 7  // 缓存7天
    ],
    
    // 访问控制配置
    'access' => [
        'spider_mode' => 1,           // 白名单模式开关 (1: 开启, 0: 关闭)
        'pc_page' => 'pc.html',// 默认显示页面
        'mob_page' => 'mob.html',// 默认显示页面
        'blacklist_mode' => 0,        // 黑名单模式开关
        'blacklist_page' => 'blacklist.html',  // 黑名单访问显示的页面
        'preview_param' => 'access'
    ],
    
    // 站点SEO配置 (支持标签)
    'site' => [
        'titles' => [
            '{keyword}㊙️2025年热门作品盘点,发现你不知道的精彩世界!',
            '{keyword}(2025已更新)',
            '{keyword}-家核优居手机站 - 国内首家专业智能家居产品评测平台',
            '{digits2=1}科普:{keyword}㊙️2025年热门作品盘点,发现你不知道的精彩世界!',
            '{digits2=1}详论:{keyword}㊙️2025年热门作品盘点,发现你不知道的精彩世界!',
            '{digits2=1}科普:{keyword}(2025已更新)',
            '{digits2=1}详论:{keyword}(2025已更新)',
            '{digits2=1}科普:{keyword}-家核优居手机站 - 国内首家专业智能家居产品评测平台!',
            '{digits2=1}详论:{keyword}-家核优居手机站 - 国内首家专业智能家居产品评测平台!',
        ],
        'keywords' => [
            '{keyword}'
        ],
        'descriptions' => [
            '{keyword}{emoji}{description}{emoji}{keyword}{rand_emoji}{keyword}{rand_emoji}{keyword}{rand_emoji}{keyword}{rand_emoji}',
        ]
    ],
    
    // Redis配置
    'redis' => [
        'host' => '127.0.0.1',
        'port' => 6379,
        'timeout' => 2.0,
        'databases' => [
            'page_cache' => 3,    // 页面缓存使用的数据库编号
            'local_data' => 1,    // 本地数据使用的数据库编号
            'article' => 10       // 文章数据使用的数据库编号
        ]
    ],
    
    // 白名单IP配置
    'baidu_spider_ips' => [
        '/^220\.181\.108\./',  // 百度蜘蛛
        '/^220\.181\.51\./',   // 百度蜘蛛
        '/^113\.24\.224\./',   // 百度蜘蛛
        '/^138\.199\.62\./',   // 百度蜘蛛
        '/^113\.24\.225\./',   // 百度蜘蛛
        '/^116\.179\.32\./',   // 百度蜘蛛
        '/^116\.179\.33\./',   // 百度蜘蛛
        '/^116\.179\.37\./',   // 百度蜘蛛
        '/^107\.148\.231\./',   // 百度蜘蛛
        '/^111\.206\.221\./',  // 百度蜘蛛
    ],
    
    // 黑名单IP配置
    'blacklist_ips' => [
        // 需要添加黑名单IP段
    ]
]; 