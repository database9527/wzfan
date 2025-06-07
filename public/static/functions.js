$(document).ready(function(){
    $("#btn-bars").click(function(){
	$(".navbar").slideToggle(500);
    });
});
$(function(){
    $('#monavber li').hover(function(){
       $(this).addClass('on');
    },function(){
       $(this).removeClass('on');
    });
});

jQuery(document).ready(function($){
var datatype=$("#monavber").attr("data-type");
    $(".navbar>li ").each(function(){
        try{
            var myid=$(this).attr("id");
            if("index"==datatype){
                if(myid=="nvabar-item-index"){
                    $("#nvabar-item-index").addClass("active");
                }
            }else if("category"==datatype){
                var infoid=$("#monavber").attr("data-infoid");
                if(infoid!=null){
                    var b=infoid.split(' ');
                    for(var i=0;i<b.length;i++){
                        if(myid=="navbar-category-"+b[i]){
                            $("#navbar-category-"+b[i]+"").addClass("active");
                        }
                    }
                }
            }else if("article"==datatype){
                var infoid=$("#monavber").attr("data-infoid");
                if(infoid!=null){
                    var b=infoid.split(' ');
                    for(var i=0;i<b.length;i++){
                        if(myid=="navbar-category-"+b[i]){
                            $("#navbar-category-"+b[i]+"").addClass("active");
                        }
                    }
                }
            }else if("page"==datatype){
                var infoid=$("#monavber").attr("data-infoid");
                if(infoid!=null){
                    if(myid=="navbar-page-"+infoid){
                        $("#navbar-page-"+infoid+"").addClass("active");
                    }
                }
            }else if("tag"==datatype){
                var infoid=$("#monavber").attr("data-infoid");
                if(infoid!=null){
                    if(myid=="navbar-tag-"+infoid){
                        $("#navbar-tag-"+infoid+"").addClass("active");
                    }
                }
            }
        }catch(E){}
    });
	$("#monavber").delegate("a","click",function(){
		$(".navbar>li").each(function(){
			$(this).removeClass("active");
		});
		if($(this).closest("ul")!=null && $(this).closest("ul").length!=0){
			if($(this).closest("ul").attr("id")=="munavber"){
				$(this).addClass("active");
			}else{
				$(this).closest("ul").closest("li").addClass("active");
			}
		}
	});
});

$(function() {
	$(window).scroll(function(){
		if($(window).scrollTop()>500){
			$("#gttop").show();
		}else{
			$("#gttop").hide();
		}
	});
	$("#gttop").click(function(){
		$("body,html").animate({scrollTop:0},1500);
		return false;
	});
});

jQuery(document).ready(function(){
    jQuery(".hm-search-button-icon").click(function() {
        jQuery(".hm-search-box-container").toggle('fast');
        jQuery(this).toggleClass("hm-search-close");
    });
});

$(function () {
    if ($("#module").length > 0) {
        var offset = $("#module").offset();
        $(window).scroll(function () {
            var scrollTop = $(window).scrollTop();
            if (offset.top < scrollTop) $("#module").addClass("following2");
            else $("#module").removeClass("following2");
        });
    }
});	

  $(function(){

        var $bottomTools = $('.bottom_tools');
        var $qrTools = $('.qr_tool');
        var qrImg = $('.qr_img');

        $(window).scroll(function () {
            var scrollHeight = $(document).height();
            var scrollTop = $(window).scrollTop();
            var $windowHeight = $(window).innerHeight();
            scrollTop > 50 ? $("#scrollUp").fadeIn(200).css("display","block") : $("#scrollUp").fadeOut(200);
            $bottomTools.css("bottom", scrollHeight - scrollTop > $windowHeight ? 40 : $windowHeight + scrollTop + 40 - scrollHeight);
        });

        $('#scrollUp').click(function (e) {
            e.preventDefault();
            $('html,body').animate({ scrollTop:0});
        });

        $qrTools.hover(function () {
            qrImg.fadeIn();
        }, function(){
             qrImg.fadeOut();
        });

    });