var zbp = new ZBP({
	bloghost: "http://demo.30blog.cc/demo/cms/",
	ajaxurl: "http://demo.30blog.cc/demo/cms/zb_system/cmd.php?act=ajax&src=",
	cookiepath: "/demo/cms/",
	lang: {
		error: {
			72: "鍚嶇О涓嶈兘涓虹┖鎴栨牸寮忎笉姝ｇ‘",
			29: "閭鏍煎紡涓嶆纭紝鍙兘杩囬暱鎴栦负绌�",
			46: "璇勮鍐呭涓嶈兘涓虹┖鎴栬繃闀�"
		}
	}
});

var bloghost = zbp.options.bloghost;
var cookiespath = zbp.options.cookiepath;
var ajaxurl = zbp.options.ajaxurl;
var lang_comment_name_error = zbp.options.lang.error[72];
var lang_comment_email_error = zbp.options.lang.error[29];
var lang_comment_content_error = zbp.options.lang.error[46];

$(function () {

	zbp.cookie.set("timezone", (new Date().getTimezoneOffset()/60)*(-1));
	var $cpLogin = $(".cp-login").find("a");
	var $cpVrs = $(".cp-vrs").find("a");
	var $addinfo = zbp.cookie.get("addinfodemocms");
	if (!$addinfo){
		zbp.userinfo.output();
		return ;
	}
	$addinfo = JSON.parse($addinfo);

	if ($addinfo.chkadmin){
		$(".cp-hello").html("娆㈣繋 " + $addinfo.useralias + " (" + $addinfo.levelname  + ")");
		if ($cpLogin.length == 1 && $cpLogin.html().indexOf("[") > -1) {
			$cpLogin.html("[鍚庡彴绠＄悊]");
		} else {
			$cpLogin.html("鍚庡彴绠＄悊");
		}
	}

	if($addinfo.chkarticle){
		if ($cpLogin.length == 1 && $cpVrs.html().indexOf("[") > -1) {
			$cpVrs.html("[鏂板缓鏂囩珷]");
		} else {
			$cpVrs.html("鏂板缓鏂囩珷");
		}
		$cpVrs.attr("href", zbp.options.bloghost + "zb_system/cmd.php?act=ArticleEdt");
	}

});

document.writeln("<script src='http://demo.30blog.cc/demo/cms/zb_users/plugin/UEditor/third-party/prism/prism.js' type='text/javascript'></script><link rel='stylesheet' type='text/css' href='http://demo.30blog.cc/demo/cms/zb_users/plugin/UEditor/third-party/prism/prism.css'/>");$(function(){var compatibility={as3:"actionscript","c#":"csharp",delphi:"pascal",html:"markup",xml:"markup",vb:"basic",js:"javascript",plain:"markdown",pl:"perl",ps:"powershell"};var runFunction=function(doms,callback){doms.each(function(index,unwrappedDom){var dom=$(unwrappedDom);var codeDom=$("<code>");if(callback)callback(dom);var languageClass="prism-language-"+function(classObject){if(classObject===null)return"markdown";var className=classObject[1];return compatibility[className]?compatibility[className]:className}(dom.attr("class").match(/prism-language-([0-9a-zA-Z]+)/));codeDom.html(dom.html()).addClass("prism-line-numbers").addClass(languageClass);dom.html("").addClass(languageClass).append(codeDom)})};runFunction($("pre.prism-highlight"));runFunction($('pre[class*="brush:"]'),function(preDom){var original;if((original=preDom.attr("class").match(/brush:([a-zA-Z0-9\#]+);/))!==null){preDom.get(0).className="prism-highlight prism-language-"+original[1]}});Prism.highlightAll()});
