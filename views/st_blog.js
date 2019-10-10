{{define "blog_js"}}
var stringProto=String.prototype;stringProto.stripTags=function(){return this.replace(/<[^>]*>/gi,"")},stringProto.decode4Html=function(){var t=document.createElement("div");return t.innerHTML=this.stripTags(),t.childNodes[0]?t.childNodes[0].nodeValue||"":""},stringProto.queryUrl=function(t){var a=this.replace(/^[^?=]*\?/gi,"").split("#")[0],e={};return a.replace(/(^|&)([^&=]+)=([^&]*)/g,function(t,a,n,i){try{n=decodeURIComponent(n)}catch(o){}try{i=decodeURIComponent(i)}catch(o){}n in e?e[n] instanceof Array?e[n].push(i):e[n]=[e[n],i]:e[n]=/\[\]$/.test(n)?[i]:i}),t?e[t]:e},function(t){t.disqus_shortname="{{.Disqus.ShortName}}",$.each(["CURRENT_PAGE","CDN_DOMAIN"],function(a,e){t[e]="";var n=$("#"+e);n&&(t[e]=n.val())})}(this),function(t){function a(a){var e,n,i={selector:null,height:200};i=$.extend(i,a),e=i.height,n=function(){var a=$(t).scrollTop(),n=$(t).height()+a;$(i.selector).find("img[data-src]").each(function(){var t=$(this);setTimeout(function(){var i,o=t.offset(),s=t.height();o.top>n+e||o.top+s<a-e||(i=t.data("src"),i&&(t.on("load",function(){t.addClass("loaded")}),t.attr("src",i),t.removeAttr("data-src")))},0)})},$(t).on("resize",n),$(t).on("scroll",n),n()}t.lazyLoad=a}(this),function(t){$(function(){$(".entry-content pre").each(function(t){var a,e,n,i=$(this),o=i.find("code");if(o.length&&o.prop("className")&&(o.hasClass("language-html")&&(a="__HTML_CODE_"+t,o.prop("id",a),e=$('<input data-id="'+a+'" type="button" class="runcode" value="在新窗口运行以上代码" />'),e.insertAfter(i)),o.html().split("\n").length>3&&o.prop("className").indexOf("language")>-1)){switch(n=o.prop("className").replace("language-","").toUpperCase()){case"XML":n="HTML";break;case"SHELL":n="BASH"}$('<b class="name">'+n+"</a>").insertBefore(o)}}),$(".entry-content input.runcode").each(function(){var a=$(this);a.click(function(a){var e,n,i;a.preventDefault(),e=$("#"+$(this).data("id")).html().stripTags().decode4Html(),n=t.open("","_preview",""),i=n.document,i.open(),i.write(e),i.close()})}),$(".entry-content > pre code").each(function(i,block){hljs.highlightBlock(block)})})}(this),function(){$(function(){$(".entry-content img[data-replace]").each(function(){var t=$(this);t.click(function(){var a,e,n,i,o=1000*(t.data("dur")||20);t.css("cursor")&&(a="/static/img/blank.gif",e=t.prop("src"),n=t.data("replace"),t.prop("src",a),t.css("cursor",""),i=new Image,i.onload=function(){t.prop("src",n),setTimeout(function(){t.prop("src",e),t.css("cursor","pointer")},o)},i.src=n)}),t.css("cursor","pointer")})})}(this),function(t){var a=function(){var e=[],n="comment_type",i=function(){if(!t.atob){return 1}try{t.postMessage("ping","*")}catch(a){return 2}return 0};return{addService:function(t){e.push(t)},clear:function(){localStorage.removeItem(n)},init:function(t){var o,r,s,c;return e.length?t.length?(o=t.data("url")||location.href,r=t.data("identifier"),o&&r?i()?void t.html("很抱歉，本站评论功能只支持这些浏览器：Chrome、Firefox、Safari、Opera、Edge、IE11+。"):(s=function(i){var c,l=e[i];l&&(c=$("<div>").prop("id",l.id),t.append(c),l.init(c,o,r),l.check(function(t){t?(c.hide(),s(++i)):(l.load(),localStorage[n]=i,a.currentServer=l)},c,o,r))},c=0,localStorage.comment_type&&(c=parseInt(localStorage[n],10)||0),c=Math.min(e.length-1,c),c=Math.max(c,0),void s(c)):void alert("没有找到评论所需标记！")):void alert("没有找到评论容器！"):void alert("没有找到可用的服务！")}}}();t.TotalThread=a}(this),function(t){var a=function(){var a=!1;return{id:"disqus_thread",check:function(a){var e,n,i=["https://c.disquscdn.com/favicon.ico","https://disqus.com/favicon.ico"],o=0,r=0,s=function(){o==i.length&&a(r==o?0:1)};for(t.__disqus_img=[],e=function(a){var e=new Image,n=setTimeout(function(){e.onerror=e.onload=null,o++,s()},2500);e.onerror=function(){clearTimeout(n),o++,s()},e.onload=function(){clearTimeout(n),o++,r++,s()},e.src=i[a]+"?"+ +new Date,t.__disqus_img[a]=e},n=0;n<i.length;n++){e(n)}},load:function(){if(!a){a=!0;var e=this;t.disqus_config=function(){this.page.url=e.url,this.page.identifier=e.identifier},$.ajax({url:t.CDN_DOMAIN+"/static/js/{{.Disqus.Embed}}",dataType:"script",cache:!0})}},init:function(t,a,e){t.html("评论完整模式加载中...<br /><br />注：如果长时间无法加载，请针对 disq.us | disquscdn.com | disqus.com 启用代理。"),this.elThread=t,this.url=a,this.identifier=e}}}();t.TotalThread.addService(a)}(this),function(t){var a=function(){var e=!1,n=function(e,n){var i,o,r=!1,s="",c=function(a){var e='<li class="post '+(a.isDeleted?"minimized":"")+'" id="post-'+a.id+'"><div class="post-content clearfix"><div class="indicator"></div><div class="avatar"><div class="user">'+(a.url?'<a href="'+a.url+'"><img data-src="'+a.avatar+'" /></a>':'<img data-src="'+t.CDN_DOMAIN+'/static/img/default_avatar.png" />')+'</div></div><div class="post-body"><header class="post-header"><span class="post-byline"><span class="author publisher-anchor-color">'+(a.url?'<a href="'+a.url+'">'+a.name+"</a>":a.name)+'</span></span><span class="post-meta"><span class="bullet bullet-first"> • </span><a data-id="'+a.id+'" href="'+location.href.replace(/#.*$/,"")+"#comment-"+a.id+'" class="time-ago">'+a.createdAtStr+'</a><span class="bullet"> • </span><a href="#" class="reply" data-next="'+s+'" data-id="'+a.id+'">回复</a></span></header><div class="post-body-inner">'+a.message+'</div></div></div><ul class="children" data-id="'+a.id+'"></ul></li>';return e},l=function(t){for(;t.length;){var a=[];t.forEach(function(t){var n=t.parent,i=e.find('ul[data-id="'+n+'"]');return i.length?void i.append(c(t)):void a.push(t)}),t=a}},d=function(a){var n=parseInt(a.parent,10),i=e.find('ul[data-id="'+n+'"]');i.prepend(c(a)),e.find(".no-result").hide(),$(t).trigger("scroll"),$(t).trigger("hashchange","scrollIntoView")};a.insertItem=d,i=function(t){var a='<div class="thread"><header><nav class="nav clearfix"><ul><li class="nav-tab tab-conversation active"><a class="publisher-nav-color"><span>'+t.total+' </span>Comments</span></a></li></ul></nav></header><section><div class="thread-new"><button class="create-post">点击发表新评论</button><span class="tips">（<a href="#reload" class="reload" title="系统检测你可能无法访问 Disqus 服务，已自动显示为评论基础模式。如需完整体验请针对 disq.us | disquscdn.com | disqus.com 启用代理~">尝试评论完整模式</a>）</span></div></section><section class="thread-post-list"><ul class="post-list" data-id="0"></ul><div class="load-more"><a href="#" class="btn">Load more comments</a></div></section></div>';e.attr("data-thread",t.thread),e.attr("data-identifier",n),e.html(a),t.comments.length||e.find(".thread-post-list").append('<p class="no-result">本文目前还没有人评论~</p>')},o=function(){if(!r){r=!0;var a;a=s?"/disqus/"+n+"/"+encodeURIComponent(s):"/disqus/"+n,$.get(a,{},function(a){var n;r=!1,a&&0==a.errno?(s?l(a.data.comments):(i(a.data),l(a.data.comments),$(t).trigger("hashchange","scrollIntoView")),n=e.find(".load-more a"),a.data.next?n.removeClass("busy"):n.hide(),s=a.data.next,$(t).trigger("scroll")):e.html('<p class="no-result">获取数据失败，请稍后再试！</p>')})}},e.on("click",".load-more a",function(t){t.preventDefault(),$(this).addClass("busy"),o()}).on("click","a.time-ago",function(t){t.preventDefault(),location.hash="comment-"+$(this).data("id")}).on("click","a.reload",function(a){a.preventDefault(),t.TotalThread.clear(),location.hash="comments",location.reload()}).on("click","a.reply, button.create-post",function(a){var n,i,o,r,s,c,l,d,u;a.preventDefault(),n=e.data("identifier"),i=e.data("thread"),o=$(this).data("next"),r=$(this).data("id"),s=420,c=520,l=t.screen.width-s/2,d=t.screen.height-c/2,u=[n,i,r,o].join("|"),t.open("/disqus/form/"+encodeURIComponent(u)+"/","_create_post","width="+s+",height="+c+",location=1,status=1,resizable=1,scrollbars=1,left="+l+",top="+d)}),$(t).on("hashchange",function(a,n){var i,o=location.hash.match(/#comment\-(\d+)/);o&&(i=e.find("#post-"+o[1]),i.length&&(e.find(".post-content.target").removeClass("target"),i.find("> .post-content").addClass("target"),n&&$(t).scrollTop(i.offset().top-90)))}),/(iPhone|Android)/.test(navigator.userAgent)&&e.addClass("mobile"),o()};return{id:"simple_thread",check:function(t){t(0)},load:function(){e||(e=!0,n(this.elThread,this.identifier))},init:function(t,a,e){t.html("评论基础模式加载中...<br /><br />注：本模式仅支持最基本的评论功能，如需完整体验请针对 disq.us | disquscdn.com | disqus.com 启用代理。"),this.elThread=t,this.url=a,this.identifier=e}}}();t.TotalThread.addService(a)}(this),function(t,a){var e=a.domain;return"{{.Qiniu.Domain}}"==e?void (location.href=location.href.replace(/(https?:\/\/[^\/]+)\//i,"//{{.Domain}}")):(function(){var t,a=location.search.queryUrl();"1"==a.clear_ls&&(delete a.clear_ls,t=$.param(a),setTimeout(function(){t?location.search=$.param(a):location.href=location.href.replace(/\?.*$/,"")},300))}(),void $(function(){lazyLoad({selector:"#content",height:100}),function(){var t=$("#content"),a=t.find("img"),e=t.width();a.each(function(){var t=$(this),a=0|t.attr("width"),n=0|t.attr("height"),i=t.prop("complete");a>e&&t.attr("height",Math.ceil(n/a*e)),t.prop("src")&&(i?t.addClass("loaded"):t.on("load",function(){t.addClass("loaded")}))})}(),function(){if("search-post"==CURRENT_PAGE){var t=$("#keyword");t.val()||t.focus()}}(),function(){var a=$(".total_thread");a.length&&setTimeout(function(){/ #comment - \d + $ /.test(location.hash)&&$("#comments").get(0).scrollIntoView();var e=setInterval(function(){var n=a.offset().top,i=$(t).scrollTop();Math.abs(n-i)<1000&&(clearTimeout(e),t.TotalThread.init(a))},150)},250)}()}))}(this,document);
{{end}}
