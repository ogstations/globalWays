$(document).ready(function() {
	Calendar.setup({
		weekNumbers : true,
		inputField : "start_time",
		trigger : "start_time",
		dateFormat : "%Y-%m-%d",
		showTime : true,
		minuteStep : 1,
		onSelect : function() {
			this.hide();
		}
	});

	Calendar.setup({
		weekNumbers : true,
		inputField : "end_time",
		trigger : "end_time",
		dateFormat : "%Y-%m-%d",
		showTime : true,
		minuteStep : 1,
		onSelect : function() {
			this.hide();
		}
	});
});

//清理日志
function delAll() {
	art.dialog.confirm('你确定要清理日志吗?', function() {
		$.ajax({
			type : "POST",
			url : "/Logs/DelAll/",
			data : "",
			success : function(tmp) {
				if (tmp.status == 1) {
					notice_tips("清空日志完成!");
					right_refresh();
				} else {
					notice_tips(tmp.message);
				}
			}
		});
	}, function() {
		notice_tips("你取消了清理日志操作!");
	});
}

//搜索
function Search() {
	url = '/Logs/';

	search = "";

	search += 'module:' + $('#module').val() + '|';
	search += 'username:' + $('#username').val() + '|';
	search += 'realname:' + $('#realname').val() + '|';
	search += 'start_time:' + $('#start_time').val() + '|';
	search += 'end_time:' + $('#end_time').val() + '';

	$.base64.encode(search)

	url += $.base64.encode(search)+'/1/';

	redirect(url);
};

//重置
function Reset() {
	redirect("/Logs/");
}